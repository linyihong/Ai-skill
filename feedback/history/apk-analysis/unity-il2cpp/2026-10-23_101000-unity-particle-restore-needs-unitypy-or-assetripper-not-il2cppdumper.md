> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；專案原始證據留在來源專案。

### 2026-10-23 - 依資產證據目標選擇 Unity 粒子分析工具

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本次文字整理與離線檢查不構成第二個專案的實測。

#### One-line Summary

分開取得貼圖、粒子模組、runtime 身分與重播證據，工具選擇以所需證據及版本支援決定。

#### Human Explanation

只匯出貼圖或取得程式骨架，不能證明已還原粒子參數與觸發鏈。資產解析與程式身分定位是不同工作，不能以單一工具完成情況概括全部。

#### Trigger

- 有粒子畫面，但只取得貼圖或型別／方法名稱。
- 要求可在 Editor 重播，卻只有資產參數。

#### Evidence

- Tool: 來源專案既有分析／驗證紀錄；本輪未重跑其執行環境。
- Sanitized excerpt: 非遊戲 Unity 展示或互動應用也能使用相同的資產／程式／runtime 證據分工。
- Evidence path: `<PROJECT_ROOT>/` 的受控分析與驗證紀錄；不複製原始物件、樣本或部署紀錄。

#### Generalized Lesson

1. 按目標拆出 texture、component modules、material 引用、runtime matching 與 replay 所需證據。
2. 可用 UnityPy 讀取支援版本的序列化資產；需要 Editor 工程關係時評估 AssetRipper。工具受版本與腳本相依限制。
3. Il2CppDumper 的型別與方法資料用於定位程式路徑，不取代資產參數。
4. 經 component 的 GameObject 與 material 引用回查資產，不套用固定物件或貼圖名稱。
5. runtime matching 與 Editor replay 是不同維度，不能因其中一項成立推定另一項。

#### Agent Action

先辨識輸入契約與前置條件，再依本條決策；不照搬來源專案的函式名、腳本路徑、固定數字或狀態碼。記錄採用／不適用的理由。

#### Goal / Action / Validation

- Goal: 分開取得貼圖、粒子模組、runtime 身分與重播證據，工具選擇以所需證據及版本支援決定。
- Action: 依 Generalized Lesson 記錄輸入、決策與反例。
- Validation or reference source: 每項宣稱附對應證據或 blocked 理由；粒子參數、runtime 匹配與重播結果分別驗證。

#### Applies When

- 授權 Unity 粒子資產分析，且能取得相關包或 runtime 觀測。

#### Does Not Apply When

- 純骨骼動畫且無 ParticleSystem。
- 只要求靜態貼圖時，不必強制補齊其他維度。

#### Validation

每項宣稱附對應證據或 blocked 理由；粒子參數、runtime 匹配與重播結果分別驗證。

#### Promotion Target

[analysis/apk/workflows/unity-asset-parse-flow.md](../../../../analysis/apk/workflows/unity-asset-parse-flow.md)。

#### Promotion Record

2026-10-26：抽象化整理並接到上述方法／入口；這是候選指引的連結，不宣告 validated／promoted，也不新增獨立重用證據。

#### Required Linked Updates

- 本分類 README 與上述方法／入口已檢查並同步。
- 索引與 RAG 經 runtime refresh 重建；原始專案證據不搬入共用庫。

#### Abstraction Check

- Swap-project test: PASS — 非遊戲 Unity 展示或互動應用也能使用相同的資產／程式／runtime 證據分工。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

依近期回饋審查，把原先依賴專案術語或操作細節的敘述改為輸入、機制、決策與驗證；修訂前內容可由 Git 歷史追溯。Status 保守維持 candidate。
