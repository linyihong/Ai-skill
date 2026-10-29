# 近期回饋抽象化覆核（2026-10-26）

範圍：以檔名日期選取 `feedback/history/` 中 2026-10-16 至 2026-10-26 的全部 12 條。這是本次日期範圍的完整覆核，不是所有歷史或按最近修改時間的全庫審查。基準 commit：`71a02433`。

## 發現與原因

- 8 條缺 `Abstraction Check`；8 條有錯誤的相對連結，影響未來 agent 找到規則或方法。來源模板的分類層級相對路徑也少一層，已同步修正模板；不只修產物。
- 部分既有 PASS 仍保留專案操作語意：章節號、固定腳本、owner 條件、資料窗口、延遲重寫，未完整表達輸入契約與反例。
- 1 條 `validated` 的 Reuse Evidence 只記使用者糾正與同專案修正，未提供獨立重用證據，改回 candidate。
- 部分提升目標指回 history 自身，或方法已寫入 analysis 卻未同步狀態與入口紀錄。

根因是把去敏與模板完整度當成可攜性；把「一次修法有效」當成通用必要步驟；把方法接線與獨立重用實測混在一起。既有 KGE 欄位檢查無法證明正文的語意抽象化。

## 逐條處置

| 日期／來源 | 原問題 | 本輪處置 |
| --- | --- | --- |
| [10/16 參數型別](../history/apk-analysis/unity-il2cpp/2026-10-16_173200-il2cpp-paytable-getmain-takes-system-char-not-string.md) | 賠付表與單一 getter 綁定；缺抽象化檢查 | 抽為簽名／表示／呼叫入口驗證，加入 char code unit 與正反對照邊界 |
| [10/16 面板裁切](../history/image-processing/common/2026-10-16_092807-dense-diagram-panel-crop-before-ocr.md) | 方法可攜，但缺 closure 證明 | 保留方法，補可攜性推演與數值／閾值非通用常數的邊界 |
| [10/16 間隙切割](../history/image-processing/common/2026-10-16_132758-regular-panel-grid-mid-gutter-cuts.md) | 方法可攜，但缺 closure；零交集判準需有雜訊邊界 | 補抽象化檢查與容差／人工核對條件 |
| [10/20 宣稱與證據](../history/development-guidance/common/2026-10-20_150000-claim-vs-proof-audit-before-a-protection-wording-switch.md) | 原事件敘述過重、英文正文、缺抽象化檢查 | 抽為逐句宣稱→控制→證據→限制與角色反例核對 |
| [10/22 部署讀回](../history/development-guidance/common/2026-10-22_093000-deployment-state-is-readback-not-narrative.md) | 把特定帳本／tag／草稿鎖設計寫成通用三道閘 | 只保留來源／讀回／unknown 與衍生紀錄原則；實作由專案政策決定 |
| [10/23 粒子資產](../history/apk-analysis/unity-il2cpp/2026-10-23_101000-unity-particle-restore-needs-unitypy-or-assetripper-not-il2cppdumper.md) | 固定物件名與線性成熟度暗示；缺 closure | 依證據目標選工具，runtime 匹配與 Editor replay 分開驗證 |
| [10/23 當前 checkpoint](../history/apk-analysis/unity-il2cpp/2026-10-23_102500-slot-iface-analysis-seated-unitypy-batch-and-commit-cadence.md) | 專案章節、產物、層數與提交頻率混入方法 | 抽為 checkpoint 優先、窗口分離與未觀測缺口；同步清理 analysis 中殘留 |
| [10/23 方法／報告邊界](../history/apk-analysis/unity-il2cpp/2026-10-23_103000-slot-iface-no-oneshot-generator-use-companion-scripts.md) | 邊界條目仍含專案腳本與參考包名；重用證據不足 | 抽為方法／實作責任分離，validated 降為 candidate |
| [10/23 回放結構](../history/apk-analysis/local-proxy/2026-10-23_134500-wire-result-glass-must-match-processresult-strip.md) | 自評 PASS 但強依賴 wire 欄位、固定窗口與補值 | 抽為展示模型／完整執行結構對齊；無合法值時拒絕，不猜值 |
| [10/23 hook 階段](../history/apk-analysis/unity-il2cpp/2026-10-23_140500-apply-shim-onleave-ui-must-cover-all-processresult-owners.md) | 自評 PASS 但把 enter／leave owner 集合相同當必要條件 | 目標×階段矩陣；capability 可不同，刻意排除必須有理由 |
| [10/23 記憶體 apply](../history/apk-analysis/unity-il2cpp/2026-10-23_141500-apply-shim-visiblegrid-must-pad-processresult-strip.md) | 自評 PASS 但補值／延遲重寫過度泛化 | 契約優先；保留 sparse 與 sentinel；重寫須先證實覆寫與排序 |
| [10/26 FX driver](../history/apk-analysis/unity-il2cpp/2026-10-26_100652-animator-play-visual-gap-use-feature-fx-api-not-timeline-first.md) | 方法有承接，但提升目標仍指 history | 更新實際 analysis 方法入口；candidate 維持，不補造第二次實測 |

## 方法承接與發現

- [IL2CPP 執行契約驗證](../../analysis/unity/workflows/il2cpp-runtime-contract-validation.md)：參數／ABI、目標×階段矩陣、展示與完整 runtime 資料結構。
- [Unity 資產解析](../../analysis/unity/workflows/unity-asset-parse-flow.md)：checkpoint 與窗口分離；報告格式屬專案。
- [Unity 驅動追蹤](../../analysis/unity/workflows/unity-fx-driver-trace.md)：feature FX 與證據分支；不靠單一 hook 命中宣稱完整因果鏈。
- 原 image-processing 方法已從 execution-flow 連回兩條候選；維持原入口，沒有另建重複方法。
- release 入口增加兩條候選指標；未把部署設計強制升為全庫規則。
- APK analysis route、Unity FX candidate sources、summary、graph、README 與索引同步；RAG 查回方法來源。

## 驗證與限制

驗證目標：12 條結構完整、相對連結可解析、無原專案腳本／章節號／產品名作通用前提，並能由方法入口與 RAG 發現。KGE 與 runtime 的實際執行結果由本次 commit／close-out 記錄；本頁不是未執行檢查的 PASS 宣告。

可攜性推演使用非來源情境：IL2CPP UI 字元查詢、多 component callback、分塊影像／虛擬列表的展示窗口，以及已有 CD 平台的部署讀回。這些是反例與適用範圍的文本推演，沒有新增裝置或部署實測，也不提升任何條目為 validated。

[回饋寫作規則](../feedback-lessons.md) 與[既有失效模式](../../enforcement/failure-patterns/feedback-lesson-without-abstraction-gate.md) 已補上「形式 PASS」與證據混淆檢查；語意可重用性仍須讀正文，不能宣稱完全機械化。
