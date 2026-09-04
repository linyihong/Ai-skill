# Development Guidance Feedback History

## 分類

| 分類 | 數量 | 說明 |
|------|------|------|
| [`common/`](common/) | 51 | 跨分類或通用 lesson（開發流程、契約治理、測試策略、後端架構、安全審計等） |
| [`controls/`](controls/) | 2 | 控制項相關 lesson |

## 來源

所有 lesson 原位於 `skills/app-development-guidance/feedback_history/`，已於 2026-05-13 搬遷至此，舊目錄已刪除。

## Recent (2026-09-04)

| Slug | Category |
|------|----------|
| `common/2026-09-04_101800-derived-state-before-blaming-shared-baseline` | 「原始碼相同」＋「可重現」不足以歸因給共用基線；過期的衍生狀態同樣滿足這兩點 |
| `common/2026-09-04_101559-build-gate-must-restore-not-only-build` | 以 `--no-restore` 建置的 gate 驗的是還原圖不是工作樹；過期圖會誤過，比誤擋更危險 |

## Recent (2026-09-01)

| Slug | Category |
|------|----------|
| `controls/2026-09-01_095200-media-entitlement-omit-playable-fields` | 未授權省略可播欄位；短效憑證只縮小重放窗口 |

## Recent (2026-08-03)

| Slug | Category |
|------|----------|
| `common/2026-08-03_211615-windows-cmd-shim-needs-shell-not-cmd-suffix` | Windows 的 `.cmd` shim 要走 shell 啟動；補 `.cmd` 副檔名已因安全性修補失效；spawn 失敗需與非零結束分開診斷 |
| `common/2026-07-30_221700-docx-text-extraction-without-converter` | 缺 converter／office suite／腳本 runtime 時，OPC 格式仍可解壓取 XML 抽文字；先探測能力再宣告不支援 |
| `common/2026-07-30_135800-run-new-ban-against-existing-tree` | 新機械 gate 先對整棵既有樹跑一次；首跑輸出是 scope 探測，不是待修清單 |
| `common/2026-07-30_135700-unreachable-test-surface-delete-over-repair` | 測試 surface 不在任何 gate 內時，先證明它已 rot，再刪除＋加禁令，而非修好設定 |
| `common/2026-07-23_071200-injectable-wall-clock-for-day-boundary-logic` | Inject wall clock / Instant params; freeze time in tests, never change OS date |
| `common/2026-07-17_171500-mpegurl-maybe-managed-mediasource-prefer-hlsjs` | mpegurl `"maybe"` + ManagedMediaSource → prefer hls.js; Simulator ≠ paint |
| `common/2026-07-15_140200-correct-sign-length-still-403-means-anti-bot-not-sign` | Sign parity OK but 403 = anti-bot, not formula |
