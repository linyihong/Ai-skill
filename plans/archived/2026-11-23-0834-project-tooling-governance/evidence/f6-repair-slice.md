# F6 修復切片：T1 補上 shebang 腳本，pre-push replay 改用同一條規則

- **日期**：2026-11-24
- **基準**：`182d2867`
- **開放**：使用者 2026-11-24「開 F6 修復切片」。範圍：[`phase3-application.md`](phase3-application.md) §8 F6。T1（no-new-shell-scripts，layout 族）原本只比對副檔名，沒有副檔名、帶 shebang 的可執行腳本抓不到。

## 1. 盤點時另外發現的問題

T1 有兩個執行點，守的是同一個不變式（Phase 3 已記錄為 G6 第 2 次），但 forbidden 範圍不同：

| 執行點 | forbidden（F6 之前） |
| --- | --- |
| pre-commit `validateNoNewShellScripts` → portable `NoNewShellScriptsRule` | 12 種副檔名 |
| pre-push replay `validateNewShellScriptsInCommit` | **只有 `.sh`** |

所以用 `--no-verify` 提交 `.py`、`.js` 等檔案時，pre-push replay 也攔不住，而 replay 本來就是用來補這個漏洞的。這和 F6 是同一把鎖的同一類缺陷（forbidden 不完整），一併修正。

## 2. 修正內容

- **內容偵測**（[`portable/kge/rule_no_new_shell_scripts.go`](../../../../scripts/ai-skill-cli/portable/kge/rule_no_new_shell_scripts.go)）：新增的檔案若內容以 `#!` 開頭，不論檔名都視為腳本入口，錯誤訊息標示 `(shebang)`。
  - 只在 adapter 提供 `CapStagedContent` 時生效。這個 capability 沒有列為 required：外部專案的 adapter 沒有提供內容時，規則仍照原本只比對副檔名，不會因為少了 capability 而整條跳過或報錯。這是 S2 的教訓，可攜規則包的行為改變不能預設外部都跟得上。
- **hook adapter 豁免**：shebang 檔案若**直接**位於 repo 設定的 `core.hooksPath` 底下，就是 git hook adapter，Go-first 政策允許它當薄入口外殼。
  - 豁免依據是 repo 設定（`git config core.hooksPath`），**不列舉檔名**，和 Phase 4 候選 T-B5 的方向一致。
  - 只豁免內容偵測；副檔名命中（例如 `.githooks/helper.sh`）任何位置都不豁免，維持原本的行為。
  - 只算直接子項目（git 只執行 hooks 目錄最上層的檔案）；`.githooks/sub/x` 和 `.githooksx/x` 都不在豁免範圍。
- **讀取方式**：pre-commit 讀 index 的 staged blob，replay 讀該 commit 的 blob。每個檔案只保留前 256 bytes，判斷 shebang 只需要第一行。
- **共用規則**：pre-commit 和 pre-push replay 都透過 `noNewScriptsContext` 建立 context、執行同一條 portable rule，forbidden 範圍只有一份。KGE workspace context 也帶入 `HookAdapterDirs`，讓 `kge check` 的判斷一致。
- **同步**：registry `shell_script_governance` 的 notes 與 rationale，`cli-modification-policy.yaml` 的 `gate.cli.no_new_shell_scripts`（description、detection_signals、enforcement_layer）。

## 3. 樣本狀態（沿用 Phase 3 §5 的分類）

| 項目 | Phase 3 | F6 之後 | 證據 |
| --- | --- | --- | --- |
| 無副檔名的 shebang 腳本（違規方向） | 缺口 | **通過** | `TestNoNewShellScriptsRule_Shebang`：`scripts/deploy`、`tools/gen`、巢狀的 `.githooks/sub/helper`、名稱相近的 `.githooksx/pre-commit` 都會被擋 |
| 合法相似寫法 | 通過（只測副檔名） | **通過** | 同一個 test：`core.hooksPath` 底下的 hook adapter、`Makefile`、`LICENSE`、`# 標題` 開頭的 markdown、Go 原始碼都不誤報 |
| 豁免不涵蓋副檔名 | — | **通過** | `.githooks/helper.sh` 仍被擋 |
| 沒有內容時降級 | — | **通過** | 沒有 `CapStagedContent` 時不產生 shebang finding |
| replay 與 pre-commit 一致 | 未發現（§1） | **通過** | `TestPushGovernanceReplayBlocksNonGoScriptsLikePreCommit`（`.py` 和 shebang 腳本以 `--no-verify` 提交後，replay 都會擋）；`TestPushGovernanceReplayAllowsHookAdapterInHooksPath` |
| pre-commit 讀 staged blob | — | **通過** | `TestRunKGENoNewShellScriptsDetectsStagedShebang` |

**突變驗證**（每次都還原）：

| # | 突變 | 結果 |
| --- | --- | --- |
| X1 | 關閉 shebang 判斷 | portable 與 app 共 3 個 test 失敗 |
| X2 | 移除 hooks 目錄豁免 | 2 個 test 失敗（hook adapter 被誤擋） |
| X3 | replay 改回只比對 `.sh` | replay 一致性 test 失敗 |

`go test ./...` 全部通過；`ai-skill enforcement lint` 為 0 FAIL。

**現有檔案掃描**：目前 tracked 檔案中，以 shebang 開頭、又沒有被副檔名規則涵蓋的，只有 `.githooks/*` 4 個 hook adapter 和 `scripts/git-hooks/*` 4 個 deprecated forwarder。這把鎖只看新增檔案，所以兩者都不受影響。之後若再新增 `scripts/git-hooks/*`，會被擋下，符合它已標為 deprecated 的狀態。

## 4. 觀察（observed，本切片不處理）

| # | 觀察 |
| --- | --- |
| V1 | **可攜規則包的行為改變**：外部專案的 adapter 若提供了 `CapStagedContent`，而且 `FileContents` 含新增檔案，但沒有設定 `HookAdapterDirs`，那麼它們新增的 hook adapter 會開始被擋。依 F2 審閱對 S2 的裁決，可攜層的影響要先盤點外部使用方式；本切片在 Ai-skill 內做了降級設計（沒有內容就不偵測），外部影響併入 S2 的唯讀盤點 |
| V2 | T1 的另一個執行點：`runtime validate` 的 `nativeToolBootstrapShellCLIDecisionValidation` 檢查既有腳本有沒有 Go 遷移決策，但**只看 `.sh`**，連 pre-commit 的 12 種副檔名都沒有涵蓋，也不看 shebang。這和 §1 是同一種「同一不變式、多份 forbidden」的問題，但守的是既有檔案（Phase 3 記錄的觸發點不同）。目前既有的 shebang 檔案只有 hook adapter 和 forwarder（§3）；有沒有既有的 `.py`、`.js` 等檔案缺少遷移決策，尚未盤點。本切片不處理 |

## 5. 審閱裁決（2026-11-24）

| # | 項目 | 裁決 |
| --- | --- | --- |
| A1 | F6 修正內容 | **通過**。四點都確認：降級設計（`CapStagedContent` 不列為 required）、豁免範圍（直接子項目，只豁免 shebang 偵測）、replay 和 pre-commit 共用同一條規則、測試強度（違規與合法樣本、降級、執行點一致性、X1–X3）。F6 不再擴大修復範圍 |
| A2 | 兩段 bin 流程 | 符合預期：`0fb5618c` 是 source 加完整 bin；`c4da45aa` 只含 bin，`BUILDINFO` 指向 source commit |
| A3 | V1 | 維持在 S2 的唯讀盤點範圍，暫不修復 |
| A4 | V2 | 維持觀察。**F6 統一 pre-commit 和 replay，不代表 T1 所有執行點都已統一**：`runtime validate` 是另一個 consumer，需要獨立的證據和裁決 |
| A5 | `f3fc5b30` 隨 F6 一起推送 | 審閱者判斷內容合理、可以保留（IL2CPP field-read lesson 和 failure pattern 都維持 candidate，沒有提前 promotion），但推送意圖要由使用者確認。使用者已在 2026-11-24 對話中確認：「f3fc5b30 是我的，沒問題」。保留在 main，不回退。內容品質和推送授權分開判斷 |
| A6 | Phase 4 | 不開放。F6 完成不代表 Phase 4 自動獲准 |

← [Evidence index](README.md)
