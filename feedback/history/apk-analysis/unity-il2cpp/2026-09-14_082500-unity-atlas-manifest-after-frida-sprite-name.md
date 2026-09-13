> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - After Frida SpriteName, build atlas manifest before imgcache MSE

Status: validated

#### One-line Summary

取得 `UIImage.get_SpriteName()` 之後，下一步應是 **UnityPy atlas manifest**（`Sprite.name → texture + rect → exported PNG`）與 **imgcache hash 直查**；全庫滑窗 MSE 只能當最後 fallback，且對 splash CDN chrome 常無 cabinet bundle 命中。

#### Human Explanation

外部建議常列出三條路：RenderDoc、Frida 遍歷 UI.Image、還原圖集。實務上 **順序** 決定精度：

1. **Frida `get_SpriteName()`** — 執行期身分（Tier-1）。已有 lesson：[`unity-ui-identity-from-loaded-objects-not-screenshot`](2026-09-09_112800-unity-ui-identity-from-loaded-objects-not-screenshot.md)。
2. **Atlas manifest** — 對 **cabinet bundle**（如 `slots.buffalo_rush`）用 UnityPy 或已匯出 `Sprite_*.png` 索引，精確命中 `high_1` 等 reel 素材。
3. **imgcache hash 直查** — splash / choose-category 素材在 CDN cache（hash 檔名），若先前 MSE pass 已記 `spriteName → hash`，直接用 hash 取 PNG，不要重跑 900+ 檔滑窗。
4. **imgcache MSE** — 僅當 2/3 皆無且需弱候選時；>35 拒絕。
5. **RenderDoc** — `SpriteName` 為 null 的 RenderTexture 合成（hero / preview_texture）。

常見錯誤：用 **功能包 PNG 清單** 或 **imgcache MSE-first** 去配 splash；choose-category id（`choose_category_*`）不在 reel bundle 內。

#### Trigger

- Frida 已 dump 大量 `spriteName`，但文件仍標 `source-pending` / `screen-reference-only`。
- Agent 想「還原素材」卻只在 imgcache 做像素搜尋。
- Cabinet bundle 已 UnityPy 匯出，卻未建 `name → rect → PNG` manifest。

#### Evidence

- Tool: `export_sprite_atlas_manifest.py`, `resolve_ui_sprites.py`, UnityPy, Frida pass-4 capture.
- Sanitized excerpt: `slots.buffalo_rush` manifest 41 Sprite；splash Frida 176 names 中 12 個 choose-category 經 hash 直查 resolved；164 為背景/非 splash 節點（identity-only）；hero 仍 unresolved。
- Evidence path: target project `docs/slots/ui-asset-restoration.md`, `bison-bash/interface-analysis/sprite-resolution-report.json`

#### Generalized Lesson

1. 有 Frida name → 先查 atlas manifest（cabinet）與 imgcache hash map（CDN），再考慮 MSE。
2. 匯出 manifest：`Sprite.name`, `textureName`, `m_Rect`, `exportedSpritePath`。
3. 全 imgcache 滑窗 MSE 成本高、易假陽性；需 size prefilter + crop map。
4. 分 plane 寫 inventory：cabinet bundle / choose-category CDN / RenderTexture composite。
5. RenderDoc 列為 composite 的 optional pass，不取代 Frida name。

#### Agent Action

Implement or run name-first resolver before claiming asset restoration complete. Document `resolutionCounts` in JSON report. Never promote MSE weak hits to `verified-matched` for identity.

#### Goal / Action / Validation

- Goal: Restorable PNG path per visible UI sprite with proven identity chain.
- Action: Frida → manifest index → hash direct → MSE fallback → mark composite unproven.
- Validation: Report lists tier per sprite; cabinet symbols hit manifest; splash chrome hits hash or stays identity-only; no blue-on-blue false verified.

#### Applies When

- Unity IL2CPP slots / choose-category splash / cabinet UI inventory.
- UnityCache PNG export available locally (gitignored).

#### Does Not Apply When

- Native Android widgets (use view dump).
- 3D mesh / Spine runtime only (need skeleton path, not Sprite atlas).

#### Validation

Re-run resolver: manifest hit count stable; hash direct rows copy readable PNG; MSE not required for hash-mapped sprites.

#### Promotion Target

- `analysis/apk/tools-and-failures.md`（Unity 表加 atlas manifest + hash-direct 列）
- `workflow/apk-analysis/execution-flow.md`（Unity canvas identity 規則補充 pass-5）

#### Required Linked Updates

- 必須：`feedback/history/apk-analysis/README.md`
- 必須：`feedback/history/apk-analysis/unity-il2cpp/README.md`
- 必須：`analysis/apk/tools-and-failures.md`
- Project scripts/docs 留在 `<PROJECT_ROOT>`（不寫入本檔 raw path）
