# Evidence Chain Validation

**Status**: `candidate-intelligence`

## 定義

Evidence chain validation 是一種 validation reasoning pattern：驗證必須沿著狀態傳播鏈一路追到最終可觀察結果，並在每個重要 state transition 上附上足夠證據。

它回答的問題不是「有沒有測試」，而是「證據是否覆蓋了 claim 所依賴的整條狀態鏈」。

## 基本形狀

```text
state source
→ propagation step
→ persisted or external state
→ readback path
→ user/business observable result
```

範例鏈路：

```text
Login → Session Created → Cookie Issued → Protected Resource → UI State
API → Domain Logic → SMTP Provider → Inbox → User Observable Email
Order → Gateway → Webhook → DB → Entitlement → UI
Producer → Topic → Consumer → DB → Business Effect
Flag Rule → Audience Match → Frontend Fetch → Rendered Feature
```

## 操作步驟

1. **列 claim**：要證明的產品行為是什麼。
2. **列 chain**：從狀態來源到最終可觀察結果，逐段列出 propagation path。
3. **標 evidence**：每段現有證據的類型、scope 和 confidence，依 [`evidence-model.md`](evidence-model.md) 判斷。
4. **找缺口**：任何沒有 readback、persisted state、external confirmation 或 user-observable state 的段落都不能被視為完整。
5. **選 depth**：依 [`evidence-depth.md`](evidence-depth.md) 決定是否需要 `live_system` 或 `independent_observation`。

## 判斷規則

- API success 只覆蓋 API segment，不自動覆蓋 DB、queue、external system 或 UI。
- Adapter success 只覆蓋 adapter segment，不自動覆蓋外部世界。
- UI success 若沒有 persisted/readback evidence，可能只是 presentation state。
- DB success 若沒有 user-observable readback，可能仍不支持完整 product claim。
- 高風險鏈路需要至少一個 readback path，critical 鏈路需要 independent observation。

## 與特化案例的關係

Identity-coupled side effect validation 是本 pattern 的特化案例，而不是獨立的一級規則：

```text
Evidence Chain Validation
├── Identity-Coupled Flow
├── Payment Flow
├── Event-Driven Flow
├── Email Flow
└── Feature Flag Flow
```

## 相關知識

- [`state-visibility-gap.md`](state-visibility-gap.md)
- [`evidence-model.md`](evidence-model.md)
- [`evidence-depth.md`](evidence-depth.md)
- [`evidence-collapse-point.md`](evidence-collapse-point.md)


## 組合、因果與發佈範圍的候選補充

選證據入口時，先問哪一層、哪個時間點最後決定宣稱的結果。局部成功或測試變紅，只有在當次輸入確實到達該層、且其他條件受控時，才能支持因果判斷。

| 宣稱範圍 | 判斷 | 可推翻的檢查 |
| --- | --- | --- |
| 多個元件組合後的 host 行為 | 採實際組合入口，包含順序造成的副作用 | 舊組合必須重現錯誤；只測單一元件不足 |
| 某守衛被突變後測試會失敗 | 先讀回最後有效層，隔離每次衍生狀態 | 突變被後續覆寫或結果來自舊範本時，測試不能證明守衛 |
| 發佈時整個分支符合不變式 | 在發佈必經入口驗當次狀態 | 帶入違規後只改無關檔案，仍應被攔截 |

適用於跨層覆寫、共享範本或多階段驗證。自足純函式、單次提交契約，或已有不可繞過全量驗證的入口，不必機械增加 host／本機全掃。啟用新發佈阻擋前，先盤點既存違規並修復或明確隔離，避免把未分類歷史直接當作本次變更新增錯誤。

此補充維持 candidate；來源事件與可攜推演不等於獨立重用實測。具體測試操作仍由專案 test strategy 決定。

### 來源回饋

- [2026-10-15_190000-mutate-the-layer-that-decides-and-reset-the-template.md](../../../../feedback/history/development-guidance/common/2026-10-15_190000-mutate-the-layer-that-decides-and-reset-the-template.md)
- [2026-10-06_090000-pipeline-behaviour-needs-host-level-test.md](../../../../feedback/history/development-guidance/common/2026-10-06_090000-pipeline-behaviour-needs-host-level-test.md)
- [2026-09-04_135746-gate-that-runs-only-at-commit-is-one-bypass-from-gone.md](../../../../feedback/history/development-guidance/common/2026-09-04_135746-gate-that-runs-only-at-commit-is-one-bypass-from-gone.md)
