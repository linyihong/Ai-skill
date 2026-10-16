> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-16 - IL2CPP paytable GetMain may take System.Char, not string

Status: candidate

#### One-line Summary

Dump 主賠付表時，若 `GetMainPayoutInfo` 對 alphabet 全回 null，先確認參數是 **System.Char（U16 valuetype）**，不要用 `il2cpp_string_new`；必要時再走 `Dictionary<char, …>` 欄位。

#### Human Explanation

Managed API 看起來像「用符號字串查表」，但 IL2CPP 簽名可能是 `GetMainPayoutInfo(char)`。把單字元包成 managed string 再 `runtime_invoke`，常見結果是整表 null，容易誤判「tiers 只能 OCR」。正確做法：`Memory.alloc` 寫入 `writeU16(codePoint)` 當 valuetype 參數；或讀 paytable 物件上的 `m_Info`／`Info` Dictionary，用 entry 的 char key 解 tiers。

#### Trigger

- Alphabet 有值，但 paytable dump 符號數 = 0。
- OCR／美術頁有 3/4/5 倍率，Frida GetMain 全空。
- 反編譯／字段顯示 `Dictionary<char, …>` 而 wrapper 卻傳 string。

#### Evidence

- Tool: Frida `il2cpp_runtime_invoke` + char valuetype arg vs string; optional Dictionary entry walk.
- Sanitized observation: string args → null tiers; U16 char args（或 Dictionary walk）→ 完整 3/4/5 multipliers，與 in-cabinet page OCR 一致。
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/*/win-calc.md` + dump script `charArg`

#### Generalized Lesson

1. 先對 method 參數型別：string ≠ char ≠ 單字元 string。
2. Char 參數用 U16 buffer，不要 `string_new`。
3. GetMain 全空時，查 `Dictionary<char,T>` 字段作回退，不要立刻改 OCR-only。
4. 規則名（如 `common`）與參數型別是兩件事；兩者都錯會雙重空表。

#### Agent Action

修 dump script 的 `charArg`；對常見 rulesName 重試；比較 OCR 與 live tiers；把 source 寫進 win-calc。

#### Goal / Action / Validation

- Goal: live-lock client paytable multipliers。
- Action: char-arg GetMain 或 Dictionary dump。
- Validation: 符號數與 OCR／美術頁一致；`$`／`#` 等 feature 可無 main tiers。

#### Applies / Does Not Apply

- Applies: 授權 IL2CPP；slots paytable dump via managed invoke。
- Does not apply: 純 managed C#／非 IL2CPP；或 API 已證實為 string key。

#### Validation

- U16 char（或 Dictionary walk）dump 後符號數 > 0，且 3/4/5 倍率與 in-cabinet paytable 頁一致；string_new 路徑必須仍能重現全 null（對照）。

#### Promotion Target

- apk-analysis unity-il2cpp Frida valuetype invoke checklist（candidate）。

#### Reuse Evidence

- 尚未有獨立 cabinet／專案的重用證據；保留為 candidate。

#### Required Linked Updates

N/A

#### Promotion Record

N/A
