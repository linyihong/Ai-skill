# Phase 3 refinement — Evidence 歸屬、分類與保留

使用者於 2026-10-06 確認的優先順序；本檔是後續 adapter 驗收契約，並非產品 PASS 記錄。
前置：[47 candidate precision](47-subtitle-candidate-detector.md)、
[既有局部驗證](evidence/2026-10-06-attribution-precision-and-acquisition-integrity.md)。
Workflow：[逐筆對帳](../../../workflow/narrative-video-production/text-evidence-multimodal-resolution.md)、
[region 身份](../../../workflow/narrative-video-production/text-evidence-regions.md)、
[欄位契約](../../../workflow/narrative-video-production/records/text-evidence.yaml)。

## Phase 3 observation 與範圍

目前已報告的反例支持：OCR acquisition 已抓到關鍵文字，主要缺口在
Evidence Attribution / Classification / Retention。這是本輪樣本的診斷，
不代表所有來源的 OCR recall 或採集完整度已驗收；資源截斷仍獨立檢查。

本輪先不提高全域 OCR 頻率，也不增加更大的 LLM agent。
先修 `raw → attribution → classification → fusion → 四態 → traceable cue`。
既有 targeted probe 保留；只有獨立原片核對證明「字幕存在但 raw OCR 未取得」，
且排除 cache、採集截斷與下游遺失後，才評估 adaptive probe／recall upgrade。

## 核心 invariant

每筆 candidate 必有 `accepted | uncertain | rejected + reason | merged + target`。
此要求從清理站開始，無須先觸發 resolution-loss anomaly。Raw 不可覆寫，
region ownership 不可借用；final cue 必須可回溯 evidence 身份與各站去向。
`uncertain` 是有效 evidence，可經 targeted probe、ASR corroboration、candidate
reconstruction 再決議；不能因不進成片就刪除。

## 優先順序與驗收

| 優先級 | 工作 | 必須同時驗收 |
| --- | --- | --- |
| P1 | 移除英文 script＋長度硬刪 | 英文短字幕保留；品牌／Logo／watermark／scene text 不自動進 dialogue；黏字仍進 boundary recovery |
| P2 | 雙語 box／region ownership | 各語 region identity、raw text、owned box 經 fusion 保留；缺可靠 box 時空間關係 unresolved；stacked 需 geometry＋temporal＋language 支持 |
| P3 | 8 部代表樣本回歸矩陣 | 每例有 expected behavior、正負例、candidate 去向與驗證範圍；缺 cache／未跑明記 unverified；分批驗證可推進，不等待八部全片重跑 |
| P4 | 時間／ASS／成片 | identity → content → region → language → temporal alignment → ASS → rendered video；不得用合併不同台詞消除 overlap |

P1–P2 每次修復使用凍結輸入與 source-frame 正負對照；不能只看污染字串消失。
不增加逐片 if 規則；保留 geometry、language、typography、temporal persistence、
region identity 與 raw text 能力，由上層 policy 使用。

## 8 部矩陣（案例槽位，待對應真實來源）

下表是測試設計，並非已跑 evidence；真實媒體身份與 artifact references 留在 consumer repo。
每例實際結果需記 PASS／FAIL／unverified，以及 scoped replay／fresh acquisition／render 範圍。

| Case | 英文 | 雙語 | 水印 | 高字幕 | 中置字幕 | ASR/OCR conflict | Expected behavior |
| --- | --- | --- | --- | --- | --- | --- | --- |
| A | ✓ | | ✓ | | | | 短英文對白保留；同窗 Logo／品牌／水印不進 dialogue |
| B | ✓ | ✓ | | | | ✓ | 語言分開、spoken/subtitle 不混；缺 box 不推 stacked |
| C | | ✓ | ✓ | ✓ | ✓ | | 各 region 自有文字與 geometry；水印擴詞不吞字幕首字 |
| D | ✓ | | | | | ✓ | Latin 黏字保留 raw，boundary recovery 可追；未收斂保留 uncertain |
| E | | | ✓ | | | | 短告別句完整對帳，未入 final 仍可查 reason／merge target |
| F | | | | ✓ | | | 高位真字幕保留；UI／鐘錶負例不得借字幕 band 升格 |
| G | | | | | ✓ | ✓ | 中置對白歸屬保留；ASR 缺支持不硬刪 OCR |
| H | ✓ | ✓ | ✓ | | | ✓ | 雙語 fusion 到 cue 的 identity trace 完整；時間修復不吞句 |

使用者報告的 `Well`、字幕首字「我」、Latin 黏字及告別句作為 regression watches；
原片／raw／candidate／final 的獨立 readback 尚未由本次文件修改驗證。
數字示例 OCR=529、candidate=208、final=40 僅說明需對帳，不能當實測或全域門檻。

## Acceptance（產品驗證保持 open）

- [x] Workflow 與 records 補 unconditional disposition／owned geometry 契約。
- [x] Plan 記錄 P1–P4、8 部矩陣與不提高全域頻率的決策。
- [ ] P1 正負 regression 在 adapter 通過並記 evidence。
- [ ] P2 缺 box／owned region／fusion 到 cue regression 通過並記 evidence。
- [ ] 八個真實來源對應矩陣，逐例 expected／actual／artifact／未驗證範圍完整。
- [ ] P4 timing／ASS／render 獨立驗收，證明無吞句。

本 refinement 不改 Phase 3A／3B blocking milestone，也不啟用 route 或 runtime projection。
