> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-04 - 檢查索引連結的 gate，必須接受從該索引真的能解析的那種寫法

Status: validated

#### One-line Summary

要求某個路徑字串出現在一份 Markdown 索引裡的 gate，若比對的是 repo 根相對路徑，就會拒絕唯一正確的相對連結，並接受寫上去會壞掉的那幾種。

#### Human Explanation

「每份文件都要從索引連過去」是常見的治理規則，實作起來也直覺：從來源取出文件路徑，然後在索引檔裡 grep 那個路徑。

問題出在**路徑的基準點不同**。gate 手上的路徑通常是 repo 根相對的（`docs/section/sub/item.md`），因為它是從別處的清單或檔案系統得到的。但索引檔裡的連結是**相對於索引自己所在的目錄**——一份位於 `docs/section/` 的索引，連到子目錄的正確寫法是 `sub/item.md`。

於是 gate 要求的字串，寫成 markdown 連結後會指到深一兩層的地方而失效；而唯一能正確解析的那種寫法，gate 不認得。作者照著能用的方式寫，然後被擋。

這類缺陷有個特徵：**在第一個真實案例出現之前完全看不出來**。若該類別長期是空的（沒有任何 Open 項目、沒有任何待歸檔文件），gate 的分支從未被執行過，任何人讀程式碼也不會覺得有問題——因為那幾種寫法「看起來」都很合理。

還有一層時序放大：這類同步檢查常設計成 commit 階段警告、push 階段阻擋。設計本身是對的，但代價是**寫文件的人只看到警告，錯誤留給下一個 push 的人**——而那個人手上沒有前因後果，很容易誤判成對方漏做。

#### Trigger

- 某個「必須從索引連結」的檢查失敗，但索引裡的連結明明存在且在文件檢視器中可以正常跳轉
- 該檢查的錯誤訊息要求的路徑形式，與索引檔裡既有連結的形式不一致
- 該類別是第一次有實際項目，或長期為空

#### Evidence

- Tool: gate 的比對邏輯；索引檔的連結寫法；連結解析驗證
- Sanitized excerpt: gate 從來源行取得 repo 根相對路徑，接受該路徑與去掉最上層目錄後的兩種形式；索引檔位於中層目錄，其連結以自身目錄為基準寫成。三種形式中只有索引自身相對的那一種能正確解析，而那正是 gate 唯一不接受的。該類別在此之前沒有任何項目，分支從未執行。
- Evidence path: 具體 gate 與文件結構留在 `<PROJECT_ROOT>`；此處只保留一般化規則。

#### Generalized Lesson

**比對「連結是否存在」時，接受的形式必須以該索引檔自身的位置為基準，因為那才是連結實際被解析的方式。**

實作上有兩種做法：

1. **接受多種拼法**——同時容許 repo 根相對與索引目錄相對的形式。改動小，適合既有 gate。
2. **真的解析連結**——從索引中抽出 markdown 連結目標，以索引所在目錄為基準解析成絕對路徑後再比對。較嚴謹，同時能抓出壞連結。

判斷是否誤判的快速方法：**把 gate 要求的字串當成連結貼回索引裡，看它指到哪。**若指不到目標檔，就是 gate 錯而不是文件錯。

修法方向也由此決定：**不要把文件改成 gate 想要的壞連結**。讓症狀消失但留下壞掉的文件，比原本的狀況更糟。

#### Agent Action

**應該做：**

1. 遇到「索引未連結」類的失敗，先驗證索引裡的連結是否真的可解析。可解析卻被擋，先懷疑 gate。
2. 檢查該 gate 的比對基準點：它手上的路徑從哪來、以什麼為基準；索引檔位於哪一層。兩者基準不同就是成因。
3. 檢查該類別是否為第一個實際案例。若是，這條分支從未被執行，缺陷存在的機率明顯偏高。
4. 放寬拼法後，以「移除索引中的連結，確認仍會失敗」驗證這是放寬而非取消。

**不應該做：**

- 不要為了讓 gate 通過而把連結改成 repo 根相對形式——那在文件中是壞連結。
- 不要因為 gate 是既有的就假設它被驗證過；長期為空的類別等於未執行的程式碼。
- 不要在確認是 gate 缺陷前，就把失敗歸因為文件作者漏做；同步類檢查的錯誤常在別人的 push 上才浮現。

#### Goal / Action / Validation

- Goal: 讓「必須從索引連結」的檢查與連結實際解析的方式一致，避免拒絕正確寫法。
- Action: 比對基準對齊索引檔位置——接受索引目錄相對的形式，或直接解析連結後比對。
- Validation: 現況（正確的相對連結）通過；移除索引中的連結後仍回報未連結，證明要求本身未被取消。

#### Applies When

- gate 以字串比對確認某路徑出現在一份 Markdown 索引中
- 索引檔不在 repo 根，且連結以自身目錄為基準書寫
- 該檢查的類別可能長期為空，直到第一個實際項目出現

#### Does Not Apply When

- 索引與被連結檔案同層，兩種基準恰好一致
- gate 已經真正解析連結而非字串比對
- 索引由工具產生，寫法由產生器統一決定

#### Validation

以來源事件反向驗證：把 gate 要求的字串當作連結貼回索引，確認它無法解析到目標；修正後正確的相對連結通過，而移除連結仍會失敗。

#### Promotion Target

- `workflow/<domain>/execution-flow.md`（文件治理章節：索引連結檢查的基準點）
- `enforcement/`（若要提升為「路徑比對必須與該檔案的解析基準一致」的通用規則）

#### Required Linked Updates

- 與 [`2026-09-04_135746-gate-that-runs-only-at-commit-is-one-bypass-from-gone`](2026-09-04_135746-gate-that-runs-only-at-commit-is-one-bypass-from-gone.md) 同源事件，且共享同一個更上層的觀察：gate 看起來在運作，不代表它證明的東西是對的。
- 依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：具體 gate、文件路徑與人員資訊留在 project docs。
- 已更新 [`feedback/history/development-guidance/README.md`](../README.md) 的 Recent 索引。
