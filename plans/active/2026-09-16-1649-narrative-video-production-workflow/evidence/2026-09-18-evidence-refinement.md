# Observation — evidence refinement loop

**Run ID**：2026-09-18-evidence-refinement  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow；**不**建權重模型）  
**原則**：Parsers collect. Refinement links, arbitrates, verifies, and learns. Script consumes canonical evidence; it is not evidence authority.

## Canonical OCR 必須有畫面位置

```yaml
ocr:
  ocr_id: ocr_00182
  timestamp: 123.42
  text: "林雪"
  region: { x: 214, y: 812, width: 356, height: 74 }
  polygon: [[214, 812], [570, 812], [570, 886], [214, 886]]
  frame_ref: frame_004821
  origin: { method: ocr, engine: ocr_adapter }
```

幾何用來區分標題／招牌／字幕／名稱標籤／手機訊息。OCR 是取得 **visual text evidence** 的方法，不是「OCR 字幕」。

## 作品級 policy，不是全域分數

先抽 5–10 個 representative scenes（keyframe + OCR + ASR + timestamps），產出 `evidence_policy`（例：dialogue 以 OCR 為 primary、spoken_content 以 ASR 為 primary、ambiguous 升級 LLM）。這是**這部作品**的 modality reliability，不是永遠誰比較準。

第一版用明示 policy（對齊 matching 的 Selection）：

```yaml
selection_policy:
  dialogue: { primary: ocr, secondary: asr }
  spoken_content: { primary: asr, secondary: ocr }
  identity_name:
    primary: [ocr, narrative_reference]
    supporting: [asr, face]
  ambiguity: { escalate_to: llm }
```

禁止第一版 `ocr: 0.83`。等真實片子累積 OCR／ASR 對錯、衝突、LLM／人工修正、獨立驗證後，才有資格從 history 演化 `policy + learned reliability`。

## 仲裁與升級

一致 → `status: accepted` 進 canonical。打架 → Ambiguity Escalation（OCR 圖 + ASR + 前後對白 + 既有 identity + face_track）。LLM 回 `selected` + `reason_codes`；**Decision 仍由 deterministic arbitration contract 決定。**

獨立審查定位為 Independent Review／Outcome Verification，不是「驗收 LLM 重寫答案」。查 identity／dialogue／subtitle drift、缺句、錯人、錯場、錯時，分類 `accept`／`minor_fix`／`re-analysis`。修正本身寫回 evidence → Future Policy。同類衝突第三次應能不再問 LLM。

## 與既有候選

[`09-visual-text-evidence.md`](../09-visual-text-evidence.md) 幾何；[`11-face-as-candidate-evidence.md`](../11-face-as-candidate-evidence.md) Face 非判定器；[`07-material-fact-extraction.md`](../07-material-fact-extraction.md) 機械採集；[`08-identity-precedes-naming.md`](../08-identity-precedes-naming.md) 名稱延後。Invariant 5／8：policy 明示；publish-ready 獨立 verifier。

## 真實片子要數

缺 bbox 的 OCR 列；被當成全域的 OCR>ASR；script 回寫 evidence；LLM confidence 當 SoT；修正是否進入 policy。沒被消費的欄位不進 schema。
