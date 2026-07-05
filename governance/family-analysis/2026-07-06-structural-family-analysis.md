# Structural Family Analysis — perf-governance / cache-read-path / play-view-dedup

**Date**: 2026-07-06 ・ **Stage**: 1（family determination only）
**Sources**（一手 plan，Vidoe-Test）：`2026-06-17-1400-performance-validation-architecture-pilot.md`、
`2026-06-22-1600-h5-redis-read-cache/_plan.md`、`2026-06-18-1430-play-view-dedup-event-db.md`
**Method**：Step 1 domain removal（perf/cache/redis/latency 詞全去除）→ Step 2 structural matrix →
Step 3 determination。**本檔不含 candidate / criterion / promotion。**

## Structural Matrix

| Structural Element | A (perf-gov) | B (cache-read-path) | C (play-view-dedup) |
|---|---|---|---|
| Observation | ✓ 每次變更都問「要不要證據」（recurring intake） | ✓ 一次性現況快照 + 可見面盤點 | ✓ incident（上層 KPI 與底層驗證不一致）+ 容量疑問 |
| Rule | ✓ regression 門檻 + 「可決策 vs 僅觀測」語意 | **?** — policy 與 implementation config 分不乾淨 | ✓ 「未來 phase 只能由 evidence 觸發」 |
| Registry | ✓ **跨 incident 持久比較基準**（持續更新、被未來比對） | ✓ 但**語意不同** — implementation inventory（namespace / script registry / 檢查清單） | **✗** — 無自有 registry，**借用 A 的** |
| Projection | ✗ | ✗ | ✗ |
| Executor | ✓ 常駐 runner，重複執行 | ✓ 一次性 pilot runner + 機制本身 | ✓ 一次性 burst + integration/BDD |
| Validation | ✓ schema 化 summary → review → 才決策 | ✓ generalization report + live-eval 強制 | ✓ capacity report + DOM gate |
| Decision Authority | **recurring gate authority**（每次變更過同一關） | **one-shot stakeholder resolution**（R1–R6 拍一次） | **evidence-triggered deferral**（E 由 F 的證據決定） |
| Observation-only Layer | ✓ 明文（advisory 語意） | ✓（pilot 標 advisory） | ✓（部分指標只觀測不 gate） |

## Structural Elements（去 domain 後的角色）

- **A** = 常駐 governance machine：intake gate（recurring）→ 持久比較基準 → 常駐 executor →
  schema 化 validation → recurring decision authority，附明文 observation-only 層。
- **B** = 一次性 optimization decision delivery：快照 → 方案裁決（one-shot resolutions）→ 實作 +
  advisory pilot → 證據確認。registry 為 implementation inventory，非比較基準。
- **C** = 一次性 correctness delivery：incident → 領域規則 → 一次性驗證 → evidence-triggered
  deferral（把未來決策交給證據）。無自有 registry。

## Shared Traits（3/3 recur — analogy，非 same-family 證明）

1. observation-first
2. evidence-before-decision
3. decisionable 與 advisory 分離

（此三條與 ECS plan EL-3 的 recurrence 同源；記 recurrence，不在本層抽 invariant。）

## Differences（決定性）

1. **Registry 語意不同**：A=跨 incident 比較基準 / B=implementation inventory / C=無（借用 A）。
2. **Decision authority 型態不同**：recurring gate / one-shot resolution / evidence-triggered deferral。
3. **Lifecycle 不同**：standing machine vs 兩個 one-shot delivery。
4. **Consumer 關係**：B、C 把 intake 填進 A 的 store、用 A 的 runner 慣例、寫進 A 的 evidence 面 —
   **B、C 是 A 的 consumer，不是 sibling**。consumer 關係本身否定 sibling 假設：三者不在同一比較軸上。

## Family Determination

```text
Result:      Different family
Reason:      Registry semantics differ; decision-authority type differs;
             lifecycle differs; B/C are consumers of A (not siblings).
Shared:      observation-first / evidence-before-decision / advisory split
Conclusion:  Analogous, but not same family.
```

**Stage-2 備註（不在本檔處理）**：A 對照 `governance/lifecycle/governance-pattern-template.md` 的
applicability **不合**（template 限 mechanical / commit-time；A 的 executor 為 agent-invoked、
budget gate deferred → A 目前 pre-mechanical）。B 的 Rule cell = `?`（insufficient）。
此兩點留給 Step 4 criterion mapping 時使用。
