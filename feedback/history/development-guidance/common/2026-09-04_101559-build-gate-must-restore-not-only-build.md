> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-04 - 以 `--no-restore` 建置的 gate 驗的是還原圖，不是工作樹

Status: validated

#### One-line Summary

當套件／專案的相依關係來自 restore 產生的鎖定檔而非組件 metadata 時，用 `--no-restore` 建置的 gate 驗證的是「上次 restore 時的相依圖」，同一棵樹在不同人的機器上會得到不同結果。

#### Human Explanation

多數現代建置工具把「相依關係解析」與「編譯」拆成兩步：restore 產生一份鎖定檔（相依圖），編譯只讀那份鎖定檔。為了速度，驗證腳本常在編譯時關掉 restore。

問題在於**鎖定檔是每個 checkout 各自的衍生狀態，不在版本控制內**。當甲在共用專案上新增一條專案 reference，乙的鎖定檔仍停在新增之前；乙的 gate 會對著舊圖編譯，然後回報一個「宣告明明就在那裡」的缺少符號錯誤。

兩個方向的後果不對稱，而且比較危險的是第二個：

- **鎖定檔少了一條現存的 reference** → 編譯失敗。吵，但看得見。
- **鎖定檔還留著一條已被移除的 reference** → 編譯成功。gate 回報綠燈，但它驗證的樹並不是它宣稱的那棵樹。

第二種情況下 gate 沒有壞掉的跡象，只是它證明的東西已經和倉庫內容脫鉤。

常見誤判：看到「宣告存在但編譯說找不到」時，直覺會往 reference 傳遞性、可見性修飾詞、專案設定去查。這些都會查不出東西，因為宣告與組件 metadata **都是對的**——錯的是那份沒人會去看的衍生鎖定檔。

#### Trigger

Gate 或本機建置回報缺少型別／命名空間／符號，但：

- 該 reference 在專案檔中明確宣告
- 中介專案編譯後的組件 metadata 確實帶著那條 reference
- 被引用專案能單獨建置成功，產出物也存在
- 同一份原始碼在其他人機器上建置正常

#### Evidence

- Tool: 建置工具的診斷層級 log；還原鎖定檔的結構化查詢
- Sanitized excerpt: 中介專案宣告了對 X 的 reference、其組件 metadata 含 X 的 AssemblyRef、X 的產出物也在中介專案的輸出目錄內；但該 checkout 的還原鎖定檔中，中介專案的 dependencies 清單沒有 X。鎖定檔時間戳早於該 reference 被加入的時間。重跑 restore 後，同一個 `--no-restore` 建置成功。
- Evidence path: 專案端的 gate 腳本與 incident 紀錄留在 `<PROJECT_ROOT>`；此處只保留一般化規則。

#### Generalized Lesson

**若相依解析與編譯是分離的兩步，gate 必須自己負責讓相依圖是最新的，不能依賴「開發者上次剛好跑過」。**

在 gate 內把 restore 當成必要前置步驟明確執行；編譯階段仍可關掉 restore，避免重複執行。判斷成本時要注意：暖 restore 通常只是數秒，而它換掉的是一個結果會因人而異的 gate。

同理適用於任何「鎖定檔／解析快取不在版本控制、但編譯讀它」的生態系。

#### Agent Action

**應該做：**

1. 檢查 gate 是否用了跳過相依解析的旗標（如 `--no-restore`、`--frozen-lockfile` 配上過期鎖定檔、離線模式）。若有，確認 gate 內另有明確的解析步驟。
2. 遇到「宣告存在但編譯找不到」時，**先查衍生狀態**（鎖定檔、解析快取、產生的中繼檔），再查語言層語意（傳遞性、可見性、別名）。先看鎖定檔的時間戳與內容，比讀五個專案檔快得多。
3. 修好之後，明確說明比較危險的是相反方向（過期圖讓 gate 誤過），而不只是「本來會失敗現在會過」。

**不應該做：**

- 不要因為「宣告在那裡但編譯說沒有」就去補一條重複的 reference。那會讓症狀消失但保留成因，而且在別人 restore 是新的的機器上會變成一條沒必要的宣告。
- 不要把 gate 內的解析步驟拿掉來省時間，除非能證明鎖定檔本身受版本控制且與樹同步。

#### Goal / Action / Validation

- Goal: 讓 gate 的結果只取決於工作樹內容，不取決於該 checkout 的衍生狀態新舊。
- Action: 在 gate 的建置階段前明確執行相依解析；建置本身維持跳過解析以免重跑。
- Validation: 在相依圖被改動過、但本地未重新解析的 checkout 上跑 gate。修正前應重現失敗，修正後應通過；並確認暖執行的額外成本可接受。

#### Applies When

- 建置系統把相依解析結果存進不受版本控制的鎖定檔或快取
- 驗證腳本為了速度使用跳過解析的旗標
- 多人／多 checkout 共用同一組專案，且 reference 圖會被修改

#### Does Not Apply When

- 相依圖完全由受版本控制的檔案在編譯期直接讀取，沒有中間解析產物
- 每次 gate 都在乾淨環境（容器／CI 全新 checkout）執行，解析必然發生
- 鎖定檔本身受版本控制且有機械閘保證它與專案檔同步

#### Validation

在一個相依圖已被他人修改、但本地尚未重新解析的 checkout 上執行 gate：修正前重現該類錯誤，修正後同一棵樹通過。反向情況（鎖定檔仍留著已移除的 reference）較難自然重現，可用暫時移除一條 reference 但不重新解析的方式驗證 gate 是否仍會誤過。

#### Promotion Target

- `workflow/<domain>/execution-flow.md`（gate 設計章節：明確列出前置解析步驟）
- `enforcement/`（若要提升為「gate 結果不得依賴不受版本控制的衍生狀態」的全庫規則）

#### Required Linked Updates

- 本 lesson 與 [`2026-09-04_101800-derived-state-before-blaming-shared-baseline`](2026-09-04_101800-derived-state-before-blaming-shared-baseline.md) 同源；該條寫的是歸因紀律，兩條互相連結。
- 依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：具體專案、人員、分支與 commit 證據留在 project docs，此處只保留一般化規則。
- 已更新 [`feedback/history/development-guidance/README.md`](../README.md) 的 Recent 索引。
