> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-08-28 - Phone-side Play focus suppress via Tasker ADB import (trial + onboarding guards)

Status: validated

#### One-line Summary

冷啟動 Play 搶焦的 **手機端持久化** 可用 Tasker「App 開啟 → ADB Wifi 循環 `am force-stop com.android.vending`」；adb 可全自動 push/import XML，但 sideload 試用過期、onboarding 第 4 步與 Run Shell 權限是常見假綠。

#### Human Explanation

Condition B（Mac 端循環 force-stop）可證明業務 UI 可達，但使用者希望 **直接點桌面 icon**。Tasker 可用 App context 觸發 + ADB Wifi action 在裝置本機重複 force-stop。實務上 agent 常誤以為「已 import profile = 已生效」；sideload APK 試用結束、首次設定未完成、或誤用 Run Shell（無 root 無法 force-stop）都會讓 profile 看似存在卻不工作。另：永久 disable Play Store 可能觸發 Pairip 類「Check Google Play is enabled」硬 gate——須保持 Play 啟用，只 suppress UI 搶焦。

#### Trigger

- 已完成 Play focus-steal A/B（`cold-start-play-focus-ab.md`），需要 **不連 Mac** 也能點原 icon 進業務 UI。
- Agent 規劃 Tasker / MacroDroid 等 on-device automation。
- Tasker profile import 回報成功，但 `mCurrentFocus` 仍落在 `UnauthenticatedMainActivity`。
- Tasker 首次啟動卡在 onboarding 或跳出 Trial Over。

#### Evidence

- Tool: adb push + `ActivityImportTaskerDataFromXml`；uiautomator dump 定位 onboarding checkbox；`dumpsys window | grep mCurrentFocus`；Tasker UI「Trial Over」對話框。
- Sanitized excerpt: import intent 成功；App 開啟事件可觸發（孪生应用提示）；試用結束後 profile 不執行；第 4 個 onboarding 項導向 Settings 關閉 placeholder notification channel；ADB Wifi 設定已開但試用阻擋；Mac 端 Condition B 仍可進 MainActivity。
- Evidence path: `<PROJECT_ROOT>/<target-app>/docs/startup-play-gate.md`、`<PROJECT_ROOT>/<target-app>/scripts/adb_tasker_import_auto.sh`（project-local；勿複製進本 lesson）。

#### Generalized Lesson

1. **分層**：L3 Mac suppress（已驗）→ L4 on-device automation（Tasker 等）→ 每層獨立驗證 focus，不可跨層宣稱。
2. **Tasker import**：XML `.prf.xml` 可经 adb 推到 `/sdcard/Tasker/profiles/` 並用 import activity 匯入；App trigger 用 `<App sr="con0">` + `pkg0`/`cls0`，非過期 `<Event code=120>` Bundle 格式。
3. **Kill 命令**：無 root 時 `Run Shell`（code 123）通常 **不能** `am force-stop` 其他 package；用 **ADB Wifi**（code 375）+ 指令 `am force-stop com.android.vending`（不含 `adb shell` 前綴）。
4. **ADB Wifi 前置**：USB 連 Mac 時執行 `adb tcpip 5555`（重開機後需重做）；首次執行需裝置端接受 USB debugging；`WRITE_SECURE_SETTINGS` grant 不等於 ADB Wifi 已授權。
5. **Onboarding**：四個 checkbox 中，overlay/battery 可经 appops 預 grant；**供應商電池優化** checkbox 常需 scroll 後點擊，且會 **開啟 Settings** 要求關閉 Tasker placeholder notification——僅點 checkbox 右緣，勿點整列文字；全部勾選後 Proceed 才可用。
6. **Trial blocker**：apkeep / sideload 的 Tasker 常為試用版；「Trial Over」對話框出現時 **所有 profile 停用**，須 Play 購買或改用 Mac adb / Termux widget 等 fallback。
7. **禁止**：`pm disable-user com.android.vending` 作為 Pairip 包的首選方案（可能觸發 license hard gate）。

#### Agent Action

1. 先完成 `cold-start-play-focus-ab.md` Condition B，確認是搶焦而非 hard kill。
2. 若做 Tasker：腳本含 permissions + onboarding 自動化 + import + **trial 檢測**；profile 用 ADB Wifi 非 Run Shell。
3. Import 後驗證：**launcher tap** → `mCurrentFocus` 為目標 MainActivity，不是 Play Unauthenticated*。
4. 遇 Trial Over：在 lesson/專案 docs 標 blocker，改推 Mac `adb_launch_without_play_login.sh` 或已驗 widget 路線；勿假裝 profile 生效。
5. onboarding 自動化：解析 uiautomator CheckBox `checked="false"` bounds；Settings 步驟關 notification switch 後 BACK 回 Tasker。

#### Goal / Action / Validation

- Goal: 可複用的 on-device Play focus suppress 自動化 guardrails。
- Action: sanitized lesson + 更新 `cold-start-play-focus-ab.md` / `tools-and-failures.md` 交叉連結。
- Validation: 試用有效時 launcher tap focus=業務 Activity；試用過期時腳本輸出 WARN 且 Mac Condition B 仍可複現。

#### Applies When

- 側載分析機無 Google 帳號 + Play 搶焦 + 需要手機端 icon 啟動。
- 已授權使用 adb 設定分析裝置自動化。

#### Does Not Apply When

- 需求是 LVL / Pairip / 簽名 crack。
- 業務在 Condition A 即 process crash（非 Play 搶焦）——另案。
- 使用者環境禁止 Tasker 付費且不接受 Mac adb fallback。

#### Promotion Target

- `analysis/apk/workflows/cold-start-play-focus-ab.md` — phone-side persistence 小節 — **done (this commit)**
- `analysis/apk/tools-and-failures.md` — Tasker trial / onboarding 失敗列 — **done (this commit)**

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md` Recent 表 — **done (this commit)**

#### Confidence / Residual Risk

- Confidence: high on trial/onboarding/Run Shell vs ADB Wifi 判讀；medium on 各 OEM Tasker onboarding UI 偏移。
- Residual: ADB Wifi 重開機後需 Mac 再跑 `tcpip 5555`；Tasker 付費與政策變更可能影響 sideload 路線。
