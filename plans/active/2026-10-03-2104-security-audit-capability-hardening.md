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

**Status**: draft — Phase 0 完成（2026-10-03，decision = proceed to Phase 1）；Q8 / Q10 已由使用者決定（2026-10-03），Phase 1–2 無 blocker；Phase 1 由主 session 執行（transport adaptation）。

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
   - **`severity` ≠ `potential_impact`**：`severity` 只在 `status: confirmed` 時有值；未確認的 finding 只填 `potential_impact` + `unresolved_fact`。不確定性必須明確表達，不得把未確認 finding 宣稱為已確認 High。
   - **`audit_execution` 執行證明**：list 層必填 `audit_execution { status, scope, coverage_ref, evidence_ref }`。`findings: []` 只有在 `audit_execution.status: completed` 時才代表「已執行、無符合條件 finding」；缺 `audit_execution` = **unknown**，不是 safe。
2. **Closure gate**（Phase 0 C1 修正位置）：`execution-flow.yaml` §gates 新增 `gate.software_delivery.security_audit_complete`（blocks `security_sensitive_completion_claim`），`artifact-gates.yaml` 只加對應 `required_evidence` 證據形狀（沿用 journey gate 雙處先例）；細化而非複製既有 `gate.software_delivery.validation_complete`「security blocker 不得隱藏」條款。裁決輸入如下：
   - `audit_execution` 缺漏或非 `completed` → block（unknown ≠ safe）
   - `confirmed` 且 `severity` ∈ {high, critical} → block
   - `needs_validation` 且 `potential_impact` ∈ {high, critical} → 需**人工審查紀錄**（審查者 + 結論）才可放行；不強制 risk acceptance、不視為已確認漏洞（Q8）
   - Risk acceptance 必須記錄 `decision_ref`、`owner`、`scope`、`expires_when`；`expires_when` = 被依賴 security control 變更（與 Phase 3 invalidation contract 同一觸發，Q10）
3. **Coverage + invalidation contract**：定義 `Entry Surface × Trust Boundary × Attack Class` 的 coverage unit 與 invalidation contract（source dependency + security control 變更 → coverage 失效）；擴充既有 [`stale-derived-state.md`](../../intelligence/engineering/anti-patterns/stale-derived-state.md) `stale_permission_state`，不另造概念。**資料存專案端**，Ai-skill 只存 contract。
4. **Verifier 補強**（Phase 0 C3 收窄）：V3 已有 evidence producer（authorization / guard 類風險可用 targeted mutation 機械枚舉），**不重寫**。本 plan 只補兩條 security 專屬規則到 `delegated-execution.md` §5：LLM 第二意見不得單獨構成 `confirmed` / `refuted`（須有 test / SAST / mutation / schema validator 等可重現證據）；歷史 intelligence 只能產生檢查假設，不能產生 finding 裁決。
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
| Trigger flow | `sd-contracts` / `sd-implementation` 判定敏感變更 → invoke `security-audit`（stance `fault_finding`）→ 產出 `security-finding-list` → `sd-validation` Verifier V3 對 finding 產證據 → completion claim 前 `gate.software_delivery.security_audit_complete` 檢查 `audit_execution` 與 open finding → pass / block / risk-acceptance decision |
| Trigger location | `workflow/software-delivery/execution-flow.yaml` §gates（gate）+ `artifact-gates.yaml` §required_evidence（證據形狀）；`ai-skill runtime capability-invoke` |
| Activation contract | 既有 `route.workflow.software-delivery` + capability invoke envelope；**不新增 route** |
| Generated surface | **不新增 surface**。Phase 0 C2：兩個 YAML 都已 `runtime_projection.enabled: true`（`workflow.software_delivery.*.contract`），Phase 2 修改會進入既有 generated surface，須 `ai-skill runtime compile` + `refresh`；目前無 Go validator 消費這兩個 contract（agent-read executable contract），故仍屬 doc-only trial、不得宣稱機械強制 |
| Validation scenarios | Phase 2 新增 ≥ 4 個 `validation/` scenario：(a) confirmed high finding 擋 closure；(b) refuted finding 附可重現證據可放行；(c) coverage 依賴的 control 變更 → 標 `needs_revalidation`；(d) `findings: []` 但無 `audit_execution` 證明 → unknown，擋 closure（未執行 ≠ 已執行但無 finding） |
| Test passing evidence | Phase 2 / 3 dogfood evidence |

**Doc-only 宣告**：Phase 1–3 不接入 runtime 機械 gate；本 plan 不得在 Phase 5 前宣稱 runtime integration 完成。接入 phase = Phase 5，entry condition = Phase 4 dogfood ≥ 2 個真實任務。Graduation deadline：2027-01-31（未達則降為 intelligence-only 並記錄）。

## Per-surface consumer 表

| Generated surface key | Named consumer(s) | Consumer 類型 |
| --- | --- | --- |
| （Phase 1–4 無新 surface）；既有 `workflow.software_delivery.execution_flow.contract` / `artifact_gates.contract` 內容擴充 | agent 經 software-delivery route 讀取 | 既有 executable YAML contract（非新 consumer） |
| Phase 5 候選：`security-audit` stance block | `ai-skill runtime capability-invoke` | CLI validator（既有 `runtime.capability_context.contract`，只改 severity） |

## Open Questions

- [ ] Q1：Coverage unit 是否需要第四軸 `Subsystem`（Cloudflare 原版有）？還是 `Entry Surface` 已涵蓋？ — *Phase 0: deferred → Phase 3（safe assumption：三軸起步，`subsystem` 為 optional 欄位）*
- [x] Q2：Finding schema 放 `workflow/software-delivery/templates/`（capability output，同 `review-report-template.md`）還是 `analysis/security/`？ — *resolved：`templates/security-finding-list-template.md`*
- [x] Q3：Closure gate 的 risk acceptance 由誰簽：使用者 decision record 即可，還是需要 `decision` asset class？ — *resolved：Decision asset class（project decision → 專案 `docs/decisions/`），`decision_ref` 指向它；簽核者 = 專案決策者，Ai-skill 不指定人*
- [ ] Q4：專案端 coverage 資料格式（YAML in repo / project-local SQLite）與 Ai-skill contract 的驗證方式（`ai-skill` 提供 validator？） — *Phase 0: deferred → Phase 3（safe assumption：專案 repo 內 YAML；Phase 1–4 不提供 validator）*
- [x] Q5：Security Light 模式下「可為空的 finding list」的最低理由欄位是什麼，才不會變成形式化填表？ — *resolved：`audit_execution { status, scope, coverage_ref, evidence_ref }`*
- [x] Q6：Reusable security intelligence（漏洞模式 / 修補 / 回歸測試）落在 `intelligence/engineering/anti-patterns/` 還是新子目錄？需走 reusable-guidance-boundary 去敏。 — *resolved：`intelligence/engineering/anti-patterns/`，不開新子目錄*
- [x] Q7：Verifier V3「可重現證據優先」是否應寫入 `plans/README.md` §Delegation loop SOP（canonical）而非 delegated-execution.md？ — *resolved：寫 `delegated-execution.md` §5（delivery 域擴充），不動 loop canonical*
- [x] Q8：`needs_validation` + `potential_impact: high` 的處置邊界：要求人工審查即可，還是一律需 risk acceptance 才能 closure？（`severity` 只屬 confirmed，已在 Decision 第 1 點凍結） — *resolved（使用者 2026-10-03）：人工審查紀錄即可放行，不強制 risk acceptance*
- [ ] Q9：Coverage invalidation 的 **dependency scope** 怎麼定義：共用 control（AuthorizationHandler、policy、query filter、middleware）修改時，哪些 coverage unit 失效？以 trust boundary 為鍵，還是需顯式 dependency 清單？ — *Phase 0: deferred → Phase 3（safe assumption：trust boundary 為鍵 + 顯式 control dependency 清單；檔案 hash 不足）*
- [x] Q10：Risk acceptance 的 `expires_when` 用什麼條件表達（時間、commit 範圍、被依賴 control 變更）？與 Q3 簽核者一併決定。 — *resolved（使用者 2026-10-03）：被依賴 control 變更即失效；不用時間期限或 commit 範圍*

> 2026-10-03 review 回寫：外部 review 確認架構方向不變、維持 draft 直接進 Phase 0、不擴大 scope；新增 Q8–Q10 與 Decision 第 1–2 點的 severity / audit_execution / risk acceptance 欄位，Q5 的「空 finding list 最低理由」由 `audit_execution` 吸收（Phase 0 確認後標 resolved）。

## Phase 0 — Pre-Build Interrogation + Architecture Compatibility Preflight

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [x] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [x] 若盤點新發現問題，已加入 §Open Questions（無新增 question；新發現為架構衝突 C1–C4，見 0.1）

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1 Subsystem 軸 | deferred → Phase 3 | 無 dogfood 證據判斷；三軸起步、`subsystem` optional，不阻擋 Phase 1 schema |
| Q2 schema 位置 | resolved | [`templates/README.md`](../../workflow/software-delivery/templates/README.md) 已把 `review-report-template.md` 定為 `code-review` capability output；[`analysis/security/README.md`](../../analysis/security/README.md) 只放觀察方法、不放產物 |
| Q3 risk acceptance 簽核 | resolved | [`domain-policies.md`](../../workflow/software-delivery/domain-policies.md) Decision asset class：project decision → 專案 `docs/decisions/`，owner = 決策者 |
| Q4 專案端資料格式 | deferred → Phase 3 | 只影響 coverage，不影響 Phase 1–2；Phase 1–4 不提供 validator（避免無 consumer surface） |
| Q5 空 list 理由 | resolved | `audit_execution` 欄位（2026-10-03 review 回寫） |
| Q6 reusable intelligence 位置 | resolved | `analysis/security/README.md` §與其他層的關係：安全 anti-patterns → `intelligence/engineering/anti-patterns/` |
| Q7 V3 規則寫哪 | resolved | `delegated-execution.md` 自述 canonical 範圍 = delivery 域角色責任 / V4–V5 擴充；loop SOP 不複製。security 證據偏好屬 delivery 域擴充 |
| Q8 needs_validation 高影響處置 | resolved | 使用者決定：人工審查紀錄即可 |
| Q9 dependency scope | deferred → Phase 3 | safe assumption 已記錄；需 dogfood 驗證 |
| Q10 `expires_when` | resolved | 使用者決定：依賴 control 變更 |

### Phase 0.1 — Preflight

- [x] 完成 [`pre-build-interrogation.md`](../../workflow/software-delivery/requirements/pre-build-interrogation.md)（見下方 Pre-build Interrogation）
- [x] 讀 software-delivery README / execution-flow / artifact-gates / delegated-execution / domain-policies；cross-cutting/review README + invocation-points；governance/cognitive-stance.md；runtime/capability-context.yaml；analysis/security/；glossary
- [x] 確認 `security-finding-list` 無其他 consumer 定義：只出現在 registry、ADR-014、cross-cutting/review README 與 archived plan 的名稱層級，無 schema → 無雙 source
- [x] 確認 delegation loop plan 的 Shared State Contract 未涵蓋 Decision 第 4 點（Q5 處理 writer/reader/owner 狀態契約，非證據偏好）；但 V3 evidence producer 已部分涵蓋 → C3
- [x] 記錄 preflight 最低格式（下表）

**Slice claim**：搜尋本地 / 遠端 commit、工作樹與 `.agent-goals/`，無 `security-audit-capability-hardening` 既有實作；`.agent-goals/` 僅一個不相關的 paused goal。
**Loop 判定**：Phase 0 屬只讀盤點 + plan 回寫，依 `delegated-execution.md` §1「純問答 / 只讀審計不觸發」由主 session 執行；Phase 1 起寫入 workflow 檔案，屬執行意圖；**transport adaptation（使用者 2026-10-03 選擇）**：Phase 1 由主 session 擔任 executor，角色降格記錄在案；Phase 2 起是否回到三角色 loop 於 Phase 1 結束時重評。

| 欄位 | 內容 |
| --- | --- |
| Trigger | 使用者要求開始本 plan Phase 0（2026-10-03） |
| Checked sources | 見上方 0.1 清單 |
| Conflicts | **C1** closure gate 位置：`artifact-gates.yaml` 的 activation 是文件 artifact 審查；completion-claim gate 的慣例在 `execution-flow.yaml` §gates（`gate.software_delivery.*_complete`），證據形狀在 `artifact-gates.yaml`（journey gate 雙處先例）→ Decision 第 2 點已修正。<br>**C2** 兩個 YAML 皆 `runtime_projection.enabled: true`，原 plan「不進 projection」描述不正確 → Runtime Execution Path 已修正（不新增 surface，但需 compile + refresh）。<br>**C3** V3 evidence producer 已涵蓋 authorization 類 targeted mutation → Decision 第 4 點收窄為兩條 security 專屬規則。<br>**C4** glossary 中 `findings` 已被 KGE（Validate→Findings）使用 → 本 plan 一律用限定詞 `security_finding`，不登記裸 `finding` |
| Interrogation | 見下方 Pre-build Interrogation |
| Open Questions 核對 | 見 0.0 表：resolved Q2/Q3/Q5/Q6/Q7/Q8/Q10；deferred Q1/Q4/Q9 |
| Decision | **proceed to Phase 1**（Q8 / Q10 已決定，Phase 2 不再 blocked） |
| Validation | grep 確認無雙 source；讀 YAML `runtime_projection`；plan diff review；commit-msg hooks |

### Pre-build Interrogation

- **Goal**：`security-audit` invoke 後有可驗證產物與 closure 判斷，caller 能回答「審過沒、審了什麼、哪些已失效」。
- **Scope**：finding template、software-delivery gate + 證據形狀、V3 兩條規則、coverage / invalidation contract、validation scenarios、dogfood。
- **Non-goals**：新 workflow domain / lifecycle phase；runtime.db 新 table；專案 coverage 資料入 Ai-skill；Phase 5 前的機械強制；自動 diff → trust boundary 偵測；SAST / SCA 工具整合。
- **Acceptance / validation target**：Phase 1 = template 存在且被 registry / invocation-points / templates README 引用 + link check；Phase 2 = ≥ 4 scenarios + `runtime compile/refresh` 無錯；Phase 3 = contract 文件 + 去敏檢查；Phase 4 = `verdict_kind: expected` 證據 ≥ 2 任務。
- **Framework discovery**：canonical source = owner-layer YAML / Markdown（execution-flow.yaml、artifact-gates.yaml、templates/、delegated-execution.md、stale-derived-state.md）；runtime.db 是 projection；registry `artifact` 欄位不變。
- **Duplication risk**：新 gate 與 `gate.software_delivery.validation_complete` 的 security 條款重疊 → 新 gate 明文為其細化並互相引用；V3 規則與 evidence producer 重疊 → 已收窄（C3）。
- **Open questions**：Q8 / Q10 已由使用者決定；Q1 / Q4 / Q9 = `safe_assumption`（Phase 3 驗證）。
- **Assumptions**：Light 模式可只填 diff-scope `audit_execution` 與空 list；三軸 coverage 足以起步。
- **Decision**：proceed（Phase 1）

## Phase 1 — Finding Schema + Template

- [ ] 定義 `security-finding-list` schema：list 層 `audit_execution`；finding 層 finding id、attack class、entry surface、trust boundary、status enum、`severity`（僅 confirmed）、`potential_impact`、`unresolved_fact`、evidence refs（可重現 / 推理）、hypothesis source（intelligence ref，僅作假設）、resolution / risk acceptance（`decision_ref`、`owner`、`scope`、`expires_when`）
- [ ] 明文三判斷分離：schema valid ≠ evidence established ≠ merge allowed
- [ ] Template 落地於 `workflow/software-delivery/templates/security-finding-list-template.md`（Q2）並接 templates README、cross-cutting/review README 與 invocation-points
- [ ] Light / Standard / Deep 對 Cognitive Mode 的映射表（不新增機制）

完成條件：schema + template 存在、被 registry artifact 欄位與 invocation-points 引用、link check 通過。

## Phase 2 — Closure Gate + Validation Scenarios

- [ ] `execution-flow.yaml` §gates 新增 `gate.software_delivery.security_audit_complete`；`artifact-gates.yaml` / `.md` 新增對應 required_evidence（C1）；與 `validation_complete` 互相引用
- [ ] `ai-skill runtime compile` + `refresh` 確認 projection 更新（C2）
- [ ] 新增 ≥ 4 個 validation scenario（見 Runtime Execution Path；含「未執行 ≠ 無 finding」）
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
- [ ] **驗證邊界明寫**：Phase 4 只驗證「workflow 依契約產生正確的阻擋**決策**」（expected verdict）；「runtime 真的**擋得住**」（actual enforcement）不在 Phase 4 範圍，留給 Phase 5。Evidence 每筆標 `verdict_kind: expected | enforced`，Phase 4 只能出現 `expected`
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
