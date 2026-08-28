> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-08-28 - Flutter Dio device-autologin as standalone SDK bootstrap (not refresh-harvest)

Status: validated

#### One-line Summary

短劇／Flutter+Dio App 常在冷啟動用 **device headers + autologin** 直接鑄 session；`pm clear` 後新 uid 是裝置身分重建證據，SDK 應還原 deviceId／fingerprint／autologin wire，而不是「手機登入後只偷 refreshToken」。

#### Human Explanation

分析目標若是「自己寫 SDK 登入」，常見錯誤是先在真機 Google/Email 登入，再把 refresh token 當唯一材料。若靜態出現 `/open/autologin`、`getApiHeaders`、`X-Device-*`，且 clear-data 後仍進入已登入業務頁並換新 uid，應優先假設 **server 依裝置身分 mint token**。Refresh 只是後續續期。

#### Trigger

- 使用者要求 standalone login／SDK，明確不要 phone-session→refresh。
- Flutter log：`开始执行自动登录`、`force_new`、AuthConfig 後立刻有 uid。
- `pm clear` 後 nickname/uid 改變但仍 `isGuest=false`（或類似非訪客態）。

#### Evidence

- Tool: `strings`/`unflutter` on `libapp.so`；adb logcat Flutter lines during clear+launch；Play focus suppress for UI reachability.
- Sanitized excerpt: header builder lists `X-Device-ID/Fingerprint/Brand/Model/Type/Physical` + `X-OS-*` + `X-App-*` + `User-Agent: <App>/…`；pref key `device_unique_id`；ANDROID_ID miss → `rd_` random；autologin mapping includes `accessToken`/`refreshToken`/`force_new`.
- Evidence path: `<PROJECT_ROOT>/<target-app>/docs/API/auth/`。

#### Generalized Lesson

1. **Bootstrap ladder**：L0 AuthConfig（OAuth 可選）→ L1 deviceId/fingerprint/headers → L2 autologin/device register → L3 optional bind → L4 refresh。
2. **SDK priority**：先做 L1+L2 replay；把 refresh-harvest 標為降級／除錯路徑。
3. **Clear-data 實驗**：clear 後新帳 = 裝置身分可重建的強信號；同 uid 則可能還有其他身分錨。
4. **unflutter 用法**：找 `*getApiHeaders*`、`*getOrGenerateDeviceId*`、`*performAutoLogin*` 與 path 字串 cluster，再補 Frida Dio dump。
5. **去敏**：文件只留 header 名、path、演算法輪廓；token／完整 client secret 不進 reusable lesson。

#### Agent Action

1. 不要先設計「只 refresh」SDK。
2. 建 `docs/API/auth/standalone-login-feasibility.md` checklist。
3. 動態只抓 headers/body **keys** + 去敏；再 offline replay mint。
4. 新技巧同輪 writeback 本 lesson 類條目。

#### Goal / Action / Validation

- Goal: 可重用的 Flutter 裝置自動登入分析順序。
- Action: static symbol+string cluster + clear-data behavior + auth docs。
- Validation: clear 後新 uid 可複現；header 名與 autologin path 可從 AOT 復現。

#### Applies When

- Flutter/Dio（或類似）+ device header interceptor + autologin/device register paths。
- Authorization covers APK analysis and SDK reconstruction for that target.

#### Does Not Apply When

- Session 只能由 SMS/OAuth 人機驗證發出且無 device mint。
- 需求是破解付費／DRM，而非授權分析。

#### Promotion Target

- `analysis/apk/workflows/` — optional later device-autologin HOWTO — deferred until wire replay exists
- `analysis/apk/tools-and-failures.md` — one failure row — **done this commit if updated**

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md` Recent — **done this commit**

#### Confidence / Residual Risk

- Confidence: high on analysis order and static header/deviceId outline; medium until Dio wire replayed.
- Residual: OEM ANDROID_ID policies; server may add signing later.
