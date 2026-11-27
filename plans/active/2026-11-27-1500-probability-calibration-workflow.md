---
id: 2026-11-27-1500-probability-calibration-workflow
plan_kind: main
status: draft
owner: workflow
owner_layer: workflow
created: 2026-11-27
parent: null
---

# Probability Calibration Workflow（`workflow/probability-calibration/` + intelligence）

**Status**: draft — 只建立 plan，未開始實作。2026-11-27 使用者審閱已裁決 Q1–Q6（見 §Open Questions）並要求 Phase 0 先做契約補強（§Phase 0.1）；Phase 0 可開始，實作開工仍以 Phase 0 Gate 為準。

**Glossary Impact**: yes — 候選詞 `calibration_contract`、`calibration_record`、`gap_ledger`、`structural_ceiling`、`holdout_probe`、`precision_target`、`holdout_policy`、`calibration_status`。Phase 1 前不登記 glossary。

## Scope（2026-11-27 審閱補充）

本 workflow 處理的是 **Probability Model Calibration & Validation**：讓機率模型／隨機產品的行為符合目標分布或規格，並以 Monte Carlo 方式驗證其統計行為。它**涵蓋**機率估計（從樣本估事件率）、模型校準（調參使模型符合目標）、Monte Carlo validation；傳統的「預測機率 calibration」（預測機率 vs 實際頻率一致性）只是其中一種特例。名稱暫不改（Q1），但 Phase 1 README 必須寫明此範圍，避免其他 consumer 因名稱判斷錯誤。

## 一句話目標

把一個消費專案（隨機性遊戲產品）累積的**機率校準方法**抽成工具中立、可重用的 workflow：
Target 定義 → 官方／參照取證 → 指標對照 → 診斷（定義／採集／產品）→ 調參 → 獨立驗證 → Gap 登錄，
並用 **Calibration Record → Failure taxonomy → Intelligence atom** 的回饋環讓它每次使用後更強。
未來遇到機率／RTP／掉落率／抽樣機問題，至少有一套現行法則可查，而不是每次從零摸索。

## Decision Rationale

### Problem & Why Now

現況盤點（2026-11-27，本 repo 內搜尋 probability／RTP／calibration／sampler／monte-carlo）：

| 位置 | 內容 | 能否直接當 workflow |
| --- | --- | --- |
| `workflow/investment/`、`validation/scenarios/investment/investment-probability-framing-v1.yaml` | 投資語境的機率**表述**紀律 | 否 — 是「怎麼說機率」，不是「怎麼校準機率」 |
| `workflow/cross-cutting/decision-support/` | 目標由人決策的兩趟式建議 | 部分 — 可承接「RTP 目標由誰定」 |
| `feedback/history/apk-analysis/**/slot-*.md` | 逆向官方 slot 行為（協議、盤面、結算） | 否 — 是**取證**側，不含校準 |
| `plans/active/**/evidence/**` 中的 probability／calibration 字樣 | 他領域證據中的偶發用語 | 否 |
| **機率校準本身** | **repo 內沒有任何 domain、workflow、intelligence atom 或 failure pattern** | — |

方法目前只存在於消費專案：一份 ~900 行、方法與十餘櫃個案流水帳混寫的校準文件，加上產品側的
取樣機設計文件與測試釘值。後果：

1. 方法（§1–5 級）與一次性結論混在一起，agent 讀不完也抽不出規則；
2. 已踩過的坑沒有被結構化（見下方 Inventory），下一個遊戲會重踩；
3. 「校準做完了」的宣稱缺統一驗收口徑（釘值＝回歸鎖，不等於校準；鏈式驗證有一關失敗卻被總結成通過）。

Why now：剛完成一次完整迭代（取樣機 Strategy 化 + 跨軸耦合 + 多目標聯合校準），失敗與修正歷程仍新鮮，
且使用者明確希望此套框架可跨遊戲共用、只改參數。

### Inventory：本案可抽象的內容（Phase 0 的種子，非最終分類）

| # | 類別 | 內容（已去專案化） | 候選落點 |
| --- | --- | --- | --- |
| I1 | Target 定義 | 受信任穩態樣本是**參照錨點**，但**不等於**產品必須追上的精確值；熱帳號／農場樣本的絕對值**不可**當目標，只取尺度不變結構（條件機率、跨軸相關、形狀占比）；目標由人決策並記錄來源。審閱修正：Target 拆三種角色（Normative／Reference Observation／Implementation，見 §Phase 0.1 C1），避免把參照樣本的隨機波動當成精確目標 | contracts + decision-support |
| I2 | 指標對照 | 官方側與產品側指標必須一一同定義（命中率、均值／注額、基礎 vs 系列拆分、可見格 vs 整盤、基準回合類型）；基準選錯會整櫃重做 | contracts |
| I3 | 診斷順序 | 旗標出現時依序排除：定義不一致 → 採集偏差 → 產品真的偏；只有第三項成立才改參數（**硬性契約**，非建議）；**不得只因 live 頻率不同改賠付表** | execution-flow + contract |
| I4 | 樣本充足 | 最小手數門檻、罕見玩法需 ≥5000 基礎手才能判「是否為零」；樣本不足只記錄不下結論。審閱修正：固定手數只是 provisional default，充足與否改由各指標的 `precision_target` 判定（§Phase 0.1 C3） | contracts |
| I5 | 驗證分層 | 釘值（固定 seed、精確回歸鎖）≠ 校準；長跑結論寫文件不當釘值（20k 手標準誤約 0.03）；優化器目標需 **holdout seed** 獨立驗證，防 overfit 自己的目標。審閱修正：holdout 不只換 seed，需分 training／validation／final_holdout 三組與揭露政策；驗證獨立性分四維記錄（§Phase 0.1 C4、C5） | verification contract |
| I6 | 結構天花板 | 每軸獨立抽樣**無法**產生跨軸相關：先證明天花板，再換抽樣結構（耦合策略），而非無限調權重。審閱拆分：「結構天花板」作為**診斷結果類別**進 execution-flow／contract（停止參數搜尋、提出結構變更需求）；「每軸獨立→無跨軸相關」這條具體機制仍是 N=1 observation atom | contract（診斷類別）+ intelligence atom（機制） |
| I7 | 單旋鈕副作用 | 加耦合會同步抬高 RTP（同類連線變多）→ 需正交旋鈕把「類別密度聚集」與「同類聚集」分開；調一旋鈕必量全指標 | intelligence atom |
| I8 | 取樣機設計模式 | 變異點做成資料（參數）+ 註冊式 Strategy；重構時**零 RNG 漂移**以其他消費者的釘值為護欄；缺參數時 fail-fast 載入 | intelligence atom（設計模式，不附程式碼） |
| I9 | 優化操作衛生 | 解析掃描優先於黑盒優化；串流輸出（避免緩衝看不到進度）；快速 JIT／plan cache 讓探針變快；seed sweep；偵測退化局部最小（如目標指標≈0） | intelligence atom |
| I10 | 誠實回報 | 鏈式驗證任一機械鎖失敗＝整體失敗，不得口頭沖掉；已知缺口必須寫入 Gap Ledger（例：絕對占比未達、單一轉輪組、合成條帶） | failure pattern + contract |
| I11 | 過期宣稱 | 舊文件的錯誤頻率（量級錯 6 倍）會被後續工作沿用 → 引用前需回查來源與日期 | failure pattern |

### Decision

新增 **`workflow/probability-calibration/`**（domain 名稱中立；slot RTP 為第一個 consumer profile），
並在 `intelligence/engineering/probability-calibration/` 放可演進的技術 atoms。**第一輪 doc-only pilot，不註冊 route、不接 runtime。**

分工（凍結意圖，Phase 0 驗證）：

1. **Ai-skill 持方法與契約；consumer 持個案流水帳與程式碼。** 個案結果、校準快照、探針程式碼留在消費專案；canonical 文件只含去專案化的規則。
2. **Calibration Contract**（target／指標定義／門檻／基準／樣本門檻／目標決策人）先於調參；沒有契約不調參。
3. **Execute ≠ Verify**：調參者（含優化器）不得自證通過；驗收用 holdout seed 與獨立 probe（沿用 `image-processing` 的契約精神，不複製其內容）。責任分三角色：**執行者**產生候選參數、**驗證者**判斷是否符合契約、**決策者**接受／拒絕最終結果；不要求三個 agent，但同一角色不得兼任執行與驗證，且 Record 必須記錄誰擔任哪個角色。
4. **Calibration Record** 每輪留一份：參數版本、量測、對目標差距、**Gap Ledger**；沿用 `traceable-evidence`（TE1 保留原始觀察、TE4 定案附理由）。
5. **回饋環**：每次校準收工必須做 FeedbackDecision；同類失敗出現 ≥3 次（跨個案）才升級成 intelligence atom／failure pattern，避免一次性事件膨脹規則庫（沿用 evidence-candidate 的 N≥3 紀律，不新建機制）。
6. **不建**自動優化器、不在 Ai-skill 放機率引擎程式碼；工具（探針、優化器）屬 consumer 或 `ai-tools/` adapter。

### Alternatives Considered

| 方案 | 裁決 |
| --- | --- |
| A. 把消費專案的校準文件整份搬進 Ai-skill | **reject** — 專案證據與去敏邊界衝突；900 行混主題違反 document-sizing；方法抽不出來 |
| B. 只寫 intelligence atoms | **reject** — 缺 Contract／驗收／Gap 登錄閉環；atoms 無法保證流程被依序執行 |
| C. 只寫 slot 專用 workflow（`workflow/slot-calibration/`） | **reject（暫）** — 方法本體（target、診斷順序、holdout、Gap Ledger）與遊戲類型無關；但保留 slot RTP 為第一 profile，是否再中立化見 Q1 |
| D. 中立 workflow + intelligence atoms + 回饋環，doc-only pilot | **accept（draft）** |
| E. 一開始就建 route／discovery signal／優化器 | **reject** — 違反 Loop-first 與 cross-cutting 的「3 個 converged case 才升格」紀律 |

### Why Not an ADR Yet

尚無第二個 consumer、無 Phase 3 dogfood 證據，容器名稱與邊界（Q1）未定。Plan + doc-only 契約穩定且 ≥2 次獨立校準使用後再評估。

### ADR Promotion Criteria（completed 時評估）

- [ ] Calibration Contract + Record 經 ≥1 次 schema revision 仍可遷移
- [ ] ≥2 次獨立校準（不同遊戲或不同玩法）以本 workflow 收工，且驗收由獨立 holdout 完成
- [ ] ≥1 條 failure pattern 因第 3 次重複而升級，證明回饋環可運作
- [ ] 未把單一語言／工具／遊戲寫進 canonical 步驟
- [ ] Open Questions 全解或有 deferred owner

### Consequences（預期）

**正面**：機率問題有現行法則可查；校準宣稱有統一驗收口徑；同類錯誤不重犯；個案證據與方法分離。
**負面**：多一個 workflow domain 要維護；去專案化抽象可能丟掉細節。
**風險**：過早抽象（只有一個 consumer）→ 以 doc-only pilot + Q2 第二 consumer 規劃緩解；canonical 文件滲入專案識別 → 以 sanitization 掃描與 Phase 1 人審緩解。

## Runtime Execution Path

**本計畫第一輪為 doc-only pilot，不接入 runtime。**

| 項目 | 內容 |
| --- | --- |
| Runtime owner | 無（不新增 gate／validator／route） |
| Trigger flow | event：agent 遇機率／RTP／抽樣校準任務 → 發現途徑：`workflow/README.md` 索引 + `workflow/workflow-routing.md` 選路表列（**不**進 `routing-registry.yaml`）→ 讀 workflow README／execution-flow → 手動套用 contract → 產出 Calibration Record。屬 `manual_activation` |
| 未來接入 | 達 3 個 converged case（跨遊戲／玩法）後，評估 `route.workflow.probability-calibration` + 一個 discovery signal；入場條件＝Phase 3 dogfood 完成 + Q2 有答案 |
| Graduation deadline | 2027-03-31：屆時升級（註冊 route + signal）或降級（歸檔為 intelligence-only） |
| Generated surface | 無（不 project 進 `runtime.db`） |
| Validation scenarios | Phase 1 以 scenario-first 方式補 ≥3 個 `validation/scenarios/probability-calibration/*.yaml`（沿用 `investment`／`image-processing` 格式），內容：holdout 缺失被判未驗證、target 取自熱樣本被擋、鏈式驗證失敗不得宣稱通過 |
| Test passing evidence | Phase 1／3 收尾填入（`ai-skill runtime refresh` + scenario validate 輸出） |

### Per-surface consumer 表

| Generated surface key | Named consumer(s) | Consumer 類型 |
| --- | --- | --- |
| （本計畫不新增 `route.*`、`runtime_projection.target_key` 或 commit-msg validator） | n/a | `manual_activation: doc-only pilot；理由見上；接入 phase 與 deadline 已列` |

## Phases

### Phase 0 — 盤點與契約凍結（Pre-Build Interrogation + Preflight）

#### Phase 0.0 — Open Questions 核對（公版，必填）

- [ ] 已讀本 plan §Open Questions 全部條目
- [ ] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [ ] `resolved` 的條目已同步勾選／附註於 §Open Questions
- [ ] 若盤點新發現問題，已加入 §Open Questions

| Open Question | 處置 | 證據／原因 |
| --- | --- | --- |
| Q1–Q6 | user-ruled（待 Phase 0 核對） | 2026-11-27 使用者審閱已給裁決（見 §Open Questions）；Phase 0 需確認裁決與 preflight 不衝突後改 `resolved` |
| Q7 | user-ruled（待 Phase 0 核對） | 同上；`10%` 定義拆 absolute／relative |

#### Phase 0.1 — 契約補強（2026-11-27 使用者審閱；吸收進既有 4 份 contract，不新增文件）

> 原則：以下全部落在 Phase 1 既有的 `calibration-contract`／`calibration-record`／`verification`／`gap-ledger` 內，不新建 workflow、phase 或獨立契約檔（沿用 minimize-governance-surface）。Phase 0.1 只凍結**欄位與語意**，不寫 Phase 1 文件本體。

- [ ] **C1 Target Contract**（→ `calibration-contract`）：Target 分三角色——`normative`（理論／規格要求）、`reference_observation`（官方或可信來源的實際觀測）、`implementation`（產品實作希望達到的值）。每個 Target 記 `source／authority`、`definition`、`sample_scope`、`uncertainty`、`decision_owner`、`provisional｜confirmed`。調參只對 `implementation` Target 收斂；`reference_observation` 只以其 uncertainty 區間作比對，不當精確值
- [ ] **C2 Objective Contract**（→ `calibration-contract`）：指標分三層——`hard_constraint`（必須達到，不可被其他指標抵銷）、`primary_objective`（主要校準目標）、`diagnostic_metric`（解釋偏差與副作用，不計入成功判定）。禁止以單一加總分數判定成功，以免掩蓋 hard constraint 缺口
- [ ] **C3 Sample Sufficiency**（→ `calibration-contract`）：每個指標宣告 `precision_target`（契約要求的估計精度）與對應估計方法——均值類用標準誤、比例類用二項信賴區間、罕見事件看**事件數**而非手數、相關係數看其不確定性、尾部看稀有事件計數。`≥5000 基礎手` 降為 provisional default，不再是通用「充足」判定
- [ ] **C3b Tolerance 定義**（→ `calibration-contract`，回答 Q7）：`max(3×SE, 10%)` 的 `10%` 拆為 `absolute_tolerance` 與 `relative_tolerance`，由各指標定義宣告採用哪一種及基準（target 值 vs 觀測值）；低基率指標（如 1% 命中率）不得套用絕對百分點容忍度
- [ ] **C4 Holdout Policy**（→ `verification`）：seed 分三組——`training_seed_set`（優化器可用）、`validation_seed_set`（候選參數驗證）、`final_holdout_seed_set`（最終驗收，不得參與調參）；`holdout_policy` 記誰產生、何時揭露（final holdout 於**候選參數凍結後**才揭露）、是否允許重跑及重跑規則。語意約束（寫入 `verification` 的 holdout policy，不新增欄位）：final holdout **不得參與參數搜尋**；驗收失敗後**不得**以同一組 final holdout 反覆調參；一旦 final holdout 被用於診斷或修正，即**失去 final holdout 身分**（降為 validation 用），下一輪驗收須重新建立新的 final holdout
- [ ] **C5 Verification Independence**（→ `verification` + `calibration-record`）：獨立性分四維——context（fresh context）、data（未見過的 seed／樣本）、measurement（不共用同一份指標計算程式，或已以獨立實作交叉核對）、decision（決策者≠執行者）。四維是**記錄維度，不是四個必勾 Gate**：Record 記錄**實際達到哪幾維**；`verification` contract 為每次驗收宣告 `required_independence`（本次驗收的最低獨立性要求），未達最低要求才判未驗證。如此不因 consumer 環境限制把所有任務判成無法驗證，也不讓 fresh context 單獨被誤認為完整獨立驗證
- [ ] **C6 Calibration Status**（→ `calibration-record`）：五個狀態是**互斥的最終分類**，每份 Record 恰好一個。`verified`＝所有 hard constraint 通過且驗證契約完成（含 C5 `required_independence`）；`partially_calibrated`＝部分目標達成、仍有已知缺口（缺口已入 Gap Ledger）；`inconclusive`＝證據不足，無法判定；`structural_blocked`＝**已確認**目前結構限制阻礙目標達成；`failed`＝驗收條件未通過，且不屬於上述其他狀態。`structural_blocked` 不得只因某次參數搜尋失敗成立，必須具備 C7 全部判定證據，證據不足時應判 `inconclusive`。判定順序（保證互斥）：證據不足→`inconclusive`；C7 證據齊備→`structural_blocked`；hard constraint 全過且驗證契約完成→ primary objective 全達為 `verified`、否則 `partially_calibrated`；其餘→`failed`。`partially_calibrated` 與 `inconclusive` 是合法收工狀態
- [ ] **C7 Structural Ceiling 診斷**（→ `execution-flow` 診斷步驟 + `gap-ledger`）：當目前模型結構無法表達目標分布時，停止參數搜尋、轉提結構變更需求；判定須記：無法表達的指標、已嘗試參數範圍、為何不是「參數未找到」、需要的抽樣能力或耦合機制。產生 C6 `structural_blocked`
- [ ] **C8 Experiment Ledger**（→ `calibration-record` 的**可選子結構**）：只記重要調參實驗——假設、變更參數、前後指標、波及的其他指標、是否改善、副作用、採納／拒絕理由。Phase 3 dogfood 證明需要獨立管理後才拆檔
- [ ] **C9 Reproducibility Reference**（→ `calibration-record`）：至少一個 `reproducibility_reference` 指回原始實驗；建議涵蓋 consumer build／commit、sampler／strategy version、parameter snapshot hash、metric definition version、seed set id、sample count、source data version、execution environment、measurement timestamp。實際值可由 consumer 提供 reference，不必全進 canonical schema
- **Gate**：C1–C9 每項的欄位名與語意已凍結並寫入本節附註；沒有任何一項導致新增 contract 檔或新 phase
- [ ] **Phase 0 語意核對 S1（C6 判定順序，2026-11-27 第三輪審閱列為候選裁決）**：判定順序保留為候選裁決，Phase 0 只核對一個邊界——「證據不足→`inconclusive`」僅指**無法對 hard constraint／驗證契約要求的指標做出可靠判定**；非必要指標（如 diagnostic metric）缺資料不得無條件覆蓋整份 Record。指標缺資料時是否阻擋 `verified`，由該指標的層級（C2）與驗證契約決定。核對結果寫入 Phase 0.0 表
- **Design freeze（2026-11-27 第二輪審閱）**：Phase 0.1 定位為契約語意補強，**不是實作授權**；C1–C9 已涵蓋審閱建議，Phase 0 期間不再擴充架構概念，只做既有治理檢查。Phase 1（四份 contract、workflow 骨架、validation scenarios）等 Phase 0 Gate 通過後才開工

- [ ] 完成 `workflow/software-delivery/requirements/pre-build-interrogation.md` 與 `plans/README.md` Architecture Compatibility Preflight，記錄於本節
- [ ] 讀 `workflow/workflow-routing.md` 與 `workflow/cross-cutting/README.md`：確認「doc-only pilot 加選路表列但不註冊 registry」是否合規；若不合規改走 `cross-cutting/` 或只加索引
- [ ] 讀 `intelligence/README.md`、`workflow/cross-cutting/traceable-evidence/`、`workflow/image-processing/`（Execute≠Verify 契約）作為重用來源；列 duplication risk
- [ ] 逐項把 Inventory I1–I11 去專案化草稿過 `enforcement/sanitization.md`；決定每項最終落點並回寫 Inventory
- [ ] 取得消費專案方同意：個案流水帳留在消費專案、canonical 只收方法（Q3）
- **Gate**：Q1／Q3 已 resolved；Inventory 每項有最終落點；preflight 無未解衝突；Phase 0.1 Gate 通過

### Phase 1 — Workflow 骨架 + Contracts + Scenarios（doc-only）

- [ ] `workflow/probability-calibration/`：README（含 §Scope 的範圍定義）、`execution-flow.md`（Target→取證→指標對照→診斷〔含 structural ceiling 出口〕→調參→獨立驗證→Gap 登錄→收工回饋）
- [ ] `contracts/`：`calibration-contract`（C1–C3b）、`calibration-record`（C5–C6、C8–C9）、`gap-ledger`（C7）、`verification`（C4–C5、釘值分層）——欄位依 Phase 0.1 凍結版
- [ ] `registry/failure-classes`：由 I10、I11、I2（基準選錯）、I3（跳過診斷改賠付）起步
- [ ] `profiles/slot-rtp.md`：第一個 consumer profile（指標詞彙、基準回合類型、系列拆分；不放個案）
- [ ] ≥3 個 validation scenarios；`ai-skill runtime refresh` 通過。候選（Phase 1 擇 ≥3）：holdout 缺失被判未驗證、target 取自熱樣本被擋、鏈式驗證失敗不得宣稱通過、final holdout 在凍結前揭露被判無效、僅 fresh context 卻宣稱完全獨立被擋、hard constraint 未達卻以總分宣稱 verified 被擋
- [ ] 更新 `workflow/README.md`、`workflow/cross-cutting/README.md`（若採用）與連動文件（`enforcement/linked-updates.md` 要求者）
- **Gate**：scenarios 通過；sanitization 掃描乾淨；獨立（fresh context）讀者只憑 README 能完成一次紙上校準走查

### Phase 2 — Intelligence atoms

- [ ] `intelligence/engineering/probability-calibration/`：I6 結構天花板、I7 單旋鈕副作用、I8 取樣機 Strategy 設計模式、I9 優化操作衛生；各 ≤300 行、單一主題、含 `適用條件／反例`
- [ ] 每個 atom 標註來源個案數（目前 N=1）與升格門檻；N<3 者標 `observation`，不當硬規則
- [ ] 更新 `intelligence/README.md` 與連動索引
- **Gate**：atoms 皆有 Counter-case／不適用條件；無專案識別字樣

### Phase 3 — Dogfood（回饋環驗證）

- [ ] 靜態走查：用 workflow 重放已完成的校準，記錄 workflow 漏掉什麼（回寫契約）
- [ ] 實戰：下一次消費專案校準任務完整走本 workflow，**驗收由 fresh-context verifier 以 final holdout seed set 執行**，產出 Calibration Record（含 C5 獨立性實際達成維度、C6 status）+ Gap Ledger
- [ ] 評估 C8 Experiment Ledger 是否需要從 Record 拆出獨立管理（有證據才拆）
- [ ] 收工 FeedbackDecision；新失敗進 registry 或 evidence candidate（N≥3 才升級）
- [ ] 若 Q2 有第二 consumer，再跑一次並評估是否中立化命名
- **Gate**：≥1 次實戰收工且驗收由獨立角色完成；回饋環至少產生 1 筆 registry 更新或明確「無新學習」紀錄

### Phase 4 — 評估與收口

- [ ] 依 Graduation deadline 決定升級（route + signal）／維持 doc-only／降級
- [ ] ADR Promotion Criteria 逐條勾核
- [ ] 執行 Plan Completion Closure（`plans/README.md`）

## Open Questions

| # | 問題 | 狀態 | 預裁決／預設 |
| --- | --- | --- | --- |
| Q1 | domain 名稱與範圍：`probability-calibration`（中立）vs slot 專用 | user-ruled 2026-11-27 | 中立名稱，保留 slot RTP profile；README Scope 寫明涵蓋 model calibration & validation（§Scope） |
| Q2 | 是否有第二種遊戲類型要用同一套（撲克、抽卡、捕魚等）？決定中立化是否值得、以及 Phase 3 第二 consumer | user-ruled 2026-11-27 | **不阻塞 Phase 0**；承認 N=1，第二 consumer 待確認，不為證明抽象而硬找案例；第一 consumer 完整走完 workflow 且標明哪些仍是 observation 即可進第一輪 dogfood |
| Q3 | 個案流水帳／快照留在消費專案，canonical 只收方法——消費專案方是否同意 | user-ruled 2026-11-27 | 同意（避免雙 SoT 與去敏衝突） |
| Q4 | Target RTP 由誰決策、如何記錄（沿用 decision-support Pass 1／2 或另列） | user-ruled 2026-11-27 | 沿用 decision-support，provisional→confirmed，明確記錄 authority（C1 `decision_owner`） |
| Q5 | 機械鎖歸屬：釘值測試與 governance 檢查由 consumer 實作，Ai-skill 只給 checklist／scenario | user-ruled 2026-11-27 | consumer 實作；Ai-skill 定義契約與 scenario |
| Q6 | 比對門檻（`max(3×SE, 10%)`、樣本 ≥5000）要當預設值還是純參數（來源為「初版待校正」） | user-ruled 2026-11-27 | provisional default，不當通用真理；充足判定改由 `precision_target`（C3） |
| Q7 | `max(3×SE, 10%)` 的 `10%` 是相對還是絕對？基準是 target 還是觀測值？ | user-ruled 2026-11-27 | 拆 `absolute_tolerance`／`relative_tolerance`，由指標定義宣告（C3b） |

## Stakeholder 同意項目

| 項目 | 狀態 |
| --- | --- |
| 在 Ai-skill 新增 `workflow/probability-calibration/`（doc-only pilot，不註冊 route） | 方向同意（2026-11-27 審閱）；Phase 1 開工待 Phase 0 Gate |
| 新增 `intelligence/engineering/probability-calibration/` atoms | 方向同意（2026-11-27 審閱，第一輪維持 observation）；Phase 2 開工待 Phase 1 Gate |
| 消費專案校準文件改為「個案流水帳 + 連結方法」而非整份搬入 | 同意（Q3） |

## 完成條件

依 `governance/lifecycle/system-upgrade-governance.md` §2 子集：

- [ ] 計畫書本身含 Decision Rationale／Runtime Execution Path／Open Questions／Phase 0.0
- [ ] 新目錄 README 與索引（`workflow/README.md`、`intelligence/README.md`）已更新
- [ ] `plans/README.md` 目前狀態表已更新
- [ ] sanitization 掃描通過；無專案識別、無本機絕對路徑
- [ ] Validation scenarios 通過；`ai-skill runtime refresh` 無新 orphan／drift
- [ ] Phase 3 實戰收工證據齊備（獨立驗證）
- [ ] 執行 Plan Completion Closure

## Watch-Out List citation

防 scope drift／over-engineering：不做自動優化器、不放機率引擎程式碼、不做 route／signal 先行、不把單一個案升為通則（N<3 一律 observation）。Phase 0.1 補強只加欄位，不新增 contract 檔、phase 或 workflow；I6–I9 atoms 第一輪維持 observation（2026-11-27 審閱確認）。

## 與其他 plans 的關係

- `2026-08-14-1101-investment-research-decision-support`：只共享「機率表述」邊界，不重疊；本計畫處理「校準」而非「表述」。
- `2026-10-14-0832-image-processing-decision-workflow`：重用 Execute≠Verify／`inconclusive` 合法等契約精神；不複製內容。
- `2026-06-16-1131-evidence-candidate-system`：N≥3 升級紀律的來源；本計畫不改其 schema。
- `workflow/cross-cutting/decision-support/`：Target 決策人與 provisional→confirmed 的承接；本計畫不計入其 converged case 數。
