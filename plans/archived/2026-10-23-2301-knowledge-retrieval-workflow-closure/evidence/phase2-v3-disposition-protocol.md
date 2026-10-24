# v3 Disposition — 單一變項的候選選擇重播

本協議在本輪 D1 評分前固定。方法由使用者基於已曝光的 O02 / O04 診斷提出，不是新獨立測試。沿用六份 raw responses、完整 ranking、request／pool／opaque mapping、原 gold、原 P0/P1 與 2000-rune proxy，**不重叫模型、不改排序／權重／top-k／摘要政策**。

## 處置規則 D1

| 原 assessment | 本次動作 | 不能推定 |
| --- | --- | --- |
| direct | 依原 ranking 進入待讀回清單；在候選層 selected-context preview 中可保留指標與摘要，但明列待 workflow 驗證 | 不等於來源可信、適用已確認或可採用 |
| background | 預設不填主要上下文，保留理由與可選背景清單 | 原 reason 不是具體 background-use 授權 |
| not-applicable | 本次不選，保留理由 | 不全域封鎖／貶抑該來源 |
| unknown | 進待確認清單，必要時另補讀；本輪不自動補讀／升格 | 不等於已確認相關 |
| 缺 assessment | 按 unknown 處置，明列 not assessed | 不沿用前一筆或因高 rank 自動 direct |

本輪沒有指定背景用途，因此 background-use overrides 固定為空。所有 assessment 都來自既有 raw，不按 expected／confuser 改標記。未評完整 pool 的候選不能因未被評估而消失：每筆 source 均有 disposition/action 與 reason，含 explicit / default 標記。

D1 保持原 source order，只篩 direct 再套**原** P0/P1；若沒有 direct，就輸出空 preview 和待確認清單，不自動退回 top1 或填背景。分類錯誤／缺理由／外來 ID 依原 parser controls 拒絕，不偷偷修復。

## 分開觀測的層

1. retrieved：原 pool 原樣保留；預標 confuser 在池內仍按原值記錄，不因 D1 被刪掉而改寫檢索分數。
2. D0：原 ranking-only P0/P1 結果保持不動，是本輪對照。
3. readback queue：D1 的 direct 來源，保留 source identity／revision／checksum 與 agent 理由。
4. selected context preview：D1 + 原 P0/P1 的候選層組裝，計 hits／MRR／confuser／runes／budget drops；仍是 pending-workflow-verification，並非已確認可用上下文。
5. verified context／adopted：本輪 not evaluated。原 Phase 3 的 source-read／scope／supersession／required_read／blocker 不被此標記取代。

只有 D1 disposition admission 是主要對照變項。另作 direct source readback audit：從固定 corpus Git blob 讀回全文並核對 checksum，再核對當前 canonical 檔案是否改變；外層已知 gold，內容判讀只能算知情診斷。此 audit 不偷偷重新篩 D1、改本次 selected-preview 分數；不把 checksum 相同當成語意適用已確認。

## 本輪驗收與限制

- 逐項核對所有 raw／request／pool hashes、原排序、metadata 與 P0/P1。新報告另存，不回改上一輪 5/5、2/4 或 retrieved confuser。
- 同六題比較：5 個指定正向來源、4 題帶預標 confuser；O06 無完整 gold，非 gold 不自動算不相關。
- D1 selected-preview 目標：正向不退步、confuser 0/4、≤2000 runes；即使達成，也只通過這個曝光資料的候選選擇 slice。
- direct queue 與已選 source 的 version/readback、unknown、背景用途、理由與 metadata 不變性皆可覆核。空結果、不足資訊或待確認不得當作 0 錯誤或已採用。
- 全 26 題語意處置、自然無答案、資訊充分性、真實 workflow consumer／採用鏈、實際 tokens、時間控制與獨立泛化仍未驗收；Phase 2 不因這個子實驗通過而完成。

不增加模型呼叫，本輪離線處置耗時不修復上一輪 runner 的時間觀察缺口。額外 API／向量 DB／正式 workflow 接線皆非本輪交付。
