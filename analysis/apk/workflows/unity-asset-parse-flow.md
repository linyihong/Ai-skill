# Unity 資產解析流程（UnityPy / AssetRipper / AssetStudio / Il2CppDumper）

本流程回答：**「UnityCache／AssetBundle 裡的靜態圖、粒子、Prefab 能不能還原？用哪把工具？」**
屬於 analysis-local procedure（HOW TO COLLECT）。工具選擇總表見 [`../tools-and-failures.md`](../tools-and-failures.md) §Unity。

專案特例、機台名、live 金額、原始 dump 路徑留在業務專案；本檔只寫可重用步驟與邊界。

## 何時啟動

- 功能頁／機台已進桌，美術在 **UnityCache**（不是只掃 APK）。
- 目標含：Sprite／Texture2D、Spine TextAsset、**ParticleSystem**、Animator、Prefab hierarchy。
- 使用者問「能不能用 Il2CppDumper／AssetRipper／AssetStudio 還原粒子／特效」。

## 工具分工（依證據目標與版本支援選擇）

| 優先 | 工具 | 產出 | 粒子／特效 |
| --- | --- | --- | --- |
| **1（預設分析）** | **UnityPy**（對 `__data` / `.unity3d`） | Texture2D／Sprite PNG；object type counts；可讀 `ParticleSystem` 各 Module（Initial／Emission／Shape…）；Material 名 | **可還原參數表 + 貼圖身分**；不產出可直接開編輯器的 Prefab YAML |
| **2（要進 Unity 重播）** | **AssetRipper** | Prefab／Material／ParticleSystem／Texture 的可匯入專案樹 | **最接近「完整粒子資產還原」**；Unity 版本／腳本 stub 可能需對齊 |
| **3（瀏覽貼圖）** | **AssetStudio** | Texture／Mesh／Audio 瀏覽與匯出 | **弱**：ParticleSystem 模組保真度通常不如 AssetRipper／UnityPy 結構化讀取 |
| **輔助（程式身分）** | **Il2CppDumper** | `dump.cs`／方法名／型別 | **不匯出** ParticleSystem／貼圖；用來寫 Frida（例如 line-payout／chip FX 類名） |

**預設不要**為了「看粒子」先開 AssetStudio。先 UnityPy 確認 bundle 內有 `ParticleSystem` + material／texture 名；若要進 Editor 重播再跑 AssetRipper。

## 已到達 checkpoint 的取證順序

1. 驗證目前功能／場景與物件身分，先取得當前可觀測窗口的畫面、資產 inventory、component／material／texture 關係。
2. 未捕獲的入場／過場獨立標 `not-captured-this-session`；不可借用當前背景宣稱另一窗口已還原。若完整入場是驗收條件，仍須補拍。
3. 以實際引用定位功能包與共用包；名稱前綴與包位置只提供搜尋線索，不是固定規則。
4. 報告檔名、層數、manifest、HTML 版型及生成腳本由專案契約決定；共用方法只要求身分可追溯、證據與宣稱一致。
5. 原始證據、報告實作與可重用方法分層維護；提交頻率遵守各 repo 的交付規則。

來源候選：[checkpoint 取證](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-23_102500-slot-iface-analysis-seated-unitypy-batch-and-commit-cadence.md)、[方法與報告邊界](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-23_103000-slot-iface-no-oneshot-generator-use-companion-scripts.md)。

## 步驟

### 1. 取得 `__data`

1. 裝置已進入目標功能（機台／頁面），確認 `UnityCache/Shared/<bundle>/<hash>/` 存在。
2. `adb pull` 該 hash 的 `__data` 到專案 gitignored 目錄（名稱由專案決定）。
3. 記錄 `bundle` slug 與 hash；不要把全文 dump commit。

### 2. 序列化資產盤點（預設使用 UnityPy）

1. `UnityPy.load(__data)`，統計 `ParticleSystem`／`ParticleSystemRenderer`／`Texture2D`／`Sprite`／`AnimationClip`／`TextAsset` 等 counts。
2. 匯出粒子相關 Texture／Sprite PNG（名稱常含 `particle`、`coins`、`firework`、`payout`）。
3. 讀每個 `ParticleSystem`：經 `m_GameObject` 取 host 名；摘 `lengthInSec`、`looping`、`playOnAwake`、Initial／Emission／Shape 關鍵 scalar。
4. 讀 `ParticleSystemRenderer.m_Materials` → material 名 → 對回 Texture。
5. 產出專案內機器可讀 dump（JSON）+ 精選 PNG（僅文件真正引用的檔）。

成功門檻：知道「有幾個粒子系統、掛在哪些 GO、用哪些 material／貼圖、預設 duration／rate」。

### 3. Live 對照（特效正在播時）

1. `adb screencap`（金額可裁切／遮罩後再進 docs）。
2. Frida：優先 `GameObject.Find`／已知 path（由盤點結果選取）；避免無界 `FindObjectsOfTypeAll` 卡住主執行緒。
3. 記錄 `isPlaying`／active 的 GO 名，對上步驟 2 的 host 名。
4. 截圖只證明 visibility；身分以 GO／material／texture 名為準。

### 4. AssetRipper（可選 — 要 Editor 重播）

1. 對同一 `__data`（或 ripper 支援的 bundle 格式）跑 AssetRipper。
2. 匯出含 ParticleSystem 的 Prefab／場景片段與 Materials。
3. 在對應 Unity 版本開專案驗證；腳本缺失時保留 stub，以視覺／粒子參數為主。
4. 標註：可重播 ≠ 已接上業務 spin／線獎狀態機（那仍靠 Frida／協議）。

### 5. AssetStudio（可選 — 只瀏覽）

1. 用來快速翻 Texture／Audio；不要當成 ParticleSystem 參數的 source of truth。
2. 若與 UnityPy 貼圖名不一致，以 UnityPy／runtime Frida 名為準。

### 6. Il2CppDumper（可選 — 寫 hook）

1. 對 `libil2cpp.so` + `global-metadata.dat` 出 `dump.cs`。
2. 搜尋 payout／chip／particle／FX 相關類，供 Frida attach。
3. **禁止**宣稱「Il2CppDumper 已還原粒子資產」。

## 還原成熟度標籤

| 標籤 | 含義 |
| --- | --- |
| `texture-only` | 只有 PNG／atlas；無 ParticleSystem 參數 |
| `params-from-bundle` | UnityPy／AssetRipper 已讀出模組預設值 + material |
| `live-matched` | 特效播時 GO／material 與 bundle dump 對上 |
| `editor-replayable` | AssetRipper 匯入後可在 Editor 目視重播 |
| `game-wired` | 另接上 spin／線獎觸發（超出本流程） |

上述是各自需證據的狀態，不是固定線性階梯；Editor 重播不自動證明 runtime 匹配，runtime 匹配也不自動證明可重播。`game-wired` 另需業務觸發鏈證據。

## 失敗判讀

| 現象 | 原因 | 下一步 |
| --- | --- | --- |
| UnityPy 有 ParticleSystem 但 GO 名全 `None` | 名稱在 GameObject，不在 component | 解 `m_GameObject` PPtr |
| 只有貼圖、以為「粒子還原完了」 | 缺 Emission／Shape／Renderer material | 補讀模組；必要時 AssetRipper |
| AssetStudio 看不到粒子曲線 | 工具偏貼圖瀏覽 | 改 UnityPy／AssetRipper |
| Il2CppDumper 後仍無 FX 圖 | 錯工具 | 回 UnityCache `__data` |
| Live 有光效但 bundle 無同名 texture | 共用 pack或 runtime atlas | 另拉 common／Frida material 名 |

## 產出格式

- 專案：`interface-analysis/*particle*.json`（或等價）+ 精選 `assets/` PNG／thumb。
- 可重用：本流程 + `tools-and-failures.md` 列；lesson 只寫 generalized 邊界，不寫機台金額。

## Related

- [`../tools-and-failures.md`](../tools-and-failures.md)
- Feedback：`slot-three-visual-planes-and-in-cabinet-anim-stack`、`unity-feature-art-lives-in-unitycache-not-apk`
- Agent orchestration：[`../../workflow/apk-analysis/execution-flow.md`](../../../workflow/apk-analysis/execution-flow.md)

動畫已播放但視覺仍缺失時，改讀 [Unity 驅動追蹤](unity-fx-driver-trace.md)，從 feature FX API 查到資料／Playables／Native；資產參數本身不證明觸發鏈。

## 候選回饋承接

粒子資產回饋主要增加證據取得程序；不因有來源回饋就必須另建 intelligence atom。

### 來源回饋

- [2026-10-23_101000-unity-particle-restore-needs-unitypy-or-assetripper-not-il2cppdumper.md](../../../feedback/history/apk-analysis/unity-il2cpp/2026-10-23_101000-unity-particle-restore-needs-unitypy-or-assetripper-not-il2cppdumper.md)
