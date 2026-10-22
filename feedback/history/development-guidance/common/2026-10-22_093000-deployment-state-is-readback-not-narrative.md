> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-22 - Deployed state is a readback, not a narrative; the omission starts where the code is written

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源是一個多人、多 agent 共用 repo 的客戶正式環境；目前只完成設計與現況比對，尚未有一次照此流程完成的真實部署。

#### One-line Summary

「正式機上跑的是什麼」必須來自讀回並寫進 repo 內的帳本，git tag 只是從帳本衍生出來的索引；遺漏多半發生在寫程式的一端，所以完成定義（DoD）要求留下給部署者的通知，部署前由會失敗的 preflight 彙整。

#### Human Explanation

很多專案的開發流程停在「review／validation／closure」，release gate 只保證「publish 前測試是綠的」。部署本身靠一份操作清單和交接文件。這樣會產生三個缺口：

1. **沒有人知道正式機現在是哪個版本。** 交接文件寫的是「當時部署了什麼」，不是「現在是什麼」；部署中斷、局部部署、有人另外更新，都不會反映在文件上。下一次部署要算差異範圍時，只能從對話或記憶重建。
2. **需要操作者注意的變更，是在寫程式的當下產生的**（例如改了一個已套用過的 migration、需要新的設定、有套用順序）。如果 DoD 沒有要求留通知，部署的人根本無從得知；部署當下再讀清單也救不回來。
3. **只寫在文件裡的規則沒有人會因為沒照做而失敗。** 「部署前先讀通知」這種行為規則，agent 讀到了才會照做；沒有使用該工具鏈的 agent 或人，入口文件裡根本沒有部署這一步。

常見誤判：

- 以為 git tag 就能記錄版本。tag 只說得出「某個 commit」，說不出每個站台各自的版本、資料庫 migration 套用到哪裡、讀回有沒有成功；而且沒部署也能打 tag。
- 以為把部署流程抄進每一份入口文件就能避免忘記。內容複製多份會各自漂移；入口文件只需要指到會失敗的 preflight。
- 只記版本、不記環境。程式版本對了、但資料庫大版本、runtime、時區和實測環境不同，一樣會出事。

#### Trigger

- 專案要部署到一個或多個客戶的正式環境，且有多人或多個 agent 輪流部署。
- 交接文件寫「未讀回版本」「不代表即時狀態」，或部署差異範圍只能靠 `git log <記憶中的 commit>..HEAD` 推算。
- 開發流程的最後一個 phase 是 closure，沒有 deployment phase。

#### Evidence

- Tool: `git rev-list --count`、`git diff --name-status <last-recorded-deploy>..HEAD -- <migration paths>`、讀健康檢查端點的原始碼。
- Sanitized excerpt: 從最後一次有記錄的部署到目前 HEAD，差了數百個 commit、上百個新 migration 檔，以及十幾個被修改的舊 migration（這次只改了註解）；既有的部署通知只涵蓋其中一個功能的 migration。其中一個站台的部署版本在交接時就標記為未讀回。健康檢查端點其實支援回報 build commit，但部署步驟沒有設定它，所以版本無法機械讀回。
- Evidence path: `<PROJECT_ROOT>` 的部署計畫與部署通知目錄（project-specific，不複製到本 repo）。

#### Generalized Lesson

1. **分三層，各有權威：**
   - **部署帳本**（repo 內，按 `<customer>/<env>` 分檔）是 canonical，只記讀回後的事實：每個站台的 commit、資料庫每個 migration context 的最後一筆、讀回方式、消化了哪些通知。不知道就寫 `unknown`，`unknown` 是合法值。
   - **git tag**（`deploy/<customer>/<env>/<site>/<timestamp>`，annotated，不可移動或刪除）是回溯用的索引，只在帳本寫入成功後產生。按站台打，因為站台常常不在同一個 commit。
   - **部署通知**（帶 owner、status、影響站台、是否擋部署的 frontmatter）是給下一次部署者的待辦。
2. **三道閘：**
   - commit／push：帳本、通知、tag 三者一致；改到「帳本 commit 之前就存在」的 migration，必須有通知點名並宣告是只改註解還是改了行為。
   - 部署前 preflight：從帳本算差異範圍並列出來；有擋部署的待確認通知、或基準版本是 `unknown` 時以非 0 結束（要繼續必須明確帶理由），並比對遠端的 tag 與帳本，因為 hook 會被 `--no-verify` 繞過。
   - 部署後 record：讀回健康檢查的 build commit 與資料庫 migration 狀態，對得上才寫帳本、建 tag；對不上就照實記錄，不建 tag。
3. **環境也要記：** 每個客戶環境一份環境檔（OS／runtime／資料庫版本、時區、站台權限、用戶端瀏覽器與裝置、外部串接、已對外承諾不能改的固定值名稱），每一項標明來源（readback／operator／customer）與 `verified_at`；機密值與實際 IP 不寫。
4. **流程文件只放指標：** 開發流程加一個 deployment phase、DoD 加「需要操作者處理時，同一個 commit 附通知」、release gate 文件說明它不等於部署閘、部署清單第 0 步是 preflight。流程內容只寫在部署文件與腳本中。

#### Agent Action

- 被要求部署或寫部署相關文件時，先找「正式機現在跑的版本」的讀回來源；找不到就明說 `unknown`，不要用交接文件的「當時部署了什麼」代替。
- 寫程式時若改到已套用的 migration、新增必要設定或套用順序，在同一個 commit 留部署通知。
- 不要把部署流程整段複製到多份入口文件；只加指向 preflight 的指標。
- 不要在 tag、帳本或環境檔寫入機密值或實際網路位址。

#### Goal / Action / Validation

- Goal: 任何人或 agent 部署前都能從 repo 內容得知各站台現況與差異範圍，並在有未處理事項時被擋下。
- Action: 帳本 + tag + 通知 + 環境檔；commit／preflight／record 三道閘；流程文件加指標。
- Validation or reference source: preflight 在目前 repo 狀態下必須停下來（該紅的要紅）；第一次真實部署完全由 preflight → 部署 → record 完成。尚未執行，所以維持 candidate。

#### Applies When

- 部署到自己無法隨時觀察的環境（客戶現場、內網機器），且部署者會輪替。
- 有資料庫 migration 或需要操作者處理的設定。

#### Does Not Apply When

- 已有 CD 平台自動記錄每次部署的 artifact 與環境版本，並可查詢目前版本（此時只需要「寫程式的一端留通知」這一條）。
- 只有單一開發者、單一環境，且每次都是整包重建。

#### Validation

- 用目前 repo 跑 preflight，必須列出全部新增與修改過的 migration，並因 `unknown` 基準與未處理的通知停下來。
- 第一次真實部署後：帳本與 tag 由腳本產生，讀回值與健康檢查端點一致。
- 刻意移動一個 deploy tag、改舊 migration 但不附通知，commit／push 閘都要紅。

#### Promotion Target

- `workflow/software-delivery/release/README.md`（目前只有一句話；補上 deployment ledger／readback／environment profile）
- `workflow/software-delivery/execution-flow.yaml` 的 operational transaction 段落（補上 deploy 的讀回與帳本要求）
- `workflow/software-delivery/closure.md` 或 DoD（「需要操作者處理時附通知」）

#### Promotion Record

尚未 promotion；保留為 candidate history。等來源專案完成第一次照此流程的真實部署，依實際結果修訂後再提出 promotion。

#### Required Linked Updates

- 已更新 `feedback/history/development-guidance/common/README.md` 與 `feedback/history/development-guidance/README.md` 的索引。
- 已依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 檢查：只保留通用方法；客戶名稱、站台、port、commit、數字細節留在專案文件。
- Promotion target 暫不修改：缺少真實部署證據，理由見 Promotion Record。

#### 修訂 2026-10-22（來源專案 Phase 1 完成後）

Status 仍為 candidate（尚未有真實部署走完流程）。實作與設計的差異：

1. **紀錄要是部署腳本的副產品，不是另一個要記得執行的指令。** 原設計是「preflight」「record」兩支腳本；使用者指出多一個指令就多一個漏掉的點，
   改為單一部署入口：通過 preflight 時自動開草稿，部署後自動讀回填入，只詢問機器判斷不了的欄位。
2. **用「未完成的紀錄草稿」當鎖。** 草稿存在或殘留 `TODO` 時，pre-commit 與 pre-push 都失敗；檢查讀檔案系統，所以未追蹤的草稿也會鎖住部署者自己的機器。
   沒有略過旗標，唯一出口是 `abort` 加理由。已知限制：刪掉未追蹤草稿就能解鎖，需要第二道痕跡或規則補強。
3. **衍生紀錄的一致性檢查不能依賴可變的鄰居。** 帳本 entry 由紀錄檔推導；若檢查器用「現在的前一筆」重算未部署站台的沿用值，之後合法回填舊 entry 會讓檢查器紅且無出口。
   修正：推導時把沿用值凍結進紀錄檔，檢查器只用紀錄本身重建。這條在任何「由 A 推導 B、並檢查 B」的閘都適用。
4. **機器讀的檔案配合執行端格式。** 部署端是 Windows／PowerShell 時，JSON（內建解析）比 YAML（需模組）適合；通知用扁平 frontmatter，兩種 shell 都能不靠套件解析。
5. **`owner` 是「要照做的人」，不是撰寫者。** 以 git author 推導會錯；由使用者指定並給預設值。
