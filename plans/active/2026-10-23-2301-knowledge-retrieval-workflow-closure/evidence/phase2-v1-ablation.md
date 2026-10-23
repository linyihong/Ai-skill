# Phase 2 v1 — 四組對照：找回增益與放寬副作用

本輪四組與兩份 dataset 已實跑，產物配對 validator 均通過。**候選品質未通過，不併入 main production 搜尋**；計畫繼續 active，不進入 Phase 3。原 baseline / labels / acceptance 不變。

## 方法與產物

- 固定 corpus `5f12fc3906051e40228959e9e6bfaed714e0658e`，darwin/arm64、Go 1.26.2；每題 1 次初次 + 20 次 warm CLI invocation。OS cache 未控制，不能稱 cold-start。
- Q source `4e84b3df`：整句 FTS phrase 改為 Unicode letter / number / underscore 詞項 OR，大小寫去重、最多 64 詞；每詞 quote，使用者字串不作 FTS 語法。未加同義詞、題目詞表或 CJK 分詞。這是量測簡單放寬的候選，不是最佳查詢策略宣稱。
- E source `fd5f520d`：feedback title / One-line Summary 支援 ATX 1–6 級標題，summary 限第一段、遇下一標題停止，跳過 fenced examples。只變 title / summary，不加入全文、來源或 lifecycle 推論；131/476 atoms 的欄位變動見 [extraction delta](phase2-v1-extraction-delta.json)。
- 候選 source / 五平台 binaries / tests 保存在實驗分支 [experiment/retrieval-qe-v1](https://github.com/linyihong/Ai-skill/tree/df95535b)。Main 沒有採用候選 Go source 或 binary。
- B0 / E binary SHA `fde3d631ddbbee662ba1e63e8807dbfa9fcb53387a8d2d2ca89732dfb24f006f`；Q / QE SHA `d59cc24c93a7b0f448b1fba1f3e61375610226edca63ffbe32ae30ad1b19cf20`，version `repo-fd5f520d`。新執行檔包含 Q/E source，但 replay query 路徑只讀指定 index，不呼叫 extraction，因此 Q 組仍使用完全相同的舊 index。
- [OLD manifest](phase2-v1-old-index-manifest.json) / [NEW manifest](phase2-v1-new-index-manifest.json) 記錄所有完整身分；source path/checksum 集合一致。NEW generator 包含 Q function，但生成索引只因 E 擷取變動；query 不參與 generator。SQLite 產物是可重建暫存，保留 manifest 與欄位差異，不建立新的 canonical 知識。

## 固定歷史資料結果

每題一個 expected source，因此本資料 case hits / expected-source hits 分母剛好相同；不可外推一般情境。前五 atoms 先截斷再按 source 去重，不補位。

| 組別 | Diagnostic hit / 14 | Recall@5 | MRR@5 | 固定比較正例 / 4 | 人工負面誤推 / 2 | Warm p95 ms | 最大 preview runes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| [B0](phase2-v1-baseline-B0.json) | 1/14 | 0.071429 | 0.071429 | 0/4 | 0/2 | 10.509 | 545 |
| [Q](phase2-v1-baseline-Q.json) | 13/14 | 0.928571 | 0.797619 | 4/4 | 0/2 | 12.618 | 3325 |
| [E](phase2-v1-baseline-E.json) | 3/14 | 0.214286 | 0.214286 | 0/4 | 0/2 | 10.445 | 698 |
| [QE](phase2-v1-baseline-QE.json) | 14/14 | 1.000000 | 0.880952 | 4/4 | 0/2 | 11.315 | 3327 |

E 相對 B0 新找回 R04 / R11，Q 找回其餘多數題但漏 R11；QE 再找回 R11。因而本版存在獨立的欄位擷取增益與 query 放寬增益；不能把全部漏找歸因於同一原因，亦不支持直接導入向量。E 單獨仍未達 8/14。

## Overlap：正向與 confuser 必須一起看

| 組別 | 正向錨點 / 5 | MRR@5 | 已標 confuser 命中 / 6 | 未審閱候選 occurrences | 最大 preview runes |
| --- | --- | --- | --- | --- | --- |
| [B0](phase2-v1-overlap-B0.json) | 0/5 | 0 | 0/6 | 0 | 4 |
| [Q](phase2-v1-overlap-Q.json) | 4/5 | 0.566667 | 6/6 | 20 | 3459 |
| [E](phase2-v1-overlap-E.json) | 0/5 | 0 | 0/6 | 0 | 4 |
| [QE](phase2-v1-overlap-QE.json) | 5/5 | 0.766667 | 6/6 | 19 | 3503 |

B0 / E 都是空結果，不報 precision=100%。Q / QE 每一題都帶入事先排除的來源。例 O01 的主題是 JSON 數字比較，QE rank 1 找到正向來源，但 rank 2 同時出現已排除的 host-pipeline lesson；O06 無完整 gold，只有已標 confuser 可確定，其他來源維持 unreviewed。這些是人工 overlap 控制，不是一般誤推率估計。

診斷 Q / QE 各有 51 個未審閱候選 occurrences；跨組、跨題計數不等於 unique source 數。[Review queue](phase2-v1-review-queue.json) 保存所有相對 B0 新增診斷來源及 overlap 來源，附 query、context、corpus revision、source SHA、既有標記與待決狀態。191 rows 包含重複實驗組；不自動把未列 expected 的來源判為 irrelevant。由於本候選已明確失敗，本輪不把未覆核清單冒充完整人工驗收；reviewer 保持空值，審閱 gate 未結清。

## Gate 處置與下一步

| 驗收 | 處置 |
| --- | --- |
| 四組 artifact pairing / corpus / datasets | PASS，兩份 dataset 的 matrix validator 均通過 |
| Diagnostic ≥ 8/14、固定比較不退步、人工負面 0/2 | Q / QE 通過此有限指標；E 未通過 diagnostic |
| 已標 confuser = 0 | Q / QE FAIL，均 6/6 |
| Preview ≤ 2000 runes | Q / QE FAIL，最大 3503；runes 仍非實際模型 tokens |
| 新增候選完整覆核 | NOT COMPLETE，保留逐筆 queue，不視為 0 錯誤 |
| p95 ≤ 500ms | 本機實驗通過；不外推其他平台或負載 |
| Intelligence 覆蓋、workflow 採用、實際 tokens、獨立泛化 | NOT EVALUATED，責任沿用主計畫 |

本版不發布為 production 改善。下一個 query 候選須先設計可測的 evidence / 詞項約束、精確名稱保護及 context 組裝上限，再以新版本保存方法和同一固定比較集；所有 v1 結果保留，已曝光資料永遠不改稱獨立 holdout。E 的有限增益值得保留為候選，但本輪不將其不足的品質結果視為整體 gate 通過。無需新增 vector 基礎設施。

## 重播與驗證

依 [execution contract](phase2-execution-contract.md) 分別在 OLD generator `830f58fc` 與 NEW generator `fd5f520d` 建 snapshot，corpus 固定上述版本；舊 generator 可攜帶 `72c53d1e` 的兩份 test-only harness，不改 production Go source。manifest 的完整 index identity 綁定同一組 snapshot；B0=Q、E=QE。新 binary 取實驗分支 df95535b 已提交產物，舊 binary 取固定 corpus，先核對 SHA，再按契約傳 env 執行八次 replay 與兩次 matrix validator。報告原始 path 是本機執行紀錄，重播者需指定自己的絕對路徑，不能依賴 /tmp 或報告 path 自動定位。

候選分支完整 `go test ./...` 通過（含 host binary smoke / SHA / source parity）；歷史 stale-control 測試仍確認 deprecated exact match 保持首位、status filter 排除 retired result、來源 checksum 改變被檢出，不再假設 OR 查詢只會回傳一筆。初次完整測試因該筆數假設失敗，修正 test assertion 後完整重跑通過，沒有以 targeted tests 代替完整驗證。

既有 CI `72c53d1e` 的 Linux / macOS / Windows tests 與 artifacts 已全部成功，run [40031262752](https://github.com/linyihong/Ai-skill/actions/runs/40031262752)。這不等同候選分支或本次 evidence commit 的最新 CI；後者需另行核對。
