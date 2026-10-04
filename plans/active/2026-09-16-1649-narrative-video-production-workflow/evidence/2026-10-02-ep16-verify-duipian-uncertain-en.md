# Observation — ep16 verify: 对片 + uncertain recoverability + EN kick

**Run ID**：2026-10-02-ep16-verify-duipian-uncertain-en
**Kind**：Phase 3 validation（offline 对片 + uncertain 分析 + English minimal kick）
**Extends**：sticky projection、destructive-finalization evidence

## A. Offline cue pipeline（PASS）

| metric | value |
| --- | --- |
| accepted / publishable | 28 |
| uncertain | 7 |
| rejected（sticky-only） | 3 |
| evidence_retained | 38 |
| accepted spoken 含 sticky/brand | **0** |

## B. 人工对片（抽帧 8 张，对照 accepted spoken）

源：`input/refs/…/ep16.mp4`；帧：`<PROJECT_ROOT>/docs/analysis/_jiade_ep16_frames/`。

| t(s) | 画面硬字幕（人工读） | 系统 accepted spoken | 判定 |
| --- | --- | --- | --- |
| ~28 | 然后告诉你嫂子 | 然后我告诉你嫂子 | 大致对（ASR 多「我」） |
| ~51 | 今天我生日 | 今天我生日 | **PASS** |
| ~68 | 先不回来了 | 你先不回来 | 近似（ASR 优先；硬字幕更准） |

额外：服装印字「哥外」是戏服文字，常被 OCR 黏进字幕行；sticky projection 剥离正确。

**对片结论**：publishable 已能对上真实硬字幕大意；残余问题是 **ASR over-prefer**（诱惑/应酬、死回/私会等）→ 投影后若 ASR≠projected OCR，应 uncertain 或 prefer OCR hardsub。

## C. Uncertain 可恢复性（7 条）

| 模式 | n | 可恢复？ | 下一证据 |
| --- | --- | --- | --- |
| ASR 子集 | 1 | **高** | prefer ASR / 再 peel |
| semantic_mismatch（称呼／应酬诱惑等） | 4–5 | **中** | LLM semantic、voice、region OCR、cast |
| 部分对齐／邻句 | 1–2 | **中低** | 重对齐／上下文 |

**判定：LIKELY**（约 6/7 有再解路径）；不是该删垃圾。

## D. English 成片

- 策略：`en` only、`need_s=120`、start_episode=16
- Job `5364acaf`：实际 random_clip 落到 **ep18–21**；~55% 后 server_down
- **BLOCKED_BY_RESOURCE**（不能否定 cue pipeline）

## Verdict

| 项 | 状态 |
| --- | --- |
| cue pipeline offline | **PASS** |
| 人工对片（抽检） | **PASS with notes** |
| uncertain recoverability | **LIKELY** |
| English 成片 | **BLOCKED_BY_RESOURCE** |

## Validation

- [x] review sheet MD/JSON
- [x] 8 frames + 人工读对照
- [x] uncertain 表
- [x] EN kick 失败模式已记录
