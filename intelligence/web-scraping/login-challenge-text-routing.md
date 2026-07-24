# Login Challenge Text Routing

## 問題

高 anti-bot SPA 登入常在**同一條 URL**上切換多種挑戰：

- 郵箱一次性碼（Email OTP）
- 簡訊一次性碼（SMS OTP）
- Authenticator（TOTP）
- 異常登入身分確認（Username／「Confirm your account」類）

若把任何「verification code」都當 TOTP 自動填，會：

1. 把郵箱碼欄位填成錯誤碼 → 失敗、增加風控摩擦  
2. 讓操作者誤以為「2FA 壞了」而非「需要人工收信」

## 判斷規則（文案優先）

| 分類 | 典型文案訊號（語言無關，舉例） | Agent 行為 |
| --- | --- | --- |
| `email_otp` | check your email、sent to your email、郵箱／信箱、遮罩 `a***@…` + code | **停機**；提示 headed 人工完成一次；**禁止**填 TOTP |
| `sms_otp` | text message、sent to your phone、短信 | **停機**；提示人工簡訊；禁止填 TOTP |
| `totp` | authenticator、authentication app、authentication code、验证器 | 僅此時用 TOTP secret 自動填 |
| `knowledge` | confirm your account、enter information associated、knowledge_check | 填 Username／handle，或點「Use password」**連結**；完成後**重新分類** |
| `unknown_code` | 僅有 verification／confirmation code，無 authenticator／email／sms 訊號 | 寧可當 email OTP 停；有 TOTP secret 也不要猜填除非產品明確允許 |

## 探測與生產分離

| 路徑 | 用途 |
| --- | --- |
| **Ephemeral／incognito profile** | 刻意觸發挑戰，抓真文案校正 classifier |
| **Persistent profile（session-first）** | 日常自動化；一帳一 profile、一帳一出口 |

不要把無痕探測當預設生產登入。

## 真 headed 條件

- `--headed` 若跑在無互動桌面的 SSH／服務工作階段，Continue／連結常無法啟用。  
- 校正文案或人工 OTP：需要**互動桌面**（例如已連線的 RDP）或明確的 `WAIT_HUMAN` 等待窗。

## 控件角色

「Use password」等可能是 **link／span**，不是 button。點擊策略：`role=link` → `role=button` → `get_by_text` → DOM 文案 fallback。

## 與既有智慧的關係

- Session-first／stealth 階梯：[`anti-bot-bypass.md`](anti-bot-bypass.md)、[`../../feedback/history/web-scraping/common/2026-07-24_093318-session-first-stealth-auth-high-antibot-spa.md`](../../feedback/history/web-scraping/common/2026-07-24_093318-session-first-stealth-auth-high-antibot-spa.md)  
- 分析方法入口：[`../../analysis/web/README.md`](../../analysis/web/README.md)  
- 本 lesson：[`../../feedback/history/web-scraping/common/2026-07-24_162800-login-challenge-text-routing-email-otp-vs-totp.md`](../../feedback/history/web-scraping/common/2026-07-24_162800-login-challenge-text-routing-email-otp-vs-totp.md)

## 不適用

- 純 OAuth／官方 API、無瀏覽器登入閘門  
- 產品保證永遠只有 Authenticator、從不出現郵箱／簡訊碼
