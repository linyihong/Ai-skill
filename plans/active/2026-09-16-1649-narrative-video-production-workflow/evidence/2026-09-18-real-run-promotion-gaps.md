# Real run — promotion-gate gaps

**Run ID**：2026-09-18-real-run-promotion-gaps  
**Kind**：去敏真實 source-analysis dogfood（不是完整 EDR run）  
**Conclusion**：observable／evidence unit 可用；narrative／identity promotion
尚不可接受。**本 run 有效，但整體 Phase 3 未 PASS。**

原始媒體、專案名、路徑、逐字稿、人物稱呼與可識別 evidence 留在外部專案。
本檔只保留可重用的計數、失效類型與驗證規則。

## Run summary

- 約兩分鐘直式影片。
- 數十個 shots／clips、數十條 transcript、近百條 OCR。
- 產出二十餘個 evidence units、少量 event candidates、兩個被升格的 story
  evidence，以及多筆 low／pending。
- Observable technical extraction、OCR、Face、Evidence Unit、Dialogue Unit、
  traceability：可繼續 dogfood。
- Narrative resolution：partial。
- Identity resolution、Story State、Story Evidence promotion：fail。

## Findings

### F1 — unresolved upstream promoted to accepted

上游 narrative role 仍 `final: null`，下游卻
`story_evidence.status: accepted`。這是 promotion lifecycle violation。
任何 blocking upstream unresolved 必須機械 stop。

### F2 — empty state claim promoted

`before`／`after` 為空、`subject`／`object` 為 null 的
`possible_state_delta` 被當 story state。它只能是 candidate；不能支持 accepted
story evidence。

### F3 — traceability mistaken for validity

`traceable: true` 只證明 refs 存在。它不證明 transcript 正確、event semantics
成立、relevance 是 high、或 state 改變。Accepted gate 必須是
resolved upstream + resolved event + valid state claim（若有）+ traceability +
independent verification。

### F4 — mechanical semantic overreach

Mechanical resolver 從一句 dialogue 直接宣告 high narrative relevance。
Mechanical 可提出 candidate／escalation，不得在未登記 deterministic policy 下
作 semantic final。

### F5 — dialogue cluster is not yet an event

Event candidates 幾乎只由 dialogue 聚合；尚未證明 action／visual／location／
participant change。新增輕量 `basis.*_refs` 可讓 verifier 看見實際支撐，
但本輪不改 workflow schema。

### F6 — vocative promoted to canonical identity

同一 voice entity 收到多個互斥稱呼，仍被標 `canonical_name: resolved`。
Vocative 是 addressee／relationship evidence，不是 uttering speaker 的名字；
缺 cast resolution 時必須 candidate／unresolved。

### F7 — OCR mention promoted to entity

Watermark variants、普通片語、情緒詞、關係稱謂被列成 name mentions，部分又成
resolved OCR entities。`name_mention` 必須先過 visual-text role、lexical／NER
與 identity resolution；mention ≠ entity。

### F8 — raw text ambiguity leaks downstream

多條 ASR／OCR 顯示疑似誤辨或斷句問題，卻被用來產生 beat、high relevance
與 story summary。Raw text 必須保留；校正只新增 candidate，不覆寫來源。

### F9 — lifecycle fields contradict

部分 records 同時出現 `resolver: llm`、`final: null`、`clarity: clear`、
`escalate_to: null`，卻已有 classification／relevance。需要統一 proposal、
resolution status、final decision 的語義。

### F10 — low archive needs reactivation

Low relevance 不刪是正確的；但關係稱謂可能被後集重新解釋。Archive 必須有
reason 且允許 identity／relationship／later-reference 觸發 reactivation。

## Classification

- `contract_gap`：**尚不能宣稱 Phase 2 contract gap**；這些下游層仍是
  observation-only，尚未進 Phase 2 workflow。先記 `promotion_gate_gap`。
- `data_insufficient`：identity／name resolution 與 ambiguous text。
- `adapter_only`：ASR normalization、OCR alignment、diarization 品質。
- `potential_design_error`：dialogue → event → state → accepted 邊界過寬；
  第二個真實 episode 後裁決。

## Next validation

下一集只驗四項，不加 detector：

1. unresolved upstream 是否確實擋 accepted；
2. state change 是否能填出具體 subject／change／before-after；
3. vocative／OCR mention 是否維持 candidate；
4. event basis 與 low archive 是否在跨集 evidence 下正確解析／reactivate。

完整 Phase 3 仍須走 brief → bible／catalog → matching → EDR → locale → QC →
publish／outcome；本 run 只覆蓋 source-analysis 子鏈。
