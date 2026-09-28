> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-28 - Payline CHIPS: strip line mirrors before scaling tags

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自一批 live payline spins：約半數 `@PAYOUT` 在「線估 + 全額 Σ CHIPS」下雙計或差額不對。

#### One-line Summary

Wire 常為每條線贏一筆 `<CHIPS>`；對獎時先按 lineWin 1:1 剝離镜像，剩餘標籤才當 extras，且 extras 可能要 ×lineBet 才對上服賠缺口。

#### Human Explanation

若把所有 CHIPS 都加成「獎勵」，線估已對平的局會出現大額負差額。若只加 residual 側、不剝離镜像，又會把線贏 CHIPS 與標籤混在一起。正確分流：镜像不计；标签计；纯标签局还要去掉 `AMOUNT==@PAYOUT` 的总额镜像（当 crumb 合计也等于 payout 时）。

#### Trigger

- `线估 ≈ @PAYOUT` 但 `Σ CHIPS ≈ @PAYOUT` → 差额 ≈ −@PAYOUT
- 有线赢时 leftover CHIPS 很小，缺口恰好是 leftover × lineBet
- 无线赢时同时存在总额 CHIPS 与 crumb 合计（双份）

#### Evidence

- Tool: batch recompute of live payline spins vs stored Rewards
- Sanitized excerpt: after stripping per-line CHIPS mirrors and scaling leftover labels by lineBet (when that hits the payout gap), residual went to zero across the positive-payout sample.
- Evidence path: keep cabinet-specific tables under `<PROJECT_ROOT>` slot docs.

#### Generalized Lesson

1. Never add raw Σ CHIPS on top of verified line Σ without partitioning.
2. Prefer multiset match of CHIPS amounts to paying-line estimates as mirrors.
3. For leftovers, try credit scales `{1, lineBet, lineBet×chipScale}` and pick the one that closes `server − lineSum`.
4. Pure-tag: if a total-mirror equals server and other crumbs also sum to server, keep one side only.

#### Agent Action

对奖差额非零时，先打印 lineWins[] 与 CHIPS[]，检查是否 1:1 镜像或 label×lineBet，再怀疑线几何。

#### Goal / Action / Validation

- Goal: `server − lineEst − creditedTags = 0` for regular (non-jackpot/FS) spins.
- Action: partition + scale in shared payline analyze.
- Validation or reference source: batch residual histogram → all zero on sample.

#### Applies When

- winMode=paylines and Rewards use multiple `<CHIPS AMOUNT>` crumbs.

#### Does Not Apply When

- Single CHIPS total with no per-line crumbs and no separate tag path.
- variable-lines / collapse cabinets that already fold chips in their own process.

#### Validation

Re-run analyze over a payout>0 sample; residual must be 0 except documented jackpot/FS/activity.

#### Promotion Target

- `workflow/development-guidance/`（若日後整理对奖 checklist）

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

N/A — candidate only.
