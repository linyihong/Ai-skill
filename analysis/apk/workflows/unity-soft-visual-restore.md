# Unity 機台 soft 視覺還原：資料先、驅動停、最後才畫

授權 Unity／IL2CPP **lab／presence** 分析中，當「畫面 API 已呼叫／invoke ok」但畫面彈回、格內容讀不回、或下一幀被蓋掉時使用。本流程回答：**為什麼不能只呼叫 `Play`／`Reset`／`SetContent` 就當還原完成**，以及建議的剝洋蔥順序。

特效驅動分支見 [Unity 驅動追蹤](unity-fx-driver-trace.md)。參數／hook 契約見 [IL2CPP 執行契約驗證](il2cpp-runtime-contract-validation.md)。端到端編排仍由 [APK workflow](../../../workflow/apk-analysis/execution-flow.md) 負責。

## 何時啟用

| 訊號 | 含義 |
| --- | --- |
| Soft mode／`SetBonus*`／Animator bool ok，畫面仍像主場或半場 | gate ≠ paint（另見 soft-gate lesson） |
| `SetContent`／`MoveContent`／`CreateContent` invoke true，但 `GetContent` 恒 null 或下一幀清空 | **內容圖／狀態模型**未進入自然進場圖，或每幀 writer 覆寫 |
| `Animator.Play("Idle")`／`SetLayout` 後瞬間彈回結算／轉輪姿態 | 資料模型仍標記 spinning／settle，或 Update／Tween／Timeline 仍在刷 |
| Lab 只改 Transform／GO active，狀態旗標未動 | 典型「只畫皮」 |

**不要**把本流程當成「一鍵 KillAll 就能合法還原正式局」。正式結果仍以授權窗內的 wire／SpinInfo／State 為準；lab 只在專案授權的 presence／replay 窗口使用。

## 核心原則：資料是王

多數機台前端是**高級播放器**：停輪圖符、獎圖、進場序列常由伺服器（或 lab 注入的等價 SpinInfo／Rewards／State）決定。前端腳本在 `Update`／行為樹／Action 佇列裡持續對齊**記憶體模型**與畫面。

| 錯誤做法 | 常見後果 |
| --- | --- |
| 只呼叫畫面 API（Animator／SetActive／SetContent） | 下一幀發現模型仍是 settle／spinning → **強制拉回** |
| 只改 Transform／Quaternion | Update／Tween／Timeline 持續覆寫 |
| 只 Kill 動畫、不改 `m_State`／SpinInfo／mode | Idle 畫面與「仍在結算」模型衝突 |

**破解思維（lab）：** 先讓**狀態變數／payload／進場 Action 圖**與目標畫面一致，再談 paint API。讀回 oracle 優先用契約上的 getter（如 content／length／mode），不要只用 screencap 或 invoke 布林。

## Unity 常見隱藏驅動（前端）

在怪罪 Timeline 之前，依證據選分支（與 [unity-fx-driver-trace](unity-fx-driver-trace.md) §3 對齊）：

### 1. 每幀硬寫 Transform／Quaternion

- 滾輪／棋盤格常在 `Update`／`FixedUpdate`／`LateUpdate` 用程式算角度或位置，**Animator 可能根本沒在跑**。
- 歸位 API 成功後，下一幀數學公式依時間／delta 再蓋掉。
- **探測：** 找出 reel／board controller 上帶 Update 的自訂腳本；短窗口 hook Update 計數或暫時 early-return（僅 lab、可解除）。

### 2. Coroutine 與 Tween（含 DOTween）

- `StartCoroutine`／`DG.Tweening`（或同類緩動）託管加速→勻速→減速→回彈。
- 呼叫 Stop／Reset 後，**既有 handle 仍可能跑**，繼續把節點拉向原終點。
- **探測：** 枚舉載入組件／image 是否出現 `DG.Tweening`（或專案實際 tween 命名空間）；必要時在授權 lab 呼叫專案確認過的 `Kill`／`KillAll` 等價 API（參數語意以該版為準，勿盲殺正式局）。
- **限制：** KillAll 可能弄亂共用 HUD／轉場；優先 Kill **目標物件綁定**的 tween。

### 3. Playables／Timeline／payload 時間軸

- 過場與大獎常由 `PlayableDirector`／graph 驅動；`m_Payloads`（或等價列表）按時間軸開關特效與骨骼。
- 單一元件改狀態，**時間軸仍在 Play** 時會被拉回。
- **探測／處置：** 見 [unity-fx-driver-trace](unity-fx-driver-trace.md)（Director `time`／`Evaluate`／`Stop`；先證實 binding 再動）。

## 建議順序（剝洋蔥）

在 lab 窗口、可回滾的前提下，**不要**一開始就狂打 paint API：

1. **對齊資料模型**
   SpinInfo／StateInfo／mode／Rewards／進場所需列表是否已是「目標畫面」該有的形狀？缺則先走自然／半自然進場 Action（Present*／Start*／Completed），或授權的 fixture inject——不要用空 Rewards 硬 SetContent 假裝閉環。
2. **停止或閘住每幀覆寫（可選、短窗口）**
   Hook 候選 controller 的 Update／LateUpdate early-return，或暫停相關 Director；記錄可解除方式。
3. **清背景計時器（若證據指向 Tween／Coroutine）**
   目標範圍 Kill；證明組件存在後再呼叫，避免「假設有 DOTween」。
4. **最後才視覺歸位／填格**
   Animator／Creator／SetContent／MoveContent／Show；以 **GetContent（或等價）非 null + screencap** 為 paint 閉環，不以 invoke true 為準。
5. **失敗分類**
   - 模型從未進場圖 → 補 Present／Start／Completed，不是換 layer 字串
   - 模型對、每幀蓋掉 → 停 Update／Kill tween
   - 模型對、無每幀、GetContent 仍 null → 回 [契約驗證](il2cpp-runtime-contract-validation.md)（ABI／層名／wrong reels 物件）

## 探測清單（Frida／IL2CPP）

| 檢查 | 作法（工具中立） |
| --- | --- |
| Tween 組件 | 列 image／assembly 名是否含 `DG.Tweening` 或專案已知 tween 庫 |
| 每幀腳本 | 在 reel／board／field 相關 klass 找 `Update`／`LateUpdate`／`FixedUpdate`；短窗口計 hit |
| Director | 找 `PlayableDirector` 實例 isPlaying／time；對照缺口窗口 |
| 狀態欄位 | 讀 mode／spinning／settle／boardGame 等旗標與 enum；寫入前先 dump 自然進場對照 |
| 內容圖 | SetContent 後立刻 GetContent；null 則標 **map 未掛**，勿報「已填格」 |

## 與 soft-gate／Creator 的關係

- Soft mode／Animator bonus：**必要但不充分** → [lab soft-gate lesson](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-27_100500-lab-soft-gate-not-visual-enter-action-queue-and-creators.md)。
- 本檔補的是：**即便 Creator／SetContent／MoveContent 都 ok**，若資料圖或每幀／Tween／Timeline 未處理，讀回與畫面仍可失敗。

## 產出

- 分支判定表：data／Update／Tween／Timeline／ABI（observed／ruled-out／not-tested）
- 還原步驟實際執行順序與解除記錄
- Paint oracle：getter 讀回 +（授權時）screencap；失敗截圖是否進 viewer 由專案政策決定

## 適用／不適用

**適用：** 授權 lab／presence；視覺與模型不一致；soft 進場半成功。

**不適用：** 只做 wire／Rewards 解析無 UI 目標；未授權對正式對局破壞性 KillAll／改餘額；非 Unity。

## 候選 lesson

- [soft 視覺還原：資料→停驅動→再 paint](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-27_110000-slot-soft-visual-restore-data-drive-then-kill-then-paint.md)
- [soft-gate ≠ 視覺進場](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-27_100500-lab-soft-gate-not-visual-enter-action-queue-and-creators.md)
