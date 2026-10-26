# Evidence Model

**Status**: `candidate-intelligence`

## 定義

Evidence model 定義不同證據類型的可信度、覆蓋範圍與適用邊界。它不是測試清單，而是幫 agent 判斷「這個證據到底證明了哪一段」。

## Evidence Characteristics

```yaml
evidence_types:
  api_response:
    confidence: low
    scope: local
    proves: controller_or_request_path_accepted

  adapter_success:
    confidence: low
    scope: boundary_local
    proves: adapter_call_returned_success

  log_record:
    confidence: medium
    scope: observed_runtime
    proves: code_path_or_operation_was_recorded

  database_state:
    confidence: high
    scope: persisted
    proves: durable_state_changed

  user_observable_state:
    confidence: high
    scope: user_visible
    proves: product_surface_reflects_state

  independent_external_confirmation:
    confidence: very_high
    scope: independent
    proves: external_system_or_user_observable_effect_confirmed
```

高可信度不代表覆蓋完整。`api_response` 可以是真實可信的，但通常只證明 request path 被接受，不證明 DB、queue、email、UI 或外部系統都成功。

## Evidence Preferences

| Domain | Avoid as final proof | Prefer |
| --- | --- | --- |
| Authentication | login API success only | real cookie/session, protected resource access, identity-specific readback |
| Email | SMTP adapter success only | inbox received, provider record, product-visible delivery state |
| Payment | API `200` / order accepted only | gateway record, webhook record, DB settlement state, entitlement state |
| Queue / event | publish call success only | topic record, consumer processing evidence, resulting business state |
| Storage | upload API success only | object exists, reachable URL/readback, metadata/ACL match |
| SSR / UI | API success only | SSR response text, CTA/state, persisted state consistency |

## 使用方式

1. 先列出 claim：「系統做了什麼」。
2. 將現有證據映射到 `evidence_types`。
3. 標出每個證據的 `scope`。
4. 若 claim 跨越多個 boundary，使用 [`evidence-chain-validation.md`](evidence-chain-validation.md) 補證據鏈。
5. 若只剩 proxy signal，視為 [`../../../anti-patterns/validation-proxy-trap.md`](../../anti-patterns/validation-proxy-trap.md)。

## 相關知識

- [`state-visibility-gap.md`](state-visibility-gap.md)
- [`evidence-depth.md`](evidence-depth.md)
- [`evidence-chain-validation.md`](evidence-chain-validation.md)

## 宣稱、角色與部署讀回的候選補充

- 保護性語句須對應操作與角色範圍；多個局部證據不能支持更廣的「無法／防止」保證。只有偵測能力時，不能當成已阻擋。
- 操作者陳述、tag 與健康端點各證明不同事實。部署現況採既有權威來源，按元件保留來源與 unknown；不因單一端點成功推定整個環境一致。
- 反例：已有完整 CD 現況平台時，另建帳本不會自然增加證據可信度；應讀原平台。

### 來源回饋

- [2026-10-20_150000-claim-vs-proof-audit-before-a-protection-wording-switch.md](../../../../feedback/history/development-guidance/common/2026-10-20_150000-claim-vs-proof-audit-before-a-protection-wording-switch.md)
- [2026-10-22_093000-deployment-state-is-readback-not-narrative.md](../../../../feedback/history/development-guidance/common/2026-10-22_093000-deployment-state-is-readback-not-narrative.md)
