# Phase 4 Readiness Gate

**Plan**: [`../2026-07-14-0856-ui-pattern-knowledge-workflow.md`](../2026-07-14-0856-ui-pattern-knowledge-workflow.md)  
**Status**: ▶ **Active**  
**Phase 4 itself**: ⏸ **Not Started**  
**Date opened**: 2026-07-14

---

## Why Readiness ≠ Observation，且 ≠ Phase 4

| Mode | Nature | Goal |
| --- | --- | --- |
| Observation | 被動記下 | 「也許有 Interaction」 |
| **Readiness（本檔）** | **主動找反例** | 證偽「Pattern+Composition 已夠」 |
| Phase 4 | 建新 Layer | 僅當 Readiness **PASS** |

Research Cycle 1 驗證的是 **方法**（Representability→Inferability→Composability）。  
尚未證明 Interaction / Orchestrability 是 **新 Knowledge Layer**，而非現有層的延伸。

**唯一前置問題**：

> Interaction 的失敗，是否**無法**用 Pattern Knowledge + Composition Knowledge 解釋？

只有答案為 **是**，才開 Phase 4。否則不為虛構的新 Layer 建架構。

Method 對齊 Cycle 1：**不是先建 Layer 再找用途；是先遇到理論解釋不了的現象，再開新 Layer。**

---

## Gates（須全部通過才可開 Phase 4）

### R1 — 存在不可表達的 Flow 問題

至少一個**真實**案例同時滿足：

1. Pattern 選型正確  
2. Composition（空間組合 / Constraints）正確  
3. **整體互動仍錯誤**

候選壓力（示意，需真實證據）：

- Bottom Sheet → Payment → Toast 時序衝突  
- Dialog 關閉後立即開另一 Overlay → focus 失敗  
- 多步驟狀態轉移無法由 Composition Constraints 表達  

**PASS**：有案例卡片；**FAIL / OPEN**：尚無（預設）。

### R2 — 不能靠 +1 Composition Constraint 解決

若 `composition_rules.yaml` +1 rule 即可修好 → **仍屬 Phase 3**（Constraint Accumulation）。

僅當需要表述 **狀態 / 事件 / 轉移（時間軸）**，且 Constraint 無法承載 → R2 PASS。

### R3 — 最小假說可說清（非先寫 Plan）

先寫**一句**（定稿前可改；空句 = 未就緒）：

> Interaction Knowledge 描述 _____ ，而非 Composition 的空間組合關係。

示意（**非正式拍板**）：  
「Interaction Knowledge 描述 Pattern 在時間軸上的合法狀態轉移，而非空間上的組合關係。」

一句都講不清 → **不到**開 Phase 4。

---

## Expected trigger shape

```text
Screen / Flow A
  PASS (Pattern + Composition)
        │
        ▼
Composition Rule +1 ?
  · 能修 → 留在 Phase 3（R2 FAIL → 不開 P4）
  · 不能修
        │
        ▼
Flow B still FAIL
  · Composition 無法表示
        │
        ▼
Interaction Hypothesis 成立（R3）
        │
        ▼
Research Cycle 2 / Phase 4 可開
```

觸發理由必須是：**現有理論第一次解釋不了真實案例**——不是「想研究 Interaction」。

---

## Outcomes

| Outcome | Meaning |
| --- | --- |
| **找不到反例** | Phase 3 能力比預期更完整；Readiness 可 Continue／收斂；**不開** Phase 4 |
| **找到且 R1∧R2∧R3** | 有扎實理由開 Research Cycle 2 / Phase 4（Interaction / Orchestrability） |

---

## Log（主動搜證）

| Date | Candidate | Pattern OK? | Composition OK? | Flow FAIL? | +1 Constraint 夠嗎？ | Disposition |
| --- | --- | --- | --- | --- | --- | --- |
| — | （尚未登記） | — | — | — | — | OPEN |

---

## Explicit non-starts

- [ ] ~~Draft Interaction Knowledge plan~~  
- [ ] ~~New entries for “flow patterns”~~  
- [ ] ~~Runtime Projection for interaction~~  

Until R1∧R2∧R3 PASS.
