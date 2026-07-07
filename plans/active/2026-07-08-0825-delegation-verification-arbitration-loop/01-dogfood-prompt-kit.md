# Dogfood Prompt Kit（三角色 loop 傳輸模板，Cursor-first）

> **本檔是 rendered transport artifact，不是 canonical contract。** Canonical contract（verifier 報告 4 欄位、仲裁三處置、3 條 role boundary invariants）在 [`_plan.md`](_plan.md) §Decision Rationale；Q2 定稿後 canonical 落點移至該處，本 kit 只跟著 render。模板刻意 self-contained（貼進 fresh session 的人不讀本 repo）——這是 transport 的本質，不是 dual source。
>
> **Tool-neutral core + Cursor-first transport**：模板 A/B/C 本體不含任何工具指令；§Cursor 傳輸備註是 Layer 3 adapter 細節（[`ai-tools/agent/cursor.md`](../../../ai-tools/agent/cursor.md)）。換 Claude Code / Codex / 其他工具時只換 §傳輸備註，模板不動。

## 使用流程（orchestrator = 你的主 session）

1. **Orchestrator** 在主 session 填好 brief（goal / acceptance / verification / context.required），填入模板 A。
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

## 驗證方式（verification — 你交付前要自己跑過）
1. <指令 / 操作 / 預期結果>
2. <...>

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

## 你驗證的對象
執行者依下面的驗收條件交付了改動：
- branch：<branch-name>
- 檢視改動：`git diff <base-branch>...<branch-name>`

## 驗收條件（acceptance — 逐條驗，這是唯一量尺）
1. <與模板 A 完全相同的條目>
2. <...>

## 驗證方式（verification — 你要實際執行，不是看 code 猜）
1. <與模板 A 完全相同的條目>
2. <...>

## 驗證紀律
1. 逐條 acceptance 驗證，每條都要有可獨立覆核的證據（實際執行結果、檔案內容、行為觀察）。
2. 你「不可以」：修改任何檔案、提出應該怎麼修的決定、擴大驗證範圍去審整個 repo。
3. diff 範圍內發現但 acceptance 沒涵蓋的問題：照記，classification 標 out-of-scope 或 observation。
4. 沒有問題也是有效結果：輸出「全數通過」+ 每條的通過證據，不可以只回「看起來沒問題」。

## 回報格式（固定，每條 finding 一列）
| # | evidence（具體檔案/行為觀察，可覆核） | acceptance_ref（對應第幾條，或 beyond-acceptance） | classification（acceptance-violation / out-of-scope / observation） | status（observed / verified / refuted） |

### Acceptance 逐條結論
| acceptance 條目 | pass / fail | 證據 |
### 讀檔紀錄
<驗證過程實際讀了哪些檔案、跑了哪些指令>
```

## 模板 C — Orchestrator 仲裁表（主 session 使用）

驗證報告回來後，逐條仲裁，記錄於被委派任務的 plan artifact（本次 dogfood 記在本檔 §Dogfood 紀錄）：

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

## Dogfood 紀錄

### 2a — software-delivery 外部 repo 任務（待跑）
- 任務：<使用者選定的真實小任務>
- brief / 仲裁紀錄 / 量測欄：<回填>

### 2b — Ai-skill 內部任務（待跑）
- 任務：<待選>
- brief / 仲裁紀錄 / 量測欄：<回填>
