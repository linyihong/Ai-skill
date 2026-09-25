# Frozen domain invariants（Phase 1）

Companion to [`_plan.md`](_plan.md)。**Architecture = ready；domain schema = experimental。**  
Phase 1 凍結這些 invariant，不凍結全部欄位細節。2026-09-16 review 採納：Loop-first／Governance-first、ERA（Evidence constrains Decision Space）、Constraint ≠ Selection。

## 定位

這是 **governed creative execution domain**，不是 ffmpeg pipeline。工具不得定義 workflow。Runtime 延後至至少一次真實 EDR dogfood 之後才有資格當 runtime candidate。

```text
Intent → Source/Evidence → Template → Constraints → Feasible set
  → Selection policy → Matching script → EDR → Assembly
  → Independent QC → Publish-ready → Outcome (evidence) → Learning
```

## 十條凍結 invariant

1. **EDR is canonical decision record.** EDR 記錄決策與證據；成片（mp4）是 output artifact，不是決策真相。驗證方向是 EDR ↔ rendered artifact，禁止從 mp4 反推「當初為什麼這樣剪」。
2. **Bible / Catalog are shared SoT.** Source bible 回答世界裡有誰／哪一集；clip catalog 回答實際有哪些可剪素材。兩者不可混成一篇散文。
3. **Catalog retrieval must return existing `clip_id`.** 統一文字庫是 canonical 查找面；向量／BM25／LLM retrieval 只是 adapter，不得發明庫外片段。
4. **Constraints define the feasible set.** `must_tags`、`entity_refs`、duration target／tolerance／band、continuity 等先過濾。未進可行集的 clip 不得被選。
5. **Selection is an explicit responsibility／policy.** 時長接近度可以是一條 selection criterion，**不得默認等於唯一的「最好」**。每個 shot 必須寫 `selection.policy` + `rationale`（例如 `duration_closest`、`preserve_character_continuity`、`human_review`）。模型可當更強 Selection Actor；runtime／workflow 仍管 Constraint 與 Verification。
6. **Assembly must be verifiable against EDR.** 時間線對不上 `shot_id`／`selected_clip_id` = 未通過。
7. **Locale correctness has independent gates.** Content（語意／專名／source residue）≠ Timing（讀得完）≠ Layout（放得下／安全區／不遮擋）。三閘不得合成一顆「字幕 PASS」。
8. **Publish-ready requires fresh verification.** Producer／agent 自驗只推進階段；宣稱 `publish-ready` 的 completion authority 必須獨立（對齊 3D：self-check ≠ completion review）。
9. **Outcome is evidence, not truth.** `supports`／`contradicts`／`insufficient_sample` 是 evidence status（目前證據是否支持該模板假設），不是「已證明模板有效」。單位仍是模板 × 窗口。
10. **Runtime projection remains deferred.** 不建 `runtime/*.yaml`、不註冊 route，直到真實 EDR dogfood 之後。

## Selection 契約（取代「時長最近 = 最好」）

```text
Need
  → Constraints  →  Feasible candidates[]
       (clip_id, matched_constraints, duration_delta, rejection_reasons)
  → Selection policy  →  selected_clip_id
  → EDR
```

空可行集 → 放寬約束、補 catalog、或標記 blocked。禁止用 heuristic 從不可行集裡硬挑一名。

## Template catalog slot（不拆 taxonomy）

v0 七個 id 可共存。每個 template 預留 `template_kind: structure | mechanism | format`。  
Phase 3 有真實 EDR 前**不**把 catalog 拆成三套。禁止執行時發明匿名模板。

## Question 分級

| 級 | 題 | Phase 1 處置 |
| --- | --- | --- |
| A 架構 | Q1 語意（模板是一級 artifact + kind slot）、Q2 成熟度邊界、Q3／Q8 SoT 形狀、Q7 所有權、Q9 受控 tag、Q11 Selection | 本 companion 凍結 |
| B 可改 v0 | Q5 outcome 欄位 | 先用建議五欄，dogfood 可修 |
| C dogfood | Q4 first profile、Q6 首輪語、Q10 band 秒數 | 不擋 Phase 1 完成 |

## Candidate invariant（Phase 3；未凍結、未進 workflow gate）

不改上方十條凍結文。升格前只當觀察契約，見 [`15-story-evidence-vs-dialogue.md`](15-story-evidence-vs-dialogue.md)、[`24-ocr-role-projection.md`](24-ocr-role-projection.md)、[`25-text-group-preserve-variants.md`](25-text-group-preserve-variants.md)、[`26-subtitle-layout-engine.md`](26-subtitle-layout-engine.md)、[`27-typography-layout-profile.md`](27-typography-layout-profile.md)、[`28-layout-review-loop.md`](28-layout-review-loop.md)、[`29-speech-timing-authority.md`](29-speech-timing-authority.md)、[`32-semantic-safe-wrap-uniform-typography.md`](32-semantic-safe-wrap-uniform-typography.md)。

11. **Story events require evidence traceability and resolved promotion.** 劇情不是對話摘要。進入 Script／EDR 的 story event／state change 必須回指 observable evidence（經 evidence_unit／event_candidate，見 [`16-evidence-unit.md`](16-evidence-unit.md)）。首份真實 run 證明 `traceable: true` 不足：blocking upstream、event semantics 必須 resolved；若宣稱 state change，subject／change／before-after 不得空；最後需 independent verifier。完整候選閘見 [`17-story-promotion-gate.md`](17-story-promotion-gate.md)。禁止「整段 ASR → LLM 摘要 → 這就是劇情」。

12. **OCR role classification MUST NOT destructively remove observable visual-text evidence; downstream consumers MUST use role-aware projections.** Watermark exclusion is a projection rule, not an evidence deletion rule。見 [`24-ocr-role-projection.md`](24-ocr-role-projection.md)。未凍結。

13. **Similarity-based merging MUST NOT discard a candidate when the differing span may carry semantic information. Only exact duplicates may be destructively merged at the evidence layer.** Timestamp 是 alignment signal，不是 identity key。見 [`25-text-group-preserve-variants.md`](25-text-group-preserve-variants.md)。未凍結。

14. **Subtitle content ≠ subtitle layout.** `layout_gate` 是 feasible-layout selection。`max_lines` 是上限不是 target；`font_size` 在 bounds。Wrap 必須 lossless 且不得切 protected span；segmentation 只能回上游。每 cue typography 統一，禁止逐行 auto-fit。見 [`26-subtitle-layout-engine.md`](26-subtitle-layout-engine.md)、[`30-max-lines-is-bound.md`](30-max-lines-is-bound.md)、[`31-font-size-hard-bounds.md`](31-font-size-hard-bounds.md)、[`32-semantic-safe-wrap-uniform-typography.md`](32-semantic-safe-wrap-uniform-typography.md)、workflow [`subtitle-layout.md`](../../../workflow/narrative-video-production/subtitle-layout.md)。未凍結。

15. **Caption cue timing for generated speech MUST be evidence-backed by a speech artifact (or ASR for source video).** TTS is an adapter. Speech timing QC ≠ layout QC. 見 [`29-speech-timing-authority.md`](29-speech-timing-authority.md)。未凍結十條；workflow 契約已落點。

