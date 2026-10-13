# Grandfather sunset — administrative closeout

Run ID: grandfather-sunset-closeout-2026-09-09

## Preflight and decision

使用者要求收尾其他計畫。本次沿用 [原調查](../02-grandfather-sunset-audit.md) 與現行
[canonical sunset rule](../../../../governance/lifecycle/system-upgrade-governance.yaml)，
重跑 `ai-skill runtime audit --json` 並讀回 registry、discovery、validator 與 scenario。

裁決：四份已 archived 計畫的 grandfather 豁免 **sunsetted**，effective date 保留 **2026-08-31**，
本輪實際 reviewed date **2026-09-09**。沒有補造準時完成紀錄，沒有延長豁免。
本輪僅處置既有 metadata／索引與證據說明；未新增 gate、變更 route trigger 或 registry coverage class。

## Declared inventory readback

| Declared item | Current audit classification | Audit basis |
| --- | --- | --- |
| route.governance.cognitive-state-evidence | auto-detected | discovery signal description references route |
| enforcement.evidence_hierarchy.contract | consumed | Go source references target_key |
| route.memory.retrieval-activation | auto-detected | discovery signal description references route |
| route.models.model-aware-routing | auto-detected | discovery signal description references route |
| route.runtime.cognitive-modes | intentionally-manual | validators_consume_by_file_path_not_route_lookup annotation |

五個 declared items 均滿足原規則 (a) 的分類條件；無需刪除 route 或 surface。
分類只是 audit 既定靜態判準，不等同 runtime observed、路由真實命中率或使用者行為改善。

## Extension checkpoint

- Audit 引入 commit `0f53e91d`，日期 2026-05-28；至 2026-08-15 為 79 天，非少於 60 天。
- High-priority wires commit `f222fdd4`，日期 2026-05-28；[原 Phase 4](../../../archived/2026-05-28-1200-gen3-runtime-trigger-audit-and-completion.md) 記錄 5/5 完成；本輪分類與其 after 欄一致。
- 既有 07-08 調查也記錄相同分類。這些是歷史證據讀回，**不是**宣稱本 session 曾在 08-15 執行 audit；未找到觸發延展的紀錄。
- 因此不啟用 2026-11-30 conditional extension；08-31 effective sunset 保留。

## Scenario limitation and disposition

[pre-2026-grandfather-coverage-v1](../../../../validation/scenarios/failure-derived/pre-2026-grandfather-coverage-v1.yaml)
原文聲稱 audit 會檢查 coverage diff／deadline，並有 Go fixture 覆蓋。
本輪讀到的實際 binding 位於 `scripts/ai-skill-cli/internal/audit/scenarios_stub_test.go`：
`TestScenarioStubsBound` 只驗五個 ID 非空，**不驗** sunset、coverage 或 extension。
因此這些原先的完整 executable-coverage 宣稱不成立；scenario 改為明示 **agent-judged** readback。

本轮人工核對：四個歷史 archive path 存在、五個 declared items 分類、日期條件、豁免狀態與 plans index 一致。
反向檢查：僅 ID stub PASS 不足判 coverage PASS；缺 item／orphan／無理由延展應判 review FAIL。
未新增 mutation test 或 runtime blocker；如未來要強制此檢查，由 final-form plan Phase 7 的 scenario runner 工作評估。
本項是明確縮限證據宣稱，**不是**把其他 orphan scenarios 算作清零。

## Linked updates and boundary

- YAML 保留 covered_plans 作歷史 inventory，加 effective/review date 與 evidence pointer。
- Routing note 改歷史敘述，manual_activation reason 保留。
- plans index 四個過時警告更新；final-form Phase 8 僅勾本 grandfather 子項，並記錄 late closure。
- 本 DVA 主計畫的 Q5–Q9、2t 與獨立驗證條件未因此完成，DVA 主計畫繼續 active。
- Runtime compile／refresh／validate 與最终 git refs 以本輪命令驗證；沒有整體 orphan 清零宣稱。
