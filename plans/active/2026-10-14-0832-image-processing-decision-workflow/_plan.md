---
id: 2026-10-14-0832-image-processing-decision-workflow
plan_kind: main
status: draft
owner: workflow
owner_layer: workflow
created: 2026-10-14
parent: null
---

# Image Processing Decision Workflow（`workflow/image-processing/`）

**Status**: draft — Phase 0 盤點草案；待 stakeholder 確認執行環境與範圍後 freeze。

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

凍結意圖（草案；Phase 0 freeze 後釘死）：

1. **完成標準 = 符合 Image Task Contract**，不是「處理流程跑完」。
2. **Constraint ≠ Selection**：契約決定可行集；Selection Policy 決定候選；LLM 不得自定驗收標準。
3. **Execute ≠ Verify**：執行工具與驗收工具分離；禁止同一模型自證 PASS。
4. **`inconclusive` 合法**：證據不足時不得硬判 PASS／FAIL。
5. **Failure taxonomy → 有限 repair**；不可無限重跑同一策略。
6. **工具選型不在 Phase 0 定死**：rembg 可作基準候選，非唯一真理。
7. **不先建 Auto Orchestrator／全能圖片 Agent**。

Companions：[`01`](01-problem-baseline-and-scope.md) · [`02`](02-image-task-contract.md) · [`03`](03-decision-model-and-selection.md) · [`04`](04-bg-removal-vertical-slice.md) · [`05`](05-independent-verification.md) · [`06`](06-failure-taxonomy-and-repair.md) · [`07`](07-dogfood-baseline-el-dorado-cutout.md)。

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
| D. Image Processing Decision Workflow + 去背垂直切片先行 | **accept（草案）** |

### Why Not an ADR Yet

尚無 `workflow/image-processing/` 目錄、無第二 consumer、無標註測試集。Plan + contracts 穩定且 ≥1 dogfood 閉環後再評估 ADR。

### ADR Promotion Criteria（completed 時評估）

- [ ] Image Task Contract + Verification Contract 經 ≥1 schema revision 仍可遷移或向後相容
- [ ] 去背垂直切片有獨立 verify（非執行器自報）且 failure≠success 可重現
- [ ] ≥1 真實 dogfood（建議沿用 slot 符號／商品圖）含 IDR + review
- [ ] 未把單一供應商／模型寫進 canonical 步驟
- [ ] Open Questions 全解或有 deferred owner

## Draft assumptions（待確認）

| # | 問題 | 草案假設（可推翻） |
| --- | --- | --- |
| A1 | 主要執行環境 | **混合**：Cursor／agent 決策 + **本機 Python**（Pillow／NumPy／scipy／rembg 等）執行；Ai-skill 持有契約與 workflow。非預設 ComfyUI／雲端 API |
| A2 | 第一階段範圍 | **A 類去背／透明 PNG 完整閉環**；B 類用困難案例測架構；C／D 契約預留、不在 P0 大擴 |
| A3 | 與影片工作流關係 | **獨立 capability**；NVP 日後 adapter 消費，不擁有決策契約 |

請 stakeholder 明確接受或改寫 A1–A3 後再 Phase 0 freeze。

## Open Questions（責任矩陣）

| # | 問題 | 狀態 | Freeze？ | 預裁決／延後 |
| --- | --- | --- | --- | --- |
| Q1 | 正式執行入口：Ai-skill skill／本機 CLI／agent adapter／混合？ | open（草案 A1） | **blocker**（隨 A1） | Stakeholder 確認 A1；未確認不得 freeze |
| Q2 | P0 是否把「對齊外部 silhouette／灰階」塞進通用去背核心？ | **pre-decided** | 非 blocker | **Consumer constraint only**；不進通用 A 核心 |
| Q3 | 驗收視覺層用何 actor？ | **pre-decided（P0）** | 非 blocker | P0 = **機械層 + 人審 gate**；第二模型／合成參考可 Phase 3+ |
| Q4 | 測試集存放：`workflow/.../fixtures/` vs 專案 private evidence？ | open | 可 deferred | Owner：Phase 0 盤點；預設可重用去敏→workflow，專案證據→consumer |
| Q5 | 商用授權核對掛哪個 phase？ | open | 可 deferred 至 Phase 0 清單 | **Phase 2 Gate 前**必須有授權檢查清單；缺清單不得宣告 Phase 2 完成 |

實作時禁止把 open 項偷偷當預設假設；僅上表 **pre-decided** 可當 freeze 不變條件。

## Phase 0 Freeze Invariants（草案；確認 A1–A3 後釘死）

1. 完成標準 = 符合 Image Task Contract，不是流程跑完。
2. Constraint ≠ Selection；Execute ≠ Verify；`inconclusive` 合法。
3. `fidelity` 以三維展開（[`02`](02-image-task-contract.md)）；禁止 `immutable=strict` 這種混用。
4. Observed ≠ reproducible；Phase 0 Gate 要求可比較 reproducible 案例（[`01`](01-problem-baseline-and-scope.md)、[`07`](07-dogfood-baseline-el-dorado-cutout.md)）。
5. Q2／Q3 預裁決如上；Q1 隨 A1；Q4／Q5 必須有 deferred owner 字串才可 freeze。
6. 不建 Auto Orchestrator；不把 rembg 寫死成唯一真理；slot 對齊不進通用核心。

## Phases

### Phase 0 — 盤點與問題基線

- [ ] Stakeholder 確認 A1–A3；處理 Q 責任矩陣中的 freeze_blocker／deferred owner
- [ ] 盤點現有圖片工具、臨時腳本、Skill、失敗案例（含 [`07`](07-dogfood-baseline-el-dorado-cutout.md)）
- [ ] 建立最小可比較測試集目錄規劃（尚可不放大量二進位）
- [ ] 將 ≥3 個測試集列從 candidate 升格為 `reproducible_case`（或註明阻礙）
- [ ] Freeze invariants → `status: in-progress`
- **Gate**（與 [`01`](01-problem-baseline-and-scope.md) 同構）：A1–A3 確認；≥5 observed 可歸類；≥3 reproducible 列齊；測試集路徑規劃；freeze invariants 已寫入

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
