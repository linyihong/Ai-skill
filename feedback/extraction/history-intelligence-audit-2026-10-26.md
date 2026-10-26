# 歷史回饋 intelligence 接線盤點（2026-10-26）

## 結論與範圍

對 `feedback/history/**/*.md` 排除 README 共 **384** 份做全量結構與直接連結盤點。修改前 12 份有 Intelligence Review，372 份缺欄位；本輪另讀三條來源正文並補承接，修改後 15 份有決定、369 份仍缺審查。**沒有宣稱這 369 份都應新建 intelligence，也沒有宣稱已逐條完成語意審查。**

可追溯清單：[384 份逐檔 CSV](2026-10-26-history-intelligence-inventory.csv)。每列保留來源、domain、review／target、intelligence 直接入／出連結與 triage。清單是本次快照，非另一份 lifecycle 權威。

| 修改後結構訊號 | 份數 | 判讀 |
| --- | --- | --- |
| review-recorded | 15 | 已有明確決定；其中三份為本輪補齊，其餘沿用前輪 |
| link-observed-needs-review | 7 | 已見 intelligence 直接連結，不能當作完全未抽取；仍須讀正文與補正式決定 |
| no-direct-link-needs-review | 362 | 未觀察到直接連結，列待審；可能已被間接承接或僅屬操作方法 |

## 本輪確認與修補

三條來源原文皆標 candidate，且以「重複後才推廣」保留在 history；判斷內容其實已符合 extraction 的 failure prevention／trade-off 門檻。缺第二專案實測應限制成熟度，不應阻止抽成候選判斷。

| 來源 | 承接決定 | 原因 |
| --- | --- | --- |
| [有效突變層與範本隔離](../history/development-guidance/common/2026-10-15_190000-mutate-the-layer-that-decides-and-reset-the-template.md) | existing | 判斷測試結果是否真的由當次有效輸入引起 |
| [Host 組合行為](../history/development-guidance/common/2026-10-06_090000-pipeline-behaviour-needs-host-level-test.md) | existing | 元件契約的證據不能覆蓋組合順序與副作用 |
| [發佈關卡涵蓋分支不變式](../history/development-guidance/common/2026-09-04_135746-gate-that-runs-only-at-commit-is-one-bypass-from-gone.md) | existing | 單次提交抽查不能證明發佈時的分支狀態 |

三者共用 [evidence-chain-validation](../../intelligence/engineering/execution/validation-reasoning/evidence-chain-validation.md)，補決策、反例、可推翻驗證與雙向來源。沒有照搬來源專案實作，也未提升為 validated。既有父 README 已有目標連結；另補 registry、index、summary、graph 供 discovery／RAG 查回。

## 機械審查與歷史邊界

以 `kge diagnose --all-feedback --root .` 執行既有 closure gate：修改前 **365** 個 error paths，修改後 **362**。每個 path 可能只回報第一個缺失，不能把 error 數當作缺欄位總數。三個 touched lessons 的增量 `kge validate` 通過。

384 與全量 gate error／pass 母體不同：舊 flat/root 路徑不一定符合目前 lesson path matcher。CSV 保留所有 384 份，不以 gate 沒報錯將舊路徑當合規。現有 gate 仍採 on-touch 與顯式全量盤點；本次不豁免、不自動填 PASS、不將舊資料全部升級，也未改變 gate 執行語意。

## 後續審查順序

369 份缺正式 review 是明列待辦，不是本輪已完成的語意修復：

1. 先讀七條已有直接 intelligence 連結的來源及目標，確認是否為相同判斷，再選 existing 或其他 disposition。
2. 再按機制群聚其餘 362 份，優先處理 failure prevention、trade-off 與高重複成本；以目標既有內容去重。
3. 每條實際修改時套用 [lesson-review](lesson-review.md) 與既有 on-touch 機械鎖；只屬操作的選 method-only，證據／邊界不足的寫具體 deferred revisit。
4. 重新執行全量診斷與接線盤點；不要把相同專案多條觀察當成獨立重用。

## 方法與驗證界限

- 直接連結盤點：解析 Markdown inline link、移除 fragment、依所在目錄正規化為 repo path；來源為所有非 README history Markdown，目標為所有非 README intelligence Markdown。只有可解析目標才列出。未掃描 reference-style、裸文字或多跳關係，因此無連結只表示「本次未觀察到」。
- 語意來源：本輪三條 lesson、既有 evidence-chain 與 extraction／review 契約；未閱讀全部 384 份正文，未重跑來源專案。
- 依賴：feedback 規則／category 與 domain README、目標父 README、linked updates、extraction pipeline、runtime／knowledge 入口已檢查。Tool mirror sync 不適用，採 reference-only。
- 驗證：增量 KGE、全量 KGE 差異、runtime compile／refresh／validate、RAG 反查與 touched Markdown links；全量歷史紅燈保留為盤點結果，不能描述成全庫綠燈。
