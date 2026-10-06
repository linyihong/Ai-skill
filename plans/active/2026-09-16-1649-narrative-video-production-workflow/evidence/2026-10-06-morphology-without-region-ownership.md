# Morphology is not region ownership

Status: adapter regression confirmation; full pack dogfood remains open.

## Finding

Substring platform vocabulary can occur inside ordinary dialogue, and a
speaker-name colon can occur in a genuine caption. Neither establishes a
platform UI or chat-panel owner. A classifier that hard-rejects these shapes
without region evidence reintroduces evidence loss even if watermark and
English retention tests already pass.

## Repair boundary

Use bounded platform-label matches rather than broad dialogue substrings.
Chat exclusion requires independent panel-layout evidence. Caption-like
geometry and relative typography can support speaker-labelled dialogue;
text-only conflicting morphology stays uncertain instead of being deleted.
Retain chat-panel and platform-label negatives alongside dialogue positives.
No global OCR frequency, ASR authority or temporal policy change is justified.

## Validation

Synthetic paired adapter checks reproduce the original false negatives and
pass after the scoped repair in the runtime product. This establishes only
the exercised classification boundary, not real-frame ownership accuracy,
locale projection, ASS layout, rendered quality or published-package quality.
Those gates remain separate and the complete dogfood checkbox stays open.

Release review also needs the monitor's window-scoped callable contract to
remain available. A green isolated classifier does not cover the integration
path that asks whether evidence exists in a requested time window.

## Open validation

- Real source-frame positive/negative pairs across representative layouts.
- Full regression replay and consumer identity/time propagation.
- Actual package execution, source audit, service activation and rendered QA.

Product paths, raw fixtures, test output and deployment status remain in
`<PROJECT_ROOT>`; no incident-specific identifiers are promoted here.
