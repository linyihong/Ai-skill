# Tools — methods（Go CLI）

機械方法入口；**不是** Selection Actor。本目錄不再放長期 `.py` / `.sh` phase gate 產品腳本。

Canonical 實作：`ai-skill image-processing`（`scripts/ai-skill-cli/internal/app/image_processing.go`）。

## Methods

| Method | CLI | Role |
| --- | --- | --- |
| Generate synthetic fixtures | `ai-skill image-processing generate-fixtures --repo <AI_SKILL_REPO>` | 合成 `neg.*` / `pos.*` + refs／consumer profile |
| Independent verify | `ai-skill image-processing verify <image> [--profile] [--source] [--consumer-profile] [--json-out] [--expect-block]` | 獨立驗收 → checks + finality |

## 組裝驗收（需要時才寫 phase）

本 workflow **只提供方法**。後續 plan phase 若需要 gate，依流程自行組裝，例如：

```text
# 1) 重新生成合成負向（可選）
<AI_SKILL_REPO>/scripts/ai-skill-cli/bin/ai-skill-<os>-<arch> \
  image-processing generate-fixtures --repo <AI_SKILL_REPO>

# 2) 對每個應擋下的負向 fixture
<AI_SKILL_REPO>/scripts/ai-skill-cli/bin/ai-skill-<os>-<arch> \
  image-processing verify \
  <AI_SKILL_REPO>/workflow/image-processing/fixtures/inputs/neg.alpha-opaque.png \
  --profile neg.alpha-opaque \
  --json-out <AI_SKILL_REPO>/workflow/image-processing/fixtures/reports/neg.alpha-opaque.verify.json \
  --expect-block
```

建議負向最小集（方法層約定，不是預先提交的 phase 腳本）：

- `neg.alpha-opaque`
- `neg.residue-halo`
- （擴張時）`neg.interior-damage`、`neg.over-erode`、`neg.consumer-layout`
- 正向對照：`pos.d-canvas-328`（**不要**加 `--expect-block`）

## Adapter 注意（非 repo 入口）

- rembg / Pillow 等第三方執行器屬 **consumer／本機 adapter**，授權見 [`../adapters/license-checklist.md`](../adapters/license-checklist.md)。
- 不得把第三方 CLI 寫進 canonical `execution-flow.md` 步驟名。
- Execute ≠ Verify：執行器不得自證 PASS；驗收一律走 `image-processing verify`。

## 禁止

- 新增長期 `.py` / `.sh` / `.rb` / `.js` 工具入口（pre-commit `validateNoNewShellScripts` 會擋）
- 預先提交 `run_phaseN_gate.*` 當產品；phase 數字只在 plan 需要時再寫
