# Gold-blind existing runner — 固定 26 題選擇與補讀交接

結論：**不知道 gold 的 fresh runner 在這批曝光診斷資料上達到選擇與輸出預算目標**。全 26 題一次 attempt、未使用人工 decisions；23/23 指定來源保留、selected confuser 0/6、完整 retrieval message 最大 1558 runes。正文補讀仍有兩題需要來源專案證據；不宣稱所有問題已解、可泛化、正式自動 selector 已交付或 Phase 2 完成。

## 固定方法與資訊邊界

[協議](phase2-blind-selector-protocol.md)、[prompt](phase2-blind-selector-prompt.txt)、[request manifest](phase2-blind-selector-manifest.json) 與 test-only Go request builder 在 `e758a3e3` 提交，先於所有 attempts。Parser／scorer 在 `742cee39` 提交，官方 score 在其後執行。準備時 Python builder 曾被 Go-first hook 阻擋，已在 attempt 前移除並改成 Go；沒有忽略該 gate。

26 個 `fork_turns=none` 新 agents，每題僅接 assigned neutral request：固定 system、原 query、固定 broad pool 去重 source 的 opaque ID／title／完整 summary。沒有 case ID、task_context（其中部分含審核提示）、source mapping、BM25、gold／confuser、人工 decisions 或既有 raw／成績。每個 request 共用同一 system，候選按 opaque ID 排序；不要求全池分類或完整 permutation。Child 只獲准讀 assigned request、寫 assigned temp raw 並回傳 JSON。

這是以指令限制檔案／網路存取，不是 OS sandbox；工具與系統外框仍存在，child 完整工具 audit 不可得。不能宣稱模型完整上下文只有 request bytes。父層與資料集／prompt 設計已見 gold，因此仍是 **gold-blind execution on exposed diagnostic data**，不是獨立 holdout／一般泛化證據。

[run](phase2-blind-selector-run.json) 保存 agent 名稱、request SHA、dispatch／raw observed timestamps、raw SHA／bytes、一次 attempt、模型／usage 不可得欄位。[原始回應](phase2-blind-R01.raw.txt) 至 `phase2-blind-O06.raw.txt` 共 26 份逐 byte 保存，未修復 JSON、改 disposition 或挑最好回答。沒有額外 API 採購／呼叫、沒有 model override；既有 runner 資源用量不可得，不當作零成本。

## 觀測與執行控制

最多三個 concurrent children，各 wave 確認前批 final 回應後才派下一批。另有持續 1 秒檔案觀察器，[完整 log](phase2-blind-selector-monitor.jsonl) 綁 SHA；記錄第一個可解析 raw 檔可見時間及 heartbeat。原始模型回應存檔是完成訊號之一，另以收到前批 final 後才派下一批的時間保存保守 final-observation 上界，沒有把檔案時間戳當推論完成時間。

| 觀測 | 值 | 限制 |
| --- | --- | --- |
| Attempts／parser 有效 | 26/26，各一次 | 0 retry、0 repair |
| 最大監測間隔 | 1.018334 秒 | 門檻 ≤30 秒，沒有續跑觀測缺口 |
| 最大 dispatch→raw 可見 | 30.262501 秒 | 含工具／排程／檔案寫入，不是 inference latency |
| 最大保守 final observed 上界 | 97.144018 秒 | 含父層工作／通知讀取間隔；每題限制 300 秒 |
| 整批觀察 | 554.426502 秒 | 含 wave 間父層工作；本輪 26 題預定上限 3000 秒 |
| Timeout／deadline events | 0 | 實際 interrupt 超時分支未被觸發，不宣稱其對抗性驗證完成 |

這一輪可確認觀測連續且所有回應在期限內；不修復舊六題 6394 秒的時間缺口，不計 production p95。精確 snapshot、實際 input/output tokens、推論 duration、帳號用量仍 unavailable。可重用的是凍結 request／schema／規則及捕捉輸出評分方法；正式 runner orchestration／中斷負例與 workflow 接點仍未交付。

## 最終 retrieval message

[results](phase2-blind-selector-results.json) 綁定 request、raw、pool、run、dataset 及 scorer SHA；`AI_SKILL_BLIND_REPORT=<new output> go test ./internal/app -run '^TestRetrievalBlindSelectorScore$' -count=1` 只評 captured raw，不叫模型。

模型 message 沿用 guard、state、pool、unassessed、unknown、deferred、trace 與 source／kind／status／完整 summary，另將 runner 的 **need_to_read 放在每個 result**，一起計入成本。來源版本／checksum、原理由及評分留外層 audit；trace 對應本次 observation 和 readback queue。O06 的具體 unknown_action 保存在 trace audit，空 message 仍可見 unknown／state；production trace resolver 尚未接線，不冒稱下游已自動讀到外層全部內容。

26/26 完整 message ≤2000 runes，最大 **1558（R08）**，0 summary truncation、0 budget deferred、沒有 P1 fallback。只是 retrieval message 的工程代理預算；selector 讀取全部 candidates 的 input、system／工具外框與整段任務 context／實際 tokens 不在這個 2000-rune gate 中，不能據此說 end-to-end 模型成本已通過。

## 選擇結果

| 指標 | 結果 | 判讀 |
| --- | --- | --- |
| 歷史正例 source hit | 18/18 | 原 diagnostic／comparison labels 不改 |
| Overlap 正向 source hit | 5/5 | O06 無完整 gold，不放正例分母 |
| 正例 MRR | 22/23 = 0.956522 | R08／R09 指定 lesson 均第 2，其餘正例第 1 |
| Retrieved overlap confuser | 6/6 | 原 fixed pool 不改 |
| Selected-output confuser | 0/6 | 僅預標來源，非一般 false-positive rate |
| 人工空池控制誤推 | 0/2 | nonce／不存在 status，不代替自然 no-answer |
| O06 | 空選，1 explicit unknown | 不刷新 discovery 當 auth 解法，也不猜 App bootstrap |

不同於人工參考 MRR=1，本輪 runner 在 R08／R09 保留有用 summary／canonical 入口在歷史 lesson 前；沒有按 gold 重排。這支持固定規則在本批資料上可執行，不證明跨任務／模型穩定性或 production 相關性已通過。

770 source occurrences 中 33 筆明列 assessments：direct 26、background 1、not-applicable 5、explicit unknown 1；737 未標記，unknown 合計 738。這是 coverage／處置現況，不是待清零缺陷。未選仍可由完整 pool／request mapping 追查。

## 來源讀回交接

[handoff audit](phase2-blind-selector-handoff.json) 在 raw／score 固定後檢查，未回寫或重新篩選原 payload。全部 26 筆準備輸出來源可沿 mapping／trace 解析，對應 21 份不同來源；scorer 實際讀取完整固定 Git blob、核對 checksum 與 current source，全部 same-version。正文 scope 審閱使用已讀過且 checksum 一致的 source evidence，只在外層驗交接，不傳給 selector。

23 個有輸出的正例中，21 題的 need_to_read 可直接對應正文中的下一個查證方法；R11／O03 的指向仍有額外前提：正文提供 fingerprint 比對欄位、隔離方式及來源專案入口，**不含具體平台版本／堆疊／指紋值**。Runner 說要讀這些內容，不能假定都在泛用正文；後續要向 owner 取得該專案 case 與本次環境資料。這兩題是可用入口加額外取證，不當完整 handoff 閉環通過。

| 已知 partial | Runner 實際補讀目的 | 正文／後續處置 |
| --- | --- | --- |
| R06 | blob URI 原因、original base URL、segment rewrite | 正文有 base 改變與 resolve／helper test，交接有用 |
| R08 | owner／projection 理由、歷史失敗背景、現行契約 | summary 導航＋lesson prose 跳步＋canonical YAML 可接續核對，不由規範反推全部歷史 |
| R11／O03 | fingerprint／版本範圍、最小隔離 | 方法在正文，具體值留來源專案；需追加任務證據 |
| R12 | 目標 statement 鎖、deferred COMMIT 與 timeout 命中確認 | 正文支持；runner 未逐一點名 readback／同 key retry，完整 safety 驗收仍須 workflow 展開 |
| R16 | gate 解析基準、合法相對路徑及驗法 | 正文含 index-relative parsing 與移除連結負例，交接有用 |

R19／R20 正確保留空池。O06 明確排除 unreachable-peer recovery，App session 候選因平台／服務／撤銷流程前提不明標 unknown、本次不輸出；要求先確認 provider 與官方重新授權入口。本輪處置完成，重新授權問題仍未解。沒有要求將 partial 摘要擴寫成完整答案，也沒有把外層全文當作模型已收到。

以上只是離線選擇→trace/source readback 的交接證據。真實 consumer 是否叫用、如何確認 applicability／supersession／required_read／blocker、版本 race 與採用仍屬 Phase 3，尚未接線。

## 驗證與剩餘工作

- `go test ./...`、request exporter、strict parser controls 及正式 26 題 scorer 通過。Non-direct／foreign／duplicate／missing selection、缺必要補讀欄位會拒絕；失敗不 fallback 成成功。
- 獨立 Python 核對全部 request 只有 system／query／cards、完整 card bytes 解碼值、mapping／membership、原 metadata、全部 hashes，以及 source hit／MRR／confuser；獨立 Go 重算 26 份整個 message runes。
- 前輪 [CI 40057004671](https://github.com/linyihong/Ai-skill/actions/runs/40057004671) 已成功；本輪 [CI 40058203066](https://github.com/linyihong/Ai-skill/actions/runs/40058203066) 撰寫時仍執行中。

Phase 2 維持未完成。下一步按既定計畫補 intelligence source-level 比較、canonical metadata／bounded graph、索引建立／refresh／修改刪移／stale／rebuild／ordering，以及正式 runner 中斷／usage 可取得性邊界；不新增 reviewer 層、不把 unknown 清零當 gate。實際 workflow 交接與採用由 Phase 3 驗收。
