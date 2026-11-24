# Phase 2 — 兩個 consumer 的唯讀評估（2026-11-23）

> **唯讀**：只讀取 consumer 的 HEAD，用 [`workflow/cross-cutting/project-tooling/`](../../../../workflow/cross-cutting/project-tooling/README.md) 的契約判斷；**沒有修改任何 consumer repo**（評估前後 `git status` 由本 phase 造成的變更為 0）。
> 具體檔名與可重跑指令在各 consumer 的 project-local `.ai-skill/dogfood/` 對照檔；本檔只留去敏後的判斷。Consumer A = 逆向分析專案（Python + Frida）、Consumer B = 業務交付 monorepo（bash + PowerShell + Node）。

## 1. Role 判定（違規句型）

| # | Consumer | 違規句型 | 實例數 | 可判定 |
| --- | --- | --- | --- | --- |
| R1 | A | Target 的 Logic 放在 Entry | 2（各 132 / 260 行，自帶 tap、截圖、OCR 流程） | 是 |
| R2 | A | Target 的 Parameters 放在 Entry | 3（入口寫死裝置序號預設值；2 支另有模組層級點擊座標） | 是 |
| R3 | A | Shared 的 Parameters 放在 Logic | 1（共用裝置模組寫死序號預設值） | 是 |
| R4 | A | Domain 的 Parameters 重複了 Shared 已有的 Parameters | 1 組（大廳模組與共用場景錨點各自定義相同的 2 個座標常數） | 是 |
| R5 | B | Domain 的 Logic 放在 Entry（且 sh / ps1 各實作一份） | 2 組（317 / 130 行入口，幾乎不呼叫 lib，自帶文字處理邏輯） | 是 |
| R6 | B | Domain 的 orchestration 放在 Entry？ | 數組 > 160 行入口（含快取、scope 判斷、依序呼叫 10–22 個 gate） | **歧義**（見 §5 A-2） |

## 2. Concept / Discovery

- **同概念群**（沿用 Phase 0 試判）：A 的「旋轉直到事件並擷取」10 支中 8 支已是「共用 helper + 單目標 sibling」；2 支（= R1）未用共用 helper。B 的「效能 lane」8 支 0 共用 import。
- **同概念的鎖**（新發現）：A 有 3 把「文件 / 資產存在 → 必須接到平台 UI」的 completeness 鎖，B 有 6 支 `frontend_integration_*` 檢查——growth rule 的 reuse evidence 同樣適用於鎖（見 §5 G6）。
- **Creation justification**：A 為 0；B 的 53 個入口中 50 個有一行用途註解，其中約半數附 plan / docs 連結（說明「做什麼」，未必說明「為什麼不用既有工具」）。
- **Capability catalog**：兩者都有 README 表格，但以**目錄**或**入口**為鍵（A：分類目錄表、入口 → 模組對照表；B：sh / ps1 腳本對照表），不是以 capability 為鍵，也沒有 state / owner / ttl 欄位。Discovery 目前只能靠讀 README 與 grep。

## 3. Placement × State × Ownership

| 項目 | 判定 |
| --- | --- |
| Ownership | 兩者皆無 owner 紀錄 → 所有 maintained 工具都是 **ownership violation**（state 不變，依契約補 owner 即可） |
| A 實驗探針區 | Research / experimental；以「未被引用 + 超過 N 天」**自動刪除**——與契約「TTL 是期限、不是自動刪除指令」有張力（見 §5 G4） |
| B 驗證延期紀錄 | 有 `expires_at`、到期即恢復完整驗證、需另一 session 寫 resolution evidence——**正例**：TTL 到期觸發重新評估，而不是默默續用或刪除 |
| B 根目錄 SQL 修復腳本 | Phase 0 疑為 Scratch 殘留 → **refuted**：被決策文件與整合文件引用，屬 Target / maintained；問題只是放在工具樹根目錄、無分類、無 owner |
| 殘留 scratch | 兩者工具樹內未發現；A 根目錄有一個不進版控的暫存目錄，符合契約位置規則 |

## 4. Lock review

| Consumer | 鎖數 | layout | thin-entry | boundary | lifecycle | completeness | drift | parity | **未涵蓋** |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| A | 24 | 6 | 0 | 2 | 1 | 10 | 2 | 1 | 2（機密 / 外洩防護） |
| B | 53 入口 / 82 lib check | ~8 | 0 | ~10 | 1（驗證延期） | ~25 | ~10 | 1 | **~25 產品程式慣例檢查**（主題、時間格式、i18n、DI 慣例、測試穩定性…） |

Compensating-lock Review Rule 逐案結論（節選）：

| 鎖 | 能否從結構消除？ | 結論 |
| --- | --- | --- |
| A：hook 副本同步鎖 | 不能——hook 必須自含、不得載入 tracked 工具（刻意的隔離政策） | **合理的 compensating lock**：原因已記錄，符合「記下原因再加鎖」 |
| A：device I/O boundary 鎖 | 部分能——R1–R3 改走共用層後違規消失 | **scope 盲區**（不掃入口層）＋ **pattern 漏洞**（抓不到「讀環境變數帶寫死預設值」與「`X, Y = 數字, 數字`」寫法） |
| B：sh / ps1 成對鎖 | 能——26 組薄外殼改為單一參數化入口後，大部分成對需求消失 | 典型 compensating lock（Q4 另案） |

兩者都**沒有 thin-entry 鎖**；A 的入口薄是慣例，B 的入口有 R5。

## 5. 契約缺口與歧義（附建議修改，**未套用**）

| # | 類型 | 發現 | 建議的契約修改 |
| --- | --- | --- | --- |
| G1 | 缺口 | 七族無法容納「機密 / 外洩 / 去敏」類鎖（A 2 把） | `lock-taxonomy.md` 新增 **safety** 族 |
| G2 | 缺口 | B 約 25 個、A 約 10 個鎖保護的是**產品程式或產品文件**的不變式，不是工具治理 | `lock-taxonomy.md` 明訂範圍：七族用於工具治理；產品不變式鎖歸「product invariant」類，不套族形狀，但**仍適用** compensating-lock Review Rule |
| G3 | 缺口 | 鎖的 forbidden pattern 只抓單一寫法，同一違規換個語法就漏 | 鎖形狀的 `forbidden` 欄補一句：須涵蓋同一違規的常見寫法，並附至少一個反向樣本測試 |
| G4 | 歧義 | 「未引用 + 逾期即自動刪除」算不算合規的 TTL 到期處置？ | `lifecycle-contract.md`：允許**未被引用、未登記 catalog** 的 Research 工具機械刪除（視為 deprecated → 移除）；被引用或已登記者到期必須明確評估 |
| G5 | 歧義 | 「用途註解 + plan / 決策連結」能否算 creation justification？ | `discovery-contract.md`：連結的文件若說明為何新增，可視為等效 justification；只有用途、沒有理由者不算 |
| G6 | 缺口 | 同概念的鎖也在重複長（A 3 把、B 6 支） | `lock-taxonomy.md` 引用 `growth-rule.md`：reuse evidence 同樣適用於鎖 |
| A-1 | 歧義 | 單一目標的**分析用畫面幾何**（OCR 裁切框、像素判斷點）算 Parameters 還是 Logic？ | `role-contract.md`：**操作幾何**（點擊座標）必須是 Parameters；**分析幾何**可留在 Target 的 Logic，但須具名、集中、單一用途 |
| A-2 | 歧義 | 入口依序呼叫多個 gate 並含快取 / scope 判斷，是 Entry 還是 Logic？ | `role-contract.md`：純依序呼叫屬 Entry；含條件、快取、範圍決策的編排屬 Logic |

## 6. 建議改動清單（不執行）

**Consumer A**

1. R1 / R2：2 支 hunt 工具的流程移到單目標區，改走共用裝置模組；點擊座標移到該目標設定檔的 controls；入口只留呼叫。
2. R2：另一個入口拿掉寫死序號與多餘的 PATH 調整（目標邏輯已會自行找工具）。
3. R3：共用裝置模組改用序號解析函式，不帶寫死預設值。
4. R4：大廳模組改從共用場景錨點取座標。
5. device I/O 鎖：scope 擴到入口層，pattern 補上 G3 的兩種寫法，加反向測試。
6. 新增 thin-entry 鎖。
7. 3 把平台接線 completeness 鎖：依 reuse evidence 第 3 次，裁決抽共用檢查或記下保留理由。
8. 既有 README 表改為以 capability 為鍵，補 state / owner。

**Consumer B**

1. 效能 lane：抽共用 harness，per-case 只留宣告（原 Phase 3 構想）。
2. 26 組薄外殼：單一參數化入口（Q4 另案）。
3. R5：2 組入口的邏輯移到 lib，sh / ps1 只留呼叫。
4. 根目錄 SQL 修復腳本：移入分類並登記 catalog、owner。
5. `lib/` 150 檔依角色分層（check / test / helper）。
6. 6 支整合測試結構檢查：依 reuse evidence 評估共用檢查引擎。
7. 腳本對照表改為以 capability 為鍵，補 owner。

## 7. 判定結果摘要

- **可判定性**：Role、Concept、Placement × State、Ownership 在兩個 consumer 都能下結論；只有 A-1、A-2 兩處歧義。
- **契約需修改**：G1–G6 共 6 處（2 處屬歧義釐清、4 處屬範圍或形狀補強）；**沒有**需要推翻的 invariant（三維度分離、concept-first、Review Rule 都成立）。
- **正例**：A 的單目標薄 preset、A 的 research promotion 表、B 的驗證延期紀錄、B 的共用 manifest（Parameters 集中）。
