Status: candidate

# Frida socket `send` 捕獲可能是 TLS record，不是 EncryptData 明文/密文

**Status:** validated
**Date:** 2026-09-09
**Domain:** apk-analysis / unity-il2cpp

## Context

從官方客戶端抓自訂遊戲 TCP 的 C2S 時，hook `libc.send` 並把 payload 當
AES 密文重放／解密。捕獲開頭是 `17 03 03`，且後兩位長度剛好等於剩餘
bytes。用遊戲 KDF 解 AES-ECB 會失敗。

## Lesson

`send()` 看到的是 **socket 寫入層**。若連線是 `tcps`／SSL wrapper，那是
TLS Application Data record，不是 `NetState.EncryptData` 的輸出。

判斷與正確抓法：

- TLS record：`17 03 03` + 2-byte length == remaining payload.
- 協定 schema 應 hook **EncryptData(string)** 或等價序列化點，只記錄
  command verb、flag 字母、值長度／形狀，不要記 token。
- Farm／replay 模板若來自 `send()`，長度會比內層 AES blob 多 5 bytes
  TLS header，且無法用遊戲 AES key 還原 CLI 字串。

## Why it matters

把 TLS record 當遊戲密文會讓 KDF「驗證失敗」、旗標字母永遠對不上，並
把裝置 mint 卡在錯誤層。先分類捕獲層再決定 decrypt vs schema parse。

## Guardrails

- 不要把完整 wire hex、token、device id 寫進 lesson。
- 長度與 flag 字母可以公開；值只能用 `<len:N>` 或常數 APK feature 列表。

## Origin

Pokerist 75.8.0 farm LoginGuest capture vs EncryptData schema probe.

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
