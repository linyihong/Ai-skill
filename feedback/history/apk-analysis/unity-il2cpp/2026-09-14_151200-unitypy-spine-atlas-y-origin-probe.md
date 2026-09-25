> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - UnityPy-exported Spine atlas may use top-left xy

Status: candidate

#### One-line Summary

從 Unity `__data` 用 UnityPy 匯出的 Texture2D PNG，搭配 Spine `.atlas` 切片時，先驗證 `xy` 是 top-left 還是 bottom-left；錯誤座標會切到完全無關的區域。

#### Human Explanation

Spine atlas 規格常把 `xy` 當 bottom-left，但 UnityPy 匯出的 PNG 列順序可能已是 top-left。同一組座標兩種 Y 解釋會得到不同圖塊：正確的是可辨識符號，錯誤的是爆炸粒子或身體部位碎片。

#### Trigger

- 從 UnityCache `__data` 匯出 atlas PNG + `.atlas` TextAsset。
- 切片結果看起來不像目標符號。

#### Evidence

- Tool: UnityPy Texture2D export + Spine atlas region crop.
- Sanitized observation: for one cabinet item atlas, treating `xy` as top-left matched known letter/wild art; bottom-left Y flip produced unrelated VFX/body tiles.
- Evidence path: `<PROJECT_ROOT>/<App>/slot/from-cache/` local export (gitignored) → curated `docs/slots/<cabinet>/assets/`.

#### Generalized Lesson

1. 先用一個已知區域（例如字母 A 或 WILD）做 A/B 兩種 Y 原點切片。
2. `rotate: true` 時 packed 寬高互換，還原後再旋轉。
3. 不要把錯誤切片提交進可分享 `assets/`。

#### Agent Action

切第一張前保存兩種 Y 原點候選，目視確認後再批量切片。

#### Goal / Action / Validation

- Goal: 可分享 symbol map 顯示正確符號，不是隨機 atlas 碎片。
- Action: dual-origin probe → choose → batch crop.
- Validation: at least one low-letter and one wild crop match paytable art.

#### Applies When

- Unity Spine/libGDX atlas + UnityPy-exported PNG.

#### Does Not Apply When

- Runtime UV/`UIAtlasRect` export already gives pixel boxes in image space.

#### Validation

- [ ] Known glyph/wild crop matches client art
- [ ] Curated assets committed only after probe

#### Related Failure Patterns

- 無

#### Promotion Target

- N/A

#### Required Linked Updates

- N/A
