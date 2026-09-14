> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - After Frida SpriteName, build atlas manifest before imgcache MSE

Status: validated

#### One-line Summary

取得 `UIImage.get_SpriteName()` 之後，下一步應是 **UnityPy atlas manifest**（`Sprite.name → texture + rect → exported PNG`）與 **imgcache hash 直查**；全庫滑窗 MSE 只能當最後 fallback，且對 splash CDN chrome 常無 cabinet bundle 命中。

#### Human Explanation

外部建議常列出三條路：RenderDoc、Frida 遍歷 UI.Image、還原圖集。實務上 **順序** 決定精度：

1. **Frida `get_SpriteName()`** — 執行期身分（Tier-1）。已有 lesson：[`unity-ui-identity-from-loaded-objects-not-screenshot`](2026-09-09_112800-unity-ui-identity-from-loaded-objects-not-screenshot.md)。
2. **Atlas manifest** — 對 **feature bundle**（如 `slots.<feature-slug>`）用 UnityPy 或已匯出 `Sprite_*.png` 索引，精確命中具名 Sprite。
3. **imgcache hash 直查** — splash / choose-category 素材在 CDN cache（hash 檔名），若先前 MSE pass 已記 `spriteName → hash`，直接用 hash 取 PNG，不要重跑 900+ 檔滑窗。
4. **imgcache MSE** — 僅當 2/3 皆無且需弱候選時；>35 拒絕。
5. **RenderDoc** — `SpriteName` 為 null 的 RenderTexture 合成（hero / preview_texture）。

常見錯誤：用 **功能包 PNG 清單** 或 **imgcache MSE-first** 去配 shared splash；shared UI resource id 不一定在 feature bundle 內。

#### 2026-09-14 Revision - custom runtime atlas handles

`UIImage.get_Sprite()` 不一定回傳 `UnityEngine.Sprite`。部分 IL2CPP client
使用自訂 `UIAtlasRect` / atlas handle；這時應先讀取 method 的宣告回傳型別，
再從 handle 取得 resource name、atlas entry id、normalized UV、pixel size、
scale / alpha 與 `UIImage.get_Texture()`。

如果 texture 是 `RenderTexture`，`ImageConversion.EncodeToPNG` 不能直接匯出。
應將 UV 乘上 texture 尺寸得到 pixel rect，再於 Unity main/render thread
callback 內執行 `RenderTexture.active` + `Texture2D.ReadPixels`。從 Frida
worker thread 建立 `Texture2D` 可能命中 Unity 的 main-thread breakpoint；
可掛到該 UI renderer 的 `OnWillRenderObject`，只在目標 resource name
相符時執行一次 bounded export。

#### Trigger

- Frida 已 dump 大量 `spriteName`，但文件仍標 `source-pending` / `screen-reference-only`。
- Agent 想「還原素材」卻只在 imgcache 做像素搜尋。
- Cabinet bundle 已 UnityPy 匯出，卻未建 `name → rect → PNG` manifest。

#### Evidence

- Tool: `export_sprite_atlas_manifest.py`, `resolve_ui_sprites.py`, UnityPy, Frida pass-4 capture.
- Sanitized excerpt: feature-bundle manifest 可精確索引具名 Sprite；shared splash resource 另由 hash cache 或 runtime atlas 處理。當 custom atlas handle 的 UV×textureSize、runtime PixelSize 與匯出 PNG 尺寸一致時，可標記為 runtime-atlas verified。
- Evidence path: `<PROJECT_ROOT>/docs/slots/ui-asset-restoration.md` 與 project-local runtime atlas evidence。

#### Generalized Lesson

1. 有 Frida name → 先查 atlas manifest（cabinet）與 imgcache hash map（CDN），再考慮 MSE。
2. 匯出 manifest：`Sprite.name`, `textureName`, `m_Rect`, `exportedSpritePath`。
3. 全 imgcache 滑窗 MSE 成本高、易假陽性；需 size prefilter + crop map。
4. 分 plane 寫 inventory：cabinet bundle / choose-category CDN / RenderTexture composite。
5. RenderDoc 列為 composite 的 optional pass，不取代 Frida name。
6. `get_Sprite()` 回傳 custom atlas handle 時，依宣告型別解析 UV；不要強制套用 `UnityEngine.Sprite` 欄位。
7. Runtime texture export 必須在 Unity main/render thread 執行，並以
   `UV × textureSize == PixelSize == PNG size` 三方一致作為驗證。

#### Agent Action

Implement or run name-first resolver before claiming asset restoration complete. Document `resolutionCounts` in JSON report. Never promote MSE weak hits to `verified-matched` for identity.

#### Goal / Action / Validation

- Goal: Restorable PNG path per visible UI sprite with proven identity chain.
- Action: Frida identity → declared sprite-handle type → manifest/hash direct；若為 custom runtime atlas，再以 UV + render-thread ReadPixels 匯出；最後才使用 MSE。
- Validation: Report lists tier per sprite；runtime export 必須同時滿足 UV 尺寸、handle PixelSize、PNG 尺寸一致；純色 MSE 假陽性不得標 verified。

#### Applies When

- Unity IL2CPP slots / choose-category splash / cabinet UI inventory.
- UnityCache PNG export available locally (gitignored).

#### Does Not Apply When

- Native Android widgets (use view dump).
- 3D mesh / Spine runtime only (need skeleton path, not Sprite atlas).

#### Validation

Re-run resolver: manifest hit count stable；hash direct rows copy readable PNG。
若走 runtime atlas，確認 handle resource name、atlas entry、UV 與 texture
尺寸已記錄，且匯出 PNG 的 width/height 等於 UV pixel rect。

#### Promotion Target

- `analysis/apk/tools-and-failures.md`（Unity 表加 atlas manifest + hash-direct 列）
- `workflow/apk-analysis/execution-flow.md`（Unity canvas identity 規則補充 pass-5）

#### Required Linked Updates

- 必須：`feedback/history/apk-analysis/README.md`
- 必須：`feedback/history/apk-analysis/unity-il2cpp/README.md`
- 必須：`analysis/apk/tools-and-failures.md`
- Project scripts/docs 留在 `<PROJECT_ROOT>`（不寫入本檔 raw path）
