# Phase 0 Convergence Audit（2026-10-04）

**Scope**：只讀既有實例，不修改 runtime、validator 或 domain schema。
**判定規則（maintainer 2026-10-04）**：

1. 每條 invariant **獨立計算**，需 ≥3 個 domain。
2. 三個 domain 必須是**獨立實作**；只引用同一份上游契約的實例合計只算 1。
3. 未達門檻者列為 **watch**：可參考，但不是新 domain 的必要條件。
4. 只有 ✅ 計入門檻；`partial`、`variant`、`absent` 都不計。

## 實例清單

| Domain | 檢查的實例 | 備註 |
|---|---|---|
| narrative-video (NV) | `workflow/narrative-video-production/records/text-evidence.yaml`、`text-evidence-multimodal-resolution.md` | 原矩陣已列 |
| 3d-character (3D) | `workflow/3d-character-production/records/candidate-record.yaml`、`identity-acceptance.yaml`、`artifact-gates.yaml` | 原矩陣已列；Phase 0 新發現 `identity-acceptance` |
| security-audit (SEC) | `analysis/security/security-coverage-ledger.md`、`workflow/software-delivery/templates/security-finding-list-template.md` | 原矩陣已列；finding list 與 ledger 屬同一 capability，合算 1 個 domain |
| legal (LG) | `workflow/legal/reference-sources.md`、`artifact-gates.md`、`research/README.md`、`negotiation/README.md` | 原矩陣已列 |
| translation (TR) | `workflow/translation/contracts/translation-decision.yaml`、`finality.yaml`、`source.yaml` | **Phase 0 新增**（0.2 掃描） |

排除：investment 的 Decision Reasoning 與 legal 共用上游 [`cross-cutting/decision-support`](../../../../workflow/cross-cutting/decision-support/README.md)，依規則 2 不另計。software-delivery delegated-execution 的 refuted 留證，屬於 SEC 同一 capability 的 verifier 協議，不另計。

## 收斂矩陣

### I1 保留原始觀察（observed／raw 不被 derived 覆寫；落選者保留）

| Domain | 判定 | 依據 |
|---|---|---|
| NV | ✅ | `text-evidence.yaml` gate `observed_candidate_resolved_layers`、`raw_text_retained_on_boundary_recovery`；`text-evidence-multimodal-resolution.md` §Evidence non-destructive resolution（rejected／uncertain／merged 仍是 evidence，禁止 hard-delete） |
| 3D | ✅ | `candidate-record.yaml` `retention.rejected_must_retain: true` |
| TR | ✅ | `translation-decision.yaml` `candidates[]` 保留 infeasible 候選並要求 `reason`（`required_when: feasible == false`）；forbidden `collapsing_candidate_space_into_single_forced_output` |
| SEC | partial | `refuted` 是 status 而非刪除；`examined_no_finding` 必填。但沒有明文的保留規則 |
| LG | partial | `negotiation/README.md`「保留每一版，不要覆蓋」：保留的是產出版本，不是觀察證據 |

**3 ✅ → Promote**

### I2 來源錨點（每筆證據指回可重查的來源與版本）

| Domain | 判定 | 依據 |
|---|---|---|
| NV | ✅ | `semantic_candidate.source_ref` 必填；`resolution.sources[]`；text_region 的 `timestamp`／`frame`／`box` |
| 3D | ✅ | `candidate-record.yaml` `provenance`（provider／model／input_reference_ids／timestamp 必填；缺項標 `unavailable`／`partial`，禁止捏造）；`identity-acceptance.yaml` `evidence.rubric_applied` + `fixed_views` |
| SEC | ✅ | `evidence.reproducible[]`；ledger 的 `evidence_ref` + `verified_at_ref`（commit 或版本，「不可只寫日期」） |
| LG | ✅ | `gate.legal.law_citation_versioned`（名稱 + 版本或最新修正日 + 查核日）；`gate.legal.source_version_pinned` |
| TR | partial | `selection.decision_basis` 必須列 artifact refs；但 source 端 `evidence_refs`、`text_origin` 都是 optional，且沒有版本 |

**4 ✅ → Promote**

### I3 來源變更即失效（宣告依賴範圍，來源變更就轉為待重驗）

| Domain | 判定 | 依據 |
|---|---|---|
| SEC | ✅ | ledger §Invalidation（`depends_on_controls` + `source_scope` → `needs_revalidation`；機械回放）；`risk_acceptance.expires_when` |
| 3D | ✅ | `identity-acceptance.yaml` `validity: current｜re_review_required｜stale` + `invalidation_rules`（依 mutation_class 分支）+ `mutation_event`；`artifact-gates.yaml` 下游 stage 要求 `validity == current` |
| LG | variant | `research/README.md` Staleness rule：引用前**當次核對**是否有新版，並記新查核日。這是**使用時重驗**，不是依賴變更觸發的轉態 |
| NV | absent | 沒有找到失效或轉態規則 |
| TR | absent | 沒有找到失效或轉態規則 |

**2 ✅ + 1 variant → Watch**（見 §研究發現 F2）

### I4 定案要附理由（candidate → resolved／accepted 必須附 reason 與 sources）

| Domain | 判定 | 依據 |
|---|---|---|
| NV | ✅ | gate `resolution_reason_required_when_resolved`；四態都有最低 trace（`reason`／`reasons[]`／`reason.code`／`merged_into`） |
| 3D | ✅ | `artifact-gates.yaml` promotion：`candidate_record.decision` 與 `reason` 齊；forbidden `promote_when_decision_or_reason_missing` |
| SEC | ✅ | Status 規則：`confirmed`／`refuted` 需要 `reproducible` 證據（LLM 第二意見不足以 refute）；`resolution.kind` + `ref`；`risk_acceptance.decision_ref` + `owner` |
| TR | ✅ | `selection.policy` 必填（I6）、`rationale`、`decision_basis` 必列 artifact；finality `needs_review` 必須有非空 `review.reason[]` |
| LG | ✅ | `gate.legal.strategy_reasoned`（Recommendation／Reason／Alternative／Trade-offs）。上游是 decision-support，但其他 4 個 domain 都沒有引用它，不違反規則 2 |

**5 ✅ → Promote**

## 獨立性檢查

| 關係 | 發現 | 判定 |
|---|---|---|
| 3D／NV／TR 共用 `schema_version: artifact-record/v1` | 這個標籤源自 3D plan 的 contracts（2026-08-31），之後 NV、TR 沿用；repo 內**沒有**定義檔，標籤只代表 record 封套慣例，不含 invariant 內容 | 不構成共用上游；但記為風險：日後若真的建 `artifact-record` schema，要重新檢查 |
| NV ↔ TR | NV 的 `confidence.type` 宣告「與 Finality 語意對齊」（詞彙對齊）；TR 的 `source.semantic_context` 可以接收 NV 的輸出（資料流） | invariant 內容是各自寫出的：TR candidate retention 2026-09-22（`1570ffee`）早於 NV 三層 2026-09-30（`58b49120`）與 non-destructive 2026-10-02（`0cc01d16`），而且 NV 的規則來自自己的事故 lesson。判定獨立 |
| 3D → TR 的 I1 | 3D `rejected_must_retain` 2026-08-31（`5de0991a`）早於 TR；但兩者的形狀不同（整筆 record 保留 vs 候選 + infeasible reason） | 判定獨立 |
| LG ↔ investment | 共用 decision-support | investment 不另計（已排除） |
| SEC ledger → stale-derived-state | 引用的是抽象 anti-pattern，不是 domain 契約 | 不影響 |

**敏感度**：如果 maintainer 改判 NV 與 TR 不獨立，I1 會降為 2（→ Watch），I4 降為 4（仍 Promote），I2 不受影響（TR 本來就只是 partial）。

## Lifecycle state ≠ epistemic status（maintainer 2026-10-04 加的檢查點）

問題：各 domain 是否把「證據走到哪個狀態」（lifecycle）與「證據有多可信」（epistemic）分成兩個維度？

| Domain | 判定 | 依據 |
|---|---|---|
| NV | ✅ | §兩種否定：`rejected`（觀測本身不成立）≠ `uncertain`（觀測成立但解讀未收斂），不得共用 discard；`confidence.type` 是另一個欄位；Evidence layer ≠ Production layer |
| SEC | ✅ | 模板 §三個判斷分開（schema valid／evidence established／closure allowed）；evidence 分 `reproducible` 與 `reasoning`；`hypothesis_source`「只提供檢查假設，不是裁決證據」；`status` 與 `resolution.kind` 是不同欄位 |
| TR | ✅ | finality I22：「Uncertainty alone MUST NOT force needs_review」；forbidden `quality_or_confidence_scalar_as_closure` |
| 3D | partial | `decision`、`validity`、`maturity` 三軸分開，禁止 `identity_score`；但沒有獨立的可信度軸 |
| LG | partial | 來源層級 A–E 是獨立的可信度軸；但 Confidence 標籤的 `provisional` 同時代表 Strategy Pass 1（lifecycle）與未查證（epistemic），一個 enum 承載兩個維度 |

**3 ✅**：已達門檻。依 maintainer 指示，**不直接新增第五條 invariant**，而是列為 I5 候選，交由 Phase 1 決定（見 §研究發現 F1）。

## 研究發現

- **F1 — I5 候選：lifecycle ≠ epistemic**。3 個獨立 domain 都有明文：resolved 不代表權重較高，rejected 也不代表原始觀察失效。LG 的 `provisional` 是反例形狀（單一 enum 承載兩軸），可作為 pattern 的 anti-pattern 範例。這也回答了 Q5：三層 lifecycle **不能**一對一對應到 evidence-hierarchy 的 Observability 軸。
- **F2 — I3 有兩種觸發形狀**：依賴變更觸發（SEC、3D：來源在 repo 內，可以偵測變更）與使用時重驗（LG：外部來源無法監看，只能每次使用前重查）。若把 I3 放寬成「derived evidence 必須宣告重驗觸發條件（變更觸發或使用時觸發）」，就會達到 3。但這是**為了過門檻而改寫定義**，Phase 0 不採用；維持 Watch，並把兩種形狀都記在 watch 條目中。是否放寬由 maintainer 決定。
- **F3 — 禁止純量分數是另一個跨域訊號**：3D 禁 `identity_score`／`scalar_quality`，TR 禁 `quality_or_confidence_scalar_as_closure`，ECS 禁 confidence 數字（屬 governance 層，不計入 domain）。目前 2 個 domain，記錄、不升級。
- **F4 — 對 visual retrieval 的意義**：I2 的錨點種類已經涵蓋 path + commit（SEC）、URL／文件 + 版本 + 查核日（LG）、timestamp + frame + bbox（NV）、模型 provider + version + 輸入引用（3D）。visual retrieval 的「頁碼 + 區域 + 文件版本」可以由既有種類組合出來，**不需要新的錨點種類**。

## 結論

| Invariant | ✅ 數 | 處置 |
|---|---|---|
| I1 保留原始觀察 | 3 | Promote（敏感度：NV／TR 若改判不獨立 → Watch） |
| I2 來源錨點 | 4 | Promote |
| I3 來源變更即失效 | 2（+1 variant） | Watch |
| I4 定案要附理由 | 5 | Promote |
| I5 候選 lifecycle ≠ epistemic | 3 | 待 Phase 1 決定 |

名稱檢查：暫定名 `source-anchored-evidence` 只描述 I2。I1、I4（以及 I5 若納入）講的是**證據在決策過程中的保存與裁決**，不是來源錨點。依 maintainer「不要為了名稱限制 pattern」的原則，建議 Phase 1 重新命名，方向是能涵蓋「保存 + 錨定 + 裁決」的名稱。
