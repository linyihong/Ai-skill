# Texture-baked feature articulation — consumer loop（去敏）

**Run ID**：2026-09-03-texture-baked-feature-articulation-consumer-loop  
**Authorization**：consumer 要求讓貼圖化眼睛可左右移動，並要求把修改流程回饋 workflow  
**Scope**：consumer adaptation + generalized workflow feedback；不是 Phase 4 profile、Phase 6 full dogfood 或 runtime-ready claim。

原始角色、貼圖、局部圖集、腳本、截圖與量測留在 consumer `<PROJECT_ROOT>`；本檔不含專案名、
本機絕對路徑或私人素材識別。

## 結論

單一貼圖內的成對臉部特徵可在不切網格、不新增 detached helper 的情況下產生可見移動：
將可動前景、被蓋住的底圖與遮擋範圍分離，在局部 texture domain 合成，再把共同表面方向
映射回各自 UV island。

consumer 驗證覆蓋：

- 兩個水平方向端點皆相對 neutral 產生可見差異。
- 左右端點彼此可區分，成對特徵未發散。
- 非目標臉部區域 readback 無漂移。
- 回到 neutral 無累積誤差。
- 與既有臉部表情及較慢的頭部動作可組合。
- 可選自然注視序列能產生快速視線跳轉、微幅修正與較慢頭部跟隨；預設關閉以保持測試可重現。

## 失敗到修正的順序

1. **直接幾何路徑**：原輪廓細節烘在表面；切割會移除 identity-critical 細節，且缺少真正的
   遮擋拓撲。退回 representation inspection。
2. **未裁切局部圖層**：可動前景越過應遮住它的邊界。將原畫開口提升為 occlusion mask。
3. **亮度式分類**：陰影表面與目標內部亮度重疊，遮罩外溢。先取樣色彩分佈，改用可分離的
   色度訊號。
4. **沿射線補底圖**：產生放射狀接縫。改用局部擴散。
5. **擴散邊界受暗輪廓污染**：底圖被平均成灰色。只有確認屬於亮底圖的像素可作固定邊界，
   解出的場只回填原 foreground 區，保留原輪廓。
6. **各 UV island 沿自己的影像軸移動**：兩個特徵在表面空間方向不一致。由 UV-to-surface
   tangents 解共同 world/local direction。
7. **各自使用最大位移**：不對稱開口使成對特徵在極限位置發散。共同採用較緊的上限，並將
   日常 UI 範圍與壓力測試範圍分開。

## Workflow 回寫

- `workflow/3d-character-production/facial-expression.md`：
  新增 Texture-baked feature articulation 的七步執行順序與 consumer-local owner 邊界。
- `feedback/history/3d-character-production/common/2026-09-03_151300-texture-baked-feature-articulation.md`：
  保存 generalized lesson。
- `validation/scenarios/3d-character-production/texture-baked-feature-articulation-v1.yaml`：
  覆蓋未裁切 overlay、未校準分類、污染補圖邊界、per-island 發散與能力誤宣稱。

## Claim boundary

此結果證明特定 consumer path 能正確消費 surface-driven articulation；不代表底層交換資產
已新增獨立眼球、眼皮拓撲或 gaze bones，也不改變其他 blocking gates。
