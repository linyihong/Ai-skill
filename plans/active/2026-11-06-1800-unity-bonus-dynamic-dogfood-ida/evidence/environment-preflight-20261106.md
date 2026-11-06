# IDA 環境 preflight

日期 2026-11-06。僅作唯讀環境與既存材料檢查，計畫維持 draft。

| 項目 | 結果 | 邊界 |
| --- | --- | --- |
| PATH 中 IDA／批次工具 | not-found-in-scope | 已查常見命令名稱 |
| 常見安裝位置 | not-found-in-scope | 系統／使用者 Applications 列檔，以及 opt、Downloads、Tools 常見根目錄候選檢查；非全磁碟搜尋 |
| IDA 版本／授權 | unknown | 尚無可執行工具可檢驗；未讀取或保存 license 內容 |
| ARM64 Hex-Rays／IDAPython | not-tested | 找到工具後仍需實際 smoke |
| 既存 binary／metadata | matched-prior-baseline | 重新計算兩者指紋，與專案既有驗收基線一致；未查目前手機版本 |
| metadata 匯出材料 | available | 先前匯出存在，不能當成本輪 IDA 匯出 |
| IDA 可重跑匯出／預測驗證 | not-tested | 本輪尚未產生 |

本地搜尋只能支持指定範圍內未找到工具，不證明沒有授權或其他主機可用。B track 的工具前置目前 blocked；A track 尚未測試。本輪未操作手機或改動 Bonus。

具體 fingerprint 與材料檢查留於 `<PROJECT_ROOT>/docs/slots/el-dorado-dices/lab/bonus-dynamic-ida-preflight-20261106.md`。

## 下一步

取得可用的 IDA 環境位置後，驗證版本、授權可用性、ARM64 反編譯與 IDAPython，再做最小靜態匯出。先前替代工具的證據保持原來源，不改標為 IDA 成果。
