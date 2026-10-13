# 10 — Phase 2 Vertical Slice（執行＋獨立驗收）

← [`_plan.md`](_plan.md)

## 範圍

打通 A 類最小閉環的 **可執行驗收層**：合成負向 fixture + independent verify + Q5 授權清單。
完整 rembg 商用 dogfood 屬 Phase 3；本階段 **不**把 rembg 當 Gate 必要依賴。

## 產物

| 產物 | 路徑 |
| --- | --- |
| Adapter 說明 | [`workflow/image-processing/adapters/`](../../../workflow/image-processing/adapters/README.md) |
| License checklist（Q5） | [`license-checklist.md`](../../../workflow/image-processing/adapters/license-checklist.md) |
| Verify／fixture tools | [`workflow/image-processing/tools/`](../../../workflow/image-processing/tools/README.md) |
| 負向輸入 | `workflow/image-processing/fixtures/inputs/neg.*.png` |
| Gate 報告 | `workflow/image-processing/fixtures/reports/phase2-gate-latest.json` |

## Gate 指令

```text
python3 workflow/image-processing/tools/run_phase2_gate.py --repo <AI_SKILL_REPO>
```

期望：exit 0；`neg.alpha-opaque` → `alpha_invalid`；`neg.residue-halo` → `edge_halo`；`finality` ≠ `accepted`。

## 已跑通（本輪）

- [x] `neg.alpha-opaque` 擋住（alpha 存在但全不透明）
- [x] `neg.residue-halo` 擋住（navy rim）
- [x] Q5 清單文件存在
- [x] Verify 輸出含 `check_id`／`status`／`failure_class`／`repair_candidate`／`hardness`
- [x] rembg 端到端 dogfood → Phase 3（[`11`](11-phase3-dogfood.md)）
- [x] `neg.interior-damage`／`neg.over-erode`／`neg.consumer-layout` → Phase 3

## 與契約的對齊

Independent verify **不是**執行器；棋盤預覽不參與機械 PASS。
Repair 建議寫在 check 的 `repair_candidate`；是否執行仍由 Selection／人審決定。
