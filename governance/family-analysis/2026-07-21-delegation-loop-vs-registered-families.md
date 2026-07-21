# Structural Family Analysis — delegation-loop (ERA) vs registered ECS families

**Date**: 2026-07-21 ・ **Stage**: 1（family determination only）
**動機**：delegation-loop 是最高頻在執行的 active plan（每日 commit、~30 個 dogfood run）。問題：它的
證據 shape 是否 map 到**已註冊的 ECS plan criterion**（若是 → 高頻 ECS feeder；若否 → 乾等的原因就
被證實是結構性的）。
**Source**：`plans/active/2026-07-08-0825-delegation-verification-arbitration-loop/_plan.md` + `evidence/`。
**Method**：Step 1 domain removal（delegation / executor / verifier / orchestrator / brief / dogfood /
APK / SD 詞全去）→ structural matrix → determination。**不含 candidate / criterion / promotion。**

## Structural Matrix（vs 唯一有 resolved pointer + template 的 registered family）

| Structural Element | delegation-loop（ERA 去 domain 後） | governance-pattern 6-step（promoted template） |
|---|---|---|
| Observation | ✓ 「預計 vs 實現的落差」 | ✓ |
| Rule | ~ **角色協議（doc-only）**，非 machine-checkable rule | ✓ policy / structural rule |
| Registry | ✗ brief / findings 是 **per-task artifact**，非跨案例持久 registry | ✓ |
| Projection | ✗ | conditional |
| **Executor** | **✗ 決定性缺口** — 人工角色協議，**無 code executor、無 commit/build-time 偵測** | ✓ **必要**（mechanical，code-detectable） |
| Validation | ~ verifier 三層（**人工** L1 重跑 / L2 讀碼 / L3 對抗），非 test-suite-as-executor | ✓ Go test / scenario yaml |
| Decision Authority | arbitration（fix / defer / reject），**與 evidence 生產顯式分離** | validator pass/fail（機械） |
| Evidence-producer ≠ decider | ✓ **核心且顯式命名為 ERA** | ✓ 隱含（validator 產、human 動作） |

## Structural Elements（去 domain 後）

- **delegation-loop = 人工 evidence-responsibility 角色協議**：producer(實作)/ independent-producer(獨立檢查)
  / decider(arbitrator) 三責任分離，per-task loop，human-run，**刻意非 mechanical**（plan 自身 Watch-Out：
  不建 orchestrator automation、不動 schema、不接 runtime）。
- **governance-pattern = mechanical governance subsystem**：violation 能被 code 在 build/commit-time 偵測。

## Differences（決定性）

1. **Executor 軸**：6-step template 的 load-bearing 要求是 **mechanical executor**（code-detectable）。
   delegation-loop **沒有**——它是人工角色協議。**這正是 perf 當初卡的同一條 applicability 軸**
   （perf 是 pre-mechanical；delegation-loop 是 deliberately non-mechanical）。
2. Registry：per-task artifact vs 跨案例持久 registry。
3. Lifecycle：human role protocol（每任務重跑）vs standing mechanical subsystem。

## Family Determination

```text
Result:      Different family
Reason:      No mechanical executor (deliberately doc-only role protocol) →
             fails the 6-step template's load-bearing Executor requirement;
             registry/lifecycle also differ. Not economics / interaction-hazard /
             perf either (those are their own subjects).
Conclusion:  delegation-loop is its OWN family (Evidence Responsibility
             Architecture), NOT an instance of any registered ECS plan.
```

## 對「ECS feeder」問題的直接結論（本分析的動機）

**delegation-loop 不會自然 feed ECS。** 它雖然是最高頻的 active plan，但 shape 不 map 到任何 registered
ECS criterion → 期望「把最高頻 plan 接成 ECS feeder」**證實不成立**。這把使用者的「不知道等到什麼時候」
從「感覺」變成**結構事實**：registered ECS plan 全是低頻來源（governance subsystem ~1–2/月、economics
事件式、interaction 事件式、perf 卡 family-undetermined），而唯一高頻的 plan 是**不同 family**。

→ ECS `phase2_gate=20` 在自然速率下**現實上是 distant / maybe-never**。這是 gate calibration 的觀察點
（門檻當初是猜的，沒有 rate 數據），**現在不改**；先讓 EL-5 起的 velocity 數據累積。此結論**不推動、也不
阻擋** Phase 1（Phase 1 已是價值）。

## Shared Trait（recurrence，analogy 非 same-family）

delegation-loop 的核心「evidence-producer ≠ decision-authority」與 ECS（candidate≠accepted / index≠
consumable）、economics（observation≠decision）、perf（result≠decisionable，EL-3）、family-analysis
（analysis≠criterion）是**同一條 recurring 分離原則的第 5 個獨立域**——而且 delegation-loop 是唯一**顯式
命名**它的（ERA）。依紀律：structures 各異 → **analogy / recurrence，非 same family、不抽 invariant**。
記 recurrence，`next_check` 繼續觀察是否某天結構也開始收斂。
