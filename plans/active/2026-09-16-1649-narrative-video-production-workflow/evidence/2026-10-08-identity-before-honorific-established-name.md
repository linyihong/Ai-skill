# Identity before honorific-established name — dogfood confirmation

## 決策

確認候選原則 [`08-identity-precedes-naming.md`](../08-identity-precedes-naming.md)
與 [`05-series-cast-canonicalization.md`](../05-series-cast-canonicalization.md)：
`series_cast` 是**已解析**身份表；名稱是後續證據，不是出場權重捷徑。

## 可重用觀察

1. **出場／`confirmed` 權重 ≠ resolved identity**
   需 identity status、identity_id、canonical_name 與 nonempty evidence_refs
  （OCR／ASR／dialogue 等可追溯引用）。缺證據 → unresolved。

2. **稱謂表面不得硬連既有角色**
   片內僅見姓＋職稱呼語、早期 cast 試驗有其他姓氏職稱時，禁止發明連結。
   可記錄 surface_evidence；不可假寫 true_name。

3. **每片 bundle 與譯名 export 分層**
   本片 `series_cast` 管來源身份；可選 `identity_translations` 是 adapter-owned
   review export（非全域詞庫、非 phrase cache 投影）。兩者都不是 production
   publish gate 的自動關閉條件。

4. **與字幕 content gate 分開**
   Reference／譯名 established 推進後，NVP timing／layout／burn 仍可 open；
   不得合成單一「字幕 PASS」。

## 未完成

真名補證、semantic／timing／layout／成片驗收、自動身份 producer 仍開放。
具體 run 統計留在 `<PROJECT_ROOT>`。

## 連動

- Translation Decision evidence：`2026-10-08-independent-name-review-binding.md`
- Lesson：`feedback/history/narrative-video-production/common/2026-10-08_113000-honorific-surface-not-established-name.md`
