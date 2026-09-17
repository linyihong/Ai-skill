# Observation — Series Cast Canonicalization

**Run ID**：2026-09-17-series-cast-canonicalization  
**Kind**：Phase 3 observation（**不是** `contract_gap`，**不改** Phase 2 workflow）  
**Status**：候選層；等真實片子計數再決定是否升 contract

## 與現況對照

Phase 2 已有 bible 的 `character_id`、受控 tags、matching `entity_refs`。  
**還沒有**：唯一 `call_name`、aliases 不得拿來寫稿、機械閘「稿件只許 call_name」、笼统詞先機械解析再開 LLM。

缺的是 **Series Cast Canonicalization**，不是整套骨架。

## 候選位置（升格後才寫進 workflow）

不進 EDR（那是單片決策）。掛在源作品：

```text
Source Bible（誰）
  → series_cast（怎麼叫：character_id / call_name / aliases）
  → Script（怎麼說；生成只許 call_name）
  → semantic_context（這句實際在說誰）
  → Matching（用哪段素材）
  → EDR（最後決策）
```

- `call_name`：稿件 canonical display name。LLM 寫稿只能出這個。
- `aliases`：理解、舊資料、源片辨識。**不是**寫稿詞彙。
- 笼统詞（妻子／總裁）：機械掃描 → 若 series_cast **唯一**可解析則程式 canonicalize；歧義 → `AMBIGUOUS` → LLM 候選或人工 → 再機械複驗。不讓 LLM 先判全部。

這與 Evidence → Constraints → Feasible Set → Selection 同構：確定的歸機械；不確定才升級。

## 與 semantic_context 的分工

| 層 | 服務誰 | 例 |
| --- | --- | --- |
| `dialogue.text` | 觀眾 | 「你真的要去？」 |
| `series_cast.call_name` | 稿件一致性 | 陳宇／林晴 |
| `semantic_context` | 機器決策 | speaker／addressee／subject_refs |

三者目的不同；不要把台詞寫成長查找句。

## 真實片子要數

1. 主詞／受詞省略的台詞比例  
2. generic role term 比例  
3. 同一角色多名稱（aliases 衝突）比例  
4. 是否真的影響 catalog retrieval  
5. 一份 series_cast 表解掉多少  
6. 純機械 resolve vs 必須 LLM repair 的比例  

幾乎每片都需要才升 contract；少數特例維持觀察。
