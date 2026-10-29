# Phase 1 structural acceptance

- **Run id**: `2026-10-29-phase1-structural`
- **Compared to**: [`phase0-preflight-baseline.md`](phase0-preflight-baseline.md)

## Checklist

| 驗收 | Result |
| --- | --- |
| 無雙 SoT | PASS — apk `unity-*`／il2cpp stubs ≈240–280B；全文在 `analysis/unity/workflows/` |
| 舊入口導向 | PASS — stubs 指向 canonical；apk README／workflows README 指引 `analysis/unity` |
| Routing 正向 | PASS — `route.analysis.unity` + `route.analysis.unity-fx` primary 為新路徑 |
| Routing 負向 | PASS — 新 route 未綁 `*.apk`／泛用 Frida；apk.workflows 已移除 Unity procedure candidates |
| Compile | PASS（既有 advisory warnings） |
| Refresh／validate | PASS after **rebuild** `ai-skill-darwin-arm64`（舊 binary 仍查 apk 路徑導致 RAG 假失敗） |
| TestUnityFX | PASS |
| assert-source 新路徑 | PASS |
| Discovery／RAG | PASS — `m_Payloads`／`PlayableDirector` 查回 `analysis/unity/workflows/unity-fx-driver-trace.md` |

## Metrics vs Phase 0 baseline

| Metric | Phase 0 | Phase 1 |
| --- | --- | --- |
| TestUnityFX | PASS | PASS |
| atoms／sources／edges | 497／481／359 | 499／483／359 |
| registry_records／summaries | 66／33 | 67／33 |
| orphan_total | 366 | 366 |
| assert-source | old path PASS | **new path** PASS |

Atoms／sources／registry +2／+2／+1：符合新增 `analysis/unity` 決策文件與 `route.analysis.unity`；非意外退化。

## Delivered tree

```text
analysis/unity/
├── README.md
├── decision-workflow.md
├── case-ledger-template.md
└── workflows/{README,unity-asset-parse,unity-fx-driver-trace,unity-soft-visual-restore,il2cpp-runtime-contract-validation}.md
```

## Explicitly deferred

- `feedback/history/**` 舊路徑批次改寫（靠 stub）
- Phase 2 5×2 dogfood
- Phase 3 機械化
