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
| **experimental** | 有 TTL、目的、輸入 / 輸出定義 | 到期必須重新評估：轉 maintained / 重設 TTL（附理由）/ deprecated（見 §Research 與 concept check） |
| **maintained** | 有 owner（缺 owner 屬 ownership violation，state 不變；見 §Ownership） | 無使用者 → 由 Selection 決定 deprecated |
| **deprecated** | 禁止新增使用；有退役處理責任方 | 移除 |

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
- **清理責任**：建立它的 agent，在該任務 close-out 時清理。
- **無法清理時**：必須回報**殘留路徑、未清理原因、後續處理方式**。不得因任務結束就假設 scratch 已消失。
- **留存超過任務**：有保留價值 → 轉為 Research / experimental（補 TTL、目的、輸入輸出）；沒有 → 刪除。不得以 Scratch / temporary 狀態長期存在，避免變成另一個沒有治理的工具層。
- 不要求自動清理工具；這是責任定義，不是機制。

## Ownership（最小契約）

- **單一記錄點**：capability catalog 表的 `owner` 欄位。工具檔頭可以放指向 catalog 的註解，但不另存一份 owner。
- `owner` 可以是團隊、角色或負責的工作流，不必是個人。
- 不建立 ownership registry、不接 CODEOWNERS 自動化；專案若已用 CODEOWNERS，catalog 的 owner 應與其一致。

**Owner 缺失是治理資訊不完整，不代表生命週期狀態改變。** 不得以缺少 owner 為理由自動降級 state。

| 狀況 | 處理方式 |
| --- | --- |
| maintained + 有 owner | 合規 |
| maintained + 無 owner | **Ownership violation**：state 不變，必須補上 owner |
| experimental + 無 owner | 可接受，但必須有 TTL 與實驗目的 |
| deprecated | 不要求維護 owner，但仍需有退役處理責任方 |

若確實找不到維護責任方，由 Selection Responsibility **明確決定**是否改為 deprecated，並留下理由；不是自動轉換。

## Capability catalog

每個功能分類的 README 放一張表，作為 discovery 的搜尋起點（[`discovery-contract.md`](discovery-contract.md)）與 ownership 記錄點。

- **主要識別單位是 capability，不是檔案。** 一個 capability 可以有多個 entry；同一 capability 只登記一列。
- **索引入口**：catalog 分散在多個分類 README 時，工具樹根 README 必須有指向各分類 catalog 的索引，讓 discovery 有單一起點，也避免同一 capability 被重複登記或找不到歸屬。

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

`Research` 是 Placement，`experimental` 是 Lifecycle State，**TTL 是期限，不是自動刪除指令**。

1. experimental 工具在 TTL 內**免做** concept check（實驗本來就需要隔離），但必須記錄目的與 TTL。
2. TTL 到期時，必須重新評估：concept check、reuse evidence（[`growth-rule.md`](growth-rule.md)）與後續去向。
3. 到期但尚未完成評估的工具，**不再視為有效的 experimental**；不得默默繼續使用。
4. 到期不代表必須刪除，去向三選一：轉為 maintained（補 owner）、繼續實驗（重設 TTL 並說明理由）、deprecated。

Research 是受治理的實驗空間，不是正式工具樹以外的永久避難所。

← [Back to project-tooling index](README.md)
