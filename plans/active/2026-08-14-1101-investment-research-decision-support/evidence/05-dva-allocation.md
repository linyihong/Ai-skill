# Dogfood 05 — DVA on allocation（O → E → V → 仲裁）

| 欄位 | 值 |
| --- | --- |
| Run ID | 05-dva-allocation |
| Date | 2026-08-14 |
| Loop | Orchestrator → Executor → Verifier → Arbitration |
| Subject | [`04-allocation-advice-fictional.md`](04-allocation-advice-fictional.md) |
| Mode | **同 session 分角色模擬**（dogfood 形狀；非正式多 Task spawn） |
| Goal | 真實爭議：Executor 若輸出「Scenario A **60%**」必須被 Verifier 挑戰 |

---

## Orchestrator brief

| 欄位 | 內容 |
| --- | --- |
| Goal | 產出可交使用者選擇的 allocation 決策支援（A/B/C），遵守策略／費用／證據強度 |
| Acceptance | (1) ≥2 方案＋trade-off＋費用；(2) recommendation strength ≤ evidence；(3) **不得**無校準地把單一方案標成精確勝率 %；(4) Verifier findings 四欄表 |
| Verification（executor 自驗） | 對照 04 intake；引用 02 A 級／03 C 級時標級別 |
| verifier_only | 挑戰任何「60%／高機率最優」類數字；檢查 D 級是否滲入 |
| Constraints | 不下單；虛構持倉；Verifier **不改建議、只產 evidence** |

---

## Executor deliverable（故意含爭議點 — 模擬失效）

> 下列為 **Executor 草稿中的問題段落**（dogfood 注入），不是最終定稿。

**Executor 主張（有瑕疵）**：

> 在虛構使用者約束下，Option A（分批加主題）為最適方案，**Scenario A 成功機率約 60%**；建議直接執行 A。

其餘部分假設 Executor 有抄 04 的方案表與費用（略）。

---

## Verifier report（只產證據，不做決定）

| # | evidence | acceptance_ref | classification | status |
| --- | --- | --- | --- | --- |
| V1 | 04／本 brief：市場語境大量依 C 級（E2）；A 級僅支撐「協議／追蹤」非「分批加碼會成功」 | acceptance (2)(3) | `acceptance-violation` | observed |
| V2 | 「60%」無統計校準來源；違反 uncertainty framing（01 H3 scar、04 已改用 Plausible） | acceptance (3)；verifier_only | `acceptance-violation` | observed |
| V3 | 「直接執行 A」跨越 Human Selection；把 Recommendation 寫成 Decision | acceptance (2)；ERA | `acceptance-violation` | observed |
| V4 | 費用 Interest 與 A/B 比較表若保留則 OK；未見 D 級大神抬轎 | acceptance (1) | `observation` | observed |
| V5 | 方案表 ≥2＋trade-off 形狀存在 | acceptance (1) | `observation` | verified |

**Verifier 結論（非裁決）**：草稿 **未**通過 acceptance (2)(3)；需降級機率宣稱並還原 human selection。

---

## Orchestrator 仲裁

| Finding | 處置 | 理由 |
| --- | --- | --- |
| V1 | **fix** | 刪除「最適／保證方向」；改「Plausible 較優（約束下）」 |
| V2 | **fix** | 移除 60%；改 Likely／Plausible／Unresolved |
| V3 | **fix** | 改「請使用者在 A vs B 選擇」 |
| V4–V5 | **reject**（無需改）／保留 | 正面觀察 |

### 定稿差額（仲裁後）

與 [`04-allocation-advice-fictional.md`](04-allocation-advice-fictional.md) 對齊：**A/B 可討論、C 降級、無 60%、無直接執行**。

---

## DVA 假說檢驗

| 檢驗 | Result |
| --- | --- |
| Verifier ≠ 複讀 Executor | **PASS**（抓到 60%／直接執行） |
| Verifier 不擅自改建議 | **PASS**（只 findings） |
| Orchestrator 才定稿 | **PASS**（fix×3） |
| 形式三角色但無爭議 | **避免**（本 run 注入爭議） |

| H | Result | 一句 |
| --- | --- | --- |
| H1 | **PASS（強化）** | 弱證據無法支撐 60% 決策強度 |
| H3 | **PASS** | 精確 % 被獨立證據否決 |
| ERA | **PASS（形狀）** | Evidence Producer（V）≠ Closure（O）；Recommendation ≠ Decision |

---

## Friction

- 同 session 分角色 ≠ fresh-context Task；正式強制 DVA 時應 spawn 分離 context（記為 Phase 2／工具適配 follow-up，**不**阻塞 Phase 1 形狀證明）。

---

## Writeback to DVA loop plan

跨域契約回饋已寫入：
[`…/4a-investment-decision-support-allocation-dva.md`](../../2026-07-08-0825-delegation-verification-arbitration-loop/evidence/4a-investment-decision-support-allocation-dva.md)
（跨域表列 + evidence index + Phase 2 checklist；**不**填 Knowledge／不偷升 stage 2）。

---

## Next

→ **06 Q8 Case A／B／C**＋Phase 1 總結。
