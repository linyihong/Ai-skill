---
id: 2026-10-03-2104-security-audit-capability-hardening
plan_kind: main
status: draft
owner: workflow
owner_layer: workflow
created: 2026-10-03
parent: null
---

# Security Audit Capability Hardening（`security-audit` invoke 補實）

**Status**: draft — 尚未進 Phase 0；未執行 Pre-build Interrogation / Architecture Compatibility Preflight。

**Glossary Impact**: yes（候選，Phase 4 前不登記）— `security_finding`、`security_coverage_unit`、`evidence_invalidation_contract`（若與既有 `stale-derived-state` invalidation contract 語意重疊，只 cross-link 不新登記）。

## 一句話目標

把已登錄但空殼的 `security-audit` capability 補成**有輸出 schema、有 closure gate、有增量覆蓋 / 證據失效契約**的開發內資安能力；**不**新增 `workflow/security/` domain，不新增 lifecycle phase。

## Decision Rationale

### Problem & Why Now

外部參考（Cloudflare security-audit skill：Reconnaissance → Coverage Ledger → Isolated Hunters → Candidate Validation → Independent Verification）引出一份「Security Development Workflow 五階段」提案。對照 repo 現況後：

| 提案假設 | Repo 現況 |
| --- | --- |
| 需要新 Security Workflow | **衝突**：[`workflow/software-delivery/README.md`](../../workflow/software-delivery/README.md) #16 與 [`analysis/security/README.md`](../../analysis/security/README.md) 明文「不另開 `workflow/security/` domain」；ADR-013：Review = capability invoke，非 phase |
| Findings schema 已有 | **缺**：`security-finding-list` 只是 [`capability-registry.yaml`](../../knowledge/runtime/capability-registry.yaml) 的 artifact 名稱，無 schema / template |
| Mechanical gate 可擋資安問題 | **未成立**：[`runtime/capability-context.yaml`](../../runtime/capability-context.yaml) Phase 1.2 stance 缺漏 = warning，非 block；artifact-gates 無 security finding gate |
| 獨立驗證可沿用 | **部分**：[`delegated-execution.md`](../../workflow/software-delivery/delegated-execution.md) Verifier V1–V5 已有角色 / 上下文獨立；僅限 `delegation.enabled` 任務 |
| Coverage / 證據失效可存 runtime.db | **缺且 layer 錯**：runtime.db 是治理狀態；專案 coverage 是 project data，Ai-skill 只應擁有 contract |

Why now：`security-audit` 已有兩個 caller slice（`sd-contracts`、`sd-implementation`）與一個 analysis consumer（media-entitlement），但 invoke 後沒有可驗證產物——caller 無法判斷「審過了沒、審了什麼、什麼已失效」。

### Decision

漸進補實既有 capability，依序：

1. **Finding schema + template**：定義 `security-finding-list` 的欄位（含 `status: candidate | confirmed | refuted | needs_validation`、evidence ref、attack class、trust boundary），並把「格式通過 ≠ 證據成立 ≠ 可合併」三判斷分開。
2. **Closure gate**：`artifact-gates.yaml` 新增 gate — 未解決的 high / critical `confirmed` 或 `needs_validation` finding 不得進 `sd-closure`（除非有明確 risk acceptance decision）。
3. **Coverage + invalidation contract**：定義 `Entry Surface × Trust Boundary × Attack Class` 的 coverage unit 與 invalidation contract（source dependency + security control 變更 → coverage 失效）；擴充既有 [`stale-derived-state.md`](../../intelligence/engineering/anti-patterns/stale-derived-state.md) `stale_permission_state`，不另造概念。**資料存專案端**，Ai-skill 只存 contract。
4. **Verifier 補強**：V3 對抗性證據優先可重現來源（test / SAST / schema validator），LLM 第二意見不得單獨構成 `confirmed` / `refuted`；歷史 intelligence 只能產生檢查假設，不能產生 finding 裁決。
5. **Stance gate 升級（gated）**：`security-audit` invoke 缺 `fault_finding` 由 warning → block，僅在 1–4 有 dogfood 證據後評估。

Light / Standard / Deep 分級**不新增機制**，映射到既有 Cognitive Mode（execution_mode / governance_mode）。

### Alternatives Considered

- A. 照提案新增 `workflow/security/` 五階段 workflow：**reject** — 違反既有 domain 決策與 ADR-013；與 software-delivery intake→closure 重複；撞 Watch-Out Wall 2（Workflow inflation）。
- B. 把 coverage ledger 做成 Ai-skill runtime.db table：**reject** — project data 進 governance runtime，違反 owner-layer 與 sanitization 邊界；且無 named consumer 會撞 `define_runtime_trigger_flow` forbidden rule (b)。
- C. 先做通用 Independent Verification Contract 給所有 review：**defer** — delegation loop plan 已在處理（Shared State Contract Promotion），重做會雙 source。
- D. 漸進補實既有 capability（**accept**）。

### Why Not an ADR Yet

Schema 與 gate 未經真實專案 dogfood；coverage unit 維度（是否需要 Subsystem 第四軸）未定；stance block 升級依賴 ADR-014（Proposed）後續。

### ADR Promotion Criteria（completed 時驗證）

- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why 全中
- [ ] ≥ 2 個真實專案任務使用 finding schema + closure gate
- [ ] Open Questions 全解
- [ ] 沒有更輕的 promotion target（多數內容可能只需停在 workflow / intelligence layer）
- [ ] 至少 1 次 coverage invalidation 實際觸發重驗的證據

### Consequences（預期）

#### 正面

- `security-audit` invoke 有可驗證產物，caller 可判斷審查狀態
- 資安留在開發流程內作為治理維度，不新增 domain
- 證據失效有明確 contract，避免「舊審查結果永久有效」

#### 負面

- software-delivery artifact-gates 增加一條 gate，closure 成本上升
- 專案端需維護 coverage 資料

#### 風險

- Finding schema 過度設計 → 小變更也被迫填完整表（緩解：Light 模式只需 diff-level finding list，可為空並附理由）
- 歷史 intelligence 污染判斷（緩解：Decision 第 4 點，intelligence 僅能產生 hypothesis）
- 執行目標程式碼驗證漏洞的沙箱需求（緩解：無 OS 層沙箱時不執行目標程式碼，finding 保持 `needs_validation`）

**Watch-Out List citation**（[`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List）：Wall 2 Workflow inflation（不新增 workflow / phase）；Wall 1 Discovery confused with Activation（不宣稱 diff → trust boundary 自動偵測已存在）；Wall 4 Telemetry explosion（coverage 只存 contract 要求的最小欄位，不做全量掃描紀錄）。

## Runtime Execution Path

| 項目 | 內容 |
| --- | --- |
| Runtime owner | `workflow/software-delivery`（gate）+ `knowledge/runtime/capability-registry.yaml`（artifact 宣告） |
| Trigger flow | `sd-contracts` / `sd-implementation` 判定敏感變更 → invoke `security-audit`（stance `fault_finding`）→ 產出 `security-finding-list` → `sd-validation` Verifier V3 對 finding 產證據 → `sd-closure` 前 artifact gate 檢查 open high finding → pass / block / risk-acceptance decision |
| Trigger location | `workflow/software-delivery/artifact-gates.yaml`；`ai-skill runtime capability-invoke` |
| Activation contract | 既有 `route.workflow.software-delivery` + capability invoke envelope；**不新增 route** |
| Generated surface | Phase 1–3 **doc-only trial**：不新增 `runtime_projection`；artifact-gates 走既有 projection（若有） |
| Validation scenarios | Phase 2 新增 ≥ 3 個 `validation/` scenario：(a) open high finding 擋 closure；(b) refuted finding 附可重現證據可放行；(c) coverage 依賴的 control 變更 → 標 `needs_revalidation` |
| Test passing evidence | Phase 2 / 3 dogfood evidence |

**Doc-only 宣告**：Phase 1–3 不接入 runtime 機械 gate；本 plan 不得在 Phase 5 前宣稱 runtime integration 完成。接入 phase = Phase 5，entry condition = Phase 4 dogfood ≥ 2 個真實任務。Graduation deadline：2027-01-31（未達則降為 intelligence-only 並記錄）。

## Per-surface consumer 表

| Generated surface key | Named consumer(s) | Consumer 類型 |
| --- | --- | --- |
| （Phase 1–4 無新 surface） | — | — |
| Phase 5 候選：`security-audit` stance block | `ai-skill runtime capability-invoke` | CLI validator（既有 `runtime.capability_context.contract`，只改 severity） |

## Open Questions

- [ ] Q1：Coverage unit 是否需要第四軸 `Subsystem`（Cloudflare 原版有）？還是 `Entry Surface` 已涵蓋？
- [ ] Q2：Finding schema 放 `workflow/software-delivery/templates/`（capability output，同 `review-report-template.md`）還是 `analysis/security/`？
- [ ] Q3：Closure gate 的 risk acceptance 由誰簽：使用者 decision record 即可，還是需要 `decision` asset class？
- [ ] Q4：專案端 coverage 資料格式（YAML in repo / project-local SQLite）與 Ai-skill contract 的驗證方式（`ai-skill` 提供 validator？）
- [ ] Q5：Security Light 模式下「可為空的 finding list」的最低理由欄位是什麼，才不會變成形式化填表？
- [ ] Q6：Reusable security intelligence（漏洞模式 / 修補 / 回歸測試）落在 `intelligence/engineering/anti-patterns/` 還是新子目錄？需走 reusable-guidance-boundary 去敏。
- [ ] Q7：Verifier V3「可重現證據優先」是否應寫入 `plans/README.md` §Delegation loop SOP（canonical）而非 delegated-execution.md？

## Phase 0 — Pre-Build Interrogation + Architecture Compatibility Preflight

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [ ] 已讀本 plan §Open Questions 全部條目
- [ ] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [ ] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [ ] 若盤點新發現問題，已加入 §Open Questions

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1–Q7 | | |

### Phase 0.1 — Preflight

- [ ] 完成 [`pre-build-interrogation.md`](../../workflow/software-delivery/requirements/pre-build-interrogation.md)
- [ ] 讀 software-delivery README / execution-flow / artifact-gates / validation / delegated-execution；cross-cutting/review invocation-points；governance/cognitive-stance.md；ADR-013 / ADR-014；analysis/security/
- [ ] 確認 `security-finding-list` 無其他 consumer 定義（避免雙 source）
- [ ] 確認 delegation loop plan 的 Shared State Contract 是否已涵蓋 Decision 第 4 點
- [ ] 記錄 preflight 最低格式（Trigger / Checked sources / Conflicts / Interrogation / Open Questions 核對 / Decision / Validation）

## Phase 1 — Finding Schema + Template

- [ ] 定義 `security-finding-list` schema：finding id、attack class、entry surface、trust boundary、status enum、severity、evidence refs（可重現 / 推理）、hypothesis source（intelligence ref，僅作假設）、resolution / risk acceptance
- [ ] 明文三判斷分離：schema valid ≠ evidence established ≠ merge allowed
- [ ] Template 落地（位置依 Q2）並接 templates README 與 cross-cutting/review invocation-points
- [ ] Light / Standard / Deep 對 Cognitive Mode 的映射表（不新增機制）

完成條件：schema + template 存在、被 registry artifact 欄位與 invocation-points 引用、link check 通過。

## Phase 2 — Closure Gate + Validation Scenarios

- [ ] `artifact-gates.yaml` / `.md` 新增 security finding gate（open high/critical confirmed 或 needs_validation → block closure，除非 risk acceptance）
- [ ] 新增 ≥ 3 個 validation scenario（見 Runtime Execution Path）
- [ ] `ai-skill runtime refresh` / validate 通過

完成條件：gate 文字 + scenarios 落地；doc-only，不宣稱機械強制。

## Phase 3 — Coverage + Evidence Invalidation Contract

- [ ] 擴充 `stale-derived-state.md`（或其子文件）定義 security coverage unit + invalidation contract（source dependency / security control 變更 → `needs_revalidation`）
- [ ] 明文：檔案 hash 不足以判斷證據有效；需追蹤共用 control（middleware、policy、query filter）
- [ ] 專案端資料格式 contract（依 Q4）；Ai-skill 不存專案 coverage 資料
- [ ] 沙箱規則：無 OS 層隔離時不執行目標程式碼，finding 留 `needs_validation`

完成條件：contract 文件落地、去敏檢查通過、被 analysis/security README 索引。

## Phase 4 — Dogfood

- [ ] ≥ 2 個真實專案任務（至少 1 個授權邊界變更）跑完 invoke → finding list → Verifier → gate
- [ ] ≥ 1 次 coverage invalidation 觸發重驗
- [ ] Evidence 存 `evidence/`（去敏，不含 host / token / 專案路徑）
- [ ] 回寫 Open Questions

## Phase 5 — Mechanical Graduation（gated）

Entry condition：Phase 4 完成且 schema 兩輪 dogfood 未需破壞性修改。

- [ ] 評估 `security-audit` stance 缺漏 warning → block（`runtime/capability-context.yaml`）
- [ ] 評估 closure gate 是否需 commit-msg / validator 機械檢查（需宣告 named consumer）
- [ ] 執行 Plan Completion Closure

## 完成條件

- [ ] Phase 1–4 完成；Phase 5 完成或明確 defer（附 owner / 重啟條件）
- [ ] 未新增 `workflow/security/` 或新 lifecycle phase
- [ ] 專案 coverage 資料未進入 Ai-skill repo
- [ ] Linked updates（registry、templates README、invocation-points、analysis/security README、glossary 若適用）已完成
- [ ] 執行 Plan Completion Closure

## Stakeholder 同意項目

- [ ] ⏳ 同意不另開 Security Workflow，以 capability 補實取代
- [ ] ⏳ 同意 coverage 資料屬專案端
- [ ] ⏳ 同意 Phase 5 機械化為 gated，不在首輪落地

## 與其他 plans 的關係

- [`2026-07-08-0825-delegation-verification-arbitration-loop`](2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)：Verifier V1–V5 owner；本 plan Decision 第 4 點只補 security 證據偏好，不重定義 loop（見 Q7）
- [`archived/2026-07-06-review-architecture-adr`](../archived/2026-07-06-review-architecture-adr/_plan.md)：ADR-013 capability invoke 模型的來源
- [`2026-06-16-1131-evidence-candidate-system.md`](2026-06-16-1131-evidence-candidate-system.md)：dogfood 案例可走 evidence candidate 索引回流本 plan
- [`archived/2026-06-10-1718-software-delivery-governance-invariants.md`](../archived/2026-06-10-1718-software-delivery-governance-invariants.md)：authority-coupled side effect / evidence shape 的前例
