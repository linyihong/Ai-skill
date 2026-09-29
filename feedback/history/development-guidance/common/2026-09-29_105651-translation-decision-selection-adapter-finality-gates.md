> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-29 - Translation Decision: Selection adapter + mechanical Finality before publish

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自一次產品字幕／配音管線對齊 `workflow/translation`（ep8 Finality／親屬 residual）的整合。

#### One-line Summary

把 LLM 當 Selection Actor 時，只注入 constraints／Candidate Space；截斷源文與親屬殘留必須用機械 Finality／Validation 閘門，不可靠 prompt 案例表或 sole fixed gloss。

#### Human Explanation

產品若只改「翻譯 prompt」，仍會補全 truncated 句、或把源語親屬稱謂留在目標語。正確拆法：Analysis／Guards 產出 notes 與 seeds → Selection 只在 feasible 集合裡選 → Independent Validation／Finality 決定 accepted／needs_review／blocked；blocked 與 kinship residue 不得進 publish 文字。

#### Trigger

- Provider 路徑仍是 freeform「translate this」而非 Selection Actor
- Prompt 開始堆 `if locale: 例句` 或固定「姐夫＝某一譯」當唯一答案
- Truncated source（句末省略或明顯截斷）被 LLM 補出未出現的內容
- 目標語仍含源語親屬 token（residue），卻當 PASS 寫入字幕／配音

#### Evidence

- Tool: product unit goldens for truncated BLOCK、kinship residue review、locale honorific structured path、Selection notes 含 Candidate Space
- Sanitized excerpt: incomplete source → Finality `blocked` before LLM；kinship token in source+target → Validation hit + blank publish；seeds 進 Candidate Space，Selection 仍必選
- Evidence path: keep episode IDs、host paths、provider keys under `<PROJECT_ROOT>` analysis only

#### Generalized Lesson

1. **Selection adapter ≠ case map**：notes 只帶 constraints、failure pattern ids、Candidate Space seeds；禁止 per-locale 例句膨脹。
2. **I21 先於 Selection**：`incomplete_source` → Finality `blocked`；pipeline 應跳過 LLM，不得 invent missing content。
3. **I23 種子不是唯一譯**：親屬／稱謂進 Candidate Space；Validation 抓 source_language_residue／kinship residue；不得把單一 gloss 當 sole answer。
4. **Finality 三元要接到 publish**：`accepted` 才寫入 cue／dub text；`blocked` 與硬 residual 應 blank 或 handoff review，不要 silent PASS。
5. **Golden 對齊 dogfood**：用 unittest.TestCase（或同等 discovery）鎖 ep8／locale 機械子集，避免 bare `test_*` 函數在 runner 裡 count=0。

#### Agent Action

產品整合 Translation Decision 時：先接 Selection notes + Independent Validation + Finality gate，再談 provider；回寫 Ai-skill 時只沉澱通用閘門規則，不帶專案 path／episode 私證。

#### Goal / Action / Validation

- Goal: truncated／kinship residue 不得以 accepted 進 publish；honorific／locale 仍走 structured／Candidate Space。
- Action: adapter notes → Selection → validate → assess_finality → blank/block on hard fails.
- Validation or reference source: mechanical goldens（truncated blocked、residue rejects、structured honorific accepted）+ existing locale／failure suites.

#### Applies When

- 字幕／配音／caption pack 等把 LLM 當翻譯執行器的產品路徑
- 對齊 `workflow/translation` Finality／failure registry 的整合

#### Does Not Apply When

- 純人工譯稿、無 Selection／publish 閘門的離線編輯
- 非 translation decision（timing／layout／burn／TTS）

#### Validation

Goldens：truncated → blocked；kinship residue → not accepted；locale honorific structured → accepted；Selection prompt 含 Candidate Space／constraints 而非 case map。

#### Promotion Target

- `workflow/translation/adapters/README.md`（Selection Actor wiring note）
- `workflow/translation/execution-flow.md`（Dogfood／product gate 對照）

#### Promotion Record

本輪已輕量更新 adapters README + execution-flow Dogfood 指向；契約 YAML 未改 schema。

#### Required Linked Updates

- `workflow/translation/adapters/README.md`
- `workflow/translation/execution-flow.md`
- `feedback/history/development-guidance/common/README.md` index
