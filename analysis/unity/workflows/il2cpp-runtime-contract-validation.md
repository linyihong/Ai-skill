# IL2CPP 執行資料契約與 hook 生命週期驗證

> **先讀** [`../decision-workflow.md`](../decision-workflow.md)，再依證據進入本 procedure。
> 主張類別（Observation／Inference／Residual unknown）沿用 [`../../reverse-engineering/evidence-contract.md`](../../reverse-engineering/evidence-contract.md)，本檔不另定義。

授權 IL2CPP 分析中，當方法回傳空、fixture 寫入後狀態不完整，或入口命中但後續行為缺失時使用。這是從 candidate lessons 抽出的驗證方法，不代表所有遊戲／應用都有相同類別、格數或補值策略。

端到端編排見 [APK workflow](../../../workflow/apk-analysis/execution-flow.md)。特效驅動另見 [Unity 驅動追蹤](unity-fx-driver-trace.md)。

## 輸入與共同紀錄

記錄版本／架構、目標方法與物件身分、正常輸入對照、fixture schema、實際呼叫入口與授權窗口。輸出是契約差異與已驗證範圍，不是單一「成功」旗標。

| 問題 | 先驗證的契約 | 不足以證明成功的訊號 |
| --- | --- | --- |
| 方法回傳空 | overload、參數型別、表示方式與呼叫 ABI | 方法位址存在／invoke 未拋例外 |
| hook 後續行為缺失 | 目標×階段的分派與適用範圍 | 只有入口總命中數 |
| fixture 寫入異常 | 展示模型與完整執行模型的差異 | 部分欄位已更新／一張畫面完整 |

## 1. 參數與呼叫入口

1. 以目前版本 metadata 確認 declaring type、overload、static／instance、參數型別、by-ref 與回傳型別。
2. 區分 `il2cpp_runtime_invoke` 與直接 Native 呼叫的參數表示；依實際入口與架構準備，不共用猜測的 ABI。
3. 顯示相同不代表型別相同：`System.Char`、managed string、數值 enum 與物件 reference 分開處理。Char 是 UTF-16 code unit；非 BMP 字元不能當單一 char 截斷。
4. 用已知有效輸入與既有錯型別／錯 overload 觀測對照；新增負向案例以離線型別檢查拒絕，不直接對進程呼叫錯誤 ABI。呼叫失敗、未初始化與資料不存在分開報告。
5. 若改讀集合欄位，先核對版本、layout、容量、有效元素與存活期；集合內容不能憑常見欄位名推測。

驗證：已知有效輸入回傳預期內容，負向對照可解釋差異；沒有這些證據就保留 unknown。

## 2. 目標×hook 階段矩陣

建立一列一個目標的矩陣，至少包含 resolved、enter、leave、後續處理、清理、刻意排除理由。

- 各階段按 capability／契約分派。某個目標需要寫資料，不代表也應強制播放 UI；階段集合不一定相同，但差異要可解釋。
- 新增目標時檢查所有階段，避免只加入口分派而返回路徑仍限於舊目標。
- 先找首個缺失階段，再查資料或 renderer；`not-resolved`、`not-hit`、`not-applicable` 與執行失敗分開。
- 事件以目標物件、執行緒／操作關聯與窗口對齊，不能用無關目標的 leave 命中填補缺口。

驗證：覆蓋需要完整後續行為與刻意不適用兩類情境；排除分支有理由，失敗回報具體階段。

## 3. 展示 fixture 與完整 runtime 結構

1. 由正常觀測或權威 schema 確認維度、座標、排序、有效範圍、隱藏緩衝區、缺值與 sentinel 語意。
2. 比較 fixture、序列化 adapter 與記憶體 apply adapter 的實際輸入／輸出，不假設相同名稱代表相同表示。
3. 缺失位置有契約定義且值可合法取得時才展開／補值／調整容量；未知缺值應拒絕或標 blocked。不得把某一專案的固定窗口、edge-clone 或填滿策略當通用規則。
4. 合法 sparse、跨格物件與佔位不得被「修補」破壞；與真正遺失資料分別測試。
5. 若後續事件覆寫資料，先證實 writer 與生命週期；優先修順序或重播契約。固定延遲 re-patch 只可作已驗證的專案策略。
6. 驗證完整狀態轉換與下一次授權操作，避免第一輪成功掩蓋下一輪卡住。

離線反例集：完整資料、僅展示子集、合法 sparse、非法缺值、不同維度。動態證據與離線一致性分開記錄。

## 產出與適用邊界

產出契約差異表、目標×階段矩陣、測試情境與觀測限制。原始 dump／fixture／物件位址留在專案受控 evidence；報告格式與腳本名稱由專案決定。

輸入契約不明時，本流程只能收斂待驗證項目，不能授權猜值、跨目標強制行為或宣稱完整 parity。

## 來源候選與發現入口

- [參數型別](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-16_173200-il2cpp-paytable-getmain-takes-system-char-not-string.md)
- [hook 生命週期](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-23_140500-apply-shim-onleave-ui-must-cover-all-processresult-owners.md)
- [記憶體 apply](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-23_141500-apply-shim-visiblegrid-must-pad-processresult-strip.md)
- [回放 adapter](../../../feedback/history/apk-analysis/local-proxy/2026-10-23_134500-wire-result-glass-must-match-processresult-strip.md)

本方法由 analysis README、既有 APK analysis route 與 summary 發現；候選來源仍保留原狀態，文件抽象化與合成反例不等於獨立專案重用。

判斷來源：[局部證據歸因](../../../intelligence/engineering/analytical-reasoning/heuristics/runtime-evidence-boundaries.md)、[表示契約對齊](../../../intelligence/engineering/analytical-reasoning/heuristics/representation-contract-parity.md)；兩者均為 candidate，依邊界使用。
