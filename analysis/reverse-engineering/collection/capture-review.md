# Capture review（HAR／pcap／inventory）

本文件是跨目標的 **已保存抓包與 inventory 審查法**。對應 [`../target-routing.md`](../target-routing.md) 的 **Capture review** 族。Desktop 落地細節見 [`../../desktop/passive-runtime-and-captures.md`](../../desktop/passive-runtime-and-captures.md)；APK 即時流量分流見 [`../../apk/traffic-triage.md`](../../apk/traffic-triage.md)。

主張仍須標 observation／inference／unknown（[`../evidence-contract.md`](../evidence-contract.md)）。

## 何時走本文件

- 使用者提供 HAR、mitmproxy 匯出、pcap、log、螢幕錄影、或既有 inventory。
- 任務是「從既有 capture 抽 API／時序／證據」，不是立刻 live 重抓。
- Historical capture **不**自動擴大成對 live 系統的新授權。

開始前確認 [`../../../enforcement/authorization-scope.md`](../../../enforcement/authorization-scope.md)：標的與允許操作（只讀檔／可再抓／禁 hook）。

## 審查節奏

```text
1 鑑定檔案身分（格式、時間、工具、digest 或 digest_unknown）
2 確認通道集合（網路／log／畫面是否同鐘）
3 做 sanitized inventory（明文 + 可觀察 metadata）
4 分層主張（外觀時序 ≠ 應用層正文 ≠ 業務語義）
5 Fail-closed：無明文則停，不發明 path/body
6 交接：映射 + residual unknowns
```

## 1. 檔案身分

| 欄位 | 說明 |
| --- | --- |
| `path` 代號 | 專案相對；reusable 文件用占位符 |
| `format` | HAR／pcap／mitm dump／log／video |
| `captured_at` / 時鐘對齊 | 多檔是否同窗 |
| `tool` | 產生工具與版本（若可知） |
| `content_digest` | 或缺則 `digest_unknown` |

## 2. Co-temporal multi-channel

理想同鐘通道：網路（pcap／代理日誌）+ 應用 log + 畫面（錄影／關鍵幀）。

| 規則 | 說明 |
| --- | --- |
| 單通道不發明缺層 | 只有 SNI／時序時，不得寫出 REST path 或 JSON body |
| 時鐘漂移要標 | 通道時間差大時，相關性降為 inference |
| 缺通道寫 unknown | `needed_authority: co-temporal <channel>` |

## 3. Sanitized inventory only

Inventory 是**可分享摘要**，不是 raw capture 入庫。

可記入 inventory 的 observation：

- 明文 HTTP 方法／path／狀態碼（已去敏 query／header）
- DNS／SNI／IP／port／時序形狀
- 已知明文應用幀標籤（專案約定且已去 payload）

不可記入 reusable 文件：

- Cookie、token、Authorization、完整 URL 含秘密
- TLS 應用資料密文當「已讀 body」
- 私人絕對路徑、帳號、裝置序號

見 [`../../../enforcement/sanitization.md`](../../../enforcement/sanitization.md)。

## 4. 分層主張（禁止跳層）

| 層 | 可宣稱 | 不可宣稱 |
| --- | --- | --- |
| 傳輸外觀 | 連了某 host／SNI、有 CONNECT、時序 | 「呼叫了某業務 API」 |
| 明文應用 | 實際 method／path／可見欄位 | 未出現在明文中的欄位 |
| 業務語義 | 有對照 seed／UI 的映射時 | 僅靠主機名推測功能 |

協議確認可另標：`static-confirmed`／`capture-confirmed`／`inferred`／`unknown`——仍須對齊三類陳述，不可用標籤掩蓋缺證據。

## 5. Fail-closed on crypto

若標準捕捉看不到明文（certificate pinning、自訂信任、僅加密位元組）：

1. **停止**把密文當 API 正文。
2. 記錄 observation：pinning／TLS 失敗／無明文。
3. **不**安裝系統信任憑證繞過、不卸 pinning、不解密他人 TLS，除非授權範圍**明確**允許且有獨立安全流程（預設否）。
4. 升級權威改走：已授權 hook、靜態協議還原、或接受 `accepted_gap`。

「看得到 SNI」≠「已確認 API」。

## 6. HAR／mitmproxy 實務

1. 列出 entries：method、host（去敏）、path、status、mime。
2. 標 `coverage: partial` 若有 filter／截斷／過大省略。
3. 對照前端／發行物 source location 時標 inference，除非 capture 內嵌可證來源。
4. mitm 原生格式需轉檔時記錄工具邊界。
5. 過大 Evidence：摘要投影進筆記，保留原檔 digest 在專案內。

## 7. 與 live 再抓的交接

| 現況 | 下一步 |
| --- | --- |
| Inventory 已夠關閉 unknowns | 走 claims → close；raw 留專案 |
| 缺明文且授權允許 | handoff Dynamic APK／desktop passive，帶 unknown 清單 |
| 授權僅歷史檔 | 不得自行開 live 捕捉 |

## 最小完成檢查

- [ ] 授權涵蓋「讀此 capture」
- [ ] 檔案身分已記
- [ ] Inventory 已去敏
- [ ] 無明文時已 fail-closed
- [ ] 未把 historical capture 當 live 授權擴大
- [ ] Findings 已分 observation／inference／unknown

← [collection/](README.md)
