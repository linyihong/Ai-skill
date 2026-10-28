# Reverse-Engineering Artifact Gates

宣稱「分析完成」前須滿足：

| Gate | 要求 |
| --- | --- |
| Authorization recorded | 具名標的 + 允許操作 |
| Artifact identity | path 代號 + digest 或 `digest_unknown` |
| Evidence discipline | findings 分開 observation／inference／unknown |
| Sanitization | 進庫或可分享產物無 token／私人 host／絕對私人路徑 |
| Engine honesty | 靜態不可宣稱 runtime 覆蓋；缺引擎不可假裝已用 MCP |
| Handoff clarity | 跨族時寫明下一站與 blocking unknowns |
| Relations delivered | 若本輪產出證據包：含節點／請求／符號與資源或行為的映射；孤立檔不算 ready（見 [`../../analysis/reverse-engineering/collection/collection-discipline.md`](../../analysis/reverse-engineering/collection/collection-discipline.md)） |
| Capture inventory sanitized | 若審查或產出 inventory：僅明文＋可觀察 metadata；無 token／密文當 body（見 [`../../analysis/reverse-engineering/collection/capture-review.md`](../../analysis/reverse-engineering/collection/capture-review.md)） |
| Fail-closed if no plaintext | pinning／無明文時不得發明 path／body；標 needed_authority 或 accepted_gap |
| Ready folder contract | 若專案約定證據目錄：映射與筆記齊備才可宣稱 handoff ready；否則標 `not_ready` |

APK 動態筆記模板仍用 [`../apk-analysis/artifact-gates/evidence-chain.md`](../apk-analysis/artifact-gates/evidence-chain.md)。
