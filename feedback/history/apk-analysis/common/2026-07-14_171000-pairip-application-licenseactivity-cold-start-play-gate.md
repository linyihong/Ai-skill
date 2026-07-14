> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-07-14 - Cold-start Play handoff: static markers ≠ observed hop (revised)

Status: validated

#### Revision (same day)

上一版把 Pairip `Application` / `LicenseActivity` 靜態命中寫成「已證實的冷啟動中間 hop」，並暗示授權過後即可進 Splash／主頁。**這是假陽性過度歸因。** 同日 adb 實測：`Splash` → 短暫業務 `Login` → `com.android.vending` `UnauthenticatedMainActivity`；**未**採樣到 Pairip `LicenseActivity`，也**進不了**主頁／閱讀頁。修正後：靜態 marker 只當**候選信號**；可驗證結論必須是 focus／截圖採樣。

#### One-line Summary

側載冷啟動被丟去 Play 時：靜態可列 Pairip／`CHECK_LICENSE`／installer 等**候選**；真正成立的技巧是 **adb focus 時間序列 + 截圖** 證明 Play handoff，且**不得**把「靜態 script exit 0」當成已證明 Pairip hop 或「可進主頁」。

#### Human Explanation

靜態常見：`com.pairip.application.Application`、`com.pairip.licensecheck.LicenseActivity`、Play Store URL、`CHECK_LICENSE`、`installer=null`。這些提高「可能有 Play／授權門檻」的先驗，但**不等于** runtime 一定會經過 LicenseActivity，也不等于绕过后再进业务页。實測可能先出現業務 Splash／Login，再被 `vending` 的未登入頁搶焦點。分析與回饋必須分開三層證據：（1）靜態候選；（2）runtime 觀測 hop；（3）是否達到目標 Activity（Home／feature）。缺（2）（3）就宣称「鏈路已證實」會產出假技巧。

#### Trigger

- 冷啟動後進 Play／Google 帳號 UI。  
- 已寫或即將寫「Pairip → LicenseActivity → Splash」之類**未採樣**鏈路。  
- 只用靜態 triage script 的 exit 0 當 validation。

#### Evidence

- Tool: `am force-stop` + `am start -W` launchable；週期讀 `dumpsys window` `mCurrentFocus`；`screencap`；靜態 aapt／DEX triage 僅作平行記錄。  
- Sanitized excerpt: Launch Splash ok → ~sub-second 業務 Login → 穩定 `vending`／Unauthenticated*；主 feature Activity 未達；Pairip LicenseActivity 未進 focus 樣本。  
- Evidence path: 目標專案 docs／capture（專有名留專案）。

#### Generalized Lesson

1. **三層證據（強制）**  
   - L1 靜態 marker（Pairip／CHECK_LICENSE／installer API）= **hypothesis**。  
   - L2 adb focus 時間序列（+截圖）= **observed hops**。  
   - L3 目標 Activity／package 仍屬業務且可操作 = **feature reachability**。  
2. **禁止**：用 L1 alone 寫「validated startup chain」或寫進 skill 當已證實 Pairip hop。  
3. **禁止**：把「無法進主頁」的觀測偷換成「繞過後進主頁」的技巧；本 lesson **不**教 bypass。  
4. **可驗證句型**：「側載冷啟動後 focus 落在 Play Unauthenticated*，無法達 Home／feature」——這句可由 adb 複核；「因此 LicenseActivity 是中間頁」——不可由未採樣推斷。  
5. **腳本**：static triage 標題必須標 triage／hypothesis；runtime probe 另做。

#### Agent Action

1. 有靜態 Pairip／Play marker → 記錄為候選，**立刻**做 focus 採樣。  
2. 專案 docs 分欄：Static candidates | Observed focus sequence | Unreached activities。  
3. 修正或標記既有過度歸因 lesson；寫 feedback 前對照 L2／L3。  
4. 使用者要求「改到能進去」而屬授權／DRM 繞過 → 拒絕；改建議 Play 正式安裝或自家測試建置。

#### Goal / Action / Validation

- Goal: 避免假陽性 startup-chain lesson；保留可複核的 Play-handoff 觀測。  
- Action: 修正本條；專案 docs 對齊 runtime；靜態 script 不加 bypass。  
- Validation: 專案 runtime 表與本條 L1/L2/L3 一致；lesson 不再宣稱未採樣的 Pairip hop。

#### Applies When

- 任何「靜態看到授權／Pairip → 斷言 runtime 鏈路」的場面。  
- 側載後被 Play 未登入頁擋住。

#### Does Not Apply When

- 已有穩定 focus 採樣證明某一中間 Activity。  
- 需要的是授權繞過以進主頁（不在 skill 範圍）。

#### Promotion Target

- `workflow/apk-analysis/execution-flow.md` § cold start：證據分層（靜態 / focus / feature）— 可選後續。  
- 本條取代同日過度歸因敘事；索引標題已反映 revised。

#### Required Linked Updates

- `feedback/history/apk-analysis/README.md` 索引说明（本輪更新）。

#### Confidence / Residual Risk

- Confidence: high on「Play handoff blocks feature reachability when L2 shows vending Unauthenticated*」。  
- Residual: Pairip LicenseActivity 可能存在極短 hop 未被 400ms 採樣抓到——只能標 unobserved，不可當已證明。
