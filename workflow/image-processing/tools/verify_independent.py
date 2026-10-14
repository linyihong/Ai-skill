#!/usr/bin/env python3
"""Independent image verification — must not be the executor."""

from __future__ import annotations

import argparse
import json
import sys
from dataclasses import asdict, dataclass
from pathlib import Path
from typing import Any

from PIL import Image


@dataclass
class CheckResult:
    check_id: str
    status: str  # pass | fail | inconclusive
    evidence: str
    failure_class: str | None = None
    repair_candidate: str | None = None
    review_required: bool = False
    hardness: str = "hard"


def opaque_count(img: Image.Image, threshold: int = 128) -> int:
    a = img.getchannel("A")
    return sum(1 for v in a.getdata() if v >= threshold)


def check_alpha(img: Image.Image) -> list[CheckResult]:
    out: list[CheckResult] = []
    if img.mode != "RGBA":
        out.append(
            CheckResult(
                check_id="alpha.present",
                status="fail",
                evidence=f"mode={img.mode}",
                failure_class="alpha_invalid",
                repair_candidate="reassemble_rgba",
                hardness="hard",
            )
        )
        return out

    out.append(
        CheckResult(
            check_id="alpha.present",
            status="pass",
            evidence="mode=RGBA",
            hardness="hard",
        )
    )

    alpha = img.getchannel("A")
    hist = alpha.histogram()
    total = img.width * img.height
    opaque = hist[255] if len(hist) > 255 else 0
    opaque_ratio = opaque / total if total else 1.0
    amin, _amax = alpha.getextrema()

    if amin >= 255 or opaque_ratio >= 0.995:
        out.append(
            CheckResult(
                check_id="alpha.not_fully_opaque",
                status="fail",
                evidence=f"alpha_min={amin} opaque_ratio={opaque_ratio:.4f}",
                failure_class="alpha_invalid",
                repair_candidate="reassemble_rgba",
                hardness="hard",
            )
        )
    else:
        out.append(
            CheckResult(
                check_id="alpha.not_fully_opaque",
                status="pass",
                evidence=f"alpha_min={amin} opaque_ratio={opaque_ratio:.4f}",
                hardness="hard",
            )
        )
    return out


def check_residue_rim(img: Image.Image) -> CheckResult:
    if img.mode != "RGBA":
        return CheckResult(
            check_id="residue.rim",
            status="inconclusive",
            evidence="no RGBA",
            hardness="hard",
        )

    px = img.load()
    w, h = img.size
    navy_opaque = 0
    transparent = 0
    for y in range(h):
        for x in range(w):
            r, g, b, a = px[x, y]
            if a == 0:
                transparent += 1
            if a >= 250 and b >= 60 and r <= 40 and g <= 50:
                navy_opaque += 1

    if navy_opaque >= 50 and transparent >= 50:
        return CheckResult(
            check_id="residue.rim",
            status="fail",
            evidence=f"navy_opaque={navy_opaque} transparent={transparent}",
            failure_class="edge_halo",
            repair_candidate="post.decontaminate",
            hardness="hard",
        )
    return CheckResult(
        check_id="residue.rim",
        status="pass",
        evidence=f"navy_opaque={navy_opaque} transparent={transparent}",
        hardness="hard",
    )


def check_interior_vs_source(out_img: Image.Image, src: Image.Image) -> CheckResult:
    """Fail when source interior opaque pixels became transparent in output."""
    if out_img.size != src.size:
        return CheckResult(
            check_id="fidelity.interior",
            status="fail",
            evidence=f"size_mismatch out={out_img.size} src={src.size}",
            failure_class="fidelity_violation",
            repair_candidate="rollback_aggressive_morphology",
            hardness="hard",
        )
    op = out_img.load()
    sp = src.load()
    punched = 0
    protected = 0
    w, h = src.size
    for y in range(h):
        for x in range(w):
            sr, sg, sb, sa = sp[x, y]
            # dark interior marker in synthetic fixtures
            if sa >= 200 and (sr + sg + sb) < 120:
                protected += 1
                _or, _og, _ob, oa = op[x, y]
                if oa < 32:
                    punched += 1
    if protected == 0:
        return CheckResult(
            check_id="fidelity.interior",
            status="inconclusive",
            evidence="no_protected_interior_marker",
            hardness="hard",
        )
    if punched / protected >= 0.5:
        return CheckResult(
            check_id="fidelity.interior",
            status="fail",
            evidence=f"punched={punched} protected={protected}",
            failure_class="interior_damage",
            repair_candidate="protect_interior",
            hardness="hard",
        )
    return CheckResult(
        check_id="fidelity.interior",
        status="pass",
        evidence=f"punched={punched} protected={protected}",
        hardness="hard",
    )


def check_subject_area_vs_source(
    out_img: Image.Image, src: Image.Image, min_ratio: float = 0.35
) -> CheckResult:
    so = opaque_count(src)
    oo = opaque_count(out_img)
    if so == 0:
        return CheckResult(
            check_id="subject.area_vs_ref",
            status="inconclusive",
            evidence="ref_opaque=0",
            hardness="hard",
        )
    ratio = oo / so
    # Clean residue but over-eroded → residue may pass; this check fails
    if ratio < min_ratio:
        return CheckResult(
            check_id="subject.area_vs_ref",
            status="fail",
            evidence=f"out_opaque={oo} ref_opaque={so} ratio={ratio:.3f}",
            failure_class="subject_missing",
            repair_candidate="relax_threshold_or_change_seg",
            hardness="hard",
        )
    return CheckResult(
        check_id="subject.area_vs_ref",
        status="pass",
        evidence=f"out_opaque={oo} ref_opaque={so} ratio={ratio:.3f}",
        hardness="hard",
    )


def check_consumer_canvas(img: Image.Image, profile: dict[str, Any]) -> CheckResult:
    want = profile.get("canvas")
    if not want or len(want) != 2:
        return CheckResult(
            check_id="consumer.canvas",
            status="inconclusive",
            evidence="no_canvas_in_profile",
            hardness="hard",
        )
    got = [img.width, img.height]
    if got != list(want):
        return CheckResult(
            check_id="consumer.canvas",
            status="fail",
            evidence=f"got={got} want={want}",
            failure_class="fidelity_violation",
            repair_candidate="post.deterministic",
            hardness="hard",
        )
    return CheckResult(
        check_id="consumer.canvas",
        status="pass",
        evidence=f"got={got}",
        hardness="hard",
    )


def suggest_finality(checks: list[CheckResult]) -> dict[str, Any]:
    hard_fails = [c for c in checks if c.hardness == "hard" and c.status == "fail"]
    hard_inconc = [
        c for c in checks if c.hardness == "hard" and c.status == "inconclusive"
    ]
    review = any(c.review_required for c in checks)

    if hard_fails:
        return {
            "disposition": "reject_or_abort",
            "finality": "rejected",
            "reason": "hard_check_fail",
            "failed_checks": [c.check_id for c in hard_fails],
        }
    if hard_inconc:
        return {
            "disposition": "human_review",
            "finality": None,
            "reason": "hard_inconclusive",
        }
    if review:
        return {
            "disposition": "queue_human_review",
            "finality": "accepted_with_review",
            "reason": "soft_review_required",
        }
    return {
        "disposition": "close",
        "finality": "accepted",
        "reason": "all_hard_pass",
    }


PROFILE_CHECKS = {
    "neg.alpha-opaque": ["alpha"],
    "neg.residue-halo": ["alpha", "residue"],
    "neg.interior-damage": ["alpha", "interior"],
    "neg.over-erode": ["alpha", "area"],
    "neg.consumer-layout": ["alpha", "consumer"],
    "pos.d-canvas-328": ["alpha", "consumer"],
    "default": ["alpha", "residue"],
}


def relpath_for_report(path: Path) -> str:
    resolved = path.resolve()
    for parent in resolved.parents:
        if (parent / "workflow" / "image-processing").is_dir() and (
            parent / "CORE_BOOTSTRAP.md"
        ).is_file():
            try:
                return str(resolved.relative_to(parent))
            except ValueError:
                break
    return path.name


def default_ref_for(path: Path, profile: str) -> Path | None:
    refs = path.parent / "refs" / f"{profile}.ref.png"
    if refs.is_file():
        return refs
    return None


def default_profile_json(path: Path, profile: str) -> Path | None:
    p = path.parent / f"{profile}.profile.json"
    return p if p.is_file() else None


def verify_path(
    path: Path,
    profile: str | None = None,
    source: Path | None = None,
    consumer_profile: Path | None = None,
) -> dict[str, Any]:
    img = Image.open(path)
    img.load()
    name = profile or path.stem
    suites = PROFILE_CHECKS.get(name, PROFILE_CHECKS["default"])
    checks: list[CheckResult] = []
    if "alpha" in suites:
        checks.extend(check_alpha(img))
    if "residue" in suites:
        checks.append(check_residue_rim(img))
    if "interior" in suites:
        src_path = source or default_ref_for(path, name)
        if src_path is None:
            checks.append(
                CheckResult(
                    check_id="fidelity.interior",
                    status="inconclusive",
                    evidence="missing_source_ref",
                    hardness="hard",
                )
            )
        else:
            src = Image.open(src_path)
            src.load()
            checks.append(check_interior_vs_source(img, src))
    if "area" in suites:
        src_path = source or default_ref_for(path, name)
        if src_path is None:
            checks.append(
                CheckResult(
                    check_id="subject.area_vs_ref",
                    status="inconclusive",
                    evidence="missing_source_ref",
                    hardness="hard",
                )
            )
        else:
            src = Image.open(src_path)
            src.load()
            checks.append(check_subject_area_vs_source(img, src))
    if "consumer" in suites:
        prof_path = consumer_profile or default_profile_json(path, name)
        if prof_path is None:
            checks.append(
                CheckResult(
                    check_id="consumer.canvas",
                    status="inconclusive",
                    evidence="missing_profile_json",
                    hardness="hard",
                )
            )
        else:
            prof = json.loads(prof_path.read_text(encoding="utf-8"))
            checks.append(check_consumer_canvas(img, prof))

    finality = suggest_finality(checks)
    return {
        "input": relpath_for_report(path),
        "profile": name,
        "checks": [asdict(c) for c in checks],
        **finality,
    }


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("image", type=Path)
    parser.add_argument("--profile", default=None)
    parser.add_argument("--source", type=Path, default=None)
    parser.add_argument("--consumer-profile", type=Path, default=None)
    parser.add_argument("--json-out", type=Path, default=None)
    parser.add_argument(
        "--expect-block",
        action="store_true",
        help="Exit 0 only if finality is not accepted (negative fixture mode)",
    )
    args = parser.parse_args()

    report = verify_path(
        args.image,
        args.profile,
        source=args.source,
        consumer_profile=args.consumer_profile,
    )
    text = json.dumps(report, indent=2)
    print(text)
    if args.json_out:
        args.json_out.parent.mkdir(parents=True, exist_ok=True)
        args.json_out.write_text(text + "\n", encoding="utf-8")

    finality = report.get("finality")
    if args.expect_block:
        if finality == "accepted":
            print("GATE FAIL: negative fixture was accepted", file=sys.stderr)
            return 2
        return 0

    if finality == "accepted":
        return 0
    return 1


if __name__ == "__main__":
    raise SystemExit(main())
