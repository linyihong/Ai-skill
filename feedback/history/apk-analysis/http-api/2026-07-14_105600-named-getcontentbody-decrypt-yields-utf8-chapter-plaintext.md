> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - Named getContentBody decrypt yields UTF-8 chapter plaintext

Status: candidate

#### One-line Summary

當章節 load 走 Java HTTP，且存在名似 `DecryptUtils.getContentBody`／`decrypt` 的 util 時：Frida 常可看到 **不透明輸入長度** → **高 printable UTF-8 正文輸出**；用 length／printable ratio 分類管線，**不要**把正文樣本寫進可分享 lesson。

#### Human Explanation

OkHttp `Response.peekBody` 在 R8 下可能失效，但章節明文仍可能在 **具名 decrypt util** 出口出現。實測同窗：`getContentBody(largeOpaque)` 與某 `decrypt(shortArgs)` overload 都產出同一量級、printable≈100% 的 UTF-8 散文；`base64Decode` 常是前段步驟。這證明「wire 可能加密／封裝，但 in-process 最終有 plaintext content path」——與「整條只走 native、Java 無明文」是不同假說。

#### Trigger

- Dynamic URL 已見 `chapter/load`（或同義）。
- App 有 `*DecryptUtils*`／`getContentBody`／類似命名。
- Response converter／peekBody hook 失敗或只見密文 JSON 欄。

#### Evidence

- Tool: Frida method hooks on decrypt util；paired chapter URL logs.
- Sanitized metrics: input/output string lengths；printable ratio；output class = UTF-8 prose（no JSON keys）。
- Never store novel／user text in Ai-skill.

#### Generalized Lesson

1. Prefer **named decrypt／content assemblers** over reinventing peekBody when Response API is obfuscated.  
2. Classify outputs with **length + printable ratio + shape**（JSON keys vs prose vs binary），not content quotes.  
3. Short decrypt inputs that yield long plaintext often mean **envelope／key args**，not “plaintext request”.  
4. Next hooks: who calls `getContentBody`；which model field stores result；whether wire JSON field is base64 blob.

#### Agent Action

1. Hook `getContentBody`／`decrypt` overloads；log inLen／outLen／printable／shape only.  
2. Correlate timestamps with `chapter/load*` URLs.  
3. Keep raw logs gitignored；project docs describe pipeline only.  
4. Do not commit chapter text to feedback／analysis.

#### Promotion Target

- Layer: analysis
- Status: promoted
- Required Linked Updates:
  - `analysis/apk/traffic-triage.md` — Response 解碼流程補「具名 getContentBody／DecryptUtils」捷徑
  - N/A: intelligence（本條以 procedure 為主）；workflow／enforcement N/A

#### Verification

Validated when chapter-load window shows getContentBody／decrypt producing high-printable UTF-8 prose with opaque large input，without relying on OkHttp peekBody.

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
