package app

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// image-processing methods: verify + generate-fixtures.
// Phase-N gate scripts are intentionally NOT shipped — compose these
// methods when a later plan phase needs a gate (see workflow tools README).

func runImageProcessing(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		_, _ = fmt.Fprintln(stderr, "usage: ai-skill image-processing <verify|generate-fixtures> [flags]")
		return ExitInvalidUsage
	}
	switch args[0] {
	case "verify":
		return runImageProcessingVerify(args[1:], stdout, stderr)
	case "generate-fixtures":
		return runImageProcessingGenerateFixtures(args[1:], stdout, stderr)
	case "help", "-h", "--help":
		_, _ = fmt.Fprintln(stdout, "usage: ai-skill image-processing <verify|generate-fixtures> [flags]")
		return ExitSuccess
	default:
		_, _ = fmt.Fprintf(stderr, "unsupported image-processing command: %s\n", args[0])
		return ExitInvalidUsage
	}
}

type ipCheckResult struct {
	CheckID         string `json:"check_id"`
	Status          string `json:"status"`
	Evidence        string `json:"evidence"`
	FailureClass    string `json:"failure_class,omitempty"`
	RepairCandidate string `json:"repair_candidate,omitempty"`
	ReviewRequired  bool   `json:"review_required"`
	Hardness        string `json:"hardness"`
}

func runImageProcessingVerify(args []string, stdout io.Writer, stderr io.Writer) int {
	fs := newFlagSet("image-processing verify", stderr)
	profile := fs.String("profile", "", "fixture/profile name (default: image stem)")
	source := fs.String("source", "", "reference/source image for fidelity checks")
	consumerProfile := fs.String("consumer-profile", "", "path to consumer profile JSON")
	jsonOut := fs.String("json-out", "", "write report JSON to path")
	expectBlock := fs.Bool("expect-block", false, "exit 0 only when finality is not accepted (negative fixture)")
	if err := fs.Parse(args); err != nil {
		return ExitInvalidUsage
	}
	rest := fs.Args()
	if len(rest) != 1 {
		_, _ = fmt.Fprintln(stderr, "usage: ai-skill image-processing verify [flags] <image>")
		return ExitInvalidUsage
	}
	imgPath := rest[0]
	report, err := ipVerifyPath(imgPath, *profile, *source, *consumerProfile)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "verify: %v\n", err)
		return ExitGeneralFailure
	}
	text, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "marshal: %v\n", err)
		return ExitGeneralFailure
	}
	text = append(text, '\n')
	_, _ = stdout.Write(text)
	if *jsonOut != "" {
		if err := os.MkdirAll(filepath.Dir(*jsonOut), 0o755); err != nil {
			_, _ = fmt.Fprintf(stderr, "mkdir: %v\n", err)
			return ExitGeneralFailure
		}
		if err := os.WriteFile(*jsonOut, text, 0o644); err != nil {
			_, _ = fmt.Fprintf(stderr, "write json-out: %v\n", err)
			return ExitGeneralFailure
		}
	}
	finality, _ := report["finality"].(string)
	if *expectBlock {
		if finality == "accepted" {
			_, _ = fmt.Fprintln(stderr, "GATE FAIL: negative fixture was accepted")
			return 2
		}
		return ExitSuccess
	}
	if finality == "accepted" {
		return ExitSuccess
	}
	return 1
}

func runImageProcessingGenerateFixtures(args []string, stdout io.Writer, stderr io.Writer) int {
	fs := newFlagSet("image-processing generate-fixtures", stderr)
	repo := fs.String("repo", "", "Ai-skill repository root")
	outDir := fs.String("out-dir", "", "output directory (default: <repo>/workflow/image-processing/fixtures/inputs)")
	if err := fs.Parse(args); err != nil {
		return ExitInvalidUsage
	}
	dir := *outDir
	if dir == "" {
		root := *repo
		if root == "" {
			root = "."
		}
		dir = filepath.Join(root, "workflow", "image-processing", "fixtures", "inputs")
	}
	manifest, err := ipGenerateNegativeFixtures(dir)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "generate-fixtures: %v\n", err)
		return ExitGeneralFailure
	}
	for name, digest := range manifest {
		_, _ = fmt.Fprintf(stdout, "wrote %s sha256=%s\n", name, digest)
	}
	_, _ = fmt.Fprintf(stdout, "wrote %s\n", filepath.Join(dir, "manifest.json"))
	return ExitSuccess
}

var ipProfileChecks = map[string][]string{
	"neg.alpha-opaque":    {"alpha"},
	"neg.residue-halo":    {"alpha", "residue"},
	"neg.interior-damage": {"alpha", "interior"},
	"neg.over-erode":      {"alpha", "area"},
	"neg.consumer-layout": {"alpha", "consumer"},
	"pos.d-canvas-328":    {"alpha", "consumer"},
	"default":             {"alpha", "residue"},
}

func ipVerifyPath(path, profile, source, consumerProfile string) (map[string]any, error) {
	img, err := ipLoadRGBA(path)
	if err != nil {
		return nil, err
	}
	name := profile
	if name == "" {
		name = strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	}
	suites, ok := ipProfileChecks[name]
	if !ok {
		suites = ipProfileChecks["default"]
	}
	var checks []ipCheckResult
	for _, suite := range suites {
		switch suite {
		case "alpha":
			checks = append(checks, ipCheckAlpha(img)...)
		case "residue":
			checks = append(checks, ipCheckResidueRim(img))
		case "interior":
			srcPath := source
			if srcPath == "" {
				srcPath = ipDefaultRefFor(path, name)
			}
			if srcPath == "" {
				checks = append(checks, ipCheckResult{
					CheckID: "fidelity.interior", Status: "inconclusive",
					Evidence: "missing_source_ref", Hardness: "hard",
				})
			} else {
				src, err := ipLoadRGBA(srcPath)
				if err != nil {
					return nil, err
				}
				checks = append(checks, ipCheckInteriorVsSource(img, src))
			}
		case "area":
			srcPath := source
			if srcPath == "" {
				srcPath = ipDefaultRefFor(path, name)
			}
			if srcPath == "" {
				checks = append(checks, ipCheckResult{
					CheckID: "subject.area_vs_ref", Status: "inconclusive",
					Evidence: "missing_source_ref", Hardness: "hard",
				})
			} else {
				src, err := ipLoadRGBA(srcPath)
				if err != nil {
					return nil, err
				}
				checks = append(checks, ipCheckSubjectAreaVsSource(img, src, 0.35))
			}
		case "consumer":
			profPath := consumerProfile
			if profPath == "" {
				profPath = ipDefaultProfileJSON(path, name)
			}
			if profPath == "" {
				checks = append(checks, ipCheckResult{
					CheckID: "consumer.canvas", Status: "inconclusive",
					Evidence: "missing_profile_json", Hardness: "hard",
				})
			} else {
				prof, err := ipLoadProfile(profPath)
				if err != nil {
					return nil, err
				}
				checks = append(checks, ipCheckConsumerCanvas(img, prof))
			}
		}
	}
	finality := ipSuggestFinality(checks)
	out := map[string]any{
		"input":   ipRelpathForReport(path),
		"profile": name,
		"checks":  checks,
	}
	for k, v := range finality {
		out[k] = v
	}
	return out, nil
}

func ipLoadRGBA(path string) (*image.RGBA, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, err := png.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	b := img.Bounds()
	rgba := image.NewRGBA(b)
	draw.Draw(rgba, b, img, b.Min, draw.Src)
	return rgba, nil
}

func ipLoadProfile(path string) (map[string]any, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func ipDefaultRefFor(path, profile string) string {
	refs := filepath.Join(filepath.Dir(path), "refs", profile+".ref.png")
	if st, err := os.Stat(refs); err == nil && !st.IsDir() {
		return refs
	}
	return ""
}

func ipDefaultProfileJSON(path, profile string) string {
	p := filepath.Join(filepath.Dir(path), profile+".profile.json")
	if st, err := os.Stat(p); err == nil && !st.IsDir() {
		return p
	}
	return ""
}

func ipRelpathForReport(path string) string {
	resolved, err := filepath.Abs(path)
	if err != nil {
		return filepath.Base(path)
	}
	cur := resolved
	for {
		parent := filepath.Dir(cur)
		if parent == cur {
			break
		}
		if _, err := os.Stat(filepath.Join(parent, "workflow", "image-processing")); err == nil {
			if _, err := os.Stat(filepath.Join(parent, "CORE_BOOTSTRAP.md")); err == nil {
				if rel, err := filepath.Rel(parent, resolved); err == nil {
					return filepath.ToSlash(rel)
				}
			}
		}
		cur = parent
	}
	return filepath.Base(path)
}

func ipOpaqueCount(img *image.RGBA, threshold uint8) int {
	b := img.Bounds()
	n := 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if uint8(a>>8) >= threshold {
				n++
			}
		}
	}
	return n
}

func ipCheckAlpha(img *image.RGBA) []ipCheckResult {
	// PNG decode into RGBA always has alpha channel; detect "no transparency intent"
	// via fully-opaque pixels — matches Python path which required mode=RGBA first.
	out := []ipCheckResult{{
		CheckID: "alpha.present", Status: "pass", Evidence: "mode=RGBA", Hardness: "hard",
	}}
	b := img.Bounds()
	total := b.Dx() * b.Dy()
	opaque := 0
	amin := uint8(255)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			_, _, _, a16 := img.At(x, y).RGBA()
			a := uint8(a16 >> 8)
			if a < amin {
				amin = a
			}
			if a == 255 {
				opaque++
			}
		}
	}
	opaqueRatio := 0.0
	if total > 0 {
		opaqueRatio = float64(opaque) / float64(total)
	}
	ev := fmt.Sprintf("alpha_min=%d opaque_ratio=%.4f", amin, opaqueRatio)
	if amin >= 255 || opaqueRatio >= 0.995 {
		out = append(out, ipCheckResult{
			CheckID: "alpha.not_fully_opaque", Status: "fail", Evidence: ev,
			FailureClass: "alpha_invalid", RepairCandidate: "reassemble_rgba", Hardness: "hard",
		})
	} else {
		out = append(out, ipCheckResult{
			CheckID: "alpha.not_fully_opaque", Status: "pass", Evidence: ev, Hardness: "hard",
		})
	}
	return out
}

func ipCheckResidueRim(img *image.RGBA) ipCheckResult {
	b := img.Bounds()
	navyOpaque, transparent := 0, 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r16, g16, b16, a16 := img.At(x, y).RGBA()
			r, g, bb, a := uint8(r16>>8), uint8(g16>>8), uint8(b16>>8), uint8(a16>>8)
			if a == 0 {
				transparent++
			}
			if a >= 250 && bb >= 60 && r <= 40 && g <= 50 {
				navyOpaque++
			}
		}
	}
	ev := fmt.Sprintf("navy_opaque=%d transparent=%d", navyOpaque, transparent)
	if navyOpaque >= 50 && transparent >= 50 {
		return ipCheckResult{
			CheckID: "residue.rim", Status: "fail", Evidence: ev,
			FailureClass: "edge_halo", RepairCandidate: "post.decontaminate", Hardness: "hard",
		}
	}
	return ipCheckResult{CheckID: "residue.rim", Status: "pass", Evidence: ev, Hardness: "hard"}
}

func ipCheckInteriorVsSource(outImg, src *image.RGBA) ipCheckResult {
	if !outImg.Bounds().Eq(src.Bounds()) {
		return ipCheckResult{
			CheckID: "fidelity.interior", Status: "fail",
			Evidence:     fmt.Sprintf("size_mismatch out=%v src=%v", outImg.Bounds().Size(), src.Bounds().Size()),
			FailureClass: "fidelity_violation", RepairCandidate: "rollback_aggressive_morphology", Hardness: "hard",
		}
	}
	b := src.Bounds()
	punched, protected := 0, 0
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			sr, sg, sb, sa := src.At(x, y).RGBA()
			r, g, bb, a := uint8(sr>>8), uint8(sg>>8), uint8(sb>>8), uint8(sa>>8)
			if a >= 200 && int(r)+int(g)+int(bb) < 120 {
				protected++
				_, _, _, oa16 := outImg.At(x, y).RGBA()
				if uint8(oa16>>8) < 32 {
					punched++
				}
			}
		}
	}
	if protected == 0 {
		return ipCheckResult{
			CheckID: "fidelity.interior", Status: "inconclusive",
			Evidence: "no_protected_interior_marker", Hardness: "hard",
		}
	}
	ev := fmt.Sprintf("punched=%d protected=%d", punched, protected)
	if float64(punched)/float64(protected) >= 0.5 {
		return ipCheckResult{
			CheckID: "fidelity.interior", Status: "fail", Evidence: ev,
			FailureClass: "interior_damage", RepairCandidate: "protect_interior", Hardness: "hard",
		}
	}
	return ipCheckResult{CheckID: "fidelity.interior", Status: "pass", Evidence: ev, Hardness: "hard"}
}

func ipCheckSubjectAreaVsSource(outImg, src *image.RGBA, minRatio float64) ipCheckResult {
	so := ipOpaqueCount(src, 128)
	oo := ipOpaqueCount(outImg, 128)
	if so == 0 {
		return ipCheckResult{
			CheckID: "subject.area_vs_ref", Status: "inconclusive",
			Evidence: "ref_opaque=0", Hardness: "hard",
		}
	}
	ratio := float64(oo) / float64(so)
	ev := fmt.Sprintf("out_opaque=%d ref_opaque=%d ratio=%.3f", oo, so, ratio)
	if ratio < minRatio {
		return ipCheckResult{
			CheckID: "subject.area_vs_ref", Status: "fail", Evidence: ev,
			FailureClass: "subject_missing", RepairCandidate: "relax_threshold_or_change_seg", Hardness: "hard",
		}
	}
	return ipCheckResult{CheckID: "subject.area_vs_ref", Status: "pass", Evidence: ev, Hardness: "hard"}
}

func ipCheckConsumerCanvas(img *image.RGBA, profile map[string]any) ipCheckResult {
	wantRaw, ok := profile["canvas"]
	if !ok {
		return ipCheckResult{
			CheckID: "consumer.canvas", Status: "inconclusive",
			Evidence: "no_canvas_in_profile", Hardness: "hard",
		}
	}
	wantArr, ok := wantRaw.([]any)
	if !ok || len(wantArr) != 2 {
		return ipCheckResult{
			CheckID: "consumer.canvas", Status: "inconclusive",
			Evidence: "no_canvas_in_profile", Hardness: "hard",
		}
	}
	wantW := int(wantArr[0].(float64))
	wantH := int(wantArr[1].(float64))
	got := []int{img.Bounds().Dx(), img.Bounds().Dy()}
	if got[0] != wantW || got[1] != wantH {
		return ipCheckResult{
			CheckID: "consumer.canvas", Status: "fail",
			Evidence:     fmt.Sprintf("got=%v want=[%d %d]", got, wantW, wantH),
			FailureClass: "fidelity_violation", RepairCandidate: "post.deterministic", Hardness: "hard",
		}
	}
	return ipCheckResult{
		CheckID: "consumer.canvas", Status: "pass",
		Evidence: fmt.Sprintf("got=%v", got), Hardness: "hard",
	}
}

func ipSuggestFinality(checks []ipCheckResult) map[string]any {
	var hardFails, hardInconc []string
	review := false
	for _, c := range checks {
		if c.Hardness == "hard" && c.Status == "fail" {
			hardFails = append(hardFails, c.CheckID)
		}
		if c.Hardness == "hard" && c.Status == "inconclusive" {
			hardInconc = append(hardInconc, c.CheckID)
		}
		if c.ReviewRequired {
			review = true
		}
	}
	if len(hardFails) > 0 {
		return map[string]any{
			"disposition": "reject_or_abort", "finality": "rejected",
			"reason": "hard_check_fail", "failed_checks": hardFails,
		}
	}
	if len(hardInconc) > 0 {
		return map[string]any{
			"disposition": "human_review", "finality": nil,
			"reason": "hard_inconclusive",
		}
	}
	if review {
		return map[string]any{
			"disposition": "queue_human_review", "finality": "accepted_with_review",
			"reason": "soft_review_required",
		}
	}
	return map[string]any{
		"disposition": "close", "finality": "accepted",
		"reason": "all_hard_pass",
	}
}

// --- fixture generation ---

func ipGenerateNegativeFixtures(outDir string) (map[string]string, error) {
	refDir := filepath.Join(outDir, "refs")
	if err := os.MkdirAll(refDir, 0o755); err != nil {
		return nil, err
	}
	manifest := map[string]string{}
	simple := map[string]func() *image.RGBA{
		"neg.alpha-opaque":    func() *image.RGBA { return ipMakeAlphaOpaque(64) },
		"neg.residue-halo":    func() *image.RGBA { return ipMakeResidueHalo(96) },
		"neg.consumer-layout": func() *image.RGBA { return ipMakeConsumerLayout(64) },
		"pos.d-canvas-328":    func() *image.RGBA { return ipMakePosDCanvas(328) },
	}
	for name, factory := range simple {
		digest, err := ipWritePNG(filepath.Join(outDir, name+".png"), factory())
		if err != nil {
			return nil, err
		}
		manifest[name] = digest
	}
	src, bad := ipMakeInteriorPair(96)
	d, err := ipWritePNG(filepath.Join(refDir, "neg.interior-damage.ref.png"), src)
	if err != nil {
		return nil, err
	}
	manifest["neg.interior-damage.ref"] = d
	d, err = ipWritePNG(filepath.Join(outDir, "neg.interior-damage.png"), bad)
	if err != nil {
		return nil, err
	}
	manifest["neg.interior-damage"] = d

	ref, eroded := ipMakeOverErodePair(96)
	d, err = ipWritePNG(filepath.Join(refDir, "neg.over-erode.ref.png"), ref)
	if err != nil {
		return nil, err
	}
	manifest["neg.over-erode.ref"] = d
	d, err = ipWritePNG(filepath.Join(outDir, "neg.over-erode.png"), eroded)
	if err != nil {
		return nil, err
	}
	manifest["neg.over-erode"] = d

	for _, stem := range []string{"neg.consumer-layout", "pos.d-canvas-328"} {
		body := map[string]any{
			"consumer_profile": "slot-symbol",
			"canvas":           []int{328, 328},
		}
		b, _ := json.MarshalIndent(body, "", "  ")
		b = append(b, '\n')
		if err := os.WriteFile(filepath.Join(outDir, stem+".profile.json"), b, 0o644); err != nil {
			return nil, err
		}
	}
	mb, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return nil, err
	}
	mb = append(mb, '\n')
	if err := os.WriteFile(filepath.Join(outDir, "manifest.json"), mb, 0o644); err != nil {
		return nil, err
	}
	return manifest, nil
}

func ipWritePNG(path string, img *image.RGBA) (string, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	if err := png.Encode(f, img); err != nil {
		_ = f.Close()
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(raw)
	digest := hex.EncodeToString(sum[:])
	if err := os.WriteFile(strings.TrimSuffix(path, filepath.Ext(path))+".sha256", []byte(digest+"\n"), 0o644); err != nil {
		return "", err
	}
	return digest, nil
}

func ipNewRGBA(size int, c color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(img, img.Bounds(), &image.Uniform{C: c}, image.Point{}, draw.Src)
	return img
}

func ipFillEllipse(img *image.RGBA, minX, minY, maxX, maxY int, c color.RGBA) {
	cx := float64(minX+maxX) / 2
	cy := float64(minY+maxY) / 2
	rx := float64(maxX-minX) / 2
	ry := float64(maxY-minY) / 2
	if rx <= 0 || ry <= 0 {
		return
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			dx := (float64(x) - cx) / rx
			dy := (float64(y) - cy) / ry
			if dx*dx+dy*dy <= 1 {
				img.SetRGBA(x, y, c)
			}
		}
	}
}

func ipFillRect(img *image.RGBA, minX, minY, maxX, maxY int, c color.RGBA) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			img.SetRGBA(x, y, c)
		}
	}
}

func ipMakeAlphaOpaque(size int) *image.RGBA {
	img := ipNewRGBA(size, color.RGBA{40, 120, 200, 255})
	ipFillEllipse(img, 12, 12, size-12, size-12, color.RGBA{220, 180, 60, 255})
	return img
}

func ipMakeResidueHalo(size int) *image.RGBA {
	img := ipNewRGBA(size, color.RGBA{0, 0, 0, 0})
	cx, cy := size/2, size/2
	ipFillEllipse(img, cx-34, cy-34, cx+34, cy+34, color.RGBA{10, 20, 80, 255})
	ipFillEllipse(img, cx-26, cy-26, cx+26, cy+26, color.RGBA{230, 190, 70, 255})
	return img
}

func ipMakeInteriorPair(size int) (src, bad *image.RGBA) {
	src = ipNewRGBA(size, color.RGBA{0, 0, 0, 0})
	cx, cy := size/2, size/2
	ipFillEllipse(src, cx-30, cy-30, cx+30, cy+30, color.RGBA{230, 190, 70, 255})
	ipFillRect(src, cx-8, cy-4, cx+8, cy+10, color.RGBA{40, 30, 20, 255})
	bad = image.NewRGBA(src.Bounds())
	draw.Draw(bad, bad.Bounds(), src, image.Point{}, draw.Src)
	ipFillRect(bad, cx-8, cy-4, cx+8, cy+10, color.RGBA{0, 0, 0, 0})
	return src, bad
}

func ipMakeOverErodePair(size int) (ref, bad *image.RGBA) {
	ref = ipNewRGBA(size, color.RGBA{0, 0, 0, 0})
	cx, cy := size/2, size/2
	ipFillEllipse(ref, cx-36, cy-36, cx+36, cy+36, color.RGBA{200, 160, 80, 255})
	bad = ipNewRGBA(size, color.RGBA{0, 0, 0, 0})
	ipFillEllipse(bad, cx-12, cy-12, cx+12, cy+12, color.RGBA{200, 160, 80, 255})
	return ref, bad
}

func ipMakeConsumerLayout(size int) *image.RGBA {
	img := ipNewRGBA(size, color.RGBA{0, 0, 0, 0})
	ipFillEllipse(img, 8, 8, size-8, size-8, color.RGBA{210, 170, 70, 255})
	return img
}

func ipMakePosDCanvas(size int) *image.RGBA {
	img := ipNewRGBA(size, color.RGBA{0, 0, 0, 0})
	cx, cy := size/2, size/2
	r := 90
	ipFillEllipse(img, cx-r, cy-r, cx+r, cy+r, color.RGBA{210, 170, 70, 255})
	return img
}
