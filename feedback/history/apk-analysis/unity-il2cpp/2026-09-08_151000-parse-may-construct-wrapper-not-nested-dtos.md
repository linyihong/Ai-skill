> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Parse may construct only the event wrapper, not nested DTOs

Status: validated

#### One-line Summary

XML `Parse` 進出窗口裡，**可能只 `new` 事件／結果 wrapper**（子類 ctor 會打到 base `.ctor`）；巢狀 apply DTO 的 `.ctor` 可以完全沒有。

#### Human Explanation

`set_*` none 之後，下一步常以為 Parse 會 `new` 整棵 DTO。應在 Parse 進出掛目標型別 `.ctor`（類別名 only）。若只有 wrapper、沒有 nested DTO ctor，代表填的是呼叫端已有的物件圖（搭配 field store），不是 Parse 內建樹。有的 Parse 進出 ctor 也是 none（null／reuse／提前 return）。

#### Trigger

- Parse 窗口已確認轉換 API 與無 `set_*`。
- 需要知道巢狀 DTO 是新建還是就地改。

#### Evidence

- Tool: Frida attach `.ctor` on wrapper + nested apply DTO types only while Parse is on stack; unique class names, not counts or fields.
- Sanitized observation: wrapper subclass ctor plus base ctor; nested apply DTO ctors absent; one Parse window had no hooked ctors.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. 子類 ctor 與 base ctor 會一起出現，不要當成兩個獨立物件。
2. nested DTO ctor none ≠ 沒有填值。
3. 不要 dump ctor 參數或呼叫次數。

#### Agent Action

Parse 填值在 setter none 之後掛 `.ctor` 類別名；只有 wrapper 就記「mutate existing graph」。

#### Goal / Action / Validation

- Goal: 分辨 Parse 內配置 vs 就地改既有 DTO。
- Action: Parse 進出 unique `.ctor` 類別名。
- Validation: wrapper 有、nested 無（或相反）可對照 getter 出現時機。

#### Applies When

- 授權 IL2CPP；靜態 `Parse(..., XElement, …)` 填 round/result 物件。

#### Does Not Apply When

- nested DTO `.ctor` 確實在 Parse 窗口出現（照實記錄類別名）。
- 純靜態。

#### Validation

Same Parse window; ctor class names only.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
