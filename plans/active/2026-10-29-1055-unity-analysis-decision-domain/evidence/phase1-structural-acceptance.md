# Phase 1 structural acceptance

- **Run id**: `2026-10-29-phase1-structural-reverify`
- **Compared to**: [`phase0-preflight-baseline.md`](phase0-preflight-baseline.md)
- **Scope**: structural only — **no** Phase 2 5×2 dogfood; **no** Phase 3 mechanization

## Phase 1 結構性驗收（6/6）

| # | 驗收 | Result | Evidence |
| --- | --- | --- | --- |
| 1 | **SoT 唯一** | **PASS** | apk stubs only: `unity-asset-parse-flow` 245B、`unity-fx-driver-trace` 242B、`unity-soft-visual-restore` 254B、`il2cpp-runtime-contract-validation` 281B；全文在 `analysis/unity/workflows/`（fx 13165B、asset 7867B、soft 7181B、il2cpp 5634B） |
| 2 | **舊入口仍有效** | **PASS** | stubs 指向 `../../unity/workflows/…` + `decision-workflow`；Go／registry／summary／scenario 已改新路徑；**feedback/history** 舊 `analysis/apk/workflows/unity-*`／il2cpp 路徑已批次改寫為 `analysis/unity/workflows/…`（`feedback/` 內 0 筆舊路徑殘留） |
| 3 | **連結有效** | **PASS** | `analysis/unity/**/*.md` 相對連結掃描：8 files、**0 broken** |
| 4 | **Routing 正向** | **PASS** | `go test … TestUnityDecisionRoutePositiveAndAPKFridaNegative/positive_*` PASS；`route.analysis.unity` → `analysis/unity/decision-workflow.md`；`route.analysis.unity-fx` → `analysis/unity/workflows/unity-fx-driver-trace.md`；訊號例：`Unity soft-enter GetContent`、`IL2CPP contract`、`Animator.Play`／`m_Payloads` |
| 5 | **Routing 負向** | **PASS** | 同 test `negative_generic_apk_frida_not_unity` PASS；下列**不得**命中 unity／unity-fx：`對這個 APK 用 Frida 抓包`、`用 Frida hook 這個 app 的 HTTP API`、`local proxy hook Netty pipeline` |
| 6 | **工具與 Discovery／RAG** | **PASS** | `runtime compile` PASS（既有 advisory）；`runtime refresh` PASS（atoms=499 sources=483 edges=359 fts=499；registry=67 summaries=33）；`runtime validate` exit 0；`TestUnityFX*` PASS；FTS `m_Payloads` → `analysis/unity/workflows/unity-fx-driver-trace.md`；`decision-workflow` 在 sources／FTS；vs Phase 0（497／481／66）為 +2／+2／+1（新 domain＋route），orphan_total 366 不變 |

## Commands（本輪重跑）

```text
cd scripts/ai-skill-cli && go test ./internal/app \
  -run 'TestUnityDecisionRoutePositiveAndAPKFridaNegative|TestUnityFX' -count=1 -v
→ PASS (UnityFX discovery/retrieval + positive/negative handoff)

ai-skill-darwin-arm64 runtime compile|refresh|validate --repo <AI_SKILL_REPO> --plain
→ exit 0

sqlite3 knowledge/runtime/sqlite/runtime-index.sqlite \
  "SELECT source_path FROM fts WHERE fts MATCH 'm_Payloads' LIMIT 5;"
→ analysis/unity/workflows/unity-fx-driver-trace.md
```

## Metrics vs Phase 0 baseline

| Metric | Phase 0 | Phase 1 reverify |
| --- | --- | --- |
| TestUnityFX | PASS | PASS |
| Routing handoff +/− | （僅 registry 敘述） | **PASS** dedicated Go test |
| atoms／sources／edges | 497／481／359 | 499／483／359 |
| registry_records／summaries | 66／33 | 67／33 |
| orphan_total | 366 | 366 |
| assert／RAG canonical | apk path | **unity** path in FTS／sources |

## Delivered tree

```text
analysis/unity/
├── README.md
├── decision-workflow.md
├── case-ledger-template.md
└── workflows/{README,unity-asset-parse-flow,unity-fx-driver-trace,unity-soft-visual-restore,il2cpp-runtime-contract-validation}.md
```

## Follow-ups（explicitly not Phase 1）

- Phase 2：專案側 5×2 soft-enter dogfood（決策是否改善調查）
- Phase 3：機械化 — **不做**，除非 Phase 2 後仍反覆漏證據／混層／錯判完成
