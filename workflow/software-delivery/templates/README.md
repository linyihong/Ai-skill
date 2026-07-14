# Software Delivery Templates

Use these templates as focused artifact shapes. Load only the template that matches the current artifact.

| Template | Use when |
| --- | --- |
| [`change-brief-template.md`](change-brief-template.md) | Capturing intake, scope, assumptions, acceptance, and validation target before implementation. |
| [`contract-template.md`](contract-template.md) | Defining domain/API/UI/consumer contracts and traceability before parallel implementation. |
| [`bdd-scenario-template.md`](bdd-scenario-template.md) | Writing behavior scenarios and acceptance examples. |
| [`implementation-plan-template.md`](implementation-plan-template.md) | Planning implementation slices, validation, and same-session closure. |
| [`review-report-template.md`](review-report-template.md) | **`code-review` capability output** after post-implementation invoke — not Validation phase output. Consumer: [`cross-cutting/review/self-review.md`](../../cross-cutting/review/self-review.md). |
| [`product-impact-alignment-template.md`](product-impact-alignment-template.md) | Aligning product impact, journey evidence, assumptions, and acceptance. |
| [`ui-governance-evidence-template.md`](ui-governance-evidence-template.md) | Classifying UI compliance evidence by governance domain, collection method, validation mechanism, evidence class, severity, and project-local design-system policy. |
| [`ui-pattern-knowledge.entry.template.yaml`](ui-pattern-knowledge.entry.template.yaml) | Core vs Extended pattern entry（selection_rules、family、near_neighbors；recipe 可 unknown）。 |
| [`ui-pattern-knowledge.composition.template.yaml`](ui-pattern-knowledge.composition.template.yaml) | Screen → contains[] pattern composition（optional flag）。 |
| [`ui-pattern-prompt-expansion.template.md`](ui-pattern-prompt-expansion.template.md) | Transient prompt-expansion checklist；勿 commit 進 canonical Knowledge。 |

Do not merge UI governance evidence back into `contract-template.md` unless the artifact is defining expected UI behavior. Compliance evidence belongs in the focused UI governance template.
Do not treat prompt-expansion checklists as Pattern Knowledge entries; seeds live under [`../ui-pattern-knowledge/`](../ui-pattern-knowledge/README.md).
