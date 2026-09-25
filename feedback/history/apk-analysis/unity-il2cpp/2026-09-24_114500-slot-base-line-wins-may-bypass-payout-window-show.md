> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-24 - Base line wins may bypass payout window Show entirely

Status: candidate

#### One-line Summary

許多 slot 小額線獎只走**轉輪上籌碼浮標／線模組**，全程不呼叫 `UISlots*PayoutWindow.Show`／`Setup`，也不呼叫對應 `*PayoutWindowViewModel.Show`／`Setup`。對窗家族設 Interceptor 獵「classic modal」時，數十～百轉 **0 enter** 是合理結果，不代表 hook 壞掉。先用一次 lab `UISlotsPayoutWindow.Show` 驗證 hook 有 `WIN_SHOW`／`FROZEN`；ViewModel 可能不是 `UnityEngine.Object`（`FindObjectsOfTypeAll` 會 AV），但仍可對 MethodInfo 掛 call-site hook。下一步改 chip presenter／提高下注等中額自然窗。

#### Human Explanation

Idle dump 可見 classic payout GO／clip，lab 也可 Create 出實例，但自然 base spin 常：

1. 只亮 payline + chip badge。
2. 從不進 `UISlotsPayoutWindow`（Show／Setup 計數為 0）。
3. Animator.Play／CrossFade 同樣可能 0 enter（自訂 Timeline／UI enable）。

因此「hook 了 PayoutWindow.Show 卻凍不到 classic」≠ 一定 hook 失敗。先確認 READY／手動 lab Show 能送出事件，再累積自然轉數。

#### Trigger

- PayoutWindow Show/Setup Interceptor READY，N 轉後事件計數仍 0。
- 截圖可見線獎籌碼，但無全屏派彩模態。
- 與 lab 強制 Show（會進 hook）對照後仍 0 自然 enter。

#### Evidence

- Tool: Frida Interceptor on `UISlotsPayoutWindow.Show`/`Setup` + spin hunt + screencap.
- Sanitized pattern: 50 base spins → 0 Show/Setup enters while chips decrease；lab Show still detectable when forced.
- Evidence path: `<PROJECT_ROOT>` cabinet `lab/classic-payout-hunt-*.json`.

#### Generalized Lesson

1. **Chip badge ≠ payout window**：兩條呈現路徑；勿假設線獎必 Show modal。
2. **0 enter 先驗證 hook**：lab 強制呼叫一次；有事件才把自然 0 當成路徑證據。
3. **下一層 hook**：line-chip／badge presenter，或提高下注等明確中額 celebration；ViewModel.Show 也可能同樣 0 enter。
4. **IL2CPP Interceptor**：attach `MethodInfo.readPointer()`；lab 強制 Show 用來驗 hook，不要只用自然 0 判斷失效。
5. **ViewModel 未必是 Unity Object**：`FindObjectsOfTypeAll(ViewModel)` 可能 AV；call-site hook 仍有效。

#### Agent Action

自然獵 modal 前記錄 Show enter 計數；0 時改路徑或提高門檻樣本，勿空轉同一 hook。

#### Goal / Action / Validation

- Goal: 分辨「沒進窗」vs「hook 失效」。
- Action: READY → lab Show 煙測 → N 轉計數 → 決策下一 hook。
- Validation: lab Show 有事件；自然計數寫入 hunt JSON。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP slots 同時有 on-reel chip 與全屏 payout window。
- Does not apply: 所有贏獎都強制開同一 payout modal 的 cabinet。

#### Related

- `2026-09-24_110000-slot-shared-payout-show-needs-setup-not-blind-sibling-invoke.md`
- `2026-09-24_113000-slot-payout-create-needs-path-rect-sound-and-settings-init.md`

#### Validation

- 先以 lab Show 取得 hook 事件，再記錄自然樣本的 enter 計數；只有兩者皆有才可把零自然命中解讀為路徑證據。

#### Promotion Target

- apk-analysis slot UI hunt：chip path vs window path；Show enter 計數。

#### Reuse Evidence

- 尚未有獨立 cabinet／專案的重用證據；保留為 candidate。

#### Promotion Record

- 尚未 promotion；僅保留為 feedback history candidate。

#### Required Linked Updates

- 已更新 `unity-il2cpp/README.md` category index；尚未更新 UI hunt SOP，待第二個呈現家族驗證後再評估。
