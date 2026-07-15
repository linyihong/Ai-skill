> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-15 - HttpGlobal header setter hook for token when OkHttp introspect fails

Status: candidate

#### One-line Summary

R8 下 `req.url()`／`peekBody` 可能 **TypeError**；bootstrap 成功與 token 長度可改 hook **vendored `HttpGlobal` header setter**（`Authorization`）或 **`SpData.setUserToken`**。

#### Human Explanation

Frida E2E 需知「bootstrap 何時完成」。混淆 OkHttp `Request` 不穩時，高語意點是 app 寫入 Bearer 的瞬間：`HttpGlobal.u("Authorization", value)` 或 prefs `setUserToken`。只記 `tokenLen`，不 `send()` 完整 token。與 `130100`（Response peek 失敗）、`133500`（sign in HttpGlobal）同族。

#### Trigger

- `req.url()` throws TypeError on obfuscated OkHttp Request
- `peekBody` 0 hits but bootstrap clearly succeeded
- Need cold-start E2E gate before downstream API calls

#### Evidence

- Tool: Frida hook HttpGlobal.u + SpData.setUserToken; pm clear + spawn
- Sanitized excerpt: Authorization setter fires once; tokenLen≈180–200 class post-bootstrap
- Evidence path: `<PROJECT_ROOT>/scripts/frida/w22c_*.js`, `w22d_*.js`

#### Generalized Lesson

```text
Token-ready detection when OkHttp API breaks:
  1. Hook HttpGlobal / Global header map setter for "Authorization"
  2. Alternate: SpData or equivalent prefs write for user token
  3. Log tokenLen only; never send() raw Bearer to host
  4. Gate downstream HTTP on tokenLen > 0 or setter fired
  5. Avoid req.url() on obfuscated Request — use path hook at Retrofit layer if needed
```

#### Agent Action

1. E2E script: wait for token setter before task API loop。
2. 與 `132500` compact send lesson 並用。
3. 專案 capture 存 redacted session env（gitignore）。

#### Goal / Action / Validation

- Goal: Reliable bootstrap-complete signal without OkHttp introspect。
- Validation: setter hook → downstream index/sign 200。

#### Applies When

- Vendored networking module (HttpGlobal pattern)
- Cold-start spawn E2E on obfuscated OkHttp builds

#### Does Not Apply When

- Standard OkHttp Request API fully usable
- Token only in memory field without setter hook point

#### Promotion Target

- `analysis/apk/workflows/frida-hook-flow.md` §device E2E executor
- `analysis/apk/workflows/headless-sdk-device-executor-flow.md`

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
- `analysis/apk/workflows/frida-hook-flow.md`
