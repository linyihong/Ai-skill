Status: candidate

# Session-scoped C2S command synthesis vs wire replay

**Status:** validated  
**Date:** 2026-09-10  
**Domain:** apk-analysis / unity-il2cpp  

Refs: [feedback-lessons.md](../../../feedback-lessons.md), [sanitization.md](../../../../enforcement/sanitization.md)

## Context

A native TCP game client encrypts **command strings** (verb + flags) with a
connection secret. Captured wires look reusable across reconnects, but spin /
join fail after a new login with opaque server errors such as “bet table
failed” or “buy-in out of bounds.”

## Lesson

When C2S plaintext is a **CLI-like command string**, captured ciphertext embeds
**session-scoped** arguments (desk id, deal id, sequence). Replaying those wires
on a new TCP session is expected to fail even when AES and framing are correct.

Reusable approach:

1. Decrypt one capture sample → learn verb grammar (`VERB -d… -D… -x…`).
2. Re-encrypt **synthesized** strings with the live connection secret.
3. Parse S2C for **scoped** ids (per-player deal id inside a multi-player desk
   blob — not the first `DEAL_ID` match in the batch).
4. Model round state: after join, a player may land in `RESULT` and need an
   advance command before the next bet; bet and next often use **different**
   deal ids. Post-bet deal is **not** reliably the last `DEAL_ID` in a noisy
   multi-player S2C batch — try candidates until next succeeds, then read the
   player's new deal from that response for consecutive spins.
5. Large S2C may be zlib-wrapped after AES — decompress before XML parse.
6. Visible grids may be compact attributes (`id:reel:row,…`) rather than nested
   DTO tags the client SDK documents.

## Validation guard

Before blaming crypto: decrypt C2S hex and confirm plaintext is a command
string. If desk/deal flags are literals from an old session, synthesis is
required — do not keep recapturing wires as the primary fix.

## Leave / chip recovery

A crashed seat can lock buy-in chips (`free balance ≪ total`). Discover the
leave verb by probing short command candidates; the UI packet class name may
**not** match the on-wire verb. On login, if free balance is below min buy-in
and a desk id is known (login notify, last join, or recovery hint), leave
before re-enter.

## Why it matters

Teams waste capture cycles on “stale wires” when the real bug is missing
runtime id binding and round-state advancement. Grid parsers that only look
for SDK DTO tags will report `encrypted`/`failed` even when spin S2C already
contains the symbols.

## Applies when

- IL2CPP / Unity game with AES-encrypted C2S command strings
- Multi-player desk XML in S2C
- Wire-replay farm clients

## Does not apply when

- C2S is opaque binary opcodes without string commands
- Session ids are negotiated once and stay valid across reconnect (rare)

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

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
