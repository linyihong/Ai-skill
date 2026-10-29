> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；專案原始證據留在來源專案。

### 2026-10-23 - 從已到達 checkpoint 取得資產證據並分開標記缺失窗口

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本次文字整理與離線檢查不構成第二個專案的實測。

#### One-line Summary

從已驗證的當前畫面取得可用證據，未觀測的入場／過場另標缺口，不把專案交件格式當成分析方法。

#### Human Explanation

為了重拍先前場景而離開已到達的 checkpoint，可能增加導航成本或失去可用窗口。分析應保留不同場景的身分與證據範圍；已載入背景不能替代未捕獲的入場素材。

#### Trigger

- 已在目標功能內，但先前入場／過場尚未捕獲。
- 資產盤點、粒子參數與畫面分層被混成固定交件清單。

#### Evidence

- Tool: 來源專案既有分析／驗證紀錄；本輪未重跑其執行環境。
- Sanitized excerpt: 任何已到達內容頁而缺少入場證據的 Unity 應用，都能採用 checkpoint 優先與窗口分離。
- Evidence path: `<PROJECT_ROOT>/` 的受控分析與驗證紀錄；不複製原始物件、樣本或部署紀錄。

#### Generalized Lesson

1. 驗證目前 checkpoint／目標物件後，先收集此窗口可證明的畫面、資產與 component 關係。
2. 將未觀測窗口標 not-captured-this-session，不借用當前背景當作其來源。
3. 依 material／texture 引用定位功能包與共用包，不假設命名或包位置。
4. 分開記錄參數已讀、runtime 已匹配、Editor 可重播等證據；依專案契約產出報告。
5. 提交粒度遵守各 repo 的交付規則，不把一次專案的章節號、檔名、層數或 push 頻率變成通用義務。

#### Agent Action

先辨識輸入契約與前置條件，再依本條決策；不照搬來源專案的函式名、腳本路徑、固定數字或狀態碼。記錄採用／不適用的理由。

#### Goal / Action / Validation

- Goal: 從已驗證的當前畫面取得可用證據，未觀測的入場／過場另標缺口，不把專案交件格式當成分析方法。
- Action: 依 Generalized Lesson 記錄輸入、決策與反例。
- Validation or reference source: 當前 checkpoint 的物件與資產可對照；未捕獲窗口明確列出；換到另一個 Unity 功能仍不需沿用原報告路徑／層數。

#### Applies When

- 已有穩定 checkpoint，且任務允許分窗口補證據。

#### Does Not Apply When

- 任務明確要求自然完整流程，或缺失入場是必要前置條件；此時仍須補拍，不能以 deferred 宣稱完成。

#### Validation

當前 checkpoint 的物件與資產可對照；未捕獲窗口明確列出；換到另一個 Unity 功能仍不需沿用原報告路徑／層數。

#### Promotion Target

[analysis/unity/workflows/unity-asset-parse-flow.md](../../../../analysis/unity/workflows/unity-asset-parse-flow.md)。

#### Promotion Record

2026-10-26：抽象化整理並接到上述方法／入口；這是候選指引的連結，不宣告 validated／promoted，也不新增獨立重用證據。

#### Required Linked Updates

- 本分類 README 與上述方法／入口已檢查並同步。
- 索引與 RAG 經 runtime refresh 重建；原始專案證據不搬入共用庫。

#### Abstraction Check

- Swap-project test: PASS — 任何已到達內容頁而缺少入場證據的 Unity 應用，都能採用 checkpoint 優先與窗口分離。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

依近期回饋審查，把原先依賴專案術語或操作細節的敘述改為輸入、機制、決策與驗證；修訂前內容可由 Git 歷史追溯。Status 保守維持 candidate。

#### Transfer Contract

- Inputs: 已驗證的當前 checkpoint 與未捕獲窗口
- Mechanism: 回頭重建先前場景可能失去可用取證窗口
- Decision: 先取得當前證據並明列未捕獲範圍
- Verification: 當前物件可對照，未捕獲項未被錯報為完成
- Counterexample: 完整入場是驗收必要條件時仍需補拍
- Transfer case: 互動展示已在內容頁但尚缺入場觀測

#### Intelligence Review

- Decision: existing
- Reason: 當前可用觀測優先與 fallback 取捨已由最高槓桿分析路徑承接。
- Target: intelligence/engineering/analytical-reasoning/highest-leverage-analysis-path.md
