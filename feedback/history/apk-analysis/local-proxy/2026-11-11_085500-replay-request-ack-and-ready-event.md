### 2026-11-11 — 重播結算後的請求回覆與就緒通知

Status: candidate

#### One-line Summary

連續重播要驗證客戶端收到與請求 ID 相符的回覆，以及原流程要求的下一輪就緒通知。

#### Trigger

保存結果經代理送達，客戶端完成本輪動畫，但操作按鈕仍停用。

#### Evidence

- Tool: 本機 TCP 代理、保存封包、手機唯讀探針與畫面。
- Sanitized excerpt: 單手重播已到完成回呼，按鈕仍停用。對照原始後續封包，看到一筆與請求 ID 配對的成功回覆，接著有同一對象的下一輪就緒通知。代理重播這個順序後，兩手均完成動畫與盤面讀回，按鈕兩次恢復。第二手從畫面實際啟動。
- Reuse Evidence: 同一目標兩手，待獨立目標重用。兩個回覆在本案一起加入，各自必要性未單獨測試。

#### Generalized Lesson

以自然封包查明每個階段的請求關聯、非同步通知及順序。回覆依測試連線的請求 ID、對象和事件序列綁定。完成驗收同時核對終點、操作恢復和下一輪實際啟動。若缺原始後續資料，保留可驗證的單輪範圍。

#### Agent Action

比對成功與卡住的封包序列，再於受控代理重播已知後續回覆。記錄哪個請求由代理接手，避免把測試操作送到實際服務。

#### Goal / Action / Validation

- Goal: 重播從輸入到下一次可操作。
- Action: 核對請求 ID 與後續狀態通知，再按順序回覆。
- Validation: 比對兩輪的完成事件和狀態，實際觸發第二輪。

#### Applies When

多階段協議的結果與就緒狀態分別傳送。

#### Does Not Apply When

協議以單一回覆更新全部狀態。依該協議驗證其完成契約。

#### Validation

核對代理回覆與操作狀態，再透過畫面啟動下一輪。

#### Transfer Contract

- Inputs: 成功的原始請求與回覆、狀態通知、客戶端狀態讀回。
- Mechanism: 請求完成與後續狀態轉換可能來自不同訊息。
- Decision: 保留觀測到的順序及身分關聯。
- Verification: 驗證兩個完整週期與實際再次操作。
- Counterexample: 單一回覆就使狀態就緒的協議。
- Transfer case: 有獨立就緒通知的裝置操作或多步工作流程。

#### Promotion Target

- analysis/reverse-engineering/collection/capture-review.md

#### Required Linked Updates

已同步抓包審查流程與分類索引。Step 6 不適用，本案先由既有方法承接。Step 7 已將畫面完成但操作仍停用的症狀寫入流程。

#### Intelligence Review

- Decision: method-only
- Reason: 單一目標的實測已能補足審查步驟，跨目標重用仍待驗。
- Target: analysis/reverse-engineering/collection/capture-review.md

#### Abstraction Check

- Swap-project test: PASS。
- Project evidence kept in: <PROJECT_ROOT>/docs/。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
