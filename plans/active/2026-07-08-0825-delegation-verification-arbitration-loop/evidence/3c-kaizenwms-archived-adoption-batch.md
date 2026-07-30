# 3c — KaizenWMS：完成計畫批次回饋收割（archived＋MES naming，2026-07-17–28）

> **專案證據邊界**：inner commit／路徑細節留 consumer；本檔只留 generalized metrics 與契約回饋。  
> **範圍**：`plans/archived/*` 與 MES launch 內已關閉、具三角色／仲裁痕跡的計畫；**不**重複寫入已由 [`2y`](2y-kaizenwms-phase2-spa-scaffold-c1b.md)／[`2z`](2z-kaizenwms-phase3-karma-stale-serve.md)／[`3a`](3a-kaizenwms-spawn-friction-skip-loop.md)／[`3b`](3b-kaizenwms-multi-round-arbitration-maturity.md) 覆蓋的同事件全文。  
> **目的**：把「完成計畫裡散落的回饋」收進本 dogfood 證據庫，供 Q3／Q5／Q8 觀察期引用。

## Run 摘要

| 欄 | 值 |
| --- | --- |
| **Harvest 觸發** | Stakeholder：檢視 `<PROJECT_ROOT>/plans` 完成計畫回饋 → 寫入 Ai-skill evidence |
| **Consumer** | KaizenWMS |
| **狀態** | 證據 only（批次對照，非單一 slice dogfood） |

## 批次對照表

| Consumer plan（泛化標題） | Loop 形態 | 關閉狀態 | 可沉澱回饋 |
| --- | --- | --- | --- |
| Shell nav groups／module isolation Phase 1 | 獨立 Task E＋V；fix→re-verify | `slice_compliant_closed` | 文件契約刀亦可跑完整 loop；F1 alias／路徑對齊屬 acceptance-violation；**已作 3a 正向對照** |
| List↔detail state restore Phase 1 | 同 session 直做、無獨立 V | **`implementation_done`**（audit 回填） | 可見 list／query 還原 ≠ surgical；事後誠實降級（同 3a F1） |
| List↔detail Phase 2 | Task E＋V；一條 isolation violation **fix** | `slice_compliant_closed` | Feature isolation 閘在 Verifier／pre-commit 抓跨 feature inject；fix 時 hook 擋 → Orchestrator same-session Executor（須註 transport） |
| QG1 contract reconciliation | E＋V | `slice_compliant_closed` | 早期正向：Verifier V1–V4＋無新 finding 關閉 |
| WMS material foundation Phase 4f1 等 | 多刀 E＋V；初驗 **fix** 後 targeted re-verify | `slice_compliant_closed` | Verifier 擋契約篩選／非原子過帳（對齊 2y C1b 精神：行為證據＞交付物） |
| MES product naming | V1 → **fix** BDD zh → V2 聚焦重驗 | `slice_compliant_closed` | 獨立 arbitration artifact 作 closure authority；並行 session 髒樹不得併入關閉 |
| Equipment standalone package | 多 evidence；slice 多為 `implementation_done` | 部分未升 closed | Observe／Edge 殘餘 → 後續 MES；**未**假稱全線合規 |
| SPA mobile Phase 1a（見 3b） | 同 session | `implementation_done` | 與 list-detail P1 **同構 process debt** |

## 失敗／不如預期（批次層）

| # | 現象 | 根因分類 | 應有行為 |
| --- | --- | --- | --- |
| **B1** | 多個「已完成」archived plan 的早期刀缺獨立 V，靠 audit 才標 `implementation_done` | **retrospective honesty** | 完成宣告前對齊 transport；audit 可降級不可改寫歷史為 closed |
| **B2** | Phase 2 list-detail：行為已過 V，isolation 閘於 commit 前才紅 → fix 走 same-session | **post-verify process gate** | 機械閘失敗仍須有書面 transport；行為 V 通過 ≠ 可跳過 isolation |
| **B3** | MES naming：並行 session 改主 plan → 若併入關閉會污染 closure | **multi-session plan contamination** | Closure authority 用獨立 arbitration／evidence 檔關閉本 slice，避免吞進無關 dirty hunks |
| **B4** | Equipment／MES 子樹大量 `implementation_done` 與少數 closed 並存 | **portfolio mixed close_kind** | 索引／handoff 須分列；禁止 portfolio 級「全綠」敘事 |

## 仲裁紀要（代表性）

| 來源（泛化） | finding | 處置 | 泛化含義 |
| --- | --- | --- | --- |
| Shell-nav Phase 1 | §6 alias 與 §3 路徑不一致 | **fix**→re-verify | 文件 acceptance 仍要獨立 V |
| List-detail Phase 2 | 跨 feature QueryState inject | **fix** | isolation 屬 acceptance／process 閘 |
| MES naming | BDD 階段敘述語系閘失敗 | **fix**→V2 | 文案閘可為 blocking acceptance |
| List-detail／SPA 1a | 無獨立 V | 標 **`implementation_done`** | 誠實降級＝有效證據 |

## 量測欄

| 指標 | 值 |
| --- | --- |
| 檢視 archived／完成主線 plan（量級） | **≥8** |
| 明確 `slice_compliant_closed` 且有獨立 V | **≥4**（含 shell-nav、list-detail P2、QG1、MES naming） |
| Audit／process 回填為 `implementation_done` | **≥2**（list-detail P1、SPA 1a） |
| 與 3a「放棄 loop」同構、其後已矯正 | **是**（shell-nav／list-detail P2／SPA 1b+） |
| 新契約標籤（本檔） | **4**（B1–B4 對應下節） |

## 契約回饋

1. **`retrospective-implementation-done`** — 完成計畫 audit 發現缺獨立 V 時，正確動作是降級狀態＋後續刀強制 loop，不是補敘 `slice_compliant_closed`。
2. **`behavior-pass-then-isolation-fail`** — Verifier 行為綠之後仍可能撞 process／isolation 閘；關閉條件須含適用機械閘，或書面 defer。
3. **`arbitration-artifact-as-closure-authority`** — 多 session 並行改 plan 時，以獨立 arbitration／evidence 檔關閉本 slice，避免吞進無關 dirty hunks。
4. **`portfolio-mixed-close-kind`** — 同一 consumer 樹可同時存在 closed 與 `implementation_done`；handoff／dogfood 索引必須分列，禁止「專案計畫都做完＝loop 都合規」。

## Evidence pointers（consumer）

- `plans/archived/2026-07-23-1422-shell-nav-groups-module-isolation/`（含 `evidence/…phase-1-nav-contract.md`）
- `plans/archived/2026-07-23-1430-list-detail-state-restore/_plan.md`
- `plans/archived/2026-07-17-2315-qg1-contract-reconciliation/_plan.md`
- `plans/active/2026-07-27-2046-mes-mvp-launch/evidence/2026-07-28-mes-product-naming-arbitration.md`
- `plans/archived/2026-07-28-0840-equipment-standalone-package/_plan.md`
- WMS foundation Phase 4f1：`plans/active/…/evidence/2026-07-22-phase-4f1-remaining-ops.md`

## Disposition

- 本輪：**批次 evidence**；補強「3a 之後 adoption 已擴散」與「完成≠合規」對照。  
- **不**視為 Phase 3／Q5 closure；**不**把 MES／Equipment 領域細節升進 Ai-skill schema。  
- 與 [`3b`](3b-kaizenwms-multi-round-arbitration-maturity.md) 分工：3b＝多輪仲裁深度；3c＝完成樹廣度＋誠實降級。
