# AI Codegen Validation Rate Parity（外部素材沉澱：AI 寫程式爆量但 Perf Test 跟不上）

Lesson date: 2026-05-27
Trigger: 使用者提供外部 2025 業界 infographic「AI 寫程式爆量的時代，Performance Test 已經跟不上了」
Status: candidate

## 觀察到的訊號

外部 infographic 摘要：
- AI codegen 把每位開發者每月程式碼產出推到 ~3.2x（4.5k → 14k 行）
- 43% AI 生成程式碼通過 QA/staging 但 production 仍需手動 debug
- 88% 公司需要 2–3 次 redeploy 才能確認 AI 修復有效
- 38% 開發者每週約 2 天在 debug 與驗證

四個 perf anti-pattern：迴圈藏 DB query、collection 無界、外部呼叫無 timeout、SQL 字串拼接。

## Lesson

當外部素材描述「某類加速工具導致驗證跟不上」時，應該識別其為**meta-tool-risk** 類型的觀察，跨工具可重用。處理路徑不是單一層級的更新，而是 6 層同步：

| 層 | 產物 |
|---|---|
| `intelligence/engineering/<domain>/` | 抽象原則 atom（為什麼這樣判斷） |
| `analysis/<domain>/` | 量化資料 + 解剖（如何觀察） |
| `enforcement/failure-patterns/` | Trigger + detection rule（如何防止） |
| `validation/scenarios/<workflow>/` | 機械化 scenario（如何驗證） |
| `workflow/<workflow>/` | 執行流程（何時執行 + 步驟） |
| `feedback/history/<domain>/` | 本 lesson（觀察記錄） |

`governance/` 為可選第 7 層，當 production-gate 需要 cross-tool 規範時加。

## 容易踩到的反模式（meta）

1. **只放 intelligence**：抽象原則沒對應觀察方法 → 沒人找得到、用不出來
2. **只放 enforcement**：detection rule 沒對應原理 → 規則更新時失去思考脈絡
3. **沒有 analysis layer**：量化資料散在 intelligence 或 enforcement 文件中 → 引用時混淆「是觀察還是判斷」
4. **沒有 workflow integration**：知識存在但 agent / reviewer 在實際流程中不知道何時觸發
5. **沒有 validation scenario**：流程只能靠人腦執行，無機械化路徑

## 對應到 Ai-skill 自身的關聯

本 repo 的 cognitive contract / hooks / runtime validation stack 就是「同步加速驗證」的具體實作。每次新加 codegen 自動化能力（例如本 session 的 Go-native hooks 取代 .sh），都應檢查 validators 是否需要同步擴張。

## 後續行動

- 候選狀態 `candidate-intelligence`：待 repo 內 first-party 觀察出現（例如：本 repo 自身或夥伴專案中遇到 AI 生成的程式碼通過 CI 但 production 出 perf bug），promote 為 `validated`
- 若高頻使用 perf-risk-gate workflow，後續加 `validatePerfRisks` pre-commit validator（Phase 7 候選工作）

## Related

- [`intelligence/engineering/ai-augmented-delivery/generation-validation-rate-parity.md`](../../../../intelligence/engineering/ai-augmented-delivery/generation-validation-rate-parity.md)
- [`analysis/ai-augmented-delivery/`](../../../../analysis/ai-augmented-delivery/README.md)
- [`enforcement/failure-patterns/ai-codegen-passes-ci-fails-production.md`](../../../../enforcement/failure-patterns/ai-codegen-passes-ci-fails-production.md)
- [`workflow/software-delivery/perf-risk-gate.md`](../../../../workflow/software-delivery/perf-risk-gate.md)
- [`validation/scenarios/software-delivery/ai-codegen-perf-risk-checklist.yaml`](../../../../validation/scenarios/software-delivery/ai-codegen-perf-risk-checklist.yaml)

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
