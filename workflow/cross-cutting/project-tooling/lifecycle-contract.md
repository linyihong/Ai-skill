# Lifecycle Contract（服務範圍與維護狀態）

定義兩個獨立維度：**Placement**（服務範圍）與 **Lifecycle State**（維護狀態）。兩者也都與 [`role-contract.md`](role-contract.md) 的 Internal Role 獨立。

## Placement（服務範圍）

| 值 | 服務對象 | 進入條件 |
| --- | --- | --- |
| **Scratch** | 只為當前任務 | 預設：任務內臨時需要 |
| **Research** | 一個尚未確認的假設 | 需要重複執行、但通用性未確認 |
| **Target** | 單一目標 / 案例 | 確認要保留，只服務一個目標 |
| **Domain** | 同一領域的多個目標 | 同領域出現同概念的重用（見 [`growth-rule.md`](growth-rule.md)） |
| **Shared** | 跨領域 | 第 2 個領域使用同一能力，且評估後上推 |

Placement 是可上推、也可退回的落點，不是必經階梯：Research 不一定升到 Target；共用邏輯只抽到「涵蓋所有使用者的最低一層」。

## Lifecycle State（維護狀態）

| 值 | 存在條件 | 到期 / 退出 |
| --- | --- | --- |
| **temporary** | 只為當前任務 | 任務結束清理 |
| **experimental** | 有 TTL、目的、輸入 / 輸出定義 | 到期必須裁決：升級（走 concept check）/ 延期（附理由與新 TTL）/ 刪除 |
| **maintained** | 有 owner | 無使用者 → deprecated |
| **deprecated** | 禁止新增使用 | 移除 |

### 組合例

| 工具 | Placement | State |
| --- | --- | --- |
| 臨時資料轉換 | Scratch | temporary |
| 尚未確認的實驗探針 | Research | experimental |
| 正式單目標工具 | Target | maintained |
| 領域共用工具 | Domain | maintained |
| 已淘汰的共用工具 | Shared | deprecated |

原則：**不是每個工具都要升級，但每個被保留的工具都要有存在理由。**「成功執行過」不等於「值得永久保留」。

## Scratch 位置與清理責任

- **位置**：專案內一個**不進版控**的固定目錄（由專案在 README 指定，並列入 ignore / exclude），或 agent 工具提供的 session 暫存區。**不得**放在正式工具樹內。
- **清理責任**：建立它的 agent，在該任務 close-out 時清理。close-out 回報須列出仍留存的 scratch 產物（若有）。
- **留存超過任務**：要嘛轉為 Research / experimental（補 TTL、目的、輸入輸出），要嘛刪除；不得以 Scratch / temporary 狀態長期存在。
- 不要求自動清理工具；這是責任定義，不是機制。

## Ownership（最小契約）

- **單一記錄點**：capability catalog 表的 `owner` 欄位。工具檔頭可以放指向 catalog 的註解，但不另存一份 owner。
- `owner` 可以是團隊、角色或負責的工作流，不必是個人。
- `maintained` 必須有 owner；沒有 owner 的工具視為 experimental（需補 TTL）或 deprecated。
- 不建立 ownership registry、不接 CODEOWNERS 自動化；專案若已用 CODEOWNERS，catalog 的 owner 應與其一致。

## Capability catalog

每個功能分類的 README 放一張表，作為 discovery 的搜尋起點（[`discovery-contract.md`](discovery-contract.md)）與 ownership 記錄點：

| 欄位 | 內容 |
| --- | --- |
| capability | 能力 / 責任（一句話；用於 concept check，不寫檔名） |
| entry | 呼叫入口 |
| placement | Scratch 以外的 placement |
| state | experimental / maintained / deprecated |
| owner | 見 §Ownership |
| ttl | 僅 experimental 必填 |

- Scratch / temporary 工具不登記。
- 機器可讀版本（YAML / JSON）deferred：README 表格足以支撐 discovery；只有當專案要用鎖檢查「新工具是否登記」時才需要，屆時與 Discovery 機械化一起評估。

## Research 與 concept check

experimental 工具在 TTL 內**免做** concept check（實驗本來就需要隔離），但必須宣告 TTL 與目的；TTL 到期的裁決即是它的 concept check 時點。

← [Back to project-tooling index](README.md)
