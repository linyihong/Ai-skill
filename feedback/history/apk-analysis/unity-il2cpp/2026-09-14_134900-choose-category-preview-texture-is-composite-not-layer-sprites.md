> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Choose-category preview_texture is a composite, not separable art sprites

Status: validated

#### One-line Summary

Splash／choose-category 中央 hero 常掛在 `preview_texture`（`SpriteName` 空）；畫面上看似主體／光暈／底座／字標分層，多半是合成預覽內的藝術層，不是多張可獨立 atlas 匯出的 UIImage sprite。

#### Human Explanation

Shared HUD chrome 通常有具名 `SpriteName`，可用 atlas／hash 還原。中央 marketing art 常走另一路徑：

1. UIImage／Renderer 名類似 `preview_texture`、`game_title`。
2. `get_SpriteName()` 為 null；loaded `Texture2D` 清單也常沒有對應的 splash art 名。
3. 視覺上可切「光暈／主體／底座／字標」帶，但 runtime 沒有對應的獨立 sprite id。

常見誤配：

- 把同場景 `left_image`／`right_image` 當成 hero 側翼（常是 lobby／play 按鈕邊條）。
- 把 3D table／menu logo `Texture2D` 當成 splash 字標。
- 把大廳 circular tile／layouts PIC 當成 splash 中央圖。
- 把進桌後 feature-bundle 背景 Sprite 當成 splash hero。

正確標註：`verified-composite-identity`／`visual-layer-in-composite`／`node-identified-texture-unresolved`；screen band crop 只做 visibility。要拆原檔需 CDN preview URL、上游行銷包，或 RenderDoc——不是 MSE。

#### Trigger

- Inventory 把中央插畫拆成多個 O-id 卻全部假陽性 PNG。
- Agent 用已載入的 table logo 或 lobby tile「完成」splash 標誌還原。
- `preview_texture` 存在但 SpriteName／Texture2D name 對不上畫面構圖。

#### Evidence

- Tool: Frida UIImage／Texture2D name sweep；cabinet bundle vs splash screen crop compare。
- Sanitized excerpt: splash 浮島構圖 ≠ feature-bundle 進桌背景；table logo 字樣與 splash 字標不同；lobby tile 為圓形入口圖。
- Evidence path: `<PROJECT_ROOT>` hero composition notes／resource map。

#### Generalized Lesson

1. 先分 plane：named chrome vs preview／ScreenTexture composites vs lobby tiles vs in-cabinet backgrounds。
2. 視覺分層 ≠ runtime 分檔；無 SpriteName 時不要發明多張 source PNG。
3. 否決清單寫進 inventory：table logo、lobby tile、button chrome、in-cabinet background。
4. `game_title` 等節點可單獨存在，仍需獨立證明 texture；不可沿用同 keyword 的已載入 Texture2D。

#### Agent Action

對 `preview_texture` 類物件標 composite；子視覺層標 visual-layer；維護 rejectedFalsePositives。還原步驟見同分類 lesson `restore-unity-preview-composite-via-main-thread-or-cdn`；不要把可執行步驟寫進 technique-free viewer HTML。

#### Goal / Action / Validation

- Goal: 正確描述 hero 組成，避免假還原。
- Action: Identity dump → plane 分類 → 否決誤配 → screen crop only。
- Validation: 文件寫明 composite／unresolved；誤配項列 rejected；不得對錯 plane PNG 標 `verified-matched`。

#### Applies When

- Unity choose-category／machine splash with central marketing preview.

#### Does Not Apply When

- In-cabinet reel／HUD sprites with proven SpriteName in feature bundle.
- Separable chrome already exported via runtime atlas.

#### Validation

同屏 chrome 可具名還原，而中央 hero 仍標 composite；至少列出兩類已否決誤配。

#### Promotion Target

- `analysis/apk/tools-and-failures.md`
- `workflow/apk-analysis/execution-flow.md`

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `analysis/apk/tools-and-failures.md`
- `workflow/apk-analysis/execution-flow.md`
