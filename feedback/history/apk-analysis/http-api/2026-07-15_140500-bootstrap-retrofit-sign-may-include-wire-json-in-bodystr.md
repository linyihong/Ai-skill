> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-15 - Bootstrap Retrofit sign may include wire JSON in bodyStr (not empty)

Status: candidate

#### One-line Summary

冷啟 `app/bootstrap` 的 Retrofit path B 簽名 canonical 可能把 **HTTP body 同份 compact Gson JSON** 放入 `bodyStr`；jadx `PostJsonBody → bodyStr=""` 對 bootstrap **不一定成立**——需 Frida `c.a.a` 原文驗證。

#### Human Explanation

Lesson `140100` 正確描述 H5 vs Retrofit 雙路徑，但 bootstrap 冷啟在裝置上實測：`canonical = timestamp + wireJson + deviceId + androidId + "" + apkMd5 + pkg`。Gson 欄位順序可能與 HashMap 插入順序不同（例：`anonymousId,distinctId,os,sysModel,firstInstall`）。離線 signer 若用空 `bodyStr` 會 `canonical_match: false` 或 host `Invalid sign`——即使後續還可能缺 headers（見 `140600`）。

#### Trigger

- Bootstrap host replay `Invalid sign` after read/upload path already green
- jadx shows PostJsonBody empty bodyStr but bootstrap still fails offline compare
- Frida signer canonical contains `"anonymousId"` JSON between millis and device UUID

#### Evidence

- Tool: Frida hook `com.newreading.net.c.a.a(String)` on `pm clear` cold start
- Sanitized excerpt: `hasJson=true`, wire bootstrap fields in canonical, `signLen=344`
- Offline: `gn_bootstrap_canon_compare.py` → `canonical_match: true` when bodyStr=wire
- Evidence path: `<PROJECT_ROOT>/api/dynamic-w27-host-bootstrap-invalid-sign-analysis.md`

#### Generalized Lesson

```text
Bootstrap sign RE:
  1. Do NOT assume PostJsonBody empty bodyStr for all Retrofit endpoints
  2. Capture cold-start bootstrap signer canonical on device (spawn + pm clear)
  3. Record Gson wire field ORDER separately from HashMap schema
  4. Offline compare: device scrub vs retrofit_canonical before host replay
  5. If canonical matches but host still fails → triage headers (140600) before server-fraud hypothesis
```

#### Agent Action

1. 專案 bootstrap signer 模組：`bodyStr = gson_wire_json`（bootstrap 專用）。
2. 交叉更新 `140100`：bootstrap 為 path B 例外子類，非推翻雙路徑模型。
3. Replay 腳本分 endpoint 記錄 bodyStr 規則表。

#### Goal / Action / Validation

- Goal: 避免 bootstrap 誤用 read/upload 或 jadx 泛化規則。
- Validation: device canonical vs offline `canonical_match: true` for bootstrap hit #1。

#### Applies When

- Guest cold bootstrap via Retrofit
- PostJsonBody annotation present but signer hook shows JSON in canonical

#### Does Not Apply When

- H5 task API (path A) — body always in canonical
- Warm authenticated Retrofit calls using `{}` empty JSON bodyStr (separate hits)

#### Promotion Target

- `analysis/apk/workflows/http-api-documentation-flow.md` §sign validation
- `analysis/apk/workflows/headless-sdk-device-executor-flow.md`

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
- Cross-ref in `http-api/2026-07-15_140100-dual-sign-canonical-h5-json-vs-retrofit-empty-bodystr.md` (bootstrap exception note)
