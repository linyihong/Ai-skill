> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-26 - Lab presence fixture：FX 列表 API 空陣列 ≠ 路徑未跑；force overlay 不能當自然閉環

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源為 Unity IL2CPP slots lab：hook 到 feature「取掉落／消失符號列表」API 在 trigger `ProcessResult` 已進入，但回傳長度 0；同輪 follow-up settle 經 `EventSink`／`SPIN` apply，未再見該 API。專案用 force `Setup(符號[])` 可出 overlay。具體柜／fixture／截圖留專案 docs。

#### One-line Summary

當 lab 用 **presence-only** 盤面／獎勵（無「將消失符號」等元資料）重放時，feature FX **列表查詢 API 可能已呼叫但回傳空**；此時 force `Setup(假列表)+Show` 只能鎖 overlay 驅動，**不能**宣稱自然列表路徑已閉環。另須分開記錄 trigger `RESULT` 與 settle `SPIN`／`EventSink` 的 apply 擁有者。

#### Human Explanation

Slots／Unity 柜常見兩段：觸發 feature（落地／通知）→ free SendNext → settle 盤。掉落／消失 FX 往往先呼叫 `Get*Symbols`／`GetDrop*` 之類方法取得 `string[]`（或等價列表），再 `Frame*Drop.Setup(list)+Show`。

Lab fixture 若只補可見格與必要 reward 標籤（presence），沒有伺服器／狀態機寫入的「消失集合」，列表 API 仍會被自然代碼路徑呼叫，但回傳 `n=0` → 空 `Setup` → 弱或無 overlay。Agent 容易誤判「hook 沒掛上」或「自然路徑未跑」。相反：force 注入假列表證明 **Show／Setup 驅動** 可用，卻把「自然列表從哪來」留開。

Follow-up settle 還可能只打到 `EventSink_SlotsRound`／`ProcessSpin`（Symbols 晚到），而列表 API 只在某一 `ProcessResult` 窗口跑——必須用 hook 事件表對齊相位，不能假設第二手一定再呼叫一次。

#### Trigger

- Hook 顯示 `GetDrop*`／`Get*Symbols` 已進入，但 `n=0`／空陣列
- force `Setup(a..e)+Show` 有 overlay，關掉 force 後無視覺
- feature 鏈：trigger RESULT 有列表呼叫；settle 僅 SPIN／EventSink apply，無第二次列表事件
- 想把「lab overlay 已鎖」寫成「自然掉落路徑完成」

#### Evidence

- Tool: Frida apply-shim Interceptor on list API + Drop `Setup`／`Show`；lab fixture with／without force list
- Sanitized excerpt: trigger `ProcessResult` → list API `n=0`；force Setup 非空 → overlay；settle apply tag=`EventSink`/`SPIN` 無第二次 list 事件
- Evidence path: `<PROJECT_ROOT>/docs/slots/*/lab/` natural-hunt notes（專案 incident）

#### Generalized Lesson

1. **分三層成熟度**：① hook 證實 API 被呼叫；② 列表非空（自然或注入）；③ overlay／粒子可見。①+③ via force ≠ ② 自然閉環。
2. **Presence fixture 預期空列表**：無 vanish／drop 元資料時，`Get*Symbols` 回空是常態，不是 hook 失敗。
3. **Force overlay 標為 lab-matched**：文件寫「驅動已鎖；自然列表仍開放」，勿升格為 live-matched natural。
4. **對齊 apply 相位**：記錄每次 `LAB_APPLIED`／等同事件的 `tag`（`ProcessResult` vs `EventSink`/`SPIN`）與 fixtureId；列表 API 可能只在其中一相。
5. **下一最小步驟**：生產 settle 自然命中、或 IL／dump 找列表填充來源再注入等價狀態——不要只加大 force 符號表假裝閉環。

#### Agent Action

- 見空列表：先查 fixture 是否 presence-only、settle 是否走到預期 ProcessResult。
- 報告分開：force overlay 證據 vs 自然 `n>0` 證據。
- 不要因為 force 截圖漂亮就關閉「自然列表」next-window。

#### Goal / Action / Validation

- Goal: 正確判定掉落／消失 FX 自然路徑是否閉環
- Action: hook 列表 API＋Drop Setup；對比 force／non-force；記錄 RESULT vs SPIN apply
- Validation or reference source: 事件表 `n`／apply tag＋截圖；自然閉環需 `n>0` 非 force

#### Applies When

- Unity IL2CPP lab replay（Frida apply-shim／fixture）獵 feature 掉落／消失／列表驅動 FX
- 有獨立 `Get*Symbols`／`Setup(string[])`（或等價）API

#### Does Not Apply When

- 掉落純粒子／Timeline，無列表 API
- 已有生產 wire／狀態明確攜帶消失集合且 desk-match 列表非空

#### Validation

- 同裝置：force 非空 Setup → overlay；non-force → list `n=0` 且弱／無 overlay
- settle 日誌無第二次 list 呼叫時，不得宣稱 settle 相位已驗證自然列表

#### Promotion Target

- [Unity 動畫／特效驅動追蹤](../../../../analysis/apk/workflows/unity-fx-driver-trace.md)（§ presence／相位）

#### Promotion Record

尚未 promotion；保留為 candidate history。本輪已把方法要點接入上述 analysis 流程草稿連結。

#### Required Linked Updates

- category README 追加本條
- `analysis/apk/workflows/unity-fx-driver-trace.md` 補 presence／RESULT–SPIN 相位注意
- 專案 evidence 留 `<PROJECT_ROOT>/docs/slots/*/lab/`（已寫 natural-hunt note）

#### Abstraction Check

- Swap-project test: PASS — 任何 Unity lab 用 presence fixture 獵列表驅動 FX 時，空回傳與 force overlay 的成熟度分層均成立
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/*/lab/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### Transfer Contract

- Inputs: presence fixture、列表 API hook、optional force Setup
- Mechanism: 無消失元資料 → 列表空；force 只證明 Show 鏈
- Decision: 分層記錄成熟度；自然閉環要非空列表或 IL 填表
- Verification: 對比 force／non-force 事件 `n` 與畫面；核對 apply tag 相位
- Counterexample: wire／狀態已帶消失集合且自然 `n>0`
- Transfer case: 其他需「查表再 Setup」的 lab feature（不只掉落）

#### Intelligence Review

- Decision: method-only
- Reason: 補強既有 FX 驅動追蹤流程的 lab／相位邊界，不另抽新 heuristic 檔
- Target: analysis/apk/workflows/unity-fx-driver-trace.md
- Revisit: none
- Discovery: not applicable — method-only
