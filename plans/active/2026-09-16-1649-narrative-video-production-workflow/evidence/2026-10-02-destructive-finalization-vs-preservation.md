# Observation — Destructive finalization vs evidence preservation

**Run ID**：2026-10-02-destructive-finalization-vs-preservation
**Kind**：Phase 3 observation → **contract_gap closed on preserve path**（不開新 workflow／不重調 OCR）
**原則**：raw≫0 且 cues=0 = **finalization failure**，不是 acquisition failure。

## 觀察（sanitized）

單集 offline rebuild（cached OCR／ASR；無整包）：

```text
OCR boxes ─┐
           ├─→ dialogue candidates → alignment → dedup/merge
ASR lines ─┘
                ↓
         spoken/subtitle rebuild
           ┌────┴────┐
        old hard-delete    live preserve
              ↓                  ↓
           cues=0          accepted + uncertain
                           evidence_retained > 0
```

| 階段（歷史 notes） | kept（約） | 含義 |
| --- | --- | --- |
| OCR boxes / ASR | 兩位數以上 | acquisition 已有資料 |
| dialogue candidates → align → dedup/merge | 仍明顯 > 0 | fusion 正常縮量 |
| spoken/subtitle rebuild（舊） | **0** | destructive hard-delete |
| live preserve | accepted＋uncertain | Evidence non-destructive |

**定性**：歷史 cues=0 是 finalization 把「未解決」當成「不存在」；不是 OCR／ASR 找不到資料。

## 兩種否定（不得共用 discard）

| 類型 | 含義 | status | 例 |
| --- | --- | --- | --- |
| A. Evidence 不成立 | 觀測本身不是對白／字幕 | `rejected` | watermark、platform UI、timestamp、noise、duplicate |
| B. Evidence 成立，interpretation 不確定 | 有對白價值，但 OCR↔ASR（等）未收斂 | `uncertain` | `semantic_mismatch`（勿強行選邊、勿刪） |

```text
Candidate
 ├─ invalid evidence → rejected
 ├─ valid + clear interpretation → accepted
 └─ valid + unclear interpretation → uncertain  # 一級結果，非垃圾桶
```

`uncertain` 必須保留 `ocr`／`asr`／`alignment` 等 evidence，供後續 voice／face／context／再採樣／LLM semantic 再解。

## Evidence layer ≠ Production layer

| 層 | 計數 | 用途 |
| --- | --- | --- |
| Evidence | `evidence_retained` ≈ accepted＋uncertain＋rejected traces | corpus／再解析 |
| Production | `publishable` ≈ accepted | 成片／burn |

禁止把 `evidence_retained` 當成可成片字幕數。

## 計數口徑（必須具名）

歷史 notes 曾出現「merge 後一組數量」與「未決另一數量」並列且單位不明。產品 notes **不得**都叫 `kept`。至少分開：

| metric | 含義 |
| --- | --- |
| `merged_cue_groups` | dedup／短窗合併後的 group／cue 單位 |
| `resolution_candidates` | 進入 spoken／subtitle rebuild 的候選數 |
| `unresolved_count` | rebuild 標為未決／uncertain 的條數 |
| `evidence_retained` / `publishable` | 見上表 |

若兩數不一致，必須能解釋單位差異；不可 silently 混用。

## Phase 3 checklist（本輪）

| 項 | 狀態 |
| --- | --- |
| Evidence preservation（resolution 不得銷毀上游 evidence） | **PASS**（live preserve） |
| Uncertain preservation（semantic conflict → uncertain，不 coerce 成 reject／delete） | **PASS／待觀察** |
| Accepted quality（accepted 是否更接近真實對白） | **待驗證** |
| Uncertain recoverability（uncertain 能否被後續 evidence 解掉） | **待驗證** |

**暫不加功能**；下一觀察點是 accepted 品質與 uncertain 原因／可恢復性，不是再開 OCR 或新 Phase。

## 分類

| 標籤 | 判定 |
| --- | --- |
| `destructive_finalization` | contract_gap（舊 rebuild）→ preserve 修正 |
| `uncertain_first_class` | contract：semantic_mismatch 等 → uncertain |
| `evidence_vs_production_layers` | 計數／消費分層 |

## 連動（不改架構）

- Workflow：[`text-evidence-multimodal-resolution.md`](../../../workflow/narrative-video-production/text-evidence-multimodal-resolution.md)（Evidence non-destructive；兩種否定；metric 具名）
- Lessons：[`cue-finalization-must-not-hard-delete-candidates`](../../../feedback/history/narrative-video-production/common/2026-10-02_172329-cue-finalization-must-not-hard-delete-candidates.md)、[`evidence-non-destructive-resolution-invariant`](../../../feedback/history/narrative-video-production/common/2026-10-02_174015-evidence-non-destructive-resolution-invariant.md)
- 產品：`resolve_fused_to_cues` preserve＋`evidence_layers` notes；metric 名稱對齊本檔

## Validation

- [x] 單集 rejection table：歷史 cues=0；live retained＞0
- [x] 本 evidence + `evidence/README.md` 索引
- [ ] Accepted quality dogfood（待驗證）
- [ ] Uncertain recoverability dogfood（待驗證）
- [ ] 產品 notes 全面改用具名 metrics（adapter 進行中）
