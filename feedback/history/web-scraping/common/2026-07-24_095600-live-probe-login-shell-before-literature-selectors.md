> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-24 - Live-probe login shells before trusting literature selectors

Status: candidate

#### One-line Summary

登入入口 URL 可能 redirect 到全新 onboarding shell，舊 `data-testid` 文獻契約會失效；實作前必須 headed live probe 欄位 `name`／按鈕文案。

#### Human Explanation

社群／大型 SPA 常把「經典多步驟 login flow」換成新的 onboarding web（SSO + email 同屏）。自動化若硬編碼舊 testid，會在「頁面已開」卻找不到欄位時誤判為指紋失敗。正確做法是 probe 真實 DOM，雙軌支援新舊殼，並避免誤點進 phone signup。

#### Trigger

- `*/flow/login` 最終 URL 變成 `/…/onboarding…` 或未知 path
- 初始 HTML／headed 首屏幾乎沒有預期 `data-testid`
- 可見 SSO（phone／Google／Apple）+「Email or username」

#### Evidence

- Tool: Patchright headed probe + session-first login
- Sanitized excerpt: redirect from classic flow path to onboarding web; fields `username_or_email` / `password`; primary CTA `Continue`; float-label intercepts clicks
- Evidence path: `<PROJECT_ROOT>/x-automation/docs/05-headed-dom-probe-2026-07-24.md`

#### Generalized Lesson

1. **Probe-before-contract**：登入自動化的 selector 契約以 live headed DOM 為準，文獻 testid 只當候選。  
2. **Detect shell**：用 URL 片段 + `input[name=…]` 分流新／舊 login 表面。  
3. **Avoid signup traps**：同屏 phone 欄存在時，優先填 email／username 路徑，偵測「Enter your phone」則 fail loud。  
4. **Click force**：float-label 容器可能攔截 pointer → `force=True` 或點 input 本體。

#### Agent Action

- 先 `probe` 輸出 testid／input name 清單，再寫 login 狀態機。  
- 不要因「找不到 LoginForm_*」直接當成 anti-bot 失敗。

#### Goal / Action / Validation

- Goal: 降低 SPA 登入殼改版導致的假陰性。  
- Action: 雙軌 login + project probe 文件；本 lesson 入 `feedback/history/web-scraping`。  
- Validation: 新殼帳密登入可到 home；status 重用 profile。

#### Applies When

- 目標有登入閘門且前端可能 A/B 或改版 onboarding  
- 任務含 browser automation login

#### Does Not Apply When

- 僅靜態表單、selector 長期穩定  
- 已使用官方 OAuth／API，不走 web login UI

#### Validation

- 另專案若 flow/login redirect，agent 先 probe 再編碼  
- Lesson 無帳號／cookie／host 實值

#### Promotion Target

- `analysis/web/spa-api-discovery-via-browser.md`（登入前加 probe note）  
- `intelligence/web-scraping/adaptive-parsing.md`（可選交叉連結）

#### Required Linked Updates

- 更新 `spa-api-discovery-via-browser.md` 一小節；索引 `feedback/history/web-scraping/README.md`。  
- 專案證據留 `<PROJECT_ROOT>/x-automation/docs/05-…`。
