---
id: 2026-10-04-0936-evidence-record-contract
plan_kind: main
status: completed
owner: linyihong
created: 2026-10-04
priority: P2
parent: null
required_for_completion: false
---

# Evidence Record Contract（domain evidence 記錄形狀的跨域收斂）

**Status**: `completed`（2026-10-04）：Phase 0–3 完成。pattern 落地於 [`workflow/cross-cutting/traceable-evidence/`](../../../workflow/cross-cutting/traceable-evidence/README.md)（TE1／TE2／TE4／TE5 invariant + TE3 watch，`pilot`）；六個 domain link back；1 個 cross-domain scenario。不升 ADR（見 §ADR Promotion Criteria 結案評估）。
Owner: framework maintainer (linyihong)
**建立日期**：2026-10-04
**Priority**：**P2**（不阻擋其他工作；是 visual retrieval research candidate 的前置）

**Glossary Impact**: no：pattern 名稱 `traceable-evidence` 已確認（Q1），但 pilot 期間不登記 glossary（比照 decision-support）；升級出 pilot 時再評估。

## Executive summary

多個 workflow 各自長出了相同的 **domain evidence 記錄形狀**：保留原始觀察、附來源錨點、從候選到定案要附理由、來源變更時要失效。這些形狀散落在各 domain，沒有共用契約，新 domain 只能重新發明一次。

本 plan **先驗證收斂，再抽取**：Phase 0 逐一核對現有實例，只有被 **≥3 個 domain 獨立實作**的 invariant，才會抽成 `workflow/cross-cutting/` 的 pattern。各 domain 再連回這份 pattern，不強制改欄位名。

觸發來源：2026-10-04 評估 ColPali 類 visual document retrieval（見 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../../architecture/ai-native-cognitive-ecosystem-system.md) §Research Candidates）。結論是 provider 介面要等 consumer 出現後再定；但不管將來接哪種感知或檢索來源，結果都必須以 candidate 身分進入同一個 evidence 流程。這個流程的記錄形狀現在就能從既有實例收斂出來，而且與模型無關。

## 邊界（與既有 evidence 機制的分工）

`evidence` 這個詞在 repo 內已經有多個 owner（ECS EL-4 曾經因為同名而誤判成重複）。本 plan 的範圍：

| 機制 | 回答的問題 | 與本 plan 的關係 |
|---|---|---|
| [`enforcement/evidence-hierarchy.md`](../../../enforcement/evidence-hierarchy.md) | 證據如何**比較與加權**（authority／freshness／validity／scope／observability） | 本 plan 定義**記錄形狀**，不重新定義加權規則；Phase 0 F1 確認 lifecycle 三層**不能**一對一對應到 Observability 軸 |
| [`governance/evidence-candidates/`](../../../governance/evidence-candidates/README.md)（ECS） | 哪個觀察要回流到**哪個 plan**（inter-plan governance routing） | 不同 layer：ECS 明確排除 domain evidence（見 ECS plan applicability 表）；本 plan 只處理 domain evidence |
| [`governance/lifecycle/plan-evidence.md`](../../../governance/lifecycle/plan-evidence.md) | 單一 plan 的 run 全文**存放在哪** | 不同 layer：plan 的儲存慣例，不涉及 domain 證據形狀 |
| [`intelligence/engineering/anti-patterns/stale-derived-state.md`](../../../intelligence/engineering/anti-patterns/stale-derived-state.md) | derived state 為什麼需要 invalidation contract（原則） | I3 目前是 Watch；watch 條目會指回這份原則 |

## Phase 0 結果（全文：[`evidence/phase0-convergence-audit.md`](evidence/phase0-convergence-audit.md)）

5 個 domain（narrative-video、3d-character、security-audit、legal，以及 Phase 0 新增的 translation）。每條 invariant 獨立計算，只有 ✅ 計入，共用上游契約的實例合算 1。

| Invariant | ✅ domain | 處置 |
|---|---|---|
| I1 保留原始觀察 | NV、3D、TR（3） | **Promote**。敏感度：若 NV／TR 改判不獨立 → Watch |
| I2 來源錨點 | NV、3D、SEC、LG（4） | **Promote** |
| I3 來源變更即失效 | SEC、3D（2）+ LG variant | **Watch**（F2：使用時重驗是另一種觸發形狀，未計入） |
| I4 定案要附理由 | NV、3D、SEC、TR、LG（5） | **Promote** |
| I5 候選 lifecycle ≠ epistemic | NV、SEC、TR（3） | **Promote**（maintainer 2026-10-04 決定納入；pattern 編號 TE5） |

獨立性：3D／NV／TR 共用 `artifact-record/v1` 版本標籤，但 repo 內沒有定義檔，標籤不含 invariant 內容；NV↔TR 有詞彙與資料流耦合，但依時間序與事故來源判定為各自寫出。細節見 audit §獨立性檢查。

## Decision Rationale

### Problem & Why Now

1. **重複發明**：多個 domain 用不同欄位名表達相同的約束；新 domain（例如將來的 visual retrieval consumer）沒有參考點，只能再發明一次。
2. **模型可替換性的真正支點**：外部評估把可替換性放在「Retrieval Provider 介面」；但對本系統來說，比較穩的支點是「感知或檢索結果一律以 candidate 進入、帶來源錨點、定案要附理由」。這一點現在就能從實例收斂，不需要等任何模型。
3. **時機**：實例已經自然出現，不是預先設計；收斂驗證的成本低（只讀現有檔案）。

### Decision

- Phase 0 逐一核對實例，產出有檔案引用的收斂矩陣（**done**）。
- 只抽取 ≥3 個獨立實例的 invariant，寫成 `workflow/cross-cutting/<name>/`（cross-cutting pattern，比照 [`decision-support/`](../../../workflow/cross-cutting/decision-support/README.md) 的形式：domain instantiate + link back）。
- 未達門檻者列為 **watch**：可參考，但**不是**新 domain 的必要條件；新 domain 缺 watch 項不算違反 pattern，除非它自己的 domain contract 已經要求。
- source anchor 的形狀要 **modality-neutral**：只涵蓋實例已經用到的錨點種類（F4：visual retrieval 可以由既有種類組合，不需要新種類）。
- **不接 runtime**、不加 validator、不強制 domain 改欄位名。

### Alternatives Considered

- **A. 先定義 Retrieval Provider 介面**：reject。沒有 consumer，形狀只能猜（見 §Research Candidates 評估要點）。
- **B. 寫進 `enforcement/evidence-hierarchy.md`**：reject。hierarchy 處理加權（P1 enforcement），記錄形狀是 domain 層的 pattern；硬塞進去會把 domain 欄位變成 repo-wide enforcement。
- **C. 擴充 ECS**：reject。ECS 是 inter-plan governance routing，已明確排除 domain evidence（narrow-applicability）。
- **D. 直接建 schema + commit-msg validator**：reject。違反 doc-only trial 先行的紀律。
- **E. Convergence audit → 只抽取已收斂的 invariant，作為 cross-cutting pattern**：**accept**。

### Why Not an ADR Yet

名稱與 I5 是否納入尚未拍板；而且有更輕的 promotion target（cross-cutting pattern）可用。

### ADR Promotion Criteria（completed 時驗證）

- [x] Open Questions 全解 — *Q1–Q10 皆 resolved（2026-10-04）*

結案評估（2026-10-04）：**不升 ADR，`adr_promotion: deferred`**。未達的條件如下（以清單記錄，不是待辦）：

- 尚無**新** domain（不在 Phase 0 矩陣內）直接採用此 pattern；外推能力未證明。
- foundational + expensive-to-reverse 不成立：pattern 是 `pilot` reference contract，可回退，不約束任何 gate。
- 有更輕的 promotion target：留在 `workflow/cross-cutting/` 即可。

重新評估時機：第一個新 domain 採用 traceable-evidence（最可能是 Gen 4 §Research Candidates 的 visual document retrieval consumer），且採用時不需要修改 pattern。

### Consequences

#### 正面
- 新 domain 有參考點；visual retrieval 等未來感知來源有明確的接入形狀
- 各 domain 的 evidence 約束可以互相對照，容易發現某個 domain 漏了 invariant

#### 負面
- 多一份 cross-cutting 文件要維護；domain link-back 需要隨 domain 演進更新

#### 風險
- **同名誤判**（EL-4 重演）：用 §邊界表與 Q1 命名處置
- **過早抽象**：用「每條 invariant 獨立 ≥3 個獨立實例」擋住
- **強迫統一欄位名**：明確非目標；pattern 定義語意，不定義欄位名
- **`artifact-record/v1` 日後被正式化**：屆時 3D／NV／TR 會變成共用上游，要重新計算獨立性

## Runtime Execution Path

**本 plan 不接入 runtime。** 不新增 `route.*`，不 project 到 `runtime.db generated_surfaces`，也不加 commit-msg validator。pattern 由 domain 文件 link back 來消費。

未來接入條件：某個 domain 的 artifact gate 需要機械驗證某條 invariant（例如「resolved 缺 sources 就擋」），而且該 domain 已有 validator 基礎設施時，在該 domain 的 plan 內評估；這不由本 plan 預建。Graduation deadline：Phase 1 須在 2026-12-31 前完成，否則重新評估本 plan 是否還有必要。

**Per-surface consumer 表**：N/A：本 plan 不新增任何 generated surface、route 或 validator。

## Watch-Out List citation

對應 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../../architecture/ai-native-cognitive-ecosystem-system.md) §Watch-Out List **Wall 2（Workflow inflation）**：pattern 只定義語意 invariant，不成為 domain 必經 stage，也不要求欄位統一；watch 項不構成義務。

## Open Questions

| # | Question | 傾向 | 處置 |
|---|---|---|---|
| Q1 | 名稱 | maintainer 暫定 `source-anchored-evidence`，並授權「來源錨點只是其中一個特徵時可以重新命名」 | **resolved（`traceable-evidence`，maintainer 2026-10-04 確認）**：暫定名只覆蓋 I2；`traceable` 涵蓋保留（TE1）、錨定（TE2）、定案理由（TE4）與分軸記錄（TE5），且不以 `evidence-` 開頭，降低與 `evidence-candidates`／`plan-evidence` 混淆。|
| Q2 | 歸屬：`workflow/cross-cutting/` 或 `intelligence/engineering/`？ | cross-cutting：它約束的是 domain artifact 形狀，比照 decision-support | **resolved（cross-cutting）**：5 個實例全部是 domain artifact 的 record／gate，不是判斷準則 |
| Q3 | I3 未達 3 個實例時，是否仍在 pattern 內列出？ | — | **resolved（maintainer 2026-10-04）**：保留為 watch，不升 invariant；watch 不是新 domain 的必要條件。Phase 0 確認 I3 = 2 + 1 variant |
| Q4 | domain link-back 的深度 | 只加連結 + 標註對應的 invariant id，不改欄位 | **resolved（Phase 2）**：每個 domain 一行，標註 TE id，未改任何欄位或 gate |
| Q5 | observed／candidate／resolved 三層與 evidence-hierarchy Observability 軸如何對應？ | — | **resolved（不能一對一）**：Phase 0 F1，resolved 不代表權重較高、rejected 不代表觀察失效；Phase 1 寫成「兩軸正交」說明，不改 hierarchy |
| Q6 | 是否需要 validation scenarios？ | 依 decision-support 先例：pattern 本身 doc-only，由 domain 既有 scenarios 覆蓋 | **resolved（補 1 個）**：decision-support 先例其實有 1 個 cross-domain heuristic scenario；比照新增 [`traceable-evidence-new-domain-adoption-v1`](../../../validation/scenarios/cross-domain/traceable-evidence-new-domain-adoption-v1.yaml)，鎖定新 domain 採用時最容易誤用的 TE1、TE5 與「watch 不是必要條件」 |
| Q7 | I5（lifecycle ≠ epistemic）達 3 個實例，是否納入 pattern？ | 納入，並以 LG 的 `provisional` 雙義 enum 當 anti-pattern 範例 | **resolved（納入，TE5）**：maintainer 2026-10-04 |
| Q8 | I3 是否放寬為「宣告重驗觸發條件（變更觸發或使用時觸發）」？放寬後會達 3 | 不放寬：這是為了過門檻改寫定義；兩種形狀都記在 watch 條目 | **resolved（不放寬）**：maintainer 2026-10-04；TE3 維持 watch，legal 列為 variant |
| Q9 | 獨立性判定 | — | **resolved（maintainer 2026-10-04）**：共用同一份上游契約的實例合算 1；NV 與 TR 判定獨立（I1 維持 3） |
| Q10 | pattern 編號與 domain 內部編號衝突（translation 已用 I1–I23） | 加前綴 | **resolved（`TE` 前綴）**：TE1–TE5 對應 audit 的 I1–I5；全文搜尋無衝突 |

## 完成條件

- [x] Phase 0 收斂矩陣每格都有檔案引用或標明 `absent`
- [x] 只抽取 ≥3 實例的 invariant；未達門檻者有明確處置（TE3 watch；Q3／Q7／Q8 resolved）
- [x] cross-cutting pattern 文件落地，含邊界表、實例表、watch 段
- [x] 所有實例 domain 已 link back；cross-cutting README 表格已更新
- [x] §Research Candidates 的前置欄已更新為 met
- [x] 執行 Plan Completion Closure（2026-10-04）

## Phase 0: Pre-Build Interrogation ✅（2026-10-04）

### Phase 0.0 — Open Questions 核對（公版，必填）

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [x] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [x] 若盤點新發現問題，已加入 §Open Questions（Q7、Q8、Q9）

| Open Question | 處置 | 證據 / 原因 |
|---|---|---|
| Q1 名稱 | still-open | 暫定名只覆蓋 I2，見 audit §結論 |
| Q2 歸屬 | resolved | 實例全是 domain record／gate |
| Q3 I3 處置 | resolved | maintainer 決策 + audit §I3 |
| Q4 link-back 深度 | still-open | Phase 2 |
| Q5 Observability 對照 | resolved | audit §研究發現 F1 |
| Q6 scenarios | still-open | Phase 3 |

### Phase 0.1 — Architecture Compatibility Preflight

| 欄位 | 內容 |
|---|---|
| Trigger | 開始 Phase 0（maintainer 2026-10-04 批准） |
| Checked sources | `workflow/cross-cutting/README.md`（三案例 promotion policy）、`decision-support/README.md`（instantiation 形式）、evidence-hierarchy、ECS plan、5 個 domain 實例 |
| Conflicts | 無。新發現：`artifact-record/v1` 標籤共用（不構成上游）；plan 改為 folder 形式以容納 `evidence/` |
| Interrogation | goal＝驗證收斂；non-goals＝runtime、validator、domain schema、Retrieval Provider；acceptance＝每格有引用 |
| Decision | proceed to Phase 1 after maintainer decides Q1／Q7／Q8 |
| Validation | runtime compile／refresh／validate；連結存在性檢查 |

- [x] 依 [`plans/README.md`](../../README.md) §Architecture Compatibility Preflight 完成最低記錄格式
- [x] 讀 [`workflow/cross-cutting/README.md`](../../../workflow/cross-cutting/README.md) 的 promotion policy 與 decision-support 的實例形式

### Phase 0.2 — Convergence audit

- [x] 逐格驗證收斂矩陣，每格附檔案引用，或標 `absent`
- [x] 掃描其他候選實例：納入 translation；排除 investment（共用 decision-support 上游）、delegated-execution（屬 security-audit 同一 capability）
- [x] 產出結論：I1／I2／I4 Promote、I3 Watch、I5 待定

## Phase 1: Pattern 文件 ✅（2026-10-04）

- [x] 建立 [`workflow/cross-cutting/traceable-evidence/README.md`](../../../workflow/cross-cutting/traceable-evidence/README.md)：每條已收斂 invariant 寫明語意、反例，以及實例表（domain → 欄位或 gate id）
- [x] Watch 段：TE3（含兩種觸發形狀）；明寫「watch 不是新 domain 的必要條件」
- [x] I5 納入為 TE5（Q7）；Q5 寫成兩軸正交說明
- [x] source anchor 種類表（只列實例已使用的錨點種類）與 §邊界表
- [x] 檢查 document sizing（單一主題，約 95 行，未超過 300 行警戒線，不拆分）

## Phase 2: Domain link-back 與索引 ✅（2026-10-04）

- [x] 各實例 domain 文件加 link back，並標註對應的 TE id（Q4）：NV `text-evidence-multimodal-resolution.md`、3D `records/README.md`、translation `README.md`、security ledger + finding list template、legal `artifact-gates.md`（TE5 反例只記為觀察，不構成 legal 的修改義務）
- [x] 更新 `workflow/cross-cutting/README.md` 的 Current concerns 表與 `workflow/README.md`（隨 Phase 1 commit，屬新文件的 linked update）
- [x] 在 `stale-derived-state.md` 加連結（TE3 watch，含兩種觸發形狀）
- [x] 更新 §Research Candidates 的前置欄狀態（met）

## Phase 3: Closure ✅（2026-10-04）

- [x] 依 Q6 補 1 個 scenario（`validation/scenarios/cross-domain/traceable-evidence-new-domain-adoption-v1.yaml`）
- [x] runtime compile／refresh／validate（無新增 route／surface；scenario 計入既有 orphan scenarios 統計，比照 decision-support）
- [x] 執行 Plan Completion Closure（2026-10-04）

## Stakeholder 同意項目

- [x] maintainer 確認範圍只限 domain evidence 記錄形狀，不碰 hierarchy、ECS、plan-evidence（2026-10-04 批准 Phase 0 時維持原定範圍）
- [x] maintainer 確認 ≥3 實例門檻適用於每條 invariant（2026-10-04）
- [x] maintainer 拍板 Q7（納入 I5）與 Q8（I3 不放寬）（2026-10-04）
- [x] maintainer 拍板 Q1 名稱（`traceable-evidence`，2026-10-04）

## 與其他 plans 的關係

- [`../2026-06-16-1131-evidence-candidate-system.md`](../../active/2026-06-16-1131-evidence-candidate-system.md)：不同 layer（inter-plan routing），見 §邊界；本 plan 不 feed ECS。
- [`../2026-09-16-1649-narrative-video-production-workflow/_plan.md`](../../active/2026-09-16-1649-narrative-video-production-workflow/_plan.md)：I1／I2／I4／I5 實例來源。
- [`../2026-08-31-1032-3d-character-production-workflow/_plan.md`](../../active/2026-08-31-1032-3d-character-production-workflow/_plan.md)：I1／I2／I3／I4 實例來源；`artifact-record/v1` 標籤起源。
- [`../2026-09-22-1000-translation-decision-workflow/_plan.md`](../../active/2026-09-22-1000-translation-decision-workflow/_plan.md)：I1／I4／I5 實例來源（Phase 0 新增）。
- [`../../archived/2026-10-03-2104-security-audit-capability-hardening/_plan.md`](../2026-10-03-2104-security-audit-capability-hardening/_plan.md)：I2／I3／I4／I5 實例來源。
- [`../2026-07-30-2101-legal-workflow-domain.md`](../../active/2026-07-30-2101-legal-workflow-domain.md)：I2／I4 實例來源；I3 variant；I5 反例形狀。
