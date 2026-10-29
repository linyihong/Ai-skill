---
id: intelligence.analysis.representation-contract-parity
title: 展示模型與執行模型的契約對齊
schema_version: knowledge-atom/v1
type: intelligence
layer: intelligence
source_path: intelligence/engineering/analytical-reasoning/heuristics/representation-contract-parity.md
summary: 展示子集與完整執行結構不必相同；轉換是否成立取決於座標、容量、缺值與 sentinel 契約，而不是畫面是否填滿。
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
# 展示模型與執行模型的契約對齊

展示子集與完整執行結構不必相同；轉換是否成立取決於座標、容量、缺值與 sentinel 契約，而不是畫面是否填滿。

## Decision Rule

| 條件 | 判斷 |
| --- | --- |
| 消費端只需展示窗口 | 保留原窗口，不擴大 |
| 完整結構必填，且缺值有合法來源 | 依契約轉換並驗證所有 adapter |
| 完整結構必填，但缺值未知 | 拒絕或 blocked，不猜值 |
| sparse／跨格佔位合法 | 保留語意，不以填滿作成功標準 |
| 後續事件可能覆寫 | 先證實 writer 與排序；固定延遲不是通用解 |

## Why

判斷以觀測覆蓋與消費端契約為基準，可以避免用一個局部成功旗標排除未測路徑，或把來源專案的修法當成普遍義務。具體取證程序留在 analysis；本檔只承接如何選路與限制宣稱。

## Boundaries

- 適用：fixture、wire adapter 與記憶體 apply 使用不同表示；也適用虛擬列表或影像緩衝窗口。
- 不適用：契約不明且沒有正常對照；此時只列未知。
- padding、擴容、重寫只是契約允許時的策略，不是普遍必要步驟。

## 驗證與狀態

套用時寫出已觀測邊界、未覆蓋段落、選定判斷與可推翻它的證據。來源 lessons 與本 atom 均為 candidate；本輪抽取與合成反例不構成第二個專案實測。

## Sources

- [2026-10-23_134500-wire-result-glass-must-match-processresult-strip.md](../../../../feedback/history/apk-analysis/local-proxy/2026-10-23_134500-wire-result-glass-must-match-processresult-strip.md)
- [2026-10-23_141500-apply-shim-visiblegrid-must-pad-processresult-strip.md](../../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-23_141500-apply-shim-visiblegrid-must-pad-processresult-strip.md)

## Related method

[執行契約驗證](../../../../analysis/unity/workflows/il2cpp-runtime-contract-validation.md)。
