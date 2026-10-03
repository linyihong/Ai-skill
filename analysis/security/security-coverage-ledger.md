# Security Coverage Ledger（資安覆蓋帳本契約）

**Status**: `candidate-analysis`（doc-only trial；plan [`2026-10-03-2104-security-audit-capability-hardening`](../../plans/archived/2026-10-03-2104-security-audit-capability-hardening/_plan.md) Phase 3）

## 目的

讓 `security-audit` 能回答三件事：**哪些攻擊面已檢查、哪些沒檢查、哪些過去的檢查已經失效**。本文件只定義契約；帳本資料屬於各專案，存在專案 repo，**不**進 Ai-skill 或 `runtime/runtime.db`。

每次 audit 的輸出（finding list）見 [`security-finding-list-template.md`](../../workflow/software-delivery/templates/security-finding-list-template.md)；本帳本是跨多次 audit 累積的覆蓋狀態，finding list 的 `audit_execution.coverage_ref` 指向這裡的 unit id。

## 何時使用

- `security-audit` 以 `standard` 或 `deep` 模式執行，且專案預期會反覆修改同一批敏感面。
- 需要增量 audit：只重驗受變更影響的部分，而不是每次從零開始。
- `light` 模式可不建帳本；若專案已有帳本，light audit 仍要檢查本次 diff 是否觸發失效（見下方 Invalidation）。

## Coverage Unit

一個 unit 是一個可追蹤的檢查單位：

```text
Entry Surface × Trust Boundary × Attack Class   （Subsystem 為選填第四軸）
```

| 欄位 | 意義 | 範例（generic） |
| --- | --- | --- |
| `entry_surface` | 外部輸入進入系統的位置 | API route、message consumer、file upload handler、CLI input |
| `trust_boundary` | 跨越的信任邊界 | `anonymous→authenticated`、`authenticated_user→other_user_resource`、`tenant→other_tenant` |
| `attack_class` | 檢查的攻擊類別 | `idor`、`authz_bypass`、`injection`、`ssrf`、`secret_exposure` |
| `subsystem` | 選填；只有 entry surface 不足以分辨責任區時才用 | billing、media delivery |

**未列入帳本的組合 = 未檢查**。不可因為「同模組其他 unit 已覆蓋」就推定為已覆蓋。

## Unit 狀態

| Status | 意義 | 可轉移到 |
| --- | --- | --- |
| `not_covered` | 已知存在、尚未檢查 | `covered`、`accepted_gap` |
| `covered` | 有 audit 證據，且依賴的 control 自證據日起未變 | `needs_revalidation` |
| `needs_revalidation` | 先前的證據可能已失效（見 Invalidation） | `covered`、`not_covered` |
| `accepted_gap` | 刻意不檢查，有 decision record | `not_covered`（decision 失效時） |

`covered` 必須附 `evidence_ref`（指向 finding list 或可重現證據）與 `verified_at_ref`（commit 或版本），不可只寫日期。

## 專案端資料格式（YAML，存在專案 repo）

```yaml
# <PROJECT_ROOT>/docs/security/coverage-ledger.yaml（路徑由專案決定）
units:
  - id: SEC-COV-<n>
    entry_surface: <surface>
    trust_boundary: <boundary>
    attack_class: <class>
    subsystem: <optional>
    status: covered               # not_covered | covered | needs_revalidation | accepted_gap
    evidence_ref: <finding list / test / scan ref>
    verified_at_ref: <commit or version>
    depends_on_controls:          # 必填（covered 時）；見 Dependency Scope
      - <shared control id>
    source_scope:                 # 直接相關的入口檔案或模組
      - <path or module>
    decision_ref: <accepted_gap 時必填>

controls:
  - id: <control id>
    kind: authorization_policy    # authorization_policy | authentication | query_filter | middleware | input_validation | secret_store | other
    source_scope:                 # 必須是路徑或模組；「某某設定」這類描述無法偵測變更
      - <path or module>
    detection: <optional: command or check that shows the control changed, e.g. grep for new consumers>
```

Phase 1–4 不提供 Ai-skill validator；格式依本文件人工 / agent 檢查。

## Dependency Scope

檔案 hash 不足以判斷證據是否有效。入口檔案沒改，但它依賴的共用 control 被改了，先前的授權檢查結果一樣可能失效。

因此每個 `covered` unit 必須列出 `depends_on_controls`：

- 以 **trust boundary** 為主鍵思考：這個邊界由哪些 control 守住？
- 列出所有參與守門的共用層，而不只 controller：authorization policy / handler、resource-based authorization、query filter（例如 owner / tenant scoping）、middleware、service 層 guard。
- 不確定是否參與守門時，列入；寧可多失效一次，不要漏失效。

## Invalidation（失效規則）

變更發生時，依序判斷：

| 觸發 | 結果 |
| --- | --- |
| 變更觸及某 control 的 `source_scope` | 所有 `depends_on_controls` 含此 control 的 unit → `needs_revalidation` |
| 變更觸及某 unit 的 `source_scope` | 該 unit → `needs_revalidation` |
| 新增 entry surface 或 trust boundary | 新增 `not_covered` unit |
| control 被移除或改由其他層負責 | 依賴它的 unit → `needs_revalidation`，並更新 `depends_on_controls` |
| 原本沒有使用者的內部介面出現第一個使用者（注入、呼叫、endpoint） | 依賴「入口」control 的 unit → `needs_revalidation`。入口以**第一個使用者**為準，不以註冊為準：依命名慣例或 assembly scanning 的註冊常常早已存在 |
| risk acceptance 的 `expires_when` 所列 control 被變更 | 該 acceptance 失效；相關 finding 回到 open |

`needs_revalidation` 的 unit 在重驗前不得被 finding list 的 `audit_execution` 引用為已覆蓋。

### 機械回放

失效判斷可以機械執行，不需要靠 agent 判讀：

1. 解析 ledger，取得每個 control 與 unit 的 `source_scope`，以及每個 unit 的 `depends_on_controls`。
2. 取得變更檔案清單（例如 `git diff --name-only <from> <to>`）。
3. `source_scope` 依**路徑邊界**比對：條目與變更路徑完全相同，或變更路徑以「條目 + `/`」開頭。只做字串前綴比對會讓 `src/Payment` 誤中 `src/PaymentHistory.cs`。
4. 被觸及的 control → 依賴它的 units；加上自身 `source_scope` 被觸及的 units → 全部標 `needs_revalidation`。

粒度取捨：以檔案或目錄為範圍會 over-invalidate（同一檔案承載兩個 control 時，只改一部分也會兩邊失效）。這是刻意的，漏失效的代價高於多重驗一次；ledger 變大後若成本過高，再把 control 拆到更細的檔案，而不是改用較寬鬆的比對。

這是 [`stale-derived-state.md`](../../intelligence/engineering/anti-patterns/stale-derived-state.md) 的 `stale_security_evidence` 變體：過去的 audit 結論是 derived state，source（control 實作）改變後必須有 invalidation contract。

## 執行目標程式碼的隔離要求

需要執行目標程式碼來確認漏洞時，必須有 OS 層隔離：限制網路、不暴露真實環境變數與 secret、限制可寫路徑與資源。缺任一條件時不執行；finding 維持 `needs_validation`，unit 不升級為 `covered`。

## 去敏邊界

- 帳本與證據存在專案端；寫回 Ai-skill 的只能是抽象化後的 pattern（見 [`enforcement/reusable-guidance-boundary.md`](../../enforcement/reusable-guidance-boundary.md)）。
- 不寫入 host、endpoint、token、真實使用者資料或專案路徑。

## 與其他層的關係

- 每次 audit 輸出 → [`security-finding-list-template.md`](../../workflow/software-delivery/templates/security-finding-list-template.md)
- Closure gate → [`execution-flow.yaml`](../../workflow/software-delivery/execution-flow.yaml) `gate.software_delivery.security_audit_complete`
- 失效的抽象原則 → [`stale-derived-state.md`](../../intelligence/engineering/anti-patterns/stale-derived-state.md)

← [回到 analysis/security/](README.md)
