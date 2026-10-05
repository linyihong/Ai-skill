# Observation — EN dogfood: source-lang gate + skip Qwen on gemini

**Run ID**：2026-10-05-en-dogfood-source-lang-vs-qwen
**Kind**：Phase 3 adapter／validation（接續 ep16 verify；true ep16-only English clip）
**Extends**：[`2026-10-02-ep16-verify-duipian-uncertain-en`](2026-10-02-ep16-verify-duipian-uncertain-en.md)、[`2026-10-02-sticky-watermark-accepted-projection`](2026-10-02-sticky-watermark-accepted-projection.md)

## 問題

1. **Local Qwen OOM**：`dub_translate_provider=local`（或 plot JSON 強制本地）會載入 Qwen3.5-9B；16GB 級 GPU 在 `Loading weights` 階段崩潰，與 clip 長度無關。
2. **源語被目標語閘掉**：`resolve_fused_to_cues` 用 `resolve_lang(cfg)`（commentary=`en`）做 script gate → `丢弃非目标语口播(en): N`，中文口播全進 rejected，publishable=0，成片 0 輸出。

## 規則

1. **Commentary lang ≠ source spoken lang**。OCR/ASR 源口播 script gate 用 `dub_asr_language`／源語（中文硬字幕預設 `zh`），禁止用 burn／commentary 目標語。
2. **`provider=gemini` 必須真的不載 Qwen**：含 `complete_json_for_plot`；不得「plot 仍本地優先」。
3. EN dogfood 最小設定：`dub_translate_provider=gemini`、`dub_episode_from=to=16`、關 virtual_streamer／progress_gif；Evidence≠Production 不變。

## Validation

- [x] offline audit：accepted 21／uncertain 14／rejected 3（prefer-OCR 後）
- [x] 清掉 rejected-only cache 後 rebuild：`cues=36`（不再 `丢弃非目标语口播(en)`）
- [x] job `2059846d` done：成功單元 1、產出 1 成片（`The Temptation of Home_…mp4`，約 4MB）
- [x] 伺服器全程存活（不再 model-load 崩潰）

## 產品落點

- `json_llm.complete_json_for_plot`：`provider=gemini` 只走 Gemini
- `spoken_text_resolution.resolve_fused_to_cues`：source script gate → `dub_asr_language`
