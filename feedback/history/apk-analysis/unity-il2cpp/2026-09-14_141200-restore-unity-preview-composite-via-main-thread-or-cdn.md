> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Restore Unity splash preview composites via main-thread dump or CDN window

Status: validated

#### One-line Summary

Named `UIImage` with empty `SpriteName`（如 choose-category `preview_texture`／`game_title`）不能走 custom-atlas UV 匯出；要在 Unity main／render-thread callback 上 `get_Texture`→`EncodeToPNG`／`ReadPixels`，或在 splash 開啟窗攔截 CDN preview URL；靜止畫面常不進 render callback。

#### Human Explanation

具名 chrome 用 `SpriteName` + `UIAtlasRect` UV 還原。中央 marketing preview 常是另一平面：

- 物件有 `Object.name`（preview／title 類），但 `get_SpriteName()` 為空。
- 沒有可用的 atlas entry／非 unit UV，Tier-1b atlas 路徑不適用。
- 視覺上可切多層，runtime 仍可能只有一張合成材質（字標或另有節點）。

還原順序（通用）：

1. **Main-thread dump（首選）**  
   Hook UI render callback（例：`UIRenderer.OnWillRenderObject`）。在 callback 內依 `Object.name` 找目標 `UIImage`，呼叫 `get_Texture()`。  
   - `Texture2D` → `ImageConversion.EncodeToPNG`  
   - `RenderTexture` → `RenderTexture.active` + `Texture2D.ReadPixels` + `Apply` + `EncodeToPNG`（略過明顯共用 2048² atlas RT）  
   禁止在 Frida worker thread 做上述呼叫（易 AV／main-thread breakpoint）。

2. **強制重繪**  
   Idle splash 可能完全不進 `OnWillRenderObject`／甚至不呼叫 `Camera.Render`。先關再開 dialog，或觸發會重繪的 UI（選擇器撥動等），再跑 dump。

3. **CDN／preview 下載窗**  
   在進入 splash 的時間窗攔截 URL／下載本體；比事後全庫 imgcache MSE 準。

4. **GPU frame（後備）**  
   texture pointer 不可讀時才用 RenderDoc（或同等）對 UI camera 擷一幀。

成功通常得到一張（或兩張）PNG，標 `verified-runtime-texture-export`／`verified-cdn-preview`。子視覺層仍標 `visual-layer-in-composite`。Screen crop 只做 visibility。

#### Trigger

- Agent 對 `preview_texture` 硬套 atlas UV 匯出失敗或假陽性。
- Worker-thread `get_Texture` AV；idle splash 上 render hook 零觸發。
- 用 table logo／lobby tile／in-cabinet BG MSE「完成」splash hero。

#### Evidence

- Tool: Frida IL2CPP UI render callback + texture encode；optional CDN capture。
- Sanitized excerpt: idle choose-category splash 上 UI render／Camera.Render hook 可不進；chrome atlas 同程序可匯出；preview 物件名存在但 SpriteName 空。
- Evidence path: `<PROJECT_ROOT>` interface-analysis composition notes（project-only）。

#### Generalized Lesson

1. Plane 分流：named chrome atlas ≠ preview composite ≠ lobby tile ≠ in-cabinet BG。
2. Preview 還原優先序：main-thread texture dump → CDN window → GPU frame → screen-isolated。
3. Idle UI 先強制重繪再斷言「texture 不可讀」。
4. 合成成功 ≠ 可拆成多張源檔；無獨立 texture 證據前保持 visual-layer。
5. Viewer HTML 只放畫面／素材對照；可執行步驟放 reusable skill／SOP，不寫進 technique-free viewer。

#### Agent Action

對空 `SpriteName` 的 preview／title `UIImage`：寫 restore plan 到 Ai-skill／workflow，不塞進 viewer HTML。執行時走 main-thread gate；失敗先重開畫面再試 CDN。

#### Goal / Action / Validation

- Goal: 取得可驗證的 preview／title PNG，且不誤標分層源檔。
- Action: Main-thread dump 或 CDN window；必要時 GPU。
- Validation: PNG 構圖對得上 screen crop；標正確 status；分層無獨立 texture 則不得 `verified-matched`。

#### Applies When

- Unity IL2CPP choose-category／machine splash with central marketing preview composites.

#### Does Not Apply When

- Has proven `SpriteName` + valid atlas handle（走 atlas UV export lesson）。
- Plain readable `UnityEngine.Sprite` with textureRect.

#### Validation

同屏 chrome 可 atlas 匯出，而 preview 走本條路徑；至少記錄是否需要 redraw gate。

#### Promotion Target

- `analysis/apk/tools-and-failures.md`
- `workflow/apk-analysis/execution-flow.md`

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `analysis/apk/tools-and-failures.md`
- `workflow/apk-analysis/execution-flow.md`
- 相關：`2026-09-14_134900-choose-category-preview-texture-is-composite-not-layer-sprites.md`
