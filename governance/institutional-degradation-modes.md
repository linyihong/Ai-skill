# 制度退化模式（Institutional Degradation Modes）

本檔記錄這套制度**最可能的退化方式**，每種附**偵測訊號**（怎麼發現正在發生）與**預防機制**（怎麼擋）。目的：當長期由較弱模型運作時，制度會慢慢從「活的約束」腐化成「形式化的擺設」——本檔讓退化可被提早發現。

> **為什麼需要**：規則不會一次崩壞，而是逐步空心化。每個退化模式都有早期訊號；漏看訊號，等到明顯時已經積重難返（本 plan 的起點——learning report 淪為只申報不執行——就是 D1 已經發生的證據）。

## D1 — 報表化（Report Hollowing）

**定義**：一個 obligation 從「產生真實行為」退化成「填格式」——欄位都填了，但語義流失，沒有對應的實際動作。

**偵測訊號**：
- 某個 report / obligation 的欄位長期是同一個「安全值」（例如 learning report 永遠 `FeedbackDecision: NONE`）。
- 申報層有 gate（驗格式），執行層沒有 gate（不驗有沒有真做）。

**預防機制**：申報必須綁一個可機械對帳的執行痕跡（例如 `Writeback: COMPLETED` 要能對到當輪 git diff）。本 plan Workstream A（deferred-feedback ledger）+ Q6（假 COMPLETED cross-check）就是 D1 的解法。

- **正例（健康）**：learning report 說 COMPLETED，git 有對應 feedback-history 的 diff。
- **反例（退化）**：連續 20 個 session 的 report 都填 NONE，但這期間明明有多次可沉澱的教訓——欄位在，語義死了。

## D2 — 繞過誘因累積（Bypass Incentive Buildup）

**定義**：遵守成本高到某個點，agent / 使用者開始找繞過方式；一旦繞過一次成功且無懲罰，繞過變常態，gate 形同虛設。

**偵測訊號**：
- `[skip-*]` opt-out marker 使用頻率上升。
- `--no-verify` 或等效繞過出現在 commit 歷史。
- Bootstrap 稅（每 session 固定注入量）持續變大（P8）。

**預防機制**：
- 繞過必須留痕且可審計（本 repo 的 pre-push governance replay 就是擋 `--no-verify` 繞過）。
- 遵守成本本身要被當債管理——bootstrap 分級（Q9）是為了降低繞過誘因，不是為了偷懶。

- **正例（健康）**：某 validator 太吵，走 Status Transition Matrix 正式降級 + 記錄理由。
- **反例（退化）**：因為某 validator 常擋，大家習慣性加 `[skip-*]`，半年後該 validator 實際上從沒生效過。

## D3 — 宣告面再膨脹（Declaration Surface Re-inflation）

**定義**：宣告了 route / surface / scenario / rule，但沒有 consumer 消費它；orphan 累積，形成「假治理感」——看起來很完整，實際沒生效。

**偵測訊號**：
- `ai-skill runtime audit` 的 orphan 總數上升（本 plan 基線 314，見 P2）。
- 新增 route / surface 但沒同時 wire discovery signal 或 validator。

**預防機制**：
- orphan 數採 ratchet（只降不升，本 plan E2）。
- 新宣告面必須在同一 plan 宣告 named consumer（`validateRuntimeTriggerWiring` 已機械擋新增未 wire 的情況）。

- **正例（健康）**：新增 route 時同批 wire 一個 discovery signal，audit 判 consumed。
- **反例（退化）**：為「架構完整性」新增 20 個 route，全部 orphan，半年後沒人記得它們做什麼。

## D4 — Rubric 腐化（Rubric Rot）

**定義**：判斷 rubric / 正反例 / 參數表過時沒人更新，弱模型照著過時的判準做出錯誤決定，卻以為自己合規。

**偵測訊號**：
- rubric 引用的 model 名 / 工具參數已不存在（見 F2 §參數保鮮 / Q11）。
- 正反例引用的檔案路徑失效（link check 報 broken）。
- rubric 的判準與現行 gate 行為矛盾。

**預防機制**：
- 參數類內容標 `unverified` 而非硬編（[`models/routing/model-tier-escalation.md`](../models/routing/model-tier-escalation.md) §參數保鮮）。
- rubric 更新責任綁 F5 edit-authority + deferred ledger（踩坑教訓回流）。
- 定期 link check（本 plan E1 scenario runner 可兼跑）。

- **正例（健康）**：新 model 上市時，F2 參數表同步更新，舊 model ID 標 deprecated。
- **反例（退化）**：rubric 還在教「用 model X」，但 model X 已退役，弱模型指定它 → 委派失敗且不知為何。

## D5 — 判斷降級成關鍵字（Judgment Collapse to Keyword-Matching）

**定義**：弱模型把需要判斷的規則簡化成關鍵字比對，繞過了規則的真實意圖。

**偵測訊號**：
- 規則用了「適當 / 必要 / 重要」等未定義觸發詞（[`weak-model-rule-authoring.md`](weak-model-rule-authoring.md) 要擋的），弱模型只好瞎猜。
- 同一類情況，不同 session 的處理天差地遠（判斷不穩）。

**預防機制**：
- 所有可執行規則走 F1 四要件（觸發條件必須可機械判斷）。
- 需要真判斷、無法降成關鍵字的 → 明文標 `behavioral-only` + escape hatch（不假裝可機械化）。

- **正例（健康）**：規則寫「staged diff 含 `enforcement/*.md`」（可機械判斷）。
- **反例（退化）**：規則寫「處理重要的 enforcement 變更時」，弱模型把「重要」理解成「檔案大」，漏掉小但關鍵的改動。

## 綜合偵測（給未來 session 的定期自檢）

每隔一段時間（建議月度，或接手一個久未維護的 repo 時）跑一次：

- [ ] `ai-skill runtime audit` orphan 數 vs 上次（升 → D3）
- [ ] `ai-skill enforcement coverage` behavioral 數 vs 上次（只升不降 → ladder 卡住，關聯 D2）
- [ ] grep commit 歷史的 `[skip-*]` / `--no-verify` 頻率（升 → D2）
- [ ] learning report 的 FeedbackDecision 分佈（長期全 NONE → D1）
- [ ] Markdown link check（broken → D4）
- [ ] rubric / 參數表的 model 名是否仍有效（失效 → D4）

## 誰會參考這裡（Inbound References）

- [`plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md`](../plans/active/2026-07-10-0955-final-form-feedback-execution-closure.md) — Workstream F / Phase 5b F6；北極星三性質正是這些退化模式的反面
- [`../architecture/ai-native-cognitive-ecosystem-system.md`](../architecture/ai-native-cognitive-ecosystem-system.md) — §Watch-Out List（process bloat / premature abstraction）與本檔互補

## 與既有層的關係

- [`../enforcement/failure-patterns/`](../enforcement/failure-patterns/README.md) — 管**具體一次失效**如何變 prevention gate；本檔管**整套制度的慢性退化模式**（更高一層）。具體 failure pattern 累積到成模式時，回連本檔。
