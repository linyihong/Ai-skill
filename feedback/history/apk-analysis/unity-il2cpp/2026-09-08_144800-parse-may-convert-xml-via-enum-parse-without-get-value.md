> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Parse may convert XML via Enum.Parse and Int64.TryParse without get_Value

Status: validated

#### One-line Summary

`Parse(XElement, …)` 期間可能用 **`Enum.Parse`、`Int*.TryParse`、`Convert.ToInt64`** 填標量，即使同窗口看不到 `XAttribute.get_Value` 或 `XmlConvert`。

#### Human Explanation

`get_Value` none 不代表沒有把 XML 轉成數字／enum。應在 Parse 進出窗口掛 BCL 轉換 API（方法名 only），不要 dump 字串或解析結果。有的 Parse 進出可以是 none（提前 return）。

#### Trigger

- 子節點 `Attribute` 有呼叫，但 `get_Value` 為 none。
- 需要知道 DTO 填值走哪條轉換，而不是再比物件 identity。

#### Evidence

- Tool: Frida on `Enum`/`Int32`/`Int64`/`UInt64`/`Convert`/`Double`/`XmlConvert`/`XAttribute.get_Value` during Parse; method names + call presence only.
- Sanitized observation: live set included `Enum.Parse`, `Int64.TryParse`, `Convert.ToInt64`, `Int32.TryParse`/`Parse`, `Double.TryParse`, `UInt64.Parse`, `Enum.ToUInt64`. `XmlConvert` and `XAttribute.get_Value` were not observed. One Parse window had no conversion calls.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. 轉換 API 可以在沒有 `get_Value` hook 的情況下出現。
2. 優先記方法名，不要 dump 參數。
3. `XmlConvert` 不是唯一路徑。
4. Parse 可為 null／無轉換。

#### Agent Action

`get_Value` none 時改掛 `Enum.Parse`／`Int64.TryParse`／`Convert.To*`。

#### Goal / Action / Validation

- Goal: 找出 XML→標量的 live API，避免誤判沒有填值。
- Action: Parse 窗口轉換方法名集合。
- Validation: 至少 `Enum.Parse` 與一種整數 Parse／TryParse。

#### Applies When

- 授權 IL2CPP；XML `Parse(XElement, …)` 填 DTO。

#### Does Not Apply When

- 確實走 `XmlConvert` 或 `get_Value`（照實記錄）。
- 純靜態。

#### Validation

Same Parse window; conversion method names; no string dumps.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
