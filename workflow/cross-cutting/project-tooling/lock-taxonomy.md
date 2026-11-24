# Lock Taxonomy（機械鎖分類與新增規則）

機械鎖（pre-commit、lint、gate、validator）**保護架構不變式，不保護檔案數量**。本檔提供鎖的分類形狀，以及新增鎖前必答的 Review Rule。

## 適用範圍：工具治理鎖 vs 產品不變式鎖

| 類別 | 保護對象 | 套用本檔的哪些部分 |
| --- | --- | --- |
| **工具治理鎖** | 工具樹本身（工具放哪、入口多薄、是否繞過共用層、生命週期） | 八族分類 + 鎖形狀 + Review Rule + 同概念重用 |
| **產品不變式鎖**（product invariant） | 產品程式或產品文件（架構分層、命名慣例、i18n、文件完整度…） | **不**套八族分類；仍適用鎖形狀的 intent / allowlist 退場、Review Rule 與同概念重用 |

兩類分開登記、分開計數；判斷一把鎖屬於哪類，看它的 intent 保護的是工具還是產品。

## 工具治理鎖：八族

| 族 | 保護的不變式 | 對應契約 |
| --- | --- | --- |
| **layout** | 工具放在專案約定的分類樹；語言不混放 | 專案層對應 |
| **thin-entry** | Entry 只做 Entry 允許的事 | [`role-contract.md`](role-contract.md) |
| **boundary** | 不繞過共用層重造能力（例：自寫 device / network I/O） | [`role-contract.md`](role-contract.md)、[`growth-rule.md`](growth-rule.md) |
| **lifecycle** | experimental 有 TTL、temporary 不留存、deprecated 不被新增使用 | [`lifecycle-contract.md`](lifecycle-contract.md) |
| **completeness** | 產物齊全才能提交（例：分析報告缺頁不得 commit） | 專案層 |
| **drift** | 衍生資料與來源一致（例：registry 與實際內容） | 專案層 |
| **parity** | 刻意的多實作保持對等（例：多平台入口） | 先過 Review Rule，多數應改為結構消除 |
| **safety** | 不提交機密、私有識別資訊或外部工作區參照；不外洩分析手法 | 專案層（例：去敏、禁止提交清單） |

### 每把鎖的形狀

```text
scope:            掃哪些路徑 / 哪些 placement × role
forbidden:        違反什麼（可還原成 role-contract 的違規句型）；涵蓋同一違規的常見寫法
allowlist:        暫時豁免的項目，每項附退場條件
intent:           一句話：保護哪個不變式
enforcement:      版控內（每個 clone 都會執行）或只在本機（只對設定它的那台機器有效）
```

- **scope 要覆蓋所有可能的出口。** 只掃單目標目錄的 boundary 鎖，會讓同樣的違規從入口層漏出去。
- **forbidden 要涵蓋同一違規的常見寫法。** 只抓一種語法（例：只抓 `X_SERIAL = "..."`，漏掉「讀環境變數並帶寫死預設值」）等於換個寫法就繞過。
- **反向樣本測試要雙向**：至少一個「違規變體會被抓到」的樣本，以及至少一個「合法但相似的寫法不會被誤判」的樣本。
- **allowlist 必須附退場條件**（遷移完成即移除），且目標是清空。退場條件寫在該筆例外旁邊，不另放。
- **allowlist 不以逐一列舉名稱維護。** 豁免集合會隨上游新增而變動時，改用屬性、介面或標記判定；逐一列舉會讓上游每新增一項就同時打破多個測試，或讓清單只增不減。（Phase 4 T-B5，兩個來源）
- **標明強制範圍（enforcement）。** 放在不進版控目錄的鎖只對那台機器有效：換一台機器、換一個 clone 就沒有這層保護。盤點或宣稱「有機械強制」時，要分開計數版控內與只在本機的鎖。（Phase 4 C6）
- **intent 必須能還原成架構不變式。** 只描述「某目錄不能有某副檔名」而說不出保護什麼責任的鎖，是在保護檔案形狀。

## Compensating-lock Review Rule（新增鎖時必答）

> **這把鎖防的重複或錯誤，能不能從結構上消除？**

- 能消除 → 先消除結構問題；鎖只守剩下的不變式，或根本不需要。
- 不能消除 → 記下原因，再加鎖。

典型 compensating lock：入口外殼被複製成兩種語言版本，於是加一把鎖檢查兩份是否成對。鎖本身沒錯，但它在**補償**「外殼被複製」這個結構問題；如果外殼收斂成單一參數化入口，這把鎖就不需要。

累積路徑提醒：一次性工具 → 沒找到舊的再寫一支 → 變成無責任的半永久工具 → **為防再犯加鎖** → 工具、外殼、validator、allowlist、hook 一起增長。Review Rule 就是要在第 4 步停下來。

## 同概念的鎖也適用 reuse evidence

[`growth-rule.md`](growth-rule.md) 的 reuse evidence 同樣適用於鎖：**依保護的同一不變式 / 能力計數，不依鎖檔案數量**。

- 第 2 把保護同一不變式的鎖 → 評估能否合併為一把（參數化 scope）。
- 第 3 把 → 明確裁決：合併為共用檢查，或記下分開的理由。
- 名稱相似但保護不同不變式的鎖，不算同概念。

## Governance check architecture（概念，不實作）

當專案的鎖數量開始隨需求線性增加時，可朝這個形狀收斂：

| 元件 | 職責 |
| --- | --- |
| single check entry | 一個執行入口，依參數選擇要跑的檢查 |
| check registry | 宣告有哪些檢查、各自的 intent |
| scope registry | 宣告每個檢查的適用範圍 |
| result contract | 統一的結果形狀（pass / fail、違規句型、退場條件、修復動作） |

- 各專案用**自己的語言與執行環境**實作；統一的是治理概念、檢查描述與結果形狀，**不是 runner**。
- **失敗結果必須附修復動作，包含鎖自己無法執行的情況。** 政策違規要說明怎麼改；鎖因工具或環境失敗而無法判斷時（讀不到檔案、依賴版本落後、子程序失敗），也要說明這是「無法檢查」而不是「違規」，以及怎麼恢復，不能以不相關的錯誤訊息結束。（Phase 4 T-B4，兩個來源）
- 單一入口不代表必須共用跨語言 runner。共用 runner 依來源計畫 Phase 4 裁決**維持 deferred**：兩個 consumer 和 Ai-skill 都用自己的語言實作單一入口與 registry。

## 反例

- 每發現一種違規就新增一個 check 檔案，再手動加進 hook 清單。
- boundary 鎖只掃某一層，違規工具搬到另一層就通過。
- allowlist 沒有退場條件，長期成為「已知違規清單」。
- allowlist 逐一列舉名稱，上游每新增一項就要同步修改清單。
- 鎖放在不進版控的目錄，卻當成整個專案的強制機制來計數。
- 鎖自己執行失敗時，只丟出不相關的錯誤，沒說明是「無法檢查」以及怎麼恢復。

← [Back to project-tooling index](README.md)
