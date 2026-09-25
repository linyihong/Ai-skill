> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - Request sign header fixed base64 length may indicate RSA-scale blob

Status: candidate

#### One-line Summary

若業務請求穩定帶 `sign`（或其他密文 header）且 **每次長度固定**（例如 base64 長度對應 ≈256 decoded bytes），優先假設 **RSA-2048 量級封裝／簽章**，不要先當成 MD5／HMAC hex；內容解密若另見 RSA block=128，代表 **sign 金鑰與內容金鑰可能不同**。

#### Human Explanation

MD5／SHA hex 長度通常 32／64；HMAC 可長但仍常隨輸入變。固定 ~344-char base64（≈258B）很像 RSA-2048 ciphertext／signature 再編碼。同時 content path 若 `decryptByPublicKey(..., 128)`，則內容側是 RSA-1024——應分開記錄兩條金鑰用途，避免「找到一把公鑰就能重放 sign」的錯假設。組裝點仍可能在 interceptor／sorted params；本 lesson 只收窄 **輸出形狀** 判斷。

#### Trigger

- 多數／全部業務 Request 都有同名 `sign` header。
- 觀測多筆 **signLen 完全相同**。
- Content crypto 另有不同 RSA block／key-size signal。

#### Evidence

- Tool: OkHttp Request dump via `toString()` when `url()`／`headers()` are R8-broken；length-only.
- No sign values in Ai-skill.

#### Generalized Lesson

1. Classify sign by **length family**: hex32／hex64／variable／fixed-b64≈256B.  
2. Fixed ≈256B decoded → investigate RSA-2048 encrypt／sign utils，not only MD5.  
3. Content RSA-1024 vs sign RSA-2048 ⇒ dual-key hypothesis.  
4. Assembly formula remains a separate gate（interceptor／canonical map lessons）。

#### Agent Action

1. Log `signLen`＋header name set for each business URL（lengths only）.  
2. Compare to `publicEncrypt`／`privateDecrypt`／`Signature` outLen in same window.  
3. Keep MD5 hooks，但用 stack／長度排除「sign=raw MD5」假說。

#### Promotion Target

- Layer: analysis
- Status: promoted
- Required Linked Updates:
  - `analysis/apk/traffic-triage.md` — request auth／sign 補固定 base64 長度分類
  - N/A: intelligence／workflow／enforcement

#### Verification

Validated when ≥10 business requests share identical signLen in the RSA-2048-size band while content RSA uses a different block size in the same app.

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
