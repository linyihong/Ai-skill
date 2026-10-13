# 2t — APK Capability Handoff（consumer Discovery 實跑回饋）

Status: **partial / 2t-A observational**（非正式 kit 模板跑完；**未**啟動 2t-B）  
Parent: [`04-apk-capability-handoff-boundary.md`](../04-apk-capability-handoff-boundary.md)  
Date: 2026-07-14

## Consumer（去敏）

| Field | Value |
| --- | --- |
| PROJECT_ROOT | authorized APK analysis workspace（消費者專案） |
| Route | `workflow/apk-analysis`（Discovery） |
| Project overlay | `.ai-skill/project/rules/apk-analysis-discovery-handoff.md` → 指向本 plan companion **04** |
| Windows | 多窗 static／MITM／Frida；含 host／HTTP stack／chapter 管線／**可重放 request-sign 公式** |

## 用了 plan 的哪一块？

| 資產 | 實際用法 |
| --- | --- |
| **`04-apk-capability-handoff-boundary.md`** | **有** — 經 project overlay 強制 Domain Boundary：Discovery ≠ SD Delegated Execution；每窗 **Capability Assessment** |
| `_plan.md` 三角色／ERA Execute loop | **無** — 依 04／overlay：**不**對純 APK Discovery 套 Orchestrator→Executor→Verifier |
| `01-dogfood-prompt-kit.md` §2t 模板 A/B/C | **未跑** — 本檔為實跑後回填，非正式 2t 啟動 checklist |
| dogfood **2t-B**（Capability Proposal → SD Intake） | **未啟動**（Assessment 持續 **No**） |

## Capability Assessment（多窗後）

| Field | Value |
| --- | --- |
| Assessment | **`no`** |
| 即使已有 | hosts、OkHttp 棧、chapter wire field、**.request-sign 可重放公式**（alg + concat；key 仅指纹） |
| 仍缺（交付物） | SDK／可交付 client／去敏契約套件／BDD — 故 **不**開 software-delivery intake |

→ 對 04 觸發條件「deliverable capability，不是 analysis-complete」：**正例**（分析深度上升仍維持 Continue Analysis）。

## F1–F4（本 observational run）

| ID | Result | Note |
| --- | --- | --- |
| **F1** | **pass** | 有 endpoint／protocol／crypto Discovery Evidence 仍不當 Capability Proposal |
| **F2** | **pass** | Assessment 明示 No；未因「公式齊」宣告 analysis-complete handoff |
| **F3** | **n/a** | 未開 SD brief（正確） |
| **F4** | **pass** | 措辭維持 candidate Discovery；未宣稱 APK Delegated Execution validated |

## 量測欄

| 指標 | 記法 |
| --- | --- |
| Discovery Evidence 類別 | catalog／hosts／proxy-bypass／HTTP stack／chapter blob／decrypt／**request-sign formula**／wire field |
| Capability Assessment | `no` |
| 2t-B 是否啟動 | **no** |
| Orchestrator 是否把 RE 細節寫進 SD brief | **0**（無 SD session） |
| 三角色套用在 Discovery？ | **0**（符合 04／overlay） |

## 契約回饋（可沉淀）

1. **Project overlay 有效**：消費者不需背整份 `_plan.md`；只 load `apk-analysis-discovery-handoff.md` → 04 即能阻止「簽章公式還原＝開 SD」誤跳。
2. **Assessment = No 的高壓案例**：當 agent 已握有 replayable sign 時，最容易誤把 Discovery 當 deliverable；本 run 靠 04 問句（*Did we form a deliverable capability?*）擋下。
3. **正式 2t 仍欠**：若要关闭 2t-A checkbox，建議下次用 kit §2t 模板补一次 orchestrator／量测表签名；本檔只作 observational seed。

## 非目標（本檔明確不做）

- 不寫 Capability Proposal
- 不升格 `workflow/apk-analysis` 正文
- 不註冊 glossary
- 不貼 host／package／key／token
