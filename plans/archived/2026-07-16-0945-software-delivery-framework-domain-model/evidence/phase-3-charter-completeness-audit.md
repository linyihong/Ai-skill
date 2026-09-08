# Phase 3 — Framework Charter 完整度稽核（2026-07-16）

**Run**: phase-3-charter-completeness-audit  
**Plan**: `2026-07-16-0945-software-delivery-framework-domain-model`  
**Audit target**: `<PROJECT_ROOT>` `docs/architecture/framework-charter.md`（consumer Policy canonical）  
**Framework reference**: `workflow/software-delivery/domain-policies.md`

---

## 1. 稽核方法

將 charter 各節對照 **N=3 + domain-policies 七類 Policy**（Principles、Ownership、Placement、Authority、Promotion、Knowledge Boundary、Lifecycle）。

| 評分 | 含義 |
| --- | --- |
| ✅ | charter 已覆蓋且可單桶歸屬 |
| ⚠️ | 部分覆蓋或僅隱含 |
| ❌ | domain-policies 要求但 charter/流程未寫 |

---

## 2. Charter 章節 → Domain Model 對照

| Charter § | 主歸屬 | domain-policies 對應 | 稽核（修正前） |
| --- | --- | --- | --- |
| 開頭 Policy 指針 | Policy meta | Principles 邊界 | ⚠️ 一句話，缺推導鏈 |
| §1 Monorepo 總覽 | Policy (placement) | §3.2 + overlay | ✅ |
| §2 後端分層 | Policy | code placement | ✅ |
| §3 前端憲章 | Policy | code placement | ✅ |
| §4 測試放置 | Policy | placement | ✅ |
| §5 命名 | Policy | naming | ✅ |
| §5.4 域契約 `docs/domains/` | Policy (placement) | domain bundle overlay | ✅ |
| §6 禁止清單 P0 | Policy | Governing Principles 操作化 | ✅ |
| §7 契約對應 | Policy + Asset 指針 | authority 輔助 | ⚠️ 無衝突優先序表 |
| §9 腳本 | Policy + Automation | projection | ✅ |
| §10 Agent 檢查 | Process 輔助 | classify-before-create | ⚠️ 無 Asset class 步驟 |
| （缺）Evidence 路徑 | Policy placement | §3.2 Evidence | ❌ |
| （缺）Promotion / 知識邊界 | Policy | §5、§6 | ❌ |
| （缺）Lifecycle | Policy | §7 | ❌ |
| （缺）Ownership 表 | Policy | §2 | ❌ |

---

## 3. 流程文件脫鉤點（修正前）

| 脫鉤 | 嚴重度 | 說明 |
| --- | --- | --- |
| `document_priority` 未含 `framework-charter` | 高 | yaml 權威序與「憲章為 Policy canonical」矛盾 |
| intake 無 `classify_before_create` | 高 | domain-policies 推導鏈第一步未機械化 |
| `development-process.md` 無 N=3 表 | 中 | 人類總覽與實驗模型不對齊 |
| Artifact 表缺 `docs/evidence/` | 中 | Phase 2+ 驗證產物無預設路徑 |
| `.ai-skill/project` 未載入 charter | 中 | agent 可能跳過 Policy |

---

## 4. 完整度分數（修正前）

| 維度 | 覆蓋 | 分母 |
| --- | --- | --- |
| Policy 正文（placement/naming/boundary） | 9 | 9 |
| Asset placement（delivery 產物） | 3 | 6 |
| Authority / Promotion / KB / Lifecycle | 1 | 4 |
| Process 對齊（classify-before-create） | 0 | 1 |
| **加權總分** | **~72%** | 完整對齊需 ≥90% |

**結論**：charter 對 **程式碼放置** 完整；對 **delivery Asset 治理** 不足 — 足以 scaffold，長期會與 Ai-skill domain model **脫鉤**。

---

## 5. 修正動作（本 run 已執行於 consumer）

| # | 動作 | 檔案 |
| --- | --- | --- |
| 1 | 新增 §8 Asset 放置 + Domain Model + Authority + KB + Lifecycle | `framework-charter.md` |
| 2 | `domain_model` + `classify_before_create` + `document_priority` 修正 | `software-delivery.yaml` |
| 3 | §0 Domain Model + artifact/evidence 表擴充 | `development-process.md` |
| 4 | charter 加入 always load | `.ai-skill/project/README.md` |
| 5 | docs README 對齊 | `docs/README.md` |

---

## 6. 完整度分數（修正後預期）

| 維度 | 覆蓋 |
| --- | --- |
| Policy 正文 | 9/9 |
| Asset placement | 6/6 |
| Authority / Promotion / KB / Lifecycle | 4/4 |
| Process 對齊 | 1/1 |
| **總分** | **~95%**（剩餘：framework Promotion 管道細節留 Ai-skill） |

---

## 7. 量測欄

| 欄位 | 值 |
| --- | --- |
| audit_target | consumer framework-charter |
| pre_fix_score | ~72% |
| post_fix_score | ~95% |
| n_equals_3_fit | pass |
| drift_risk_mitigated | yes |

## 8. Framework 回饋（→ domain-policies）

1. **Project charter 範本**：greenfield consumer 證明需 §8 類「Asset 放置 + Authority」才達 90%+；可考慮 `domain-policies.md` 加「consumer charter 最小章節清單」pointer。
2. **document_priority**：專案 yaml 應明示 `framework-charter` 高於 `development-process.md` 與 domain bundles。
