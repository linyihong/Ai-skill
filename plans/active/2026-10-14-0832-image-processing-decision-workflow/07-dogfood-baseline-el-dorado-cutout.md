# 07 — Dogfood Baseline：El Dorado Dice Cutout

← [`_plan.md`](_plan.md)

## 來源（消費者專案）

- 專案：`pojiepokerist`／cabinet `el-dorado-dices`
- 任務表象：賠付表彩色頭像去背、置中、對齊 Unity 灰階 Sprite
- 執行環境：Cursor agent + 本機 Python（`rembg`、Pillow、NumPy、scipy）
- 產出：`paytable-symbols/high_*.png` 等（業務 repo；本檔只留抽象教訓）

本檔是 **Phase 0 問題基線**，不是「已通過的去背工作流證據」。

## 契約若當時寫清楚，應長怎樣

| 欄位 | 實際需要 |
| --- | --- |
| `task_class` | A（去背）+ D（置中／縮放到固定 canvas） |
| `background` | transparent |
| `immutable` | 臉／牙／羽／寶石細節；禁止生成式重畫 |
| `consumer_profile` | `slot-symbol`：328²、face-axis、可選 Sprite silhouette |
| `fidelity` | strict（前景） |

當時缺失：沒有獨立 verify；沒有 failure_class；repair 靠同一腳本加參數。

## 觀察到的失敗 → 分類

| 觀察 | failure_class | 層級 |
| --- | --- | --- |
| rembg 輸出即當完成，仍有 navy／green halo | `edge_halo` | L3 |
| 棋盤格預覽當透明證據 | 驗收方法錯（非 class） | L3 |
| binary opening 非金色區 → 牙齒透明洞 | `interior_damage` | L4 |
| 牙齒黑塊被錯誤 inpaint 成綠矩形 | `interior_damage`／`fidelity_violation` | L4 |
| 羽冠被 paytable UI 裁平仍硬修 | 契約未區分「源圖已裁切」 | L1 |
| 全案同一 rembg→硬 alpha→rim 剝離 | 無候選路由 | L2 |
| 變量名遮蔽致後續圖失敗（工程事故） | `tool_error` | 執行工程 |
| UI 白線／殘字當主體邊緣 | `background_residue`／分析漏 `text_overlay` | L1/L2 |

## 有效做法（可升為候選，不當唯一真理）

| 做法 | 候選角色 |
| --- | --- |
| rembg 作初始 matte | `seg.generic` baseline |
| 僅剝 rim 暗邊／色暈，不動內部 | `post.decontaminate`／受限 rim fix |
| 最大連通域、丟小浮島 | 確定性 post（有面積上限） |
| Sprite silhouette 作 **consumer** 硬裁 | profile constraint，非通用核心 |
| 白／黑／棋盤預覽 + 灰階並排 | 人審輔助；機械仍查 alpha |

## 明確反模式（應進 registry 候選）

1. **執行成功即 PASS**
2. **對主體內部做 morphology opening「清乾淨」**
3. **暗像素一律當背景填掉**（破壞口腔／陰影）
4. **無界重跑同一閾值**
5. **用生成式整圖重畫取代去背**
6. **把業務對齊（臉軸／Sprite）寫死成通用去背步驟**

## 建議納入 Phase 0 測試集的抽象案例

| 案例 id | 特徵 | 期望壓力 |
| --- | --- | --- |
| `illust-hard-edge` | 插畫硬邊、高飽和 | 基礎 A |
| `illust-glow-halo` | 自帶光暈底 | `edge_halo` |
| `illust-interior-holes` | 牙縫／孔洞應保留 | `interior_damage` |
| `ui-overlay-text` | 圖上有 UI 線／字 | `text_overlay` |
| `soft-fringe` | 柔邊／半透明 | B 類預告 |

（實際像素另存；注意去敏與授權。）

## 對本 plan 的含義

El Dorado 經驗證明：**缺的是決策與驗收閉環，不是多一個 rembg 包裝 Skill**。P0 垂直切片必須能擋住「有 PNG 但 halo／interior_damage」並迫使換策略或人審。
