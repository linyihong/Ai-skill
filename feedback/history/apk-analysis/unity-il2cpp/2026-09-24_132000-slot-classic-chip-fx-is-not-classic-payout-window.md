> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-24 - Classic chip FX class is not the classic payout window

Status: validated

#### One-line Summary

名稱含 `Classic` 的 **ChipsWin FX**（例如 `UISlots*ChipsWinClassicFX`）與 **PayoutWindow classic modal**（例如 `*payout_window_classic`）常是兩條呈現路徑。Classic FX 多半 **Create-on-demand**（idle `FindObjectsOfTypeAll` 為 0），要掛 `Create` 才能獵自然誕生；只掛 resident window `Show` 會漏掉 chip classic。文件與 hunt 必須分列，勿把 idle classic GO 當成 ClassicFX 已開。

#### Human Explanation

Slot 小額／中額「經典」呈現可能是：

1. 轉輪旁／線上 **籌碼 FX**（`*ChipsWinClassicFX.Create`→`Setup`→`Show`）。
2. 全屏／半屏 **PayoutWindow** prefab（`UISlots*PayoutWindow` + classic path）。

兩者 class／GO／Animator 都不同。Idle dump 有 classic window GO 不代表會走 window Show；反之 ClassicFX idle 0 實例也不代表功能不存在。

#### Trigger

- 文件把 chip badge 與 classic modal 混成同一 residual。
- 只 hook window Show，數十轉 0 enter，卻常見線獎籌碼。
- `FindObjectsOfTypeAll(ClassicFX)` 為 0，誤判 class 未載入。

#### Evidence

- Tool: IL2CPP class dump + Frida FindObjectsOfTypeAll + Show/Create Interceptor.
- Sanitized pattern: ClassicFX Create/Setup/Show surface；idle instance count 0；line payout module resident GO separate.
- Evidence path: `<PROJECT_ROOT>` cabinet `lab/classic-chipfx-inventory-*.json`.

#### Generalized Lesson

1. **拆名**：`*ClassicFX` ≠ `*payout_window_classic`。
2. **Create-on-demand**：idle 0 → hook `Create`（靜態回傳實例）再 `Show`。
3. **Hunt 分列**：chip classic 與 window classic 分開計數與證據。
4. **Lab 驗 hook**：對 resident 模組（如 line payout）強制 Show；對 ClassicFX 用 Create 或等自然 Create。
5. **HUD Last Win ≠ 已進 ClassicFX**：餘額／上次贏得可更新，而 Create/Show/SetPayout 計數仍為 0 — 改視覺截圖或找 Animator／Spine 路徑，勿空轉同一 Interceptor。

#### Agent Action

開 classic residual 前先列 chip FX vs window 兩表；hook 集合同時含 Create（FX）與 Show（window）。

#### Goal / Action / Validation

- Goal: 不誤判「classic 缺失」其實只是看錯 class。
- Action: dump 兩家族 → 分 hook → 分 JSON。
- Validation: inventory 明寫 Create-on-demand 與不同 GO／class。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP slots 同時有 chip FX 與 payout window。
- Does not apply: 單一 class 涵蓋所有派彩呈現的 cabinet。

#### Related

- `2026-09-24_114500-slot-base-line-wins-may-bypass-payout-window-show.md`
- `2026-09-24_113000-slot-payout-create-needs-path-rect-sound-and-settings-init.md`

#### Promotion Target

- apk-analysis slot UI：classic FX vs classic window 分路。
