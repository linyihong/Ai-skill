# Phase 4 Readiness Gate

**Plan**: [`../2026-07-14-0856-ui-pattern-knowledge-workflow.md`](../2026-07-14-0856-ui-pattern-knowledge-workflow.md)  
**Status**: ▶ **Active**（設計凍結 — 不再往前設計 Phase 4）  
**Phase 4 / Research Cycle 2**: ⏸ **Not Started** · Interaction Knowledge = **Not yet justified**  
**Date opened**: 2026-07-14

---

## Research rhythm（凍結）

```text
Research Cycle 1
        │
        ▼
Knowledge Evolution Method (Emerging)
        │
        ▼
Phase 4 Readiness
        │
        ├── 找不到 R1
        │      └── Pattern + Composition 足夠（好結果）
        │
        └── 找到 R1
               │
               ▼
        驗證 R2
               │
               ▼
        驗證 R3
               │
               ▼
        才成立 Interaction Knowledge
               │
               ▼
        Open Research Cycle 2
```

Method：**不是下一層看起來合理就建下一層；是上一層第一次解釋不了真實現象，才建下一層。**  
不再設計 Phase 4 本體，直到 R1→R2→R3 有序通過。

---

## Research Cycle 2 開啟條件（鎖定一句）

> 發現一個**經過驗證（validated）**、且**無法由既有 Knowledge Layer 表達（cannot be represented）**的 Flow 問題。

不是「看起來麻煩／很複雜／很多 Rule」。

---

## Gates = 必要且有序（非平行）

```text
R1
│
├─ FAIL → 留在 Phase 3 / 繼續搜證（無新 Layer）
│
└─ PASS
      │
      ▼
R2
│
├─ FAIL → 新增 Composition Constraint（Constraint Accumulation；仍 Phase 3）
│
└─ PASS
      │
      ▼
R3
│
├─ FAIL → 繼續 Readiness（假說未成形；Not yet justified）
│
└─ PASS
      │
      ▼
Open Research Cycle 2 / Phase 4
```

每一個 FAIL 都有明確去向：升格 vs 補 Rule 不再模糊。

### R1 — Pattern ✅ + Composition ✅，Flow 仍 ❌

真實案例同時：選型正確、空間組合／Constraints 正確、**整體互動仍錯誤**。  
FAIL = 尚無此案例（預設）→ 不進入 R2。

### R2 — 不能靠 composition_rules +1 修復

+1 Constraint 能修 → R2 FAIL → 寫入 `composition_rules.yaml`，**留在 Phase 3**。  
需要狀態／事件／轉移且 Constraint 無法承載 → R2 PASS。

### R3 — 一句定義 Interaction Knowledge（先於 Plan）

草稿欄（空 = R3 FAIL）：

> Interaction Knowledge 描述 _______________________________ 。

示意（非正式）：時間軸合法狀態轉移，而非空間組合。  
講不清 → 繼續 Readiness；**Not yet justified**。

---

## Domain vs Method 成果

| 層 | 成果 |
| --- | --- |
| **Domain（UI）** | Pattern Knowledge · Composition Knowledge |
| **Method（Architecture）** | Layer Growth Rhythm · Constraint Accumulation · Governed Trace Termination · **Readiness-before-New-Layer** |

後四項構成 **Knowledge Evolution Method**（🟡 Emerging — first validation = Cycle 1；待第二研究線）。

---

## Final maturity（收尾判定）

| 對象 | 狀態 | 下一步 |
| --- | --- | --- |
| Pattern Knowledge | 🟢 Stable | 維護與擴充 Pattern |
| Composition Knowledge | 🟢 Stable | 新 Screen 持續 dogfood |
| Knowledge Evolution Method | 🟡 Emerging | 等待第二條獨立研究線驗證 |
| Phase 4 Readiness | ▶ Active | 持續收集 R1 反例 |
| Interaction Knowledge | ⚪ **Not yet justified** | 尚未證成新 Layer（≠「缺觀察」） |

「Not yet justified」≠ Observation：缺的是**開新 Layer 的理由**，不是缺被動觀察。

---

## Log（主動搜證）

| Date | Candidate | R1 | R2 | R3 | Disposition |
| --- | --- | --- | --- | --- | --- |
| — | （尚未登記） | — | — | — | OPEN |

---

## Explicit non-design

- [x] Readiness 設計完整 → **停止擴寫 Phase 4**  
- [ ] ~~Interaction Knowledge schema / plan draft~~  
- [ ] ~~Runtime Projection for interaction~~  

Until ordered R1→R2→R3 PASS.
