# Edit Decision Record（EDR）

**Invariant 1：** EDR 記錄決策與證據；成片（mp4／其它容器）是 output artifact，不是決策真相。
驗證方向：EDR ↔ rendered artifact。禁止從成片反推「當初為什麼這樣剪」。

結構化檔（YAML 或同等）是 canonical；Markdown 是人讀投影。沒有 EDR 不得宣稱 cut-ready 或 publish-ready。

欄位 SoT：[`records/edit-decision-record.yaml`](records/edit-decision-record.yaml)。
可填示範：[`records/examples/sanitized-matching-and-edr.yaml`](records/examples/sanitized-matching-and-edr.yaml)。

## 最小身份

`film_id`、`title`、`profile`（可空）、`bible_id`、`matching_script_id`、主 `narrative_template_id`、`shots[]`（含 `selected_clip_id`、入出點）、`mutations[]`、`qc`、`publish`、`locale_packs[]`、`outcome`（發布後）。

`qc.producer_self_check` 可推進到 cut-ready。`qc.fresh_reviewer` 才允許 publish-ready（Invariant 8）。

## Mutation

換 clip、改腳本、重做 bible 必須寫 `mutations[]`（原因 + 影響的 `shot_id`）。歷史 selected 不得默默覆蓋。

對白的 `dialogue.semantic_context` 若存在，抄自腳本／matching，作為決策痕跡；不得從成片口型反推 speaker。
