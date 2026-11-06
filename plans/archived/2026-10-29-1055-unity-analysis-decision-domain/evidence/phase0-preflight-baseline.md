# Phase 0 preflight + pre-migration baseline

- **Run id**: `2026-10-29-phase0-preflight`
- **Recorded at**: 2026-10-29T02:15:17Z (approx.)
- **Git HEAD**: `c0879ebb`（當時工作樹另有本 plan 未提交變更）
- **Decision**: **proceed**

## 1. RE evidence-contract（已讀）

Source: [`analysis/reverse-engineering/evidence-contract.md`](../../../../analysis/reverse-engineering/evidence-contract.md)

| Observation | 本輪 Phase 0 |
| --- | --- |
| Observation／Inference／Residual unknown 三分開 | 確認存在；Unity decision-workflow **只引用**，不另造哲學 |
| Residual unknown 最小欄位 | question／blocking／needed_authority／supporting_ids／status |
| 與 APK evidence-chain 接縫 | workflow gate 管層次；本契約管主張可否覆核 |

**Unity 補充形狀（實作時寫入 decision-workflow，不複製 RE 定義）**：方法命中、getter、物件引用、畫面、資產引用；裁決狀態 `supported`／`blocked` 等另欄。

## 2. Candidate files（存在性）

| Path | Status |
| --- | --- |
| `analysis/apk/workflows/unity-asset-parse-flow.md` | present |
| `analysis/apk/workflows/unity-fx-driver-trace.md` | present（canonical today） |
| `analysis/apk/workflows/unity-soft-visual-restore.md` | present |
| `analysis/apk/workflows/il2cpp-runtime-contract-validation.md` | present |
| `knowledge/runtime/routing-registry.yaml` → `route.analysis.unity-fx` | present |
| `knowledge/summaries/unity-fx.md` | present |
| `validation/scenarios/runtime/unity-fx-discovery-wiring-v1.yaml` | present |
| `scripts/ai-skill-cli/internal/app/unity_fx_discovery.go` | present；`unityFXSource = analysis/apk/workflows/unity-fx-driver-trace.md` |

## 3. Q2 — il2cpp 檔適用範圍

讀 `il2cpp-runtime-contract-validation.md` 開頭與章節：主體為授權 IL2CPP 參數／hook 階段／fixture 形狀驗證；交叉連到 Unity FX；**未**重定義 Observation／Inference。

**裁決**：遷移至 `analysis/unity/workflows/`；通用證據哲學改**引用** RE evidence-contract；apk 留 stub。

## 4. Discovery 訊號所有權（現況）

`route.analysis.unity-fx` user_signals：`Animator.Play`、`m_Payloads`、`PlayableDirector`、`Unity FX`、`Unity 特效`、`Unity 動畫`；context：`**/unity-fx-driver-trace.md`。

`route.analysis.apk.workflows`：Frida hook／API 文件化／local proxy／HLS 等；**不**擁有 `*.apk` 泛用 glob（屬 parent workflow）。

**Phase 1 約束**：新 `route.analysis.unity` 只加 Unity／IL2CPP／soft／GetContent／UnityPy 等領域訊號；禁止綁定泛用 Frida／`*.apk`。

## 5. Consumers — 硬編碼 `analysis/apk/workflows/{unity-*,il2cpp-runtime-contract-validation}.md`

Phase 1 **必須更新**（結構性；否則雙 SoT 或斷線）：

| 類 | Paths |
| --- | --- |
| Registry／generated reports | `knowledge/runtime/routing-registry.yaml`, `model-checklists.md`, `model-context-report.md`, `runtime-report.md` |
| Discovery／CLI | `scripts/ai-skill-cli/internal/app/unity_fx_discovery.go` (+ tests), `scripts/ai-skill-cli/README.md`, `scripts/README.md` |
| Validation | `validation/scenarios/runtime/unity-fx-discovery-wiring-v1.yaml` |
| Workflow | `workflow/apk-analysis/execution-flow.md`, `execution-flow.yaml`, `README.md` |
| Analysis hubs | `analysis/apk/README.md`, `analysis/apk/workflows/README.md`, `analysis/README.md`（新增 unity 入口） |
| Indexes／graphs／summaries | `knowledge/indexes/README.md`, `knowledge/graphs/apk-analysis-pilot.yaml`, `knowledge/summaries/unity-fx.md`, `knowledge/summaries/apk-workflows.md` |
| Intelligence | `intelligence/.../runtime-evidence-boundaries.md`, `representation-contract-parity.md` |
| Onboarding／RE collection | `runtime/onboarding/apk-analysis-completion.md`, `analysis/reverse-engineering/collection/collection-pitfalls.md` |
| Extraction maps | `feedback/extraction/recent-review-map.md`, CSV inventory rows（可後續批次） |

Phase 1 **明確不批次改**（plan 既有決定）：`feedback/history/**` 大量舊路徑 — 靠 apk stub 保持可點；可選後續清理。

## 6. 遷移前可比較基線（命令結果）

> Phase 1 結構性驗收必須與此節比較；僅「顯示通過」而無本基線則**不可**宣稱回歸正常。

### 6.1 Unity FX discovery unit tests

```text
cd scripts/ai-skill-cli && go test ./internal/app -run TestUnityFX -count=1
→ ok  github.com/linyihong/Ai-skill/scripts/ai-skill-cli/internal/app  ~0.9s
EXIT: 0
```

### 6.2 runtime validate + assert-source（pre-migration canonical）

```text
ai-skill-darwin-arm64 runtime validate --repo <AI_SKILL_REPO> \
  --assert-source analysis/apk/workflows/unity-fx-driver-trace.md --plain
→ exit 0
  runtime_index_native: atoms=497 sources=481 edges=359 fts=497
  knowledge_runtime_native: registry_records=66 summaries=33 graphs=35
  runtime_audit_warning: orphan_total=366 (routes=43, surfaces=82, scenarios=241)
```

同輪亦對 `--keyword m_Payloads`／`PlayableDirector` 跑過 validate（exit 0；輸出為全量 validate 摘要，與上列同級指數）。

### 6.3 基線指標快照（供 Phase 1 對照）

| Metric | Pre-migration value |
| --- | --- |
| TestUnityFX | PASS |
| assert-source `analysis/apk/workflows/unity-fx-driver-trace.md` | PASS |
| FTS atoms／sources／edges | 497／481／359 |
| registry_records／summaries | 66／33 |
| orphan_total（warning） | 366 |

Phase 1 後期望：

- TestUnityFX PASS（source 常數改為 `analysis/unity/workflows/unity-fx-driver-trace.md`）
- assert-source **新路徑** PASS；舊路徑可僅 stub 存在
- FTS／registry 數量無意外大幅退化；orphan 可 ± 小幅波動但須說明
- Unity 任務可查回新 canonical；一般 APK 路由不被誤導（負向：不因本變更把 Frida 泛用意圖鎖到 unity）

## 7. Conflicts / non-conflicts

| Item | Finding |
| --- | --- |
| 雙 SoT 風險 | 可控：Phase 1 stub-only |
| Layer 錯置 | 無：決策＋procedure 屬 analysis；編排屬 workflow |
| RE 契約衝突 | 無：引用即可 |
| Routing 搶訊號 | 可控：沿用 de-dup 紀律 |
| Feedback history 斷鏈 | 接受 stub；不阻塞 proceed |
| 需 ask user | **無** |

## 8. Decision

**proceed** → 進入 Phase 1（結構性遷移＋驗收）。不等 Phase 2；不做 Phase 3。
