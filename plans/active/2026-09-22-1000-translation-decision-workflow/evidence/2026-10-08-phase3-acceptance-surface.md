# Phase 3 acceptance surface — scoped consumer

## 決策

採用 [dogfood acceptance adapter](../../../workflow/translation/adapters/dogfood-acceptance.md)
的 Reference、Semantic、Acceptance、Target Lexical 分責。保留既有 lifecycle、
Core registry 與 locale manifestations；不新增 phase、Agent 或日文 workflow。
依敘事重要性宣告必要 assertions，不強迫每句填滿重型 semantic schema。

## 消費端證據

Scoped consumer 比較獨立覆核提供的概念／關係 assertions，不從 raw text
假裝推導語意。語序不同不構成失敗；已宣告動作、受事、否定、模態、時間或
數量變更，以及無來源擴張，會保留逐項原因。未知 reference 不建立固定譯名。
Review 綁定 source、candidate、locale、reviewer 與 evidence refs；舊候選、
不同 locale、自評或缺 provenance 不可關閉必要 gates。

Mechanical 與 lexical 分欄：shared-script 詞彙風險不冒充 script illegality，
合法漢字姓名不因非片假名而一律失敗。內容 acceptance 投影還需要獨立的
canonical Finality decision；部分 assertion pass 不能自關完整 semantic gate。
Timing／visual 保持 consumer-owned 未驗狀態，不由 text review 自動提升。

凍結 local actor 輸出唯讀重放揭示：既有機械 subset 的 accepted 標籤，
不能取代缺少的 identity／semantic／lexical acceptance evidence。
Child native exit 與候選 quality 分開；原始 baseline 未改動。
Consumer 負例與既有回歸驗證通過，不等於語意 producer 或影片已驗收。

唯讀 per-series Reference consumer 已接入 scoped replay：明確 mention 依本片
角色證據查詢；候選權重、來源身份與目標譯名覆核分開。缺少實際角色資料時
保留 unresolved，不建立固定詞表。既有接受名稱須綁定實體、當前來源名稱、
locale 與獨立覆核；alias 歧義與名稱衝突不採 first-match。此接線只驗讀取與
證據保留能力，尚未建立自動身份 producer 或更新 production 發布路徑。

## 獨立譯名覆核（延續）

在本片 `series_cast`／review export 接線之後，獨立覆核確認：

- 有 OCR／ASR 呼語證據的個人名，可 accept 單一 locale 譯名進 established
- 衝突譯名與拉丁殘留必須 reject，不得因「看起來像」或 en 同用而接受
- 姓＋職稱稱謂表面（無真名）全部 defer：identity 維持 unresolved，
  established_names 不增加

凍結重放後 established 名稱數等於獨立 accept 數；content／artifact Finality
仍可 blocked。詳見同目錄
[`2026-10-08-independent-name-review-binding.md`](2026-10-08-independent-name-review-binding.md)。

## 未完成邊界

目前是 Phase 3 dogfood 的 acceptance projection，尚未取代產品發布路徑，
未新增自動角色解析或可靠自然語言 semantic producer。稱謂真名與必要
semantic assertions 的獨立覆核、代表影片 timing／visual 仍開放。
Canonical TDR、語意與整批成片驗收仍開放。具體來源、測試、run 與統計留在專案。
