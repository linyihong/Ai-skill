# Extracted — See [`workflow/apk-analysis/execution-flow.md`](../../../../workflow/apk-analysis/execution-flow.md) (Section 1: Safety and Sanitization)

### 2026-04-30 - 登入限流要避免 tight-loop，優先 session reuse

Status: candidate

#### One-line Summary

遇到 login too frequently，不要盲目旋轉單一參數；先重用 session 並記錄風控維度。

#### Human Explanation

登入限流可能不是單一 request 欄位造成，而是伺服器用 device、User-Agent、IP、時間窗、App fingerprint、帳號狀態等多維度計算。一直換 device id 或 tight-loop login 可能讓問題更嚴重。穩定做法是同一輪測試重用 session，記錄每次登入嘗試時間與參數，必要時用 device/session pool。

#### Trigger

- API 回 login too frequently。
- 多個測試每次都重新登入。
- 改某個 device/body 欄位後結果不穩定。

#### Evidence

授權測試中曾做過參數 probe，無法把限流穩定歸因於單一欄位；session reuse 明顯降低重複登入風險。

#### Generalized Lesson

登入流程測試要有節流與重用策略。不要為每個測試方法重新登入；不要在沒有證據時假設旋轉單一 device 欄位即可繞過限制。

#### Agent Action

設計 live integration 或 runner 時，優先共用 session/context，記錄 login attempt metadata。遇到限流時先停止 tight-loop，再分析時間窗與風控維度。

#### Promotion Target

- `WORKFLOW.md`
- `DOCUMENTATION.md`

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
