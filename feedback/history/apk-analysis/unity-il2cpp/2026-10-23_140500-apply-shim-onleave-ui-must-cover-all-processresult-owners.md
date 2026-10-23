> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-23 - Apply-shim onLeave UI must cover all ProcessResult owners

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源為多柜共用 Frida apply-shim；具體柜名留專案 lab。

#### One-line Summary

多柜共用的 Frida `ProcessResult` Interceptor 若只在某一 owner 的 `onLeave` 跑 line-chip／celebration UI，其他柜會出現 `LAB_APPLIED`（Symbols／Payout 已 patch）卻永遠沒有 `LAB_LINE_CHIP_FX`／forcePaint。

#### Human Explanation

Apply-shim 常在 `onEnter` 對所有柜的 `ProcessResult` 寫 Symbols／Rewards／Payout，但 stampede／FS／line-chip 等「結果後 UI」習慣放在 `onLeave`。若 `onLeave` 用 `owner === OneCabinet` 閘住，第二、三柜只拿到盤面 patch，線獎／芯片 FX 路徑從不啟動——易被誤判成「Rewards 注入失敗」或「該柜沒有 LinePayoutModule」。

先確認：`LAB_APPLIED` 是否出現、`onLeave` 是否對該 owner 呼叫 `tryLabUi`／等價函式、再查 Rewards dict 與 `FindLinePayoutModule`。

#### Trigger

- 柜 A 的 apply-shim 已能 `LAB_LINE_CHIP_FX`，柜 B 只有 `LAB_APPLIED`
- `openLineChipFx`／forcePaint fixture 旗標已開，仍無 line-chip 事件
- 程式碼搜尋 `onLeave` 只比對單一 `owner` 字串

#### Evidence

- Tool: Frida apply-shim Interceptor on `ProcessResult` enter/leave
- Sanitized excerpt: enter 對多 owner hook；leave 僅 `OneCabinet` → 他柜缺 `LAB_LINE_CHIP_FX`
- Evidence path: `<PROJECT_ROOT>/docs/lab/` apply-shim 笔记；柜机 lab README

#### Generalized Lesson

1. **Apply 與 UI 分閘**：`onEnter` patch 與 `onLeave` celebration 的 owner 集合必須一致（或 leave 用「任一已 hook 的 ProcessResult」）。
2. **新柜注册 hook 时**：同时检查 leave UI 闸，不要只加 enter 的 `maybeApply`。
3. **排障顺序**：`LAB_APPLIED` → leave 事件／`LAB_LINE_CHIP_FX` → Rewards dict → LinePayoutModule 查找。
4. **Rewards null ≠ 无 FX 路径**：forcePaint／ShowPreview 可在 leave 补救；前提是 leave 有跑。

#### Agent Action

- 扩展共用 shim 时：对所有 `ProcessResult` owners 跑 leave UI，或抽 `isResultApplyOwner(h)`。
- 复测：第二柜 fixture `openLineChipFx=true` 必须出现 line-chip 事件（或明确模块查找失败，而不是静默）。
- 不要把单一柜的 stampede-only leave 闸当成全库默认。

#### Goal / Action / Validation

- Goal: 多柜 apply-shim 在 Symbols patch 后也能进入线奖 FX 路径
- Action: leave UI 覆盖所有 ProcessResult owners
- Validation or reference source: 第二柜 log 含 `LAB_LINE_CHIP_FX` 或等价步骤表

#### Applies When

- 共用 Frida apply-shim / lab host 挂多个 `Slots*.ProcessResult`
- 线奖／chip／celebration 依赖 ProcessResult leave

#### Does Not Apply When

- 仅单柜专用 script，无多 owner 共享 Interceptor
- FX 完全由原生 Rewards 驱动且不走 leave forcePaint

#### Validation

- 至少两个 owner：enter 都有 `LAB_APPLIED`，leave 都有 UI 事件或明确失败步骤
- Code review：`onLeave` 无单一柜硬编码闸（除非注释说明 stampede-only）

#### Promotion Target

- `feedback/history/apk-analysis/unity-il2cpp/`（本条）
- 专案 apply-shim 注释／lab README（已／应同步）

#### Promotion Record

尚未 promotion；保留为 candidate history。

#### Required Linked Updates

- 专案 apply-shim：leave 覆盖多 ProcessResult owner（本轮已改）
- category README 索引追加本条

#### Abstraction Check

- Swap-project test: PASS — any multi-cabinet Frida ProcessResult apply-shim must run leave UI for every hooked owner, not only the first cabinet that needed stampede
- Project evidence kept in: `<PROJECT_ROOT>/docs/` cabinet lab notes（apply-shim A/B）
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
