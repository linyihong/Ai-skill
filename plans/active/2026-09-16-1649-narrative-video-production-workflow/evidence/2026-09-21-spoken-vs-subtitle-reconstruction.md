# Observation — spoken vs subtitle reconstruction

**Run ID**：2026-09-21-spoken-vs-subtitle-reconstruction  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**換 ASR 模型；**不**擴 workflow）  
**Extends**：[`2026-09-18-narrative-representation-gap.md`](2026-09-18-narrative-representation-gap.md)、[`2026-09-18-evidence-refinement.md`](2026-09-18-evidence-refinement.md)

## 兩種不同錯誤

| 通道 | 常對 | 常錯 |
| --- | --- | --- |
| OCR | 畫面字 | 可能是 sanitization／改寫，不是原音 |
| ASR | 有人在說、時間窗 | 字詞 |

禁止強迫「哪一個是真的」。兩個都是 observation，**object 不同**：`spoken_text` vs `subtitle_text`。

## 機械證據（constraint，不是答案）

保存 `duration_ms`、`char_count`、估計語速、若引擎有則 **word timestamps**。字數用來**排除**不合理候選（過長句對不上 speech window），不是鎖定答案。OCR 與 ASR 先做 temporal alignment（同一 speech window）再比字。

## `text_relation` 候選

- `semantic_mismatch`：窗對齊但用詞不同（例 ASR 不確定 vs 字幕成語）
- `subtitle_sanitization`：字幕可能在修飾髒話／敏感詞；兩邊都保留

## Spoken Text Reconstruction

任務名不是 ASR correction。輸入：對齊後的 ASR／OCR、timing、speaker、前後 1–3 句 local context、同一 narrative window、已解析 episode state／詞彙（角色名、成語、替代詞）。輸出多個 `reconstruction.candidates`，各附 reason。

禁止為劇情通順而發明後半句。reconstructed text 須 evidence-supported，否則 `unresolved`。Raw ASR／OCR 永不覆寫。

最終列同時存 `text.spoken.selected`＋candidates，以及 `text.subtitle.observed`。不確定就 unresolved。

## 本 round 順序

1. ASR word timestamp／speech duration  
2. OCR ↔ ASR temporal alignment  
3. spoken／subtitle 分家  
4. conflict candidate＋relation  
5. LLM reconstruction  
6. mechanical verification  

然後才進 narrative window／event assembly。詞彙證據接 [`18-episode-vs-knowledge-accumulation.md`](../18-episode-vs-knowledge-accumulation.md)（Learning Inbox，不直接當知識）。

## 候選紀錄形狀（去敏）

```yaml
speech_evidence:
  duration_ms: 1800
  asr: { text: "<asr_span>", char_count: 4 }
  ocr: { text: "<ocr_span>", char_count: 4 }
  prosody: { estimated_speech_rate: 3.2 }  # chars per sec
text_resolution:
  spoken:
    selected: "<idiom_or_unresolved>"
    candidates:
      - { text: "<ocr_span>", source: ocr, status: supported }
      - { text: "<asr_span>", source: asr, status: low_confidence }
  subtitle: { observed: "<ocr_span>" }
  relation: { type: semantic_mismatch }  # or subtitle_sanitization
  resolution: { status: resolved }       # or unresolved
```

observable：ASR（raw_text、timestamps、word timestamps、duration、speaker_id）、OCR（raw_text、bbox、timestamp、visual_style、persistence）、video signals。linking：alignment、candidates、spoken／subtitle relation、conflict、resolution。narrative 只消費 resolved spoken／subtitle 對，不消費 raw。
