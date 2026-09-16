# Script and shot list

把主模板的 beat 落到 `shot_id`，再交給匹配腳本選 clip。
本檔不負責選材（那是 [`matching-script.md`](matching-script.md)）。

EDR 欄位：`template_beats[]`（`beat_id` ↔ `shot_id`）、`hook`、可選 `dialogue`。

欄位 SoT：[`records/dialogue-semantic-context.yaml`](records/dialogue-semantic-context.yaml)。

## Dialogue ≠ Retrieval Context

有對白時可寫：

```text
dialogue.text              → 給觀眾的短台詞（可省略主詞）
dialogue.role              → 可選：spoken / reveal / narration / other
dialogue.semantic_context  → 給機器的結構化語義補充
```

`semantic_context` **不是**字幕、不是把台詞寫長、不是旁白。它支援查找約束、角色對應、分鏡驗證、翻譯脈絡、EDR 決策痕跡。生命週期比「查找語」長，因此不叫 `search_text`／`lookup_text`。

有 `semantic_context` 時必填：`speaker`、`summary`。  
有意義才填：`addressee`、`subject_refs`、`object_refs`、`intent`。已知角色不可省；未知則整欄省略，禁止填 `unknown`。

這是 **Constraint / retrieval context**，不是 Selection。選 clip 仍走 Need → Constraints → 可行集 → selection.policy。

Phase 3：先 optional。真實片子裡「幾乎每段對白 shot 都要靠它才能對上素材」才考慮升成必填；少數特例維持 optional。觀察名：`dialogue_semantic_ambiguity`（不是立刻 `contract_gap`）。

推進：每個將上成片的 beat 有 `shot_id`；每個 `shot_id` 在 matching script 有一列。
失敗 rollback：`script_author`。
