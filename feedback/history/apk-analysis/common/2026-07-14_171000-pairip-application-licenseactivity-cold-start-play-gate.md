> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-14 - Pairip Application + LicenseActivity as cold-start Play gate

Status: validated

#### One-line Summary

側載後一開 App 就被導向 Play／要登入時，優先用**靜態 triage**對照：manifest `Application` 是否為 Pairip wrapper、是否註冊 Pairip `LicenseActivity` + Play Store URL、`CHECK_LICENSE`／Billing、installer 歸屬 API，並用裝置 `installer=` 交叉驗證——**不要**先假設是 Splash 業務登入。

#### Human Explanation

部分 Play 發行建置會把 `android:name` 設成 Pairip 的 `Application`，並註冊 Pairip `LicenseActivity`（常伴隨 `https://play.google.com/store`）。授權失敗時，授權／商店 remediation 可能發生在業務 Splash **之前**。另常見搭配：`CHECK_LICENSE` permission、BillingClient、`getInstallerPackageName`／`getInstallSourceInfo`。側載安裝在裝置上常顯示 `installer=null`（Play 安裝則為 `com.android.vending`）。這條鏈路解釋「為什麼還沒進首頁就被丟去商店」；靜態對照即可定位，無需 runtime 繞過。

#### Trigger

- 冷啟動立刻進 Play Store / Google 帳號 UI，業務 Splash 幾乎看不到。
- Manifest 出現 `com.pairip.application.Application` 或 `com.pairip.licensecheck.LicenseActivity`。
- Artifact 來自第三方 XAPK／側載，`pm list packages -i` 顯示 `installer=null`。

#### Evidence

- Tool: `aapt dump badging`、`aapt dump xmltree … AndroidManifest.xml`、DEX marker 計數、可選 `adb shell pm list packages -i`。
- Sanitized excerpt: manifest flags `pairip_application=true`、`pairip_license_activity=true`、`play_store_url=true`；device `installer=null` 與側載一致；腳本 exit 0。
- Evidence path: 目標專案 docs／scripts（專有名、package、本機路徑留專案；本 lesson 不寫）。

#### Generalized Lesson

1. **Cold-start Play-gate triage（在 traffic／Frida 之前）**  
   - 讀 launchable-activity **與** `application android:name`。  
   - 若 Application 是 Pairip wrapper，把「授權門檻」排在 Splash 業務邏輯之前。  
   - 搜尋 Pairip `LicenseActivity`、Play Store URL、`CHECK_LICENSE`、Billing / Integrity / installer API。
2. **Device correlation**：同 package 查 `installer=`；`null` ≠ Play 歸因，對照静态门闸比猜 Splash login 更快。
3. **腳本角色**：自動化上述 marker 報告即可；**禁止**把「驗證腳本跑通」定義成 license bypass。可重用技巧是 triage，不是繞過。
4. **防守（自己的 App）**：客戶端導流只是 UX； entitlement 與內容 API 應由後端依 Play Integrity `appLicensingVerdict`（等）裁決。

#### Agent Action

1. 出現冷啟動商店導流 → 先跑 Play-gate static triage（aapt + DEX markers + 可選 installer）。  
2. 將鏈路寫進目標專案 docs（Pairip → LicenseActivity → Splash），再決定 capture window。  
3. 寫 Ai-skill lesson 前去敏：禁止真實 package、裝置 serial、本機絕對路徑、bypass PoC。

#### Goal / Action / Validation

- Goal: 正確定位冷啟動 Play／授權門檻，避免誤判為一般登入頁。  
- Action: 靜態 marker 對照 + installer 交叉驗證；專案內保存 map 與 triage script。  
- Validation: triage script exit 0 且 flags／installer 與「側載→商店導流」現象一致（本輪已驗證）。

#### Applies When

- 側載／第三方 artifact 冷啟動異常導向 Play。  
- Manifest 顯示 Pairip Application / LicenseActivity。  

#### Does Not Apply When

- 僅業務 LoginActivity（無 Pairip／CHECK_LICENSE／商店 remediation）。  
- 需要的是授權繞過或重打包——不在本 lesson 範圍，且不應寫入 skill。

#### Promotion Target

- `analysis/apk/traffic-triage.md` 或 cold-start checklist：在 network triage 前加 Play-gate static markers（可選，後續）。  
- `workflow/apk-analysis/execution-flow.md` §1：可引用「授權／installer gate 先於 Splash」。

#### Required Linked Updates

- N/A（本輪僅 history lesson；promotion 未改 workflow 正文）。

#### Confidence / Residual Risk

- Confidence: high（manifest flags + device installer + script assert）。  
- Residual: Pairip 版本差異可能改 class 名；仍應以 aapt/xmltree 實讀為準。
