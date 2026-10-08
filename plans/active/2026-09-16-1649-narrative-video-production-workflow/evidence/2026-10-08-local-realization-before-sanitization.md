# Realization evidence before sanitization

## Confirmed failure boundary

Wrong-language actor output can be erased by response cleaning before the
target validator sees it. An empty-result retry then repeats the same request
instead of repairing the actual failed candidate. Conversely, removing source
characters from a mixed-script full sentence can leave a short proper-name
fragment that appears script-valid but no longer realizes the dialogue.
Formatting cleanup is not target admission or semantic reconstruction.

## Scoped adapter repair

Preserve non-target candidate text through formatting and clause collection.
Run target validation on the complete candidate; pass the actual failure reason
and full candidate to a bounded repair actor. Script legality may reject a
realization, but must not license deleting the dialogue to obtain legal residue.
Revalidate repair output before content processing and cache admission. Keep
original candidate, repaired candidate, source identity and failure disposition
as evidence distinct from published phrase entries and timed cues.

Per-source bounded attempt history belongs in the entity/episode evidence
bundle, not a global continuously growing store. Write retained attempts after
phrase updates so an earlier cache snapshot cannot erase them. Subsequent phrase
and timed projection writes must preserve attempt evidence. Mechanical
`target_script_admitted` is deliberately not semantic Finality `accepted`.

## Verification and remaining gates

Frozen real actor traces distinguish wrong-script generation, response stripping,
name erasure and ineffective repair. Paired fixtures fail before repair and
verify retained candidate routing, invalid-residue non-promotion, failure-reason
binding, successful-repair retention, projection persistence and bounded history.
Existing orthography, modality, region and locale regressions remain separate.

Changed realization is tested in a fresh bounded local-only child with raw
response traces and isolated cache. One repaired candidate becomes a complete
target-script sentence; other candidates remain rejected with their full attempt
evidence preserved and without published phrase entries. Target-script admission
does not certify honorific identity, action preservation or semantic adequacy.
Native exit and semantic/content gates remain
separate. Successful routing and retention do not prove that a multilingual actor
can produce an adequate realization or preserve every entity/action.
Content, timing, layout, full corpus and governed TDR acceptance remain open.
Concrete source text, test names, paths, run IDs and counts stay in `<PROJECT_ROOT>`.
No canonical lifecycle stage, fixed locale glossary or corpus PASS is introduced.
