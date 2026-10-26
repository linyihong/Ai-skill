> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-26 - 列表 FX API 可能在 StateInfo.Paid／SpinsCount 早退；filter 0 次 ≠ Where 濾光

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源為 Unity IL2CPP slots lab：列表 API 在 Paid＝true 或 SpinsCount＝0 時直接回空；predicate hook 0 次；lab 覆寫這兩欄後同路徑自然 `n>0`。普通付費轉仍會呼叫列表 API，但恒早退。具體柜／fixture 留專案 docs。

#### One-line Summary

當 feature 掉落／消失列表 API 回空且 **Where／predicate 從未進入** 時，先用 dump／disasm 查 **StateInfo（或等價）上的 Paid／剩餘 spins 早退閘**；不要假設「全被濾掉」或缺 SPINS reward。Lab 可暫時覆寫閘門證明列表＋Show 鏈；**普通付費轉呼叫列表 API 仍可能恒早退**，生產非空需未付 feature 態。

#### Human Explanation

列表驅動 FX 常先 `Get*Symbols`／`GetDrop*` 再建候選、再 LINQ Where。若早退在建候選之前（`Paid==true` 或 `SpinsCount==0`），filter MethodInfo 永遠不跑——看起來像「hook 錯 RVA」或「Where 濾光」。Presence settle／SendNext 後的 State 也常已是付費完結（Paid＋spins＝0），所以 force group／洞／SPINS inject 仍不夠。分開三件事：① 早退閘 ② 候選來源（Spin↔State reel-set 等）③ allowlist／predicate。

#### Trigger

- 列表 API `n=0` 且 DropFilter／Where hook **0 次呼叫**
- force group／洞／SPINS 注入後仍空
- lab 覆寫 Paid＝false＋SpinsCount≠0 後突然 `n>0`＋自然 Show
- 未 arm 的普通付費轉也進列表 API，但 dropGate 恒為 paid／spins＝0

#### Evidence

- Tool: Frida on list API `onEnter`／`onLeave`＋StateInfo 欄位快照；dump.cs／Instruction.parse 早退
- Sanitized excerpt: Paid＠field／SpinsCount＠field → return null before Where；force 兩欄 → 自然非空列表
- Evidence path: `<PROJECT_ROOT>/docs/slots/*/lab/` pass-early／setup-drop IL notes

#### Generalized Lesson

1. **Predicate 0 次 → 先查早退**：在 Where 之前找 State／round 閘（常見 Paid、剩餘 spins／series）。
2. **診斷標籤**：`paid_true`／`spinsCount_zero`／`pass_early`（通過閘）寫進列表 API 事件，避免只報 `n=0`。
3. **Lab force 閘門 ≠ 生產閉環**：覆寫 Paid／SpinsCount 只證明「通過閘後列表鏈可用」；文件標 lab-matched。
4. **普通付費轉可呼叫仍早退**：trace-only 多手若全 `paid_true`，下一刀改獵 **未付 feature 中段**（FS／respin in-progress），不要加普通轉次數假裝覆蓋。
5. **與 presence 空列表 lesson 疊加**：無消失元資料時列表可空；即使有閘通過，仍可能因 reel-set 差分／allowlist 得空——見 sibling reel-set lesson。

#### Agent Action

- 見 filter 0 次：讀 StateInfo（或 action 上的 state）Paid／spins，再決定是否 disasm。
- 報告分開：閘門 lab force vs 生產 unpaid 樣本。
- 勿把「普通轉有進 GetDrop」當成「已有 pass_early」。

#### Goal / Action / Validation

- Goal: 正確解釋列表 API 空回傳與 filter 沉默
- Action: 快照閘門欄位；必要時 lab 覆寫對照；生產改獵 unpaid feature
- Validation or reference source: 同 fixture force 閘 → `n>0`；non-force paid 轉 → 恒早退

#### Applies When

- Unity IL2CPP slots／feature 列表 API 依賴 StateInfo（或等價）付費／剩餘次數
- Lab presence replay 獵掉落／消失 FX

#### Does Not Apply When

- 列表 API 無 State 早退、純讀盤面洞／payload
- 已證實 Where 有進入且 keep＝false（真濾光）

#### Validation

- force Paid／SpinsCount → filter 開始有呼叫且 `n>0`；還原後早退
- 未 arm 付費轉：列表 API 可進、閘門標籤恒 paid／spins0

#### Promotion Target

- [Unity 動畫／特效驅動追蹤](../../../../analysis/apk/workflows/unity-fx-driver-trace.md) §2.1

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- category README 追加本條
- `analysis/apk/workflows/unity-fx-driver-trace.md` §2.1 補 Paid／SpinsCount 早退與 pass_early 診斷
- 與 [2026-10-26_113002…](2026-10-26_113002-lab-presence-fixture-empty-fx-list-vs-force-overlay.md) 交叉引用（不重複全文）
- 專案 evidence：`<PROJECT_ROOT>/docs/slots/*/lab/`

#### Transfer Contract

- Inputs: 列表 API hook、StateInfo 欄位、optional lab force
- Mechanism: Paid／spins 早退 → 無 Where → `n=0`
- Decision: 先閘門後 filter；生產要 unpaid feature
- Verification: force 閘對照＋診斷標籤
- Counterexample: Where 已進入的真濾光
- Transfer case: 其他「剩餘次數／付費旗標」門控的 feature FX 列表

#### Intelligence Review

- Decision: method-only
- Reason: 補強 FX 驅動追蹤的閘門診斷，不另抽 heuristic 檔
- Target: analysis/apk/workflows/unity-fx-driver-trace.md
- Revisit: none
- Discovery: not applicable — method-only

#### Abstraction Check

- Swap-project test: PASS — 任何以 State 付費／剩餘次數門控列表 FX 的 Unity lab／live 追蹤均適用
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/*/lab/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
