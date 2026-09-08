> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Slot spin fixtures should pair JSON, manifest, and reel crop

Status: validated

#### One-line Summary

解析 slot symbol DTO 後，每個可接受樣本應用同一 `sampleId` 綁定完整座標 JSON 與同轉 reels-only 截圖，再由 manifest 和 validator 證明 visible grid 沒有手抄漂移。

#### Human Explanation

只有 terminal grid 或只有停輪截圖，都不足以長期證明 protocol symbol ID 對應哪個圖案。可重跑的證據包需要保留完整 returned rows、明確標示 visible row range、從座標推導 row-major visible grid，並把同一轉的去敏 reel crop 放在旁邊。多次 spin 應各自成 fixture；attach-state event、duplicate callback 或無法確認時序的 screenshot 不可混成成功樣本。

完整畫面常含帳號、餘額、下注或贏分；可分享證據只保留 reels-only crop，瞬時數值 overlay 必須遮蔽並在 fixture 記錄 redaction。

#### Trigger

- 已能讀出 `SymbolData.ID` 與 reel/row 座標，要建立可供後續 session 或開發使用的樣本。
- 同一 symbol mapping 需要多轉截圖交叉驗證。
- HTML 報告、JSON fixture 和圖片開始分散，容易出現路徑或 grid 漂移。

#### Evidence

- Tool: narrow runtime probe, stopped-state screenshot, JSON fixture, package manifest, structural validator.
- Sanitized observation: accepted fixtures each covered the declared coordinate rectangle; validator re-derived visible rows from coordinates and compared them with both stored JSON and the HTML sample index.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/<cabinet-id>/`.

#### Generalized Lesson

1. 一個 spin 一個 `sampleId`；JSON 與 screenshot 必須一對一。
2. JSON 保留完整 returned row range，另外存明確 orientation 的 visible grid。
3. Manifest 集中管理 symbol map、sample paths、art-only variants 與 unknowns。
4. Screenshot 只提交 reels-only crop；原始完整圖與 log 留在受控、gitignored evidence。
5. Validator 至少檢查 unique coordinates、declared layout coverage、visible-grid derivation、symbol-map coverage、PNG path，以及 HTML grid 與 fixture 一致。
6. 無法時序配對的 capture 記為 rejected，不可為了增加樣本數而 promotion。

#### Agent Action

新 slot cabinet 第一次取得 parsed symbol array 後，先建立 package skeleton、fixture schema 與 validator，再做 bounded multi-spin capture。每次 capture 完成立刻保存同 ID 的 JSON/crop；結束前用 manifest 驅動驗證，不靠人工逐檔記憶。

#### Goal / Action / Validation

- Goal: 讓其他 session 能重建 symbol mapping，並能檢查每個對應的證據強度。
- Action: 建立 manifest + per-spin JSON + reels-only PNG + visual report。
- Validation or reference source: validator 從完整 coordinates 重算 visible grid，並確認 report 顯示資料與 fixture 相同。

#### Applies When

- Unity/IL2CPP 或其他 runtime 已能輸出 symbol ID 與 reel/row 座標。
- 畫面可在同一 spin 停輪後截取。
- 需要把 target evidence 轉成可共享的 parser/layout fixture。

#### Does Not Apply When

- 沒有可靠座標，只能辨識整張畫面。
- screenshot 無法確認與哪個 parsed event 同轉。
- 任務只需一次性 runtime triage，不需要 durable symbol mapping。

#### Validation

- Fixture schema validation 通過。
- `(ReelID, Row)` 唯一且完整覆蓋 declared layout。
- 從 `symbols` 推導的 visible grid 等於 fixture 與 HTML 顯示。
- 每個 accepted sample 的 screenshot path 存在且已人工完成去敏 review。

#### Promotion Target

- Candidate for `workflow/apk-analysis/artifact-gates/evidence-chain.md` after a second cabinet/package dogfood confirms the same contract.

#### Required Linked Updates

- Updated `feedback/history/apk-analysis/unity-il2cpp/README.md`.
- Updated `feedback/history/apk-analysis/README.md`.
- Project-specific identifiers, sample values, images, and runtime evidence remain only in project docs.
