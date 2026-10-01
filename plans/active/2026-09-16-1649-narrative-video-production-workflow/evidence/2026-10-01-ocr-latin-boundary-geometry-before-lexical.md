# Observation — Latin sticky tokens need geometry-before-lexical recovery

**Run ID**：2026-10-01-ocr-latin-boundary-geometry-before-lexical

**Kind**：Phase 3 dogfood → mechanical capability gap（adapter；非新 Phase／非 LLM 斷詞）

**Plan**：[`../42-ocr-boundary-and-script-aware-normalization.md`](../42-ocr-boundary-and-script-aware-normalization.md)

**Workflow**：[`../../../workflow/narrative-video-production/text-evidence-ocr-boundary.md`](../../../workflow/narrative-video-production/text-evidence-ocr-boundary.md)

**Corpus**：[`../../../workflow/narrative-video-production/records/latin-boundary-regression.yaml`](../../../workflow/narrative-video-production/records/latin-boundary-regression.yaml)

## 去敏情境

雙語硬字幕。Join 接縫修好後，仍見 **單 part** 黏串：`Justlethervolunteer`、`appointmentfortoday`、`commitmentremainsvalid` 等。`parts` 長度常為 1；僅有整句 bbox。

## 機械觀察

| 層 | 結果 |
| --- | --- |
| Join token-seam | 多 part Latin 已正常（非本檔回歸主軸） |
| Engine parts | 常 `parts=[整串]`，無 word boxes |
| 整句 bbox | 「有座標」但不足以指出詞界 |
| Closed-class 當第一刀 | 漏詞表／誤判單字；且跳過幾何 |
| 合法單字 | `Unexpectedly` 等不應 suspicious |

## Invariant

1. Recovery 順序：OCR boxes → **intra-box geometry** → lexical candidates（含 closed-class）→ ASR／context → resolver。
2. Closed-class = candidate generator，不是 truth generator。
3. Suspicious 閘門：非凡 Latin 都拆；dictionary exact-match → ok。
4. 案例進 regression corpus，禁止 raw 字串硬編碼特例。

## Regression（固定 raw）

見 `latin-boundary-regression.yaml`：`appointmentfortoday`、`tomasturbate`、`Workhard`、`aghostis`、`alsoromantic`、`Justlethervolunteer`、`commitmentremainsvalid`、`Itcan'tbesuchacoincidence`、`Ihaveahusband's` + negative singles。
