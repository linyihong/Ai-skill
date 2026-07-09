# Dogfood Prompt Kit（三角色 loop 傳輸模板，Cursor-first）

> **本檔是 rendered transport artifact，不是 canonical contract。** Canonical contract（verifier 報告 4 欄位、仲裁三處置、3 條 role boundary invariants）**已定稿於 [`plans/README.md`](../../README.md) §Delegation「派發 → 獨立驗證 → 仲裁（loop SOP）」子節**（Q2 resolved，2026-07-08，由 dogfood 2b 委派落地）；[`_plan.md`](_plan.md) §Decision Rationale 為決策紀錄。本 kit 只跟著 canonical render。模板刻意 self-contained（貼進 fresh session 的人不讀本 repo）——這是 transport 的本質，不是 dual source。
>
> **Tool-neutral core + Cursor-first transport**：模板 A/B/C 本體不含任何工具指令；§Cursor 傳輸備註是 Layer 3 adapter 細節（[`ai-tools/agent/cursor.md`](../../../ai-tools/agent/cursor.md)）。換 Claude Code / Codex / 其他工具時只換 §傳輸備註，模板不動。

## 使用流程（orchestrator = 你的主 session）

1. **Orchestrator** 在主 session 填好 brief（goal / acceptance / verification / context.required），填入模板 A。**填 brief 教訓（2b）**：若任務目標是 reusable 文件（README / workflow / enforcement），acceptance 必含 tool-neutral 措辭條款，否則 executor 寫入工具專屬詞不算違規。**填 brief 教訓（2026-07-08）**：在 `verification_backfill`（或 brief 附錄）為每條 acceptance 標 `executor`（happy path 測試）或 `verifier_only`（負面 / 架構 / 禁止事項）；避免 verifier 只重跑 executor 自寫測試。
2. 開一個 **全新** Cursor chat（executor），貼模板 A。executor 交付 branch + 自驗報告。
3. 開 **另一個全新** Cursor chat（verifier，不可沿用 executor 的 chat），貼模板 B（含同一份 brief 的 acceptance/verification + executor 的 branch/diff 位置）。
4. Verifier 報告帶回主 session，orchestrator 用模板 C 逐條仲裁，記錄量測欄位。
5. `fix` 項：把模板 C 產出的補充指示貼回 **executor 原 chat**（保留其 context），修完重跑步驟 3。

**鐵則**：主 session 全程不寫 code、不自己驗。忍不住動手 = 越界信號，記進模板 C 的量測欄。

---

## 模板 A — Executor prompt（貼進全新 Cursor chat）

```text
你是本任務的「執行者」。你只憑下面這份 brief 完成任務，不需要也不應該去找其他規劃文件。

## 任務目標（goal）
<一句話目標>

## 驗收條件（acceptance — 做到什麼算完成）
1. <條目，可觀察、可驗證>
2. <...>

## 驗證方式（verification — 你交付前要自己跑過；限 happy path）
1. <指令 / 操作 / 預期結果>
2. <...>

## 測試範圍（由 orchestrator 指定）
- **你負責（executor）**：<happy path 整合測試 / 自驗命令>
- **不由你寫（verifier_only）**：<負面 case、架構禁止事項檢查 — 留給獨立 verifier>

## 必要 context（先讀這些，應該就夠了）
- <path/to/file>：<為什麼需要>
- <...>

## 工作紀律
1. 在新 branch 工作：`<branch-name>`。
2. 只做 acceptance 範圍內的改動；發現範圍外的問題「記下來回報」，不要修。
3. 只改必須改的行，遵循周圍既有 code style；不加順便的重構或清理。
4. 若 brief 資訊不足以完成，停下來明說缺什麼，不要猜。

## 交付格式（回覆最後必須包含）
### 摘要
<做了什麼，2-4 句>
### Diff 位置
branch：<branch-name>（已 commit，未 merge）
### 自驗結果
| verification 條目 | 結果 pass/fail | 證據（輸出/截圖/觀察） |
### 範圍外發現（只記不修）
- <發現 / 無>
### 讀檔紀錄（誠實列出）
- 「必要 context」清單內：<列出>
- 清單外另外讀的：<列出 / 無 — 這欄用來改進 brief，不是考核你>
### 阻塞
<無 / 缺什麼>
```

## 模板 B — Verifier prompt（貼進另一個全新 Cursor chat）

```text
你是本任務的「獨立驗證者」。你的立場是 fault-finding：主動找碴，但「只產出證據，不做決定」。

**禁止降級**：不可只重跑 executor 的自驗命令就結案。你必須完成下面三層驗證。

## 你驗證的對象
執行者依下面的驗收條件交付了改動：
- branch：<branch-name>
- 檢視改動：`git diff <base-branch>...<branch-name>`

## 驗收條件（acceptance — 逐條驗，這是唯一量尺）
1. <與模板 A 完全相同的條目>
2. <...>

## 驗證方式（verification — L1：重跑 executor 自驗，必要但不充分）
1. <與模板 A 完全相同的條目>
2. <...>

## verifier_only case（L3：對抗性驗證 — 你必須執行或補寫測試）
1. <負面 / 邊界 case；若 executor 未寫對應測試，你可補寫並執行>
2. <架構 / 禁止事項靜態檢查（grep、classpath、契約欄位）>

## 三層驗證紀律
1. **L1 重跑**：實際執行 `verification` 列出的命令，記錄 pass/fail 與輸出摘要。
2. **L2 讀碼審查**：讀 diff，對照 acceptance 與架構禁止事項；**不依賴** executor 測試即可發現的違規也要記 finding。
3. **L3 對抗性**：逐條執行 `verifier_only` case；未覆蓋 → `acceptance-violation`（即使 L1 全綠）。
4. 你「不可以」：提出應該怎麼修的仲裁決定、擴大驗證範圍去審整個 repo。你「可以」：為 `verifier_only` case **補寫並執行**測試（在 branch 上或獨立驗證腳本），只產證據。
5. diff 範圍內發現但 acceptance 沒涵蓋的問題：照記，classification 標 out-of-scope 或 observation。
6. 沒有問題也是有效結果：輸出「全數通過」+ 每條的通過證據（含 L2/L3），不可以只回「看起來沒問題」。

## 回報格式（固定，每條 finding 一列）
| # | evidence（具體檔案/行為觀察，可覆核） | acceptance_ref（對應第幾條，或 beyond-acceptance） | classification（acceptance-violation / out-of-scope / observation） | status（observed / verified / refuted） |

### Acceptance 逐條結論
| acceptance 條目 | pass / fail | 證據（含 L1/L2/L3 哪一層抓到） |
### verifier_only 覆蓋表
| case # | 已執行 / 已補測 | pass / fail | 證據 |
### 讀檔紀錄
<驗證過程實際讀了哪些檔案、跑了哪些指令、補了哪些測試>
```

## 模板 C — Orchestrator 仲裁表（主 session 使用）

驗證報告回來後，逐條仲裁，記錄於被委派任務的 plan artifact（dogfood 量測全文見 [`evidence/`](evidence/)）：

```markdown
### 仲裁紀錄（<日期> / <任務>）
| finding # | 處置（fix / defer / reject） | 理由 | 後續 |
|---|---|---|---|
| 1 | fix | 違反 acceptance 2 | 補充指示已回 executor，重驗 |
| 2 | defer | 真實但超出本次 scope | 轉 <observation / 新 plan / evidence candidate> |
| 3 | reject | 覆核不成立：<理由> | 標 refuted，留證據 |

### fix 項補充指示（貼回 executor 原 chat）
針對驗證發現的問題修正，仍只憑原 brief 的 acceptance 為準：
- finding 1：<具體要修什麼、對應哪條 acceptance>
修完後重新輸出「交付格式」全部欄位。

### 量測欄（dogfood evidence，每輪必填）
| 指標 | 值 |
|---|---|
| verifier 差集（verifier 抓到、executor 自驗沒抓到） | <n 條 + 列舉> |
| verifier 降級（只跑 L1、未做 L2/L3？） | <是 / 否> |
| 仲裁分佈 | fix <n> / defer <n> / reject <n> |
| orchestrator 越界（動手寫 code？被迫回讀 diff 細節？） | <無 / 描述 — Q1 信號> |
| verifier 報告自足性（不回讀 diff 即可仲裁？） | <是 / 否，缺哪個欄位> |
| executor 讀檔差集（context.required 以外） | <無 / 列舉 — 回饋 brief> |
| 契約缺漏回饋 | <無 / 修模板哪一段 → v2> |
```

## Cursor 傳輸備註（Layer 3，僅此節綁工具）

- **Fresh context = 全新 chat**：executor 與 verifier 各開一個新 chat；verifier 絕不沿用 executor 的 chat（Cursor 同 workspace 的 chat 之間不共享對話記憶，滿足 fresh-context invariant）。
- **交接靠 git，不靠對話**：executor 交付 commit 到 branch；verifier 用 `git diff <base>...<branch>` 取改動。不要用「上一個 chat 說了什麼」傳遞。
- **Cursor 的 codebase 自動索引**會讓 executor/verifier 可能讀到 context.required 以外的檔——這不違反協議（協議管「brief 是否自足」，不禁讀），但「讀檔紀錄」要誠實填，作為 brief 改進信號。
- **在 Ai-skill repo 內 dogfood 時**（Phase 2b）：Cursor 走 `.cursor/rules`，不經 Claude Code PreToolUse bootstrap gate；但 brief 的 context.required 仍應含該任務相關的 canonical 規則檔，讓兩種工具行為一致。
- 換工具：Claude Code 用 Agent tool（可選 worktree isolation）跑同一份模板 A/B 文字；模板本體不改。
- **Task subagent transport**（2c / 2d）：orchestrator 用 Cursor `Task` spawn Executor / Verifier；**省略 `model`**，子 agent 繼承主 session（stakeholder 2026-07-08，控制 token 成本）。
- **Execute 意圖 hook allowlist**（2d 契約回饋）：orchestrator 應可寫 `<PROJECT_ROOT>/docs/plans/**`、`.ai-skill/project/**`、`tests/**`；只 deny 內層實作路徑（如 `manageCode/server/**`）。否則 orchestrator 連 plan patch 也被擋，被迫讓 Executor 代寫外層 artifact。

## Dogfood 證據索引（精簡）

> **全文與仲裁表**在 [`evidence/`](evidence/)（見 [`evidence/README.md`](evidence/README.md)）。本 kit 只保留傳輸模板；新 run 一律新增 `evidence/<run-id>-<slug>.md`。

| Run | 狀態 | 關鍵量測（摘要） | 全文 |
|---|---|---|---|
| **2a** | ✅ | violation **0**；fix 0 / defer 5 / reject 2；orchestrator 越界 **無** | [`evidence/2a-software-delivery-review-invoke.md`](evidence/2a-software-delivery-review-invoke.md) |
| **2b** | ✅ | violation **0**；fix 1 / defer 1 / reject 1；brief 自足 ★★★★☆+ | [`evidence/2b-plans-sop-expansion.md`](evidence/2b-plans-sop-expansion.md) |
| **2a-external** | ✅ | violation **0**；orchestrator pre-loop 越界 **1**；6 tests + 2 bug fix | [`evidence/2a-external-sync-adapter-step6.md`](evidence/2a-external-sync-adapter-step6.md) |
| **2c** | ✅ | violation **2**/8 slices；spawn ~18；slice 紀律後越界 **0** | [`evidence/2c-tiered-archive-platform.md`](evidence/2c-tiered-archive-platform.md) |
| **2d** | ✅ | consumer overlay + backfill；gate 後 manageCode 越界 **0** | [`evidence/2d-outbound-sync-phase3.md`](evidence/2d-outbound-sync-phase3.md) |
| **2d′** | 證據 only | integration UX fail **1**；`remote_absent_delete` fix | [`evidence/2d-prime-externalrepoc-module-alignment.md`](evidence/2d-prime-externalrepoc-module-alignment.md) |
| **2h** | 證據 only | violation **≥2**；RBAC + api-surface gate 回饋 | [`evidence/2h-externalrepoc-common-url-verification-gaps.md`](evidence/2h-externalrepoc-common-url-verification-gaps.md) |
| **2e** | 完成 | Research grandfather sunset | [`02-grandfather-sunset-audit.md`](02-grandfather-sunset-audit.md) |
