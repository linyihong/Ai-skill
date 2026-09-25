> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-15 - One Collapse level can pay multiple symbols concurrently

Status: candidate

#### One-line Summary

同一 Rewards `level` 的 chips 可能對應**多個**同時中獎符號（不同 per-way 金額混在同一 chip 列表）；先把 chip multiset partition 成多組 ways，再對所有中獎格做一次 union tumble——不要假設「一 level = 一 symbol」。

#### Human Explanation

多 level Collapse sim 若只挑「ways 數 == chip 數且金額一致」的單一符號，會在某一層留下 `no-ways-match`，總額低於 Payout。實況是該層盤面可同時形成兩條（或以上）variable-lines 線：chip 列表按 per-way 金額分組（例如部分 1000、部分 2400），各組對應不同 symbol×length。Tumble 時用中獎格集合聯集移除（共用 WILD 只刪一次），再套該層 `NewSymbols`。

#### Trigger

- Multi-level sim matches early levels then fails with mixed chip amounts in one level.
- `Σ chips == Payout` but simulated multi-level total < Payout and failed level has heterogeneous chip values.

#### Evidence

- Tool: chip-multiset partition + concurrent ways + union tumble.
- Sanitized observation: after level-1 tumble, level-2 chips split into two per-way amounts matching two symbols on the intermediate board; sum with level 1 equaled Payout.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/<cabinet>/win-calc.md` concurrent Collapse-level sample.

#### Generalized Lesson

1. Per Collapse level, treat chips as a multiset to partition across candidate symbols—not a single uniform ways count.
2. When multiple symbols match, tumble the union of winning cells (shared wild once).
3. Prefer this explanation before inventing new paytable rules or assuming a missed RESULT.

#### Agent Action

Extend multi-level recompute to concurrent pays per level; keep documenting level tables with multi-symbol rows when chips are heterogeneous.

#### Goal / Action / Validation

- Goal: close ways/Payout gaps when one reward level has mixed chip amounts.
- Action: partition chips by per-way amount → multi-symbol hits → union tumble.
- Validation: simulated Σ == Payout on a live concurrent-level sample; prior single-symbol multi-level samples still match.

#### Related

- `2026-09-14_164000-slot-multi-level-collapse-one-result.md`
- `2026-09-14_162000-slot-rewards-chips-per-way-and-collapse.md`

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
