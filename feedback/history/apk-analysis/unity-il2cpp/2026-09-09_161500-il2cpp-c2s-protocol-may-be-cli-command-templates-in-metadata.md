# IL2CPP C2S 協定可能是 metadata 字串表裡的 CLI 命令模板

**Status:** validated
**Date:** 2026-09-09
**Domain:** apk-analysis / unity-il2cpp

## Context

分析一個 Unity IL2CPP 遊戲的自訂 TCP session 協定時，先前假設 C2S request
是 XML／protobuf，因為 `ParseResult(XContainer)` 與 `XDocument` 出現在 send
路徑附近。實際上 `XContainer`/`XDocument` 只用於 **S2C response**。

## Lesson

對 IL2CPP 客戶端，C2S 請求的明文常常是 **command-line 風格的命令字串**
（例如 `VERB -n{0} -Y{1}`、`ACTION -d{0} -b{1} -a{2}`），`{n}` 是位置參數。
這些模板通常以 **完整字面量** 存在 `global-metadata.dat` 字串表裡，可用
Il2CppDumper 產出的 `stringliteral.json` 直接撈出來，不需要反編譯 native
body 或動態 hook：

- 在 `stringliteral.json` 搜 `^[A-Z][A-Z0-9_]{2,}( -|\s|$)` 加上 ` -[A-Za-z]\d?(\{\d+\}|\s|$)` 這類 flag pattern，命令表會整批浮現。
- 命令 verb 與 flag 名是 **constant**；真正的變數只有 `{n}` 填入的值。
- 共用尾旗標（例如所有 login 都有的某個 flag）通常由 base packet 的
  `OnSendPacket` 附加，會讓 wire 比 verb 模板長很多。

## Why it matters

先看字串表能在幾分鐘內拿到整個協定的 opcode/flag 結構（request 端 schema），
把「reconstruct plaintext」從反編譯工作降級成查表；只剩 **加密金鑰推導**
與 **少數變數值** 需要 Ghidra RVA 反編譯或 names/length-only 動態驗證。

## Guardrails

- 只保留 verb/flag schema；不要把真實填入值、token、device id、secret、
  金鑰或完整 wire hex 寫進可重用 lesson 或 commit。
- 字串表證明「模板存在」，不證明某封包當下用哪個模板；仍需與 send-path
  packet class（`SendPacket` arg class）或動態證據對齊。

## Origin

Static Il2CppDumper（metadata v31）對一個 casino 類 IL2CPP client 的授權分析。
Target-specific 命令表、host、金鑰留在該專案 evidence，不進本 lesson。
