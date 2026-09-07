> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-07 - Leave live table may send offers then leave then menu layout

Status: validated

#### One-line Summary

從進行中的機台回到列表，C2S 往往是「離桌優惠封包 → 真正離桌 opcode → 刷新選單 layout」，不要把中間的 retention dialog 當成沒有網路、也不要把優惠封包當成唯一離桌指令。

#### Human Explanation

選單上的離開按鈕可能先打一發 offers/retention RPC，畫面上才出現「看廣告拿籌碼 / 直接退出」。使用者按退出（略過 rewarded video）之後才送 leave opcode，接著再要一份 menu layout。只 hook leave 型別名會漏掉第一段；只看到 offers 會誤以為已經離桌。不要 dump 優惠內容或廣告 SDK payload。

**2026-09-08 revision:** `GetLeaveTableOffers` 仍可能先送出，即使畫面只是「否／是」確認、沒有 rewarded-video overlay。不要用彈窗樣式判斷有沒有 offers 封包。

#### Trigger

- UI 已出現離桌確認／看廣告彈窗，但 feature leave opcode 尚未出現。
- 回到 lobby 後列表有刷新。

#### Evidence

- Tool: Frida names-only `SendPacket` + EncryptData length/first-char (no bodies).
- Sanitized excerpt: offers packet (alpha token starting with `L`) then leave packet (`S`) then menu-layout packet (`G`).
- Evidence path: `<PROJECT_ROOT>/<App>/docs/protocol.md`

#### Generalized Lesson

1. 離桌 UI 可能拆成 offers → confirm → leave → layout refresh。
2. 記錄三段 type 名 + 字串格式分類即可；不要印 command 正文。
3. 略過 rewarded video 的按鈕與「領取」按鈕座標要分開驗證，避免誤觸廣告路徑。

#### Agent Action

離開手勢短窗只掛輕量 `SendPacket`。禁止 dump offers XML／字串。

#### Goal / Action / Validation

- Goal: 對上離桌 C2S 序列，而不是單一 opcode。
- Action: 菜單離開 → 略過廣告 → 回到列表；對照三次 `SendPacket` 型別名。
- Validation: 同一序列可再走一次（允許穿插 ping／buddy）。

#### Applies When

- 長連線 session；機台內有明確 leave 選單；授權 Frida attach。

#### Does Not Apply When

- 直接 force-stop／殺行程離桌。
- HTTP REST 單一 DELETE。

#### Validation

同一 Guest session 再進另一台櫃時，enter opcode 應仍是既有 play-now／join 序列。

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
- Project 細節留 `<PROJECT_ROOT>` docs
