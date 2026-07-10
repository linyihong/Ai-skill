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

## Dogfood 紀錄

### 2a — software-delivery 任務 ✅（2026-07-08，Cursor session demo）

- **任務**：chat-only「Delegation loop ↔ Review invoke 整合備註」— 評估 `workflow/software-delivery` 如何定位 review invoke，以及 delegation 驗證閉環是否從 software-delivery 入口可發現（**read-only，零 commit / 零檔案寫入**）。
- **Transport**：Cursor 主 session orchestrator + **Task subagent** 模擬 executor / verifier fresh context（本輪使用者授權「就跑一次流程」；**偏離** plan 原文 2a「外部 repo + 人類貼全新 chat」— 記為 transport adaptation，模板 A/B 文字未改）。
- **Repo**：Ai-skill（software-delivery workflow 域；非外部 repo — 同樣記為 scope adaptation）。
- **Executor**（agent `c630a61a`）：acceptance 1–5 自驗全 pass；讀檔含 `intake.md`（brief 未列，見 F3）；交付完整 Artifact + 自驗表。
- **Verifier**（agent `0a3e197d`，fresh）：acceptance 1–5 獨立重驗全 pass；**0 acceptance-violation**；findings ×7（皆 observation / beyond-acceptance）。
- **Loop 關閉**：無 `fix` 項 → 無需 re-delegate / 重驗；所有 findings 已仲裁。

#### Brief（orchestrator 撰寫，摘要）
- **goal**：產出 chat-only 整合備註（Review invoke 定位 + delegation SOP 可發現性 + verifier↔fault_finding 對應）。
- **acceptance**：5 條（SD README 行號引用 / plans SOP 可發現性 / registry 證據 / tool-neutral / 自驗表）。
- **verification**：讀 context.required + grep `code-review`/`fault_finding`。
- **context.required**：`workflow/software-delivery/README.md`、`execution-flow.md`（thin index）、`plans/README.md` L62–95、`capability-registry.yaml`（code-review）。

#### 仲裁紀錄（2026-07-08 / SD Review invoke 整合備註）
| finding | 處置 | 理由 | 後續 |
|---|---|---|---|
| F1 SOP L68 未點名 `code-review` id，executor 以 registry 例證略強於字面 | defer | 語意對齊成立；點名 id 屬整合推論，非本次 acceptance 要求 | 若未來 cross-link 文件可補一句 explicit id |
| F2 自驗表 L4 稱 L58 出現在 §2，Artifact 正文未引用 | reject | 覆核：tool-neutral 正文確實無工具詞；meta 自辯措辞不準但不影響 acceptance 4 pass | 標 refuted；executor 交付格式可改進措辞 |
| F3 brief 未列 `intake.md` 但 acceptance 2 路徑需它 | defer | 暴露 **orchestrator brief 缺漏**（非 executor 失誤）；executor 自行補讀後結論正確 | brief v2 若再跑應列 `intake.md`；記入契約回饋 |
| F4 executor 讀取超出 context.required 範圍 | defer | 結論不受影響；讀檔紀錄誠實性議題輕微 | 回饋 brief 邊界即可 |
| F5 Artifact 提及 ADR-013 不在 context.required | reject | beyond-acceptance 整合推論，無 acceptance 影響 | 標 refuted |
| F6 registry 檔案級 `status: candidate` vs 條目 `active` | defer | executor 正確報告條目級；檔案級為 metadata 細節 | observation only |
| F7 可發現性斷層為真（SD README ↔ plans §Delegation 無連結） | defer | 任務核心發現，超出「修 brief 交付」範圍 | 轉 evidence candidate / 未來 cross-link plan（不屬本輪 fix） |

#### 量測欄
| 指標 | 值 |
|---|---|
| verifier 差集（executor 自驗沒抓到） | 7 條（皆 beyond-acceptance / observation；acceptance-violation **0**） |
| 仲裁分佈 | fix **0** / defer **5** / reject **2** |
| orchestrator 越界 | **無** — 未寫 code / 未 commit；仲裁全憑 verifier 報告，**未回讀**原始檔驗證行號 |
| verifier 報告自足性 | **是** — Q1 正向證據 ×2（2b + 本輪） |
| executor 讀檔差集 | **有** — `intake.md`（brief 未列，F3） |
| 契約缺漏回饋 | orchestrator brief 缺 `intake.md` 於 context.required（acceptance 2 路徑依賴） |
| transport adaptation | Task subagent 代替人類全新 chat；Ai-skill 代替外部 repo — 雙項偏離 plan 原文 2a，已記錄 |

### 2b — Ai-skill 內部任務 ✅（2026-07-08，agent transport）

- **任務**：plans/README.md §Delegation loop SOP 擴充（本 plan Phase 1 真實待辦，委派執行）。
- **Transport**：Claude Code Agent — executor（worktree 隔離，branch `worktree-agent-adeec…`）+ verifier（fresh agent，唯讀）；orchestrator = 主 session。兩個 agent 首則回覆均輸出 Bootstrap Receipt（bootstrap 注意事項實測通過）。
- **Executor**：commit `af26064`（35 行純新增，單檔）；自驗 8/8 pass；資訊不足時未猜測。fix commit `2d5bc60`（單行）。
- **Verifier round 1**：acceptance 8/8 pass；findings ×3（F1 invariant 措辭一字 drift / F2 consumer 層出現「PreToolUse」工具專屬詞 / F3 省略括號補述），全部 beyond-acceptance 觀察級，0 violation。
- **Fix 重驗（delta）**：pass — 單行替換；替換 gate id `gate.bootstrap.receipt_present` 經 grep 驗證存在於 canonical yaml；acceptance 1–8 無回退。

#### 仲裁紀錄（2026-07-08 / SOP 擴充）
| finding | 處置 | 理由 | 後續 |
|---|---|---|---|
| F1 措辭 drift（`dogfood evidence`→`evidence`） | defer | 由 Q2 決策收斂：README 落地即成 canonical，_plan.md 降為決策紀錄，無雙源同步義務；一般化措辭反而正確 | Q2 resolved 回寫 _plan.md |
| F2 「PreToolUse」入 consumer 層 | fix | 暴露 brief 契約缺漏（reusable doc 目標缺 tool-neutral 條款）→ brief v2 追加 acceptance 9 後回派 | 重驗 pass；教訓寫入本 kit §使用流程 |
| F3 省略括號補述 | reject | 覆核不成立：省略符合 Layer 3 邊界；「結構化 findings」由 4 欄位表承載，無語意損失 | 標 refuted 留證據 |

#### 量測欄
| 指標 | 值 |
|---|---|
| verifier 差集（executor 自驗沒抓到） | 3 條（皆 beyond-acceptance 觀察級；acceptance 內 0） |
| 仲裁分佈 | fix 1 / defer 1 / reject 1 |
| orchestrator 越界 | 無 — 未寫任何實作 diff；整合（merge / worktree 清理）與 plan 回寫屬 orchestrator artifact 職責 |
| verifier 報告自足性 | **是** — 仲裁全憑報告引文完成，orchestrator 未回讀 diff（Q1 正向證據 ×1） |
| executor 讀檔差集 | 無 — 只讀 context.required（brief 自足 ★★★★☆+） |
| 契約缺漏回饋 | brief v1 缺 reusable-doc tool-neutral 條款（F2 暴露）→ v2 補 acceptance 9；已回寫模板使用說明 |
| executor 範圍外發現 | plans/README.md L261 歷史狀態表含工具詞（歷史紀錄描述，不動）；_plan.md checkbox 回寫屬 orchestrator（已處理） |

### 2a-external — 外部 monorepo sync-adapter Step 6 ✅（2026-07-08，Cursor Task transport）

> **專案證據邊界**：sub-plan 檔名、class/test 名、remote host、live 環境細節留於 `<PROJECT_ROOT>` 被委派 sub-plan §8 execution records；Ai-skill 只保留 generalized dogfood metrics（依 [`enforcement/sanitization.md`](../../../enforcement/sanitization.md) + [`reusable-guidance-boundary.md`](../../../enforcement/reusable-guidance-boundary.md)）。

- **任務**：Admin 直寫 CRUD → remote outbound sync（delegated sub-plan Step 6）；含 mock 單測 + test-data-marker live IT。
- **Transport**：Cursor orchestrator + **Task subagent** executor/verifier（偏離 plan 原文「人類開全新 chat」— 同 2a demo 記為 transport adaptation；**外部 repo** 條件成立：`<PROJECT_ROOT>` 外層 + `<INNER_REPO>` Java 內層）。
- **Repo**：`<PROJECT_ROOT>`（plan/brief）+ `<INNER_REPO>/manageCode`（實作 submodule）。
- **Slices**：mock（6-1–6-4）→ live（6-5）；各一輪 executor + verifier。

#### Executor / Verifier 軌跡
| Slice | Executor (Task) | Deliverable commits | Verifier (Task) | Outcome |
|---|---|---|---|---|
| Mock | `<TASK-id-m>` | `<HASH-a>`（orchestrator 越界，見下）、`<HASH-b>` | `<TASK-id-v-m>` | L1–L3 pass；acceptance-violation **0** |
| Live | `<TASK-id-e>` | `<HASH-c>` | `<TASK-id-v-e>` | L1 rerun pass；live IT pass |

#### Brief（orchestrator 撰寫，摘要）
- **goal**：Admin create/update/delete/purge 在 `approved` + `remote_id` 時 outbound push；mock 覆蓋 + live 僅 test-data marker 列。
- **acceptance**：§6 條（create/update push、delete guard、unit tests、live IT、無 production sync-remote 污染）。
- **verification**：`mvn test` mock + live IT profile；verifier 獨立重跑 L1。
- **context.required**：admin service layer、outbound sync adapter、remote envelope helper、delegated sub-plan §6。

#### 實作品質差集（相對「單 session 直接改 code」反事實）
| 項目 | 直接實作（本 session 早期） | Loop 後 |
|---|---|---|
| 單測 | 可能延後 | 5 mock tests（service-layer outbound sync suite） |
| Live 驗收 | 可能跳過 | 1 live IT；僅 test-data marker |
| 已知 bug 關閉 | — | remote-id null guard（`<HASH-b>`）；remote envelope 裸 SPI unwrap（`<HASH-c>`，否則 live create 假 502） |
| Plan 證據 | 無 | `<PROJECT_ROOT>` sub-plan §8 + verification_backfill |

#### 仲裁紀錄（2026-07-08 / sync-adapter Step 6）
| finding | 處置 | 理由 | 後續 |
|---|---|---|---|
| O1 batchImport outbound 未補 service-layer test | defer | beyond mock acceptance 字面；batchImport 路徑已 wire | 若後續 slice 需覆蓋可開 observation |
| O2 purge* outbound 未單獨 mock | defer | purge 共用 delete push；delete test 已 cover guard | observation |
| O3 502 envelope 路徑未 mock | defer | live IT 已驗證 unwrap fix；mock 502 非 acceptance 必須 | observation |
| V1 orchestrator 在 loop 建立前直接改 admin service layer | defer | **Q1 負向信號**：`<HASH-a>` 由 orchestrator 手寫；後續 mechanical reminders（project overlay hooks、BDD gate）補上 | 記入越界指標；非 verifier finding |

#### 量測欄
| 指標 | 值 |
|---|---|
| verifier 差集（verifier 抓到、executor 自驗沒抓到） | **acceptance-violation 0**；beyond-acceptance observation **3**（O1–O3）；implementation gap **2**（guard + envelope，均由 executor 在交付前測試/IT 暴露，verifier L1 確認） |
| verifier 降級（只跑 L1、未做 L2/L3？） | Mock：**否**（L1–L3）；Live：**L1 rerun only**（delta 小，合理） |
| 仲裁分佈 | fix **0** / defer **4** / reject **0** |
| orchestrator 越界（動手寫 code？被迫回讀 diff 細節？） | **有（1）** — loop 前 `<HASH-a>` 直接改 Java；loop 後 **未**回讀 diff 細節即可仲裁 |
| verifier 報告自足性 | **是** — 兩輪仲裁憑 verifier 交付 + commit hash；orchestrator 未為仲裁 grep 原始碼 |
| executor 讀檔差集 | **輕微** — live slice 讀 envelope helper（brief 已列） |
| 契約缺漏回饋 | brief 可明示「orchestrator 不得 pre-loop 直接 commit 實作」→ 已用 `<PROJECT_ROOT>` project overlay + sessionStart hook 機械化 |
| transport adaptation | Task subagent 代替人類全新 chat（同 2a demo） |
| **協調成本（量化）** | Task spawn **4**（2 executor + 2 verifier）；外層 plan commit **6**；內層 `<INNER_REPO>` commit **3** |
| **品質成本（量化）** | 新增 test **6**（5 mock + 1 live IT）；acceptance-violation **0**；live 安全約束 **100%**（僅 test-data marker） |

#### 結論摘要（orchestrator 負擔 vs 品質 — 供 Q3）
| 維度 | 相對單 session 直接實作 | 證據 |
|---|---|---|
| Orchestrator **寫 code 負擔** | **↓**（loop 後為 0；loop 前 1 次越界） | `<HASH-a>` vs `<HASH-b>`/`<HASH-c>` 全由 executor |
| Orchestrator **協調/文件負擔** | **↑** | brief、spawn×4、§8 backfill、6 外層 commit |
| 交付品質 | **↑** | 6 tests、2 bug fix、0 acceptance-violation、可審計軌跡 |
| Verifier 邊際價值（本任務） | **中等** — 差集主要在「強制 IT + 結構化 defer」，非大量 acceptance catch | envelope/guard 由 executor 測試暴露；verifier 提供獨立 L1 與 observation 分類 |
| 適用邊界 | **多步 acceptance + live 驗收** 划算； trivial 單檔 fix 不划算 | 與 plan §Trade-offs advisory 一致 |

### 2c — 外部 monorepo tiered archive platform（全線 A–D）✅（2026-07-08，Cursor Task transport）

> **專案證據邊界**：slice 級 inner commit、EC2 host、PG endpoint、class 名留於 `<PROJECT_ROOT>` tiered plan §12 + inner `Data/docs/tiered-data-archive.md`；Ai-skill 只保留 generalized dogfood metrics。

- **任務**：分層資料歸檔平台 Phase A / A′ / B / C / D 全線交付（8 slices）；首個 domain adapter = verification-code retention。
- **Transport**：Cursor orchestrator（主 session）+ **Task subagent** executor / verifier（每 slice 一輪；fix 項再 spawn）；**外部 repo** 成立（`<PROJECT_ROOT>` 外層 plan + `<INNER_REPO>/manageCode`）。
- **Repo**：`<PROJECT_ROOT>` `feature/<slug>` + `<INNER_REPO>` 同 branch。
- **Brief 來源**：外層 tiered plan §12 `delegation.brief` + `.ai-skill/project/rules/plan-delegation-execution-loop.md`（orchestrator **不讀** `manageCode/server/**`，僅仲裁 dispute）。

#### Slice 軌跡（摘要）
| Slice | Phase | Executor deliverable | Verifier | Fix 輪 |
|---|---|---|---|---|
| 1 | H2 e2e | retention IT + export-failure | L1–L3 pass → **policy_snapshot** gap | fix + 重驗 |
| 2 | Doc align | inner/outer doc 對齊 | L1–L3 pass | — |
| 3 | Live PG | live PG retention IT | L1–L3 pass | — |
| 4 | Worker ops | health + live IT hardening + runbook | L1–L3 pass | — |
| 5 | A′ Sqlite | `SqliteArchiveWriter` | L1–L3 pass | — |
| 6 | B | per-domain warm.backend + dynamic cron | L1–L3 pass | — |
| 7 | C | `ColdArchivePromoter` + scheduler | L1–L3 pass | — |
| 8 | D | `ArchiveQueryRouter` + Admin history | **manifest 未登記** | fix + 重驗 |

#### 仲裁紀錄（2026-07-08 / tiered archive 全線 — 代表性 findings）
| finding | 處置 | 理由 | 後續 |
|---|---|---|---|
| S1 policy_snapshot 未 assert | fix | acceptance-violation — e2e 缺 D4 meta-audit 欄位 | executor 補 assertion；重驗 pass |
| S8 Phase D migration 未入 manifest | fix | acceptance-violation — deploy 會漏跑 RBAC migration | manifest + down SQL；重驗 pass |
| Worker 僅模組存在、測試服未部署 | defer（Slice 4 前）→ 後續 ops 關閉 | 真實 gap；Slice 4 brief 含 live deploy 驗收 | test EC2 worker UP + cron 觸發紀錄 |
| Verifier INFO：verification plan §4.4.5 未同步 Phase D | defer | beyond-acceptance 文件債 | 轉 feature plan 修訂 |
| Orchestrator 早期 session 直接讀 manageCode | defer | 越界信號；採三角色紀律後 **0** 實作 diff | project overlay `plan-delegation-execution-loop.md` 機械化 |

#### 量測欄（8-slice 彙總）
| 指標 | 值 |
|---|---|
| verifier 差集（acceptance-violation） | **2**（S1 policy_snapshot、S8 manifest）— 皆 `fix` 後重驗 pass |
| verifier 降級（只跑 L1？） | **否** — 各 slice 均 L1–L3；live deploy slice 含 L3 對抗性（cron 觸發、health） |
| 仲裁分佈（代表性 violation） | fix **2** / defer **3+** / reject **0** |
| orchestrator 越界（寫 code？回讀 diff？） | **Slice 模型後：無實作 diff**；早期規劃 session 有讀碼，記 defer；仲裁 **未** 為 dispute 外回讀 diff |
| verifier 報告自足性 | **是** — Slice 8 manifest gap 全憑 verifier 報告觸發 fix，orchestrator 未自行 grep manifest |
| Task spawn 成本 | **~18**（8 executor + 8 verifier + 2 fix 輪） |
| 外層 plan commit | **~12**（§12 backfill + brief）；內層 **~12** feature commits |
| 品質信號 | acceptance-violation **2/8 slices**（25%）；皆在 merge 前捕獲；test env deploy 驗證 scheduler 實際觸發 |
| vs 2a-external（單 Step 6） | 協調成本 **↑↑**（8×）；越界寫 code **↓**（slice 紀律後為 0）；verifier catch **↑**（多 phase 邊界） |

#### 契約回饋（寫回 canonical / project overlay）
1. **L2/L3 價值在多 slice 平台計畫中放大** — manifest / meta-audit 欄位類問題 L1 全綠仍漏。
2. **Orchestrator outer-only commit** — 外層 §12 + brief 先 commit 再 spawn，軌跡可審計（C1–C5 合規關閉）。
3. **Task transport = 2a/2a-external 的 agent 化** — 不必人類開新 chat；fresh context 由 subagent 保證。
4. **多 slice 任務值得 loop**；單檔 typo 不值得 — 與 advisory 適用邊界一致。
5. **Schema promotion 仍不建議（證據累積中）** — 2c 支持 doc-only 延續；Phase 3 Q5 決策待後續收斂（見 `_plan.md` §Phase 3）。
### 2d — 外部 monorepo outbound sync Phase 3（4 slices）✅（2026-07-08，Cursor Task transport）

> **專案證據邊界**：inner commit、class 名、live 環境細節留於 `<PROJECT_ROOT>` active main plan §執行紀錄；Ai-skill 只保留 generalized dogfood metrics（依 [`enforcement/sanitization.md`](../../../enforcement/sanitization.md)）。

- **任務**：出站同步平台 Phase 3 — rollback、`sync_failure` 隊列 + reconcile worker、Admin 監控 UI、`WORKER_BATCH` scheduler（4 slices）。
- **Transport**：Cursor orchestrator + **Task subagent** executor / verifier（每 slice 一輪；fix 再 spawn）；**省略 `model`**。
- **Repo**：`<PROJECT_ROOT>` 外層 + `<INNER_REPO>/manageCode` 內層；雙 repo commit。
- **Consumer overlay**：`<PROJECT_ROOT>/.ai-skill/project/rules/plan-delegation-execution-loop.md` — 在 canonical L1–L3 之上擴 **L1–L4 外層驗收 + V1–V4 內層（含 V4 產出物）+ `slice_kind` + C1–C5 合規關閉**。
- **機械 gate**：`check-plan-delegation-orchestrator.py` + Cursor hooks（五事件）+ BDD `gate.plan_delegation_orchestrator`、`gate.consolidation.gherkin_canonical_placement`。

#### Slice 軌跡（摘要）

| Slice | slice_kind | 重點 | 關閉類型 |
|---|---|---|---|
| 3.0-A | implementation | `failure_kind`、補償 delete、內層 IT | `implementation_done` |
| 3.0-A′ | outer_acceptance | Gherkin rollback + L2/L3 外層 | `slice_compliant_closed` |
| 3.0-B | combined | `sync_failure` 隊列 + reconcile worker | `slice_compliant_closed` |
| 3.0-C | combined | sync_status、pending_delete、monitor UI、batch scheduler | `slice_compliant_closed` |

#### 相對 canonical / 2c 的新信號

| # | 觀察 | 對本 plan 的意義 |
|---|---|---|
| 1 | **implementation + outer_acceptance 拆 slice**（A / A′）比單一 combined 好關閉 | consumer overlay `slice_kind` 有效；canonical 可 advisory：user-visible 行為勿僅用 inner JUnit 關閉 |
| 2 | **verification_backfill + deliverables[]** Execute 前填好，Verifier V4 抓 feature / manifest 遺漏 | **Q7 正向**：backfill 像 execution 前 acceptance→evidence 映射 primitive |
| 3 | **機械 gate 生效**：gate 後 orchestrator **零** manageCode 實作 diff | **Q5 支持** Layer 3 consumer gate 夠用、schema promotion 仍不急 |
| 4 | **hook 副作用**：Execute 意圖下 orchestrator 連外層 plan 有時被擋 | 契約缺口 → kit §Cursor 傳輸備註 allowlist（見上） |
| 5 | **combined mega-slice**（3.0-C，8 deliverables）Verifier V4 負載高 | 建議拆 2–3 Verifier 輪；與 2c「多 slice 值得 loop」一致 |
| 6 | **deploy / migration** 由 orchestrator 做，不在 loop 內 | loop 邊界：runtime deploy 應列 brief acceptance 或標 `beyond-loop` |
| 7 | **Gherkin 唯一目錄 gate** 防 `.feature` 再分散 | consumer enforcement 範例；workflow dogfood 清單候選 |

#### 量測欄（4-slice 彙總）

| 指標 | 值 |
|---|---|
| Slices | **4**（2 combined + 1 implementation + 1 outer_acceptance） |
| Task spawn | **~8–12**（每 slice executor + verifier；含 fix 輪） |
| acceptance-violation（merge 前） | **少數** — 多為 deliverable / outer tier 遺漏（V4 抓到） |
| verifier 降級（只跑 L1？） | **初期有**；補強 overlay L1–L4 + V4 後改善 |
| orchestrator 越界寫 manageCode | **0**（gate 生效後） |
| orchestrator 被迫回讀 diff 仲裁 | **罕見**（爭議時定點 Read） |
| 外層 L1–L3 linked 才關 user-visible slice | **是**（C1b 紀律有效） |

#### vs 2c tiered archive

| 維度 | 2c | 2d |
|---|---|---|
| 領域 | 資料歸檔平台 | 同步 / 失敗處理 / Admin 監控 |
| 外層證據 | 以 inner IT 為主 | **L1 Gherkin + L2 BDD + L3 外層 IT** 為關閉條件 |
| slice 模型 | 8 小 slice | 4 slice + **slice_kind 混用** |
| 新 primitive 信號 | 多 slice 值得 loop | **backfill + slice_kind + deploy 邊界** |

#### 契約回饋（寫回 canonical / consumer overlay）

1. **Consumer overlay 模式成立** — canonical 保持 L1–L3 + 三角色；雙 repo + Gherkin 外層用 project overlay 擴 tier，不必一次進 schema。
2. **Execute 意圖 hook allowlist** — 見 kit §Cursor 傳輸備註。
3. **`slice_kind`**（implementation / outer_acceptance / combined）應進模板 A「測試範圍」段（advisory）。
4. **Deploy 不屬 Production leg** — brief 需明示或 defer。
5. **Q5 仍維持 doc-only** — 2d + 2c + consumer 機械 gate = 證據累積中，尚不足以 promote schema。

### 2d′ — ExternalRepoC 9j2 module alignment follow-on（2026-07-09，證據 only）

> **專案證據邊界**：inner commit、class 名、live 環境細節留於 `<PROJECT_ROOT>` active main plan §執行紀錄；Ai-skill 只保留 generalized dogfood metrics（依 [`enforcement/sanitization.md`](../../../enforcement/sanitization.md)）。**本節為 §2d 同一 consumer（ExternalRepoC）的延續 run**，非新 consumer、非 Phase 3 closure。

- **任務**：9j2 出站同步平台 — 模組 **01 app-url** / **02 bookmark** M1–M8 對齊、sync-remote 外層驗收、雙 feature branch 合併至可部署分支、live 刪除語義與測試落點治理。
- **Transport**：Cursor orchestrator + **Task subagent**（Executor / Verifier，多 slice）；部分對齊與外層 acceptance 由 orchestrator 直接寫 `<PROJECT_ROOT>`（hook allowlist 內）。
- **Repo**：`<PROJECT_ROOT>` 外層 + `<INNER_REPO>/manageCode` 內層；雙 repo commit。
- **Consumer overlay**：沿用 2d — `plan-delegation-execution-loop.md`、verification backfill、`slice_kind`、C1–C5；新增 project rules：`pre-push-build-gate`、`test-acceptance-placement`、`9j2-sync-module-alignment`。

#### 相對 §2d 的新信號

| # | 觀察 | 對本 plan 的意義 |
|---|---|---|
| 1 | **平行 feature branch 下 user-visible 已關閉但 UI 不可見**（sync-remote 在未合併分支；可部署分支缺按鈕） | 延伸 2d #1：`slice_compliant_closed` 若含 user-visible 行為，brief 應加 **integration gate**（合併至可部署分支）或標 `beyond-loop: merge`；否則 orchestrator 與使用者驗收脫節 |
| 2 | **遠端已缺席時 delete 非 idempotent**（遠端「未找到」阻本地刪） | L3 / `verifier_only`：`remote_absent_delete` 應進 backfill；mock IT 綠燈仍可能漏 — ERA「哪種證據能關哪個狀態」 |
| 3 | **Live IT teardown 須雙邊**（本地 DB + 遠端清 test-data marker 列） | backfill 應明示 `live_delete_policy` / teardown owner（executor vs orchestrator beyond-loop） |
| 4 | **首輪模組對齊曾跳過 formal Verifier**，使用者糾正後補跑 | 反模式實例：doc/align 回合仍須 V1–V4 或書面 `defer`；僅 executor 自驗 ≠ loop 關閉 |
| 5 | **inner `src/test` commit block 已機械化**（新測試僅 outer `tests/`） | 相對 2g「漸進遷移」：ExternalRepoC 選整路徑 deny + BDD gate；consumer 自理 gate 模式 ×2 家族內變體 |
| 6 | **pre-push build gate**（push 可部署分支前 client build + compile） | Q5：release-time consumer gate 與 orchestrator deny gate 互補；仍 doc-only |
| 7 | **V5 runtime smoke defer**（dev captcha in-memory） | 延伸 2d #6：runtime tier 須 `V5: defer(reason)`，不可默認已驗收 |
| 8 | **merge / push / stack restart 由 orchestrator，不在 Executor loop** | 延伸 deploy 邊界 → **integrate / release leg** 為 beyond-loop，但 user-visible closure 依賴它 |

#### 量測欄（follow-on 彙總）

| 指標 | 值 |
|---|---|
| 模組 slice | **2**（01 + 02 對齊；含外層 acceptance 補強） |
| integration UX fail（合併前） | **1** — sync-remote UI 不可見直至 feature 合併 |
| acceptance-violation（live delete） | **1** — `remote_absent_delete`；fix 後 idempotent |
| verifier 降級（首輪跳過） | **1 次** — 使用者糾正後補 Verifier |
| orchestrator 越界寫 manageCode | **0**（gate 生效後）；idempotent delete fix 經 Executor |
| 新 consumer 機械 gate | **2** — pre-push build、inner src/test block |
| 模型自然落位（Phase 3 穩定性） | **是** — overlay + backfill + ERA 分工未改形狀 |

#### 契約回饋（寫回 canonical / consumer overlay）

1. **`integration gate`** — user-visible `slice_compliant_closed` 應列合併至可部署分支，或標 `beyond-loop: merge`（advisory，模板 A acceptance 段）。
2. **`remote_absent_delete`** — 出站 sync 模組 backfill 建議列 `verifier_only` 負面 case（遠端 404 / 未找到仍允許本地刪）。
3. **`live_delete_policy`** — live tier teardown 雙邊 owner 寫入 backfill，避免 `[sync-test]` 殘留。
4. **Release-time gate** — pre-push build 與 commit-time placement gate 可並存於 consumer overlay，不必進 schema。
5. **Q5 仍維持 doc-only** — 2d′ 強化 consumer gate 證據，不改 schema promotion 判斷。

### 2h — ExternalRepoC common-url Execute 验证不严（2026-07-09，證據 only）

> 全文：[`evidence/2h-externalrepoc-common-url-verification-gaps.md`](evidence/2h-externalrepoc-common-url-verification-gaps.md)

- **摘要**：03 common-url Execute 后用户手验暴露 `sync-remote` static-resource；回溯 F1–F6（RBAC 漏 `admin_role_menus`、V5 仅 `list`、combined 拟 defer L1–L3、verifier 降級、04 user-feedback 产品面事后修正）。
- **量測**：acceptance-violation **≥2**；新 consumer gate **≥2**（RBAC 三连、api-surface smoke）；stakeholder 纠偏 **≥2**。
- **契约回饋**：`V5-api-surface`、`RBAC-triple`、`combined-no-inner-close`、`restart-aware-runtime`、`spec-before-execute`（见 evidence §契约回饋）。

### 2i — ExternalRepoC user-feedback S0–S4 Execute（2026-07-09，證據 only）

> 全文：[`evidence/2i-externalrepoc-user-feedback-pull-execute.md`](evidence/2i-externalrepoc-user-feedback-pull-execute.md)

- **摘要**：04 user-feedback **S0→S4 完整 Execute**；用户 Stop 中断 Verifier → orchestrator resume；**2h 教训迁移成功**（RBAC gate、api-surface、L1–L3 不 defer）；新信号：inventory 发现 mapper 指错表、`sync_jobs` 与 `sync_failure` 分表、verifier `resource_exhausted` fallback。
- **量測**：slice **5**；Executor spawn **5**；verifier 降級 **1**；stakeholder 纠偏（Execute 前）**≥4**；orchestrator 写 manageCode **0**。
- **契约回饋**：`resume-after-stop`、`verifier-fallback`、`s0-inventory-gate`、`platform-queue-typing`、`mapping-defer-pattern`、`2h-lessons-transfer`（见 evidence §契约回饋）。
