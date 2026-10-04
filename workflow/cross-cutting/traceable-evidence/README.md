# Traceable Evidence（domain evidence 記錄形狀的跨域 pattern）

外部觀察（OCR、ASR、模型產出、稽核推理、法規查證、未來的視覺檢索……）進入 domain 時，記錄要讓人能**追溯**：原始觀察還在、指得回來源、每次定案都有理由，而且「走到哪個狀態」與「有多可信」分開記。

> **狀態**：`pilot`（reference contract）。本檔定義**語意**，不定義欄位名；各 domain 用自己的欄位實作，再連回本檔。
> **不是**必跑 stage、**不**註冊 `route.*`、**不**加 validator。收斂證據：
> [`evidence/phase0-convergence-audit.md`](../../../plans/active/2026-10-04-0936-evidence-record-contract/evidence/phase0-convergence-audit.md)。

## 收錄門檻

每條 invariant **獨立計算**，需 **≥3 個獨立 domain** 實作才收錄。共用同一份上游契約的實例合算 1。未達門檻者列為 **watch**。

| 類型 | 意義 | 對新 domain 的效力 |
| --- | --- | --- |
| **Invariant**（TE1、TE2、TE4、TE5） | 已有足夠的獨立實例支持，可作為跨域參考契約 | 參考契約：新 domain 處理外部觀察時優先採用；`pilot` 期間不是強制義務 |
| **Watch**（TE3） | 已有局部實例，但還沒證明跨域普遍 | **不是**必要條件。缺 watch 項不算違反本 pattern，除非該 domain 自己的 contract 已經要求 |

編號用 `TE` 前綴，避免與 domain 內部編號（例如 translation 的 I1–I23）混淆。

## 與其他 evidence 機制的分工

| 機制 | 回答的問題 | 與本檔的關係 |
| --- | --- | --- |
| [`enforcement/evidence-hierarchy.md`](../../../enforcement/evidence-hierarchy.md) | 證據如何**比較與加權** | 本檔只定義記錄形狀；TE5 讓加權所需的資訊不被 lifecycle 狀態吃掉 |
| [`governance/evidence-candidates/`](../../../governance/evidence-candidates/README.md)（ECS） | 哪個觀察要回流到**哪個 plan** | 不同 layer；ECS 只收 governance 證據，不收 domain evidence |
| [`governance/lifecycle/plan-evidence.md`](../../../governance/lifecycle/plan-evidence.md) | plan 的 run 全文**存放在哪** | 不同 layer；plan 儲存慣例 |
| [`stale-derived-state.md`](../../../intelligence/engineering/anti-patterns/stale-derived-state.md) | derived state 為什麼需要 invalidation contract | TE3（watch）的抽象原則 |

## Invariants

### TE1 — 保留原始觀察

原始觀察不得被推導結果覆寫；落選、被拒、被合併的候選要留下紀錄與理由，不得從紀錄中消失。

- **反例**：用 resolved 文字覆寫 raw 觀察；finalization 時把沒選上的候選直接刪掉；只留最終答案、不留候選空間。
- **為什麼**：日後重判、稽核或 TE3 的重驗，都需要原始觀察；刪掉就無法知道當初看到了什麼。

| Domain | 實作 |
| --- | --- |
| narrative-video | [`text-evidence.yaml`](../../narrative-video-production/records/text-evidence.yaml) gates `observed_candidate_resolved_layers`、`raw_text_retained_on_boundary_recovery`；[`text-evidence-multimodal-resolution.md`](../../narrative-video-production/text-evidence-multimodal-resolution.md) §Evidence non-destructive resolution |
| 3d-character | [`candidate-record.yaml`](../../3d-character-production/records/candidate-record.yaml) `retention.rejected_must_retain` |
| translation | [`translation-decision.yaml`](../../translation/contracts/translation-decision.yaml) `candidates[]` 保留 infeasible 候選與 `reason`；forbidden `collapsing_candidate_space_into_single_forced_output` |

### TE2 — 來源錨點

每筆證據都要指回**可重查**的來源，並帶上足以判斷是否過時的版本資訊。缺項要明確標示（例如 `unavailable`），不得捏造。

- **反例**：只寫「依官方範本」不寫版本；只寫日期不寫 commit；補一個看起來合理的 seed 或設定值。

| Domain | 實作 |
| --- | --- |
| narrative-video | `semantic_candidate.source_ref` 必填；`resolution.sources[]`；region 的 `timestamp`／`frame`／`box` |
| 3d-character | `candidate-record.yaml` `provenance`（缺 seed 標 `unavailable`，禁止捏造）；[`identity-acceptance.yaml`](../../3d-character-production/records/identity-acceptance.yaml) `evidence.rubric_applied` + `fixed_views` |
| security-audit | [`security-finding-list-template.md`](../../software-delivery/templates/security-finding-list-template.md) `evidence.reproducible[]`；[`security-coverage-ledger.md`](../../../analysis/security/security-coverage-ledger.md) `evidence_ref` + `verified_at_ref`（不可只寫日期） |
| legal | [`artifact-gates.md`](../../legal/artifact-gates.md) `gate.legal.law_citation_versioned`、`gate.legal.source_version_pinned` |

**錨點種類**（只列已有實例使用的種類；新來源先試著組合，不要直接發明新種類）：

| 種類 | 組成 | 實例 |
| --- | --- | --- |
| Repo artifact | path + commit 或版本 | security-audit |
| 外部文件 | 名稱 + 版本或修正日 + 查核日 | legal |
| 時間媒體 | 來源 + timestamp／frame + 區域（bbox） | narrative-video |
| 模型產出 | provider + model（+ version）+ 輸入引用 + timestamp | 3d-character |

例：視覺文件檢索的結果是「文件 + 版本 + 頁碼 + 區域」，可以由「外部文件」加上「區域」組合出來，不需要新增種類。

### TE4 — 定案要附理由

候選走到 resolved／accepted／confirmed／refuted 這類定案狀態時，必須附上**理由**與**依據的證據引用**。只有結論、沒有理由的定案無效。

- **反例**：resolved 但沒有 `reason`／`sources`；用「LLM 第二意見」refute；只寫「看起來比較自然」而不列依據的 artifact。

| Domain | 實作 |
| --- | --- |
| narrative-video | gate `resolution_reason_required_when_resolved`；四態都有最低 trace（`reason`／`reasons[]`／`reason.code`／`merged_into`） |
| 3d-character | [`artifact-gates.yaml`](../../3d-character-production/records/artifact-gates.yaml) promotion 要求 `decision` 與 `reason` 齊；forbidden `promote_when_decision_or_reason_missing` |
| security-audit | finding Status 規則：`confirmed`／`refuted` 需要 `reproducible` 證據；`resolution.kind` + `ref`；`risk_acceptance.decision_ref` + `owner` |
| translation | `selection.policy`（I6）、`rationale`、`decision_basis` 必列 artifact；[`finality.yaml`](../../translation/contracts/finality.yaml) `needs_review` 必須有非空 `review.reason[]` |
| legal | `gate.legal.strategy_reasoned`（四欄，來自 [`decision-support`](../decision-support/README.md)） |

### TE5 — Lifecycle 狀態 ≠ 可信度

「證據走到哪個狀態」（lifecycle：observed → candidate → resolved／rejected）與「證據有多可信」（epistemic：直接觀察、推導、推論，可重現與否）是**兩個獨立維度**，要分開記錄。

- resolved 不代表權重比 observed 高；它可能仍然只是推論。
- rejected 不代表原始觀察失效；被拒的可能是解讀，而不是觀察本身。
- 不確定本身不是狀態轉移的理由；要轉狀態，必須有明確的理由（TE4）。

**反例（anti-pattern）**：用一個 enum 同時表達兩個維度。例如 legal 的 Confidence 標籤 `provisional` 同時代表「還在 Strategy Pass 1」（lifecycle）和「前提未查證」（epistemic）；讀者無法分辨一條 `provisional` 的建議是「流程還沒走到」還是「查了但證據不足」。

| Domain | 實作 |
| --- | --- |
| narrative-video | §兩種否定：`rejected`（觀察不成立）≠ `uncertain`（觀察成立、解讀未收斂）；`confidence.type` 另立欄位；Evidence layer ≠ Production layer |
| security-audit | 模板 §三個判斷分開（schema valid／evidence established／closure allowed）；`reproducible` vs `reasoning` 證據；`hypothesis_source` 不是裁決證據 |
| translation | finality I22：uncertainty alone ≠ needs_review；forbidden `quality_or_confidence_scalar_as_closure` |

部分實作（不計入門檻）：3d-character 把 `decision`／`validity`／`maturity` 分軸並禁止 `identity_score`，但沒有獨立的可信度軸；legal 的來源層級 A–E 是可信度軸，但 Confidence 標籤混了兩個維度（見上方反例）。

**與 evidence-hierarchy 的關係**：lifecycle 三層**不能**一對一對應到 hierarchy 的 Observability 軸（direct／derived／memory／inference）。兩軸正交；加權規則仍由 hierarchy 擁有。

## Watch

### TE3 — 來源變更即失效（watch，不是必要條件）

先前的結論依賴某些來源；來源一變，結論就要轉為「待重驗」，重驗前不得被當作仍然有效。目前只有 2 個獨立實例，**未達門檻**，依 maintainer 決定不放寬定義來湊數。

| Domain | 觸發形狀 | 實作 |
| --- | --- | --- |
| security-audit | **依賴變更觸發**：宣告依賴範圍，變更命中就轉態 | ledger §Invalidation（`depends_on_controls` + `source_scope` → `needs_revalidation`；可機械回放）；`risk_acceptance.expires_when` |
| 3d-character | **依賴變更觸發** | `identity-acceptance.yaml` `validity: current｜re_review_required｜stale` + `invalidation_rules` + `mutation_event` |
| legal（variant，不計入） | **使用時重驗**：外部來源無法監看，每次引用前重查並更新查核日 | [`research/README.md`](../../legal/research/README.md) Staleness rule |

選擇觸發形狀的參考：來源在 repo 內、可以偵測變更 → 依賴變更觸發；來源在外部、無法監看 → 使用時重驗。抽象原則見 [`stale-derived-state.md`](../../../intelligence/engineering/anti-patterns/stale-derived-state.md)。

## 觀察中（未達 watch）

- **禁止純量分數當定案依據**：3d-character 禁 `identity_score`／`scalar_quality`，translation 禁 `quality_or_confidence_scalar_as_closure`。目前 2 個 domain。ECS 禁 confidence 數字屬 governance 層，不計入。

## Domain 如何採用

1. 對照 TE1、TE2、TE4、TE5，在自己的 record／gate 中找出對應欄位；沒有對應時，建議補上或寫明不適用的理由。
2. 錨點先從上方「錨點種類」組合；確實無法組合時，在 domain 文件寫明新種類與理由。
3. 在 domain 文件加一行連回本檔，並標註實作了哪幾條 TE。
4. 不必改既有欄位名。

## 驗證

- 原始觀察與落選候選是否保留，且附理由？（TE1）
- 每筆證據是否指得回來源與版本？缺項是否明確標示而非捏造？（TE2）
- 每次定案是否附理由與證據引用？（TE4）
- lifecycle 狀態與可信度是否分開記錄？有沒有單一 enum 承載兩個維度？（TE5）
- 若 domain 宣告實作 TE3：觸發形狀是否與來源可監看性相符？

行為鎖定：[`traceable-evidence-new-domain-adoption-v1`](../../../validation/scenarios/cross-domain/traceable-evidence-new-domain-adoption-v1.yaml)（新 domain 採用時的 TE1、TE5 與 watch 誤用）。
