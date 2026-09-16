# Script and shot list

把主模板的 beat 落到 `shot_id`，再交給匹配腳本選 clip。
本檔不負責選材（那是 [`matching-script.md`](matching-script.md)）。

EDR 欄位：`template_beats[]`（`beat_id` ↔ `shot_id`）、`hook`（開場承諾是否在前 N 秒兌現）。

推進：每個將上成片的 beat 有 `shot_id`；每個 `shot_id` 在 matching script 有一列。
失敗 rollback：`script_author`。
