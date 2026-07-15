> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-15 - Frida device E2E: compact send payloads and off-main-thread in-app HTTP

Status: candidate

#### One-line Summary

Device E2E 用 Frida `send()` 回傳 host 時，勿塞完整 API JSON；in-app HTTP 勿在 Android main thread 執行，否則 host parse 失敗或 `NetworkOnMainThreadException`。

#### Human Explanation

Frida spawn + Python `frida_tools.repl` 常見兩個閉環斷點：(1) script 把完整 response body 經 `send()` 丟給 host，Python 端 `ast.literal_eval`／regex 在超大字串或嵌套引號上失敗，且 pipe 可能延遲到 repl 結束才 flush；(2) 在 `Java.scheduleOnMainThread` 或 UI callback 內直接打 `HttpURLConnection`／OkHttp，觸發 `NetworkOnMainThreadException`。實務上 steps 只送 `{path, http, status}` 等精簡欄位，必要時 host 用 regex fallback + 固定 sleep；網路改在 Frida `setTimeout`／`setInterval` 執行緒跑，或沿用 app 既有 signed header helper。

#### Trigger

- Python host 收不到 Frida `message` 或 parse 失敗，但 device 上 API 其實已成功
- `send({ type: "step", raw: hugeJsonString })` 導致 `literal_eval` / JSON decode error
- Frida log 出現 `NetworkOnMainThreadException` after in-script HTTP replay

#### Evidence

- Tool: Frida 17.x spawn mode; Python `frida_tools.repl` subprocess + threaded stdout reader
- Sanitized excerpt: compact `{path, http, status}` steps parse reliably; full task-center JSON in `send()` breaks host parser; moving HTTP off main thread clears exception
- Evidence path: `<PROJECT_ROOT>/scripts/frida/w22d_device_e2e.js`, `<PROJECT_ROOT>/scripts/headless/run_w22d_e2e.py`（project-only；不含 host／token）

#### Generalized Lesson

```text
Frida device E2E host bridge:
  1. send() only compact, stable fields (path, http code, business status, tokenLen)
  2. Do not embed full API response bodies in send() unless host uses robust JSON framing
  3. If Python parse fails, add regex fallback on stdout + fixed post-tap sleep before repl exit
  4. In-script HTTP: use Frida setTimeout/setInterval thread, NOT Android main thread
  5. Prefer app stack helpers (e.g. existing H5 header builder) over raw HttpURLConnection when sign is fragile
```

#### Agent Action

1. 設計 E2E 時先定 host contract：每 step 最大 payload 大小與欄位白名單。
2. 除錯時把 raw body 寫 device log／project capture，不要經 `send()` 送回 host。
3. 見 `NetworkOnMainThreadException` → 立刻移出 main thread，不要加 `StrictMode` 繞過。
4. 與 `142600`（CLI subprocess for Java RPC）互補：本條處理 **message bridge** 與 **in-app HTTP thread**。

#### Goal / Action / Validation

- Goal: Reliable Frida→Python E2E verdict without silent parse loss
- Action: Compact send schema + off-main-thread HTTP + regex/sleep fallback
- Validation or reference source: Host prints all expected steps; no `NetworkOnMainThreadException`; project doc records http/status only in commit message

#### Applies When

- Spawn-mode Frida driving in-app HTTP for signed API replay
- Python subprocess wrapping `frida -U -f` with stdout parsing

#### Does Not Apply When

- Pure RPC via `rpc.exports` with small return values
- Host uses Frida Python API `script.on('message')` with structured JSON and size limits already enforced
- Traffic relay only (no in-script HTTP)

#### Promotion Target

- N/A (candidate; link from project headless-rewards roadmap if promoted)

#### Required Linked Updates

- N/A
