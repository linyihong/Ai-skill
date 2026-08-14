# Dogfood 06 — Q8 Case A／B／C ＋ Phase 1 總結

| 欄位 | 值 |
| --- | --- |
| Run ID | 06-q8-routing-and-phase1-close |
| Date | 2026-08-14 |
| Goal | Semantic routing 邊界＋收束 Phase 1 H1–H8／Q12 證據 |

---

## Q8 cases（decision object／task semantics）

### Case A

| 欄位 | 內容 |
| --- | --- |
| Utterance | 「我想投資某家公司，幫我分析值不值得」 |
| Decision object | **證券／公司基本面與市場評價** |
| Task type | `name-diligence`（或 need-framing→name-diligence） |
| Route | **`investment`** |
| Result | **PASS**（未被 legal 吸走） |

### Case B

| 欄位 | 內容 |
| --- | --- |
| Utterance | 「我想投資某家公司，幫我看投資協議」 |
| Decision object | **合約／股權權利義務文件** |
| Task type | legal `review`／`draft` |
| Route | **`legal`** |
| Result | **PASS**（不因「投資」字面進 investment） |

### Case C（mixed — 不預設答案）

| 欄位 | 內容 |
| --- | --- |
| Utterance | 「我準備投資這家公司，幫我分析公司本身＋投資協議風險」 |
| Objects | (1) 公司／證券分析 (2) 協議法律風險 |
| 觀察到的分解選項 | 見下表 |
| Result | **MIXED／餵 Q12**（非單一自動路由） |

| 分解策略 | 含義 | Phase 1 觀察 |
| --- | --- | --- |
| Single primary=`investment`，legal 當附件問句 | 易漏協議深度／Red tier | **不足**（協議是獨立 decision object） |
| Single primary=`legal`，公司分析當背景 | 易漏市場 diligence | **不足** |
| **Primary + secondary**（先 framing 拆成兩 deliverable） | investment diligence ＋ legal review 各跑 lifecycle | **Phase 1 建議記錄為此** |
| 正式 multi-domain runtime | 單一 detector 同時兩 active routes | **未實作**；需 Q12／另 plan |

**Case C 裁決（dogfood）**：語意不足以單選 → **先 framing**；產出兩個子任務（investment `name-diligence` + legal `review`），**不是** `contains("投資")→investment`。

### H8

| Result | **PASS（A/B）**；Case C → **Q12 still-open 但有證據** |
| --- | --- |
| Consequence | lean-promote **Semantic Route Disambiguation**；**reject** keyword precedence |

### Q12（Phase 1 證據，不關閉）

| 問題 | Phase 1 證據 |
| --- | --- |
| 要不要 multi-domain runtime？ | Case C 需要 **至少** primary+secondary／雙 lifecycle；是否升 runtime multi-route **仍 open** |
| 下一步 | Phase 2+ 或獨立 routing plan；**本 plan 不偷升** |

---

## Phase 1 總結

### Investment workflow 能不能跑？（形狀層）

| Run | 狀態 |
| --- | --- |
| ① theme-research | done |
| ② name-diligence | done |
| ③ periodic-sweep | done |
| ④ allocation-advice | done |
| ⑤ DVA | done（同分角色模擬） |
| ⑥ Q8 A/B/C | done |

→ **能跑通 Phase 1 實驗室路徑**（尚未建 `workflow/investment/` 檔案與 route — 屬 Phase 2–5）。

### H1–H8 彙總

| H | 跨 run 結果 | Candidate consequence |
| --- | --- | --- |
| H1 Evidence→Decision | **PASS**（含 DVA 強化） | **lean-promote** Evidence-to-Decision Gate |
| H2 DS Lifecycle | **PASS** | **lean-promote** Decision Support Lifecycle |
| H3 Uncertainty Framing | **PASS**（01 MIXED→02/04/05 修正） | **promote framing labels**；**reject** probability-first glossary |
| H4 Decision Depth | **PASS** | **defer**（需更多 depth 階） |
| H5 Observation loop | **PASS**（03） | **lean-promote** |
| H6 Source Authority | **PASS strong**（02/03） | **lean-promote**；**reject** Expert Knowledge |
| H7 State boundary | **PASS** | **defer**（allocation 已用虛構；真實設定檔待 Q10） |
| H8 Semantic routing | **PASS A/B**；C→Q12 | **lean-promote** disambiguation |

### 特別值得帶去 Phase 2 的疤

1. 假精確 %（01／Executor）必須可被 Verifier 殺死  
2. Paywall／C 級轉述易膨脹  
3. Case C 證明 single-route 不夠講「投資」混合句  

### Phase 1 → Phase 2 gate

- [x] ①–⑥ evidence 齊  
- [x] H1–H8 有 PASS/FAIL/MIXED 而非打勾劇  
- [ ] **未做**：實作 `analysis/investment/`、`workflow/investment/`（Phase 2–3）  
- [ ] Q10 schema、Q12 multi-domain runtime — **仍 open**

**Phase 1 status：✅ CLOSED（dogfood lab）** — 可進 Phase 2 方法層實作，**停止**再靠架構討論改 Phase 0。
