# Consumer A：暫存檔登記框架（自家實際套用的唯讀評估）

- **日期**：2026-11-24
- **起因**：使用者 2026-11-24 指定以 Consumer A 的工具樹作為第三個案例（本計畫最早就是從整理這個工具樹開始的）。盤點時發現 Consumer A 當天自己推上了一個暫存檔登記框架。使用者裁決：**唯讀評估這個框架**，當作實際套用的候選案例；Phase 2 建議的 R1–R4 只標成仍未處理。
- **性質**：Consumer A 自己的決策和改動，已經合併到它的上游主幹。**不是**本計畫對 consumer 的修改。這和 [`consumer-b-cli-pilot.md`](consumer-b-cli-pilot.md) 同類。框架有沒有參考本契約，從 commit 看不出來，判為未知。
- **方式**：全程唯讀，沒有修改 Consumer A 的任何檔案。repo 名稱、路徑、SHA 不寫在這裡（去敏規則）。

## 1. 框架做了什麼（Consumer A）

| 層（Role contract） | 內容 |
| --- | --- |
| Parameters | 一份 policy 檔，只有純 literal（不 import、沒有運算式），鎖用 `ast.literal_eval` 讀取。內容包括：類別名稱對應目錄樣板、清理規則（由上往下比對、第一條命中的生效）、全域保留名稱、保留字 |
| Logic | 路徑配發（`temp_dir` / `temp_path` / `mkdtemp` / `mkstemp` / `temp_file` / `tool_path`，未登記的類別在執行期報錯）、清理規劃與套用（sweeper） |
| Entry | 一支 16 行的清理 CLI（`--list`、`--unregistered`、`--apply`） |
| Lock | 一把 pre-commit 鎖：用到暫存檔卻不走框架（`tempfile`、`/tmp`、`TMPDIR`、手寫工作目錄、未登記類別）就擋。測試 32 項，違規和合法相似寫法兩個方向都有；豁免標記必須寫原因 |

另外把散落各處的暫存寫法遷到框架；原有的大廳工作目錄 helper 改成框架之上的薄包裝；順手讓一支單目標 runner 改用共用裝置模組。

## 2. 對照契約

| 契約項目 | 判定 | 說明 |
| --- | --- | --- |
| Scratch 位置：不進版控的固定目錄，README 有寫明，並列入 ignore | **符合** | 工作根目錄和每個目標的 capture 目錄都由 `.gitignore` 排除（用 `git check-ignore` 確認過）；README 新增一節說明 |
| Scratch 不放在正式工具樹 | **符合** | 工作根目錄在資料輸出區，不在 `scripts/` 底下 |
| 清理責任、殘留回報 | **符合，而且做得比要求多** | 契約只定責任，不要求機制；框架加了機制。`--unregistered` 會列出沒被任何規則管到的檔案，對應「回報殘留」 |
| Scratch 不得長期存在 | **有張力** | 工作根目錄底下有 `keep` 規則的長期內容：輔助 binary 的存放區，以及幾組素材庫。契約說有保留價值的要轉成 Research / experimental，但那條是寫給**工具**的；這裡是**資料產物**。見 §3 C1 |
| 機械清理的邊界（G4 修訂後的三個條件，類比套用） | **2／3** | (1) 只作用在指定區域：符合；(2) 排除無法確認使用情況的項目：大致符合（`keep_if_referenced` 只掃 docs/scripts，未登記的只回報不刪）；(3) **先隔離、可恢復：不符合**，套用時直接 `rmtree` / `unlink`。緩解措施：預設 dry-run、第一次需要人工 `--apply` 才會啟用每日清理、被刪的都是 gitignore 的輸出 |
| Role 分層 | **正例** | Parameters 是純資料，鎖不 import 它而是直接解析；Logic 集中；Entry 很薄 |
| Concept-first / 先找再建 | **正例** | 同一個責任（配發暫存路徑）原本散在多處，現在收斂成一個能力；既有的 helper 改成薄包裝，沒有另起一套 |
| Lock taxonomy | **符合** | 雙向測試；豁免需要寫原因 |
| Ownership（Q9） | **缺口不變** | 沒有 owner 欄位，和 Phase 2 判定相同。`managed_by` 是「清理責任委派給哪支工具」，不是維護 owner |
| 鎖逐檔蔓延（Phase 0 A5） | **持續** | 檢查檔從 24 個變成 25 個，pre-commit 多一行呼叫，形狀和 A5 相同 |

## 3. 契約問題（候選，交給 Phase 4）

| # | 問題 |
| --- | --- |
| C1 | Scratch 與機械清理條款是以「工具」為對象寫的，但實際案例管的是**暫存資料和輸出產物**。契約要不要區分 scratch 工具和 scratch 資料，還是明文寫出資料產物不適用、或可以類比適用。單一來源 |
| C2 | G4 第 3 條（先隔離、可恢復）放在 gitignore 的、可重新產生的輸出上，是否過嚴？還是「預設 dry-run 加人工啟用」可以算等效緩解？單一來源 |

## 4. Phase 2 建議項目的現況（重新量測）

用 project-local 對照檔裡可重跑的指令量測：

| 項目 | Phase 0 / 2 | 2026-11-24 | 狀態 |
| --- | --- | --- | --- |
| R1 / A1 非薄入口 | 3 支（31／132／260 行） | 3 支（31／138／266 行，暫存遷移各多 6 行） | 仍未處理 |
| R2 / R3 / A3 寫死的裝置序號預設值 | 4 檔 | 4 檔（同一批檔案） | 仍未處理 |
| 鎖檔數 | 24 | 25 | 新增的是本框架的鎖 |

## 5. 對 Q11 的意義（觀察，不是裁決）

現在有三個「實際套用」的候選案例：

1. Ai-skill 自家套用（本計畫做的）；
2. Consumer B 的 CLI 試點（consumer 自己做，已合併）；
3. Consumer A 的暫存登記框架（consumer 自己做，已合併）。

後兩個的共同疑點相同：都是 consumer 自己的決策，不是本計畫的套用；而且都只套用了契約的一部分（B 是 Role 分層加先找再建，A 是 Scratch、Role、concept-first）。它們算不算收斂案例，由 Phase 4 依 Q11 裁決。

## 6. 後續：由 Consumer A 自己的 agent 修改，本計畫重新檢查

使用者 2026-11-24 提議：Consumer A 的修改由使用者另開一個在 Consumer A 內工作的 agent 執行，完成後回報，再由本計畫唯讀重新檢查。Ai-skill 仍然不直接修改外部 repo。交付的 brief 只涵蓋 Phase 2 建議清單第 1–6 項（R1–R4、裝置 I/O 鎖範圍、thin-entry 鎖）；C1、C2 是尚未裁決的契約問題，不要求 consumer 照做。重新檢查的結果另外記錄。

← [Evidence index](README.md)
