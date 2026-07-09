# 2c — 外部 monorepo tiered archive platform（全線 A–D）（2026-07-08，Cursor Task transport）

> **專案證據邊界**：slice 級 inner commit、EC2 host、PG endpoint、class 名留於 `<PROJECT_ROOT>` tiered plan §12 + inner `Data/docs/tiered-data-archive.md`；Ai-skill 只保留 generalized dogfood metrics。

## Run 摘要

- **任務**：分層資料歸檔平台 Phase A / A′ / B / C / D 全線交付（8 slices）；首個 domain adapter = verification-code retention。
- **Transport**：Cursor orchestrator（主 session）+ **Task subagent** executor / verifier（每 slice 一輪；fix 項再 spawn）；**外部 repo** 成立（`<PROJECT_ROOT>` 外層 plan + `<INNER_REPO>/manageCode`）。
- **Repo**：`<PROJECT_ROOT>` `feature/<slug>` + `<INNER_REPO>` 同 branch。
- **Brief 來源**：外層 tiered plan §12 `delegation.brief` + `.ai-skill/project/rules/plan-delegation-execution-loop.md`（orchestrator **不讀** `manageCode/server/**`，僅仲裁 dispute）。

## Slice 軌跡（摘要）

| Slice | Phase | Executor deliverable | Verifier | Fix 輪 |
|---|---|---|---|---|
| 1 | H2 e2e | retention IT + export-failure | L1–L3 pass → **policy_snapshot** gap | fix + 重驗 |
| 2 | Doc align | inner/outer doc 對齊 | L1–L3 pass | — |
| 3 | Live PG | live PG retention IT | L1–L3 pass | — |
| 4 | Worker ops | health + live IT hardening + runbook | L1–L3 pass | — |
| 5 | A′ Sqlite | `SqliteArchiveWriter` | L1–L3 pass | — |
| 6 | B | per-domain warm.backend + dynamic cron | L1–L3 pass | — |
| 7 | C | `ColdArchivePromoter` + scheduler | L1–L3 pass | — |
| 8 | D | `ArchiveQueryRouter` + Admin history | **manifest 未登記** | fix + 重驗 |

## 仲裁紀錄（2026-07-08 / tiered archive 全線 — 代表性 findings）

| finding | 處置 | 理由 | 後續 |
|---|---|---|---|
| S1 policy_snapshot 未 assert | fix | acceptance-violation — e2e 缺 D4 meta-audit 欄位 | executor 補 assertion；重驗 pass |
| S8 Phase D migration 未入 manifest | fix | acceptance-violation — deploy 會漏跑 RBAC migration | manifest + down SQL；重驗 pass |
| Worker 僅模組存在、測試服未部署 | defer（Slice 4 前）→ 後續 ops 關閉 | 真實 gap；Slice 4 brief 含 live deploy 驗收 | test EC2 worker UP + cron 觸發紀錄 |
| Verifier INFO：verification plan §4.4.5 未同步 Phase D | defer | beyond-acceptance 文件債 | 轉 feature plan 修訂 |
| Orchestrator 早期 session 直接讀 manageCode | defer | 越界信號；採三角色紀律後 **0** 實作 diff | project overlay `plan-delegation-execution-loop.md` 機械化 |

## 量測欄（8-slice 彙總）

| 指標 | 值 |
|---|---|
| verifier 差集（acceptance-violation） | **2**（S1 policy_snapshot、S8 manifest）— 皆 `fix` 後重驗 pass |
| verifier 降級（只跑 L1？） | **否** — 各 slice 均 L1–L3；live deploy slice 含 L3 對抗性（cron 觸發、health） |
| 仲裁分佈（代表性 violation） | fix **2** / defer **3+** / reject **0** |
| orchestrator 越界（寫 code？回讀 diff？） | **Slice 模型後：無實作 diff**；早期規劃 session 有讀碼，記 defer；仲裁 **未** 為 dispute 外回讀 diff |
| verifier 報告自足性 | **是** — Slice 8 manifest gap 全憑 verifier 報告觸發 fix，orchestrator 未自行 grep manifest |
| Task spawn 成本 | **~18**（8 executor + 8 verifier + 2 fix 輪） |
| 外層 plan commit | **~12**（§12 backfill + brief）；內層 **~12** feature commits |
| 品質信號 | acceptance-violation **2/8 slices**（25%）；皆在 merge 前捕獲；test env deploy 驗證 scheduler 實際觸發 |
| vs 2a-external（單 Step 6） | 協調成本 **↑↑**（8×）；越界寫 code **↓**（slice 紀律後為 0）；verifier catch **↑**（多 phase 邊界） |

## 契約回饋（寫回 canonical / project overlay）

1. **L2/L3 價值在多 slice 平台計畫中放大** — manifest / meta-audit 欄位類問題 L1 全綠仍漏。
2. **Orchestrator outer-only commit** — 外層 §12 + brief 先 commit 再 spawn，軌跡可審計（C1–C5 合規關閉）。
3. **Task transport = 2a/2a-external 的 agent 化** — 不必人類開新 chat；fresh context 由 subagent 保證。
4. **多 slice 任務值得 loop**；單檔 typo 不值得 — 與 advisory 適用邊界一致。
5. **Schema promotion 仍不建議（證據累積中）** — 2c 支持 doc-only 延續；Phase 3 Q5 決策待後續收斂（見 `_plan.md` §Phase 3）。
