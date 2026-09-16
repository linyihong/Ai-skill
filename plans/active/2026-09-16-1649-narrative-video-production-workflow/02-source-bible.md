# Source bible, clip catalog & matching script

Companion to [`_plan.md`](_plan.md)。這層回答三件事：源作品**是誰**、庫裡**有哪些可剪片段**、匹配腳本如何用**同一套文字庫**查出 **可行候選**，再依**明示 selection policy** 選定。不是搜尋引擎產品、不是向量 DB 產品、也不是單片 EDR。

## 適不適合進 workflow？

**適合。** 「AI 分析源片 → 擷取片段 → 畫面寫成 tags → 統一文字庫查出可行 clip → 再依明示 selection policy 選定」。這比「只有人物名單」或「時長最近自動當選」更接近可治理的素材庫。

四層不要混：

| 層 | 生命週期 | 回答的問題 |
| --- | --- | --- |
| **Source bible** | 一部源作品共用 | 集數、人物、場景、Canonical 事實的**穩定 id** |
| **Clip catalog** | 同作品共用、可增量 | 抽出的片段：入出點、時長、畫面描述、tags |
| **Unified text index** | 與 catalog 同一 SoT 的可查文字面 | 用同一套詞去找 clip，不必每次重看長片 |
| **Matching script** | 一條成片一份 | 約束、可行集、明示 policy、選中 clip |
| **EDR** | 一條成片一份 | 實際採用的入出點、字幕、發布 |

人物 id 解決「講對人」；clip catalog 解決「剪對畫面、夠不夠長」。

## 分析進庫（ingest）

分析源片時（人工抽或 AI 抽，策略可換）每條素材至少寫：

| 欄位 | 用途 |
| --- | --- |
| `clip_id` | 穩定主鍵 |
| `episode_id` | 連 bible |
| `source_in` / `source_out` / `duration_s` | 可重跑的擷取窗 |
| `visual_description` | **畫面裡有什麼**（誰、動作、物件、空間、光線／情緒短語） |
| `tags[]` | 統一詞表上的標籤；能連 `character_id`／`place_id` 的用 id，其餘用受控 tag |
| `dialogue_excerpt` | 可選；不是把整集 ASR 貼進 catalog |
| `duration_band` | 如 `hook`（≤3s）／`beat`（3–12s）／`hold`（12s+）或專案自定 bands |
| `ingest.origin` | `human`／`model`；模型抽的必須能被駁回 |

禁止：沒有入出點的「印象標籤」；沒有 `clip_id` 的散文；把整部片當一條 clip。

## 統一文字庫（給 AI 查）

Canonical 查找面是 **同一份結構化 catalog 的文字欄**（id、aliases、visual_description、tags、duration）。Agent／未來工具用同一套詞查，查到的結果必須是 `clip_id` 列表，不能是「再去片裡找找」。

未來若做向量檢索，只能當 adapter：**仍必須回傳 catalog 裡已存在的 `clip_id`**，不得發明庫外片段。Workflow 不規定 embedding 模型。

查詢最小契約（每個需要畫面的 shot）：

```text
need
constraints: must_tags / entity_refs / duration_target / duration_tol / duration_band / continuity
→ feasible candidates[]
     clip_id, matched_constraints, duration_delta, rejection_reasons
→ selection:
     policy: duration_closest | preserve_character_continuity | human_review | …
     rationale: …
→ selected_clip_id
```

Constraint 決定誰進可行集。Selection policy 才決定選誰。時長接近度可以是 criterion，**不得默認等於唯一的「最好」**。空可行集 → blocked，禁止從集外硬挑。

未來模型可當更強 Selection Actor；governance 契約不變。

## Matching script 如何用庫

頭：`bible_id`、`catalog_id`、`narrative_template_id`、目標成片時長、發布語  
身：每個 `shot_id` 含 constraints、feasible `candidates`、`selection.policy`／rationale、`selected_clip_id`  
未解析名稱、零可行候選、或選中 clip **違反約束**（不是「不是最近」）→ QC 失敗

通過後才寫入 EDR 的實際入出點；EDR 可微調入出點但必須仍指向同一 `clip_id`（或記 mutation：換 clip）。

## 與 bible 實體的關係

- tags 優先連 `character_id`／`episode_id`，避免「小明」「男主」兩套詞。
- 同一人物可有很多不同時長的 clip；腳本要短 hook 就選 `duration_band: hook`，要反應戲就選更長 band。
- `clip_hint` 不再當主查找；真正可剪的是 catalog 列。

## drop

- 把向量 DB／推薦引擎當 source of truth
- 無時長的「好鏡頭」分數；把 `duration_delta` 最小當成唯一自動選材規則
- 每條短片複製一份互不相通的 clip 庫（應掛在源作品下增量）
