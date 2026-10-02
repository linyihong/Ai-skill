> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-02 - Evidence non-destructive resolution (Phase 3 positive dogfood)

Status: candidate

#### Reuse Evidence

與 [`cue-finalization-must-not-hard-delete-candidates`](./2026-10-02_172329-cue-finalization-must-not-hard-delete-candidates.md) 同族；本條補上 **正向驗證**：單集 offline rebuild 證明 preserve 後 evidence 可保留，並把 invariant 命名為 Evidence non-destructive resolution。尚無跨專案獨立重用，維持 candidate。

#### One-line Summary

Resolution 只能把 candidate 標成 accepted／uncertain／rejected／merged（附 reason），不得無理由消失；`evidence_retained` ≠ publishable，Evidence layer 與 Production layer 必須分開。

#### Human Explanation

Dogfood 已定性：raw OCR／ASR 有資料時 cues=0，是舊 spoken／subtitle rebuild 在 finalization 前 destructive hard-delete，不是 acquisition miss。Live resolve 呈現健康鏈：raw→fused（大幅減少可接受，只要每步有 mechanical reason）→resolve_out→accepted＋uncertain，且 uncertain 是一級狀態。接下來觀察重點是 uncertain 的可解釋原因與 accepted 是否更接近真實對白，而不是再開 OCR 或新 workflow。

#### Trigger

- 單集 rejection table 顯示 raw≫0、歷史 cues=0，但 live preserve 後 evidence_retained＞0
- 有人把 `evidence_retained` 當成可成片字幕數
- 討論要為 cues=0 再開新 workflow／再加 OCR
- uncertain 被當垃圾桶或直接刪除

#### Evidence

- Tool: product offline single-episode rejection table（cached OCR／ASR rebuild）
- Sanitized：raw≈160 → fused≈40 → resolve_out≈38 → accepted≈30＋uncertain≈8；歷史 spoken rebuild 彙總機械／近音／LLM／未決→共0
- Evidence path: `<PROJECT_ROOT>` analysis only

#### Generalized Lesson

1. **Invariant — Evidence non-destructive resolution**：任一 candidate 在 resolution 只能變成 `accepted`｜`uncertain`｜`rejected`｜`merged`（`merged_into`），且必須有 `reason.code`；禁止無 trace 消失。
2. **兩層消費**：Evidence layer（accepted＋uncertain＋rejected／merged 可審計）≠ Production layer（publishable≈accepted）。Story／identity／translation 各自從 Evidence store 取用，不得在 dialogue finalization 依 story relevance 刪 evidence。
3. **raw→fused 縮量正常**：時間重疊、同句合併、duplicate、對齊、grouping 可解釋即可；要查的是每步 discard／merge reason，不是「為什麼只剩 N」。
4. **uncertain 是一級狀態**：OCR／ASR 語意衝突、跨語言、和諧詞、人名、黏字、UI 殘留等都應進 uncertain corpus，供後續 episode／voice／face／context 再解。
5. **不開新 workflow**：cues=0 已定性為 destructive finalization architectural bug；preserve＋三態 classify 是修正，不是新 Phase／Agent。
6. **Phase 3 觀察點**：accepted 是否更接近真實對白；uncertain 是否保留可再解資訊 — 兩者成立即證明 Evidence→Resolution→Finality 開始工作。

#### Agent Action

遇 raw≫0／cues=0：先單集 rejection table 與 live preserve 對照；寫入／強化 multimodal-resolution 的 non-destructive invariant＋Evidence／Production 分層；產品輸出 stage／status 計數時分開 `evidence_retained` 與 `publishable`；勿為 cues=0 開新 workflow，勿把 uncertain 當刪除。

#### Goal / Action / Validation

- Goal: resolution 可審計、非破壞；Evidence 與 Production 計數可分開讀。
- Action: lesson＋workflow strengthen；產品 notes／counts 對齊兩層；必要時單集 dogfood 對照。
- Validation: live resolve 在 raw＞0 時 evidence_retained＞0；每條非 accepted 有 status＋reason；文件禁止把 retained 等同 publishable。

#### Applies When

- spoken／subtitle／dialogue cue finalization
- OCR／ASR／fuse 已有 raw evidence
- Phase 3 dogfood 對照歷史 cues=0

#### Does Not Apply When

- 真的沒有任何 OCR／ASR／fuse candidate（仍屬 acquisition）
- 純排版／燒錄 typography（publish layout）

#### Validation

- Rejection table：歷史 destructive → 0；live preserve → retained＞0
- Schema／notes 同時出現 evidence 與 publishable 計數
- 無「為 cues=0 新開 workflow」的 plan／phase

#### Promotion Target

- `workflow/narrative-video-production/text-evidence-multimodal-resolution.md`
- 交叉：[`cue-finalization-must-not-hard-delete-candidates`](./2026-10-02_172329-cue-finalization-must-not-hard-delete-candidates.md)

#### Promotion Record

尚未 promotion；本輪同步強化 multimodal-resolution。

#### Required Linked Updates

- 更新 common README 索引
- 強化 multimodal-resolution：invariant 命名、merged、Evidence vs Production、Phase 3 正向證據指向
- 產品：resolve notes／counts 分開 evidence_retained 與 publishable（若尚未）
- 專案幀／計數證據留 `<PROJECT_ROOT>`；本 lesson 不含片名／本機路徑
