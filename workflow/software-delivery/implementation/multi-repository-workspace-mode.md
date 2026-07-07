# Multi-Repository Workspace Mode（多倉庫處理模式）

> **Cognitive Slice**：`sd-implementation` 子面（orchestrator git + sibling application repos + project workflow overlay + integration adapter slice）。  
> **Consumer**：專案透過 `<PROJECT_ROOT>/docs/*-development-workflow.yaml`、`.ai-skill/project/sibling-repos.json` 與 `.ai-skill/project/rules/` 落地；本檔為 **canonical 模式**，不複製專案路徑進 application git。

| slice 欄位 | 值 |
|---|---|
| `id` | `sd-implementation.multi-repository-workspace-mode` |
| `purpose` | 當 workspace 含 **外層 orchestrator git**（plan / 驗收 / workflow）+ **一個或多個 sibling application git** 時，規定 Plan-First、manifest、commit sanitization、隊友 merge、integration slice 的實施順序 |
| `type` | `execution` |
| `tags` | implementation, multi-repo, sibling-repos, integration-adapter, plan-first, workspace-orchestrator |
| `load_when` | 多倉庫 workspace；`sibling-repos.json` 登記的應用 repo；隊友已 merge 業務面；本團隊接 sync / adapter / 驗收；ninej2 / SPI / webhook 類整合 |
| `do_not_load_when` | 單倉 greenfield；純外層文件；純 consolidation backfill（見專案 overlay consolidation hub） |
| `owner_layer` | workflow |
| `dependencies` | `sd-intake` §Plan-First、`sd-test-strategy`、`sd-validation`、`sd-closure` |

**範圍說明**：orchestrator + **1** 個 sibling 是多倉庫模式的 **N=1 特例**，不是另一套「雙倉庫模式」。一律讀本檔。

## 1. 多倉庫分工（N ≥ 1 sibling）

| 層 | 典型位置 | 進哪個 git | 職責 |
| --- | --- | --- | --- |
| **外層 orchestrator** | `docs/plans/`、`docs/features/`、`tests/`、`.ai-skill/project/` | workspace 根 repo | Plan artifact、驗收規格、workflow YAML、overlay rules、BDD / Java 可執行測試 |
| **Sibling application #1..N** | 各 repo 根（例：`browserManage/manageCode/`、`video-platform/`、`server_doc/`） | 各自獨立 git | 應用碼、SQL migration、內層 `Data/docs/<feature>.md`、內層 `.cursor/rules/` |

**Manifest（專案 overlay）**

- 登記：`<PROJECT_ROOT>/.ai-skill/project/sibling-repos.json` — `repos` 陣列列出 workspace 下所有 sibling 目錄名。
- 新增 repo：改 manifest → `bash .ai-skill/project/scripts/install-sibling-repo-githooks.sh` → README 一行 **本 repository 提交規範**。
- 參考實作：Vidoe-Test（5 siblings）、Brower（目前 1 sibling：`browserManage`）。

**規則**

- Plan **只**在外層 `docs/plans/active/`（advisory Plan-First，見 [`intake.md`](../intake.md) §Plan-First Ordering）。
- Sibling markdown **不得**引用外層 `docs/plans/`、`tests/bdd`、`tests/integration`、workspace 名稱、`.ai-skill/`（專案 commit sanitization）。
- Sibling 功能定義：各 repo 內 `Data/docs/<feature-slug>.md`（或專案約定路徑）。
- Sibling **可**互引（Vidoe-Test：`video-platform` ↔ `server_doc`）；**不可**引 orchestrator 外層路徑。

專案 overlay 範例：

| 專案 | orchestrator workflow | sibling manifest |
| --- | --- | --- |
| Brower | `docs/browser-manage-development-workflow.yaml` | `browserManage` |
| Vidoe-Test | `docs/framework-development-workflow.yaml` | `video-platform`, `server_doc`, … |

## 2. Commit sanitization（跨 sibling 一致）

| 元件 | 位置 | 說明 |
| --- | --- | --- |
| Template hooks | `.ai-skill/project/templates/sibling-repo-githooks/` | workspace only，不 commit 進 sibling remote |
| 安裝腳本 | `.ai-skill/project/scripts/install-sibling-repo-githooks.sh` | rsync template → 各 sibling `.githooks/` |
| 規則正文 | `.ai-skill/project/rules/sibling-repos-boundary.md` | manifest + forbidden 列表 |
| Cursor 委派 | `.ai-skill/project/rules/check-sibling-repo-before-commit.py` | 讀 manifest，委派到目標 sibling `.githooks/` |
| Opt-out | `COMMIT_SANITIZATION_SKIP=1` | 緊急用 |

每個 sibling：`.githooks/` 在 `.gitignore`；`core.hooksPath` 指向本機 hooks。

## 3. 兩種工作模式

| 模式 | 何時 | Sibling application | 外層 plan / tests |
| --- | --- | --- | --- |
| **自行開發**（目前預設） | 本團隊端到端功能 | 可改；`feature/<slug>` 本地 commit | 同 feature 分支補 plan + 驗收 |
| **隊友 merge + 整合切片** | 他人已 merge 業務面到 sibling 預設分支 | **只接** integration（sync SDK、adapter 掛接） | 修訂 plan 登記 evidence commit；補驗收 |

隊友 merge 時 **不**假設 plan 已更新：先 **T1 核對 commit**（`git show <sha>`），再 **T2 修訂外層 plan**，再 Execute。

## 4. Integration adapter 切片（典型：9j2 sync）

當業務面（Part A）已存在，本團隊交付出站 / 入站同步（Part B）：

```text
Discover → Interrogate → Draft/Revise Plan（外層 sub-plan）
    ⟲ Preflight（commit 對照、模組邊界、parity §）
Execute（sibling feature/*）
    → common：NineJ2ModuleSync adapter
    → admin/merchant：Service 掛接點（如 approve 後 push）
Validate
    → sibling：mvn test / 專案約定
    → orchestrator：tests/backend-java/、tests/bdd/、tests/integration/
Closure → 修訂 plan checkbox、sibling Data/docs、preflight log
```

| 步驟 | 動作 | 驗收 |
| --- | --- | --- |
| 0 | 外層 sub-plan 標明 `teammate_commits`、`sync_policy`、`inner_feature_doc` | plan frontmatter 可讀 |
| 1 | 核對預設分支上 Part A（表、API、UI） | `git log` / sibling feature doc |
| 2 | Adapter 骨架（若未有） | `NineJ2ModuleSync` + unit test |
| 3 | **掛接點**（如 `AdminXxxService.approve`） | mock SPI → `remote_id` |
| 4 | sibling doc §出站 / 入站分工更新 | 與程式一致 |
| 5 | 外層 plan Step 勾選 + parity 阻礙項更新 | sub-plan `status` |

**Policy 對照**

| Policy | 掛接時機 | 範例 |
| --- | --- | --- |
| `PUSH_ON_GATE` | Admin approve 後 | app-url、common-url |
| `INSTANT_PUSH` | 本地持久化成功後 | bookmark |
| `PULL_ON_DEMAND` / 入站按鈕 | Admin 手動 action | app-url `sync-remote` |

## 5. Feature branch 生命週期（sibling）

與專案 overlay 對齊（純本地預設不上傳）：

| 階段 | Sibling | Orchestrator |
| --- | --- | --- |
| 開工 | `git checkout -b feature/<slug>` from 預設分支 | 同 slug feature branch |
| 開發 | application commits | plan / tests commits |
| 驗收 | 專案約定單元 / 整合測試 | orchestrator `tests/` |
| 上傳（使用者明確要求） | merge → 預設分支 → push | merge → `main` → push |

Agent：**不**主動 push / merge 到 protected 分支，除非使用者明確要求。

## 6. 同工作階段閉環（本切片）

Implementation 完成時 **同一批次** 移動：

| 產物 | 位置 |
| --- | --- |
| Adapter + Service 掛接 | sibling repo |
| `Data/docs/<feature>.md` §sync | sibling |
| sub-plan Steps / status | orchestrator |
| `preflight-feedback-log.md`（若有 T0→T2） | orchestrator |
| 驗收（自行開發） | orchestrator `tests/` |

缺少 plan checkbox 或 sibling doc 同步 → 不得宣稱 sub-plan 完成。

## 7. 與 Plan-First / Preflight 的關係

- **Plan-First**：integration 任務仍須外層 sub-plan；隊友 commit 用 `teammate_commits` / plan revision 登記。
- **Preflight T1**：`git show` 驗證模組歸屬（例：app-url commit ≠ common-url endpoint）。
- **Preflight T2**：修訂 plan 後再寫程式。

## 8. 專案引用

專案在 `docs/*-development-workflow.yaml` 的 `ai_skill_workflow_slices` 與 `loading_surfaces` 應使用：

```yaml
multi_repository_workspace:
  source: "Ai-skill workflow/software-delivery/implementation/multi-repository-workspace-mode.md"
  load_when:
    - workspace orchestrator git + sibling-repos.json manifest
    - integration adapter or sync SDK slice after teammate merge
```

專案 overlay **只使用** `multi_repository_workspace`；不再提供 `dual_repository_integration` 鍵。

## Plan revisions

| 日期 | 變更 |
| --- | --- |
| 2026-07-07 | 初稿：多倉庫處理模式；sibling-repos manifest、跨 repo sanitization、Vidoe-Test / Brower 參考 |
| 2026-07-07 | 整合：移除雙倉庫獨立模式與 `dual_repository_integration` YAML 鍵；N=1 併入多倉庫；檔名 `multi-repository-workspace-mode.md` |
