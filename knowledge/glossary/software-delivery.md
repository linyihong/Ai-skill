# Software Delivery Framework Glossary

> Domain glossary for Software Delivery Framework Primary Model（N = 3）。  
> Platform-wide terms remain in [`ai-skill.md`](ai-skill.md)。Resolution priority：[`README.md`](README.md) §Vocabulary Resolution Priority。

上游：[`plans/active/2026-07-16-0945-software-delivery-framework-domain-model/_plan.md`](../../plans/active/2026-07-16-0945-software-delivery-framework-domain-model/_plan.md)

---

## delivery_artifact

```yaml
term: delivery_artifact
status: candidate
owner-layer: workflow-orchestration
meaning: >
  A project-layer instance of a framework_asset produced to complete a delivery
  intent. Not a fourth primary domain concept; a lifecycle/tag on Asset at
  project scope (e.g. change brief, BDD file, evidence run output).
affects:
  - workflow/software-delivery/README.md
  - workflow/software-delivery/domain-policies.md
  - plans/active/2026-07-16-0945-software-delivery-framework-domain-model/_plan.md
aliases:
  - artifact
anti-meaning: >
  Not synonymous with framework_asset. Templates and workflow slices are
  framework assets, not delivery artifacts.
introduced-by: plans/active/2026-07-16-0945-software-delivery-framework-domain-model/_plan.md
related-terms:
  - { type: derived_from, target: framework_asset }
  - { type: related_to, target: framework_asset }
```

## delivery_policy

```yaml
term: delivery_policy
status: candidate
owner-layer: workflow-orchestration
meaning: >
  Operational rules governing framework_asset classes: ownership, placement,
  promotion, authority precedence, knowledge boundary, lifecycle. Derived from
  governing_principles; canonical body in workflow/software-delivery/domain-policies.md.
affects:
  - workflow/software-delivery/domain-policies.md
  - workflow/software-delivery/contracts.md
  - workflow/software-delivery/artifact-gates.md
anti-meaning: >
  Not the same as metadata/recovery domain policies or platform enforcement
  rules. Not governing_principles (which guide policy creation).
introduced-by: plans/active/2026-07-16-0945-software-delivery-framework-domain-model/_plan.md
related-terms:
  - { type: derived_from, target: governing_principle }
  - { type: related_to, target: delivery_process }
  - { type: related_to, target: framework_asset }
  - { type: related_to, target: governing_principle }
```

## delivery_process

```yaml
term: delivery_process
status: candidate
owner-layer: workflow-orchestration
meaning: >
  Temporal sequencing of software delivery work: which Process stage typically
  produces or consumes which asset class. Canonical surfaces include
  execution-flow.md and sd-* cognitive slices. Process does not own asset
  definitions or placement rules.
affects:
  - workflow/software-delivery/execution-flow.md
  - workflow/software-delivery/README.md
  - governance/cognitive-slice-taxonomy.md
anti-meaning: >
  Not a list of what may be created in a stage; describes typical
  produce/consume timing only.
introduced-by: plans/active/2026-07-16-0945-software-delivery-framework-domain-model/_plan.md
related-terms:
  - { type: related_to, target: delivery_policy }
```

## framework_asset

```yaml
term: framework_asset
status: candidate
owner-layer: workflow-orchestration
meaning: >
  Any managed entity in the Software Delivery Framework domain: knowledge,
  contracts, evidence, decisions, templates, process components, runtime
  contracts, capabilities, etc. Primary domain concept (N=3). Asset class
  taxonomy is an ontology view over framework_asset, not a fourth core.
affects:
  - workflow/software-delivery/README.md
  - workflow/software-delivery/domain-policies.md
  - plans/active/2026-07-16-0945-software-delivery-framework-domain-model/evidence/phase-0-classification-matrix.md
aliases:
  - asset
anti-meaning: >
  Not limited to project deliverables or markdown files. Workflow YAML and
  cognitive slice definitions are framework assets at framework layer.
introduced-by: plans/active/2026-07-16-0945-software-delivery-framework-domain-model/_plan.md
related-terms:
  - { type: related_to, target: delivery_policy }
  - { type: related_to, target: delivery_artifact }
```

## governing_principle

```yaml
term: governing_principle
status: candidate
owner-layer: workflow-orchestration
meaning: >
  Abstract belief that guides delivery_policy creation (e.g. One Asset One
  Owner, Classification before Creation). Meta-layer above Policy; not
  operational policy text and not a fourth primary domain concept alongside
  framework_asset, delivery_policy, delivery_process.
affects:
  - workflow/software-delivery/domain-policies.md
  - workflow/software-delivery/README.md
anti-meaning: >
  Not a delivery_policy entry. Not stored as enforceable gate text without
  operationalization in domain-policies.md.
introduced-by: plans/active/2026-07-16-0945-software-delivery-framework-domain-model/_plan.md
related-terms:
  - { type: related_to, target: delivery_policy }
```
