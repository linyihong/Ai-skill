Status: candidate

# Horizontal Classification Dimension Thinking

## 觀察

在將 `java-tsv-trim-split-trailing-empty.md` 分類時，最初直接放入 `analytical-reasoning/failure/`，因為該目錄已有 failure 子層。但這個分類忽略了知識的「語言特定性」——Java `String.trim()` 行為是 Java 標準庫的語言特定知識，不屬於跨語言的分析技術失敗模式。

## 教訓

分類知識時，不應只檢查現有子層是否能容納，而應先思考：

1. **橫向思考**：這份知識是否屬於全新的分類維度？
   - 語言特定知識 → `language-specific/<lang>/`
   - 框架特定知識 → `framework-specific/<framework>/`
   - 平台特定知識 → `platform-specific/<platform>/`
2. **只有當不屬於新維度時**，才檢查現有子層是否能容納。

## 已套用的改善

- `knowledge-update-flow.md` Step 2.4 已加入橫向維度決策樹
- `intelligence/engineering/` 下新增 `language-specific/` 維度
- `intelligence/engineering/language-specific/java/failure/` 已建立

## 適用範圍

任何需要將知識文件分類到 `intelligence/engineering/` 下的場景。

## 觸發信號

- 知識內容涉及特定語言的標準庫行為
- 知識內容涉及特定框架的 API 行為
- 知識內容涉及特定平台的 runtime 特性
- 直覺上「放在現有目錄好像不太對」但說不出原因

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
