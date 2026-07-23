# Delegated Execution（委派執行 loop 的 software-delivery 落地）

> **Cognitive Slice**：`sd-delegated-execution`（maturity: **candidate** — 泛化自一個外部 consumer（ExternalRepoC）多 slice 真實 run + Ai-skill 內部 dogfood；第二個獨立 consumer 證據前不升 enforcement、不接 runtime）。

| slice 欄位 | 值 |
|---|---|
| `id` | `sd-delegated-execution` |
| `purpose` | 已宣告 delegation 的交付任務進入執行時，定義**誰該做什麼**：角色 × 證據責任矩陣、verification backfill、deliverables 清單、slice 關閉狀態與 Verifier 驗證動作 |
| `type` | `execution` |
| `tags` | delegation, execution-loop, verification-backfill, role-responsibility |
| `load_when` | 交付任務已宣告 `delegation.enabled: true` 進入執行；orchestrator 需拆 slice / 寫 brief / 判定合規關閉 |
| `do_not_load_when` | surgical 小修、單 session 無委派（直接 implementation + validation）；純 intake / 純 contract 階段 |
| `owner_layer` | workflow |
| `canonical_source` | Loop 契約：[`plans/README.md`](../../plans/README.md) §Delegation loop SOP（本檔不複製 verifier 4 欄位 / 仲裁三處置正文）；本檔 canonical 範圍 = delivery 域的角色責任 / backfill / 關閉狀態 |
| `dependencies` | `sd-test-strategy`（BDD closure / 證據語意）、`sd-validation`（completion evidence）、`sd-implementation`（execution mode）、review capability invoke（[`README.md`](README.md) §Review invoke） |

## 1. 何時走 loop（advisory）

| 情境 | 建議 |
|---|---|
| 使用者表達執行意圖（「開始執行 plan / sub-plan / slice / 幫我做 Step N」） | 進入 orchestrator 模式 — **mandatory 三角色 loop**（brief → executor → verifier → 仲裁） |
| sub-plan / 任務宣告 `delegation.enabled: true` | 走 loop（brief → executor → verifier → 仲裁） |
| **loop 未關閉**（backfill 仍有未關閉行、仲裁未完） | **維持** orchestrator 模式——不因跨 turn / 跨 session 而失效 |
| 跨 session、可驗收 slice、主 session 需保持規劃 / 仲裁位 | 走 loop |
| surgical 小修、單 session、無 delegation | 不走；直接 implementation + validation |
| 純問答 / 只讀審計（不進入 Execute） | 不觸發 |
| **Transport adaptation**：使用者明確要主 session 當 executor | 允許，但**必須在 plan 註明**（角色降格是記錄在案的例外，不是默認滑移） |
| **Spawn／bootstrap 機械 gate 卡住 Task**（dogfood **3a**） | fresh Executor／Verifier 被 `gate.bootstrap.receipt_present`／primary_source **deny** 非-Read 工具 → 看起來像「三角色壞了」→ **不得**取消 loop；書面 transport fallback；同 session 驗證最高 `implementation_done` |

## 1b. Orchestrator 執行順序（不可跳過）

1. 讀 active plan 執行段 + 本 slice（**不讀實作源碼**）。
2. 補 `delegation.brief`（含 slice 類型）+ verification backfill（每條 acceptance → tier + owner，§3）。
3. **Commit plan 變更**（無 git 錨點不得派發）。
4. 派發 **Executor**（fresh session / agent；`context.required` 由 executor 自己讀）。
5. 派發 **Verifier**（另一個 fresh context；V1–V4，§5）。
6. 仲裁 fix / defer / reject → `fix` 再派 executor → **重新驗證** → 關閉狀態 + C1–C5 寫入 plan 執行紀錄 → **commit** → 才開下一 slice。

## 2. 角色 × 證據責任矩陣（誰該做什麼）

四責任閉環（Specification → Production → Evidence → Decision）在 delivery 域的分工。**每條 backfill 行必標 `owner`**：

| 歸屬 | 產出責任 | 不可做 |
|---|---|---|
| `orchestrator`（主 session） | plan / slice 拆分、brief（acceptance + verification + `context.required` + `deliverables`）、**執行前的 acceptance-spec 骨架**（人類可讀 scenario + spec-alignment gate 骨架）、仲裁表、合規關閉宣告 | 不寫 implementation diff；不為探路讀實作源碼；不代 executor 讀 `context.required` |
| `executor`（獨立 fresh session / agent） | 依 brief 實作 + **implementation 層證據**（單元 / 內層整合測試）+ brief 指定的整合測試與非測試產出物（doc / migration / config） | 不改 acceptance / scope；不自判 Done；資訊不足停下回報，不猜 |
| `verifier_only` | 負面 / 邊界 / 架構禁止項測試——**executor 自寫測試之外**的對抗性補充 | — |
| `verifier`（另一個 fresh session） | 驗證動作 V1–V4（§5）+ 逐條 acceptance 結論；只產證據 | 不修檔、不提修法、不下判決、不兼任 executor session |

## 3. Verification backfill（執行前必設計的「驗收 → 證據」映射）

Plan 從規劃進入執行前，orchestrator 在被委派任務的 plan artifact 內建立 backfill 表：**每條 acceptance → 至少一行證據映射，每行必標 `tier` + `owner` + 狀態（`linked` / `deferred` / `pending`）**。

證據 tier 語意引用既有 slices（**不另建 taxonomy**）：

| tier | 對應既有概念 | 關閉權威 |
|---|---|---|
| acceptance-spec | [`test-strategy.md`](test-strategy.md) Journey Specification / BDD scenario | user-visible acceptance 的**驗收權威** |
| spec-alignment | 規格 ↔ 代碼 ↔ 測試路徑的機械對齊 gate | 同上（輔助鎖） |
| integration | [`validation.md`](validation.md) 整合 / journey 證據 | 同上 |
| live | [`validation.md`](validation.md) live system proof（brief 宣告時 mandatory，不准 skip） | 按 brief |
| inner | implementation 層單元 / mock 測試 | **實現證據**——不得單獨關閉 user-visible slice |
| deliverable | 非測試產出物（doc / migration / config / 流程錨點） | Verifier V4 核對 |
| runtime（candidate，consumer 2026-07-09） | 真實部署路徑證據：migration 已執行、服務可起、主路徑 API/UI 可走——介於 integration 與 live 之間 | Verifier V5；含 schema/API/UI 交付面時 mandatory，缺 → `runtime-omission` 不得合規關閉 |

**回填鐵則**：

1. 每條 **user-visible** acceptance → 至少一行 acceptance-spec + 一行 spec-alignment 或 integration（或整行 `deferred` + follow-up slice id）。
2. 每條非測試產出物 → `tier=deliverable` 或 brief `deliverables[]`（見 §4）。
3. `tier=inner` alone 不得標 `linked` 並關閉 user-visible slice。
4. 合規關閉前：無裸 `pending`；`deferred` 必附 follow-up。
5. **Loop 外動作明示**：runtime deploy / migration / 環境操作若不在 executor 的 Production leg 內，必須列入 acceptance（由 executor 交付）**或明標 `beyond-loop`**（orchestrator 執行並記錄於 plan）——不得隱含（consumer 2d 證據：deploy 隱含在 loop 外導致邊界模糊）。

## 4. Brief `deliverables[]`（防「測試綠但東西沒交」）

Orchestrator 在 brief 中與 acceptance 並列宣告非測試產出物，執行前隨 plan commit：

```yaml
deliverables:
  - id: D1
    artifact: <path-or-process-anchor>
    owner: executor | orchestrator
    check: <一句可核對的完成判準>
```

## 5. Verifier 驗證動作（V1–V4）

| 動作 | 內容 | 對應 canonical 三層 |
|---|---|---|
| V1 重跑 | 獨立執行 brief `verification`（不信 executor 自驗輸出） | L1 replay |
| V2 讀碼 | diff 對照 acceptance 與架構禁止項 | L2 inspection |
| V3 對抗 | 執行或補寫 `verifier_only` 負面 / 邊界 case；反例來源可用 evidence producer 機械枚舉（見下方 V3 evidence producer） | L3 adversarial |
| V4 產出物與流程 | **對照 brief `deliverables` + backfill 逐條核對**：該存在的檔案 / 步驟是否交付；必須產出核對表，不能只看測試綠 | （delivery 域擴充） |
| V5 運行態 smoke（candidate，consumer 2026-07-09 incident） | **真實部署路徑是否可行**：schema/migration 已 apply（非僅檔案存在）、服務可起、使用者主路徑可走（API/UI 依賴鏈）。backfill 有 `runtime` tier 行時 **mandatory 不可跳**；纯 refactor 無 schema/API/UI 變更可書面 defer | （delivery 域擴充） |

**V1 與 V4 的區別**：V1 問「測試過沒過」；V4 問「**該交付的東西有沒有交齊**」。**V4 與 V5 的區別**：V4 問「migration **檔案**在不在」；V5 問「migration **跑了沒**」——build 綠 + 測試綠 + 檔案齊 仍可能 runtime 打不開（consumer 實證：mock IT/build/grep 全綠、dev DB 未跑 migration → UI 開啟即 missing column）。

**V3 evidence producer（使用者裁決 2026-07-09）**：V3 的反例不必只靠 verifier 自行想像（imagination-driven）。對 boundary / boolean / nullability / authorization / invariant / guard 類風險，targeted mutation（[`test-strategy.md`](test-strategy.md) §Mutation Testing / Test Effectiveness Check）是 V3 的一種 **evidence producer**——機械枚舉行為區分點（mechanical falsification），**不是新的驗證層**（不設 V6）。使用契約：

1. **Survived mutant 本身只是資訊，不是 finding**。Verifier 必須轉譯為 semantic-gap finding：`evidence` = mutation + behavioral implication（如「`>` → `>=` survived ⇒ boundary `price==100` 未被測試區分」）+ 建議的 `verifier_only` case；`acceptance_ref` / `classification` 照 canonical 契約。Orchestrator 只讀 finding，不需知道 mutation engine 存在。
2. **不得以 mutation score 作 KPI 或關閉門檻**；通過標準沿用 test-strategy：殺掉代表真實風險的 mutant、過濾 equivalent mutants（equivalent mutant 對應仲裁 `reject` + `refuted` 留證）。
3. Producer 可替換：mutation 只是「Behavioral Falsification」producer family（mutation / fault injection / property-based / model-based）之一，皆產出同型 evidence——「此行為未被驗證區分」。family 通用化 gated on plan Q9（forming abstraction，observe-only）；未 graduate 前本 slice 只承載 targeted mutation 這一種。

**Delivery 域 finding 分類擴充**（candidate；映射回 canonical `classification` family，第二 consumer 證據前不進 canonical enum）：

| 擴充值 | 語意 | canonical 映射 |
|---|---|---|
| `deliverable-omission` | brief / backfill 要求的產出物缺交 | `acceptance-violation` family（deliverables 已宣告時） |
| `process-omission` | 宣告的流程步驟漏做（branch、commit 錨點、交付表欄位） | `observation` / `acceptance-violation` 依 brief 宣告而定 |
| `runtime-omission`（consumer 2026-07-09） | runtime tier 行未驗或部署路徑不可行（migration 未跑 / 服務未起） | `acceptance-violation` family（runtime tier 已宣告時） |

## 6. Slice 關閉狀態（合規判定）

**Slice 類型**（orchestrator 在 brief 標明；consumer 2d 證據：implementation + outer_acceptance 拆開比單一 combined 好關閉）：

| slice 類型 | 內容 | 合規關閉要求 |
|---|---|---|
| `implementation` | 僅實作 + inner 證據 | inner `linked`；user-visible acceptance 可 `deferred`（必附 follow-up slice id） |
| `outer_acceptance` | 補 acceptance-spec / spec-alignment / integration 外層鏈 | 外層 tier 全 `linked` |
| `combined`（user-visible 行為預設） | 實作 + 外層驗收同批 | 外層 tier `linked`；inner 為輔助證據。**單輪 deliverables 過多（mega-slice）→ 拆多輪 verifier 或拆 slice** |

| 狀態 | 條件 | 語意 |
|---|---|---|
| `implementation_done` | inner 證據 linked；user-visible acceptance `deferred` 且有 follow-up slice id | 內層交付完成，**不等同**合規關閉 |
| `slice_compliant_closed` | 下表 C1–C5 全部成立 | 合規關閉，可開下一 slice |

| # | 條件 |
|---|---|
| C1 | brief 每條 acceptance 在 backfill 為 `linked`（或已記錄的 `deferred` + follow-up） |
| C1b | user-visible acceptance 至少一行 acceptance-spec / spec-alignment / integration tier `linked`（不可僅 inner） |
| C2 | Verifier 已跑 V1–V4（+ V5 於 runtime tier 適用時）；無未仲裁的 `acceptance-violation` / `deliverable-omission` / `runtime-omission` |
| C3 | 仲裁 `fix` 項：re-delegate 修復 → **重新驗證** → pass |
| C4 | 無裸 `pending`；Verifier 未降級為「僅重跑測試、跳過 V2/V4」 |
| C5 | Orchestrator 書面確認關閉狀態並寫入 plan 執行紀錄 |

## 7. Orchestrator 錨定紀律

1. **先 commit plan 變更，再派發**：brief / backfill / 仲裁表寫完 → commit → 才 spawn 下一個 executor 或開下一 slice（避免 plan 紀錄與工作區脫節、合規關閉無 git 錨點）。**多 slice（2p）**：優先用 plan 正文 **slice 累積表**追加列，當前 `frontmatter.brief` 只保留本 slice 最小集，降低整份覆寫造成的外層 commit 密度。
2. **條件式讀檔**：常態只讀 plan / brief / backfill / 交付表 / findings；僅仲裁爭議時按 verifier 引用的具體路徑+行號**定點讀**，不掃目錄、不回讀整份 diff。**Verifier 報告缺四欄表時**：先要求補表，再仲裁（勿默認散文可過）。
3. Orchestrator 與 executor 各 commit 自己的層（plan artifact vs 實作 repo）；orchestrator 不代 executor commit 實作。
4. **多 todo / 一口氣（2p）**：使用者「Don't stop until all todos」= **多輪** brief→E→V，不是單 Task 包辦。

**主 session 禁止（loop 內）**：

1. 為寫碼而 Read / Grep / Edit 實作源碼（brief 裡列給 executor 的路徑：只記路徑，不打開讀）。
2. 未派發 executor 就直接寫實作。
3. 未 commit plan 變更就派發。
4. 同一個 session 兼任 verifier（必須新開 fresh context）。

**被 gate 擋下時**：不要 retry 直改實作；補 brief → commit plan → 派發 executor。

## 8. Anti-patterns（review lens）

| failure | 徵象 | recovery |
|---|---|---|
| 自證循環 | executor 寫實作 + 自寫測試，verifier 只重跑同一批測試 | verifier 必跑 V2–V4；`verifier_only` 對抗補充 |
| 測試綠即關閉 | deliverables / 流程步驟缺交仍宣告 pass | V4 核對表 mandatory（C2/C4） |
| inner-only 關閉 | user-visible slice 只憑單元測試關閉 | C1b block；改 `implementation_done` + follow-up |
| orchestrator 探路讀碼 | 「為了寫 brief」大量讀實作源碼 | brief 只記路徑給 executor 讀；越界記入量測欄 |
| verifier 兼任 | 同一 session 先執行後驗證 | fresh-context invariant（canonical SOP invariant 1） |
| **single-agent skip verifier**（consumer 2j 2026-07-10） | orchestrator 只 spawn **一个** Task 包办 implement + 自验 mvn；**0** Verifier session；用 `delegation.enabled: false` 当不 loop 理由 | Execute 意图 → loop mandatory；consumer **verifier-after-executor** gate；须标 `implementation_done` 非 `slice_compliant_closed` |
| **spawn-friction / gate-misattribution**（dogfood **3a** 2026-07-23） | Cursor Task spawn 失敗或貴 → 主 session 直做；或把 Ai-skill **candidate／bootstrap** 說成「擋三角色」 | spawn 摩擦 ≠ 豁免；candidate ≠ deny-loop；書面 transport fallback；同 session 驗證 ≤ `implementation_done`；見 `evidence/3a-…` |
| **mechanical-bootstrap × Task cold-start**（dogfood **3a** 主軸） | `preToolUse` 對 **尚未 Receipt** 的 fresh Task deny 寫入／Shell → orchestrator 放棄 E／V | 區分 bootstrap gate vs 三角色專用 deny；Task brief 強制冷啟動 bootstrap；失敗走 transport fallback（consumer §2.1）；與 2j「該擋略過卻沒擋」對偶 |
| **batch-todo single Task**（dogfood 2p 对照） | 用户「一口气 / all todos」被解读为一次 Task 做完多 slice | **多 todo = 多轮** brief→E→V；禁止合并 |
| **transport inner-only path proof**（dogfood 2q） | Verifier 只重跑 mock/unit，默认已切正式 API 却从未 V5-A/sync-remote；对外说「功能通」 | backfill 强制 `tier=runtime` + features↔L3；无密钥 defer follow-up；consumer `gate.plan_transport_runtime_evidence` |
| **SPA unit-only / stale serve**（dogfood 2z） | Karma／TestBed 绿即关 SPA slice；或长跑 `ng serve` watcher 失效仍手验旧 bundle；HMR 被误当验收 | C1b 要求 browser e2e／runtime UI；V1 独立重跑 Playwright；V5-U：冷启动或确认 watcher 后再验；HMR ≠ evidence |
| combined mega-slice | 單一 slice 塞過多 deliverables，verifier V4 單輪負載過高、核對品質下降 | 拆 2–3 輪 verifier 或拆 slice（§6） |
| deploy 隱含在 loop 外 | runtime deploy / migration 無人宣告歸屬，orchestrator 默默代做 | backfill 鐵則 5：列 acceptance 或明標 `beyond-loop` |
| 綠燈假象（missing constraint） | build 綠 + 測試綠 + 檔案齊，但 runtime 打不開——既有證據只能排除 implementation failure，**無任何證據約束 runtime 可行性**，Close 在過大的可行集裡做出 | backfill 補 `runtime` tier 行 + Verifier V5；未來勿再往 V6/V7 疊 checklist，先問「缺的是哪一種 constraint」 |
| mutation score KPI | 以 mutation % 作為品質指標或關閉門檻；跑全量 engine 追分數 | targeted mutation only（risk-triggered）；survived mutant → semantic-gap finding；殺掉代表真實風險的 mutant 即可（§5 V3 evidence producer） |

## 9. 機械 gate 模式（Layer 3 adapter，optional）

Role boundary 靠行為維持不住時，consumer 可在**工具層**落機械 gate（不動本 slice、不動 delegation schema）。已驗證的事件生命週期模式：

```text
執行意圖偵測（prompt / active plan 宣告）→ arm orchestrator lock
  → orchestrator lock 生效期間：deny 主 session 對實作路徑的寫入與探路讀取
  → subagent（executor）啟動 → grant executor lock（可寫實作）
  → subagent 結束 / session 結束 → 清除 locks
```

- **Deny 範圍只限實作路徑**：orchestrator 自有 artifact 路徑（plan / overlay / 外層測試骨架）**必須 allowlist**——否則 orchestrator 連 plan patch 都被擋，被迫讓 executor 代寫外層 artifact = 角色反轉（consumer 2d 契約回饋）。
- **Verifier spawn 追踪（consumer 2j 2026-07-10）**：除挡 orchestrator 写实作外，consumer 可追踪 `executor` subagent 完成後是否 spawn **独立 Verifier**——未完成 Verifier 就 spawn 下一 Executor 或宣告 slice closed → deny Task spawn（`verifier_required_before_next_executor`）。
- **Execute 意图 > `delegation.enabled: false`**：用户 Execute 意图触发 loop；frontmatter `false` **不是**豁免理由——Execute 前须改 `true` 并补 backfill（2j 负向证据）。
- **Bypass 必須書面記錄**（plan 或使用者明示），對應 §1 的 surgical / transport-adaptation 例外。
- 實作細節（hook 事件名、腳本、state 檔）屬 consumer 工具層；已驗證實例：外部 consumer（ExternalRepoC）editor-hook 五事件 gate + BDD 10/10（2026-07-08）。

## Provenance / 升級條件

- 來源：外部 consumer（ExternalRepoC）多 slice 真實 run（2026-07-08，含 consumer 層機械 gate）+ Ai-skill 內部 dogfood；決策紀錄見 [`plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md`](../../plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md)。
- 維持 doc-only / advisory；**不新增 enforcement rule、不接 runtime、不動 delegation schema**。
- 升級條件：第二個獨立 consumer 的真實使用證據（分類擴充進 canonical enum、backfill 模板化、機械 gate 泛化各自獨立評估，一次一階）。
