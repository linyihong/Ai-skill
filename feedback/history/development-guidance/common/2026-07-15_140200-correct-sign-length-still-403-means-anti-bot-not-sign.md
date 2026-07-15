> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-15 - Correct offline sign length but HTTP 403 usually means anti-bot, not sign formula

Status: candidate

#### One-line Summary

主機離線 signer 產出的 `sign` **長度與 in-app fingerprint 一致**，但 API 仍 **HTTP 403**（edgesuite／challenge HTML）時，優先判 **WAF／anti-bot／TLS fingerprint**，不要繼續改 canonical 公式。

#### Human Explanation

Agent 易在 403 時無限調 sign。若 `--sign-only` 已對齊 in-app key FP 與 sign len，403 多半是 **外部 JVM／桌面 IP** 被擋，非簽名錯誤。解法階梯：(1) device executor（Frida in-app HTTP）；(2) adb forward + on-device proxy；(3) 移動 IP／補齊 client header 族——不是再猜 HMAC。

#### Trigger

- Mac／CI headless POST → 403 HTML `Access Denied`
- `sign_len` matches device export; body JSON matches schema
- Same request via device Frida → HTTP 200

#### Evidence

- Tool: offline signer + host curl vs device executor
- Sanitized excerpt: sign_len=344 FP match; host 403; device 200 status=0
- Evidence path: `<PROJECT_ROOT>/api/dynamic-w21b-*.md`

#### Generalized Lesson

```text
403 triage after sign RE:
  IF sign_len + key_FP match in-app
    AND wire body schema validated
    AND host gets 403 HTML (not JSON business error)
  THEN classify as anti-bot / edge block
  THEN pivot to device executor before more sign tweaks
```

#### Agent Action

1. 專案 probe 腳本分 `--sign-only` 與 `--live`；live 失敗時標 WAF tier。
2. 交叉引用 `094500` anti-bot gateway、`095400` device proxy bypass。
3. Private SDK 架構預設 device executor 為 production path。

#### Goal / Action / Validation

- Goal: 停止錯誤的 sign-RE 迴圈。
- Validation: A/B host 403 vs device 200 同 payload。

#### Applies When

- Building external SDK against CDN/WAF-protected API
- Sign offline parity already proven

#### Does Not Apply When

- 403 with JSON `Invalid sign` business code (sign path wrong)
- TLS handshake failure before HTTP response

#### Promotion Target

- `analysis/apk/workflows/headless-sdk-device-executor-flow.md`
- `analysis/apk/traffic-triage.md`

#### Required Linked Updates

- `feedback/history/development-guidance/README.md`
- `analysis/apk/traffic-triage.md`
