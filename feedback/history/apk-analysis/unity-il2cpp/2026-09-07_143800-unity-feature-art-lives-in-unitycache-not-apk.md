> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-07 - Unity feature art often lives in UnityCache, not the install APK

Status: validated

#### One-line Summary

Unity IL2CPP 的功能美術（機台／關卡包）常在執行期寫入 `UnityCache/Shared/<bundle>/…/__data`；只掃 APK 裡的 `.unity3d` 或用物件名過濾，會把畫面上已顯示的圖誤判成「拿不到」。

#### Human Explanation

安裝包可能只帶 engine、core chrome、localization。真正功能包用 Addressables／AssetBundle 下載。Android 上常見落點是外部 `files/UnityCache/Shared/`：每個 bundle 一個目錄，底下再一層 hash，檔名 `__data`（本體）與 `__info`（很小的標頭）。`__data` 通常可直接給 UnityPy 當 AssetBundle 讀，不必先當加密容器。

常見誤判：APK split 沒有對應功能名的 bundle → 結論「美術都在 CDN、無法落地」。使用者已經把功能頁打開時，cache 裡可能已有部分包；但 **畫面上的每一張縮圖也不保證都已成獨立 bundle**——目錄名與 UI 標題常對不齊，未點進的機台可能還沒下載。

#### Trigger

- APK 是 Unity IL2CPP（`libil2cpp.so` + `global-metadata.dat` + `.unity3d`）。
- 靜態只抽出 core chrome，或功能頁截圖上的標題在 APK bundle 清單裡找不到。
- 需要把畫面上的 2D 美術抽成 PNG，而不是先做 HTTP MITM。

#### Evidence

- Tool: `adb` + UnityPy on UnityCache `__data`
- Sanitized excerpt: Shared 目錄出現 `feature.common` 與若干 `feature.<machine>` 包；`__data` 可 load，Texture2D/Sprite 可存 PNG。可見 UI 標題集合 ⊃ cache 目錄名集合。
- Evidence path: target project `docs/art-extraction.md`（具體 App 證據不進本檔）

#### Generalized Lesson

1. 分三層：APK 內 AssetBundle、UnityCache 已下載包、仍只在記憶體／尚未落地的遠端資源。
2. 功能頁打開後先 `find` `UnityCache/Shared`，用目錄名對 UI，再 `pull` `__data`。
3. 不要用 APK 物件名 substring（例如功能關鍵字）當「這個畫面所有圖」的完備集合。
4. UnityCache 未出現的標題：記 `not-yet-cached`，不要寫成「無法匯出」。

#### Agent Action

下次對 Unity 目標抽美術時：APK 掃描之後，只要使用者已進功能頁，就列 UnityCache 並匯出 `__data`。截圖標題與 cache 目錄做差集。不要先假設 pinning／無法拿圖。

#### Goal / Action / Validation

- Goal: 判斷畫面上的 Unity 美術能否落地匯出。
- Action: 對照 APK bundles vs `UnityCache/Shared` vs screenshot labels；UnityPy 匯出 Texture2D/Sprite。
- Validation or reference source: 至少一個 cache `__data` 匯出非零 PNG；差集標題標成 not-yet-cached。

#### Applies When

- Unity IL2CPP / AssetBundle / Addressables。
- 授權裝置可讀 App 外部 files（通常不需 root；內部 data 才要）。

#### Does Not Apply When

- Flutter / 純 Java 資源。
- 自訂加密、UnityPy 無法識別的 blob（先當失敗，不要硬解）。
- 需要的是協定／RNG，不是美術檔。

#### Validation

同一功能頁再開一次：已開啟過的機台包應仍在 Shared；新點的機台應多一個目錄。匯出 PNG 尺寸與畫面縮圖不必 1:1（atlas）。

#### Promotion Target

- `analysis/apk/tools-and-failures.md`（Unity cache 工具列）
- `analysis/apk/README.md`（指向該列即可）

#### Required Linked Updates

- 必須：`feedback/history/apk-analysis/README.md` 分類與 Recent
- 必須：本 category `README.md`
- 必須：`analysis/apk/tools-and-failures.md` 增加 Unity IL2CPP 小節
- 已檢查：`workflow/apk-analysis/execution-flow.md` 不改（仍是流量 triage 主線；美術落地是旁路）
- Project incident 證據留在 `<PROJECT_ROOT>` target docs
