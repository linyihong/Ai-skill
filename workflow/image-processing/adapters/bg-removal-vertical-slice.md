# Adapter — Background Removal Vertical Slice（Phase 2）

實作計畫 companion [`04`](../../../plans/active/2026-10-14-0832-image-processing-decision-workflow/04-bg-removal-vertical-slice.md) 的 **本機最小閉環**。
步驟 id 對齊 [`execution-flow.md`](../execution-flow.md)；此處只寫 adapter 綁定。

## Loop（有限）

```text
0-1 Contract + Input Validation
2   Analysis (risk_flags; low confidence → ask)
3   Constraints → Feasible Candidates
4   Selection (IDR)
5   Execute adapter:
      - baseline: seg.generic → optional rembg bind
      - post.decontaminate / post.deterministic as needed
      - assemble RGBA; previews are NOT evidence
6   Independent Verify → ai-skill image-processing verify
7-8 Disposition / Repair (max 2; must change candidate)
9   Finality + IDR
```

## Candidate 綁定（可替換）

| candidate_id | 本機綁定（例） | 備註 |
| --- | --- | --- |
| `seg.generic` | rembg 通用模型（可選；consumer 本機） | **非**唯一路徑；未安裝時用合成／外部 matte |
| `post.decontaminate` | 本機後處理（rim 統計） | 需契約 `edge_treatment=rim_decontaminate` |
| `post.deterministic` | 尺寸／清零透明 RGB | |
| verify | `ai-skill image-processing verify` | 與執行器分離；見 [`../tools/README.md`](../tools/README.md) |

## 驗收組裝（方法，非預置 phase 腳本）

```text
ai-skill image-processing verify …/neg.alpha-opaque.png --profile neg.alpha-opaque --expect-block
ai-skill image-processing verify …/neg.residue-halo.png --profile neg.residue-halo --expect-block
```

必須：上述負向被 verify 擋下（finality ≠ accepted）。

## 禁止

- 執行腳本自報 PASS
- 把 rembg 寫進 `execution-flow.md` 步驟名
- 無授權打勾就用於商用 dogfood（見 [`license-checklist.md`](license-checklist.md)）
