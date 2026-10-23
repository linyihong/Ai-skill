> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-23 - Apply-shim visibleGrid must pad ProcessResult strip

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源為 Frida apply-shim + Unity ProcessResult strip lab。

#### One-line Summary

Lab fixture 只有可見窗 `visibleGrid`（例如 5×3）時，apply-shim 必須依 `layout.returnedRowRange` 展開成與 live `ProcessResult` 相同的全高 strip，並補齊上游缺格／`"?"` 洞；否則桌上會留灰空格，易被誤判成 cascade「爆掉未補落」。

#### Human Explanation

許多柜的 `SpinInfo.Symbols` 是帶上下緩衝的 strip（常見 Row −2..4＝35 cells）。Fixture 為 desk-match 方便只存可見 15 格。若 shim 只寫 `reel:0..2`，off-window 與 tall-wild 省略格保持上游空洞／錯誤 ID，UI 就出現「上面有符號、下面兩格空」——在 paylines 柜上這不是合法 cascade 行為。

Wire RESULT builder 已有 pad；apply-shim 必須對齊同一契約。Celebration／line-chip 後可再 `LAB_REPATCH` 一次，避免 FX 清格後空洞殘留。

#### Trigger

- apply-shim `LAB_APPLIED` 後可見區出現灰空格
- fixture 只有 `visibleGrid`，`layout.returnedRowRange` 寬於可見窗
- log `nSym`＜預期 strip，或 `patched` 遠小於 strip cell 數

#### Evidence

- Tool: Frida apply-shim `setFixture`／`patchSpin`
- Sanitized excerpt: pad 後 `cells=35`、`patched=35`、`missing=[]`；pad 前可見區灰空
- Evidence path: `<PROJECT_ROOT>/docs/` cabinet lab apply-shim notes

#### Generalized Lesson

1. **同一 strip 契約**：wire builder 與 apply-shim 對 `returnedRowRange` 的 pad 規則必須一致。
2. **缺格要補**：上游 tall-wild 常省略 Symbols 元素；fixture 要實 ID 時需 grow 陣列或覆寫 `"?"`，不能只 skip。
3. **排障**：先分 paylines vs collapse；灰空＋paylines → 優先查 pad／missing，不是「補落動畫沒播」。
4. **FX 後重寫**：line-chip／celebration 後排程短延遲 re-patch。

#### Agent Action

- 读 fixture `layout.returnedRowRange`（或等价）再展开 `visibleGrid`
- `patchSpin` 回报 `missing`；非 `"?"` 缺格则 grow `Symbols[]`
- 有 openLineChipFx 时 schedule re-patch

#### Goal / Action / Validation

- Goal: lab 强制盘面可见区无意外灰空
- Action: pad strip + grow holes + optional re-patch
- Validation or reference source: `cells`／`patched` 对齐 strip；截图可见 5×3 满格（tall `"?"` 除外）

#### Applies When

- Frida／IL2CPP apply-shim 写 `SpinInfo.Symbols`
- Fixture 以 `visibleGrid` 为主且 strip 高于可见窗

#### Does Not Apply When

- 柜契约本就是可见窗 only
- 故意 sparse／巨像 `"?"` 空洞（pad 会破坏 Spine）

#### Validation

- 同 fixture：pad 前后 `patched`／可见空格对比
- 故意 `"?"` 洞仍保留，不被 edge-clone 填满

#### Promotion Target

- `feedback/history/apk-analysis/unity-il2cpp/`（本条）
- 与 wire GLASS strip lesson 互补

#### Promotion Record

尚未 promotion；保留为 candidate history。

#### Required Linked Updates

- 专案 apply-shim 已 pad／grow／re-patch
- category README 索引追加本条
- 关联 wire lesson：`local-proxy/…wire-result-glass-must-match-processresult-strip`

#### Abstraction Check

- Swap-project test: PASS — any Unity ProcessResult strip cabinet whose lab fixture ships visibleGrid-only needs the same pad/grow contract in apply-shim as on the wire builder
- Project evidence kept in: `<PROJECT_ROOT>/docs/` cabinet lab notes
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
