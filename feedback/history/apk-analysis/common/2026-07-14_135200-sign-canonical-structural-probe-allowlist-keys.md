> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-14 - Sign-input structural probes: allowlist keys; do not print JWT-shaped tokens

Status: candidate

#### One-line Summary

描述 request-sign canonical 時只記錄 **長度／分隔符／布林旗標**；用 `key=` regex 時必須 **allowlist**，否則 JWT／長 base64 會被誤當 key 名並洩漏進 log。

#### Human Explanation

找到 signer entrypoint 後，常見下一動是 dump canonical 字串統計。若用 `(?:^|\\s)([A-Za-z0-9_]+)=` 這類寬鬆 regex，JWT 或 base64url blob 會被當成「第二個 key」。另：廣 hook `StringBuilder.toString` 可能 stack overflow。正确做法：布林旗標（startsWith `timestamp=`、`amp=0`、hex32 個數、space 數、LEFT/RIGHT 長度）＋固定 allowlist key 名。

#### Trigger

- Signer 輸入變長、含空白而非 `&`
- Descriptor log 出現超長「key」或 `eyJ`／`ZXlK` 外形
- `StringBuilder` 廣 hook 後 process crash／stack overflow

#### Evidence

- Tool: Frida structural counters on signer `(String)→String`
- Sanitized excerpt: `startsTs=true space=1 amp=0 hex32n=2 eq=2`；allowlist 前曾誤抓 JWT-shaped token
- Evidence path: `<PROJECT_ROOT>/<App>/api/dynamic-w*.md`（結構 only）

#### Generalized Lesson

```text
Canonical structure probe:
  1. Never log signer plaintext / JWT / tokens
  2. Counts: amp, space, eq, hex32n, quote/brace/colon
  3. key= regex → allowlist only (timestamp, path, body, …)
  4. Avoid global StringBuilder.toString hooks
```

#### Agent Action

1. Project notes 只留結構表。
2. Capture scripts scrub JWT-shaped runs before sharing logs。

#### Applicable / Not applicable

- Applicable: HTTP request-sign RE with Frida descriptors
- Not applicable: already have full formula and no need to probe shape

#### Linked Updates

- N/A — technique lesson; workflow wording unchanged

#### Validation

- [x] No secrets in lesson
- [x] Diff review scrubbed JWT-shaped false keys from commit set

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
