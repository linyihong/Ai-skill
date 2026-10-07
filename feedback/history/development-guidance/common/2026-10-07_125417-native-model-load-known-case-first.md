> 遵守 [feedback lesson 規則](../../../feedback-lessons.md)；具體執行證據留在專案文件。

# 原生模型載入失敗先比對既有案例

Status: candidate

## Reuse Evidence

尚未有獨立專案／環境的重用證據，維持 candidate；不能把局部修正當通用套件相容保證。

## One-line Summary

模型載入原生崩潰時先比對既有失敗指紋與版本範圍，再做最小隔離驗證，不從完整產品流程重新猜原因。

## Human Explanation

裝置初始化成功、小型量化運算成功或權重檔存在，都不能證明完整模型能載入。
原生存取違例也可能發生在權重映射／切片讀取，而非推論或量化核心。
只保存最後一個 Python 例外會漏掉程序直接終止的證據。

## Trigger

載入權重進度中斷、程序非正常退出、原生 storage／slice／memory-map 堆疊，
或相同錯誤在重試後再現；錯誤碼單獨相同不足以判定同一原因。

## Evidence

- Tool: 擁有者限定的子程序、原生堆疊、退出狀態與單變量讀取方式對照。
- Sanitized excerpt: 讀取策略對照能區分權重 materialization 與生成階段；具體版本和結果不寫入泛用正文。
- Evidence path: `<PROJECT_ROOT>/docs/analysis/2026-10-07-qwen-native-loading.md`；對應套件清單與測試報告亦留在該專案。

## Generalized Lesson

重用失敗指紋及驗收方法，不重用未核對的全域原因結論。

## Agent Action

1. 先搜尋專案已知案例；比對失敗階段、原生堆疊、平台、套件版本、模型與權重格式及修正是否仍在。
2. 指紋與範圍一致：先驗既有相容修正和小範圍回歸，不反覆跑整片或同一會崩潰的 loader。
3. 版本／模型／堆疊不同、修正缺失或最小驗證失敗：保留差異，重新開隔離診斷；不得套用舊案例直接宣稱修好。
4. 需重現時，用自有子程序啟用原生堆疊記錄；父程序保存 stdout、stderr、PID、退出碼和 timeout 分類，只清理自有程序。
5. 一次只改讀取 backend、dtype、device placement 或套件其中一項。映射讀取、逐段讀取與整份讀入 RAM 不可視為等價。
6. 相容 adapter 限定權重路徑及生命週期；回歸驗證其他檔案不受影響、正常及異常退出恢復 hook。
7. 完整載入→最小生成→實際產品輸出分開驗收。父程序顯示／編碼失敗與子模型失敗亦分開判讀。
8. 不擅自改用雲端、下載另一模型、廣泛升級、全精度／CPU 回退或終止其他工作來掩蓋失敗。

## Goal / Action / Validation

- Goal: 重用已有診斷，縮小重試成本且避免錯誤泛化。
- Action: 保存可比對指紋與修正範圍，建立索引及專案操作入口。
- Validation or reference source: 核對專案原生堆疊、單變量對照、adapter 回歸及實際輸出證據；文件鏈接與 runtime 索引驗證另做。

#### Applies / Does Not Apply

### Applies When

有可追溯的既有原生載入案例，或 loader 在產生業務輸出前非正常退出。

### Does Not Apply When

一般可捕捉的 Python 設定例外、下載／權重完整性錯誤、生成 OOM，或語義／字幕版面問題；需按各自證據處理。

## Validation

相同指紋能找到原紀錄；不同版本／堆疊會被要求重驗；最小載入成功不被升格成產品品質通過。

## Promotion Target

目標：專案操作入口及既有 Translation Decision 的原生載入隔離證據。

## Promotion Record
尚未 promotion 至全域 workflow／enforcement；維持 candidate，等待獨立重用。

## Required Linked Updates

- 已更新 domain／common 索引，既有計畫證據加入本條連結；專案操作規則連回具體診斷。
- 已檢查 reusable-guidance boundary：版本、主機、完整堆疊與 live 結果留專案；保留英文僅為工具／API／搜尋術語。
- Step 6 不適用：單一候選方法不再拆 intelligence atoms，無新分類或架構變更。
- Step 7 不適用：本條整理已驗產品原生失敗，未證明新增跨代理治理失效。
- 全域 workflow promotion 不適用：尚無獨立重用；runtime refresh／validate 與 commit／push／readback 必須完成。
