# 經驗查找決策紀錄

Consumer：[`sd-validation`](../validation-retrieval.md)。本紀錄存專案驗證文件；填入可覆核事實，不保存完整對話或秘密。

## 任務與觸發

- 任務／claim／風險：
- workflow／route／required_read（實際讀取紀錄）：
- 當前 blocker 與 owner gate：
- 動作：lookup / requery / skip / disabled-fallback
- 事件、前提／版本變動或略過理由（含先前紀錄引用）：
- query／filters／prepare、watch、complete trace：
- 選擇方式：current-agent / existing-runner；raw 保留期限與去敏入口：
- pool／unassessed／explicit unknown／deferred；空輸出處置：

## 每個準備使用的來源

| 欄位 | 紀錄 |
| --- | --- |
| Canonical source／need_to_read | |
| 候選版本與 checksum | |
| 實際回讀版本與 checksum／段落或證據入口 | |
| 採用前再次核對的版本與 checksum | |
| domain／project／platform／前提 | 已確認事實、缺口；不可只填「符合」 |
| 證據有效性／supersession／lifecycle／authority | 明確宣告來源；缺值 unknown |
| 決策 | adopt / reject / unknown |
| 理由與本次 workflow 用途 | |
| 版本變更的差異與重驗結果 | 不變也須記核對；重驗中再變則停止沿用 |
| 執行後結果／支持的 claim 邊界 | 尚未執行填 not evaluated |

## 預算及 gate 保持

- 實際 message artifact／SHA／納入範圍：
- model／tokenizer／version／tokens／上限2000：
- 無法取得的欄位與 admission 處置：
- 正文補讀成本與其他任務 context（另報）：
- required_read／activation／blocker 原 gate 是否保持及證據：
- 未解事項、下一個取證動作與 owner：

找到來源不表示適用；採用參考不表示規則升級；未知不填零。
