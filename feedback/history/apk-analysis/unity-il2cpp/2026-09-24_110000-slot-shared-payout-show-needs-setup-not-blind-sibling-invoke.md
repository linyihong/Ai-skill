> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-24 - Shared payout Show without settings can open the wrong sibling window

Status: candidate

#### One-line Summary

多個慶祝窗共用 `*PayoutWindowAnimatorController.Show`／`UISlots*PayoutWindow.Show` 時，**對全部實例盲呼叫 Show** 可能打開**另一個 sibling**（例如 Flash），不是目標 classic／BHS 窗。具名 GO `SetActive` + `Animator.Play(show clip)` 也可能「成功」但無完整 chrome——通常還缺 ViewModel／Settings `Setup`。Frida 端：`FindObjectsOfTypeAll` 高頻 poll 可弄死 script；Python `rpc.exports` 鍵名會被轉成全小寫。

#### Human Explanation

Idle dump 可見 `*_payout_window_classic` clip／GO 與共用 controller。Lab 強制開窗時常見失誤：

1. `FindObjectsOfTypeAll(AnimatorController)` → 對每個實例 `Show(null,null)` → 畫面變成 **Flash／其他 payout**，金額可能是 placeholder `0`。
2. 只對目標 GO `SetActive(true)` 或 `Play(classic_show hash)` → 回報 ok，但無標題／CTA／完整模態。
3. 正確方向：走該窗的 **Create + Setup(Settings／ViewModel) + Show**，或只對**綁定目標 skin 的那一個** controller／window 實例呼叫；強制錯窗後用對稱的 `Hide`／`SetActive(false)` 清場，勿狂點 Continue。

附帶 Frida 操作：

- 預設關閉 clip poll（`FindObjectsOfTypeAll` tick）；短獵才手動開。
- Host 呼叫 RPC 時用 **全小寫** export 名（`setpollenabled`），勿假設 snake／camel 保留。

#### Trigger

- Lab「開 classic」卻截到 Flash／BHS／TOTAL。
- GO Play／SetActive 日誌成功，截圖仍是桌面或錯窗。
- 開啟 Animator clip poll 後 Frida script destroyed／App 重啟。
- Python `exports_sync.set_poll_enabled` → `unable to find method`。

#### Evidence

- Tool: Frida IL2CPP FindObjectsOfTypeAll + Show／Hide／Play；host screencap.
- Sanitized pattern: blind Show → wrong sibling window (amount placeholder)；classic GO Play alone → no chrome；Hide clears forced Flash；poll tick unstable.
- Evidence path: `<PROJECT_ROOT>` cabinet `lab/classic-payout-lab-attempt-*.json`（project-local）。

#### Generalized Lesson

1. **共用 Show ≠ 目標窗**：先辨識實例所屬 skin／GO 名，再呼叫；禁止全實例盲 Show。
2. **開窗三件套**：Create／Setup(Settings)／Show；缺 Setup 常無 chrome。
3. **錯窗用 Hide 清**：對稱 API 優於 UI 連點。
4. **FindObjectsOfTypeAll 當 poll 危險**：預設關；短時手動開並準備 detach。
5. **Frida Python export 全小寫**：定義 RPC 時用 lowercase-only keys。

#### Agent Action

Lab 強制開窗前列出目標 GO／controller 實例名；只對匹配項 Setup+Show；截圖驗證窗種後再入庫；RPC 用 lowercase。

#### Goal / Action / Validation

- Goal: 目標慶祝窗可 lab 重現且不誤開 sibling。
- Action: 定位 skin → Create/Setup/Show（或單實例 Show）→ screencap → 錯窗則 Hide。
- Validation: 截圖標題／CTA 屬目標窗種；金額已遮或為明確 placeholder 並標 `replay-lab`。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP UI 多窗共用 AnimatorController／PayoutWindow 家族。
- Does not apply: 單一專用 window class、無 sibling Show 路徑的 cabinet。

#### Related

- `2026-09-24_085100-slot-ui-harvest-lab-fixture-apply-shim-not-live-payout-forge.md`
- `2026-09-21_111500-slot-celebration-window-hunt-continue-cta-not-midreel-gold.md`
- `2026-09-14_170500-frida-chain-harness-nonblocking-stdout.md`

#### Validation

- 螢幕截圖與實例名稱必須同時指向目標 window；若顯示 sibling 或無內容，記為 lab miss 而非成功。

#### Promotion Target

- apk-analysis slot UI lab：強制開窗須 Setup／實例鎖定；Frida poll／RPC naming。

#### Reuse Evidence

- 尚未有獨立 cabinet／專案的重用證據；保留為 candidate。

#### Promotion Record

- 尚未 promotion；僅保留為 feedback history candidate。

#### Required Linked Updates

- 已更新 `unity-il2cpp/README.md` category index；尚未更新 UI lab SOP，待重用驗證後再處理。
