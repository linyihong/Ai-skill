# 驗證階段的經驗查找與採用

本文件由 [`validation.md`](validation.md) 載入，適用 `sd-validation`。它規定 workflow 如何使用候選；搜尋、selector 與 cache 規格仍由 [CLI contract](../../scripts/ai-skill-cli/docs/command-contract.md) 擁有。

## 何時查找

先完成 route 契約指定的 required_read 並列出未解 blocker，再決定下表動作。route 已命中仍須依任務查經驗；detector miss 的 Bridge advisory 不會啟用 workflow，也不會讓候選成為必讀規則。

| 事件 | 動作 | 留存證據 |
| --- | --- | --- |
| 實作完成，形成最低足夠驗證方案前 | 查與 claim／風險／證據缺口相關的既有經驗 | query、任務範圍、事件與 trace |
| 驗證失敗，證據鏈斷點不明 | 查失敗訊號與缺口，先讀回再改方案 | observed failure、補讀目的、採用／拒絕理由 |
| 前提、任務、失敗類型或來源版本改變 | 重查或重新讀回並重驗適用性 | 舊／新版本、改變的前提、新決策 |
| 同一任務、前提相同、來源 checksum 未變且已有決策紀錄 | 可略過重複 query | 引用舊 trace，當次 checksum 核對與理由 |
| 純格式改動且現有 checklist 已完整覆蓋 | 可略過經驗查找 | diff 範圍與 checklist；不略過 required_read |

沒有先前紀錄、未核對版本或風險已變時，不可用「已查過」略過。

## 執行與回讀

1. 使用既有 production `runtime lookup index/prepare`，只加任務可證明的 filters；缺 metadata 為 unknown。將固定 request 交給現有 agent／runner，再由 watcher 與 `lookup complete` 核對。失敗不修補 raw、不回退舊 decisions。
2. direct 只是待讀回來源。background 預設不填主要上下文；unknown 可補查或維持未知，不能當成不相關。空輸出須保留 pool／unassessed／unknown 與 trace。
3. 讀回每個準備使用的 canonical source 全文或與用途相關的段落，按 need_to_read 補足資訊；歷史 lesson 如指向專案證據，還須回到該證據。衍生 DB 的 checksum 不會提高它的規則權威。
4. 檢查 domain、專案、平台、版本與前提；分開相關性、證據有效性、supersession 與 lifecycle／authority。只有 metadata 或正文明確宣告才能判定取代關係，沒有欄位為 unknown。
5. 記錄候選 checksum、實際回讀 checksum，以及採用前再次核對的 checksum。A→B 時重驗 B，保留差異及理由；變化只觸發重驗，不能直接宣告結論失效。重驗中再次改變則停止沿用舊判定。
6. 對每個輸出來源留下 `adopt`／`reject`／`unknown`，說明本次用途、前提與證據。候選智慧可作本次判斷參考，不能升格為全域契約；歷史決策可用於追溯，但被取代者不能當現行規則採用。
7. 採用後執行對應驗證，再留下實際結果與 claim 邊界。找到、讀回、適用、採用、驗證成功是不同狀態。

## 失敗與既有 gate

- 檢索關閉、缺索引、stale、unavailable 或 timeout 時，回到既有直接來源／人工查找並記原因；必要證據仍缺則保持 blocker。不得拿舊 payload 當成功結果。
- required_read 由 route／契約確定性載入，不受排名、預算或 feature 開關影響。
- 採用經驗不能結清 blocker；只有原 owner gate 要求的證據或明確授權處置才能結清。
- 來源前提不足時保留 unknown 與下一個取證入口，不猜測平台、認證流程或專案事實。

## Context 與紀錄

使用 [查找決策紀錄](templates/retrieval-decision-template.md)，放專案文件／驗證紀錄，僅保存必要去敏事實。raw request/response 預設最多保留7日；可覆核的去敏決策紀錄依專案 retention 保存，不建立對話資料庫。

依任務核定的成本契約記錄完整實際送入的 message 範圍。可取得對應 model／tokenizer／version 時記錄實際計數；若使用者明確允許暫緩 token 量測，記 `deferred-by-user`、授權引用與仍適用的預算，照常使用既有 agent 執行，不以缺計數阻擋。這不表示 token 上限已通過，也不免除其他採用檢查。固定 encoding 的參考計數、runes 與實際模型用量分開標示，不相互代替。全文補讀與任務其餘 context 另列，不能冒稱整個任務只用了 retrieval message 的 tokens。

本流程由 workflow 與現有 agent 執行，不新增 activation、authority、全域 trigger 或通用自動採用引擎。
