# Fixtures staging（Phase 0 planning）

← [`../_plan.md`](../_plan.md) · cases [`cases/`](cases/)

## 目錄規劃

| 路徑 | 用途 | 何時 |
| --- | --- | --- |
| `plans/active/2026-10-14-0832-image-processing-decision-workflow/fixtures/` | **Staging**：案例規格、合成生成說明；二進位可晚放 | Phase 0–1 |
| `workflow/image-processing/fixtures/` | Canonical 可重用去敏／合成圖 | Phase 1 skeleton 起 |
| Consumer 專案 `evidence/`／私有資產 | 未去敏業務原圖、Sprite 對照 | 永遠不進 Ai-skill canonical |

## 命名

- 正例／壓力：`cases/<id>.md` + 未來 `inputs/<id>.png`
- 負向（verify 必擋）：`cases/neg.<id>.md`（對齊 [`05`](../05-independent-verification.md)）

## Reproducible case 最低欄位

每份 `cases/*.md` 必須有：`case_id`、`kind`（`reproducible_case`）、契約摘要、期望 check／`failure_class`、重跑步驟、`input_hash`（無二進位時用 `PENDING_BINARY`）。

## 索引

| case_id | kind | 壓力 | 二進位 |
| --- | --- | --- | --- |
| [`illust-glow-halo`](cases/illust-glow-halo.md) | reproducible_case | `edge_halo` | PENDING |
| [`illust-interior-holes`](cases/illust-interior-holes.md) | reproducible_case | `interior_damage` | PENDING |
| [`neg.alpha-opaque`](cases/neg.alpha-opaque.md) | reproducible_case（負向） | `alpha_invalid` | PENDING（可合成） |
| [`neg.residue-halo`](cases/neg.residue-halo.md) | reproducible_case（負向） | `edge_halo` | PENDING |

Observed-only 敘事仍見 [`07`](../07-dogfood-baseline-el-dorado-cutout.md)；不列於此索引當回歸基準。
