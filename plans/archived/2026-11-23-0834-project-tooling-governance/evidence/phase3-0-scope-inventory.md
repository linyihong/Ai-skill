# Phase 3.0 — Ai-skill 工具樹範圍盤點（唯讀）

- **日期**：2026-11-23
- **基準**：`9a4e93ed`
- **範圍**：`scripts/`、`.githooks/`、[`enforcement/enforcement-registry.yaml`](../../../../enforcement/enforcement-registry.yaml)
- **性質**：唯讀盤點，沒有修改任何工具。發現事項都標為 `observed`，在後續步驟驗證前不視為違規。

## 1. 工具樹規模

| 區塊 | 檔案（含 test） | 非 test 行數 | 角色概況 |
| --- | --- | --- | --- |
| `cmd/ai-skill` | 1 | 11 | Entry，只呼叫 `app.Run` |
| `cmd/releasebuild` | 2 | 118 | 獨立 binary，不經 dispatch（可被找到 ≠ 可被路由） |
| `cmd/evidencecandidate` | 2 | 224 | 獨立 binary，不經 dispatch |
| `internal/app` | 131（74 test） | 24,715 | dispatch 加上全部命令的 Logic |
| `internal/{audit,glossary,pathutil,planvalidate}` | 22 | 1,863 | 領域模組 |
| `portable/kge` | 38（4 test） | 3,933 | 可直接複製給其他專案的驗證引擎，README 明文規定不得呼叫 git 或 hook |
| `.githooks/*` | 4 個 hook | 各 21–22 行 | shell adapter：選擇對應平台的 binary，再轉呼叫 `ai-skill hooks run` |
| `scripts/git-hooks/*` | 4 個 hook | 各 5 行 | 舊路徑，只轉呼叫 `.githooks/` |

`internal/app` 最大的檔案：`hooks.go` 3,699 行、`runtime.go` 3,249 行、`enforcement.go` 1,214 行、`runtime_compiler.go` 1,198 行、`discovery.go` 1,080 行。

## 2. 入口與能力

- **dispatch**：`app.go` 有 16 個頂層命令，只做 switch 分派：version、doctor、init-project、goals、close-loop、hooks、sync-cursor-bundle、runtime、roo、copilot、glossary、scan-checkboxes、enforcement、plans、kge、router、image-processing。另有 2 個獨立的 `cmd/` binary。
- **子命令**（`case` 數，粗略）：hooks 26、enforcement 19、runtime 13、plans 10、goals 8、kge 7。
- **既有的 catalog 候選**：三份文件各涵蓋一部分，都沒有 placement、state、owner 欄位。
  - [`docs/command-contract.md`](../../../../scripts/ai-skill-cli/docs/command-contract.md)：命令表
  - [`scripts/README.md`](../../../../scripts/README.md)：舊腳本到 Go 的遷移對照表
  - enforcement registry：鎖與 executor 的綁定表

## 3. 鎖（validator）

| Hook 階段 | 數量 | 說明 |
| --- | --- | --- |
| pre-commit | 10 | author email、diff whitespace、runtime compile/validate、no-new-shell-scripts、sanitization（block）與 incident score（warn）、review-architecture drift、canonical ownership drift、release binary parity |
| commit-msg | 32 | 多數是 `validateX → runKGEX → kge.Rule` 三段；Cognitive Mode、plan-tree、plan sync、feedback closure 等 |
| pre-push | 4 | push governance replay、KGE check、Go test preflight、GitHub workflow history |
| registry | 44 個 rule entry | `coverage`、`executors`（檔案 + symbol + hook 階段 + block/warn）、`rationale`、`verification_status` |

**初步分類**（正式結果以 Lock review 為準）

- **工具治理鎖**：約 7 把
  - no-new-shell-scripts（layout）
  - release binary parity（drift）
  - sanitization 與 author email（safety）
  - cli_doc_sync（drift）
  - enforcement registry sync / transition（drift）
  - CLI preflight 與 workflow history（completeness）
- **其他**：大多是產品不變式鎖，保護對象是 governance 文件、plan、Cognitive 契約與 runtime 資料。

## 4. Go-first 與 bin 重建（確認）

- **政策**：[`scripts/README.md`](../../../../scripts/README.md) §Go-first，以及 [`runtime/cli-modification-policy.yaml`](../../../../runtime/cli-modification-policy.yaml) 中的 `gate.cli.no_new_shell_scripts` 和 `gate.cli.bin_rebuild_for_HEAD`。
- **pre-commit**：`validateNoNewShellScripts` 擋下新增（Added）的非 Go 腳本；`validateReleaseBinaryParityStaged` 規定 CLI source 有變更時，5 個 binary、`BUILDINFO`、`SHA256SUMS` 必須在同一個 commit 一起 stage，可用 `[skip-release-binary-parity]` 豁免。
- **pre-push**：`TestRepoLocalBinariesMatchChecksumsAndCurrentSource` 比對的是「最後一個動到 CLI source 的 commit」，不是 HEAD。
- **待驗證**：pre-commit 要求「同一 commit」，以往經驗則是「之後另開一個只含 bin 的 commit」，兩者不一致。Phase 3 第一次改 source 時要先驗證實際流程，不預設哪一種正確。

## 5. 觀察事項（observed，尚未驗證）

| # | 觀察 | 相關契約 |
| --- | --- | --- |
| O1 | `hooks.go` 同時放 hook 編排和 27 個 validator 的 Logic；依 A-2，編排本身含條件決策，算 Logic，所以不是 Role 違規。問題在 discovery：同一個檔案裡有多種能力 | discovery、catalog |
| O2 | `validateX` 只是轉呼叫 `runKGEX`（例：`validateTokenBudget` 只有一行），形成三段間接呼叫。registry 以 `validate*` symbol 登記 executor，這可能就是 wrapper 存在的原因，屬於補償結構候選 | Review Rule |
| O3 | `scripts/git-hooks/*` 和 `.githooks/*` 是兩套重複的入口外殼；`scripts/README.md` 仍把前者寫成「最小 adapter」，與實際不符 | lifecycle（應為 deprecated 並寫明移除條件）、drift |
| O4 | 共有 28 種 `[skip-*]` 豁免標記，都沒有退場條件。單次 commit 的豁免算不算 allowlist，契約沒有定義 | lock-taxonomy（契約歧義候選） |
| O5 | registry 只有 `rationale`（約等於 intent），沒有 scope、forbidden、allowlist；`verification_status` 已追蹤 scenario 與 regression，可對應雙向樣本 | lock-taxonomy 四欄 |
| O6 | 治理檢查架構已有一部分做好：registry 對應 check registry、`hooks run` 對應單一入口、`Check` struct 對應 result contract。scope registry 還沒有 | lock-taxonomy §Governance check architecture |

## 6. 切片提案（2026-11-23 已裁決：S1 + 五點邊界）

| 方案 | 內容 | 評估 |
| --- | --- | --- |
| **S1（建議）**：hook pipeline 切片 | catalog 列全部 16 個頂層命令和 2 個獨立 binary，共 18 列；Role 只檢查 `.githooks/*`、`scripts/git-hooks/*`、`cmd/ai-skill`、`app.go` 與 hooks 的編排；Lock review 對工具治理鎖做完整四欄、雙向樣本和 Review Rule，產品不變式鎖只分類計數 | 涵蓋 O2–O6，改動量可控 |
| S2：整個工具樹 | 檢查全部 24.7k 行的 Role，並審查全部 46 把鎖的四欄 | 超出一個 phase 的合理範圍，也違反「不擴大實作」 |
| S3：只做 catalog | 不做 Role 和 Lock review | 無法驗證 Phase 3 的 Acceptance |

**S1 不包含**：`runtime.go` 等大型 Logic 檔的拆分、registry schema 變更（不新增欄位，四欄只寫在 evidence）、KGE 三段間接呼叫的重構。這些只記錄為後續工作。

### 裁決後的邊界（2026-11-23）

1. **產品不變式鎖**：全部分類登記，但不算已審查，明列「未審、後續處理」。只有工具治理鎖做完整審查。
2. **對帳**：先建立 hook → executor → registry 對照。§3 的 46 和 44 計數單位不同，「約 7 把」只是初步數字，要等對帳後才能確定。
3. **O2、O3**：Role 是否通過，和入口是否重複，分開判斷。兩套外殼可能各自都是薄入口；一行轉呼叫本身不算違規。
4. **O4**：維持為契約問題。單次 skip 的有效期間和 allowlist 的退場條件是不同問題，先記錄作用範圍、授權方式與失效時點。
5. **bin 流程**：§4 的待驗證項，要在第一次修改 Go source 前，依現行 hook 實作和建置文件確認。
6. **O6**：只是既有架構的盤點結果，不承諾「只補 scope registry 就完整」。

← [Evidence index](README.md)
