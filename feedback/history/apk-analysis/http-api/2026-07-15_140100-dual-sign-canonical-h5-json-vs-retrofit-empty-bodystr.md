> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-15 - Dual sign canonical paths: H5 JSON body vs Retrofit PostJsonBody empty bodyStr

Status: candidate

#### One-line Summary

同一 App 可能同時存在 **兩條簽名 canonical**：H5 任務 API 把 **JSON body 字串** 納入 canonical；Retrofit `@Body PostJsonBody` 路徑在簽名時 **`bodyStr=""`**，但 wire 仍送 JSON——混用會 `Invalid sign`。

#### Human Explanation

離線 signer 常只還原一條路。實務觀測：check-in／award 走 WebView H5 helper（path A），bootstrap／read upload 走 OkHttp Retrofit interceptor（path B）。Path B 的 Gson 序列化 body **不進** canonical `bodyStr`；Path A 的 `getH5HeaderData(bodyJson, timestamp)` 把 body 當簽名材料。手動 `HttpURLConnection` 用錯 path 會簽名長度對但 server 拒絕。

#### Trigger

- Offline sign length matches in-app FP，但 replay 回 `Invalid sign`
- Same app: task/center/* vs app/bootstrap use different sign helpers
- jadx: `PostJsonBody` + empty body in interceptor vs H5 `getNewSign` bridge

#### Evidence

- Tool: Frida hook both sign entry points; compare canonical first arg
- Sanitized excerpt: path A canonical includes JSON body; path B `bodyStr=""` with non-empty wire JSON
- Evidence path: `<PROJECT_ROOT>/api/dynamic-w7.md`, `dynamic-w19-*.md`

#### Generalized Lesson

```text
Dual-path sign triage:
  1. Catalog endpoints by sign helper class (H5 bridge vs Retrofit interceptor)
  2. For each path, record: wire body shape + canonical bodyStr rule
  3. Host SDK: gn_h5_sign vs gn_retrofit_sign (or equivalent) — never mix
  4. Device replay: prefer app's own header builder per path
  5. Validation matrix: ≥1 path A + ≥1 path B replay on device
```

#### Agent Action

1. 專案 roadmap 明確標 A/B 路徑與端點歸屬。
2. 與 `141800`（sorted-map canonical）、`133500`（vendored HttpGlobal）並讀。
3. Replay 腳本按 path 選 signer module。

#### Goal / Action / Validation

- Goal: 避免單一路徑 signer 誤覆蓋全 App API。
- Validation: bootstrap (B) + check-in (A) both `status=0` on device executor。

#### Applies When

- Hybrid native REST + embedded H5 rewards
- Retrofit PostJsonBody coexists with H5 signed XHR

#### Does Not Apply When

- Single global sign canonical for all endpoints (verify first)

#### Promotion Target

- `analysis/apk/workflows/http-api-documentation-flow.md` §sign validation
- `analysis/apk/workflows/headless-sdk-device-executor-flow.md`

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md`
- `analysis/apk/workflows/http-api-documentation-flow.md`
