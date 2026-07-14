> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-14 - Cross-check request-sign with offline key reconstruct vs in-app signer fingerprint

Status: candidate

#### One-line Summary

驗證內嵌 PKCS8 request-sign 时：在 **gitignore scratch** 重建 key，对本机 synthetic canonical 簽名，再 Frida 调 app signer，只比对 **outLen + sha256_8(sign)**，不记录 sign／canonical／key。

#### Human Explanation

公式落地后需要证明「反编译拿到的 key 真的是 runtime 在用的那把」。把 key 或 sign 明文写进 repo／chat 会越界。可行办法：scratch 目录重组 PKCS8 → 检查长度／已知指纹 → 签一条无专用 canonical → Frida 调同一入口 → 比较摘要。匹配即閉じ dry-run；不匹配再查 Base64／字符集／拼接。

#### Trigger

- 已还原 SHA256withRSA + 内嵌 PKCS8
- 需要证明 offline key ≡ in-app key
- 不能提交私钥或 sign 原文

#### Evidence

- Tool: local Python cryptography + Frida invoke signer
- Sanitized excerpt: outLen=344；sha256_8 match=true；key FP 与静态 decode 一致
- Evidence path: `<PROJECT_ROOT>/<App>/api/dynamic-w*.md` + gitignored `.tmp_static/dryrun/`

#### Generalized Lesson

```text
Sign dry-run ladder:
  1. Reconstruct key only in gitignored scratch
  2. Verify key len + hash prefix vs dynamic decode
  3. Sign synthetic canonical offline
  4. Frida call same signer; compare outLen + hash(sign)
  5. Never commit key/sign/canonical plaintext
```

#### Agent Action

Document match in project notes；scratch stays ignored；Assessment 仍可保持 No。

#### Applicable / Not applicable

- Applicable: embedded private-key request-sign RE
- Not applicable: server-only HMAC with no client private key

#### Linked Updates

- Extends `http-api/2026-07-14_141500-request-sign-may-be-sha256withrsa-pkcs8-header-concat`

#### Validation

- [x] Lesson contains no key/sign material
