# 2v — External greenfield consumer Phase 2 preflight（2026-07-16）

> **專案證據邊界**：brief YAML、ng 版本、build 輸出留 `<PROJECT_ROOT>` plan `evidence/2026-07-16-dual-plan-preflight.md`；Ai-skill 只保留 generalized metrics。

## Run 摘要

- **任務**：greenfield consumer（SPA-first、無 `src/`/`web/`  yet）Phase 2 scaffold 的 **ERA preflight** — Orchestrator 寫 brief + verification backfill；**未** spawn Executor/Verifier。
- **觸發**：domain-model plan Phase 3 execute + stakeholder 指定 consumer 雙 plan 測試。
- **配對**：[`phase-3-external-greenfield-consumer-execute.md`](../../2026-07-16-0945-software-delivery-framework-domain-model/evidence/phase-3-external-greenfield-consumer-execute.md)（classify-before-create）。

## Orchestrator checklist

| # | 步驟 | 結果 |
| --- | --- | --- |
| 1 | 讀 active plan Phase 2（不讀實作源碼） | pass |
| 2 | brief + backfill 寫入 consumer plan evidence | pass |
| 3 | plan commit 錨點 | **pending**（Execute 前） |
| 4 | Executor spawn | 未執行 |
| 5 | Verifier spawn（fresh context） | 未執行 |
| 6 | 仲裁 + close | 未執行 |

## Brief 摘要（generalized）

| 欄位 | 內容 |
| --- | --- |
| `slice_id` | `phase-2-spa-scaffold` |
| `slice_type` | `combined` |
| `delegation.enabled` | `true` |
| acceptance 條數 | 4（serve、shell+i18n、mock、charter 合規） |
| verification 命令 | `npm run build` + headless unit |
| deliverables | SPA tree + `docs/evidence/phase-2-scaffold/` |
| backfill | 5 行；含 `verifier_only` runtime + charter grep |

## Consumer overlay 對照

| 項 | 狀態 |
| --- | --- |
| `delegation_loop` in portable yaml | active |
| `delegation.enabled: false` 非豁免 | documented |
| consumer `delegated-execution.md` | 鏡像 Ai-skill sd-delegated-execution |
| `scripts/verify.sh` | OK（pre-scaffold skip backend/frontend） |

## 相對 2n 的差異

| 指標 | 2n（ExternalRepoC） | **2v（本 run）** |
| --- | --- | --- |
| E+V loop | 6/6 | **0**（preflight） |
| 程式碼起點 | 既有 monorepo | **greenfield**（零 web/src） |
| brief 落點 | sub-plan slices | consumer plan evidence |
| domain model 步驟 0 | 未強制 | **classify-before-create** 已跑 |

**判讀**：2v 證明 **greenfield consumer** 可在第一次 Execute 前完成 loop 契約準備；完整 E+V 留待使用者「開始執行 Phase 2」。

## 量測欄

| 指標 | 值 |
| --- | --- |
| run_kind | external_preflight |
| e_v_loop | 0/0 |
| brief_pre_backfill | yes |
| orchestrator_manageCode_diff | 0 |
| consumer_verify | OK |
| close_kind | n/a |

## 契約回饋

1. **greenfield preflight** — brief 可放在 consumer `plan/evidence/` 而非 Ai-skill；符合專案證據邊界。
2. **classify-before-create** — Orchestrator 步驟 0 與 ERA 步驟 2 可合併於同一 consumer evidence 檔。
3. **2v 不關 Phase 3** — schema promotion 仍 open；本 run 增強「第三個 consumer 形狀 = greenfield SPA-first」信號。

## 關聯

- Domain model：phase-3-external-greenfield-consumer-execute
- 正向對照：2n
- Consumer：`<PROJECT_ROOT>` material-foundation plan Phase 2
