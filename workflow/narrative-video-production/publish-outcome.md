# Publish outcome

流量用來檢驗**模板假設**，不是作業去重，也不是「已證明模板有效」。Invariant 9。

單位：**模板 × 窗口**（同一平台、同一窗口定義）。

## v0 欄位（Q5；可修）

`impressions`、`views`、`avg_watch_pct`、`engagements`、窗口天數。

## evidence_status

| 值 | 何時 |
| --- | --- |
| `supports` | 同模板 ≥ 2 部、同一窗口定義，目前證據支持該假設 |
| `contradicts` | 同上，目前證據反對該假設 |
| `insufficient_sample` | 單部爆款、樣本不足、窗口未滿 |

這是 **evidence status**，不是 truth status。禁止無觀察的綜合品質分；禁止把平台原始 dump 整包提交進本知識庫。去敏後才回寫模板 evidence 索引。
