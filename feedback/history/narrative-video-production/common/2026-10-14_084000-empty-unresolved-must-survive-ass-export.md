> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-14 - Empty unresolved must survive ASS export

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自 REVIEW caption matrix：cache 已寫
等長 unresolved 空槽，ASS 略過空 text 後 index zip 整段錯位，被誤判成「亂譯」。

#### One-line Summary

`text_source=unresolved`／空字串仍是時間軸事件；ASS／burn 不得默默刪行，
REVIEW 也不得用 dense Dialogue index 對齊 pivot。

#### Human Explanation

Projection 契約若只保證 cache／JSON cue 數，交付物（ASS、burn 窗）仍可能壓縮
掉空槽。下游用「第 N 條字幕」對 ZH 時，從第一個被刪空槽起全部錯位——看起來像
語意崩潰，根因卻是 **export 長度 ≠ source_index 長度**。正確做法：空槽輸出
timed placeholder（或明確禁止 index zip，改 `source_index`／start join）；並把
「cache 等長但 ASS 短一截」做成機械檢查。

#### Trigger

- locale cache `len(cues)==len(pivot)` 且存在 `text_source=unresolved`／`text==""`
- ASS／burn Dialogue 行數 < cache cues
- ZH↔locale 大量 `start` 不一致，但各自內部單調

#### Evidence

- Tool: postfix multi-locale REVIEW matrix；one id pack showed 43 cache cues
  including empty unresolved vs 42 ASS dialogues and cascading time skew
- Plan evidence:
  `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/2026-10-14-ass-empty-unresolved-skew-and-semantic-review.md`

#### Generalized Lesson

1. **Disposition-preserving export**：unresolved empty is still a timed event.
2. **Join key**：`source_index` or media time — never dense ASS ordinal.
3. Cue-count gate must compare **export surface** used by REVIEW／burn, not only
   cache JSON.

#### Agent Action

修 `build_dub_ass`（或等價）保留空／unresolved 時間槽；加 cache↔ASS 長度
回歸；REVIEW 腳本改 time／source_index join；文件標明 index zip 為反模式。

#### Goal / Action / Validation

| | |
| --- | --- |
| Goal | 缺譯／拒譯不造成假「亂譯」錯位 |
| Action | ASS placeholder + join-key 契約 + 長度檢查 |
| Validation | 合成 unresolved 空槽 → ASS 等長；skew 檢測歸零 |

#### Applies / Does Not Apply

- Applies: locale projection → ASS／burn／matrix REVIEW
- Does not apply: pure semantic antonym when timelines already aligned

#### Validation

- Synthetic empty unresolved → ASS event count matches cache
- Matrix REVIEW join uses start／source_index
- Linked from Translation Decision evidence README

#### Promotion Target

- `workflow/translation/` projection／adapter notes（when promoting）
- Translation Decision plan evidence（linked）

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- `feedback/history/narrative-video-production/common/README.md` 索引
- `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/README.md`
