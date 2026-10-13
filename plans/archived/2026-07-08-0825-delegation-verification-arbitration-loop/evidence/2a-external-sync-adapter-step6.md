# 2a-external — 外部 monorepo sync-adapter Step 6（2026-07-08，Cursor Task transport）

> **專案證據邊界**：sub-plan 檔名、class/test 名、remote host、live 環境細節留於 `<PROJECT_ROOT>` 被委派 sub-plan §8 execution records；Ai-skill 只保留 generalized dogfood metrics（依 [`enforcement/sanitization.md`](../../../../enforcement/sanitization.md) + [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md)）。

## Run 摘要

- **任務**：Admin 直寫 CRUD → remote outbound sync（delegated sub-plan Step 6）；含 mock 單測 + test-data-marker live IT。
- **Transport**：Cursor orchestrator + **Task subagent** executor/verifier（偏離 plan 原文「人類開全新 chat」— 同 2a demo 記為 transport adaptation；**外部 repo** 條件成立：`<PROJECT_ROOT>` 外層 + `<INNER_REPO>` Java 內層）。
- **Repo**：`<PROJECT_ROOT>`（plan/brief）+ `<INNER_REPO>/manageCode`（實作 submodule）。
- **Slices**：mock（6-1–6-4）→ live（6-5）；各一輪 executor + verifier。

## Executor / Verifier 軌跡

| Slice | Executor (Task) | Deliverable commits | Verifier (Task) | Outcome |
|---|---|---|---|---|
| Mock | `<TASK-id-m>` | `<HASH-a>`（orchestrator 越界，見下）、`<HASH-b>` | `<TASK-id-v-m>` | L1–L3 pass；acceptance-violation **0** |
| Live | `<TASK-id-e>` | `<HASH-c>` | `<TASK-id-v-e>` | L1 rerun pass；live IT pass |

## Brief（orchestrator 撰寫，摘要）

- **goal**：Admin create/update/delete/purge 在 `approved` + `remote_id` 時 outbound push；mock 覆蓋 + live 僅 test-data marker 列。
- **acceptance**：§6 條（create/update push、delete guard、unit tests、live IT、無 production sync-remote 污染）。
- **verification**：`mvn test` mock + live IT profile；verifier 獨立重跑 L1。
- **context.required**：admin service layer、outbound sync adapter、remote envelope helper、delegated sub-plan §6。

## 實作品質差集（相對「單 session 直接改 code」反事實）

| 項目 | 直接實作（本 session 早期） | Loop 後 |
|---|---|---|
| 單測 | 可能延後 | 5 mock tests（service-layer outbound sync suite） |
| Live 驗收 | 可能跳過 | 1 live IT；僅 test-data marker |
| 已知 bug 關閉 | — | remote-id null guard（`<HASH-b>`）；remote envelope 裸 SPI unwrap（`<HASH-c>`，否則 live create 假 502） |
| Plan 證據 | 無 | `<PROJECT_ROOT>` sub-plan §8 + verification_backfill |

## 仲裁紀錄（2026-07-08 / sync-adapter Step 6）

| finding | 處置 | 理由 | 後續 |
|---|---|---|---|
| O1 batchImport outbound 未補 service-layer test | defer | beyond mock acceptance 字面；batchImport 路徑已 wire | 若後續 slice 需覆蓋可開 observation |
| O2 purge* outbound 未單獨 mock | defer | purge 共用 delete push；delete test 已 cover guard | observation |
| O3 502 envelope 路徑未 mock | defer | live IT 已驗證 unwrap fix；mock 502 非 acceptance 必須 | observation |
| V1 orchestrator 在 loop 建立前直接改 admin service layer | defer | **Q1 負向信號**：`<HASH-a>` 由 orchestrator 手寫；後續 mechanical reminders（project overlay hooks、BDD gate）補上 | 記入越界指標；非 verifier finding |

## 量測欄

| 指標 | 值 |
|---|---|
| verifier 差集（verifier 抓到、executor 自驗沒抓到） | **acceptance-violation 0**；beyond-acceptance observation **3**（O1–O3）；implementation gap **2**（guard + envelope，均由 executor 在交付前測試/IT 暴露，verifier L1 確認） |
| verifier 降級（只跑 L1、未做 L2/L3？） | Mock：**否**（L1–L3）；Live：**L1 rerun only**（delta 小，合理） |
| 仲裁分佈 | fix **0** / defer **4** / reject **0** |
| orchestrator 越界（動手寫 code？被迫回讀 diff 細節？） | **有（1）** — loop 前 `<HASH-a>` 直接改 Java；loop 後 **未**回讀 diff 細節即可仲裁 |
| verifier 報告自足性 | **是** — 兩輪仲裁憑 verifier 交付 + commit hash；orchestrator 未為仲裁 grep 原始碼 |
| executor 讀檔差集 | **輕微** — live slice 讀 envelope helper（brief 已列） |
| 契約缺漏回饋 | brief 可明示「orchestrator 不得 pre-loop 直接 commit 實作」→ 已用 `<PROJECT_ROOT>` project overlay + sessionStart hook 機械化 |
| transport adaptation | Task subagent 代替人類全新 chat（同 2a demo） |
| **協調成本（量化）** | Task spawn **4**（2 executor + 2 verifier）；外層 plan commit **6**；內層 `<INNER_REPO>` commit **3** |
| **品質成本（量化）** | 新增 test **6**（5 mock + 1 live IT）；acceptance-violation **0**；live 安全約束 **100%**（僅 test-data marker） |

## 結論摘要（orchestrator 負擔 vs 品質 — 供 Q3）

| 維度 | 相對單 session 直接實作 | 證據 |
|---|---|---|
| Orchestrator **寫 code 負擔** | **↓**（loop 後為 0；loop 前 1 次越界） | `<HASH-a>` vs `<HASH-b>`/`<HASH-c>` 全由 executor |
| Orchestrator **協調/文件負擔** | **↑** | brief、spawn×4、§8 backfill、6 外層 commit |
| 交付品質 | **↑** | 6 tests、2 bug fix、0 acceptance-violation、可審計軌跡 |
| Verifier 邊際價值（本任務） | **中等** — 差集主要在「強制 IT + 結構化 defer」，非大量 acceptance catch | envelope/guard 由 executor 測試暴露；verifier 提供獨立 L1 與 observation 分類 |
| 適用邊界 | **多步 acceptance + live 驗收** 划算； trivial 單檔 fix 不划算 | 與 plan §Trade-offs advisory 一致 |
