# Dogfood 2 — multi-tenant payment API, deep audit + invalidation replay（`<CONSUMER_PROJECT_2>`）

**Date**: 2026-10-03 · **Mode**: `deep` · **verdict_kind**: `expected` · **Roles**: auditor = main session; verifier = fresh subagent（只讀、不繼承推理）

## Target

`<CONSUMER_PROJECT_2>`：多租戶支付 API。範圍是共用授權 handler 與 policy、claims 推導（Host / Tenant / 使用者類型）、租戶解析與全域查詢過濾器、JWT、匿名的註冊 / 登入 / 驗證碼 / refresh 端點、一個支付閘道的存款回呼，以及餵給這些 control 的 tracked 設定。

選擇原因：授權面最豐富（同一組 JWT 兩類使用者、Host header 決定租戶、共用 handler），且有一個只改共用 control、不改任何 controller 的歷史 commit，可做 coverage invalidation 的正向回放。

## 鏈路是否走通

| 步驟 | 結果 |
| --- | --- |
| invoke `security-audit`（deep） | 初稿 5 findings → verifier 後 12 findings（全部 `needs_validation`；無可重現執行環境） |
| `audit_execution` | `completed`；not_covered 經 verifier 修正 |
| Coverage ledger（專案端 YAML） | 初稿 8 units / 7 controls → 12 units / 13 controls |
| Fresh verifier | 主動嘗試推翻最嚴重的 finding 失敗；指出 1 個 over-stated、2 個 under-stated、1 個證據路徑錯誤、7 個遺漏（其中一個是 tracked 設定含簽章金鑰）、not_covered 3 處不誠實 |
| 仲裁 | 全部 `fix`；orchestrator 自行查證 4 個關鍵主張（secret 只看 key 名稱，不輸出值） |
| Gate（expected） | **block**：3 個 needs_validation 且 potential_impact critical / high、無 human_review。首次在真實任務走到 block 路徑 |

## Coverage invalidation replay（Phase 4 必要項）

以 HEAD 定義的 ledger，對兩段真實歷史機械回放（解析 ledger + `git diff --name-only`，可重跑腳本留在專案端）：

| Diff | 結果 |
| --- | --- |
| 拆分使用者模型的歷史 commit（改了 handler、claims 推導、租戶解析、查詢過濾器、JWT 簽發；**0 個 controller**） | 7 個 control 被觸及 → 12 個 units 中 10 個 `needs_revalidation`；未失效 2 個（與這些 control 無關的 withdraw 與 secrets 設定） |
| 本次同步的 9 個前端 / 文件 commit | 0 個 control、0 個 unit（負向對照正確） |

這是**事後回放**：units 以 HEAD 定義，回答「若 ledger 當時存在，這個 diff 會讓哪些證據失效」。它證明 invalidation 規則在真實 diff 上可機械執行、且能抓到檔案 hash 規則抓不到的情況；它不是在真實開發流程中被動觸發的失效。

## 這次 dogfood 證明的事

1. **檔案 hash 規則確實不夠**：最關鍵的授權變更沒有碰任何 controller；以端點檔案判斷會讓全部覆蓋維持有效，control 依賴則抓到 10 / 12。
2. **目錄型 `source_scope` 必須依路徑邊界比對**：第一版回放用字串前綴，一個目錄範圍誤中同層、名稱以該目錄名開頭的檔案。改成「完全相同或後接 `/`」才正確。
3. **檔案 / 目錄粒度會 over-invalidate**：同一檔案承載兩個 control（租戶解析與 client IP）時，只改其中一部分也會讓兩邊的 units 失效；目錄範圍（entity 目錄）會牽動看似無關的 unit。依「寧可多失效」可接受，但成本會隨 ledger 變大。
4. **Verifier 的對抗價值在 deep 模式更明顯**：稽核者漏掉 tracked 設定裡的簽章金鑰（潛在 critical）與驗證碼依設定放行，並把它們放進 not_covered 的「範圍外」項目。這是 not_covered 被用來排除範圍內缺陷的實例。
5. **Placeholder 實作要當成 finding 線索**：一個「簽章驗證占位，正式上線前請替換」的類別被正式回呼路徑使用，而且根本沒有被呼叫。
6. **Over-stated 也會被抓**：一個 privilege-escalation finding 的 policy 在 HEAD 沒有任何端點使用，verifier 將其降為 latent。

## 寫回 Ai-skill 的修正

- Coverage ledger 契約：目錄 `source_scope` 依路徑邊界比對；新增「機械回放」程序與粒度取捨說明。
- Finding list template：not_covered 不能用來排除範圍內的缺陷；餵給受稽核 control 的 tracked 設定（金鑰、開關）屬於範圍內；placeholder / stub 實作列為必查線索。

## 尚未證明

- `enforced` verdict（Phase 5）。
- 在真實開發流程中（非回放）由新 diff 被動觸發失效並完成重驗。
- Template schema 在不需修改欄位的情況下再跑一輪（Phase 5 entry condition，見 plan）。

## 去敏

本檔不含專案名稱、類別名稱、路徑、host、人名或任何 secret 值；finding list、ledger 與回放腳本留在 consumer repo。
