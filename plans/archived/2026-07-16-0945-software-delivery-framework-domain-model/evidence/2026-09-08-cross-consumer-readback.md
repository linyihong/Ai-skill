# 跨 consumer Domain Model 結案覆核

Run ID: sd-domain-model-closeout-2026-09-08
日期：2026-09-08
範圍：既有來源、分類與局部機械檢查；不是產品功能驗收或 fresh-context 三角色驗證。

## 來源與重現邊界

使用者指定兩個 consumer。A 為多 repository workspace；B 為 domain-bundle monorepo。
本機綁定、各 repo HEAD、20 份來源的 SHA-256、working-tree 狀態與檢查輸出，保存於
repo 根 gitignored `local/plan-evidence/sd-domain-model-closeout-2026-09-08/`。
公開紀錄使用 A/B 代稱；專案名稱、類別、私有路徑及業務內容留在 consumer。

方法是先讀專案 Policy／Process，再沿具名實例追到 contract、測試、執行器或既有 evidence。
下列六列是兩個專案內的樣本，不能當成六個獨立專案，也不是全庫覆蓋率。
B 是既有 greenfield dogfood 的後續演進觀察，不計作新第三個 consumer。

## 案例矩陣

| ID | 可回查觀察 | Asset → Policy → Process | 判定與上限 |
| --- | --- | --- | --- |
| A1 | 行為規格帶 Test／Code 引用；規格、測試、實作分處外層與 sibling；抽樣引用檔案存在 | Contract／implementation／test assets → workspace placement／traceability → 實作與驗證 | 可解釋跨 repo 放置；檔案存在不等於業務行為 PASS |
| A2 | 測試放置 Policy 在 overlay；guard 讀 pattern config；測試同時包含允許及拒絕案例 | Automation 也是受管理 Asset；其行為投影 placement Policy，於 commit Process 執行 | 複製 guard/config 至臨時目錄執行既有 self-test，exit 0；未執行會 reset/stage consumer 的整套測試 |
| A3 | 既有 incident 記錄 inner test 綠但使用者路徑失敗；要求外層 journey evidence | Evidence／lesson → 證據權威與 promotion Policy → 驗證／收口 | 支持 evidence 與 decision 分工；本次未重跑瀏覽器 journey，歷史紀錄不升為本輪 runtime PASS |
| B1 | Charter 宣告 domain bundle、BDD、plan、evidence 各自放置與 owner；領域目錄有 brief／invariants／API contract | 同一 bundle 含多種 Asset → charter placement → intake／contract／implementation | 目錄不同不要求第四核心；evidence 可在 plan 內，不因預設 docs/evidence 不存在而判缺件 |
| B2 | 權限歸屬計畫具 inventory、ownership manifest、checker 與 archived evidence；目前 checker 通過 | Policy 定義／manifest／checker assets → owner 約束 → 註冊與驗證 | 目前 read-only ownership checker exit 0；支持專案級投影有 consumer，不代表全域 placement 已機械化 |
| B3 | 已解 Decision 優先於 blueprint 的 Open/Draft 段落；plan evidence 隨計畫歸檔 | Decision／plan／evidence → authority／lifecycle → 決策與結案 | 可用既有三核心解釋狀態演進；只確認宣告與抽樣產物，不聲稱所有決策都正確 |

## 反例檢查與限定

1. **Intent 可以被保存**：brief／accepted plan 是承載 intent 的 Asset；不能沿用早期
   「Intent 一律不持久化」作為排除第四核心的理由。這些樣本仍可由 Asset + Process 輸入解釋。
2. **Automation 具有雙重觀察面**：checker 檔案是 Asset，執行語意是 Policy／Process 的投影；
   兩種描述不是兩個 owner，也不要求新增第四核心。
3. **一份 YAML 可同時承載 Policy 與 Process**：按段落責任分類，不能依副檔名或檔案數判 N。
4. **專案已有 gate 不等於通用 gate 完成**：A/B 的 selector、owner manifest 與部署方式不同；
   本輪不把任何 project-specific checker 複製進 Ai-skill。
5. **既有來源可能有陳舊敘事**：A 的 workflow 仍標 draft、B charter 保留演進段落。
   本輪採具體實例與執行器讀回，未將整份專案文件視為全面合規證明。

## 結論

六個抽樣案例未提出無法由 Asset／Policy／Process 解釋的穩定第四核心。
這支持本計畫的有限完成條件，不證明 N=3 對所有未來專案永遠成立，亦不證明模型能力提升。
既有 README、Domain Policies 與 glossary 已足以承載結論，無需新增第一級文件樹或 ADR。

## 檢查紀錄

| 檢查 | 結果 | 證據上限 |
| --- | --- | --- |
| A/B source manifest | 20/20 路徑存在，保存 hash 與 repo revision | 來源可回查；不等於20項行為測試 |
| A guard/config 既有 self-test | PASS，exit 0 | 臨時副本中的分類正負案例；不是 live staged-hook proof |
| B 現行 ownership checker | PASS，exit 0 | 當下 source／manifest 一致性；不是服務端授權整合測試 |
| Consumer 修改 | 無 | 不修改產品，也不操作外部服務 |
