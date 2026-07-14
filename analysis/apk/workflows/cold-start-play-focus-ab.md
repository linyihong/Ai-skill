# Cold-start Play focus A/B（HOW TO DO）

`analysis/apk/workflows/cold-start-play-focus-ab.md`

側載／商店未登入裝置上，冷啟動常被誤判為「App 起不來」。本流程用 **adb focus 時間序列** 做 A/B 對照：允許 Play 搶焦 vs 抑止 `com.android.vending` 搶焦。

> **Intelligence**
> - `intelligence/engineering/analytical-reasoning/heuristics/play-focus-steal-vs-hard-kill.md`
> - Feedback: `feedback/history/apk-analysis/common/2026-07-14_171000-pairip-application-licenseactivity-cold-start-play-gate.md`

**範圍：** 觀測與截圖。**不是** Pairip／LVL／簽名繞過；禁止改包與授權 hook。

## 前置

- `adb` 已連裝置；目標 package 已安裝。
- 知道 launchable activity（`aapt dump badging`）。
- 記錄：`pm list packages -i` 的 `installer=`、`dumpsys account` 的 `Accounts: N`。

## 靜態候選（L1，可選）

```bash
aapt dump badging <base.apk> | grep -E "package:|launchable-activity|CHECK_LICENSE|BILLING"
aapt dump xmltree <base.apk> AndroidManifest.xml | grep -iE 'pairip|CHECK_LICENSE|LicenseActivity|play.google.com/store'
```

L1 只產生假說（例如 Pairip Application／LicenseActivity）。**不得**寫成已證實 runtime hop。

## Condition A — 允許 Play

```bash
adb shell am force-stop <package>
adb shell am force-stop com.android.vending
sleep 1
adb shell am start -W -n <package>/<launchable-activity>
# 立刻 ≤100ms 輪詢直到穩定或超時（建議 5–8s）:
# dumpsys window | grep mCurrentFocus
# 變化時 screencap
```

預期常見結果（尤其 `Accounts: 0`）：短暫業務 Activity → `com.android.vending`／Unauthenticated*。

## Condition B — 抑止商店搶焦

冷啟動後（或與啟動並行）循環：

```bash
adb shell am force-stop com.android.vending
# 間隔約 100–150ms，維持 15–30s
```

同時再做 `mCurrentFocus` 採樣與截圖。預期：業務 Home／主頁可維持；可能可點進 feature Activity。

記錄必須寫明：**抑止搶焦後業務 UI 可達**，不可寫「繞過 CHECK_LICENSE／Pairip」。

## 證據分層（寫進專案 docs）

| Layer | 內容 |
| --- | --- |
| L1 | 靜態 marker（假說） |
| L2 | A/B focus 時間序列 + 截圖 |
| L3 | feature Activity 是否可操作（在 B 組） |

## 失敗判讀

見 [`../tools-and-failures.md`](../tools-and-failures.md)「冷啟動後進 Play 未登入頁」列。

## 結束條件

- A、B 至少各一輪可複現 focus 表。
- 專案 docs 分開 Static／Observed／Reachability。
- 未宣稱未採樣的 Pairip LicenseActivity hop。
