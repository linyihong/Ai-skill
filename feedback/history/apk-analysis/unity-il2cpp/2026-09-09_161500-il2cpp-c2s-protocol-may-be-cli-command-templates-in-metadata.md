Status: candidate

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
