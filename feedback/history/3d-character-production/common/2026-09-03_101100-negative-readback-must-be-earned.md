> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-03 - 負向 readback 必須先證明可觀測

Status: validated

#### One-line Summary

記錄 `fail`／`partial` 前，必須先證明目標 artifact 已實際產出、且 readback 工具能在已知正例上讀到訊號；否則「未產出」與「探針失明」會被誤記成「驗證失敗」。

#### Human Explanation

Gate 讀到 0 或 no-change 時，至少有三種互斥成因：artifact 根本沒被產生、readback
工具讀不到、真的變形失敗。三者在 record 上長得一樣，但修復位置完全不同。若不先排除
前兩者，團隊會反覆在最貴的那層（重做資產）打轉，而真正的成因可能只是一個預設關閉的
開關，或一個不支援稀疏儲存格式的解析器。

#### Trigger

- Stage record 長期停在 `fail`／`partial`，但沒有任何欄位記錄 artifact 是否曾被產出。
- 產出流程含 opt-in 開關（環境變數、feature flag）且預設關閉。
- readback 工具自行解析容器格式，未在已知正例上校準。
- 同一結論（「動不了」）反覆出現，但每次歸因到不同層。

#### Evidence

- Tool: artifact 內容探針、DCC 端與匯出端的雙向比對。
- Sanitized excerpt: DCC 內部顯示形變資料存在，匯出檔經自製探針讀取卻全為零；探針補上稀疏儲存格式（如 glTF sparse accessor）支援後，同一檔案讀出非零位移。另一條路徑上，形變資料本身因 opt-in 旗標預設關閉而從未被產生。
- Evidence path: 具體探針輸出與逐階段量測留在 `<PROJECT_ROOT>` 的專案 artifacts。

#### Generalized Lesson

負向證據和正向證據一樣需要被賺取。記錄 stage 失敗前，先關掉兩個替代解釋：

1. **產出確認**：artifact 內確實存在該項目的非空資料；opt-in 開關的實際取值要被記錄，不能只看預設值。
2. **探針校準**：readback 工具先在已知正例上讀到預期訊號，才有資格宣稱 0。
3. 兩者都成立，`fail` 才是對資產本身的判斷。

未產出、未觀測與已觀測失敗，是三種不同狀態，不應共用同一個 `fail`。

#### Agent Action

在 record 或診斷紀錄中分開承載「artifact 是否產出」與「readback 結果」。遇到全零／
無變化時，先跑一個已知會動的對照項驗證探針，再回頭判定資產。自製解析器必須覆蓋容器
格式的壓縮與稀疏分支，否則其負向結果不得作為證據。

#### Goal / Action / Validation

- Goal: 避免把「沒做」或「沒看到」誤記成「做了但失敗」。
- Action: 在 acceptance record 增加產出確認欄位；readback 工具附帶已知正例的自我校準。
- Validation or reference source: 負向 scenario 需覆蓋「artifact 缺項」與「探針讀不到但資產正常」兩種情形，且兩者不得產生與真實失敗相同的 record。

#### Applies When

- 任何以匯出檔／二進位容器作為 stage 證據的驗收。
- 產出步驟含 opt-in 開關，或 readback 由自製工具完成。

#### Does Not Apply When

- 探針為上游權威實作且已有回歸覆蓋，且 artifact 產出無條件執行。
- 純探索性觀察，未寫入 record、不影響 eligibility。

#### Validation

- Record 能區分 `not_produced`、`not_observed` 與 `observed_fail`。
- 探針在已知正例上讀到非零，才允許以 0 作為 fail 證據。
- opt-in 產出旗標的實際取值被記錄在 evidence，而非假設。

#### Promotion Target

- `workflow/3d-character-production/`
- `validation/scenarios/3d-character-production/`

#### Required Linked Updates

- 同步 facial expression acceptance record 的產出確認欄位與 stage 文件。
- Reusable lesson 只保留 generalized rule；具體檔案、量測與工具路徑留在 consumer project。
