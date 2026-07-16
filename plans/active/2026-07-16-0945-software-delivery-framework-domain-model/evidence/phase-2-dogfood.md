# Phase 2 Dogfood — 3 New Artifact Questions

驗證 Primary Model（N=3）能否回答「新產物放哪、誰擁有、哪個 Process stage」而**不需開新目錄**。

Canonical policy：[`workflow/software-delivery/domain-policies.md`](../../../workflow/software-delivery/domain-policies.md)

---

## Q1：新功能要產出 UI governance evidence（accessibility scan + screenshot）

| 維度 | 答案 | 依據 |
| --- | --- | --- |
| **Asset class** | Evidence（project delivery asset） | domain-policies §3.2 |
| **Placement** | `docs/evidence/`（專案）；raw screenshots 可 gitignored `.local/` | domain-policies §3.2、§6 |
| **Owner** | 驗證 / review 執行者（專案） | domain-policies §2 |
| **Process stage** | `sd-validation` / UI governance slice（Ship 前） | execution-flow + artifact-gates §5.1 evidence shape |
| **新目錄？** | 否 — 用既有 `docs/evidence/` + template | — |

---

## Q2：replacement 變更要建 parity inventory

| 維度 | 答案 | 依據 |
| --- | --- | --- |
| **Asset class** | Plan deliverable | domain-policies §3.2 |
| **Placement** | `docs/planning/` 或併入 delivery-pack | domain-policies §3.2 |
| **Owner** | 發起人 / PM（專案） | domain-policies §2 |
| **Process stage** | `sd-intake`（code 前 blocking gate） | intake.md Parity Gate；Process 只定時序 |
| **新目錄？** | 否 | — |

---

## Q3：APK 分析 session 產出可重用 hardening lesson

| 維度 | 答案 | 依據 |
| --- | --- | --- |
| **Asset class** | Knowledge（framework）— 須 promotion | domain-policies §5、§6 |
| **Placement** | 先 `feedback/history/development-guidance/`；promote 後 `controls/` 等 | domain-policies §3.1、§5 |
| **Owner** | framework maintainer（Ai-skill） | domain-policies §2 |
| **Process stage** | 非 linear workflow stage；closure / promotion pipeline | promotion policy |
| **新目錄？** | 否 — 用既有 feedback → intelligence/workflow 管道 | — |

---

## Grep：placement dual source-of-truth（2026-07-16）

在 `workflow/software-delivery/**/*.md` 搜尋 `docs/planning`、`放置位置`、`放置規則`：

- **Canonical 全文**：僅 `domain-policies.md` §3
- **artifact-gates.md §2**：pointer 表，無雙寫路徑表
- **perf-governance.md**：`docs/evidence/perf/` 為 domain-specific 子路徑範例，不與 §3.2 衝突

**結論**：workflow slice 無 placement dual source-of-truth。
