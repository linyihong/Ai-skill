# First Consumer Names A Shared Capability After Itself（第一個使用者以自己命名共用能力）

Status: candidate
Class: `source-of-truth-duplication`

## Trigger

某個切片要建立一項**機制性**能力 —— 分區與清理、重試、匯出、快取失效、批次排程、稽核落庫 —— 而它目前只有一個使用者。

實作自然以那個使用者命名：`audit_ensure_partitions`、`order_export_job`、`invoice_retry_policy`；表名、欄位名寫死在 routine 內部；介面只收該使用者需要的參數。

## Failure Mode

能力本身是通用的，命名與簽章卻不是。第二個使用者出現時，通用化的成本已經不落在程式碼上，而落在**已上線的資料結構與運維程序**上：

| 通用化路徑 | 代價 |
| --- | --- |
| 改造原能力 | 要遷移正在使用中的資料結構（已分區的表、已排程的 job、已存在的歷史列） |
| 並存兩套 | 兩套生命週期／兩套設定／兩份文件，且兩者的差異無人檢查（見 [`cross-boundary-agreement-never-mechanically-checked.md`](cross-boundary-agreement-never-mechanically-checked.md)） |
| 第二個使用者自建 | 最省事，也最常發生 —— 於是同一個概念在系統裡有兩個實作，而它們會漂移 |

**這與「過早抽象」不同，差別在時序證據。** 過早抽象是為想像中的使用者付成本；本 pattern 的前提是**第二個使用者已經可指名**：另一張表的鍵已經預備成同樣形狀、另一份計畫已經在排程上、另一個模組已經有同樣的需求記錄在案。

實際觀察到的形態：稽核事件的分區與保留設計完整（月 range partition、預建未來月份、到期 detach/archive/drop），但 routine 與 port 皆以 `audit_` 命名並寫死該表。同時另一張表的主鍵早已為同一目的預備成 `(id, creation_time)`，且系統已排定三個月後要建立通用歸檔框架。

## Risk

- 遷移活的資料結構，風險與成本都遠高於當初多收兩個參數
- 並存兩套機制後，設定、文件與運維程序各自分裂，且沒有任何檢查會因它們不一致而失敗
- 第二個使用者選擇自建時，系統從此有兩個「同一件事」的實作，且通常無人記得它們本該是一個

## Required Agent Action

**在第一個使用者的設計階段就問「這個能力預期有第二個使用者嗎」，並讓答案影響命名與簽章。**

1. **判定它是機制還是領域邏輯。** 分區、保留、重試、匯出、排程屬機制；「訂單如何計價」屬領域。機制預設會有第二個使用者。
2. **找第二個使用者的既有證據。** 是否已有另一張表／模組的鍵、欄位或計畫為同一目的預備？有具名證據就不是想像。
3. **有具名證據時，收參數而非寫死。** 表名、時間鍵、期間、模組識別作為參數；命名用中性詞（`partition_prune` 而非 `audit_prune`），第一個使用者只是第一個註冊對象。
4. **成本要兩邊都估。** 「現在多收兩個參數」對比「事後遷移已上線結構」或「並存兩套」。前者通常小一個數量級 —— 但若第二個使用者只是猜測，這個比較不成立，此時寫死是對的。
5. **把判定寫成決策而非留在腦中。** 記錄「已知第二個使用者是誰」與「因此命名為何中性」，否則接手者會看到一個過度一般化的介面而不知原因，可能把它改回專屬。

## Prevention Gate

建立一項機制性能力時：

- 這是機制還是領域邏輯？
- 第二個使用者**可指名**嗎？有沒有已存在的鍵／欄位／計畫指向同一需求？
- 若有：現在收成參數要多少工？事後遷移活資料要多少工？
- 我的命名是否讓第二個使用者看起來像個外來者？
- 這個判定寫在哪裡？接手者看得到理由嗎？

## 驗證

1. 機制性能力的設計紀錄中，明確回答「是否已有可指名的第二個使用者」
2. 若有，介面以參數承載使用者差異，且命名不含第一個使用者的專屬詞
3. 若判定為單一使用者而寫死，該判定與其理由同樣記錄，日後出現第二個使用者時可追溯
4. 不存在「同一機制的兩套實作」而無任何檢查比對兩者

## Linked Rules

- [`cross-boundary-agreement-never-mechanically-checked.md`](cross-boundary-agreement-never-mechanically-checked.md) — 並存兩套後的漂移形態
- [`framework-duplication-without-interrogation.md`](framework-duplication-without-interrogation.md) — 同家族：改動 framework 前的 source-of-truth discovery
- [`preflight-scope-narrower-than-the-repository.md`](preflight-scope-narrower-than-the-repository.md) — 找既有實作／既有使用者的搜尋要求
- [`../failure-learning-system.md`](../failure-learning-system.md) — `source-of-truth-duplication` 分類
