# Observation — OCR probe is discovery, not exclusion

**Run ID**：2026-10-01-ocr-probe-discovery-not-exclusion
**Kind**：Phase 3 observation → **contract_gap candidate**（workflow 已吸收；產品 adapter 待驗）
**原則**：Probe 只能說「目前假設區有沒有命中」，不能說「影片有沒有字幕」。

## 觀察（sanitized）

對一段含硬字幕的 random-clip 來源做 mechanical probe（預設 bottom／subtitle-like 幾何閘）：

| 訊號 | 值 |
| --- | --- |
| Probe `subtitle_like` | 0 |
| Probe `reason` | `no_subtitle_like` |
| Probe `sufficient` | false |
| 同 clip 後續／強制 OCR 採集 | 大量非空 CJK dialogue 行（完整 evidence 存在） |

結論：**default layout hypothesis 錯了**，不是無字幕。若產品把 `no_subtitle_like` 當終止條件，會在其他機器上表現成「無法產生字幕」。

## 錯誤塌縮

```text
bottom probe → no subtitle_like → STOP / 無字幕
```

## 目標 recovery

```text
bottom probe → no match → inconclusive
  → global layout discovery → upper-middle（或其他）candidate
  → targeted OCR → dialogue evidence
```

## 指標缺口

僅 `subtitle_like: 0` 無法區分 probe miss 與真無字。需要分層：`probe.*`、`discovery.*`、`coverage.*`、`layout.discovered_regions[]`。

## 分類

| 標籤 | 判定 |
| --- | --- |
| `ocr_probe_exclusion_false_negative` | contract_gap candidate |
| 與 13 mechanical probe | 補強：miss 必須有 recovery；LLM 仍不寫死 crop |

## Linked

- Workflow：`workflow/narrative-video-production/text-evidence-ocr-discovery.md`
- Plan：`45-ocr-discovery-layout-probe.md`
- Prior：`2026-09-18-mechanical-visual-text-probe.md`
