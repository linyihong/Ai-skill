> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-13 - Selection scaffolding must not become destination text

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自 locale pack dogfood：非對白的
digit／punct pivot 進入 Selection 後，dst 出現 adapter 區塊標題字樣。

#### One-line Summary

OCR／排版屑當源句時，I21 必須 BLOCK；否則模型可能把 Selection prompt 的
「Guards / candidate space」類標題寫進譯文。

#### Human Explanation

翻譯鏈若把 `:0`、`(1`、純數字等屑送進 LLM，輸出不一定是空或亂碼，而可能是
prompt 裡的章節標題。這比 residual_cjk 更糟——機械 script gate 可能仍「通過」
（全是拉丁字母）。正確鏈：incomplete_source／非對白 crumb → BLOCK，不呼叫
Completion；並在 admit 前掃描已知 scaffolding 字樣作為安全網。

#### Trigger

- pivot／src 為短數字、裸標點、或無字母／CJK 的符號屑
- dst 含 `Guards`、`candidate space`、`CONSTRAINT[`、`CANDIDATE_SPACE` 等 adapter 標記
- repair loop 剛注入過 Failure Evidence selection_notes

#### Evidence

- Tool: postfix multi-episode locale matrix（REVIEW）；one pack showed EN dst
  equal to Selection section header while ZH pivot was punct／digit crumbs
- Sanitized pattern only；episode tables stay in `<PROJECT_ROOT>` analysis
- Plan evidence: `plans/active/2026-09-22-1000-translation-decision-workflow/evidence/2026-10-13-locale-projection-and-selection-leak.md`

#### Generalized Lesson

1. **I21 covers crumbs**，不只 ellipsis／頓號截斷。
2. **Scaffolding string deny-list** on dst before cache／ASS write.
3. Acquisition should not promote digit／punct crumbs to publishable pivot；
   translation BLOCK is the last line of defense, not the first.

#### Agent Action

擴充 `is_incomplete_source`（或等價）覆蓋 digit／punct crumbs；新增 dst 掃描
拒絕 scaffolding；回歸測試用合成 src=`:0` 斷言 BLOCK／空輸出。

#### Goal / Action / Validation

| | |
| --- | --- |
| Goal | 非對白屑不得變成「合法」譯文 |
| Action | I21 + scaffolding deny-list + 測試 |
| Validation | 合成 crumb → 不呼叫／不寫入；既有對白句不受影響 |

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
