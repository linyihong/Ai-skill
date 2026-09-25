> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-21 - Slot cabinets have three visual planes; in-cabinet animation is ID→prefab, not strip art

Status: candidate

#### One-line Summary

同一 cabinet 至少有三個視覺平面（大廳 tile、choose-category splash、進桌 cabinet）。進桌不是把 PNG 排成 5×3：伺服器給符號 ID／座標，客戶端在 Unity UI 層樹上 spawn Spine／Animator／ParticleSystem。

#### Human Explanation

分析或重做 slot 介面時，常把三張「都有野牛／標題」的圖當成同一套零件：

1. **大廳入口 tile** — 圓形／小圖；配色可與 splash 不同（例如暗色霓虹 vs 夕陽橙）。
2. **Choose-category splash** — 中央常是 `preview_texture` 合成預覽（見同分類 `choose-category-preview-texture-is-composite-not-layer-sprites`）。
3. **進桌 cabinet** — feature bundle 的 UI hierarchy：背景 Sprite、Jackpot 橫條、五 reel 底板、金框、共用 HUD。

進桌動畫也不是「一張長 reel 圖在捲」：

| 層 | 技術 | 角色 |
| --- | --- | --- |
| Layout | Unity UI / RectTransform | 固定機台骨架（背景、reel 容器、框、Jackpot、共用 HUD） |
| Symbols | Spine skeleton（常見 authored 60 fps） | 每格符號 idle／落地／中獎演出 |
| State | Animator + OverrideController + clips | 窗口、線數、bonus 狀態切換 |
| FX | ParticleSystem + SpriteRenderer | 加速光、金幣、煙火、線閃 |
| Audio | AudioClip + sound_store | SFX／BGM bed；慶祝窗字級常與 `*_big_win`／`*_huge_win`／`*_super_win` 同名 |

伺服器 spin 結果只提供 strip／符號 ID 與 reel/row；客戶端查表後在 `items`／`animated_items` 等節點掛對應 prefab。協議 ID（短字元）≠ Spine／GameObject 資源名（`high_*`／`wild_*`）。

#### Trigger

- 用 splash／lobby tile 裁切去「還原」進桌符號或背景。
- 以為 layer-stack 示意頁的色塊等於可匯出的分層 PNG。
- 重做遊戲時先拼靜態 5×3 圖，卻缺少 Spine／狀態機／粒子與 shared HUD 邊界。
- 把 lobby 暗色圓標與 splash 暖色合成當成同一張源檔的色偏版本。

#### Evidence

- Tool: UnityCache／Addressable object counts；prefab path dump；同轉 screenshot ↔ 協議 ID 對照。
- Sanitized pattern: 同一 cabinet 的 tile／splash／in-table 構圖與配色可分離；reel prefab 分 `reels_content`、`animated_content`、`combination`、`lines`、`frame`。
- Evidence path: `<PROJECT_ROOT>` cabinet `interface-analysis/` layer-stack viewer + asset inventory（project-local）。

#### Generalized Lesson

1. **先標 plane**：lobby tile ≠ splash preview ≠ in-cabinet stack。三平面各自 inventory，禁止跨平面標 `verified-matched`。
2. **進桌層級用邏輯層（L0…Ln）描述骨架**；不要宣稱已還原精確 Canvas draw-call，除非有 sibling index／sorting／mask dump。
3. **動畫棧分開**：Spine＝符號表演；Animator＝狀態／窗口；Particle＝瞬間特效；**Audio＝SFX／BGM**（見同日 `slot-audio-fourth-stack-and-win-tier-sfx-names`）。缺一就不能說「動畫做完了」。
4. **資料流**：C2S/S2C 符號 ID → client lookup → instantiate／enable prefab；靜態 PNG 只服務 L0–L2／L4 等底板，不驅動 L3 動態。
5. **Shared HUD**（玩家列、下注、旋轉鍵）屬共用 slot UI，不要算進單機台 feature-bundle 資產清單。

#### Agent Action

開任何 slot UI 任務時先問「現在是哪一個 plane？」；進桌重建 checklist 必須同時列：hierarchy、Sprite 底板、Spine、Animator、Particle、協議 ID map、shared HUD 邊界。Viewer HTML（technique-free）只標位置與層名，還原步驟放專案／apk-analysis guidance。

#### Goal / Action / Validation

- Goal: 正確規劃重做／分析範圍，避免混平面與假還原。
- Action: Plane 分類 → 進桌 hierarchy → 動畫三棧 → ID→prefab 對照。
- Validation: 文件分三平面；進桌 inventory 含 Spine／Animator／Particle 計數或明確 not-applicable；符號表同時列協議 ID 與資源名。



#### Amendment 2026-09-21 (idle live dump)

停輪 idle 時 `Spine.Unity.SkeletonAnimation` 的 `FindObjectsOfTypeAll` 可為 **0**。不要因此判定「沒有 Spine」：同場仍有 per-symbol `*_spine_animated_controller`（AnimatorOverrideController）與 `item_*_animated` Animator／GameObject，以及 `AccelerateStart`／`AccelerateProcess`／`AccelerateStop` clip 名。

Agent action：idle 先列 AnimationClip／RuntimeAnimatorController／Animator GO 名；要 track clip（idle／win／land）需在旋轉中或另讀 Override 映射，不能只靠 SkeletonAnimation 實例計數。進桌 checklist 另加 AudioClip／sound_store（見 `2026-09-21_111200-slot-audio-fourth-stack-and-win-tier-sfx-names.md`）。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP slot cabinets with Addressable／UnityCache feature packs + choose-category splash。
- Does not apply: 純 Web 靜態 slot 示意；無 Spine 的簡單 Sprite reel；非 slot 大廳遊戲。

#### Related

- `2026-09-21_111200-slot-audio-fourth-stack-and-win-tier-sfx-names.md`
- `2026-09-21_111500-slot-celebration-window-hunt-continue-cta-not-midreel-gold.md`
- `2026-09-14_134900-choose-category-preview-texture-is-composite-not-layer-sprites.md`
- `2026-09-14_141200-restore-unity-preview-composite-via-main-thread-or-cdn.md`
- `2026-09-14_153800-slot-in-reel-vs-paytable-art-presentations.md`
- `2026-09-08_162200-symbol-dto-id-may-be-one-char-strip-not-resource-name.md`

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
