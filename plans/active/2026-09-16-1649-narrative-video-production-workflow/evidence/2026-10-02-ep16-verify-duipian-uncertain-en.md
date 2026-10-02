# Observation — ep16 verify: 对片 + uncertain recoverability + EN kick

**Run ID**：2026-10-02-ep16-verify-对片-uncertain-en
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

源：`input/refs/…/ep16.mp4`；帧目录：`<PROJECT_ROOT>/docs/analysis/_jiade_ep16_frames/`。

| t(s) | 画面硬字幕（人工读） | 系统 accepted spoken | 判定 |
| --- | --- | --- | --- |
| ~28 | 然后告诉你嫂子 | 然后我告诉你嫂子 | **大致对**（ASR 多「我」） |
| ~51 | 今天我生日 | 今天我生日 | **PASS** |
| ~68 | 先不回来了 | 你先不回来 | **近似**（用了 ASR；硬字幕更准） |

额外观察：服装印字「哥外／强哥外卖」是 **戏服文字**，常被 OCR 黏进字幕行；sticky projection 剥离开是对的（projection 非 delete）。

**对片结论（样本级）**：publishable 路径已能对应真实硬字幕大意；仍有 **ASR 优先过度** 个案（诱惑/应酬、死回/私会、先不回来了/你先不回来）→ follow-up：投影后若 ASR 与 projected OCR 冲突，应 uncertain 或 prefer projected OCR（硬字幕）。

## C. Uncertain 可恢复性（7 条）

| # | 现象 | 可恢复？ | 下一证据 |
| --- | --- | --- | --- |
| 0 | ASR 子集（算了 ⊂ 更长 OCR） | **高** | prefer ASR／再 peel |
| 1–4 | semantic_mismatch（安哥/袁晏/诱惑/唐哥） | **中** | LLM semantic、voice、region OCR、cast |
| 5 | 嫂子…日子 vs 哇嫂子 | **中低** | 邻句上下文／再对齐 |
| 6 | 同文 OCR vs 完全不同 ASR | **中** | 时间窗错位？重对齐 |

**判定**：`LIKELY`（约 6/7 有明确再解路径）；**不是**该删的垃圾。

## D. English 成片

- 策略：单集焦点 ep16、`commentary_langs=[en]`、`dub_clip_need_s=120`（防 OOM）
- Job：见 kick 日志／`_jiade_ep16_en_kick.json`（本 run 启动后轮询）

## Verdict

| 项 | 状态 |
| --- | --- |
| cue pipeline offline | **PASS** |
| 人工对片（抽检） | **PASS with notes**（ASR over-prefer） |
| uncertain recoverability | **LIKELY** |
| English 成片 | **IN_PROGRESS / 见 kick 结果** |

## Validation

- [x] review sheet MD/JSON
- [x] 8 frames + 人工读 3 关键对照
- [x] uncertain 表
- [ ] EN job final outputs（轮询后补勾）
