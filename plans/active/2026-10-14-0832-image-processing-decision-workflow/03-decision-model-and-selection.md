# 03 — Decision Model & Selection Policy

← [`_plan.md`](_plan.md)

## 決策迴路（Loop）

```text
Input Contract
    → Image Task Analysis          # 特徵／風險／信心（可推翻）
    → Constraints                  # 契約 + consumer profile
    → Feasible Candidates          # 被允許的策略／工具組合
    → Selection Policy             # 品質／保真／速度／成本／授權
    → Execute                      # adapter：分割／matting／像素運算
    → Independent Verify           # 與執行分離
    → Repair | Human Review | Abort
    → Evidence & Closeout          # IDR
```

分析結果是 **路由依據**，不是不可推翻真相。低信心 → ask／human_review。

## Image Task Analysis（最小輸出）

| 欄位 | 例 |
| --- | --- |
| `subject_count` | single／multi／unknown |
| `edge_type` | hard／soft／mixed |
| `content_type` | photo／illustration／ui_composite |
| `risk_flags` | hair, glass, shadow, thin_structure, text_overlay, glow_halo |
| `confidence` | 0–1 或 low／med／high |
| `needs_user_subject` | bool |

## Candidate Registry（示意，非綁定實作）

| candidate_id | 適用訊號 | 角色 |
| --- | --- | --- |
| `seg.generic` | 單一硬邊主體 | 通用分割 |
| `seg.portrait_matting` | 人像／髮絲 | matting |
| `seg.interactive` | 主體選錯／多主體 | 點選／框選 |
| `post.decontaminate` | 邊緣色污染 | 去污染 |
| `post.mask_fix` | 形狀錯（非顏色錯） | 修遮罩，不修邊緣色 |
| `post.deterministic` | 格式／尺寸／合成 | Pillow／OpenCV |
| `gen.edit` | 契約允許的局部重繪 | **非**預設去背 |

P0 以 `seg.*` + `post.*` 為主；`gen.edit` 預設不在可行集。

## Constraint → Feasible Candidates

約束示例（保真度三維見 [`02`](02-image-task-contract.md)）：

- `fidelity` preset=`strict` 或任一 `subject_*=preserve` → 排除 `gen.edit`
- `edge_treatment=none` → 排除 `post.decontaminate`／激進 matte refine
- `edge_treatment=rim_decontaminate` → 僅允許 rim 內 `post.decontaminate`；超出 `max_rim_px` 的候選不可行
- `background=transparent` → 必須產出 alpha 路徑
- `immutable` 含「牙齒／配件孔洞」→ 排除激進 hole-fill／interior opening（與 rim 授權無關）
- 授權未核可的權重／API → 排除

## Selection Policy（維度）

依序或加權（Phase 1 定公式；P0 可用簡單字典序）：

1. **契約符合**（硬性）
2. **保真度風險**（越低越好）
3. **預期邊緣品質**（對 risk_flags）
4. **成本／延遲**
5. **授權清晰度**

Selection Actor 可以是規則引擎或 LLM，但 **只能在 Feasible Candidates 內選**，並寫入 IDR。

## Image Decision Record（IDR）最小欄位

| 欄位 | 說明 |
| --- | --- |
| `contract_ref` | 契約摘要／hash |
| `analysis` | 分析輸出 |
| `feasible` | 候選 id 列表 |
| `selected` | 選中候選 + 理由 |
| `rejected` | 未選候選 + 原因 |
| `execution` | adapter、參數摘要、輸出路徑／hash |
| `verification` | check 列表（見 [`05`](05-independent-verification.md)） |
| `repairs` | 嘗試過的 repair + 結果 |
| `finality` | `accepted` \| `accepted_with_review` \| `rejected` \| `aborted` |

## Gate（Phase 1）

能指出：誰負責 Constraint、誰負責 Selection、LLM 不能做什麼。
