# Extracted — See [`workflow/apk-analysis/execution-flow.md`](../../../../workflow/apk-analysis/execution-flow.md) (Section 1: 開始前確認, session recovery)

### 2026-04-30 - Session refresh 要還原 App 的真實登入/裝置流程

Status: candidate

#### One-line Summary

Token 過期不一定有 refresh-token；要看 App 實際怎麼重新取得 session。

#### Human Explanation

有些 App 收到 invalid token 後，不會走標準 refresh-token API，而是清掉舊 token，回到啟動或裝置登入流程重新拿 session。若只拿舊 token 重算簽章，簽章可能正確但 token 仍無效。分析 session 問題時，要同時看 response interceptor、token store、device identity、login request builder 與 signing path。

#### Trigger

- API 回 no token、token expired、invalid token。
- 重新簽 request 仍失敗。
- App 重啟後又能成功。

#### Evidence

授權分析中曾確認：舊 session 失效後，需要按 App 的裝置登入流程取得新 token，而不是只重用舊 token 或單純重算 request signature。

#### Generalized Lesson

Session refresh 要從 App 內的 token invalidation、device identity、login body、request signing 與 token storage 一起還原。不要假設一定有 OAuth-style refresh token。

#### Agent Action

遇到 token/session 問題時，要求檢查：response interceptor 對錯誤碼的處理、token 存放位置、device id 來源、login endpoint/body、簽章 canonical path、成功後 token 寫回位置。

#### Promotion Target

- `WORKFLOW.md`
- `TOOLS.md`

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
