> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-01 - 分離資產證據與 stage PASS

Status: validated

#### One-line Summary

3D 資產 gate 必須分開記錄實際幾何、可見 readback、neutral baseline 與 stage eligibility；任一必要證據只有 `partial` 時不得宣稱 PASS。

#### Human Explanation

Morph delta、匯出成功或畫面上的輔助圖形，只能證明局部現象。它們不能單獨證明
目標資產本身已正確變形、使用者可見結果成立，或 neutral 狀態沒有漂移。若 workflow
把這些證據混成單一布林值，局部成功會掩蓋材質、映射或 identity 的 blocking failure。

#### Trigger

- 幾何資料顯示有變形，但 runtime／viewer 看不到預期表情。
- 可見差異來自 detached helper、overlay 或其他非目標資產。
- surface/material 尚未驗收，卻因 rig 或 export 成功而準備推進。
- record 支援 `partial`，但 gate 只定義「有 fail 才 fail」。

#### Evidence

- Tool: DCC inspection、runtime readback、fixed-view visual comparison。
- Sanitized excerpt: 局部幾何證據與可見證據互相矛盾時，stage 必須保持未通過。
- Evidence path: 具體 consumer evidence 留在 `<PROJECT_ROOT>` 的專案文件或 artifacts。

#### Generalized Lesson

每個 facial／deformation stage 至少分開保存：

1. 目標資產的 geometry readback。
2. 同一 asset revision 的 visible/runtime readback。
3. neutral baseline／identity validity。
4. surface/material eligibility。
5. 明確的 stage decision。

`partial` 表示證據或能力尚未閉合，不是較低強度的 PASS。Stage gate 應採
`pass_requires`，而不是只用 `fail_when` 推導成功。

#### Agent Action

先確認 stage ownership 無循環，再用 record 欄位分別承載 geometry、visible、
baseline 與 eligibility。Helper-only、export-only 或單側 readback 填 `partial`／`fail`；
只有所有 required evidence 都 pass 才允許下一 stage。

#### Goal / Action / Validation

- Goal: 防止局部資產證據被誤升為 stage completion。
- Action: 採獨立 facial acceptance record、明確 `pass_requires`，並讓 surface／identity 在下游 gate 重讀。
- Validation or reference source: 負向 scenario 必須覆蓋 partial、helper-only、surface fail 與 stale identity。

#### Applies When

- 可驅動 3D 角色的 rig、facial expression、material 或 runtime consumption 驗收。
- 一個 stage 同時依賴幾何與使用者可見結果。

#### Does Not Apply When

- 純 exploration 且明確不做 promotion、completion 或 downstream eligibility claim。
- 非目標資產的 debug visualization，且沒有被當成 acceptance evidence。

#### Validation

- 任一 required evidence 為 `partial`、missing 或 fail 時，下一 stage 被阻擋。
- Detached helper／overlay 不可滿足 real asset geometry 或 visible readback。
- Export success 不可覆蓋 surface、identity 或 runtime consumption failure。
- Stage graph 不要求尚未進入的下游 stage 反向完成上游 gate。

#### Promotion Target

- `workflow/3d-character-production/`
- `validation/scenarios/3d-character-production/`

#### Required Linked Updates

- 已同步 execution record、artifact gates、stage docs 與負向 scenarios。
- Reusable lesson 只保留 generalized rule；具體資產、路徑與量測留在 consumer project。
