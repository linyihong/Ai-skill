# Decision Support Stage（cross-domain execution pattern）

一個跨 domain 的 execution pattern：在 **Intake 之後、Execution 之前**插入一個
分析與建議階段，讓 workflow 從 *instruction following* 變成 *decision support*。

```text
Intake  →  Decision Support  →  Research / Verify  →  Execution  →  Review
```

> **狀態**：`pilot`。本檔定義 **generic contract**（管線、格式、規則、domain 需提供什麼）。
> 目前 **2／3 converged cases**（`workflow/legal/strategy/`、`workflow/investment/strategy/`）。依
> [`../README.md`](../README.md) §Slice promotion policy，達到 **3 個** converged case
> 之前**不**註冊 `route.*`、**不**寫進各 workflow 的必跑 stage、也**不**宣稱已是全庫能力。

## 問題

Workflow 的預設行為是照使用者說的做。但使用者陳述的往往是**一個合法選項**，不是
**最有利的選項**——而選項之間的差異常常很大：

| Domain | Instruction following | Decision support |
| --- | --- | --- |
| legal | 「準據法寫日本法」→ 照寫 | 分析成本／可執行性／對方接受度 → 建議 + 替代 + trade-off |
| travel-planning | 「星期五出發」→ 照排 | 分析票價／人潮／天氣 → 「星期三出發便宜且避開人潮，要換嗎？」 |
| software-delivery | 「用 DDD」→ 照做 | 分析團隊規模／領域複雜度／演進成本 → Modular Monolith 可能更合適 |
| investment | 「買這檔」→ 照查 | 分析風險／稅／匯率／費用 → 提出配置建議與 trade-off（human selection） |

差別不在輸出多寡，而在**使用者拿到的是決策依據還是執行結果**。

## Generic 管線（六步）

| 步 | 名稱 | 內容 |
| --- | --- | --- |
| 1 | **Requirements** | 使用者真正想達成的結果（不是他指定的做法） |
| 2 | **Domain Analysis** | 該結果在本 domain 受哪些規則／限制約束；哪些可變、哪些不可變 |
| 3 | **Risk Analysis** | 照預設或對方版本做，使用者會承擔什麼；機率與影響 |
| 4 | **Interest Analysis** | 哪種安排對使用者較有利；**以及其他方的真實利益**（可交換什麼） |
| 5 | **Trade-off Analysis** | 每個方案的成本、可行性、阻力、時間 |
| 6 | **Recommendation** | 建議 + 理由 + 替代方案 + 代價 + 需使用者決策的點 |

**第 4 步是最常被跳過的一步。**只算「對使用者最有利」會得出對造／現實不可能接受的方案。
沒有 Interest Analysis 的建議不可執行。

## Decision Reasoning 格式（generic contract）

任何建議**不得**只給結論。每個決策點四欄齊備：

```markdown
### <決策點>
- **Recommendation**: <建議方案>
- **Reason**: <綁定具體因素，不是原則性敘述>
- **Alternative**: <替代方案 1、2>
- **Trade-offs**: <選建議方案放棄什麼；誰會反對；反對時退到哪個 alternative>
- **Confidence**: provisional | confirmed
- **需使用者決策**: <是／否，若是說明決策點>
```

## Two-pass 規則

建議常依賴**尚未查證的前提**（法規現況、票價、效能數據、稅率）。因此本 stage 跑兩趟，
中間夾 domain 的查證步驟：

| Pass | 位置 | 產出 | Confidence |
| --- | --- | --- | --- |
| **Pass 1** | Research 之前 | Decision Register：決策點 + 選項 + 利益分析 + **待查證前提清單** | `provisional` |
| **Pass 2** | Research 之後 | 以已核實前提收斂 + 使用者確認 | `confirmed`（需附來源） |

單趟版本會產出「有理由但前提沒查」的建議——這是本 pattern 明文禁止的失效模式。

## Optimization Suggestion 規則（instruction 與最佳方案衝突時）

使用者已明確指定做法，但分析顯示有更好的選項時：

| 步 | 行為 |
| --- | --- |
| 1 | **講一次**：說明更好的選項、量化差異（成本／時間／風險）、以及使用者原方案的合理之處 |
| 2 | **明確詢問**：要換還是照原方案 |
| 3 | **使用者維持原方案** → 照原方案執行，記錄「已知較優方案 + 使用者選擇維持」，**不重複勸說** |
| 4 | **不確定或未回應** → 依原 instruction 執行，把建議留在產出的建議欄位 |

**邊界**：
- 不得以「為你好」為由**擅自改變**使用者指定的內容。
- 不得把建議重複三次或每個段落都提醒——講一次、記錄、往前走。
- 不得為了展示分析能力而製造無意義的替代方案（optimization theater）。
  差異不顯著時就說「使用者方案合理，無明顯更優選項」。

## Domain 要提供什麼（instantiation contract）

一個 domain 要實例化本 pattern，必須提供四項：

| # | 項目 | 說明 | legal 的實例 |
| --- | --- | --- | --- |
| 1 | **Decision point inventory** | 本 domain 反覆出現的決策點清單 | 九個決策點（準據法／管轄／付款／違約金／IP／保密／驗收／不可抗力／終止） |
| 2 | **Playbooks** | 每個決策點的推理因素、常見結論與**反轉條件** | [`../../legal/strategy/decision-playbooks.md`](../../legal/strategy/decision-playbooks.md) |
| 3 | **Verification source** | Pass 1 的待查證前提由誰查證 | [`../../legal/research/README.md`](../../legal/research/README.md) |
| 4 | **Risk / depth gate** | 什麼情況下**不得**給建議 | [`../../legal/risk-classification.md`](../../legal/risk-classification.md)（Red tier 不出策略建議） |

**第 4 項不可省略。**沒有 depth gate 的 decision support 會在高風險領域產出越權建議。

## Instantiations

| Domain | 實例 | 狀態 |
| --- | --- | --- |
| legal | [`workflow/legal/strategy/`](../../legal/strategy/README.md)（Stage 3a／3b） | ✅ converged case #1 |
| investment | [`workflow/investment/strategy/`](../../investment/strategy/README.md)（Pass 1／2） | ✅ converged case #2（2026-08-14；Phase 1 dogfood＋四項 instantiation；**未**註冊 `route.workflow.investment`） |
| travel-planning | — | candidate（既有 intake 已有分派性質，未抽出 decision support） |
| software-delivery | — | candidate（architecture-fit 分析已有部分素材，未按本 contract 組織） |

Abstraction review（何者可升 generic follow-up、何者 investment-only）：見 investment plan [`evidence/07-phase4-abstraction-review.md`](../../../plans/active/2026-08-14-1101-investment-research-decision-support/evidence/07-phase4-abstraction-review.md)。**本檔不實作 7 個 generic primitive。**

## Anti-patterns

| Anti-pattern | 症狀 |
| --- | --- |
| Ask-only | 把選項原封丟回使用者，不做分析 |
| Conclusion-only | 給建議但無理由、替代方案或 trade-off |
| One-sided interest | 只算使用者利益，方案在現實中不可能成立 |
| Provisional as confirmed | 未查證即標已確認 |
| Playbook as answer | 直接套 playbook 常見結論，未綁本案因素 |
| Repeated pushback | 使用者已決定後仍反覆勸說 |
| Optimization theater | 製造無意義替代方案以展示分析能力 |
| Silent override | 未經同意就改掉使用者指定的做法 |
| No depth gate | 在高風險領域照樣給建議 |

## Promotion 條件

達成下列全部才考慮把 Decision Support 升為各 workflow 的正式 stage 並註冊 route：

- [ ] **3 個** domain 有 converged instantiation（目前 **2／3**：legal、investment）
- [ ] 每個 instantiation 都提供上表四項（含 depth gate）
- [ ] 至少 2 個真實任務證明 Optimization Suggestion 規則未退化成 repeated pushback 或 silent override
- [ ] Generic 管線在 3 個 domain 間未出現分歧定義（否則先拆 bounded context）

未達成前：各 domain 自行實例化並連回本檔，**不**在 `workflow/README.md` 宣稱為全庫 stage。

## 驗證

- 決策點是否四欄齊備？
- Interest Analysis 是否分析了其他方的利益？
- Pass 1 建議是否標 `provisional` 並列出待查證前提？
- Optimization Suggestion 是否只講一次並記錄使用者決定？
- 本 domain 是否有 depth gate（什麼情況不給建議）？
