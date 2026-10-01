# Source vs publish timebase

分析軸與發布軸必須分離。這是 **contract refinement**（小型 invariant），不是新 Phase、不是新 runtime route。

## Invariant

**All source understanding and evidence extraction MUST use the canonical source timebase. Publish-time speed transforms MUST NOT be treated as new evidence.**

公式（constant speed）：

```text
publish_t = source_t / rate
```

禁止用「先把成片壓到 1.5x／1.25x／局部 ramp，再對該媒體重跑 OCR／ASR／scene／dialogue／story／clip matching」當成正式分析路徑。

## 兩條軸

| 軸 | timebase | 做什麼 | 不做什麼 |
| --- | --- | --- | --- |
| SOURCE／ANALYSIS | `canonical`（源片 1× 時鐘） | OCR、ASR、scene、dialogue、story event、clip catalog／matching、EDR decisions | 把 sped media 當 evidence source |
| PUBLISH／PLAYBACK | `publish`（經 transform） | video／audio／caption／TTS 投影、burn、最終 MP4 | 用 wall-clock `t` 與 1× 片直接互比畫面 |

Clip catalog 的 `source_in`／`source_out`／`duration_s` 天然屬 source timeline。Matching 要的是「需要 N 秒素材」；選完再做 publish transform。

## Timeline transform（EDR／Timeline IR）

優先抽象成 transform，不要把契約寫死成「只會有 1×／1.5×」：

```yaml
timeline_transform:
  type: constant_speed   # 未來可擴局部 ramp
  source_timebase: canonical
  target_timebase: publish
  rate: 1.5              # publish_t = source_t / rate
```

EDR／Timeline IR 應能同時保存：

- `source`／canonical 時戳（cue、clip、ASR／OCR observation）
- `publish` 投影時戳（若已 assemble）
- 上述 `timeline_transform`（或等價欄位）

## 不是同一套 timing policy

下列產物**不一定**共用同一個 transform 實作細節，但必須**明示各自 policy**，且不得 silently 混用 timebase：

| 產物 | 常見政策 |
| --- | --- |
| source video | constant_speed → publish video |
| 原聲／混音 | 同 rate，或明示保留 1× |
| TTS／narration | 同 rate，或重新 timing |
| caption | `source` 窗 → `publish` 窗 recalculated |

## QC 拆分

| QC | 問什麼 | 失敗例 |
| --- | --- | --- |
| `temporal_integrity` | video／audio／caption 是否遵循**宣告的**同一套（或明示分岔的）transform | video 1.5×、audio 1.0× → A/V sync failure |
| `presentation_comfort` | publish 後是否仍可讀／可聽／口型可接受 | 機械正確但口型過快、CPS 過高 |

`timing_gate`（locale）驗的是 **publish timebase** 的可讀性與 temporal integrity。
Evidence analysis（OCR／ASR validity、story）驗的是 **source timebase**。

### CPS 必須在正確軸上驗

Source 通過 CPS **不蘊含** publish 通過。例：source 10 字／2s = 5 CPS；`rate=1.5` 後窗變成 ~1.33s → ~7.5 CPS。Publish CPS／cue 窗必須在 publish 軸重算或等價縮放後再判。

## 比較成片的正確方式

比 1× 參考片與 sped 發布片時，必須對齊 **對應 source 時刻**：

```text
publish t=8.0s  ↔  source t=12.0s   (rate=1.5)
```

禁止：

```text
publish t=8.0s  ↔  source t=8.0s   ❌ 誤判「畫面不對」
```

口型快但 A/V／字幕互鎖 → 多半是 `presentation_comfort`，不是 alignment bug。

## 與既有契約的關係

- Timeline IR／coverage：[`captions-and-locales.md`](captions-and-locales.md)
- Assemble／QC：[`assemble-and-qc.md`](assemble-and-qc.md)
- EDR 欄位：[`records/edit-decision-record.yaml`](records/edit-decision-record.yaml)
- Locale `timing_gate`：仍三閘獨立；本檔只補 **timebase 標註**
