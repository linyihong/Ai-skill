> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-07 - Slot S2C may apply via ProcessSpin GameAction not ParseResult

Status: validated

#### One-line Summary

Unity 老虎機若 `ParseResult` 沒打在下注封包上，改 hook 機台 controller 的 `SendSpin` / `ProcessSpin`：C2S 回傳 packet 型別，S2C 常回 `GameAction[]` 驅動轉輪動作。

#### Human Explanation

信封 RPC（時間／好友）走 `ParseResult` + XML 文件物件；功能旋轉的結果卻可能先進 UI 層再拆成 action graph。只 hook `Packet.ParseResult` 會誤判「沒有 S2C schema」。`SendSpin` 的 retval class 對到已知名的 C2S packet；稍後 `ProcessSpin`（或 round-commit）回傳 action 陣列，轉輪 stop 是獨立 action 型別，不是把 XML dump 出來。

Idle 會大量呼叫 `get_*TableInfo`／jackpot 探測；hook 過濾必須排除 getter，否則 log 不可用。

#### Trigger

- 已證明 C2S 動作封包，但該型別沒有 `ParseResult` hit。
- HUD／轉輪已更新。

#### Evidence

- Tool: Frida names-only on methods matching spin/process (no string bodies).
- Sanitized excerpt: controller `this` is a cabinet subclass; send-spin returns the bet packet type; process-spin / round-sink return a game-action array; reel-stop is a cabinet-prefixed reel-result action wrapping a reel model.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/protocol.md`

#### Generalized Lesson

1. `ParseResult` 沒打中功能封包 ≠ 沒有結果路徑。
2. 先找 `Send*` 回傳 packet class，再找 `Process*` 回傳 action／model。
3. Hook 排除 `get_`／jackpot-support 探測。
4. 不要印 action 欄位值或 XML。
5. 套用路徑清楚後：對 reel / spin-info **model** 做 `il2cpp_class_get_fields`（名＋型別）。那是 decrypt 之後的客戶端 DTO，不是線上欄位表。

#### Agent Action

短窗、names-only。禁止 dump 明文與金鑰。

#### Goal / Action / Validation

- Goal: 對上 S2C **套用** 的型別名（不是猜 JSON）。
- Action: 一次旋轉；記錄 send vs process 兩段 retval class。
- Validation: 同一手勢可再現 send→process 順序（允許穿插 reel getter）。

#### Applies When

- IL2CPP Unity；授權 attach；存在 spin/process 風格方法名。

#### Does Not Apply When

- 結果已在 `ParseResult` retval DTO 上。
- 純 HTTP JSON converter。

#### Validation

重複一次旋轉應再見到 process 回傳 action 陣列。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
- Project 細節留 `<PROJECT_ROOT>` docs
