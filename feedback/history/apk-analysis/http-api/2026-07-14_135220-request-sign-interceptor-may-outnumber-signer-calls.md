> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-14 - Request-sign interceptor invokes may outnumber signer calls

Status: candidate

#### One-line Summary

OkHttp／vendored HTTP interceptor 次數 **不必等於** signer `(String)→String` 次數；sign 可能條件重算或快取，勿假設 1:1。

#### Human Explanation

找到 interceptor → signer 後，後續大量業務 URL 可能只有 interceptor 日誌、沒有 signer 輸入。若解讀成「這些请求未簽名」會錯；更常見是 header 沿用先前算出的 sign，或僅在 body／timestamp 變時重算。驗證：冷啟動前段應對上多次 signer；之後比對 header 仍有固定長 sign，同時 signer hook 安靜。

#### Trigger

- Signer outLen 固定；早期多次 hit，後期 interceptor ≫ signer
- 業務 URL 仍帶 sign header（長度穩定）

#### Evidence

- Tool: Frida dual hook interceptor + signer
- Sanitized excerpt: cold-start signer batch then many interceptor paths without new signer calls
- Evidence path: `<PROJECT_ROOT>/<App>/api/dynamic-w*.md`

#### Generalized Lesson

```text
Sign frequency check:
  1. Count signer vs interceptor separately
  2. Confirm sign header still present/length-stable when signer quiet
  3. Treat cache/conditional resign as default hypothesis
```

#### Agent Action

Document both counts in analysis window; do not require signer hit on every URL for “signed traffic”.

#### Applicable / Not applicable

- Applicable: mobile HTTP request-sign RE
- Not applicable: per-request HMAC always visible in same stack frame

#### Linked Updates

- N/A — extends entrypoint lessons；no workflow edit required

#### Validation

- [x] Sanitized
