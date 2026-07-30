# Legal Strategy Analysis（Strategy Recommendation Engine）

本 slice 是 `workflow/legal/` 的**共同能力**，不屬於某一種合約。它回答的問題不是
「法律規定是什麼」，而是：

> **怎麼安排，對使用者最有利？有哪些替代方案？各自的代價是什麼？**

這是把 workflow 從「流程自動化」提升為 **legal decision support** 的關鍵 stage：
不是問使用者「法院選哪裡」，而是分析雙方所在地、執行成本、語言、法律熟悉度、判決可執行性，
提出建議與替代方案，再讓使用者決策。

> **本 slice 是 cross-domain pattern 的 domain instantiation。**
> Generic contract（六步管線、四欄格式、two-pass 規則、Optimization Suggestion 規則、
> anti-patterns、promotion 條件）在
> [`../../cross-cutting/decision-support/README.md`](../../cross-cutting/decision-support/README.md)。
> 本檔只放 **legal 專屬**的部分：決策點清單、法律風險 depth gate 的接法、與其他 legal slice 的分工。
> Generic 規則有變動時改 cross-cutting 檔，不在此複製。legal 是目前唯一的 converged case（1／3）。

> **邊界**：策略建議仍受 [`../risk-classification.md`](../risk-classification.md) 約束。
> **Red tier 不產出策略建議**——只做升級。Yellow tier 的建議必須標明哪些前提需律師覆核。

## 為什麼要兩趟（Pass 1 / Pass 2）

策略推理依賴法律前提（某法域的強制法、判決可執行性、範本是否可增修）。若在查證前就給
建議，等於用未核實的假設做決策。因此本 slice 跑兩趟，中間夾 Research：

```text
Stage 1 Intake
      ↓
Stage 3a  Strategy Pass 1（framing）
          → 列出決策點、選項、利益分析、必須查證的前提 → Decision Register
      ↓
Stage 6 Applicable Law Research（查證被標記的前提）
      ↓
Stage 3b  Strategy Pass 2（close-out）
          → 用已核實前提收斂 → Recommendation + Trade-offs → 使用者確認
      ↓
Stage 8 Produce（draft / review / negotiation）
```

Pass 1 的建議一律標 `provisional`；只有 Pass 2 經查證後的建議可標 `confirmed`。

## 分析管線（六步）

| 步 | 名稱 | 產出 |
| --- | --- | --- |
| 1 | **Requirements** | 使用者真正想達成的商業結果（不是他說要的條款） |
| 2 | **Legal Analysis** | 這個結果在相關法域受哪些規則約束；哪些是強制、哪些可約定 |
| 3 | **Risk Analysis** | 若照對方版本／預設做法，我方會承擔什麼；發生機率與影響 |
| 4 | **Interest Analysis** | 哪一種安排對使用者較有利；對手方的真實利益是什麼（可交換什麼） |
| 5 | **Trade-off Analysis** | 每個方案的成本、可執行性、談判難度、時間 |
| 6 | **Recommendation** | 建議 + 理由 + 替代方案 + 代價 + 需使用者決策的點 |

**第 4 步是最常被跳過的一步**。只分析「我方風險」會得出「所有條款都要對我方最有利」的
無用結論。要同時分析對手方為什麼會接受，建議才可執行。

## Decision Reasoning 格式（強制）

任何策略建議**不得**只給結論。每個決策點必須四欄齊備：

```markdown
### <決策點名稱>
- **Recommendation**: <建議方案>
- **Reason**: <為什麼，綁定具體因素：成本／可執行性／強制法／慣例>
- **Alternative**: <替代方案 1、2>
- **Trade-offs**: <選建議方案放棄了什麼；對方可能反對什麼；反對時退到哪一個 alternative>
- **Confidence**: provisional | confirmed（confirmed 需附 Research 來源）
- **需使用者決策**: <是／否，若是說明決策點>
```

範例（台灣公司 vs 日本公司的爭議解決地）：

```markdown
### 爭議解決地
- Recommendation: 新加坡仲裁（SIAC），仲裁語言英文，獨任仲裁人
- Reason: 我方為台灣公司、對方為日本公司，選任一方本國法院對另一方成本不對等，
  對方難接受；仲裁在雙方資產所在地的可執行性優於外國法院判決；SIAC 對雙方皆中立
- Alternative: (1) 台灣法院專屬管轄（我方成本最低但對方大概率拒絕）；
  (2) 東京地方法院（對方接受度高，我方翻譯／律師／執行成本顯著上升）；
  (3) JCAA 仲裁（對方接受度高於 SIAC，中立性略低）
- Trade-offs: 仲裁費用高於訴訟，小額爭議不經濟 → 建議加金額門檻，門檻下先協商／調解；
  若對方堅持日本，退到 alternative (3) JCAA 而非 (2) 東京法院
- Confidence: provisional（待查證：SIAC 現行規則版本、對方資產所在地是否為紐約公約締約國）
- 需使用者決策: 是 —— 仲裁費用可接受度、是否設金額門檻
```

## 策略 playbooks

九個反覆出現的決策點（準據法、管轄／爭議解決、付款條件、違約金、智慧財產、
保密期間、驗收條款、不可抗力、終止）的推理因素與常見結論，見
[`decision-playbooks.md`](decision-playbooks.md)。

Playbook 是**推理起點，不是答案**。每次仍要依本案的 Interest Analysis 調整。

## Optimization Suggestion（使用者已指定做法時）

使用者說「準據法就寫日本法」「爭議解決寫東京法院」而分析顯示有更好選項時，依
[`../../cross-cutting/decision-support/README.md`](../../cross-cutting/decision-support/README.md)
§Optimization Suggestion 規則：**講一次**（含量化差異與使用者原方案的合理之處）→
**明確詢問** → 使用者維持原方案就照做並記錄，**不重複勸說**、**不擅自改動**。

legal 專屬補充：若使用者堅持的方案觸及**強制法**或會讓條款不可執行（而非只是成本較高），
這不是 optimization 而是**可行性問題**——必須明說該方案可能無效，並依
[`../risk-classification.md`](../risk-classification.md) 判斷是否升 Yellow／Red。

## 產出：Decision Register

Pass 1 與 Pass 2 共用同一份表，Pass 2 更新 Confidence 與最終建議：

```markdown
## Decision Register
| # | 決策點 | Pass 1 建議（provisional） | 待查證前提 | Pass 2 建議（confirmed） | 使用者決定 |
| - | --- | --- | --- | --- | --- |
| 1 | 準據法 | | | | |
| 2 | 爭議解決地 | | | | |
| 3 | 付款條件 | | | | |
| ... | | | | | |
```

每個 `#` 在產出中必須有對應的 Decision Reasoning 四欄區塊。

## 與其他 slice 的關係

| Slice | 分工 |
| --- | --- |
| [`../../cross-cutting/decision-support/README.md`](../../cross-cutting/decision-support/README.md) | **Generic contract owner**：管線、四欄格式、two-pass、Optimization Suggestion、anti-patterns |
| [`../intake.md`](../intake.md) | 蒐集事實。**不做**建議 |
| [`../jurisdiction.md`](../jurisdiction.md) | 定義法域相關的**五個變數與依賴**。本 slice 負責**推薦怎麼選** |
| [`../research/README.md`](../research/README.md) | 查證前提。本 slice 標記「要查什麼」，research 回答 |
| [`../risk-classification.md`](../risk-classification.md) | 決定策略建議能做到多深（Red 不做） |
| [`../draft/README.md`](../draft/README.md) | 把已確認的 Decision Register 落成條款文字 |
| [`../review/README.md`](../review/README.md) | 審閱時反向使用：對方條款對應到哪個決策點、偏離我方最佳方案多遠 |
| [`../negotiation/README.md`](../negotiation/README.md) | 用 Trade-offs 欄產生讓步順序 |

## 常見失效模式

| 失效 | 症狀 | 防呆 |
| --- | --- | --- |
| Ask-only | 只問「法院選哪裡」，把決策丟回使用者 | Decision Reasoning 四欄強制 |
| Conclusion-only | 給「建議台灣法院」但無理由與替代方案 | 同上 |
| One-sided interest | 只算我方風險，得出對方不可能接受的方案 | 管線第 4 步 Interest Analysis 必填 |
| Unverified confidence | Pass 1 的建議標成 confirmed | Confidence 欄 + Research 來源要求 |
| Strategy after draft | 先寫完條款再補策略說明 | Stage 順序：3 在 8 之前 |
| Playbook as answer | 直接套 playbook 常見結論，未做本案調整 | Reason 欄必須綁本案具體因素 |
| Repeated pushback | 使用者已決定後仍反覆勸說改方案 | Optimization Suggestion 規則：講一次、記錄、往前走 |
| Silent override | 未經同意就把使用者指定的準據法／管轄改掉 | 同上；建議與執行分離 |
| Optimization theater | 差異不顯著仍硬造替代方案 | 明說「使用者方案合理，無明顯更優選項」 |

## 驗證

- 每個決策點是否四欄齊備（Recommendation / Reason / Alternative / Trade-offs）？
- Interest Analysis 是否分析了**對手方**為什麼會接受？
- Pass 1 的建議是否標 `provisional`，且列出待查證前提？
- Pass 2 的 `confirmed` 建議是否附 Research 來源？
- 需使用者決策的點是否明確列出，而非代替使用者決定？
