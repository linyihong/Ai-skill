# Observation — visual text evidence

**Run ID**：2026-09-17-visual-text-evidence  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow）  
**current_support**：partial（plan／locale 有 `text_origin: ocr`；源片分析未要求 timestamped visual text）  
**candidate_concept**：`visual_text_evidence`

## 與現況

- 已有：ASR vs OCR vs 腳本，禁止混成一條無標記對白（locale／caption）。
- 沒有：Video → OCR → timestamped text → clip catalog／narrative／identity 的明確一級產物。
- 不把 OCR 改成「必須由 LLM 做」。它屬 Material Fact Extraction 的 evidence source。

## 兩個證據通道

| 通道 | 回答 | 例欄 |
| --- | --- | --- |
| `speech_text` | 有人說了什麼 | ASR + start/end + `speaker_id` |
| `visual_text` | 何時何地出現什麼字 | 見下方 canonical shape；OCR 只是一種取得方式 |

資料室 SoT 是 **visual text evidence**，不是 OCR dump。可與 ASR 互證，也可單獨提供 ASR 沒有的資訊（例如畫面「三年後」）。可成為 [`08-identity-precedes-naming.md`](../08-identity-precedes-naming.md) 的 `name_evidence`，但須先過 spatial／persistence 過濾。

## Canonical shape（機械層；role 只當 candidate）

```yaml
visual_text_evidence:
  id: vtx_00182
  text: "林雪"
  timestamp: 123.42
  frame_ref: frame_004821
  geometry:
    box: { x: 214, y: 812, width: 356, height: 74 }
    polygon: [[214, 812], [570, 812], [570, 886], [214, 886]]
    normalized_box: { x: 0.214, y: 0.812, width: 0.356, height: 0.074 }
  spatial_features:
    region: bottom
    relative_area: 0.026
    aspect_ratio: 4.81
    near_edge: true
  temporal_features:
    first_seen: 123.2
    last_seen: 126.0
    frame_coverage: 0.04
  acquisition: { method: ocr, engine: ocr_adapter }
  role: { candidate: subtitle }   # parser 不得寫成已判定
```

`normalized_box` 跨 1920×1080／1280×720／1080×1920 才可比較。`spatial_features` 由 box 算出。浮水印要靠 **同一 region 跨 shot 高 persistence**（例 `frame_coverage: 0.94`），不是只看面積。

`acquisition.method` 可為 `ocr`／`subtitle_file`／`scene_text`／`ui_text`／`human`。

禁止：缺 normalized 幾何；parser `role: watermark` 當事實；把右上角常駐 LOGO／台標拿去跟 ASR 做人名或對白仲裁。

## 真實片子要數

名稱、地點、時間、對白、手機訊息、劇情提示、場景文字、浮水印：哪些穩定出現、哪些被 bible／catalog／matching／EDR **消費**。穩定且被消費才考慮升 contract。
