---
id: 2026-07-08-0825-delegation-verification-arbitration-loop-04
plan_kind: companion
status: active
parent: 2026-07-08-0825-delegation-verification-arbitration-loop
created: 2026-07-14
last_updated: 2026-07-14
---

# Domain vs Workflow — APK Analysis ↔ Software Delivery（Capability Handoff）

> **Companion** of [`_plan.md`](_plan.md)。本檔是 stakeholder 裁決紀錄 + dogfood 預註冊；**不是** workflow SOP promotion。canonical handoff 產物格式仍以 [`workflow/apk-analysis/artifact-gates/feature-handoff.md`](../../../workflow/apk-analysis/artifact-gates/feature-handoff.md) 為現況；本檔定義的 **Capability Handoff** 是其上界的 **Domain Boundary 契約候選**（觸發 = deliverable capability，不是 analysis-complete）。

## Stakeholder 裁決（2026-07-14）

1. **Domain Layer 與 Workflow Layer 分離** — APK Analysis 與 Software Delivery 是兩個 Domain；各自有 Decision Semantics 同構於 ERA（Evidence → Feasible Set → Selection），但回答**不同問題**。
2. **不對稱成熟度保留** — 不得把 `software-delivery` 已驗證的 Delegated Execution 宣稱成 `apk-analysis` 的 workflow。APK 側目前成熟的是 Domain Knowledge + Discovery route（candidate）；SD 側成熟的是 Delegated Execution（validated）。
3. **Handoff 是 Domain Boundary，不是 workflow 步驟** — 兩條 route 之間用正式契約銜接；SD **不**需要知道 Frida / mitm / hook。
4. **觸發條件** =「形成可交付能力（deliverable capability）」，**不是**「分析完成」。

## Domain Layer

```text
APK Analysis Domain
      │  發現可交付能力
      ▼
Capability Handoff Boundary   ← Domain Boundary（本檔契約）
      │
      ▼
Software Delivery Domain
```

| Domain | Decision 問句（ERA） | 成熟度（現行證據） |
|---|---|---|
| **APK Analysis** | Evidence → *Can we explain the APK?* | **Domain Knowledge 成熟**；Discovery／RE workflow = **candidate**；**無**三角色 / Verifier / ERA orchestration 真實 run |
| **Software Delivery** | Evidence → *Can we ship this capability?* | **Delegated Execution validated**（本 plan 2a–2s 累積；sd-delegated-execution） |

相同 Decision Semantics，不同問題 → **不同 Domain**。不得因「都用證據」就把兩條 workflow 合併成同一 family。

## Workflow Layer（不對稱）

```text
APK Analysis Route（candidate）
  ├── Discovery
  ├── Reverse Engineering
  ├── Evidence Collection
  ├── API Mapping
  ├── UI Mapping
  └── Capability Assessment          ← handoff 評估出口
           │
           │  Capability Proposal（僅當 Yes）
           ▼
Software Delivery Route（validated）
  ├── Specification
  ├── Delegation
  ├── Implementation
  ├── Verification
  └── Delivery
```

| 路線 | 定位寫法（強制） | 禁止寫法 |
|---|---|---|
| `workflow/apk-analysis` | **Current Workflow（candidate）** — Discovery / RE；自累 dogfood 與跨 consumer 證據 | 「已採用 Delegated Execution」「三角色是 APK SOP」 |
| `workflow/software-delivery` | **Delegated Execution（validated）** — 能力交付 | 把 Frida/mitm 塞進 SD execution model |

## Evidence vs Delivery Artifact（分類鐵則）

仍屬 **Discovery Evidence**（留在 APK Analysis；**不**觸發 handoff）：

- API catalog / protocol / crypto flow / UI flow / event model
- hook／proxy／pcap 日誌、重建所需但尚未可交付的觀察

跨到 **Delivery Artifact**（才可進 Software Delivery Intake）的訊號 — 開始出現可重用交付物意向：

- SDK / Client / Contract / OpenAPI / BDD / 測試套件 / 可重用 library

既有 [`feature-handoff`](../../../workflow/apk-analysis/artifact-gates/feature-handoff.md) 仍是 APK **artifact gate**（8 面向重建輸入）。與本契約關係：

| 概念 | 層 | 問句 |
|---|---|---|
| Feature Reconstruction Handoff | APK workflow artifact gate | 「後續能否 *重建* 此 feature？」 |
| **Capability Proposal / Capability Handoff** | Domain Boundary | 「是否已形成可 *交付* 的可重用能力？」 |

Feature handoff **可餵** Capability Proposal，但 ** alone 不足** 構成 Delivery Intake（預註冊 F1，見下）。

## Capability Handoff 契約（候選）

```text
APK Analysis Output
  → Capability Assessment
  → Can become reusable artifact?
       No  → Continue APK Analysis
       Yes → Software Delivery Intake
```

**Handoff 輸入不是 APK**，是 **Capability Proposal**：

```yaml
# capability-proposal (candidate schema — dogfood 前 doc-only)
capability:
  type: sdk | client | contract | openapi | bdd_suite | library | other
evidence:           # Discovery Evidence 引用（去敏 path）
  - api_catalog
  - protocol
  - auth_flow
  # …依 type 最小集
deliverables:       # 預期 Delivery Artifact
  - client_sdk
  - contract
  - tests
non_goals:
  - frida_hooks      # SD 不消費
  - mitm_setup
intake_owner: software-delivery
```

**Software Delivery Intake 只回答**：現在有一個能力要交付（Specification → Delegation → …）。**不**載入 APK RE 機械細節。

## Dogfood 預註冊 — **2t**（可加入，但不對稱）

> 與本 plan Phase 2「跨域 / 穩定性」相容；**不**填 Knowledge 格；**不**把 stage 2 3/3 偷換成 APK；**不**視為 Q5/Q6/Q8 Phase 3 closure。

### 為何可以加入 dogfood

- 真實 APK 專案在跑 → 滿足「真實任務、不 manufacture」。
- 本 plan 的價值不只測三角色，也測 **Domain Boundary + Decision Semantics 分離**（第九輪「Decision Semantics 不等於 Workflow」的直接實例）。
- 若能力升到可交付 → 可測 **Handoff → SD Intake → Delegated Execution** 的銜接，而不污染 APK candidate workflow。

### 為何不能直接套三角色到 APK

| 缺什麼（APK） | 已有什麼（SD） |
|---|---|
| Delegation loop / 三角色 evidence | 多次 dogfood + External consumer |
| Verifier / independent evidence run | Research / Architecture 跨域同構 |
| Orchestration 契約 | Governance collision / ERA |

無證據即複製 = 違反 promotion discipline。

### 2t 雙軌設計

| Track | 名稱 | 驗什麼 | 不驗什麼 |
|---|---|---|---|
| **2t-A** | APK Discovery（candidate workflow） | Discovery Evidence 鏈是否支撐 Decision「能解釋 APK？」；Capability Assessment 是否誠實標 No/Yes | Delegated Execution、三角色 mandatory |
| **2t-B** | Capability Handoff → SD Intake（僅 Yes 時） | Capability Proposal 自足性；SD 能否只憑 Proposal 開 brief（零 Frida context）；其後用 **既有 SD loop** 交付 | 「APK 分析本身走完 E→V→A」 |

2t-B 啟動後的交付腿 = **既有 software-delivery dogfood 形狀**（可另編 slice id，證據檔鏈到本 2t）。2t-A **單獨也可關閉**（若本輪未形成 deliverable capability — 那是有效負結果，支持邊界紀律）。

### 預註冊判準（run 前鎖定）

| ID | 判準 | Pass | Fail（邊界訊號） |
|---|---|---|---|
| **F1** | 僅有 API catalog / protocol / UI flow 等 Discovery Evidence **不得**當 Capability Proposal 過 Intake | Intake 拒絕或降回 Continue Analysis | 把 endpoint 表當 SDK deliverable 開 SD |
| **F2** | Handoff 觸發 = deliverable capability，**不是** analysis-complete | Assessment 明示 Yes/No + type | 「分析做完所以 handoff」無 deliverables |
| **F3** | SD Intake brief 的 `context.required` **不含** Frida/mitm/RE SOP；只含 Proposal + 去敏契約 | Executor 可 brief-only | SD session 被迫讀 hook 腳本才能開寫 |
| **F4（asymmetry）** | 證據敘述不得宣稱 APK 已 validated Delegated Execution | 措辭維持 candidate | 跨域表把 APK 標成與 SD 同 maturity |

### 量測欄（預留）

| 指標 | 預期記法 |
|---|---|
| Discovery Evidence 件數（去敏類別） | catalog / protocol / UI / crypto / … |
| Capability Assessment | `no` / `yes:<type>` |
| F1–F4 | pass/fail + 一句證據 |
| 2t-B 是否啟動 | yes/no；若 yes → 連結 SD evidence / brief path |
| Orchestrator 是否把 APK RE 細節寫進 SD brief | 0 為目標 |

### 執行啟動條件

1. 使用者指定真實 APK consumer project（`<PROJECT_ROOT>` 級，細節留 consumer plan）。
2. Orchestrator 開 `evidence/2t-apk-capability-handoff.md`（run 時新增）+ 本檔 checkbox。
3. 先跑 2t-A；Assessment = Yes 才開 2t-B。
4. **禁止**為湊 dogfood 而假造 Capability Proposal。

## 對本 plan 既有裁決的對齊

- **ERA**：兩域同 Decision Semantics、不同問句 → Domain 分開合理。
- **Promotion discipline / falsification ladder**：APK workflow 維持 candidate；不提前泛化。
- **Adoption stage 2**：2t **不**替代 Knowledge 格；跨域表 APK 列為 **Domain Knowledge + candidate workflow**，不是「四責任已驗證」。
- **既有 feature-handoff**：保留；本契約是 boundary 升層候選，升格前不動 `workflow/` 正文強制度（dogfood 證據後再 linked-update）。

## Open follow-ups（不升格，登記）

- [ ] Dogfood **2t-A** 真實 APK run（evidence 檔待建）
- [ ] 若 Yes → **2t-B** Capability Proposal + SD Intake brief
- [ ] 證據夠 → 評估是否將 Capability Handoff 寫回 `workflow/apk-analysis`（boundary 專節）+ SD intake pointer（獨立 linked-update；本 companion 不自 promote）
- [ ] Glossary 候選：`capability_handoff` / `deliverable_capability` / `discovery_evidence` — graduate 時才註冊

## 與其他檔

- Parent：[`_plan.md`](_plan.md) §Stakeholder / Phase 2 **2t**
- Kit：[`01-dogfood-prompt-kit.md`](01-dogfood-prompt-kit.md) §2t 摘要
- Evidence index：[`evidence/README.md`](evidence/README.md)
- APK artifact gate（現況）：[`feature-handoff.md`](../../../workflow/apk-analysis/artifact-gates/feature-handoff.md)
- SD execution（validated）：[`delegated-execution.md`](../../../workflow/software-delivery/delegated-execution.md)
