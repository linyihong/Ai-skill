# 2d — 外部 monorepo outbound sync Phase 3（4 slices）（2026-07-08，Cursor Task transport）

> **專案證據邊界**：inner commit、class 名、live 環境細節留於 `<PROJECT_ROOT>` active main plan §執行紀錄；Ai-skill 只保留 generalized dogfood metrics（依 [`enforcement/sanitization.md`](../../../../enforcement/sanitization.md)）。

## Run 摘要

- **任務**：出站同步平台 Phase 3 — rollback、`sync_failure` 隊列 + reconcile worker、Admin 監控 UI、`WORKER_BATCH` scheduler（4 slices）。
- **Transport**：Cursor orchestrator + **Task subagent** executor / verifier（每 slice 一輪；fix 再 spawn）；**省略 `model`**。
- **Repo**：`<PROJECT_ROOT>` 外層 + `<INNER_REPO>/manageCode` 內層；雙 repo commit。
- **Consumer overlay**：`<PROJECT_ROOT>/.ai-skill/project/rules/plan-delegation-execution-loop.md` — 在 canonical L1–L3 之上擴 **L1–L4 外層驗收 + V1–V4 內層（含 V4 產出物）+ `slice_kind` + C1–C5 合規關閉**。
- **機械 gate**：`check-plan-delegation-orchestrator.py` + Cursor hooks（五事件）+ BDD `gate.plan_delegation_orchestrator`、`gate.consolidation.gherkin_canonical_placement`。

## Slice 軌跡（摘要）

| Slice | slice_kind | 重點 | 關閉類型 |
|---|---|---|---|
| 3.0-A | implementation | `failure_kind`、補償 delete、內層 IT | `implementation_done` |
| 3.0-A′ | outer_acceptance | Gherkin rollback + L2/L3 外層 | `slice_compliant_closed` |
| 3.0-B | combined | `sync_failure` 隊列 + reconcile worker | `slice_compliant_closed` |
| 3.0-C | combined | sync_status、pending_delete、monitor UI、batch scheduler | `slice_compliant_closed` |

## 相對 canonical / 2c 的新信號

| # | 觀察 | 對本 plan 的意義 |
|---|---|---|
| 1 | **implementation + outer_acceptance 拆 slice**（A / A′）比單一 combined 好關閉 | consumer overlay `slice_kind` 有效；canonical 可 advisory：user-visible 行為勿僅用 inner JUnit 關閉 |
| 2 | **verification_backfill + deliverables[]** Execute 前填好，Verifier V4 抓 feature / manifest 遺漏 | **Q7 正向**：backfill 像 execution 前 acceptance→evidence 映射 primitive |
| 3 | **機械 gate 生效**：gate 後 orchestrator **零** manageCode 實作 diff | **Q5 支持** Layer 3 consumer gate 夠用、schema promotion 仍不急 |
| 4 | **hook 副作用**：Execute 意圖下 orchestrator 連外層 plan 有時被擋 | 契約缺口 → kit §Cursor 傳輸備註 allowlist（見 [`01-dogfood-prompt-kit.md`](../01-dogfood-prompt-kit.md)） |
| 5 | **combined mega-slice**（3.0-C，8 deliverables）Verifier V4 負載高 | 建議拆 2–3 Verifier 輪；與 2c「多 slice 值得 loop」一致 |
| 6 | **deploy / migration** 由 orchestrator 做，不在 loop 內 | loop 邊界：runtime deploy 應列 brief acceptance 或標 `beyond-loop` |
| 7 | **Gherkin 唯一目錄 gate** 防 `.feature` 再分散 | consumer enforcement 範例；workflow dogfood 清單候選 |

## 量測欄（4-slice 彙總）

| 指標 | 值 |
|---|---|
| Slices | **4**（2 combined + 1 implementation + 1 outer_acceptance） |
| Task spawn | **~8–12**（每 slice executor + verifier；含 fix 輪） |
| acceptance-violation（merge 前） | **少數** — 多為 deliverable / outer tier 遺漏（V4 抓到） |
| verifier 降級（只跑 L1？） | **初期有**；補強 overlay L1–L4 + V4 後改善 |
| orchestrator 越界寫 manageCode | **0**（gate 生效後） |
| orchestrator 被迫回讀 diff 仲裁 | **罕見**（爭議時定點 Read） |
| 外層 L1–L3 linked 才關 user-visible slice | **是**（C1b 紀律有效） |

## vs 2c tiered archive

| 維度 | 2c | 2d |
|---|---|---|
| 領域 | 資料歸檔平台 | 同步 / 失敗處理 / Admin 監控 |
| 外層證據 | 以 inner IT 為主 | **L1 Gherkin + L2 BDD + L3 外層 IT** 為關閉條件 |
| slice 模型 | 8 小 slice | 4 slice + **slice_kind 混用** |
| 新 primitive 信號 | 多 slice 值得 loop | **backfill + slice_kind + deploy 邊界** |

## 契約回饋（寫回 canonical / consumer overlay）

1. **Consumer overlay 模式成立** — canonical 保持 L1–L3 + 三角色；雙 repo + Gherkin 外層用 project overlay 擴 tier，不必一次進 schema。
2. **Execute 意圖 hook allowlist** — 見 kit §Cursor 傳輸備註。
3. **`slice_kind`**（implementation / outer_acceptance / combined）應進模板 A「測試範圍」段（advisory）。
4. **Deploy 不屬 Production leg** — brief 需明示或 defer。
5. **Q5 仍維持 doc-only** — 2d + 2c + consumer 機械 gate = 證據累積中，尚不足以 promote schema。
