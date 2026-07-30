> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-30 - 缺 document converter 時的 .docx 文字抽取備援

Status: candidate

#### One-line Summary

`.docx` 是 ZIP + XML，所以在沒有 `pandoc` / office suite / Python 的環境裡，仍可用平台內建解壓 + 任一可用 runtime 做標籤剝除來取得可審閱的全文。

#### Human Explanation

文件處理指引通常預設 `pandoc` 或 office suite 可用。實際 agent 環境常常三個都不在：
converter 未安裝、office suite 未安裝、而系統內建的 Python 是 app store 佔位程式
（`--version` 直接以非零碼結束）。此時容易誤判為「無法讀取此檔」而卡住，或退回請使用者
自行轉檔。

關鍵事實是 `.docx`（以及 `.xlsx` / `.pptx`）是 **Open Packaging Convention 壓縮檔**：
主文內容在 `word/document.xml`。只要環境有任一種解壓能力與任一種可執行 runtime，
就能取得全文。**先探測環境能力、再選路徑**，比假設工具鏈存在可靠。

抽取時的品質關鍵在於**先標記結構再剝標籤**：直接把所有標籤替換成空字串會讓段落、
表格與儲存格全部黏成一行，長文件會失去可讀性與條號對應關係。正確順序是先把
段落／列／儲存格／換行／定位標記轉成佔位符，再移除其餘標籤，最後還原 XML 實體。

#### Trigger

需要讀取 `.docx`（或其他 OPC 格式）內容進行審閱、比對或摘要，但 converter 與 office suite
都不可用。

#### Evidence

- Tool: 平台內建壓縮工具（解壓）+ 環境既有 JavaScript runtime（文字處理）
- Sanitized excerpt: 一份約 1.7 MB 的合約文件（含多張內嵌圖片），`word/document.xml`
  約 190 KB；結構標記後剝除標籤，產出約 10 KB、340 行的可審閱純文字，段落與表格邊界保留。
- Evidence path: 抽取腳本置於工作階段 scratchpad，非專案內容；不進本庫也不進業務專案。

#### Generalized Lesson

處理專有格式文件前，**先探測環境能力，不要假設工具鏈存在**：

1. 探測：converter → office suite → 腳本 runtime，逐一確認可執行（不只確認可解析路徑，
   要實際跑一次最小指令，因為佔位程式會出現在 PATH 上）。
2. 若都不可用而檔案為 OPC 格式（`.docx` / `.xlsx` / `.pptx`）：用平台內建解壓取出
   主文 XML，再用任一可用 runtime 剝標籤。
3. 剝標籤前**先把結構轉成佔位符**（段落、表格、列、儲存格、換行、定位、圖片），
   再移除其餘標籤，最後還原 XML 實體。順序反了會失去文件結構。
4. 抽取結果是**唯讀審閱用**。要修改該格式檔案時不可用此法產出——編輯需走
   解壓 → 改 XML → 重新封裝的完整流程，並保留原始封裝結構。

#### Agent Action

下次遇到讀不到專有格式文件：

- **先探測、再宣告能力不足**。不要在未實測 runtime 的情況下回報「環境不支援」。
- 探測時對每個候選工具跑一次最小可驗證指令；PATH 上存在 ≠ 可執行。
- 若走 XML 抽取路徑，在回覆中說明「此為文字抽取，非完整轉檔」，並註明圖片、
  頁首頁尾、註解、追蹤修訂等內容可能未涵蓋。
- **不要**把抽取腳本留在業務專案目錄；放工作階段暫存區。

#### Goal / Action / Validation

- Goal: 在缺 converter 的環境仍能取得可審閱的文件全文，而不把環境限制轉成使用者的工作。
- Action: 能力探測 → OPC 解壓 → 結構標記 → 標籤剝除 → 實體還原。
- Validation or reference source: 抽取後檢查行數與字元數是否與文件規模相稱、
  段落與表格邊界是否保留、是否出現整份黏成單行的徵象；必要時與文件已知章節標題交叉比對。

#### Applies When

- 目標檔案為 OPC 壓縮格式（`.docx` / `.xlsx` / `.pptx`）。
- 只需**讀取**內容（審閱、比對、摘要、抽取條號）。
- 環境有任一解壓能力與任一腳本 runtime。

#### Does Not Apply When

- 需要**修改**並回存該檔案（改用解壓 → 改 XML → 重新封裝的完整流程）。
- 需要保真轉檔或版面渲染（需 office suite 或 converter）。
- 舊二進位格式（非 OPC，例如 `.doc`）——必須先轉檔。
- 需要讀取追蹤修訂、註解、頁首頁尾等非主文內容時，本法覆蓋不完整。

#### Validation

- 抽取檔的規模與來源 XML 規模相稱（過小通常表示標籤剝除吃掉了文字）。
- 隨機抽查數個已知章節標題可在輸出中找到，且順序與文件一致。
- 表格區域仍可辨識列與儲存格邊界。

#### Promotion Target

- `intelligence/engineering/heuristics/`（若「先探測環境能力再選路徑」在多個情境重複出現，
  可提升為通用 heuristic，本 lesson 為其具體個案）
- 暫不進 `workflow/`：此為工具環境備援技巧，非執行流程階段。

#### Required Linked Updates

- 無需連動更新：本 lesson 未改變任何既有 workflow 階段、gate 或模板；它補的是
  工具不可用時的備援判斷，屬 lesson 層。
- 依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：
  本檔不含檔名、當事人資訊、本機絕對路徑或具體腳本內容；只保留可重用的判斷順序與驗證方式。
