# Family Analysis（structural working analysis — Stage 1 layer）

**Observation 與 candidate 之間的中間分析層。** 存放「這些案例是不是同一個 structural family」的
工作分析（working analysis），累積可比對的 Structural Matrix。

## 這層是什麼、不是什麼

| 是 | 不是 |
|---|---|
| Stage 1 分析產物（matrix + family determination） | **Evidence Log**（EL 只收真實 observation；matrix 是對 observation 的分析，型態不同，不得混入） |
| 可累積：第 4、5 個案例直接加進既有 matrix | **governance-pattern-library-draft**（那是 Stage 2 之後、進 criterion 的 library evidence） |
| 只回答「同不同 family」 | candidate / criterion mapping / promotion（一律不出現在本層） |

## 標準流程（每個 candidate family 都走；Step 4 才第一次碰 criterion）

```text
Step 1  Domain removal        去掉 domain/implementation 詞（Method A — veto test）
Step 2  Structural matrix     只畫 shape：角色 / 狀態 / transition / authority /
                              decision / observation / registry / validation
Step 3  Family determination  三值輸出：Same family / Different family / Insufficient evidence
------------------------------------------------------------------ 本層到此為止
Step 4  Criterion mapping     （library draft / 對應 plan 的事）
Step 5  Candidate             （ECS 的事；criteria_hits 由人標）
```

兩個互補方法：

- **Method A — Domain removal（veto test）**：「拿掉 domain metrics，治理形狀還成立嗎？」
  falsification heuristic，只能剔除、不能 admit（canonical 定義暫存於 ECS plan EL-3 方法備註，
  provisional home 未升格）。
- **Method B — Structural matrix → family determination**：Method A 過了之後，用固定元素表逐格比對，
  判 same/different/insufficient。「shape survives」≠「same family」——兩方法各管一段。

## 規則

1. 每份分析一個 dated 檔：`YYYY-MM-DD-<slug>.md`。
2. 內容只有五節：Structural Matrix / Structural Elements / Shared Traits / Differences / Family Determination。
3. 結論只有三值。**不寫 candidate / criterion / promotion**。
4. Shared traits 若跨案例 recur，記 recurrence（可回饋 EL 作 observation），不在本層抽 invariant。

## 層級定位（2026-07-06 maintainer）

每層回答**不同的問題**；本層的價值 = Observation 與 Candidate 之間終於有一個「**可分析但不可承諾**」的空間——很多東西可以先停在 Analysis，不必急著分類。

| Layer | 回答的問題 | 允許輸出 | 不允許輸出 |
|---|---|---|---|
| Observation（EL） | 發生了什麼？ | observation | analysis、candidate |
| **Family Analysis（本層）** | 形狀是不是同源？ | analysis | criterion、candidate |
| Candidate（ECS inbox） | 有沒有既有語言可以描述？ | candidate | promotion |
| Plan Evidence | 這個 plan 接不接受？ | evidence | analysis |
| ADR / Promotion | 值不值得抽成共用能力？ | primitive | observation |

## 觀察點（deferred — 現在不抽象）

若未來 **非 governance 域**（UI architecture / runtime portability / tool economics / software delivery）
也自然長出 `Observation → Structural Matrix → Family Determination`，屆時再問：**Family Analysis 是
project-specific working layer，還是 AI-native Cognitive System 的通用 Analysis primitive？**
目前證據集中在 governance 系列 → 維持 working layer，不升格。

## Index

| 日期 | 分析 | 結論 |
|---|---|---|
| 2026-07-06 | [perf-governance vs cache-read-path vs play-view-dedup](2026-07-06-structural-family-analysis.md) | **Different family**（B、C 為 A 的 consumer） |
