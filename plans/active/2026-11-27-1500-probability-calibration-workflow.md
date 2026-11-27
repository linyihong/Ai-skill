---
id: 2026-11-27-1500-probability-calibration-workflow
plan_kind: main
status: in-progress
owner: workflow
owner_layer: workflow
created: 2026-11-27
parent: null
---

# Probability Calibration Workflow（`workflow/probability-calibration/` + intelligence）

**Status**: in-progress（Phase 2 完成；Phase 3 dogfood 進行中：consumer 回饋 2 輪，尚無滿足 Gate 的實戰收工）— 2026-11-27 使用者明確授權 Phase 1（範圍：四份 contract、README＋execution-flow、slot-rtp profile、failure classes、≥3 scenarios、索引連動、sanitization＋runtime refresh＋Phase 1 Gate；限制：不新增 contract 檔、不註冊 route／signal、不實作引擎／優化器、N=1 不升格、Gate 未過不進 Phase 2）。2026-11-27 使用者審閱已裁決 Q1–Q6（見 §Open Questions）並要求 Phase 0 先做契約補強（§Phase 0.1）；2026-11-27 Phase 0 完成、Gate 通過（§Phase 0 紀錄）；

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
| I10 | 誠實回報 | 鏈式驗證任一機械鎖失敗＝整體失敗，不得口頭沖掉；已知缺口必須寫入 Gap Ledger（例：絕對占比未達、單一轉輪組、合成條帶） | workflow-local failure class（連回 `evidence-suppression`）+ contract |
| I11 | 過期宣稱 | 舊文件的錯誤頻率（量級錯 6 倍）會被後續工作沿用 → 引用前需回查來源與日期 | workflow-local failure class（連回 `stale-derived-state`） |

### Design Fit Check（命名／分層依據，2026-11-27 使用者要求「看設計能不能抽象」）

使用者裁決規則：設計若偏老虎機就用老虎機命名；除非能抽出一層對每種遊戲都通用的模型，才用中立。
逐項檢驗 Inventory：**可以抽出中立 core，slot 專屬內容可乾淨隔離成 profile → 維持中立命名（與 Q1 一致）。**

| 層 | Inventory 項 | 為何屬於此層 |
| --- | --- | --- |
| **中立 core** | I1 target 定義、I3 診斷順序、I4 樣本充足、I5 驗證分層／holdout、I9 優化衛生、I10 誠實回報、I11 過期宣稱 | 原則只依賴「目標分布 + 可重現產生器 + 一組指標」，與盤面、線、符號無關 |
| **core 原則 ＋ profile 實例** | I2 指標對照、I6 結構天花板、I7 單旋鈕副作用 | 原則通用（指標須同定義；產生器可表達集合有上限；調一旋鈕必量全指標）；實例（可見格 vs 整盤、跨軸相關、同類連線抬高 RTP）屬 slot |
| **中立設計模式 ＋ profile 內策略** | I8 取樣機 Strategy | 「變異點做成資料 + 註冊式 Strategy + 零 RNG 漂移護欄」通用；條帶式、加權格位、耦合條帶等具體策略屬 slot |
| **slot profile 專屬** | base／feature（免費轉、respin）拆分、線賠與賠付倍數分檔、注額縮放參數、基準回合類型、熱帳號旋轉遞減 | 依賴盤面、賠付結構與官方採集管線 |

**Q2 紙上檢驗（非實證）**：以兩個非 slot 場景代入 core，檢查契約是否仍成立。

| 場景 | Target 結構（C1 角色） | core 是否適用 | 需要的差異 |
| --- | --- | --- | --- |
| 抽卡／掉落（含保底） | 公告機率＝`normative`，**無** `reference_observation` | 適用：指標＝單抽率、期望抽數；結構天花板＝獨立抽樣無法表達保底計數器 | 診斷順序**無「採集偏差」步驟**（參照是規格不是樣本） |
| 牌類（洗牌／發牌公平性） | 理論分佈＝`normative`，無 `reference_observation` | 適用：指標＝牌型頻率；驗收＝偏差檢定 + holdout seed | 同上；I1 熱樣本警告不適用 |
| slot（本案） | 官方採集＝`reference_observation`，產品目標＝`implementation` | 適用，且是唯一需要完整診斷順序的情境 | — |

**設計結論**（不新增欄位，沿用 C1 角色，符合 Phase 0.1 Design freeze）：依 Target 是否具 `reference_observation` 裁剪——I3 有則走「定義不一致 → 採集偏差 → 產品偏」全序、無則只走「定義不一致 → 產品偏」；I1 熱樣本限制同樣只在具 `reference_observation` 時適用。此差異只有代入非 slot 場景才看得到。

**誠實限制**：目前只有 1 個真實 consumer，中立 core 是**紙上驗證**，不是實證；slot 衍生原則維持 `candidate-intelligence`（N=1），待第二個真實 consumer（任何類型）走完才升級。Phase 1 只落地 core + `slot-rtp` profile，**不**預建其他 profile 空殼。

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
| C. 只寫 slot 專用 workflow（`workflow/slot-calibration/`） | **reject** — §Design Fit Check 顯示 target 定義、診斷順序、holdout、Gap Ledger 與遊戲類型無關，可抽出中立 core；slot 內容隔離為 profile 即可 |
| D. 中立 workflow + intelligence atoms + 回饋環，doc-only pilot | **accept（draft）** |
| E. 一開始就建 route／discovery signal／優化器 | **reject** — 違反 Loop-first 與 cross-cutting 的「3 個 converged case 才升格」紀律 |

### Why Not an ADR Yet

尚無第二個 consumer、無 Phase 3 dogfood 證據；中立 core 僅紙上驗證（§Design Fit Check）。Plan + doc-only 契約穩定且 ≥2 次獨立校準使用後再評估。

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
| Validation scenarios | Phase 1 以 scenario-first 方式補 ≥3 個 `validation/scenarios/probability-calibration/*.yaml`（格式沿用 `validation/scenarios/investment/*.yaml`；`image-processing` 無 scenario 先例），內容：holdout 缺失被判未驗證、target 取自熱樣本被擋、鏈式驗證失敗不得宣稱通過 |
| Test passing evidence | Phase 1／3 收尾填入（`ai-skill runtime refresh` + scenario validate 輸出） |

### Per-surface consumer 表

| Generated surface key | Named consumer(s) | Consumer 類型 |
| --- | --- | --- |
| （本計畫不新增 `route.*`、`runtime_projection.target_key` 或 commit-msg validator） | n/a | `manual_activation: doc-only pilot；理由見上；接入 phase 與 deadline 已列` |

## Phases

### Phase 0 — 盤點與契約凍結（Pre-Build Interrogation + Preflight）

#### Phase 0.0 — Open Questions 核對（公版，必填）

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [x] `resolved` 的條目已同步勾選／附註於 §Open Questions
- [x] 若盤點新發現問題，已加入 §Open Questions（盤點未產生新 Open Question；兩處修正見 §Phase 0 紀錄 F1／F2）

| Open Question | 處置 | 證據／原因 |
| --- | --- | --- |
| Q1 | resolved | 使用者裁決 + preflight 佐證：`enforcement/failure-patterns/first-consumer-names-a-shared-capability-after-itself.md` 支持「能力以中立名稱、第一 consumer 只當 profile」 |
| Q2 | resolved | 使用者裁決：不阻塞；N=1 已在 Inventory 與 Phase 2 標示 |
| Q3 | resolved | 使用者裁決；與 `reusable-guidance-boundary`／sanitization 一致（個案留 consumer） |
| Q4 | resolved | 使用者裁決；`workflow/cross-cutting/decision-support/` 的 provisional→confirmed 兩趟式可直接承接 C1 `decision_owner`；本計畫不計入其 converged case |
| Q5 | resolved | 使用者裁決；與 `image-processing` 先例一致（Ai-skill 只給 contract，不提交產品腳本） |
| Q6 | resolved | 使用者裁決；C3 已把門檻降為 provisional default |
| Q7 | resolved | 使用者裁決；C3b |
| S1 | resolved | 見 §Phase 0 紀錄 S1 裁決 |

#### Phase 0.1 — 契約補強（2026-11-27 使用者審閱；吸收進既有 4 份 contract，不新增文件）

> 原則：以下全部落在 Phase 1 既有的 `calibration-contract`／`calibration-record`／`verification`／`gap-ledger` 內，不新建 workflow、phase 或獨立契約檔（沿用 minimize-governance-surface）。Phase 0.1 只凍結**欄位與語意**，不寫 Phase 1 文件本體。

- [x] **C1 Target Contract**（→ `calibration-contract`）：Target 分三角色——`normative`（理論／規格要求）、`reference_observation`（官方或可信來源的實際觀測）、`implementation`（產品實作希望達到的值）。每個 Target 記 `source／authority`、`definition`、`sample_scope`、`uncertainty`、`decision_owner`、`provisional｜confirmed`。調參只對 `implementation` Target 收斂；`reference_observation` 只以其 uncertainty 區間作比對，不當精確值
- [x] **C2 Objective Contract**（→ `calibration-contract`）：指標分三層——`hard_constraint`（必須達到，不可被其他指標抵銷）、`primary_objective`（主要校準目標）、`diagnostic_metric`（解釋偏差與副作用，不計入成功判定）。禁止以單一加總分數判定成功，以免掩蓋 hard constraint 缺口
- [x] **C3 Sample Sufficiency**（→ `calibration-contract`）：每個指標宣告 `precision_target`（契約要求的估計精度）與對應估計方法——均值類用標準誤、比例類用二項信賴區間、罕見事件看**事件數**而非手數、相關係數看其不確定性、尾部看稀有事件計數。`≥5000 基礎手` 降為 provisional default，不再是通用「充足」判定
- [x] **C3b Tolerance 定義**（→ `calibration-contract`，回答 Q7）：`max(3×SE, 10%)` 的 `10%` 拆為 `absolute_tolerance` 與 `relative_tolerance`，由各指標定義宣告採用哪一種及基準（target 值 vs 觀測值）；低基率指標（如 1% 命中率）不得套用絕對百分點容忍度
- [x] **C4 Holdout Policy**（→ `verification`）：seed 分三組——`training_seed_set`（優化器可用）、`validation_seed_set`（候選參數驗證）、`final_holdout_seed_set`（最終驗收，不得參與調參）；`holdout_policy` 記誰產生、何時揭露（final holdout 於**候選參數凍結後**才揭露）、是否允許重跑及重跑規則。語意約束（寫入 `verification` 的 holdout policy，不新增欄位）：final holdout **不得參與參數搜尋**；驗收失敗後**不得**以同一組 final holdout 反覆調參；一旦 final holdout 被用於診斷或修正，即**失去 final holdout 身分**（降為 validation 用），下一輪驗收須重新建立新的 final holdout
- [x] **C5 Verification Independence**（→ `verification` + `calibration-record`）：獨立性分四維——context（fresh context）、data（未見過的 seed／樣本）、measurement（不共用同一份指標計算程式，或已以獨立實作交叉核對）、decision（決策者≠執行者）。四維是**記錄維度，不是四個必勾 Gate**：Record 記錄**實際達到哪幾維**；`verification` contract 為每次驗收宣告 `required_independence`（本次驗收的最低獨立性要求），未達最低要求才判未驗證。如此不因 consumer 環境限制把所有任務判成無法驗證，也不讓 fresh context 單獨被誤認為完整獨立驗證
- [x] **C6 Calibration Status**（→ `calibration-record`）：五個狀態是**互斥的最終分類**，每份 Record 恰好一個。`verified`＝所有 hard constraint 通過且驗證契約完成（含 C5 `required_independence`）；`partially_calibrated`＝部分目標達成、仍有已知缺口（缺口已入 Gap Ledger）；`inconclusive`＝證據不足，無法判定；`structural_blocked`＝**已確認**目前結構限制阻礙目標達成；`failed`＝驗收條件未通過，且不屬於上述其他狀態。`structural_blocked` 不得只因某次參數搜尋失敗成立，必須具備 C7 全部判定證據，證據不足時應判 `inconclusive`。判定順序（保證互斥）：證據不足→`inconclusive`；C7 證據齊備→`structural_blocked`；hard constraint 全過且驗證契約完成→ primary objective 全達為 `verified`、否則 `partially_calibrated`；其餘→`failed`。`partially_calibrated` 與 `inconclusive` 是合法收工狀態
- [x] **C7 Structural Ceiling 診斷**（→ `execution-flow` 診斷步驟 + `gap-ledger`）：當目前模型結構無法表達目標分布時，停止參數搜尋、轉提結構變更需求；判定須記：無法表達的指標、已嘗試參數範圍、為何不是「參數未找到」、需要的抽樣能力或耦合機制。產生 C6 `structural_blocked`
- [x] **C8 Experiment Ledger**（→ `calibration-record` 的**可選子結構**）：只記重要調參實驗——假設、變更參數、前後指標、波及的其他指標、是否改善、副作用、採納／拒絕理由。Phase 3 dogfood 證明需要獨立管理後才拆檔
- [x] **C9 Reproducibility Reference**（→ `calibration-record`）：至少一個 `reproducibility_reference` 指回原始實驗；建議涵蓋 consumer build／commit、sampler／strategy version、parameter snapshot hash、metric definition version、seed set id、sample count、source data version、execution environment、measurement timestamp。實際值可由 consumer 提供 reference，不必全進 canonical schema
- **Gate**：C1–C9 每項的欄位名與語意已凍結並寫入本節附註；沒有任何一項導致新增 contract 檔或新 phase
- [x] **Phase 0 語意核對 S1（C6 判定順序，2026-11-27 第三輪審閱列為候選裁決）**：判定順序保留為候選裁決，Phase 0 只核對一個邊界——「證據不足→`inconclusive`」僅指**無法對 hard constraint／驗證契約要求的指標做出可靠判定**；非必要指標（如 diagnostic metric）缺資料不得無條件覆蓋整份 Record。指標缺資料時是否阻擋 `verified`，由該指標的層級（C2）與驗證契約決定。核對結果寫入 Phase 0.0 表
- **Design freeze（2026-11-27 第二輪審閱）**：Phase 0.1 定位為契約語意補強，**不是實作授權**；C1–C9 已涵蓋審閱建議，Phase 0 期間不再擴充架構概念，只做既有治理檢查。Phase 1（四份 contract、workflow 骨架、validation scenarios）等 Phase 0 Gate 通過後才開工

- [x] 完成 `workflow/software-delivery/requirements/pre-build-interrogation.md` 與 `plans/README.md` Architecture Compatibility Preflight，記錄於本節
- [x] 讀 `workflow/workflow-routing.md` 與 `workflow/cross-cutting/README.md`：確認「doc-only pilot 加選路表列但不註冊 registry」是否合規；若不合規改走 `cross-cutting/` 或只加索引
- [x] 讀 `intelligence/README.md`、`workflow/cross-cutting/traceable-evidence/`、`workflow/image-processing/`（Execute≠Verify 契約）作為重用來源；列 duplication risk
- [x] 逐項把 Inventory I1–I11 去專案化草稿過 `enforcement/sanitization.md`；決定每項最終落點並回寫 Inventory
- [x] 取得消費專案方同意：個案流水帳留在消費專案、canonical 只收方法（Q3）

#### Phase 0 紀錄（2026-11-27）

| 欄位 | 內容 |
| --- | --- |
| Trigger | 使用者指示繼續；執行本 plan Phase 0 |
| Checked sources | `plans/README.md` §Preflight、`pre-build-interrogation.md`、`workflow/workflow-routing.md`、`workflow/README.md`、`workflow/cross-cutting/README.md`、`cross-cutting/traceable-evidence/`、`cross-cutting/decision-support/`、`image-processing/`（README、`contracts/verification.yaml`、`finality.yaml`、`image-decision-record.yaml`、`registry/failure-classes.yaml`）、`intelligence/README.md`、`intelligence/engineering/anti-patterns/`、`enforcement/failure-patterns/`、`enforcement/sanitization.md`、`validation/scenarios/` |
| Interrogation | Goal：可重用的機率模型校準與驗證方法；Scope：doc-only workflow + 4 contracts + failure-class registry + slot RTP profile + scenarios；Non-goals：route／signal、優化器、引擎程式碼、個案資料；Acceptance：Phase 1 Gate（scenarios、sanitization、紙上走查）；Framework discovery：canonical = workflow-local YAML／Markdown，無 runtime projection、無 generated surface；Assumptions：見 F1–F4 |
| Routing 合規 | 合規。先例：`image-processing`（列於選路表 §常見歧義、未註冊 route）、`narrative-video-production`（選路表標「未註冊」）。不走 `cross-cutting/`：本案是有自身 lifecycle 與產出的 domain workflow，不是跨 slice 共用 concern |
| Duplication risk | 無雙 SoT。Execute≠Verify、`inconclusive` 合法沿用 `image-processing` 精神，但為不同 domain 的獨立契約；I5「釘值≠校準」是 `validation-proxy-trap` 的領域實例、I10 是 `evidence-suppression` 的領域實例、I11 是 `stale-derived-state` 的領域實例 → 均以 workflow-local failure class **連回**既有 anti-pattern，不新增 enforcement failure pattern。TE1／TE2／TE4 由 C8／C9／C1 實作並連回 traceable-evidence |
| Terminology watch | C6 `calibration_status` 與 image-processing `finality` 皆含 `inconclusive` 等詞，語意相近但所屬記錄不同；依 terminology-collision 紀律只記錄、不統一 |
| Sanitization | 本 plan 無專案名、路徑、host、個人資料；「slot」為遊戲類型泛稱。I9／I11 的數值（20k 手 SE≈0.03、量級錯 6 倍）Phase 1 只能以「示例值」出現 |
| Decision | proceed（待使用者授權 Phase 1） |
| Validation | 本節 diff review；commit hooks |

**修正（F）**：
- **F1** Runtime Execution Path 原寫「沿用 `investment`／`image-processing` 格式」：`image-processing` 沒有 validation scenarios，只有 `validation/scenarios/investment/*.yaml` 可當格式先例 → 已改正
- **F2** Phase 2 原寫 atoms 標 `observation`：intelligence 層既有狀態只有 `candidate-intelligence`／`validated-intelligence` → atoms 用 `candidate-intelligence` 並註明來源個案 N=1，不新增狀態值
- **F3** Phase 1 contracts 檔案格式依 `image-processing` 先例：`artifact-record/v1` YAML、`runtime_projection.enabled: false`；failure classes 用 `registry/v1`
- **F4** I10／I11 落點由「failure pattern」改為 workflow-local failure class + 連回既有 anti-pattern（見 Duplication risk）
- **F5**（Gate 通過後的補記，2026-11-27 使用者再問命名／分層）：新增 §Design Fit Check，確認中立命名成立；唯一語意補充是 I3 診斷順序依 Target 是否具 `reference_observation` 裁剪，**沿用 C1 角色、不新增欄位／契約／phase**，未改動 Phase 0 Gate 結論

**S1 裁決（C6 證據不足的範圍）**：指標是否「必要」直接由 C2 層級決定，不新增欄位——`hard_constraint` 與 `primary_objective` 是必要指標，`diagnostic_metric` 不是。必要指標未達 C3 `precision_target` → `inconclusive`；diagnostic metric 樣本不足只寫入 Gap Ledger（標 `insufficient_evidence`），不影響狀態。所以 hard constraint 與驗證契約都已完成、只有 diagnostic 缺資料時，仍可判 `verified`

**Gate 判定**：Q1／Q3 resolved ✓；Inventory 每項有最終落點 ✓；preflight 無未解衝突（F1–F4 已就地修正）✓；Phase 0.1 Gate ✓ → **Phase 0 Gate 通過**。Phase 1 開工仍待使用者明確授權（方向同意≠開工授權）
- **Gate**：Q1／Q3 已 resolved；Inventory 每項有最終落點；preflight 無未解衝突；Phase 0.1 Gate 通過

### Phase 1 — Workflow 骨架 + Contracts + Scenarios（doc-only）

- [x] `workflow/probability-calibration/`：README（含 §Scope 的範圍定義）、`execution-flow.md`（Target→取證→指標對照→診斷〔含 structural ceiling 出口；「採集偏差」步驟僅在 Target 具 `reference_observation` 時適用，見 §Design Fit Check〕→調參→獨立驗證→Gap 登錄→收工回饋）
- [x] `contracts/`（`artifact-record/v1` YAML、`runtime_projection.enabled: false`，依 image-processing 先例）：`calibration-contract`（C1–C3b）、`calibration-record`（C5–C6、C8–C9）、`gap-ledger`（C7）、`verification`（C4–C5、釘值分層）——欄位依 Phase 0.1 凍結版
- [x] `registry/failure-classes`：由 I10、I11、I2（基準選錯）、I3（跳過診斷改賠付）起步；另加 I5（釘值當校準，連回 `validation-proxy-trap`）
- [x] `profiles/slot-rtp.md`：第一個 consumer profile（指標詞彙、基準回合類型、系列拆分；不放個案）
- [x] ≥3 個 validation scenarios；`ai-skill runtime refresh` 通過。候選（Phase 1 擇 ≥3）：holdout 缺失被判未驗證、target 取自熱樣本被擋、鏈式驗證失敗不得宣稱通過、final holdout 在凍結前揭露被判無效、僅 fresh context 卻宣稱完全獨立被擋、hard constraint 未達卻以總分宣稱 verified 被擋
- [x] 更新 `workflow/README.md`、`workflow/cross-cutting/README.md`（若採用）與連動文件（`enforcement/linked-updates.md` 要求者）
- **Gate**：scenarios 通過；sanitization 掃描乾淨；獨立（fresh context）讀者只憑 README 能完成一次紙上校準走查
- **Phase 1 Gate 紀錄（2026-11-27）**：4 scenarios 結構合規（schema 必填鍵、type enum）且 YAML 可解析；sanitization grep 乾淨；`ai-skill runtime refresh` success。Fresh-context 讀者（只讀 `workflow/probability-calibration/`，0 個目錄外檔案）以合成掉落率＋保底案例完成一輪紙上走查，判 `inconclusive` → **Gate 通過**。走查就地修正 4 處語意缺陷：(1) parity 未解／獨立驗證未執行或低於 `required_independence` 原本落到 `failed`，改入 `inconclusive`；(2) `precision_target` 對必要指標改為必填；(3) contract 加 `contract_version`／`frozen_at`；(4) Gap Ledger 加 `process_gap`，execution-flow 補「契約前已調參」處理與 b 步略過須記錄。**使用者裁決（2026-11-27）**：接受 (1)，判定界線凍結於 `calibration-record.yaml` §`status_boundary`（驗證鏈未完成或無效＝`inconclusive`，完成且確認未達標才是 `failed`）；Phase 1 Gate 維持通過，不重開。延到 Phase 3 實證（observation，不擴設計）：有狀態／非 iid 機制（保底）的 estimator 與樣本單位、分布比較型 estimator、holdout 大小與 precision_target 的連動、ad hoc 量測的 seed_set 歸類、「低基率」定義、純規格 Target 的 tier 指引、保底是否屬結構天花板

### Phase 2 — Intelligence atoms

- [x] `intelligence/engineering/probability-calibration/`：I6 結構天花板、I7 單旋鈕副作用、I8 取樣機 Strategy 設計模式、I9 優化操作衛生；各 ≤300 行、單一主題、含 `適用條件／反例`；每個 atom 標 N=1 與升格門檻，用既有 `candidate-intelligence`（不新增狀態值）；更新 `intelligence/README.md` 與連動索引
- **Gate**：atoms 皆有 Counter-case／不適用條件；無專案識別字樣
- **Phase 2 Gate 紀錄（2026-11-27，使用者授權 Phase 2）**：4 atoms（39–44 行）皆 `candidate-intelligence`、標 N=1、分開「原則」與「單一案例觀察」，各有適用條件／反例與升級條件（≥3 獨立案例、至少一例非 slot）；每個 atom 標來源契約或 Inventory 項；與 `shotgun-debugging`／`validation-proxy-trap` 只連結不重建；保底與估計方法仍列 Phase 3 觀察；無 route／signal／優化器／引擎；sanitization grep 乾淨 → **Gate 通過**。驗證責任：atoms 只整理候選知識，實證屬 Phase 3 dogfood

### Phase 3 — Dogfood（回饋環驗證）

> **Dogfood 回饋（2026-11-27，consumer 回報；待 plan 擁有者複核）**
>
> **這一輪的性質**：消費專案在 Phase 1 落地的同一個 session 內做了一輪「新增長連線目標」的調參，**不是**照 workflow 從步驟 0 開始的實戰，而是事後重放（retrofit）。結果：沒有找到可採用的候選、沒有換參數。**不滿足**下方實戰 Gate（無 final holdout、無獨立驗收）。三條 candidate lesson 見 [`feedback/history/probability-calibration/common/`](../../feedback/history/probability-calibration/common/README.md)。
>
> **Consumer 側已自我登記的 `process_gap`**：(PG1) 契約未先凍結——條件化長連線比值在使用者授權該缺口當目標後才在探針裡定義；(PG2) 沒有獨立驗證，所有候選只量在 training seed；(PG3) 執行者自行寫了 `partially_calibrated`，違反「Tune 階段禁止寫最終 status」，已更正為「執行者初判、暫為 `inconclusive`」。
>
> | 步驟 | 重放發現 | 類型 | 建議落點（由 owner 決定） |
> | --- | --- | --- | --- |
> | 0 契約 | 凍結後新增一個目標／指標沒有程序：該改版本重開一輪，還是就地修？ | 缺口 | `calibration-contract`：寫明「新增目標或指標＝新 `contract_version`＋新一輪，舊 Record 保留舊版本」 |
> | 2 指標對照 | 參照側約 380 個單位、產品側約 10 萬單位，直接比「最大值」誤報成缺口；對照清單有分子／分母／基準／聚合單位，沒有**比較單位數 n**；estimator enum 沒有順序統計量 | 缺口（證實 Phase 1 延後的「分布比較型 estimator」） | `execution-flow` 步驟 2 加 n；estimator 加「以參照 n 重抽」；lesson `order-statistic-needs-equal-sample-size` |
> | 1／3 取證與診斷 | 參照是熱樣本；a→b→c 順序與 `hot_sample` 規則**有效**，擋住了把絕對值當目標。但規則只說「取尺度不變結構」，沒說怎麼構造 | 證實＋小補強 | profile／contract：條件化比值＝以前置事件數為分母，precision 看前置事件數；lesson `scale-invariant-metric-for-biased-reference-sample` |
> | 3s 結構天花板 | 搜尋得到的前緣只能填 `parameter_ranges_tried`，**不算** `why_not_parameter_not_found`；契約沒說哪種論證算數 | 缺口 | `gap-ledger`：列出可接受的論證型別（解析上界、縮減參數空間的窮舉、鬆弛），取樣前緣不算；否則維持 `inconclusive` |
> | 4 調參 | 一個 3×3 格點加兩輪貪婪搜尋無法逐點寫進 C8 的單一實驗格式 | 證據（不拆檔） | C8 允許「掃描級」條目（軸、點數、前緣表引用）；C8 獨立成檔仍無證據（N=1） |
> | 4 調參（atoms） | 先做旋鈕篩選（無效／動目標／動硬約束）省了整輪優化；候選用同種子比較，前緣點要在 validation seed 重量 | 補充 | `single-knob-side-effects` 加旋鈕篩選；`optimization-probe-hygiene` 加共同亂數與前緣點重量 |
> | 7 狀態 | Record 只有一個 `calibration_status`，沒有放「執行者初判」的位置，執行者就直接填了最終值；「沒有任何候選可驗」也沒有範例，結果直覺上像 `partially_calibrated`，實際依 `evaluation_order` 第一條是 `inconclusive` | 缺口（本輪實際踩到） | `execution-flow` 步驟 7 加範例；`status_reason` 約定「執行者初判」標記為非最終；新增 scenario「executor 自寫最終 status 被擋」 |
> | 8 缺口登錄 | 「刻意沒調、等人決定是否算目標」的缺口沒有對應 `kind`，目前只能借 `known_model_limitation`＋`open` | 觀察（延後） | 若第二個案例重現再評估 |
>
> **計畫本身的小地方**：(1) §Stakeholder 同意項目第二列仍寫「Phase 2 開工待 Phase 1 Gate」，已過時；(2) §完成條件的勾選框全空，其中索引、`plans/README.md` 狀態列、sanitization、scenarios 已有證據可勾；(3) §Watch-Out 寫 atoms「維持 observation」，與 F2 的 `candidate-intelligence` 用語不一致；(4) 頁首狀態行結尾多一個「；」；(5) 建議新增 Q8：**同一 consumer 的多輪校準算幾個「獨立案例」**——建議算 1，否則 N≥3 與 ADR「≥2 次獨立校準」很容易被同一個專案灌水。
>
> **N 的計法**：本輪與先前的耦合重調同屬同一個 consumer，所有 lesson 與 atoms 仍是 N=1，不升級。

> **Dogfood 回饋 #2（2026-11-27，consumer 回報；與上一輪不同：這一輪**已在 consumer 主幹採用變更**）**
>
> **這一輪的性質**：決策者要求把某玩法的單轉線均值、基礎手線獎與整場尾端壓低，「略低於參照值沒關係」。調參用單旋鈕（元素占比）完成，已合併進 consumer 主幹並更新其釘值測試。同樣**不是**照 workflow 從步驟 0 開始：契約未先凍結、執行者自行量測與採用、無 final holdout、無 fresh-context verifier，所以**不滿足**實戰 Gate。執行者**不寫最終 status**（依上一輪更正的規則）；初判為 `inconclusive`（C5 實際達成維度＝無），最終由 verifier／decider 定。四條 candidate lesson 見 [`feedback/history/probability-calibration/common/`](../../feedback/history/probability-calibration/common/README.md)。
>
> **Consumer 側 `process_gap`**：(PG1) 契約後補；(PG2) 無獨立驗證，量測全在固定 seed 與 training 同源；(PG3) 採用先於驗收，契約沒有承載這種狀態。
>
> | 步驟 | 重放發現 | 類型 | 建議落點（由 owner 決定） |
> | --- | --- | --- | --- |
> | 0 契約／C1 | 實作目標＝`參照觀測值 × 縮放係數`；係數由決策者確認；契約沒有欄位承載係數與其 `decision_owner` | 缺口 | `calibration-contract` C1：`implementation` Target 可記 `derived_from` ＋ `scale` ＋ 決策者；lesson `owner-one-sided-tolerance…` |
> | 0 契約／C3b | 決策者的偏好是單側（略低可接受）；容忍度只有 absolute／relative，沒有方向，驗收無法機械分辨「偏低被允許」與「偏低是失敗」 | 缺口 | C3b 加 `direction`（兩側／僅下界／僅上界）；單側許可仍要記偏差量，入 Gap Ledger |
> | 2 指標對照 | 整體指標（整場）與成分指標（線獎、附加獎）分開對照；參照側自己就有約 5% 殘差，成分對齊後產品整體仍低約 20% | 缺口 | `execution-flow` 步驟 2 先算參照殘差；lesson `aggregate-metric-residual…`；`gap-ledger` 殘差缺口範例（`definition_mismatch`） |
> | 4 調參 | 找出主旋鈕是元素占比；動它前沒有預測稀有事件，事後發現「連續 k 同類」大獎率降一個量級（占比降 40%、k=5 預測倍率 0.078、實測約 0.125） | 補充＋證實 I7 | `single-knob-side-effects` 補預測式與第二個實例；`execution-flow` Tune：動旋鈕前把稀有事件列為診斷指標 |
> | 5 驗證 | 回歸測試共用同一條亂數流，參數變更後無關案例翻紅；改成每案例種子後穩定，再以多種子量出約 0.07% 的潛在缺陷（與本次調參無關） | 缺口 | `verification` 釘值分層：記錄測試種子獨立性；`sampler-strategy-as-data`（零 RNG 漂移護欄）補「案例獨立種子」 |
> | 7 狀態 | 變更已採用，但 Record 只有 `calibration_status`，沒有「採用／未採用」；採用不等於驗收，但 consumer 的釘值測試綠燈容易被讀成驗收 | 缺口（語意） | Record 加 `adoption`（`adopted_ref`／`not_adopted`）與 status 分開；`pinned_value_claimed_as_calibration` failure class 補「採用後測試綠＝驗收」變體 |
> | 8 缺口登錄 | G1 整場尾端低於縮放後參照（open，等決策者確認偏低量是否可接受）；G2 稀有大獎率降一個量級（open，等決策者）；G3 成分占比偏離參照（`known_model_limitation`）；G4 參照殘差（`definition_mismatch`）；G5 無關路徑的潛在缺陷（consumer 缺陷清單，非本 workflow 範圍）；G6 `process_gap`（PG1–PG3）。**`accepted_risk` 需要決策者**，執行者不得自標 | 證實 `accepted_risk_requires_decision_owner` | 無 |
>
> **N 的計法**：與上一輪、先前的耦合重調同屬一個 consumer，全部維持 N=1，不升級。回應上一輪提出的 Q8：本輪同樣支持「同一 consumer 的多輪校準算 1 個獨立案例」，否則第三輪就會讓 N≥3 變成單一專案的內部重複。
>
> **對 plan 的建議（待 owner）**：(1) 若採納本輪 4 個缺口，只在既有 contract 補欄位（`direction`、`scale`／`derived_from`、`adoption`、測試種子獨立性），不新增 contract 檔、不新增 phase；(2) 優先順序建議：`adoption` 與 `direction` 先補（影響判定語意），殘差與種子獨立性可延後；(3) 下一次實戰請改走完整流程：先寫 Contract，再調參，最後交 fresh-context verifier 以 final holdout 驗收。

- [x] 靜態走查：用 workflow 重放已完成的校準，記錄 workflow 漏掉什麼（回寫契約）——已由 consumer 重放一輪（見上方回饋表）；契約回寫由 owner 決定
- [ ] 實戰：下一次消費專案校準任務完整走本 workflow，**驗收由 fresh-context verifier 以 final holdout seed set 執行**，產出 Calibration Record（含 C5 獨立性實際達成維度、C6 status）+ Gap Ledger
- [ ] 評估 C8 Experiment Ledger 是否需要從 Record 拆出獨立管理（有證據才拆）
- [ ] 收工 FeedbackDecision；新失敗進 registry 或 evidence candidate（N≥3 才升級）
- [ ] 若 Q2 有第二 consumer，再跑一次並評估是否中立化命名
- **Gate**：≥1 次實戰收工且驗收由獨立角色完成；回饋環至少產生 1 筆 registry 更新或明確「無新學習」紀錄

### Phase 4 — 評估與收口

- [ ] 依 Graduation deadline 決定升級（route + signal）／維持 doc-only／降級；ADR Promotion Criteria 逐條勾核；執行 Plan Completion Closure（`plans/README.md`）

## Open Questions

| # | 問題 | 狀態 | 預裁決／預設 |
| --- | --- | --- | --- |
| Q1 | domain 名稱與範圍：`probability-calibration`（中立）vs slot 專用 | resolved（user-ruled；Phase 0 核對 2026-11-27） | 中立名稱，保留 slot RTP profile；README Scope 寫明涵蓋 model calibration & validation（§Scope） |
| Q2 | 是否有第二種遊戲類型要用同一套（撲克、抽卡、捕魚等）？決定中立化是否值得、以及 Phase 3 第二 consumer | resolved（user-ruled；Phase 0 核對 2026-11-27） | **不阻塞 Phase 0**；紙上檢驗（抽卡、牌類）見 §Design Fit Check；承認 N=1，第二 consumer 待確認，不為證明抽象而硬找案例；第一 consumer 完整走完 workflow 且標明哪些仍是 observation 即可進第一輪 dogfood |
| Q3 | 個案流水帳／快照留在消費專案，canonical 只收方法——消費專案方是否同意 | resolved（user-ruled；Phase 0 核對 2026-11-27） | 同意（避免雙 SoT 與去敏衝突） |
| Q4 | Target RTP 由誰決策、如何記錄（沿用 decision-support Pass 1／2 或另列） | resolved（user-ruled；Phase 0 核對 2026-11-27） | 沿用 decision-support，provisional→confirmed，明確記錄 authority（C1 `decision_owner`） |
| Q5 | 機械鎖歸屬：釘值測試與 governance 檢查由 consumer 實作，Ai-skill 只給 checklist／scenario | resolved（user-ruled；Phase 0 核對 2026-11-27） | consumer 實作；Ai-skill 定義契約與 scenario |
| Q6 | 比對門檻（`max(3×SE, 10%)`、樣本 ≥5000）要當預設值還是純參數（來源為「初版待校正」） | resolved（user-ruled；Phase 0 核對 2026-11-27） | provisional default，不當通用真理；充足判定改由 `precision_target`（C3） |
| Q7 | `max(3×SE, 10%)` 的 `10%` 是相對還是絕對？基準是 target 還是觀測值？ | resolved（user-ruled；Phase 0 核對 2026-11-27） | 拆 `absolute_tolerance`／`relative_tolerance`，由指標定義宣告（C3b） |

## Stakeholder 同意項目

| 項目 | 狀態 |
| --- | --- |
| 在 Ai-skill 新增 `workflow/probability-calibration/`（doc-only pilot，不註冊 route） | ✅ 2026-11-27 使用者授權 Phase 1 |
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
