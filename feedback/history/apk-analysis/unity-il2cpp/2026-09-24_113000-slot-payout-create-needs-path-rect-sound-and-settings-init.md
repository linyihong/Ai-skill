> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-24 - Payout window Create needs Path+Rect+SoundStore; Settings needs object init

Status: candidate

#### One-line Summary

`UISlots*PayoutWindow.Create` 工廠常是 `(_Group, _ID, _Rect, _SoundStore, _Path, _ApplyMenuBarOffset)`：`_Path` 才是 prefab／skin 資源名，`_ID` 是層級唯一 id。Rect 與 SoundStore 應從**已存活 donor** 複製；空 Path／錯 ABI 會回 `null` 或 assert。`SlotsPayoutWindowSettings` 用 `il2cpp_object_new` 後必須 `il2cpp_runtime_object_init` 再 `Setup`。另：GO 上有 `*_classic_controller` **RuntimeAnimatorController 資產名**，不代表場景裡存在同名的 `*PayoutWindowAnimatorController` **元件**——對 Flash／FS 元件盲 `Show` 仍會開錯 sibling。

#### Human Explanation

Lab 強制開 classic／skin 專窗時，常見卡點：

1. **把 skin 字串塞錯參數**：`_ID`≠`_Path`。正確是唯一 `_ID` + `_Path=<prefab name>`（例如 path 固定為 classic prefab 名，id 用 `lab_<ts>`）。
2. **Rect／SoundStore 瞎填**：從同層已掛 parent 的 payout／video window 抄 `m_FinalRect` 與 `get_SoundStore`；缺 SoundStore 時 Create 常失敗或後續 Show 無內容。
3. **Settings 未 init**：`object_new`  alone → `Setup` breakpoint／exception；先 `runtime_object_init` 再寫 Payout／Multiplier／Title。
4. **資產名 ≠ 元件實例**：`get_runtimeAnimatorController` 可讀到 `*_classic_controller`，但 `FindObjectsOfTypeAll(*PayoutWindowAnimatorController)` 可能只列 Flash／FS。對後者 `Show` 會開 Flash placeholder，不是 classic chrome。
5. **Create+Setup+Show 成功仍可能無可見 chrome**：`m_IsHidden=0` 只表示節點未標 hidden；無自然派彩／ViewModel 綁定時截圖仍是桌面——標 `lab miss`，勿當 live proof。

#### Trigger

- Create 回傳 null／breakpoint，但 idle dump 已有 classic GO／clip。
- Setup 一呼叫就 assert，Show 卻「成功」。
- Animator controller **資產名**含 classic，controller **元件列表**卻只有 Flash／FS。
- Create/Setup/Show 日誌全綠，截圖無模態。

#### Evidence

- Tool: Frida IL2CPP `Create`／`Setup`／`Show` + `FindObjectsOfTypeAll` + screencap.
- Sanitized pattern: path+donor Rect/Sound → Create ok；Settings init → Setup ok；no classic AnimatorController *component* in scene; Show may leave hidden=0 without chrome.
- Evidence path: `<PROJECT_ROOT>` cabinet `lab/classic-payout-lab-attempt-*.json`（project-local）。

#### Generalized Lesson

1. **Dump Create 參數名**：先確認 `_ID` vs `_Path`，再掃 skin 字串。
2. **Donor 複製 Rect／SoundStore／Parent**：不要硬編螢幕像素當 Rect。
3. **Settings：`object_new` + `runtime_object_init`** 再填欄位。
4. **分清 RuntimeAnimatorController 資產 vs `*AnimatorController` 元件**：元件列表鎖定目標名；無元件則勿盲 Show sibling。
5. **日誌成功 ≠ 可見 chrome**：截圖分類；無模態就記 boundary，改自然獵或更深 ViewModel 路徑。

#### Agent Action

Create 前 dump overload；用 donor Rect/Sound；Settings init；只對匹配元件 Show；截圖驗證後再入庫。

#### Goal / Action / Validation

- Goal: Lab 能穩定 Create 目標 skin 實例，且不誤開 sibling。
- Action: Create(path+donor) → init Settings → Setup → Show → screencap；錯窗 Hide。
- Validation: Create 回傳 alive；截圖為目標窗種或明確標 `no chrome / lab miss`。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP 多 skin 共用 PayoutWindow 工廠＋Settings Setup。
- Does not apply: 單一 prefab、無 `_Path` 工廠、非 IL2CPP UI。

#### Related

- `2026-09-24_110000-slot-shared-payout-show-needs-setup-not-blind-sibling-invoke.md`
- `2026-09-24_085100-slot-ui-harvest-lab-fixture-apply-shim-not-live-payout-forge.md`

#### Validation

- 確認 Create 回傳存活實例、Settings 初始化後 Setup 無 exception，並以截圖區分可見目標 chrome 與 no-chrome lab miss。

#### Promotion Target

- apk-analysis slot UI lab：Create ABI／Settings init／controller 資產 vs 元件。

#### Reuse Evidence

- 尚未有獨立 cabinet／專案的重用證據；保留為 candidate，避免把單一 factory ABI 當作通用介面。

#### Promotion Record

- 尚未 promotion；僅保留為 feedback history candidate。

#### Required Linked Updates

- 已更新 `unity-il2cpp/README.md` category index；尚未更新 UI lab SOP，待相同類型 factory 的重用驗證後再評估。
