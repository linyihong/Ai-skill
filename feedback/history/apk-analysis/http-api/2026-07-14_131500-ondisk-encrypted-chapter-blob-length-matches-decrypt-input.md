> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - On-disk encrypted chapter blob length matches content-decrypt input

Status: candidate

#### One-line Summary

當讀者把章節密文落到 app-private files（常為 `<bookId>/<chapterId>.<ext>`）時：先量測 **檔案 bit 長度是否等於** 具名 content-decrypt／`getContentBody` 的 **inLen**；相等即可把「decrypt 輸入來源」鎖定為磁碟 blob，不必先解出 wire 欄位名。

#### Human Explanation

Wire JSON 欄位名可能因 Gson/`Type`／Kotlinx／converter 難抓；但許多閱讀 App 會把同一不透明字串寫入本地檔。清掉章節目錄可逼出重新 load。若檔案整體是 printable／base64，解碼後常有固定 magic（例如 ASCII `data`）＋二進位 header＋密文——仍只記 magic／長度，不寫 payload。磁碟路徑也可解釋「為何 reader 從 cache 讀卻仍走 encrypt decrypt 管線」。

#### Trigger

- 具名 content decrypt／`getContentBody` 已有穩定 inLen。
- App files 下出現 per-chapter blobs；清檔後 reader 觸發重新網路／寫檔。
- SQLite chapter 列可能**沒有** content 欄，只留 `filePath`／download flag。

#### Evidence

- Tool: size of on-device chapter file vs Frida `getContentBody` inLen；optional FileOutputStream path hook.
- Sanitized: path **pattern** only（`…/books/<id>/<id>.ext`），extension／magic／lengths；no ciphertext samples.

#### Generalized Lesson

1. Compare **file size == decrypt inLen** before chasing wire JSON.  
2. Clearing the chapter directory is a legitimate force-reload knob.  
3. Base64-file → decode → look for short ASCII magic + length fields；document magic only.  
4. Model field names（e.g. `content`）remain hypotheses until converter/`setContent`/`fromJson` hits.

#### Agent Action

1. Locate chapter cache dirs under app `files/`；log write paths with Frida if needed.  
2. Delete cache → reopen chapter → correlate new file size with decrypt inLen.  
3. Keep file bytes out of Ai-skill／shared notes.

#### Promotion Target

- Layer: analysis
- Status: promoted
- Required Linked Updates:
  - `analysis/apk/traffic-triage.md` — Response 解碼補「磁碟 blob 長度 ≡ decrypt inLen」
  - N/A: intelligence／workflow／enforcement

#### Verification

Validated when a force-cleared chapter reload recreates a file whose byte length equals the same-window content-decrypt input length, with reader util still on the decrypt stack.

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
