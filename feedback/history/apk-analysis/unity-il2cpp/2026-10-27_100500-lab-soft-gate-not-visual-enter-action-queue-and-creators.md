> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-27 - Lab soft gate／模式旗標 ≠ 視覺進場；用 Action 佇列 + Creator 建內容

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源：bonus／board 類 lab 只寫 soft mode 或 Animator `SetBonusMode` 後，軌道路徑／符號仍缺；改走非 null 回呼的 Present／Show Action，並用 Creator `CreateContent+Show` 後截圖才見內容。

#### One-line Summary

寫 soft mode／bonus 旗標或只呼叫 Animator 式 `SetBonusMode` **不代表**畫面已進 bonus／board；要另走 **GameAction／Present*Action.Start**（避免同步 Present 傳 null Finished 卡死），並用 **Creator CreateContent+Show**（單靠 SetupItemSymbol／設格子 ID 常無視覺）。fixture 屬性別名對不上時驅動會靜默跳過。

#### Human Explanation

Unity 機台常把「邏輯已在 bonus」與「chrome／reels／棋子已上屏」拆開。Lab 若只翻布林或切 Animator bool，截圖仍像主場。同步 `Present*(…, Finished=null)` 可能讓主執行緒／Frida invoke 掛死，應改建構對應 Action 後 `.Start`。符號管線上，`SetupItemSymbol` 或改 model 長度常只改內部狀態；真要看到骰／道具，需走 board／field Creator 的 `CreateContent`＋`Show`，再 `SetContent`／Drop。`SupportCell==false` 不表示 Create 一定失敗。另外 betting／fixture 用的 wire 鍵名（如 `FROM_*`／`TO_*`）常與 IL property（`Start*`／`Finish*`）不同——parser 只認一邊會讓 Move／Show 驅動靜默 no-op。畫面上看起來像「棋子」的符號，也可能只是格內 high 符號，不是軌道上的 piece GO。

#### Trigger

- soft mode／SetBonusMode 回傳 ok，但截圖無 bonus chrome／符號／棋子
- Frida 同步 Present／Show 帶 null callback 後腳本無回應
- GetReelLength／GetSymbol 有值或 SetupItemSymbol ok，但 GetContent 空、畫面無 glyph
- fixture 有 MAKE_MOVE／移動類 reward，但 Move*Action 從未建構（attrs 讀不到）
- 截圖「有頭像／棋子樣」但軌道 piece 的 m_Shown／世界座標未變

#### Evidence

- Tool: Frida apply-shim；Creator CreateContent／Show；Present*Action.Start；screencap
- Sanitized excerpt: soft+animator ok → 無符號；length restore＋CreateContent+Show → 多數格可見；wire `FROM_/TO_` 對上後 MoveAction 才 armed
- Evidence path: `<PROJECT_ROOT>/docs/slots/*/lab/` enter-board／enhance apply notes

#### Generalized Lesson

1. **Gate ≠ paint**：soft／mode 旗標與 Animator bonus bool 是前置，不是 presence 閉環。
2. **Present／Show**：優先佇列 Action `.Start`；避免同步 API + null Finished（易 hang）。
3. **符號**：Setup／寫 ID／改 reel length 之後仍空 → 找 `*Creator.CreateContent`＋`Show`（必要時 SetContent／DropItems）；勿把 SupportCell false 當硬失敗。
4. **長度**：公開 `SetMax` 類 API 可能腐化 length；優先還原已知安全欄位（min／max／origin）再 `Setup`，並以 GetReelLength 驗證。
5. **Fixture attrs**：同時接受 wire 鍵與 IL／DTO 別名；驅動前 log 解析到的 from／to（或等價），缺值則標 skipped。
6. **身分**：格內符號 ≠ 軌道／路徑 piece；用物件指標／shown 旗標／世界座標區分，不要靠截圖像不像。

#### Agent Action

- 進場 checklist：mode → switcher／bg／fg → Present*Action → labels → creators → piece Move／Show；每步截圖或 dump。
- 寫 fixture parser 時列出 attrs 鍵別名表，並做「無移動鍵 → 明確 skipped」回報。
- 宣稱「符號／棋子已上屏」前必須 screencap 或等價 paint oracle，不能只靠 invoke ok。

#### Goal / Action / Validation

- Goal: lab 視覺進場可重現，且不把 soft gate 誤報完成
- Action: Action 佇列＋Creator；attrs 別名；paint 驗證
- Validation or reference source: 截圖有目標 chrome／符號；Move 驅動有 from／to；無 Present hang

#### Applies When

- Unity／IL2CPP lab 用 Frida 強進 bonus／board／雙 reels 視覺
- fixture／betting-record 驅動 Move／Show／Create

#### Does Not Apply When

- 只驗證 Rewards／CELL 解析、不要求畫面
- 非 Unity 或無可 Creator／GameAction 的管線

#### Validation

- 同 fixture：只 soft → 無符號；加 Creator → 多數格可見
- Present(null) 路徑 vs Action.Start：後者可返回
- 錯 attrs 鍵 → move skipped；別名齊 → ctor／start true

#### Promotion Target

- analysis/unity/workflows/unity-fx-driver-trace.md

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- category README 追加本條
- `analysis/unity/workflows/unity-fx-driver-trace.md` §2.1 補 soft-gate／Creator／attrs 別名要點與本條連結
- 專案 evidence 留 `<PROJECT_ROOT>/docs/slots/*/lab/`；本 lesson 不含柜机／sample 專名

#### Transfer Contract

- Inputs: soft／bonus 旗標、Present／Show API、Creator、fixture attrs、screencap
- Mechanism: 邏輯模式與視覺／內容工廠分離；null 回呼阻塞；wire 鍵≠IL 名；glyph 像≠piece
- Decision: 旗標後必須 Action＋Creator；attrs 雙別名；paint 才算進場
- Verification: 截圖／GetContent／move from-to 日誌可推翻「已進場」宣稱
- Counterexample: 純 Rewards dump、無 UI 目標
- Transfer case: 任何「mode ok 但畫面空」的 Unity feature lab

#### Intelligence Review

- Decision: method-only
- Reason: 操作邊界與既有 FX／presence lab 流程增量，併入 analysis/unity/workflows/unity-fx-driver-trace.md
- Target: analysis/unity/workflows/unity-fx-driver-trace.md
- Revisit: none
- Discovery: not applicable — method-only

#### Abstraction Check

- Swap-project test: PASS — 任一 Unity lab「旗標／Animator ok 但畫面無內容」均可套用 Action＋Creator＋attrs 別名＋paint 驗證
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/*/lab/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
