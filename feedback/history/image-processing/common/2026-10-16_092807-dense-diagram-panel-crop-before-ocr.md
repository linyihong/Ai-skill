> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-16 - Dense multi-diagram pages: panel crop-first before OCR

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。

#### One-line Summary

對密佈多圖示頁，先裁成「單一標籤 + 單一圖板」面板再 OCR／路徑抽取；禁止用整頁 OCR 或等分網格結果直接改寫正式映射。

#### Human Explanation

整頁 Vision／OCR 會把相鄰圖示、共用雙標籤格、背景雜訊一起讀進同一判斷，常見假陽性是「中間橫線」與整排 lineId 錯位。等面積 N×M 切格看起來像「先切再讀」，但若圖示間距不均或一格畫兩條線，切格本身就錯，後面 OCR 再準也救不回來。正確順序是：契約先凍結「標籤與幾何以圖板為準」→ 內容對齊的面板裁切 → 每面板獨立讀標籤與路徑 → 另一條驗證路徑（人審或獨立腳本）對照後才寫入正式資料。這與 image-processing IP1／IP3 一致：完成＝符合契約，執行器不得自證 PASS。

#### Trigger

- 一頁上有多個標號小圖（圖示網格／對照表）
- 部分格子共用同一圖板、兩個標號
- 整頁 OCR 或等分網格後大量 line／cell 與肉眼不一致
- 想用一次全頁分析直接 remap 正式幾何表

#### Evidence

- Tool: panel crops + per-panel vision／OCR vs equal-area page tessellation + whole-page OCR
- Sanitized excerpt: 未對齊內容的等分裁切混入鄰格並漏切共用標籤；內容對齊的面板保留其標籤與幾何關係
- Evidence path: project capture notes under `<PROJECT_ROOT>` slot paytable docs（細節不進本檔）

#### Generalized Lesson

1. **Panel crop-first**：每個分析單元 = 一個標號（或明確的雙標號共用板）+ 其圖板；去掉頁面裝飾與相鄰圖示。
2. **等分網格只是 fallback**：僅在已證明圖示間距均勻、一格一標號時可用；否則先用標號錨點或人工框。若版面本身是規則 M×N 長方形陣列，改用 **mid-gutter 矩形格切**（切割線落在內容帶中間、外緣留 pad），見 [`regular-panel-grid-mid-gutter-cuts`](2026-10-16_132758-regular-panel-grid-mid-gutter-cuts.md)——不要為每個標號切細條。
3. **OCR 角色**：面板內讀標號／輔助對齊，不是整頁 remap 權威。
4. **Verify ≠ Execute**（IP3）：抽取腳本輸出是 draft；正式寫入前要獨立核對（人審面板裁切，或第二工具對同一面板）。
5. **雙標號共用板**：契約註明 shared panel；一次裁切、兩次路徑讀取，禁止拆成兩個等面積半格硬猜。

#### Agent Action

- 先做：凍結 Image Task Contract（意圖＝讀標號＋幾何；immutable＝不得改圖板像素當「修正」）→ 產出面板裁切 → 逐面板 OCR／路徑 → 獨立核對清單。
- 不要做：整頁 OCR 分數當 remap；等分 N×M 未驗證邊界就寫正式表；用執行模型自評「看起來對」當 accepted。

#### Goal / Action / Validation

- Goal: 密佈圖示頁的標號→幾何映射可重現且可獨立核對。
- Action: 面板裁切 → 每面板抽取 → draft JSON → 獨立核對後再寫正式表。
- Validation or reference source: 每面板人審或第二工具一致；draft 與正式表 diff 有對照紀錄；對齊 `workflow/image-processing` IP1／IP3。

#### Applies When

- 單頁多個標號圖示／對照表
- 標號 OCR 與路徑抽取可分離
- 需要把結果寫入正式配置（非僅預覽）

#### Does Not Apply When

- 單主體去背／修圖（走既有 A／B／C 契約，不是圖示網格）
- 已有機器可讀幾何權威且圖示僅作顏色／展示對照
- 圖示為不規則拼貼且無法定義面板邊界（先人工框或 abort）

#### Validation

同一頁用「面板裁切」與「等分網格／整頁 OCR」各跑一次：若後者改寫了前者已核對的映射，後者不得進入正式表。

#### Promotion Target

- `workflow/image-processing/execution-flow.md`（Input Validation／Analysis 補 panel-crop 指引）
- `workflow/image-processing/README.md`（何時啟用本 lesson）

#### Promotion Record

尚未 promotion；保留為 candidate history。本輪已在 execution-flow 加短指引並連回本檔。

#### Required Linked Updates

- [x] `feedback/history/image-processing/README.md` + `common/README.md`
- [x] `workflow/image-processing/execution-flow.md` 短指引
- [ ] 獨立第二專案重跑後再標 `validated`

#### Abstraction Check

- Swap-project test: PASS — 技術圖表、表單圖示或流程面板都可用內容對齊裁切與獨立驗證，不需沿用原頁標號或幾何。
- Project evidence kept in: `<PROJECT_ROOT>/` 的原圖、裁切與獨立核對紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

補齊抽象化檢查。文中的矩陣數量、顏色、標號與零交集示例均不是通用常數；抗鋸齒、雜訊與圖例侵入時，需依影像契約定容差或人工核對，不能只靠固定閾值自證通過。仍維持 candidate。

#### Transfer Contract

- Inputs: 密佈圖板、標籤與保留幾何的契約
- Mechanism: 全頁混讀或錯切會污染單元身分
- Decision: 依內容邊界裁切後分板抽取，再獨立核對
- Verification: 裁切單元與正式映射的對照
- Counterexample: 已有機器可讀幾何權威時不以 OCR 取代
- Transfer case: 技術圖紙中的多標籤面板

#### Intelligence Review

- Decision: method-only
- Reason: 本條主要是圖板裁切與 OCR 取證順序，已在 image workflow，不強制建立重複智慧。
- Target: workflow/image-processing/execution-flow.md
