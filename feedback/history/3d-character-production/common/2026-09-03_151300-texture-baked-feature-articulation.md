> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-03 - 貼圖化特徵先以遮擋約束的局部合成驗證

Status: validated

#### One-line Summary

當應動的臉部特徵被烘在貼圖而非獨立幾何時，先分離可動前景、被遮住的底圖與開口遮罩，再以 UV-to-surface mapping 統一成對方向；未裁切 overlay 或直接切網格都不是預設解。

#### Human Explanation

「眼睛應該轉動」不代表資產裡一定存在可旋轉的眼球。若虹膜、眼白、眼線與眼皮都烘在同一
張貼圖，直接旋轉頂點只會拉花；切洞放球則會破壞原本負責角色辨識的眼皮與睫毛。單純在畫面
上疊一個可動圖層也不夠，因為前景移到邊界時會蓋到本應遮住它的表面。

可行的局部適配是把原本隱含在繪圖裡的圖層關係顯式化：可動特徵是 foreground、其下方需要
補回 background、畫好的開口是 occlusion mask。這保留原輪廓，同時讓前景在邊界自然消失。

#### Trigger

- 應動的臉部特徵只存在於 diffuse/base-color texture。
- 形變控制存在，但可見結果是貼圖拉伸或沒有獨立運動。
- 幾何切割會移除眼皮、嘴唇、眼線或其他 identity-critical 細節。
- 成對特徵在 texture atlas 中使用不同旋轉的 UV islands。

#### Evidence

- Tool: texture pixel probe、UV-to-surface tangent probe、局部 inpaint、互動 consumer readback。
- Sanitized excerpt: 初始局部移動能產生位移，但遮擋遮罩外溢、底圖被暗輪廓平均成灰色，且兩個 UV island 沿各自影像長軸移動時在 3D 中方向不一致。改用取樣後的色度分類、只以確認的底圖像素作擴散邊界、由 UV 映射解共同世界方向，並取成對構造的較小位移上限後，雙向端點、回中、區外零漂移與既有表情組合均通過。
- Evidence path: 原始素材、局部圖集、截圖與 consumer 測試留在 `<PROJECT_ROOT>`。

#### Generalized Lesson

1. 先讀 representation，不從語意名稱推測一定有幾何或骨骼。
2. 在修改前保存 neutral consumer baseline。
3. 將可動前景、底圖與遮擋遮罩分成三個 artifact。
4. 分類閾值須由實際像素取樣；陰影區通常使固定亮度閾值失效。
5. inpaint 只採用確認屬於底圖的像素作邊界，不能讓輪廓線與抗鋸齒污染擴散。
6. 成對 UV islands 先映射到共同 world/local direction，再各自轉回 texture direction。
7. 成對構造共享位移幅度並取較緊邊界；極限值與日常控制值分開。
8. consumer 驗證至少包含：雙向差異、成對一致、區外不變、回到 neutral、與既有動作組合。

#### Agent Action

遇到 texture-baked facial feature 時，先產出 representation probe 與局部預覽條，再決定是否
值得接入 consumer。只有在 neutral parity、遮擋與共同方向都成立後才加入互動控制；若只在
特定 viewer 成立，record 明記其 consumer-local method，不冒充底層資產已有獨立 rig。

#### Goal / Action / Validation

- Goal: 在不破壞 identity-critical 表面細節的前提下，讓貼圖化特徵可控地運動。
- Action: 執行 foreground/background/mask 分離、校準補圖、UV-to-surface 方向映射與共同限幅。
- Validation or reference source: 局部預覽須含 untouched original 與 neutral composite；consumer 測試須驅動真實控制項並驗證目標 ROI 變化、非目標 ROI 不變及回中無漂移。

#### Applies When

- 可動特徵烘在 texture atlas，且交付面允許 runtime 局部 texture update。
- 原輪廓比新增通用幾何更能保留角色辨識。

#### Does Not Apply When

- 資產已有可用的獨立幾何、骨骼與遮擋拓撲；此時應優先修復既有 rig。
- 目標 runtime 不允許動態 texture update，或遠景 mipmap 一致性是 blocking requirement。
- 要求底層交換格式本身攜帶 gaze rig；viewer-local adaptation 不能冒充 asset-level 能力。

#### Validation

- 原圖與 neutral composite 在目標 ROI 內達 parity。
- 兩個方向的端點彼此可區分，成對特徵不發散。
- 非目標 ROI 無可見漂移，回中無累積誤差。
- 既有表情、頭部或骨架動作仍可組合。

#### Promotion Target

- `workflow/3d-character-production/facial-expression.md`
- `validation/scenarios/3d-character-production/`

#### Required Linked Updates

- workflow 加入 texture-baked feature articulation 的執行順序與 owner 邊界。
- scenario 覆蓋未裁切 overlay、各 UV island 獨立方向、污染 inpaint 邊界與 viewer-local 能力誤宣稱。
- 專案素材、角色名稱、絕對路徑與局部像素證據只留在 consumer project。
