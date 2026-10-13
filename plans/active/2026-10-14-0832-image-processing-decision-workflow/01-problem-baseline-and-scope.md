# 01 — Problem Baseline & Scope

← [`_plan.md`](_plan.md)

## 要解決的四層失效

| ID | 失效 | 症狀 | 錯誤完成訊號 |
| --- | --- | --- | --- |
| L1 | 任務判斷錯誤 | 去背做成生成新圖、裁掉主體、改細節 | 「看起來有圖了」 |
| L2 | 工具選擇無規則 | 硬邊商品與髮絲／插畫走同一 pipeline | 「用了最強模型」 |
| L3 | 缺少結果驗收 | 有 PNG／棋盤格就當成功 | exit code 0 |
| L4 | 失敗無修復策略 | 同一參數重跑；越修越壞 | 「再試一次 rembg」 |

**核心原則**：驗收標準是 **輸出符合任務契約**，不是 **圖片處理流程跑完**。

## 任務類型註冊（建置優先序）

| Class | 名稱 | Priority | 驗收焦點 |
| --- | --- | --- | --- |
| A | 背景移除與透明化 | **P0** | 主體完整、背景真透明、無明顯殘留／光暈 |
| B | 精細遮罩與局部修復 | P1 | 髮絲／半透明細節、邊緣自然、不誤傷主體 |
| C | 修復與內容編輯 | P2 | 修改區正確；未授權區保真 |
| D | 確定性像素處理 | P1 | 尺寸／格式／座標／色彩空間符合規格 |

優先序是 **工作流建置順序**，不是技術難度排序。P0 未穩定前不大擴 C。

## 範圍邊界

**In（P0）**

- 通用去背決策契約與最小垂直切片
- Independent verification（機械 + 人審／inconclusive）
- Failure taxonomy + 有限 repair
- Dogfood：含 slot 符號類合成／插畫案例（當 consumer 壓力測試）

**Out（P0）**

- 影片全流程、字幕、分鏡
- 商業 SaaS 自動去背產品化
- 強制 ComfyUI／特定雲端 API
- 把「對齊灰階 Sprite／臉軸」寫進通用核心（見 Q2：consumer constraint）

## Phase 0 盤點清單（待填）

| 項目 | 現況（草案） | 待補 |
| --- | --- | --- |
| 本機工具 | Pillow、NumPy、scipy、rembg（agent 臨時腳本） | 授權／版本／入口 CLI |
| Ai-skill Skill | 無正式 image-processing skill | 不先建 Skill；先契約 |
| Prompt | 無標準 | 禁止以 Prompt 代替契約 |
| 失敗案例 | El Dorado cutout（[`07`](07-dogfood-baseline-el-dorado-cutout.md)） | 再收 3–5 張去敏代表圖 |
| 統一入口 | **無**（混合） | Q1 |

## 與現有治理的對照

| 既有原則 | 本工作流落地 |
| --- | --- |
| Loop-first | Analysis → Selection → Execute → Verify → Repair → Closeout |
| Constraint vs Selection | Contract 裁可行集；Policy 選候選 |
| Runtime state machine | 處理狀態、重試上限、human_review、abort |
| Canonical state | IDR／verify report 為權威；文件鏡像不代替 |
| Independent verification | 執行器 ≠ 驗收器 |
| Failure learning → guards | 僅經覆核案例可升 registry |
| Lazy loading | 只載入當前 task class 契約 |

## Gate（Phase 0）

- [ ] A1–A3 假設被接受或改寫
- [ ] 至少 5 個失敗可歸到 L1–L4／failure_class
- [ ] 測試集規劃就位（路徑 + 案例表；二進位可延後）
