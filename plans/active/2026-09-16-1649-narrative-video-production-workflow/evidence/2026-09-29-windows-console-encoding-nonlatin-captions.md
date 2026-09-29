# Evidence — Windows console encoding must not abort non-Latin caption jobs

Status: observed  
Date: 2026-09-29  
Linked plan: narrative-video-production（caption／locale pack consumer）

## Finding

On Chinese Windows, the default console encoding is often GBK／cp936.  
Diagnostic `print` of cue／title text in a non-Latin target locale（e.g. Thai）raises `UnicodeEncodeError` and can abort an entire locale export even when caption files themselves are written as UTF-8.

## Rule

- Caption／ASS／SRT **file I/O** stays UTF-8.  
- **Console／log prints** that may include target-locale text must be encode-safe（replace／omit），never job-fatal.  
- Content correctness gates（Translation Decision／content_gate）are unrelated to console codec failures.

## Product follow-up

Product dub ASS path adopted a safe print helper for fit／build diagnostics so Thai（and other non-GBK）locales continue.

## Sanitization

No hosts、paths、or episode payloads.
