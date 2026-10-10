# Shipped-Artifact Evidence 契約

本文件定義跨目標（APK 靜態、native、JS/Electron、firmware 等）的**取證品質**。它不規定 APK 動態捕捉的步驟順序——那是 [`../../workflow/apk-analysis/artifact-gates/evidence-chain.md`](../../workflow/apk-analysis/artifact-gates/evidence-chain.md) 的職責。兩者應同時成立：workflow gate 管「記什麼層次」，本契約管「每個主張的權威與未知」。

## 三類陳述（必須分開寫）

| 類別 | 定義 | 可宣稱 | 不可宣稱 |
| --- | --- | --- | --- |
| **Observation** | 引擎／工具直接回傳、可複驗的事實 | digest、manifest 欄位、反編譯文字、靜態 xref、pcap 位元組存在 | 「使用者一定會走到這裡」 |
| **Inference** | 由 observation 推出的假說 | 「此 method 可能負責簽章」 | 無 observation 的因果 |
| **Residual unknown** | 尚缺權威／環境才能關閉的問題 | 列所需權威（runtime hook、第二版本對照、授權裝置） | 把 unknown 寫成否定句假裝已證偽 |

寫 findings 時，每一條主張旁標 `observation` / `inference` / `unknown`。混寫會讓後續 agent 把假說當證據。

## Artifact 身分（最低欄位）

對每個被分析的發行物，記錄：

| 欄位 | 說明 |
| --- | --- |
| `path` 或專案相對代號 | 可復現；reusable docs 用占位符，不用私人絕對路徑 |
| `content_digest` | 至少一種穩定摘要（例如 SHA-256）；缺則標 `digest_unknown` |
| `declared_version` / package id | 若可得 |
| `analysis_engine` | 工具族＋版本或「手動 jadx／ghidra」；可選 MCP 橋接須記實際 advertised 能力，不鎖死某 npm 版本為唯一真相 |
| `analyzed_at` | 日期即可 |

同一結論不得跨 digest 混用（例如 A 版的 method body 套到 B 版）。

## Authority 與 limitations

| 規則 | 說明 |
| --- | --- |
| 靜態 ≠ 執行 | 反編譯完整不代表語意完整；`body_status: not_available`／partial 必須保留 |
| 靜態 xref ≠ runtime call | reflection、動態載入、JNI 橋接保持 unknown，直到動態證據 |
| Provider 邊界 | 引擎不支援的格式／架構標 unknown，不得填 false |
| 覆蓋率 | 分頁／截斷／budget 省略時標 `coverage: partial` 並記省略原因 |
| 失敗也是證據 | 超時、取消、缺 JAR／JDK 時保留 partial observation 與 remediation，不刪除過程 |

## Residual unknown 最小欄位

| 欄位 | 說明 |
| --- | --- |
| `question` | 尚未回答的具體問題 |
| `blocking` | 是否阻擋當前 goal |
| `needed_authority` | 例如「授權裝置上的 Frida hook」「第二版 artifact」「HAR 完整頁」 |
| `supporting_ids` | 相關 observation 代號（專案內） |
| `status` | `open` / `resolved` / `accepted_gap` |

## 與 APK evidence-chain 的接縫

| 情境 | 用哪份 |
| --- | --- |
| 寫 APK 單次分析筆記、pcap→hook→decrypt 層次 | apk `evidence-chain` slice |
| 跨 APK／native／Electron 共用「主張可否覆核」 | **本契約** |
| 靜態 JADX 結束要升級動態 | 本契約保留 unknowns → [`../apk/traffic-triage.md`](../apk/traffic-triage.md)／Frida workflows |

## 可選引擎

本地 MCP／CLI 反編譯橋接（若已配置）可作為 `analysis_engine` 之一。方法步驟必須也能用通用 JADX／Ghidra／手動流程執行。引擎缺席時標 `engine_unavailable`，改走 CLI／手動，**不得**把「沒裝 MCP」當成分析完成。

安裝與 client 註冊見 `ai-tools/`（落地中），不寫進本檔。

## Sanitization

進 Ai-skill 或可分享文件前：去 host、token、payload 正文、私人路徑。只保留抽象 pattern。見 [`../../enforcement/sanitization.md`](../../enforcement/sanitization.md) 與 [`../../enforcement/reusable-guidance-boundary.md`](../../enforcement/reusable-guidance-boundary.md)。

## 與 Evidence Candidate System

[`plans/active/2026-06-16-1131-evidence-candidate-system.md`](../../plans/active/2026-06-16-1131-evidence-candidate-system.md) 管的是「哪個 framework plan 該看的 candidate」。本契約管「一次 reverse 分析的主張品質」。**不要**把 shipped-artifact 原始證據丟進 ECS registry。

← [回到 reverse-engineering/](README.md)
