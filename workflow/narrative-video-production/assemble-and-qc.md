# Assemble and QC

對 EDR 的符合性，不是教渲染。Invariant 1、6、8。

## Assemble vs EDR

成片時間線必須對得上每個 `shot_id` 的入出點與 `selected_clip_id`。
對不上 = 未通過。不得用「成片看起來對」代替欄位對帳。

實際窗可微調，仍須同一 `clip_id`；換 clip → `mutations[]`。

## Timeline IR（編輯檔）先於 MP4

Assemble 產出／消費的 canonical 是 **Timeline IR**（見 [`captions-and-locales.md`](captions-and-locales.md) §Timeline projection），不是直接猜 MP4。
每條 selected caption／clip／voice 要有 trace；機械 QC（coverage／timing／bilingual group）PASS 後才 render。
發現「evidence 有、成片沒有」時，先查 Timeline／coverage report，不要只倒帶猜 fuse。

## Source vs publish timebase

Matching／EDR／clip catalog 用 **source** 時軸；playback speed 是 **publish** transform（`publish_t = source_t / rate`），不得把 sped 成片當 OCR／ASR／story 新證據。Timeline IR／EDR 應帶 `timeline_transform`（或等價）。契約：[`source-publish-timebase.md`](source-publish-timebase.md)。

| QC | 驗什麼 |
| --- | --- |
| `temporal_integrity` | 宣告的 transform 下 video／audio／caption 是否一致（含 A/V 同 rate） |
| `presentation_comfort` | publish 後可讀／可聽／口型體感（機械對 ≠ 體感可接受） |

比較 1× 參考片與 sped 發布片時，必須對齊 mapped source time，禁止同 wall-clock 秒互比。

## QC 角色

| 檢查 | 誰 | 最高成熟度 |
| --- | --- | --- |
| producer self-check | 作者／同一 agent | `cut-ready` |
| independent verification | 未參與該片 matching／assemble 的人 | `publish-ready` |

`qc.status: pass` 若只有 producer 簽名，不得當 publish-ready。fresh reviewer 必須留下角色、reviewed artifact ids、各 blocking gate verdict 與 review timestamp；例外時再記 exception 與 rollback target。這是可審查的 manual evidence，不能把欄位存在誤稱為 runtime 已機械驗證。
無觀察 → 不得填 pass。
