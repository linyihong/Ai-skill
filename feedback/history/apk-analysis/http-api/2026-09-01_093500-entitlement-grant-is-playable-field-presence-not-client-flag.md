> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-09-01 - Entitlement grant is playable-field presence, not a client flag

Status: candidate

#### One-line Summary

鎖定內容若仍下發 playlist／key，只靠 `can_play` 或 UI 鎖，屬控制面洩漏；防守應在未授權時**省略**可播欄位。跨 identity 的 playlist 才證明伺服器綁帳，本機 registry 不是。

#### Human Explanation

三類觀察可同時成立：(1) 未解鎖列可播欄位為空，spend 後才出現（伺服器強制）；(2) 業務碼／旗標說不可播，但 JSON 仍帶 URL 與 wrapped key（控制面洩漏，client 必須 fail-closed 不當下載授權）；(3) 列表／詳情已有**完整 source URL**，同一列卻有價格、試看時間或 unlock 提示，而 transport 可直接播放（前端假象）。不要把 `previewURL` 名稱當成證據：它可能為空，真正要追的是 player 最後採用的欄位。解鎖帳號 A 可播、帳號 B 同內容仍無可播欄位，才是 entitlement bind。

#### Trigger

- Parser treats non-empty media URL as downloadable while business code says unlock-first
- UI lock overlay but detail JSON already has playable URL field
- Preview UI exists, but the player actually consumes a full-source field rather than a distinct preview asset
- Local download registry used as proof another identity is entitled

#### Evidence

- Offline fixtures + playlist `hasSrc` reports; SDK `entitledToPlay()` vs `hasPlaylistUrl()`
- Project matrix: `<PROJECT_ROOT>/apk-analysis-sdk/docs/plans/integration/entitlement-layer-comparison.md`

#### Generalized Lesson

```text
Classify entitlement by field presence across identities, not UI:
  empty playable field until spend → server grant
  URL/key present + can_play false → leak; exporter must refuse
  full source present + price/free-time/lock metadata + transport works → ui-only
  preview label or previewURL field alone → inconclusive; trace the player-consumed field
  local registry ≠ server bind; probe the other identity's playlist
```

#### Agent Action

1. SDK parsers expose grant boolean separately from URL presence.
2. HLS export refuses non-entitled PLAY envelopes.
3. Live probes record classification labels only (no URL/token).

#### Goal / Action / Validation

- Goal: 不要把「有 URL」當成「已授權」。
- Action: 比較 empty-field vs flag-with-URL vs UI-only。
- Validation: unit fixtures; optional live isolation probe.

#### Applies When / Does Not Apply When

- Applies: 短劇／HLS／playlist 授權欄位設計與 SDK 匯出。
- Does not apply: 未授權標的的繞過步驟。

#### Promotion Target

- ✅ 觀察方法：`analysis/security/media-entitlement-control-plane.md`
- 第一方設計 lesson：`feedback/history/development-guidance/controls/2026-09-01_095200-media-entitlement-omit-playable-fields.md`
- this lesson remains the apk-analysis classification record

#### Required Linked Updates

- `feedback/history/apk-analysis/http-api/README.md` index
- Step 6 intelligence extraction: **否**（已抽 analysis atom，不另建 intelligence anti-pattern）
- Step 7 failure-learning: **否**（不是 agent 失效，是產品授權語意）
