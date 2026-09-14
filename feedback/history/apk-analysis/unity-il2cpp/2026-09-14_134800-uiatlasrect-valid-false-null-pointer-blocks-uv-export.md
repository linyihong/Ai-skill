> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - UIAtlasRect may expose PixelSize while valid=false and Pointer=null

Status: validated

#### One-line Summary

Custom atlas handle 即使有 resource name 與 `PixelSize`，若 `valid=false` 且 atlas `Pointer`／entry id 為 null，則 UV 常是 unit rect，不可做 runtime `ReadPixels` 匯出；應改標 `screen-isolated` 或等 packed on-screen entry。

#### Human Explanation

自訂 atlas handle 還原流程假設：handle 已綁定 atlas page，可讀 UV，再從 active `RenderTexture` 裁切。

實務上會遇到「半載入」狀態：

- `Name`／`PixelSize`／`Loaded`／`Scale` 有值
- `valid=false`，on-screen／enabled counter 為 0
- atlas Pointer 為 null → 無 atlas address、無 entry id
- UV／`GetUV` 回傳 `(0,0,1,1)` unit rect

此時若仍跑 main-thread ReadPixels，會裁到整張 atlas 或失敗。另：`get_Sprite()` 可能回傳不可讀指標，但 UIImage 上的 sprite 欄位仍可能握有同一 handle——可用欄位 fallback，仍要以 `valid`＋Pointer 當 export gate。

Icons font 裡語意相近的「時間」PUA glyph（例如沙漏）也不等於 countdown UIImage sprite；必須比對字形後再決定路徑。

#### Trigger

- Frida 已命中 SpriteName／UIImage name，但 runtime atlas export 一直 empty。
- Handle dump 顯示 PixelSize 正確，UV 卻是 unit rect。
- Agent 把 imgcache 弱 MSE 或 Icons font 同義 glyph 當成 source。

#### Evidence

- Tool: Frida IL2CPP field／method dump on custom atlas handle；main-thread export hook。
- Sanitized excerpt: 小尺寸 timer-style UIImage 有 name + PixelSize，但 `valid=false`、Pointer null；同場景其他 chrome 則有 atlas entry 可匯出。
- Evidence path: `<PROJECT_ROOT>` interface-analysis runtime notes。

#### Generalized Lesson

1. Export gate：`valid==true` 且 atlas Pointer／entry id 可讀，且 UV 不是 unit rect。
2. `get_Sprite()` 失敗時可嘗試 UIImage sprite 欄位 fallback；仍過 gate 才匯出。
3. 未過 gate → 標 `screen-isolated`／`runtime-atlas-blocked`；不要把 MSE 或單位 UV 裁切升格為 verified source。
4. 語意相近的 Icons-font glyph 必須做字形比對，不可只靠名稱含 time／clock。

#### Agent Action

Dump handle 時同時記錄 `valid`、Pointer、entry id、UV、PixelSize。未過 gate 就停止 ReadPixels。

#### Goal / Action / Validation

- Goal: 避免假匯出；只在 packed handle 上宣稱 runtime atlas PNG。
- Action: Gate → export；否則 screen-isolated + next-step。
- Validation: 通過 gate 仍須 `UV×size == PixelSize == PNG`；未通過不得標 `verified-runtime-atlas-export`。

#### Applies When

- Unity IL2CPP custom atlas handles used by UIImage／UIRenderer chrome.

#### Does Not Apply When

- Plain `UnityEngine.Sprite` with readable textureRect.
- Non-atlas preview composites（另見 preview-texture composite lesson）。

#### Validation

同屏至少一顆通過 gate 的 chrome 可匯出；未通過者 JSON 含 blocked reason 與 gate 欄位。

#### Promotion Target

- `analysis/apk/tools-and-failures.md`
- `workflow/apk-analysis/execution-flow.md`

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `analysis/apk/tools-and-failures.md`
- `workflow/apk-analysis/execution-flow.md`
