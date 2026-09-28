# Observation — natural boundary before length balance

**Run ID**：2026-09-28-natural-boundary-before-length  
**Kind**：contract  
**Extends**：[`2026-09-28-break-candidate-schema.md`](2026-09-28-break-candidate-schema.md)

## 反例

「公司裁员名单就在领导电脑里。」被切成「电脑｜里」，因為切點最靠近理想字數。逗號與句號沒有先成為 clause boundary。

## Contract

Clause segmentation 先於 line break。標點是高優先候選，短句可合併。`prefer_at` 只在同一 boundary tier 內比較。`电脑里` 這類 lexical split 以 `hard_violation` 離開可行集。
