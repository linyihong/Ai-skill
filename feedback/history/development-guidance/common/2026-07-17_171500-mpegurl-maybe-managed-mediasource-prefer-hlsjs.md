> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-17 - mpegurl "maybe" + ManagedMediaSource prefer hls.js

Status: candidate

#### One-line Summary

對 Apple WebKit，`canPlayType("application/vnd.apple.mpegurl") === "maybe"` 且存在 `ManagedMediaSource` 時，應優先走 hls.js / blob playlist，不要為了模擬器方便把 `"maybe"` 提升成 native `<video src=*.m3u8>`。

#### Human Explanation

現代 iPhone Safari 常回報 HLS 能力為 `"maybe"`，同時暴露 `ManagedMediaSource`。已知可播的 production 路徑在此條件下走 hls.js（`blob:` playlist），畫面可解。若 agent 為了 iOS Simulator（也常回 `"maybe"`）把 `"maybe"` 當成 native preference，真機會改走 CDN / native `<video src>`，出現 `readyState` 正常、時間前進、但 `videoWidth=0`（有聲無畫面）。

模擬器常缺 HEVC 硬解，attach 成功也不能當 paint oracle。驗證必須用實體裝置量 `videoWidth/Height`，或以 production 同機 A/B 對照。

#### Trigger

- 修改 immersive / HLS player 的 native vs hls.js 分支
- `canPlayType("…mpegurl")` 在目標裝置為 `"maybe"`
- `typeof ManagedMediaSource !== "undefined"`
- 回歸症狀：封面卡住、有聲音無畫面、`videoWidth===0`
- 僅用 Simulator / attach (`readyState`、duration) 宣告修復成功

#### Evidence

- Tool: physical-device Safari remote automation + production A/B
- Sanitized excerpt: Device reported mpegurl `"maybe"` and `ManagedMediaSource`. Production painted via `blob:`; test host forced native CDN src and stayed at `videoWidth=0` until native preference returned to `"probably"`-only.
- Evidence path: Project-specific episode IDs, hosts, and screenshots stay in `<PROJECT_ROOT>` project feedback / evidence; this lesson only records the capability-branch rule.

#### Generalized Lesson

| Capability signal | Prefer |
| --- | --- |
| mpegurl `"probably"` | Native `<video src>` may be OK |
| mpegurl `"maybe"` + ManagedMediaSource | hls.js / blob (or same path as known-good production) |
| Simulator attach green, `videoWidth=0` | Not a paint pass — need physical device or non-HEVC fixture |

Do not widen `"maybe"` → native solely to make Simulator attach tests green.

#### Validation

- Same physical device: production vs candidate build, compare `currentSrc` shape (`blob:` vs CDN) and `videoWidth`
- Contract/unit markers: native preference gated on `"probably"` only
- Optional: Simulator may still assert attach, but must not be the sole close signal for paint

#### Related

- Failure pattern: [`enforcement/failure-patterns/attach-or-simulator-green-as-paint-oracle.md`](../../../enforcement/failure-patterns/attach-or-simulator-green-as-paint-oracle.md)
- Sibling lesson: [`2026-06-08_141100-blob-manifest-uri-rewrite-test.md`](2026-06-08_141100-blob-manifest-uri-rewrite-test.md)

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
