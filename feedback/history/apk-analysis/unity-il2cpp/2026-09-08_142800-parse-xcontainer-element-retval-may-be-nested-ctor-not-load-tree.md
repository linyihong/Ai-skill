> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Parse XContainer.Element retval may be nested ctor, not Load tree

Status: validated

#### One-line Summary

`Parse` 參數上的 `XContainer.Element` 回傳值可以是 **巢狀 `XElement` ctor**，即使同窗口 Load 樹 `XElement` pool 非空也不必 EQ；`Elements`／`GetElements` 可能回傳 **iterator**，不是 `XElement`。

#### Human Explanation

知道 Parse 呼叫 `Element` 之後，下一跳是回傳物件身份。Load 樹上的 Root／Element 與 Parse 參數本身都可能不是這個回傳值。應同時 pool：Load 樹節點、ctor `this`、Parse args。`Element` 可為 null。不要 dump `XName`。

#### Trigger

- Parse 已證實走 `XContainer.Element`／`Elements`。
- 需要判斷子節點是 Load 樹還是又一次 copy。

#### Evidence

- Tool: Frida; Parse-arg `this` only; return-register class + EQ vs tree/ctor/arg pools.
- Sanitized observation: `Element` hit ctor pool and missed Load-tree pool (tree pool nonempty). `Elements`/`GetElements` were compiler-generated iterators (`NE`). Some `Element` returns were null.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. `Element` 回傳 identity ≠ Load 樹 identity。
2. `Elements`／`GetElements` 先記 iterator class，再掛 `MoveNext`／`Current`（若需要）。
3. null `Element` 不是 hook 失敗。
4. 不要 dump `XName`／XML。

#### Agent Action

Parse 期間只在 args 上比 `Element` 的 return register 與 tree／ctor pool。

#### Goal / Action / Validation

- Goal: 避免把 `Element()` 寫成 Load 樹 hop。
- Action: 三 pool EQ/NE + 回傳 class。
- Validation: ctor hit + tree miss，或明確 iterator class。

#### Applies When

- 授權 IL2CPP；XML Load 後 `Parse(XElement, …)` 再 `XContainer.Element`。

#### Does Not Apply When

- `Element` 回傳確實 EQ Load 樹（記樹 hop）。
- 純靜態。

#### Validation

Same Decrypt window; Load-tree pool nonempty; Parse-time `Element` tagged ctor not tree.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
