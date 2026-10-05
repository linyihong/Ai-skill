# Text evidence — OCR Discovery / Layout Probe

何時讀：硬字幕 OCR 在預設 bottom band 報 `no_subtitle_like`／`insufficient`；
字幕實際在 upper-middle／非預設區；或產品把 probe miss 當成「影片沒字幕」而停掃。
Plan：[`45-ocr-discovery-layout-probe.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/45-ocr-discovery-layout-probe.md)。
銜接：[`13-mechanical-visual-text-probe.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/13-mechanical-visual-text-probe.md)、
[`text-evidence-regions.md`](text-evidence-regions.md)、[`24-ocr-role-projection.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/24-ocr-role-projection.md)、
[`text-evidence-acquisition-loop.md`](text-evidence-acquisition-loop.md)（跨模態 Monitor／時間窗 gap）、
[`text-evidence-subtitle-candidate.md`](text-evidence-subtitle-candidate.md)（發現後的 precision／雜訊過濾）。

> **執行契約，不重新解釋契約。** Probe = discovery，不是 exclusion。不開新 Phase／OCR Agent。

## 一句話

**OCR Probe 是 discovery，不是 exclusion。** 未命中預設字幕區必須保留 recovery path；
字幕位置由 evidence-driven layout discovery 決定；mechanical OCR 負責實際採集；
LLM 只做 hypothesis／selection，不直接決定 crop 或斷言「沒有字幕」。

## 禁止的語意塌縮

| 錯誤 | 正確 |
| --- | --- |
| `no_subtitle_like` → 影片沒字幕／停止 OCR | `status: inconclusive`；`reason: no_match_in_current_region` |
| 固定 bottom probe 失敗 → `subtitle.exists=false` | 只記錄「目前假設區未命中」 |
| LLM 回 `subtitle_y: 0.73` → 寫死 crop | LLM 只產 `layout_hypothesis`／role classification |
| 一開始每幀 full OCR | 低成本 global discovery 抽樣 → layout candidates → targeted OCR |

## 兩層 probe（採納順序）

```text
① Mechanical global discovery   低成本全畫面取樣（代表 frame）
        ↓
   Layout candidates（bbox／script／area／persistence／center）
        ↓
   enough evidence?
      │ yes → Targeted OCR（scan_profile regions）
      │ no  → LLM／resolver layout classification（escape／selection）
        ↓
   Layout hypothesis → Targeted OCR → Evidence → coverage validate
        ↓
   insufficient? → rediscovery（非終止）
```

**採納**：Mechanical discovery → LLM／resolver layout classification → Mechanical targeted OCR。
**拒絕**：LLM 先猜位置 → 機械只掃那裡（LLM 錯＝資料遺失）。

## Layout Discovery Probe（不是最終字幕 OCR）

抽樣代表 frame（均勻比例或 scene-aware）。Full-frame **lightweight** OCR 只收 evidence：

- `bbox`／`normalized_box`（相對 frame；禁止絕對 px 當規則）
- `text`／`script`／`area_ratio`／`aspect_ratio`／`center`
- **typography**：`height_ratio`／`width_ratio`／`estimated_font_size_band`
- **temporal**：persistence／text_stability／box_stability／replacement_rate（若可得）
- color／visual attributes（若可得）

產出 `subtitle_region_candidate[]`（含 watermark／title／livestream_chat／platform_ui 候選），**不要**只問 `cy >= 0.68`。

流程（機械先）：OCR boxes → geometry clustering → region behavior → `region_role` → subtitle candidate；LLM 只對 unknown／衝突 region 看少量代表幀。詳見 [text-evidence-subtitle-candidate.md](./text-evidence-subtitle-candidate.md) 的 Layout & Typography Evidence，以及 lesson `2026-10-02_172938-ocr-layout-typography-are-evidence-features`。

## LLM 的正確位置

優先讓 LLM 判：**哪一組 detections／regions 像 dialogue／watermark／title／livestream_chat／platform_ui**（給 sample frames + mechanical boxes），
而不是猜單一 `y=`，也不是當 OCR engine。輸出是 hypothesis：

```text
layout_hypothesis:
  subtitle_regions: [{ region, likelihood }, …]
  text_types: { dialogue, watermark, livestream_chat, platform_ui, … }
  ocr_priority / asr_priority
  exclusion_hints: […]   # 為何某區不應進 dialogue evidence
```

清楚機械特徵不必打 LLM（見 13／24）。

## Per-source OCR profile（observation，可版本化）

作品級累積 `source_ocr_profile`（**不是** canonical global rule）：

```text
source_ocr_profile:
  version: 1
  regions[]:
    region_id, normalized_box
    behavior: persistence | text_repetition | position_stability | change_rate | sticky?
    role: subtitle_candidate | watermark_candidate | livestream_chat | platform_ui | unknown
  scan_policy: include_regions / exclude_roles / sticky_spans[]
```

後續集直接用 profile → mechanical targeted OCR；**layout drift** 才觸發 rediscovery → profile v2。
對齊 Observation → Accumulation → Validation → Promotion；**禁止**單片 dogfood 立刻寫死 `watermark_y`／全域刪字字典。

### Sticky / persistent region

若 region：長時存在、bbox 穩、字級穩、文本高度重複、不隨對白變 → `sticky: true` → 預設 `watermark_candidate`（機械即可）。
正式 OCR 應盡量 **空間排除** sticky layer；若 OCR 仍把 sticky 與對白黏成一字串，走 **role-aware／sticky projection**（保留 observed，只投影 spoken）— 見 sticky-watermark evidence。
本集高重複 span 可進 `scan_policy.sticky_spans`（source-local），不得晉升為全域詞表除非跨源 fixture 驗證。

### Profile impact（必記，防「修這部壞那部」）

`profile_change: v1 → v2` 至少對照：

| 指標 | 用途 |
| --- | --- |
| OCR／dialogue candidates before／after | 掃描寬度是否塌縮 |
| watermark_rejections／sticky peels | sticky 是否生效 |
| accepted／uncertain／rejected | Evidence≠Production 分層 |
| small fixture regression | 難例 clip corpus（非整部上一部） |

risk 用 fixture 差計量，不用「感覺上一部怪怪的」。詳見 evidence [`2026-10-05-source-adaptive-ocr-region-profile`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/evidence/2026-10-05-source-adaptive-ocr-region-profile.md)。

## Coverage 指標（必記）

不要只記 `subtitle_like: 0`。至少區分：

| 層 | 例 |
| --- | --- |
| probe（假設區） | `bottom.subtitle_like` |
| discovery | `upper_middle.dialogue_candidate_count` |
| coverage | `sampled_frames`／`frames_with_text`／`frames_with_dialogue_candidate`／`subtitle_region_coverage` |
| layout | `discovered_regions[]`：`y_range`、`role`、`evidence_count` |

Agent 應能讀出：**不是沒字幕，是 default layout hypothesis 錯了。**

## 在 lifecycle 的位置

插在 execution-flow **6b** 最前：Layout Discovery／scan_profile → Targeted OCR raw → boundary／regions → …
不取代 boundary、Language Relation、Text Resolution。

## Adapter 驗收（產品後改）

1. 預設 bottom miss → `inconclusive`，不得當 exclusion／STOP。
2. Recovery：global discovery 或 expand 後能建立非 bottom `scan_profile` 並 targeted OCR。
3. LLM 不得單獨寫死 crop；不得在無 mechanical discovery 時斷言無字幕。
4. Metrics 同時露出 probe miss 與 discovery hit。
