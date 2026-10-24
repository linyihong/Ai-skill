<!-- archival-provenance -->
<!-- Fixed-version historical URLs and corpus paths below retain their original active location. -->

# Phase 1 審核包與 Gate 狀態

日期：2026-10-24。Baseline 已建立；**Phase 1 gate 完整通過尚待核驗**。本次不執行 Phase 2、不修改 baseline 題目、labels、split、acceptance 或量測結果。此段保留當時原範圍 gate 狀態；後續使用者接受 baseline 後，主 plan 已縮限 Phase 1，剩餘責任與執行控制見 [phase2-execution-contract.md](phase2-execution-contract.md)，不回溯宣稱原完整 gate 通過。

## 可獨立讀取的固定版本

公開 GitHub repo，固定 revision `8ed2acd0c7796988347eed942ea9712017e5fcd7`：

- [phase1-baseline.md — 原始報告](https://github.com/linyihong/Ai-skill/blob/8ed2acd0c7796988347eed942ea9712017e5fcd7/plans/active/2026-10-23-2301-knowledge-retrieval-workflow-closure/evidence/phase1-baseline.md)
- [baseline-results.json — 逐題結果與量測指紋](https://github.com/linyihong/Ai-skill/blob/8ed2acd0c7796988347eed942ea9712017e5fcd7/plans/active/2026-10-23-2301-knowledge-retrieval-workflow-closure/evidence/baseline-results.json)
- [baseline-cases.json — 題目、來源、split、expected sources 與 acceptance](https://github.com/linyihong/Ai-skill/blob/8ed2acd0c7796988347eed942ea9712017e5fcd7/plans/active/2026-10-23-2301-knowledge-retrieval-workflow-closure/evidence/baseline-cases.json)
- [phase1-protocol.md — 評分與留出資料使用規則](https://github.com/linyihong/Ai-skill/blob/8ed2acd0c7796988347eed942ea9712017e5fcd7/plans/active/2026-10-23-2301-knowledge-retrieval-workflow-closure/evidence/phase1-protocol.md)
- [replay harness — 評分與漏找分類實作](https://github.com/linyihong/Ai-skill/blob/8ed2acd0c7796988347eed942ea9712017e5fcd7/scripts/ai-skill-cli/internal/app/retrieval_baseline_test.go)

## 指標重算

本次從 cases 的 expected_sources 與 results 的實際 sources 重算，未只引用 summary。Dataset SHA 與 results 記錄相符。

| 組別 | 至少命中一份來源的題數 / 正例題數 | 命中 expected sources / 全部 expected sources | Source-level Recall@5 |
| --- | --- | --- | --- |
| Diagnostic | 1 / 14 | 1 / 14 | 0.071429 |
| 預留正例 | 0 / 4 | 0 / 4 | 0 |

兩者數值相同是因本版每正例只有一份 expected source；一般情況不能互換。MRR 只平均正例，重複 source 去重後在原 top 5 atoms 內計分，不補第 6 名。人工負面控制 0/2 誤推不加入正例分母，也不能推估自然情境誤推率。

## 漏找證據的強度

13 題診斷漏找的自動分類都是 `query_processing_or_indexed_fields`：runner 查到 source row 存在，但該 source 對原始 quoted FTS query 的 match 為 0。這能區分「source row 完全不存在」與「查詢 / 已索引欄位不足」，**不能**逐題證明是哪一項造成。

Source row 存在不代表正文或摘要完整進入 FTS。已讀到 R04 / R11 的舊標題與摘要格式、以及 generator 的固定 heading 規則，足以證明存在欄位擷取風險；人工 FTS controls 能證明片語與 CJK tokenizer 行為，但尚無 query-only / extraction-only 的逐題對照增益。先前摘要把兩者合稱主要缺口，不能當作逐題因果歸屬完成。

Phase 2 的預先實驗矩陣（未執行）：

| 組別 | Query | Extraction / corpus | 用途 |
| --- | --- | --- | --- |
| B0 | 既有 | 既有 | 不可覆寫的 baseline |
| Q | 只改 query 處理 | 保持既有 index records | 量測 query-only 增益 |
| E | 保持既有 query | 只修 heading / summary 擷取；不新增 corpus | 量測 extraction-only 增益 |
| QE | query 修正 | extraction 修正；同一 corpus | 量測合併增益與交互作用 |

新增正文覆蓋必須另組實驗，不能混入 E 假裝只有 extraction 改變。全部組別固定同一 dataset、source revision、top-k 與評分；記錄 generator / query 版本、每題前後排名、hit / miss 與 latency / preview 成本。只用 diagnostic 設計修法，凍結方法後才評估預留組一次。

## Acceptance 與尚未通過事項

已凍結於 cases：diagnostic ≥ 8/14、預留 Recall / MRR 不低於 baseline、人工控制誤推 0/2、warm CLI p95 ≤ 500ms、preview estimate ≤ 2000 UTF-8 runes。最後一項是代理量尺，**不是實際模型 token budget 已驗收**；採用、required-read、activation、blocker 的安全表現仍為 `not evaluated`。

預留 baseline 0/4 使「不退步」門檻很弱，不能單憑它宣稱泛化改善。Phase 1 未涵蓋全庫 intelligence relevance、自然語意無答案或真實 supersession 採用；這些限制沒有因人工 controls 通過而消失。Phase 2 前仍需核對此初版範圍與代理量尺是否足夠作交付 acceptance，不靜默降低原計畫要求。

## 預留資料曝光紀錄

| 時點 | 已曝光內容 | 用途 / 限制 |
| --- | --- | --- |
| 建資料與核來源 | 所有題目、來源與 expected labels | Author-labelled，非盲測；同家族不跨 split |
| 初次與重跑 baseline | aggregate；診斷集逐題 misses | 固定 baseline；未依結果改題或改門檻 |
| 保存原始結果 | 完整 JSON 含預留個案輸出 | 可審核，未密封；不能宣稱不可見的獨立測試集 |
| Phase 1 方法診斷 | diagnostic misses 與 R04 / R11 格式 | 不用預留個案結果設計 query / extraction 修法 |
| 本次審核 | 從完整 JSON 重算分母與 aggregate | 已處理原始結果，保守視為非密封預留組；不主張嚴格獨立 holdout |

Phase 2 只可稱此資料為固定 regression / reserved comparison set。若要以獨立 holdout 支持泛化，須另收未參與診斷與方法選擇的資料；Phase 4 也必須重查獨立性，不能沿用名稱來保證。

## 驗證狀態

本次只做 artifact 重算、SHA 與公開遠端路徑確認；未把重算當 fresh corpus replay 或獨立審查通過。當前 CI 詳情應直接查 [run 40027219933](https://github.com/linyihong/Ai-skill/actions/runs/40027219933)，不可引用上一輪成功來替本次 run 背書。
