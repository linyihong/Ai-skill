# Profile — Slot RTP

第一個 consumer profile。只放**詞彙、指標定義慣例與 slot 特有的診斷注意事項**；個案結果、校準快照、參數值與探針程式碼留在 consumer repo。

> 狀態：N=1 consumer。本檔內容是 profile 慣例，不是通用規則。

## Target 結構（C1）

| 角色 | slot 中通常是 |
| --- | --- |
| `normative` | 規格或公告的 RTP |
| `reference_observation` | 受信任穩態帳號的官方採集樣本（附樣本範圍與 SE） |
| `implementation` | 產品實作要逼近的值 |

熱帳號／農場帳號樣本的**絕對值**不得當 Target（常有旋轉數遞減等採集偏差），只取尺度不變結構：條件機率、跨軸相關、形狀占比。

## 指標詞彙（I2 Metric Parity）

逐一核對兩側定義，基準選錯常導致整輪重做：

| 指標 | 必須對齊的定義點 |
| --- | --- |
| RTP | 回報÷注額；基準回合類型；是否含系列玩法 |
| 命中率 | 任一中獎 vs 指定賠付；以回合或以線計 |
| 均值／注額 | 注額縮放參數是否一致 |
| base vs feature 拆分 | 免費轉、respin 等系列玩法是否拆開計 |
| 盤面範圍 | 可見格 vs 整盤（含不可見位置） |
| 賠付分檔 | 線賠與倍數分檔的歸類 |

## 典型 tier 配置（建議，由 contract 決定）

| 指標 | 常見 tier |
| --- | --- |
| 總 RTP | `hard_constraint` |
| base／feature RTP 拆分、命中率 | `primary_objective` |
| 獎勵形狀占比、跨軸相關 | `primary_objective` 或 `diagnostic_metric` |
| 罕見玩法觸發率 | `diagnostic_metric`（樣本常不足，判斷看事件數而非手數） |

## Slot 特有診斷注意事項（candidate-intelligence, N=1）

- **跨軸相關**：每軸獨立抽樣的結構無法產生跨軸相關；先依 C7 證明結構天花板，再換耦合式抽樣結構，不要無限調權重。
- **單旋鈕副作用**：增加同類聚集常同步抬高 RTP；需要把「類別密度聚集」與「同類聚集」分成正交旋鈕，且每調一個就量全部指標。
- **抽樣策略**：條帶式、加權格位、耦合條帶等策略應做成「參數化資料 + 註冊式 Strategy」；重構時以其他 consumer 的釘值作零漂移護欄。

以上三條待第二個真實 consumer 驗證後才可能升級。
