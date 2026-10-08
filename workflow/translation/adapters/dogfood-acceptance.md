# Dogfood 驗收維度與證據綁定

本文件細化既有 Translation Decision 的 adapter 驗收，不新增日文 workflow、
固定譯詞庫或 runtime route。資料責任仍屬於
[`reference-resolution`](../contracts/reference-resolution.yaml)、
[`validation`](../contracts/validation.yaml) 與
[`finality`](../contracts/finality.yaml)。

## 四個責任

| 層面 | 證據與責任 | 禁止的捷徑 |
| --- | --- | --- |
| Reference / Identity | 角色、指涉、關係與來源；既有驗收名稱是候選與一致性證據 | 看字形猜人物、未決身份直接寫固定對照 |
| Semantic | 比較來源與目標的角色、事件、受事、關係、否定、模態、時間、數量與意圖 | 通順、同語系或同一模型自評直接當語意通過 |
| Acceptance | 執行、機械、語意、時間、畫面與 Finality 分欄 | 指令 exit 0 或機械 pass 推成內容／成片 pass |
| Target Lexical | 依 locale 檢查詞彙存在、語意等價與上下文自然性 | 日文含漢字即錯、漢字合法即自然、逐字換形即翻譯 |

Reference Resolution 保留於既有 lifecycle 的正式 Expression Analysis 之前。
可先偵測人物／稱謂候選以觸發解析，但不得把候選偵測當身份已確定。
解析後的實體與關係再供分析與 Candidate Space 使用。
沒有充分證據時保留 unresolved；缺少必要 reference → blocked，
有可追溯但有歧義的 reference → needs_review + 明確原因。

Semantic validation 使用既有 Core F1、F10、F11、F16 等概念，不新建 locale
語意 taxonomy。語序重排、自然改寫或可由語境恢復的省略，不可只靠字詞差異
判失敗；需要來源／目標證據與關係綁定。未分析不等於語意 pass。
依 Translation Context／敘事重要性宣告本次必要 semantic assertions，
關鍵 dialogue／劇情事件可驗人物、動作、受事、否定與 speech-act；
其他內容不強迫填滿重型結構。必須標示 coverage／未驗維度，
只驗部分 assertions 的 pass 不得宣稱完整 semantic acceptance。

ja-JP shared-script 是 F18／JA-F12 的 lexical/locale constraint；
合法漢字姓名可能成立，不得要求所有外國姓名一律片假名。
script、lexical、identity 與 semantic 結果必須能分別追溯。

## 驗收紀錄

Adapter 每筆結果至少保留 source identity、target locale、候選 identity、
producer／版本、reviewer／證據引用，以及以下獨立維度：

- execution：completed／failed／timeout／running，不代替品質。
- mechanical：非空、文字系統與既有硬限制；缺輸出不得 pass。
- reference、lexical、semantic：pass／fail／review／not_evaluated；各有原因。
- temporal、visual：由 consumer 擁有；未實測記 not_evaluated，不自動 pass。
- finality：依 canonical PASS／REVIEW／BLOCK；必須列出阻塞／覆核原因。

接受 review evidence 前，核對 source、candidate、locale 與 evidence provenance。
舊候選的覆核不得套在新候選；Selection Actor 的自評不得自關必要 gates。
摘要可以投影既有欄位，但不得用另一套結果 schema 改寫 canonical 契約。
尚缺必要驗證 → blocked + missing_required_validation；並非無理由一律 REVIEW。
翻譯內容驗收不代替 NVP timing/layout gates；成片驗收需 consumer 的證據。
目前作為 Phase 3 dogfood acceptance surface；跑過代表影片後才評估是否
升格 canonical schema。未知不是錯誤，但不得投影成已決定。

## 重跑與驗證

先保存 baseline，再針對已知 failure 改變修復條件；沒有新證據或新條件時
不得把相同 deterministic retry 視為修復。保留完整失敗候選與原因。
中途報告只能是暫定；正式摘要須讀取 terminal child exit 與完整 artifact。

必備負例：空字串機械假陽性、只有人名的殘片、缺 reference、動作／受事／
否定／模態遺失、無來源的語意擴張、舊候選覆核錯用、日文合法漢字姓名、
shared-script lexical 未決，以及內容通過但時間／畫面未驗。
Fixtures 驗的是 evidence consumer，不證明自然語言分析 producer 已可靠；
兩者與真實語料、ASS、burn 的驗收分開。
