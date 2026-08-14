# Dogfood 4a — Investment Decision Support allocation DVA（跨域回饋）

| 欄位 | 值 |
| --- | --- |
| Run ID | 4a-investment-ds-allocation-dva |
| Date | 2026-08-14 |
| Source plan | [`2026-08-14-1101-investment-research-decision-support`](../../2026-08-14-1101-investment-research-decision-support/_plan.md) |
| Canonical run | [`…/evidence/05-dva-allocation.md`](../../2026-08-14-1101-investment-research-decision-support/evidence/05-dva-allocation.md) |
| Domain | **Investment / Decision Support**（非 Coding／非 Knowledge） |
| Mode | **同 session 分角色模擬**（**非正式** fresh Task spawn） |

> **紀律**：本 run **不**填 adoption stage 2 的 Knowledge 格；**不**宣稱「Investment 已驗證完整 Delegated Execution」；只回饋 **pattern／契約** 訊號。

---

## 給本 plan 的契約回饋（該吸收什麼）

### 1. 四責任仍成立（topology 可變）

| 責任 | Investment 實例化 |
| --- | --- |
| Specification | Orchestrator brief（acceptance：證據強度門檻、禁未校準 %、保留 human selection） |
| Production | Executor allocation 草稿 |
| Independent Evidence | Verifier findings（四欄） |
| Decision / Arbitration | Orchestrator fix／defer／reject → 定稿 |

角色名仍是 O／E／V，但 **acceptance 量尺是「recommendation strength ≤ evidence」與「Recommendation ≠ Decision」**，不是 code green。→ 支援「pattern ≠ topology」：Decision Support 消費同一 loop，量尺換成 DS／ERA。

### 2. `verifier_only` 對「未校準機率」有效（L3 類）

Brief 預標：挑戰「60%／高機率最優」。Executor 注入該失效 → Verifier 打 `acceptance-violation` → Orchestrator **fix**。

**回饋**：跨域 brief 應把「禁止無校準精確勝率／把建議寫成決策」列為 **verifier_only**（或同等 acceptance），否則 Verifier 易只做 L1 形狀檢查而漏過數字幻覺。

### 3. ERA 訊號：Evidence 約束 Decision Space；Recommendation ≠ Selection

Verifier 抓到三類越界：

| 失效 | ERA 對應 |
| --- | --- |
| C 級證據卻寫「最適」 | Decision strength > evidence strength |
| 「60%」無校準 | 假精確把 uncertainty 壓成點估計 |
| 「直接執行 A」 | Recommendation 僭越 Human Selection／Closure |

→ 與 ERA v2（Evidence → Decision Space → Preference → Selection）**同構**；Investment 多露出 **Recommendation** 這一層，值得在跨域敘事中顯式分開。

### 4. 必須誠實降級的限制（反膨脹）

| 限制 | 含義 |
| --- | --- |
| 同 session 分角色 | **≠** fresh-context Independent Verification（對照 2o／3a） |
| 無 L1「重跑命令」 | Allocation artifact 是 markdown；Verifier 主要是 **L2 讀稿＋L3 verifier_only** |
| 爭議為注入 | 證明「有爭議時 loop 能修」；**不**證明生產環境自然產出爭議率 |

若把 4a 寫成「Investment 完整三角色已驗證」→ **overclaim**（與 2w Travel「消費協議、不偷升」同紀律）。

### 5. 對 kit／delegated-execution 的可選補強（doc-only candidate）

- Acceptance 模板加一條：**uncalibrated probability / success-% forbidden unless source-backed**  
- Acceptance 模板加一條：**recommendation must leave human selection explicit**  
- Cross-domain 表加 Investment／DS 列（本檔）  

**不**因此動 schema／runtime。

---

## 量測欄（精簡）

| 指標 | 值 |
| --- | --- |
| Findings | 5（violation×3，observation×2） |
| 仲裁 | fix×3 |
| Verifier 改建議？ | 否 |
| Fresh Task spawn？ | **否** |
| Pattern-held？ | **是（降級宣稱）** |

---

## 與跨域表的建議落點

| Domain | Production | Evidence | Decision | 建議證據狀態 |
| --- | --- | --- | --- | --- |
| **Investment／Decision Support** | Allocation／research author | Independent DS verifier（證據強度／uncertainty／human selection） | Orchestrator／使用者 | **形狀已跑（4a）— role-sim；非 fresh-Task 完整驗證** |

**不**填 Knowledge；**不**推進 stage 2 至 3/3。
