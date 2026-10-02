# Text evidence — Evidence Acquisition & Escalation Loop

何時讀：OCR／ASR／Visual 已能各自 probe，但「第一次沒找到」被當成不存在；
跨模態密度不合理（例 ASR 對白密、OCR 幾乎空）；或 Story 端才發現事件過少、才回頭懷疑採集。
Plan：[`46-evidence-acquisition-escalation-loop.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/46-evidence-acquisition-escalation-loop.md)。
銜接：[`text-evidence-ocr-discovery.md`](text-evidence-ocr-discovery.md)、
[`text-evidence-subtitle-candidate.md`](text-evidence-subtitle-candidate.md)、
[`text-evidence-multimodal-resolution.md`](text-evidence-multimodal-resolution.md)、
[`12-evidence-refinement.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/12-evidence-refinement.md)。

> **執行契約，不重新解釋契約。** Probe 負責發現；**Candidate Detector 負責 precision**；Monitor 負責可疑；Escalation 負責加大採集；
> LLM 只在 escalation 後、仍衝突／不足時做 arbitration。不開新 Phase／新 Agent／runtime。

## 一句話

**Probe 不負責證明不存在；Probe 負責發現證據。** 當 OCR、ASR、Visual 或跨模態出現不合理缺口／衝突時，必須先 escalation，改變採集策略；只有 evidence 仍不足或語義衝突時，才交 LLM。
**但「畫面有文字」≠字幕存在**——coverage／escalation 只在 `subtitle_like` 候選上計算（見 Candidate Detector）。
**且 Source-level evidence ≠ window-level evidence**——全片 `hardsub=True`／cues cache hit 不能單獨滿足「這個 clip 時間窗有可用對白」（見下方 Scope invariant）。

## 位置（在 6b 最前）

```text
Source profiling → Initial acquisition (OCR | ASR | Visual)
  → Subtitle Candidate Detector  (subtitle_like | non_subtitle | uncertain)
  → Evidence Monitor  (confirmed | inconclusive | suspicious)  // densities on subtitle_like
       ├─ sufficient → Text Resolution / Story consumers
       └─ suspicious → Escalation Policy (Level 0–4) → re-acquire → Monitor again
```

本檔管 **acquisition 充分度**；[`text-evidence-multimodal-resolution.md`](text-evidence-multimodal-resolution.md) 管 observed→candidate→resolved。
[`12-evidence-refinement.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/12-evidence-refinement.md) 是更外層 linking／story；本檔是其 acquisition 閘。

## 三態（每模態）

| status | 含義 | 不得做的事 |
| --- | --- | --- |
| `confirmed` | 覆蓋與 expected 一致、跨模態無高嚴重缺口 | — |
| `inconclusive` | 目前假設區／參數未命中（例 region miss） | 不得寫 `exists=false`／STOP |
| `suspicious` | 機械 anomaly（跨模態缺口、語言錯配、layout drift…） | 不得直接進 Story；必須 escalation 或明確 defer |

單模態 `subtitle_like=0` 最多是 `inconclusive`；若同時 ASR dialogue 高 → 升為 **OCR `suspicious`**（`ocr_asr_coverage_gap`）。

## Expected evidence（先於 anomaly）

| 源型（profiling） | expected OCR | expected ASR |
| --- | --- | --- |
| 硬字幕劇 | high | supporting |
| 純口播無硬字幕 | low | high |
| 雙語硬字幕 | high + multi-script | supporting |
| 外語字幕＋中文口播 | latin／other high | zh supporting |

`OCR low`  alone ≠ anomaly。`expected OCR=high` ∧ `OCR low` ∧ `ASR high` → anomaly。

## Mechanical suspicion signals（非窮盡）

| id | 條件（摘要） | 典型 escalation |
| --- | --- | --- |
| `ocr_asr_coverage_gap` | ASR dialogue density 高、**subtitle_like** OCR 極低（整段或時間窗）；忽略 clock／mail／UI | widen OCR region／↑ sampling／targeted window |
| `ocr_scene_text_only` | OCR 有字但全是 non_subtitle_like | **no_escalation**（或僅記 scene_text）；不得當 recovered dialogue |
| `asr_ocr_coverage_gap` | OCR 對白密、ASR 幾乎空 | ASR retry／lang／segmentation |
| `asr_language_mismatch` | OCR 大量 CJK、ASR 判成 English（或反向） | language redetect／retry |
| `layout_drift` | 對白 cy／band 中段突變 | rediscovery／reprofile |
| `latin_boundary_suspicious` | 黏字串／假拉丁 | character-level／geometry recovery |
| `visible_text_without_dialogue_role` | boxes 多但 dialogue_candidate=0 | role／layout expand（非 exclusion） |
| `text_resolution_conflict` | OCR↔ASR 語義衝突簇 | 進 multimodal resolution；必要時 LLM |

## Escalation levels

| Level | 成本 | 例 |
| --- | --- | --- |
| 0 | cheap probe | 少 frame、default regions |
| 1 | broaden | 更多 frame、全畫面／多區 |
| 2 | targeted | upper-middle／center／bbox／時間窗重掃 |
| 3 | enhanced | 更高解析、crop 放大、char-OCR、ASR retry |
| 4 | LLM arbitration | 僅在仍 suspicious／語義衝突 |

**禁止** 把 Level 4 當第一個 probe。每一次 escalation 必須留下 before／after coverage 與 `resolution.status`（`recovered`｜`still_suspicious`｜`deferred`）。

## evidence_coverage（最小欄位）

```yaml
evidence_coverage:
  expected: { ocr: high|low|unknown, asr: high|low|unknown }
  ocr: { status, dialogue_coverage, layout_confidence }
  asr: { status, dialogue_coverage, language_consistency }
  cross_modal: { agreement, signals: [] }
  overall: { status: sufficient|insufficient, escalation_required: bool }
  escalations: [{ trigger, level, action, before, after, resolution }]
```

`overall.insufficient` → **不得**宣稱「本集無字幕／無對白」進入 Story；最多記 `acquisition_blocked`／繼續 escalation。

## Scope invariant（必守）

> **Source-level evidence cannot satisfy window-level evidence requirements.**

| scope | 回答什麼 | 不得用來斷言什麼 |
| --- | --- | --- |
| `source`（全片／全集 profiling） | 這部 source 有沒有硬字幕／口播／語系 | 「這次 clip／時間窗已有可用對白」 |
| `time_window`（`start_s`–`end_s`） | 這次要用的這一段有沒有可用 cues | 「整集沒有字幕」 |

反過來也成立：window OCR 找不到 ≠ source 沒有字幕。

**Cache 命中本身不能代表當前 scope 充分。**
`dialogue_cues` cache hit + `slice` → 0 句，只能記 `cue_coverage.status=empty|sparse`（帶 window scope），不得把 source `hardsub=True` 當成 window 已滿足。

```yaml
cue_coverage:
  status: sufficient | sparse | empty | unknown
  scope:
    type: time_window
    start_s: <float>
    end_s: <float>
  signals:
    cached_cues: <int>
    asr_dialogue: present | none | unknown
    source_hardsub: bool
    source_speech: bool
```

### Window-local fallback（保留全片 probe）

保留全片 probe（source profiling）與 cues cache；**不要**為了 scope 一致而拿掉全片 probe。

```text
source probe → hardsub/speech profile
  → read dialogue_cues cache
  → slice to clip window
       ├─ cues > 0 → continue
       └─ cues == 0
            → assess window_evidence（勿無條件全畫面 OCR）
                 ├─ insufficient（無 ASR／無應有訊號）→ 記錄，可 skip／defer
                 └─ suspicious（source_hardsub ∧ window ASR／應有對白）
                      → escalation: window_ocr（只掃該時間窗）
                      → 更新／補充 cues → 再 slice
```

| window_evidence | decision |
| --- | --- |
| `cached_cues=0` ∧ `asr_dialogue=present` ∧ `source_hardsub=true` | `suspicious` → `window_ocr` |
| `cached_cues=0` ∧ `asr_dialogue=none` ∧ `source_hardsub=true` | 多半 `insufficient`（該窗可能真無對白）；**不得**只因 source hardsub 就全窗重 OCR |
| `cached_cues=0` ∧ 空／壞 cache ∧ `source_hardsub` ∧ `source_speech` | `suspicious` → `window_ocr` |

禁止：把 Level-2／全畫面 OCR 當 window 0 cues 的第一刀。

## Adapter 驗收（產品）

1. 單次 probe miss → `inconclusive`，不是 exclusion。
2. ASR 對白密 + **subtitle_like** OCR 空／窗內缺口 → `suspicious` + 至少一級 OCR escalation，並寫 before／after。
3. OCR 僅 clock／document／UI → 不得觸發「有字＝recovered」；coverage 仍可因缺失 subtitle_like 而 escalate。
4. Story／event 過少不得當第一個 OCR 修復觸發；應在 acquisition Monitor 攔下。
5. LLM 只出現在 Level 4 或 multimodal resolution，不寫死 crop／不斷言無字幕。
6. 每次 escalate／no_escalation 留 `probe_decision.reason`（見 Candidate Detector）。
7. Source `hardsub=True`／cues cache hit **不得**單獨滿足 window；window 0 cues 必須產出 `cue_coverage`（含 scope），suspicious 才 `window_ocr`，再 slice。
8. `slice_subtitles_to_window` 本身不是修復點；修復在 scope 分層與 window-local escalation。
