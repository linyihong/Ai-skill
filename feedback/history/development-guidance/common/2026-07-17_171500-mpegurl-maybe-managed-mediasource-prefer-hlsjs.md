> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-17 - mpegurl "maybe" + ManagedMediaSource prefer hls.js

Status: validated

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
