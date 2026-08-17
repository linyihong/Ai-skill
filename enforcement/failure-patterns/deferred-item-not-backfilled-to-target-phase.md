# Deferred Item Not Backfilled To Target Phase（延後項只寫在證據檔，未回填目標階段）

Status: candidate
Class: `traceability-gap`

## Trigger

多階段計畫執行中，某個 acceptance 項目被判定為**轉出**而非完成：依賴尚未存在、根因與另一階段相同、現在做等於重複既有的錯誤做法。

Agent 在該 slice 的證據文件裡完整寫下轉出理由，並在關閉判定中誠實標記 `deferred`。

## Failure Mode

轉出被記錄在**來源**（本 slice 的證據檔），卻沒有回填到**目的地**（目標階段的 checklist）。

這看起來已經很完整 —— 理由充分、揭露誠實、Verifier 也確認沒有謊報。問題在於這兩份文件的讀取時機完全不同：

| 文件 | 何時被讀 |
| --- | --- |
| 本 slice 證據檔 | 本 slice 執行與驗證期間；關閉後幾乎不再被打開 |
| 目標階段 checklist | 幾週後、往往是**另一個 session** 開始該階段時 |

執行目標階段的人照 checklist 做事。轉出項不在上面，於是不會被做，而且**沒有任何東西會提醒他** —— 該階段的 checklist 自身完整自洽，Exit 條件也會滿足。

長流程計畫在中途換 session 是常態（上下文壓縮、額度、跨日），這使得「只有前一個 session 知道」等於「沒有人知道」。

## Risk

- 轉出項永久遺失，且因為當初的揭露是誠實的，事後追查時不會被歸類為隱瞞 —— 更難發現
- 計畫可以在所有階段都「完成」的情況下，整體仍缺一塊
- 缺口在最後歸檔階段才浮現，屆時原始理由已不在任何人的上下文裡

## Required Agent Action

**轉出是一次雙向寫入：來源記理由，目的地記工作。**

1. **標記 `deferred` 的同時，就在目標階段的 checklist 加一列。** 不留到關閉時再說 —— 關閉時最容易漏。
2. **目的地那列要能獨立被執行。** 寫明做什麼、依賴什麼、為何當時不做，不能只寫「見 Phase N 證據檔」。
3. **若轉出理由是「依賴另一個尚未建立的東西」，該依賴本身也必須是目標階段的一個 checklist 項**，且排在轉出項之前。
4. **關閉 slice 前用轉出項編號反查目標階段。** 搜尋該編號在計畫檔中是否出現於目的地；只出現在來源就是還沒做完。

## Prevention Gate

宣告 slice 關閉前：

- 本 slice 有哪些 `deferred`？每一個在目標階段的 checklist 上都找得到嗎？
- 一位只讀目標階段 checklist 的執行者，會做到這些項目嗎？
- 轉出項所依賴的前置工作，也在目標階段列出來了嗎？
- 若這個計畫從此換人接手，只靠計畫檔，缺口會被發現嗎？

## 驗證

1. 以轉出項編號搜尋計畫檔，同時命中來源證據檔與目標階段 checklist
2. 目標階段的該列可獨立閱讀執行，不需回頭讀來源證據檔才知道要做什麼
3. 關閉 commit 的 diff 同時包含本階段的關閉區塊與目標階段的 checklist 新增列
4. 最終歸檔時，所有曾標記 `deferred` 的項目皆可追蹤到完成或明確取消

## Linked Rules

- [`decision-revised-without-contract-authority-update.md`](decision-revised-without-contract-authority-update.md) — 同家族：單一事實需寫入多個落點，且落點的讀者不同
- [`shallow-component-traceability-validation.md`](shallow-component-traceability-validation.md) — 追溯性驗證深度
- [`../failure-learning-system.md`](../failure-learning-system.md) — `traceability-gap` 分類
