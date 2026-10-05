# Text evidence — OCR regions, subtitle groups, bilingual pairs

何時讀：硬字幕可能多語疊字、多 box、或 OCR 把多語黏成一字串時；在 Language Relation Gate／Text Resolution 之前。
欄位：[`records/text-evidence.yaml`](records/text-evidence.yaml)。
Plan：[`41-bilingual-ocr-regions-and-subtitle-groups.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/41-bilingual-ocr-regions-and-subtitle-groups.md)。
銜接：[`text-evidence-language-relation.md`](text-evidence-language-relation.md)。

> **執行契約，不重新解釋契約。** LLM 不猜語言行別；機械 region／group 先於 alignment。

## 在 lifecycle 的位置

```text
OCR Detection
  → Raw OCR Evidence
  → Text Segmentation / Boundary Recovery  (見 text-evidence-ocr-boundary.md)
  → text_region (text, box, language_candidate, raw/derived, …)
  → Region Classification (dialogue vs non-dialogue candidate)
  → Subtitle Grouping (incl. bilingual_pair)
  → Language Relation Gate (region ↔ ASR)
  → Text Resolution
```

對應 execution-flow **6b** 的前半；6b 後半仍是 Language Relation Gate。

## `text_region`（必填／可選）

| 必填 | 可選 |
| --- | --- |
| `id`、`text`、`language_candidate` | `box`／`normalized_box`、`box_size`、`timestamp`／frame、`persistence`、`visual_style`、`text_role` |

`language_candidate`：Latin-heavy→en；CJK-heavy→zh／ja／ko；Thai→th；Arabic→ar；未知→unknown。只作 candidate。

每個 region 的文字必須對得上同一 detection box 的 raw observation。
Lexicon／boundary segmentation 的 `parts` 數量可以不同於 boxes；不得靠陣列
等長假設配對，或失配後把整行 aggregate text 複製到每個 box。否則水印文字
會取得 dialogue region 的 geometry／role，破壞後續 projection 與 coverage。
若 box-local raw 不可得，需保留 grouping scope 並標記 attribution 未決。

回歸驗證必須同時涵蓋新格式與缺少 box-local raw 的舊格式；新格式 fixture
通過不能證明舊快取安全。分別驗證 non-dialogue 排除與真字幕保留，不能只
以污染字串消失作為成功條件。Source-local band 是輔助證據：主要帶的出現
頻率不能單獨否定另一個已有獨立字幕角色支持的帶；低頻、位置切換與 box
jitter 需有保留測試。未知角色仍保持未決，不因保留要求自動升格。

已投影的 box-local evidence 必須保留其 geometry、region identity、整個 source
的 persistence 與 role；下游不可重新展開 raw boxes 而撤銷 projection。定向
重探亦須傳遞這些欄位，不可用只含 text/time 的摘要取代。短窗本身不足以
推翻整個 source 的持續性證據；不明 scene text 應待確認而非直接 publish。

## Subtitle Grouping

同時間窗＋空間合理（常見：stacked、相似寬度、bottom band）→ `subtitle_group`。
`type`：`mono`｜`bilingual`｜`multi`｜`unresolved`。
bilingual 時記 `language_relation`（如 en+zh）與 `spatial_relation: stacked`。

## 與 ASR／spoken

| 情況 | spoken／dialogue primary | 異語 region |
| --- | --- | --- |
| bilingual ZH+EN，ASR=zh | zh dialogue region + ASR | EN = translation evidence（subtitle） |
| 僅 EN OCR，ASR=zh | ASR spoken；EN = cross_language subtitle | 見 language-relation |
| OCR 含 SALE／UI | 排除出 dialogue group | 保留 raw；role projection |

禁止：翻譯字幕直接當 `spoken_text`；未分 region 就把雙語黏字串當單一 OCR 證據。

## Projection（locale pack／burn）

`spoken_timed` 只投影 **spoken**。Dialogue cue／locale `cues[]` 必須仍帶：

- `subtitle_group`（含 per-language `regions`）
- `speech`（ASR timing＋text）
- `alignment`（speech ↔ subtitle_group）

不得為了 burn 方便把上述結構壓扁後丟棄。目標語字幕優先取同語 region evidence。詳見 plan [`41-…`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/41-bilingual-ocr-regions-and-subtitle-groups.md) §Projection。

## 推進條件

| 條件 | 失敗 |
| --- | --- |
| 多語疊字窗有 ≥2 `text_region` 或明示 split 失敗原因 | 不得用單一字串進 Language Relation 當唯一 OCR |
| bilingual group 有 per-region language_candidate | 不得 LLM 猜行別 |
| spoken 來自口播語 dialogue evidence | EN-only hardsub 當 zh spoken = 閘失敗 |
