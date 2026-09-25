> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-14 - Variable-lines fixtures are reel-major jagged, not row-major

Status: candidate

#### One-line Summary

`variable-lines` 機台的 `visibleGrid` 是 **外層=轉軸、內層=該軸上→下符號**（可 jagged）；HTML / validator 若當固定盤面 row-major 會把軸數畫錯，且 orientation 常數不可再硬編碼成只有一種。

#### Human Explanation

固定 paylines / all-ways 盤面是矩形 row-major：`outer=row-top-to-bottom`，`inner=reel-left-to-right`。variable-lines 的 RESULT 是每軸高度不同，fixture 用一條 outer array 存一根軸；長度對應 `reelHeights`。若產生器仍用「第一列長度當轉軸數」，會得到錯誤欄數（例如把某軸高度 5 當成 5 軸）。

Spine 角色 high 符號常是 atlas 零件，不能當 paytable 完整圖；中期可用官方賠付表木框裁切補完整构图。

#### Trigger

- `winMode=variable-lines` 或 `layout.heightModel=variable-per-reel`。
- `visibleGrid` 各 outer 長度不等，或 outer 長度 = reelCount。
- symbol-map 樣本表頭軸數 ≠ 機台 reelCount。

#### Evidence

- Sanitized: one 6-reel cabinet had fixture outer lengths matching `reelHeights` product vs HUD line count; generator previously labeled columns from first inner length.
- Fix path: sample `gridLayout=reels-jagged` in symbol-map generator; schema allows reel-major orientation pair; validator derives jagged grid from `reelHeights`.

#### Generalized Lesson

1. `visibleGridOrientation` 必須與實際儲存一致；variable-lines 用 `outerArray=reel-left-to-right` + `innerArray=row-top-to-bottom`。
2. HTML `renderGrid` 對 jagged：欄 = outer 長度（軸），列 = max 高度，短軸空格占位。
3. High 角色若只有 Spine 零件，優先賠付表／同轉截圖完整构图，不要提交 atlas 碎片當主圖。

#### Agent Action

新增 variable-lines 樣本時：寫對 orientation + `reelHeights`；重跑 `generate_symbol_map.py` 與 `validate_spin_package.py`；確認表頭「转轴 0..N-1」且 N=reelCount。

#### Goal / Action / Validation

- Goal: symbol-map 樣本盤面 x 軸 = reelCount；符號圖是完整可辨識構圖。
- Action: fix orientation/schema/generator/validator；paytable crop highs。
- Validation: package validate PASS；模擬 detection 顯示 `cols=reelCount` + `reels-jagged`。

#### Promotion Target

- `workflow` / project slots docs：fixture orientation + generator jagged path（已落到 target repo scripts/docs）。

#### Related

- `2026-09-14_144100-slot-variable-lines-is-third-win-mode.md`
- `2026-09-14_151200-unitypy-spine-atlas-y-origin-probe.md`

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
