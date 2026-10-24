# v3 D1 — disposition admission 重播

結論：六題曝光診斷的候選選擇 slice 通過；Phase 2 維持未完成。沒有新增模型呼叫、修改 production 搜尋或 workflow consumer。先固定的協議為 `df323fa2`；test-only scorer 為 `fd8a8cc8`。

## 固定輸入與唯一變項

沿用原六份 raw、request、opaque mapping、候選池、排序、來源 metadata、gold 與 P0/P1。只增加 D1：direct 才進待讀回清單與候選層 preview，background 不自動填入，not-applicable 本次不選，unknown／未標記均待確認。不增加背景用途 override。舊 agent-results 與所有原始回應保持不動。

[results](phase2-v3-disposition-results.json) 記錄 protocol、D0 results、pool、manifest、run、逐題 request／raw 及 scorer SHA；`scorer_revision` 指向已提交 scorer。permutations 中每題分別保存處置 trace 與原完整排序，沒有更改排名。253 筆來源出現次數全部保留：direct 9、background 15、not-applicable 8、unknown 221（其中 1 筆明確 unknown、220 筆 not assessed）。同一來源在不同問題的處置不升格成全域標記。

## 分層結果

以下分母均固定在同六題：R03、R08、O02、O03、O04、O06；五個指定正向來源、四題帶預標 confuser。O06 無完整 gold，不把額外來源一律算錯。

| 層／方法 | 指定正向 | MRR（5 正例） | confuser 題數 | 最大 P0 / P1 runes |
| --- | --- | --- | --- | --- |
| retrieved 完整池 | 5/5 | 不以處置評分 | 4/4 | 不適用 |
| D0 舊 ranking-only preview | 5/5 | 0.866667 | 2/4 | 1922 / 1777 |
| D1 direct 待讀回清單 | 5/5 | 0.866667 | 0/4 | 不適用 |
| D1 selected preview，P0 | 5/5 | 0.866667 | 0/4 | 1922 |
| D1 selected preview，P1 | 5/5 | 0.866667 | 0/4 | 1624 |
| verified context / adopted | not evaluated | not evaluated | not evaluated | not evaluated |

| 題目 | direct 數 | P0 / P1 runes | 指定來源位置 | 處置觀察 |
| --- | --- | --- | --- | --- |
| R03 | 1 | 528 / 528 | 1 | derived-state lesson 保留 |
| R08 | 3 | 1922 / 1624 | 3 | summary、canonical YAML 仍在歷史 lesson 前，沒有以 gold 重排 |
| O02 | 1 | 570 / 568 | 1 | viewport direct 保留；KPI background 不填入 |
| O03 | 1 | 486 / 486 | 1 | native-load direct 保留；不承接其因果結論 |
| O04 | 1 | 618 / 582 | 1 | nested-checkout direct 保留；描述性日誌 background 不填入 |
| O06 | 2 | 1043 / 1043 | 無完整 gold | peer confuser 不選；兩份 auth 線索仍未確認適用 |

12 份最終 preview 均未因預算丟棄 direct；P1 的摘要政策與原協議一致。runes 計 compact JSON results array，未包含外層診斷／處置 trace，也不是實際模型 tokens。不能由較短輸出宣稱資訊充分。

## 來源讀回：身份確認與適用性分開

[readback audit](phase2-v3-disposition-readback.json) 保存 9 份 direct 全文讀回的 revision／checksum、當前 worktree checksum 與 scope 備註。固定 Git blob 均符合 pool checksum；9/9 當前檔案也一致。外層已知 gold，這只是知情診斷，不是第二道 admission filter；沒有改 D1 分數。

R03 要先確認 derived state；O02 要有真實 viewport 量測；O04 要證明 shared-pool nested checkout。R08 三份來源的既有權威層級保持不同。O03 仍為 candidate、無獨立 reuse 證據，需版本／平台／模型／錯誤指紋與排除條件核對。

尤其 O06：APK session 與 guest bootstrap 的來源前提，沒有被「peer 已連線、token 撤銷」這個 query 證明。兩筆 direct 只提供待研究線索，適用性 unknown，不能宣稱已找到有效解法。checksum 相同只確認讀到同一份來源，不確認適用性或採用安全。

## 驗證與未完成項

- `go test ./...` 通過；新增 admission controls 驗證缺標記、unknown、background 不 fallback，direct 保持原排序且不修改 response。
- opt-in D1 replay 通過，原 parser 仍檢查 raw 完整排列與 SHA；另用獨立 Python 核對 12 份 direct-only 順序、來源 metadata、不掉 direct；獨立 Go JSON 序列化重算全部 payload runes。
- 審閱前輪 scorer CI `40054758419` 已成功。本輪 scorer CI `40055737331` 在本報告撰寫時執行中，不能當作已通過。

這是使用者根據已曝光失敗提出的方法，並非獨立泛化證據。全 26 題語意處置／審閱、自然無答案、資訊充分性、實際 tokens、受控 runner 時間與真實 workflow 採用鏈仍待完成。新離線重播不修復上一輪時間監測缺口。下一個落地門檻是 workflow 消費待讀回清單，逐筆核對 scope／版本／supersession／required_read 並記錄採用或拒絕，而不是把 preview 改名成 verified context。
