# Phase 3 — Ai-skill 自家套用（切片 S1）

- **日期**：2026-11-23
- **基準**：`9df9621f`
- **切片**：S1 加五點邊界（見 [`_plan.md`](../_plan.md) §Phase 3 套用切片）
- **改動範圍**：只有文件，即 catalog 與 README 的流程描述。Go source、hook、registry 都沒有改。
- **狀態標記**：`verified` = 已對照 HEAD 的程式碼或 git 歷史確認；`observed` = 尚未確認。

## 1. Bin 流程預查（verified）

| 來源 | 要求 |
| --- | --- |
| pre-commit `validateReleaseBinaryParityStaged` | 有 CLI source 變更（`scripts/ai-skill-cli/` 下非 test 的 `.go`、`go.mod`、`go.sum`）時，同一個 commit 必須 stage 5 個 binary、`BUILDINFO`、`SHA256SUMS`；豁免標記 `[skip-release-binary-parity]` |
| pre-push `TestRepoLocalBinariesMatchChecksumsAndCurrentSource` | `BUILDINFO.source_commit` 必須等於「最後一個 CLI source commit」，範圍是 cmd、internal、portable、go.mod、go.sum，排除 `_test.go` |
| 歷史實例 | `bb927a7e`（source 加上完整 bin）→ `4f86bbf4`（只含 bin 的重建）；現行 `BUILDINFO` 指向 `bb927a7e` |

**結論**：
- 一定要兩個 commit：第一個帶 source 和完整 bin，第二個只放 bin，重建時用 `--commit <source commit>`。
- 舊說法「先只 commit source」會被 pre-commit 擋下。
- 已修正的文件：[`scripts/ai-skill-cli/README.md`](../../../../scripts/ai-skill-cli/README.md)。
- 還沒同步：[`runtime/cli-modification-policy.yaml`](../../../../runtime/cli-modification-policy.yaml) 的 `gate.cli.bin_rebuild_for_HEAD` 仍寫用「HEAD SHA」，列為後續工作 F4。

## 2. Capability catalog

18 列已寫入 [`scripts/ai-skill-cli/README.md`](../../../../scripts/ai-skill-cli/README.md) §Capability catalog，[`scripts/README.md`](../../../../scripts/README.md) 加了指向它的入口。

- **placement**：17 列是 Domain（服務 Ai-skill 治理領域）；`kge` 的 Logic 在 Shared 的 `portable/kge`。
- **`evidencecandidate`**：維持 Research / experimental。依使用者裁決（2026-11-23）設 TTL 2026-12-23，來源 plan 的 observation period 若更早結束就以那天為準；它的 phase2 gate 是計數門檻，沒有固定日期。catalog 已連到實驗目的與輸入／輸出 contract。到期時三選一：轉為 maintained、附理由續期，或 deprecated；不會自動刪除。
- **`scripts/git-hooks/*`**：判定為 deprecated。移除條件是「確認沒有 clone 還在使用舊路徑」；由於使用情況無法機械確認，依 G4 排除自動清理。

## 3. Lock 對帳：hook → executor → registry

| Hook 階段 | 執行點 | 有 registry 綁定 | 沒有綁定 |
| --- | --- | --- | --- |
| pre-commit | 8 個 validator 函式 + 2 個 `runtime compile/validate` | 6 | `warnSanitizationIncidentScore`、`validateReviewArchitectureNavigationDriftStaged` |
| commit-msg | 32 | 26 | `warnPlanArchivalLinkIntegrity`、`warnPlanEvidenceLineNumberCitations`、`validateEnforcementRegistryTransition`、`validateEnforcementRuleRegistrySync`、`validatePlanEvidenceConvention`；`validateCommitMessageGovernance` 列在 `internal_helper_allowlist` |
| pre-push | 4（replay、KGE check、Go preflight、workflow history） | workflow history 以 `runPrePushHook` 綁定 | `validatePushGovernanceReplay` |

- registry 有 44 個 entry，coverage 分布：mechanical 24、behavioral_only 13、not_mechanizable 5、research_required 2；其中 25 個 entry 有 executor。
- 46 和 44 計數單位不同：一個 entry 可以綁多個 executor（例：`cognitive_mode_governance` 綁 8 個），而 13 + 5 + 2 個 entry 本來就沒有 executor。
- **發現 L1（verified）**：registry lint 的 `lintOrphanExecutors` 有兩個限制：只掃 `hooks.go` 裡的宣告，而且只看 `run*` 和 `validate*` 前綴。上表 8 個沒有綁定的 symbol，不是宣告在其他檔案，就是用 `warn*` 前綴，所以 registry 和 lint 都看不到它們。這是 lock-taxonomy「scope 要覆蓋所有出口」的實例，與 Phase 0 的 A2 屬於同一型。

## 4. Role（切片內）

薄入口判定和重複入口分開判斷。

| 入口 | Role 判定 | 重複 / 生命週期 |
| --- | --- | --- |
| `cmd/ai-skill/main.go`（11 行） | Entry，通過 | — |
| `app.go` dispatch | Entry，通過：依命令名分派，屬於參數解析，沒有條件決策 | — |
| `.githooks/*`（21–22 行） | Entry，通過：條件分支只用來找出對應平台的 binary，屬於「取得必要設定」，不是範圍或快取決策（A-2 的邊界案例，記錄判準） | 4 份檔案只差 hook 名稱；commit-msg 多傳 `"$@"` |
| `scripts/git-hooks/*`（5 行） | Entry，通過 | 和 `.githooks` 重複；責任是相容舊的 `core.hooksPath`，判定為 deprecated |
| hooks 編排（`runPreCommitHook` 等） | 有條件決策，依 A-2 屬於 Logic，而它本來就放在 `internal/app`，通過 | — |

- **O2（verified，不算違規）**：一行的 `validateX` wrapper 有實際責任：registry 用 `validate*` symbol 登記 executor，lint 也只認 `run*` 和 `validate*`，wrapper 讓登記名稱保持穩定。Review Rule 結論：這是 registry 命名慣例造成的補償結構，依切片延後到 KGE 呼叫鏈重構（F3）。
- **O8（verified）**：`hookAdapterContent` 只有 test 在用，正式安裝已改走 `core.hooksPath=.githooks`。它產生的內容與 committed 的 `.githooks` 不同：解析順序是先用裸 `bin/ai-skill`，而 `.githooks` 的註解指出裸 `bin/ai-skill` 常是過期的本機 alias。這等於第三份外殼形狀（F5）。
- **結果**：切片內沒有 Role 違規。

## 5. Lock review：工具治理鎖（依不變式計數，共 7 把）

| # | 族 | intent | scope | forbidden | allowlist / 豁免 | 樣本狀態（見下方圖例） | Review Rule 結論 |
| --- | --- | --- | --- | --- | --- | --- | --- |
| T1 | layout | 自動化一律用 Go，不新增其他腳本入口 | staged **Added** 路徑；`runtime validate` 另查既有 `.sh` 的 Go 決策 | 12 種副檔名。**缺口**：沒有副檔名、帶 shebang 的可執行腳本抓不到（只比對副檔名；verified） | `[skip-go-migration]`，需附 migration plan | **通過**：`TestNoNewShellScriptsRule`（.sh、.py、.rb、.js 會被擋；.go、.md、.yaml 不誤判）。**缺口**：無副檔名的 shebang 腳本（讀程式碼確認，沒有樣本） | 不能消除：這是限制未來新增的政策。兩處檢查守同一不變式（G6 第 2 次），但觸發點不同（新增 vs 既有），記錄為分開保留的理由 |
| T2 | drift | 提交的 binary 與 source 一致 | pre-commit 的 staged CLI source；pre-push 的 git 歷史 | 只交部分 bin、`BUILDINFO` 指向錯誤的 commit | `[skip-release-binary-parity]` | **通過**：`TestValidateReleaseBinaryParityStaged`（全部 stage 時通過、缺任一檔時失敗）、`TestReleaseBinaryParityOptOut` | 不能消除：repo-local binary 是既有決策；兩階段是因為 commit hash 要等 commit 建立後才存在 |
| T3 | drift | CLI 新增命令或 hook 時同步更新 command-contract | 只看 `internal/app/*.go`。**缺口**：`cmd/` 和 `portable/` 不在範圍內（verified） | 6 個字面 pattern 加上 `run*Hook`。**缺口**：`app.go` 新增頂層 `case "x":` 抓不到（verified） | `[skip-cli-doc-sync]` | **通過**：engine 3 個 + adapter 1 個 test（能偵測新 hook、skip 標記有效）。**缺口**：頂層 `case`、`cmd/`、`portable/`（讀程式碼確認，沒有樣本） | **可以部分消除**：若 command-contract 的命令表改由 dispatch 產生，這把鎖就不需要（F2，不實作） |
| T4 | drift | registry 與 executor 雙向綁定 | `runtime validate` 和 commit-msg；orphan executor 只看 `hooks.go`。**缺口 L1** | orphan rule、orphan executor、缺少 symbol | `internal_helper_allowlist`：這是 helper 分類清單，不是暫時豁免，所以不適用退場條件 | **通過**：rule 方向（`TestLintOrphanRule_*`、`TestEnforcementRegistryLint*`）。**尚未驗證**：executor 方向沒有專屬 test。**缺口**：L1（讀程式碼確認） | 不能消除：registry 本身就是為了防止「有規則但沒有 executor」。L1 是鎖的正確性問題，不屬結構消除（F1） |
| T5 | safety | 共享層不提交私有 token 和泛用敏感 pattern | 所有 staged 內容；push replay 再掃一次 | private_entities 的各種寫法 + 泛用 pattern | project-local 層、placeholder、pattern 設定檔自身 | **通過**：sanitization 6 個 + incident score 2 個 + push replay 1 個（會擋的和會放行的都有） | 不能消除：任何內容都可能夾帶敏感資料 |
| T6 | safety | 選用此功能的 clone 只用允許的 author / committer email | 每次 commit 的身分 | 不在 allowlist 內的 email | 本機 git config；未設定時不限制 | **通過**：4 個 test（未設定時略過、會擋的、會放行的、讀 info 檔） | 不能消除：身分設定是每台機器各自的 |
| T7 | completeness | 改到 CLI、hook 或 workflow 時，push 前 `go test` 必須通過 | 相對 upstream 的變更路徑 | 測試失敗；workflow history 紅燈只發警告 | 沒有豁免 | **不適用**：不是 pattern-based 鎖。範圍判斷由 `TestHasCLICIPreflightChange` 涵蓋，**通過** | 部分屬於補償：CI 太慢，所以在本機重跑一次。記錄理由後保留 |

**樣本狀態圖例**（2026-11-23 於 `294bcd6c` 執行 `go test`，表中列名的 29 個 test 全部 PASS）：

| 狀態 | 意義 |
| --- | --- |
| 通過 | 已有測試且本次執行通過 |
| 缺口 | 已確認鎖抓不到某種常見寫法，或 scope 漏掉某個出口；目前沒有對應樣本 |
| 尚未驗證 | 沒有找到能證明該方向的測試 |
| 不適用 | 不是 pattern-based 鎖，不強制雙向樣本 |

**完成審查不等於驗證通過。** 7 把鎖都已審查完畢，其中 T1、T3、T4 的審查結果是發現缺陷：T1、T3 有已確認的 forbidden 或 scope 缺口，T4 另有一個方向尚未驗證。只有 T2、T5、T6 是雙向樣本都已通過；T7 不適用雙向樣本，範圍判斷的 test 已通過。

**結案裁決（2026-11-23，使用者）**：缺口移交，不在 Phase 3 補樣本。T4 executor 方向與 L1 → F1，T3（`app.go` 頂層 `case`、`cmd/`/`portable/` 出 scope）→ F2，T1（無副檔名 shebang 腳本）→ F6。Phase 3 結論為「S1 套用與審查完成，驗證缺口已移交」；這些方向仍是未驗證，不因結案而視為通過。

**未納入的項目**：`validateStagedDiffWhitespace` 和 T2 同屬 registry 的 `pre_commit_diff_integrity`，但它保護的是 repo 內容，不是工具樹，所以歸為產品不變式鎖。同一個 registry entry 底下有兩種不變式，這件事記錄為分類觀察。

## 6. 產品不變式鎖：已分類，未審查，後續處理

- **mechanical 的 18 個 entry**：cognitive_mode_governance、capability_context_governance、plan_governance、plan_archival_link_integrity、plan_tree_governance、evidence_hierarchy、runtime_trigger_wiring、bootstrap_entry_thinness、runtime_yaml_projection、runtime_index_freshness、markdown_yaml_sync、glossary_governance、dirty_repo_close_loop、feedback_learning_report、feedback_lesson_closure、linked_updates、workflow_activation、validation_scenario_governance。
- **還有**：`pre_commit_diff_integrity` 裡的 whitespace 部分，以及 §3 中沒有登記的產品類 validator（plan evidence、review-architecture drift、push replay）。
- **behavioral_only / not_mechanizable / research_required 的 20 個 entry**：沒有 executor，不屬於機械鎖，只登記。
- **契約仍要求**它們補齊 intent、allowlist 退場條件與 Review Rule，本 phase 沒有做。

## 7. O4：`[skip-*]` 豁免標記（契約問題，不判違規）

| 面向 | 現況（verified，通則） |
| --- | --- |
| 作用範圍 | 單一 commit：標記必須獨占一行，寫在該 commit 的 message 裡 |
| 授權方式 | commit 作者自己寫；沒有第二人核可。例外：`[skip-go-migration]` 依 scripts/README 須附 migration plan |
| 失效時點 | 隨該 commit 結束，不會延續到下一個 commit。push replay 是否也認得這些標記，尚未確認（observed） |
| 數量 | 28 種 |

allowlist 是一份持續存在、要逐步清空的清單；skip 是單次的作者聲明，兩者不同。契約目前只定義了 allowlist 的退場條件。是否要另外定義單次豁免，例如必須附理由，或是要計數與觀察，留待 Phase 4 裁決。

## 8. 後續工作（只記錄，本 phase 不實作）

| # | 項目 | 類型 |
| --- | --- | --- |
| F1 | **已完成（2026-11-23，見 [`f1-repair-slice.md`](f1-repair-slice.md)）**。修正 L1：覆蓋集合改從「hook 實際註冊或呼叫了哪些檢查」建立。若只是擴大成掃全部檔案再加上 `warn*`，換個命名方式仍會漏掉。那 8 個 symbol 要逐一判定是 enforcement executor、advisory warning 還是 helper，再決定怎麼登記。T4 executor 方向的樣本一併補上 | 鎖正確性（Go 改動 + 兩個 commit 的 bin 流程） |
| F2 | **已完成（2026-11-24，見 [`f2-repair-slice.md`](f2-repair-slice.md)）**；做法改為宣告式命令表加狀態比對。原寫法：由 dispatch 產生 command-contract 命令表，消除 T3（含 F1 切片發現的誤報：diff 比對命中 test fixture 字串，見 [`f1-repair-slice.md`](f1-repair-slice.md) R3） | 結構消除 |
| F3 | KGE 呼叫鏈與 registry 命名（O2） | 結構，依切片延後 |
| F4 | **已完成（隨 F1）**。修正 `cli-modification-policy.yaml` 的 bin 重建說明。**跟下一次 Go source 修改一起處理**，讓 canonical policy 和已確認的兩個 commit 流程一致 | 文件 drift（runtime YAML） |
| F5 | 移除或對齊 `hookAdapterContent`（O8） | 死碼 / drift |
| F6 | T1 的 forbidden 補上無副檔名的 shebang 腳本 | 鎖缺口（G3 型） |
| F7 | 審查產品不變式鎖（§6） | 後續 phase 或 plan |
| F8 | 低優先：plan-evidence 的行號引用警告會把 finding ID（例：`L1`）誤判成行號。補對照樣本：真正的行號引用要提醒，finding ID 不能誤報 | 鎖誤報（advisory） |

← [Evidence index](README.md)
