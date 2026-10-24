# Phase 3：supersession consumer 回放與 Windows 暫緩

## 範圍與結果

使用者明確指示 Windows 暫時不用理會。本輪沒有 SSH、主機探測或遠端操作；先前 native 當前環境仍 unknown，Windows 工作標 `deferred-by-user`，只有使用者明確恢復 Windows 調查時才重啟。原始專案文件已讀回，不再混同於主機當前修正驗證。暫緩不表示修正通過，也不阻擋本輪 workflow 採用／拒絕鏈。

補齊既定 supersession consumer 情境：兩次現有 agent 執行 production prepare → watcher → complete → source readback → decision。沒有新 runner／selector 方法、黑名單、embedding、模型 API 或計數工程。

| Query intent | 待讀回來源 | 實際處置 |
| --- | --- | --- |
| 現行使用 | cross-cutting review/checklist.md | 讀現行 Contract Governance／Reusable Guidance Boundary，採用為本輪 evidence review 的現行清單；舊 stable inbound path 只作背景，不填主要 payload |
| 歷史追溯 | software-delivery/review-checklist.md | 讀其遷移宣告，再讀遷移前 Git blob；歷史 canonical 字樣只作歷史證據，不恢復成現行 owner |

同一真實 ownership 遷移在不同 intent 下得到不同用途。候選、實際回讀、採用前 SHA 一致，full summary／need_to_read 保留，沒有 budget deferral 或 background fill。Canonical status／scope／replaces 未宣告仍為 unknown；consumer 讀正文的明確遷移宣告，沒有虛構 metadata，也沒有把相關性當權威。

## 受控曝光的邊界

兩份原文是本 repo 的真實來源，位元組及遷移關係未改；歷史內容由真實 `4a68b432^` Git blob 讀回。本輪將它們複製到隔離 corpus，新增兩個**僅供測試的 summary 索引入口**，讓正式 source-level CLI 同時提供兩來源。Fixture summary 以獨特 query marker 限制 pool，不寫 canonical summaries、不宣稱 production 已索引這兩份文件，也不調參。

此處是 source-informed current-agent 的受控 consumer integration，不是 gold-blind selector 品質／泛化試驗。兩題不加入前輪的兩個獨立真實任務分母。它補上「搜尋候選進來後，現行採用與歷史追溯須分開處置」的既定負面情境，不新增「自主通用 supersession 演算法」驗收。

原始 request／raw bytes、正規化 transport bundle／watcher／complete、[source／fixture／binary／index 指紋](phase3-supersession-manifest.json) 與[回讀決策](phase3-supersession-decisions.json) 均保存。外層路徑正規化不回改原 request／raw。Fixture 可按 manifest 由固定 corpus、兩份來源與兩份 summary 重新建立；未改 repo production index。未知 tokens 仍依使用者指示暫緩。

## Phase 3 gate 判定

既定範圍內的 current-agent consumer 驗收已完成：

- [x] 查找／失敗後重查／相同前提與 checksum 下合理略過，以及不當略過拒絕：見 [真實 dogfood](phase3-dogfood.md) 與 events。
- [x] 兩項不同的實際任務：runtime.db 年齡警告的 evidence-scope 採用／無據修補拒絕；native 已知案例的原始證據回讀／未核對修法拒絕。Unknown 是安全終態，不要求靠猜測變成修復 PASS。
- [x] Domain／project／前提／lifecycle／authority 分開；來源不升格、blocker 不因搜尋命中結清。
- [x] 來源 A→B 重驗及 C 再變時停止沿用；版本變化不直接推翻結論。
- [x] 已取代的歷史追溯與現行採用差異：本輪兩個 intent，實際 CLI 候選與來源回讀。
- [x] Required-read、route hit/miss、feature-off 直接來源 fallback、缺索引／stale／空 pool／timeout／晚到輸出的邊界：前輪實際 consumer 回放及既有 gate regression 分開記錄。
- [x] 實際 message／trace／來源版本／用途、去敏／retention 及使用者核定的成本處置可回查；runtime validate、來源 SHA 與 request/raw/complete 身分鏈通過。

本結論只限現有 agent 的兩個有範圍真實任務及明列的受控負例；不支持自動化適用性判定、跨任務泛化、Windows 原生修正、產品品質或實際 token 上限通過。過去報告的未完成判定保留為當時狀態，這份補證與使用者暫緩指示後才結清 Phase 3。整份主計畫仍 active，Phase 4 條件式方法決策與整體收尾尚未完成。
