> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Multi-level Collapse can pay several cascades inside one RESULT

Status: validated

#### One-line Summary

當 `Rewards` 出現多個 `level`（各含 chips + 可選 `SlotsRewardCollapse.NewSymbols`）時，付費 cascade 可能全部序列化在**同一個** `ProcessResult`：先算 level N ways，再依 NewSymbols tumble，再算 level N+1；不要只因單一 Symbols 快照 ways < Payout 就改公式。

#### Human Explanation

Chips-per-way 與單階段 ways 對齊後，仍可能看到 `Σ chips == Payout` 但單次 ways 只覆蓋第一段。原因是字典／level 鍵代表 cascade 層次：每一層付完後用該層 Collapse 的 NewSymbols 在客戶端填格，下一層 ways 在中間盤面計算，卻仍打包在同一 RESULT。Tumble 驗證模型：移除該層中獎格 → 倖存符號沉底 → NewSymbols（依 Row）置頂。多 RESULT 付費鏈仍可能存在，但多 level Collapse 常已解釋「看似漏段」的 Payout。

#### Trigger

- Single-pass ways < Payout 但 `Σ SlotsRewardChips == Payout` 且 Rewards 有 ≥2 distinct `level`。
- Collapse present yet idle window shows only one ProcessResult.

#### Evidence

- Tool: Rewards level walk + tumble sim + ways recompute.
- Sanitized observation: two reward levels in one RESULT; level-1 ways matched first chip group; after tumble, level-2 ways matched second chip group; sum equaled Payout.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/<cabinet>/win-calc.md` multi-level Collapse sample.

#### Generalized Lesson

1. Group Rewards by `level` before concluding ways/Payout mismatch.
2. If ≥2 levels with Collapse NewSymbols, simulate tumble between levels; compare each chip group to that board’s ways.
3. Do not rewrite the base ways formula when multi-level Collapse explains the gap.
4. Keep paid multi-RESULT as a separate hunt; multi-level Collapse is not proof of multi-RESULT.

#### Agent Action

Extend recompute to optional multi-level Collapse simulation; document level table next to Payout; only escalate to “missed RESULT” after levels fail to explain.

#### Goal / Action / Validation

- Goal: explain Payout > single-pass ways when Collapse levels > 1.
- Action: level-grouped Rewards + tumble (bottom survivors / top NewSymbols) + per-level ways.
- Validation: simulated Σ level amounts == Payout on a live multi-level sample.

#### Related

- `2026-09-14_162000-slot-rewards-chips-per-way-and-collapse.md`
- `2026-09-14_163000-slot-collapse-often-single-result.md`
