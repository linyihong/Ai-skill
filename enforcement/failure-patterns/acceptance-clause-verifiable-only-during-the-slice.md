# Acceptance Clause Verifiable Only During The Slice（驗收條款只在切片執行當下可驗）

Status: candidate
Class: `validation-gap`

## Trigger

一個切片的工作是**替換機制**：把靜態門檻換成動態查詢、把手寫對照表換成推導、把舊 API 換成新 API、把兩份宣告收斂成一份。

驗收條款自然寫成「轉換前後行為等價，且以機械方式驗證」。

## Failure Mode

該比對的**另一側就是被移除的東西**。切片完成後，「之前」不存在了，比對沒有對象。

於是那個檢查會走向兩種結局，兩種都不好：

| 結局 | 症狀 |
| --- | --- |
| 被改寫成別的斷言 | 測試名稱與檔案還在，看起來仍在守著等價性，實際只斷言了較弱的東西（例如「沒有殘留的舊寫法」） |
| 靜默變成空轉 | 迴圈的來源集合為空，測試恆綠（見 [`cross-boundary-agreement-never-mechanically-checked.md`](cross-boundary-agreement-never-mechanically-checked.md) 第 7 條） |

而驗收條款仍寫著「已機械驗證」。日後讀到它的人會假設存在一個常駐守衛 —— 它不存在，而且**無法**存在。要重新取得該保證，唯一方式是從版本歷史還原被移除的一側再比對。

實際觀察到的形態：驗收要求「轉換前後角色集合逐一相同，機械驗證」。比對在轉換前的子切片落地並抓到一項真實分歧；轉換子切片移除了 `RequireRole` 宣告，該比對隨之被改寫為「沒有 policy 還讀角色」。等價性保證此後只存在於 git 歷史。

## Risk

- 驗收文字承諾一個常駐保證，實際只有一次性證據；差異在文件上不可見
- 後續變更（例如編輯那張推導表）不再被任何等價性檢查攔住，而所有人都以為有
- 若該條款被 Verifier 判為 pass，pass 是正確的（當下確實驗過），但它同時**掩蓋**了保證已消失這件事

## Required Agent Action

**轉換型切片的驗收條款要區分「當下驗過」與「持續守著」，並在條款上標明。**

1. **撰寫 brief 時判斷每個條款的存活期。** 若條款的驗證依賴切片將要移除的東西，標為 `transitional`。
2. **`transitional` 條款必須指名接手的常駐不變式。** 「等價性由 X 驗過一次；此後由 Y 守著另一個不變式」—— 並明說 X 與 Y 不是同一件事。
3. **一次性證據要留在可追溯處。** 執行紀錄寫下比對結果與發現的分歧，而非只寫「已驗證」。
4. **不要為了讓檢查繼續存在而弱化它。** 把等價性比對改寫成較弱的斷言並保留原檔名，比讓它消失更糟：檔名會讓人以為保證還在。若要保留，換名並說明它現在斷言什麼。
5. **關閉切片時複查每個 `transitional` 條款。** 確認紀錄裡有當下的證據，且接手的不變式確實存在且會失敗。

## Prevention Gate

寫下「轉換前後等價，機械驗證」這類條款時：

- 比對的另一側是什麼？這個切片會移除它嗎？
- 移除之後這個檢查還剩什麼？它會變成空轉，還是被改寫成更弱的斷言？
- 誰接手守著？那個不變式和等價性是同一件事嗎（通常不是）？
- 一年後讀到這條款的人，會以為有守衛嗎？文件上看得出沒有嗎？

## 驗證

1. Brief 中每個依賴「被移除的一側」的條款標為 `transitional`，並指名接手的常駐不變式
2. `transitional` 條款的一次性證據記錄於執行紀錄（含比對結果與發現的分歧），不只寫「已驗證」
3. 關閉紀錄明確區分「當下驗過」與「此後由何者守著」，且該區別對未讀過對話的人可見
4. 沒有任何測試以原名保留卻已改為斷言較弱的命題

## Linked Rules

- [`cross-boundary-agreement-never-mechanically-checked.md`](cross-boundary-agreement-never-mechanically-checked.md) — 空轉檢查與 mutation 要求
- [`proxy-signal-substituted-for-target-state.md`](proxy-signal-substituted-for-target-state.md) — 上位家族：以間接訊號替代目標狀態
- [`refactor-parity-feedback-miss.md`](refactor-parity-feedback-miss.md) — 同型：replacement／refactor parity 的回饋落點
- [`../failure-learning-system.md`](../failure-learning-system.md) — `validation-gap` 分類
