# Unity case ledger template（精簡）

複製到**業務專案** evidence 使用。本庫不保存 raw 專案證據。

必要六欄；其餘按需展開。未知可留白，但須標明**哪些未知會影響判斷**。

## 必要欄位

### 1. `target_outcome`

要達成的可觀察結果（一句）。

### 2. `acceptance_oracle`

如何區分成功／失敗／證據不足（與產生結果的那次 API 呼叫可分開）。

### 3. `evidence_scope`

build／fixture／物件身分／時間窗口／掃描範圍；自然 vs force 是否隔離。

### 4. `hypotheses`

| id | statement | class (obs/inf/unk) | status (supported／ruled-out-in-scope／not-tested／blocked) |
| --- | --- | --- | --- |
| H1 | | | |
| H2 | | | |

### 5. `next_experiment`

| field | value |
| --- | --- |
| question | |
| expected_if_h1 | |
| expected_if_h2 | |
| falsifier | |
| observation / result | |
| decision_after_result | |
| side_effect_and_rollback | |

### 6. `conclusion`

oracle 結果、已排除／仍 open、blocked 原因、下一最小步驟。

## 可選展開

完整 hook 事件表、完整 build fingerprint、screencap 清單、回滾腳本——僅在該案需要時附加。

← [`decision-workflow.md`](decision-workflow.md)
