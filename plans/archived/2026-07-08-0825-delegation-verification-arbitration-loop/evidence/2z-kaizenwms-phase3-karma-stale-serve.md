# Dogfood 2z — KaizenWMS Phase 3：Karma-only 假綠 + stale `ng serve`（browser / runtime 缺口）

**Date**: 2026-07-21  
**Consumer**: `kaizenwms`（Bitbucket；`plans/active/2026-07-14-1610-wms-material-foundation`）  
**Loop**: Phase 3 曾標 `slice_compliant_closed`（unit/mock）；stakeholder 糾偏後補 Playwright + 機械門檻  
**同構**：[`2q`](2q-externalrepoc-transport-inner-only-runtime-gap.md)（inner-only 假綠）、[`2k`](2k-externalrepoc-push-post-close-runtime-gaps.md) / [`2m`](2m-externalrepoc-phase-g-mirror-batch-retrofit.md)（stale runtime JVM）、[`2y`](2y-kaizenwms-phase2-spa-scaffold-c1b.md)（同 consumer，C1b 遞進）

> **專案細節邊界**：commit SHA、selector、feature 檔名留 consumer evidence；本檔只留 generalized dogfood metrics 與契約回饋。

## What happened

| 現象 | 細節（泛化） |
| --- | --- |
| Verifier / close | Phase 3 List/Detail/Timeline Mock：Karma + TestBed 綠 → 宣告合規關閉 |
| Stakeholder | 「沒在瀏覽器跑過 BDD＝假過」；三角色驗證感覺是假的 |
| Runtime 手驗 | 長跑 `ng serve` 後改碼，畫面仍像舊版；操作按鈕「沒反應」 |
| 糾偏 | Consumer 加 Playwright M-E2E、`verify-frontend-e2e` 進每次 `verify-frontend`；DoD：Karma alone ≠ SPA Done；C1c 類 browser evidence |

## Failure taxonomy（對齊既有標籤）

| ID | Finding | Class | 對照 |
| --- | --- | --- | --- |
| F1 | user-visible SPA acceptance 只靠 **FE unit / mock** 關閉 | **inner-only / missing constraint**（缺 browser path proof） | 2q transport inner-only |
| F2 | Verifier V1 只重跑 Karma → **複讀自驗**，未獨立開瀏覽器 | **verifier 降級** | 2j/2q「有 Verifier 仍假綠」 |
| F3 | 長跑 `ng serve` **檔案監看失效 / 進程僵死**（含 EMFILE 類 watcher 耗盡）→ HMR／live reload **不觸發**，仍 serve 舊 bundle | **runtime stale**（對照 stale JVM） | 2k F1 / 2m stale-jvm |
| F4 | 部分「按鈕無效」實為 **業務 enablement**（狀態∩位置）或錯誤被吞，unit 未覆蓋瀏覽器旅程 | **acceptance-gap**（非純 HMR） | 2k post-close UI gap |

## Feedback to this plan / SD contract

1. **`unit-green ≠ SPA path proof`** — SPA 可見 acceptance 的 C1b／合規關閉，至少一行須是 **browser e2e**（或同等 runtime UI smoke），不可僅 Karma／TestBed。與 2q「features + outer integration」同構：缺的是 **路徑約束**，不是少 spawn 一次 Verifier。
2. **`verifier-must-rerun-browser`** — SPA slice 的 V1 必須含獨立重跑 Playwright／瀏覽器門檻；只信 Executor 的 unit 綠 = 降級（記入量測欄）。
3. **`stale-dev-server-checklist`（V5-U 候選）** — 與 `stale-jvm-v5-a-checklist` 並列：手驗或 Verifier runtime 前檢查 watcher／端口／必要時 **cold restart** `ng serve`（或等效）；勿假設 HMR 永遠有效。
4. **`HMR ≠ acceptance evidence`** — Angular HMR／live reload 是開發便利，**不是**獨立驗證；機械門檻應自起乾淨 serve（如 Playwright `webServer` + `CI=1` 不 reuse）。
5. **Consumer 已落地（參考，非 Ai-skill schema）**：`verify-frontend-e2e`、FeRef 必須連 `tests/frontend/**/e2e/`、`SKIP_E2E` 不可關 slice。

## Angular HMR 為何「更新了卻還是舊的」（給 stakeholder）

Angular `ng serve` **有** live reload／HMR，但本次不是「框架沒開 hot reload」，而是：

1. **Watcher 壞了仍在聽 port** — 進程活著、舊 compile 結果繼續對外；改檔不重建（EMFILE／監看耗盡常見於 macOS + 大樹）。
2. **HMR 覆蓋面有限** — 路由、`app.config`／interceptor、部分 provider、資產路徑等常需 **全量 rebuild**；HMR 失敗時若未觸發完整 reload，瀏覽器仍跑舊模組。
3. **錯把「舊分頁」當真相** — Mock in-memory store／已開分頁未 hard refresh，看起來像功能沒更新。
4. **驗證應用冷啟動** — Playwright／CI 應自己起 serve，不 reuse 開發者可能已 stale 的實例。

**處置**：懷疑 stale 時 kill port 上的 `ng serve` 再起；合規證據以 **冷啟動 e2e** 為準。

## Evidence pointers（consumer）

- `plans/.../evidence/2026-07-21-phase-3-browser-e2e.md`
- `plans/.../evidence/2026-07-21-phase-3-arbitration.md`（當時僅 unit；後補 browser 門檻）
- Consumer docs：`docs/workflow/delegated-execution.md` C1c；`scripts/verify-frontend-e2e`

## Disposition

- 本輪：**evidence + plan 索引**；建議 SD `delegated-execution` 反模式表加一行「SPA unit-only / stale serve」（advisory，不動 schema）。
- **不**視為本 plan Phase 3 closure；**不**升 enforcement。
