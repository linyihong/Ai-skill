---
id: 2026-06-22-1009-phase3-dogfood-glossary-spike
plan_kind: spike
status: in-progress
owner: linyihong
created: 2026-07-06
parent: 2026-06-22-1009-plans-system-portability-and-delivery-integration
required_for_completion: false
sub_plan_reason: >
  Phase 3 dogfood 載體（可獨立 archive 的 spike）：承載一份真實 delegation
  brief（註冊 glossary 詞條 plan_profile），用來驗證 (a) Phase 2 delegation
  validator 接受真實完整 brief、(b) brief 的獨立性（Brief Independence Score）。
  非 03 completion 條件（required_for_completion: false），dogfood 完可 archive/移除。
delegation:
  enabled: true
  brief:
    goal: >
      在 Ai-skill glossary（knowledge/glossary/ai-skill.md）新增 `plan_profile`
      詞條，關閉 runtime audit 對它的 glossary-coverage 缺口。
    acceptance:
      - knowledge/glossary/ai-skill.md 出現 `## plan_profile` 區塊，含標準 yaml
        block（term / status: candidate / owner-layer / meaning / affects /
        anti-meaning / related-terms / introduced-by），置於正確英數字母排序位置。
      - meaning 需表達：plan_profile = plan-validation engine 的 **portable
        capability 邊界**——「任何採用 canonical plan tree 的 repo 都應成立」的
        plan-structure 規則集合（unique id / parent resolution / archive
        ordering / required-sub completion / schema compatibility），與
        Ai-skill-only governance overlay（例如 delegation）**刻意區分**。
      - anti-meaning 需表達：plan_profile **不是**「validator 類型」直覺分類，
        也不含 Ai-skill-only workflow 特性；它由 contract/dependency/execution-
        context 分類**推導**得出。
      - （brief v2，agent dogfood 回饋補上）`introduced-by` = 01 的
        `01-external-repo-plan-system-shared-binary.md`（plan_profile 概念起源，
        非本 spike/03）；`owner-layer` = `validation-governance`。原 brief 未 pin
        這兩欄 → agent 只能推斷，`introduced-by` 推成 03（錯）。
    verification:
      - 執行 `scripts/ai-skill-cli/bin/ai-skill-darwin-arm64 runtime audit`，
        確認 `glossary candidate plan_profile ... not in glossary_terms` 警告
        消失。
      - 執行 `runtime refresh` 後 `runtime validate` 乾淨（glossary 改動需 refresh
        index）。
    context:
      required:
        - knowledge/glossary/ai-skill.md
  execution:
    modes:
      - human
      - agent
    constraints:
      - 只改 knowledge/glossary/ai-skill.md（+ glossary 編輯造成的 runtime index refresh）；不動其他檔。
---

# Phase 3 Dogfood — Glossary `plan_profile` Registration（spike）

**Status**: `in-progress`
**Owner**: linyihong
**Parent**: [`_plan.md`](_plan.md)

## 用途
sub-plan 03 Phase 3 的 dogfood 載體。本 spike 的 `delegation.brief` 就是被委派的
任務本體：註冊 glossary 詞條 `plan_profile`（一個真實的 runtime-audit coverage
缺口，×46 使用未註冊）。

- **驗 (a)**：Phase 2 delegation validator 對「真實完整 brief」pass（enabled + 4
  必填齊）。
- **驗 (b)**：brief 是否自足——human / agent 僅憑 `brief`（+ `context.required`）
  能否完成，記 **Brief Independence Score**（見 03 §Phase 3）。

## Dogfood 結果（回填）
- **agent 路徑：★★★★☆（2026-07-06）**。乾淨 general-purpose agent、worktree 隔離、僅餵 brief（無對話上下文）。
  - **只讀 `context.required`（glossary 檔）本身，未讀 main plan / 未 read whole repo** → 依 Brief Independence Score 為 ★★★★☆（brief + context.required）。
  - 內容正確:meaning/anti-meaning 直接對上 acceptance;yaml 欄位形狀、`excludes: delegation` 互倒數連結由同檔既有 `delegation` 詞條推得。
  - **缺漏（→ brief v2 已修）**:brief 未 pin `introduced-by` / `owner-layer` → agent 推斷,`introduced-by` 推成 03（應 01）。`owner-layer` 推成 `validation-governance`（其實比 delegation 的 `plan-governance` 更準,採納）。
  - **正向信號（capability/workflow 分層成立）**:agent 回報「brief 未提 worktree,首次 Edit 對 shared-checkout 失敗才改 worktree copy」——worktree 是 **Layer 3 tool/execution** 細節,**本就該不在 tool-neutral brief 裡**（歸 `execution.constraints` + `ai-tools/`）。agent 靠 operational discovery 處理掉,證明 brief 保持 tool-neutral 是對的,不是缺陷。
  - **採納**:agent 產出（修正 `introduced-by`→01 後）已 land 進 canonical glossary,關閉真實 audit 缺口（`glossary candidate plan_profile` 警告消失、`runtime validate` 乾淨）→ dogfood 淨正值,非僅觀察。
- **human 路徑：<待你回填 — Score + 觀察>**（brief 已在對話中交付你）

## 收尾
dogfood 完成後：brief 缺漏回饋修正 03 schema/範例；本 spike 可 archive（不影響
main tree completion，`required_for_completion: false`）。
