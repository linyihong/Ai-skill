# APK 靜態 JADX 取證路徑

授權 APK 上，用**靜態**反編譯回答 class／method／manifest／靜態 refs 問題。遵守 [`../reverse-engineering/evidence-contract.md`](../reverse-engineering/evidence-contract.md)。

本路徑**不執行** APK、不做 native-library 動態分析、不做 Frida／MITM。那些目標走 [`traffic-triage.md`](traffic-triage.md) 與 [`workflows/`](workflows/README.md)。

## 何時用

| 適合 | 不適合（應 handoff） |
| --- | --- |
| Manifest／permission／component 盤點 | 業務 API 實際流量 |
| 搜 class／method 名、讀反編譯 body | Flutter AOT／Unity IL2CPP runtime |
| 靜態 incoming refs | 需 hook 才看得到的解密／簽章 |
| 為動態 hook 找高語意符號 | Proxy bypass／TUN 行為 |

## 步驟（工具中立）

1. **授權與 artifact 身分**
   確認授權；記錄 package／版本／digest（或 `digest_unknown`）。
2. **Package 概覽**
   Manifest、permission、component 摘要。對照原始 XML 若摘要可疑。簽章驗證若未做，標 `not_performed`。
3. **搜尋候選**
   依功能關鍵字／已知 class 片段搜尋。空 query 清點 class 名成本高——有目標再搜。
4. **Class inventory**
   讀 fields／methods／inner types；overload 先列清再選。
5. **Method 反編譯**
   指定 overload（勿默認第一個）。記錄 `complete`／`partial`／`not_available`（native／abstract）。反編譯文字是 derived representation。
6. **靜態 refs**
   Incoming static relationships；非 runtime call。reflection／動態載入保持 unknown。
7. **寫 Evidence**
   每條主張標 observation／inference／unknown；保留引擎限制。

### 可選引擎

- 通用：`jadx` GUI／CLI、`apktool`（resources／smali）、`aapt`。
- 若已配置本地 MCP／CLI Android inspect 族工具，可用同等步驟；以當時 `tools/list` 為準。見 `ai-tools/`（落地中）。

## Handoff 到動態主線

出現任一信號就升級，不要延長純靜態：

| 信號 | 下一站 |
| --- | --- |
| 需要實際 request／response／明文 | [`traffic-triage.md`](traffic-triage.md) |
| Java hook 或 OkHttp 路徑 | [`workflows/frida-hook-flow.md`](workflows/frida-hook-flow.md) |
| `libapp.so`／Flutter | blutter／unflutter + Frida（tools-and-failures） |
| Unity／IL2CPP／cache 素材 | Unity 相關 lessons／tools 表 |
| Local proxy／loopback | [`workflows/local-proxy-hook-flow.md`](workflows/local-proxy-hook-flow.md) |

Handoff 時把殘餘 unknown 原樣帶入動態筆記（需要哪種 runtime 權威）。

## 失敗判讀（摘要）

| 現象 | 處理 |
| --- | --- |
| Overload 歧義 | 先 class inventory；拒絕猜測 index |
| Method body 不可用 | 標 not_available；考慮 smali 或動態 |
| 只有 native stub | handoff native／Frida，不假裝 Java 層已解 |
| 引擎／JDK 缺失 | `engine_unavailable`；改 CLI 或記錄 blocked |

細節併入 [`tools-and-failures.md`](tools-and-failures.md) 維護。

← [回到 analysis/apk/](README.md)
