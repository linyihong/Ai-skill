> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Collapse NewSymbols often client-applied without 2nd RESULT

Status: validated

#### One-line Summary

`SlotsRewardCollapse.NewSymbols` 常在**同一** `ProcessResult` 內交付；客戶端 tumble 後不一定再來一次 RESULT。要抓多段付費 cascade，hook 需持續聽多個 RESULT，但不要假設每次 Collapse 都有第二段 server 盤面。

#### Human Explanation

單階段 ways 與 `Σ SlotsRewardChips` 已可對上 Payout，同時 Rewards 裡仍可帶 Collapse。實機常見：只有一個 RESULT，Collapse 帶 NewSymbols，idle 視窗內無第二段。另一樣本出現兩個 RESULT 但皆零赔，證明鏈路可用，卻不是付費 cascade。歷史 under-count 更可能是舊 `once` hook 漏掉後段，而非 ways 公式錯誤。

#### Trigger

- Rewards 含 Collapse 且要解釋是否還有後續 RESULT。
- Designing multi-stage capture idle timeout.

#### Evidence

- Tool: multi-RESULT Frida chain + ways/Rewards recompute.
- Sanitized observation: paid spin matched ways with 16 chip rows and Collapse newSymbolCount>0 in one stage; another spin had two zero-pay RESULT stages.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/slots/<cabinet>/win-calc.md` cascade section.

#### Generalized Lesson

1. Record `cascadeStages[]` until idle; primary grid = last stage.
2. Do not require a 2nd RESULT to acknowledge Collapse.
3. Treat paid multi-RESULT as optional evidence; keep hunting separately from single-stage Collapse proof.
4. Prefer chain capture over inventing pays when historical Payout > ways(single snapshot).

#### Agent Action

Use idle-based multi-RESULT capture; document single-stage Collapse as normal; keep open item for paid multi-stage.

#### Goal / Action / Validation

- Goal: correct cascade capture expectations.
- Action: chain hook + fixture cascadeStages; sample with Collapse-only single RESULT.
- Validation: Collapse + chips + ways exact on one stage; chain plumbing shown on ≥2 RESULT sample.

#### Related

- `2026-09-14_162000-slot-rewards-chips-per-way-and-collapse.md`
