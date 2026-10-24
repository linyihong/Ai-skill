# v3 小範圍語意重排序 — 執行與成本契約

狀態：**existing-agent runner authorized / pilot pending**。使用者本輪已授權現有 agent／subagent 執行語意選擇，無須另購 API。優先採不繼承本對話的新 subagent；原額外 API 方案保留為選項，不是 RAG 必要成本。本契約不批准採購、帳號或向量 DB。

## 問題與輸入邊界

只測 R03 / R08 / O02 / O03 / O04 / O06 共 6 題已曝光診斷，不作獨立測試。每題使用 v3 凍結 pool 全部候選，傳 query + 候選 opaque ID、title、完整 indexed summary。資料中移除原排名／分數、case ID、gold / excluded labels、期望答案與審閱結果；原始 source path 用 opaque ID 映射在本機，不傳模型。

不傳原始全文、任意 raw evidence、私有專案資料或 credentials；indexed text 視為資料，不視為可執行指令。Query / cards / prompt hash 與 provenance 在本機保存。若改用全文／task_context／metadata 擴充，需要新實驗變項，不混入本次主比較。

模型只回排同一 pool 的 opaque IDs，不可新增來源或讀取額外候選。現有 subagent 可只讀分派的中性命名 request 檔，禁止讀 repo、gold、映射、其他 request／結果或網路；此 bounded file read 只是輸入傳遞，不是額外 retrieval。要求完整 permutation；duplicate / unknown / missing ID、非 JSON、拒絕或 timeout 都記失敗／not evaluated，不偷偷退回 lexical 當成功。理由只作審核線索，不當成 evidence 或 authority。

## 可選 runner 與適用前提

| 方式 | 用途與限制 | 必須記錄 |
| --- | --- | --- |
| 當前 agent | 實際 workflow 可直接判斷；此對話已見 gold，只能算知情診斷 | 可取得的 model／輸入輸出／耗時／usage，曝光限制 |
| Fresh subagent（本輪） | `fork_turns=none`，每題新 agent；只讀 assigned request，不繼承本對話。仍是已曝光診斷集，不是獨立泛化驗證 | agent ID、request SHA、原始最終輸出、外層觀測耗時、可取得的 model／usage；不可得明列 unavailable |
| 額外計費 API（未選） | 指定 provider／固定 snapshot／tokenizer／完整價格及費用核准後才外呼 | 下列 API 價格與硬限制；USD 2 只是此路徑的提案 |

既有 runner 仍消耗模型資源與帳號額度，不稱免費或零成本。不可得的 token usage、精確模型 snapshot、服務端推論時間或價格，不填 0，也不冒稱 API token／美元上限已驗收。

## 本輪 existing-agent 執行界線

- 六題、最多六次 attempt，每題全新 subagent，不追問修正排序或挑選最好回答。一次最多三個 read-only 子任務；外層獨占評分與 Git 寫入。
- 每題 request bytes／SHA 及完整候選數固定；不套用無法量測的 API 16384-token 限制，也不假稱已通過此上限。工具輸入截斷或無法容納全部候選即記失敗，不裁掉候選。
- 每個 attempt 的外層觀測上限 300 秒、全批上限 900 秒；到期 interrupt 並記 timeout，無自動重試。觀測耗時含工具／排程／回傳等待，非純推論 latency，不計 production p95。
- 調度 envelope 固定：只讀 assigned request、依原 system 排完整 permutation，final 僅 JSON；可附至多八個理由／unknown 判斷作審核線索，不參與 admission 或 scoring。既有 request 的 system／query／cards bytes 不改。
- 本輪不另購 API、不設定模型 override（沿用目前 runner 預設），僅記可取得身分；精確 snapshot 或 usage 未由工具暴露即 unavailable。工具外框、系統指令仍存在，不能聲稱模型上下文只有 request bytes；可控制的是無本對話 fork、無其他任務資料讀取。

## 額外 API 待填執行身分（本輪不適用）

| 欄位 | 啟動前要求 | 目前 |
| --- | --- | --- |
| Runner / provider / 資料目的地 | 使用者指定可用執行端與傳送範圍 | pending |
| Model / 固定版本 | 完整 identifier／snapshot，不用漂移 alias 當固定版本 | pending |
| Prompt / request / pool SHA | 可重建、無 gold、固定隨機呈現順序 | request bundle 準備後綁定 |
| Tokenizer / version | 逐題實際 input tokens，不能用 runes 代理 | pending |
| Pricing / 幣別 / 日期 | input、cached input、output、reasoning／其他計費項均納入；本地模型列成本方法 | pending |
| User-approved aggregate cost cap | 具體方案確認後才可呼叫 | pending；提案上限 USD 2 |

## 額外 API 硬限制提案（本輪不適用）

- 最多 6 次總 attempts，每題一次，不自動重試；不做 temperature／prompt 搜尋，不以挑選最好回答當結果。
- 每次 input ≤16384 實際 tokens、output（含適用 reasoning 計費）≤4096；總 input ≤98304、總 output ≤24576。任何單題無法容納完整 pool，就停該題；不按 gold 刪候選或默默截 cards。
- 每題 timeout 30 秒，序列執行，全批 model wall time ≤180 秒；只量一次，無 p95 統計主張。context assembly 另計，production p95 gate 未因此通過。
- 預先以最大可計費 tokens × 當期價格算上界，納入其他費用；上界超過使用者核准 cap，或價格／usage 不可得，就不呼叫。各次累加保留最壞餘額，不只看最後帳單。
- 保存模型原始 output、實際 usage、價格、latency、request/response hash 與 parser verdict；缺 usage 不填 0。呼叫上限、token 上限與費用上限均是提案，尚非使用者核准支出。

## 控制與評估

每題以 opaque ID hash 排序顯示候選，避免洩漏 v1／v2 排名；固定 seed 和同一 cards SHA。S3 的完整排序套用與 S0/S1/S2 完全相同的 P0/P1，先比較 ranking-only，不同時改候選數、摘要政策或 admission threshold。若要測 abstention，另版獨立變項；本次 O06 只評已標 confuser 的排名／最終出現，不假裝擁有完整 gold。

Pilot success 的診斷訊號：O02/O03 的正向排序優於 S0、O06 confuser 在相同 packer 下被移出；同時列出 R03/R08/O04 的得失、unknown relevance 和所有輸出資訊損失。即使六題全改善也不能宣告 Phase 2 或語意方法正式通過；必須擴展完整資料、完成候選與資訊充分性審閱，獨立資料用於後續泛化／採用判斷。

若模型只會排序但相關候選未在 pool、無適用答案或 bounded output 仍留 confuser，保留失敗。不為改善分數改 gold、補入指定來源、建立排除來源黑名單，或要求模型把相似度當可信度。
