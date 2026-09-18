# Observation — mechanical visual-text probe

**Run ID**：2026-09-18-mechanical-visual-text-probe  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow）  
**原則**：LLM does not decide where text is. LLM classifies what role detected text plays, and only when mechanical features are ambiguous.

固定比例 band（例 bottom 72–98%、top 0–12%）是第一層搜尋，不是版式真相。不要為了「這部片字幕在 68%」讓 LLM 改 OCR crop（成本、延遲、新片不確定、crop 錯就看不到字）。

## Probe 與 coverage

```text
Probe Strategy: bottom band → top band → optional center → fallback full-frame
OCR detections → evidence coverage check
  sufficient → continue
  insufficient → expand probe
```

不是「bottom 有 OCR 就結束」。

## 機械 candidate（不必 LLM）

| 例 | 特徵 | `role.candidate` | `resolver` |
| --- | --- | --- | --- |
| 「某某短劇」normalized 右上小 box、persistence 0.97 | 清楚 | watermark | mechanical |
| 「你真的要去？」下方、persistence 0.02 | 清楚 | subtitle | mechanical |
| 中央 4 秒「林雪」 | 人名牌／字幕／招牌／文件／UI | （未定） | llm_vision |

`role.candidate` 仍是 candidate，不是 canonical 已判定。清楚時 **不必** 叫 LLM。

## LLM 不得回寫 probe

```yaml
observation:
  probe: bottom_band
  result: insufficient
  llm_observation:
    likely_subtitle_region: { y: 0.65 }
```

禁止立刻變成 `subtitle_y: 0.65`。多部片子累積後才：Observation → Accumulation → Policy candidate → Validation → Promotion。對齊本庫 observation → registry → validation。

| 問題 | 第一版 |
| --- | --- |
| OCR 掃哪裡 | Mechanical Probe（可 fallback） |
| bbox／normalized | 必存 |
| 浮水印／字幕候選 | 先機械特徵 |
| ASR↔OCR | Evidence arbitration |
| 不確定文字角色 | LLM vision escalation |
| LLM 改 crop／全局規則 | 不要 |
| 多片後改善 probe | Evidence → Observation → Validation → Policy |

## 真實片子要數

固定 band 漏字時是否 expand；LLM 是否被拿去改 crop；機械清楚案例是否仍打 LLM。沒被消費的探針參數不進 schema。
