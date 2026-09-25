# Observation — caption composition and scored semantic breaks

**Run ID**：2026-09-25-caption-composition-semantic-breaks  
**Kind**：contract（製片反例延伸；已寫進 workflow）  
**Extends**：[`2026-09-25-semantic-safe-wrap-uniform-typography.md`](2026-09-25-semantic-safe-wrap-uniform-typography.md)

## 反例

- 為字幕大小切句，把「談話」拆成 cue／行尾「談」與下一行「話」。機械上可能 fit，語意壞。
- 兩行以字數均分為目標，切開人名或動詞結構。
- 同一 cue 上行／下行各自縮字。

## Contract

Content integrity 先於 layout fit。Speech segmentation 改 cue；line wrap 只從
scored `semantic_break_candidates` 選 offset，且不得進 `protected_spans`。禁止
equal-length 目標。Typography 整 cue 統一；到 min 仍 overflow → 重切 Speech Unit。
