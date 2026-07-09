# 2d′ — ExternalRepoC 9j2 module alignment follow-on（2026-07-09，證據 only）

> **專案證據邊界**：inner commit、class 名、live 環境細節留於 `<PROJECT_ROOT>` active main plan §執行紀錄；Ai-skill 只保留 generalized dogfood metrics（依 [`enforcement/sanitization.md`](../../../../enforcement/sanitization.md)）。**本 run 為 §2d 同一 consumer（ExternalRepoC）的延續**，非新 consumer、非 Phase 3 closure。

## Run 摘要

- **任務**：9j2 出站同步平台 — 模組 **01 app-url** / **02 bookmark** M1–M8 對齊、sync-remote 外層驗收、雙 feature branch 合併至可部署分支、live 刪除語義與測試落點治理。
- **Transport**：Cursor orchestrator + **Task subagent**（Executor / Verifier，多 slice）；部分對齊與外層 acceptance 由 orchestrator 直接寫 `<PROJECT_ROOT>`（hook allowlist 內）。
- **Repo**：`<PROJECT_ROOT>` 外層 + `<INNER_REPO>/manageCode` 內層；雙 repo commit。
- **Consumer overlay**：沿用 2d — `plan-delegation-execution-loop.md`、verification backfill、`slice_kind`、C1–C5；新增 project rules：`pre-push-build-gate`、`test-acceptance-placement`、`9j2-sync-module-alignment`。
- **相關 kit 章節**：[`01-dogfood-prompt-kit.md`](../01-dogfood-prompt-kit.md) → `### 2d`（baseline consumer）

## 相對 §2d 的新信號

| # | 觀察 | 對本 plan 的意義 |
|---|---|---|
| 1 | **平行 feature branch 下 user-visible 已關閉但 UI 不可見**（sync-remote 在未合併分支；可部署分支缺按鈕） | 延伸 2d #1：`slice_compliant_closed` 若含 user-visible 行為，brief 應加 **integration gate**（合併至可部署分支）或標 `beyond-loop: merge`；否則 orchestrator 與使用者驗收脫節 |
| 2 | **遠端已缺席時 delete 非 idempotent**（遠端「未找到」阻本地刪） | L3 / `verifier_only`：`remote_absent_delete` 應進 backfill；mock IT 綠燈仍可能漏 — ERA「哪種證據能關哪個狀態」 |
| 3 | **Live IT teardown 須雙邊**（本地 DB + 遠端清 test-data marker 列） | backfill 應明示 `live_delete_policy` / teardown owner（executor vs orchestrator beyond-loop） |
| 4 | **首輪模組對齊曾跳過 formal Verifier**，使用者糾正後補跑 | 反模式實例：doc/align 回合仍須 V1–V4 或書面 `defer`；僅 executor 自驗 ≠ loop 關閉 |
| 5 | **inner `src/test` commit block 已機械化**（新測試僅 outer `tests/`） | 相對 2g「漸進遷移」：ExternalRepoC 選整路徑 deny + BDD gate；consumer 自理 gate 模式 ×2 家族內變體 |
| 6 | **pre-push build gate**（push 可部署分支前 client build + compile） | Q5：release-time consumer gate 與 orchestrator deny gate 互補；仍 doc-only |
| 7 | **V5 runtime smoke defer**（dev captcha in-memory） | 延伸 2d #6：runtime tier 須 `V5: defer(reason)`，不可默認已驗收 |
| 8 | **merge / push / stack restart 由 orchestrator，不在 Executor loop** | 延伸 deploy 邊界 → **integrate / release leg** 為 beyond-loop，但 user-visible closure 依賴它 |

## 量測欄

| 指標 | 值 |
|---|---|
| 模組 slice | **2**（01 + 02 對齊；含外層 acceptance 補強） |
| integration UX fail（合併前） | **1** — sync-remote UI 不可見直至 feature 合併 |
| acceptance-violation（live delete） | **1** — `remote_absent_delete`；fix 後 idempotent |
| verifier 降級（首輪跳過） | **1 次** — 使用者糾正後補 Verifier |
| orchestrator 越界寫 manageCode | **0**（gate 生效後）；idempotent delete fix 經 Executor |
| 新 consumer 機械 gate | **2** — pre-push build、inner src/test block |
| 模型自然落位（Phase 3 穩定性） | **是** — overlay + backfill + ERA 分工未改形狀 |

## 契約回饋（寫回 canonical / consumer overlay）

1. **`integration gate`** — user-visible `slice_compliant_closed` 應列合併至可部署分支，或標 `beyond-loop: merge`（advisory，模板 A acceptance 段）。
2. **`remote_absent_delete`** — 出站 sync 模組 backfill 建議列 `verifier_only` 負面 case（遠端 404 / 未找到仍允許本地刪）。
3. **`live_delete_policy`** — live tier teardown 雙邊 owner 寫入 backfill，避免 `[sync-test]` 殘留。
4. **Release-time gate** — pre-push build 與 commit-time placement gate 可並存於 consumer overlay，不必進 schema。
5. **Q5 仍維持 doc-only** — 2d′ 強化 consumer gate 證據，不改 schema promotion 判斷。
