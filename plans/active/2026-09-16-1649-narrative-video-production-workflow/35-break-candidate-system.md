# Candidate: Break candidate system

Companion to [`33-caption-composition-semantic-breaks.md`](33-caption-composition-semantic-breaks.md)。**caption／layout 能力契約，不改 narrative lifecycle。**  
觀察：[`evidence/2026-09-28-break-candidate-system.md`](evidence/2026-09-28-break-candidate-system.md)。

`CjkBreakPolicy` 一類 heuristic 保留為 **v0**：candidate generation、feature、penalty。
它們不是最終斷句，也不是 `NEVER BREAK`。Mechanical 產生可行集；LLM 在集內做語意選擇並寫
`break_evidence`；mechanical 再驗 glyph、lossless、protected span。接受／拒絕進入 evidence。
Learned pattern 停在 observation，經 review 才 promotion。LLM 不得直接改 canonical policy。
單集「談｜話」不當特例寫進字表。
