# 跨目標採集紀律

本文件定義授權標的上的 **操作紀律**：如何採、何時停、交付什麼才算可交接。主張如何標 observation／inference／unknown 見 [`../evidence-contract.md`](../evidence-contract.md)。

抽象自授權業務專案採集 SOP；範例用語為 runtime dump／package extract／saved capture，不綁定單一產品。

## 1. Probe before attach

在 hook、dump、或開長窗捕捉前，先列可觀測根：視窗／場景／模組／進程顯示名。

| 規則 | 說明 |
| --- | --- |
| 先列名再 attach | 用 list／probe／inventory 取得實際 runtime 名 |
| 靜態名 ≠ runtime 名 | package／prefab／資源檔名不得直接當 attach 目標 |
| 進程識別用顯示事實 | 包名、顯示名、PID 可能不一致；以當下可觀測事實為準 |

## 2. Platform consistency

同一功能的 runtime 證據、靜態包體、抽出資源必須來自**同一平台／同一發行物族**。

- 換平台（例如手機包 ↔ PC 包）＝整套重採，不得混算幾何或資源路徑。
- 懷疑「解碼壞了」前，先用同平台數值／路徑對照，排除跨平台混用。

## 3. Transient full-dump

短命 UI、一次性動畫、短暫對話框、隨即消失的狀態：

1. 立刻 attach／開啟捕捉視窗
2. **一次全讀**（幾何、組件、文字、可見資源引用、相關狀態）
3. **先落盤**
4. 確認檔案寫入成功
5. 才報告或開始分析

界面不會等分析者；「來不及了」通常是採集節奏失敗，不是目標不存在。

## 4. Two-phase capture

| 階段 | 做什麼 | 不做什麼 |
| --- | --- | --- |
| A 採集 | 全量落盤（字串／樹／多通道 raw） | 依「命中率」預過濾 |
| B 分析 | 在已落盤資料上過濾、對照、分主張 | 回頭要求目標再停一次才補采（除非另開授權窗） |

採集期過濾掉的資料是**永久損失**；命中率不是預過濾的理由。

## 5. Static + runtime dual source

殼層可能由代碼在 runtime 組裝，靜態樹遠小於 runtime 樹。

| 要回答的問題 | 優先來源 |
| --- | --- |
| 內容／版面／資源綁定 | 靜態包或抽出資源 |
| 殼層／動態掛載／實際可見樹 | runtime dump |
| 兩邊衝突 | 分寫 observation；缺一邊標 residual unknown |

## 6. Deliver relations（不是孤立檔）

交接前必問：下游能否用這些檔**拼出對應關係**？

最低交付物類型（依目標選適用者）：

- 節點／符號 ↔ 資源或資產路徑
- 請求／幀 ↔ 畫面或功能 seed
- 靜態結構 ↔ runtime 觀測

只有資產清單、沒有映射關係 → **未 ready**。

## 7. Evidence folder contract + ready gate

業務專案內建議固定證據目錄（名稱可本地化，職責固定），例如：

```text
evidence/<feature>/
  markdowns/     # 筆記、REFERENCE_MAP、結論
  dumps/         # runtime／靜態樹
  protocols/     # 抓包摘要、inventory（已去敏）
  assets/        # 抽出資源（僅引用相關）
```

| Ready 條件 | 失敗時 |
| --- | --- |
| 目錄齊、映射檔存在、主張已分級 | 標 `not_ready`；不得宣稱 handoff 完成 |
| 無關檔已移出或標 ignored | 目錄本身是交付物的一部分 |

## 8. Search hygiene

- 搜尋符號／字串時**排除本輪產物**，避免自我匹配上一輪 trigger／fixture。
- 記憶／字串針必須夠特異（帶產品或功能前綴）；裸通用名會污染跨 app 命中。
- 勿因省事排除整個目標樹，導致漏掃相關靜態來源。

## 9. Impl-retry ceiling → capture gap

| 規則 | 說明 |
| --- | --- |
| 同症狀實作上限 | 建議 ≤2 次；第 3 次停止實作 |
| 轉採集 | 查缺哪一層證據（runtime／靜態／協議／映射） |
| 禁頂替 | 不得用截圖猜測、自製資源、或未證實綁定頂呈現 |

缺件時標 `needed_authority`，不要盲做。

## 10. Human recognition export

不確定的視覺／音訊資產：匯出給**人**辨認，寫回標註。禁止模型反復讀大圖燒 token 當「查證」。

## 11. Numeric visual QA

可用可覆核指標（亮度／覆蓋率／區塊形狀指紋等）替代「看起來對不對」。指標是 observation；「與官方一致」若無對照基線則為 inference 或 unknown。

## 12. Mistake → reusable writeback

- 一個可檢索症狀只允許在流程裡踩一次；第二次同坑＝流程失敗。
- 寫回 reusable lesson 前對照**已知正確樣本**。
- 禁止籠統「不可行／不要試」擋掉仍可用的權威路徑；改寫具體條件與替代權威。

去敏與 reusable 邊界見 [`../../../enforcement/sanitization.md`](../../../enforcement/sanitization.md)、[`../../../enforcement/reusable-guidance-boundary.md`](../../../enforcement/reusable-guidance-boundary.md)。

## 最小檢查

- [ ] 已 probe 再 attach
- [ ] 平台一致
- [ ] 瞬態已先落盤（若適用）
- [ ] 採集／分析兩階段未混用
- [ ] 交付含映射關係
- [ ] 實作重試未超限，或已轉採集缺口

← [collection/](README.md)
