# Text evidence — Multimodal Text Evidence Resolution

何時讀：OCR＋ASR（或硬字幕＋口播）已過 Language Relation Gate，要做 spoken／subtitle
reconstruction，且可能出現 ASR 字面怪、跨語言字幕語義清晰、或近音錯字。

Plan companion：[`43-multimodal-text-evidence-resolution.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/43-multimodal-text-evidence-resolution.md)。
前置：[`text-evidence-language-relation.md`](text-evidence-language-relation.md)、
[`text-evidence-regions.md`](text-evidence-regions.md)。

## 在管線中的位置

插在 Language Relation Gate **之後**、narrative consumer **之前**（不是新 publish stage）：

```text
… → Language Relation Gate
  → Multimodal Text Evidence Resolution
       Evidence Normalization
       → Phonetic Evidence | Semantic Evidence
       → Candidate Generation
       → Context Resolution
       → Independent Verification
  → Resolved spoken_text / subtitle_text
```

## 核心契約

1. **ASR observed ≠ spoken meaning SoT**；保存 phonetic／timing／quality。
2. **OCR observed ≠ 自動 spoken**；跨語言時另產 `semantic_candidate`／`semantic_anchor`（provenance）。
3. **LLM 只選 candidate**，必須寫 `resolution.reason`＋`sources`。
4. **異常只觸發重建**，不機械替代表定案。
5. **權重看 evidence convergence**，禁止全域 OCR>ASR。
6. **三層命名**：`observed` → `candidate` → `resolved`；禁止 silently overwrite observed。

## observed / candidate / resolved

| 層 | 含義 | 例 |
| --- | --- | --- |
| observed | 設備實際看到／聽到 | ASR「…拍皮」；OCR `Last Lot` |
| candidate | evidence 推導的候選 | 「…拍卖品」（`semantic_reconstruction`） |
| resolved | resolution 接受、供敘事消費 | 同上（若 evidence 收斂） |

`拍皮→拍卖品`、`不偿→補償`、`秦舍→禽獸` 同一能力名：**semantic_reconstruction**（不是 typo correction）。

## semantic_anchor（OCR 領域術語）

英文硬字幕常是領域固定術語，不是逐字 gloss：

```yaml
semantic_anchor:
  source: "Last Lot"          # OCR observed（保留）
  domain: auction
  candidate_meaning:
    - "最後一件拍賣品"
```

路徑：OCR anchor ＋ ASR phonetic／lexical anomaly ＋（可選）scene domain
→ Candidate Generation → Resolution。
**OCR 非必要**：僅 ASR＋scene 也可產 candidate，但 confidence 應低於有 anchor 的收斂。

## 最小欄位

| 物件 | 必填／建議 |
| --- | --- |
| `asr` | `observed_text`, `language`; 建議 `phonetic`, `timing`, `quality.lexical_confidence` |
| `ocr` | `text`, `language`, `text_role`; 建議 `region_refs` |
| `semantic_anchor` | `source`, `domain`; 建議 `candidate_meaning[]` |
| `semantic_candidate` | `text`, `from: ocr_translation\|ocr_normalize\|semantic_reconstruction\|context_hint`, `source_ref` |
| `candidates[]` | `text`, `evidence[]`, `status` |
| `resolution` | `status`, `text`（若 resolved）, `reason[]`, `sources[]`, `confidence.type` |

`confidence.type`：`evidence_supported`｜`needs_review`｜`unresolved`（與 Finality 語意對齊即可，不強制同一 enum）。

## evidence_policy（摘要）

| 情境 | lexical／OCR semantic | ASR |
| --- | --- | --- |
| 有 dialogue subtitle | high | supporting |
| 無字幕 | — | primary |
| 跨語言／雙語字幕 | subtitle_semantics high | supporting |
| ASR lexical anomaly | ocr_semantics high；開 phonetic reconstruction | suspicious |

## 案例族（同一框架）

| 模式 | 典型訊號 | 走向 |
| --- | --- | --- |
| ASR 錯音 | phonetic 近、字面怪 | phonetic candidates＋context |
| ASR 怪／字幕對 | OCR semantic 清晰 | semantic_candidate＋phonetic support |
| 跨語言對齊 | EN OCR＋ZH ASR | `cross_language_translation`；非 conflict |
| 同語言 match | 字面一致 | `same_language_match` |
| 領域術語錨 | OCR `Last Lot`＋ASR「拍皮」 | `semantic_anchor` → `semantic_reconstruction` |

## 禁止

- 翻譯字幕直接寫入 `spoken_text` 且無 ASR／reason
- 未過 Language Relation Gate 就跑 sanitization／和諧覆蓋
- 用 OCR 顯示詞做同音展開（仍遵守 22）
- 無 evidence 時 LLM 自由改寫
- **Hard-delete candidates**：OCR／ASR／fuse 已產生的 candidate 不得因 script／lexicon／LLM
  未通過而從 evidence graph 消失（`continue`／drop 且不留 trace）
- **Cue finalization 混入 story relevance**：是否「劇情重要」不得決定是否保留 dialogue cue

## Evidence non-destructive resolution（invariant）

**命名**：Evidence non-destructive resolution（不是「再抓更多 OCR」）。

任一 candidate 在 resolution 只能落到：

| status | 含義 | 最低 trace |
| --- | --- | --- |
| `accepted` | 可當 dialogue／subtitle cue | `reason`／`sources` |
| `uncertain` | 未收斂，但**一級狀態**（非垃圾桶） | `reasons[]`（例 ocr_asr_conflict） |
| `rejected` | 明確非對白／字幕 | `reason.code`（例 watermark_region） |
| `merged` | 併入其他 cue | `merged_into` |

**禁止**：無理由從 evidence graph 消失（裸 `continue`／drop／hard-delete）。
`rejected`／`uncertain`／`merged` 仍是 evidence，不是「刪掉」。

Production eligibility 必須消費 resolver 的實際 decision。若 nested resolution
是 `unresolved`／`uncertain`／`rejected`，不得因外層 status 缺省或 stale accepted
而進 publishable；adapter 應在 finalization 同步狀態，derived cache 也驗證此一致性。

### Evidence layer ≠ Production layer

```text
Evidence layer
├── accepted
├── uncertain
└── rejected / merged（可審計）

Production layer
└── publishable ≈ accepted   # 成片／burn 只吃這一層
```

**禁止**把 `evidence_retained`（accepted＋uncertain＋…）解讀成「可成片字幕數」。
Story／identity／translation／matching 各自從 Evidence store 再判斷；不得在 dialogue finalization 依 story relevance 刪 evidence。

### 兩種否定（不得共用 discard）

| 類型 | 含義 | status | 例 |
| --- | --- | --- | --- |
| Evidence 不成立 | 觀測本身不是對白／字幕 | `rejected` | watermark、platform UI、timestamp、noise |
| Evidence 成立、interpretation 不確定 | 有對白價值但未收斂 | `uncertain` | `semantic_mismatch`；**禁止**強行選邊或刪除 |

```text
Candidate → invalid evidence → rejected
         → valid + clear → accepted
         → valid + unclear → uncertain   # 一級結果
```

### 具名 metrics（禁止混用 `kept`）

| metric | 含義 |
| --- | --- |
| `merged_cue_groups` | dedup／短窗合併後 group 數 |
| `resolution_candidates` | 進入 spoken／subtitle rebuild 的候選數 |
| `unresolved_count` | rebuild 標 uncertain／未決的條數 |
| `evidence_retained` / `publishable` | Evidence vs Production |

兩數不一致時必須能解釋單位；見 plan evidence [`2026-10-02-destructive-finalization-vs-preservation`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/evidence/2026-10-02-destructive-finalization-vs-preservation.md)。

### 健康漏斗（Phase 3 正向證據）

```text
raw → fused → resolve_out
                ├─ accepted
                └─ uncertain
→ evidence_retained > 0   （即使 publishable < retained）
```

- **raw→fused 大幅減少可以正常**：時間重疊、同句合併、duplicate、對齊、grouping — 每步需可解釋 mechanical reason。
- **accepted＋uncertain 並存是健康**：舊邏輯「不確定→刪掉」才是 bug。
- **cues=0 且 raw≫0**：先定性 destructive finalization；**不開新 workflow**。單集 rejection table 對照 live preserve。

### Resolution-loss monitor（定向重探，不凍結全片 coverage）

`raw → usable → fused → accepted` 是 evidence transitions，不是把 OCR／ASR
壓成單一分數。每一層必須記具名計數與非 accepted 的 disposition：

```yaml
resolution_funnel:
  raw_ocr: N
  usable_candidates: N
  fused_candidates: N
  accepted: N
  uncertain: N
  rejected: N       # each has reason.code
  merged: N         # each has merged_into
  anomaly_signals: [resolution_loss]
```

若 `candidate → accepted` 的縮減不被 `uncertain`／具 reason 的 `rejected`／
具 target 的 `merged` 解釋，系統必須標 `resolution_loss`。它不是直接 BLOCK：

```text
resolution_loss → targeted re-probe (OCR region | ASR window | frame sample)
                → re-resolution → recovered | still_suspicious | deferred
```

OCR、ASR、geometry 是互補 evidence：OCR 主要提供 visual text，ASR 提供
speech timing／phonetic corroboration，geometry 提供 region identity。禁止以
固定 OCR>ASR、ASR>OCR 或「附近無 ASR 即 reject」取代 resolution。

Coverage 也必須對 expected evidence 計算：speech windows 與
`subtitle_like` windows 分別比較 resolved coverage；全片 wall-clock 空白本身
不能直接判漏字幕。

### Derived artifact integrity 與 cache reuse

`dialogue_cues`／coverage artifact 寫入需遵守：write temp → schema validate →
atomic rename。JSON 或 schema 無效、source hash 不符、或 OCR／ASR／region profile／
fusion／resolution version 不相容時，cache 是 `invalid`，不得被當成正常 cache hit。
invalid cache 觸發對應 scope 的 rebuild／re-probe；不要求整片盲目重跑。

Lesson（問題）：[`cue-finalization-must-not-hard-delete-candidates`](../../feedback/history/narrative-video-production/common/2026-10-02_172329-cue-finalization-must-not-hard-delete-candidates.md)。
Lesson（正向 invariant）：[`evidence-non-destructive-resolution-invariant`](../../feedback/history/narrative-video-production/common/2026-10-02_174015-evidence-non-destructive-resolution-invariant.md)。
跨域 pattern：本檔實作 [`traceable-evidence`](../cross-cutting/traceable-evidence/README.md) TE1（non-destructive resolution）、TE2（`source_ref`／`sources[]`）、TE4（四態最低 trace）、TE5（兩種否定）。

## candidate → final（不是 PASS／DELETE）

Finalization 只回答「這條 evidence 能不能當 dialogue／subtitle cue」，輸出上表四態。

最低審計欄位（即使 `accepted=[]`／`cues=[]` 也要有）：

```yaml
resolution:
  candidates: { ocr: N, asr: M, fused: K }
  accepted: []
  uncertain: []
  rejected:
    - id: …
      reason: { code: watermark_region | … }
  merged:
    - id: …
      merged_into: cue_…
  stage_counts:
    mechanical: …
    cross_modal: …
    lexicon_or_phonetic: …
    llm: …
  layers:
    evidence_retained: …    # accepted + uncertain + rejected + merged traces
    publishable: …          # ≈ accepted only
```

跨語言（例：Latin OCR＋CJK ASR）字面不等 **不是** automatic reject；應走 cross-modal
resolution（見上表「跨語言對齊」），必要時標 `uncertain`，而不是刪除。

診斷優先：raw≫0 且 final=0 時先做**單集 rejection table**，不要用整包重跑／模型重載
把資源 OOM 與 resolver bug 混成同一個 failure。下一觀察點是 **uncertain 原因分佈** 與 **accepted 是否更接近真實對白**，不是再開 OCR。

**Accepted 品質（Phase 3）**：`status=accepted` 不得表示 OCR 整串可 publish。若 OCR 對白與 sticky watermark／logo／platform UI 黏在同一字串，必須先做 **role-aware projection**（見 watermark exclusion-is-projection evidence）；投影失敗則降為 `uncertain`，保留 observed，**禁止**把黏串當 clean `spoken_selected`。詳見 [`2026-10-02-ep16-accepted-uncertain-audit`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/evidence/2026-10-02-ep16-accepted-uncertain-audit.md) 與 [`2026-10-02-sticky-watermark-accepted-projection`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/evidence/2026-10-02-sticky-watermark-accepted-projection.md)。

Adapter：`sticky_ocr_projection` 在 `resolve_fused_to_cues` 出口前 post-pass — observed 全留；只投影 spoken／status。中文硬字幕片源 peel 後 **spoken 以 projected OCR 為權威**；ASR 不得靠 loose Jaccard 覆蓋語義衝突（見 sticky evidence 規則 3）。

## 產品落點

Adapter 應在 fuse／spoken resolution 產出上述 candidates 與 reason；locale content
消費 **resolved pair**，不消費 raw ASR 字面當唯一 SoT。Record：
[`records/text-evidence.yaml`](records/text-evidence.yaml)。
`resolve_fused_to_cues`（或同等 finalizer）不得對 wrong-script／watermark／empty-select
路徑裸 `continue`；應寫入 `uncertain`／`rejected` 並保全 `subtitle.observed`。
