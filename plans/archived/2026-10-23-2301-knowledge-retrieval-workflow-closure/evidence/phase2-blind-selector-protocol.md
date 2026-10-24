# Gold-blind existing runner — 26 題固定選擇與讀回交接

在新 attempt 前固定。使用者要求既有 runner 不知道 gold 時按固定規則產出選擇與 unknown。本輪沿用 fresh subagent runner，不使用人工 decisions 作 selector 輸入；不另購 API，不指定 model override。資料集與 prompt 設計已曝光，因此是 **gold-blind execution on exposed diagnostic data**，不是獨立泛化驗證。

## 固定輸入與輸出

26 題各一個 fork_turns=none agent，僅讀 assigned neutral request。request 只有共用 system、原 query、固定池 source 去重後的 opaque ID／title／完整 indexed summary。排除 task_context（部分帶審核提示）、case ID、source mapping、原 rank／score、gold／confuser、人工 decisions／sufficiency、先前 raw／成績；候選呈現按 opaque ID 排序，不按預期來源。池空也執行，記 abstention，非自然 no-answer。

每題一次、無 retry／修復／挑最好答案。JSON schema 固定 selected（ordered opaque IDs）、assessments（最多八筆：id、relation、reason、need_to_read）、unknown_action。direct 才能 selected，每個 direct 必須列於 selected 且有具體補讀目的；background 不填主要上下文，not-applicable 本次不用，unknown 保留前提與處置。未標記是 unassessed，不要求全池分類。不知道 task domain／前提時不可假定跨域資料適用；可空選。摘要不需回答全部問題，但補讀應說明查什麼而不是捏造全文內容。

## Runner 控制與原始輸出

最多三個同時執行、每題新 agent；child 不讀 repo／其他檔／網路、不操作 Git，只可讀 assigned request 並寫 assigned temp raw response；final 亦回傳同一 JSON。這個例外輸出位置只是保存原始模型產物。外層記 agent ID、request SHA、dispatch／首次觀察完成時間、raw SHA／bytes、parser verdict；精確模型 snapshot、token usage、child 工具 audit 不可得即 unavailable，不冒稱 OS sandbox 隔離。

每題外層期限 300 秒、全批期限 3000 秒，連續監測間隔 ≤30 秒；到期 interrupt，不重試，已產出的無效 raw 仍保存。若工作中斷造成 >30 秒觀測缺口，本輪 timing-control gate 記未證明；不要用後續補計時修復。包含工具／排程等待的外層 elapsed 不是推論 latency 或 production p95。用帶持續時間的 clock 觀測與 UTC timestamps 留紀錄；timeout／parser failure 不 fallback 成成功。

## 模型 retrieval message 與交接驗收

沿用 final-output 的整物件 ≤2000 runes、P0 完整摘要、guard／state／pool／unassessed／unknown／deferred／trace；每個 result 在 source／kind／status／summary 外新增 **need_to_read**，這個補讀目的是 consumer 必要資訊，一起計成本。維持 selector 順序、超額從尾端 deferred、不截摘要或偷切 P1。trace 綁 raw／request／pool／版本，audit、評分、hash 與未選理由留工具層。實際 tokens／整段選擇模型輸入成本仍不可得，不能稱已驗收。

外層只在原始輸出凍結後評分：分開 retrieved confuser、selected confuser、23 正例 hit／MRR、2 人工空集合、O06 unknown 處置、budget drops。每個準備輸出來源都需讀回固定 corpus 並核對當前 checksum，審閱與 query 的關係、need_to_read 是否能指向正文可查的缺口；不改 raw 或先替模型修判斷再評分。

handoff 有效＝所選來源可解析／版本可核對，補讀目的對應 query 且正文能支持下一個查證動作；不要求摘要升成完整答案。對已知六個 partial 分開評估補讀目的是否覆蓋所缺內容，不將 outer 讀過全文算成下游已收到。若額外任務前提缺失，明列 unknown。這只是離線 handoff evidence，真實 workflow 採用、required_read／blocker／supersession／版本 race 仍 Phase 3。

成功要求：所有 attempt 可追查、parser 有效；23 正例來源不漏、selected 預標 confuser 0/6、2 人工控制不誤推、O06 不把未知 auth 線索當已確認解法；整個 retrieval payload 在預算、所選來源審閱與 handoff 無未處置缺口。未通過即保留失敗，不再增加一層審核或強制 unknown 清零。Phase 2 的 intelligence、索引 lifecycle 與其他既有 gates 保持。
