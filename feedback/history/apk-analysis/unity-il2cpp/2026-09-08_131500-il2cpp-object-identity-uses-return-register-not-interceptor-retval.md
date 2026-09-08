> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - IL2CPP object identity uses return register, not Interceptor retval

Status: validated

#### One-line Summary

比對 IL2CPP **物件回傳值** 與下一個 `Byte[]` 參數時，用 **return register**（arm64：`this.context.x0`）保存指標；`Interceptor` `onLeave` 的 `retval` 常不是有效 managed object。

#### Human Explanation

`onLeave(retval)` 對許多 IL2CPP 方法看起來像 NativePointer，但拿去 `il2cpp_object_get_class` 會得到非物件（低位址／null）。同一瞬間讀回傳暫存器才是 `Byte[]` 實例。若用失效的 `retval` 做 pointer EQ，會誤判成「Decrypt 與 GetString 不是同一塊陣列」，其實是 hook 取錯指標。

#### Trigger

- 要證明 `Byte[]` 回傳值是否原樣傳進 `Encoding.GetString`（禁止讀內容）。
- `retval` 無法通過 `object_get_class`，或一律 EQ 失敗。

#### Evidence

- Tool: Frida Interceptor; class-name on stored pointer; EQ/NE only.
- Sanitized observation: Interceptor retval was not a managed array; the return register pointer was `Byte[]` and matched GetString’s array argument. A nested native transform return was a different object.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. 物件 identity：clone **return register**，不要信任 IL2CPP 的 Interceptor `retval`。
2. 先用 `il2cpp_object_get_class` 確認兩邊都是 `Byte[]`，再 EQ；不要 dump 內容或印出位址。
3. 內層 Transform／ToArray 可能是另一個物件；EQ 失敗時分別保存各層回傳。
4. arm64 用 `x0`；其他 ABI 用對應的 return register。

#### Agent Action

Identity 探針一律存 `ptr(this.context.<return-reg>)`；`nobj` 就改 register，不要寫成「有 copy」。

#### Goal / Action / Validation

- Goal: 避免把 hook 取錯指標當成業務端 copy。
- Action: class-name + EQ/NE；不讀 buffer。
- Validation: Decrypt 回傳 class=`Byte[]` 且與 GetString 參數 EQ；Interceptor retval 路徑為 nobj。

#### Applies When

- 授權 IL2CPP／Frida；arm64 或已知 return register。

#### Does Not Apply When

- 回傳值是 valuetype／blittable（不在 x0 當物件）。
- 純靜態、無 live identity 需求。

#### Validation

nobj on Interceptor retval vs Byte[]+EQ on return register in the same window.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
