> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Unity inline icon may be a private-use font glyph

Status: candidate

#### One-line Summary

畫面上的小圖示若不在 active `UIImage`／Sprite 清單，先檢查相鄰 `UIText` 是否含 Private Use Area codepoint；它可能是字型字形，不是遺失的 PNG。

#### Human Explanation

數值前的貨幣、籌碼、票券或狀態標記，外觀像獨立 icon，但 Unity client 可能把它編進文字字串。若只搜尋 Sprite、cache PNG 或相似圖片，容易把外觀相近但語義無關的素材誤配上去。

這類圖示應保留兩種不同產出：

1. **身分證據**：文字內容中的 codepoint、實際字型、size/style，以及 `Font.GetCharacterInfo` 回傳的 glyph UV。
2. **可見性對照**：畫面 crop 或 contact sheet，明確標成 screen reference。

動態字型 atlas 可因 session、字型大小或已請求字元集合而重建；UV 是當次執行證據，不應被當成跨版本固定 source rect。

#### Trigger

- 畫面有小 icon，但 active `UIImage.get_SpriteName()` 找不到對應項目。
- 同一 icon 緊貼數值，且會隨文字顏色、大小或狀態一起變化。
- Sprite/cache 搜尋只得到外觀相近、用途不相符的候選。

#### Evidence

- Tool: IL2CPP runtime object inspection、`UIText.get_Text()`、`UIFont.get_Font()`、`Font.GetCharacterInfo()`。
- Sanitized excerpt: 某個 Private Use Area codepoint同時出現在多個數值欄位；指定字型與實際 size/style 可回傳有效 glyph UV。
- Evidence path: 具體 codepoint、字型名稱、UV 與畫面 crop 留在 `<PROJECT_ROOT>` 的 UI resource map。

#### Generalized Lesson

1. Sprite 清單沒有圖示時，不要直接升級到像素猜測；先檢查 `UIText`／rich text。
2. 對 Private Use Area 字元記錄 codepoint、font、size、style、texture size 與 `CharacterInfo.uv`。
3. 同一 codepoint 在不同字型或 size 下可能尚未進 atlas；`GetCharacterInfo=false` 不代表整個 UI 沒有該 glyph。
4. 只有 codepoint/font/UV 與畫面位置一致時，才能標 `verified-generated-glyph`。
5. 字型 glyph 不是獨立 source PNG；screen crop 只能作 visibility reference。
6. 外觀相近的 ticket、coin 或 chip PNG 必須拒絕，除非有 runtime object identity 鏈。
7. Runtime body/style font 名稱不一定是 PUA icon 的來源檔；要對候選字型做 cmap 查 codepoint，另見 [`2026-09-14_113500-pua-icon-glyph-may-live-in-separate-icons-font.md`](2026-09-14_113500-pua-icon-glyph-may-live-in-separate-icons-font.md)。

#### Agent Action

遇到「看得到 icon、找不到 Sprite」時，先 dump 相關文字物件與 codepoint，再解析實際字型的 glyph metadata。文件中把 `generated glyph`、`standalone image` 與 `screen reference` 分開標示。

#### Goal / Action / Validation

- Goal: 判斷畫面小圖示是獨立素材或文字字形，避免錯配圖片。
- Action: `UIText` 內容 → codepoint → font/size/style → `GetCharacterInfo` → texture/UV → 畫面位置交叉比對。
- Validation or reference source: codepoint 出現在目標文字；`GetCharacterInfo=true`；UV 在字型 texture 範圍內；同一符號可在多個欄位重現。

#### Applies When

- Unity／IL2CPP 自訂 UI、rich text、金額或狀態文字。
- 圖示隨文字樣式一起變化，且沒有 active Sprite identity。

#### Does Not Apply When

- 已有 active `UIImage`／Sprite 或 custom atlas handle 可直接證明獨立素材。
- 多 draw-call、3D mesh、Spine 或 shader effect；這些仍需對應的 render path。

#### Validation

- 搜尋 target `UIText` 的原始字串並列出非一般文字 codepoint。
- 以實際 font、size、style 呼叫 `Font.GetCharacterInfo`。
- 驗證 UV 非零且落在 texture bounds。
- 確認 viewer 將 crop 標為 screen reference，沒有宣稱為 standalone PNG。

#### Promotion Target

- `workflow/apk-analysis/execution-flow.md`
- `workflow/apk-analysis/execution-flow.yaml`
- `runtime/onboarding/apk-analysis-completion.md`

#### Required Linked Updates

- 更新 `feedback/history/apk-analysis/README.md`
- 更新 `feedback/history/apk-analysis/unity-il2cpp/README.md`
- 更新 `workflow/apk-analysis/README.md`
- Project-specific codepoint、字型名稱、UV 與畫面證據留在 `<PROJECT_ROOT>`
