> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-27 - Soft 視覺還原：先資料與進場圖，再停每幀／Tween，最後才 paint

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源：授權 Unity 機台 lab 中 soft 進場／填格 API（SetContent／MoveContent／CreateContent）invoke 成功，但 GetContent 恒 null、或 Animator／Transform 歸位後被下一幀拉回；對照「資料模型／進場 Action 未完成」與「Update／Tween／Timeline 仍在刷」兩類機制後收斂還原順序。

#### One-line Summary

機台前端常是**對齊記憶體模型的播放器**；只呼叫 `Play`／`Reset`／`SetContent` 不夠。Lab 應先對齊 SpinInfo／State／mode／進場 Action 圖，必要時短窗口停 Update 或 Kill 目標 Tween／Director，**最後**才 paint，並以 getter 讀回（如 GetContent）而非 invoke true 當閉環。

#### Human Explanation

伺服器（或 lab 注入的等價 payload）決定停輪與獎圖時，前端腳本會持續檢查「模型是否仍在 spinning／settle」。只改畫面會被下一幀強制拉回。另有三類前端套路：程式在 Update 裡硬寫旋轉／位移；Coroutine／DOTween 類 handle 在背景繼續跑；PlayableDirector／payload 時間軸仍在 Play。這與「soft gate 已開但 Creator 未跑」可同時存在：即使 Create／SetContent／Move 都回 true，內容圖未掛或每幀覆寫時，GetContent 仍可為 null、截圖仍錯。正確順序是剝洋蔥：資料與自然／半自然進場 →（可選）停驅動 → 清 tween／Director → paint → 讀回驗證。

#### Trigger

- SetContent／MoveContent／Show invoke true，GetContent（或等價）仍 null
- Animator.Play／SetLayout／SetActive 後畫面瞬間彈回
- Soft mode ok 且 layer GetLayer 已解析，填格仍不 stick
- 懷疑 DOTween／Coroutine／PlayableDirector，但尚未證明

#### Evidence

- Tool: Frida／IL2CPP lab apply；assembly 枚舉；Update hit 計數；getter 讀回
- Sanitized excerpt: layer parent Move ok + SetContent true → GetContent 0；對照需 Present／Start 完成或停每幀覆寫後再探
- Evidence path: `<PROJECT_ROOT>/docs/slots/*/lab/` 與 interface-analysis 文字筆記（失敗截圖是否進 viewer 由專案政策）

#### Generalized Lesson

1. **Data is king：** 目標畫面所需的 SpinInfo／State／Rewards／mode／進場 Action 未就緒前，不要宣稱 soft paint 完成。
2. **Invoke ≠ map：** SetContent／MoveContent true 只表示呼叫成功；以 GetContent（或契約 getter）+ 授權 screencap 當 paint oracle。
3. **每幀覆寫：** 找 reel／board controller 的 Update／LateUpdate；lab 可短窗口 early-return，必須可解除。
4. **Tween／Coroutine：** 先證實組件／命名空間存在，再對**目標範圍** Kill；避免 blind KillAll 打壞共用 HUD。
5. **Timeline／payloads：** Director 仍在 Play 時改單元件無效；依 unity-fx-driver-trace 處置。
6. **順序鎖定：** 資料／進場圖 → 停驅動 → Kill tween／Director → paint API；不要顛倒。
7. **分支記錄：** data／Update／Tween／Timeline／ABI 分開標 observed／ruled-out／not-tested。

#### Agent Action

- Soft 半成功時載入 [unity-soft-visual-restore.md](../../../../analysis/unity/workflows/unity-soft-visual-restore.md)，勿只換 layer 字串或加大 force 表。
- 探測清單：assembly 有無 tween 庫；Update hit；Director isPlaying；mode／spinning 旗標；SetContent 後立即 GetContent。
- 專案 viewer：未還原成功前不要把失敗 soft 截圖當 matched 證據（若專案有此政策）。

#### Goal / Action / Validation

- Goal: soft 視覺還原路徑可解釋、可驗證，避免「API ok＝已還原」
- Action: 按剝洋蔥順序執行並記錄分支
- Validation or reference source: getter 讀回非空或明確 data／driver 缺口；換專案不依賴原柜机 API 名

#### Applies When

- 授權 Unity／IL2CPP 機台 lab／presence 視覺還原
- Soft 進場或填格半成功（API ok、讀回／畫面失敗）

#### Does Not Apply When

- 僅 wire／Rewards 解析、無 UI 目標
- 未授權對正式對局做破壞性 KillAll／改結果／改餘額
- 非 Unity 或無每幀／Tween／Director 證據且契約 ABI 已解釋失敗

#### Validation

- 同 fixture：只 paint → GetContent null 或彈回；補進場資料／停驅動後讀回改善或缺口可歸類
- 無 tween 組件時不呼叫 KillAll
- 分支表至少標一項 observed 與一項 not-tested

#### Promotion Target

- [analysis/unity/workflows/unity-soft-visual-restore.md](../../../../analysis/unity/workflows/unity-soft-visual-restore.md)
- 交叉：[unity-fx-driver-trace.md](../../../../analysis/unity/workflows/unity-fx-driver-trace.md)

#### Promotion Record

2026-10-27：新增 workflow 並掛本條為候選來源；不宣告 validated／promoted。

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md` Recent
- `feedback/history/apk-analysis/unity-il2cpp/` 本條
- `analysis/apk/README.md` Unity 入口
- `analysis/unity/workflows/unity-fx-driver-trace.md` 交叉連結
- 專案 evidence 留 `<PROJECT_ROOT>/docs/slots/*/lab/`；本 lesson 不含柜机／sample 專名

#### Abstraction Check

- Swap-project test: PASS — 任何授權 Unity 機台 lab 在 soft paint API ok 但讀回／畫面失敗時，都可套用「資料→停驅動→paint」順序。
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/*/lab/` 與專案 interface-analysis 政策。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### Transfer Contract

- Inputs: soft 半成功觀測、狀態旗標、Update／Tween／Director 探測、getter 讀回
- Mechanism: 播放器對齊模型；每幀／Tween／Timeline 覆寫；invoke≠content map
- Decision: 先資料與進場圖，再停驅動，最後 paint
- Verification: GetContent（或等價）或可歸類缺口；非僅 invoke true
- Counterexample: 純協議分析、未授權破壞性操作
- Transfer case: bonus／board soft 進場填格後讀回仍空

#### Intelligence Review

- Decision: method-only
- Reason: 本條主要沉澱 soft 還原操作順序 procedure；局部證據歸因仍用既有 heuristics。
- Target: analysis/unity/workflows/unity-soft-visual-restore.md
