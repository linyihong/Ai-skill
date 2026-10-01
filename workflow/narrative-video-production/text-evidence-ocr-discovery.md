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

- `bbox`／`normalized_box`
- `text`／`script`／`area`／`aspect_ratio`／`center`
- persistence／color（若可得）

產出 `subtitle_region_candidate[]`（含 watermark／title 候選），**不要**只問 `cy >= 0.68`。

## LLM 的正確位置

優先讓 LLM 判：**哪一組 detections 像 dialogue／watermark／title**（給 sample frames + mechanical boxes），
而不是猜單一 `y=`。輸出是 hypothesis：

```text
layout_hypothesis:
  subtitle_regions: [{ region, likelihood }, …]
  text_types: { dialogue, watermark, … }
  ocr_priority / asr_priority
```

清楚機械特徵不必打 LLM（見 13／24）。

## Per-source OCR profile

作品級可累積 `source_ocr_profile`（primary／secondary subtitle bands、watermark zones）。
後續集直接用 profile → mechanical targeted OCR；**layout drift** 才觸發 rediscovery。
對齊 Observation → Accumulation → Validation → Promotion；禁止單次 LLM 觀察立刻改全局規則。

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
