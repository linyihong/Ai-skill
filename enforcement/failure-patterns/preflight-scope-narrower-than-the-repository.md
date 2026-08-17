# Preflight Scope Narrower Than The Repository（preflight 掃描範圍窄於實際 repo）

Status: candidate
Class: `traceability-gap`

## Trigger

Agent 在動工前做盤點（preflight）：找出現有實作、確認架構相容、界定變更範圍。掃描時只覆蓋「產品程式碼」的目錄 —— 例如只看 `src/`。

而該 repo 另有承載真實邏輯的位置：`tests/`（含開發用工具、seeding、fixture 產生器）、`scripts/`、`tools/`、`docs/` 的可執行片段。

## Failure Mode

Preflight 回報「此機制尚不存在，需新建」，而它其實已經存在於未被掃描的目錄裡 —— 通常還帶著一份寫得很清楚的理由，說明為何放在那裡而不放在 agent 正打算放的位置。

於是：

- 重複建置，產生同一事實的第二份宣告（正是多數計畫明確要避免的）
- 錯過既有實作已解決的子問題（例如「不重抄 key」已由讀 manifest 解決）
- 與既有設計理由正面衝突而不自知；該理由可能比 agent 的規劃更好

觀察到的實例：計畫排定由各模組 pack 宣告 role→permission，preflight 只掃 `src/` 後回報「尚未落地」。實際上 `tests/` 下的開發用 seeder 已有完整矩陣，讀 manifest 取 key、fail-closed，且其註解明確反對 pack 放置（pack 由出貨的 Migrator 套用，會替每個客戶決定產品層級的選擇）。

## Risk

- 產生第二份真相，且兩份會漂移（見 [`cross-boundary-agreement-never-mechanically-checked.md`](cross-boundary-agreement-never-mechanically-checked.md)）
- 推翻一個有更好理由的既有設計，而該理由從未進入討論
- 浪費的不只是重建的工，還包括之後兩份並存所需的協調

## Required Agent Action

**Preflight 的搜尋範圍以 repo 為界，不以「產品程式碼目錄」為界。**

1. **對每個「這個機制不存在」的結論，掃描範圍必須覆蓋 `tests/`、`scripts/`、`tools/` 等非產品目錄。** 開發用工具經常是某個機制的第一個實作。
2. **以概念名搜尋，不只以型別名搜尋。** 既有實作可能叫別的名字；搜「suggested」「recommended」「default」「matrix」這類詞比搜預定的類別名有效。
3. **找到既有實作時，先讀它的放置理由再決定移動它。** 註解裡的「為什麼在這裡而不在那裡」是設計資訊，不是雜訊。
4. **若既有理由與計畫衝突，停下來提出，而不是照計畫續建。** 兩者都可能對；決定權不在執行者手上。
5. **搬移既有實作時，區分「可達性」與「偏好」。** 若正式環境需要它（而它目前的位置在部署中不可達），移動有硬理由；若只是覺得該放別處，就不是。

## Prevention Gate

寫下「需新建」之前：

- 我搜過 `tests/`、`scripts/`、`tools/` 了嗎？
- 我搜的是概念詞還是我預想的名字？
- 如果它已經存在，最可能藏在哪個我沒看的地方？
- 找到了嗎？找到的話，它為什麼在那裡？

## 驗證

1. Preflight 紀錄中列出實際掃描的路徑範圍，且包含非產品目錄
2. 「不存在」的結論附帶搜尋詞與範圍，可被他人複驗
3. 若發現既有實作與計畫衝突，該衝突寫入證據文件並提交決策，而非由執行者逕行取捨

## Linked Rules

- [`cross-boundary-agreement-never-mechanically-checked.md`](cross-boundary-agreement-never-mechanically-checked.md) — 重複建置的後果
- [`deployment-scoped-mechanism-verified-in-wrong-venue.md`](deployment-scoped-mechanism-verified-in-wrong-venue.md) — 可達性 vs 偏好的另一面
- [`../failure-learning-system.md`](../failure-learning-system.md) — `traceability-gap` 分類
