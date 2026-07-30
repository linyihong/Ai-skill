# 3b — KaizenWMS：多輪獨立 Verifier 成熟度（ops-chrome／scanner／device slices，2026-07-27–30）

> **專案證據邊界**：commit SHA、selector、feature 檔名留 consumer `<PROJECT_ROOT>` plan；本檔只留 generalized dogfood metrics 與契約回饋。  
> **Consumer 主來源**：`plans/active/2026-07-23-1500-spa-mobile-desktop-devices/_plan.md`（slices `phase-1b-1-ops-chrome-lite`、`phase-1b-2-scanner-bar`、Phase 2 device 改版多刀）。  
> **同構／前序**：[`2y`](2y-kaizenwms-phase2-spa-scaffold-c1b.md)／[`2z`](2z-kaizenwms-phase3-karma-stale-serve.md)／[`3a`](3a-kaizenwms-spawn-friction-skip-loop.md)（同 consumer；3a 後已能穩定走獨立 Task E＋V）。  
> **本 run 主訊號**：**不是又發現「跳過 Verifier」**，而是 loop 跑起來之後的**仲裁成熟度**——多輪 fix→re-verify、角色表 ownership、空閘／vacuous e2e、契約自傷、假證據拒絕。

## Run 摘要

| 欄 | 值 |
| --- | --- |
| **Consumer** | KaizenWMS（Cursor Task；`verifier_transport: independent_subagents`） |
| **時間窗** | 2026-07-27 → 2026-07-30（§2.1 fallback **未**動用） |
| **代表 slices** | ops-chrome-lite（Verifier **三輪**）、scanner-bar（**兩輪**）、Put Away／Transfer／Issue 單欄（各至少一輪＋fix round） |
| **關閉形態** | 多刀皆達 `slice_compliant_closed`（C1–C5 書面宣告） |
| **對照負向（同 plan）** | Phase **1a** 同 session 直做 → 回填 **`implementation_done`**（不得假稱 closed） |

## 因果鏈（成熟度，非失敗主線）

```text
3a 糾偏後 → 可穩定 spawn 獨立 E／V
  ↓
多 slice 完整 O→E→V＋多輪仲裁
  ↓
暴露「loop 內」契約缺口（非 transport 缺口）：
  · backfill／Arbitration 表誰可寫
  · gate 掃描範圍 < 契約 shell 範圍 → 假綠
  · orchestrator 補契約字面自傷（缺手動例外）
  · vacuous e2e／為製造 submitting 窗口而植假 async
  · G5 守住後 C1／C4 變 Orchestrator own debt
  ↓
契約回饋（本檔）→ consumer brief 明寫角色邊界；SD 反模式候選
```

## 失敗／不如預期（taxonomy）

| # | 現象 | 根因分類 | 應有行為 |
| --- | --- | --- | --- |
| **F1** | Gate 只掃 app shell SCSS，lib dashboard 仍含硬編碼 breakpoint → gate 綠、契約破 | **gate-scope ≠ acceptance-scope**（假證據） | Gate／mutation 範圍必須覆蓋 brief 定義的 shell／lib 範圍 |
| **F2** | Orchestrator 補「跨帶自動收合」未寫手動例外 → Executor 照字面實作 → UX 覆蓋使用者選擇 | **orchestrator-contract-self-injury** | 補契約時寫清優先序；fix 歸契約作者，不怪 Executor |
| **F3** | Executor 改寫了 Orchestrator 的 verification backfill 列 | **role-boundary process drift**（內容可為真） | Brief 明示：Executor **不得**編輯 backfill／Arbitration 表 |
| **F4** | G5 守住後 backfill 長期裸 `pending` → C1／C4 FAIL | **orchestrator-own bookkeeping debt** | 關閉前 Orchestrator **必須**自行刷新 linked／數字 |
| **F5** | Tablet viewport 上「內容溢出」斷言 vacuous；後改 phone 幾何才具鑑別性 | **vacuous acceptance evidence** | e2e 須含 non-vacuity gate（溢出／觸及 bar 等可觀測前提） |
| **F6** | 「提交中 disable」產品端零綁定；逼 e2e＝逼植假 async | **planted-async-false-evidence** | 拒絕假證據；記 `deferred`＋解除條件（真 API），勿為綠而造窗 |
| **F7** | HANDOFF 內 commit SHA 因 history 重寫失效 | **stale-handoff-ref** | Closure 用可驗證 HEAD／路徑證據；書寫 SHA 須可覆核 |
| **F8** | Verifier 工作單寫 `git stash` 還原壞版本 → 乾淨樹不可照字面執行 | **brief-wording vs discriminating probe** | 鑑別性還原用 `checkout <bad> -- file` 類可逆步驟；偏差須回報 |

## 仲裁紀要（泛化）

| 叢集 | 典型處置 | 契約含義 |
| --- | --- | --- |
| Gate／契約範圍不一致 | **fix** | 假綠不可關閉 |
| 契約自傷（缺例外列） | **fix**（orchestrator 認帳） | Production 無過失 |
| Backfill 被 Executor 改 | **accept 內容／記 process** | 下一刀 brief 寫死邊界 |
| Integration 0 ran／skipped | **defer**＋誠實 backfill | 不得冒充第二條證據 |
| Vacuous／假 async | **fix 幾何**／**defer 契約列** | 空斷言≠ linked |
| Verifier 自撰 mutation／probe | **accept** | L3 成熟訊號（對齊 Q9 方向，仍不升 schema） |
| Finding **refuted**（編譯／cascade 證明無視覺回歸） | **reject／refuted** | 仲裁可駁回，不強迫修 |

## 量測欄

| 指標 | 值（量級） |
| --- | --- |
| 獨立 Task E＋V（本窗代表 slices） | **≥5** closed |
| ops-chrome 獨立 Verifier 輪次 | **3** |
| scanner-bar 獨立 Verifier 輪次 | **2** |
| §2.1 same-session fallback 使用 | **0**（本窗正向切片） |
| 同 plan 仍標 `implementation_done`（1a） | **1**（誠實降級） |
| 反覆出現的 process finding | **G5／J12／K11 同源**（backfill ownership） |
| Verifier 降級為「只重跑 V1」 | **0**（宣告 C4） |
| 契約回饋條數（本檔） | **8**（見下） |

## 契約回饋（寫回 SD／kit／consumer overlay）

1. **`backfill-arbitration-orchestrator-owned`** — `verification_backfill` 與 Arbitration 表的 writer＝Orchestrator；Executor 只寫執行紀錄。Brief 必填禁令（KaizenWMS 已示範）。
2. **`gate-scope-covers-acceptance-shell`** — 機械閘掃描範圍須 ⊇ acceptance 所稱 shell／shared lib；否則 mutation 證明「空閘」。
3. **`orchestrator-owns-contract-gaps`** — 仲裁補列造成的行為衝突，歸 Specification／Arbitration，不記 Executor acceptance-violation。
4. **`vacuous-e2e-non-vacuity-gate`** — 溢出／遮擋類 acceptance 必須先證明「溢出／遮擋可發生」，再斷言結果。
5. **`no-planted-async-for-submitting`** — 同步 mock 無 submitting 窗口時，不得為 e2e 植入人工 delay；用 `deferred`＋真 API 解除條件。
6. **`c1-c4-after-g5-is-orchestrator-debt`** — 禁止 Executor 改 backfill 之後，裸 `pending`／過期數字是 **Orchestrator 關閉阻塞**，不是 Verifier 可略過的 observation。
7. **`discriminating-revert-probe`** — L3／fix 鑑別性：還原「壞」版本須可逆、可覆核；工作單勿假設 dirty stash 工作流。
8. **`implementation_done-sibling-honesty`** — 同 plan 內未走獨立 V 的刀必須維持 `implementation_done`；不得因鄰近刀已 closed 而回填假合規。

## Evidence pointers（consumer）

- `plans/active/2026-07-23-1500-spa-mobile-desktop-devices/_plan.md` — Arbitration 表 F／G／H／J／K／L／M 叢集；Closure C1–C5
- Phase 1a process note（同檔）— `implementation_done` 誠實降級
- Portable overlay：`docs/workflow/delegated-execution.md` §2／§7（consumer）

## Disposition

- 本輪：**evidence + 索引**；強化 Q3（多輪 fix 差集品質）與 Q8（Evidence Producer ≠ Closure Authority：Verifier 產證、Orchestrator 關 C5）。  
- **不**視為本 plan Phase 3／Q5 closure；**不**升 schema。  
- 建議 SD `delegated-execution` 反模式表加：`executor-edits-backfill`、`vacuous-e2e`、`planted-async`、`gate-under-scope`（advisory）。
