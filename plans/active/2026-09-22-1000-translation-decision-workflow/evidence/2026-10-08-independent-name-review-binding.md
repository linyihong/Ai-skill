# Independent name-review binding — honorific defer & conflict retention

## 決策

Translation Decision dogfood 的 Reference／Acceptance 分責下：目標譯名
established 只能來自**獨立覆核**，不得由 MT、phrase cache 或「與英文一致」
捷徑關閉。稱謂表面與個人名分開處理。

## 可重用規則（去敏）

1. **Address／honorific surface ≠ canonical identity**
   姓＋職稱／呼語可有片內 OCR／ASR 證據，但仍可 identity=unresolved。
   此時 locale 候選只可 `needs_review`／`deferred`，不可 `accepted`。

2. **Source identity ≠ target-name acceptance**
   來源身份 resolved 後，`validation_status`／Finality 仍可 not_evaluated／
   blocked，直到 independent accept 綁定 locale＋evidence_refs。

3. **衝突與殘留全留後再裁決**
   同一 identity／locale 的多個 MT 形式保留為候選；獨立覆核可 accept 一筆、
   reject 衝突筆與拉丁殘留。禁止 first-match 或自動選「看起來像」的一筆。

4. **Review binding**
   established 需要：`status=accepted`、`reviewer_role=independent`、
   reviewer id、source_name 與當前 identity 一致、locale 精確、evidence_refs
   非空。Selection／overnight extract 角色不得自關閘門。

## 觀測（sanitized）

凍結 local 輸出唯讀重放：獨立覆核 accept 兩筆個人名譯名後，
`established_names` 計數恰為 2；稱謂表面三候選全部 defer，不增加
established。content／artifact 仍可不 accepted。baseline 未改、無重跑翻譯。

## 未完成

稱謂真名證據、必要 semantic assertions、timing／visual／ASS／burn、
production 發布路徑仍開放。具體片名與路徑留在 `<PROJECT_ROOT>` analysis。

## 連動

- Adapter：[`dogfood-acceptance.md`](../../../workflow/translation/adapters/dogfood-acceptance.md)
- NVP 身份先於名稱：plan companion `08-identity-precedes-naming.md`
- Lesson：`feedback/history/narrative-video-production/common/2026-10-08_113000-honorific-surface-not-established-name.md`
