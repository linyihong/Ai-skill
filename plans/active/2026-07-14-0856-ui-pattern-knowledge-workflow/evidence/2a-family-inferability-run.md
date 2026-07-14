# 2a — Phase 2 Family Inferability Run

**Plan**: [`../`](../../2026-07-14-0856-ui-pattern-knowledge-workflow.md)  
**Scenarios**: [`selection-scenarios.yaml`](selection-scenarios.yaml)  
**Date**: 2026-07-14  
**Method**:
1. **rule-trace** against `entries/*.yaml` selection_rules + near_neighbors  
2. **LLM blind run**（Task agent；只讀 entries + pattern-index；不讀本檔 / expected fields）— 2026-07-14

## Gate

成功 = **10/10 PASS**。本檔為第一輪閉環（rule-trace + blind）。

## Run matrix（rule-trace）

| ID | Prompt（短） | Expected | Trace（entry keys hit） | Result |
| --- | --- | --- | --- | --- |
| S1 | 刪除確認 + 背景壓暗 | `scrim`（composition；主面 dialog） | scrim.`when`: modal_or_temporary_overlay_present, need_visual_separation；neighbor modal_dialog；`not_when` 排除 sole surface | **PASS** |
| S2 | 已儲存短暫提示 | `toast`（排除 scrim） | toast.`when`: non_blocking_feedback；scrim.`not_when`: feedback_only_toast_or_snackbar | **PASS** |
| S3 | 確定刪除嗎？ | `modal_dialog` | dialog.`when`: destructive_or_irreversible, confirm_or_cancel；sheet.`not_when`: destructive_confirm_primary；neighbor rule 明寫 Sheet≠Dialog | **PASS** |
| S4 | 條款短閘門確認 | `modal_dialog` | dialog.`when`: needs_blocking_attention, short_focused_content；toast.`not_when`: requires_user_choice | **PASS** |
| S5 | 選付款方式 | `bottom_sheet` | sheet.`when`: payment_share_or_picker_list, multiple_actions_or_options；dialog.`not_when`: choose_payment_or_share_targets | **PASS** |
| S6 | 分享到多個目標 | `bottom_sheet` | sheet.`intent`: share_or_export；dialog 排除 option list；toast 排除 multi-action | **PASS** |
| S7 | 左側常駐導航 | `drawer` | drawer.`when`: persistent_or_semi_persistent_nav, navigation_destinations；sheet.`not_when`: persistent_navigation | **PASS** |
| S8 | 側欄篩選與列表並存 | `drawer` | drawer.`when`: filter_or_settings_panel_beside_content, side_panel_with_main_content；declare persistent/temporary in recipe | **PASS** |
| S9 | 已儲存成功 | `toast` | toast.`family`: **feedback**；overlay neighbors 指向 dialog/sheet 不當選 | **PASS** |
| S10 | 背景完成輕提示 | `toast` | toast.`when`: success_or_soft_status_ack, no_required_decision；scrim 排除 | **PASS** |

**Score (rule-trace)**: 10/10 PASS

## Blind LLM run

| ID | Blind primary | Match expected |
| --- | --- | --- |
| S1 | scrim | ✓ |
| S2 | toast | ✓ |
| S3 | modal_dialog | ✓ |
| S4 | modal_dialog | ✓ |
| S5 | bottom_sheet | ✓ |
| S6 | bottom_sheet | ✓ |
| S7 | drawer | ✓ |
| S8 | drawer | ✓ |
| S9 | toast | ✓ |
| S10 | toast | ✓ |

**Score (blind)**: 10/10 PASS

## Family boundary note

| Pattern | Family | 在本輪的角色 |
| --- | --- | --- |
| scrim / modal_dialog / bottom_sheet / drawer | `overlay` | Overlay Decision + compositional scrim |
| toast | `feedback` | **不**進 Overlay Decision；S2/S9/S10 驗 Family 是否有用 |

若把 toast 誤標成 `overlay`，S9/S10 會退回「哪個 overlay」競爭，Family 邊界失效。本輪刻意把 toast 改為 `feedback`。

## 剩餘（不擋本 gate）

- `<PROJECT_ROOT>` project alias / outer↔inner 雙鏈（plan 標次要）

## Verdict

Phase 2 **主成功條件**（5 entries + 10 Selection Scenarios）在 **rule-trace + blind LLM** 兩層皆 **10/10**。  
下一動作：stakeholder 確認關閉 Phase 2，或開 Phase 3 composition。
