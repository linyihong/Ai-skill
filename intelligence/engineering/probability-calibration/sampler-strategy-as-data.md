# Sampler Strategy as Data（抽樣變異點做成資料 + 註冊式 Strategy）

**Status**: `candidate-intelligence`
**Evidence**: N=1（單一 slot consumer 的取樣機重構）
**Source**: plan Inventory I8；profile 層的具體策略分類見 [`workflow/probability-calibration/profiles/slot-rtp.md`](../../../workflow/probability-calibration/profiles/slot-rtp.md)

## 原則

需要對多個產品或玩法做校準時，把抽樣機的**變異點**（抽樣方式、權重、耦合規則）做成參數化資料，抽樣行為以註冊式 Strategy 選擇；校準只改資料，不改引擎程式碼。重構時以既有 consumer 的固定 seed 釘值作為「零 RNG 漂移」護欄；缺參數時在載入階段 fail-fast，而不是以預設值靜默執行。

本 atom 是設計模式說明，不附程式碼，也不指定語言或引擎。

## 判斷訊號

- 每換一個產品或玩法就要改抽樣程式碼
- 校準參數散落在程式碼常數中，無法被 calibration record 的 `parameter_snapshot_hash` 指到
- 重構後無法證明未改變既有產品的隨機序列

## 單一案例觀察（N=1，非通用規則）

- 該案把條帶式、加權格位、耦合條帶等策略註冊為同一介面；重構期間以其他產品的釘值確認 RNG 序列未漂移。

## 適用條件

- 有兩個以上產品或玩法共用同一抽樣基礎
- 校準需要重現與比對參數快照

## 反例／不適用

- 只有一個產品、短期原型：直接寫死可能更便宜
- 抽樣行為需要任意程式邏輯、無法用有限參數表達：強行資料化只會造出一套隱形腳本語言
- 釘值護欄只證明「沒漂移」，**不**證明校準正確（見 verification 契約的驗證分層）

## 與既有知識的關係

- 釘值≠校準：[`validation-proxy-trap`](../anti-patterns/validation-proxy-trap.md) 與 failure class `pinned_value_claimed_as_calibration`
- 可重現性欄位：C9 `reproducibility_reference`（契約）

## 升級條件

≥3 個獨立 consumer 採用此模式並各自留下零漂移護欄證據，且至少一例不是 slot。
