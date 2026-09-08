> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Nested DTOs may have their own Parse/ParseAll beside the wrapper Parse

Status: validated

#### One-line Summary

Round wrapper 的 `Parse` 旁邊，巢狀 DTO 常有自己的 **`Parse` / `ParseAll` / `ParseSub*`**；它們多半在 wrapper Parse **前／後**觸發，而不是在 `Process*` apply 裡。

#### Human Explanation

ctor 時間桶之後，應列出巢狀型別上非 getter 的 fill 方法（`Parse*`、`Fill*`、`CopyFrom`）並用同一 PRE/IN/POST/APPLY 窗口掛上。某型別沒有 `Parse*` 是有效發現（可能只有 ctor + field store）。不要 dump 參數。

#### Trigger

- Nested `.ctor` 已確認在 Parse 前／後。
- 需要方法名，不是再比物件 identity。

#### Evidence

- Tool: Frida enumerate nested-class methods matching Parse/Fill/Load/Create/Copy/Add; hook in PRE/IN/POST/APPLY; names only.
- Sanitized observation: live set included `Parse`, `ParseAll`, and a `ParseSub*` helper around wrapper Parse; APPLY none; at least one nested apply DTO had no `Parse*` in metadata.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Wrapper `Parse` 不是唯一 XML 入口。
2. `ParseAll` 常配座標／獎勵陣列。
3. 缺 `Parse*` 的 DTO 不要硬找。

#### Agent Action

nested ctor 之後 enumerate `Parse*`；用同一時間桶驗證。

#### Goal / Action / Validation

- Goal: 命名巢狀 XML fill 方法。
- Action: 方法名清單 + PRE/IN/POST/APPLY 命中。
- Validation: 至少一個 `Parse*` 在 PRE 或 POST。

#### Applies When

- 授權 IL2CPP；wrapper Parse + nested apply DTO。

#### Does Not Apply When

- 巢狀填值全部 inline 在 wrapper Parse（IN 會看到）。
- 純靜態未掛 live。

#### Validation

One round; method names only.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
