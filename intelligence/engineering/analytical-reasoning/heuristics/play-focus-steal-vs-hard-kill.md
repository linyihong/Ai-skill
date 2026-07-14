# Play 搶焦 vs 硬殺（冷啟動）

`intelligence/engineering/analytical-reasoning/heuristics/play-focus-steal-vs-hard-kill.md`

## 問題

側載或裝置無 Google Account 時，冷啟動後畫面常變成 Play「登录／Unauthenticated」頁。分析者易做成錯誤結論：

1. 「靜態有 Pairip／CHECK_LICENSE → runtime 一定先過 LicenseActivity」  
2. 「沒帳號就完全進不了業務 UI／分析失敗」  
3. 「一定要先登 Google 才能做冷啟動 capture」

## 啟發式

| 信號 | 較可能解讀 | 下一步 |
| --- | --- | --- |
| Focus 落到 `com.android.vending`／Unauthenticated* | **Play UI 搶焦** | 做 A/B：允許 Play vs 循環 `force-stop` vending |
| ≤100ms 採樣看到短暫業務 Home／Login | 業務進程**有**起來 | 不要只信 400ms+ 粗採樣 |
| B 組（抑止 vending）後 Home／Reading 可維持 | 搶焦 ≠ 進程硬殺 | 可在 B 組做 UI／API 短窗 |
| B 組仍立刻進程死亡、無業務 Activity | 可能另有硬閘 | 另查 kill reason／logcat，勿硬套本表 |
| 靜態有 Pairip LicenseActivity 但 focus 從未見到 | **未觀測 hop** | 標 unobserved；勿寫進 validated chain |

## 決策表

| 條件 | 行動 |
| --- | --- |
| 只需證明「會被丟去 Play」 | Condition A focus 表 + 截圖即可 |
| 需無帳號驗證業務 UI／開始 capture | Condition B 抑止搶焦後再測（標明非 license crack） |
| 靜態 triage script exit 0 | 只當 L1；必須 L2 |
| 使用者要改包／hook 過授權 | 拒絕；改做 A/B 觀測或正式 Play／測試建置 |

## Anti-patterns

- 用 L1 Pairip 字符串推「已證實」中間頁。  
- 把 B 組成功寫成「繞過 Pairip／LVL」。  
- 用「必須先登帳號」當唯一阻塞，而未做 B 組對照。

## 相關

- HOW TO DO：`analysis/apk/workflows/cold-start-play-focus-ab.md`
- Evidence-first：本檔是 cold-start UI 路由；流量主線仍見 `analysis/apk/traffic-triage.md`
