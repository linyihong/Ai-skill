> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-15 - HTTP 200 JSON `Invalid sign` may mean missing initHeaders, not bad RSA

Status: candidate

#### One-line Summary

TLS 已通、回 **HTTP 200 + JSON `status=2` / `Invalid sign`** 時，在 sign canonical 已與裝置對齊後，優先檢查是否缺少 **冷啟 Global.initHeaders() 全量 header 模板**（`bigdataSession`、`afid`、螢幕/硬體欄位等），不要先假設伺服器 IP 封鎖或 RSA 錯誤。

#### Human Explanation

與 `140200`（403 edgesuite = anti-bot）不同：此類失敗已到 **業務 JSON 層**。實測矩陣：精簡 8 欄 headers + 正確 sign → `Invalid sign`；**裝置真實 sign 原樣重放 + 精簡 headers** 仍失敗；同一 sign + **完整 cold header template**（去 Authorization）→ `status=0`。因此 `Invalid sign` 在此端點可能是 **header 不足** 的泛化拒絕碼，易誤導無限改 canonical。

#### Trigger

- Host bootstrap: HTTP 200, `message=Invalid sign`, sign_len=344, key FP matches
- Device Frida bootstrap succeeds; offline `canonical_match: true`
- Adding only `deviceId`/`androidId`/`sign`/`platform` insufficient

#### Evidence

- Tool: A/B header matrix on same millis+sign+body
- Sanitized matrix:
  - minimal headers → biz=2 Invalid sign
  - device sign + minimal → biz=2
  - device sign + full initHeaders template → biz=0
  - fresh offline sign + full template + synthetic device → biz=0 new uid
- Evidence path: `<PROJECT_ROOT>/api/dynamic-w27-host-bootstrap-invalid-sign-analysis.md` §W27c

#### Generalized Lesson

```text
Invalid-sign JSON 200 triage (after canonical parity):
  IF offline canonical_match == true
    AND sign_len + key_FP OK
    AND response is JSON (not edgesuite HTML)
  THEN before server-fraud hypothesis:
    1. Export device Global.initHeaders() / header_template on cold start
    2. Replay with FULL template minus Authorization for bootstrap
    3. Binary-search required header subset only after full template passes
  Distinct from 140200: 403 HTML = TLS/WAF; 200 JSON Invalid sign = sign OR headers
```

#### Agent Action

1. Headless SDK：bootstrap facade 必須合併 cold header builder，不可只送 sign/deviceId。
2. Live test 分層：canonical parity test + full-header replay test。
3. 文件 Capability Assessment 在 header 未齊前不得標 host bootstrap Done。

#### Goal / Action / Validation

- Goal: 避免誤判「主機無法建遊客」而放棄 host-only farm。
- Validation: minimal vs full header matrix; at least one `status=0` with synthetic identity on host。

#### Applies When

- Cold guest bootstrap / first-token endpoints
- Rich mobile client headers observed in W1b-style probes
- Business JSON errors despite sign offline parity

#### Does Not Apply When

- HTTP 403 edgesuite (see `140200`)
- `canonical_match: false` — fix signer first
- Authenticated endpoints where Authorization is required in canonical

#### Promotion Target

- `analysis/apk/workflows/headless-sdk-device-executor-flow.md`
- `analysis/apk/workflows/http-api-documentation-flow.md` §header template

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
- `feedback/history/development-guidance/README.md` (cross-ref vs 140200)
