# v3 Existing-agent Pilot — 六份 fresh-context 語意排序

**六個 attempt 已完成，無額外計費 API；語意排序有診斷增益，但品質及時間 gate 未通過，Phase 2 維持未完成。** 本輪不修改 production 搜尋，不新增 vector DB，也不以既有 runner 的資源消耗冒稱免費。

## Runner 與資料隔離

成本契約於 `157fd14e`、執行前 [agent runner 協議](phase2-v3-agent-runner.md) 固定後派發。每題一個新的 `collaboration.spawn_agent`、`fork_turns=none`，不指定 model override；兩批至多三個 read-only agents，outer agent 獨占評分與 Git。使用者已授權這個方案，不再把 API 價格／USD 2 提案當成既有 agent runner 的必要前提。

六份 request bytes／SHA、pool 和 opaque-to-source 映射與 `858da1b7` 完全一致。Ranker 只被分派中性 SHA 檔名的單題 request；dispatch envelope 禁止讀 repo、來源全文、gold、映射、其他結果或網路。不繼承此對話，不把 case ID 或 lexical rank／score 送給 ranker。工具／系統指令仍存在；parent 無法取得完整 child read trace，故實際讀取邊界未獨立稽核，不宣稱完全無先驗或獨立 holdout。

原始回應保存為 parent 工具可見 final text 加一個儲存換行，沒有改排名或 assessments。[Run manifest](phase2-v3-agent-run.json) 記 agent ID、request／dispatch／raw SHA、時間與缺失欄位；模型 snapshot、token usage、價格、純推論時間和完整 child tool trace 均 unavailable。沒有用模型自述猜測版本。

| 案例 | Agent | 候選 sources | 原始 final text | Permutation |
| --- | --- | --- | --- | --- |
| R03 | /root/rank_a | 5 | [raw](phase2-v3-agent-R03.raw.txt) | valid 5/5 |
| R08 | /root/rank_b | 55 | [raw](phase2-v3-agent-R08.raw.txt) | valid 55/55 |
| O02 | /root/rank_c | 64 | [raw](phase2-v3-agent-O02.raw.txt) | valid 64/64 |
| O03 | /root/rank_d | 49 | [raw](phase2-v3-agent-O03.raw.txt) | valid 49/49 |
| O04 | /root/rank_e | 57 | [raw](phase2-v3-agent-O04.raw.txt) | valid 57/57 |
| O06 | /root/rank_f | 23 | [raw](phase2-v3-agent-O06.raw.txt) | valid 23/23 |

六份皆為 JSON、完整 permutation，無 duplicate／unknown／missing ID，沒有重試、追問修正或挑最好結果。總計 253 個 source occurrences 被排序；不是 253 個獨立來源或測試題。

## 外層評分：同一 pool 與 P0／P1

Test-only scorer `bcb73045` 的 `retrieval_semantic_agent_test.go` 先驗 raw／request／pool hashes 與 permutation，再使用既有 `retrievalV3Pack`。Assessments 只作線索，不作 admission filter；不能看到 background 後才臨時刪除來源。

[完整結果](phase2-v3-agent-results.json) 保存 12 份最終 payload、完整 source order、正向／confuser rank、hash 與原 assessments。為保持其他 payload metadata 不變，`rank` 欄仍保留該 source 的原 v1 bm25 score；**array order 才是 semantic rank**，沒有捏造 semantic confidence／score。

以下只比較同六題：五題有指定正向來源、四題有預標 confuser；O06 沒有完整 gold。不得把分母改寫成全 26 題或原 overlap 六題。

| 選擇方法 | Packer | 正向 hits / 5 | MRR（五題） | Confuser / 4 | 最大 runes |
| --- | --- | --- | --- | --- | --- |
| S0 v1 lexical | P0 full | 5/5 | 0.533333 | 4/4 | 1950 |
| S0 v1 lexical | P1 compact | 5/5 | 0.533333 | 4/4 | 1800 |
| S1 weighted | P0 full | 3/5 | 0.400000 | 4/4 | 1949 |
| S1 weighted | P1 compact | 3/5 | 0.400000 | 4/4 | 1799 |
| S2 phrase tier | P0 full | 3/5 | 0.400000 | 4/4 | 1949 |
| S2 phrase tier | P1 compact | 3/5 | 0.400000 | 4/4 | 1799 |
| S3 fresh subagent | P0 full | 5/5 | 0.866667 | 2/4 | 1922 |
| S3 fresh subagent | P1 compact | 5/5 | 0.866667 | 2/4 | 1777 |

沒有新方法的全資料 acceptance／獨立泛化驗證，人工 nonce/status 兩題亦未重跑 S3。資訊充分性、實際 tokens 與 source-read／adoption 一律 not evaluated，不能由命中與 JSON 長度代替。

## 問題焦點與剩餘錯誤

| 案例 | S3 位置／結果 | 判讀 |
| --- | --- | --- |
| R03 | 正向 rank 1，S0 為 2 | 能區分本地衍生狀態與不相干基線題材；另有摘要不足者明列 unknown |
| R08 | 指定 lesson rank 3，與 S0 相同；canonical 契約仍在前 | 不把相關 YAML 誤判不相關；gold 沒重標，兩種 payload 均保留指定 lesson |
| O02 | 正向 rank 1；confuser rank 2、標 background | 了解 viewport 問題，但 ranking-only packer 仍把背景 lesson 帶入；品質 gate FAIL |
| O03 | 正向 rank 1；confuser rank 46 | 固定池保留正向後，語意排序能找回 v2 權重排到 33 的來源 |
| O04 | 正向 rank 1；confuser rank 2、標 background | 看懂 now() 是描述性日誌，但沒有 admission 決策，仍帶入 age-gate lesson |
| O06 | confuser rank 5、標 not-applicable，P0/P1 均未輸出 | 看懂連線拒絕已排除；其他 reauthentication 候選的 domain／適用性仍未驗證，不推定已找到可採用解法 |

這證明現有 runner 可以完成語意理解工作，不必把所有理解寫成 Go 權重公式。也顯示 **重排不等於拒絕**：背景候選在其餘來源更不相關時仍可能排第二。不能為通過 0-confuser gate 而事後將 assessments 改為過濾器。

下一個獨立變項應是 workflow／agent 的明確 candidate disposition（可讀回、背景、不適用、unknown 與理由），再核對來源版本、前提及採用；需先固定規則與評分，分開記 retrieved、selected context、adopted，不以相似度或 agent 標記當權威。本輪沒有偷接 workflow 或宣告 Phase 3 通過。

## 資源與時間 gate 的限制

沒有額外 API 呼叫／採購，但既有 agent 仍消耗執行資源與帳號額度，無可取得 token usage 或價格，不能換算美元或填 0。額外 API 的 USD 2、tokenizer、價格 gates 本輪不適用；若未來改選該 runner 才需滿足。

外層時間紀錄不完整：前三題的 dispatch-to-parent-observation 為 69.880259 秒；後三題直到續跑後才落盤，為 6394.516939 秒；batch observation 為 6468.738843 秒。這些包含排程、工具與續跑／parent 記錄間隔，不是子模型純推論耗時。

本輪外層 deadline 沒有持續監測，因此 **300 秒／題、900 秒／批的時間 gate 未通過／未證明受控**。不能把不可得的實際完成時間換成較漂亮的數字，也不把已完成結果偽稱模型 timeout。保留為語意品質診斷，不作 latency、timeout 或 production p95 證據。之後若驗證時間上限，須使用不中斷的調度監測與可取得的完成事件時間，另跑，不修改本輪原始時間。

## 驗證與處置

完整 `go test ./...` 與 pre-push preflight 通過。Parser controls 覆蓋缺 ID、重複、外來 ID、Markdown fences、trailing output、未知欄位與無效 assessment；不默默 fallback。原始 6 request / pool／mapping 不變，所有 raw／run hash 已核對；公開 JSON report 的字串 hyphen 採標準 Unicode escape，解碼後資料完全相同，避免既有公開檔名被誤判為金鑰；以獨立 Go JSON encoder 重算 12 份 final payload，均與 report 相同且 ≤2000 runes。

重算只需 `AI_SKILL_RETRIEVAL_AGENT_REPORT=<new absolute report path>`，在 CLI module 執行 `go test ./internal/app -run '^TestRetrievalAgentPilotScore$' -count=1 -v`。它只讀已保存回應，不呼叫模型或 API。不能以重算當成第二次模型重複試驗。

處置：runner 方案可行，pilot 已執行且有有限增益；0-confuser、完整審閱、資訊充分性、時間控制與全資料 gate 皆未通過。Phase 2 繼續 active，production 搜尋與採用鏈保持未交付。先前 v3 harness CI [40045239412](https://github.com/linyihong/Ai-skill/actions/runs/40045239412) 已三平台 tests / artifacts 全部成功；新 scorer CI 需另查，不沿用舊 run。
