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
