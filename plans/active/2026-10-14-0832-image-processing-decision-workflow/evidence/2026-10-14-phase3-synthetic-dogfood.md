# Phase 3 — Synthetic dogfood + expanded negatives

## Validation

- [x] Five negative fixtures blocked (`finality` ≠ `accepted`)
- [x] rembg `seg.generic` dogfood produced IDR with independent verify
- [x] Metrics (elapsed／bytes) recorded
- [x] Guard promotion explicitly **not** performed（需人覆核）

## 量測欄

| 項 | 值 |
| --- | --- |
| Gate 指令 | `python3 workflow/image-processing/tools/run_phase3_gate.py --repo <AI_SKILL_REPO>` |
| Negatives | `neg.alpha-opaque`, `neg.residue-halo`, `neg.interior-damage`, `neg.over-erode`, `neg.consumer-layout` |
| Dogfood sources | 5 synthetic category proxies under `workflow/image-processing/fixtures/dogfood/sources/` |
| rembg subject | `dogfood.object-hard` |
| rembg elapsed | see `fixtures/dogfood/reports/dogfood.object-hard.rembg.idr.json` → `metrics.elapsed_sec` |
| rembg finality | `accepted_with_review`（無前景真值；強制人審） |
| Machine summary | `workflow/image-processing/fixtures/reports/phase3-gate-latest.json` |

## 候選比較（本輪）

| Candidate | 結果 | 註 |
| --- | --- | --- |
| `seg.generic`（rembg） | 產出 RGBA；機械 alpha 可 pass | 升為唯一真理 **禁止** |
| `gen.edit` | 未選 | fidelity strict 排除 |
| `post.decontaminate` | 未自動套用 | repair_candidate 僅建議 |

## Failure learning

| 觀察 | 處置 |
| --- | --- |
| 假透明／rim／內部咬穿／過度侵蝕／canvas 錯 | 已有負向 fixture + failure_class |
| 升 registry guard | **未升** — 需人覆核真實案例後才可 |

## 範圍邊界

- 合成圖 = 類別代理，**不是**業務 El Dorado 原圖。
- 人像髮絲等真實照片 dogfood 可另開 consumer evidence；本 run 用 `dogfood.soft-fringe` 代理。
