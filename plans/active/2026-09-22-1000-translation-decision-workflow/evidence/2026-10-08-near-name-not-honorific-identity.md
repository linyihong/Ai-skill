# Near-name tokens do not close honorific identity

## 決策

稱謂表面（姓＋職稱）的真名探查：**同片個人名／OCR↔ASR 近形衝突／水印短串
≠ 自動 identity bind**。無同窗共現或獨立覆核等同證據時維持
`defer_no_bind`；honorific locale 譯名繼續 deferred。

## 觀測（sanitized）

- 稱謂 token 占絕對多數；regex 常把「稱謂＋後續詞」誤切成假人名。
- 同系列出現個人名與拉丁 MT 形式，但與稱謂表面 **0** 同窗共現。
- OCR／ASR 對鄰近 cue 可給近形不同姓／名；兩側都要保留，不可 silent merge。
- 廠標／浮水印短串可看起來像人名，不得寫入 cast／established。

## 未完成

名覆核隊列（其他个人名／稱謂）、semantic／timing／ASS／burn、production
publish 路徑仍開放。具體證據留在 `<PROJECT_ROOT>` analysis。

## 連動

- Lesson：`feedback/history/narrative-video-production/common/2026-10-08_133500-near-name-not-honorific-identity.md`
- 既有：`2026-10-08-independent-name-review-binding.md`、honorific-surface lesson
