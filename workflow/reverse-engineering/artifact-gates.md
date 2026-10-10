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

APK 動態筆記模板仍用 [`../apk-analysis/artifact-gates/evidence-chain.md`](../apk-analysis/artifact-gates/evidence-chain.md)。
