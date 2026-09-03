> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-03 - 互動式交付物需驅動真實控制項驗收

Status: validated

#### One-line Summary

當資產最終透過互動介面被消費時，離線算圖通過數值門檻不等於使用者看得見；acceptance 必須驅動真實控制項並逐項斷言效果，且阻擋旗標要由 record 推導而非硬編碼。

#### Human Explanation

離線算圖能證明資產本身會變形，卻繞過了使用者實際接觸的那一層。消費端可能因為預設疊加
的除錯圖層遮住觀察區、事件處理順序把輸入值歸零、或一個早期寫死的 blocking 常數而完全
看不到效果。這些缺陷不會出現在任何離線畫面比較裡，卻會讓「已通過」的 stage 對使用者
呈現為失敗。

#### Trigger

- Stage 已依離線畫面比較宣告 PASS，但使用者回報看不出差異。
- 消費端有預設開啟的除錯 overlay、標記或輔助幾何。
- 阻擋條件寫成硬編碼常數，而非讀取 acceptance record。
- 控制項的事件處理同時負責「停止既有狀態」與「套用新值」。

#### Evidence

- Tool: 以自動化瀏覽器／runtime 驅動真實控制項，對觀察區做像素 readback。
- Sanitized excerpt: 離線比較已達門檻，但互動介面上除錯標記預設遮住觀察區；同時強度控制項的處理器先呼叫重置函式，導致其後讀到的輸入值恆為 0。兩者皆僅在驅動真實 UI 時才顯現。
- Evidence path: 逐控制項檢查清單與截圖留在 `<PROJECT_ROOT>` 的專案 artifacts。

#### Generalized Lesson

驗收層級要對齊交付物被消費的層級。若最終消費面是互動介面：

1. 驗收腳本必須操作真實控制項，而非直接寫入底層狀態。
2. 每個控制項各自斷言可觀測效果，並包含回復到基準態。
3. 預設視圖不得被除錯輔助物遮蔽；輔助物的顯示狀態要與其開關一致。
4. 阻擋旗標必須逐項由 acceptance record 推導；全域硬編碼的 block 會在缺陷修復後繼續阻擋。

離線畫面比較是資產層證據，不是消費層證據；兩者都要，且不可互相替代。

#### Agent Action

在 acceptance 中加入「消費面 readback」：以自動化方式載入真實介面、逐一驅動控制項、
斷言觀察區像素變化與回復。發現硬編碼 block 時，改為讀取對應 record 欄位並支援逐項
放行，讓已驗收項目自動解除阻擋。

#### Goal / Action / Validation

- Goal: 防止資產層 PASS 與使用者實際感知脫節。
- Action: 建立 per-control 互動檢查清單並納入 stage 證據；blocking 改為 record-derived。
- Validation or reference source: 負向 scenario 需覆蓋「資產正常但介面遮蔽」「控制項無作用」「缺陷已修但硬編碼 block 仍阻擋」。

#### Applies When

- 資產最終由互動 viewer、runtime 或應用程式介面消費。
- Stage 證據目前只有離線算圖或數值差異。

#### Does Not Apply When

- 交付物本身即為離線產物（靜態算圖、批次輸出），無互動消費面。
- 尚未存在消費面的早期探索，且不宣稱 stage eligibility。

#### Validation

- 每個消費面控制項都有對應的可觀測斷言與回復斷言。
- 預設載入狀態下，觀察區未被除錯輔助物遮蔽。
- 移除某項 blocking 缺陷後，對應控制項在不加旁路參數的情況下即可使用。

#### Promotion Target

- `workflow/3d-character-production/`
- `validation/scenarios/3d-character-production/`

#### Required Linked Updates

- 同步 facial expression acceptance record 的消費面 readback 欄位與 stage 文件。
- Reusable lesson 只保留 generalized rule；具體介面、選擇器與截圖留在 consumer project。
