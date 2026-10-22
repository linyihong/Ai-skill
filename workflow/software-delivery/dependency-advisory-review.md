# 相依元件公告查核

本流程補充 `security-audit`，供已授權專案的 CVE 詢問、相依元件升級、資安審查及發布檢查使用。這是 agent 讀取的文件流程，尚無專用 CI 或機械阻擋器；不能宣稱已自動持續監控。

## 觸發與時效

- 使用者詢問近期漏洞、變更 SDK／runtime／直接或間接相依、基底映像或部署環境，以及每次發布前，必須執行查核。
- 一般不涉及相依或部署的修改，可引用仍有效的既有證據，記錄引用理由；不能把每次 session 啟動都變成完整掃描。
- 記錄 `checked_at`、來源及公告日期。初次盤點對照所有已知且適用的公告；後續查核從上次成功時間起補查，並保留未解決項目，不限「最近 30 天」。
- 發布前必須重新取得當次公告與相依掃描結果。新公告、版本／lockfile／映像／OS／功能或代理設定變更，立即使相關舊證據失效。
- 網路、資料源、restore 或工具失敗記為 `partial`／`unknown`，不能把空輸出、零 finding 或命令成功退出當成安全證明。

## 查核順序

1. **盤點實際組成**：SDK、shared runtime、目標框架、直接與間接套件、lockfile、native library、容器基底／OS，以及正式部署模式。版本宣告、解析版本、執行版本分開記錄；SDK 版本不等於 runtime 版本。
2. **取得官方公告**：以供應商安全公告、維護者 advisory 與官方 release metadata 為主要依據，交叉核對 CVE／GHSA、受影響範圍、修補版、OS／架構及必要功能。新聞只用來發現線索；同一 CVE 的多個識別碼去重。
3. **執行生態系掃描**：包含間接相依，保存命令、工具版本、資料源、時間、退出碼、涵蓋專案和未涵蓋項目。重新整理或還原相依；僅用快取時註明時效限制。掃描通過不涵蓋所有 shared runtime 或 OS 公告。
4. **逐項判定適用性**：分開記錄官方嚴重度、版本是否命中、元件是否存在、必要功能與外部可達性。可用 `affected`、`not_affected`、`unknown`；`not_affected` 需反證及失效條件，不能只因沒搜尋到一個 symbol 就排除。
5. **決定處置**：已知遭利用或可達的 Critical／High 優先隔離、修補並升級處理；不等待一般發布。其他已命中版本亦排修補。無修補版時記錄暫時控制、owner、期限及重驗條件。延後不等於接受風險；接受風險須有具名決策及有效期。
6. **驗證修補**：最小相容版本變更、restore／build、受影響路徑回歸與掃描重跑。runtime／容器修補需更新部署產物，重啟並讀回實際版本；只改 manifest 或開發機安裝不足以關閉。

## .NET 查核面

```sh
dotnet --info
dotnet --list-runtimes
dotnet list <solution-or-project> package --vulnerable --include-transitive --format json
```

依實際 SDK 確認命令支援；不要假設掃描 finding 會讓 CLI 回傳非零碼。若接 CI，必須解析報告並分辨 finding、掃描失敗與未覆蓋專案。

- 比對 `global.json`、`TargetFramework`、套件解析結果與官方 .NET／ASP.NET Core servicing 公告。
- framework-dependent：修補每個部署主機的相容 shared runtime／Hosting Bundle，重啟並讀回。
- self-contained／single-file／Native AOT：核對實際封裝模式，更新建置工具或 runtime pack，重新 publish、部署及驗證；只更新主機 runtime 不能取代。
- 容器：重建使用已修補基底的映像，核對 digest 與部署實際映像；不能只更新相同 tag 的文字。
- NuGet 套件升級與 runtime 修補分開驗證；不能用 `netX.0` 或所有套件尾碼一致來推定安全。

## 證據與關閉

在專案 repo 保存查核日期、組成清單、官方來源、CVE 判定表、掃描結果、未涵蓋項目及修補／部署讀回。敏感 host、credentials、套件源認證與原始私人資料不得進共用文件。

輸出沿用 [`security-finding-list-template.md`](templates/security-finding-list-template.md)：官方嚴重度不能直接冒充專案 confirmed severity；缺部署或可達性證據時保留 `needs_validation` 與未決事實。正式發現的獨立覆核與關閉依該模板。

發布判定：查核失敗、範圍內高風險狀態未明、或未解決的適用 Critical／High，不能宣稱資安查核通過；依既有 security gate 補驗證、修補或具名風險決策。低嚴重度項目亦需 owner、期限與追蹤位置。

延伸：[覆蓋帳本](../../analysis/security/security-coverage-ledger.md)、[發布入口](release/README.md)。
