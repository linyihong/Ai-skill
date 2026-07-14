> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-14 - Cold-start Play handoff vs vending focus-steal (revised)

Status: validated

#### Revision (same day, second)

1. 靜態 Pairip hop 過度歸因 → 已撤回。  
2. 「無 Google 帳號就無法進業務頁／分析假的」→ **也不成立**。同日實測：循環 `am force-stop com.android.vending` **只**抑止商店搶焦後，`Accounts: 0` 仍可維持業務 Home（含列表 UI）並進入 Reading Activity。標明：**這不是 Pairip／LVL crack**，是去掉 Play UI 搶焦；防守必須假設攻擊者會做。

#### One-line Summary

側載冷啟動：用 ≤100ms focus 證明「允許 Play → Unauthenticated 搶焦」；再用「抑止 vending」對照證明「無帳號也可站穩業務 UI」——兩者都要採樣，禁止只用靜態 Pairip 或「必須有帳號」當結論。

#### Human Explanation

常見現象是冷啟動後 focus 落到 Play Unauthenticated*（尤其裝置無 Google Account）。這常被誤讀成「App 完全進不去」或「一定要先登帳號才能分析」。細採樣可見業務 Home 可能先短暫出現；若循環 force-stop Play，Home 可維持並可導航到 Reading。靜態 Pairip／CHECK_LICENSE 仍只是候選。可重用結論是 **Play UI 搶焦 vs 業務 UI 可達** 的對照實驗設計，不是授權繞過菜譜。

#### Trigger

- 冷啟動進 Play 未登入頁。  
- 爭論「沒帳號就不能驗證／技巧是假的」。  
- 靜態 script exit 0 被當成已證實 Pairip hop。

#### Evidence

- Tool: `am start -W`；≤100ms `mCurrentFocus`；`screencap`；對照組循環 `am force-stop com.android.vending`。  
- Sanitized excerpt: （A）允許 Play：短暫 Home → Unauthenticated*。（B）抑止 Play：Home 維持＋目錄 UI；Reading Activity focus；Accounts=0。Pairip LicenseActivity 未採樣。  
- Evidence path: 目標專案 docs／capture／scripts。

#### Generalized Lesson

1. **三層證據**：L1 靜態假說／L2 focus／L3 feature（可操作）。  
2. **對照實驗**：Condition A 允許 Play；Condition B 抑止 vending 搶焦。分別記錄 focus＋截圖。  
3. **措辭**：B 成功時寫「抑止商店搶焦後業務 UI 可達」，**禁止**寫成「繞過 Pairip／CHECK_LICENSE」。  
4. **採樣**：≤100ms。  
5. **防守**：客戶端依賴「跳 Play 登入」不夠；內容／API 需後端 entitlement（Integrity 等）。  
6. **禁止**：APK 改包、授權 hook、教人永久破解 LVL。

#### Agent Action

1. 先做 A（允許 Play）再做 B（抑止搶焦），寫進專案 docs。  
2. 更新 skill lesson 時分開 A/B 證據。  
3. 使用者要 DRM bypass → 拒絕；可做 B 類觀測。

#### Goal / Action / Validation

- Goal: 可複核的冷啟動／搶焦／無帳號 UI 可達模型。  
- Action: A/B adb 對照 + 去敏 lesson。  
- Validation: Accounts=0 下 B 組 Reading／Home focus 可複現；A 組仍落到 Unauthenticated*。

#### Applies When

- 側載＋Play 未登入搶焦。  
- 需要無帳號驗證業務 UI 是否存在。

#### Does Not Apply When

- 需求是授權／簽名破解。  
- 已確認業務進程在 A 組就被殺掉（非僅搶焦）——另案分析。

#### Promotion Target

- `analysis/apk/workflows/cold-start-play-focus-ab.md`（HOW TO DO）— **done**  
- `intelligence/engineering/analytical-reasoning/heuristics/play-focus-steal-vs-hard-kill.md`（HOW TO THINK）— **done**  
- `analysis/apk/tools-and-failures.md` 失敗判讀列 — **done**  
- `workflow/apk-analysis/execution-flow.md` § Reset baseline 交叉連結 — **done**

#### Required Linked Updates

- 同上 Promotion；另更新 `analysis/apk/workflows/README.md`、`intelligence/.../heuristics/README.md`。

#### Confidence / Residual Risk

- Confidence: high on A/B 對照（搶焦 vs 抑止後 UI 可達）。  
- Residual: 抑止 Play 可能影響 Billing／後續授權 API；不代表正式 Play entitlement。
