> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-15 - Cold-start guest token may come from dedicated bootstrap endpoint, not login

Status: candidate

#### One-line Summary

`pm clear` 冷啟後，遊客 session 可能來自 **專用 bootstrap 端點**（`authLen=0`），而 **`user/login` 可能 0 次**；離線 SDK 勿假設必須走 login POST。

#### Human Explanation

常見誤判：任何 token 都必須經 `user/login`。實務上冷啟序可能是 `app/bootstrap`（或同族）在 **空 Authorization** 下回傳嵌套 `user { token, uid }`，之後才有 `user/switch`、`device/update` 等。`user/login` 可能只服務顯式帳號（郵箱／OAuth），不在遊客 SDK 主線。與 `143400`（sentinel header on login）互補：本條是 **端點選擇**，不是 header sentinel。

#### Trigger

- Cold-start capture after `pm clear`：bootstrap path 有 `authLen=0`，login path 未出現
- jadx：`startBoot`／`BootStrpModel` 含嵌套 `UserInfo`
- SDK 卡在 chicken-egg「沒 token 不能 login」

#### Evidence

- Tool: Frida cold-start path hook + jadx `RequestService`
- Sanitized excerpt: bootstrap → `data.user.token` + `uid`; login count = 0 in same window
- Evidence path: `<PROJECT_ROOT>/api/dynamic-w22a-*.md`

#### Generalized Lesson

```text
Guest session RE order:
  1. pm clear + spawn; log ALL hwycclient paths in first 30s
  2. If login absent but bootstrap present with authLen=0 → token likely from bootstrap
  3. Parse nested user object in bootstrap response (not only top-level token field)
  4. Keep user/login for explicit account flows; do not force it into guest SDK
  5. Cross-check SpData / HttpGlobal Authorization setter after bootstrap
```

#### Agent Action

1. 專案 doc 記錄冷啟 API 序與 response JSON path。
2. 交叉引用 `143400`、`143500`（session type vs entitlement）。
3. Ai-skill 不寫 package／host／token 樣本。

#### Goal / Action / Validation

- Goal: 鎖定遊客 token 主路徑，避免錯打 login。
- Action: cold-start path matrix + bootstrap response parse。
- Validation: device E2E 用 bootstrap token 打通下游 task API。

#### Applies When

- Rewards／任務 SDK 目標為 guest／TEMP 帳號
- Cold-start 觀測到 bootstrap 族端點

#### Does Not Apply When

- OAuth-only 無 bootstrap 族
- Login 在冷啟序明確出現且 bootstrap 無 user 嵌套

#### Promotion Target

- `analysis/apk/workflows/http-api-documentation-flow.md` §session bootstrap
- `analysis/apk/traffic-triage.md` §Session / Token

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
- `analysis/apk/traffic-triage.md`
