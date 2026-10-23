> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-23 - Line-win / coin particles restore via UnityPy or AssetRipper; not Il2CppDumper

Status: candidate

#### One-line Summary

進桌線獎飛幣／火花是 bundle 內 `ParticleSystem` + material／texture；UnityPy 可讀模組預設值，AssetRipper 才接近 Editor 重播。Il2CppDumper／純 AssetStudio 貼圖瀏覽不能當成「粒子已還原」。

#### Human Explanation

畫面上中獎籌碼周圍的金色火花、弧光、飛幣，常被誤以為「再 dump 一次 C# 或開 AssetStudio 就能還原」。實際資產在 UnityCache `__data`：

| 層 | 內容 |
| --- | --- |
| Host GO | 例：`coins_fx`、`sparks`／`sparks_left`／`sparks_right`、`arc_fx_*`、`front_fx` |
| Component | `ParticleSystem`（Initial／Emission／Shape…）+ `ParticleSystemRenderer` |
| Material／貼圖 | 例：`slots_particle_02`、`slots_particle_34`、line-payout coins／firework textures |

UnityPy 對 `__data` 可直接 `read()` 出上述模組 scalar 與 material 名。AssetRipper 另產可匯入 Prefab。Il2CppDumper 只給驅動類名（方便 Frida）。AssetStudio 適合翻貼圖，粒子曲線保真弱。

Live 截圖證明 visibility；身分以 GO／material 名對 bundle dump。`isPlaying` 與觸發時序仍要特效播時的 Frida，且避免無界 `FindObjectsOfTypeAll`。

#### Trigger

- 問「能不能用 Il2CppDumper／AssetStudio／AssetRipper 還原粒子」卻未先 UnityPy 盤點。
- 只匯出 `particle_*.png` 就標粒子還原完成（停在 `texture-only`）。
- 用 Dumper 的 `dump.cs` 宣稱已有 ParticleSystem 資產。

#### Evidence

- Tool: UnityPy on UnityCache `__data`；optional AssetRipper；Frida／screencap while FX plays.
- Sanitized pattern: ParticleSystem count + named hosts (`coins_fx`／`sparks*`／`arc_fx*`) + material names; mature labels `params-from-bundle`／`live-matched`／`editor-replayable`.
- Evidence path: project-local `interface-analysis/*particle*`（勿把金額寫進 reusable lesson）。

#### Generalized Lesson

1. **粒子還原優先序：** UnityPy 參數表 →（可選）AssetRipper Editor 重播 → Frida live match；Il2CppDumper／AssetStudio 不是主路徑。
2. **成熟度標籤：** `texture-only` ＜ `params-from-bundle` ＜ `live-matched` ＜ `editor-replayable` ＜ `game-wired`。
3. **Component 無名：** 經 `m_GameObject` 取 host；不要因 `m_Name` 空判定「沒有粒子」。
4. **共用 pack：** line-payout 貼圖可能在 `slots.common`；feature pack 與 common 都要對 material 名。

#### Agent Action

粒子任務先 pull `__data` + UnityPy inventory；文件寫明工具邊界與成熟度；專案證據留專案目錄。步驟見 [`analysis/apk/workflows/unity-asset-parse-flow.md`](../../../analysis/apk/workflows/unity-asset-parse-flow.md)。

#### Goal / Action / Validation

- Goal: 正確選擇工具，避免假還原。
- Action: Cache dump → UnityPy modules → live GO match → optional AssetRipper。
- Validation: 文件含 ParticleSystem count／host／material；未把 Il2CppDumper 標成資產來源。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP + Addressable／UnityCache feature packs with ParticleSystem FX。
- Does not apply: 純 Spine 符號動畫（無 Particle）；Web 靜態 mock；非 Unity 客戶端。

#### Related

- `2026-09-21_091600-slot-three-visual-planes-and-in-cabinet-anim-stack.md`
- `2026-09-07_143800-unity-feature-art-lives-in-unitycache-not-apk.md`
- [`analysis/apk/workflows/unity-asset-parse-flow.md`](../../../analysis/apk/workflows/unity-asset-parse-flow.md)

#### Validation

新粒子任務文件必須標成熟度標籤，且未把 Il2CppDumper／純 AssetStudio 貼圖瀏覽寫成 ParticleSystem 資產來源；UnityPy dump 含 host GO＋material 名（或明確 blocked）。

#### Promotion Target

- [`analysis/apk/workflows/unity-asset-parse-flow.md`](../../../analysis/apk/workflows/unity-asset-parse-flow.md)
- [`analysis/apk/tools-and-failures.md`](../../../analysis/apk/tools-and-failures.md) §Unity IL2CPP / AssetBundles

#### Required Linked Updates

- Updated: `analysis/apk/workflows/unity-asset-parse-flow.md`（新建）
- Updated: `analysis/apk/workflows/README.md`、`analysis/apk/README.md`、`analysis/apk/tools-and-failures.md`
- Updated: `feedback/history/apk-analysis/README.md`、`feedback/history/apk-analysis/unity-il2cpp/README.md`
- Not applicable: workflow/apk-analysis execution-flow（orchestration 未改觸發條件）
