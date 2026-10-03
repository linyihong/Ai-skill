# Dogfood 1 — refresh-token family foundation（`<CONSUMER_PROJECT>`）

**Date**: 2026-10-03 · **Mode**: `standard` · **verdict_kind**: `expected` · **Roles**: auditor = main session; verifier = fresh subagent（不繼承稽核推理，只讀）

## Target

`<CONSUMER_PROJECT>` 已發佈的一個內部 auth 切片：refresh-token family 的持久化與 rotation（opaque 256-bit credential、只存 digest、generation + digest 的 compare-and-swap、絕對期限、撤銷）。這個切片故意不對外開放 endpoint；公開的 refresh / logout 屬於下一個切片。

選擇原因：真實授權邊界（token replay、撤權後續期、跨 site / Host）、已有獨立 V1–V4 紀錄可當可重現證據、下一個切片會改動同一組 control，可在 dogfood 2 觸發 coverage invalidation。

## 鏈路是否走通

| 步驟 | 結果 |
| --- | --- |
| invoke `security-audit`（fault_finding） | 產出 finding list，完全依 template 欄位 |
| `audit_execution` | `completed`；`execution_environment` 註明沒有隔離 DB、未執行目標程式碼，可重現證據引用既有紀錄 |
| Coverage ledger（專案端 YAML） | 初稿 6 units + 5 controls → verifier 後 10 units |
| Fresh verifier（V3 對抗） | 10 個 item，2 個 acceptance-violation、8 個 observation；另列 8 個遺漏 |
| 仲裁 | 全部 `fix`；orchestrator 自行重讀 DI module 確認 reachability 錯誤後才修 |
| Gate（expected） | pass：`audit_execution` completed、無 confirmed high、needs_validation 最高 medium |

Finding 終版：1 confirmed（low，replay 不撤銷 family；deferred 非 accepted）、4 needs_validation（revoke 無 user binding、SQL Server CAS 未執行、response loss 使 family 失效、issue check-then-insert）、1 candidate（改密碼不失效，超出現有 acceptance）。

## 這次 dogfood 證明的事

1. **獨立 verifier 有實際價值**：稽核者宣稱「service 沒有 DI 註冊」是錯的——專案用 assembly scanning 依命名慣例自動註冊。稽核者只 grep 型別名稱，看不到慣例註冊。Verifier 讀 DI module 才發現。這個錯誤會讓 `expires_when` 的觸發條件（「註冊呼叫者」）在一開始就已成立而沒人發現。
2. **Deferral ≠ risk acceptance**：稽核者把 brief 裡的「延後到下一切片」當成 risk acceptance，並自行填了 owner。Verifier 指出沒有任何決策者接受風險。Template 原本沒有區分這兩者。
3. **`audit_execution.not_covered` 有用但會被低估**：初稿 4 項，verifier 補出 4 項（每使用者 family 無上限、revoke vs rotate 競態、nested transaction、clock skew）。
4. **`examined_no_finding`** 讓 verifier 有東西可反駁：5 項中 2 項被指出過度宣稱（soft delete 不在 barrier test、digest-only 斷言偏弱）。
5. **Unit 失效條件要綁「第一個使用者」而不是「註冊」**：ledger 的 rotation-entry control 改寫為「第一個注入 / endpoint」，並附可機械執行的 detection（grep 介面在擁有者以外的引用）。

## 寫回 Ai-skill 的修正

- Template：deferral 與 risk acceptance 分開；confirmed finding 不帶 `potential_impact`；reachability 宣稱需涵蓋慣例 / 反射註冊並區分「已註冊」與「有使用者」；新增 `execution_environment` 與 `examined_no_finding`。
- Coverage ledger 契約：control 可帶 `detection`；失效觸發以第一個使用者為準。

## 尚未證明

- Coverage invalidation 實際觸發（需要下一個切片加入第一個使用者）→ dogfood 2。
- Gate 的 block 路徑在真實任務中出現（本次為 pass）。
- `enforced` verdict（Phase 5）。

## 去敏

本檔不含專案名稱、路徑、host、人名或 credential；專案端 finding list 與 ledger 留在 `<CONSUMER_PROJECT>` repo。
