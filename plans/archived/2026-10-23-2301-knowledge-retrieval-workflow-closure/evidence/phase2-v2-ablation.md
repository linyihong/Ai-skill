# Phase 2 v2 — 最終輸出重評：預算通過，候選品質仍失敗

**v2 實驗已完成，候選拒收；Phase 2 維持未完成。** 方法先固定於 main `146f15e0` 的 [v2 protocol](phase2-v2-protocol.md)，再執行重播，未依結果修改權重、shortlist、labels 或 acceptance。正式搜尋未採用本版。

## 產物與評分對象

候選保存在 [experiment/retrieval-qe-v2，固定版本 dead83a5](https://github.com/linyihong/Ai-skill/tree/dead83a5)。Production source commit 為 `58b8b5bf65f9b1fc3e650833feece037155788a4`，五平台 binary refresh 為 `b06a1054`；後續 `dead83a5` 只補 test trace，不改 production source。完整 Go tests 通過，包括 SHA / source parity / host binary smoke。

- Corpus 保持 `5f12fc3906051e40228959e9e6bfaed714e0658e`，datasets / labels 完全不變。仍是已曝光的固定比較資料，沒有獨立泛化驗證。
- B0 / Q 使用 v1 [OLD manifest](phase2-v1-old-index-manifest.json)，E / QE 使用同一份 v1 [NEW manifest](phase2-v1-new-index-manifest.json)。未重建或偷偷更新 extraction。舊 binary SHA 保持 `fde3d631ddbbee662ba1e63e8807dbfa9fcb53387a8d2d2ca89732dfb24f006f`。
- Q / QE binary SHA `e4ccdca67cc941773359fb4fe1808b0b196820a0c667aac3e4e0f64613adf0b3`，version `repo-58b8b5bf`、完整 source commit 見各報告 version readback。新 binary 只查指定 index，沒有在 replay 中重新 extraction。
- 四組／兩份資料，共八次 replay，兩次 matrix validator 均 PASS。Darwin/arm64、Go 1.26.2；每題 1 次初次 + 20 次 warm invocation，非 cold-start 實驗。
- v2 reports 新增 `consumer_results`：實際 pinned binary 的完整最終 QueryResult 欄位。命中、MRR、confuser、runes 均由此最終輸出計算，不以 shortlist 前候選充當命中。空 JSON results 在 CLI 的 omitempty roundtrip 後記錄為 nil，原 rune proxy 為 4，沿用既有協議。
- JSON 中部分反斜線以標準 Unicode escape 編碼，保留程式碼範例的原字串；轉換前後 JSON decode 物件逐欄相等，未修改實際 summary 或結果。R04 的 JSON 字串比較範例可回到固定 corpus 原始來源核對。

## 最終歷史資料結果

下列是完成 phrase / fallback、欄位權重、source 去重、top1 與 2000-rune 限制後的 results；不是 raw candidate 指標。v2 最多回傳一筆，因此本資料的 MRR@5 與 Recall@5 相同，不能據此宣稱所有原始候選排名改善。

| 組別 | Diagnostic hit / 14 | Recall@5 | MRR@5 | 固定比較正例 / 4 | 人工負面誤推 / 2 | Warm p95 ms | 最大 preview runes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| [B0](phase2-v2-baseline-B0.json) | 1/14 | 0.071429 | 0.071429 | 0/4 | 0/2 | 11.810 | 545 |
| [Q](phase2-v2-baseline-Q.json) | 11/14 | 0.785714 | 0.785714 | 4/4 | 0/2 | 12.066 | 841 |
| [E](phase2-v2-baseline-E.json) | 3/14 | 0.214286 | 0.214286 | 0/4 | 0/2 | 10.597 | 698 |
| [QE](phase2-v2-baseline-QE.json) | 12/14 | 0.857143 | 0.857143 | 4/4 | 0/2 | 11.868 | 840 |

B0 / E 分數延續 v1；v1 Q / QE 為 13/14、14/14，本版各損失 2 題指定來源命中。固定比較資料 4/4 與 nonce / status 0/2 不外推一般誤推率。

## Overlap：成本縮小並不等於候選正確

| 組別 | 正向錨點 / 5 | MRR@5 | 已標 confuser / 6 | 未審閱 occurrences | 最大 preview runes |
| --- | --- | --- | --- | --- | --- |
| [B0](phase2-v2-overlap-B0.json) | 0/5 | 0 | 0/6 | 0 | 4 |
| [Q](phase2-v2-overlap-Q.json) | 2/5 | 0.400000 | 4/6 | 0 | 724 |
| [E](phase2-v2-overlap-E.json) | 0/5 | 0 | 0/6 | 0 | 4 |
| [QE](phase2-v2-overlap-QE.json) | 3/5 | 0.600000 | 3/6 | 0 | 693 |

QE 正向 O01 / O04 / O05 命中；O02 / O03 / O06 則只留下已標 confuser。v1 QE 為正向 5/5、confuser 6/6；v2 同時減少正向與 confuser，不能只報後者下降。B0 / E 仍是空結果，不報 precision=100%。

## 排序與輸出限制各自造成什麼

四份 trace 分別對照 [baseline Q](phase2-v2-baseline-Q-trace.json)、[baseline QE](phase2-v2-baseline-QE-trace.json)、[overlap Q](phase2-v2-overlap-Q-trace.json)、[overlap QE](phase2-v2-overlap-QE-trace.json)。Trace 呼叫固定 source 的 candidate function；每題 bound 後全部 QueryResult 欄位與實際 pinned CLI 結果比對相等。它是同 source 診斷，不冒稱另一份獨立 binary replay。預算前排名記錄為 atom rank；前五 atoms 再依原 source scoring 去重，不補位。

| QE 指標 | 預算前前五 atoms | 最終 consumer results |
| --- | --- | --- |
| Diagnostic expected hits | 14/14 | 12/14 |
| Overlap 正向錨點 | 4/5 | 3/5 |
| Overlap 已標 confuser | 6/6 | 3/6 |

| 案例 | 可核對的變化 | 判讀 |
| --- | --- | --- |
| R04 / R11 | 完整 phrase、正向 rank 1、最終保留 | 精確名稱／片語保護有執行；沒有 OR 補滿其餘位置 |
| R03 | 正向由 v1 rank 2 到 v2 預算前 rank 4，最終 top1 丟失 | 加權排序與 shortlist 均有影響；top1 不是 relevance gate |
| R08 | lesson 在預算前 rank 2，最終選 canonical executable-contract YAML | 指定來源漏找；該 YAML 可能相關，不自動標成 irrelevant，也不事後改 gold 讓分數通過 |
| O02 | 正向 rank 2、confuser rank 1，最終只留 confuser | 輸出縮減實際傷害正向召回 |
| O03 | 正向由 v1 rank 3 降至 v2 atom rank 33 | 欄位加權本身已造成排序退步，不能全歸因 top1 或摘要截短 |
| O06 | 已排除的 peer / ECONNREFUSED lesson 仍在 rank 1 | 詞彙匹配仍無法理解否定與主要問題；沒有使用六題來源黑名單 |

本 corpus 的 v2 輸出無摘要截短事件；成本下降主要來自 top1 與不重複整段 query 的 reason。2000-rune guard 與 Unicode / JSON escape 截短由合成測試驗證，不能宣稱已驗證真實被截短摘要仍保有充分資訊。所有最終輸出均在預算內；實際模型 tokens / workflow context 組裝仍 not evaluated。

## Gate 與審閱處置

| Gate | 結果 |
| --- | --- |
| 四組身分、資料一致性、兩次 matrix | PASS |
| Q / QE diagnostic ≥ 8/14、固定比較不退步、人工負面 0/2 | PASS，限定本固定比較資料 |
| 候選最終 preview ≤ 2000 runes、p95 ≤ 500ms | PASS，本機工程代理；最大 841 runes / 13.994ms |
| Overlap 已標 confuser = 0 | FAIL：Q 4/6、QE 3/6 |
| 新增候選完整覆核 | NOT COMPLETE，不以低未審閱數當成品質通過 |
| Intelligence 覆蓋、採用安全、真實 tokens、獨立泛化 | NOT EVALUATED |

[Review queue](phase2-v2-review-queue.json) 共 39 個跨組 occurrences；其中 Q / QE 的診斷各 2 個非 gold 來源保持 pending-review。Overlap 最終來源都已有預標記，因此 unreviewed=0，但仍有明確 confuser，不能視為已通過人工品質驗收。v1 191 rows 繼續保留未完成；依使用者此次指示，已拒收版本不要求全部補審。

下一輪先處理問題焦點／已排除前提與排序證據，避免將 title 權重或 top1 當成通用適用性判定；需另版固定方法，保存本次排序退步與 shortlist 損失作回歸比較。候選規則與 context metadata 的責任需沿用現有 query / workflow 邊界，不新增通用 trigger framework。本輪沒有新增 vector，也沒有 production 採用。

## 重播與執行驗證

依 [v2 protocol](phase2-v2-protocol.md) 與 [execution contract](phase2-execution-contract.md) 指定同一組 OLD / NEW index、固定 corpus 與 datasets，Q / QE 使用 source `58b8b5bf` 的 binary。實驗分支 replay 另設 `AI_SKILL_RETRIEVAL_CAPTURE_RESULTS=1` 保存實際 final fields；matrix 讀本次四份 report，兩份 dataset 分開驗證。

Trace 使用 `AI_SKILL_RETRIEVAL_TRACE_REPORT`、`AI_SKILL_RETRIEVAL_FINAL_REPORT`、`AI_SKILL_RETRIEVAL_CASES`、`AI_SKILL_RETRIEVAL_INDEX`、`AI_SKILL_RETRIEVAL_INDEX_MANIFEST`，執行 `go test ./internal/app -run '^TestRetrievalSelectionTrace$' -count=1 -v`；四次 trace 的 final fields equality 皆通過。依重播者位置指定絕對路徑，不從 report path 自動尋找 binary；generator revision 仍是 v1 manifest 的固定值。

從保存的 consumer_results 以 Go 標準 JSON encoder 另行重算全部 104 份 final payload，長度與報告逐題一致，所有來源順序一致；兩組已準備發布的 JSON reports 再跑 matrix validator 均通過。凍結 protocol、datasets 與 v1 queue 未變，v2 queue 的 39 筆 source SHA 均按固定 corpus 核對。

候選完整 `go test ./...` 及 pre-push preflight 通過。v1 CI [40032549355](https://github.com/linyihong/Ai-skill/actions/runs/40032549355) 已包含三平台 tests / artifacts 全部成功；v2 CI [40033981207](https://github.com/linyihong/Ai-skill/actions/runs/40033981207) 是另外一個執行，不能沿用 v1 的成功作背書，最新狀態另行讀回。
