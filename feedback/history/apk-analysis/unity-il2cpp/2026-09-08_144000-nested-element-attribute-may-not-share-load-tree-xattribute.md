> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Nested Element Attribute may not share Load-tree XAttribute

Status: validated

#### One-line Summary

Parse **參數 wrapper** 上的 `Attribute` 可以與 Load 樹共用 `XAttribute`，但 **`Element`／`Current` 子節點**上的 `Attribute` 不必共用（樹 pool 非空仍 NE）。

#### Human Explanation

不要把「某一層 Attribute EQ 樹」推到整棵 copy。子節點若是新 ctor，屬性物件也可能是新的。Parse 期間子節點上也可能看不到 `get_Value`。不要 dump `XName` 或值。

#### Trigger

- Wrapper `Attribute` 已 EQ Load-tree `XAttribute`。
- 下一跳是子 `XElement` 的 `Attribute` identity。

#### Evidence

- Tool: Frida; Parse-time `Element`/`get_Current` kids vs pre-Parse Load-tree Attribute pool; class + EQ/NE.
- Sanitized observation: child `Attribute` returned `XAttribute` with tree=NE while tree pool was nonempty; some Attribute returns were null; `get_Value` not observed on those children before Parse returned.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Wrapper 層 Attribute 共用 ≠ 子節點 Attribute 共用。
2. 每一層 copy 都要分開 pool。
3. 子節點 `get_Value` none 不代表沒有屬性；可能用別的讀取或延後讀。
4. 不要 dump `XName`／XML。

#### Agent Action

對 `Element`／`Current` 子物件再跑一次 Attribute EQ，不要沿用 wrapper 結論。

#### Goal / Action / Validation

- Goal: 避免把一層 Attribute 共用寫成整棵樹共用。
- Action: 子節點 Attribute vs Load-tree pool。
- Validation: tree pool nonempty + child Attribute NE。

#### Applies When

- 授權 IL2CPP；Parse 先 copy wrapper 再 `Element`／foreach 子節點。

#### Does Not Apply When

- 子節點 Attribute 確實 EQ 樹。
- 純靜態。

#### Validation

Same Parse window; child Attribute tagged tree=NE with nonempty tree pool.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
