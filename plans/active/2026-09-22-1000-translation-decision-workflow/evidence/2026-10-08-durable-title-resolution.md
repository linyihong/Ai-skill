# Durable title resolution is identity-scoped, not a fixed glossary

## Observation and cause

An export manifest or short-lived search-result cache does not guarantee reuse
of a selected localized title. If title realization runs before export-cache
lookup, repeated episodes/jobs can generate different names despite successful
media reuse. Country/language labels duplicated outside the active locale
catalog can drift independently from the actual selected target constraint.

## Adapter boundary

- Bind durable selections to stable source entity identity and existing target
  language codes. A normalized source-name fallback is weaker identity;
  never merge distinct known IDs solely because their names match.
- Keep all localized selections and catalog-derived language labels in one
  inspectable registry. Source title, normalized aliases, selected title,
  provenance, admission status and validation scope remain distinct fields.
- Reuse admitted selections before catalog lookup or model realization.
  Same-language identity is also a resolution path, not an excuse to skip
  target orthography. Custom/deprecated locales must not silently alias into
  an unrelated active target.
- Store core title separately from episode/part suffixes and marketing tags.
  Upper/lower presentation must not alter another episode's selected identity.
- Validate fresh and stored outputs. Do not freeze empty, wrong-script,
  contract-failed or generic placeholder fallbacks as accepted names.
- Serialize read→resolve→write across actors; atomic replacement and retained
  corrupt-file evidence prevent lost updates and silent registry destruction.
- A local-only constraint applies to realization, retries and validation
  actors even when cloud credentials exist. A registry hit needs no provider.

This is a per-entity selected-output registry, not a cross-expression fixed
translation glossary. A retrieval hit is not a new semantic decision, and
mechanical admission is not independent semantic or full Finality PASS.
First-title semantic quality, human overrides, shared-device synchronization
and historical-name migration need separately scoped evidence.

## Target orthography limitation

Native script conversion can pass common tests while leaving individual
Traditional characters in a default Simplified deliverable. Representative
many-to-one characters and phrase exceptions should exercise fresh realization,
registry storage and cache reuse, not only script-family detection.
Pinned character/phrase resources can repair this boundary without a provider
call; preserve resource provenance/license and distinguish this subset from a
complete regional vocabulary engine. Never normalize raw acquisition evidence.

## Verification and plan disposition

Consumer fixtures cover active catalog projection, disk reload without actors,
cross-process single resolution, known-ID isolation, part separation, invalid
candidate non-admission, write/corrupt-file retention and local-only negative
paths. Existing derived-output and title-contract regressions remain green.
An old grounding fixture accidentally invoked a model; isolate provider actors
explicitly so offline validation cannot become an unbounded runtime run.
Concrete code, counts, logs and environment details remain in `<PROJECT_ROOT>`.

Supports I12 title/context and target realization adapter seams only. Full
semantic review, complete TDR/finality, corpus content/timing/render acceptance
remain open. No new provider, runtime route, fixed per-title glossary or Phase 4
capability was introduced. Existing workflow and adapter contracts were checked;
no lifecycle or schema definition change is warranted by this scoped result.
