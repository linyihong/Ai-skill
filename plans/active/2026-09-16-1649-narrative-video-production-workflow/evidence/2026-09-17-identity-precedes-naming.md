# Observation — identity precedes naming

**Run ID**：2026-09-17-identity-precedes-naming  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow）  
**原則**：Identity precedes naming; naming is an evidence-backed resolution, not an input requirement.

角色在真實世界是連續的；名稱常延遲出現。禁止第一集就要求 `character_id` + 名字，也禁止 LLM 硬猜。

## 層（發現 ≠ 已解析）

| 層 | 此時合法 | 尚未要求 |
| --- | --- | --- |
| Observed entity | `entity_id` + face／speaker／descriptor | 名字 |
| Identity cluster | 跨集「可能同一人」的 observations | `character_id` |
| Character identity | `character_id` | canonical name 仍可空 |
| Named character | `canonical_name` + `name_evidence` | — |
| Series cast | 已解析表：`call_name`／aliases | 不得當成第一集發現層 |

`name.status: unresolved` 完全合法。

## Evidence 不改歷史

第一集 observation 保持「當時不知道名字」。第二集名稱證據產生 **link**，不把 ep01 原始列改成「林晴」。

`identity_links` 獨立：`from` entity → `to` character，`relation: same_identity`，附 face／voice／narrative／explicit_name 等 evidence 與 `accepted`。不是「AI 覺得是她」。

Face／voice **不是**唯一 identity key（雙角色、變裝、回憶、配音、遮臉、音質差）。它們只當 evidence，避免早期錯誤污染 series_cast。掛點與禁止捷徑見 [`2026-09-17-face-as-candidate-evidence.md`](2026-09-17-face-as-candidate-evidence.md)。

## 與 series_cast 的關係

[`05-series-cast-canonicalization.md`](../05-series-cast-canonicalization.md) 的 `call_name`／aliases 掛在**已解析** `character_id` 上，並應能回指 `identity_evidence`（entity 列）與 `canonical_name.evidence_refs`。兩層不要混。

## 真實片子要數

第一集 unnamed entity 比例；後來命名是否用 link 而非覆寫；face／voice 衝突時有沒有當 identity key。畫面文字當 name evidence 見 [`2026-09-17-visual-text-evidence.md`](2026-09-17-visual-text-evidence.md)。沒有普遍需求就不進 workflow。
