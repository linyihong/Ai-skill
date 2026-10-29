# Unity／IL2CPP 分析方法

`analysis/unity/` 是 Unity／IL2CPP **決策與取證程序**的 canonical 目錄。
端到端編排與完成 gate 仍在 [`workflow/apk-analysis/`](../../workflow/apk-analysis/README.md)。通用 Android／Frida／流量方法仍在 [`../apk/`](../apk/README.md)。

## 閱讀順序

1. [`decision-workflow.md`](decision-workflow.md) — **決策入口**（目標、假設、實驗、驗收）
2. 依證據選 [`workflows/`](workflows/README.md) 內單一 procedure
3. 需要帳本時複製 [`case-ledger-template.md`](case-ledger-template.md) 到**業務專案**（六必要欄；勿把 raw dump 寫進本庫）
4. 主張品質：[`../reverse-engineering/evidence-contract.md`](../reverse-engineering/evidence-contract.md)

## 目錄

```text
analysis/unity/
├── README.md
├── decision-workflow.md
├── case-ledger-template.md
└── workflows/
    ├── unity-asset-parse-flow.md
    ├── unity-fx-driver-trace.md
    ├── unity-soft-visual-restore.md
    └── il2cpp-runtime-contract-validation.md
```

## 範圍

**負責：** Unity／IL2CPP 證據收斂決策、資產解析、驅動追蹤、soft 視覺還原、執行契約驗證。

**不負責：** 泛用 Frida／traffic triage（`analysis/apk`）；agent 全流程編排（`workflow/apk-analysis`）；HOW TO THINK heuristics（`intelligence/...`）；專案 raw evidence。

## 與 apk 的關係

舊路徑 `analysis/apk/workflows/unity-*.md` 與 `il2cpp-runtime-contract-validation.md` 僅為 **stub**，指向本目錄。新工作請直接讀本域。

## 可選作業節奏（非 gate）

整理本週新增證據、反覆失敗假設、下一個最小實驗——**不做** commit hook／排程強制。

## Inbound

- `route.analysis.unity` → primary `decision-workflow.md`
- `route.analysis.unity-fx` → `workflows/unity-fx-driver-trace.md`
- APK workflow 在 Unity 視覺／契約缺口時鏈入本域
