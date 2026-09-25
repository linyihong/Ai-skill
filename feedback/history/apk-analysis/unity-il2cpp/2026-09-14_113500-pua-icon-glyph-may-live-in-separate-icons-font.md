> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - PUA icon glyph may live in a separate Icons font file

Status: candidate

#### One-line Summary

`UIText` 的 body font 名稱不一定是 Private Use Area icon 的來源檔；要對候選 `.otf` / `.ttf` 做 cmap 查 codepoint，找出真正含該 glyph 的 icon font，再匯出 HTML 可用 PNG。

#### Human Explanation

金額欄位可能同時混用數字字型與 icon 字型。執行期 `UIFont` / style 常回報主要文字字型（例如 Condensed Semibold），但 `U+F800` 這類貨幣／籌碼符號可能只存在於另一個 Icons Regular 檔案。

若把 body font 名稱直接寫成 glyph source，會誤導後續還原與 HTML import。正確做法是：

1. 先從文字內容取得 Private Use Area codepoint。
2. 對可用字型檔做 cmap 查詢，確認哪個檔案真正定義該 codepoint。
3. 以該字型檔匯出 glyph PNG，或用 `@font-face` 直接載入字型。
4. 執行期 atlas UV / screen crop 仍只當 visibility evidence。

#### Trigger

- 已確認 inline icon 是字型 glyph，但仍找不到可 import 的獨立 PNG。
- Runtime 回報的 font 名稱與交付包內實際含該 codepoint 的字型檔不一致。
- Viewer 需要可重用的 `<img>` 或 `@font-face` 資產。

#### Evidence

- Tool: font cmap inspection（例如 fontTools）、local recovered font packages、runtime codepoint identity。
- Sanitized excerpt: `U+F800` 出現在金額文字；`KamaGames Icons Regular` cmap 對應 glyph `kg_chip`；同目錄 Condensed Semibold 不含該 codepoint。
- Evidence path: 字型檔與匯出 PNG 留在 `<PROJECT_ROOT>` shared fonts / UI resource map。

#### Generalized Lesson

1. Runtime body-font 名稱 ≠ 一定是 PUA icon source file。
2. 用 cmap 驗證候選字型是否含目標 codepoint，再標記 `sourceFontFile`。
3. HTML import 優先用字型檔或由其渲染的透明 PNG；screen crop 不升級為 source。
4. 文件中分開記錄：body/style font、icon source font、generated PNG、screen reference。

#### Agent Action

找到 PUA icon 後，除了記 runtime font/UV，還要在可用字型檔中做 cmap 確認，並產出 HTML 可引用資產。

#### Goal / Action / Validation

- Goal: 取得可 import 的 icon glyph source，且不誤指 body font。
- Action: codepoint → cmap across candidate fonts → render or `@font-face` → update resource map.
- Validation or reference source: only the font whose cmap contains the codepoint is labeled source; PNG dimensions and codepoint are recorded; screen crops stay labeled reference.

#### Applies When

- Unity UI mixes text fonts and icon fonts.
- Private Use Area marks appear next to amounts or stake labels.

#### Does Not Apply When

- Icon already has an independent Sprite / atlas handle identity.
- No recovered or authorized font file is available; then keep runtime UV + screen reference only.

#### Validation

- cmap hit for the target codepoint in exactly the claimed source font.
- Generated PNG or `@font-face` path is referenced by the viewer.
- Resource map distinguishes body font, source font file, and screen reference.

#### Promotion Target

- `workflow/apk-analysis/execution-flow.md`
- `feedback/history/apk-analysis/unity-il2cpp/2026-09-14_111500-unity-inline-icon-may-be-private-use-font-glyph.md`

#### Required Linked Updates

- 更新 `feedback/history/apk-analysis/README.md`
- 更新 `feedback/history/apk-analysis/unity-il2cpp/README.md`
- 更新既有 PUA glyph lesson 的 Linked / companion note
- Project font paths 留在 `<PROJECT_ROOT>`
