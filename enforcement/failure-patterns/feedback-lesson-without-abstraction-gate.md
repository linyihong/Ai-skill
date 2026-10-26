# Feedback Lesson Without Abstraction Gate

Status: candidate
Class: `scope-drift`

## Trigger

當 agent 準備把「本輪專案裡學到的東西」寫進 `<AI_SKILL_REPO>/feedback/history/**`，或使用者指出「回饋太專案特例／沒抽象化／別人用不了」時，使用此 pattern。

## Failure Mode

Agent 跳過 Abstraction Gate，直接把柜机名、產品名、sample id、裝置／host、一次 live 結果或專案路徑寫進共用 lesson；或只做去敏（刪路徑／token）卻**沒有**把原因與驗證改寫成換專案仍成立的規則。結果是：Ai-skill 看起來有回饋，但其他專案的 agent 無法套用，或被專案細節誤導。

常見變體：

1. **未抽象化** — Generalized Lesson 仍綁定單一產品／柜机／樣本。
2. **假去敏** — 刪了絕對路徑，但仍用可識別專案專名當規則主詞。
3. **錯層** — 把只適用於一個 target 的 runbook 當成 technique lesson 推進 `feedback/history/`。
4. **形式 PASS** — 寫了 Swap-project PASS，但正文仍依賴專案腳本、章節號、固定資料形狀或無條件修法。
5. **證據混淆** — 同專案修正成功或方法已接入口，卻冒充獨立重用驗證。
6. **漏檢查** — 沒填 `#### Abstraction Check`，或自評 FAIL 仍試圖 commit。

## Risk

- 其他 agent 複製錯誤的專案假設，浪費時間或做出錯誤自動化。
- 共用庫被專案 incident 汙染，違反 [`reusable-guidance-boundary.md`](../reusable-guidance-boundary.md)。
- 使用者以為「已回饋 Ai-skill」，實際上沒有可重用技巧。

## Required Agent Action

1. 停止繼續往 `feedback/history/` 寫專案細節。
2. 讀 [`../../feedback/feedback-lessons.md`](../../feedback/feedback-lessons.md) §Abstraction Gate、[`reusable-guidance-boundary.md`](../reusable-guidance-boundary.md)、[`sanitization.md`](../sanitization.md)。
3. 跑四步 Abstraction Gate；`Swap-project test: FAIL` → 只更新專案文件，不開／不改共用 lesson。
4. `PASS` 時重寫 One-line Summary／Generalized Lesson／Applies 邊界為中性詞，Evidence 只留摘要＋`<PROJECT_ROOT>/…`。
5. 填完整 `#### Abstraction Check` 三行；提交時讓 KGE `feedback-lesson-closure` 通過。
6. 依 feedback-lessons 的「自評 PASS 的實質檢查」，重建輸入、機制、決策與反例；不把一次修法當跨專案必要步驟。
7. 若問題是「agent 反覆忘記抽象化」而不只是單次內容錯誤，更新本 pattern 或相關 checklist，勿只改那一條 lesson。

## Prevention Gate

| Check | Required answer |
| --- | --- |
| Gate run | 寫檔前是否已完成 Abstraction Gate 四步表？ |
| Swap-project | 拿掉專名後的摘要是否仍成立？（必須能寫成 `PASS — …`） |
| Evidence split | 專案證據是否只在 `<PROJECT_ROOT>`，lesson 內無 raw incident？ |
| Mechanical | `#### Abstraction Check` 是否齊全且非 FAIL？ |
| Layer | 這是 technique lesson、failure pattern，還是應留專案 runbook？ |

## Validation

- 新 lesson 含 `Abstraction Check` 且 Swap-project 為 PASS。
- `git commit` 在 Ai-skill 上通過 `feedback-lesson-closure`（無不當 `[skip-feedback-lesson-closure]`）。
- 全文搜尋無本機絕對路徑、無未宣告的產品／柜机專名作為規則主詞。
- 另一個不熟悉原專案的 agent 只讀 lesson 能說出「何時套用／何時不套用」。

## Related

- [`../../feedback/feedback-lessons.md`](../../feedback/feedback-lessons.md) §Abstraction Gate
- [`reusable-guidance-boundary.md`](../reusable-guidance-boundary.md)
- [`sanitization.md`](../sanitization.md) / [`sanitization-mechanical.md`](../sanitization-mechanical.md)
- [`skill-local-feedback-bypass.md`](skill-local-feedback-bypass.md)（漏做 failure-learning 的 complementary pattern）
- KGE：`scripts/ai-skill-cli/portable/kge/rule_feedback_lesson_closure.go`（Abstraction Check）

## 審查邊界

格式 gate 只驗證結構，不能證明語意抽象化。對已存在的條目也須讀正文、核對狀態與方法承接位置；結果應分開報告「格式通過」「可攜性推演」「獨立實測」，不把其中一項代替其他項。
