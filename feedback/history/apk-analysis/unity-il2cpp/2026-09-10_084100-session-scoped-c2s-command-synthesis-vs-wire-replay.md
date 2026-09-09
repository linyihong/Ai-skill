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
   deal ids — take the post-bet deal from bet S2C, do not reuse the bet input.
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
