---
id: 2026-10-23-2301-knowledge-retrieval-workflow-closure
plan_kind: main
status: in-progress
owner: larrylin/codex-session
created: 2026-10-23T23:01:00+09:00
parent: null
owner_layer: knowledge
---

# 既有知識檢索與 Workflow 採用閉環強化

**Status**: `in-progress`。Phase 0 盤點完成，證據見 [phase0-inventory.md](evidence/phase0-inventory.md)；Phase 1 已按「歷史 lesson 檢索基準」縮限核對通過；原完整範圍的剩餘責任見下表。Phase 2 前執行控制見 [phase2-execution-contract.md](evidence/phase2-execution-contract.md)；Q / E / QE v1 及 v2 對照已完成，候選品質均未通過；v2 成本受控但 confuser 仍 3/6，見 [v2 實驗報告](evidence/phase2-v2-ablation.md)；未採用到 production。Production 搜尋與採用鏈尚未實作。

目標：讓既有 intelligence、決策與失敗經驗更穩定地被找到，經來源與適用性確認後支援 workflow。先改善現有能力，再用漏找證據決定是否加入向量搜尋與 reranking。

## Decision Rationale

### Problem & Why Now

知識能累積不等於執行時能找回。現有 Discovery Bridge 推薦 route，與 workflow 執行中的經驗查找責任不同；只補 detector miss path，無法覆蓋「workflow 已命中但 intelligence 沒找到」。固定歷史 lesson 重播基準已建立：診斷命中 1/14、固定比較正例 0/4；這只界定已量測案例的漏找，尚不能外推一般檢索品質或預估改善幅度。

撰寫時已核對的現況（不是 benchmark 結果）：

| 現有能力 | 來源與限制 |
| --- | --- |
| FTS5 / BM25 與 layer、type、status 過濾 | [runtime.go](../../../scripts/ai-skill-cli/internal/app/runtime.go) 的 `runRuntimeIndexQuery` 已存在；不另建同功能引擎 |
| FTS 欄位 | 同檔 schema 包含 title、summary、tags、when_to_read、validation_signal 等；不能當成全庫正文均已索引 |
| Discovery Bridge | [discovery.go](../../../scripts/ai-skill-cli/internal/app/discovery.go) 與 [hooks.go](../../../scripts/ai-skill-cli/internal/app/hooks.go)：確定性路由未鎖定時給候選建議 |
| 索引與來源邊界 | [SQLite lookup cache](../../../knowledge/runtime/sqlite/README.md)：可重建；候選需要讀回 canonical source |
| Runtime source 分類 | [runtime/README.md](../../../runtime/README.md)：internal config 與 executable contract 分開治理，不把所有 DB 都當可刪除 cache |

### Decision

重用既有檢索能力、分開消費入口：Discovery Bridge 查候選 route；workflow 已啟動後查相關 intelligence、決策與經驗。只有盤點確認重複實作或可量測收益時，才調整 Bridge 呼叫接點；workflow lookup 的交付不以 Bridge 重構為前提。先盤點與量測，再補全文覆蓋、查詢處理、metadata 過濾和有界關聯展開；向量與 reranking 以獨立增益實驗決定。

固定邊界：

- `required_read`、primary source 與規則依賴確定性載入，不受 top-k、搜尋失敗或排名影響。
- Discovery 維持 advisory，不滿足 activation triggers、不自動鎖 route、不修改 routing registry。
- 相似度只表達相關性；不作可信度、規則權重或 lifecycle promotion 判定。
- 索引只衍生自有身分與版本的來源；不把檢索摘要寫成第二份 canonical knowledge。
- 搜尋失敗可回退直接讀取與明確路由，但不得繞過既有 blocker。
- 本計畫不建 hosted vector service、通用 agent framework、自動知識升格或自主改規則系統。

### Alternatives Considered

| 方案 | 處置與原因 |
| --- | --- |
| 直接導入向量 DB + reranker | 暫不採用；無基準，無法辨識覆蓋缺口或評估新增成本 |
| 只增加 routing keywords | 可修精確路由個案，但不解決 workflow 內部找經驗 |
| 全部塞進 Discovery Bridge hook | 不採用；擴張 miss-path 責任並增加每次工具呼叫成本 |
| 重用 SQLite / BM25 後分階段接入 | 採用；既有來源、驗證、CLI 與 cache 邊界可延續 |

### Why Not an ADR Yet / ADR Promotion Criteria

本案是待驗證提案。僅在跨 session、跨專案的採用證據顯示形成難以逆轉的基礎決策、能說明 why 且無更輕承載位置時，才依 decision promotion 流程評估 accepted ADR；plan 文件完成不構成升格條件。

### Consequences

預期改善來源可發現性與採用可追溯性；代價是索引體積、refresh、metadata 維護與查詢成本增加。正文 chunking 可能丟上下文，過度過濾可能排除適用知識；用整份來源回讀、unknown 狀態與負面案例控制。對應 [Gen 4 Watch-Out List](../../../architecture/ai-native-cognitive-ecosystem-system.md#watch-out-listgen-4-路上會撞到的牆) 的 wall 1、2、4、6：不混淆 discovery / activation、不膨脹 workflow、不無界記錄 telemetry、不建自主 optimizer。

**Glossary Impact: no** — no new framework vocabulary introduced；BM25、Recall@5、MRR、embedding 為既有技術術語。本計畫不註冊新 framework 名詞。

## Runtime Execution Path

Runtime owner 為現有 CLI `internal/app`；knowledge 層管理查找來源與關聯，workflow 管理何時使用經驗，governance / enforcement 保有適用性與規則權重。以下為待實作路徑，不能當作已接線證據。

| 入口事件 | Query / detector → 載入 → 動作 → 證據 |
| --- | --- |
| 任務無確定 route | 現有 detector miss → 保留 Bridge 現有 ranking / advisory → 提供 route / primary source；只有符合 Phase 3 接入條件才改用共用查詢，並記錄比較證據；不 activation |
| Workflow 已啟動且需要相關經驗 | workflow 明確查找步驟 → `runtime query` 共用檢索與 context filters → 讀 canonical source、核對回讀版本與適用性 → 採用或拒絕並記理由與實際採用版本 |
| 來源新增、修改、刪除 | 現有 refresh / validate → 重建或更新索引、檢查 checksum → 清除舊候選與失效 source 指向 → freshness / rebuild evidence |
| 索引不存在、過期或查詢逾時 | 顯式 unavailable / stale → 回退 source lookup 與 required reads → 不使用靜默舊結果支持結論 → 記錄 fallback |

不預先新增 `route.*`、validator 或 `runtime_projection.target_key`。Phase 0 若確認必要，先更新 plan 的 per-surface consumer 表，列 named consumer 與 wiring；任何新 contract 必須先定 owner/source 類型並驗證 compile 後的真實 consumer。

| Generated surface key | Named consumer(s) | Consumer 類型 |
| --- | --- | --- |
| 既有 `runtime.discovery.config`（不新增） | `discovery.go` | Bridge config lookup |
| 既有 `runtime-index.sqlite` 中 atoms / fts / sources | `runtime query`、後續共用檢索介面 | CLI / workflow lookup；Bridge 新接點為條件式範圍，若啟用需 Phase 3 驗證 |

## Open Questions

| ID | 問題與預設方向 | 狀態 / 解決期限 |
| --- | --- | --- |
| Q1 | 哪些 intelligence / 決策未進索引？正文覆蓋需要整篇還是 section chunk？先盤點 sources / atoms / fts 與 generator | resolved；見 evidence/phase0-inventory.md §來源覆蓋 |
| Q2 | 中文、混合術語、精確識別字與長句的查詢行為如何？先實測 tokenizer 與 query literal handling | resolved（初版）；整句 phrase / CJK / legacy 欄位限制見 Phase 1；不聲稱已評估一般語意搜尋 |
| Q3 | domain、project、supersession、證據版本哪些已有 canonical 欄位？避免另造 metadata；缺值預設 unknown | resolved；重用 domains / replaces / checksum，缺值 unknown；見 Phase 0 證據 |
| Q4 | 首批 consumer workflow 選哪個？先選有歷史漏找、可重跑的既有流程 | resolved；首批 sd-validation，觸發 / 重查 / 略過見 Phase 0 證據 |
| Q5 | 是否存在 Bridge 重複實作或可量測接入收益？若需調整，如何維持 Light budget？採 bounded lookup、dedup、timeout | resolved；本輪保留 Bridge，不改呼叫點；重開條件與 budget 見 Phase 0 證據 |
| Q6 | 向量與 reranking 是否值得？剩餘語意漏找才啟動；供應商、模型、持久化方案不預選 | deferred；Phase 4，無漏找證據則記錄不啟用 |

## Phase 0 — Pre-Build Interrogation / Architecture Compatibility Preflight

### Phase 0.0 — Open Questions 核對（公版，必填）

逐條核對本 plan §Open Questions，標記處置並回寫：

- [x] 已讀本 plan §Open Questions 全部條目
- [x] 對每條標記 `resolved`（附 Phase 0 證據）/ `still-open` / `deferred`（附原因）
- [x] `resolved` 的條目已同步勾選 / 附註於 §Open Questions
- [x] 若盤點新發現問題，已加入 §Open Questions

- [x] 記錄 trigger、checked sources、conflicts、interrogation、Q1–Q6 處置、execution decision 與 validation；未完成不得開始 implementation。
- [x] 盤點 generator、query、graph lookup、refresh / validation、Bridge hook、metadata 與 memory retrieval 邊界；列可重用能力、來源覆蓋和 duplication risk。
- [x] 核對 current DB schema / generated surfaces 與 canonical config 類型；不可把舊 YAML pointer 當唯一 source，也不可另建 runtime mirror。
- [x] 選定首批 workflow、實際查找 / 重查 / 略過條件、結果消費者與 baseline 命令；分開凍結 Phase 2 檢索及 Phase 3 採用驗收範圍與成本。

Candidate files（只代表調查與可能修改，不承諾全改）：`internal/app/runtime.go`、`discovery.go`、`hooks.go` 與對應 tests；`knowledge/runtime/sqlite/README.md`；`metadata/schema.md`；`governance/workflow-activation-engine.md`；Phase 0 選定 workflow 的 README / execution-flow；必要的 routing / summaries / graphs / scenarios。實作前讀各層 README、適用契約與 linked updates。Gate：Q1/Q3/Q4 已解，無未處理 source ownership 衝突。

## Phase 1 — 歷史 Lesson 檢索基準（縮限驗收）

- [x] 建立 20 題初版基準（18 題有歷史 lesson 來源的重播、2 題明確標註人工控制；非 20 個獨立原始提問），標 `case_id`、問題、query intent（如歷史追溯 / 當前採用）、task context、應找到的 source IDs / paths、相關程度與適用理由；無答案案例標空集合，不虛構真實案例。
- [x] 包含精確名稱、中文 / 英文 / 混合詞、不同措辭、workflow 已命中情境與無答案控制；過期 / 退休狀態 / 不適用 filter 以另列人工 fixture 探測，真實 supersession / 採用情境仍待 Phase 3。完整分布與缺口見 Phase 1 報告。
- [x] 分開報告 route 推薦與知識查找；本輪只量 workflow knowledge lookup，Bridge unchanged，因此 route discovery 標 `not evaluated`，不混用結果；每個漏找分類為未索引、查詢處理、排序、語意落差、filter、觸發 / consumer 缺失。
- [x] 凍結 corpus revision、查詢、工具版本、metadata、成本與結果；約 70% 診斷、30% holdout，相關改寫不得跨組，holdout 不用於調參。
- [x] 依下表分開定義檢索與採用指標；候選是否相關須依 query intent 判斷，歷史追溯可合理命中已取代版本，不得為降低不可採用候選比例而隱藏歷史。小樣本同報分子 / 分母，不宣稱統計顯著。

| 驗收層 | 指標 | 驗收階段 |
| --- | --- | --- |
| 檢索 | Recall@5、MRR@5、依 intent 定義的無答案誤推率、候選適用性分布、p95 latency、context tokens | Phase 2 |
| 採用 | 來源回讀、適用性判定、錯誤採用數、required-read / activation / blocker 保持、採用時版本核對 | Phase 3 |

未接線或未量測的指標一律標 `not evaluated`，不可填 0 或當成通過；分布至少區分適用、不適用、已取代與 unknown，分布本身不等於檢索錯誤率。

Gate（本輪縮限）：歷史重播來源、dataset SHA、expected labels、split 與評分協議可覆核；B0 在固定 corpus、顯式 binary / index 身分下重現 baseline。此範圍已核對通過。檢索增益與副作用依 Phase 2 驗收，採用與實際 tokens 依 Phase 3 驗收；剩餘責任不能由縮限 gate 代替。沒有增益時只交付診斷，不聲稱搜尋強化完成。

Phase 1 審核補充（2026-10-24）：[phase1-audit.md](evidence/phase1-audit.md) 分開核對 case hit 與 source Recall，13 題目前只支持 `query_processing_or_indexed_fields` 粗分類，未逐題證實因果。現有預留組保守作固定 comparison / regression set，不主張嚴格獨立 holdout；獨立泛化驗證須另收資料。

- [x] 依使用者 2026-10-24 審核處置縮限 Phase 1：只驗收 18 題歷史 lesson 重播 + 2 題人工控制的來源、資料一致性、評分與固定比較基礎；使用者已核對資料與 harness，本輪以顯式 binary / index 重播 B0 再確認分數。這不表示原完整範圍或 Phase 2 品質已通過。

本輪不驗收全庫 intelligence recall、獨立 holdout、一般自然語意誤推率、真實 supersession 採用、實際模型 token 上限或採用安全。這些責任未取消：

| 剩餘項目 | 階段與交付 | 責任 |
| --- | --- | --- |
| Query / extraction 增益與擴大查詢副作用 | Phase 2：產物綁定的四組對照、另版 overlap controls、準備輸出的新增來源逐筆審閱；未選候選保留 coverage／trace，不要求 unknown 清零 | 本計畫 query / extraction 執行者整理，larrylin 或指定 reviewer 覆核 |
| Intelligence 本文覆蓋與 relevance | Phase 2：另版 source-level intelligence 比較資料；未量測不得由 lesson 結果宣稱 intelligence 改善 | 本計畫索引執行者與首批 workflow owner |
| 自然語意無答案、真實 supersession、來源變更與採用安全 | Phase 3：有範圍與標註的 consumer replay，source-read / applicability / required-read / blocker 證據；人工控制不替代真實案例 | 首批 sd-validation maintainer 與 consumer 任務 owner |
| 實際模型 token budget | Phase 3：接入 consumer 前固定 tokenizer / version、計入的 context 範圍與 ≤ 2000 實際 tokens；未量測維持 not evaluated | Consumer / context 組裝執行者；Phase 2 僅驗 ≤ 2000 runes 工程代理值 |
| 獨立泛化驗證 | Phase 2 若要宣稱泛化，或 Phase 4 語意實驗前，另收未參與診斷 / 方法選擇的資料；現有組只作 comparison / regression | 該階段實驗 owner 整理，larrylin 或指定 reviewer 核對資料使用歷程 |

縮限 Phase 1 gate 通過只允許進入有控制的 Phase 2 實驗，不代表搜尋功能已改善或整個計畫可完成。

## Phase 2 — 補既有全文搜尋、metadata 與關聯查找

最新範圍以 [final-output 協議](evidence/phase2-final-output-protocol.md) 為準：只要求輸出來源逐筆審閱；未選／unknown 保留處置與追查，取消全池清零 gate。 已完成 [26 題知情診斷](evidence/phase2-final-output.md)：23/23 指定來源、selected confuser 0/6（retrieved 仍 6/6），凍結模型 retrieval payload 最大 1322 runes，完整摘要無丟棄。資訊充分性為 17 個正例足供下一步、6 partial；O06 未解。這是 current-agent exposed diagnostic，不是受控 runner／泛化／採用 gate 通過。

歷史盤點依 [readiness 契約](evidence/phase2-readiness-contract.md)：unknown 對 consumer 可見並區分 unassessed，空 preview 不表示無知識；優先 P0 完整摘要。依最新範圍完成全資料最終輸出審閱／資訊充分性、成本時間與索引生命週期 gates 前，不安排 Phase 3 接線。[全資料盤點](evidence/phase2-readiness.md) 顯示 26 題 770 source occurrences 中 738 unknown；P0 array 最大 1922，但本版可見 envelope R08 為 2208 runes，完整 context 成本未通過，不能由六題結果推定完成。

本輪按 [disposition 協議](evidence/phase2-v3-disposition-protocol.md) 重用原六份 raw，只新增候選處置；direct 待讀回，background 預設不填主要上下文，not-applicable 本次不選，缺標記／unknown 不當成已確認。分開 retrieved、待讀回、selected preview、verified/adopted，不改前輪分數。 結果見 [disposition 重播](evidence/phase2-v3-disposition.md)：同六題 preview 5/5、0/4、≤1922 runes，retrieved 仍 4/4；9 份 direct 來源身份核對一致，verified/adopted 未驗收，Phase 2 維持未完成。

v3 先按 [固定候選協議](evidence/phase2-v3-protocol.md) 分離 pool、選擇與 packer，再準備 [6 題語意排序 pilot 成本契約](evidence/phase2-v3-semantic-contract.md)。已完成 26 題固定池與 156 組離線比較，見 [v3 preflight](evidence/phase2-v3-preflight.md)：S0 + full 多筆為 14/14、≤1982 runes，但 confuser 仍 6/6。焦點／否定／排序失敗已有證據，可測 reranking，不以向量搜尋或 Phase 2 通過為前提；只作曝光資料診斷。六份語意 requests 已備；使用者已選現有 agent／subagent runner，按 [agent runner 協議](evidence/phase2-v3-agent-runner.md) 完成六個 fresh context 診斷，見 [agent pilot](evidence/phase2-v3-agent-pilot.md)：5/5 指定正向、2/4 confuser，ranking-only 品質仍失敗，時間監測未證明受控；不可得 usage 明列。額外 API 非必要且本輪不選；模型／usage 不可得欄位明列，不以美元方案阻擋此 runner。

v1 結果：B0 / Q / E / QE 診斷命中 1 / 13 / 3 / 14（各 /14）；Q / QE overlap confuser 均 6/6、preview 超額，故候選未採用且 Phase 2 gate 未通過。逐題產物與審閱待辦見 [phase2-v1-ablation.md](evidence/phase2-v1-ablation.md)。v2 已按 [phase2-v2-protocol.md](evidence/phase2-v2-protocol.md) 完成最終輸出重評：QE diagnostic 12/14、overlap 正向 3/5 / confuser 3/6，成本通過但品質仍失敗，見 [phase2-v2-ablation.md](evidence/phase2-v2-ablation.md)。下一輪處理問題焦點與已排除前提，保留排序退步及 shortlist 損失。v1 已拒收，不再為結清而要求完整審完其 191 rows；準備採用的版本仍須完成審閱。保留 v1 證據與原門檻，不導入 vector。

- [x] 依 [執行控制契約](evidence/phase2-execution-contract.md) 先跑固定 corpus 的 B0 / query-only / extraction-only / combined 對照：顯式 query binary 路徑 / SHA、index manifest、generator revision / source SHA、index inputs / corpus SHA；四組配對 validator 必須通過，逐題記錄增益與成本；新增正文另組實驗，不混入 extraction-only。對照組與資料獨立性規則見 Phase 1 審核包。
- [ ] 以另版 [overlap-controls-v1.json](evidence/overlap-controls-v1.json) 約束查詢放寬：selected-output 已標 confuser 命中 0，retrieved 原值另報；準備輸出來源逐筆覆核，未選候選保留數量與追查入口。依使用者最新審閱明確取代原全池覆核要求：unknown 可為合理終態，須結清處置而非逐筆清零，不以 8/14 單獨證明有用。
- [ ] 依 Phase 1 原因補索引覆蓋與查詢處理，保留精確識別字搜尋；正文擴充只涵蓋選定 workflow、intelligence、決策與整理後知識，不預設索引所有 raw evidence。
- [ ] 若需 chunk，保留 stable source identity、section anchor、source revision / checksum；重複命中按來源合併，避免一篇文件占滿 top-k。
- [ ] 重用 canonical domain / scope / status / supersession 欄位；硬限制只用明確適用條件，缺欄位顯示 unknown，避免過濾造成靜默漏找。
- [ ] 重用既有 graph lookup，有界展開來源、依賴與相關決策；記錄 edge 類型與深度，不推定 relation 等於支持或授權。
- [ ] 回傳來源、命中理由、版本 / freshness、適用性待查事項與回讀入口；相關性分數和可信 / 適用性欄位分開。
- [ ] 驗證建立、refresh、修改、刪除、搬移、stale、缺 metadata、缺索引、rebuild 與 deterministic ordering；用 Phase 1 固定資料比較。

Gate：通過縮限後凍結的檢索層 acceptance、產物矩陣核對與候選審閱 gate；Phase 2 成本上限明確是 ≤ 2000 runes 代理值，不宣稱實際 tokens 已驗收。採用層保持 `not evaluated`，留待 Phase 3；cache 可從 canonical source 重建；不得修改 canonical 結論來讓 benchmark 通過。

## Phase 3 — 接入 Workflow 採用流程與條件式 Bridge 調整

- [ ] 優先重用既有檢索能力。只有 Phase 0 確認重複實作，或比較實驗證明可量測收益，才抽共用介面並調整 Bridge 呼叫接點；記錄啟用 / 不調整的證據。Bridge 保留 miss-path 與 advisory 邊界，workflow lookup 不依賴 detector miss，也不以 Bridge 重構為交付前提。
- [ ] 在首批 workflow 寫明「事件 → 查找動作 → 證據」：初次查找事件（如形成方案前、失敗診斷時）、重查條件（如關鍵前提、來源或任務範圍改變）、可略過條件及可覆核理由；同時定義 context filters、讀回與採用檢查。用觸發、重查、合理略過與不當略過案例驗證，不新增通用 trigger framework。
- [ ] 採用前檢查 domain / project / 前提、證據版本與失效判定、supersession、候選 / 已驗證 / 契約身分；checksum 變動只表示需重驗，不單獨判斷結論失效。
- [ ] 採用時核對實際回讀來源的 revision / checksum；與搜尋候選版本不一致時，重新驗證適用性並記錄實際採用版本。驗證期間再次變動則重驗或停止採用，不能沿用舊判定。
- [ ] 記錄 query、查找 / 重查 / 略過理由、候選版本、回讀與實際採用版本、採用 / 拒絕 / unknown 理由與本次 workflow 用途；只存必要去敏資料，設定 retention 與 token cap，不建立全量對話庫。
- [ ] 建立正負 scenarios：miss advisory、route hit 後找經驗、低排名 required_read 仍載入、相似但不適用、已取代的歷史追溯與當前採用差異、查詢版本 A 到採用版本 B 的重驗、查找 / 重查 / 略過、stale / unavailable fallback、無答案與 timeout。
- [ ] 執行真實 consumer dogfood；報告找到 → 讀回 → 適用 → 採用各階段結果；關閉 feature 後能回到既有 lookup，既有 activation gate 行為一致。

Gate：通過採用層 acceptance；未實際評估項目不得填 0。至少兩個獨立真實任務的完整採用 / 拒絕鏈，加上負面回放；只測 helper 或 mock 不算 runtime integration。完成 compile / refresh / validate（按受影響來源）及 consumer 行為證據。

## Phase 4 — 條件式語意搜尋 / Reranking 實驗與收尾

- [ ] 依缺口分開決定 candidate generation 與 reranking：候選漏找才評估 vector；已入池但焦點／否定／排序失敗，可獨立測語意重排序，不要求先加向量。Phase 2 v3 只先做小範圍診斷 pilot；本階段泛化／採用仍需下列獨立資料與成本 gate，不能用 pilot 代替。
- [ ] Phase 4 啟動時先確認仍有獨立留出資料；若原 holdout 已參與 Phase 2 後的診斷、方法選擇或調參，須另建未參與這些活動的留出集，相關改寫不得跨組。記錄各集使用歷史，不將已看過的測試資料重新宣稱為獨立 holdout。
- [ ] 實驗開始前凍結模型 / 版本、資料邊界、reindex 成本、獨立 holdout、最小有效改善題數與 latency / tokens budget；無此凍結不啟動。
- [ ] 新增方法須改善預先標記語意漏找，holdout 不退步、安全錯誤為 0、成本在 budget 內才採用；不達標則保留 BM25 並記錄拒絕原因。holdout 看過後再調參需新留出集。
- [ ] 若無剩餘語意缺口，記錄 skip 的證據；可完成核心計畫，不因未導入 vector 阻擋。
- [ ] 彙整結果、已知限制與 deferred owner / restart 條件；同步 README / 契約 / generated consumers，執行 Plan Completion Closure。

## Stakeholder 同意項目

使用者已要求撰寫計畫並提供五點審閱補強，方向為修訂後進入 Phase 0；已依後續分段執行要求完成 Phase 0 盤點；已建立 Phase 1 基準版；production 搜尋 implementation 尚未開始。執行啟動時需確認首批 workflow、基準資料與成本 budget；涉及外部模型傳送資料或付費服務時另行確認具體方案。Phase 4 是條件式選項，不是預先採購或導入承諾。

## 完成條件

- [ ] Phase 0–3 gates 全通過，Phase 4 為有證據的採用 / 拒絕 / skip 決策。
- [ ] 真實問題 baseline、保留集比較與消費者採用鏈可重跑；未達標項不得以文件完成替代。
- [ ] required_read、activation、source authority、freshness 與適用性負面案例通過。
- [ ] 每個新增 surface 都有 named consumer 與 runtime evidence，無 orphan config / index。
- [ ] 驗證連結、linked updates、diff 與適用 tests / validators；完成歸檔、狀態索引、commit / push / readback。

## 與其他 plans 的關係

Bridge 呼叫接點的變更為條件式範圍；不調整時記錄理由並保留現有行為，不阻擋 workflow lookup 交付。延續 [Discovery Bridge](../2026-06-06-1700-workflow-activation-discovery-bridge.md) 的 advisory 能力，不把它未完成的 Deep Discovery / empirical work 默認為完成；Phase 0 需劃清重疊 work item 的 owner 與證據歸屬。與 [Feedback Execution Closure](../2026-07-10-0955-final-form-feedback-execution-closure.md) 互補：該計畫處理經驗沉澱，本計畫處理沉澱後的查找與使用，不接管其 roadmap。

後續多次量測的完整記錄依 [plan evidence 規則](../../../governance/lifecycle/plan-evidence.md) 放本資料夾 `evidence/` 並建 README Run 索引；Phase 0 盤點已存 [evidence/phase0-inventory.md](evidence/phase0-inventory.md)，完整索引見 [evidence/README.md](evidence/README.md)；Phase 1 baseline 見 [evidence/phase1-baseline.md](evidence/phase1-baseline.md)；本輪僅量測既有能力，不代表 Phase 2 搜尋強化或 Phase 3 採用層已通過。
