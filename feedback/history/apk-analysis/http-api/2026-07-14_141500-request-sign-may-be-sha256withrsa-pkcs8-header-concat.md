> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-14 - Request sign may be SHA256withRSA over header/body concat with embedded PKCS8

Status: candidate

#### One-line Summary

部分業務 `sign` header = **Base64(SHA256withRSA(canonical))**，private key 以 **PKCS8** 內嵌於 client；canonical 常為 `timestamp=`+millis+body+device 標頭+Authorization+**APK 簽名憑證 MD5**+packageName。

#### Human Explanation

Signer entrypoint 可能只是 `Signature.getInstance("SHA256withRSA")` + `initSign(PKCS8)`，不是 Cipher 塊加密。Frida 見 Base64 256→固定 sign 長度時應優先查 `java.security.Signature`。Canonical 組裝常在 OkHttp interceptor：query 同步寫 `timestamp`，字串拼接裝置／Authorization／安裝包簽名 MD5（uppercase hex）。觀測到的「空白切兩段」可能只是 `Bearer ` 空白，不是協定分隔符。文件與 feedback **只留公式與 key 長度／指紋**，永不貼 PKCS8。

#### Trigger

- Frida：`Base64.encode` inBytes=256、sign 固定長（≈344 for RSA-2048）
- Stack：`Signature`／`SHA256withRSA`／PKCS8EncodedKeySpec
- Interceptor 對 URL 加 `timestamp` query 且 header 寫 `sign`

#### Evidence

- Tool: jadx single-class + Frida Base64/KEYMAT length probe
- Sanitized excerpt: alg=SHA256withRSA；PKCS8 decode outBytes≈1.2KB；sign outLen=344；canonical concat fields as above
- Evidence path: `<PROJECT_ROOT>/<App>/api/dynamic-w*.md`（公式 only）

#### Generalized Lesson

```text
Request-sign reverse:
  1. Confirm Signature alg vs Cipher
  2. Map interceptor canonical field order (do not invent & separators)
  3. Fingerprint embedded key (len + hash prefix); never commit key material
  4. Treat Bearer space as observational split only
```

#### Agent Action

1. Project docs：公式 + FP；decompile scratch gitignore。
2. Capability Assessment 仍 No 直到有交付物 client。

#### Applicable / Not applicable

- Applicable: Android OkHttp apps with RSA request-sign
- Not applicable: HMAC-only or server-issued nonce schemes without local private key

#### Linked Updates

- Extends `http-api/2026-07-14_131600-request-sign-fixed-b64-length-may-be-rsa-scale` and vendored HttpGlobal entrypoint lesson

#### Validation

- [x] No private key / token in lesson body
