> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-01 - Source vs publish timebase separation

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自叙事成片 dogfood：同內容 1× 與 constant-speed 發布片被拿同一 wall-clock `t` 互比，誤判「畫面／字幕錯位」。

#### One-line Summary

OCR／ASR／story／matching 必須留在 canonical source timebase；speed 只是 publish timeline transform（`publish_t = source_t / rate`），不得當新 evidence；QC 要分 temporal integrity 與 presentation comfort。

#### Human Explanation

把發布片加速後，同一口白會出現在較早的 wall-clock。若再對加速片重跑分析，或拿兩邊相同秒數對畫面，會把「時間軸變換」誤判成字幕／ASR／剪輯錯位。正確做法：在 1× 建 evidence 與 EDR；clip catalog 用 source_in／out；最後才投影到 publish；比較時用 `source_t = publish_t * rate`。口型看起來快但 A/V 互鎖，多半是 presentation comfort，不是 sync bug。Source CPS 通過也不等於 publish CPS 通過。

#### Trigger

- 對 sped／publish 媒體重跑 OCR／ASR／scene／matching
- 用同一 wall-clock `t` 比較不同 speed 的兩條成片
- 把「口型快」直接當 A/V alignment failure
- `timing_gate`／CPS 只在 source 軸驗、卻宣稱 publish 可讀

#### Evidence

- Tool: product dogfood；1× 參考成片 vs constant-speed 發布片對照
- Sanitized：`timeline_transform.type=constant_speed`；compare via mapped source time
- Paths／titles：`<PROJECT_ROOT>` only

#### Generalized Lesson

1. **Analysis 軸 ≠ Publish 軸**：理解劇情只在 canonical source。
2. **Speed = adapter／transform**，不是新證據源。
3. **Clip matching 用 source duration**；選完再壓 publish。
4. **EDR／Timeline IR 保存** source、publish、transform。
5. **QC 拆分**：temporal_integrity（同宣告 transform）vs presentation_comfort（可讀／口型）。
6. **`timing_gate.timebase = publish`**；evidence analysis `timebase = source`。
7. **禁止**先加速再分析；允許局部 rate／ramp 時仍保留同一 invariant。

#### Agent Action

寫／改產片契約時補 timebase；product adapter 明示 transform metadata；比較成片先換算時刻；locale CPS 在 publish 軸驗。

#### Validation

- Fixture：同一 cue `source 12–14` → `rate 1.5` → `publish 8–9.33`，無重 OCR。
- 比較協議：publish 8s 只對 source 12s，不對 source 8s。
- CPS：source pass + rate>1 時 publish 可能 fail（刻意）。

#### Goal / Action / Validation

- Goal: 加速只影響發布投影，不污染 evidence／matching／EDR 決策軸。
- Action: workflow invariant + EDR 欄位 + lesson／plan evidence；product 對齊 transform 元資料。
- Validation: mapped-time 對照 + 不對 sped media 重跑分析。

#### Applies When

- 成片有 playback speed／constant-speed／未來局部 ramp
- 雙成片（1× 參考 vs 發布）人工或機械對照
- locale timing／CPS 與 OCR／ASR 同管線

#### Does Not Apply When

- 無 speed transform（全程 1× publish）
- 純音訊／無時間軸投影的交付

#### Promotion Target

- `workflow/narrative-video-production/source-publish-timebase.md`
- `workflow/narrative-video-production/captions-and-locales.md`
- `workflow/narrative-video-production/assemble-and-qc.md`
- `workflow/narrative-video-production/records/edit-decision-record.yaml`
- Plan companion under NVP active plan

#### Required Linked Updates

- Evidence：`plans/active/2026-09-16-1649-narrative-video-production-workflow/evidence/2026-10-01-source-publish-timebase.md`
- `feedback/history/development-guidance/common/README.md` 索引
- NVP README／execution-flow 入口列
