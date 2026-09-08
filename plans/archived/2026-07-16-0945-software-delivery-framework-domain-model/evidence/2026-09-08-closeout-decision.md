# Domain Model 結案裁決

Run ID: sd-domain-model-closeout-decision-2026-09-08

## Pre-build Interrogation / Architecture Compatibility Preflight

| 欄位 | 結論 |
| --- | --- |
| Trigger | 使用者要求補讀兩個 consumer 後收尾指定計畫 |
| Goal / acceptance | 複核原 acceptance、補可回查證據、處置 OQ 與選做項目、歸檔 |
| Scope | 計畫／evidence／inbound links；README 補充 intent 表示邊界 |
| Non-goals | 修改 consumer、新增全域 gate、執行產品驗收、升格 ADR、收其他計畫 |
| Checked sources | plans README、原 plan／evidence、SD README／Domain Policies／intake／contracts／artifact-gates／execution-flow、glossary、registry、runtime owner-layer contract、consumer Policy／Process／checker |
| Conflicts | 早期 Phase 0「唯一授權」已過時；若干 relative links 已錯位；Intent 不持久化的早期理由過強。更新 current disposition，保留歷史證據 |
| Source of truth | Policy 留 workflow；glossary 留 knowledge；plan evidence 不取代 canonical；生成索引用 runtime refresh |
| Duplication risk | 不建立第二份 taxonomy／runtime source；Automation 僅做設計裁決 |
| OQ 核對 | OQ-1–5 全部 resolved；選做 generic automation 明確 deferred，見主計畫 |
| Decision | proceed：依原 scope 完成，不需擴充模型或新 gate |
| Validation | 來源 hash、局部 checker、Markdown links、runtime refresh/validate、plan validation、staged diff、hooks 與 push readback |

## Phase 3b 處置

| 項目 | 裁決 | 原因／重啟條件 |
| --- | --- | --- |
| Asset class → required policy fields | 評估完成；generic implementation deferred | owner-layer YAML 可以描述；需先有可跨 overlay 的 asset declaration 與誤判案例，不能用固定目錄猜 class |
| 新 enforcement rule_class | 本輪不新增 | 本輪沒有新 executor 或強制行為；現有 plan_governance 管 lifecycle。consumer gate 不能冒充 Ai-skill coverage |
| Test-first | 原則保留；新 gate/scenario implementation deferred | 真實需求重啟時先定 positive／negative／overlay cases，再實作 executor；本輪無新機械行為需鏡像測試 |

Deferred owner：software-delivery workflow maintainer。追蹤位置：主計畫 §3b（永久保存）。
重啟入口：正常開發出現無法由既有 Policy／overlay 解決的重複分類或 placement 失效，
且有明確 asset declaration、可觀察驗收與 executor consumer 時，另開有界 implementation plan。
不設自動重啟，不把每個新 Asset 當成增加 gate 的理由。

## 原 acceptance 與 promotion

原七條 acceptance 依 Phase 0/2 歷史證據、canonical pointer 讀回與本輪兩個 consumer 抽樣成立。
新增觀察不回填為當年的量測結果；歷史 95% 只指原 40 概念矩陣。

ADR：不 promotion。模型仍可演化，README／Policy／glossary 是更輕且足夠的落點；
本次不把專案個案變成平台不可逆 invariant，也不自動提升 glossary candidate 狀態。

Runtime：沒有新增 route、target_key 或機械 gate。現有 SD 執行／產出 contract 保持原狀；
完成標記只描述此計畫交付。Phase 3b 的通用 automation 並未完成。

## Linked updates

- 主計畫與 evidence 目錄整體移到 archived，Run 索引同步。
- plans/README 新增結案列；SD README／Domain Policies／execution-flow 與 glossary 引用改新位置。
- 另一 active plan 的歷史 evidence 僅修 inbound link，不改其狀態或裁決。
- 已檢查 intake、contracts、artifact-gates 與 registry，無行為改動，無需改契約或 coverage。
- runtime refresh 更新生成索引；reference-first，tool mirror sync 不適用。

## 驗證狀態

來源抽樣與 consumer 局部檢查已通過。Repository 檢查：

- 15 份 touched／歸檔 Markdown，282 個本地檔案連結：0 missing（不含 heading anchor 語意檢查）。
- `plans validate`：45 plans、0 findings、0 blocking。
- `runtime refresh` 與 `runtime validate`：exit 0；現有 runtime.db age／orphan warning 不屬本次新增 gate 的完成證據。
- `git diff --check`：PASS；commit hooks 於正式提交時執行。
- `execution-flow.md` 只修證據／導航連結；YAML 行為不變，使用有理由的 markdown-only 同步例外。
- Touched 索引／導航內既有失效連結同步修復；未變更其它計畫的實際狀態。
Git commit/push 的真實結果以 git refs 與最終回報為準，不以此段提前宣稱。
