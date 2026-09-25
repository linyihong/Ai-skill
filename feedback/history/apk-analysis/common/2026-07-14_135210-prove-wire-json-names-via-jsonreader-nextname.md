> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-14 - Prove wire JSON names with JsonReader.nextName when setters are stripped

Status: candidate

#### One-line Summary

R8／Gson field reflection 常使 model **沒有 setX**；證 wire 欄位名優先 hook **`JsonReader.nextName`**（或 converter），不要只找 setter。

#### Human Explanation

章節／媒體 blob 模型可能只有 `getContent`、欄位 `content`，無 `setContent`。Gson `ReflectiveTypeAdapter` 直接灌 field。Hook setter→0 hit，容易誤判「不是這個欄位」。`JsonReader.nextName` 在 Retrofit Gson converter 路徑上可見真實 JSON key；再以 value 長度／magic prefix 對上落地 blob（勿記錄全文）。

#### Trigger

- Model 有 field／getter，無 setter；Frida setter hook 0
- Retrofit + GsonResponseBodyConverter in stack
- DEX 有多個相似字串（如 `content` vs `chapterContent*`）需判誰是 wire

#### Evidence

- Tool: Frida `com.google.gson.stream.JsonReader.nextName`
- Sanitized excerpt: `nextName=content` via ReflectiveTypeAdapter ← GsonResponseBodyConverter；no `chapterContent` in same window
- Evidence path: `<PROJECT_ROOT>/<App>/api/dynamic-w*.md`

#### Generalized Lesson

```text
Wire field proof ladder:
  1. JsonReader.nextName (length/prefix only on nextString)
  2. Converter / fromJson Type hooks
  3. Field reflection / getter after parse
  4. Do not require setters under R8
```

#### Agent Action

1. Prefer nextName before concluding alternate DEX string is the wire key.
2. Keep blob magic／lengths in project notes only.

#### Applicable / Not applicable

- Applicable: Retrofit+Gson apps with R8
- Not applicable: handwritten JSONObject.getString pipelines (hook those instead)

#### Linked Updates

- Complements `common/2026-06-22_130100-r8-obfuscated-okhttp-response-needs-converter-hook`（converter layer）and `http-api/2026-06-22_142900-wire-json-field-names-differ-from-gson-bean-paths`（wire≠bean nesting）

#### Validation

- [x] Sanitized; no host／package／payload

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
