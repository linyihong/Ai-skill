# Observation — break candidate system

**Run ID**：2026-09-28-break-candidate-system  
**Kind**：contract（製片反例；已寫進 caption／layout 契約）  
**Extends**：[`2026-09-25-caption-composition-semantic-breaks.md`](2026-09-25-caption-composition-semantic-breaks.md)

## 反例

- `dangling_start` 把「懂不能當行首」寫死，分不出 `看懂` 這個 lexical unit。
- 字元合法的「談｜話」通過 heuristic，語意上拆開「談話」。
- 為這一集再加一條字表特例，不會變成可累積的 break evidence。

## Contract

Script heuristic 產生候選與 penalty。LLM 只在 feasible set 內選擇，並留下
`break_evidence`。Learned pattern 先當 observation；review 後才 promotion 進 policy。
本 repo 沒有 `CjkBreakPolicy` 程式；契約落在 caption layout，不改 narrative lifecycle。
