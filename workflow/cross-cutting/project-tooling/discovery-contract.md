# Discovery Contract（先找再建）

解決的問題：agent 不是故意重複造輪，而是**沒有先找**，或找到了卻沒辨識出是同一件事。本契約是建立新工具前的**行為不變式**，不是表單。

> 狀態：behavioral invariant（pilot）。是否機械化 deferred，於來源計畫 Phase 4 依 dogfood 證據重新評估。

## 建立新工具之前

依序完成，任何一步已能解決就停：

1. **搜尋專案既有工具與 capability catalog**（各功能分類 README 的能力表）。
2. **搜尋 Domain / Shared 是否已有相同能力**——以能力 / 責任搜尋，不只以檔名或關鍵字。
3. **確認能否以參數解決**：既有工具加參數、加 config 值、加 strategy 選項。
4. **確認能否擴充既有 harness**：共用 Logic 已存在時，只長一個只放差異的 sibling。
5. **仍需新增時**，給出 creation justification，並決定 Placement 與 Lifecycle State。

「同概念」的判定見 [`growth-rule.md`](growth-rule.md) §Concept check。

## 合理新增的理由（只有四種）

| 理由 | 例 |
| --- | --- |
| 新的執行責任 | 既有工具都不承擔這個責任 |
| 新的資料來源或介面 | 同樣的分析，但輸入來自全新的系統或協定 |
| 現有抽象無法合理涵蓋 | 硬塞會讓既有工具的 Parameters 變成邏輯藏身處 |
| 實驗需要隔離 | 尚未確認可行，不應污染正式工具 → Research |

「這次比較快」「找不到」「舊的不好讀」不是理由；後兩者應該記為既有工具的問題（catalog 缺漏或可讀性），而不是另起爐灶。

## Creation justification（輕量寫法）

一兩句，放在**工具檔頭註解**或**新增它的 commit message**，擇一即可：

```text
New: <能力 / 責任>。Why new: <四種理由之一 + 一句說明>。Placement/State: <例：Research / experimental, TTL 2026-12-31>。
```

不需要另開文件、不需要審批表。重點是讓下一個 agent 能在搜尋時看到「這個工具為什麼存在」。

## 與其他責任的關係

| 責任 | 本契約中的角色 |
| --- | --- |
| Selection | agent 自己完成上方 5 步並留下 justification |
| Constraint | 專案可用鎖檢查「新工具是否登記 catalog」等**已明確違反**的不變式；不檢查 justification 寫得好不好 |
| Verification | review / verifier 檢查是否有同概念的既有能力被忽略 |

**明確不做**：在 agent 寫工具前插入強制問卷或審批流程；自動判定「是否同概念」並自動合併。

## 反例

- 任務需要解析一份日誌，agent 直接寫一支新腳本，而專案的分析分類下已有同責任的解析工具，只差欄位設定。
- 第三次需要相同的效能量測流程，仍各自複製計時、環境檢查、報告寫檔，而不是抽一個 harness。
- 為了避免「新增腳本」，把一段全新的執行責任硬塞進既有工具的設定檔。

← [Back to project-tooling index](README.md)
