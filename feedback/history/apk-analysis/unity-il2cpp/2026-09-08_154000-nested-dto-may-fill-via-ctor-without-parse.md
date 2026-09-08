> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - A nested DTO may have no Parse* and fill only via ctor plus field stores

Status: validated

#### One-line Summary

沒有 `Parse*` 的巢狀 apply DTO，可能只暴露 **`.ctor` + 幾個非 XML 的 helper**；Parse 前後 live 也可能只有 `.ctor`，填值仍靠 field store。

#### Human Explanation

`Parse*` 掃描 miss 之後，應列出該 class 全部非 `get_`/`set_` 方法再掛同一時間桶。Helper 名叫 `GetReward*` 不代表 Parse 窗口會呼叫。不要 dump helper 回傳。

#### Trigger

- Nested `Parse*` 清單缺某一 apply DTO。
- 仍看到該型別 `.ctor` 在 PRE/POST。

#### Evidence

- Tool: Frida enumerate all non-get/set methods on the DTO; hook PRE/IN/POST/APPLY; names only.
- Sanitized observation: method set was `.ctor` plus reward-query helpers; live around Parse was `.ctor` only; APPLY none.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. 缺 `Parse*` 要改列全方法，不要當 hook 失敗。
2. `Get*` helper 可存在卻不在 fill 窗口。
3. 不要 dump helper 回傳值。

#### Agent Action

`Parse*` miss 時 enumerate 該 class 非 accessor 方法並用同一時間桶。

#### Goal / Action / Validation

- Goal: 確認無 Parse 的 DTO 如何填。
- Action: 方法名清單 + PRE/IN/POST 命中。
- Validation: LIST 有 `.ctor`；live 可為僅 `.ctor`。

#### Applies When

- 授權 IL2CPP；部分巢狀 DTO 有 ParseAll、另一部分沒有。

#### Does Not Apply When

- 該 DTO 確實有 live `Parse*`（照實記錄）。
- 純靜態。

#### Validation

One round; method names only.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
