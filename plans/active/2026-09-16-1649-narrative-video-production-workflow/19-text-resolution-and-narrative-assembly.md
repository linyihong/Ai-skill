# Candidate: Text resolution and narrative assembly

Companion to [`_plan.md`](_plan.md)。**候選 representation，不是 Phase 2
workflow schema，也不是新增 Story Agent。** 去敏觀察：
[`evidence/2026-09-18-narrative-representation-gap.md`](evidence/2026-09-18-narrative-representation-gap.md)。

首份真實 run 顯示兩個不同缺口：

1. ASR／visual-text selection policy 已存在，但 resolved text 沒有成為 narrative
   consumer 的輸入；
2. event candidate 幾乎只是 `dialogue_cluster`，缺少跨 evidence 的情境組裝。

因此不是放寬 Story promotion threshold，也不是追求 event 數量。候選路徑：

```text
observable evidence
  → text resolution
  → evidence unit
  → narrative window
  → event assembly
  → narrative significance + candidate relations
  → story event
  → typed story-state claim（若有）
```

## Text resolution

Narrative 不得直接讀 raw ASR，也不得全域宣告「OCR 永遠優先」。先確認 visual
text 是 subtitle candidate，再按時間／片段對齊、完整性與語境選擇：

```yaml
text_resolution:
  selected: { text: resolved_span, source: visual_text }
  alternatives:
    - { text: raw_asr_span, source: asr }
  status: resolved
  evidence_refs: [...]
```

Raw evidence 永不覆寫。`status: unresolved` 時，下游不得把破碎文字當 final event
semantics。

`spoken_text` ≠ `subtitle_text`。OCR 準確不等于原音；ASR 有時間不等于字對。重建任務是
**Spoken Text Reconstruction**，不是 OCR 優先也不是改寫 raw ASR。見
[`21-spoken-vs-subtitle-reconstruction.md`](21-spoken-vs-subtitle-reconstruction.md)。

## Narrative window

`narrative_window` 是 context carrier，不是 event，也不是固定秒數。候選邊界可由
temporal proximity、相鄰 shots、speaker continuity、dialogue／subtitle
continuity、action continuity 與 scene continuity 組成。Local relevance 只描述
單一 evidence；Narrative relevance 必須在 window 內判斷。

## Event assembly

`dialogue_cluster` 降為 evidence grouping。第一輪只觀察少量語意 family：
`interaction_change`、`relationship_signal`、`action_candidate`、
`situation_change`、`revelation_candidate`、`transition_candidate`。這些不是
封閉 enum，也不能只憑文字標籤升格。

每個 event candidate 仍須有 `basis.*_refs`。多句 medium evidence 可在同一
window 組成 high-significance candidate；單句 high 不保證是 event。

## Narrative relation

Event 不是平行孤島。Relation 只可作 evidence-backed candidate：
`continuation`、`escalation`、`interruption`、`consequence`、`revelation`、
`setup`。每條需 `from`、`to`、evidence refs、resolution status；時間相鄰本身
不等於因果。Graph 是 carrier，不是 truth。

Beat 可由一組 events／relations 衍生，但本輪不把 Beat 凍成必經 schema。
Dialogue ≠ Event；Event ≠ Beat；Beat ≠ State Change。若宣稱 state change，
沿用 [`17-story-promotion-gate.md`](17-story-promotion-gate.md) 的具體、
可驗證 claim gate。

## 下一輪只驗四件事

1. role-qualified text resolution 是否真的被 narrative consumer 使用；
2. narrative window 是否能把分散的 medium evidence 組成 event candidate；
3. event relation 是否能表達 escalation → interruption → continuation，而不
   把 temporal proximity 當因果；
4. typed state claim 是否能描述 relationship／interaction／situation change。

Phase 3 仍不改 workflow、不加 detector、不加 Agent。第二個真實 episode 後再
決定哪些 representation 值得 promotion。
