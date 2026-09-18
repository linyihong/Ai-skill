# Candidate: Episode evidence vs knowledge accumulation

Companion to [`_plan.md`](_plan.md)。**候選閉環，不是完整 Knowledge DB，也不是 Phase 2 schema。**
缺口不是 `identity_resolution` 欄位不夠，而是缺少
Episode → Learning → Verification → Knowledge Promotion。
觀察：[`evidence/2026-09-18-episode-vs-knowledge-accumulation.md`](evidence/2026-09-18-episode-vs-knowledge-accumulation.md)。
Identity 原則：[`08-identity-precedes-naming.md`](08-identity-precedes-naming.md)。
升格閘：[`17-story-promotion-gate.md`](17-story-promotion-gate.md)。

Agent 不得「本集分析出 speaker = 暱稱 → 直接寫進 identity_resolution」。
本集觀察與跨集知識累積必須分開。

```text
episode analysis
  → episode evidence          # 這一集發生了什麼
  → accumulation gate
       ├── keep in episode
       └── learning candidate # 尚未證實，可提出、不可直接寫庫
            → independent verification
                 ├── reject
                 ├── keep candidate
                 └── promote
                      ├── knowledge store     # identity / narrative / retrieval
                      └── mechanical registry # 可重複、可機械驗證的 pattern
```

寫碼前先分類：

| 這是什麼 | 寫入哪裡 |
| --- | --- |
| 本集觀察 | Episode Evidence |
| 尚未證實的新發現 | Learning Inbox candidate |
| 已驗證的跨集知識 | Knowledge Store（需 provenance／status／validity） |
| 可重複且可機械驗證的模式 | Mechanical Registry |

機械與語意走不同路徑：watermark 區域反覆出現 → pattern candidate → validation → mechanical rule；稱呼／關係 → semantic candidate，不得一次 resolved。

最小 dogfood：先只做 identity domain 的 Episode Evidence → Learning Inbox → Verification → Promotion。不建完整 Knowledge DB，不讓 agent 直接改 registry／knowledge 檔。Phase 3 **不**改 workflow。
