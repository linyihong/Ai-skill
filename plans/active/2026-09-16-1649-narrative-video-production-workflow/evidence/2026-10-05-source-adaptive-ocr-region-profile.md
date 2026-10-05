# Observation — source-adaptive OCR / region profile (not per-series rules)

**Run ID**：2026-10-05-source-adaptive-ocr-region-profile
**Kind**：Phase 3 adapter／contract_gap（dogfood《春风拂我心》；抽象後可重用）
**Extends**：[`2026-10-02-sticky-watermark-accepted-projection`](2026-10-02-sticky-watermark-accepted-projection.md)、[`2026-10-01-ocr-probe-discovery-not-exclusion`](2026-10-01-ocr-probe-discovery-not-exclusion.md)、[`45-ocr-discovery-layout-probe`](../45-ocr-discovery-layout-probe.md)

## 問題（sanitized）

短劇硬字幕片源常同時有：

- 底部／中置對白字幕
- 中置／全寬 **sticky** 免責水印（高重複、跨 scene）
- 平台 UI／直播留言

固定 OCR band 或「這部片手調 watermark_y／字典刪字」會造成：

1. 修 A 片 → 怕壞 B 片（無 impact report）
2. OCR 已把水印與對白黏成一字串後，字典刪字不穩
3. probe `subtitle_like=0` 被誤讀成「沒字幕」（本例仍 `has_hardsub=true`、窗內 7–8 句、cues≈40）

本例價值：**不是 cues=0**，而是 **sticky layer 污染 accepted 品質** + **缺 source-local profile**。

## 目標架構（契約）

```text
Source Video
  → Mechanical sparse probe (representative frames; may be wider than final OCR)
  → OCR boxes
  → Region clustering
  → Region behavior profiling (persistence / repetition / stability / change_rate)
  → source_ocr_profile (observation; versioned; NOT global rule)
  → scan policy: exclude sticky layers spatially when possible
  → targeted OCR + ASR fusion
  → sticky / role-aware projection (spoken) if layers already glued in text
  → accept / uncertain / reject
  → profile impact + small fixture regression
```

## 規則

1. **Global contract ≠ Source observation**：單片發現不得寫死 `watermark_y=0.45` 進全域規則；寫進 `source_ocr_profile`（可版本化 v1→v2）。
2. **Region first, text second**：優先用空間／時序把 watermark region 與 dialogue region 分開，再 OCR／理解；字典刪黏字串是 fallback，不是主路徑。
3. **Sticky = behavior**：長時存在 + bbox 穩 + 文本高度重複 + 不隨對白變 → `sticky: true`／`watermark_candidate`；不需要 LLM。
4. **Probe 可寬、正式 OCR 可窄**：probe 任務是學「文字在哪、怎麼出現」；正式掃依 profile。
5. **LLM 只做少量校準**：對 ambiguous regions 看代表幀標 role；一部 source 一次低成本 calibration，其後機械執行。
6. **Profile impact**：v1→v2 必須產出 candidates／accepted／uncertain／rejected 前後對照；有小 fixture regression 時標 risk。
7. **不開新 Phase**：接在既有 OCR discovery → regions → sticky projection → resolution 鏈。

## Dogfood（sanitized）

| 觀測 | 值 |
| --- | --- |
| probe | hardsub=true, speech=true, source=ocr；mechanical `subtitle_like=0` / dialogue_cand=0 / inconclusive |
| layout discovery | center 帶 `dialogue_subtitle` 候選；另有 upper_middle／bottom |
| window subs | 每 30s 窗約 7–8 句（非空） |
| cues | publishable≈40；多條含高重複免責水印黏串 |
| 錯誤做法 | 為本片寫死中置 y 或全域「仅供娱乐…」字典 |
| 正確方向 | source profile sticky region + spatial exclude + text projection from learned sticky spans |

## Validation

- [x] 本 evidence + README 索引
- [x] workflow ocr-discovery 補 source profile／sticky／impact
- [x] feedback lesson（candidate）
- [x] 產品：本集自學 sticky span peel（無全域單片字典）+ 春风拂我心 抽檢污染下降（text peel 9→0）
- [x] 產品：`source_ocr_profile` discovery behavior + fuse 前 drop pure sticky（probe 純水印行 2 條）
- [x] 產品：OCR 插入字 sticky peel（bounded subsequence）+ sticky fragment drop；filter `impact` 計數
- [ ] profile v1/v2 跨版 fixture regression corpus（後續 adapter）

## 產品落點

1. `sticky_ocr_projection`：本批 OCR observed → episode-local sticky spans peel（spoken）；保留 observed。
2. `source_ocr_profile`：layout region `behavior.sticky` → `scan_profile.sticky_spans`／`exclude_boxes`；fuse 前 drop **純** sticky 行；黏串對白仍交給 projection。
3. **禁止**單片名 watermark 字串寫進全域常數。
