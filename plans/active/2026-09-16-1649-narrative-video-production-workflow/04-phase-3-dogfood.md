# Phase 3 dogfood protocol

Companion to [`_plan.md`](_plan.md)。**不新增架構、不補 Q12/Q13、不接 runtime。**  
Phase 2 落點：`34f778d4`（workflow）、`8757a578`（plan 狀態）。

目標不是「成功做出一部片」，而是：一部真實片子能否把 **素材選擇、敘事結構、EDR 決策、locale、QC、publish 狀態** 留成可驗證決策鏈。

## 核心原則

拿一部真實片子，證明 Phase 2 contract 能承載真實決策。卡住時先分類，禁止直接加欄位或加 phase。

| 分類 | 含義 | 本 round 允許 |
| --- | --- | --- |
| `contract_gap` | Phase 2 欄位／閘無法表達真實決策 | 記缺口；不改 workflow 除非使用者授權 |
| `dialogue_semantic_ambiguity` | 短台詞省略主詞 | optional `dialogue.semantic_context` 已落地；升必填等計數 |
| `shot_unit_semantic_gap` | 分鏡單元（dialogue／action／visual）缺機器可讀語義 | **不是**立刻 `contract_gap`。候選：每個單元 = `text` + `semantic_context`。本 phase **不**擴 workflow。見 [`evidence/2026-09-17-shot-unit-semantic-context.md`](evidence/2026-09-17-shot-unit-semantic-context.md) |
| `character_naming_gap` | 角色多名稱／笼统詞／寫稿詞彙不受控 | **不是**立刻 `contract_gap`。候選 **Series Cast Canonicalization**。Phase 3 **不改** workflow。見 [`evidence/2026-09-17-series-cast-canonicalization.md`](evidence/2026-09-17-series-cast-canonicalization.md) |
| `material_fact_gap` | 素材事實層未先於語義推理 | **不是**立刻 `contract_gap`。L0 機械證據 → L1 身份／context → L2 決策。分析器當外部 evidence candidates。見 [`evidence/2026-09-17-material-fact-extraction.md`](evidence/2026-09-17-material-fact-extraction.md) |
| `data_insufficient` | bible／catalog／locale 還沒填夠 | 補資料，不改契約 |
| `adapter_only` | ffmpeg／TTS／模型／GUI 問題 | 留在外部工具；canonical 不吸收 |
| `design_error` | invariant 本身擋不住或互相矛盾 | 停手，回 Phase 1 護欄討論 |

## 觀察鏈（依序，不可跳過分類）

```text
real brief → source bible → clip catalog → template
  → matching script → feasible candidates → explicit selection policy
  → EDR → locale packs → QC / independent verification
  → publish-ready → outcome
```

對照檔：[`workflow/narrative-video-production/execution-flow.md`](../../workflow/narrative-video-production/execution-flow.md)。

特別確認：每個 shot 有可行集 + `selection.policy`；查找只回既有 `clip_id`；成片對 EDR 而非反推；三閘分開；`publish-ready` 有獨立 verifier。Q4／Q6／Q10 只觀察，不在本 phase 凍結。另計：semantic unit、series_cast、以及 L0 機械證據哪些真的被 catalog／matching 消費。**本 phase 不因這些觀察擴 workflow**（dialogue optional 維持現況；分析器不進 contract）。

## 本庫 vs 外部專案

- **執行與原始媒體**在 `<PROJECT_ROOT>`（非本庫）。
- **回寫本庫**只允許去敏摘要：進 [`evidence/`](evidence/README.md)。禁止片名／路徑／host／金鑰／未授權肖像。
- 虛構示範 [`sanitized-matching-and-edr.yaml`](../../workflow/narrative-video-production/records/examples/sanitized-matching-and-edr.yaml) **不算** Phase 3。

## Phase 3 PASS 最低條件

1. 外部有一部真實片子走過觀察鏈（outcome 可停在 `insufficient_sample`）。
2. 本庫 `evidence/` 有一份去敏 run：鏈上每站 `pass`／卡住分類。
3. 未註冊 route、未開 runtime projection、未把 provider 寫進 workflow。
