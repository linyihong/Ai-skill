> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-24 - Slot UI harvest lab: post-decrypt fixture apply-shim, never forge live payouts

Status: candidate

#### One-line Summary

要重複拍 slot 慶祝／FS 窗時，用**自有 fixture + 解密後／apply 前 Frida shim**（或自有 gateway）驅動 UI；不要 MITM 正式服偽造派彩。Parse RESULT 常有空 `Symbols`；`set_Count` 等 managed 參數 ABI 必須對齊（例如 `Int64` 勿當 `Int32`）。

#### Human Explanation

Live RNG 很難精準重現 FS stampede、BHS、Flash 等短窗。正確 closed-loop：

1. 從**曾 live 過**的 settle 做成去敏 fixture（符號 ID + reel/row；無餘額／派彩）。
2. 在 **ProcessResult／SPIN apply**（有 `nSym>0`）覆寫 `SpinInfo.Symbols` ID。
3. 需要開窗時，在**遊戲執行緒**（例如 ProcessResult `onLeave`）呼叫 Create／Setup／Show；JS `setTimeout` 或錯執行緒常 `breakpoint`／無效。
4. 子類 `Process*` 可能是空 `RET` stub——改呼叫基類實作。
5. 呼叫 managed setter 前查 IL2CPP 參數型別；`set_Count(Int64)` 用 `int32` NativeFunction 會弄壞 ABI、軟鎖旋轉。

Boundary：**presence-only**；標 `replay-lab`；正式服派彩權威不變。

#### Trigger

- 想「每次旋轉都出 stampede」卻去改正式 wire／cryptor。
- 在 `Parse` `onLeave` patch，RESULT 的 `Symbols` 仍 null，`patched:0`。
- Create window 回傳 ok 但無畫面；或 Show 後無法再旋轉。
- 子類 `ProcessStartFreeSpin` Interceptor／呼叫「成功」但 prologue 是 `RET`。

#### Evidence

- Tool: Frida IL2CPP apply-shim + host screencap burst（no Continue during target window）。
- Sanitized pattern: fixture grid → apply patch → Create/Setup/Show → stampede atlas；post-dismiss spin still spends bet。
- Evidence path: `<PROJECT_ROOT>` cabinet `replay-lab-plan`／`lab/`／redacted `replay-lab-*-live.png`（project-local）。

#### Generalized Lesson

1. **UI harvest lab ≠ live forgery**：fixture／自有入口；禁止偽造正式 payout。
2. **Patch 時機 = apply 面**：等 `Symbols` 陣列非空（ProcessResult／EventSink／ProcessSpin），不要死守 Parse leave。
3. **開窗在遊戲執行緒**：Process* `onLeave` 呼叫 Create／Show；避免錯執行緒。
4. **先看 prologue**：子類空 stub 就改基類。
5. **NativeFunction ABI 對齊 IL 型別**：`Int64`／valuetype／bool 傳錯會軟鎖。
6. **入庫前去敏**並標 `replay-lab`。

#### Agent Action

先做 fixture + apply-shim MVP；開窗前核對 method 參數型別與 stub；Show 後用一次真實旋轉確認未 soft-lock。

#### Goal / Action / Validation

- Goal: 可重播目標窗（如 FS stampede）供截圖／對桌，不碰正式派彩。
- Action: fixture → arm → spin → patch → Create/Setup/Show → burst → redact。
- Validation: `patched≈cellCount`；burst 見目標窗；dismiss 後仍可扣注旋轉。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP slot cabinets；自有裝置／自有 lab。
- Does not apply: 未授權環境；要證明 live RNG 權重／RTP 的場合（仍以正式服為準）。

#### Related

- `2026-09-21_111500-slot-celebration-window-hunt-continue-cta-not-midreel-gold.md`
- `2026-09-21_091600-slot-three-visual-planes-and-in-cabinet-anim-stack.md`
- `2026-09-14_170500-frida-chain-harness-nonblocking-stdout.md`
- `2026-09-08_165800-slot-spin-fixtures-pair-json-manifest-and-reel-crop.md`

#### Validation

- 在自有 lab 以 fixture replay 驗證 patch 命中、目標窗可見，且 dismiss 後真實旋轉仍可完成；不得以正式派彩或帳務結果作為驗證。

#### Promotion Target

- apk-analysis slot capture SOP：fixture replay lab／apply-shim 邊界與 ABI 檢查。

#### Reuse Evidence

- 尚未有獨立 cabinet／專案的重用證據；保留為 candidate，等待以相同 safety boundary 完成第二次驗證。

#### Promotion Record

- 尚未 promotion；僅保留為 feedback history candidate。

#### Required Linked Updates

- 已更新 `unity-il2cpp/README.md` category index；尚未更新 capture SOP，因尚缺獨立重用證據。
