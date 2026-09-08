> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Parsed symbol DTO ID may be a one-character strip code

Status: validated

#### One-line Summary

Live `SymbolData.ID` 可能是 **長度 1、大小寫敏感的 reel-strip 字元**，不必等於 Unity resource / prefab 名。座標欄若是 valuetype，要按 inline 讀，陣列還可能含畫面外的額外 row。

#### Human Explanation

掛 wrapper Parse 後讀巢狀 symbol 陣列時，先印 `string_length`。若幾乎都是 `L1`，把 ID 當 strip alphabet，再用同一次 spin 的停輪畫面做格子對位。不要假設 ID 等於 `high_1` 這類美術名。Point 若 `class_is_valuetype`，從父物件 offset 連續讀兩個 int，不要當 managed pointer。

不要 dump 下注、賠付、帳號或原始 XML。

#### Trigger

- 已有美術 resource 名，要驗證是否等於 parsed symbol ID。
- ID 讀出來像亂碼、err，或 Point pointer 讀取拋錯。

#### Evidence

- Tool: Frida field offsets + `il2cpp_string_length` / `il2cpp_string_chars`; valuetype Point inline ints; same-spin screenshot after Parse.
- Sanitized observation: IDs were length-1 and case-sensitive; Point was valuetype; array included negative and extra rows around the visible window.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` symbol notes

#### Generalized Lesson

1. Protocol symbol ID ≠ client resource ID，直到對位證明。
2. 長度 1 的 ID 優先當 strip code。
3. Valuetype 座標不要當 reference type 解。
4. 同一 round 才把格子對上截圖；晚幾秒的畫面可能是下一轉。

#### Agent Action

讀 ID 長度；valuetype 則 inline；只在 Parse 後短等待截停輪畫面。

#### Goal / Action / Validation

- Goal: 證明 parsed ID 與畫面符號的對應層。
- Action: ID+reel+row + 同轉截圖。
- Validation: 可見格子與 ID 網格一致，或明確標未出現的符號。

#### Applies When

- 授權 IL2CPP；slots/reel 類 parsed symbol 陣列。

#### Does Not Apply When

- ID 已證實是 resource 字串（照實記錄）。
- 沒有畫面可對位。

#### Validation

One live round; names/IDs of strip alphabet only; no payloads.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
