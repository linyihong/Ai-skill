> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Slot symbols may have in-reel and paytable presentations

Status: candidate

#### One-line Summary

同一協議符號 ID 可能有兩套美術：**轉軸互動**（Spine／Unity 盤面）與 **賠付表靜態**（木框／說明頁）；必須分開登錄，不可互相覆蓋。

#### Human Explanation

賠付表上的高價值符號常是完整木框肖像；盤面旋轉時則是 Spine 合成角色或無框 Sprite。兩者協議 ID 相同，但裁切來源與外觀不同。若只用賠付表裁切當同轉匹配圖，或只用 atlas 零件當唯一圖，都會丟資訊。

#### Trigger

- 賠付表符號有木框／說明排版，同轉截圖上的符號沒有相同邊框。
- High 符號在 Unity 是 Spine skeleton，賠付表卻是靜態完整圖。

#### Evidence

- Sanitized: one cabinet curated `assets/in-reel/` vs `assets/paytable/` plus `manifest.symbolMap.*.artPresentations`.
- HTML symbol table shows two thumb columns; sample grid uses in-reel for matching.

#### Generalized Lesson

1. `image` / 同轉網格預設用 `inReel`。
2. `paytableImage` / `artPresentations.paytable` 單獨存賠付表構圖。
3. Validator 兩邊圖檔都要存在；產生器表格兩列都要顯示。

#### Agent Action

發現外觀差異時立刻開兩目錄並寫 `artPresentations`，不要先決定「哪套才算真符號」。

#### Goal / Action / Validation

- Goal: 協議 ID 對應的互動圖與賠付表圖都可追溯。
- Action: dual folders + manifest fields + HTML dual thumbs.
- Validation: package validate PASS；HTML 含「转轴互动」「赔付表静态」。

#### Promotion Target

- Project slots onboarding / symbol-map generator（已落到 target repo）。

#### Related

- `2026-09-14_153000-slot-variable-lines-reel-major-jagged-grid.md`
- `2026-09-14_151200-unitypy-spine-atlas-y-origin-probe.md`

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
