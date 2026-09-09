> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-09 - Unity on-screen UI identity comes from loaded objects, not screenshot matching

Status: validated

#### One-line Summary

Unity 畫面上「用了哪個 UI 檔」要從**當下已載入的** `Texture2D` / `Sprite` / Addressable key 與 active GameObject 路徑取得；Android view dump 與對 hash 快取做像素比對都不是身分來源。

#### Human Explanation

Unity IL2CPP 客戶端把介面畫在自己的 canvas 上。Android `uiautomator dump` 往往只剩空的 `FrameLayout` / `SurfaceView` 殼，沒有 `ImageView` src、沒有按鈕文字、沒有檔名。截圖仍可證明「畫面上有什麼區塊」，但不能告訴 agent 那個區塊對應哪個 Unity 物件。

功能包（例如某一關／某一機台的 reel 包）也**不等於**進桌前的 splash／共用 HUD。包已在 UnityCache，只證明那個 bundle 存在，不證明當前 modal 的中央圖、側欄活動卡、開始／關閉按鈕來自該包。

另一層常見陷阱是 app `cache/` 底下的 **hash 檔**（不是 `UnityCache/Shared/<bundle>/__data`）：解碼後可能是 PNG 或帶寬高 header 的 raw RGBA，但檔名沒有物件名。像素搜圖成本高、誤配多，最多當輔助，不能當完備盤點。

#### Trigger

- 前台是 Unity `Activity`，要盤點目前畫面用了哪些圖／按鈕／HUD。
- `uiautomator dump` 幾乎沒有 text / resource-id。
- 已從 UnityCache 匯出大量 PNG，但仍對不上 splash／共用 chrome。
- 裝置 `cache/` 出現大量無副檔名 hash 檔，想用截圖去對。

#### Evidence

- Tool: `uiautomator dump`（空樹）、UnityCache `__data` 物件清單、hash image cache 解碼、live screenshot 當 visiblity 證據（非身分）。
- Sanitized excerpt: 前台 Unity activity 的 Android dump 只有數個無標籤 node；功能 bundle 的 Texture2D 集合不含 splash／側欄 HUD；hash cache 可解出像素但沒有 Unity 物件名。
- Evidence path: target project interface-analysis / art-extraction notes（具體 App 證據不進本檔）

#### Generalized Lesson

1. **身分來源（優先）**：畫面仍開著時列出已載入 `Texture2D.name`、`Sprite.name`、Addressable key、active RectTransform／GameObject 路徑。只記 name、bundle/key、content hash。
2. **可見性來源（次要）**：去敏 screenshot 用來確認某物件有沒有畫出來，以及 layout 分區。
3. **不要**把 Android XML、lobby tile icon、或功能 bundle 的 PNG 清單，直接當成當前 splash／modal 的完備素材表。
4. **不要**把 hash image cache 的像素比對當主線。沒有物件名的檔只能標 `decoded-unidentified`。
5. 功能包、共用 HUD 包、CDN／hash cache 是三層；缺層就標 `unproven`，不要用搜圖硬配。

#### Agent Action

下次要回答「這頁用了哪張圖」時：先確認 Unity 前台 → 做 loaded-object dump（Frida／IL2CPP 或 Addressables hook）→ 用 screenshot 勾選哪些 name 可見。禁止先開一輪截圖對 hash／對 bundle PNG。`uiautomator` 空樹時記 `android-hierarchy-not-applicable`，改走 Unity 身分路徑。

#### Goal / Action / Validation

- Goal: 證明畫面上每個可見 UI 物件的 Unity 身分，而不是只重建外觀。
- Action: 在目標畫面停留時 dump loaded names；對照 UnityCache slug 與（若有）hash cache；screenshot 只做 visiblity。
- Validation or reference source: Android dump 無標籤；至少一個可見區塊對上 loaded object name，或明確標 `unproven`；功能 bundle 未命中的中央／HUD 圖不得寫成該 bundle 已證明素材。

#### Applies When

- Unity IL2CPP / Addressables / Unity UI canvas。
- 授權可讀外部 files，且允許短窗口 runtime probe（不輸出餘額、帳號、完整 CDN URL）。

#### Does Not Apply When

- 原生 Android / Compose / Flutter 畫面（那些才優先用 view dump / widget tree）。
- 已有 Addressable key 或 prefab 路徑的 live 對照，只需 screenshot 做 visiblity。
- 目標是協定／RNG，不是 UI 素材身分。

#### Validation

同一 splash／modal 再開一次：loaded name 集合應穩定（或只差動畫 frame）；hash cache 檔名仍不可當身分。把 dump 裡沒有的可見區塊標 `unproven`，不要用最高相關像素當命中。

#### Promotion Target

- `analysis/apk/tools-and-failures.md`（Unity 工具列 + 失敗表）
- `workflow/apk-analysis/execution-flow.md`（UI evidence：Unity 身分規則）

#### Required Linked Updates

- 必須：`feedback/history/apk-analysis/README.md` 分類數量與 Recent
- 必須：本 category `README.md`
- 必須：`analysis/apk/tools-and-failures.md`
- 必須：`workflow/apk-analysis/execution-flow.md` 加一條 Unity canvas identity 規則
- 已檢查：`analysis/apk/README.md` 已指向 tools-and-failures，不需重複正文
- Project incident 證據留在 `<PROJECT_ROOT>` target docs
