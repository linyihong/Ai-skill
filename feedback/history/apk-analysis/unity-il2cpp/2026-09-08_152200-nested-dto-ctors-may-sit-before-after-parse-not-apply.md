> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Nested DTO ctors may sit immediately before/after Parse, not in apply

Status: validated

#### One-line Summary

XML `Parse` 進出沒有巢狀 DTO `.ctor` 時，下一步應看 **同 thread、Decrypt 短窗口的 Parse 前／後**；`Process*` apply 裡也可能完全沒有 nested ctor。

#### Human Explanation

只掛 Parse 進出會漏掉「Parse 前後立刻 `new` 巢狀 DTO」。應把 nested `.ctor` 分成 PRE（Decrypt 窗口且不在 Parse）、IN、POST（Parse leave 後短 dt）、APPLY（Process* 深度>0）。記類別名，不要記次數。APPLY none 不代表沒有 DTO，只代表 apply 用已建好的物件。

#### Trigger

- Parse 窗口 nested `.ctor` 為 none，但 getter 之後看得到那些型別。
- 需要知道物件圖是誰建的。

#### Evidence

- Tool: Frida nested `.ctor` with PRE/IN/POST/APPLY buckets; unique class names only.
- Sanitized observation: nested apply DTOs in PRE and POST; one nested type in IN; APPLY none.
- Evidence path: `<PROJECT_ROOT>/<App>/docs/` protocol notes

#### Generalized Lesson

1. Parse IN none ≠ 沒有 `new`；查 PRE/POST。
2. APPLY none ≠ hook 失敗。
3. 不要 dump ctor 參數或次數。

#### Agent Action

nested ctor 用四桶時間，不要只掛 Parse body。

#### Goal / Action / Validation

- Goal: 定位巢狀 DTO 配置相對於 Parse／apply。
- Action: PRE/IN/POST/APPLY unique 類別名。
- Validation: 至少一個桶非 none，且 APPLY 可為 none。

#### Applies When

- 授權 IL2CPP；Decrypt→XML Parse→Process* apply。

#### Does Not Apply When

- nested ctor 確實只在 Parse IN（照實記錄）。
- 純靜態。

#### Validation

One live round; bucket names only.

#### Promotion Target

- N/A this round（category index only）

#### Required Linked Updates

- `feedback/history/apk-analysis/unity-il2cpp/README.md`
- `feedback/history/apk-analysis/README.md`
