### 2026-11-10 — 從前段事件找回回放差額

Status: candidate

#### Reuse Evidence

同一目標有對帳補值與跨段原生累加的對照。獨立目標尚未重用。

#### One-line Summary

重播總額不足時，查前段過渡事件是否已有會累積到後段的原生獎勵。

#### Human Explanation

資料庫的逐手投影可能從中途開始，把過渡階段已得到的獎勵攤入後段結算。逐手補值能重現畫面金額，但會遮蔽實際獎勵來源。

#### Trigger

每段事件看似正確，最終總額仍少一筆，或開始狀態已有額外獎勵。

#### Evidence

- Tool: 唯讀保存封包解碼、同版 managed 型別與獨立 HUD 讀回。
- Sanitized excerpt: 過渡段的原生獎勵總和與先前人工補值總和一致。先播過渡段，再播保存的逐手事件，最終原生累加和視窗金額吻合。
- Evidence path: <PROJECT_ROOT>/docs/slots/<cabinet>/lab/。

#### Generalized Lesson

沿同一活動的已驗證事件關聯向前找過渡幀，對照起始餘額與逐段變化。保留先前只播中段的對照。以原生型別重播過渡與中段，分別讀回事件小計、執行期總額和最終視窗。若入口、出口或個別演出需測試介入，將它們與原生總額驗收分開。

#### Agent Action

缺少同活動的過渡資料時保留差額，對帳補值另行標記。

#### Goal / Action / Validation

- Goal: 識別回放總額的前段來源。
- Action: 比對中段回放和含過渡的原生回放。
- Validation or reference source: 保存的分段結果、原生讀回與視窗畫面。

#### Applies When

有可關聯的多段資料、明確金額口徑和相同版本的原生執行環境。

#### Does Not Apply When

資料只能以記錄 ID 鄰近推測關聯，或累計與分段金額口徑不可比。

#### Validation

去掉所有手動補值。若過渡段小計、後續小計或最終視窗金額任一對不上，前段來源假說仍未成立。

#### Promotion Target

- analysis/unity/workflows/unity-dynamic-flow-validation.md

#### Promotion Record

候選，待獨立目標重用。

#### Required Linked Updates

更新 Unity 動態流程程序與 feedback 分類索引。原始封包、影片與記錄識別留在業務專案。

#### Transfer Contract

- Inputs: 前段與後段記錄、獎勵事件與結算總額。
- Mechanism: 前段獎勵累積入後段顯示，整理後的逐段資料可能合併金額。
- Decision: 以原生事件重播驗證來源，再判斷是否需要顯示對帳。
- Verification: 排除手動補值後核對原生小計與最終視窗。
- Counterexample: 匯率換算或稅額差異也會造成表面差額。
- Transfer case: 任務過場先發獎勵，關卡結果只列後段分數。

#### Intelligence Review

- Decision: method-only
- Reason: 單一目標驗證，適合由既有程序承接。
- Target: analysis/unity/workflows/unity-dynamic-flow-validation.md

#### Abstraction Check

- Swap-project test: PASS — 任務過場或分段結算可沿同一活動關聯重播前段。
- Project evidence kept in: <PROJECT_ROOT>/docs/slots/<cabinet>/lab/
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
