> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-26 - Animator.Play visual gap → feature FX API before Timeline/IDA

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源為 Unity IL2CPP lab：強制 `Animator.Play` 有 hash／matched 但仍無目標倍率／掉落視覺；改 hook 櫃機專屬 FX `Setup`／`Play*` 後截圖可見。具體柜名／資產留專案 docs。

#### One-line Summary

當 lab 強制 `Animator.Play`（或 SetActive）回報成功卻仍有「視覺缺口」時，優先在 `dump.cs` 找同功能的 feature FX／`SlotsGameAction` API 並 hook／invoke，不要先假設 Timeline／Playables 或開 IDA。

#### Human Explanation

現代 Unity 柜常同時有 Animator 狀態名與高層 C# 驅動（例如 `*FX.Setup(int)`、`PlayMultiplier`、`Show`、柜專屬 `*GameAction.Start`）。Animator 狀態可能只是 FX 節點內部參數的一部分；外層若未寫入倍率／未 `Setup` 符號列表，盲目 `Play(hash)` 會「方法成功、畫面不變」。

Timeline／PlayableDirector、硬體斷點、AssetRipper 仍有價值，但成本高且本機常缺工具。先做靜態類名搜尋（`*HighWild*`、`*DropItems*`、`*PlayAction*`、`AddPayloads`）再 Frida hook 這些方法，通常更快鎖定真正驅動路徑。第三方行為樹（NodeCanvas／BehaviorDesigner）也要用 dump 字串證實，不要預設存在。

#### Trigger

- `Animator.Play`／`StringToHash` ok、`matched>0`，截圖仍無目標 chrome
- 符號／物件仍在位移或粒子續跑，但標籤／掉落未出現
- 猜測 Timeline／m_Payloads 寫入，但資產盤點 `PlayableDirector`／`TimelineAsset` 計數為 0

#### Evidence

- Tool: Il2CppDumper `dump.cs` + Frida apply-shim Interceptor／`runtime_invoke`
- Sanitized excerpt: feature FX `Setup(n)`＋`Play*` 後 `scanned/setup/played>0` 且截圖出現倍率標；盲目 Animator 路徑無標
- Evidence path: `<PROJECT_ROOT>/docs/slots/*/lab/` apply-shim／drive-path notes

#### Generalized Lesson

1. **視覺缺口 ≠ Animator 失敗**：先分「狀態機未進」與「外層 FX API 未帶參數」。
2. **dump.cs 類名優先於 Timeline／IDA**：搜 `*FX`、`*Drop*`、`*PlayAction*`、`AddPayloads`、柜名前綴；確認 `PlayableDirector` 是否真在映像／資產中。
3. **Hook 高層 API**：對 `Setup`／`Play*`／`Show`／`GameAction.Start` 設 Interceptor；可選短 backtrace（預設關，避免拖垮 settle）。
4. **Lab force 對齊真實 API**：`FindObjectsOfTypeAll`＋`Setup(參數)`＋`Play*`；必要時 `SetActive(true)`；延遲到轉輪停定再拍。
5. **第三方插件要證實**：NodeCanvas／BehaviorDesigner／xNode 僅在 dump／字串命中後再 hook `OnUpdate`。
6. **Assembly 掃描位置**：柜機 C# 在 `Assembly-CSharp`；勿把 feature 類掃描寫進僅掃 `UnityEngine*` 的迴圈。

#### Agent Action

- Animator 視覺缺口時：先 `dump.cs` 找同功能 FX／GameAction，再 hook／force 那些方法。
- 不要一開始就要求 IDA／AssetRipper／HW breakpoint（工具常缺）。
- Trace 開著時預設不要 ACCURATE backtrace；需要時再開。
- 成功路徑寫入專案 lab README，並把可重用規則回寫本 history（本條）。

#### Goal / Action / Validation

- Goal: 補齊 Animator-only lab 抓不到的倍率／掉落等視覺
- Action: dump 找 FX API → Frida hook／invoke → settle 後截圖
- Validation or reference source: hook 事件＋`setup/played>0`＋截圖可見目標 chrome

#### Applies When

- Unity IL2CPP 柜機／slots lab，有 Il2Cpp dump 與 Frida
- 已證實 Animator 路徑不足以驅動目標視覺

#### Does Not Apply When

- 目標確實由 Timeline／PlayableDirector 驅動（dump／資產有類型）
- 純 Spine／原生 C++ 改骨骼且無對應 C# FX 包裝（再升級 native）

#### Validation

- 同 fixture：Animator-only 無 chrome；FX API force 有 chrome
- READY／hook 表顯示 feature 方法已解析（非全 false）

#### Promotion Target

- [Unity 驅動追蹤](../../../../analysis/apk/workflows/unity-fx-driver-trace.md)。

#### Promotion Record

2026-10-26：方法已接入上述 analysis 流程與 APK workflow；仍維持 candidate，入口接線與本輪文件驗證不算獨立重用證據。

#### Required Linked Updates

- 專案 apply-shim：feature FX force／trace 旗標（本輪已改）
- 專案 cabinet lab／next-windows／assets 索引（本輪已改）
- category README 追加本條

#### Abstraction Check

- Swap-project test: PASS — 任何 Unity IL2CPP 功能視覺在 Animator.Play 成功仍缺口時，均可改找 feature FX／GameAction API
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/*/lab/` 與 `assets/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS

#### 修訂 2026-10-26

補上實際方法入口與證據狀態；不再把 history 自身當作提升目標。
