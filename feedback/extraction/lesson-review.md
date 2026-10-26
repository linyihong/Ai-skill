# 回饋抽象化與 intelligence 承接流程

執行契約與欄位定義以 [feedback-lessons.yaml](../../enforcement/feedback-lessons.yaml) 為準；寫作入口是 [feedback-lessons.md](../feedback-lessons.md)。本頁說明如何執行既有 extraction pipeline，不建立另一套 lifecycle。

## 每次新增或修改 lesson

1. 讀來源內容與鄰近 intelligence，區分操作方法、判斷取捨、專案事實與規則。
2. 寫 `Transfer Contract`：輸入／前置條件、原因機制、決策、可推翻的驗證、反例與另一個可攜情境。僅寫 PASS 或換掉專案名稱不足以完成。
3. 寫 `Intelligence Review`，依下表決定承接。不是每條 lesson 都要另建 atom，但每條都要有決定。
4. 需要新 atom 時可先抽為 candidate-intelligence；HOW TO THINK 留 intelligence，HOW TO DO 留 analysis／workflow。不以缺乏第二次實測為理由漏掉候選判斷，也不把抽取當成 validated。
5. 同一提交補來源／目標雙向連結、目標父層 README、必要 metadata／route／summary／graph，執行 compile／refresh／KGE 與反向查回。
6. 缺條件時保持 candidate 或明確 deferred；禁止宣稱已抽取卻只留下不存在或未 staged 的目標。

| Decision | 何時選 | 機械要求 |
| --- | --- | --- |
| `extract` | 新的可攜判斷，現有 atom 尚未承接 | intelligence 下的具體 Markdown、id/type/layer/status metadata、反向來源、父 README、registry 與 Discovery summary；來源 candidate 時新 atom 也須 candidate |
| `existing` | 已有相同判斷 | 指向既有 intelligence 具體文件並補來源反向連結與必要內容；不建重複 atom |
| `method-only` | 增量主要是操作或既有規則 | 指向 analysis／workflow／enforcement 的具體文件，保留來源與父 README |
| `deferred` | 邊界或證據不足，仍可留下通用候選觀察 | Reason、Target: none 與可觸發的 Revisit 條件；不要只寫「以後再說」 |

若結果是 project-only，留在專案文件；不可用這個標籤讓專案特例通過共用 history 的抽象化閘。

## 寫法

欄位使用單行 `- Key: value`，不重複 key；Target／Reference 是 repository-relative `.md` 路徑，不用 `<placeholder>` 或 Markdown link 包裝。

```text
#### Transfer Contract
- Inputs: <已知輸入與前置條件>
- Mechanism: <原因機制>
- Decision: <什麼條件下選哪條路>
- Verification: <如何推翻或證實>
- Counterexample: <不適用情境>
- Transfer case: <不同情境如何套用；不是獨立實測宣告>

#### Intelligence Review
- Decision: extract|existing|method-only|deferred
- Reason: <選擇原因>
- Target: <repo-relative.md；deferred 時為 none>
- Revisit: <僅 deferred 必填的觸發條件>
- Discovery: <僅 extract 必填的 knowledge/summaries/*.md>
```

`validated/promoted` 的 `Reuse Evidence` 另須有 `Independent context`、`Result`、`Reference`。Reference 指向另一本已追蹤的去敏證據紀錄；可以記錄受控外部證據的來源，不把原始私人日誌放進共用庫。格式與路徑可機械驗證，是否真的獨立仍須正文審查。

## 機械鎖範圍

- commit-msg 與 pre-push 的既有 `feedback_lesson_closure` 驗證以上欄位與接線。
- 檢查讀 git index 的 lesson、target、README 與 evidence；工作目錄或 untracked 檔案的修復不能充數。
- `kge validate --all-feedback` 是既有資料的顯式修復盤點；平常只擋新增／修改條目，不把全部舊 history 一次判成已修好。
- 機械鎖保證有決策、必要內容及可解析承接，不保證作者的語意判斷正確；不會自動 promote 或製造重用證據。

本輪實際分類見 [近期判斷表](recent-review-map.md)。
