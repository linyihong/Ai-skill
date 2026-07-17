# 2x — Consumer player variant matrix：同型 fixture 假綠與 A-fix-B-break（2026-07-17）

> **專案證據邊界**：片源 URL、episode id、host、commit、測試檔名與 release 結果留在 `<PROJECT_ROOT>`；本檔只保留 generalized failure shape、量測與契約回饋。

## Run 摘要

- **任務**：修復 H5 player 黑屏，並以瀏覽器整合測試證明可播放。
- **模式**：主 session 以 Mode A 連續診斷、實作、自驗、部署；fresh Executor / Verifier 均為 **0**。
- **結果**：先以單一未加密直向 fixture 建立 oracle，後續使用者實測才暴露加密 HLS 與橫向電影回歸。補成「方向 × 裝置 × media path」矩陣後，桌面／手機的橫向與直向路徑才同時成立。
- **定位**：負向 evidence only；不修改 ERA working model，不升格 schema / workflow。

## 假綠機制

| 假綠來源 | 為何無法關閉 slice |
|---|---|
| 單一 fixture 通過 | fixture 的 codec、encryption、aspect ratio 與失敗案例不等價 |
| 封面可見或時間前進 | 只能證明 UI / audio / timeline 活著，不能證明 video frame 已 decode |
| unsupported banner 出現 | fallback UI 正常，不等於「使用者要求的可播放路徑」已恢復 |
| 桌面單一路徑通過 | mobile UA、native HLS、MSE/hls.js 與 fullscreen 分支仍可能不同 |
| 只重跑剛新增的測試 | 實作者選定的量尺可能把相鄰 variant 排除在 acceptance 外 |
| deploy smoke 通過 | generic smoke 未必包含 media decode、key retrieval 或 orientation journey |

## Root-cause 差集

本 run 不是單一 bug，而是測試量尺逐輪才補齊：

1. **Frame oracle 缺失**：idle cover / poster 被當成非黑屏證據；真正的 decode oracle 應要求 `videoWidth > 0 && videoHeight > 0`。
2. **Media-path 等價假設錯誤**：未加密 HLS、AES-key HLS、browser-native HLS 與 MSE/hls.js 不是同一路徑。
3. **Capability probe 過寬**：browser 回報「可能支援」不等於 native HLS 可播放；需要以實際 decode 路徑或更嚴格 capability 判準決定 fallback。
4. **Orthogonal regression 未列入 backfill**：修直向時未把既有橫向影片列為 `verifier_only` regression，造成 A-fix-B-break。
5. **環境約束未顯式化**：key endpoint 的 rate limit / runtime config 可讓 encrypted fixture 黑屏，但 UI 層測試不會指出原因。

## 契約回饋候選

### 1. Acceptance-equivalence gate

Orchestrator 在選 fixture 前應回答：fixture 是否與 user-reported failure 在下列維度等價？

| 維度 | 例 |
|---|---|
| codec family | broadly supported vs capability-gated |
| encryption | clear HLS vs key-backed HLS |
| aspect / layout | portrait vs landscape |
| device mode | desktop vs mobile UA / viewport |
| playback engine | native HLS vs MSE/hls.js |
| user-visible oracle | cover/banner/timeline vs decoded frame |

若不等價，fixture 只能產生局部 evidence，不得代表整個 player slice。

### 2. Variant matrix before execution

User-visible player change 的 verification backfill 應至少包含：

- primary failing variant（executor happy path）；
- 一個與修改軸正交的既有 variant（`verifier_only` regression）；
- 適用時的 desktop + mobile runtime；
- decoded-frame oracle，而非只驗 DOM、poster、audio 或 fallback banner。

矩陣不是要求全排列；orchestrator 以「本次修改可能切換哪些 runtime branches」挑最小區分集。

### 3. Verifier 必須挑戰 fixture 等價性

L3 adversarial 不只問「新測試有沒有過」，還要問：

1. 測試 fixture 是否真的走 user-reported path？
2. oracle 是否證明 user-visible capability，而非替代狀態？
3. 是否有一個相鄰 branch 可反證 A-fix-B-break？
4. generic deploy smoke 是否缺少本 slice 的 runtime authority？

## 量測欄

| 指標 | 值 |
|---|---|
| Executor fresh context | **0** |
| Verifier fresh context | **0** |
| stakeholder 重新定位 root cause | **≥2** |
| 初始 authority | 單 fixture、局部 decode / fallback |
| 最終最小矩陣 | portrait + landscape × desktop + mobile（4 格） |
| 額外 encrypted-path probe | key response + decoded-frame browser journey |
| 最終 decode oracle | `videoWidth > 0 && videoHeight > 0` |
| 模型自然落位 | **是** — missing constraint / partial authority；ERA 不需修改 |

## 對 Phase 3 的意義

- **Predictive stability：支持**。現有模型可直接解釋：implementation evidence 沒有約束所有 runtime branches，closure 發生在過大的 feasible set。
- **Q5：不新增 schema promotion 壓力**。consumer 可先以 backfill matrix + fresh Verifier 補洞。
- **Q8：增強**「Evidence Producer ≠ Closure Authority」與 constraint responsibility；同一 session 選 fixture、寫測試、判 pass，正是自證循環。
- **不視為 Phase 3 closure**；只作新案例無需修改模型即落位的 observation。
