# Wire Path vs Signing Canonical Path

Status: candidate

## Context

在分析加密／簽章 API 時，live request 可能同時存在兩個不同的 path 概念：

- 實際 HTTP wire URL path：request 送到 server 的 path。
- 簽章 canonical path：App 內部餵給 `eh` / signature generator 的 path material。

這兩者不一定完全相同。<target-app> gossip live smoke 中，HTTP wire path 使用 `/v1/api/public/` 才會回到 encrypted JSON；但 `eh` 產生仍使用 App 內部 `api/public/?...` canonical material。若把兩者混成同一個值，server 可能回 HTML/error page，表面上像是授權、簽章或 decrypt 缺口。

## Rule

遇到 encrypted API 回 HTML 或非 JSON 時，不要先歸咎於 identity/signing/decrypt 缺失。先分別驗證：

1. wire URL path 是否與實際 app request family 一致；
2. signing canonical path 是否與 App crypto helper 取用的 material 一致；
3. 兩者是否需要分開建模與文件化。

## Evidence

- <target-app> guest login 文件與既有實作使用 `/v1/api/public/` 作為 wire path。
- 既有 `DartEncryptAESProvider` / `GuestLoginClient` 註解顯示 `eh` signing path 使用 `api/public/`，不是 `v1/api/public/`。
- Gossip live smoke 從 `/api/public` 改為 `/v1/api/public/` 後，categories/articles/detail 由 HTML failure 進入可解析 response，並能下載 detail content image。

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
