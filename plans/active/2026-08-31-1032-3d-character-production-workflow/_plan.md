---
id: 2026-08-31-1032-3d-character-production-workflow
plan_kind: main
status: in-progress
owner: workflow
owner_layer: workflow
created: 2026-08-31
parent: null
---

# 3D Character Production Workflow（`workflow/3d-character-production/`）

**Status**: in-progress — Phase 3 **implementation landed ≠ validated**（`6567ae4e`）。
Detector no_match **EXPECTED**。下一階段若繼續：execution dogfood（停／放行），不是再擴寫 workflow。
Phase 4／5 未授權。Invalidation SoT：`workflow/3d-character-production/records/identity-acceptance.yaml`。

**Glossary Impact**: yes — 候選詞 `character_identity_lock`、`asset_maturity_gate`、
`deformation_acceptance_set`、`runtime_ready_character_pack`。Phase 0 glossary collision
check（2026-08-31）：`knowledge/glossary/ai-skill.md` **無同名 entry**。唯一近鄰是
`change_intent_lock`（software-delivery implementation intent，非角色辨識）。四詞 **不**
與 `change_intent_lock` 合併；Phase 5 再登記，避免未落地 workflow 先膨脹 glossary。

## 一句話目標

建立一套工具中立、可重跑、可驗證的 3D 角色製作 workflow，將「參考圖丟進生成工具後憑感覺
反覆修改」改成「角色規格鎖定 → 候選生成 → Mesh QA → Rig → Face → Outfit/Animation →
格式匯出 → Runtime 驗收」的階段化生產流程。

首個 dogfood 使用既有 VRM 角色專案，但專案名稱、私有路徑、原始素材與一次性操作紀錄不進入
reusable workflow；workflow 只吸收可跨角色重用的 gate、artifact shape 與判斷規則。

## Decision Rationale

### Problem & Why Now

目前外部角色專案已經有參考圖、圖生 3D mesh、Blender 工程、Humanoid rig、VRM 匯出腳本、
表情腳本與動作清單，但成品仍反覆偏離目標。主要缺口不是「沒有建模工具」：

1. **角色一致性沒有被獨立驗收**：臉型、髮型、挑染、比例、服裝等辨識點散落在文件與圖片，
   每次生成容易把「像同一角色」和「技術上可匯出」混成一個模糊結果。
2. **生成成功被誤當成資產成功**：GLB/FBX/VRM 能輸出，不代表拓撲可修、肢體已分離、權重正確、
   表情可驅動或能承受目標動作。
3. **修補發生得太晚**：模型進入 rig、表情或動畫後才發現 mesh 不適合，造成高成本重做。
4. **沒有固定比較基線**：不同 run 的輸入、工具設定、輸出與 rejection reason 沒有統一 artifact，
   無法知道品質差異來自 reference、生成、retopology、rig 還是 export。
5. **現有 software-delivery workflow 不擁有角色資產品質**：它可以管理 Blender Python、
   exporter 或 runtime code，但不應定義角色辨識度、變形品質與 asset maturity。

### Decision

新增 `workflow/3d-character-production/`，domain 邊界是**可驅動 3D 角色資產的生產與驗收**，
不是泛稱的「AI 建模」，也不是純影片生成。

核心決策：

- AI 生成、手工建模與外包是可替換的 acquisition strategy，不是 workflow 名稱。
- 流程以 **asset maturity gate** 推進；任何階段失敗都回到最早失效 owner，不用下游腳本掩蓋。
- Identity 拆兩閘：**Specification / Reference Lock**（生成前 prerequisite）與
  **Identity Acceptance**（候選 attribute-level gate）。Lock ≠ 候選通過。
- Identity 是 attribute pass/fail + blocking list（must-preserve / allowed-variation /
  forbidden-drift），**禁止** scalar quality score。
- 每個候選保留 provenance、accepted/rejected、reason 與下一步，不只保存勝出的檔案。
- 可匯出只算格式 gate。第一版 completion = **`runtime-ready`**：所有 blocking gates PASS +
  target runtime contract PASS + known defects 已分類。**不是**「沒有任何問題」。
- VRM 是首個 profile；full-body 為 required，half-body 為 optional/derived。不把 workflow
  寫死為 VRM、Blender 或特定生成服務。第二個輸出 profile **defer** 到第二個真實案例。
- Stage handoff **不**另建 `handoff.yaml`：由 `artifact-gates.yaml` 同時表達 artifact、
  acceptance、next-stage eligibility、rollback owner。
- 首輪不建立 `analysis/3d-character-production/` 或 `intelligence/3d-character-production/` 空殼。

### Domain Boundary

| 任務 | 本 workflow | 其他 route |
| --- | --- | --- |
| 角色 reference、mesh、rig、表情、服裝、動作與 VRM/GLB/FBX **資產**驗收 | Primary | — |
| Viewer 開啟、humanoid mapping、expression trigger、motion play、outfit switch | Asset consumption（本 workflow） | — |
| Blender Python、exporter、viewer 或 runtime **程式**修改 | Acceptance contract 由本 workflow 定義 | 實作走 `software-delivery` |
| Inventory UI、角色選擇 API、權限、network sync、application state | 不適用 | `software-delivery` |
| 純 2D 圖片或不可驅動影片生成 | 不適用 | 未來 media/creative workflow |
| 機器學習模型訓練、fine-tuning、推論部署 | 不適用 | software/ML workflow（若建立） |
| Live2D / 第二輸出 profile | 首輪不宣稱支援 | Q6 deferred：等第二個真實案例 |

混合任務採順序 handoff：先由本 workflow 定義資產品質與輸出 contract，再由
`software-delivery` 實作自動化。本 workflow 只回答 **Can the asset be consumed correctly?**；
應用行為 **Can the application correctly implement its behavior?** 不歸本 domain。

### Alternatives Considered

- **A. 建立 `ai-modeling` skill**：reject。名稱同時可能指 3D、影像生成與機器學習，routing
  語意不穩；且把工具 prompt 當主體，無法解決下游品質 gate。
- **B. 把規則放進外部角色專案 checklist**：reject。只能改善單一角色，無法跨專案重用，
  也不能形成 workflow routing 與 validation scenarios。
- **C. 擴充 `software-delivery`**：reject。角色資產品質不是軟體交付 domain model 的附屬項；
  會把 asset acceptance 與 code acceptance 混成雙 source-of-truth。
- **D. 直接建立完整 3D intelligence/analysis/workflow 三層**：reject。現有 evidence 只有一個
  主要專案，先建立 workflow + dogfood，避免未驗證抽象。
- **E. `3d-character-production` workflow + VRM first profile（accept）**。

### Why Not an ADR Yet

目前只有一個主要 dogfood，尚未證明 asset maturity gates 可跨不同角色、不同生成來源與不同
輸出格式穩定重用。先以 plan + workflow + scenarios 驗證；至少兩個額外角色案例通過後，再判斷
是否存在值得 promotion 的 foundational architecture decision。

### ADR Promotion Criteria（completed 時評估）

- [ ] 至少 3 個角色案例完成全流程，其中至少 1 個不是 VRM。
- [ ] 至少 2 種 acquisition strategy（生成、手工、外包任兩種）共用同一 gate。
- [ ] 重做能被定位到明確 owner stage，而非依賴人工直覺。
- [ ] Open Questions 全部 resolved 或有具體 deferred owner。
- [ ] 沒有較輕的 workflow/intelligence promotion target。

### Consequences

#### 正面

- 生成品質、可修復性、可驅動性與 runtime 可用性分開驗收。
- 每次失敗留下可比較 evidence，減少「再生成一次看看」。
- VRM、GLB、FBX 或未來 profile 可共享上游角色與 mesh gates。
- 自動化腳本有明確 contract，不再用 exporter success 代替產品品質。

#### 負面

- 前期會增加規格、候選紀錄與 QA 成本。
- 部分視覺判斷仍需人工或獨立 reviewer，不能全部機械化。
- 第一次導入可能暴露既有資產無法通過中間 gate，需要回退重做。

#### 風險

- **量化幻覺**：為視覺品質硬造單一分數。緩解：使用多維 rubric + blocking defects，不建立
  無證據的綜合分數。
- **工具鎖定**：把特定生成服務或 DCC 操作寫成 canonical。緩解：核心只定 input/output/evidence；
  工具差異放 profile 或 adapter。
- **流程過重**：小型探索被全套 gate 拖慢。緩解：三級 maturity；`exploration` / `prototype`
  可省略部分 gate；**只有 `runtime-ready` completion claim** 需要完整 blocking set。
- **Runtime God Object**：把應用 UI／API／網路／效能全塞進 3D workflow。緩解：資產消費
  vs 應用行為分界見 §Domain Boundary。
- **Deformation set 膨脹**：每個角色無限 pose/camera/expression。緩解：最小充分集合（暴露
  該 stage 主要 failure class），不是全覆蓋。
- **專案證據污染**：把私有素材、路徑或角色細節寫進 reusable docs。緩解：外部 evidence 留在
  consumer project，本庫只保存抽象 scenario 與去敏 lesson。

## Proposed Workflow Shape

```text
workflow/3d-character-production/
  README.md
  execution-flow.md
  intake.md
  character-specification.md
  reference-governance.md
  candidate-generation.md
  mesh-quality.md
  rigging-and-deformation.md
  facial-expression.md
  outfit-and-animation.md
  export-and-runtime-validation.md
  artifact-gates.md
  profiles/
    README.md
    vrm.md
  execution-flow.yaml
  artifact-gates.yaml
```

首輪不建立 `tools/adapters/`；只有實際出現第二種工具路徑且差異無法由 profile 表達時才新增。

## Runtime Execution Path

| 環節 | 計畫內容 |
| --- | --- |
| Runtime owner | `knowledge/runtime/routing-registry.yaml` 的候選 `route.workflow.3d-character-production` |
| Event / signal | 使用者明確要求 3D 角色建模、character rig、retopology、blendshape、VRM、角色換裝或角色動作資產驗收 |
| Ambiguity gate | 裸「AI 建模」不得單獨鎖 route；先裁決 3D character vs ML model vs image/video generation |
| Loaded source | `workflow/3d-character-production/execution-flow.md` + intake / artifact gates；依 stage lazy-load focused files |
| Runtime action | 單一路由鎖定後由既有 `gate.workflow.primary_source_read` 強制讀 primary source |
| Evidence | workflow-context positive/counter cases + stage-gate scenarios + external dogfood report |

### Per-surface consumer 表

| Generated surface key | Named consumer(s) | Consumer 類型 |
| --- | --- | --- |
| `route.workflow.3d-character-production` | `DetectWorkflows`、`workflowPrimarySourceGate` | detector + Go validator |
| `workflow.3d_character_production.execution_flow.contract` | runtime routable lookup、projection validator | lookup + validator |
| `workflow.3d_character_production.artifact_gates.contract` | runtime routable lookup、projection validator | lookup + validator |

若 Phase 0 發現無法建立足夠精準的 discovery signal，第一版改為
`manual_activation: { reason: "domain vocabulary remains ambiguous" }`，不得為了完成 route 而加入
會吸走 ML 任務的寬鬆關鍵字。

## Pre-build Interrogation

- **Goal**：讓 3D 角色資產生產可重跑、可比較、可在早期拒絕錯誤候選。
- **Scope**：角色規格、reference、候選、mesh、rig、face、outfit/action、export/runtime gates。
- **Non-goals**：不建立 3D 生成模型；不教完整 Blender 操作；不保存私有素材；不承諾自動美術評分；
  本輪不實作 workflow。
- **Acceptance / validation target**：scenario-first、routing counter cases、artifact schema、
  deformation set、viewer/runtime readback、外部 dogfood。
- **Framework discovery**：workflow 是 canonical execution owner；project evidence 留在 consumer repo；
  runtime DB 僅為 projection。
- **Duplication risk**：需與 software-delivery、未來 ML workflow、creative media workflow 做語意裁決。
- **Assumption**：首個 profile 為 VRM；上游 gates 不依賴特定格式；full-body 為驗收主路徑。
- **Decision**：Phase 3 workflow core 已寫；route／投影仍 Phase 5。

## Open Questions

- [x] **Q1（blocker）**：domain 是否確認使用 `3d-character-production`，而非較窄的
  `vrm-character-production`？
  → **`resolved`（YES）**：通用 domain；VRM 只是 first profile。縮成 `vrm-character-production`
  會在第二個格式出現時碎 domain。
- [x] **Q2（blocker）**：第一版 completion target 是 `runtime-ready`，還是先做到
  `prototype`？
  → **`resolved`**：第一版 dogfood 目標 = **`runtime-ready`**，定義為
  `blocking gates PASS` + `target runtime contract PASS` + `known defects 已分類`。
  不是「所有東西完美」。允許 `runtime-ready with known defects`；blocking 未過 = 不得 completion。
- [x] **Q3**：Identity 用人工 rubric 還是固定 turnaround/render？
  → **`resolved`（兩者都要，非二選一）**：Specification / Reference Lock 的 required
  attributes + attribute-level rubric（判斷規則）+ fixed-view evidence（turnaround/render）。
  Rubric = decision rule；render = evidence。禁止 AI scalar score。兩閘僅
  Specification / Reference Lock 與 Identity Acceptance（不另設第三個 Identity Contract 閘）。
- [x] **Q4**：full-body vs half-body？
  → **`resolved`**：full-body = **primary required** profile；half-body = optional/derived。
  Rig / deformation / outfit / animation 的 workflow 價值主要在 full-body。
- [x] **Q5**：製作者自驗 vs fresh reviewer？
  → **`resolved`（雙層）**：self-check = **stage progression**；fresh reviewer =
  **completion claim**。不得用自驗宣稱 `runtime-ready`。
- [x] **Q6**：第二個 graduation profile？
  → **`deferred`**：等第二個真實（非 VRM）案例出現再選 GLB/FBX/Live2D。
  **Owner** = workflow maintainer at second-profile intake（屆時選定格式；Phase 0 不預選）。
  ADR 的「至少 1 個非 VRM」是 promotion 條件，**不是** Phase 4 現在就要寫第二個 profile。
- [x] **Q7**：provenance 粒度？
  → **`resolved`**：目標是 **candidate lineage 可解釋的足夠重跑**，不是保存服務全部內部
  參數。最小欄位：`acquisition_strategy` / `provider` / `model` / `model_version` /
  `input_reference_ids` / `generation_settings` / `seed` / `timestamp`。無 seed →
  `seed: unavailable`；settings 不全 → `generation_settings.availability: partial`。禁止硬補。
- [x] **Q8**（review 新增）：Stage handoff 是否獨立 artifact？
  → **`resolved`**：**不**新增 `handoff.yaml`。`artifact-gates.yaml` 同時表達 artifact +
  acceptance gate + next-stage eligibility + rollback owner。

## Phase 0 — Architecture Compatibility Preflight

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved` / `still-open` / `deferred`
- [x] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [x] 盤點新發現問題已加入 §Open Questions（Q8）

| Open Question | 處置 | 證據 / 原因 |
| --- | --- | --- |
| Q1 Domain 名稱 | resolved | stakeholder 2026-08-31：`3d-character-production` |
| Q2 第一版 maturity | resolved | `runtime-ready` = blocking + runtime contract + classified defects |
| Q3 Identity | resolved | Lock + rubric + fixed-view；僅兩閘，無第三 Contract 閘 |
| Q4 Body scope | resolved | full-body required；half-body optional/derived |
| Q5 Reviewer | resolved | self-check = stage；fresh = completion |
| Q6 Second profile | deferred | 等第二案例；owner = workflow maintainer at intake |
| Q7 Provenance | resolved | lineage-sufficient；unavailable/partial 合法 |
| Q8 Handoff contract | resolved | 併入 artifact-gates.yaml，不新增檔 |

### Phase 0.1 — Current Architecture Check

- [x] Routing registry：無 `route.workflow.3d-*` / VRM / blendshape 等價 route（grep 2026-08-31）。
- [x] 與 software-delivery 邊界：程式實作歸 SD；資產 acceptance 歸本 workflow（review 確認保留）。
- [x] Candidate files：`workflow/3d-character-production/` 為新建；compiler / `validation/scenario.schema.json` / `runtime.db` 既有。
- [x] 既有 3D/VRM intelligence：repo 內無第二份 character-production source（僅本 plan + consumer 專案）。
- [x] 第一版 routing：精準 user_signals；裸「AI 建模」不得鎖 route。若 detector 詞彙不夠穩 →
  `manual_activation`，不為 recall 吸走 ML / 影像生成。
- [x] Linked-update matrix：見 §Linked Updates。
- [x] Glossary collision：四候選詞無同名；近鄰 `change_intent_lock` 不合併。

完成條件：Q1/Q2 resolved ✓；source-of-truth、route 邊界與第一版 maturity 明確 ✓。
**Decision = Phase 0 closed**（stakeholder 2026-08-31 二次確認，無 architecture blocker）。
Phase 1 scenarios 已寫；**仍不寫** workflow 檔案。

### Phase 0.2 — Frozen semantics（裁決後不得 silently 改寫）

| 概念 | 凍結定義 |
| --- | --- |
| `exploration` | 探索生成／拓撲；可無完整 gate；不得宣稱產品可用 |
| `prototype` | 可開啟、可基本動作；允許未分類缺陷；不得 completion claim |
| `runtime-ready` | 全部 **blocking** gates PASS + target runtime contract PASS + known defects **已分類**（含 severity / owner / defer-or-fix） |
| `runtime-ready with known defects` | 合法子集：僅 non-blocking 缺陷；blocking 未過則仍是 prototype |
| Specification Lock | 規格與 Reference Set 齊全；未鎖不得比較候選 |
| Identity Acceptance | 候選對 must-preserve attributes 的 pass/fail；blocking fail = reject |
| Deformation Acceptance Set | **最小充分** pose/camera/observation，暴露 failure class；非全 pose 覆蓋 |
| Asset vs app | 資產可被正確消費 vs 應用正確實作行為 |
| Production vs completion review | 作者自驗推進 stage；fresh reviewer 才能 completion |

建議最小 deformation set（Phase 2 可微調，不可無上限膨脹）：Neutral、T-pose/A-pose、Arms up、
Knee bend、Shoulder rotation、一個 extreme facial expression。

## Phase 1 — Test-First Scenarios

完成條件：下列 YAML 存在、符合 `validation/scenario.schema.json` 形狀；在 Phase 3/5 落地前
**預期 fail by absence**（不得假裝 detector 或 gate 已綠）。

- [x] Routing positive + counters（ML、純圖/影片、裸「AI 建模」）：
      [`validation/scenarios/runtime/workflow-detector-3d-character-production-v1.yaml`](../../../validation/scenarios/runtime/workflow-detector-3d-character-production-v1.yaml)
- [x] Specification/Reference Lock：
      [`validation/scenarios/3d-character-production/specification-lock-blocks-generation-v1.yaml`](../../../validation/scenarios/3d-character-production/specification-lock-blocks-generation-v1.yaml)
- [x] Identity Acceptance（blocking fail 禁 mesh/rig；禁 scalar）：
      [`validation/scenarios/3d-character-production/identity-acceptance-blocks-mesh-v1.yaml`](../../../validation/scenarios/3d-character-production/identity-acceptance-blocks-mesh-v1.yaml)
- [x] Downstream mutation → Identity re-review／unaffected（非 Phase 0 blocker）：
      [`validation/scenarios/3d-character-production/identity-downstream-mutation-invalidation-v1.yaml`](../../../validation/scenarios/3d-character-production/identity-downstream-mutation-invalidation-v1.yaml)
- [x] Mesh gate：
      [`validation/scenarios/3d-character-production/mesh-gate-blocks-rig-v1.yaml`](../../../validation/scenarios/3d-character-production/mesh-gate-blocks-rig-v1.yaml)
- [x] Deformation gate：
      [`validation/scenarios/3d-character-production/deformation-gate-blocks-animation-v1.yaml`](../../../validation/scenarios/3d-character-production/deformation-gate-blocks-animation-v1.yaml)
- [x] Export / runtime readback：
      [`validation/scenarios/3d-character-production/export-runtime-readback-v1.yaml`](../../../validation/scenarios/3d-character-production/export-runtime-readback-v1.yaml)
- [x] Provenance：
      [`validation/scenarios/3d-character-production/provenance-blocks-promotion-v1.yaml`](../../../validation/scenarios/3d-character-production/provenance-blocks-promotion-v1.yaml)

## Phase 2 — Artifact Contracts

欄位契約索引：[`01-artifact-contracts.md`](01-artifact-contracts.md)。**不**把 Phase 1 scenario 的 heuristic
再寫成第二套規則；Phase 3 只 promote 這些 record，不另抄。

- [x] Character Specification：[`contracts/character-specification.yaml`](contracts/character-specification.yaml)
- [x] Reference Set：[`contracts/reference-set.yaml`](contracts/reference-set.yaml)
- [x] Identity Acceptance record（attribute pass/fail + blocking；禁 scalar）
- [x] Identity invalidation 表（唯一 SoT）：[`contracts/identity-acceptance.yaml`](contracts/identity-acceptance.yaml)
- [x] Candidate Record：[`contracts/candidate-record.yaml`](contracts/candidate-record.yaml)
- [x] Mesh QA Report：[`contracts/mesh-qa-report.yaml`](contracts/mesh-qa-report.yaml)
- [x] Deformation Acceptance Set：[`contracts/deformation-acceptance-set.yaml`](contracts/deformation-acceptance-set.yaml)
- [x] Runtime-ready Character Pack：[`contracts/runtime-ready-character-pack.yaml`](contracts/runtime-ready-character-pack.yaml)
- [x] `artifact-gates.yaml` eligibility + rollback owner：[`contracts/artifact-gates.yaml`](contracts/artifact-gates.yaml)
      （`runtime_projection.enabled: false`；不是 executable-contract）

## Phase 2.5 — Contract probe（非正式 dogfood）

Phase 2 合約寫完後、Phase 3 寫滿 workflow 前：用既有 VRM 角色專案填一次 artifact shape。
目的只回答「這些欄位在真實資產上填得出來嗎？」；**不是** completion dogfood。

- [x] 至少填一份 Specification + 一份 Candidate Record（含 rejected 或 unknown）。
- [x] 標出填不出的欄位 → 本輪 **無缺欄**；未觀察項填 fail/hold，禁止空欄硬過、禁止發明 PASS。
- [x] Evidence 留 consumer repo；本庫結論見 [`evidence/2026-08-31-phase-2.5-contract-probe.md`](evidence/2026-08-31-phase-2.5-contract-probe.md)。

## Phase 3 — Workflow Core

- [x] README、execution flow 與 focused stage files（`workflow/3d-character-production/`）。
- [x] Maturity：`exploration → prototype → runtime-ready`（Phase 0.2）。
- [x] Gate failure rollback owner（`records/artifact-gates.yaml`）。
- [x] 工具名不進 core；VRM profile 留 Phase 4。
- [x] `artifact-gates.md` + YAML（eligibility 只讀欄位；invalidation 不搬出 identity-acceptance.yaml）。
- [x] `runtime_projection.enabled: false`；**未**登記 route／glossary。

## Phase 4 — VRM First Profile

- [ ] 定義 Humanoid mapping、scale/axis、required expressions、spring bones、material 與 metadata。
- [ ] 定義 **資產消費** readback：viewer open、humanoid mapping、expression trigger、
  motion play、outfit switch。不含 inventory UI / 選角 API / 權限 / 同步。
- [ ] full-body = 完成定義主路徑；half-body 另列較窄 eligibility，不得與 full-body 混稱完成。
- [ ] 保留 Blender 為實作範例，不將其升為 workflow 必要工具。
- [ ] 不在本 Phase 寫第二個輸出 profile（Q6）。

## Phase 5 — Routing、Runtime 與 Linked Updates

- [ ] 新增精準 route 或明確採 manual activation。
- [ ] 新增 execution/artifact YAML projections 與 named consumers。
- [ ] 更新 `workflow/README.md`、`workflow/workflow-routing.md`、knowledge index/summary。
- [ ] 視 Glossary Impact 結果更新 glossary。
- [ ] 執行 runtime compile、refresh、validate 與 workflow-context counter cases。

## Phase 6 — External Dogfood（full；非正式 probe 見 Phase 2.5）

- [ ] 以現有 VRM 角色專案從目前 maturity 重新進場，不假裝從零開始。
- [ ] 對每個既有 artifact 標記 pass/fail/unknown 與 owner stage。
- [ ] 至少拒絕一個「可輸出但不可進下一階段」的候選，驗證 gate 有實際區分力。
- [ ] 完成固定 deformation set、expression/action 與 viewer readback。
- [ ] 專案 evidence 留在 consumer repo；只將去敏 lesson 回寫本庫。

## Phase 7 — Independent Review and Close Loop

- [ ] Fresh reviewer 依 acceptance 與 counter cases 審查，不只重跑命令。
- [ ] 對 findings 逐項 fix/defer/reject；fix 項重驗。
- [ ] 檢查文件大小、語言、去敏、工具中立與 linked updates。
- [ ] Commit/push/readback/clean status。
- [ ] 使用者確認 workflow 是否達到可用門檻；未確認則 plan 保持 active。

## Stakeholder 同意項目

| 項目 | 狀態 |
| --- | --- |
| 建立 workflow，而非單一 `ai-modeling` skill | ✅ 2026-08-31 review 保留 |
| Domain 名稱為 `3d-character-production` | ✅ Q1 |
| 通用 workflow + VRM first profile | ✅ |
| 第一版 completion = `runtime-ready`（blocking + runtime + classified defects） | ✅ Q2 |
| Identity = Lock + Acceptance；attribute rubric + fixed evidence；無 scalar score | ✅ Q3 |
| full-body required；half-body optional | ✅ Q4 |
| self-check = stage；fresh review = completion | ✅ Q5 |
| 第二 profile defer | ✅ Q6 deferred |
| Provenance = lineage-sufficient | ✅ Q7 |
| Handoff 併入 artifact-gates.yaml | ✅ Q8 |
| Dogfood 使用既有 VRM 角色專案，evidence 不進 reusable docs | ✅ |
| 首輪不建立 analysis/intelligence 空殼 | ✅ |
| Routing 精準 auto-detect；歧義則 manual activation | ✅ |
| Phase 2.5 小 probe，Phase 6 才 full dogfood | ✅ |
| Phase 0 close；進 Phase 1；不新增 Q9+ | ✅ 2026-08-31 二次 review |
| Identity PASS 非永久；Phase 1 測 downstream mutation | ✅ scenario 已寫 |
| Phase 2 contracts 不複製 scenario 規則；invalidation 單一 SoT | ✅ 2026-08-31 |

## Watch-Out List citation

依 [`architecture/ai-native-cognitive-ecosystem-system.md`](../../../architecture/ai-native-cognitive-ecosystem-system.md)
§Watch-Out List：

- 不為單一 dogfood 建完整新 layer hierarchy。
- 不把特定工具、供應商或專案 workaround promotion 成 framework contract。
- 不用 runtime projection 代替實際 consumer。
- 不用單一品質分數包裝未驗證的主觀判斷。
- 不把 workflow 做成 3D 百科；只保存執行順序、gate、artifact 與 handoff。
- 不把 runtime 消費驗收膨脹成應用行為（UI / API / 權限 / 同步）。
- 不用單一 Identity score；不用無限 pose 集合代替最小充分 deformation set。
- Rollback-owner 先當 workflow 行為 dogfood，**現在不** promotion 成 cross-domain pattern。

## 完成條件

- [x] Phase 0 裁決完成並關閉（Q1–Q8；2026-08-31）。
- [x] Phase 1 Test-First Scenarios 已寫入 validation（含 identity invalidation）。
- [x] Phase 2 Artifact Contracts 已寫入 `contracts/`（不投影、不建 workflow）。
- [x] Phase 2.5 contract probe PASS ≠ workflow PASS。
- [x] Phase 3 workflow core（無 route／無投影）。
- [ ] Phase 4–7 全部完成，或 deferred 項有 owner/entry condition。
- [ ] Routing positive/counter 在 Phase 5 登記後通過（現況 detector 應 no_match）。
- [ ] External dogfood 證明至少一個早期 gate 能阻止下游返工。
- [ ] Runtime surfaces 有 named consumer，或明確採 manual activation。
- [ ] Linked updates、runtime projection、readback 與 repository close-loop 完成。
- [ ] ADR promotion criteria 已評估，不因單一案例提前 promotion。

## Linked Updates

| Surface | 原因 |
| --- | --- |
| `workflow/README.md` | 新 workflow 人類索引 |
| `workflow/workflow-routing.md` | 3D character / ML model / media / software automation 歧義裁決 |
| `knowledge/runtime/routing-registry.yaml` | route canonical 或 manual activation |
| `knowledge/indexes/README.md` | task intent 導航 |
| `knowledge/summaries/` | summary-first route |
| `knowledge/glossary/ai-skill.md` | 新 framework vocabulary（若確認） |
| `validation/scenarios/` | routing 與 stage-gate test-first scenarios |
| `plans/.../contracts/` | Phase 2 record shapes；Phase 3 再 promote |
| `runtime/runtime.db` | compiler projection；不得手改 |

## 與其他 plans 的關係

- [`2026-07-30-2101-legal-workflow-domain.md`](../2026-07-30-2101-legal-workflow-domain.md)：
  參考其 domain boundary、scenario-first、runtime wiring 與「不建空殼層」策略；不複製法律專屬
  intake/risk model。
- Software Delivery Framework Domain Model 相關 plan：本 workflow 擁有角色 asset
  acceptance；software-delivery 只擁有其自動化程式的交付。
- Workflow Activation Engine 相關 archived plans：沿用既有 detector 與
  `gate.workflow.primary_source_read`，本 plan 不另建 activation engine。
- [`2026-06-16-1131-evidence-candidate-system.md`](../2026-06-16-1131-evidence-candidate-system.md)：
  Candidate Record 的 accept/reject/reason 形狀相近；本 domain **不**把 evidence-candidate
  系統當 runtime 依賴，只借用形狀，避免過早抽共用 primitive。
