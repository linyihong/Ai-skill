> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-23 - Unity iface analysis techniques stay in Ai-skill; HTML report scripts stay in the business repo

Status: validated

#### Reuse Evidence

使用者糾正：要分享到 Ai-skill 的是 **Unity 分析技巧**；`interface-analysis` HTML 生產器屬於業務專案。隨後在業務 repo 補上 `generate_interface_analysis.py`（對齊 buffalo 薄骨架），並把本 lesson／`unity-asset-parse-flow.md` 改成明確邊界，避免把專案 CLI 寫成 Ai-skill 義務。

#### One-line Summary

Ai-skill 只沉澱 UnityCache／UnityPy／粒子成熟度／seated-vs-splash 等**分析技巧**；§3e HTML／atlas CLI 留在業務 `scripts/`，缺了就在業務 repo 補生產器，不要把專案腳本路徑 promote 成通用 workflow 清單。

#### Human Explanation

先前誤把「業務 onboarding 的 Pass-5 腳本表」寫進 Ai-skill，好像那些 `.py` 是跨專案標準。正確分層：

| 層 | 放什麼 |
| --- | --- |
| **Ai-skill** | UnityPy 盤點、ParticleSystem host／material、成熟度標籤、進桌優先於 splash、勿把進桌當 hero |
| **業務 repo** | `generate_interface_analysis.py`、`generate_symbol_map.py`、atlas／resolve／merge／thumbs、具體 HTML 契約 |

缺 HTML 生產器時：在業務 repo 依既有參考包（薄骨架／深度手寫）補腳本並 regenerate，**不是**在 Ai-skill 假裝已有一鍵工具。

#### Trigger

- 使用者說「技巧是 Unity 分析；HTML 腳本是業務專案的」。
- Agent 把專案 `scripts/slots/*.py` 路徑寫進 Ai-skill workflow 當必跑清單。
- §3e 交件是手搓淺色 stub，與參考包視覺不一致，且業務其實缺／漏跑 HTML 生成器。

#### Evidence

- Tool: 業務 `generate_interface_analysis.py`（對齊 buffalo dark shell）；Ai-skill workflow 邊界改寫。
- Sanitized excerpt: Unity techniques ≠ project HTML CLI；missing generator → add in business repo from reference packages.
- Evidence path: `<AI_SKILL_REPO>/analysis/apk/workflows/unity-asset-parse-flow.md`；業務 `<PROJECT_ROOT>/scripts/slots/generate_interface_analysis.py`（名稱因專案而異）。

#### Generalized Lesson

1. **Feedback 到 Ai-skill = 可重用分析判斷**，不是拷貝專案腳本 API。
2. **報告 HTML／驗證 CLI 屬專案**：缺了就在專案補；參考薄／深兩類既有頁面。
3. **Unity 技巧可跨專案**：UnityPy counts、粒子參數、成熟度、seated batch（見同目錄 101000／102500）。
4. **契約欄位（inheritsShared 等）**若是專案 onboarding 硬規則，寫在專案 docs；Ai-skill 只抽象「分開 shared chrome vs cabinet delta」。

#### Agent Action

寫 Ai-skill lesson 前先問：這是 Unity／Frida **方法**，還是專案 **腳本路徑**？方法→feedback／workflow；腳本→業務 repo。§3e 交件前查業務是否已有 HTML 生成器；沒有就補，有就跑，禁止無契約 stub。

#### Goal / Action / Validation

- Goal: 技巧分層正確；業務有可用的 §3e HTML 生產器。
- Action: 修正本 lesson + workflow 邊界；業務新增／執行 `generate_interface_analysis.py`。
- Validation: Ai-skill 正文無強制專案腳本檔名清單；業務 `--check` 對生成 HTML 通過；參考包可 `handMaintained`。

#### Applies When

- Unity IL2CPP／UnityCache UI／粒子分析要沉澱到 Ai-skill，且業務另有報告殼。

#### Does Not Apply When

- 純業務腳本 bugfix，無跨專案方法可抽。

#### Validation

新 lesson／workflow 若出現具體 `scripts/slots/<name>.py` 當「全庫必跑」，視為邊界違規，改為「執行業務專案既有報告腳本」。

#### Promotion Target

- [`analysis/apk/workflows/unity-asset-parse-flow.md`](../../../analysis/apk/workflows/unity-asset-parse-flow.md)（Unity vs 業務腳本邊界）
- 業務 onboarding／`scripts/` README（HTML 生產器；非 Ai-skill 路徑）

#### Promotion Record

2026-10-23：workflow 改為邊界說明；業務補 `generate_interface_analysis.py` 並 regenerate 目標機台。本檔由「無一鍵生成器」改寫為「分層＋缺則在業務補」。

#### Required Linked Updates

- Updated: 本 lesson；`unity-asset-parse-flow.md`；`feedback/history/apk-analysis/unity-il2cpp/README.md`
- Updated (business repo): `scripts/slots/generate_interface_analysis.py`、onboarding §3e、目標機台 HTML
- Not applicable: enforcement failure-pattern（一次分層誤寫，lesson 已吸收）
