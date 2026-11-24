# Fixed Offset Object Read Without Metadata Check（沿用固定偏移讀取 live 物件而未驗證 metadata）

Status: candidate
Class: `validation-gap`

## Trigger

Agent 要從 live runtime 物件讀取欄位（例如 IL2CPP、Unity、JVM 或其他有 metadata 的 runtime），並且出現任一情況：

- 名稱比對回傳 null，agent 轉而沿用其他 probe 或其他類別的固定偏移。
- 讀到的數值全為 0 或與畫面已知值不符，但沒有 error。
- Probe 在解引用後讓 target process 崩潰或 attach 失敗，agent 立即重跑同一支 probe。

## Failure Mode

Agent 用「別處驗證過的偏移」替代「此處 metadata 中的實際偏移」，並直接解引用。

- 偏移來自另一個類別或另一個 build，在當前類別上指向相鄰欄位（陣列、集合或其他物件）。
- 解引用成功，回傳合法但錯誤的數值，沒有例外，因此 agent 把它當成「真實值為 0」或「功能尚未觸發」。
- 讀取同時伴隨其他危險操作（dump 集合、遍歷 dictionary），讓 target process 崩潰。崩潰後 agent 又在不確認 target 狀態的情況下重新 attach，形成重複失敗。

共同點：agent 沒有先查詢目前 class 的 metadata，也沒有驗證指標所指物件的類別。

## Risk

- 把錯誤數值寫進 manifest、commit、goal 或交接文件，成為後人採信的錯誤事實（例如錯誤的 enum id、錯誤的線數）。
- Target process 崩潰，中斷正在進行的狀態（例如 live 遊戲場景、已進入的功能），且之後的探測結果不可比較。
- 把「沒讀到」誤判為「不存在」，因而錯過正確的欄位名稱。

## Required Agent Action

1. **先列 metadata，不讀物件。** 對目標 class 沿繼承鏈列出所有欄位（名稱、偏移、型別）。以名稱搜尋時，同時考慮自動屬性的 backing field 形式。
2. **偏移只能來自當前 class 的 metadata。** 不得直接複製其他 probe 的偏移；若別處的偏移是唯一線索，必須先在當前 class 的 metadata 中找到同一欄位名稱才能使用。
3. **解引用前驗證物件類別。** 讀取指標後，確認其 runtime class 名稱與預期相符；不符就停止，不讀值。
4. **Probe 只做一件事。** 讀純量時不順便 dump 集合或 dictionary。
5. **失敗後先確認狀態，不重跑。** Probe 導致 attach 失敗、target 崩潰或結果異常時，先確認 target process 是否仍存在、畫面是否仍在預期狀態，並記錄哪一步的假設不成立；不得以換偏移或重複 attach 方式直接重試。
6. **宣告值之前先交叉驗證。** 數值若要寫進文件，必須有第二個來源（enum 名稱、已知資產名稱、metadata 偏移與 class 驗證）支持。

## Prevention Gate

- **文件層（已存在）：** 本 pattern 與 [`feedback/history/apk-analysis/unity-il2cpp/`](../../feedback/history/apk-analysis/unity-il2cpp/README.md) 的對應 lesson，供 `unity-il2cpp` 類任務在進入 probe 前讀取。
- **Validator 層（待定）：** 理想上 probe 模板應強制三段式：metadata 列表 → class 驗證 → 純量讀取。目前尚無機械檢查；在 validator 建立前，此 pattern 僅為行為要求，Status 保持 `candidate`。
- **Promotion 條件：** 第二個獨立目標再次觀察到同類失誤，且 validator 草案存在時，才考慮提升為 `validated`。

## 驗證

- 在 probe 執行前，agent 的記錄中應出現 metadata 欄位列表（含偏移與型別）。
- 解引用前的 class 名稱驗證結果應可追溯。
- 數值進入文件前，應能指出第二個支持來源。
- 失敗後的回應應包含 target 狀態確認，而不是直接重跑。

## Linked Rules

- [`feedback/history/apk-analysis/unity-il2cpp/2026-11-24_103116-il2cpp-field-name-and-offset-verify-before-deref.md`](../../feedback/history/apk-analysis/unity-il2cpp/2026-11-24_103116-il2cpp-field-name-and-offset-verify-before-deref.md)
- [`failure-learning-system.md`](../failure-learning-system.md)
- [`proxy-signal-substituted-for-target-state.md`](proxy-signal-substituted-for-target-state.md)（同屬「用代理訊號替代目標狀態」的家族）
