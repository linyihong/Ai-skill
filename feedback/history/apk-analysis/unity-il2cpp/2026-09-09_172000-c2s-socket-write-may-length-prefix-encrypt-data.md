Status: candidate

# C2S socket write may length-prefix EncryptData ciphertext

**Status:** validated
**Date:** 2026-09-09
**Domain:** apk-analysis / unity-il2cpp

## Context

A native replay encrypts the same command string the client passes to
`EncryptData(string)` / AES and writes that blob to the game TCP socket.
The server answers S2C hello on connect, then stays silent after the
login/command ciphertext. Idle keepalive on the same socket is one byte
longer than a single AES block.

## Lesson

`EncryptData` is **not** the on-wire PDU. A later `WriteToBuffer` /
`SendPacket` step often **length-prefixes** the ciphertext before
`Stream.Write`.

Typical split (little-endian IL2CPP `byte[].max_length` compared to 240):

- Short: `[u8 cipherLen][aes…]` — keepalive of 17 bytes is `0x10` plus one
  AES-128 block.
- Long: `[0xFF][u32be cipherLen][aes…]` when cipher length **> 240**.

S2C may use a **different** header (for example 4-byte length + session
tag). Do not assume C2S is “raw AES because tcpdump showed no 4-byte
length on small packets.”

Verify by disassembling the send buffer helper: look for `EncryptData`
then `cmp` against `0xF0`, `strb 0xFF`, and `BlockCopy` at offset 1 or 5.

Replay that starts at `libc.send` already includes this prefix (and may
also be TLS). Replay that starts at EncryptData must add it.

## Why it matters

Omitting the prefix makes a large login blob look like an endless
unframed stream; the server never emits the next S2C frame. Matching AES
keys and CLI flags is not enough.

## Guardrails

- Do not commit ciphertext, keys, or device IDs.
- Document prefix rules as lengths and opcode bytes (`FF`, `F0`), not
  payloads.

## Origin

IL2CPP custom game TCP: EncryptData vs WriteToBuffer send path.

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
