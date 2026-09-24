> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-24 - Device UI harvest via owned cryptor lab gateway + fixture loop

Status: validated

#### One-line Summary

要可控重放 slot 特效／畫面時，用**自有 lab game gateway**（與 client 相同 AES framed 協定）反喂去敏 fixture，並用 Frida 只改寫 game TCP `connect` 到 lab；不要對正式站做 MITM，也不要用 empty-secret stub 假裝已覆蓋裝置路徑。

#### Human Explanation

稀有窗／動畫若只等 live RNG，成本高。較穩的閉環是：

1. 把已採集、去敏的 RESULT／SPIN 形 fixture 存進 playlist。
2. 起一台 **lab-owned `connectionSecret`** 的 gateway：解密 C2S、在 `SLOTS_BET`（等）回加密 S2C。
3. 控制口可 `loop` 同一 fixture，重複觸發同一盤面／動畫路徑。
4. 裝置端用 Frida 把 `:2043` `connect` 指到 lab（CDN/HTTPS 仍走正式）。
5. Cold start 進 lab session；**不要**混用正式站 session key。

Plaintext push-only stub（empty secret）只適合 offline client 驗線；裝置真機需要 cryptor 對稱路徑。

#### Trigger

- 想「同一結果打一百次」抓特效，卻還在燒真金等 RNG。
- 只做 Frida apply-shim 改 Symbols，卻希望「整段協定／多種 S2C」都由我們控。
- 試圖 MITM 正式 GAMESERVER 密文（越權、易碎、且違反 lab 邊界）。

#### Evidence

- Tool: Node TCP gateway + AES helpers + Frida libc `connect` redirect + fixture playlist／control HTTP。
- Sanitized pattern: lab secret；`PAYOUT=0`；log 只打 verb + fixture id；session 標 `replay-lab`。
- Evidence path: project `pokerist-client` lab gateway + cabinet replay-lab docs（project-local）。

#### Generalized Lesson

1. **UI harvest 優先自有 gateway + fixture loop**，不是正式站注入。
2. **裝置路徑必須 cryptor 對稱**；plaintext stub ≠ 真機協定覆蓋。
3. **只導 game TCP**；CDN／HTTPS 可仍官方。
4. **lab secret 與正式 session 隔離**；cold start；secret 不進 git。
5. Frida apply-shim 可並存（改記憶體物件），但不能取代「可控 S2C」閘道。

#### Agent Action

做裝置動畫／稀有窗反查時：先落地 cryptor lab gateway + fixture arm/loop + connect redirect；用 desk-match 驗 GLASS；再上機 smoke。禁止 production MITM 與實帳派彩偽造。

#### Goal / Action / Validation

- Goal: 可控重放採集結果以反查 UI／動畫。
- Action: gateway BET→fixture；loop；Frida redirect；去敏。
- Validation: client 測試 round-trip desk-match；裝置至少一轉 fixture 盤面（presence-only）。

#### Applies / Does Not Apply

- Applies: 自有 cryptor 已知的 Unity／自訂 TCP slot session；已有去敏 fixture。
- Does not apply: 未知加密、禁止改裝置網路、或目標是攻擊正式經濟系統。

#### Related

- `2026-09-24_085100-slot-ui-harvest-lab-fixture-apply-shim-not-live-payout-forge.md`
- `2026-09-21_111500-slot-celebration-window-hunt-continue-cta-not-midreel-gold.md`
- `2026-09-24_135500-slot-continue-cta-scorer-needs-matte-green-and-left-band.md`

#### Promotion Target

- apk-analysis slot capture SOP：lab gateway／fixture loop／device redirect 邊界。
