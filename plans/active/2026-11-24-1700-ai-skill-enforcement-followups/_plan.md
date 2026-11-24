---
id: 2026-11-24-1700-ai-skill-enforcement-followups
plan_kind: main
status: draft
owner: linyihong
owner_layer: enforcement
created: 2026-11-24
parent: null
---

# Ai-skill Enforcement Follow-ups（Project Tooling 計畫的遺留項目）

**Status**: draft — 只建立，尚未開始。每一項都要等使用者開放才做。

## 一句話目標

承接 [`2026-11-23-0834-project-tooling-governance`](../../archived/2026-11-23-0834-project-tooling-governance/_plan.md) 過程中發現、但不屬於 Project Tooling 契約的 Ai-skill 自身治理問題，逐項裁決或修復。

## 來源

[`phase4-decision-package.md`](../../archived/2026-11-23-0834-project-tooling-governance/evidence/phase4-decision-package.md) §7。每一項的原始證據在來源計畫的 evidence 目錄。

## 項目

| # | 項目 | 狀態 | 來源 | 需要先決定的事 |
| --- | --- | --- | --- | --- |
| R1 | registry 有 46 個 rule_class，超過 `rule_class_hard_limit` 40；`current_count` 仍寫 40 | observed | F1 §5 | 調高上限、合併 class，還是維持超限並寫明理由 |
| R2 | `executor_kind_spec.enum` 只列 4 種，registry 實際用了更多種，lint 沒有檢查 | observed | F1 §5 | 補 enum 加 lint，還是承認 kind 是自由欄位 |
| S1 | core-bootstrap 的 obligation severity 只是宣告，dispatcher 不照它執行 | observed | F2 §5、§7 | severity 是描述用還是有約束力 |
| S2 | portable/kge 的 `CLIDocSyncRule` 要不要在可攜層 deprecated | 盤點完成：本機範圍內沒有外部使用者（observed） | F2 §7、S2 盤點 | 要不要 deprecated；若要，README 的複製範例也要一起改 |
| V2 | `runtime validate` 的既有腳本檢查（`nativeToolBootstrapShellCLIDecisionValidation`）只看 `.sh` | observed | F6 §4 | 要不要和 T1 共用同一份 forbidden 範圍；先盤點既有的非 `.sh` 腳本有沒有缺遷移決策 |
| W1 | `warnCLIDocSync` 在 2026-12-24 再評估是否移除 | 待執行紀錄來源 | F2 §3、§7 | 先決定「它在觀察期間確實執行過」用什麼紀錄證明 |

## 完成條件

- [ ] 每一項都有使用者裁決（修復、接受現況並寫明理由，或 deferred 並附重啟條件）
- [ ] 需要修改程式的項目都通過 Ai-skill hooks 並已推送
- [ ] Plan Completion Closure 執行完畢

**Glossary Impact**: no new framework vocabulary introduced

## 與其他 plans 的關係

- 來源：[`2026-11-23-0834-project-tooling-governance`](../../archived/2026-11-23-0834-project-tooling-governance/_plan.md)（已結案）。
