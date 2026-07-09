# Plan Evidence Accumulation（計畫證據累積）

> **Status**：validated（2026-07-09，dogfood 2d′ 落地 + `validatePlanEvidenceConvention` commit-msg validator）
> **Canonical source**：本檔為 plan 證據目錄與引用規則的人類可讀 rule；機械檢查見 `scripts/ai-skill-cli/internal/app/plan_evidence.go`，obligation `obligation.commit.plan_evidence_convention`。

## 目的

多 session / 多 slice 的 plan（dogfood、delegation loop、Phase 證據評估）會持續累積量測與契約回饋。若全文塞在單一 kit 或 `_plan.md` 內：

- 檔案膨脹、transport 模板與證據混雜；
- 引用依賴**行號**（`L449`、`kit §2d′`），編輯後易漂移；
- IDE / agent 讀檔快取與磁碟不一致時，定位錯誤。

**解法**：每個 plan folder 使用 **`evidence/` 子目錄** 存放證據全文；plan 本體與 kit 只留**摘要 + 檔案路徑連結**。

## 目錄佈局

```text
plans/{active|archived}/<YYYY-MM-DD-HHMM-slug>/
├── _plan.md                 # 決策與 checkbox；引用 evidence/ 檔案，不寫行號
├── 01-<companion>.md        # 可選：transport kit、audit brief 等（NN- 前綴）
└── evidence/
    ├── README.md            # 必填：索引 + 引用規則
    └── <run-id>-<slug>.md   # 各 run 全文（量測欄、契約回饋、仲裁表）
```

**與 flat cluster 的關係**：頂層 `<slug>.md` + `<slug>-dogfood-evidence.md` 應 **folderize** 為上列佈局；dogfood 全文進 `evidence/`，不要與 `_plan.md` 同檔堆疊。見 [`plans/README.md`](../../plans/README.md) §扁平多檔 → 資料夾。

## 引用規則（強制紀律）

| 規則 | 說明 |
|---|---|
| **檔案路徑** | 引用寫 `evidence/<file>.md` 或 markdown 連結，**禁止** `kit L449`、`README L62–95` 類行號 |
| **標題錨點** | 需定位時寫檔內 `## 量測欄` 或表格 `#` 欄，不用絕對行號 |
| **索引同步** | 每個 `evidence/*.md`（除 `README.md`）必須列於 `evidence/README.md` 的 Run 索引表 |
| **去敏** | 遵循 [`enforcement/sanitization.md`](../enforcement/sanitization.md)；class 名、live host、inner commit 細節留 consumer project plan |
| **新 run** | 新增 evidence 檔 + 更新 `evidence/README.md` 同一 commit（或 README 已含該檔連結） |

## README.md 必填結構

`evidence/README.md` 至少含：

1. **引用規則**（可複製 [`plans/templates/plan-evidence/README.md`](../../plans/templates/plan-evidence/README.md) 首段）
2. **Run 索引**表：`| Run ID | 檔案 | 狀態 | 摘要 |`

## 機械驗證（commit-msg）

| Validator | Severity | 規則 |
|---|---|---|
| `validatePlanEvidenceConvention` | **block** | staging `.../evidence/**` 時：`evidence/README.md` 必須存在；每個 `evidence/*.md`（除 README）須在 README 內被引用；README 須含「引用規則」與「Run 索引」 |
| `validatePlanTreeFolderConvention` | warning | `evidence/` 內檔名**豁免** `NN-` 前綴；`plans/.../<slug>/evidence/<file>.md` 深度視為合法（不觸發 depth≥3 warning） |
| `warnPlanEvidenceLineNumberCitations` | **warning**（不 block） | 同 plan folder 內 staged `.md` 出現 `\bL\d+\b` 行號引用時提醒改用檔案路徑 |

Opt-out：`[skip-plan-evidence]` 獨立 trailer 行（緊急遷移 only）。

手動全庫檢查：`ai-skill plans validate --root <repo>` 目前覆蓋 portable plan-tree engine；plan evidence 規則由 **commit-msg** 在變更當下強制。

## 範例

- 索引 + 規則：[`plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/evidence/README.md`](../../plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/evidence/README.md)
- Run 全文：[`evidence/2d-prime-externalrepoc-module-alignment.md`](../../plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/evidence/2d-prime-externalrepoc-module-alignment.md)
- 模板：[`plans/templates/plan-evidence/README.md`](../../plans/templates/plan-evidence/README.md)
