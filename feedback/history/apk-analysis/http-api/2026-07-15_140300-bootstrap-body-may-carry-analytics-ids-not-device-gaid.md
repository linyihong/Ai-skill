> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-15 - Bootstrap body may carry analytics IDs separate from device GAID

Status: candidate

#### One-line Summary

Bootstrap／冷啟 POST body 可能含 **`anonymousId`／`distinctId`**（analytics SDK）與 **`firstInstall`**（prefs），與 header 的 **GAID／androidId** 分離——合成 identity 時勿只造 GAID。

#### Human Explanation

jadx 常見 `RequestApiLib` 組 HashMap：`SensorsDataAPI.getAnonymousId/getDistinctId` + `AppUtils.getOSInfo/getModel` + `SpData.isFirstInstall()`。離線 SDK 若只合成 `deviceId` 而漏 analytics 欄位，bootstrap 可能仍 200 或 silently 風控。`anonymousId` 規則常 **不等於** GAID hex。

#### Trigger

- Bootstrap `@Body HashMap` with analytics field names in jadx
- Live `Request.toString` shows distinctId == anonymousId on first install
- Synthetic profile missing analytics fields

#### Evidence

- Tool: jadx static + Frida `RequestApiLib` hook
- Sanitized excerpt: five-field body; analytics IDs 16hex class; firstInstall boolean
- Evidence path: `<PROJECT_ROOT>/api/dynamic-w22b-*.md`, `scripts/headless/gn_bootstrap_body.py`

#### Generalized Lesson

```text
Bootstrap body RE:
  1. Static: find boot/start API builder method (HashMap put chain)
  2. Classify fields: analytics IDs | device meta | install flag
  3. SDK: separate bootstrap_body() from device header identity
  4. Open: synthetic analytics ID generation rules (project evidence)
  5. Sign: still path B empty bodyStr if PostJsonBody
```

#### Agent Action

1. 專案 helper 明確列出 bootstrap body schema table。
2. 與 W13 device identity doc 分檔，避免混欄位。
3. Ai-skill 不寫真實 ID 值。

#### Goal / Action / Validation

- Goal: 完整還原 bootstrap wire body。
- Validation: static jadx fields == live Frida MAP window。

#### Applies When

- Cold-start bootstrap with analytics SDK on classpath
- Retrofit HashMap body on boot endpoint

#### Does Not Apply When

- Bootstrap is empty body or query-only

#### Promotion Target

- `analysis/apk/workflows/http-api-documentation-flow.md` §request body

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
