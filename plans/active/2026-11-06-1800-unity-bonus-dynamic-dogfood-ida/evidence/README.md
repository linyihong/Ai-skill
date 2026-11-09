# Evidence — 動態 dogfood 與 Ghidra

Plan: [主計畫](../_plan.md)。

## 引用規則

引用 evidence 檔案路徑，不使用行號。每筆保留可獨立閱讀的去敏實驗、資料形狀、關鍵事件／數值及裁決。跨庫引用用於追溯原始材料，移除專案存取後仍能理解如何取得與判讀證據。原始 trace、位址、畫面與裝置細節留業務專案。

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
| 2026-11-09-guard | [隔離 lab 中止](guarded-cancellation-20261109.md) | scoped PASS | 驗證完整委派所有權、兩次失效及重進復跑 |
| 2026-11-09-reuse | [立即重用時間對照](immediate-reuse-20261109.md) | immediate timing FAIL；delayed scoped PASS | 三組對照，終點一致而時間不同 |
| 2026-11-09-lap | [完整外圈回放](full-lap-replay-20261109.md) | scoped PASS | 升級格起訖，22 段逐格、原速影片與位置序列 |

## Workflow 技巧對照

- [靜態基線](ghidra-static-baseline-20261109.md) 保存方法邊界、位址換算與指令裁決。
- [自然對照](natural-runtime-20261109.md) 保存代表段落、接收者配對及巢狀事件。
- [相鄰格回放](dynamic-replay-20261109.md) 保存根物件與渲染子物件的實際位置序列。
- [到站](stop-replay-20261109.md) 和 [獎勵鏈](reward-chain-replay-20261109.md) 保存分層 oracle 與正負對照。
- [單純重設](cancellation-reset-20261109.md) 和 [隔離中止](guarded-cancellation-20261109.md) 保存失敗反證、完整委派所有權驗證與復跑。

- [立即重用](immediate-reuse-20261109.md) 保存三組實際收尾時間與基線反例。

各頁的去敏數據與判讀依據隨本計畫保存。相關取證方法見 [Unity 動態流程驗證](../../../../analysis/unity/workflows/unity-dynamic-flow-validation.md)。
