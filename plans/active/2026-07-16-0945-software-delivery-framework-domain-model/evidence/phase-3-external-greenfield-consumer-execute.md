# Phase 3 Execute — External Greenfield Consumer（2026-07-16）

**Run**: phase-3-external-greenfield-consumer-execute  
**Plan**: `2026-07-16-0945-software-delivery-framework-domain-model`  
**Consumer**: `<PROJECT_ROOT>`（見 `local/plan-evidence/` run boundary）  
**Paired ERA run**: [`2v-external-greenfield-consumer-phase2-preflight.md`](../../2026-07-08-0825-delegation-verification-arbitration-loop/evidence/2v-external-greenfield-consumer-phase2-preflight.md)

**Method**: classify-before-create dry-run + consumer artifact audit + Phase 2 產出預分類；**未派 E/V**（preflight）。

---

## 1. 全庫 artifact 抽查（N=3）

| 路徑模式 | 主歸屬 | 備註 |
| --- | --- | --- |
| `docs/architecture/*` | Policy | charter、pre-commit、api pipeline |
| `docs/workflow/software-delivery.yaml` | Policy + Process | phases + rules 投影 |
| `docs/workflow/delegated-execution.md` | Process | ERA loop 消費者鏡像 |
| `docs/domains/<domain>/*` | Asset | domain bundle overlay |
| `docs/testing/*` | Asset (Policy-shaped) | test-strategy = contract-like |
| `docs/Reference Project/**` | Asset (read-only Knowledge) | 不 promote |
| `plans/active/**` | Asset | Plan deliverable |
| `scripts/verify*` | Automation | Policy 投影 |

**覆蓋**：抽查 39 個 `docs/` + `plans/` 檔；**38/39 單一主歸屬**；yaml 雙桶為已知投影（與 Phase 0 一致）。

---

## 2. classify-before-create（Phase 2 將產出）

| 將建立 | Asset class | Placement | Owner | 需新目錄？ |
| --- | --- | --- | --- | --- |
| Angular workspace | Implementation deliverable | `web/<spa-app>/`（charter §3） | Executor | 否（charter §3） |
| Mock adapter | Implementation deliverable | under `web/.../src/app/` | Executor | 否 |
| Scaffold evidence | Evidence | `docs/evidence/phase-2-scaffold/` | Verifier | 否 |

**結論**：Orchestrator 步驟 0 可在不開新第一級目錄下完成分類。**pass**

---

## 3. Placement dual-write（consumer）

| 檢查 | 結果 |
| --- | --- |
| placement 全文 duplicate | 無 — charter §2–§5 + yaml artifacts 分工 |
| domain 檔自帶 placement 表 | 無 |
| `docs/planning/` 或 `docs/contracts/` 扁平目錄 | 未使用 — domain bundle 合法 overlay |

---

## 4. Mechanical

| Check | Result |
| --- | --- |
| consumer `scripts/verify.sh` | OK |
| framework `domain-policies.md` §3.2 domain-bundle | 解釋 consumer 佈局 |

---

## 5. 量測欄

| 欄位 | 值 |
| --- | --- |
| run_kind | external_execute_preflight |
| classify_before_create | pass |
| n_equals_3_fit | pass |
| placement_dual_write | pass |
| e_v_loop | 未執行（brief 在 consumer plan evidence） |

## 6. 契約回饋

1. **domain-bundle + charter** 足以讓 greenfield consumer 不需 `docs/planning/`。
2. **Execute 前** consumer plan evidence 承載 brief 比 Ai-skill evidence 更合適（專案細節邊界）。
3. **Primary Model 指針** 已加 consumer charter 一行（`<AI_SKILL_REPO>` 占位）。
