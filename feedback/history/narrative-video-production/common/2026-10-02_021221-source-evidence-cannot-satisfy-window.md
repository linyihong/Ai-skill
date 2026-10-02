> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-02 - Source evidence cannot satisfy window evidence

Status: promoted

#### Reuse Evidence

Promoted into acquisition-loop workflow in the same turn (scope invariant + window-local fallback contract). Independent product dogfood of the adapter remains open on the Phase 3 checklist.

#### One-line Summary

全片／source 探針證明「有字幕」不能代替「當前時間窗有可用 cues」；cache 命中後切窗為空必須依 window evidence 決定是否做 window-local OCR。

#### Human Explanation

短劇配音常先做全片 hardsub／speech profiling，再讀整集 `dialogue_cues` 緩存並切到 random-clip 時間窗。若把 source `hardsub=True` 或 cache hit 當成 window 已滿足，會在「畫面其實有字、但這 30 秒 cues 為空」時直接失敗，或反過來為了修一個空窗就無條件全畫面重 OCR。正確做法是保留兩層：source profile 回答「這部片有什麼」；window coverage 回答「這段夠不夠用」。

#### Trigger

- Source／全片 probe：`hardsub=True`（甚至 `speech=True`）
- `dialogue_cues` cache hit
- `slice` 到 clip window 後 0 句
- 或人工抽幀可見硬字幕，但 pipeline 報無可用對白

#### Evidence

- Tool: workflow contract review + mechanical path analysis（probe 全片抽樣 vs excerpt 時間切窗）
- Sanitized excerpt: source hardsub true + cues cache hit + window excerpt 0 → raise；slice 公式本身在有重疊時可保留句子
- Evidence path: `<AI_SKILL_REPO>/plans/active/.../evidence/2026-10-02-evidence-scope-source-vs-window.md`

#### Generalized Lesson

**Source-level evidence cannot satisfy window-level evidence requirements.**
Evidence 必須帶 `scope`（`source` vs `time_window`）。Cache hit ≠ 當前 scope 充分。Window 0 cues 時先產出 `cue_coverage`，僅在 suspicious（例如 source hardsub ∧ window ASR present，或空／壞 cache ∧ speech）才 escalation `window_ocr`；insufficient 不得只因 source hardsub 就重 OCR。不要把修復點設在時間切窗公式，也不要拿掉全片 probe。

#### Agent Action

- 更新／實作 acquisition loop 時：保留全片 probe + cache；補 window fallback，而不是改 slice 或縮成只探 window。
- 看到「hardsub=True 但節錄 0 句」：先查 cue_coverage／window signals，不要直接等同「探針壞了」或「全片無字幕」。
- 寫 reusable 規則時用 scope 用語，不寫單一劇名／路徑。

#### Goal / Action / Validation

- Goal: 防止 source／window 證據範圍塌縮；空窗時有條件 window-local recovery。
- Action: 寫入 workflow invariant + plan dogfood checklist；產品 adapter 實作 assess → window_ocr → 再 slice。
- Validation or reference source: workflow §Scope invariant；plan 46 scope dogfood 勾選；產品空窗不再僅因 source hardsub raise，且 insufficient 不觸發無條件 OCR。

#### Applies When

- 有 source profiling + 整集 cues cache + 時間窗節錄的配音／解說管線
- Monitor／escalation 需要區分「沒找到」的層級

#### Does Not Apply When

- 單次全片離線轉寫、不做 clip window
- 純手動時間軸、無 cache／probe 分層

#### Validation

- 契約文件出現 scope invariant 與 window_evidence 決策表
- 產品：window 0 + suspicious → 只掃該窗並可恢復；window 0 + insufficient → 不無條件 OCR
- slice 單元在重疊窗仍能保留句子（回歸）

#### Promotion Target

- `workflow/narrative-video-production/text-evidence-acquisition-loop.md`
- `plans/active/.../46-evidence-acquisition-escalation-loop.md`

#### Promotion Record

- 2026-10-02：已寫入 acquisition-loop §Scope invariant 與 plan 46 Scope dogfood；本 lesson status=`promoted`。

#### Required Linked Updates

- Updated: `text-evidence-acquisition-loop.md`、`46-evidence-acquisition-escalation-loop.md`、`evidence/2026-10-02-evidence-scope-source-vs-window.md`
- Domain index: `feedback/history/narrative-video-production/README.md`（新建）
- Project-specific run evidence stays in plan `evidence/`（sanitized）；no raw host／path／title in this lesson
