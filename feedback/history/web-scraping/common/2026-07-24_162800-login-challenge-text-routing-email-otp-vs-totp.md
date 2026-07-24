> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-24 - Login challenge text routing (email OTP vs TOTP vs identity)

Status: candidate

#### One-line Summary

多步驟登入挑戰必須**先依頁面文案分類**（email OTP／SMS OTP／Authenticator TOTP／identity knowledge），再決定自動填碼或停機；禁止用 TOTP 硬填郵箱碼；無痕 profile 可刻意觸發挑戰以便校正分類器。

#### Human Explanation

高 anti-bot SPA 登入常在同 URL 切換多種挑戰。若把「verification code」一律當 Authenticator，會把郵箱一次性碼誤填成 TOTP，既失敗又消耗帳號信任。正確做法是文案優先分流；冷／無痕環境更容易先出現 identity confirm（例如 Confirm your account），然後才是郵箱 OTP——這是觸發驗證分類器的合法探測路徑，不是日常發文路徑。

#### Trigger

- 瀏覽器自動化登入含 2FA／異常登入確認
- 頁面同時可能出現郵箱碼、簡訊碼、Authenticator 碼
- 需要驗證分類器是否認得真實挑戰文案

#### Evidence

- Tool: ephemeral（incognito）browser probe + persistent session path comparison
- Sanitized signals: page copy markers for email／SMS／authenticator／identity-confirm；headless Continue often stays disabled without interactive desktop; 「Use password」 may be link not button
- Project evidence stays under `<PROJECT_ROOT>` dump／docs（no secrets in this lesson）

#### Generalized Lesson

1. **文案優先分類**：`email_otp`／`sms_otp`／`totp`／`knowledge`／`unknown_code`；email／SMS 命中則**立即停機**並提示人工一次，**不**呼叫 TOTP。  
2. **unknown_code**：僅有「verification code」而無 Authenticator 字樣時，寧可當 email OTP 停，不要猜填 TOTP。  
3. **Identity confirm ≠ OTP**：確認帳號／Username 步驟先填 handle 或走「Use password」連結；過關後再重新分類。  
4. **探測 vs 生產**：用 ephemeral profile＋獨立出口刻意觸發挑戰以抓真文案；日常作業仍 session-first（一帳一 profile）。  
5. **真 headed**：SSH／服務工作階段的 `--headed` 若無互動桌面（RDP 斷線），Continue／連結常點不動；校正分類器需互動桌面或 `WAIT_HUMAN`。  
6. **控件角色**：文案按鈕可能是 `link`／`span`，不要只 `get_by_role("button")`。

#### Agent Action

- **先做**：實作／審查登入狀態機時加入 challenge classifier；email／SMS → fail with operator message；TOTP only when authenticator markers present。  
- **不要做**：把專案帳號、郵箱、IP、host 寫進 Ai-skill；把無痕探測當預設生產路徑；宣稱可自動收郵箱 OTP。

#### Goal / Action / Validation

- Goal: 可重用的登入挑戰文案分流，降低誤填與帳號摩擦。  
- Action: 更新 `analysis/web` 登入閘門檢查點；`intelligence/web-scraping` 增加判斷原子；本 lesson。  
- Validation: 單元測試用去敏文案 fixture；live probe dump 只留專案側。

#### Applies When

- SPA／stealth 瀏覽器登入含多種驗證碼挑戰  
- 需要 session 持久化與偶發人工 OTP  

#### Does Not Apply When

- 純 OAuth／官方 API、無瀏覽器登入  
- 單一固定 Authenticator、從不出現郵箱／簡訊碼  

#### Validation

- Classifier tests：email／SMS／TOTP／knowledge 文案互不誤判  
- 誤把 email OTP 當 TOTP 的路徑在 code review 被擋  

#### Promotion Target

- `analysis/web/README.md`（登入閘門挑戰分類表）  
- `intelligence/web-scraping/login-challenge-text-routing.md`（新 atom）  
- `intelligence/web-scraping/README.md`（索引）  
- `intelligence/web-scraping/anti-bot-bypass.md`（交叉連結）  

#### Required Linked Updates

- 本 round 同步更新上述 Promotion Target 與本 domain README 索引。
