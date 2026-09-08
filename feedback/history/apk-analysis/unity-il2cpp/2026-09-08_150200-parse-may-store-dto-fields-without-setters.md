> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Parse may store DTO fields without calling set_*

Status: validated

#### One-line Summary

XML `Parse` 填 DTO 時，**BCL 轉換 API 出現不代表會走 property `set_*`**；目標型別可能沒有 setter，或現有 setter 在 Parse 窗口完全不觸發。

#### Human Explanation

IL2CPP 常把欄位編成直接 store。掛 `set_*` 得到 none 時，不要當成「沒填值」或 hook 失敗——先列出該 class 實際有哪些 `set_*`，再對照 Parse 進出。none + 已知轉換 API = 欄位寫入，不是再追 getter/setter 鏈。

#### Trigger

- Parse 窗口已看到 `Enum.Parse` / `Int*.TryParse` / `Convert.To*`。
- 對 apply DTO 掛 `set_*` 結果為 none。

#### Evidence

- Tool: Frida enumerate `set_*` on apply DTO classes; attach only while Parse is on stack; log method names, not arguments.
- Sanitized observation: some apply DTOs expose **no** `set_*`. Sibling types may list a few setters that still do not fire during Parse.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. 先 enumerate 該 class 的 `set_*` 集合，再解讀 none。
2. 有轉換、無 setter = 直接 field store。
3. 不要 dump setter 參數。

#### Agent Action

Parse 填值下一步掛 `set_*` 名稱；none 時停止追 property，改記 field store 或靜態欄位清單。

#### Goal / Action / Validation

- Goal: 分辨 property fill vs field store。
- Action: Parse 窗口 unique `set_*` 名稱（可為 none）。
- Validation: READY 列出實際存在的 setter；Parse 進出對照。

#### Applies When

- 授權 IL2CPP；XML Parse 填 managed DTO。

#### Does Not Apply When

- 明顯走 `set_*`（照實記錄名稱）。
- 純靜態、未掛 Parse 進出。

#### Validation

Same Parse window; setter names only; no value dumps.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
