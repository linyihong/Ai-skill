> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Frida chain harness must not block forever on stdout readline

Status: validated

#### One-line Summary

等待 instrumentation 子行程輸出並設有 idle／chain timeout 時，**不要**對 stdout 用會永久阻塞的 `readline()`；改用 `select`／短超時讀取，否則 deadline 檢查永遠跑不到，hunt 會在「已有 RESULT、尚未 CHAIN_DONE」卡住。

#### Human Explanation

多段 RESULT 捕獲常在第一段之後靠 idle timer 印出 chain-done。若 Python 側用阻塞 `readline()` 等下一行，而 Frida 暫時無輸出（timer 未到、或輸出延遲），迴圈無法回到 timeout 判斷，整輪 capture 會掛死。非阻塞／短超時讀取可讓 chain timeout 與 overall deadline 真正生效，即使 chain-done 標記漏發也能帶已收集 stages 繼續 screencap。

#### Trigger

- Subprocess Frida／probe 等待多事件標記 + idle timeout。
- Hunt 日誌停在單一 RESULT／skip dump，卻從不寫 fixture。

#### Evidence

- Tool: Python recapture harness + Frida chain idle marker.
- Sanitized observation: hunt died after one RESULT with no fixture write; after switching to timed stdout reads, same path completed chain-done → settle → screencap.
- Evidence path: `<PROJECT_ROOT>/<App>/scripts/slots/` recapture helper (stdout wait loop).

#### Generalized Lesson

1. Pair every “wait for marker” loop with a wall-clock deadline that is checked even when no line arrives.
2. Prefer `select`／poll with ≤250ms slices over blocking `readline()`.
3. On timeout, proceed with partial stages rather than hang; log which marker was missing.
4. Keep instrumentation idle timers, but do not trust them as the only exit path.

#### Agent Action

When writing Frida／probe collectors with multi-event windows, implement timed stdout reads before launching long hunts.

#### Goal / Action / Validation

- Goal: reliable multi-stage capture loops that cannot hang on quiet stdout.
- Action: replace blocking readline with select-timeout reads; honor chain timeout.
- Validation: capture completes to fixture write after a single-stage RESULT even if chain-done is delayed.

#### Related

- `2026-09-14_163000-slot-collapse-often-single-result.md`
