> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-13 - Residue admit order and projection alignment

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自多語 locale projection：admit 前拒
絕 residual 字形會少寫 timed cue，下游按 index 對齊整段錯位。

#### One-line Summary

對非源語目標：先剝殘留源文字／再 admit；reject 不得默認縮短 cue 序列；
mono 外語 hardsub 不得蓋過 MT（除非真實 bilingual 證據）。

#### Human Explanation

三條容易混在一起：

1. **Hardsub prefer**：目標語畫面上的 OCR 看起來「已經是譯文」，但若 spoken
   是源語、group 是 mono，那常是外語硬字幕殘留，不是這句對白的譯文。
2. **Admit 順序**：泰文等譯文若夾一個漢字姓，先 strip／轉寫再跑 script gate；
   先拒後剝會變成 missing realization。
3. **Projection 長度**：uncertain／rejected 仍要保留時間軸對齊（placeholder 或
   disposition-aware join），不能讓 `len(target_cues) < len(source_cues)`  silently
   錯位。

#### Trigger

- `text_source=ocr_region` 遠多於 `mt` 且 spoken 為源語
- `realization_rejected` + `projection_loss` + target cue count 短一截
- residual_cjk／latin_leftover／en_honorific 後整句消失

#### Evidence

- Tool: multi-locale packs on a REVIEW caption matrix；EN hardsub override and
  TH residual_cjk reject-before-strip reproduced；other locales still drop on
  honorific／latin leftover
- Plan evidence: `…/evidence/2026-10-13-locale-projection-and-selection-leak.md`

#### Generalized Lesson

1. OCR override only with **bilingual／owned translation evidence**，不是「有目標語 OCR」。
2. **Normalize then admit** for peelable residue.
3. **Disposition-preserving projection**：source_index 對齊優於壓縮後的 dense list。

#### Agent Action

修 prefer-region 條件；strip／locale enforce 移到 reject 前；投影寫入與
disposition 等長或明確 target_index；加 cue-count／skew 機械檢查。

#### Goal / Action / Validation

| | |
| --- | --- |
| Goal | 語系包與源 cue 時間軸可對齊；MT 不被錯誤 hardsub 取代 |
| Action | gate 順序 + 投影對齊契約 |
| Validation | 合成 bilingual vs mono hardsub；reject 案例 cue 數不漂移 |

#### Applies / Does Not Apply

- Applies: locale Selection／realization／projection dogfood；caption REVIEW packs
- Does not apply: pure layout／timing without translation；identity-only name review with no dst text

#### Validation

- Synthetic crumb／residue cases assert BLOCK or strip-then-admit
- Cue-count skew and scaffolding dst strings are mechanically detectable
- Plan evidence linked from Translation Decision `evidence/README.md`

#### Promotion Target

- `workflow/translation/` adapters／failure-registry binding（when promoting）
- Translation Decision plan evidence（already linked）

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- `feedback/history/narrative-video-production/common/README.md` 索引
- `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/README.md`
- plan evidence `2026-10-13-locale-projection-and-selection-leak.md`
