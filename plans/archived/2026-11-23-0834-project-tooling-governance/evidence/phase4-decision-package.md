# Phase 4 裁決資料包

- **日期**：2026-11-24
- **開放**：使用者 2026-11-24「開 Phase 4」。
- **用途**：把 Phase 0–3、F1/F2/F6、S2 和三個 consumer 案例的證據，整理成 Phase 4 清單要裁決的事項。本檔只整理證據和建議；裁決結果寫在 §6。

## 1. 案例清單

| # | 案例 | 性質 | 由誰改 | 狀態 | 證據 |
| --- | --- | --- | --- | --- | --- |
| K1 | Ai-skill 工具樹 | 實際套用（本計畫） | 本計畫 | 已推送；T1/T3/T4 缺口已由 F6/F2/F1 修復 | [`phase3-application.md`](phase3-application.md)、F1/F2/F6 |
| K2 | Consumer A 唯讀評估 | 唯讀 | — | 完成 | [`phase2-readonly-evaluation.md`](phase2-readonly-evaluation.md) |
| K3 | Consumer B 唯讀評估 | 唯讀 | — | 完成 | 同上 |
| K4 | Consumer B CLI 試點 | 實際套用（consumer 自己決定做） | Consumer B | 已合併，並在正式 pre-push 路徑上執行 | [`consumer-b-cli-pilot.md`](consumer-b-cli-pilot.md) |
| K5 | Consumer A 暫存登記框架 | 實際套用（consumer 自己決定做） | Consumer A | 已合併 | [`consumer-a-temp-registry.md`](consumer-a-temp-registry.md) §1–§5 |
| K6 | Consumer A 依 brief 修改 | 實際套用（**照本契約推導出的清單**） | Consumer A 的 agent | 已推送；重新檢查通過。限制：鎖不在版控內、R1 沒在真機驗證 | 同上 §7 |

## 2. 每條契約 invariant 的狀態

✓ = 成立；◐ = 成立但有附註；✗ = 有反例；— = 不適用或沒涵蓋。

| Invariant | K1 Ai-skill | K2 A 唯讀 | K3 B 唯讀 | K4 B 試點 | K5 A 暫存 | K6 A brief | 需修改？ |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 三維度分離（Role / Placement / State 不互推） | ✓ catalog 18 列 | ✓ 可判定 | ✓ 可判定 | — | ✓ | — | 否 |
| Discovery：先找再建 | ◐ | ✗ justification 為 0 | ◐ 有用途、少有「為何不用既有」 | ✓ 沿用既有 attestation | ✓ 既有 helper 改成薄包裝 | ✓ 改用共用裝置模組 | 否（可判定，consumer 採用度低） |
| Role：Entry / Logic / Parameters | ✓ 無違規 | ✓ 抓到 R1–R4 | ◐ R6 歧義（A-2，已在 Q12 釐清） | ✓ | ✓ | ✓ R1–R4 修完 | 否；**候選 C3**：Parameters 放在工具樹外 |
| Lifecycle：Placement × State、Scratch、TTL | ✓ | ◐ G4 張力（已修訂） | ✓ 正例（`expires_at`） | — | ◐ Scratch 符合；G4 符合 2／3 | — | **候選 C1、C2**（單一來源） |
| Ownership | ✓ catalog 有 owner 欄 | ✗ 沒有 owner | ✗ 沒有 owner | — | ✗ 沒有 owner | — | 否：契約可判定，但兩個 consumer 的採用度都是 0，**這條在外部沒有收斂** |
| Growth rule：concept-first、使用證據第 1／2／3 次 | ✓ F2 改成宣告式命令表 | ✓ 概念分組 | ✓ | ✓ 7 閘收斂 | ✓ 一個能力取代多處 | — | 否 |
| Lock taxonomy：四欄形狀、Review Rule、雙向樣本 | ◐ 7 把已審；T1/T3/T4 缺口已修 | ✓ 可判定 | ✓ 可判定 | ✓ | ✓ 雙向測試 | ✓ 雙向測試；allowlist 附移除條件 | **候選 T-B4、T-B5、C5、C6** |

**結論**：沒有任何 invariant 被推翻。外部案例沒有覆蓋到、或覆蓋了但沒收斂的，是 Ownership（採用度 0）和 Lifecycle 的機械清理邊界（只有 K5 一例，而且不完全符合）。

## 3. Q11：哪些案例算「收斂案例」

| 選項 | 計數規則 | 結果 |
| --- | --- | --- |
| (a) 唯讀評估也算 | K1 + K2 + K3 | 3／3 |
| (b) 只算實際套用；consumer 自己決定做的也算；同一個 consumer 只算一次 | K1、B（K4）、A（K5 + K6） | **3／3** |
| (c) 只算照本契約推導出來的實際套用 | K1、K6 | 2／3 |

**建議 (b)**。理由：
- 唯讀評估只證明「契約判斷得出來」，沒證明「照契約改得動」。這正是 Q11 開題時點出的差別。
- K4、K5 雖然不是本計畫推動的，但都在外部 repo 合併，而且符合契約的 invariant。把它們排除，就等於要求案例必須由本計畫推動，這不是收斂門檻的原意。
- 計數附帶條件：每個案例都要標出涵蓋了哪幾份契約（見 §2）。**3／3 不等於整份契約都收斂**：Ownership 和機械清理邊界在外部仍未收斂。

## 4. 契約修改候選

| # | 候選 | 來源數 | 建議 |
| --- | --- | --- | --- |
| T-B4 | 失敗結果必須附修復動作，**包含鎖自己無法執行的情況** | 2（B、Ai-skill） | **接受**，寫進 lock-taxonomy 的 result contract |
| T-B5 + C5 | allowlist 用逐一列舉名稱維護是反例；每筆例外要附移除條件 | 列舉反例 2 個來源（B、Ai-skill）；附移除條件的正例 1 個來源（K6） | **接受**，合併成一條 allowlist 規則 |
| C6 | 不進版控的鎖只對那台機器有效；契約要求標明鎖的強制範圍（版控內 / 只在本機） | 1（K6，但涵蓋 Consumer A 全部鎖） | **接受為澄清**：不是新規則，只要求盤點鎖時標明這一欄 |
| 比對層級 | pattern-based 的鎖要宣告比對層級（文字 / AST / 狀態集合） | 1（Ai-skill 內 2 把鎖） | 觀察 |
| C1、C2 | Scratch 條款區分工具和資料；G4 第 3 條對可重新產生的輸出是否過嚴 | 1 | 觀察 |
| C3 | Parameters 放在工具樹外的資料目錄 | 1 | 觀察 |
| C4 | validator 會讀的 key 要保留命名空間 | 1 | 已寫成 feedback lesson（candidate），契約不改 |
| C7 | 外部 OCR 程式算不算裝置 I/O | 1 | 屬專案自己的定義，契約不改 |
| O4 | 單次 `[skip-*]` 豁免是否要附理由或計數 | 1（Ai-skill） | 觀察 |
| T-B2 | 是否在 Ai-skill 補量測 hook 的重複 I/O | 1 次量測 + 1 次觀察 | 不補；維持候選 |

## 5. Q5、Q7、runtime 接入

| 項目 | 證據 | 建議 |
| --- | --- | --- |
| Q5 共用 runner | T-B1：B（Node）和 Ai-skill（Go）各自用自己的語言實作單一入口 + registry，都可行；K6 用 Python | **維持 deferred**：沒有共用 runner 的需求證據 |
| Q7 discovery 機械化 | Discovery 可判定，但兩個 consumer 的 justification 採用度都很低；沒有任何案例證明機械化 discovery 有效 | **維持 deferred** |
| runtime 接入（discovery signal 或 `init-project` 範本） | 依 (b) 達到 3／3，但 Ownership 和機械清理邊界在外部未收斂 | **本計畫不接入**；維持 doc-only pilot，狀態從「0／3」更新為依 Q11 的計數。真要接入，另開 plan 並附 Per-surface consumer 表（原 entry condition）。Graduation deadline 2027-02-28 不變 |

## 6. 裁決（使用者 2026-11-24，全部採用建議）

| 事項 | 裁決 |
| --- | --- |
| Q11 | **只算實際套用**；consumer 自己決定做的也算；同一個 consumer 只算一次。結果 **3／3**（K1、B、A）；每案標明涵蓋的契約；Ownership 與機械清理邊界在外部仍算未收斂 |
| 契約修改 | **T-B4、T-B5+C5、C6 接受**，寫進 `lock-taxonomy.md`；其餘候選維持觀察（見 §4） |
| Q5、Q7、runtime 接入 | **全部維持 deferred**；doc-only pilot；要接入時另開 plan；graduation deadline 2027-02-28 |
| 結案 | **結案**；§7 遺留項目移到後續計畫（只建立、不開始） |

## 7. 不屬於本契約的遺留項目

以下是過程中發現的 Ai-skill 自身治理問題。它們不是 Project Tooling 契約的裁決事項，計畫結案時要另外安排去處：

| # | 項目 | 來源 |
| --- | --- | --- |
| R1 | rule_class 有 46 個，超過上限 40 | F1 |
| R2 | `executor_kind_spec.enum` 沒有 lint 檢查 | F1 |
| S1 | core-bootstrap 的 severity 欄位只宣告、dispatcher 不執行 | F2 |
| S2 | `CLIDocSyncRule` 要不要在可攜層 deprecated（盤點已完成：沒有外部使用者） | F2、S2 盤點 |
| V2 | `runtime validate` 的既有腳本檢查只看 `.sh` | F6 |
| — | `warnCLIDocSync` 2026-12-24 再評估，前提是先決定執行紀錄的來源 | F2 §7 |

← [Evidence index](README.md)
