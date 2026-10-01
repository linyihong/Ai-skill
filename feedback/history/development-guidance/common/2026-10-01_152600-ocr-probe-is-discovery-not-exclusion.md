> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-01 - OCR Probe is discovery, not exclusion

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自硬字幕 layout variation dogfood：預設 bottom／`subtitle_like` 幾何閘報 miss，同 clip 全量／targeted OCR 仍採到大量 dialogue。

#### One-line Summary

OCR Probe 只能標示「目前 region 假設未命中」（`inconclusive`），不得當成「影片沒有字幕」而 STOP；必須走 Mechanical global discovery → layout candidates →（必要時）LLM／resolver selection → mechanical targeted OCR；LLM 不寫死 crop。

#### Human Explanation

固定 `cy >= 0.68` 這類 bottom 假設在「字幕偏上」的片子會假陰性。把 `no_subtitle_like` 解讀成 exclusion，會直接丢掉可採集的硬字幕。正確做法是低成本全局 discovery 產 layout candidates，再定向掃；LLM 幫選哪組 box 像 dialogue／watermark，而不是猜一個 y 百分比。

#### Trigger

- Probe `reason=no_subtitle_like`／`sufficient=false` 被用來停止 OCR 或標記無字幕
- 只掃 bottom、無 global discovery recovery
- LLM 輸出單一 `subtitle_y` 並寫死 crop
- Metrics 只有一個 `subtitle_like`，無法看出 discovery 區有命中

#### Evidence

- Tool: product dogfood mechanical probe vs force OCR on same clip
- Sanitized：`ocr_probe_exclusion_false_negative`
- Paths／titles：`<PROJECT_ROOT>` only

#### Generalized Lesson

1. **Probe ≠ exclusion**：miss → `inconclusive` + `no_match_in_current_region`。
2. **Recovery 必備**：global discovery → layout candidates → targeted OCR。
3. **LLM = selection／hypothesis**，不是 OCR engine，不單獨決定掃區真理。
4. **採納** discovery→classify→targeted；**拒絕** LLM-first crop。
5. **Per-source profile** + layout drift 才 rediscovery。
6. **分層 metrics**：probe／discovery／coverage／layout regions。
7. **adapter_only** 產品落地；不開新 Phase／OCR Agent。

#### Agent Action

改 OCR probe 時：先對齊 `text-evidence-ocr-discovery.md`；禁止把 `no_subtitle_like` 映射成無字幕；補 recovery 與分層 coverage 欄位；LLM 只分類 detections／hypothesis。

#### Validation

- Bottom miss 案例：status 為 inconclusive，且存在 discovery／targeted 路徑
- 同素材 targeted OCR 能採到 dialogue，而 probe 仍可記錄 bottom miss
- 無「僅 LLM y=… → 唯一 crop」路徑

#### Goal / Action / Validation

- Goal: layout variation 不因 default band miss 而丢失字幕 evidence。
- Action: Ai-skill workflow／plan／evidence writeback；產品 adapter 後續對齊。
- Validation: dogfood probe miss + OCR hit；adapter 驗收清單。

#### Applies When

- 硬字幕 OCR mechanical probe／band expand／subtitle_like／scan profile

#### Does Not Apply When

- 純 ASR 路徑且明確不依賴畫面硬字幕
- 已確認片源無任何 on-screen dialogue text（需有 discovery 證據，不是單次 bottom miss）

#### Promotion Target

- `workflow/narrative-video-production/text-evidence-ocr-discovery.md`
- Plan companion `45-ocr-discovery-layout-probe.md`
- 補強 `13-mechanical-visual-text-probe.md`

#### Required Linked Updates

- Evidence：`…/evidence/2026-10-01-ocr-probe-discovery-not-exclusion.md`
- `feedback/history/development-guidance/common/README.md` 索引
- NVP `_plan.md` mechanical probe checkbox 註記
- `workflow/…/README.md`、`execution-flow.md` 入口

#### Closure

Workflow 契約已寫入；產品代碼本輪不改。Close-loop 以 Ai-skill commit／push 為準。

#### Notes / Links

- Prior observation：`evidence/2026-09-18-mechanical-visual-text-probe.md`
