---
id: 2026-10-14-0832-image-processing-decision-workflow
plan_kind: main
status: in-progress
owner: workflow
owner_layer: workflow
created: 2026-10-14
parent: null
---

# Image Processing Decision Workflow（`workflow/image-processing/`）

**Status**: in-progress — **Phase 0 freeze**（2026-10-14；見 [`08`](08-phase-0-freeze-invariants.md)）。下一階段：Phase 1 doc-only 骨架。

**Glossary Impact**: yes — 候選詞 `image_task_contract`、`image_decision_record`（IDR）、`feasible_candidates`、`selection_policy`、`independent_image_verification`、`failure_class`、`repair_candidate`。Phase 1 前不登記 glossary。

## 一句話目標

建立 **governed image processing decision** 總綱：Input Contract → Image Task Analysis → Constraints → Feasible Candidates → Selection Policy → Execute → Independent Verify → Repair → Evidence Closeout。模型負責理解與選擇；確定性工具負責可重複操作；獨立驗收判斷「是否符合任務契約」，不以「工具成功輸出檔案」當完成。

## Decision Rationale

### Problem & Why Now

近期 El Dorado Dice 符號去背／置中（Cursor agent + 本機 Python：`rembg`／Pillow／scipy）暴露的不是「模型不夠聰明」，而是缺決策閉環：

| 失效層 | 實例（見 [`07`](07-dogfood-baseline-el-dorado-cutout.md)） |
| --- | --- |
| 任務判斷 | 把「透明去背 + 對齊灰階對照」做成過度侵蝕／填色／改主體細節 |
| 工具選擇 | 全案同一條 rembg → 硬 alpha → dilate/erode，無依特徵路由 |
| 驗收缺失 | PNG 寫出即當成功；棋盤格預覽被當成透明通道證據 |
| 修復策略 | 同一閾值／opening 重跑；牙齒被 opening 咬穿後才後退 |

這與 Ai-skill 既有 Loop-first／Governance-first／Constraint≠Selection／Independent verification 同一類問題。**不**建議只補「去背 Skill」或加幾個圖片 Prompt；也**不**併入 narrative-video 當子步驟。應做成 **cross-cutting capability workflow**，供影片、商品圖、角色素材、Logo、slot 符號共用。

### Decision

新增 **`workflow/image-processing/`**（獨立 main plan；NVP／slots／3D 為 future consumers）。

凍結意圖（**Phase 0 freeze**；權威細節 [`08`](08-phase-0-freeze-invariants.md)）：

1. **完成標準 = 符合 Image Task Contract**，不是「處理流程跑完」。
2. **Constraint ≠ Selection**：契約決定可行集；Selection Policy 決定候選；LLM 不得自定驗收標準。
3. **Execute ≠ Verify**：執行工具與驗收工具分離；禁止同一模型自證 PASS。
4. **`inconclusive` 合法**：證據不足時不得硬判 PASS／FAIL。
5. **Failure taxonomy → 有限 repair**；不可無限重跑同一策略。
6. **工具選型不在 Phase 0 定死**：rembg 可作基準候選，非唯一真理。
7. **不先建 Auto Orchestrator／全能圖片 Agent**。

Companions：[`01`](01-problem-baseline-and-scope.md) · [`02`](02-image-task-contract.md) · [`03`](03-decision-model-and-selection.md) · [`04`](04-bg-removal-vertical-slice.md) · [`05`](05-independent-verification.md) · [`06`](06-failure-taxonomy-and-repair.md) · [`07`](07-dogfood-baseline-el-dorado-cutout.md) · [`08`](08-phase-0-freeze-invariants.md) · [`fixtures/`](fixtures/README.md)。

### Domain Boundary

| 任務 | 本 workflow | 其他 |
| --- | --- | --- |
| 任務分類、契約、候選策略、選擇、執行閘、獨立驗收、失敗修復、IDR／證據 | Primary | — |
| 影片 timing／字幕／分鏡 | Consumer | `narrative-video-production` |
| 3D／角色生產管線 | Consumer | `3d-character-production` |
| Slot／paytable 符號資產規範 | Consumer profile | 業務 repo（如 `pojiepokerist/docs/slots`） |
| 具體模型權重、ComfyUI graph、雲端 API | Adapter only | `ai-tools/`／專案腳本 |
| 生成式「整圖重畫」當去背替代 | **Out of scope for P0** | 僅 C 類且契約允許時 |

### Alternatives Considered

| 方案 | 裁決 |
| --- | --- |
| A. 只加 rembg Skill／Prompt | **reject** — 無契約／驗收／修復 |
| B. 併入 NVP execution-flow | **reject** — 圖片能力跨領域 |
| C. 一開始建大型圖片 Agent／編排器 | **reject** — 違反 Loop-first；過早綁工具 |
| D. Image Processing Decision Workflow + 去背垂直切片先行 | **accept**（Phase 0 freeze） |

### Why Not an ADR Yet

尚無 `workflow/image-processing/` 目錄、無第二 consumer、無標註測試集。Plan + contracts 穩定且 ≥1 dogfood 閉環後再評估 ADR。

### ADR Promotion Criteria（completed 時評估）

- [ ] Image Task Contract + Verification Contract 經 ≥1 schema revision 仍可遷移或向後相容
- [ ] 去背垂直切片有獨立 verify（非執行器自報）且 failure≠success 可重現
- [ ] ≥1 真實 dogfood（建議沿用 slot 符號／商品圖）含 IDR + review
- [ ] 未把單一供應商／模型寫進 canonical 步驟
- [ ] Open Questions 全解或有 deferred owner

## Assumptions（已接受 — 2026-10-14）

| # | 問題 | 裁決 |
| --- | --- | --- |
| A1 | 主要執行環境 | **接受 — 混合**：agent 決策 + **本機 Python** 執行；Ai-skill 持契約與 workflow。非預設 ComfyUI／雲端 API |
| A2 | 第一階段範圍 | **接受** — A 類去背閉環；B 壓架構；C／D 預留 |
| A3 | 與影片工作流關係 | **接受** — 獨立 capability；NVP 等為 consumers |

## Open Questions（責任矩陣 — freeze 時點）

| # | 問題 | 狀態 | Freeze？ | 預裁決／延後 |
| --- | --- | --- | --- | --- |
| Q1 | 正式執行入口 | **resolved via A1** | 已解除 | 混合；CLI／skill 名稱 Phase 1 再定 |
| Q2 | silhouette／灰階進通用核心？ | **pre-decided** | — | Consumer constraint only |
| Q3 | 驗收視覺層 actor | **pre-decided（P0）** | — | 機械 + 人審 |
| Q4 | 測試集存放 | **deferred + owner** | 已滿足 freeze 條件 | Staging＝plan `fixtures/`；canonical＝`workflow/image-processing/fixtures/`（Phase 1）；私有＝consumer。Owner：Phase 1 skeleton |
| Q5 | 授權核對 | **deferred + owner** | 已滿足 freeze 條件 | Stub 在 [`08`](08-phase-0-freeze-invariants.md)；**完整清單 Phase 2 Gate 前必達**。Owner：Phase 2 prep |

權威 freeze 敘述：[`08`](08-phase-0-freeze-invariants.md)。

## Phases

### Phase 0 — 盤點與問題基線 — **FREEZE**

- [x] Stakeholder 確認 A1–A3；Q1 隨 A1 resolved；Q4／Q5 有 deferred owner（[`08`](08-phase-0-freeze-invariants.md)）
- [x] 盤點現有圖片工具、臨時腳本、Skill、失敗案例（[`01`](01-problem-baseline-and-scope.md)、[`07`](07-dogfood-baseline-el-dorado-cutout.md)）
- [x] 建立最小可比較測試集目錄規劃（[`fixtures/README.md`](fixtures/README.md)）
- [x] ≥3 個 `reproducible_case` 規格凍結（`PENDING_BINARY`；見 `fixtures/cases/`）
- [x] Freeze invariants → `status: in-progress`
- **Gate**：通過 — 見 [`01`](01-problem-baseline-and-scope.md)／[`08`](08-phase-0-freeze-invariants.md)

### Phase 1 — Image Task Contract + Decision Model（doc-only）

- [ ] 建立 `workflow/image-processing/` 骨架（README、execution-flow、contracts、registry）
- [ ] 落地任務類型 A–D、輸入／輸出不變條件、候選策略註冊、選擇責任、驗收狀態、失敗分類
- [ ] 靜態走讀：不依賴特定模型也能判定契約是否完整
- **Gate**：契約可執行；無模型名寫死在 canonical 步驟

### Phase 2 — 去背最小垂直切片

- [ ] Input validation → analysis → candidate → mask/matte → assemble → verify → repair（有限次）
- [ ] rembg（或當時基準候選）+ 確定性後處理作 **candidate**，非唯一路徑
- [ ] Independent verification 產 `check_id`／`status`／`failure_class`／`repair_candidate`
- [ ] [`05`](05-independent-verification.md) 負向 fixture 矩陣至少跑通 `neg.alpha-opaque`、`neg.residue-halo`
- [ ] Q5 授權檢查清單已存在
- **Gate**：代表性案例可追溯驗收；負向 fixture 擋住假成功；失敗不被誤報成功

### Phase 3 — Dogfood + failure learning

- [ ] 真實圖：一般物件、人像髮絲、插畫／UI 合成、多主體、高風險邊緣
- [ ] 比較候選方法；人工覆核後才可升級 guard
- **Gate**：品質／修復成功率／成本／耗時有證據（`evidence/`）

### Phase 4 — 擴充 B／C／D（按需）

- [ ] 每個新任務類型：獨立契約片段 + fixtures + 驗收標準
- **Gate**：共用 decision loop 不需重寫；只加 registry／adapter

## Non-goals（Phase 0–2）

- 不建全能 Auto Orchestrator
- 不讓 LLM 自行決定驗收標準或永久規則
- 不為每種任務先做龐大 Skill
- 不把一次工具成功當成流程成熟
- 不把生成式整圖重畫當預設去背
- 不把 slot「對齊 Sprite 臉軸」寫成通用去背必經步驟（屬 consumer profile）

## Consequences

### 正面

- 圖片爭議可沿 Contract → Analysis → Candidates → Selection → Verify → Repair 追溯
- 換模型／工具只換 adapter；workflow 保留
- 與 translation／ERA 同一套 Constraint／Selection／Independent verification 語言

### 負面

- 每張圖多一份分析／驗收成本
- 初期需人審補 `inconclusive`

### 風險

- **Workflow inflation**：把 Pillow 逐步碼寫進 execution-flow → 緩解：canonical 到閘與紀錄；步驟留 adapter
- **過早綁 rembg**：緩解：Phase 0–1 禁寫死；Phase 2 標 `baseline_candidate`
- **驗收假陽性**：遮罩面積當品質證明 → 緩解：無真值時最多 `inconclusive`／異常偵測

## Closeout rules（草案）

宣告本 plan 任務完成前必須：

1. 有 IDR（或等效決策紀錄）含 selected candidate、參數摘要、verify 結果
2. 所有 `fail` 有 `failure_class` 與處置（repair／human_review／abort）
3. `git status` clean 且 Ai-skill writeback（若改了 workflow）已 commit／push／readback
4. Dogfood evidence 已鏈到 `evidence/`（若本輪有實測）
