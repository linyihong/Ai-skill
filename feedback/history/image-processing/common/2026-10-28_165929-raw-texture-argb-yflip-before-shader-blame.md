> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-28 - Raw texture dumps: try ARGB + Y-flip before blaming shaders

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本輪在行動端影像快取的 typ=5 方圖上，以 ARGB 通道解讀 + 垂直翻轉後，與同畫面的 live 截圖中段對齊且色調正常；細節與 hash 留在專案文件。

#### One-line Summary

從記憶體／磁碟快取撈到的 raw 四通道紋理若出現綠膚、洋紅天空、整體偏暗，先假設 **通道順序（ARGB／BGRA／RGBA）與 Y-flip** 解錯，再談缺 shader、缺 `_Tint` 或「防盜混淆」。

#### Human Explanation

引擎畫面上的成品常是 Texture + Material + Shader 的合成結果，因此逆向時很容易把「怪色調」直接歸因於沒套 shader 或故意把資料塞進通道遮罩。實務上更常見、也更便宜的失敗是：快取裡存的是 **尚未依顯示慣例重排的位元組**——例如把 ARGB 當 RGBA 讀，再疊上引擎紋理常見的 **Y 軸顛倒**。這種誤讀會穩定產生「健康膚色變綠、藍變洋紅、紅藍對調」的外觀，看起來很像缺著色，但其實只要改通道索引與翻轉就能對回 live 截圖的中段色調。真正的 shader／Material 遺失通常還會伴隨錯誤的邊緣合成、缺少自發光層或整塊仍是遮罩灰階；通道誤讀則是整張圖的色相系統性錯位，線稿與構圖仍清楚。

#### Trigger

- 從 app 影像快取、`Texture2D` raw bytes、或類似 header+RGBA 區塊 dump 出圖
- 構圖／標題可辨，但膚色偏綠、天空／雲偏洋紅、紅藍對調
- 有人懷疑「防盜演算法」或「一定要還原 Material `_Color`」
- 同一資源的螢幕截圖色調正常

#### Evidence

- Tool: header+payload square texture decode；channel-order sweep（RGBA／ARGB／BGRA）× optional Y-flip；masked NCC vs live crop（避開 UI chrome）
- Sanitized excerpt: 誤用 RGBA 時膚色／天空系統性錯色；改 ARGB + Y-flip 後與 live 中段對齊且無 jackpot／volatility 疊層
- Evidence path: project lobby-tile clean-pull notes under `<PROJECT_ROOT>/docs/…`（細節不進本檔）

#### Generalized Lesson

1. **先分三層**：快取 raw 位元組 ≠ GPU 最終著色 ≠ 螢幕 UI 疊層。怪色先查 raw 解碼，再查 shader，最後才剝 overlay。
2. **通道掃一遍**：至少試 `RGBA`、`ARGB`（取 `[1,2,3]` 為 RGB、`[0]` 為 A）、`BGRA`；各搭配 `flip_v` on/off。
3. **用 live 中段當裁判**：與螢幕截圖比對時遮掉頂／底 UI chrome；目視優先於單一 NCC 分數（分數可能因裁切／縮放偏低，但色相對了就能定案）。
4. **不要過早結論「加密／防盜」**：明文 header + 固定 `w*h*4` payload 且通道修正後可對齊 live，就不是加密問題。
5. **Shader／Material 路徑仍保留**：通道修正後仍缺自發光、整張像遮罩、或只有灰階通道有意義時，再去讀 `material.GetColor`／RenderDoc。

#### Agent Action

- 先做：確認 payload 長度＝`w*h*4`（或已知格式）→ 通道×翻轉矩陣解碼 → 與 live 中段目視／遮罩相關 → 寫入乾淨資產。
- 不要做：一看到綠膚就寫「需要 shader 還原演算法」；把 live 截圖的 jackpot／volatility 當原圖；只試 BGR swap 就放棄。

#### Goal / Action / Validation

- Goal: raw 紋理解碼色調與螢幕成品中段一致，且無 UI chrome。
- Action: ARGB／BGRA／RGBA × Y-flip 掃過 → 選與 live 對齊者落地。
- Validation or reference source: 目視膚色／天空／品牌色正確；遮罩區與 live 相關上升；錯誤解碼變體明顯錯色。

#### Applies When

- 四通道 raw 紋理／影像快取 dump
- 引擎紋理可能 Y-flip 或非 RGBA 記憶體順序
- 有（或可再截）同畫面 live 參考

#### Does Not Apply When

- 已確認是多通道遮罩貼圖（R/G/B 各表不同資料）且設計上就不是可直出的最終色
- 只有壓縮圖（ASTC／ETC）尚未正確解壓
- 問題僅是 UI overlay，底圖色調已正確

#### Validation

同一 blob 用錯誤通道解碼應再現系統性錯色；正確通道＋必要翻轉後應與 live 中段色相一致。換另一張同格式快取圖重複一次。

#### Promotion Target

- `workflow/image-processing/execution-flow.md`（raw texture decode 檢查清單）
- 相關專案腳本的 decode 預設與文件說明（project-local）

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- [x] `feedback/history/image-processing/common/README.md`
- [x] `feedback/history/image-processing/README.md`（category 已列 common）
- [x] `workflow/image-processing/execution-flow.md` 短指引（raw texture channel sweep）
- [x] 專案側 decode 預設改為 ARGB+Y-flip（證據不進本檔）

#### Transfer Contract

- Inputs: 已知寬高的四通道 raw 紋理 + 可選 live 參考圖
- Mechanism: 通道順序／Y 軸慣例誤讀會造成系統性色相錯位，外觀酷似缺 shader
- Decision: 先掃 ARGB／BGRA／RGBA × flip，對齊 live 中段後再追究 Material
- Verification: 錯誤變體仍錯色；正確變體與 live 中段一致且無 chrome
- Counterexample: 真·通道遮罩貼圖或未解壓的 GPU 壓縮格式
- Transfer case: 任一 Unity／行動端影像快取的 header+RGBA 區塊 dump

#### Intelligence Review

- Decision: method-only
- Reason: 解碼檢查順序已寫入 image-processing execution-flow；尚未升格為全庫 enforcement
- Target: `workflow/image-processing/execution-flow.md`
- Revisit: 第二個無關專案或第二種快取格式重跑通過後改 `validated`
- Discovery: none

#### Abstraction Check

- Swap-project test: PASS — 任何四通道 raw 紋理 dump 都可先做通道×翻轉矩陣，再決定是否需要 shader／Material
- Project evidence kept in: `<PROJECT_ROOT>/docs/` 的 clean-tile／decode 紀錄
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
