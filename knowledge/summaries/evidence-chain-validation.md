# Evidence chain validation

| 欄位 | 值 |
| --- | --- |
| Source path | [Evidence chain 候選智慧](../../intelligence/engineering/execution/validation-reasoning/evidence-chain-validation.md) |
| Summary | 宣稱範圍決定證據入口：host composition、mutation effective layer／template isolation、branch invariant release gate 均須驗到最後決定結果的層與時間點。 |
| When to read | 元件測試綠但最終行為錯誤、突變結果疑似污染，或 commit gate 未涵蓋發佈時。 |
| Do not use for | 自足局部契約不必一律提升為整體測試；接線與合成反例不是獨立重用證據。 |
| Validation signal | 舊組合可重現、有效突變可讀回、無關變更仍觸發分支不變式檢查。 |
| Last checked | 2026-10-26 |
