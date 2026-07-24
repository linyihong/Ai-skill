# SPA 后台 API 发现 — 浏览器自动化

> 适用：`analysis/web/` 六步流程中 **步骤 2–4**（网站评估 → 工具 → 策略），目标为 **登录后 REST API 契约**，而非 HTML 爬取。

## 何时用

| 条件 | 推荐 |
| --- | --- |
| 目标为 SPA（Vue/React），业务数据走 XHR/fetch | ✅ 浏览器 |
| 仅静态 HTML | httpx 即可 |
| 需登录；表单无 CAPTCHA 或 CAPTCHA 可人工一次配合 | Playwright / CDP |
| 已有 HAR 导出 | 解析 HAR，不必自动化 |

**不要**在未 live 验证时用 monorepo 内 Controller 推断他人线上 API。

## 方法 A：Playwright（推荐，Agent 可脚本化）

```javascript
import { chromium } from 'playwright';

const browser = await chromium.launch({ headless: true });
const page = await browser.newPage();
const captured = [];

page.on('response', async (res) => {
  const url = res.url();
  if (!url.includes('/api/')) return;
  captured.push({
    method: res.request().method(),
    url,
    status: res.status(),
    body: (await res.text()).slice(0, 4000),
  });
});

await page.goto('https://target.example/#/auth/login', { waitUntil: 'networkidle' });
await page.locator('input[name="username"]').fill(process.env.USER);
await page.locator('input[name="password"]').fill(process.env.PASS);
await page.locator('button').filter({ hasText: /登录|Login/i }).click();
await page.waitForTimeout(3000);

// 遍历 hash 路由或侧栏，触发各模块 API
for (const hash of ['#/module/a', '#/module/b']) {
  await page.goto(`https://target.example/${hash}`, { waitUntil: 'networkidle' });
  await page.waitForTimeout(1000);
}
```

要点：

1. 监听 **`response`**，不只 `request`
2. 登录后 token 常在 `localStorage` / Cookie；后续请求 Header 由前端 interceptor 注入
3. 从 **`bootstrap*.js` + lazy chunks** regex 提取 `"/api/..."` 作 path 索引，再逐 path 触发

## 方法 B：Chrome DevTools 远程调试（CDP）

与 Playwright 同级：**驱动真实 Chrome**，适合需 DevTools Protocol、已有 `RWD_BROWSER_PATH` 的集成仓（例如 Vidoe-Test `player-complete-browser.mjs`）。

```javascript
import { spawn } from 'node:child_process';

const port = 9222;
spawn('/Applications/Google Chrome.app/Contents/MacOS/Google Chrome', [
  `--remote-debugging-port=${port}`,
  '--headless=new',
  '--no-first-run',
  'about:blank',
]);

// 连接：Playwright connectOverCDP(`http://127.0.0.1:${port}`)
// 或 puppeteer.connect / chrome-remote-interface
```

| 对比 | Playwright launch | CDP 连接已有 Chrome |
| --- | --- | --- |
| 安装 | `npm i playwright` | 系统 Chrome + debug port |
| Session/Cookie | 自动 | 与真实浏览器一致 |
| CI | 友好 | 需预装 Chrome |
| 典型用途 | 通用 API 抓取 | 与 H5/播放器等同进程调试 |

## 方法 C：人工 HAR

1. Chrome → Network → Preserve log  
2. 登录并点击各菜单  
3. Export **HAR with content**  
4. 从 HAR 提取 method、url、request/response JSON  

无 CAPTCHA 自动化时，HAR 往往最快。

## JS Bundle Path 扫描

```bash
# 下载入口 JS → bootstrap chunk → 提取引号内 /api/ 路径
curl -sS 'https://target/jse/index-*.js' | rg -o '"/api/[^"]+"'
```

注意：bundle 内可能有 **Vben 模板残留**（如 `/api/user/auth/login` demo path）；以 **Network live** 为准。

## CAPTCHA

| 情况 | 策略 |
| --- | --- |
| 登录页 **无** captcha 字段 | 账密自动化即可（低 anti-bot SPA） |
| 图形 captcha | 人工一次提供 code，或 HAR；见 `intelligence/web-scraping/anti-bot-bypass.md` §CAPTCHA |
| reCAPTCHA / hCaptcha / FunCAPTCHA | Stealth + headed；**session-first**（一次过关后持久化 profile）；打码服务或 `WAIT_HUMAN`；**不**假设能「破解」 |
| TLS／行为指纹导致「Something went wrong」 | 不要只改 delay；升级 stealth fork、对齐真实 Chrome channel、避免无头冷登录 |

相关 lesson：[`../../feedback/history/web-scraping/common/2026-07-24_093318-session-first-stealth-auth-high-antibot-spa.md`](../../feedback/history/web-scraping/common/2026-07-24_093318-session-first-stealth-auth-high-antibot-spa.md)

## 产出检查清单

- [ ] Envelope 形状（`code`/`success`/`data`）
- [ ] Token 字段名与 Header 格式
- [ ] 分页字段（`page` vs `currentPage` vs `items`）
- [ ] 各模块 list item 字段（从 live JSON，非 TypeScript 猜测）
- [ ] 标注 404/502 未挂载 endpoint
- [ ] 文档 **不含** 密码、token、hash 样本值

## 参考案例

- **console.9j2.cn**（2026-07-02）：Playwright 登录；真实前缀 `/api/user/*`；`/api/admin/*` 不存在 — Brower `docs/analysis/9j2-console/`

## 相关

- [README.md](./README.md) — Web 分析总览  
- [sources-and-tools.md](./sources-and-tools.md) — 工具对照  
- [../../intelligence/web-scraping/anti-bot-bypass.md](../../intelligence/web-scraping/anti-bot-bypass.md) — CAPTCHA / anti-bot  
