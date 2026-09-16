# Matching script

一條成片一份。結構化檔是 SoT；Markdown 是人讀投影（同 EDR）。
Invariant 4–5：Constraint ≠ Selection。

填 [`records/matching-script.yaml`](records/matching-script.yaml)。概念閉環示範：[`records/examples/sanitized-matching-and-edr.yaml`](records/examples/sanitized-matching-and-edr.yaml)。

## 每個需要畫面的 shot

```text
need
  → constraints (must_tags / entity_refs / duration_target / duration_tol / duration_band / continuity)
  → feasible candidates[]
       clip_id, matched_constraints, duration_delta, rejection_reasons
  → selection:
       policy: duration_closest | preserve_character_continuity | human_review | …
       rationale: …
  → selected_clip_id   # 必須 ∈ feasible candidates
```

- **Constraint** 決定誰進可行集。未進集者不得被選。
- **Selection policy** 才決定選誰。`duration_closest` 可以是一條 policy，**不是系統默認的「最好」**。
- 空可行集 → `blocked`：放寬約束、補 catalog、或停。禁止從集外硬挑。
- 選中 clip **違反約束**（不是「不是最近」）→ QC 失敗。
- 未來模型可當更強 Selection Actor；本契約不改。

頭欄：`bible_id`、`catalog_id`、`narrative_template_id`、目標時長、發布語。

通過後才把實際入出點寫入 EDR；EDR 可微調窗但須仍指向同一 `clip_id`（或 `mutations[]` 換 clip）。
