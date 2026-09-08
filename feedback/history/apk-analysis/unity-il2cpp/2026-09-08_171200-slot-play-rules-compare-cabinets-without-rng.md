> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Slot play-rules should compare cabinets without RNG or paytable

Status: validated

#### One-line Summary

每個 slot cabinet 套件除了 symbol fixture 外，還要有一份固定欄位的 client-visible play-rules（win mode、線數、符號角色、feature 狀態、result DTO 名），用來比較遊戲差異；不要把停輪算法或數字賠付表寫進去。

#### Human Explanation

Symbol JSON 證明某一轉停了什麼。玩法比較需要另一層：這台是 paylines 還是 all-ways、線數、wild/scatter 角色、結果是 server DTO 套到畫面。這兩層分開後，後續 cabinet 可以對表，而不會把 HUD 筆記散落在 protocol 雜記裡。

常見誤判是把「可以分析玩法」做成還原 RNG、帶權重轉輪或 live 賠付數字。那些通常是 server-authoritative，也不該進可分享套件。另一個誤判是用通用 25 線模板填幾何：沒有 client config 或 sanitized reward 點位就保持 `hud-count-only`。

#### Trigger

- 已有 cabinet 套件（art + parsed spin），使用者要玩法／公式。
- 準備第二台遊戲，需要和第一台對差異。
- HUD 同時出現 25 線、50 線、ALL WAYS。

#### Evidence

- Tool: cabinet package schema, play-rules JSON, structural validator.
- Sanitized observation: required keys include winMode, paylineCount, symbol roles, feature status, and `resultApply.authority=server-result-dto`; numeric pay and RNG listed as out of scope.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/play-rules.md` and `<cabinet-id>/play-rules.json`.

#### Generalized Lesson

1. Play-rules 是每個 cabinet 套件的必備檔，不是 README 裡的一段附註。
2. 比較用固定 keys：grid、winMode、paylineCount、featureSet、result-apply 欄位名。
3. 線路徑只從 client config 或 sanitized reward 點位複製；HUD 線數 ≠ 已證明幾何。
4. Composition / paytable 美術不可覆蓋 live 停輪（例如美術把 wild 畫在外側，live 仍可能在中間 reel）。
5. Scatter 觸發次數未證明前保持 null；單顆 scatter 未進 bonus sink 只能記 observed，不能當公式。
6. Validator 應拒絕缺少 play-rules 或 `authority` 不是 server result DTO 的套件。

#### Agent Action

新 cabinet 在第一轉 fixture 前就建立 play-rules stub。問「算法／公式」時先填這層，明確拒絕 RNG 與數字賠表。第二台遊戲用同一 JSON 欄位做 diff。

#### Goal / Action / Validation

- Goal: 後續 session 能對兩台遊戲看出結構差異，而不誤以為已還原出獎算法。
- Action: schema + per-cabinet JSON/MD + package validator。
- Validation or reference source: validator requires play-rules.json; all-ways 必須 `paylineCount=null`。

#### Applies When

- 授權分析 in-app slots，且結果由 server DTO 驅動 client 顯示。
- 需要跨 cabinet 比較，或要把玩法寫進可分享套件。

#### Does Not Apply When

- 任務只做一次性 HUD 截圖、不建套件。
- 目標是實作非官方對官方伺服器的預測或下注繞過。

#### Validation

- 套件目錄含 play-rules.md / play-rules.json。
- JSON 含 comparisonKeys 與 outOfScope。
- 沒有 live 下注／賠付數字。

#### Promotion Target

- Candidate for `workflow/apk-analysis/artifact-gates/` after a second cabinet dogfood uses the same keys.

#### Required Linked Updates

- Updated `feedback/history/apk-analysis/unity-il2cpp/README.md`.
- Updated `feedback/history/apk-analysis/README.md`.
- Project cabinet names, protocol IDs, and live counts remain only in project docs.
