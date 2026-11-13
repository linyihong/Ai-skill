# External reference gap observations — Cloudflare security-audit skill（2026-11-13）

**Date**: 2026-11-13 · **Kind**: external-reference comparison（**不是 dogfood**）· **Status**: `observed`（未經真實 audit 驗證）· **verdict_kind**: n/a

## 這份紀錄是什麼

把本 plan 當初參考的外部 skill（`https://github.com/cloudflare/security-audit-skill`，branch `main`，`skills/security-audit/`）逐檔對照我們的契約後，留下的五筆差異觀察。

- **只記觀察，不記結論**。每筆指向雙方的真實 artifact；是否採用由日後的真實 audit 證據決定。
- **不計入 Phase 5 restart trigger**，也不是 graduation 證據（見 `_plan.md` §Phase 5 — Deferred）。
- **沒有修改任何契約欄位**。先改 template 欄位會讓「至少一輪真實 audit 不需修改 template 欄位」這個 graduation criterion 重新起算。
- 歸檔狀態不變；本檔只是 evidence 目錄的追加紀錄。

## 蒐集方式與限制

- 外部端：讀取 18 個檔案（`SKILL.md`、`RECONNAISSANCE.md`、`HUNTING.md`、`VALIDATION-AND-REPORTING.md`、`ATTACK-CLASSES.md`、10 份領域 companion、`report-schema.json`、兩支 `validate-*.cjs`）。內容經 subagent 讀取後轉述，**未逐字核對原文**；未記錄外部 repo 的 commit ref。
- 我方：直接讀 [`security-finding-list-template.md`](../../../../workflow/software-delivery/templates/security-finding-list-template.md)、[`security-coverage-ledger.md`](../../../../analysis/security/security-coverage-ledger.md)、[`execution-flow.yaml`](../../../../workflow/software-delivery/execution-flow.yaml) 的 `gate.software_delivery.security_audit_complete`、[`delegated-execution.md`](../../../../workflow/software-delivery/delegated-execution.md) §5 的 V3 security 證據規則。`artifact-gates.yaml` 未讀。

## Observations

欄位沿用 Evidence Log 慣例：`strength ∈ hard/soft`；`decision_blocked ∈ yes/no`。

| id | surface | 外部 artifact | 我方 artifact | observation | strength | decision_blocked |
| --- | --- | --- | --- | --- | --- | --- |
| XR-1 | coverage-completeness | `HUNTING.md`（coverage-critic waves） | ledger §Coverage Unit；template §Audit Execution | 外部在每波 hunter 之後由 fresh agent 讀 ledger 與 source，回報 `missing_units` / `reassign_ids`；standard / deep 另需一個 distinct final-clean critic 也不接受新工作才算覆蓋完成。我方 unit 由 auditor 自行列出，verifier 審 finding 與 `not_covered`，沒有專責找「漏列 unit」的步驟。 | soft | no |
| XR-2 | unit-derivation | `RECONNAISSANCE.md`（`architecture.md`，上限約 1,000 字） | ledger §Coverage Unit | 外部先產出 principal、entry surface、trust boundary、起始路徑的摘要，unit 由 surface × boundary × subsystem × attack class 的實質組合推導。我方定義了 unit 的軸與狀態，沒有定義 unit 從哪個產物推導出來。 | soft | no |
| XR-3 | over-reporting-guard | `SKILL.md`（exclusions、anti-patterns）；hunter result 的 `hardening` 陣列 | template §Audit Execution（`not_covered` 誤用）、§Findings | 外部明列不算漏洞的情況（縱深防禦缺口、checklist 偏差、同一 principal 的自我影響、把 crash 加強成 code execution），並把非 finding 的建議放進獨立的 `hardening`。我方現有規則防的是少報（`not_covered` 不得用來排除範圍內缺陷）；沒有對應的多報排除規則，也沒有 hardening 落點。 | soft | no |
| XR-4 | severity-anchoring | `SKILL.md`（severity anchors）；`validate-findings.cjs`（overall 不得高於 impact） | template §Findings、§Status 規則 | 外部五級（含 `informational`）各有定義，並區分 likelihood / impact / overall，機械檢查 overall 不高於已證明的 impact。我方 `severity` 是四級 enum，沒有各級定義，也沒有「不得超過已證明影響」的規則。 | soft | no |
| XR-5 | trace-and-dedup | `report-schema.json`（`trace[]`：entrypoint → propagation → sink，逐筆 file / line / scope）；`HUNTING.md`（fingerprint + root cause 合併） | template §Findings（`evidence.reproducible` / `reasoning`、`SEC-<n>`） | 外部每筆紀錄有結構化 trace 與穩定 fingerprint；同一 root cause 的多條入口合併成一筆，跨次執行以 fingerprint 對應。我方 evidence 是自由 ref，id 是流水號，沒有同 root cause 的合併規則，也沒有跨次 audit 對應同一 finding 的鍵。 | soft | no |

五筆 `strength` 都是 `soft`：差異本身可由文件對照確認，但「我方因此在真實 audit 中漏報或誤判」尚無任何實例。

## 並用時的對接觀察（非候選）

- 外部 ledger 沒有 `depends_on_controls`。把外部 ledger 直接當我方帳本使用時，§Invalidation 的 control 依賴失效無法運作。
- 外部 `validate-findings.cjs` 與 `validate-coverage-ledger.cjs` 各自只讀單一檔案；`result_fingerprints` 與 findings 的跨檔一致性沒有被任一腳本檢查。

## 會讓某一筆升級為 `verified` 的事件

以下任一情況發生在真實 `security-audit` 使用中，才把對應觀察改標 `verified` 並評估契約修改；否則維持 `observed`。

| id | 事件 |
| --- | --- |
| XR-1 / XR-2 | 事後發現某個 entry surface 或 trust boundary 從未被列為 unit，且不是 `not_covered` 明列的面 |
| XR-3 | finding list 出現沒有可達邊界違反的項目被標為 finding，或 hardening 建議沒有地方放而被塞進 findings |
| XR-4 | 兩次 audit 對同類 finding 給出不同 `severity`，或 `severity` 高於可重現證據所證明的影響 |
| XR-5 | 同一 root cause 被記成多筆，或增量 audit 無法判斷某筆 finding 是否為前次的同一筆 |

反向同樣成立：若下一次真實 audit 沒有出現上列事件，記為 null result，該筆維持 `observed`，不因外部有此做法而採用。
