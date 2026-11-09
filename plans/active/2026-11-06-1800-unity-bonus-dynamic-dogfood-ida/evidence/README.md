# Evidence — 動態 dogfood 與 IDA

Plan: [主計畫](../_plan.md)。

## 引用規則

引用 evidence 檔案路徑，不使用行號；具體執行證據留業務專案。

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
| --- | --- | --- | --- |
| 2026-11-06-planning | [計畫與 workflow 補強](planning-baseline.md) | docs-ready | A／B oracle 與實測步驟已定義，尚未實測 |

| 2026-11-06-environment | [環境 preflight](environment-preflight-20261106.md) | B tooling blocked | 常見位置未找到 IDA；版本／授權 unknown，材料匹配既有基線；未實測動態流程 |
| 2026-11-09-ghidra | [Ghidra 靜態基線](ghidra-static-baseline-20261109.md) | static baseline PASS | ARM64 樣本反編譯與兩次一致匯出；Runtime 待驗 |
| 2026-11-09-natural | [自然 Runtime 對照](natural-runtime-20261109.md) | scoped PASS | 首個靜態預測獲支持，人工回放待測 |
| 2026-11-09-replay | [相鄰格回放](dynamic-replay-20261109.md) | scoped PASS | 四次中途移動、同一行程重進、獨立渲染位置驗收 |
| 2026-11-09-stop | [到站回放](stop-replay-20261109.md) | scoped PASS | 普通與特殊格到站各兩次，獎勵鏈待測 |
| 2026-11-09-reward | [+1 鏈回放](reward-chain-replay-20261109.md) | scoped PASS | 有／無獎勵資料兩組驗收，重啟重進後重跑通過 |
| 2026-11-09-cancel | [重設與取消對照](cancellation-reset-20261109.md) | reset ruled out；recovery PASS | 三次遲到收尾，等待收尾後重進兩次成功 |
