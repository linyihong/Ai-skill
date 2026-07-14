> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - Request sign may live in vendored HttpGlobal interceptor, not app GnIntercept

Status: validated

#### One-line Summary

當 App 自有 `*Intercept` 未見組 `sign`，但每個業務請求仍有固定長度 `sign` header 時：優先在 **vendored HTTP 套件**（如 `*.http.HttpGlobal$a.intercept`）與同棧 **`*.net.*.a(String)→String`（outLen=固定 base64）** 找組裝點；並 hook **套件內自帶 Base64**，不要只 hook `android.util`／`java.util.Base64`。

#### Human Explanation

App 層 interceptor 可能只管 timing／logging／body stringify，真正的 request signing 常落在共用 HTTP 模組。實證模式：interceptor → `sign(canonical)→fixedLenB64`；static init 可能先 `Base64.decode` 大段內嵌 key material。第三方廣告 SDK 的 RSA-OAEP（例如 32→256）容易誤判成業務 sign——必須用 stack package 排除。

#### Trigger

- 穩定 `sign` header 長度族（固定 base64）已確認。
- App `Interceptor`／`Request.Builder.addHeader("sign")` hooks 空窗。
- DEX／loaded classes 出現 `HttpGlobal`／`NRKeyManager`／`com.*.http.common.Base64`。

#### Evidence

- Tool: Frida on vendored interceptor + `a(String)→String` outLen correlate to known `signLen`.
- Sanitized: class／method names, inLen／outLen only; no canonical／sign／key bytes.

#### Generalized Lesson

1. Treat app interceptor as **suspect, not default** for custom sign.  
2. Prefer fixed outLen method as sign function; then walk callers to interceptor.  
3. Always enumerate **library Base64** helpers beside platform ones.  
4. Disambiguate ad／analytics RSA from business sign via stack root package.

#### Agent Action

1. Enumerate `*http*`／`*KeyManager*`／custom `Base64` types.  
2. Hook candidate `String→String` with constant outLen≈signLen.  
3. Confirm caller is global OkHttp interceptor before claiming formula.

#### Promotion Target

- Layer: analysis
- Status: promoted
- Required Linked Updates:
  - `analysis/apk/traffic-triage.md` — request auth／sign 補 vendored HttpGlobal／custom Base64
  - N/A: intelligence／workflow／enforcement

#### Verification

Validated when hooks show every business request gets `sign` from a vendored `intercept`→`a(String)` with constant outLen matching observed header length, while app-owned interceptor does not set `sign`.
