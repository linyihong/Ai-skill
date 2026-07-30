# Phase 0 — Architecture Mapping（Knowledge Governance Engine）

**Run ID**: `p0-map`  
**Date**: 2026-07-30  
**Goal**: 驗證 Architecture Hypothesis — **不是**盤點功能清單。  
**Scope**: Commit Message / README-or-doc-sync / Linked-updates proxies；**不含** Plans engine 真搬遷。  
**Method**: 讀 `runtime/core-bootstrap.yaml` `per_commit_obligations` + `commitMsgValidatorRegistry` + pre-commit gates；對每類 rule 標 Context / Projection / Capability / Adapter。

---

## 0. 既有「近 Context」事實（重要）

今日 commit-msg 路徑已有類 Context 物件：

```text
commitMsgCtx { text, staged, root, modes }
```

加上 `commitMsgValidatorRegistry`（obligation id → func）。

| 觀察 | 對 KGE 的含義 |
| --- | --- |
| Context 形狀已存在 | Context Builder 不是從零發明；是 **正規化／去 Git 耦** |
| Registry 已存在但綁 Git hook | Capability Registry 要把「obligation id → Go func」升成 **Rule → Capability → Implementation**，且可被 CLI adapter 呼叫 |
| 多數 validator 簽名已是 `(text, staged, root)` 或 `(modes, …)` | 與 `Validate(Context)→Findings` 同構 |

→ Architecture 不是空想；是在 **抽離 Adapter**。

---

## 1. Capability 候選（收斂用工作集）

> Phase 0 假設：**有限集合**。若映射後不斷爆炸 → 否證「Capability 可收斂」。

| Capability ID | 提供什麼 | 典型消費者 |
| --- | --- | --- |
| `cap.commit_msg` | 完整 commit message 字串 | cognitive presence、token budget 文字部分、opt-out trailers |
| `cap.modes` | 從 msg 解析的 cognitive modes map | cost / floors / governance consistency / adaptive |
| `cap.staged_paths` | staged 相對路徑列表 | memory subdir、plan status、多數 co-change gates |
| `cap.staged_content` | 依 path 讀取檔案內容（root-relative） | bootstrap thinness、sanitization、ownership drift |
| `cap.staged_diff` | cached diff 文字／hunk 掃描 | cli-doc-sync（新 case / 新 hook 偵測） |
| `cap.repo_fs` | 路徑是否存在於 working tree（非 diff） | markdown-yaml sibling exists? |
| `cap.path_cochange` | 「若 A 在 staged，B 也應在 staged」約束 | cli-doc-sync、md↔yaml sync（**輕量 linked-update proxy**） |

**刻意不放進 Phase 0 capability 核心**：

| 候選 | 為何延後 |
| --- | --- |
| `cap.plan_tree` | Plans domain — Later pack |
| `cap.runtime_db_query` | 與 `runtime/` session engine 交界；避免 KGE 吞噬 phase machine |
| `cap.full_linked_updates_graph` | `enforcement/linked-updates.md` 大多仍 **behavioral**；機械面只有 proxy |

**初步收斂判定**：首批三域大約落在 **7 個 cap** 內；多 rule 共享同一 cap → Registry 合理（見五問 C/D）。

---

## 2. Validator → Architecture 映射表（焦點三域 + 對照）

### 2.1 Commit Message 域（Rule A 池）

| Validator / Obligation | Kind | Context 需求 | Projection? | Dependency? | 主要 Capabilities | Adapter 透明？ |
| --- | --- | --- | --- | --- | --- | --- |
| `cognitive_mode_block`（gate） | validation | msg | **No** | No | `commit_msg` | Yes — CLI 也可掃一段文字 |
| `cognitive_cost` | validation | modes | **No** | No | `modes` | Yes |
| `token_budget` | validation | modes + msg | **No** | No | `modes`, `commit_msg` | Yes |
| `activation_signals` | validation | modes + msg | **No** | No | `modes`, `commit_msg` | Yes |
| `capability_snippet` | validation | modes + msg | **No** | No | `modes`, `commit_msg` | Yes |
| `adaptive_triggers` | validation | modes + msg | **No** | No | `modes`, `commit_msg` | Yes |
| `execution_mode_floors` | validation | modes + staged | **No** | No | `modes`, `staged_paths` | Yes（CLI: 傳 path 列表） |
| `governance_mode_consistency` | validation | modes + staged + msg | **No** | No | `modes`, `staged_paths`, `commit_msg` | Yes |
| `memory_mode_subdir` | validation | modes + staged | **No** | No | `modes`, `staged_paths` | Yes |

**域結論**：幾乎全部 **無 Projection**。偶發需要 `staged_paths` 做 floor 檢查，仍非 Graph。

### 2.2 README / Doc-sync 域（含 CLI contract、md↔yaml、bootstrap entry）

| Validator | Kind | Context 需求 | Projection? | Dependency? | Capabilities | Adapter 透明？ |
| --- | --- | --- | --- | --- | --- | --- |
| `cli_doc_sync` | validation | staged + cached diff + msg opt-out | **No**（path 約束即可） | **輕量** path_cochange（Go↔command-contract.md） | `staged_paths`, `staged_diff`, `commit_msg`, `path_cochange` | **Yes if** diff 由 Adapter 注入 Context，不在 Rule 內叫 `git` |
| `markdown_yaml_sync` | validation | staged + sibling exists on disk | **No** | **輕量** companion pair | `staged_paths`, `repo_fs`, `path_cochange` | Yes — Rule 不需知 Git |
| `bootstrap_entry_thinness` | validation | staged + file content | **No** | No | `staged_paths`, `staged_content`, `commit_msg` | Yes |

**域結論**：這就是計畫中的「README Sync」類 — **paths + content / co-change**；**不需要** plan_tree 或全庫 dependency graph。

> 註：repo 內未必有名為 `readme-sync` 的單一函式；**cli_doc_sync / markdown_yaml_sync / bootstrap_entry_thinness** 是同族「文件必須跟著 source」的機械實例。

### 2.3 Linked-updates 域（機械 proxy vs behavioral）

| 層 | 狀態 | KGE 含義 |
| --- | --- | --- |
| `enforcement/linked-updates.md` 全文 | 大多 **behavioral** | 不可妄想 Phase 0 一次機械化 |
| `markdown_yaml_sync` / `cli_doc_sync` | **機械 proxy** | 適合當 Linked-updates **子集** Rule |
| `glossary_retro_own` | 機械（framework 詞彙） | 偏 glossary co-change |
| `review_architecture_doc_drift` / `canonical_ownership_drift` | pre-commit 內容掃描 | `staged_paths` + `staged_content`；仍非全 graph |
| `sanitization_scan` | pre-commit content | `staged_content` |

**域結論**：Linked-updates 進 KGE 應標成 **「co-change / content policy rules」**，不是「載入完整 linked-updates graph」。Dep Provider **optional** 且多數時候用 `path_cochange` 就夠。

### 2.4 Explicitly deferred（對照用，不當 Phase 0 migration）

| Validator 族 | 為何 Later |
| --- | --- |
| plan_status / checkbox / archival / plan_tree_* / plan_evidence | Plans domain 最重；混進 Phase 0 會分不清架構 vs domain |
| runtime_yaml_projects / runtime_trigger_wiring / runtime_index_freshness | 緊貼 `runtime/`；先釐清 KGE vs phase machine 再碰 |
| enforcement_registry_* | Layer 2.5 自我治理；應 **對接** Capability Registry，不當第一個業務 plugin |

---

## 3. 五問判定

| # | 問題 | 判定（Phase 0 中期） | 證據摘要 |
| --- | --- | --- | --- |
| A | Context 是否足夠？ | **正向**（工作假設成立） | 大多數焦點 rule 可由 `commitMsgCtx` 超集覆蓋；真正缺口是 **把 `git diff` 從 Rule 挪到 Adapter→Context**（cli-doc-sync） |
| B | Projection 是否真 Optional？ | **強烈正向** | Commit-msg 域 0% 需要 Projection；doc-sync / linked proxy 也不需 typed projection — 路徑約束即可。真正需要 projection 的多在 **Plans / runtime**（Later） |
| C | Capability 是否收斂？ | **正向** | 首批映射收在 ~7 caps；未看到「一 rule 一 cap」爆炸 |
| D | Registry 是否合理（共享）？ | **正向** | 整簇 cognitive validators 共享 `modes`+`commit_msg`；doc-sync 共享 `staged_paths`+`path_cochange` |
| E | Adapter 是否透明？ | **條件正向** | 概念上：Rule 只收 Context。**現況缺口**：`validateCLIDocSync` 內部直接 `exec.Command("git", …)` — Mini Spike 必須證明可改為「Adapter 預填 `staged_diff`」。其餘多數已接近透明 |

**總結**：Architecture Hypothesis **尚未否證**。進 Phase 1 的架構風險偏低；剩餘工程是 **抽 Git 出 Rule** + **正式 Capability 表**，不是重畫邊界。

---

## 4. Mini Spike（紙上 walkthrough）— 納入 Phase 0

### Rule A — 無 Projection

**選**：`obligation.commit.cognitive_cost`（`validateCognitiveCost`）

```text
Adapter(commit-msg | CLI --msg-file)
  → InputSnapshot{ commit_msg }
  → ContextBuilder → Context{ capabilities: modes? }  
       （modes 可由同一 Adapter 從 msg 解析後放入）
  → Capability Registry: rule.cognitive_cost requires [cap.modes]
  → Dispatch Validate(ctx) → Findings
  → Exit Policy: block if non-empty
```

**觀察**：自然；無 Projection；無 Dependency；Rule 零 Git。

### Rule B — 需「連結／共變」但非 Graph Projection

**選**：`obligation.commit.cli_doc_sync`（`validateCLIDocSync`）  
（作為 README/doc-sync + linked-update proxy 的代表）

```text
Adapter(commit-msg | CLI --staged --diff)
  → InputSnapshot{ staged_paths, staged_diff, commit_msg }
  → ContextBuilder → Context{ … }
  → (skip Projection)
  → (optional) Dependency/Cochange Provider: pair
        scripts/ai-skill-cli/internal/app/*.go
        ↔ scripts/ai-skill-cli/docs/command-contract.md
  → Capability Registry: requires [cap.staged_paths, cap.staged_diff, cap.path_cochange, cap.commit_msg]
  → Dispatch Validate(ctx) → Findings
```

**關鍵改造點（Phase 1 契約）**：Rule **不得** `exec git`；只讀 `ctx.StagedDiff`。若 Adapter 沒提供 diff → Finding=`capability_missing`（不是再回落呼叫 Git）。

**觀察**：Dispatcher / Capability / Context 仍自然；「需要的是 optional co-change，不是 Projection Builder 必經」。

### Spike 結論

| 檢查 | 結果 |
| --- | --- |
| Rule A 自然？ | Yes |
| Rule B 自然？ | Yes（條件：diff 進 Context） |
| Projection 必經？ | **No** — 兩條都不需要 |
| Capability 失控？ | No |
| 下一步 | 可寫 Capability 最小 schema 草案；**不需**先搬 Plans |

---

## 5. 對 Phase 0 Exit checklist 的進度

| Exit 項 | 狀態 |
| --- | --- |
| Context Builder 沒有一直加特例 | **暫過** — 特例風險集中在「誰提供 diff」（屬 Adapter，不屬無限 Context 欄位） |
| Projection 只有少數 rule 需要 | **暫過** — 焦點三域≈0；需要者在 Plans/runtime Later |
| Capability 有限且穩定 | **暫過** — 7-cap 工作集 |
| 同 Rule 多 Adapter 概念成立 | **暫過** — 條件：先剝 Git out of Rule B |
| Mini Spike A+B | **紙上完成**；可選後續極薄 stub（仍不改 hooks 大結構） |
| Plans 不當本階段 migration | **確認** |

→ **建議**：完成本 run 的 Capability schema 短草案後，即可評估關閉 Phase 0 → Phase 1。尚未自動關閉 Exit（待你簽核）。

---

## 6. 風險與非目標（本 run）

- 未實作 Go stub / 未改 `hooks.go` 結構（符合 Phase 0 約束）。
- 未宣稱 linked-updates.md 已機械化。
- 未把 sanitization / ownership drift 詳表化（屬同 capability 族，可 Phase 1 pack 擴充）。
- Plans / runtime_* validators 僅作 deferred 對照。

## 7. Next（建議）

1. 使用者簽核本 mapping 五問判定。  
2. 補一頁 `p0-capability-schema-draft.md`（欄位：id / inputs / produced_by_adapter / shared_by_rules）。  
3. 簽核後 → Phase 1：Context/Finding/Rule/Capability contracts + 抽 Rule A/B 薄路徑（CLI+hook 同 Context）。
