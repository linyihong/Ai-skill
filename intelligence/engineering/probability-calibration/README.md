# Probability Calibration Intelligence

機率模型校準的**候選判斷知識**。流程與契約在 [`workflow/probability-calibration/`](../../../workflow/probability-calibration/README.md)；本目錄只放「為什麼這樣判斷」，不重新定義契約欄位或 status 規則。

> **狀態**：全部為 `candidate-intelligence`，證據 N=1（單一 slot consumer）。**不是**通用硬規則；每個 atom 都列出「單一案例觀察」與升級條件。升級為 `validated-intelligence` 需 ≥3 個獨立收斂案例（至少一例非 slot）。

| Atom | 主題 | 來源 |
| --- | --- | --- |
| [`structural-ceiling-before-search.md`](structural-ceiling-before-search.md) | 先證明結構天花板，再擴大搜尋 | Inventory I6；C7 |
| [`single-knob-side-effects.md`](single-knob-side-effects.md) | 調一旋鈕、量全部指標；必要時拆正交旋鈕 | Inventory I7；C8 |
| [`sampler-strategy-as-data.md`](sampler-strategy-as-data.md) | 抽樣變異點資料化 + 註冊式 Strategy + 零漂移護欄 | Inventory I8；C9 |
| [`optimization-probe-hygiene.md`](optimization-probe-hygiene.md) | 解析掃描優先、串流、seed sweep、退化解偵測 | Inventory I9；verification roles |

## 邊界

- 不建 runtime route、signal 或自動選擇機制
- 不提供優化器或機率引擎實作
- 估計方法（有狀態機制、分布比較）、保底機制是否屬結構天花板：Phase 3 觀察項，**不**在此下結論
- 與既有 anti-pattern 重疊時只連結：[`shotgun-debugging`](../anti-patterns/shotgun-debugging.md)、[`validation-proxy-trap`](../anti-patterns/validation-proxy-trap.md)
