> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-26 - 掉落列表可能是 Spin↔State reel-set 差分，不是桌上洞；Group 有≠有 reel-set

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來源：通過 State 早退閘後，列表由 per-reel `GetReelSymbols(BySpinInfo|ByStateInfo)` 建集合再差分／allowlist；feature group 可 `GetGroup` 成功但 `GetReelSet(child)` 回 null → state 側空 strip。滿盤仍可出水果 id。專案細節留 docs。

#### One-line Summary

通過早退閘後，掉落／消失候選常來自 **Spin 側 reel-set strip 與 State 側 reel-set strip 的集合差分**（再經字元 allowlist），**不是**讀桌上 `?` 洞。`Groups[name]` 存在仍可能 **ChildGroups→reel-set 無映射** → State 側整軸空陣列。

#### Human Explanation

`GetReelSymbolsBySpinInfo`／`ByStateInfo` 多半轉成 `GetReelSymbols(group, child, reel)` → `GetGroup(group).GetReelSet(child)` → 依 reel-set 名取 strip。Lab 把 State group 設成 feature 名時，若該組的 child 鍵沒有 reel-set，State 回 `[]`，差分≈整份 Spin／base strip，再被 allowlist 收成小寫水果等。Agent 易誤判「滿盤不應掉落」或「缺洞所以有列表是 bug」。

#### Trigger

- 通過 Paid／SpinsCount 閘後 `n>0`，但桌上 `holeCount=0`
- spin 側 strip 很長、state 側每軸 `n=0`
- `GetGroup(feature)` found＝true，但 `GetReelSet(child)` empty／null（試 base child 名仍空）
- allowlist 為連續小寫字母區間，拒高階／特殊符號

#### Evidence

- Tool: Frida on BySpin／ByState／GetReelSymbols／GetGroup／GetReelSet；dump 簽名鏈
- Sanitized excerpt: spin group/child＝base → 非空 strip；state feature group → GetReelSet null → core n＝0；差分∩allowlist＝水果 id
- Evidence path: `<PROJECT_ROOT>/docs/slots/*/lab/` get-drop-symbols-source notes

#### Generalized Lesson

1. **先畫呼叫鏈**：BySpin／ByState → GetReelSymbols(group,child,reel) → GetReelSet → ByReelSetName。
2. **洞不是充分條件**：`holeCount` 與列表可獨立；滿盤＋非空列表在「State strip 空」時合理。
3. **Group found ≠ reel-set**：對 feature group 同時探 `GetReelSet` 多個常見 child；皆 null 則記「組無 strip 綁定」。
4. **Lab force group／child**：把 child 寫成與 group 同名可能仍空；對照 child＝base 名仍空 → 配置問題，非單純誤寫 child。
5. **Allowlist**：靜態字串 `Contains`／區間；高階符號被拒是預期，不是列表源錯。

#### Agent Action

- 通過閘後 hook reel-set 鏈，對照 spin vs state 每軸 `n`。
- 勿要求桌上有洞才接受非空掉落列表。
- 文件寫清：差分語義＋「feature 組無 ChildGroups 映射」根因（若適用）。

#### Goal / Action / Validation

- Goal: 解釋通過閘後列表 id 從何而來
- Action: 探 GetReelSymbols／GetReelSet；記錄 group／child／reelSet／n
- Validation or reference source: state `n=0`＋spin 非空 → 差分∩allowlist 對上列表 API ids

#### Applies When

- Unity IL2CPP slots 以 SpinInfo／StateInfo＋SlotsInfo reel-set 表驅動掉落列表

#### Does Not Apply When

- 列表直接掃 SpinInfo.Symbols 洞或 Rewards payload 集合
- 無 Groups／reel-set 配置模型

#### Validation

- 同次 GetDrop：spin core n≫0、state core n＝0、列表＝allowlist(spin−state)
- GetGroup(feature)=true 且 GetReelSet(*)=null 可複現

#### Promotion Target

- [Unity 動畫／特效驅動追蹤](../../../../analysis/unity/workflows/unity-fx-driver-trace.md) §2.1

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

- category README 追加本條
- `unity-fx-driver-trace.md` §2.1 補 Spin↔State reel-set 差分
- 交叉 [2026-10-26_160018…](2026-10-26_160018-fx-list-api-may-early-out-on-stateinfo-paid-spinscount.md)
- 專案 evidence：`<PROJECT_ROOT>/docs/slots/*/lab/`

#### Transfer Contract

- Inputs: 通過早退閘的列表 API、SlotsInfo reel API
- Mechanism: spin strip − state strip（state 可因無 reel-set 綁定而空）
- Decision: 探 GetReelSet；勿用洞解釋列表
- Verification: 同次 spin／state `n` 對照列表 ids
- Counterexample: 列表純掃 Symbols／Rewards
- Transfer case: 其他 group／child 切換 reel 表的 feature FX

#### Intelligence Review

- Decision: method-only
- Reason: 補強列表填充來源診斷
- Target: analysis/unity/workflows/unity-fx-driver-trace.md
- Revisit: none
- Discovery: not applicable — method-only

#### Abstraction Check

- Swap-project test: PASS — 任何 Spin／State＋reel-set 表驅動的列表 FX 均可套用差分與「Group≠reel-set」檢查
- Project evidence kept in: `<PROJECT_ROOT>/docs/slots/*/lab/`
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
