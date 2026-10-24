# 檢索與 workflow 採用計畫：完成閉環

## 完成範圍

使用者已分別接受縮限 Phase 1、Phase 2 實際 production 回放及 Phase 3 既定 current-agent consumer 驗收。Phase 4 不啟用 vector、不新增額外 reranker／模型服務，保留現有 agent 語意選擇；依據與 deferred owner／重啟條件見 [決策與交接](phase4-decision-and-handoff.md)。不宣稱沒有任何語意缺口。

| Gate | 結果與證據 |
| --- | --- |
| Phase 0／1 | [盤點](phase0-inventory.md)、[baseline](phase1-baseline.md)、[執行契約](phase2-execution-contract.md)；固定診斷基準與資料身分，非全庫品質 |
| Phase 2 | [production E2E](phase2-production-wiring.md)、[結果](phase2-production-results.json)、[原 CI](phase2-production-ci.json)；來源、confuser、完整 payload、unknown、lifecycle／失敗處置，限制仍保留 |
| Phase 3 | [實際任務](phase3-dogfood.md)、[source migration consumer](phase3-supersession-consumer.md)；找到→讀回→前提→方法採用／修法拒絕，unknown 可為有效終態 |
| Phase 4 | [方法決策](phase4-decision-and-handoff.md)；未啟動新增方法實驗，沒有新 holdout／模型／成本實驗 PASS |
| Canonical／consumer | workflow/software-delivery/validation.md → validation-retrieval.md → decision template；runtime lookup prepare/watch/complete；source version／authority／required_read 與 blocker 保持 |
| Archive continuity | 整個 evidence folder 搬至 archived，同步 test-only harness 及文件連結；原 dataset／request／raw／payload／prompt 的 non-Markdown bytes 逐檔核對不變 |

## 本輪 closure 檢查

- Archive 後跑既有 Go suite，確保 replay harness 讀取歸檔後證據；不是重開模型／corpus 實驗。原本只有 test-only 路徑改動；歸檔 gate 實際攔下凍結 JSON 內的 Markdown-shaped 字串後，另修正 archival validator 僅掃 `.md`，保留所有 rename targets 的 inbound 驗證。該必要 closure 修正已有 red→green regression，依規則重建五平台 binaries；沒有新增檢索方法。
- 依 closure 規則執行 runtime refresh／validate、plans validate、diff 與連結核對；本輪只調 plan／test fixture 連結，不新增 canonical runtime contract，故不重編譯 runtime.db。
- 所有不可得／暫緩資料仍按原紀錄處置；原 runtime.db age warning 不為消除 warning 而改 DB。
- 歸檔目錄／plans 狀態／活躍引用同步；歷史固定 commit URLs 及原實驗 raw 不重寫。
- 最終檢查見 [final-validation](final-validation.json)；test-only CI readback 存本目錄的 final-ci 產物；依原始 commit 身分辨識各次驗證，後續純文件 readback commit 不宣稱重跑 CI。
- 完成 commit／push／遠端 ref 與 clean status 讀回後才回報整體結清；local goal 完成後刪除，不留臨時 goal 當長期 roadmap。

## 歸檔 validator 修正

原 validator 把搬移的 JSON／JSONL／TXT／B64 全加進 Markdown 掃描輸入，從 evidence container 目錄解析原始來源摘要內的相對連結，會迫使人改寫凍結資料。修正位於 portable KGE 的文件篩選，只解析 `.md`；rename map 仍保留所有 evidence，因此真正 Markdown 指向已搬移 JSON 的 inbound links 仍阻擋。新 integration fixture 實際 Git archive staging，修正前因 raw strings 失敗、修正後通過，反向 inbound 負例仍能檢出。原生 `.md` inbound/outbound 與 KGE 規則測試亦通過。這不是 opt-out 或新的治理 gate；未繞過 hook，245份既有非 Markdown bytes 仍未改。

## 保留限制

完成的是明訂資料與 current-agent 執行範圍；不支持未曝光任務泛化、全 corpus recall、自然無答案誤推率、任意 metadata 遷移自動辨識、實際總 context/token usage 或 Windows 修正／產品品質 PASS。這些是 [有 owner 的 deferred 項目](phase4-decision-and-handoff.md)，不是隱藏完成項，也不是另建的普遍 blocker。
