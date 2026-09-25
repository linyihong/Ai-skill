# Extracted — See [`workflow/apk-analysis/execution-flow.md`](../../../../workflow/apk-analysis/execution-flow.md) (Section 2: Quick Start, step 6) and [`intelligence/engineering/analytical-reasoning/heuristics/hook-selection.md`](../../../../intelligence/engineering/analytical-reasoning/heuristics/hook-selection.md)

### 2026-04-30 - 高語意 hook 優先於 socket bytes

Status: candidate

#### One-line Summary

能 hook request/response 物件，就不要先從 socket bytes 開始拼。

#### Human Explanation

socket、TLS read/write、`send`/`recv` 事件很多，容易卡 App，也需要自己重組 HTTP、解壓縮、切分 body。高語意 hook 例如 request options、response interceptor、decrypt function，通常事件少、內容接近業務語意，更適合建立 API 文件與測試 fixture。

#### Trigger

低層 socket / TLS hook 事件量大，容易造成 App 卡頓，也需要自行重組 HTTP。

#### Generalized Lesson

優先找 request options、response interceptor、decode/decrypt function。只有在高語意點找不到或需要補證據時，才降到 socket / TLS 層。

#### Agent Action

看到 socket hook 卡頓、ANR、輸出爆量時，停止擴大低層 hook，改回靜態搜尋 request builder / interceptor / decoder，或縮小 hook 條件。

#### Promotion Target

已整理到 `WORKFLOW.md`。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
