# Text evidence — Subtitle Candidate Detector

何時讀：OCR escalation／full-frame discovery 已提高 recall，但把鐘錶、郵件、手機 UI、招牌、浮水印當成字幕進 corpus；
或 ASR 對白密時，只因「畫面有任何文字」就擴大 OCR。
Plan：[`47-subtitle-candidate-detector.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/47-subtitle-candidate-detector.md)。
銜接：[`text-evidence-acquisition-loop.md`](text-evidence-acquisition-loop.md)、
[`text-evidence-ocr-discovery.md`](text-evidence-ocr-discovery.md)、
[`text-evidence-regions.md`](text-evidence-regions.md)。

> **執行契約，不重新解釋契約。** Loop 負責 recall；本檔負責 precision。
> 不開新 Phase／Agent／runtime。不把 escalation 拿掉。

## 一句話

**「畫面有文字」≠「字幕存在的 evidence」。**
先做 Subtitle Candidate Detection，再決定要不要擴大探針；非字幕候選不得進 dialogue／subtitle corpus，也不得單獨觸發 OCR escalation。

## 流程（插在 Probe 與 Monitor 之間）

```text
Initial Mechanical Probe / OCR text
        │
        ▼
Subtitle Candidate Detector（mechanical multi-evidence）
        │
   ┌────┼────────────────┐
   ▼    ▼                ▼
subtitle-like  non-subtitle  uncertain
   │         │                │
   │         │                └─ optional expand probe（仍 mechanical）
   │         └─ keep as scene_text；不進 dialogue corpus；不觸發 OCR escalate
   └─ Evidence Monitor（OCR×ASR coverage on subtitle-like only）
            │
            └─ Escalation 僅當「字幕存在 evidence」不足／可疑
```

## 多證據（全部 mechanical；不可單靠 cy）

每個 OCR box 產出 `text_candidate`，**位置是 likelihood evidence，不是硬閘**：

| signal | subtitle_like 傾向 | non_subtitle 傾向 |
| --- | --- | --- |
| `vertical_zone` | lower／upper_middle 帶狀 | 角落小區、任意孤立 |
| `geometry` | 水平文字帶：`aspect_ratio` 高、高度小、寬度適中 | 近方形小塊（鐘錶／icon）、極窄豎條 |
| `text_morphology` | dialogue_like（CJK／問句／對白長度） | `clock_like`／`document_like`／`ui_like`／`url_like`／純數字 |
| `temporal_persistence` | 短窗出現後消失（對白節奏） | 長時間不變（手錶時間、固定 UI） |
| `spatial_stability` | x／y／w 近似穩定 | 隨鏡頭物件大幅漂移 |
| `asr_support` | 與 ASR dialogue 時間／語義可對上 | 無 ASR 對應且 morphology 可疑 |

```yaml
text_candidate:
  text: "…"
  geometry:
    normalized_box: { x, y, w, h }   # 相對 frame；禁止絕對 px 當規則
    aspect_ratio: …
    area_ratio: …                    # bbox_area / frame_area
    center_x: …
    center_y: …
  typography:
    height_ratio: …                  # bbox_h / frame_h
    width_ratio: …
    estimated_font_size_band: small|medium|large   # 由相對高度推估，非 pt/px 閾值
  temporal:
    first_seen: …
    last_seen: …
    persistence_s: …
    text_stability: high|low|unknown      # 同區文字是否幾乎不變
    box_stability: high|low|unknown
    position_stability: high|low|unknown
    replacement_rate: low|medium|high     # 聊天高；硬字幕低
  region:
    region_id: …
    region_role: subtitle_candidate|livestream_chat|platform_ui|watermark|badge|nameplate|title|unknown
  evidence:
    vertical_zone: strong|weak|non_subtitle_like
    geometry: strong|weak
    typography: strong|weak|non_subtitle_like
    morphology: dialogue_like|clock_like|document_like|ui_like|chat_like|unknown
    temporal_persistence: short_burst|sticky|unknown
    asr_support: strong|none|unknown
  exclusion_evidence: []   # 例 small_typography, rapidly_changing_text, avatar_present, badge_present
  status: subtitle_like|non_subtitle_like|uncertain
  reason: []   # 例 clock_pattern, no_subtitle_geometry, livestream_chat_layout
```

## Layout & Typography Evidence（非硬規則）

**禁止**：`font_size > X → subtitle` 或僅 `position: bottom → subtitle`。
**必須**：typography + geometry + temporal + region_role 聯合；單特徵不得定案。

| 特徵組合 | 偏向 |
| --- | --- |
| small band + lower_left 多行 + high `replacement_rate` + 徽章／頭像 | `livestream_chat` → non_subtitle |
| 固定頂欄／側欄 + sticky UI 關鍵字 | `platform_ui` → non_subtitle |
| medium/large band + 相對穩定區 + 同文持續數秒 + dialogue morphology | `subtitle_candidate` |
| 特徵衝突或僅部分吻合 | `uncertain` → 可送少量代表幀給 LLM region classifier |

Negative evidence（為何不是字幕）必須保留在 `exclusion_evidence`，不可當垃圾丟棄。

## Morphology（先機械，不打 LLM）

| pattern | 例 | status 偏向 |
| --- | --- | --- |
| `time_pattern` | `12:35`、`00:32` | clock_like → non_subtitle |
| `email_pattern` / `document_keyword` | `To:` `Subject:` `From:` `@` | document_like |
| `ui_keyword` | `PLAY` `MENU` `暫停` | ui_like |
| `chat_layout` | 多行小字＋用戶名冒號＋高替換 | chat_like → non_subtitle |
| `url_pattern` | `http` `www.` | non_subtitle |
| `numeric_density` 高且無 CJK／字母詞 | 純時間／分數 | non_subtitle |
| dialogue_like | 對白長度＋CJK／標點問句 | subtitle_like |

## Escalation 收緊（相對 acquisition-loop）

| 情況 | 決策 |
| --- | --- |
| OCR 無 text，且無 ASR dialogue／無 prior subtitle_like | `no_escalation`（可 inconclusive） |
| OCR 有 text，但全部 `non_subtitle_like`（clock／mail／UI） | **不**因「有字」escalation；記 `text_candidate_is_non_subtitle` |
| ASR dialogue 高 + 無 subtitle_like candidate | **可** escalation（existence signal 來自 ASR＋缺失 subtitle_like） |
| ASR dialogue 高 + 僅 clock_like OCR | **不**把 clock 當 recovered；仍可因缺失 subtitle_like 而 escalate |
| subtitle_like uncertain（高位／形狀勉強） | widen vertical search／↑ frames（Level 1–2） |

每次決策保留：

```yaml
probe_decision:
  status: escalate|no_escalation
  evidence:
    ocr_text_found: bool
    subtitle_like_count: n
    non_subtitle_counts: { clock_like: n, document_like: n, ui_like: n }
    asr_dialogue: high|low|none
  decision:
    reason: []   # 為何升／不升
```

## 與既有檔邊界

| 檔 | 管什麼 | 本檔 |
| --- | --- | --- |
| ocr-discovery | region miss ≠ 無字幕；layout discovery | 發現後的 **precision／role** |
| acquisition-loop | OCR×ASR coverage Monitor／escalation levels | escalation 前置：**subtitle existence** 與 corpus 過濾 |
| regions／role projection | text_region／subtitle_group | candidate `status` 餵入 role；非字幕不進 dialogue group |

## Adapter 驗收（產品）

1. `12:35`／郵件頭／`PLAY` 不得計入 `dialogue_candidate`／subtitle corpus。
2. ASR 密 + 僅 clock OCR → coverage gap 仍可 escalate，但 before／after 不得把 clock 算成 OCR dialogue recovered。
3. Escalation／no_escalation 都寫 `probe_decision.reason`。
4. LLM 仍只在 Level 4／仍 uncertain 後；不用 LLM 當第一層字幕／雜訊分類。
