# 採集踩雷速查（跨目標）

症狀 → 可能原因 → 修法。只列**去掉產品／引擎專名後仍成立**的列。Unity／Frida／IL2CPP 細節見 [`../../apk/workflows/`](../../apk/workflows/README.md) 與 `feedback/history/apk-analysis/`。

操作紀律正文見 [`collection-discipline.md`](collection-discipline.md)；抓包審查見 [`capture-review.md`](capture-review.md)。

## 入口與命名

| 症狀 | 可能原因 | 修法 |
| --- | --- | --- |
| attach／dump「找不到視窗／根」 | 用靜態資源名當 runtime 名 | Probe-before-attach；先 list 實際名 |
| 掛上錯誤進程 | 包名 ≠ 顯示名；或掛到自家複刻進程 | 用當下可觀測 PID／顯示名；確認目標發行物 |
| 「解碼／切圖壞了」傳說 | 跨平台混用幾何或資源 | Platform consistency；同平台重採 |

## 瞬態與兩階段

| 症狀 | 可能原因 | 修法 |
| --- | --- | --- |
| 界面沒了，什麼都沒抓到 | 先分析再落盤；或反覆短 attach | Transient full-dump：一次全讀先落盤 |
| 預期字串／節點永遠沒出現 | 採集期預過濾針表 | Two-phase：Phase A 全量；Phase B 再濾 |
| 針命中無關 app | 針名過泛 | Search hygiene：特異前綴 |

## 靜態／runtime／交付

| 症狀 | 可能原因 | 修法 |
| --- | --- | --- |
| 靜態樹遠小於畫面 | 殼層 runtime 組裝 | Dual source；殼看 runtime |
| 實作說「缺資源」但其實有檔 | 只交孤立檔、無映射 | Deliver relations；未 ready 不 handoff |
| 「查不到」過早 | 只掃一個靜態來源 | 相關靜態來源全掃後才標 unknown |
| 組件／API「存在」但行為沒發生 | 存在 ≠ 會執行／會播 | 分寫 observation；要 runtime 證據再升格 |

## 網路／capture

| 症狀 | 可能原因 | 修法 |
| --- | --- | --- |
| 從 SNI 寫出 REST path／body | 跳層發明 | Capture-review 分層；禁發明 |
| 屬性全空、正則全滅 | 轉義未還原或錯誤 unescape | 依格式還原；避免破壞編碼的全域 unescape |
| 搜到「協議」但其實是本輪產物 | 自我匹配 | 排除本輪 dumps／fixtures |
| TLS 密文當 API 正文 | 未 fail-closed | 停捕；標 needed_authority |
| 只有代理綠燈就當目標狀態 | 代理信號 ≠ 目標事實 | 見 evidence-hierarchy；補目標側證據 |

## 實作與查證迴圈

| 症狀 | 可能原因 | 修法 |
| --- | --- | --- |
| 同 UI bug 改三次仍錯 | 實作頂替採集 | Impl-retry ceiling → 查 capture gap |
| 用截圖／猜綁定頂呈現 | 缺件仍硬做 | 停做；開採集或標 unknown |
| 模型反復讀圖仍不確定 | 錯誤查證手段 | Human recognition export |
| 「看起來差不多」爭議 | 無對照指標 | Numeric visual QA + 基線 |
| 「運行時讀 X 不可行」擋路 | 籠統禁令 | 寫具體條件；對照已知正確樣本 |
| 說要重採但盤上已有 | 未翻已落盤資料 | 先搜既有 evidence 目錄 |

## 元級

| 症狀 | 可能原因 | 修法 |
| --- | --- | --- |
| 第二次踩同一坑 | 未寫回流程／lesson | Mistake → reusable writeback（去敏） |
| 錯誤結論比沒結論更貴 | 把 inference 寫成 observation | 強制三類陳述 |

## 域內加深（不在本表展開）

| 主題 | 入口 |
| --- | --- |
| Frida／hook 安全與參數契約 | `analysis/apk/workflows/frida-hook-flow.md`、`il2cpp-runtime-contract-validation.md` |
| Local proxy／MITM 導流 | `analysis/apk/workflows/local-proxy-hook-flow.md`、`traffic-triage.md` |
| Unity 資產／特效 | `unity-asset-parse-flow.md`、`unity-fx-driver-trace.md`、`unity-soft-visual-restore.md` |
| 歷史案例 | `feedback/history/apk-analysis/` |

← [collection/](README.md)
