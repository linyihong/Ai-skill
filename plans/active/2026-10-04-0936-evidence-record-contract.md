---
id: 2026-10-04-0936-evidence-record-contract
plan_kind: main
status: draft
owner: linyihong
created: 2026-10-04
priority: P2
parent: null
required_for_completion: false
---

# Evidence Record Contract（domain evidence 記錄形狀的跨域收斂）

**Status**: `draft`：Phase 0（convergence audit）尚未開始。
Owner: framework maintainer (linyihong)
**建立日期**：2026-10-04
**Priority**：**P2**（不阻擋其他工作；是 visual retrieval research candidate 的前置）

**Glossary Impact**: deferred：名稱是 §Open Questions Q1，本 draft 不註冊任何 term；Phase 1 定名後再決定是否登記到 `knowledge/glossary/ai-skill.md`。

## Executive summary

多個 workflow 各自長出了相同的 **domain evidence 記錄形狀**：保留原始觀察、附來源錨點、從候選到定案要附理由、來源變更時要失效。這些形狀散落在各 domain，沒有共用契約，新 domain 只能重新發明一次。

本 plan **先驗證收斂，再抽取**：Phase 0 逐一核對現有實例，只有被 **≥3 個 domain 獨立實作**的 invariant，才會抽成 `workflow/cross-cutting/` 的 pattern。各 domain 再連回這份 pattern，不強制改欄位名。

觸發來源：2026-10-04 評估 ColPali 類 visual document retrieval（見 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Research Candidates）。結論是 provider 介面要等 consumer 出現後再定；但不管將來接哪種感知或檢索來源，結果都必須以 candidate 身分進入同一個 evidence 流程。這個流程的記錄形狀現在就能從既有實例收斂出來，而且與模型無關。

## 邊界（與既有 evidence 機制的分工）

`evidence` 這個詞在 repo 內已經有多個 owner（ECS EL-4 曾經因為同名而誤判成重複）。本 plan 的範圍：

| 機制 | 回答的問題 | 與本 plan 的關係 |
|---|---|---|
| [`enforcement/evidence-hierarchy.md`](../../enforcement/evidence-hierarchy.md) | 證據如何**比較與加權**（authority／freshness／validity／scope／observability） | 本 plan 定義**記錄形狀**，不重新定義加權規則；記錄欄位應該讓五軸可以被判讀 |
| [`governance/evidence-candidates/`](../../governance/evidence-candidates/README.md)（ECS） | 哪個觀察要回流到**哪個 plan**（inter-plan governance routing） | 不同 layer：ECS 明確排除 domain evidence（見 ECS plan applicability 表）；本 plan 只處理 domain evidence |
| [`governance/lifecycle/plan-evidence.md`](../../governance/lifecycle/plan-evidence.md) | 單一 plan 的 run 全文**存放在哪** | 不同 layer：plan 的儲存慣例，不涉及 domain 證據形狀 |
| [`intelligence/engineering/anti-patterns/stale-derived-state.md`](../../intelligence/engineering/anti-patterns/stale-derived-state.md) | derived state 為什麼需要 invalidation contract（原則） | 本 plan 若收斂出 invalidation invariant，就是這個原則的記錄形狀實作 |

## 初步收斂矩陣（Phase 0 必須逐格驗證，以下只是 draft 判讀）

| Invariant | narrative-video text evidence | 3d-character candidate record | security coverage ledger | legal reference sources |
|---|---|---|---|---|
| **I1 保留原始觀察**（observed／raw 不得被 derived 覆寫；落選者保留） | ✅ `observed_candidate_resolved_layers`、`raw_text_retained_on_boundary_recovery` | ✅ `retention.rejected_must_retain` | ? `needs_revalidation` 是否保留舊 evidence_ref | ? |
| **I2 來源錨點**（每筆證據指回可重查的來源與版本） | ✅ `source_ref`、bbox／timestamp | ✅ `provenance`（provider／model／inputs／timestamp） | ✅ `evidence_ref` + `verified_at_ref`（commit，不可只寫日期） | ✅ 版本 + 查核日 |
| **I3 來源變更即失效**（宣告依賴範圍，來源變更就轉為待重驗） | ? | ? | ✅ §Invalidation + 機械回放 | ? stale 警示，是否有轉態規則 |
| **I4 定案要附理由**（candidate → resolved／accepted 必須附 reason 與 sources） | ✅ `resolution_reason_required_when_resolved` | ✅ 缺 decision／reason 就禁止 promotion | ? `accepted_gap` 的 `decision_ref` | ? Decision Reasoning 四欄是否涵蓋 |

Draft 判讀：I2 可能 4/4、I4 可能 3/4、I1 可能 2–3/4、I3 可能只有 1–2/4。**未達 3 個實例的 invariant 不抽取**，留在原 domain（I3 目前由 security ledger + stale-derived-state 持有）。

## Decision Rationale

### Problem & Why Now

1. **重複發明**：四個 domain 用不同欄位名表達相同的約束；新 domain（例如將來的 visual retrieval consumer）沒有參考點，只能再發明一次。
2. **模型可替換性的真正支點**：外部評估把可替換性放在「Retrieval Provider 介面」；但對本系統來說，比較穩的支點是「感知或檢索結果一律以 candidate 進入、帶來源錨點、定案要附理由」。這一點現在就能從實例收斂，不需要等任何模型。
3. **時機**：實例已經自然出現 4 個，不是預先設計；收斂驗證的成本低（只讀現有檔案）。

### Decision

- Phase 0 逐一核對實例，產出有檔案引用的收斂矩陣；再掃描是否有其他實例。
- 只抽取 ≥3 實例的 invariant，寫成 `workflow/cross-cutting/<name>/`（cross-cutting pattern，比照 [`decision-support/`](../../workflow/cross-cutting/decision-support/README.md) 的形式：domain instantiate + link back）。
- source anchor 的形狀要 **modality-neutral**：涵蓋現有實例已經用到的錨點種類（path + commit、URL + 版本 + 查核日、timestamp + bbox），不預先發明實例沒有用過的欄位。
- **不接 runtime**、不加 validator、不強制 domain 改欄位名。

### Alternatives Considered

- **A. 先定義 Retrieval Provider 介面**：reject。沒有 consumer，形狀只能猜（見 §Research Candidates 評估要點）。
- **B. 寫進 `enforcement/evidence-hierarchy.md`**：reject。hierarchy 處理加權（P1 enforcement），記錄形狀是 domain 層的 pattern；硬塞進去會把 domain 欄位變成 repo-wide enforcement。
- **C. 擴充 ECS**：reject。ECS 是 inter-plan governance routing，已明確排除 domain evidence（narrow-applicability）。
- **D. 直接建 schema + commit-msg validator**：reject。N 尚未驗證，而且違反 doc-only trial 先行的紀律。
- **E. Convergence audit → 只抽取已收斂的 invariant，作為 cross-cutting pattern**：**accept**。

### Why Not an ADR Yet

收斂本身還沒驗證（矩陣仍是 draft 判讀），名稱與歸屬也未定（Q1／Q2）；而且有更輕的 promotion target（cross-cutting pattern）可用。

### ADR Promotion Criteria（completed 時驗證）

- [ ] ≥1 個**新** domain（不在 Phase 0 矩陣內）直接採用此 pattern，並且不需要修改 pattern（證明有外推能力）
- [ ] foundational + cross-session + cross-project + expensive-to-reverse + explains-why 全中
- [ ] Open Questions 全解
- [ ] 沒有更輕的 promotion target（per ADR-007）

### Consequences

#### 正面
- 新 domain 有參考點；visual retrieval 等未來感知來源有明確的接入形狀
- 各 domain 的 evidence 約束可以互相對照，容易發現某個 domain 漏了 invariant

#### 負面
- 多一份 cross-cutting 文件要維護；domain link-back 需要隨 domain 演進更新

#### 風險
- **同名誤判**（EL-4 重演）：用 §邊界表與 Q1 命名處置
- **過早抽象**：用 ≥3 實例門檻與 Phase 0 逐格驗證擋住
- **強迫統一欄位名**：明確非目標；pattern 定義語意，不定義欄位名

## Runtime Execution Path

**本 plan 不接入 runtime。** 不新增 `route.*`，不 project 到 `runtime.db generated_surfaces`，也不加 commit-msg validator。pattern 由 domain 文件 link back 來消費。

未來接入條件：某個 domain 的 artifact gate 需要機械驗證某條 invariant（例如「resolved 缺 sources 就擋」），而且該 domain 已有 validator 基礎設施時，在該 domain 的 plan 內評估；這不由本 plan 預建。Graduation deadline：Phase 0 須在 2026-12-31 前完成，否則重新評估本 plan 是否還有必要。

**Per-surface consumer 表**：N/A：本 plan 不新增任何 generated surface、route 或 validator。

## Watch-Out List citation

對應 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List **Wall 2（Workflow inflation）**：pattern 只定義語意 invariant，不成為 domain 必經 stage，也不要求欄位統一。

## Open Questions

| # | Question | 傾向 | 處置 |
|---|---|---|---|
| Q1 | 名稱：`evidence-record` 與 ECS／plan-evidence 同用 "evidence"，會不會重演 EL-4 誤判？候選：`evidence-record`／`evidence-provenance`／`source-anchored-evidence` | 取最抽象、又不與既有 owner 撞名的名稱；Phase 0 結束時依實際收斂的 invariant 決定 | open |
| Q2 | 歸屬：`workflow/cross-cutting/`（domain instantiate）或 `intelligence/engineering/`（判斷準則）？ | cross-cutting：它約束的是 domain artifact 形狀，比照 decision-support | open |
| Q3 | I3（invalidation）若未達 3 個實例，是否仍在 pattern 內以「watch」列出？ | 列為 watch、不列 invariant，並指回 security ledger + stale-derived-state | open |
| Q4 | domain link-back 的深度：只加連結，還是要求 domain 文件標註「instantiates I-x」？ | 只加連結 + 標註對應的 invariant id，不改欄位 | open |
| Q5 | observed／candidate／resolved 三層與 evidence-hierarchy 的 Observability 軸（direct／derived／memory／inference）如何對應？ | Phase 1 寫成對照表，不改 hierarchy | open |
| Q6 | 是否需要 validation scenarios？ | 依 decision-support 的先例決定；若 pattern 是 doc-only，則由 domain 既有 scenarios 覆蓋 | open |

## 完成條件

- [ ] Phase 0 收斂矩陣每格都有檔案引用或標明 `absent`
- [ ] 只抽取 ≥3 實例的 invariant；未達門檻者有明確處置（Q3）
- [ ] cross-cutting pattern 文件落地，含邊界表與實例表
- [ ] 所有實例 domain 已 link back；cross-cutting README 表格已更新
- [ ] §Research Candidates 的前置欄已更新為 met
- [ ] 執行 Plan Completion Closure

## Phase 0: Pre-Build Interrogation

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [ ] 已讀本 plan §Open Questions 全部條目
- [ ] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [ ] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [ ] 若盤點新發現問題，已加入 §Open Questions

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1 名稱 | still-open | 待收斂結果 |
| Q2 歸屬 | still-open | 待收斂結果 |
| Q3 I3 處置 | still-open | 待矩陣驗證 |
| Q4 link-back 深度 | still-open | — |
| Q5 Observability 對照 | still-open | — |
| Q6 scenarios | still-open | 待讀 decision-support 先例 |

### Phase 0.1 — Architecture Compatibility Preflight

- [ ] 依 [`plans/README.md`](../README.md) §Architecture Compatibility Preflight 完成最低記錄格式
- [ ] 讀 [`workflow/cross-cutting/README.md`](../../workflow/cross-cutting/README.md) 的 promotion policy 與 decision-support 的實例形式

### Phase 0.2 — Convergence audit

- [ ] 逐格驗證收斂矩陣，每格附檔案引用（`path` + 欄位或 gate id），或標 `absent`
- [ ] 掃描其他候選實例（例如 narrative-video 的 non-destructive finalization feedback、software-delivery delegated execution、translation failure-pattern 的 provenance），納入或排除都要附理由
- [ ] 產出結論：哪些 invariant ≥3 個實例、哪些不足

## Phase 1: Pattern 文件

- [ ] 依 Q1／Q2 結論建立 `workflow/cross-cutting/<name>/README.md`：每條已收斂 invariant 寫明語意、反例，以及實例表（domain → 欄位或 gate id）
- [ ] source anchor 種類表（只列實例已使用的錨點種類）
- [ ] Observability 對照（Q5）與 §邊界表
- [ ] 檢查 document sizing

## Phase 2: Domain link-back 與索引

- [ ] 各實例 domain 文件加 link back，並標註對應的 invariant id（Q4）
- [ ] 更新 `workflow/cross-cutting/README.md` 的 Current concerns 表
- [ ] 視需要在 `stale-derived-state.md` 加連結（若 I3 被收錄或列為 watch）
- [ ] 更新 §Research Candidates 的前置欄狀態

## Phase 3: Closure

- [ ] 依 Q6 決定是否補 scenarios
- [ ] runtime compile／refresh／validate（確認沒有誤增 surface）
- [ ] 執行 Plan Completion Closure

## Stakeholder 同意項目

- [ ] maintainer 確認範圍只限 domain evidence 記錄形狀，不碰 hierarchy、ECS、plan-evidence
- [ ] maintainer 確認 ≥3 實例門檻適用於每條 invariant（而非整份 pattern）
- [ ] maintainer 拍板 Q1 名稱

## 與其他 plans 的關係

- [`active/2026-06-16-1131-evidence-candidate-system.md`](2026-06-16-1131-evidence-candidate-system.md)：不同 layer（inter-plan routing），見 §邊界；本 plan 不 feed ECS。
- [`active/2026-09-16-1649-narrative-video-production-workflow/_plan.md`](2026-09-16-1649-narrative-video-production-workflow/_plan.md)：observed／candidate／resolved 三層的主要實例來源。
- [`active/2026-08-31-1032-3d-character-production-workflow/_plan.md`](2026-08-31-1032-3d-character-production-workflow/_plan.md)：provenance 與 retention 實例來源。
- [`archived/2026-10-03-2104-security-audit-capability-hardening/_plan.md`](../archived/2026-10-03-2104-security-audit-capability-hardening/_plan.md)：invalidation 實例來源（coverage ledger）。
- [`active/2026-07-30-2101-legal-workflow-domain.md`](2026-07-30-2101-legal-workflow-domain.md)：版本 + 查核日的來源錨點實例。
