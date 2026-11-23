# Role Contract（工具內部角色）

描述**一個工具內部**各段程式負責什麼。與 Placement（服務範圍）、Lifecycle State（維護狀態）是三個獨立維度，見 [`README.md`](README.md)。

## 三種角色

| Role | 允許 | 不允許 |
| --- | --- | --- |
| **Entry** | 解析參數、取得必要設定、呼叫 Logic、傳遞結果與結束狀態 | 領域邏輯；重複實作 device / network / 檔案 I/O；資料處理 |
| **Logic** | 實際執行邏輯 | 寫死應由 Parameters 提供的值（目標 ID、座標、裝置或環境識別、門檻） |
| **Parameters** | 設定值、策略選擇、per-target 數值 | 大量條件分支、流程描述、可執行邏輯 |

### 判定細則

- **幾何與數值**：點擊座標、以及任何**會隨目標調整**的參數（門檻、ID、裝置識別、等待時間）屬 Parameters。**演算法本身的幾何**（例：OCR 裁切比例、像素判斷區域）可以留在 Logic，但必須具名、集中在一處，不得散落在流程中。
- **入口編排**：Entry 只是**依序呼叫**多個 Logic 仍屬 Entry；一旦包含**條件決策、快取判斷或範圍選擇**（例：依變更範圍決定跑哪些檢查、命中快取就略過），那段編排就是 Logic，應移到 Logic，Entry 只呼叫它。

### Placement 不決定 Role

Target、Domain、Shared 的工具都可以有自己的 Entry / Logic / Parameters，但**不要求同時具備三種**。例如一個共用模組可以只有 Logic，由各 Target 的 Entry 呼叫。

### Entry 維持薄入口

Entry 是最容易長胖的地方：agent 為了「先跑起來」常把整段流程寫在入口檔。判斷方式不看行數門檻本身，而看是否出現 Entry 不允許的職責；行數只是篩選訊號。

### Parameters 不是邏輯藏身處

Config-first 不等於 config-everything。當設定開始出現「如果 A 則做 B」的分支、或描述步驟順序，就是邏輯外流到 Parameters，應回到 Logic，或改用 Strategy（見 [`growth-rule.md`](growth-rule.md) §Config-first ≠ Config-everything）。

## 違規判定

### 描述格式

```text
<placement> 的 <應有 role> 放在 <實際 role 或位置>
```

例：

- 「Target 的 Logic 放在 Entry」——單一目標的旋轉 / 擷取流程整段寫在入口檔。
- 「Shared 的 Parameters 放在 Logic」——共用裝置模組內寫死某台裝置的識別值。
- 「Domain 的 Logic 放在 Parameters」——設定檔用條件欄位描述流程分支。

此格式**不依賴任何語言或目錄結構**，同一句話可用在 Python、Node、PowerShell 或 SQL 工具上。

### 判定順序

1. 先判 Placement（這段能力服務誰）→ 見 [`lifecycle-contract.md`](lifecycle-contract.md)。
2. 再判這段程式應有的 Role。
3. 比對實際所在的 Role / 位置；不一致即違規，用上方格式記錄。

### 常見訊號（篩選用，不是判定本身）

| 訊號 | 可能的違規 |
| --- | --- |
| 入口檔直接呼叫 subprocess / socket / 檔案寫入 | Logic 放在 Entry |
| 入口檔或共用模組有模組層級的座標、ID、序號常數 | Parameters 放在 Logic / Entry |
| 同一個「讀環境變數 + 寫死預設值」出現在多處 | Parameters 未集中，且可能混在 Shared 的 Logic |
| 設定檔欄位名像 `if_*`、`when_*`、`steps` | Logic 放在 Parameters |

## 跨語言適用原則

- 契約只用 Role / Placement / State 三個詞描述；**目錄名、檔名慣例屬專案層**，由專案在自己的 README 對應。
- 機械鎖可以用語言專屬的 pattern 偵測訊號，但鎖的**意圖描述**必須能還原成上方的違規句型，否則屬於「保護檔案形狀而非架構」（見 [`lock-taxonomy.md`](lock-taxonomy.md)）。
- 同一專案用多種語言時，每種語言對應同一組角色，不各自發明分層。

## 不在本契約範圍

- 工具該放在哪個範圍、何時升級 → [`lifecycle-contract.md`](lifecycle-contract.md)、[`growth-rule.md`](growth-rule.md)
- 是否該新增工具 → [`discovery-contract.md`](discovery-contract.md)

← [Back to project-tooling index](README.md)
