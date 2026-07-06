---
id: 2026-07-06-review-architecture-adr-phase0b
parent: 2026-07-06-review-architecture-adr
phase: 0b
status: complete
created: 2026-07-06
purpose: >
  Phase 0b 泛化 / 反例測試證據。成功標準：找出 Role primitive 的邊界（含不支持 Role 的案例），
  而非只找支持 Role 的案例。Review 問題與 Role 問題分開判定。
---

# Phase 0b — Perspective / Role Generalization Evidence

## 兩個問題的分離（Phase 0b 前提）

| 問題 | Phase 0a 狀態 | Phase 0b 目標 |
|---|---|---|
| **Review 是不是 Workflow Phase？** | **已證明：否** | 不需再找 Review case |
| **Role 是不是 Runtime Primitive？** | **未證明** | 用非 Review 案例 + **反例** 判定 |

ADR-013 目前足以回答 Review；**不足以**回答 Role。Phase 0b 不追加 Review 論述，只測 **Perspective Switch 是否值得 primitive 化**。

---

## 測試方法

1. **支持案例**：需 perspective switch + 不同 capability 的自然組合（Review、Planning、Debugger…）
2. **反例（刻意）**：Refactoring、Documentation、Test Authoring — 若 `context.objective` 即可，則 **不** 應升級 Role
3. **Slice 衝突檢查**：若 activity 已被 ADR-009 workflow slice 擁有，新增同名 Role 是否 **重複 taxonomy**（Planning / Architecture / Validation）

**Primitive 晉升標準（不可替代性）**：

- 不是「很好用」，而是 **Capability + Context 無法穩定表達**
- 且 **≥3 個 activity 共享同一 primitive 語意**（非 Review-only）
- 且有 **bounded catalog** 可拒絕 Tester/Writer/Refactorer 膨脹

---

## 主矩陣：Activity × Placement

| Activity | 固定 Workflow Phase？ | Capability + Context 即可？ | Role / Perspective 更自然？ | 結論 |
|---|---|---|---|---|
| **Review** | ❌ 否（跨 slice） | △ 需 `perspective: reviewer` 才穩 | ✅ 強（停止實作、findings-only） | **Perspective 候選**；primitive 待定 |
| **Planning** | ❌ 否（但 `sd-intake` / plans 已擁有） | ✅ `plan-drafting` + `objective: plan` | △ 有 switch，但與 intake slice **重疊** | **Capability**；不升 Role |
| **Debugging** | ❌ 否 | △ `debug-trace` + `objective: root-cause`；`FORENSIC` mode 重疊 | △ 有 switch（停 forward fix） | **Perspective 候選 — Phase 0b.5 重開**（0b 過早 dismiss） |
| **Architecture fit** | ❌ 否（`architecture/` slice） | ✅ fit-analysis capability | ❌ slice 已是 Architect 階段 | **Slice**；Review 用 `architecture-review` + perspective |
| **Validation / evidence** | ❌ 否（`sd-validation` slice） | ✅ evidence-acquisition capabilities | ❌ slice 已是 Validator 階段 | **Slice**；Release review 用 perspective |
| **Refactoring** | ❌ 否 | ✅ `implementation` + `execution_mode: preparatory_refactoring` 或 `objective: refactor` | ❌ Refactorer role **重複** execution_mode | **Capability** — **反例** |
| **Documentation** | ❌ 否 | ✅ `doc-sync` + `objective: author` / linked-updates | ❌ Writer role 易膨脹 | **Capability** — **反例** |
| **Test authoring** | ❌ 否（`sd-test-strategy` slice） | ✅ `test-design` + `objective: test-first` | △ 弱 separation（常同一人 TDD） | **Capability** — 傾向反例 |
| **Implementation** | ❌ 否（`sd-implementation`） | ✅ default coding path | ❌ Implementer 為 default，非 switch | **Default path**；非 Role 候選 |

圖例：✅ 強 · △ 中等 · ❌ 弱 / 拒絕

---

## 三類候選深度分析（非 Review）

### 1. Planner

| 維度 | 分析 |
|---|---|
| Perspective switch? | 是 — 分解目標 vs 寫 code |
| 現有框架 | [`plans/README.md`](../../../plans/README.md)、[`intake.md`](../../../workflow/software-delivery/intake.md) §Plan-First、`sd-intake` slice |
| Role 問題 | **Planner role 與 sd-intake 雙重 taxonomy** — slice 已表達「規劃階段」 |
| Capability 表達 | `invoke: plan-drafting` + `context.objective: plan` + plan artifact 產出 |
| **結論** | **Capability + workflow slice**，非 Role primitive |

### 2. Debugger

| 維度 | 分析 |
|---|---|
| Perspective switch? | 是 — 停 forward feature work；假設驅動；證據優先 |
| 現有框架 | `cognitive_mode: FORENSIC / RECOVERY`（ADR-008）；[`validation-reasoning/`](../../../intelligence/engineering/execution/validation-reasoning/)；analysis workflow |
| Role 問題 | 與 **FORENSIC mode** 高度重疊 — mode 已編碼「如何查」 |
| 與 Reviewer 相似點 | 對抗式、停止 ship、產 findings |
| 與 Reviewer 相異點 | 常仍在改 code（fix loop）；artifact 是 trace 非 review report |
| **結論** | **Perspective 候選 — Phase 0b.5 重開**；與 Review 可能共享 `fault_finding`；**不足以** 單獨建立 Role primitive |

### 3. Refactoring（刻意反例）

| 維度 | 分析 |
|---|---|
| 需要 Implementer → Refactorer？ | **否** |
| 現有框架 | [`execution-modes.md`](../../../workflow/software-delivery/implementation/execution-modes.md) — `preparatory_refactoring` **已是 implementation execution_mode** |
| Capability 表達 | `context.objective: refactor` + `behavior_change: false` |
| **結論** | **強反例** — Role 會 duplicate execution_mode axis |

---

## Perspective Switch 模式檢驗（D1 必要條件）

若 D1 成立，應觀察到 **跨 activity 的統一模式**：

```text
Workflow (caller slice)
  → Perspective Switch
  → Different Capability
  → Different Artifact discipline
```

| Activity | 符合統一模式？ | 備註 |
|---|---|---|
| Review | ✅ | 最清晰 |
| Debugger | △ | 模式類似但 artifact/mode 不同 |
| Planner | ❌ | slice 已覆蓋；switch 邊界模糊 |
| Architect | ❌ | slice 覆蓋 |
| Validator | ❌ | slice 覆蓋 |
| Refactorer | ❌ | execution_mode 覆蓋 |
| Writer | ❌ | objective 覆蓋 |

**僅 1 個強符合（Review）+ 1 個弱符合（Debugger）→ 不滿足「≥3 共享 primitive 語意」promotion rule。**

---

## 反例彙總（Role 邊界）

下列 activity **不應** 使用 `cognitive_role` primitive（若建立也應拒絕 catalog entry）：

| Activity | 應使用 | 理由 |
|---|---|---|
| Refactoring | `execution_mode` / `context.objective` | 已有 preparatory_refactoring |
| Documentation | `doc-sync` capability + linked-updates | Writer role 無不可替代性 |
| Test authoring | `sd-test-strategy` + test capabilities | TDD 與 implement 同人 |
| Planning | `sd-intake` + plans artifact | slice 已擁有 |
| Architecture analysis | `architecture/` slice | slice 已擁有 |
| Evidence validation | `sd-validation` slice | slice 已擁有 |

---

## Phase 0b 對五題必答的更新

| # | 問題 | Phase 0b 結論 |
|---|---|---|
| 1 | Review 跨 Workflow？ | **Yes**（0a 已證明） |
| 2 | Review 需 Persona 切換？ | **Yes** |
| 3 | Persona 需 Runtime Primitive？ | **No（傾向）** — `context.perspective` 足夠表達 Review / Debug 視角 |
| 4 | Capability 足以表達 Reviewer Context？ | **Yes** |
| 5 | Primitive 泛化到其他 Domain？ | **No** — 多數 activity 由 **slice 或 execution_mode 或 context.objective** 覆蓋；僅 Review（+ 部分 Debugger）呈現 perspective switch，不足以支撐 runtime primitive |

---

## Phase 0b 對 D1 / D2 的推導

| 結果 | 推導 |
|---|---|
| 若僅 Review（+ 弱 Debugger）適合 perspective | → **D2** |
| 若 Planner + Architect + Reviewer + Debugger + Validator 皆呈現同一 primitive 語意 | → **D1** |
| **Phase 0b 實際結果** | 多數 activity 已被 **slice / mode / objective** 軸表達；**2 個 perspective 候選不足以 primitive 化** → **推薦 D2** |

### D2 建議 envelope（Phase 0b draft）

```yaml
invoke:
  capability: code-review          # 或 architecture-review, debug-trace, …
  context:
    perspective: reviewer          # bounded enum: reviewer | debugger | author | default
    caller_slice: sd-implementation
    objective: optional            # refactor | plan | test-first — 非 role
```

- **`perspective`**：對抗式 / 審查式 **視角**（少數值，防膨脹）
- **`objective`**：任務目標變體（refactor、plan…）— **不是 role**
- **`execution_mode`**（ADR-008）：FORENSIC 等 — **不是 role**
- **`workflow slice`**：Planning / Architecture / Validation 階段 — **不是 role**

---

## Phase 0c 輸入（供 stakeholder review）

1. **Accept ADR-013 with D2** — Review gap 用 cross-cutting review capabilities + `context.perspective`
2. **Reject D1** — 泛化不足；taxonomy 與 slice/mode 衝突；primitive 成本 > 收益
3. **保留 open**：`perspective` enum 是否僅 `reviewer | debugger` 兩值起點
4. **仍 reject A, C** — 不變

---

## 與既有框架的對照證據

| 既有機制 | 已覆蓋的「像 Role」語意 | 對 D1 的影響 |
|---|---|---|
| ADR-009 cognitive slice | Planning、Architecture、Validation 階段 | 削弱 Planner/Architect/Validator role |
| ADR-008 cognitive_mode | FORENSIC、RECOVERY（debug 類） | 削弱 Debugger role |
| implementation execution_mode | preparatory_refactoring | 削弱 Refactorer role |
| plans/ + intake Plan-First | 規劃 artifact 循環 | 削弱 Planner role |
| validation_capability (glossary) | 證據取得 executable layer | 與 Role 不同軸；D2 可並列 |

---

## Phase 0b 成功標準 — 自檢

- [x] 刻意尋找 **不支持 Role** 的案例（Refactoring、Documentation、Test authoring）
- [x] 說明 Role 的 **適用邊界**（非十頁空論）
- [x] 分離 Review 已證明 vs Role 待證明
- [x] 產出可驅動 Phase 0c 的 **自然推導**（傾向 D2，非偏好投票）
