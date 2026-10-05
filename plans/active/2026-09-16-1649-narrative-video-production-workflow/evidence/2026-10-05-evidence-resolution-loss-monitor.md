# Observation — Evidence resolution loss must be observable and recoverable

**Run ID**: 2026-10-05-evidence-resolution-loss-monitor
**Kind**: Phase 3 dogfood observation → contract/adapter gap
**Scope**: text-evidence acquisition, multimodal resolution, and derived cue-cache integrity.

## Generalized observation

An evidence pipeline can acquire a substantial candidate corpus, then produce a
small publishable cue set without a mechanically explainable disposition for
the difference. Treating every non-accepted candidate as absent destroys the
ability to distinguish three different conditions:

1. evidence is invalid (`rejected` with a concrete reason);
2. evidence is valid but unresolved (`uncertain`);
3. evidence was consolidated (`merged` with a target).

The defect is a resolution-loss anomaly, not proof that either OCR or ASR is
globally insufficient. OCR visual text, ASR speech timing/phonetics, and
geometry/region identity are complementary evidence dimensions; none should be
promoted to unconditional truth.

## Contract refinement

Record a named funnel (`raw_ocr`, `usable_candidates`, `fused_candidates`,
`accepted`, `uncertain`, `rejected`, `merged`). A large transition into
`accepted` is suspicious when the remaining candidates have no traceable
disposition. The response is a scope-limited re-probe—region OCR, ASR window,
or frame sample—then re-resolution. It is not a blanket rerun and not a fixed
whole-video coverage threshold.

Coverage is compared against expected evidence: resolved cues over speech
windows and over `subtitle_like` windows. A visually quiet or speech-free
timeline interval is not, by itself, subtitle loss.

## Watermark and cache implications

Sticky text must be classified from spatial persistence, repetition, and visual
features before dialogue resolution. Role-aware projection preserves observed
evidence while preventing watermark/UI material from contaminating the
publishable dialogue text.

Derived cue and coverage artifacts are cache inputs only after schema validation
and atomic publication. Cache provenance must identify source content plus OCR,
ASR, region-profile, fusion-policy, and resolution versions. An invalid or
incompatible artifact triggers a targeted rebuild and cannot be accepted as a
healthy cache hit.

## Linked updates and validation

- Workflow: `text-evidence-multimodal-resolution.md` adds the resolution-loss
  monitor and cache-integrity contract.
- Record: `records/text-evidence.yaml` adds funnel, cache metadata, artifact
  integrity, and corresponding gates.
- Plan companion and this evidence index identify the Phase 3 validation work.

Adapter validation now covers atomic cue/coverage publication and cache reuse:
the fixture confirms that invalid cue integrity or invalid coverage JSON is
rejected rather than treated as a healthy cache hit. The adapter emits the
named funnel and marks an unexplained contraction as `targeted_required`.

The adapter now runs a bounded window OCR/re-resolution pass with before/after
counts, preserving failed or unresolved observations. Fixtures distinguish
explained duplicate reduction from missing dispositions, unrelated texts with
shared timestamps, quiet timeline gaps, and failed acquisition.

## Attribution and eligibility findings

Lexicon segmentation can change the number of parts while detection boxes
remain unchanged. A positional pairing fallback that copies the aggregate
line into each box gives watermark text a dialogue box's geometry. Box-local
raw text must retain ownership regardless of segmentation count. The region
contract and a mixed watermark/dialogue fixture now capture this requirement.

A second gap is status projection: a nested unresolved or rejected decision
cannot become publishable through an absent/stale outer status. Recovery must
pass the same finalization gate as initial resolution, including episode-level
role projection; an increased cue count alone is not recovery evidence.
An accounted `merged` observation can still point to an uncertain target.
Retry outcomes must inspect the target's final status after eligibility gates;
accounting closure alone must not be reported as recovered dialogue.

Another attribution hazard is conversion through a text/time-only summary or
replaying raw boxes over already-projected evidence. Both can discard region
identity and source-level persistence, making a short retry incorrectly
rehabilitate sticky or scene text. Preserve box-local metadata across retries
and require attributable dialogue support at finalization. Source-local changing
text bands can support mid-frame captions; short retries must not establish
that source profile by themselves. Ambiguous roles remain uncertain.
Bounded retry scheduling should prioritize attributable dialogue losses over
unresolved role noise; preserve the latter for review without starving the
actual omission windows. Source-profile matching must tolerate box jitter.

Validation remains partial: scoped acquisition executes and preserves unresolved
observations, but accepted-content quality and remaining loss must be checked
independently. No complete subtitle coverage or publish-ready claim is made
from fixture passes or increased cue counts.

## Cross-source regression validation boundary

回歸驗證需保留互相獨立的兩類 assertion：non-dialogue 不進 accepted，
以及已有可信角色支持的真字幕不因污染修復被排除。須涵蓋 box-local raw
完整與舊快取缺失兩種格式，避免只測修復後的理想輸入。主要 changing-text
band 的相對頻率不是否定第二字幕帶的充分證據；低頻帶、位置變更、jitter
皆需獨立測例。舊 evidence replay 與 fresh acquisition 是不同驗證層，
兩者皆不能以正常退出或 cue 數增加代替內容、時間與成片驗收。具體執行
結果與未通過的 fixture 留在專案證據；cross-source 品質驗收仍未關閉。

詞庫擴詞可能形成第三類 loss：把反覆出現的對白首字吸收到 watermark
term，即使 raw observation 保留完整字幕，derived dialogue 已缺字。
擴詞應由同一 watermark region 的 raw 支持，不能跨 box 依 aggregate
prefix 復現數定案。Regression 必須同時檢查擴詞邊界與真字幕首字保留。

## 分階段修復與跨語系驗證缺口

修復需凍結 raw observation、derived cache、policy version 與測試基準，
每次只改一項 attribution 行為並重跑原 assertion。既有污染詞庫也需從
raw box 重新驗證；僅阻止未來擴詞，不能修好已存在的跨 box 長詞。
無 box-local 支持的衍生長詞應隔離並保留 provenance，不可繼續侵蝕對白。

box／parts 數相同仍不證明文字歸屬。舊格式缺少可靠 ownership 時，整組
保留 unresolved，而非複製 aggregate 到多個 box。理想格式的 positive
fixture 需明示真正的 raw ownership，另保留舊格式 negative fixture。
獨立支持的第二字幕帶不可只因頻率低被否定；反面測例須阻擋未決文字
借用字幕角色，以免修漏字卻引入水印。

單語回歸通過不等於跨語系能力通過。Latin script／單詞長度不是 junk
的充分條件；須分開測英文短字幕、水印、品牌、UI 與未知文字。雙語
script-run 只支持語言候選，沒有 box 時空間關係應 unresolved，不能
聲稱已確認上下疊行。文字／時間摘要不可取代原始 region identity。

目前此驗證層仍 partial：機械重播、fresh acquisition、完整 resolver、
獨立 frame annotation 與成片檢查是不同證據。快取盤點需明示缺樣本，
命中 junk heuristic 的 Latin box 不可未看畫面就稱為被漏掉的英文字幕。
新增失敗測例揭露跨語系及空間歸屬缺口；不以既有測例通過關閉品質閘。

## Retention 修復的驗收邊界

禁止 script／長度／全大寫本身決定硬刪；取消此規則不表示所有 Latin
文字都是 dialogue。Positive 要覆蓋短字幕及黏字仍可進 boundary recovery，
negative 要覆蓋品牌、logo、水印、UI、場景文字不被直接升成 accepted。
未決 evidence 的保留是一個正確 disposition，不是內容品質失敗，也不是
成片驗收通過。`merged` 必須有實際可解析 target 與 target_status。
文字／時間相同仍不足以聲稱屬於某個 accepted region；需帶來源 identity
並確認 target 的 attribution 支持，不能把 retention ledger 又變成誤歸屬。

原始 visual observation ledger 應覆蓋融合前的排除，不只追蹤已清理候選。
明示計量單位與來源，分開報 raw-ledger 的 disposition 完整性、resolver
候選的未決數與 accepted 品質，避免補齊 uncertain 後把 loss 警告洗掉。
已捕捉到下游 loss 的案例不能推論整個 corpus 的 acquisition recall 足夠。

雙語的語言候選與空間候選分開：缺 box、共用 aggregate box、缺有效
共時窗口時，spatial relation 保持 unresolved。有 box 與窗口也先產候選，
不能替代整條 pipeline 保留 region identity 的驗證。代表案例矩陣需列
feature、獨立 annotation、expected behavior 及缺樣本狀態；不以每部
影片跑過一次取代能力驗收。Temporal／ASS／成片閘仍待 identity 穩定。
