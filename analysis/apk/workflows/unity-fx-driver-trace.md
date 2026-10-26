# Unity 動畫／特效驅動追蹤

授權 Unity／IL2CPP 分析中，當 `Animator.Play` 命中但目標文字、掉落、位移或粒子不同步時，使用本流程。視覺缺口只證明目前觀測不足；不能直接推論為 Timeline、Native 或 `m_Payloads` 寫入。

資產解析先讀 [Unity 資產流程](unity-asset-parse-flow.md)；端到端編排與完成門檻仍由 [APK workflow](../../../workflow/apk-analysis/execution-flow.md) 負責。原始 dump、物件位址與工作幀留在專案受控 evidence。

## 1. 建立可比較窗口

- 記錄版本／架構、同一 fixture、feature checkpoint、時間基準、目標 GameObject／component 身分及工作幀。
- 分開觀察 Animator state、文字、active／renderer、位移、粒子；保存觸發前、變更中、settle 後的事件與畫面。
- `matched>0` 或 invoke 成功不代表目標視覺已還原。記錄各層 resolved／hooked／hit／unknown；未解析或太晚 attach 的 hook 無法排除路徑。

## 2. 先追 feature 自有 API 與資料

1. 在 `dump.cs`／metadata 搜尋 feature 前綴、`*FX`、`Setup`、`Play*`、`Show`、`GameAction`、`AddPayloads`；feature 類可能在 `Assembly-CSharp`，不能只掃 `UnityEngine*`。
2. 確認 overload、參數型別、this 物件、方法位址；對少量候選 hook，記錄參數形狀、物件身分與相對時間，先觀察正常觸發。
3. 追 `Setup → Play/Show → payload 更新 → renderer/文字/Transform`；欄位可由 constructor 或直接 store 寫入，不保證存在 getter／setter。
4. 只有專案授權的 lab replay 才依正常呼叫順序帶入參數；Unity 物件操作需在合適主執行緒窗口。保留自然觸發與人工 force 的證據差異。

同 fixture 比較 Animator-only 與 feature API 路徑；高層路徑已解釋缺口即可收斂，未解釋部分再升級。

## 3. 依證據選分支

| 分支 | 啟動證據 | 追蹤點與限制 |
| --- | --- | --- |
| 資料驅動 | payload／列表／狀態在缺口窗口改變 | 找真正持有者與 writer；參照、集合項目、項目欄位分開追 |
| Timeline／Playables | 已載入 Director／graph、資產 binding 或呼叫鏈 | `PlayableDirector.Play/Evaluate`、`PlayableGraph.Play/Evaluate`；graph 可由程式建立，不能只靠 `.playable` 檔名排除 |
| 粒子／顯隱 | 粒子續播、物件被關閉、renderer 狀態改變 | `ParticleSystem.Play/Stop/Simulate`、`GameObject.SetActive`、renderer enabled、`Object.Destroy`；`SetActive` 不是假設存在的 `set_active` |
| 骨骼／位移 | Transform 持續改變但 Animator 路徑不足 | 自訂 Update／LateUpdate、tween、Spine 或 Native writer；抽樣後才擴大 Transform setter hooks |
| 第三方節點 | 型別 namespace、組件、資產或 runtime stack 命中 | 證實 NodeCanvas／BehaviorDesigner／xNode 等實際版本與節點方法，再解析該方法；不預設皆有 `OnUpdate` |

UnityEngine 方法可能是 managed wrapper、icall、injected entry 或 Native 實作；先解析當前版本的可執行位址與簽名。全域高頻 hook 必須限制物件、時間、事件數並可解除。

## 4. 靜態 Native 與資產交叉驗證

- Il2CppDumper 的 `dump.cs` 是型別／方法骨架，不是 C# 函式本體。對同版 binary＋metadata 產生 `script.json`／`il2cpp.h`，依工具版本用 `ida.py`／`ida_with_struct.py`、`ghidra.py`／`ghidra_with_struct.py` 載入名稱與結構；不要假定產物叫 `script.py` 或所有型別都能完整還原。
- IDA／Ghidra 追候選方法 xref、field offset 的 store 指令與呼叫者；缺 setter 不代表沒寫入，inline／stripped code 也可能缺 xref。ImHex 用於位元組／結構檢查，不能當成同等的 Native 反編譯器。
- UnityPy 先盤點相關載入包；需要重建 Editor 資產關係時用 AssetRipper 檢查 Prefab、AnimationClip、Director binding、Timeline track、MonoBehaviour 與 renderer。匯出工程可能缺腳本、動態配置或相依包，不能保證完整遊戲可重播。
- 單一包中 Director／Timeline 計數為 0 只排除該範圍；記錄掃描範圍，必要時補 shared／下載包與 runtime graph。Cinemachine 存在不等於 FX 由 Timeline 驅動。

## 5. Payload 寫入監看

1. 證實 declaring type、instance／static／value-type 儲存方式、runtime field offset 與物件存活期；只有確認為 instance field 才用 object base＋offset。不要用猜測的 C++ layout。
2. `m_Payloads` 若是 reference，監看欄位 slot 只抓參照替換；List 的 items buffer、元素與元素欄位是其他位址，擴容／重建後須重新解析。
3. 先 hook 已知 writer；仍無來源時，檢查 Frida／OS／CPU 是否支援 thread hardware watchpoint，再用 `Thread.setHardwareWatchpoint(id, address, size, 'w')`（在對應 thread 物件上呼叫）。限制 size／alignment、硬體槽數與相關寫入執行緒；新執行緒與物件釋放要重新處理。
4. `Process.setExceptionHandler` 只處理例外，本身不設 watchpoint。只接管可確認屬於本次監看的例外；記錄 context 後解除對應 watchpoint，再安全恢復，無關例外交還系統。退出、timeout、物件失效都要清理；不要無條件 return true 吞例外。
5. `MemoryAccessMonitor` 是每頁首次 access 通知，不能當成持續、逐欄位寫入記錄。未命中只能說本次窗口／執行緒未觀測到，不能證明沒有 writer。

## 6. Native backtrace 與收斂

僅在已篩選 hook 的 `onEnter` 或可辨識 watchpoint exception 內，短窗口啟用 `Thread.backtrace(context, Backtracer.ACCURATE)`；預設關閉、限量、去重。符號缺失或 unwind 不完整要記為限制；FUZZY 結果可能有誤判。

保存 module 名稱、build ID／版本、runtime base、PC 與 module-relative offset。`PC - base` 不是檔案 offset；在 IDA／Ghidra 要依 image base／segment mapping 對齊並核對指令，不可直接把差值當檔案位置。

完成紀錄必須包含：

- 同 fixture 的前後工作幀／settle 畫面、事件順序與目標物件身分。
- 已證實 driver、候選與 unknown 分開；人工觸發成功不等於已證實自然完整因果鏈。
- 每個分支的 observed／ruled-out-in-scope／not-tested／blocked，以及 coverage 限制與下一個最小追蹤點。
- 工具版本、hook 解析結果、停止條件與解除結果；沒有裝置或 binary 時只交付方法與待驗證項目。

## 發現與維護鎖

本流程須同時保留 analysis README、APK workflow Markdown／YAML、routing registry、summary 與知識索引的入口。`runtime refresh/validate` 的 Unity discovery 檢查驗證 route candidate、detector、Discovery Bridge 提名與 summary；RAG 檢查要求 `m_Payloads` 與 `PlayableDirector` 可查回本 canonical source。缺任一環應阻擋驗證，不以新增一份 scenario 文件代替可執行檢查。

Discovery Bridge 仍是 advisory，不因此自動啟用或略過授權／來源讀取；本鎖保證接線與查回，不能保證模型已執行每一步。

## 參考

- [Frida JavaScript API](https://frida.re/docs/javascript-api/) 與 [hardware watchpoint 說明](https://frida.re/news/2024/09/06/frida-16-5-0-released/)
- [Il2CppDumper 產物與匯入腳本](https://github.com/Perfare/Il2CppDumper)
- [PlayableDirector.Evaluate](https://docs.unity3d.com/6000.0/Documentation/ScriptReference/Playables.PlayableDirector.Evaluate.html)
- [既有 feature FX candidate lesson](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-26_100652-animator-play-visual-gap-use-feature-fx-api-not-timeline-first.md)；其 promotion 狀態維持 candidate，本流程加入方法與驗證要求，並不新增獨立實測證據。
