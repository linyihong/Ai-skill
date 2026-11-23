# Growth Rule（同概念怎麼長）

新工具可以長，但只在**確認沒有同概念能力**時才長。同概念時，加參數，或抽出共用邏輯後在同一功能分類下長一個 sibling。

## Concept check（Selection Responsibility）

查 capability catalog（[`lifecycle-contract.md`](lifecycle-contract.md) §Capability catalog），回答「是否已有同概念的能力？」

| 判定 | 允許的動作 | 禁止 |
| --- | --- | --- |
| 同概念、只差數值 / 目標 | 在既有工具加參數，或加 config / strategy 值 | 複製一個新工具 |
| 同概念、流程有差異 | 先把共同 Logic 抽到該分類的共用模組，再在**同一功能分類下**長 sibling（只放差異） | 在別的分類或根目錄另起爐灶；sibling 重寫共用 Logic |
| 新概念（catalog 無對應） | 開新功能分類 + 新工具，並登記 catalog | 不登記 catalog |

### 什麼叫同概念

**概念 = 能力或責任，不是相似的程式碼。**

- 判定準則：**產出物**與**核心責任**兩者都相同，才是同概念。只差目標 ID、座標、門檻、站點等，屬參數。
- 兩個工具都有迴圈、讀 JSON、輸出報告，不代表同概念。
- 程式碼完全不同，但承擔相同責任，可能就是同概念。
- 只有產出物相同、責任不同（例：同一批資產的擷取、去背、匯出、同步各是 pipeline 的一段），是**不同概念**。

此準則在兩個 consumer 的 5 個工具群試判中都判得出來，見來源計畫 [`evidence/phase0-baseline.md`](../../../plans/active/2026-11-23-0834-project-tooling-governance/evidence/phase0-baseline.md) §Concept 試判。

## Config-first ≠ Config-everything

| 差異型態 | 落點 |
| --- | --- |
| 流程相同、資料不同 | Config（Parameters） |
| 流程部分不同、核心契約相同 | Strategy（例：多種 OCR 引擎共用輸入 / 輸出契約，各自實作） |
| 執行責任不同 | 獨立模組或工具（新概念） |
| 實驗尚未穩定 | Research |

不要為了避免新增工具，把所有差異塞進一個超大設定檔；那只是把 script sprawl 換成 configuration sprawl（見 [`role-contract.md`](role-contract.md) §Parameters 不是邏輯藏身處）。

## Reuse evidence（Governance Constraint）

計數對象是**同一能力 / 責任的出現次數**，不是相似程式碼。

| 次數 | 必須採取的行動 |
| --- | --- |
| 第 1 次 | 依概念決定 Placement |
| 第 2 次 | 進行重用與抽象化評估（走上方 concept check 表） |
| 第 3 次 | 明確裁決：抽取共用能力，或寫下保留重複的理由 |
| 第 2 個領域使用 | 評估是否上推 Shared |

分工：concept check 決定「是不是同一件事、該怎麼長」；reuse evidence 確保抽象化決策**不能無限期逃避**。觸發的是決策，不是強制合併——表面相同不代表未來變化方向相同。

「保留重複的理由」寫在 catalog 該能力列的備註，或相關 commit message；一句即可。

## Promotion ladder（抽共用時的落點）

沿 [`lifecycle-contract.md`](lifecycle-contract.md) 的 Placement 上推：`Scratch → Research → Target → Domain → Shared`，數值差異落在 Parameters / Strategy。

- 每階可退回。
- 共用 Logic 抽到「涵蓋所有使用者的最低一層」即可，不預先上推到 Shared。
- 上推只改 Placement；Role 結構（Entry / Logic / Parameters）照 [`role-contract.md`](role-contract.md) 重新檢查。

## 反例

- 第三個目標需要相同的「旋轉直到事件並擷取」流程，又複製一支完整腳本，只改了座標與事件條件。
- 8 個效能案例各自重寫計時、環境斷言與報告輸出，0 個共用模組。
- 為了「共用」，把三個責任不同的 pipeline 階段合併成一個工具，再用設定旗標切換。

← [Back to project-tooling index](README.md)
