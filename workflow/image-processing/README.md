# Image Processing Decision Workflow

`workflow/image-processing/` 是 **cross-cutting governed image processing decision** capability：
Input Contract → Analysis → Constraints → Feasible Candidates → Selection → Execute → Independent Verify → Repair → IDR Closeout。

> **狀態**：Phase 0–4 landed（plan [`2026-10-14-0832`](../../plans/active/2026-10-14-0832-image-processing-decision-workflow/_plan.md)）。
> **沒有** `route.workflow.image-processing`；YAML `runtime_projection.enabled: false`。
> 工具名（如 rembg）只出現在 candidate registry／adapter，**不**寫進 canonical lifecycle 步驟。
> B／C／D：[`contracts/task-class-*.yaml`](contracts/) lazy-load（step `0b`）。

## 一句話責任邊界

| 元件 | 問句 |
| --- | --- |
| Contract | 這張圖要做成什麼？什麼不可改？ |
| Analysis | 主體／邊緣／風險訊號是什麼？（可推翻） |
| Constraints | 哪些候選被契約／profile／授權排除？ |
| Selection | 在可行集內選哪個策略組合？ |
| Execute | Adapter 如何產出候選輸出？（非驗收） |
| Verify | 是否符合契約？（執行器不得自證 PASS） |
| Repair／Finality | 換策略、人審，還是 abort？ |
| IDR | 決策與證據能否追溯？ |

## Mechanical invariants（摘要）

| ID | 規則 |
| --- | --- |
| IP1 | 完成 = 符合 Image Task Contract，不是流程跑完 |
| IP2 | Constraint ≠ Selection；LLM 不得自定驗收標準 |
| IP3 | Execute ≠ Verify；禁止同一模型自證 PASS |
| IP4 | `inconclusive` 合法；無真值時統計不得單獨證明品質 |
| IP5 | `fidelity` 三維展開（structure／color／edge）；禁止 `immutable=strict` |
| IP6 | Failure → 有限 repair；同一 candidate 連敗必須換策略或人審 |
| IP7 | Slot silhouette／face-axis 屬 consumer profile，非通用 A 核心 |
| IP8 | 未註冊 runtime route；手動啟用本目錄 |

跨域證據形狀參考：[`traceable-evidence`](../cross-cutting/traceable-evidence/README.md)。

## 何時讀哪個檔

| 認知階段 | 檔案 |
| --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) |
| Task contract | [`contracts/image-task-contract.yaml`](contracts/image-task-contract.yaml) |
| Class B／C／D | [`task-class-b.yaml`](contracts/task-class-b.yaml) · [`c`](contracts/task-class-c.yaml) · [`d`](contracts/task-class-d.yaml) |
| Decision record | [`contracts/image-decision-record.yaml`](contracts/image-decision-record.yaml) |

| Verification | [`contracts/verification.yaml`](contracts/verification.yaml) |
| Finality | [`contracts/finality.yaml`](contracts/finality.yaml) |
| Task classes | [`registry/task-classes.yaml`](registry/task-classes.yaml) |
| Candidates | [`registry/candidates.yaml`](registry/candidates.yaml) |
| Failure classes | [`registry/failure-classes.yaml`](registry/failure-classes.yaml) |
| Fixtures | [`fixtures/README.md`](fixtures/README.md)（合成負向 inputs + reports） |
| Adapters | [`adapters/README.md`](adapters/README.md)（含 Q5 license checklist） |
| Verify tools | [`tools/README.md`](tools/README.md) |

## Methods（Go CLI；無預置 phase gate 腳本）

```text
<AI_SKILL_REPO>/scripts/ai-skill-cli/bin/ai-skill-<os>-<arch> \
  image-processing generate-fixtures --repo <AI_SKILL_REPO>

<AI_SKILL_REPO>/scripts/ai-skill-cli/bin/ai-skill-<os>-<arch> \
  image-processing verify <image> --profile <name> [--expect-block] [--json-out <path>]
```

組裝流程與注意事項見 [`tools/README.md`](tools/README.md)。Phase-N gate 只在後續 plan 需要時依該流程寫，不預先提交產品腳本。

## 明確不做

- 註冊 `route.workflow.*`／runtime 投影
- 把具體供應商寫進 execution-flow 步驟名
- 建 Auto Orchestrator／全能圖片 Agent
- 把業務 Sprite 對齊寫成通用去背必經步驟
- 以 rembg 安裝與否當作驗收條件
- 新增長期非 Go 腳本入口（`.py`／`.sh`／…；pre-commit 機械鎖）
