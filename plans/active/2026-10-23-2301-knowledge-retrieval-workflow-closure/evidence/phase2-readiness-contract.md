# Phase 2 剩餘驗收與 consumer preview 契約

此契約延續使用者對 890de637 的審閱。只建立離線 coverage／preview 原型與全資料盤點，不接 Phase 3 consumer、不新增模型判斷、不回改既有六題分數。

## Consumer 可見欄位

- 每次 preview 必須同時提供 candidate 數、assessed 數、unassessed 數、explicit unknown 數、unknown 總數、background／not-applicable／direct 數。missing assessment 與 explicit unknown 分開，兩者均不得自動進主要上下文。
- 空 preview 狀態區分：池內無候選、尚待 assessment、已評池內無 direct。任何狀態都不代表全庫「沒有相關知識」。有 direct 時仍顯示未確認數與待讀回狀態，不能因非空就隱藏 unknown。
- 每題 coverage 綁來源結果的 SHA 與完整 decision trace 入口；summary counts 不能替代 audit queue。這輪只是 preview prototype，真實 workflow 是否讀取這些欄位仍 not evaluated。
- 預設固定 P0 完整 indexed summary，保留原 direct 順序和 2000-rune results-array 邊界；不自動切 P1、不自動填 background／unknown、不按 gold 裁剪。
- results array、consumer-visible envelope 各自重算 runes。舊 2000-rune gate 只涵蓋 results array；若新增 envelope 後超過，不宣稱整體 context 在預算內，也不偷偷縮摘要。真正接線前須固定整體 budget 與實際 tokenizer。

## Phase 2 驗收順序與責任

| 項目 | 現況與下一份證據 | Owner / gate |
| --- | --- | --- |
| 全 26 題候選處置與逐筆審閱 | 只有六題 sparse assessments，最大 8 筆理由的 runner 並非全量評估；先輸出每題 coverage，未標記維持 unknown。新全量 assessment 須另版 protocol／requests／輸出 schema，不覆寫舊 raw | 檢索執行者；larrylin 或指定 reviewer 覆核；未結清不通過 |
| 資訊充分性 | source hit 不等於能回答問題；逐題列必要資訊、P0 實際保留內容、缺失／不足／unknown。外層已見 gold，須標曝光診斷 | 檢索執行者整理、reviewer 覆核；不得用 Recall 代替 |
| Intelligence／整理知識覆蓋 | 另版 source-level 比較資料；不可由 lesson 20 題宣稱全庫改善 | 索引執行者與 sd-validation owner |
| Metadata／graph／索引生命週期 | 重用 canonical 欄位及有界 graph；建立、refresh、修改／刪除／搬移、stale、rebuild、ordering 均需實際 evidence | CLI／索引執行者；未驗不得勾選主 plan |
| 成本與時間 | 分離 lookup／選擇／packing，記連續受控外層時間、失敗與 timeout；不可得 usage 明列。前輪時間缺口不能由離線新計時修復 | Runner 執行者；新 attempts 需先凍結監控／中斷契約 |
| 候選品質 | 在準備採用的版本上重跑完整固定資料與 overlap controls，保留 retrieved 與 selected 指標、未審閱數；已拒收版本不強迫補審 | 檢索執行者與 reviewer；不改 gold、不用未知當零錯誤 |
| Phase 3 接線 | Phase 2 gates 通過後，才安排 sd-validation 真實 consumer；readback、適用性、版本 race、supersession、required_read／blocker、採用理由與實際 tokens | Workflow owner；本輪 not evaluated |

完整資料仍是曝光比較集，不是獨立泛化驗證。六題通過不縮減全資料、資訊充分性、成本／時間與既有 lifecycle gate。既有 raw parser／scorer不改，此原型不增加 generated runtime surface 或正式 consumer。
