### 2026-11-10 — 在轉場的第一個可見畫格驗收

Status: candidate

#### Reuse Evidence

同一目標完成直接關閉與準備就緒後關閉兩種回放。獨立目標尚未重用。

#### One-line Summary

延後既有遮罩的原生關閉，直到下一畫面的可見內容與狀態已讀回。

#### Human Explanation

使用者已點擊、下一場景已開始準備和畫面已可安全顯示是三個時點。終點截圖正確仍可能漏掉中途出現的空白或替代內容。

#### Trigger

切換畫面短暫顯示舊資料、空白或錯誤布局，稍後恢復。

#### Evidence

- Tool: 原生確認回呼探針、顯示狀態讀回、連續錄影。
- Sanitized excerpt: 保留原生遮罩時先準備下一畫面。等顯示物件與標籤讀回成功，再放行同一確認回呼，第一個可見畫面與預期一致。
- Evidence path: <PROJECT_ROOT>/docs/slots/<cabinet>/lab/。

#### Generalized Lesson

把輸入確認、下一畫面準備和遮罩關閉分開記錄。隔離回放可暫留同一原生遮罩，讓場景準備留在其後。核對可見物件與標籤，再於引擎執行緒放行原本的關閉回呼。確認一次點擊只放行一次，並以連續錄影驗收第一個可見畫格及後續動作。記下失敗的回復方式。

#### Agent Action

先保留空白或錯畫面的時序證據，再測試是否存在遮罩與內容的交接差。

#### Goal / Action / Validation

- Goal: 排除轉場中間的錯畫面。
- Action: 在資料與顯示均就緒後放行原生關閉。
- Validation or reference source: 獨立畫面時間序列與前後讀回。

#### Applies When

既有遮罩可獨立於背景場景準備保留，並有安全的回呼與失敗回復契約。

#### Does Not Apply When

遮罩關閉是內容準備不可拆分的先決條件，或無法安全暫留及重入回呼。

#### Validation

對照立即關閉與等待讀回後關閉。若第一個可見畫格仍錯，保留失敗，不以最終截圖替代。

#### Promotion Target

- analysis/unity/workflows/unity-dynamic-flow-validation.md

#### Promotion Record

維持候選，待獨立目標重用。

#### Required Linked Updates

更新 Unity 動態流程程序與 feedback 分類索引。原始影像留業務專案。

#### Transfer Contract

- Inputs: 遮罩、原生確認回呼、下一畫面內容、可見性證據。
- Mechanism: 背景準備與遮罩關閉是獨立的時序。
- Decision: 以讀回決定是否放行遮罩。
- Verification: 比對第一個可見畫格及稍後畫面。
- Counterexample: 場景只有在遮罩關閉後才配置 view。
- Transfer case: 關卡載入過場後切換到道具格。

#### Intelligence Review

- Decision: method-only
- Reason: 單一目標，先由既有驗證程序承接。
- Target: analysis/unity/workflows/unity-dynamic-flow-validation.md

#### Abstraction Check

- Swap-project test: PASS — 同法可驗收非同步載入道具格或場景。
- Project evidence kept in: <PROJECT_ROOT>/docs/slots/<cabinet>/lab/
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
