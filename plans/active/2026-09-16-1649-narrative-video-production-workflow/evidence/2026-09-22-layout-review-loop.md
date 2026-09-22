# Observation — layout review is a bounded loop

**Run ID**：2026-09-22-layout-review-loop  
**Kind**：Phase 3 observation（**不是** AI 自動調字幕；**不**改 workflow）  
**Extends**：[`2026-09-22-typography-layout-profile.md`](2026-09-22-typography-layout-profile.md)、[`2026-09-18-episode-vs-knowledge-accumulation.md`](2026-09-18-episode-vs-knowledge-accumulation.md)

## 兩層責任

Render 後 AI 看 visual evidence，標 fail／affected_constraints／adjustment_candidates。Engine 套 candidate 再 render。AI 不寫 profile。

## 階段（任一 FAIL 只重跑該層）

| 階段 | 問什麼 |
| --- | --- |
| mechanical_fit | 出框、超 max_lines、碰 forbidden、過小過大 |
| cue_layout | 單 cue 可行 layout |
| visual_review | 擁擠／空／擋臉／舒服 |
| temporal_stability | 字級／位置／斷行跳變（例 max_font_delta） |
| scene_consistency | 多 cue 代表幀 |
| episode_review | 整集視覺語言 |

## Budget 與 taxonomy

每層 max iterations、全 loop 總上限。低於 min 不是繼續試，而是 `contract_gap` 或 `data_insufficient`。

## Learning

單次「44 比 48 舒服」= observation。Accepted adjustment 進 Learning Inbox；跨集 evidence_count 後才 `promoted` preferred。對齊 Observation → Verification → Learning → Promotion。
