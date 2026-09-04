> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-04 - 只在 commit 階段跑的 gate，繞過一次就等於永久失效

Status: validated

#### One-line Summary

若某個 gate 只掛在 commit 而不掛在 push，它檢查的是「這次提交的內容」而不是「分支的狀態」，任何一次繞過或跳過都不會再被任何後續動作抓回來。

#### Human Explanation

驗證通常分階段：commit 跑便宜的檢查、push 跑昂貴的。分階段本身沒問題，問題是**分派**——若某類檢查只出現在 commit 階段，它就從「分支不變式」降級成「單次提交的一次性檢查」。

兩條路都會讓它靜默失守：

- **繞過**：`--no-verify`、hook 未安裝、hook 自身有缺陷。該次提交不受檢查，而 push 不會重驗，違規就此定居。
- **路徑過濾**：commit 階段常依「本次 staged 路徑」決定要不要跑某個 lane。一個沒有 staged 對應路徑的提交（例如只改文件）不會觸發，但它仍可能與既有違規共存於同一棵樹。

兩者的共同結果是：**沒有任何時刻會對「分支現在的樣子」執行這些檢查**。違規不會被報出來，只會累積，直到某個人的提交剛好 staged 到對應路徑時才一次爆開——而那個人通常與成因無關。

常見誤判：看到 lane 名稱寫著「pre-push 跑重的、commit 跑輕的」，會直覺以為 push 是 commit 的超集。實際可能相反——push 只跑重的，輕的一個都沒跑。腳本註解甚至可能寫著「always run」，而該 lane 剛好不在那個 always 的清單裡。

#### Trigger

- 分支上同時存在多個違規，且都能被本機檢查抓到，卻沒有任何一次 push 擋下來
- 違規全部來自同一段期間的工作，之後無人察覺
- 某人做了一次不相干的提交後，突然被一批與自己無關的違規擋住

#### Evidence

- Tool: 驗證入口腳本的 lane 條件判斷；版本控制歷史
- Sanitized excerpt: 一組結構檢查掛在 lane 清單 `all | commit | static`，而 push 使用的 lane 不在其中；commit 階段又只在「本次 staged 路徑符合」時才進入該 lane。結果是五個各自獨立的違規在同一段期間進入主線並持續存在，直到另一個人的提交剛好 staged 到相關路徑才被擋下。把 push lane 加入清單後，以重新注入其中一個違規確認 push 階段確實會回報。
- Evidence path: 具體腳本與 incident 紀錄留在 `<PROJECT_ROOT>`；此處只保留一般化規則。

#### Generalized Lesson

**分階段驗證時，要先分清楚每個檢查是「提交級」還是「分支級」，分支級的必須同時掛在 push（或任何對外發佈的閘）上。**

判斷方法：問「如果這個檢查被跳過一次，還有什麼會再抓到它？」若答案是「沒有」，它就必須出現在對外發佈那一關。

成本通常不是理由——會被歸進「commit 階段」的檢查本來就便宜；把它們加進 push 只是重跑一次，而換到的是「分支狀態受檢」這件事本身。

順序上有個陷阱：若主線目前已有既存違規，先補閘會立刻擋住所有人的下一次發佈。**必須先清乾淨，再補閘**，這個順序不能反。

#### Agent Action

**應該做：**

1. 稽核驗證入口的 lane 對應表，逐一列出每個檢查實際會在哪些階段執行——讀條件判斷，不要讀 lane 的名稱或註解。
2. 特別檢查 push／發佈階段使用的 lane 是否被排除在「總是執行」的清單之外。註解寫 always 不代表該 lane 在清單裡。
3. 補閘前先跑一次完整檢查，把既存違規一次列完再動手；修完才加閘。
4. 補完後以「重新注入一個違規」確認該階段確實會擋，而不是只確認它通過。

**不應該做：**

- 不要因為 commit 階段有跑就認定該檢查已受保障；commit 階段是可繞過且可能依路徑跳過的。
- 不要在主線仍有既存違規時就把閘加到發佈階段——那會把清理成本轉嫁給下一個發佈的人。
- 不要用「push 只跑重的」當作分派原則；輕重是成本分類，提交級／分支級才是責任分類。

#### Goal / Action / Validation

- Goal: 讓分支級不變式在對外發佈時真的被檢查，而不是只在某些提交上被抽查。
- Action: 稽核 lane 對應表；先清既存違規，再把分支級檢查加進發佈階段的 lane。
- Validation: 重新注入一個該類違規，確認發佈階段會回報；並確認補閘前主線已為綠。

#### Applies When

- 驗證分成多個階段或 lane，且不同階段執行不同的檢查子集
- commit 階段的檢查會依變更路徑決定是否執行
- 本機 hook 可被跳過，或曾經有過會靜默通過的缺陷

#### Does Not Apply When

- 所有檢查都在單一必經關卡執行（例如僅由伺服器端流水線把關，本機不設閘）
- 該檢查本質上只對「這次變更」有意義（例如提交訊息格式），不是分支不變式
- 發佈階段本來就會重跑完整檢查集

#### Validation

以來源事件反向驗證：補閘前，注入的違規在發佈階段不會被回報；補閘後同一個違規會被擋下。另可用歷史佐證——若同類違規曾在主線累積且無任何發佈被擋，即為該檢查未涵蓋發佈階段的證據。

#### Promotion Target

- `workflow/<domain>/execution-flow.md`（gate 設計章節：提交級／分支級的分派原則）
- `enforcement/`（若要提升為「分支級不變式必須在對外發佈關卡受檢」的全庫規則）

#### Required Linked Updates

- 與 [`2026-09-04_101559-build-gate-must-restore-not-only-build`](2026-09-04_101559-build-gate-must-restore-not-only-build.md) 同族：該條談 gate 驗到的是不是當前工作樹，本條談 gate 有沒有在該跑的時候跑。
- 與 [`2026-09-04_140100-index-link-gate-must-accept-the-resolvable-form`](2026-09-04_140100-index-link-gate-must-accept-the-resolvable-form.md) 同源事件。
- 依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：專案、分支、人員與 commit 證據留在 project docs。
- 已更新 [`feedback/history/development-guidance/README.md`](../README.md) 的 Recent 索引。
