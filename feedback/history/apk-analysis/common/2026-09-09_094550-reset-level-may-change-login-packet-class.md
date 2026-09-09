> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-09 - Reset level can change the live login packet class

Status: validated

#### One-line Summary

`force-stop only` 與 `clear app data` 的自動登入，可能走上不同的 login **packet class**（例如 token 恢復 vs 建帳）。一個窗口不能代表兩種 reset。

#### Human Explanation

同一 Guest manager 仍可能：清資料後送「建立身分」封包，下次 force-stop 改送「token 恢復」封包。中途 TCP 重連也可能再走 token 類，而不是建帳類。把 force-stop 序列當成 first-run，會把 session recovery 誤寫成註冊流程。

只記 class 名與時序，不要序列化 body。

#### Trigger

- 只做過 force-stop login 就要寫「Guest 登入」。
- 或清資料後假設下一次冷啟動仍是同一 packet class。

#### Evidence

- Tool: names-only Frida on SendPacket / login managers across reset levels。
- Sanitized excerpt: first-run used a create-guest packet class after SendLogin; later force-stop and in-session reconnect used a token packet class; table observe/join only when a table had been persisted.
- Evidence path: login bootstrap windows under `<PROJECT_ROOT>`

#### Generalized Lesson

1. 每個 login 窗口標 reset level。
2. 至少分開：first-run / clear-data、force-stop recovery、in-session reconnect。
3. 記錄 packet **class**，不要假設 manager 名等於 wire 型別。
4. 下游 feature（觀桌／進桌）可能只在「有 persisted table」的 recovery 出現。

#### Agent Action

寫 login 結論前對照至少兩個 reset；缺的標 open，不要合成一條路徑。

#### Goal / Action / Validation

- Goal: login Discovery 不把 recovery 當成 first-run。
- Action: 分窗口表（reset × packet class × 後續 feature）。
- Validation: 文件有 ≥2 reset；class 名不同處有寫。

#### Applies When

- 授權分析自動登入／Guest／token 恢復。

#### Does Not Apply When

- 目標沒有本機 session（每次都是全新帳密表單）。

#### Validation

兩次 spawn 的 SendPacket class 名可對上表。

#### Promotion Target

- `workflow/apk-analysis/execution-flow.md` Reset baseline 表

#### Required Linked Updates

- 必須：apk-analysis README Recent + common 計數
- 具體 packet 類名留專案文件
