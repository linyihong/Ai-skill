# Translation Decision — Execution Flow

Canonical lifecycle。欄位 SoT 在 [`contracts/`](contracts/)。**不要**寫 provider／prompt／固定譯詞表。

## Lifecycle

```text
0. Bind TranslationContext     → locale + content.type + speaker/addressee/scene
1. Locale Resolution           → bind only; ≠ Language Detection
2. Content-Type Resolution
3. Reference / Identity Resolution  ← dialogue kinship／pronouns（I23）
4. Title Structure Analysis    → when content.type=title
5. Expression + Semantic Analysis
   ├── lexical / pragmatic / syntactic / discourse / cultural
   └── semantic_structure + roles + modality
6. Target-Locale Realization seeds → name／script／kinship Candidate Space
7. Registry binding + Guards → failure-registry-binding → core ∪ locale patterns
8. Apply Constraints / Guards  → Feasible Candidates（含 CORE-F21 truncated）
9. Selection + Target Realization
   ├── semantic / syntax（I14–I15）
   ├── register / character_voice / politeness / gendered（I24）
   └── naturalness（I16／I18；≠ grammatical）
10. Independent Review
11. Validation（分欄）
12. Finality — PASS / REVIEW / BLOCK（I21–I22）
13. (post) Failure Learning → Governance Review（I13）
```

## Finality ternary

| Status | When |
| --- | --- |
| **PASS** `accepted` | gates pass + `review.required=false` + source complete |
| **REVIEW** `needs_review` | explicit `review.reason[]` only |
| **BLOCK** `blocked` | truncated source／hard fail／missing required context |

禁止：不確定 → 一律 REVIEW。詳見 [`finality.yaml`](contracts/finality.yaml)、[`13`](../../../plans/active/2026-09-22-1000-translation-decision-workflow/13-ep8-contract-revision.md)。

## Failure registry binding（step 7–8）

```text
Locale Resolution
        ▼
failure-registry-binding.yaml     # binding / index only
     /                    \
    ▼                      ▼
Core failure-patterns    locale/<target>/failure-patterns
        \                    /
         ▼                  ▼
           Constraint Space
```

| 讀 | 路徑 | 角色 |
| --- | --- | --- |
| Binding | [`registry/failure-registry-binding.yaml`](registry/failure-registry-binding.yaml) | 載入哪些 registry（`pattern_ids`） |
| Core SoT | [`registry/failure-patterns.yaml`](registry/failure-patterns.yaml) | 跨語言 failure concepts |
| Locale SoT | [`registry/locale/<locale>/`](registry/locale/README.md) | manifestations（`manifests: F*`） |

新語言：加 locale 檔 + binding entry；**不**改主鏈。  
舊檔名 `locale-failure-taxonomy.yaml` → deprecated pointer。

## 禁止

- 固定 姐夫=義兄さん／女强人=バリバリ… 當唯一答案（I10／I23）
- 補全 truncated source（I21）
- `if locale: += 例句` prompt 膨脹
- Phase 4 route／provider 寫進本檔

## Dogfood

- Ep8（PASS／REVIEW／BLOCK retune）：[`ep8-ja-walkthrough.yaml`](examples/ep8-ja-walkthrough.yaml)
