# 本機 token 計數工具與參考結果

使用者已明確授權「有需要可以直接安裝」。已在獨立本機環境 `<USER_CACHE>/ai-skill/token-counter-venv` 安裝 tiktoken 0.14.0，Python 3.12.3。未改動系統 Python、repo CLI、release binary 或新增付費模型呼叫。

## 本次量測

固定 encoding 為 `o200k_base`，對已凍結的26份 `phase2-prod-*.payload.json` **完整 UTF-8 原始文字**計數，不重新序列化、不改原檔。每份留下 SHA、runes 與 reference_tokens；encode/decode roundtrip 全部一致。

最大值為 **449 reference tokens，R09**。逐題結果見 [計數產物](phase3-reference-token-counts.json)。這不重評 Phase 2 的成功門檻，也不回改既有 runes 結果。

模型對應 encoding 仍未知，故不能稱此數值為目前 agent 的實際 input usage；未計系統訊息、tool wrapper、來源全文補讀、selector 輸入輸出與其他對話 context。實際用量維持 unknown，token 驗收依 [使用者處置](phase3-token-disposition.md) 暫緩，不阻擋 Phase 3。

## 重現方式

使用獨立環境的 Python，載入 `tiktoken.get_encoding("o200k_base")`，對 artifact 的 UTF-8 文字執行 `len(enc.encode(text, disallowed_special=()))`。必須確認 encoding、tiktoken 版本及輸入 SHA 一致。這是離線文字計數；encoding 詞表首次載入需要下載，本輪已完成。未建立新的 repo script entrypoint。

官方來源：[openai/tiktoken](https://github.com/openai/tiktoken)。本機安裝與26份計數已實際執行，並非僅提供安裝指令。
