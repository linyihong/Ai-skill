# Sub-flow: periodic-sweep

定期巡檢：**delta reassessment**，非全量重寫（H5）。  
本 flow 是 **report producer**；何時觸發由使用者／外部排程決定。

## 步驟

1. Intake：watchlist 來源（user-local）、上次 as-of、視窗  
2. Risk tier  
3. Research delta：新／變／失效；未變指向前次 brief  
4. 筆記對照僅列有更新者  
5. 產出 Sweep brief；若出現配置建議 → 轉 `allocation-advice`（勿在 sweep 偷做完整配置）  

## 產出

```markdown
## Sweep brief
- Window: <from> → <as-of>
- New / Changed / Invalidated
- Unchanged: (pointers only)
- Open questions
- Suggested follow-up task types
```

方法：[`news-trend-summary.md`](../../../analysis/investment/news-trend-summary.md) §Delta。
