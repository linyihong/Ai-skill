> 遵守 [feedback-lessons](../../../feedback-lessons.md) 與 [共用規則索引](../../../../enforcement/README.md)。

### 2026-11-06 — 核對 IL2CPP 布局、複製身分與 view 拓撲

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。單一目標的版本差異、自然／回放窗口與欄位修復讀回支持方法候選；跨目標驗證待補。

#### One-line Summary

以同版布局、物件複製與 view 快取建立觀測契約，再由獨立位置驗證格位和座標單位。

#### Human Explanation

欄位偏移沿用舊 build，可能讀到有效但型別不同的物件，造成一連串模型或渲染假故障。顯示層也可能使用保留 ID 的物件副本，令 pointer-only 配對漏失。邏輯 model 數量與畫面欄列數之間需要獨立映射。

#### Trigger

Runtime 顯示正常，但回放查不到 model、layer 或內容；Create 回傳成功仍無法和模型配對；List index 與畫面位置矛盾。

#### Evidence

- Tool: 同版 metadata 匯出、ARM64 指令、Frida 被動 trace、Transform 位置 getter。
- Sanitized excerpt: 較新版物件布局改變後，按欄位名稱解析可恢復先前失敗的讀回；複製 cell 保留 ID；獨立 view 位置支持邏輯 model 到可見格位的關聯。
- Evidence path: `<PROJECT_ROOT>/docs/slots/<cabinet>/lab/`。原始 trace、位址、樣本計數與截圖保留於業務專案。

#### Generalized Lesson

先對齊 binary／metadata fingerprint、方法 RVA 和執行期型別。以欄位名稱搜尋目前 class 與繼承鏈，cache key 包含 class 身分，欄位缺失須顯示失敗。配對 copied objects 時同時記來源與結果的 ID、指標及生命週期；ID 的適用窗口也要驗證。先找已存在 model 的 ID 關聯，再查後續寫入。用 controller→model→cell ID→view 的關係與獨立 view 位置建立格位，最後才比較待檢驗的 Rect。

view 的內容集合與已處理 ID 快取須分開讀回。內容清空後，快取仍保留 ID 時，更新流程可能跳過重建。以單一格的 ID 失效作最小實驗，並驗證其他格保持原狀。涉及 Unity 物件建立的操作應在引擎主執行緒回呼中驗證；RPC 執行緒上的失敗先記為執行契約問題。

模型 Rect 的單位與 creator/view 的 pixel Rect 分別記錄。active、shown 或內容數量不足以證明可見；應同時核對 Renderer 與 world position。若位置偏差符合尺寸再次縮放的數值關係，可提出單位混用假說，仍須獨立的換算／畫面驗證。

world 座標軸的排序須經 camera 投影或可見單格標記校準，才能轉為 top-first／left-first 格位。僅依 world Y 大小推定上下排可能顛倒。以保留備份的單格實驗校準，並同步修正程式、regression 預期與舊分析 summary 的有效性。

單格座標單位修復與全格畫面驗收分別記錄。即使 model Symbol、內容型別及點數都正確，Item 的尺寸／錨點或殘留物件仍可能令畫面不符。手動回放沒有 child hook 事件時，須另驗證觀測契約，不能直接推定 child 方法未執行。

設定中的 pivot 與實際 RectTransform pivot 應分別讀回。若二者不同，單獨更新尺寸未必恢復位置。先記錄原 Rect，再用目標既有公開 layout 方法作最小修復，核對 pivot、bounds 與畫面；不可將固定 pivot 值當成跨專案規則。

探索性修改可能留下非初始 model offset、重複內容或失效物件。正式驗收前須核對完整前置狀態；前置不符時拒絕修改，必要時重建乾淨場景。將 model／typed content 讀回、獨立資源與位置讀回、可見畫面分別保存，並保留延後讀回以檢查動畫覆寫。

#### Agent Action

進入稀有自然窗口前掛齊必要呼叫點、輸入／輸出及內容註冊，保留整段 phase。phase 切換後可能仍有自然動畫 callback，須以來源關聯判斷。先修復觀測契約，再設計最小介入。

#### Goal / Action / Validation

- Goal: 降低布局錯誤與身分錯配造成的渲染誤判。
- Action: 以欄位反射修復讀回，追複製 ID，核對 model 和 view 位置。
- Validation or reference source: 對修復前後的相同 getter 比對；對同版 Create 來源／結果及內容註冊重跑關聯。

#### Applies When

IL2CPP 可取得同版 metadata 或可信 Runtime 欄位反射，且對目標有授權。

#### Does Not Apply When

ID 跨生命週期被重用且無獨立關聯，或 view 的 world 座標處於旋轉／透視／動畫中，尚未驗證格位排列。

#### Validation

核對現版和舊版不同 class layout 的 cache、繼承欄位與缺失欄位失敗。對 copied ID 的模型關聯同時測試時間先後、指標相同與不同情況。對 view 位置使用多個時間點及實際可見格位驗證。

#### Promotion Target

- `analysis/unity/workflows/il2cpp-runtime-contract-validation.md`，先連結 candidate，待獨立目標重用後評估 promotion。

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

已更新 category／domain feedback index，並於 Runtime 契約程序加入 candidate 方法入口。已依 reusable-guidance-boundary 檢查，業務原始證據保留專案。runtime schema、routing 與其他工具入口未變，無需新增治理規則。

#### Transfer Contract

- Inputs: 同版 binary／metadata、Runtime class、Create 和註冊事件、view 位置。
- Mechanism: 版本位移、保留 ID 的副本、邏輯 model 與畫面拓撲之差。
- Decision: 契約核對失敗先修資料來源，關聯成立後再比較渲染。
- Verification: 用獨立 getter 與未參與配對的內容／位置讀回反駁假說。
- Counterexample: 跨窗口 ID 重用；位置被動畫改寫。
- Transfer case: 另一個使用 pooled cell 和複製 view 的 Unity inventory grid，尚未實測。

#### Intelligence Review

- Decision: method-only
- Reason: 現階段是單一目標的方法候選，先由 Unity analysis 承接。
- Target: analysis/unity/workflows/il2cpp-runtime-contract-validation.md

#### Abstraction Check

- Swap-project test: PASS — 換成其他 Unity 列表或物件格子，仍能用布局、ID 生命週期與獨立位置驗證。
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/<cabinet>/lab/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
