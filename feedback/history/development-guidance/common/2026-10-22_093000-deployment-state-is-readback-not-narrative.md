> 遵守 [共用規則索引](../../../../enforcement/README.md) 與 [回饋規則](../../../feedback-lessons.md)；專案原始證據留在來源專案。

### 2026-10-22 - 部署現況以讀回事實與來源為準

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本次文字整理與離線檢查不構成第二個專案的實測。

#### One-line Summary

部署現況應可追溯到各元件的讀回證據；操作者陳述、歷史文件與 git tag 分別標示來源。

#### Human Explanation

交接敘述可能隨局部部署或中斷而過期。版本、migration 與 runtime 也可能不同步，因此不能只看 tag 或單一健康端點推定整個環境一致。

#### Trigger

- 多元件或多人交接部署，現況只能由聊天／歷史文件推測。
- 變更需要操作者處理，但交付物沒有通知或讀回位置。

#### Evidence

- Tool: 來源專案既有分析／驗證紀錄；本輪未重跑其執行環境。
- Sanitized excerpt: 自動 CD 與人工部署都能採用「讀回事實、來源與未知分開」的原則，不需共用帳本格式或腳本語言。
- Evidence path: `<PROJECT_ROOT>/` 的受控分析與驗證紀錄；不複製原始物件、樣本或部署紀錄。

#### Generalized Lesson

1. 先找既有 CD／部署平台的現況權威；已有可查平台時不另建競爭帳本。
2. 按元件記錄已部署 artifact、必要 schema／runtime 版本、來源與 verified_at；未知保留 unknown。健康端點只能證明它實際讀回的範圍。
3. 在變更產生處留下必要操作、依賴與復原通知；部署入口依專案政策驗證，部署後自動讀回優先。
4. 若需要本地紀錄或 tag，先定義權威、衍生關係與失敗復原；標籤格式、unknown 是否阻擋、審核角色與例外流程由專案政策決定。
5. 衍生紀錄使用當次凍結輸入重建，避免依賴之後可能變動的前一筆紀錄。
6. 部署成功陳述與版本已讀回分開記錄；不得把尚未實跑的部署設計寫成通用已驗證流程。

#### Agent Action

先辨識輸入契約與前置條件，再依本條決策；不照搬來源專案的函式名、腳本路徑、固定數字或狀態碼。記錄採用／不適用的理由。

#### Goal / Action / Validation

- Goal: 部署現況應可追溯到各元件的讀回證據；操作者陳述、歷史文件與 git tag 分別標示來源。
- Action: 依 Generalized Lesson 記錄輸入、決策與反例。
- Validation or reference source: 在來源專案驗證部署中斷、局部讀回、unknown 與成功路徑；本輪沒有實跑部署，特定閘設計仍是待驗證候選。

#### Applies When

- 部署現況需交接、元件可能不同步，或變更有操作者前置作業。

#### Does Not Apply When

- 已有完整 CD 現況與通知能力；此時整合原平台，不強制另建帳本／tag／草稿鎖。

#### Validation

在來源專案驗證部署中斷、局部讀回、unknown 與成功路徑；本輪沒有實跑部署，特定閘設計仍是待驗證候選。

#### Promotion Target

[workflow/software-delivery/release/README.md](../../../../workflow/software-delivery/release/README.md)。

#### Promotion Record

2026-10-26：抽象化整理並接到上述方法／入口；這是候選指引的連結，不宣告 validated／promoted，也不新增獨立重用證據。

#### Required Linked Updates

- 本分類 README 與上述方法／入口已檢查並同步。
- 索引與 RAG 經 runtime refresh 重建；原始專案證據不搬入共用庫。

#### Abstraction Check

- Swap-project test: PASS — 自動 CD 與人工部署都能採用「讀回事實、來源與未知分開」的原則，不需共用帳本格式或腳本語言。
- Project evidence kept in: `<PROJECT_ROOT>/` 的受控驗證紀錄。
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

依近期回饋審查，把原先依賴專案術語或操作細節的敘述改為輸入、機制、決策與驗證；修訂前內容可由 Git 歷史追溯。Status 保守維持 candidate。
