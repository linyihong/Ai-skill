# 2e — 跨域 run：grandfather sunset audit（Research/Audit 域）✅（2026-07-08–09，Claude Code Agent transport）

> **還原註記（2026-07-10）**：本 run 紀錄原載於 kit §2e，於併發 plan 回寫（b6481e5）中遺失，自 commit 66f58ed 還原並依 evidence/ 慣例落檔。


> **Stage 2 裁決 run（Q6 / Q7(b) / Q8）**。任務真實性：`governance/lifecycle/system-upgrade-governance.yaml` §`pre_2026_05_28_doc_only_completion` 的 post_sunset_evaluation_rule 要求 4 個 covered plans 在 **2026-08-31** 前完成 wire-or-downgrade 處置——本 run 產出其調查與建議前置，是 deadline 驅動的自然待辦（非 manufactured）。**誠實標記**：run 由使用者發起（deliberate stage-2 probe），非全然偶發；「自然性」判準落在 (a) 任務本身真實、(b) 不把 sd 詞彙硬套進 brief——brief 以 Research 域原生語彙撰寫（Question → Investigation → Evidence → Recommendation），無 slice_kind / backfill / deliverables 欄位。
>
> **候選盤點紀錄（Q6/Q8 佐證）**：economics 計畫（durability observation，禁主動推）、pattern-library T3A（authority surface decision 明文 do-not-pre-decide）、interaction-hazard（evidence-accumulation-only invariant）皆不可作為 run 載體——所有非 delivery active plans 均處 evidence-gated 暫停，佐證「manufacture 不可行、只能等真實任務」的紀律成立。

#### Brief（orchestrator 撰寫，Research 域原生）
- **問題（question）**：4 個 pre-2026-strengthened doc-only plans 的 orphan surfaces 現況為何？sunset 前應各自採取什麼處置？
- **驗收（acceptance）**：(1) 每 plan 一節：宣告 orphan_surfaces 的現況事實（route 存在？誰消費？附檔案+行號或指令輸出，不可憑印象）；(2) 對照 post_sunset_evaluation_rule + 各 wire_plan 提示判定 wired / orphan；(3) 每 plan 處置建議（升 auto-detected / 補 wire 含最小動作 / 降 orphan 含移除清單 / manual_activation 註記）+ 理由；(4) conditional_extension_trigger 兩條件現況核對；(5) 不確定處明標 unverified。
- **自查（verification）**：`ai-skill runtime audit` 取得分類現況；grep registry / discovery yaml / hooks.go 覆核每條 consumer 主張。
- **產出**：`02-grandfather-sunset-audit.md`（本資料夾 companion，調查報告，處置決定保留給 maintainer）。
- **結果**：✅ loop 完整走完（2026-07-09）。調查者（fresh agent、worktree）產 252 行報告（commit `c8ff035`）：**5/5 surfaces 已 wired**（flag 條款與補 wire 同日 2026-05-28 落地、條款從未回頭更新——過時宣告）；延展條件兩者皆不成立；sunset 只剩行政收尾。中途遇 session limit 中斷 → resume 續跑完成交付（transport resilience 實證：deliverable 未損）。事實查核者（另一 fresh agent、唯讀）：全部引文逐條命中、5 surfaces 獨立重跑 audit 一致、無虛引無「提到當消費」；findings ×2 皆 observation 級。

#### 仲裁紀錄（2026-07-09 / grandfather sunset audit）
| finding | 處置 | 理由 | 後續 |
|---|---|---|---|
| F1 manual_activation 註記尾端行號差 1（L2581→實為 L2580） | defer | 語意主張不受影響；查核表已記正確值 | 隨 sunset 行政收尾批次修 |
| F2 related_scenario 檔案實際存在（報告標 unverified 未讀） | defer | 非矛盾（報告只稱未讀）；補充事實供 maintainer | 併入 sunset 收尾檢視 |

#### 量測欄
| 指標 | 值 |
|---|---|
| verifier 差集 | 2（皆 observation；**0 violation**——調查報告品質經對抗性查核成立） |
| 仲裁分佈 | fix 0 / defer 2 / reject 0 |
| orchestrator 越界 | 無實作 diff；仲裁全憑查核報告。Borderline 記錄：中斷恢復時做過報告**章節骨架** grep（recovery 決策用，未讀正文） |
| verifier 報告自足性 | 是（Q1 第 3 個正向） |
| 調查者讀檔差集 | 有：gen3 audit plan + 工具 git history（延展條件的直接證據來源，**正當**）→ 契約回饋：此類 audit brief 的 context.required 應含「宣告來源 plan」 |
| 契約缺漏回饋 | brief v1 未列 gen3 plan；research 域 brief 教訓：調查對象的「宣告寫入時點的來源 plan」是必要 context |

#### Q6 / Q7 / Q8 跨域觀察（本 run 的核心產出）
| 問題 | 觀察 |
|---|---|
| **Q6 四責任閉環** | **自然成立**：Specification（調查 brief）→ Production（調查報告）→ Evidence（查核 findings）→ Decision（仲裁 defer×2）→ 回饋（brief 教訓）。角色 topology 自然不同（調查者/查核者/orchestrator + **maintainer 治理處置為第二層 decision**——Research 域特有的 decision 兩級分層，sd 域無此明顯分層）。**Pattern held, topology differed** = Q6 預測的正確形態 |
| **Q7 backfill 出現/缺席** | **結構化 backfill（tier+owner 映射表）明確缺席且不需要**；出現的是**弱形式**——acceptance 條目內嵌證據要求（「附檔案+行號」「不可憑印象」「明標 unverified」）。判讀：Verification Backfill 是 sd 域強形式；跨域共同候選是 **evidence-first acceptance**（acceptance 自帶證據標準），非 backfill artifact 本身 |
| **Q8 證據責任結構** | **四問全部自然重現**：誰產證據（調查者=主張+引文 / 查核者=獨立重取+對抗查漏）；哪種證據關哪狀態（引文全命中+獨立重跑一致 → 報告可信）；哪種不能單獨關（**調查者自查 necessary but insufficient——同構於「inner 證據不能單獨關 user-visible slice」：自產證據不能自我關閉**）；誰依證據決策（orchestrator 仲裁 findings；治理處置另屬 maintainer）。同構度高 + 一個域 variant（decision 兩級） |

