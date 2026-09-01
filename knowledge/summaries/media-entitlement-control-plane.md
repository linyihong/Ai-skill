## analysis.media-entitlement-control-plane

| 欄位 | 值 |
| --- | --- |
| Atom ID | `analysis.media-entitlement-control-plane` |
| Source path | `analysis/security/media-entitlement-control-plane.md` |
| Lifecycle | `candidate` |
| Summary | 付費媒體授權以可播欄位存在與否分類（server-grant／control-plane-leak／ui-only）；第一方設計先省略欄位，再短效憑證，再分段 license。TTL 只縮小重放窗口。 |
| When to read | 付費播放 API 同時有價格與媒體 URL；旗標拒絕仍帶 URL；設計短效播放憑證；`security-audit` 觸及 playlist 發放。 |
| Do not use for | 未授權取得步驟；把 preview 欄位名當證據；以本機下載 registry 證明他帳已授權。 |
| Context cost | ~280 tokens |
| Estimated full cost | ~1800 tokens |
| Validation signal | 能回答未授權是否省略可播欄位、player 消費欄位、跨 identity 是否仍空、設計層 1–3 是否到位。 |
| Last checked | 2026-09-01 |
