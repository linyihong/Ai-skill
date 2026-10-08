# Translation Repair Loop（Loop Interfaces）

本文件把 product dogfood 已驗證的 `selection_notes`／`repair_of` 升成
**workflow 責任**：校正能力屬於 Loop／Governance，不屬於任何單一 Actor
（ChatGPT／Claude／Gemini／本地 Qwen／DeepSeek 都只是 Selection／Repair Actor）。

不新增 runtime route、不新增固定譯詞庫、不取代
[`execution-flow.md`](../execution-flow.md) 主鏈。與
[`dogfood-acceptance.md`](dogfood-acceptance.md) 互補：後者管驗收分欄；
本文件管 **failure → constraint／candidate → repair → re-validate**。

## 責任切分（硬邊界）

| 層 | 擁有 | 禁止 |
| --- | --- | --- |
| Governance | failure registry、constraints、Finality、escalation policy、telemetry | 讓 Actor 自改 registry／自關 PASS |
| Independent Validator | mechanical／lexical／declared-semantic／reference binding 分欄 | 用「我覺得對」當唯一證據 |
| Selection／Repair Actor | 在 Feasible Candidate Space 內選表面、整句重寫 | 發明身份、補 truncated source、把通順當語意 pass |
| Human／Strong Model Arbitration | 僅 **unknown failure** 或 registry 缺口 | 每條錯譯都呼叫強模型 |

Constraint Responsibility ≠ Selection Responsibility。
Actor 可仲裁提案，**不可**自行修改世界狀態（不得直接寫 accepted／改 SoT）。

## Loop 形狀

```text
Local/Cloud Actor → Candidate
        ↓
Independent Validation（mechanical ∪ reference ∪ lexical ∪ declared semantic）
        ├── PASS → Finality（仍受 timing／visual／artifact 分欄）
        └── FAIL → Failure Evidence
                        ↓
                 Failure Classification（core ∪ locale registry）
                        ↓
              known? → Constraint / Candidate refinement
                        ↓
                   repair_of + selection_notes
                        ↓
                   Actor Repair → Validation ↺
              unknown? → Escalate Arbitration → Governance Review → registry writeback
```

## Loop Interfaces（第一代產品契約）

### 1. `selection_notes` — Selection Space Guard

告訴 Actor：**本輪 decision space 已被治理層縮小**；負責 Selection，不是背答案。

最小欄位（可用純文字投影，但語義必須可還原）：

| 欄位 | 含義 |
| --- | --- |
| HARD CONSTRAINT | locale／script／established name／必保 action |
| FORBIDDEN | Latin residue、已 reject 表面、反義謂語詞根 |
| CANDIDATE_SPACE | 已獨立 accept 的譯名／允許的 honorific 層 |
| COVERAGE | 本輪只宣告哪些 semantic assertions |

禁止：用越來越長的「案例句子表」取代 registry／constraint。

### 2. `repair_of` — Repair Context

上一輪 **完整失敗候選**（不是 scrub 後殘片）+ Independent Validation 失敗宣告。
Actor 必須整句重寫；不得只改一個 token 卻宣稱已修。

### 3. Failure Evidence（應累積的不是單句對錯）

```yaml
failure:
  pattern_id: F* | locale-manifest:*   # registry reference
  manifestation: ...                   # locale-specific surface
  evidence_refs: [...]
  violated_constraint: [...]
  repair_direction: [...]              # feeds next selection_notes
  known: true|false
```

「`ทำแบบนี้` 錯了」只是 manifestation。
可重用資產是 pattern + violated_constraint + repair_direction。

## Governance Bootstrap → Local Autonomous Loop

| Phase | 行為 | 目標 |
| --- | --- | --- |
| A — Arbitration Bootstrap | Local Actor 失敗 → Human／Strong Model 分類 failure、寫 constraint／candidate／repair rule | 教 loop，不是教詞表 |
| B — Local Self-Repair | known pattern → 本地套用 guards → repair → Independent Validation | 降低 escalation rate |
| Escalate | unknown／regression／registry 缺口 → Arbitration → Governance Review | 不是每錯都 escalate |

## Loop Telemetry（不要把成功率誤當成「模型學會了」）

至少觀測：

- first-pass acceptance
- repair success rate
- repeated failure rate
- known／unknown failure ratio
- escalation rate
- independent-review rate
- constraint reuse rate
- regression recurrence

Autonomous capability 上升的證據是：**escalation↓、known repair↑、unknown↓**，
而不是 prompt 變長或單次 dogfood 全綠。

## Local Arbitration Actor（可選升格）

本地模型可輸出仲裁提案：

```yaml
arbitration:
  failure_pattern: ...
  affected_dimension: identity_realization|semantic|lexical|...
  proposed_action: repair|escalate|needs_review
  candidate_space: [...]
  rationale: ...
```

仍必須由 Mechanical Validator + Governance Rules 決定是否接受。
禁止 Actor 直接 `finality=accepted`。

## Fine-tuning 優先級

預設 **不**用 fine-tune 記「陈小姐→Nona Chen」。
要學的是決策程序：identity／reference → 分離 name vs honorific → candidate space → select。
Fine-tune 僅在 loop telemetry 顯示程序穩定、且仍有系統性 selection 偏差時再評估。

## Product wiring checklist

1. Selection Actor 只吃 `selection_notes`／Candidate Space（見 adapters README）。
2. Validation 失敗保留完整候選 → `repair_of`。
3. Failure Evidence 寫入／對齊 registry（known）或 escalation（unknown）。
4. 重跑必須改變條件或證據；相同 deterministic retry ≠ 修復。
5. content Finality ≠ artifact／timing／visual／publish PASS。

## Linked updates

- [`execution-flow.md`](../execution-flow.md) §Repair Loop
- [`adapters/README.md`](README.md)
- [`dogfood-acceptance.md`](dogfood-acceptance.md)（驗收分欄不變；重跑規則交叉引用）
- Translation Decision plan evidence（本輪 dogfood）
