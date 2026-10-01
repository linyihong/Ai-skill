# Candidate: OCR Discovery / Layout Probe（probe ≠ exclusion）

Companion to [`13-mechanical-visual-text-probe.md`](13-mechanical-visual-text-probe.md)、
[`09-visual-text-evidence.md`](09-visual-text-evidence.md)、
[`24-ocr-role-projection.md`](24-ocr-role-projection.md)。
**不是新 Phase、不是新 Agent、不推翻 Phase 1/2。**
Workflow：[`text-evidence-ocr-discovery.md`](../../../workflow/narrative-video-production/text-evidence-ocr-discovery.md)。
Evidence：[`evidence/2026-10-01-ocr-probe-discovery-not-exclusion.md`](evidence/2026-10-01-ocr-probe-discovery-not-exclusion.md)。

## 問題

產品 Mechanical Probe 用固定 bottom 假設（例 `cy >= 0.68` + `area`）計算 `subtitle_like`。
Miss 時 `reason=no_subtitle_like`／`sufficient=False` 被當成 **「這支影片沒有字幕」** 而停掃或放棄 OCR 路徑。
語意過強：實際只是「目前 region 假設未命中」。硬字幕在 upper-middle 等 layout variation 是跨作品常態。

Dogfood：同一 clip 上 probe 報 `subtitle_like=0`／`no_subtitle_like`，全量 targeted／expanded OCR 仍可採到大量 dialogue 行（含 CJK 對白）。證明缺口在 **discovery／recovery**，不是「無字」。

## 核心 refinement

```text
OCR Probe = discovery，不是 exclusion。
未命中預設字幕區 → inconclusive + recovery path，不是 subtitle.exists=false。
```

升格現有「Mechanical visual-text probe」為：

```text
Mechanical global discovery
  → Layout candidates
  → (optional) LLM／resolver layout classification
  → Mechanical targeted OCR
  → coverage validate → rediscovery if needed
```

## 與 13 的關係

| 13 已有 | 本檔補強 |
| --- | --- |
| LLM 不決定掃區／不寫死 `subtitle_y` | 明訂 miss ≠ 無字幕（`inconclusive`） |
| coverage 不足才 expand | **必有** global discovery recovery；expand 失敗仍不可 exclusion |
| 機械 role candidate | Layout candidates + per-source `source_ocr_profile` |
| Observation→Promotion 改 probe | 同上；另要求 coverage／layout metrics 分層記錄 |

## 禁止

- `no_subtitle_like` → STOP／`exists: false`
- 一開始每幀 full OCR 當唯一策略
- LLM 先猜 y → 機械只掃那裡（無 global discovery 兜底）
- 只暴露單一 `subtitle_like` 整數、無法區分 probe miss vs discovery hit

## Phase 3 分類

| 標籤 | 判定 |
| --- | --- |
| `ocr_probe_exclusion_false_negative` | **contract_gap candidate**；workflow 已吸收 discovery 契約；**產品 adapter 另驗**（本輪只 writeback Ai-skill） |

## 產品後續（本輪不做）

Adapter 應把 `no_subtitle_like` 改成 region-scoped inconclusive，並接 recovery loop／scan_profile／分層 metrics。
細節見 workflow 檔 Adapter 驗收。

## 產品 adapter 狀態（2026-10-01）

Windows product host 已落地第一刀：`ocr_probe`／`dialogue_source.probe_hardsub`／`ocr` frame-diff watch。
Dogfood：同一 clip 上 `subtitle_like=0` 但 `dialogue_candidates` → `has_hardsub=True`。
仍待：per-source `source_ocr_profile` 持久化、LLM layout classification escape（非 crop 寫死）。
