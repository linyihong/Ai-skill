> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-24 - Session-first stealth auth for high anti-bot SPA

Status: candidate

#### One-line Summary

對同時具備 FunCAPTCHA／同等挑戰、TLS／JA3 指紋與多步驟同 URL 登入的 SPA，優先「有頭 stealth 瀏覽器一次登入 + 持久化 session」，不要用 vanilla headless 每次帳密冷啟動。

#### Human Explanation

高防護社群／消費級 SPA 常把 web client 當第一方信任面：登入在驗證密碼前就可能因 fingerprint 失敗。帳密填寫「成功」卻回 Something went wrong、或新 IP 必出 Arkose 級挑戰，是預期現象而非 selector 寫錯。正確分析應先定 anti-bot 等級與 session 策略，再談發文／抓包。

#### Trigger

- 任務需要登入後的寫入動作（發文、送表單）或登入後 API 發現
- 目標為 JS SPA、登入步驟同 URL 切換、出現 Arkose／FunCAPTCHA／Akamai 指紋訊號
- Vanilla Playwright／Selenium 冷登入反覆失敗

#### Evidence

- Tool: browser analysis against public anti-bot literature + `analysis/web` six-step application
- Sanitized excerpt: multi-step SPA login; TOTP-capable 2FA; FunCAPTCHA-class iframe; TLS fingerprinting; success cookies include long-lived session + CSRF pair
- Evidence path: `<PROJECT_ROOT>/x-automation/docs/00-ai-skill-analysis-flow.md`（專案對照；不含 secrets）

#### Generalized Lesson

1. **Session-first**：冷啟動登入成本高時，分析與實作預設為 headed one-shot → persistent user-data-dir 或 `storage_state` → 後續 run 只做 session 健康檢查。  
2. **工具階梯**：`sources-and-tools` 的 HTTP → Playwright → Stealth 在 L4–L6（CAPTCHA + 行為／TLS 指紋）應明示 **stealth Chromium fork／StealthyFetcher 級**；不要停在「加 `--disable-blink-features=AutomationControlled` 就夠」。  
3. **同 URL 狀態機**：步驟偵測用 `data-testid`／文案，不依賴 location change；2FA 類型同樣靠文案分流。  
4. **登入後寫入**：先 UI 動作驗證 session，再依 `spa-api-discovery-via-browser` 做 Network／HAR 加速；勿一開始就重放未 live 驗證的 mutation。  
5. **CAPTCHA**：可見 FunCAPTCHA／enforcement 模式預設 `WAIT_HUMAN`（或合規打碼服務），不假設可程式「破解」。

#### Agent Action

- **先做**：Read `analysis/web/README.md`；填六步表；定 anti-bot 等級；選 stealth + headed + persistence。  
- **不要做**：無頭狂重試帳密；把 project cookie／TOTP／host 寫進 Ai-skill；把 browser_review 當 ToS compliance proof。

#### Goal / Action / Validation

- Goal: 可重用的高 anti-bot SPA 登入／session 分析決策，避免錯誤工具階梯。  
- Action: 更新 `analysis/web` 工具表與 CAPTCHA 策略；沉澱本 lesson；專案側保留目標站細節。  
- Validation or reference source: 專案文件六步對照表完整；lesson 無 secrets／無特定帳號；`anti-bot-bypass` 含 session-first 條目。

#### Applies When

- 登入閘門伴有 CAPTCHA／指紋／CDN bot manager  
- 需要維持登入態做後續讀寫或 API 發現  
- 官方 API 不可用或刻意不採用、改走瀏覽器

#### Does Not Apply When

- 靜態 HTML、無登入、L0–L2 保護（HTTP client 足夠）  
- 已有穩定 OAuth／官方 user-context API 且授權允許  
- 僅一次性人工操作、無需自動化

#### Validation

- 另專案遇到同級 anti-bot SPA 時，agent 先提出 session-first 而非冷登入迴圈  
- `analysis/web/sources-and-tools.md` 極高欄含 stealth fork／headed persistence  
- `git diff` 無本機路徑、token、帳號

#### Promotion Target

- `analysis/web/README.md`（步驟 3–4 補 session-first）  
- `analysis/web/sources-and-tools.md`（極高 anti-bot 工具列）  
- `analysis/web/spa-api-discovery-via-browser.md`（CAPTCHA 表）  
- `intelligence/web-scraping/anti-bot-bypass.md`（L4–L6 + session-first）  
- `intelligence/web-scraping/README.md`（atoms 索引若需要）

#### Required Linked Updates

- 已同步更新上述 Promotion Target（本 round）。  
- `feedback/history/web-scraping/README.md` 索引列已加。  
- Project evidence 留在 `<PROJECT_ROOT>/x-automation/docs/`；已依 reusable-guidance-boundary 檢查。
