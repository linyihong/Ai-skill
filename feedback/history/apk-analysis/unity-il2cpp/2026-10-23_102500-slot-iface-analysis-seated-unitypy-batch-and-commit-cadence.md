> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-23 - Seated interface-analysis: UnityPy batch + honest splash stub + commit per segment

Status: candidate

#### One-line Summary

裝置已進桌時，先做進桌 §3e（UnityCache UnityPy 盤點＋粒子參數＋L0–L7 map＋遮罩 live），splash 標 deferred stub；每完成一段就 commit／push；線獎材料可能是 feature-pack 專屬名，勿假設只有 `slots.common`。

#### Human Explanation

對已 onboard 的 slot 補 `interface-analysis/` 時，常見兩種浪費：

1. **堅持先做 splash**，即使手機已 seated → 空等或誤把進桌背景當 hero。
2. **只貼 curated PNG** 就當 UI 還原完成，或把線獎粒子當成 common 貼圖拷貝。

成功路徑（本輪驗證）：

| 步驟 | 做法 |
| --- | --- |
| 1 | `adb screencap` → `assets/in-cabinet-live.png`（**遮罩**底欄金額） |
| 2 | pull `UnityCache/Shared/<pack>/<hash>/__data` |
| 3 | UnityPy：`bundleObjectCounts` + `ParticleSystem` 經 `m_GameObject` 取 host + `ParticleSystemRenderer` materials |
| 4 | 寫 `asset-inventory.json`、`particle-system-unitypy-dump.json`、`in-cabinet-resource-map.json`（L0–L7）、animation／fx／audio 配套、`layer-stack.html` |
| 5 | Entrance：若未在 splash，只寫 stub + lobby tile + shared chrome 指針，**不要假造** hero 還原 |
| 6 | `manifest.interfaceAnalysis` 掛齊 §3e 必填欄 |
| 7 | **每段結束 commit／push**（專案 docs 一段；Ai-skill lesson／workflow 另段） |

線獎粒子：host 常仍叫 `coins_fx`／`sparks*`，但 **material 可能是** `slots_<cabinet>_line_payout_fx_coins` 這類 feature 專屬名，不是永遠 `slots_common_line_payout_module_coins_gold`。

成熟度：停輪只有參數 → 標 `params-from-bundle`；未見播映勿標 `live-matched`。

#### Trigger

- 使用者要求「像 bison interface-analysis 一樣分析」且裝置已進桌。
- 補 §3e 時想一次做完 splash＋進桌＋線獎 live。
- 假設所有線獎貼圖都在 `slots.common`。

#### Evidence

- Tool: adb screencap；UnityPy on `__data`；§3e file set；git commit per segment.
- Sanitized pattern: seated → in-cabinet first；ParticleSystem count＋host＋cabinet-specific materials；splash deferred stub；maturity label honesty.
- Evidence path: project `docs/slots/<id>/interface-analysis/`（勿寫機台金額進 lesson）。

#### Generalized Lesson

1. **進桌優先於 splash（若已 seated）**：進桌線可完整；入場線可 stub，標 `not-captured-this-session`。
2. **UnityPy 一批產出 inventory＋粒子參數**：比先翻 AssetStudio／Dumper 快且可機械寫 JSON。
3. **線獎 material 先讀 feature pack**：可能 cabinet-prefixed；再對 common。
4. **成熟度誠實**：`params-from-bundle` ≠ `live-matched`；無線獎幀就寫 not-observed。
5. **分段 commit／push**：專案 `interface-analysis` 一段；可重用技巧進 feedback／workflow 另段，避免大包混 dirty。

#### Agent Action

§3e 任務：確認 plane（splash vs seated）→ UnityPy batch → 必填 HTML／JSON → manifest 索引 → 遮罩 live → commit／push → 若有可重用技巧再寫 feedback lesson 並 push Ai-skill。

#### Goal / Action / Validation

- Goal: 在已進桌條件下快速交可審的 interface-analysis，並沉澱可重用步驟。
- Action: Seated batch above；splash stub；per-segment git close-loop。
- Validation: `manifest.interfaceAnalysis` 齊；particle dump 有 host／material；splash 狀態明示 deferred 或 live；專案與 Ai-skill 各自 clean push。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP slot cabinets with UnityCache feature packs；補 §3e／粒子參數。
- Does not apply: 尚未進櫃的 onboard；純協議／worker 對獎、不做 UI 還原。

#### Related

- `2026-10-23_101000-unity-particle-restore-needs-unitypy-or-assetripper-not-il2cppdumper.md`
- `2026-09-21_091600-slot-three-visual-planes-and-in-cabinet-anim-stack.md`
- [`analysis/apk/workflows/unity-asset-parse-flow.md`](../../../analysis/apk/workflows/unity-asset-parse-flow.md)

#### Validation

新機台 §3e：若裝置已 seated，文件必須有進桌 inventory＋粒子（或 N/A）與 splash 狀態欄；不得把進桌背景標成 splash hero；commit 訊息不混 Ai-skill 與業務 docs。

#### Promotion Target

- [`analysis/apk/workflows/unity-asset-parse-flow.md`](../../../analysis/apk/workflows/unity-asset-parse-flow.md)（補「seated batch / commit cadence」小節）
- Slot onboarding §3e 完成檢查（業務 repo 契約；Ai-skill 只保留通用步驟）

#### Required Linked Updates

- Updated: 本 lesson；`feedback/history/apk-analysis/README.md`；`unity-il2cpp/README.md`
- Updated: `analysis/apk/workflows/unity-asset-parse-flow.md`（seated batch 步驟）
- Not applicable: enforcement rules（行為未改權重）
