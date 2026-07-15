# Headless SDK Device Executor Flow（裝置執行器流程）

`analysis/apk/workflows/headless-sdk-device-executor-flow.md` 描述當 **離線 signer 已對齊** 但 **主機直連被 anti-bot／WAF 擋（HTTP 403）** 時，如何用 **裝置端 Frida executor** 完成 private SDK 的協議驗證與日常執行。

> **Intelligence Extracted**
> See:
> - `feedback/history/development-guidance/common/2026-07-15_140200-correct-sign-length-still-403-means-anti-bot-not-sign.md`
> - `feedback/history/apk-analysis/http-api/2026-07-15_140100-dual-sign-canonical-h5-json-vs-retrofit-empty-bodystr.md`
> - `feedback/history/apk-analysis/common/2026-07-15_132500-frida-e2e-compact-send-and-off-main-thread-http.md`

## 適用條件

| 條件 | 說明 |
| --- | --- |
| Sign RE Done | 離線 `sign_len` + key fingerprint 與 in-app 一致 |
| Host blocked | 桌面／JVM 直連 → 403 HTML（非 JSON `Invalid sign`） |
| Device works | 同 payload 經 in-app HTTP → HTTP 200 |
| Goal | Private SDK：identity → bootstrap → rewards，不依賴長期 session 匯出 |

## 架構

```text
gn_identity.new_profile()          # 合成 GAID + androidId（專案）
  → path B: POST bootstrap         # Retrofit sign, bodyStr=""
  → GnSession { token, uid }       # 專案 session 物件
  → path A: task/center/*          # H5 getH5HeaderData sign
  → capture/sdk/*.env (gitignore)
```

App／Frida **僅協議發現**；production runner = **device executor**（adb + Frida spawn）。

## 步驟 1 — 403 分級（勿再改 sign）

```text
IF sign-only parity OK
  AND host 403 edgesuite/challenge HTML
  AND device replay 200
THEN anti-bot tier — pivot here
ELSE continue sign RE or schema RE
```

## 步驟 2 — 雙簽名路徑模組化

| Path | 典型 API | Canonical rule | Host module |
| --- | --- | --- | --- |
| **A** H5 | check-in, award/receive | JSON body in canonical | `gn_h5_sign` |
| **B** Retrofit | bootstrap, read upload | `bodyStr=""`, wire JSON separate | `gn_retrofit_sign` |

**禁止** 用 path A signer 打 path B 端點（或反之）。

## 步驟 3 — 冷啟 bootstrap E2E

1. `pm clear` + Frida spawn（乾淨遊客態）。
2. 等 token ready：`HttpGlobal` Authorization setter 或 `SpData.setUserToken`（見 `140400` lesson）。
3. 只 log `tokenLen`／`uid`，不 `send()` raw token。
4. Bootstrap body：analytics IDs + `firstInstall` + device meta（見 `140300`）。

## 步驟 4 — 下游 API（裝置端）

每請求 fresh app header helper（如 `getH5HeaderData(body, timestamp)`），刪 `localTime` 若 H5 同族要求。

- HTTP 在 **Frida 執行緒**（`setTimeout`），非 Android main thread。
- `send()` 只送 `{path, http, status}` 精簡 steps。
- Host：`run_*_e2e.py` spawn wrapper + regex fallback parser + 固定 sleep。

## 步驟 5 — Session 持久化（gitignore）

| 產物 | 用途 |
| --- | --- |
| `GnSession.from_frida_export()` | 合併 profile + token |
| `save_env` / `save_json` | 除錯；非長期架構依賴 |
| `to_probe_session()` | 接主機 probe（仍可能 403） |

## 步驟 6 — Mac headless（可選、常 blocked）

主機 `probe_*.py` 保留作 sign 驗證與未來 Akamai 繞過實驗；**不以** 其成功作為 SDK Done 門檻。

## 成功產出

```text
bootstrap: tokenLen>0, uid present
downstream: task/center/index + sign → http=200, status=0
session: capture/sdk/latest_session.env (gitignore)
docs: project dynamic-w22* + roadmap Phase marks
```

## 可重用 lesson 索引

| Slug | 主題 |
| --- | --- |
| `140000` | Bootstrap not login |
| `140100` | Dual sign A/B |
| `140200` | Sign OK but 403 = anti-bot |
| `140300` | Bootstrap analytics IDs |
| `140400` | HttpGlobal token hook |
| `132500` | Compact send + off-main-thread HTTP |

---

← [回到 analysis/apk/workflows/](README.md)
