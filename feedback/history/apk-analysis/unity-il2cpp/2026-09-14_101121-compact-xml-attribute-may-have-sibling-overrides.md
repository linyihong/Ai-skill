> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-09-14 - Compact XML attribute may have sibling entries

Status: validated

#### One-line Summary

XML 中的 compact list attribute 可能只承載主要項目，旁邊的同類 sibling
elements 仍會由官方 parser 合併；不能把 attribute 單獨視為完整集合。

#### Human Explanation

自製 parser 常見捷徑是：只要 compact attribute 解析出任何資料就立即
return。這會讓合法的混合編碼看起來像「封包缺格」或「伺服器漏資料」。
官方 client 的結果 DTO 可能完整，因為它同時讀 attribute 與 sibling
elements，而不是靠畫面層推算缺值。

#### Trigger

- 解密與 XML framing 都正常，但 compact list 偶爾少數項目。
- 官方 post-parse DTO 比同一轉的 compact attribute 多。
- 缺項比例穩定地呈現 partial，而非整包 parse failure。

#### Evidence

- Tool: runtime hook at XML input、nested parser overload entry/return、以及
  post-parse DTO 三個邊界。
- Sanitized excerpt: 一個輸入中 compact attribute 有 `N` 項，sibling
  element overload 回傳 `K` 項，最終 parser 回傳 `N+K` 項。
- Evidence path: project-specific raw captures留在 `<PROJECT_ROOT>`，不放入本 lesson。

#### Generalized Lesson

解析同一 container 時，應收集所有 schema-valid representation，再以穩定
identity（例如座標或 id）合併與檢查衝突。不可因 compact attribute 非空就
跳過 sibling elements。

#### Agent Action

1. 同一事件同步 hook XML boundary 與 post-parse DTO。
2. hook compact-item parser 與 sibling-element parser 的回傳數量。
3. 若 `attribute count + sibling count = DTO count`，先實作兩來源合併，
   不要用鄰格、歷史 frame 或 UI 推測補值。
4. 合併必須保留事件／玩家 scope；重複 identity 值衝突時維持 partial
   並記錄診斷。

#### Goal / Action / Validation

- Goal: 判定 partial 是資料遺失，還是自製 parser 漏讀混合表示法。
- Action: 三邊界同轉計數，並只記 schema 與必要 identity。
- Validation or reference source: 完整 DTO 數必須等於 compact 與 sibling
  去重合併後的數量；加入 compact-only、sibling-only、mixed、duplicate
  conflict 測試。

#### Applies When

- XML 或類 XML protocol 允許 compact attributes 與 nested/sibling entries。
- 可觀察官方 parser 或可信 post-parse DTO。

#### Does Not Apply When

- Schema 明確保證 attribute 是唯一且完整的 representation。
- sibling elements 屬於不同事件、不同玩家或不同 frame。

#### Validation

以同一 container 建立測試：compact list 少一項、sibling element 提供該
項，合併後應完整；若兩來源同 identity 不同值，測試應拒絕靜默覆寫。

#### Promotion Target

- `analysis/apk-analysis/`（累積更多跨 target 證據後再 promotion）

#### Required Linked Updates

- 已更新 `feedback/history/apk-analysis/unity-il2cpp/README.md` 與
  `feedback/history/apk-analysis/README.md`。
- 已依 reusable-guidance boundary 去除 target、class、payload、裝置與路徑細節。
