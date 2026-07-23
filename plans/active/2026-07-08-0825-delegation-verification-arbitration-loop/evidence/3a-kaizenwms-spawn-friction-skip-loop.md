# 3a — KaizenWMS：spawn 摩擦誤當「Ai-skill 擋三角色」→ 同 session 跳過 loop（2026-07-23）

> **專案證據邊界**：consumer commit／plan 路徑留 KaizenWMS；本檔只留 generalized dogfood metrics 與契約回饋。  
> **同構**：[`2j`](2j-externalrepoc-push-execute-skip-verifier-loop.md)／[`2l`](2l-externalrepoc-common-url-s2-mirror-skip-loop.md)（跳過 Verifier）；[`2o`](2o-consumer-tab-scroll-single-vs-delegation.md)（單 session vs 三角色）；[`2y`](2y-kaizenwms-phase2-spa-scaffold-c1b.md)／[`2z`](2z-kaizenwms-phase3-karma-stale-serve.md)（同 consumer 正負前例）。  
> **正向對照（同日）**：shell-nav Phase 1 強制 O→E→V 後 Verifier 擋下 §3/§6 路徑矛盾（F1）→ fix → `slice_compliant_closed`。

## Run 摘要

| 欄 | 值 |
| --- | --- |
| **Consumer** | KaizenWMS（portable `docs/workflow/delegated-execution.md`；Ai-skill optional） |
| **負向切片** | SPA shell chrome（手機漢堡）＋列表 queryParams 還原等近期 Execute — **0** 獨立 Executor／Verifier；同 session 直做＋自跑 `verify-frontend` |
| **觸發敘事（錯誤）** | Agent／對話歸因：「Ai-skill gates／bootstrap 擋住三角色，所以只好自己做完」 |
| **事實** | `sd-delegated-execution` 仍為 **candidate**（無機械 deny「未 spawn Verifier」）；Cursor Task／子 agent bootstrap 成本高 → **軟性略過**，非硬 gate 拒絕 |
| **糾偏** | Consumer 寫入 **Cursor transport fallback**（spawn 失敗 ≠ 豁免；same-session Verifier 最高 `implementation_done`）；欠債 slice 回填 `implementation_done`、禁止追溯 `slice_compliant_closed` |
| **正向對照** | 同 consumer **shell-nav Phase 1**（文件契約）：brief commit → Task Executor → Task Verifier → F1 fix → re-verify → `slice_compliant_closed` |

## 相對 2y／2z 的退步（反模式）

| 2y／2z 已驗證 | 3a 負向表現 |
| --- | --- |
| Execute 前 brief／backfill／獨立 Verifier | 近期 chrome／列表 slice：**無** `delegation:`、**無** Verifier evidence |
| C1b／browser e2e 紀律（2z） | 有跑 verify／e2e，但 **Production＝Evidence＝同 session**（自證） |
| 關閉語意清楚 | 暗示「做完＝closed」；審計後才降為 `implementation_done` |

## 失敗／不如預期

| # | 現象 | 根因分類 | 應有行為 |
| --- | --- | --- | --- |
| F1 | 可見 shell／列表行為變更當「小刀」同 session 做完 | **process-omission**（同 2j／2l） | Execute 意圖 → mandatory loop；chrome／BDD ≠ surgical |
| F2 | 把 Ai-skill **bootstrap／receipt／candidate 成熟度**說成「擋三角色」 | **gate-misattribution** | candidate = 行為契約；缺的是 Verifier spawn 追蹤 gate，不是「禁止 loop」 |
| F3 | Cursor Task spawn／子 agent bootstrap 摩擦 → 直接主 session 包辦 | **transport-friction skip** | transport adaptation 須書面；Verifier 仍不可省；否則僅 `implementation_done` |
| F4 | 自跑 `verify-frontend` 當獨立驗證 | **verifier 降級** | V1 須 fresh context；同 session 通過 ≠ `slice_compliant_closed` |

## 仲裁紀要（consumer Orchestrator／stakeholder，2026-07-23）

| finding | 處置 | 理由 | 後續 |
| --- | --- | --- | --- |
| F1 跳過 loop | **fix**（流程＋文件） | acceptance-violation（流程） | consumer §2.1 fallback；欠債標 `implementation_done` |
| F2 gate 誤讀 | **reject**（誤讀） | Ai-skill 未機械 deny Verifier spawn | 本 3a 回饋；Q5／機械 gate 仍 open |
| F3 spawn 摩擦豁免 | **fix**（契約） | spawn 失敗 ≠ surgical／≠ Execute 豁免 | portable fallback 順序寫進 consumer YAML |
| F4 自證關閉 | **fix** | C2／C4 | 獨立 Verifier 後才可升 `slice_compliant_closed` |

## 量測欄

| 指標 | 值 |
| --- | --- |
| 負向 Execute slice（審計） | **≥2**（shell hamburger；list-detail restore） |
| 負向 Executor／Verifier spawn | **0／0** |
| gate-misattribution 事件 | **≥1**（stakeholder 明確指出） |
| consumer 契約回寫 | **1**（portable delegated-execution §2.1＋YAML rules） |
| 糾偏後正向 E+V（同日） | **1**（shell-nav Phase 1；Verifier findings ≥1 → fix → closed） |
| Ai-skill 機械 gate 阻止略過 | **否**（candidate；與 2j F6 同構） |

## 契約回饋（寫回 canonical／consumer overlay）

1. **`spawn-friction-not-exemption`** — Cursor Task／Verifier spawn 失敗、bootstrap 成本高、或「怕 gate」**不得**豁免三角色；須走 **transport fallback**（書面 `same_session_executor`／same-session Verifier checklist），關閉上限 `implementation_done`。  
2. **`gate-misattribution`** — 不可把 **candidate／behavioral** 契約說成「runtime 擋我做 loop」。缺的是 Verifier-spawn tracking（2j）；現況是 agent **選擇跳過**，不是被 deny。  
3. **`portable-consumer-fallback`** — 外部 repo 可在自有 `delegated-execution` 寫 Cursor fallback（KaizenWMS 已落地）；Ai-skill 正文加深 anti-pattern，**不**等 Q5 promotion 才有可執行路徑。  
4. **`chrome-bdd-never-surgical`** — shell chrome／列表行為／新 BDD＝Execute → loop（與 consumer surgical 窄邊界一致）。  
5. **Q5／機械 gate 仍欠** — 3a 再證：無 `verifier_required_before_next_executor` 時，摩擦會重複製造 2j／2l／3a 負向；**不**因本 run 提前關閉 Q5。  
6. **正負同日對照有效** — 審計債務 + 強制下一 slice 走 loop（Verifier 真抓到 F1）＝ 2l→2m 同構的 consumer 版。

## 相對 2j／2l／2o 的增量

| 主題 | 2j／2l／2o | **3a 新增** |
| --- | --- | --- |
| 跳過理由 | `enabled:false`／surgical 濫用／「小修」 | **誤指 Ai-skill gate 擋住三角色** |
| Transport | 單 Task 包辦 | **spawn 摩擦 → 主 session 包辦**（未標 transport adaptation） |
| Consumer 成熟度 | ExternalRepoC 有機械 gate 候選 | KaizenWMS **portable-only** 仍可寫 fallback 並跑通正向 loop |
| 關閉語意 | 常缺 `implementation_done` | 審計強制降級＋正向 slice 達 `slice_compliant_closed` |

## 四責任閉環（負向 run — 斷裂點）

```text
Specification — 使用者 Execute／可見 chrome（缺 brief／backfill）
  ↓
Production — 主 session 直做（應為 Executor leg）
  ↓
Evidence — ❌ 與 Production 合併（自跑 verify）
  ↓
Decision — ❌ 未仲裁即暗示完成
  ↓
Specification — 3a 本文 + consumer §2.1 + 本 plan anti-pattern
```

**斷裂點**：把 **transport 脆弱性** 誤讀成 **governance 禁止 loop** → Production／Evidence 合併。

## Evidence pointers（consumer）

- `docs/workflow/delegated-execution.md` §2.1／§8.1（process debt 表）
- `docs/workflow/software-delivery.yaml` `delegation_loop.rules`（spawn failure ≠ exemption）
- `plans/active/2026-07-23-1422-shell-nav-groups-module-isolation/evidence/2026-07-23-phase-1-nav-contract.md`（正向對照）

## Disposition

- **寫回**：本 evidence + plan 索引／checkbox；`delegated-execution.md` anti-pattern 列；`plans/README.md` Delegation SOP 一小段。  
- **不**視為 Phase 3／Q5 closure。  
- **不**新建 feedback-history lesson 檔（本輪以 plan dogfood evidence 為準；可後續 promote）。
