# News & Trend Summary Template（新聞與趨勢摘要）

把新聞／趨勢編成可覆核摘要，服務 `theme-research`、`name-diligence`、`periodic-sweep`、`allocation-advice` 的 Research 段。

## 何時使用

- 需要「最近發生什麼」且之後要接 recommendation  
- Sweep：只寫 **delta**，避免全量重寫（H5）

## 摘要模板

```markdown
## Scope
- Subject: <theme | ticker | portfolio-theme>
- Window: <from> → <as-of>
- Markets: <TW | US | both>

## Headlines（每則一行）
| Date | Claim (one line) | Authority | Locator | Stance | Notes |
| --- | --- | --- | --- | --- | --- |
| YYYY-MM-DD | … | A–D | url or excerpt-id | supports/weakens/context | paraphrase? |

## Trend synthesis（≤5 bullets）
- …（每點可回溯到上表列）

## Uncertainty labels
- Base / Bull / Bear（或同等 framing）— **禁止**無校準精確勝率 %
- What would change the view

## Delta vs prior（sweep only）
- New: …
- Changed: …
- Invalidated: …
- Unchanged (omit full rewrite): …
```

## 規則

| 規則 | 說明 |
| --- | --- |
| 一則一 claim | 勿把整篇新聞糊成一句「利多」 |
| Authority 必填 | 見 [`sources-and-tools.md`](sources-and-tools.md) |
| 轉述降級 | Paywall／二手摘要 → C＋`paraphrase` |
| Sweep = delta | 未變段落指向前次 brief，不重貼（H5） |
| 不關閉決策 | 摘要結尾不寫「因此應執行…」；決策留給 workflow＋human selection |

## Uncertainty framing（H3）

**Prefer**：`low / medium / high confidence`、情境權重敘事、前提清單。  
**Reject**：無校準的 `60% 上漲`、`高機率最優`（DVA 05 已擋）。

若必須引用他人給出的百分比：標 `source-stated-%`＋authority，並說明本報告 **不採納為校準機率**。
