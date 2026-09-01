# Media Entitlement Control Plane（媒體授權控制面）

**Status**: `candidate-analysis`

## 目的

判斷付費／解鎖媒體是**伺服器授權**還是**用戶端假象**，並在第一方設計／`security-audit` 時檢查控制面是否 fail-closed。本方法是觀察與審查路線，不是未授權存取步驟。

## 何時觸發此分析

- 列表、詳情或播放 API 同時出現價格、解鎖提示、試看時長與媒體 URL／playlist／wrapped key。
- 業務碼或 `can_play` 類旗標為否，但 JSON 仍帶可播欄位。
- 第一方產品要設計播放授權、短效媒體憑證或 HLS／DASH license。
- SDK 匯出／下載把「有 URL」當成「已授權」。
- `security-audit` 觸及付費播放、membership、或 playlist 發放。

## 分類（觀察面）

依**可播欄位是否出現**與**跨 identity** 分類，不要依 UI overlay 或欄位名稱：

| 標籤 | 觀察 | 結論 |
| --- | --- | --- |
| `server-grant` | 未授權列省略 playlist／src／unwrap key；授權後才出現 | 控制面在伺服器 |
| `control-plane-leak` | 旗標／業務碼拒絕，但 URL 或 wrapped key 仍在 | 用戶端必須 fail-closed；第一方不得這樣發放 |
| `ui-only` | 完整 source 已在列表／詳情，且 transport 可播；價格或鎖只在 UI | 不是伺服器鎖 |
| `inconclusive-preview-label` | 僅有 `preview*` 名稱 | 追 player 實際採用的欄位 |
| `identity-bound` | 授權 identity 有可播欄位，另一 identity 同內容仍省略 | 帳號綁定成立；本機 registry 不能當證明 |

試看資產必須是**獨立短媒體**，不可與正片共用同一可播 URL。

## 第一方設計層（審查面）

由外到內疊加；內層不能取代外層：

1. **省略可播欄位**：未授權回應不發 playlist、直鏈、unwrap key。
2. **短效、可撤銷憑證**：播放 URL 綁 session／裝置、TTL 短、過期重簽。TTL 只縮小持有者重放窗口，不能當唯一控制。
3. **分段 license**（有餘力）：每個 segment／key 再驗 entitlement；一條長壽命 manifest 等於把整包交給持有者。
4. **CDN 簽名**是輔助，不是授權來源：憑證一旦離開 API 就離開你的控制面。

用戶端加密標頭、混淆、本機鎖 overlay 不是授權邊界（見 `client-encrypted-header-not-boundary` lesson）。

## Audit Steps

```text
1. 列出可播欄位
   playlist / src / sourceURL / unwrap key / license URL
   與價格、free-time、lock 旗標分開記

2. 對未授權 identity 看欄位存在與否
   空欄位 → server-grant 候選
   有 URL 卻拒絕旗標 → control-plane-leak

3. 追 player 消費欄位
   不要用 preview 名稱當證據

4. 跨 identity 對照同一內容
   本機下載 registry ≠ 伺服器綁定

5. 第一方設計審查
   省略 → 短 TTL 重簽 → 分段 license
   預覽與正片拆資產
```

## Failure Signals（紅旗）

| Signal | 可能問題 |
| --- | --- |
| 拒絕碼與可播 URL 同 envelope | 控制面洩漏 |
| 列表已有完整 source，詳情只加鎖 UI | 用戶端假象 |
| 試看與正片同一 URL | 預覽標籤失效 |
| 長 TTL 或不綁 session 的媒體 URL | 持有者重放窗口過大 |
| SDK `hasUrl` 當下載授權 | 把洩漏當 grant |

## 證據蒐集規範

依 [`enforcement/sanitization.md`](../../enforcement/sanitization.md)：報告只留分類標籤與欄位**存在／空缺**；不寫 host、token、signed query、raw URL。具體樣本留專案文件。

## 驗證

| 問題 | 期望答案形式 |
| --- | --- |
| 未授權時可播欄位是否省略？ | yes/no + 哪個欄位 |
| 旗標拒絕時是否仍帶 URL／key？ | yes/no |
| Player 實際消費哪個欄位？ | 欄位名（泛化） |
| 另一 identity 是否仍省略？ | yes/no 或未測 |
| 第一方是否省略後才發短效憑證？ | 層級 1–3 勾選 |

## 與其他層的關係

- 觀察分類 lesson：[`feedback/history/apk-analysis/http-api/2026-09-01_093500-entitlement-grant-is-playable-field-presence-not-client-flag.md`](../../feedback/history/apk-analysis/http-api/2026-09-01_093500-entitlement-grant-is-playable-field-presence-not-client-flag.md)
- 設計 lesson：[`feedback/history/development-guidance/controls/2026-09-01_095200-media-entitlement-omit-playable-fields.md`](../../feedback/history/development-guidance/controls/2026-09-01_095200-media-entitlement-omit-playable-fields.md)
- 執行與修補：[`workflow/software-delivery/`](../../workflow/software-delivery/README.md) invoke `security-audit`；**不**另開 `workflow/security/`
- 授權邊界：[`enforcement/authorization-scope.md`](../../enforcement/authorization-scope.md)

← [回到 analysis/security/](README.md)
