# Assemble and QC

對 EDR 的符合性，不是教渲染。Invariant 1、6、8。

## Assemble vs EDR

成片時間線必須對得上每個 `shot_id` 的入出點與 `selected_clip_id`。
對不上 = 未通過。不得用「成片看起來對」代替欄位對帳。

實際窗可微調，仍須同一 `clip_id`；換 clip → `mutations[]`。

## QC 角色

| 檢查 | 誰 | 最高成熟度 |
| --- | --- | --- |
| producer self-check | 作者／同一 agent | `cut-ready` |
| independent verification | 未參與該片 matching／assemble 的人 | `publish-ready` |

`qc.status: pass` 若只有 producer 簽名，不得當 publish-ready。
無觀察 → 不得填 pass。
