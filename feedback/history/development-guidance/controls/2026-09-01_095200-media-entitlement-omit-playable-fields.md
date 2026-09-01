> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-09-01 - 媒體授權先省略可播欄位，再發短效憑證

Status: candidate

#### One-line Summary

付費播放的控制面是未授權時**不發** playlist／key；短效媒體憑證縮小重放窗口，不能取代省略，也不能單獨擋住已授權持有者。

#### Human Explanation

用戶端旗標、鎖 overlay、請求簽名都不是授權。正確第一層是伺服器對未授權 identity 省略可播欄位。第二層才是短 TTL、可撤銷、盡量綁 session 的播放憑證。TTL 到期後必須重打授權 API。分段 license 更強，但成本較高。試看必須是獨立短資產。

#### Trigger

- 設計或審查付費播放、membership、HLS／DASH 發放
- `security-audit` 觸及 playlist 或 media URL
- SDK 準備把播放 URL 當下載授權

#### Evidence

- Sanitized excerpt: 三類控制面（省略至授權後才出現、拒絕旗標仍帶 URL、列表已有完整 source）已在 apk-analysis http-api lesson 抽象
- Evidence path: 專案比對矩陣留 `<PROJECT_ROOT>` 文件，不進本檔

#### Generalized Lesson

```text
Layer 1: omit playable fields until entitled (fail-closed)
Layer 2: short-lived, revocable, session-bound media credentials
Layer 3: per-segment license when cost allows
Preview asset ≠ full-source URL
TTL ≠ reuse-proof; it only shrinks the replay window
Client flags / overlays / request wrappers are not the grant
```

#### Agent Action

1. 第一方實作：未授權 envelope 不含可播 URL／key。
2. 發放後使用短 TTL 重簽；不要把長壽命 CDN URL 當授權。
3. SDK 匯出：grant boolean 與 URL 存在分開；非 grant 拒匯出。
4. 未授權第三方標的：只做欄位存在分類，不寫取得步驟。

#### Goal / Action / Validation

- Goal: 把媒體授權設計成伺服器 fail-closed，而不是 UI 鎖。
- Action: 省略欄位 → 短效憑證 →（可選）分段 license。
- Validation: 未授權 identity 的 playlist 為空；授權後 URL 有過期；預覽與正片不同資產。

#### Applies When / Does Not Apply When

- Applies: 第一方付費媒體、SDK 播放匯出、`security-audit` 審查播放契約。
- Does not apply: 未授權標的的取得步驟；公開免登入媒體。

#### Promotion Target

- ✅ `analysis/security/media-entitlement-control-plane.md`
- `workflow/software-delivery/` 載入點（README + validation + security-audit invoke）
- 不新建 `workflow/security/` domain

#### Required Linked Updates

- `feedback/history/development-guidance/README.md` 與 `controls/README.md`
- `analysis/security/README.md`、`analysis/README.md`
- `knowledge/summaries/` atom
- apk-analysis 對應分類 lesson 的 Promotion Target
- Step 6 intelligence extraction: **否**（先留 analysis atom）
- Step 7 failure-learning: **否**（產品控制面，不是 agent 失效）
