> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-09 - Radio toggle may not drop a custom TCP game session

Status: validated

#### One-line Summary

`svc` 開關 Wi‑Fi／數據或 shell 發的 airplane broadcast，常常**不斷**已建立的自訂 TCP session。重連／relogin hook 沒打到，不代表沒有重連路徑。

#### Human Explanation

Unity 等客戶端常把業務 session 放在非 443 的長連線上（見 custom-TCP lesson）。系統無線電切換可能讓該 socket 在逾時前仍 ESTABLISHED，Connector 類 disconnect/reconnect 不會跑。shell uid 也常不能發 `AIRPLANE_MODE` broadcast。

要觀測 in-session recovery，應在授權 lab 對**該 session 的 TCP** 做重置（例如對已觀測的 game port 拒絕/重設），而不是先假設「沒有 SessionConfirm／Reconnect」。不要把具體 firewall 指令寫進可重用文件當攻擊步驟。

#### Trigger

- attach 後關 Wi‑Fi／數據十幾秒，disconnect/reconnect 方法零命中。
- `ss` 仍顯示同一個遠端 game port ESTABLISHED。

#### Evidence

- Tool: `svc wifi|data` vs resetting the established game-session TCP; Frida names-only reconnect hooks。
- Sanitized excerpt: radio toggle produced no Connector disconnect; resetting the session port produced Disconnected → Reconnect → login-token-class send.
- Evidence path: login bootstrap notes under `<PROJECT_ROOT>`

#### Generalized Lesson

1. 先確認業務面是自訂 TCP 還是 HTTPS。
2. 無線電切換無 hook ≠ 無重連；查 `ss`/`netstat` 該 port 是否還 ESTAB。
3. 授權範圍內重置**該 session**，再記錄 packet **class 名**，不要 dump body。
4. 不要用未授權的 firewall／DoS 配方當通用 runbook。

#### Agent Action

重連窗口失敗時先看 socket 是否還在，再考慮 session-port reset；不要升級成讀 token。

#### Goal / Action / Validation

- Goal: 分清「沒斷線」與「沒有重連 API」。
- Action: radio toggle 陰性時改 session TCP reset。
- Validation: disconnect 與 relogin packet class 有 names-only 命中。

#### Applies When

- 授權動態分析；已證明業務在長生命 TCP。

#### Does Not Apply When

- 純 HTTPS／短請求 API。
- 未授權對他人主機或網路做干擾。

#### Validation

同一 attach：radio 無事件；session TCP reset 後有 Reconnect + login packet class。

#### Promotion Target

- `workflow/apk-analysis/execution-flow.md` reset / session windows
- `unity-il2cpp/2026-09-07_144800-unity-game-session-may-be-custom-tcp-not-tls443.md`

#### Required Linked Updates

- 必須：apk-analysis README Recent + common 計數
- 具體 host／port／packet 名留專案文件
