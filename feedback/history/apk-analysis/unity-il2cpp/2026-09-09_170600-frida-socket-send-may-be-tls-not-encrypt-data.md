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
