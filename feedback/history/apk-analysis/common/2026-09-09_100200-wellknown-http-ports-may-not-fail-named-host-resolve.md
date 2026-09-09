> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-09 - Well-known HTTP(S) ports may not fail named host resolve

Status: validated

#### One-line Summary

擋裝置 OUTPUT 上常見的 HTTP／HTTPS port，**不能**當成已觸發 metadata 裡的 host-resolve Failed／Timeout 回呼。Success 仍可能接著連業務 TCP。

#### Human Explanation

Unity 客戶端常有「HTTP helper 找 host」再「自訂 TCP session」兩段（見 custom-TCP lesson）。實驗室若只拒絕 well-known web port，helper 仍可能在短時間內走 Success，然後 Connect 業務 port。Failed／Timeout 方法沒打到，不代表那些方法不存在，只代表**這次干擾沒打中 helper 實際用的通道**（別的 port、IPv6／另一條 filter、快取、或規則沒蓋到該 uid）。

靜態存在的 session-confirm 類 packet 也可能從不出現在自動 Guest／token 恢復路徑上。不要把「刻意打失敗」預設成會打到那個 packet。

不要把具體 firewall 指令寫成可重用攻擊步驟。

#### Trigger

- 目標是觀察 named resolve-fail／timeout 回呼。
- 只擋了常見 web port 之後，仍看到 Success，接著業務 TCP login packet class。

#### Evidence

- Tool: names-only Frida on resolve Success／Failed／Timeout plus login packet classes; device OUTPUT drops on well-known HTTP(S) ports, including after clear-data.
- Sanitized excerpt: Success still ~sub-second after helper enter; Failed／Timeout zero hits; session-confirm class never constructed on that path.
- Evidence path: login bootstrap notes under `<PROJECT_ROOT>`

#### Generalized Lesson

1. 先分清 HTTP helper 與業務 TCP port（`ss`／names-only Connect）。
2. well-known web port 陰性 ≠ Failed 回呼不存在。
3. session-confirm 類型若只有 ctor／ParseResult 且自動登入零命中，不要用它當失敗窗的預期事件。
4. 授權範圍內才干擾網路；不要輸出 host／URL／body。

#### Agent Action

resolve-fail 窗口若仍 Success，記錄「干擾未命中 helper」，改對照業務 session port 或標記 open；不要升級成 dump URL。

#### Goal / Action / Validation

- Goal: 分清「沒打到 helper」與「沒有 fail API」。
- Action: 對照 resolve 方法時間戳與後續 TCP Connect／login class。
- Validation: Failed／Timeout 零命中時 Success 與業務 Connect 仍可並列為證據。

#### Applies When

- 授權動態分析；metadata 同時有 HTTP resolve helper 與自訂 TCP login。

#### Does Not Apply When

- 純 HTTPS 短請求、沒有分開的 resolve helper。
- 未授權對他人主機或網路做干擾。

#### Validation

clear-data 與 force-stop 兩窗：擋 well-known HTTP(S) 後仍 Success → 業務 TCP login class；Failed／Timeout 未 live。

#### Promotion Target

- `workflow/apk-analysis/execution-flow.md` reset / session windows
- `analysis/apk/tools-and-failures.md`
- `unity-il2cpp/2026-09-07_144800-unity-game-session-may-be-custom-tcp-not-tls443.md`
- `common/2026-09-09_094520-radio-toggle-may-not-drop-custom-tcp.md`

#### Required Linked Updates

- 必須：apk-analysis README Recent + common 計數
- 具體 host／port／packet 名留專案文件
