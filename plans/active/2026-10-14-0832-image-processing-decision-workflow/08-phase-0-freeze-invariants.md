# 08 — Phase 0 freeze — boundary invariants（2026-10-14）

← [`_plan.md`](_plan.md)

Stakeholder continuation（本輪「繼續」）：**接受 A1–A3 草案** → Phase 0 **freeze**，准許進入 Phase 1（doc-only）。
本檔只釘 **邊界與責任**；不重寫 plan、不提前實作 adapter。

## Verdict

| 項 | 裁決 |
| --- | --- |
| Phase 0 | **freeze** |
| Phase 1 | **開跑**（`workflow/image-processing/` 契約骨架；不加平台能力） |
| Plan 重做 | **不需要**（gap-fix B1–B3 已吸收） |

## Accepted assumptions

| # | 裁決 |
| --- | --- |
| A1 | **接受**：混合 — agent 決策 + 本機 Python 執行；Ai-skill 持契約／workflow；非預設 ComfyUI／雲端 |
| A2 | **接受**：P0 = A 類去背閉環；B 困難案例壓架構；C／D 契約預留 |
| A3 | **接受**：獨立 capability；NVP／slots／3D 為 future consumers |

## Q 矩陣（freeze 時點）

| # | 狀態 | 說明 |
| --- | --- | --- |
| Q1 | **resolved via A1** | 入口 = 混合（契約在 Ai-skill；執行 adapter／本機腳本）。正式 CLI／skill 名稱 Phase 1 再定，不阻塞 freeze |
| Q2 | pre-decided | Consumer constraint only |
| Q3 | pre-decided（P0） | 機械 + 人審 |
| Q4 | **deferred with owner** | Staging：`plans/active/2026-10-14-0832-image-processing-decision-workflow/fixtures/`；Phase 1 遷入 `workflow/image-processing/fixtures/`；專案私有證據留 consumer。Owner：Phase 1 skeleton |
| Q5 | **deferred with owner** | Phase 0 列出檢查項（本檔 §License checklist stub）；**完整清單必須在 Phase 2 Gate 前**落地。Owner：Phase 2 prep |

## Frozen invariants

1. 完成標準 = Image Task Contract，不是流程跑完。
2. Constraint ≠ Selection；Execute ≠ Verify；`inconclusive` 合法。
3. `fidelity` 三維展開（[`02`](02-image-task-contract.md)）；禁止 `immutable=strict`。
4. Observed ≠ reproducible（[`01`](01-problem-baseline-and-scope.md)、[`07`](07-dogfood-baseline-el-dorado-cutout.md)）。
5. Slot silhouette／face-axis **不**進通用 A 核心。
6. rembg 最多 `baseline_candidate`；禁止寫死為唯一路徑。
7. 不建 Auto Orchestrator／全能圖片 Agent。
8. 負向 fixture 必須能擋住假成功（[`05`](05-independent-verification.md)）。

## Decision loop（凍結形狀）

```text
Input Contract
    → Image Task Analysis
    → Constraints (+ consumer profile)
    → Feasible Candidates
    → Selection Policy
    → Execute (adapter)
    → Independent Verify
    → Repair | Human Review | Abort
    → IDR / Evidence Closeout
```

## License checklist（Q5）

Phase 2 完整清單：[`workflow/image-processing/adapters/license-checklist.md`](../../../workflow/image-processing/adapters/license-checklist.md)。
場次打勾仍是使用端責任；清單文件存在即滿足 Phase 2 Gate。
## Phase 1 准許／禁止

| 准許 | 禁止 |
| --- | --- |
| 建 `workflow/image-processing/` README、execution-flow、contracts、registry 草案 | 綁死供應商名到 canonical 步驟 |
| 靜態走讀契約完整性 | 實作完整執行 pipeline 當 Phase 1 完成條件 |
| 遷 fixtures 目錄規劃進 workflow | 把 El Dorado 未去敏原圖 commit 進 Ai-skill |

## Exit → Phase 1

當 `_plan.md` `status: in-progress` 且本檔 Verdict = freeze，Phase 1 checklist 可開始。
