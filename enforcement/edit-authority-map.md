# 編輯權限對照表（Edit-Authority Map）

本檔把 [`rule-weight.md`](rule-weight.md) 的 P0–P3 權重體系，翻成一張**顯式對照表**：某類檔案，較弱的模型可以自行改、還是必須先問使用者、還是禁改。目的是讓弱模型**不必自己從 P0–P3 推導**——推導容易錯，查表不會錯。

> **怎麼用**：動任何檔案前，先在下表找它屬於哪一類，照該列的權限行事。找不到對應類別 → 當作「需確認」處理（保守預設）。
>
> **術語**：表中「rule_class」= enforcement 規則的分類條目，定義在 [`enforcement-registry.yaml`](enforcement-registry.yaml)（每條規則對應一個 `rule_class` id + coverage 等級）。判斷「有沒有動 rule_class」= 你的改動有沒有新增 / 刪除 / 改變某條規則的 `rule_class` 條目或其 coverage；純改規則說明文字通常不動 rule_class。不確定 → 當作有動，走「需確認」。

## 權限三級定義

| 級別 | 意義 | 弱模型行為 |
|---|---|---|
| **自改（autonomous）** | 屬於被交代任務的正常產出，可回退，錯了容易發現 | 直接改，在回報中說明改了什麼 |
| **需確認（confirm-first）** | 改動有較高影響 / 較難回退 / 可能超出 scope | 先向使用者說明要改什麼與為什麼，等明確同意再改 |
| **禁改（forbidden-without-explicit-instruction）** | 不可逆 / source-of-truth 完整性 / 需治理程序 | 除非使用者明確要求，否則不動；即使被要求也先說明風險 |

## 對照表

| 檔案類別 | 路徑範例 | 權限 | 觸發特別注意 |
|---|---|---|---|
| Plan 正文（active） | `plans/active/**` | **自改** | 遵守 plan template 必填章節 + checkbox sync |
| Feedback lesson（新增） | `feedback/history/<domain>/*.md` | **自改** | 先去敏（[`sanitization.md`](sanitization.md)）；抽象化非專案細節 |
| Workflow / analysis / intelligence 正文 | `workflow/**` `analysis/**` `intelligence/**` | **自改** | 改後跑 linked-updates；動 execution-flow 先讀對應 README |
| Governance 規則正文 | `governance/**/*.md` | **自改** | 套 [`../governance/weak-model-rule-authoring.md`](../governance/weak-model-rule-authoring.md) 四要件 |
| Enforcement 規則正文 | `enforcement/*.md` | **自改** | 若動 rule_class → 必須同步 `enforcement-registry.yaml`（否則 commit 被 `validateEnforcementRuleRegistrySync` 擋） |
| README 索引 | 各層 `README.md` | **自改** | 只改索引 / 連結；內容 truth 在被指向的檔 |
| Enforcement registry | `enforcement/enforcement-registry.yaml` | **需確認** | coverage 轉換走 Status Transition Matrix；動前讀 [`enforcement-registry.md`](enforcement-registry.md) |
| Runtime YAML canonical | `runtime/*.yaml` | **需確認** | 改後必須 `ai-skill runtime compile + refresh`；`runtime_projection` 規則嚴格 |
| Metadata schema | `metadata/**` | **需確認** | schema 欄位變動影響多個 consumer |
| Bootstrap 入口 | `CLAUDE.md` `AGENTS.md` `GEMINI.md` `CORE_BOOTSTRAP.md` | **需確認** | thin pointer 規則（`validateBootstrapEntryThinness`）；新 obligation 加到 canonical source 不加這裡 |
| Go CLI 原始碼 | `scripts/ai-skill-cli/**/*.go` | **需確認** | 改後必須 `go test ./...` + rebuild `bin/`（否則 pre-push 擋）；動 hooks.go 須同步 registry |
| Committed 二進位 | `scripts/ai-skill-cli/bin/**` | **需確認** | 只透過 `releasebuild` 重建，不手改；stamp latest CLI source commit |
| **runtime.db** | `runtime/runtime.db` | **禁改（除 compile 產生）** | 只由 `ai-skill runtime compile` 生成；不手改 SQLite |
| **Constitution / ADR** | `constitution/ADR-*.md` | **禁改** | accepted decision；提案走 plan 的 Decision Rationale，completed 才 promote |
| **P0 安全規則** | `sanitization.md` `authorization-scope.md` 及其強制邏輯 | **禁改** | 削弱 safety / source integrity 屬 P0 violation；即使被要求也先說明 |
| **Git history** | `.git/**`（rewrite / force-push） | **禁改** | 不可逆、影響所有 clone；見 [`../governance/judgment-rubrics.md`](../governance/judgment-rubrics.md) R3 |
| 其他 repo（consumer project） | Ai-skill 外的專案 | **需確認** | 跨 repo 改動先確認 scope 與該 repo 的規則 |

## 完成判準（二值）

動檔前能回答：這個檔屬於表中哪一類、對應權限是什麼？

- **正例（改既有檔）**：要改 `enforcement/rule-weight.md` → 查表屬「Enforcement 規則正文 = 自改」，但注意欄提醒「動 rule_class 要同步 registry」→ 確認本次沒動 rule_class，自改，跑 linked-updates。
- **正例（新增檔）**：要新增 `governance/foo.md` → 查表屬「Governance 規則正文 = 自改」→ 自改，但注意欄要求套 [`../governance/weak-model-rule-authoring.md`](../governance/weak-model-rule-authoring.md) 四要件；新增後把它加進 `governance/README.md` 索引（linked-updates）。
- **反例（新增檔）**：新增 `constitution/ADR-099-xxx.md` 宣稱某決策 accepted → 違反「Constitution / ADR = 禁改」；ADR 只能在 plan completed + 通過 promotion criteria 後建立，不可自行新增。
- **反例（改既有檔）**：要改 `runtime/runtime.db` → 沒查表，直接用 SQLite 工具手改 → 違反「禁改（除 compile 產生）」，改動不會進 canonical YAML，下次 compile 被覆蓋，且可能過不了 freshness 檢查。

## 誠實條款

表中沒有的檔案類別、或無法判斷屬於哪類 → **不要猜**，當作「需確認」處理並向使用者說明你的不確定。保守預設優於錯誤自改。

## 誰會參考這裡（Inbound References）

- [`plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md`](../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md) — Workstream F / Phase 5b F5
- [`../governance/judgment-rubrics.md`](../governance/judgment-rubrics.md) — R3（停下問使用者）引用本表判斷 scope

## 與既有層的關係

- [`rule-weight.md`](rule-weight.md) — 本表是其 P0–P3 體系的顯式投影；權重衝突的完整判斷仍以 rule-weight 為 canonical。
- [`dependency-reading.md`](dependency-reading.md) — 「自改」類仍須完成其回寫閉環（diff / linked-updates / commit / push / readback）。

← [回到 enforcement 索引](README.md)
