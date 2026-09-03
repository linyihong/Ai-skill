# Facial expression 首次 stage PASS — 兩個 contract 缺口

**Run ID**：2026-09-03-facial-expression-first-pass-consumer-loop
**Authorization**：consumer 要求「先做一個微笑」並在互動 viewer 上確認
**Scope**：consumer dogfood + contract 回饋。未跑 execution dogfood 七條，未碰 Phase 4／5／6。

Filled instances 與量測留在 consumer `<PROJECT_ROOT>`。本檔不含專案名、絕對路徑、素材檔名。

## 結論

Facial expression observation 首次由 `partial` 轉為 **PASS**：真實臉部幾何雙側嘴角上揚，
geometry 與 visible readback 同時成立，neutral baseline 無漂移。

過程暴露 **兩個 contract 缺口**，均已回寫 workflow：

1. record 無法區分「表情從未產出」與「產出但驗證失敗」。
2. stage 證據只有離線算圖，未涵蓋資產實際被消費的互動介面。

其餘 gate 不變：眼部表情仍因貼圖化眼睛而無法通過 visible readback；pack 仍是 **prototype**。

## 缺口一：負向 readback 未被賺取

該表情長期記為 `fail`，實際成因與資產無關：

| 表象 | 實際成因 | 修復位置 |
|---|---|---|
| 匯出物中形變資料全為零 | 產出步驟的 opt-in 旗標預設關閉，資料從未被建立 | 產出設定 |
| 自製探針讀出全為零 | 解析器未支援容器格式的稀疏儲存分支 | 驗證工具 |

兩者在舊 record 上都長成 `fail`，指向「重做資產」這個最貴的層。真正的動作是打開一個旗標、
以及讓探針支援稀疏分支。

**回寫**：observation 新增 `artifact_state`，record 新增 `probe_calibration`；
`negative_evidence_requires` 要求 artifact 已產出且探針在已知正例上讀到訊號，
否則 no-change 只能是 `partial`。

## 缺口二：消費面未納入證據

離線算圖已達可見門檻後，consumer 在互動 viewer 上仍**看不出效果**。三個缺陷都只在
驅動真實 UI 時才顯現：

- 除錯標記預設疊在觀察區上，遮住嘴部。
- 強度控制項的處理器先呼叫重置函式，導致其後讀到的輸入值恆為 0。
- blocking 寫成硬編碼常數，缺陷修復後仍繼續阻擋已驗收的表情。

**回寫**：observation 新增 `consumer_readback`；`forbidden_evidence` 新增
`offline_render_only_when_consumer_surface_exists` 與 `uncalibrated_probe_no_change`；
stage 文件新增 Negative readback 與 Consumer-surface readback 兩節。

## 方法

- 資產層：DCC 端與匯出端雙向比對，探針先於已知正例校準。
- 消費層：自動化瀏覽器驅動真實控制項，逐項斷言觀察區像素變化與回復基準態。
- 強度掃描定出可見度與貼圖形變之間的上限，取樣點由人工挑選而非只看數值。

## 對應 lessons

- `feedback/history/3d-character-production/common/2026-09-03_101100-negative-readback-must-be-earned.md`
- `feedback/history/3d-character-production/common/2026-09-03_101200-interactive-surface-readback.md`

## 不是什麼

- 不是 facial expression stage decision 整體 PASS（眼部 observation 仍未通過）
- 不是把 live 角色升成 `runtime-ready`
- 不是 Phase 4 profile／Phase 5 route／Phase 6 完整 dogfood
