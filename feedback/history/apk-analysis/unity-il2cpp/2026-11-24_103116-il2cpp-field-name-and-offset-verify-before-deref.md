> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-11-24 - IL2CPP object reads: verify field name and offset against metadata before dereferencing

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。目前只在單一 Unity IL2CPP 目標的單一 session 中觀察到，並已用同一目標的 metadata 與 runtime class 名稱交叉驗證。

#### One-line Summary

Frida 讀 IL2CPP 物件欄位時，先用 metadata 確認欄位名稱、型別與偏移，並在解引用前驗證指標所指物件的 class 名稱；不要沿用其他 probe 的固定偏移，也不要用自動屬性的顯示名稱去比對。

#### Human Explanation

Unity 的 C# 自動屬性在 IL2CPP metadata 裡通常是 `<Name>k__BackingField`，不是 `Name`。用顯示名稱去比對會得到 null，看起來像「欄位不存在」，但其實只是名稱不符。更危險的是固定偏移：同一 build 裡，類別的相鄰欄位可能是陣列、集合或另一型別的物件。若直接沿用其他類別的偏移，解引用後會讀到合法但錯誤的數字（例如 0），不會報錯，也可能在後續 dump 集合時讓目標 process 崩潰。

#### Trigger

- 需要從 live IL2CPP 物件讀一個 enum、int 或 flag，但先前 probe 用名稱比對回傳 null。
- 準備把其他類別或其他 build 的固定偏移複製進新 probe。
- 讀到的數值全為 0 或與已知畫面值不符，卻沒有任何 error。

#### Evidence

- Tool: Frida（IL2CPP runtime API：class／field metadata、`field_static_get_value`、`object_get_class`、`class_get_name`）
- Sanitized excerpt: 同一控制器的 backing field 名稱與名稱比對不符；繼承鏈上的陣列欄位與單一物件欄位位置相鄰，錯用偏移時得到全 0 的數值。enum 字面值可由 static literal 直接映射為名稱。
- Evidence path: `<PROJECT_ROOT>/docs/slots/` 專案文件（具體目標與數值留在專案，不在本庫）

#### Generalized Lesson

對 IL2CPP 物件的欄位讀取分三步，且第一步必須不讀物件：

1. **Metadata-only 列舉**：沿著 `class_get_parent` 走完整條繼承鏈，列出每個欄位的名稱、偏移與型別。以 `<Name>k__BackingField` 形式搜尋自動屬性。這一步只讀 class metadata，不碰 live object，不會因錯誤指標而崩潰。
2. **Enum 名稱映射**：若欄位是 enum，用 static literal 讀取（`il2cpp_field_static_get_value`）建立數值到名稱的表，不要憑字串猜測。
3. **解引用前驗證 class**：從欄位取得指標後，先用 `object_get_class` + `class_get_name` 確認 class 名稱符合預期，再讀取純量（int、enum）。Probe 內只讀純量；不要在同一個 hook 中順便 dump dictionary、array 或集合，避免一次錯誤把 target process 拖垮。

若任一步驗證失敗，停止該 probe，回到第 1 步重新確認 metadata，不要換偏移再試。

#### Agent Action

- 看到「欄位名稱比對為 null」時，先列出 metadata，再決定是名稱差異（backing field）還是欄位不存在。
- 不要把其他 class 或其他 build 的固定偏移當作事實；偏移必須來自目前 class 的 metadata。
- 讀取 live object 前，先做 class 名稱驗證；驗證失敗就不讀值。
- 同一個 probe 只做一件事（讀純量）；dump 集合與讀純量分開跑。

#### Goal / Action / Validation

- Goal: 從 live IL2CPP 物件取得正確的 enum／int 值，且不讓 target process 因錯誤解引用而崩潰。
- Action: metadata-only 列舉 → enum 名稱映射 → class 驗證後讀純量。
- Validation or reference source: 讀到的 enum 名稱與 pack／資源名稱一致；讀到的數值與 class 驗證結果一致；metadata 偏移與 runtime 讀值的物件類別相同。

#### Applies When

- Unity IL2CPP target，且要讀的欄位是 C# 自動屬性或繼承自基底類別的欄位。
- 同一專案內有多個相似類別（例如多個遊戲控制器共用基底）。

#### Does Not Apply When

- 目標是非 IL2CPP 的 runtime（例如 Java／ART 層）。
- 同一 build 已由 metadata 確認某類別的偏移，且 probe 已在同一 class 上驗證過，可直接沿用。
- 只需要 method 名稱或 class 存在性，不涉及欄位讀值。

#### Validation

- Metadata 列表包含目標欄位名稱、偏移、型別，且欄位名與 `<Name>k__BackingField` 一致。
- Runtime class 名稱驗證通過後才讀取。
- 讀到的值在 enum 映射表中有對應名稱，且名稱與專案已知資產對得上。
- 若 probe 導致 target 崩潰，重新啟動後從第 1 步重跑，並記錄哪一步的假設不成立。

#### Promotion Target

- `analysis/unity/workflows/`（待單一 technique 證據足夠後，再評估是否新增 procedure）
- `enforcement/failure-patterns/fixed-offset-object-read-without-metadata-check.md`（agent 失誤面）

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- 更新 [`unity-il2cpp/README.md`](README.md) 索引列。
- 更新 [`apk-analysis/README.md`](../README.md) 分類數量。
- 新增 failure pattern（見 `enforcement/failure-patterns/fixed-offset-object-read-without-metadata-check.md`），並在 `enforcement/failure-patterns/README.md` 加入索引。
- 已依 [`reusable-guidance-boundary.md`](../../../../enforcement/reusable-guidance-boundary.md) 抽象化：目標遊戲名稱、類別名、數值、裝置與路徑都只留在專案文件。

#### Transfer Contract

- Inputs: 一個 live IL2CPP process、目標物件的類別名稱、要讀取的欄位名稱（可能是自動屬性）。
- Mechanism: IL2CPP metadata 以 backing field 名稱保存自動屬性；固定偏移只在同一 class 與同一 build 內有效；錯誤指標會合法地讀到錯誤數值，沒有例外。
- Decision: 先做 metadata 列舉，再做 enum 映射，最後在 class 驗證後讀純量。任一步失敗就停止，不換偏移重試。
- Verification: 欄位在 metadata 中可找到；runtime class 名稱正確；enum 名稱與專案資產一致。
- Counterexample: 同一 class 的偏移已由 metadata 確認且已驗證過的情境，沿用不算錯。
- Transfer case: 任何 Unity IL2CPP 遊戲在讀取繼承欄位或自動屬性時，都可套用同樣的三步順序。

#### Intelligence Review

- Decision: deferred
- Reason: 目前只有單一目標、單一 session 的觀察；依 candidate 規則，不應提升為 intelligence atom。
- Target: none
- Revisit: 第二個獨立 Unity IL2CPP 目標再次出現「名稱比對為 null」或「固定偏移讀到全 0」時，重新評估是否 extract 到 `analysis/unity/workflows/`。

#### Abstraction Check

- Swap-project test: PASS — 拿掉遊戲名稱、類別名與數值後，「先 metadata、再 enum 映射、最後 class 驗證後讀純量」的規則仍成立。
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/` 專案文件
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
