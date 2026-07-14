> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - RSA public unwrap to short key then named content decrypt

Status: validated

#### One-line Summary

章節明文管線常見兩段式：**`RSAUtils.decryptByPublicKey`(大輸入)→短輸出（約數十 char／bytes）**，再進 **具名 `getContentBody`／對稱 decrypt**；Frida 用 inLen／outLen 對齊時間線，不要假設短 RSA 輸出是正文。

#### Human Explanation

接續「getContentBody → UTF-8 plaintext」：短 `decrypt` 輸入（約 55）可來自 **RSA 公鑰解密大正文 blob** 的結果，再被 MD5／對稱 decrypt 使用。呼叫棧常經 `ReaderUtils.readContent`（或同義 reader helper）而非 OkHttp Response 物件。這與「整包 AES-only」或「RSA 直接出全文」都不同——要分別記錄 unwrap 與 content decrypt 兩步。

#### Trigger

- `getContentBody`／content decrypt 出口已見 UTF-8 正文。
- 同窗出現 `decryptByPublicKey`／`publicDecrypt` with **inLen ≫ outLen**（outLen 像 key／IV／密鑰包裝）。
- MD5 util 對長度≈outLen 的緩衝做 `byteToString`／`encrypt32`。

#### Evidence

- Tool: Frida hooks on RSA＋MD5＋getContentBody；reader helper in stack.
- Sanitized lengths only（example order：RSA in≈tens of KB → out≈55；MD5 55→32；getContentBody KB→KB prose）。
- No key／ciphertext／novel text in Ai-skill.

#### Generalized Lesson

1. Timeline-correlate RSA／KDF／content-decrypt by lengths.  
2. Label stages: `rsa-unwrap` → `key-material` → `content-decrypt` → `utf8-body`.  
3. Prefer reader util stack（`readContent`）over Response.peekBody for final plaintext.  
4. Short RSA output is usually **key material**, not chapter text.

#### Agent Action

1. Hook RSA decrypt＋getContentBody＋MD5 in one script；log inLen／outLen／stack.  
2. Map who supplies RSA input（wire field vs file vs memory）.  
3. Keep working keys out of docs； project notes name fields only when known.

#### Promotion Target

- Layer: analysis
- Status: promoted
- Required Linked Updates:
  - `analysis/apk/traffic-triage.md` — Response 解碼補 RSA-unwrap→short-key→content-decrypt
  - N/A: intelligence／workflow／enforcement

#### Verification

Validated when one reader window shows RSA large→short then getContentBody producing UTF-8 prose, with reader util on the stack.
