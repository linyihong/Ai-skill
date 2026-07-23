# 3a — KaizenWMS／Cursor：機械 bootstrap gate 讓 Task subagent 跑不起來 → 三角色被放棄（2026-07-23）

> **專案證據邊界**：consumer commit／plan 路徑留 KaizenWMS；本檔只留 generalized dogfood metrics 與契約回饋。  
> **同構**：[`2j`](2j-externalrepoc-push-execute-skip-verifier-loop.md)／[`2l`](2l-externalrepoc-common-url-s2-mirror-skip-loop.md)（跳過 Verifier）；[`2o`](2o-consumer-tab-scroll-single-vs-delegation.md)（單 session）；[`2y`](2y-kaizenwms-phase2-spa-scaffold-c1b.md)／[`2z`](2z-kaizenwms-phase3-karma-stale-serve.md)（同 consumer）。  
> **本 run 主反例（stakeholder 指定）**：不是「又少 spawn 一次 Verifier」 alone，而是 **通用機械 gate（Bootstrap Receipt／workflow primary_source）作用在 fresh Task subagent 上 → 看起來「三角色被 Ai-skill gate 卡住」→ orchestrator 乾脆不跑 E／V**。  
> **糾偏（consumer 已修）**：portable `delegated-execution` **§2.1 Cursor transport fallback**（spawn／gate 失敗 ≠ 豁免；同 session 驗證 ≤ `implementation_done`）。  
> **正向對照（同日）**：明示走 Task E＋V 的 shell-nav Phase 1 → Verifier 擋 F1 → `slice_compliant_closed`。

## 因果鏈（主反例）

```text
Execute 意圖 → 政策要求 O→E→V（獨立 Task）
  ↓
Cursor spawn fresh Executor／Verifier Task
  ↓
subagent 尚未滿足 gate.bootstrap.receipt_present
  （及／或 workflow primary_source Read）
  ↓
preToolUse 機械 deny 非-Read 工具  ←── 通用 bootstrap gate，非「三角色 deny gate」
  ↓
看起來像「Ai-skill 擋三角色驗證」
  ↓
orchestrator 放棄 spawn／改同 session 直做＋自驗
  ↓
Production＝Evidence（loop 斷裂）
```

**關鍵誤判**：把 **transport／bootstrap L3** 的攔截，讀成 **delegation loop 被禁止**。  
**事實**：`sd-delegated-execution` 仍為 **candidate**（沒有「未 spawn Verifier → deny」閘）；真常擋工具的是 **Bootstrap Receipt** 等通用 hook（見 Cursor adapter `preToolUse`）。

## Run 摘要

| 欄 | 值 |
| --- | --- |
| **Consumer** | KaizenWMS（Ai-skill optional；Cursor + project bootstrap pointer） |
| **Stakeholder 觸發** | 「三角色驗證好像被 Ai-skill gate 卡住，是不是要修，不然都會被擋著」 |
| **負向結果（糾偏前）** | 近期 Execute（shell chrome／list-detail 等）**0** 獨立 E／V；同 session 直做 |
| **根因（本 3a）** | **mechanical-gate × subagent cold-start** → 假「三角色不可用」信號 → 放棄 loop |
| **糾偏** | consumer §2.1 fallback 寫清：Task／bootstrap 失敗時的書面 transport；**不得**暗示 `slice_compliant_closed` |
| **正向對照** | 同日 shell-nav Phase 1：brief → Task E → Task V → F1 fix → closed |

## 失敗／不如預期

| # | 現象 | 根因分類 | 應有行為 |
| --- | --- | --- | --- |
| **F0（主）** | fresh Task 被 bootstrap／primary_source 機械 gate 卡住 → 不跑或半死 | **mechanical-gate × delegation-transport collision** | subagent brief 強制 bootstrap 讀＋Receipt；或 adapter 對 Task 冷啟動有明確路徑；失敗則 **§2.1 fallback**，不是放棄 loop |
| F1 | 可見 chrome／列表當同 session 「小刀」做完 | process-omission（2j／2l 同構） | Execute → loop；chrome／BDD ≠ surgical |
| F2 | 對外說「Ai-skill 擋三角色」 | **gate-misread** | 擋的是 bootstrap 工具權，不是三角色專用 deny |
| F3 | 自跑 verify 當獨立 Verifier | verifier 降級 | fresh Verifier；同 session ≤ `implementation_done` |

## 仲裁紀要

| finding | 處置 | 理由 | 後續 |
| --- | --- | --- | --- |
| F0 gate×subagent | **fix**（契約＋adapter 認知） | 通用機械 gate 不應默認殺死委派 transport | 3a 回饋；consumer §2.1 已落地；Cursor adapter 須寫清 Task 冷啟動 |
| F1 跳過 loop | **fix** | 流程 violation | process debt → `implementation_done` |
| F2 誤讀 | **reject**（敘事） | 無三角色專用機械 deny | 本檔＋SOP 用語修正 |
| F3 自證 | **fix** | C2／C4 | 獨立 V 後才可升 closed |

## 量測欄

| 指標 | 值 |
| --- | --- |
| stakeholder 明確指向「gate 卡住三角色」 | **1** |
| 負向 Execute（無獨立 E／V） | **≥2** |
| 真「Verifier-not-spawned」機械 deny 存在？ | **否**（candidate） |
| 常擋工具的機械層 | **Bootstrap Receipt／workflow primary_source（Cursor preToolUse）** |
| consumer 糾偏文件 | **§2.1 transport fallback** |
| 糾偏後正向 Task E＋V | **1**（shell-nav Phase 1） |

## 契約回饋（寫回 canonical／adapter／consumer）

1. **`mechanical-bootstrap-vs-delegation-transport`（主）** — `gate.bootstrap.receipt_present`／workflow `primary_source` 等 **通用**機械閘，作用在 **fresh Cursor Task** 上時，會產生「三角色被擋」的假信號。契約必須區分：  
   - **A** 三角色專用 deny（尚未存在於 Ai-skill candidate）  
   - **B** bootstrap／workflow 冷啟動 deny（存在）  
   B 失敗 → **transport fallback**，**不是**「loop 取消」。
2. **`subagent-cold-start-bootstrap-tax`** — Executor／Verifier Task brief **必須**要求首步 Read CORE_BOOTSTRAP＋`core-bootstrap.yaml`＋Receipt（Ai-skill repo）；consumer／portable 路徑則寫明：gate fail-open 或 §2.1 checklist。Kit／Cursor adapter 應有一句「Task 冷啟動必過 bootstrap，否則看起來像三角色壞了」。
3. **`spawn-or-gate-fail-not-exemption`** — Task spawn 失敗 **或** 機械 gate 連續 deny ≥N → 書面 `transport: same_session_*`；Verifier 仍要有；關閉 ≤ `implementation_done`。
4. **`portable-consumer-fallback`** — KaizenWMS 已示範不依賴 Q5 promotion 即可寫可執行 fallback（§2.1）。
5. **與 2j F6 對偶** — 2j：機械閘 **沒擋**「跳過 Verifier」。3a：機械閘 **擋了** subagent 工具 → **誘發**跳過。Q5 若只做 verifier-spawn tracking 而不處理 **bootstrap×Task**，仍會再出 3a。
6. **不關閉 Q5** — 3a 是 transport／adapter 反例＋行為契約回饋，不是 Shared State promotion 證據。

## 相對 2j／2l 的增量

| 主題 | 2j／2l | **3a** |
| --- | --- | --- |
| 跳過近因 | 故意單 Task／surgical／`enabled:false` | **機械 bootstrap gate → Task 不可用 → 放棄 spawn** |
| Gate 角色 | 「該擋略過卻沒擋」（缺口） | 「不該被讀成三角色禁令卻擋了冷啟動」（誤傷／誤讀） |
| 修復面 | verifier-after-executor tracking | **transport fallback ＋ adapter 冷啟動說明** |

## 四責任閉環（斷裂點）

```text
Specification — Execute／mandatory loop
  ↓
Production — ❌ Task 冷啟動被 bootstrap 機械 gate 卡住 → 改主 session
  ↓
Evidence — ❌ 與 Production 合併
  ↓
Decision — ❌ 未仲裁即「做完」
  ↓
Specification — 3a + consumer §2.1 + Cursor adapter 冷啟動註記
```

## Evidence pointers（consumer／adapter）

- KaizenWMS：`docs/workflow/delegated-execution.md` §2.1；`software-delivery.yaml` `delegation_loop.rules`
- 正向：`…/1422-shell-nav-groups-…/evidence/2026-07-23-phase-1-nav-contract.md`
- Cursor：`ai-tools/agent/cursor.md`（preToolUse＝機械 deny；`AI_SKILL_REPO` 不可解析時 fail-open）

## Disposition

- 更新本 evidence（本修訂以 F0 為主軸）＋ anti-pattern／`plans/README`／Cursor adapter 一句。  
- **不**視為 Phase 3／Q5 closure。
