> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-16 - Regular diagram pages: mid-gutter rectangular panel grid

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本輪在密佈對照表頁以「綠線／切割線落在 gutter 中段、不碰灰格／彩格」驗證通過；細節留在專案 capture，不進本檔。

#### One-line Summary

當一頁圖示本來就是規則的 M×N 長方形格子時，用「內容帶偵測 + 切割線落在相鄰內容中間（gutter 正中）+ 外緣留呼吸空間」一次切成 M×N 面板；禁止為了單一標號把同一長方形拆成細條。

#### Human Explanation

密佈對照表若版面本身是齊整的長方形陣列（例如 5×4＝20 格），正確裁切單位是**格子**，不是每個數字一條細帶。細帶裁切會失去上／中／下列上下文，雙標號共用板也被拆壞。切割線應畫在兩側內容之間的空隙正中，並在外緣留 pad；成功條件很單純：切割線像素不要碰到灰色或其他顏色的內容方塊。這與「panel crop-first」不衝突——此時「面板」＝整格長方形；僅在格子間距明顯不均、無法定義規則陣列時，才退回標號錨點裁切。

#### Trigger

- 對照表／圖示頁肉眼是規則的列×行長方形陣列
- 部分格子內有多條標號共用同一圖板
- 依標號裁成細條後無法判斷上／中／下或鄰格內容被切進裁切框
- 需要先畫切割線／綠線核對，再產出面板檔

#### Evidence

- Tool: content-band detection (grey idle + neon stroke) → mid-gutter edges → overlay eye-check
- Sanitized excerpt: per-label thin strips split a shared top cell and lost row context; 依內容帶推導的規則矩陣保留共用標籤；分割線未穿過已核對的內容帶
- Evidence path: project paytable capture / panel-grid audit under `<PROJECT_ROOT>`（細節不進本檔）

#### Generalized Lesson

1. **先判斷版面**：規則 M×N 長方形 → 切格子；不規則／間距扭曲 → 標號錨點或人工框（見 dense panel-crop lesson）。
2. **切割線 = gutter 正中**：水平／垂直分割落在相鄰內容帶之間的中點；外緣另加呼吸 pad，不要貼齊內容外框「切太死」。
3. **成功判準**：切割線鄰域的灰格／彩格 ink fraction ≈ 0；用綠線／色線 overlay 做人審一次。
4. **一格一檔**：多標號共用板輸出同一個矩形；線號對應用格內 OCR，不要為每個 lineId 再切細條。
5. **Draft only**：格子裁切後的路徑抽取仍是 draft；正式映射要人審或獨立驗證。

#### Agent Action

- 先做：確認是否為規則陣列 → 偵測內容帶 → 中點 gutter 切線 + 外緣 pad → overlay 核對未切到方塊 → 再抽標號／路徑。
- 不要做：預設每個標號一條細帶；切割線貼齊內容外緣；用等分全頁（含標題／頁腳 chrome）當格子邊界。

#### Goal / Action / Validation

- Goal: 規則對照表頁的面板裁切可重現，且保留格內上／中／下上下文。
- Action: mid-gutter M×N 裁切 → overlay 核對 → 格內讀標號。
- Validation or reference source: divider ink fraction ≈ 0；overlay 人審；與 dense panel-crop lesson 的適用邊界一致。

#### Applies When

- 圖示／對照表為規則列×行長方形
- 需要保留同一格內多條線的相對上下位置

#### Does Not Apply When

- 間距明顯不均、拼貼、或無法定義穩定 M×N
- 單主體去背／修圖（非圖示網格）

#### Validation

同一頁產出切割 overlay：切割線不得穿過灰／彩內容方塊；抽一格含上中下三條水平線時，三條須同框可見。

#### Promotion Target

- `workflow/image-processing/execution-flow.md`（Dense multi-diagram 段補「規則陣列 → mid-gutter grid」）
- 既有 lesson `dense-diagram-panel-crop-before-ocr`（交叉連結適用邊界）

#### Promotion Record

尚未 promotion；保留為 candidate history。本輪已在 execution-flow 加短指引並連回本檔。

#### Required Linked Updates

- [x] `feedback/history/image-processing/common/README.md`
- [x] `workflow/image-processing/execution-flow.md` 短指引
- [x] 交叉連結 dense-diagram panel-crop lesson
- [ ] 獨立第二專案重跑後再標 `validated`

#### Abstraction Check

- Swap-project test: PASS — 不同文件的規則圖示矩陣均可按內容間隙裁切；格數、顏色與閾值依輸入推導。
- Project evidence kept in: `<PROJECT_ROOT>/` 的原圖、裁切與獨立核對紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

補齊抽象化檢查。文中的矩陣數量、顏色、標號與零交集示例均不是通用常數；抗鋸齒、雜訊與圖例侵入時，需依影像契約定容差或人工核對，不能只靠固定閾值自證通過。仍維持 candidate。
