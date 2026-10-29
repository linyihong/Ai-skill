---
id: intelligence.analysis.runtime-evidence-boundaries
title: 局部執行證據的歸因邊界
schema_version: knowledge-atom/v1
type: intelligence
layer: intelligence
source_path: intelligence/engineering/analytical-reasoning/heuristics/runtime-evidence-boundaries.md
summary: 方法呼叫或 hook 命中只證明所在邊界；先找第一段未驗證契約，再決定要檢查參數、高層 driver 還是底層路徑。
domains: [apk-analysis, analytical-reasoning]
tags: [runtime-contract, evidence-boundary, diagnostic-decision]
status: candidate
priority: P2
confidence: low
stability: evolving
context_cost:
  estimated_tokens: 650
  load_strategy: lazy
  cacheable: true
when_to_read: 當局部成功或資料表示差異可能導致錯誤歸因時。
validation: 以來源條件與反例核對，未有獨立重用實測不提升狀態。
---
# 局部執行證據的歸因邊界

方法呼叫或 hook 命中只證明所在邊界；先找第一段未驗證契約，再決定要檢查參數、高層 driver 還是底層路徑。

## Decision Rule

| 觀測 | 優先判斷 | 不能推論 |
| --- | --- | --- |
| 呼叫結果為空，型別仍未核對 | 先確認簽名與參數表示 | 目標沒有資料 |
| enter 命中但後續事件缺失 | 查目標×階段的能力與分派 | 所有 owner 都必須做相同 UI 處理 |
| Animator 命中但視覺缺口仍在 | 比較 feature API 與當前證據支持的 driver | 一定是 Timeline／Native |
| 已有直接 Native／Timeline 證據 | 直接追該已證實分支 | 仍必須機械重跑每個高層候選 |

## Why

判斷以觀測覆蓋與消費端契約為基準，可以避免用一個局部成功旗標排除未測路徑，或把來源專案的修法當成普遍義務。具體取證程序留在 analysis；本檔只承接如何選路與限制宣稱。

## Boundaries

- 適用：同一目標、窗口與 fixture 可對齊，局部成功仍有未解釋結果。
- 不適用：只有無關目標或不同版本的命中數；先補觀測身分。
- 完整契約已證實時不重複猜參數；未解析 hook 不能用來排除路徑。

## 驗證與狀態

套用時寫出已觀測邊界、未覆蓋段落、選定判斷與可推翻它的證據。來源 lessons 與本 atom 均為 candidate；本輪抽取與合成反例不構成第二個專案實測。

## Sources

- [2026-10-16_173200-il2cpp-paytable-getmain-takes-system-char-not-string.md](../../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-16_173200-il2cpp-paytable-getmain-takes-system-char-not-string.md)
- [2026-10-23_140500-apply-shim-onleave-ui-must-cover-all-processresult-owners.md](../../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-23_140500-apply-shim-onleave-ui-must-cover-all-processresult-owners.md)
- [2026-10-26_100652-animator-play-visual-gap-use-feature-fx-api-not-timeline-first.md](../../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-26_100652-animator-play-visual-gap-use-feature-fx-api-not-timeline-first.md)

## Related method

[執行契約驗證](../../../../analysis/unity/workflows/il2cpp-runtime-contract-validation.md)。
