# Security Finding List: <feature / change>

> **`security-audit` capability output**（registry artifact `security-finding-list`）。由 `sd-contracts` 或 `sd-implementation` invoke（`stance: fault_finding`）後產出；不是 Validation phase output。Invoke 時機見 [`cross-cutting/review/invocation-points.md`](../../cross-cutting/review/invocation-points.md)。

## 三個判斷分開

| 判斷 | 問題 | 誰回答 |
| --- | --- | --- |
| Schema valid | 本文件欄位是否齊全、enum 是否合法？ | 格式檢查 |
| Evidence established | 每個 finding 的 `status` 是否有足夠證據支撐？ | Verifier（V3）+ 可重現證據 |
| Merge / closure allowed | 依 gate 規則可否完成交付？ | `gate.software_delivery.security_audit_complete` |

格式通過不代表證據成立；證據成立不代表可以合併。

跨域 pattern：[`traceable-evidence`](../../cross-cutting/traceable-evidence/README.md)。`evidence.reproducible` 實作 TE2；Status 規則與 `resolution` 實作 TE4；本節的三判斷分開實作 TE5。

## Audit Execution（必填）

`findings: []` 只在 `status: completed` 時代表「已執行、未發現符合條件的 finding」。缺本段或 `status` 不是 `completed` = **unknown**，不是 safe。

```yaml
audit_execution:
  status: completed          # completed | partial | not_run
  mode: light                # light | standard | deep（對照下方 Audit Mode）
  scope: diff                # diff | module | trust_boundary | full
  scope_ref: <commit range / PR / module list>
  coverage_ref: <project coverage unit id, or none>   # 見 analysis/security/security-coverage-ledger.md
  evidence_ref: <where the audit trace / tool output / notes live>
  execution_environment: >   # 有沒有執行目標程式碼、有沒有隔離環境；沒執行就寫明證據來自既有紀錄
    <e.g. no isolated DB; target code not executed; reproducible evidence cites recorded run X>
  not_covered:               # 明列未檢查的面，不可省略為空泛「其他」
    - <entry surface / attack class not examined, or none>
```

`not_covered` 只列「範圍外」或「沒時間看」的面，**不能**用來把範圍內的缺陷歸類掉。常見誤用：把 tracked 設定裡的金鑰寫成「secrets 不在範圍」、把可預測的驗證碼寫成「rate limit 不在範圍」。餵給受稽核 control 的設定（簽章金鑰、驗證開關、proxy header 規則）屬於範圍內。

標示為 placeholder / stub / 「正式上線前替換」的實作，是必查線索：確認它有沒有被正式路徑使用，以及是否真的被呼叫。

### Reachability 宣稱

「沒有呼叫者」「沒有註冊」「外部不可達」會直接影響 severity 和失效條件，所以要分開寫，且各自有依據：

- **已註冊**：除了搜尋型別名稱，還要檢查依命名慣例、assembly scanning、反射或設定檔的註冊；只 grep 型別名會漏掉這些。
- **有使用者**：介面在擁有者以外是否被注入、呼叫或掛到 endpoint。
- 失效條件綁「第一個使用者出現」，不要綁「被註冊」——慣例註冊常常早就存在。

## Findings

每個 finding 一個 block。未確認的 finding 不填 `severity`；confirmed 的 finding 不填 `potential_impact`（嚴重度已由 `severity` 表達，未來可能升級的理由寫在註解或 resolution）。

```yaml
findings:
  - id: SEC-<n>
    status: candidate        # candidate | confirmed | refuted | needs_validation
    attack_class: <e.g. idor, authz_bypass, injection, secret_exposure, ssrf>
    entry_surface: <API route / handler / message consumer / CLI input>
    trust_boundary: <e.g. authenticated_user→other_user_resource>
    subsystem: <optional>
    severity: null           # 只在 status: confirmed 時填 critical | high | medium | low
    potential_impact: high   # critical | high | medium | low — 未確認時的潛在影響
    unresolved_fact: >
      <尚未確認的事實；confirmed / refuted 時寫 none>
    evidence:
      reproducible:          # test / SAST / mutation / schema validator / runtime trace
        - <ref>
      reasoning:             # 讀碼推理；不能單獨構成 confirmed / refuted
        - <ref>
    hypothesis_source: <intelligence ref, or none>   # 只提供檢查假設，不是裁決證據
    resolution:
      kind: none             # none | fixed | human_review | risk_acceptance
      ref: <fix commit / review record / decision record>
      reviewer: <human_review 時必填>
      risk_acceptance:       # kind: risk_acceptance 時必填
        decision_ref: <project docs/decisions/... >
        owner: <project decision owner>
        scope: <accepted finding scope>
        expires_when: <depended-on security control(s) whose change voids this acceptance>

examined_no_finding:         # 檢查過但沒發現問題的面；給 verifier 反駁用，不可省略
  - topic: <what was examined>
    note: <evidence that it holds; mark which part is test-backed and which is static reasoning>
```

### Deferral ≠ risk acceptance

計畫或 brief 寫「延後到下一個切片」只是 **deferral**：它說明何時處理，不代表有人接受風險。只有決策者在 decision record 明確接受風險時，才能填 `resolution.kind: risk_acceptance`；`owner` 必須是該 record 裡的決策者，不能由稽核者推定。只有 deferral 時用 `kind: none`，並在 `ref` 指出 deferral 與「必須在何時之前處理」。

### Status 規則

| Status | 需要的證據 | `severity` |
| --- | --- | --- |
| `candidate` | 發現者的假設；尚未驗證 | 不填 |
| `needs_validation` | 已檢查但缺可重現證據或缺執行環境；`unresolved_fact` 必填 | 不填；用 `potential_impact` |
| `confirmed` | 至少一個 `reproducible` 證據 | 必填 |
| `refuted` | 至少一個 `reproducible` 反證；LLM 第二意見不足以 refute | 不填 |

- 發現 finding 的 agent 不能同時擔任該 finding 的 verifier（見 [`delegated-execution.md`](../delegated-execution.md) §5）。
- `hypothesis_source` 指向的歷史 intelligence 只能讓 agent 去檢查某個假設；裁決只看當前原始碼與證據。
- 無 OS 層隔離（網路、環境變數、可寫路徑、資源限制）時，不執行目標程式碼來確認漏洞；finding 維持 `needs_validation`。

## Audit Mode（對照 Cognitive Mode，不是新機制）

| Mode | 典型觸發 | Cognitive Mode 對照 | 最低要求 |
| --- | --- | --- | --- |
| `light` | 一般程式修改，未觸及敏感面 | `FAST` / `NORMAL`，governance `LIGHT` / `STANDARD` | diff-scope `audit_execution` + 既有規則檢查；可為空 list |
| `standard` | 修改授權、認證、資料存取、secret、付費 / entitlement 等敏感面 | `NORMAL` / `DEEP`，governance `STRICT` | trust boundary 分析、針對性檢查、獨立 verifier |
| `deep` | 正式稽核或高風險變更 | `DEEP` / `FORENSIC`，governance `STRICT` | coverage 單位逐一標記、多輪驗證、報告 |

Cognitive Mode 定義見 [`models/cognitive-modes/README.md`](../../../models/cognitive-modes/README.md)。

## Closure 判斷輸入

`gate.software_delivery.security_audit_complete` 依下列規則判斷（gate 本體在 [`execution-flow.yaml`](../execution-flow.yaml)）：

| 條件 | 結果 |
| --- | --- |
| `audit_execution` 缺漏或 `status` ≠ `completed` | block（unknown ≠ safe） |
| `confirmed` 且 `severity` ∈ {critical, high}，`resolution.kind` ∈ {none, human_review} | block |
| `confirmed` 高嚴重度且 `resolution.kind: risk_acceptance` 欄位完整 | pass（記錄 acceptance） |
| `needs_validation` 且 `potential_impact` ∈ {critical, high} 且無 `human_review` 紀錄 | block |
| `needs_validation` 高潛在影響 + `human_review`（reviewer + 結論） | pass |
| 其他 | pass |

## Traceability

- **Caller slice**: <sd-contracts | sd-implementation>
- **Change Brief / Contract**: <link>
- **Coverage units touched**: <project coverage ids, or none>
- **Verifier report**: <link, or not run + reason>
