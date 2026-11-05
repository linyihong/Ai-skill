# IDA Static Analysis Adapter 導入方案

狀態為待執行。這是主計畫 Phase 2 的工具接入方案，以下兩個接入階段沿用既有 Unity 決策入口與 Selection Responsibility。

| 層 | 責任 |
| --- | --- |
| 靜態分析 | IDA Pro／ARM64 Hex-Rays、IDAPython 與 IL2CPP metadata 建立函式、控制流、參數傳遞和欄位讀寫的證據。 |
| 執行期觀測 | slots-lab／Frida 記錄實際執行順序、物件身分、生命週期與獨立讀回。 |
| 決策與驗證 | Ai-skill 依既有流程裁決證據、選擇最小區分實驗並設定停止條件。 |

IDA database 和匯出物屬於分析證據。決策記在既有 case ledger。商業專案保存 binary、metadata、IDB、原始反編譯與 trace。Ai-skill 保存通用契約和去識別化的方法回饋。

## 接入階段 1 — 驗證環境並建立方法索引

- [ ] 確認授權的 IDA Pro／ARM64 Hex-Rays 可開啟 ELF 並反編譯指定函式，記錄 IDA、decompiler 和 metadata 工具版本。
- [ ] 從同一目標 build 取得 ARM64 `libil2cpp.so` 與 `global-metadata.dat`。記錄 app build、ABI、兩個 SHA-256、ELF build ID（若有）及 image base。以 Il2CppDumper 或 Cpp2IL 擇一產生候選映射。
- [ ] 為 Item／Dice 的 `CreateContent`、`Create/5` 和 reel model 的 `AddCells` 建立 `method identity → RVA → IDA EA` 索引。逐一核對可執行區段、函式入口與呼叫點。比對 Frida 時使用同 build 的 module-relative 位址，保留位址換算依據。
- [ ] 驗收時以至少一個 CreateContent 和 AddCells 核對 metadata、IDA 指令及同版 Runtime 位址。版本或位址無法對齊時記為 blocked，保留候選與缺失來源。

## 接入階段 2 — 可重跑的靜態證據匯出

- [ ] 用單一 IDAPython 入口依目標清單匯出函式清單、直接 caller／callee XREF、指定欄位讀寫及有限深度呼叫鏈。對間接呼叫標記解析狀態。ARM64 ABI 只先匯出 Soft50 需要的參數、回傳與 Rect 指令片段。
- [ ] 每筆包含 build fingerprint、方法身分、RVA、指令位址、證據類型及未解析原因。區分反編譯推論、指令直接支持的觀測與待 Runtime 驗證的主張。Runtime 關聯另記 session/stamp、phase、事件時間及 in-window 物件關聯。
- [ ] 同一 binary、目標清單與工具版本重跑兩次，排序後的證據資料一致。人工抽查一條 caller XREF、一處欄位寫入與一個 Create 回傳值。缺符號、反編譯失敗或間接分派都要在匯出中顯示。
- [ ] 用 Soft50 產出一個帶證據的被動 hook 候選。交付包含入口或出口選擇、要記錄的欄位、兩個假設的預測與停止條件。

## 第一個案例 — Soft50 的格位定位

本次自然進盤已捕捉到 CreateContent、Create/5 與 GetPath。可見 CreateContent 有 `cellId`、Symbol 和 FP Rect，格位解析仍回報 `cell-id-only-no-model-hit`。這個結果留下模型寫入時點與 model instance 身分兩個問題。

1. 優先追 Item／Dice CreateContent 的 cell 參數及回傳值，並追 AddCells 寫入哪個 model instance 和 `m_Cells` 欄位。其次追 Rect 的建立／覆寫與 cellId 的產生／保存。
2. 先用靜態寫入點選擇最小探針。若寫入發生在 CreateContent 之後，預期同一 model 在寫入前找不到該 cell、寫入後能解析 reel/row。若 lookup 查的是不同 model，預期 AddCells 寫入 instance 與 lookup instance 不同，且 cell 出現在前者。
3. 同版 Runtime 被動記錄 CreateContent onLeave、候選寫入點之後的 model 身分與成員、以及獨立讀回的 reel/row。cellId 僅在已驗證的窗口內使用。自然側讀回不足時記 `data_insufficient`；同場 soft 配對成立後才比較 Rect。Rect 寫入由既有 Soft50 契約另行裁決。

完整 Runtime correlation 自動化與唯讀 MCP 放到後續評估。現階段的驗收依上面兩個接入階段。

工具參考：[Hex-Rays 批次反編譯](https://docs.hex-rays.com/9.1/user-guide/decompiler/batch) · [IDAPython API](https://python.docs.hex-rays.com/)。
