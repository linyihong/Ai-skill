# Lock Taxonomy（機械鎖分類與新增規則）

機械鎖（pre-commit、lint、gate、validator）**保護架構不變式，不保護檔案數量**。本檔提供鎖的分類形狀，以及新增鎖前必答的 Review Rule。

## 七族鎖

| 族 | 保護的不變式 | 對應契約 |
| --- | --- | --- |
| **layout** | 工具放在專案約定的分類樹；語言不混放 | 專案層對應 |
| **thin-entry** | Entry 只做 Entry 允許的事 | [`role-contract.md`](role-contract.md) |
| **boundary** | 不繞過共用層重造能力（例：自寫 device / network I/O） | [`role-contract.md`](role-contract.md)、[`growth-rule.md`](growth-rule.md) |
| **lifecycle** | experimental 有 TTL、temporary 不留存、deprecated 不被新增使用 | [`lifecycle-contract.md`](lifecycle-contract.md) |
| **completeness** | 產物齊全才能提交（例：分析報告缺頁不得 commit） | 專案層 |
| **drift** | 衍生資料與來源一致（例：registry 與實際內容） | 專案層 |
| **parity** | 刻意的多實作保持對等（例：多平台入口） | 先過 Review Rule，多數應改為結構消除 |

### 每把鎖的形狀

```text
scope:            掃哪些路徑 / 哪些 placement × role
forbidden:        違反什麼（可還原成 role-contract 的違規句型）
allowlist:        暫時豁免的項目，每項附退場條件
intent:           一句話：保護哪個不變式
```

- **scope 要覆蓋所有可能的出口。** 只掃單目標目錄的 boundary 鎖，會讓同樣的違規從入口層漏出去。
- **allowlist 必須附退場條件**（遷移完成即移除），且目標是清空。
- **intent 必須能還原成架構不變式。** 只描述「某目錄不能有某副檔名」而說不出保護什麼責任的鎖，是在保護檔案形狀。

## Compensating-lock Review Rule（新增鎖時必答）

> **這把鎖防的重複或錯誤，能不能從結構上消除？**

- 能消除 → 先消除結構問題；鎖只守剩下的不變式，或根本不需要。
- 不能消除 → 記下原因，再加鎖。

典型 compensating lock：入口外殼被複製成兩種語言版本，於是加一把鎖檢查兩份是否成對。鎖本身沒錯，但它在**補償**「外殼被複製」這個結構問題；如果外殼收斂成單一參數化入口，這把鎖就不需要。

累積路徑提醒：一次性工具 → 沒找到舊的再寫一支 → 變成無責任的半永久工具 → **為防再犯加鎖** → 工具、外殼、validator、allowlist、hook 一起增長。Review Rule 就是要在第 4 步停下來。

## Governance check architecture（概念，不實作）

當專案的鎖數量開始隨需求線性增加時，可朝這個形狀收斂：

| 元件 | 職責 |
| --- | --- |
| single check entry | 一個執行入口，依參數選擇要跑的檢查 |
| check registry | 宣告有哪些檢查、各自的 intent |
| scope registry | 宣告每個檢查的適用範圍 |
| result contract | 統一的結果形狀（pass / fail、違規句型、退場條件） |

- 各專案用**自己的語言與執行環境**實作；統一的是治理概念、檢查描述與結果形狀，**不是 runner**。
- 單一入口不代表必須共用跨語言 runner。是否提供共用 runner 依來源計畫 Phase 4 證據決定。

## 反例

- 每發現一種違規就新增一個 check 檔案，再手動加進 hook 清單。
- boundary 鎖只掃某一層，違規工具搬到另一層就通過。
- allowlist 沒有退場條件，長期成為「已知違規清單」。

← [Back to project-tooling index](README.md)
