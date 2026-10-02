> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-02 - Cue finalization must not hard-delete candidates

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自 narrative dogfood：OCR／ASR 已有大量 raw evidence，spoken／subtitle 重建後 `final.cue` 數歸零，且 resolution 只留下「機械／近音／LLM／未決」彙總，沒有 per-candidate reject reason。

#### One-line Summary

`candidate → final cue` 只能 classify（accepted／uncertain／rejected），禁止 PASS／DELETE；raw counts 與淘汰原因必須可追溯，且不得與 story relevance 混關。

#### Human Explanation

Evidence pipeline 在 OCR／ASR 已「有證據」之後，若 hard-cut 要求同時過齊 geometry、script、lexicon、LLM，任一失敗就丟棄，會得到 precision 看起來很高、recall 直接歸零。對 dogfood 而言「真的沒字幕」與「resolver 不敢接受」完全不同；後者必須留下 `uncertain`／`rejected`，否則下游 story／LLM 永遠看不到候選。跨語言硬字幕（Latin OCR＋CJK ASR）更不能用字面一致當硬閘。

#### Trigger

- OCR／ASR raw 明顯 > 0，但 `dialogue_cues`／final cues = 0
- notes 僅見「機械0 近音0 LLM0 未決N／共0」，沒有 rejection table
- Latin hardsub 或 wrong-script 路徑直接 `continue` 丟棄，不寫 status
- 把「劇情是否重要」提前拿來決定是否保留 dialogue cue

#### Evidence

- Tool: product dialogue evidence／spoken text resolution dogfood
- Sanitized：多集 OCR／ASR 兩位數以上 → final cues 0；重建彙總 mechanical／phonetic／LLM 皆 0
- Evidence path: `<PROJECT_ROOT>` analysis notes only（不寫本機絕對路徑）

#### Generalized Lesson

1. **三態 classify，禁止 DELETE**：每個 candidate 必須落到 `accepted`｜`uncertain`｜`rejected`；`rejected`／`uncertain` 仍保留在 resolution 旁路（或同 doc 的 `rejected[]`），不得從 evidence graph 蒸發。
2. **分層判斷**：（L1）mechanical worth-keeping；（L2）cross-modal「是否同一句」（允許 OCR≠ASR，尤其跨語言）；（L3）僅 conflict／uncertain 才 LLM。LLM 不是唯一放行闸。
3. **保全計數**：至少記錄 `candidates.ocr`／`candidates.asr`／`accepted`／`uncertain`／`rejected`＋`reason[]`；最終 `cues=[]` 時仍要能回答「淘汰在哪一層」。
4. **Cue ≠ story**：dialogue cue finalization 只回答「是不是對白／字幕 evidence」；story importance 屬下一層 extraction，不得反向刪 cue。
5. **診斷先於整包**：raw≫0 且 final=0 時，先做單集 rejection table；不要用全量 pack／重載模型把 OOM 與 resolver bug 混成同一 failure。
6. **對齊既有契約**：遵守 [`text-evidence-multimodal-resolution.md`](../../../../workflow/narrative-video-production/text-evidence-multimodal-resolution.md) 的 observed→candidate→resolved，禁止 silently overwrite／丢弃 observed。

#### Agent Action

看到 raw evidence≫0、final cues=0：停止整包重跑；定位 finalization／hard-cut 的 delete／`continue`；補 resolution contract＋rejection table；單集驗證後再動資源／OOM。不要先加新 AI 探針。

#### Goal / Action / Validation

- Goal: candidate→final 可審計，且跨語言硬字幕不會因字面不一致被全滅。
- Action: 改 finalization 為三態；delete 路徑改寫 rejected／uncertain；單集排出 stage 淘汰表。
- Validation or reference source: 單集跑後 doc 含非空 `rejected` 或 `uncertain`（當 raw>0）；notes／summary 含 stage counts；`cues=0` 時仍能指出主要 `reason`。

#### Applies When

- Multimodal OCR＋ASR（或硬字幕＋口播）已進入 spoken／subtitle resolution
- Final cues 用於對白／字幕 corpus，而非僅 debug dump

#### Does Not Apply When

- 尚未產出任何 OCR／ASR candidate（仍屬 acquisition／probe）
- 純單一模態且無 fuse／resolution 階段

#### Validation

- 以 raw OCR／ASR 計數 > 0 的單集重跑：不得只剩 `cues=[]` 而無 `rejected`／`uncertain`／stage counts。
- 跨語言案例（Latin OCR＋非 Latin ASR）不得因字面不等而整批消失。
- 整包 OOM 與 resolver 歸零分開記 failure。

#### Promotion Target

- Workflow：`workflow/narrative-video-production/text-evidence-multimodal-resolution.md`
- 銜接：`text-evidence-subtitle-candidate.md`（candidate status 語意）

#### Promotion Record

尚未 promotion；保留為 candidate history。本輪先強化 multimodal-resolution 契約條目並連到本 lesson。

#### Required Linked Updates

- 更新 `feedback/history/narrative-video-production/common/README.md` 索引列
- 強化 `workflow/narrative-video-production/text-evidence-multimodal-resolution.md`：禁止 hard-delete；要求 uncertain／rejected 保全
- Project evidence 留在 `<PROJECT_ROOT>` analysis／job notes；本 lesson 不含專案名／路徑
