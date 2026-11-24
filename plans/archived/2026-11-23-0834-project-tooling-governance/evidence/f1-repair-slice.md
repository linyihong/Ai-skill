# F1 修復切片：hook 實際呼叫的檢查必須全部登記

- **日期**：2026-11-23
- **基準**：`0e8c95d1`
- **開放**：使用者 2026-11-23「開 F1 修復切片」。範圍沿用 [`phase3-application.md`](phase3-application.md) §8 的 F1；F4 依先前裁決隨這次 Go 修改一起處理。
- **登記方式裁決**（使用者 2026-11-23）：沒有合適既有 rule_class 的 meta gate，新增一個 `enforcement_self_governance` class 收納；rule_class 數超過 hard_limit 的問題只記錄為觀察。

## 1. 修正內容

**L1（已修正）**：`lintOrphanExecutors` 原本只掃 `hooks.go` 內 `run*` / `validate*` 開頭的宣告。現在另外建立 hook-call coverage：

- **根**：`runPreCommitHook`、`runCommitMsgHook`、`runPrePushHook` 三個 git hook 入口，加上 commit-msg 的 `commitMsgValidatorRegistry`。
- **集合**：用 Go AST 解析 `internal/app` 的非 test 原始碼，收集根直接呼叫、且在 package 內有宣告的函式。**不依命名**：`warnX`、`attachX`、`checkX` 只要被 hook 入口呼叫，就會進入集合。
- **判定**：集合中的每個 symbol 都必須綁定到某個 rule_class 的 executor，或列入 `internal_helper_allowlist`；否則回報 `orphan_executor`（FAIL）。
- **防止集合悄悄縮成零**：根被改名或刪除時回報 `hook_call_coverage_root_missing`（FAIL）；原始碼無法解析時回報 `hook_call_coverage_unavailable`（FAIL）。
- **`--registry` shadow**：`lintWithOverride` 原本只複製 `hooks.go`，現在複製 `internal/app` 全部非 test 原始碼，否則宣告在其他檔案的 callee 會解析不到。

原本的 `hooks.go` 命名掃描保留（兩者取聯集），避免既有覆蓋退步。

## 2. 分類：hook 呼叫但未登記的 17 個 symbol

hook-call coverage 在真實 repo 找到 17 個未登記 symbol：Phase 3 已知的 8 個，加上這次才浮現的 9 個。

**Phase 3 已知的 8 個**

| symbol | 呼叫點 | 分類 | 登記到 |
| --- | --- | --- | --- |
| `validateReviewArchitectureNavigationDriftStaged` | pre-commit | enforcement executor（block） | `capability_context_governance` |
| `validatePlanEvidenceConvention` | commit-msg registry | enforcement executor（block） | `plan_governance`（source 加 `governance/lifecycle/plan-evidence.md`） |
| `validateEnforcementRegistryTransition` | commit-msg registry | enforcement executor（block） | `enforcement_self_governance`（新） |
| `validateEnforcementRuleRegistrySync` | commit-msg registry | enforcement executor（block） | `enforcement_self_governance`（新） |
| `validatePushGovernanceReplay` | pre-push | enforcement executor（block） | `enforcement_self_governance`（新） |
| `warnSanitizationIncidentScore` | pre-commit | advisory warning（warn） | `sanitization` |
| `warnPlanArchivalLinkIntegrity` | commit-msg 直接呼叫 | advisory warning（warn） | `plan_archival_link_integrity` |
| `warnPlanEvidenceLineNumberCitations` | commit-msg 直接呼叫 | advisory warning（warn） | `plan_governance` |

8 個中沒有 helper：5 個會擋下 commit 或 push，3 個是只提醒的 advisory。

**新浮現的 9 個**

| symbol | 呼叫點 | 分類 | 登記到 |
| --- | --- | --- | --- |
| `attachKGECheck` | pre-push | enforcement executor（block） | `enforcement_self_governance`（新） |
| `attachKGEAdvisoryCount` | commit-msg | advisory warning（warn） | `enforcement_self_governance`（新） |
| `githubWorkflowHistoryCheck` | pre-push | advisory warning（warn） | `github_workflow_history`（原本只綁 `runPrePushHook` 並在 notes 寫「wraps」） |
| `planValidateShadowCheck` | commit-msg | advisory warning（warn，不影響 exit code） | `plan_tree_governance` |
| `Run` | pre-commit | helper | allowlist（CLI 重入，用來跑 `runtime compile/validate`；擋下的是 hook 本身） |
| `compactSummary` | pre-commit / pre-push | helper | allowlist（輸出格式） |
| `canonicalOwnershipDriftOptOut` | pre-commit | helper | allowlist（`[skip-*]` 標記解析） |
| `releaseBinaryParityOptOut` | pre-commit | helper | allowlist（同上） |
| `reviewArchitectureDocDriftOptOut` | pre-commit | helper | allowlist（同上） |

**判定說明**：

- `validatePushGovernanceReplay` 和已在 allowlist 的 `validateCommitMessageGovernance` 一樣，都是轉呼叫其他檢查。但 pre-push 只有 `runPrePushHook` 登記在 `github_workflow_history`（warn），replay 的 block 行為在 registry 中沒有任何對應，所以判為 enforcement executor，不判為 helper。
- `planValidateShadowCheck` 檢查的是兩套 plan-tree 引擎是否一致，屬於工具 parity 的觀察；它會發出 warning，所以歸 advisory，掛在它鏡像的 `plan_tree_governance` 底下。

## 3. 樣本狀態（Phase 3 §5 的三類標示）

| 項目 | 狀態 | 證據 |
| --- | --- | --- |
| T4 rule 方向 | 通過 | 原有測試（Phase 3 已跑） |
| T4 executor 方向（違規變體） | 通過 | `TestLintHookCallCoverage_UnboundCalleeDetectedRegardlessOfName`：沒有 `run`/`validate` 前綴的 `checkFoo`、宣告在其他檔案並經 registry literal 呼叫的 `validateBar`，都會被抓到 |
| T4 executor 方向（合法相似寫法） | 通過 | `TestLintHookCallCoverage_LegalLookAlikesNotFlagged`：已綁定、已列 allowlist、只被非根函式呼叫、stdlib selector 呼叫、根本身、只在 test 檔宣告的函式，都不會誤報 |
| 根消失 | 通過 | `TestLintHookCallCoverage_RootMissingFails` |
| 真實 registry | 通過 | `TestLintHookCallCoverage_RealRepo`：coverage 涵蓋全部 8 個 symbol，且沒有 orphan、root-missing、unavailable 任何一種 finding；`ai-skill enforcement lint` 結果為 0 FAIL |
| 突變驗證 | 通過 | 停用 `lintHookCallCoverage` 後，違規變體和根消失兩個測試都失敗，恢復後通過。`RealRepo` 在突變下仍會通過，因為它鎖的是 registry 綁定，不是 coverage 偵測，這點屬預期 |

`go test ./...` 全部通過（基準 `0e8c95d1` 加上本切片的工作區修改）。

## 4. F4（隨本次 Go 修改處理）

[`runtime/cli-modification-policy.yaml`](../../../../runtime/cli-modification-policy.yaml) 的 `gate.cli.bin_rebuild_for_HEAD` 已改成兩個 commit 的流程：先提交 source 加完整 bin（pre-commit），再提交只含 bin 的重建、`--commit` 指向 source commit（pre-push），`enforcement_layer` 改為 pre-commit + pre-push。內容與 [`scripts/ai-skill-cli/README.md`](../../../../scripts/ai-skill-cli/README.md) 以及 Phase 3 §1 已確認的流程一致。

## 5. 觀察（observed，本切片不處理）

| # | 觀察 |
| --- | --- |
| R1 | registry 實際有 46 個 rule_class（新增前為 45），超過 `governance_thresholds.rule_class_hard_limit` 40；`current_count` 欄位仍寫 40。依裁決只記錄，另案處理 |
| R2 | `executor_kind_spec.enum` 只列 4 種，但 registry 實際還用了 `pre_commit_validator`、`pre_commit_staged_content_scanner`、`runtime_validator` 等，lint 沒有檢查 enum。本切片沿用相鄰項目已在用的 kind，沒有新增 kind |
| R3 | T3（`cli_doc_sync`）出現誤報：它比對 staged diff 的字串 `+func runPreCommitHook`，結果命中 `_test.go` 裡 fixture 原始碼字串，而且不排除 test 檔。這是 T3 在合法相似寫法方向的新缺口（Phase 3 只記錄了漏抓方向），併入 F2。本次 commit 沒有改任何命令契約，用 `[skip-cli-doc-sync]` 豁免 |

← [Evidence index](README.md)
