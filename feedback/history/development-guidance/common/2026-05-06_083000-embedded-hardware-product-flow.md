# Embedded hardware product flow belongs in app development guidance
# Extracted — See [`workflow/software-delivery/execution-flow.md`](../../../../workflow/software-delivery/execution-flow.md)

Status: candidate

## Lesson

Some products are not mobile/web/backend apps, but they still follow the same contract-first development discipline. Embedded firmware and hardware-backed products need product behavior, BDD, domain models, public interfaces, error handling, implementation slices, and tests, plus hardware-specific contracts for datasheets, protocols, board context, driver boundaries, target validation, and bring-up evidence.

## Rule

When a project involves firmware, sensors, boards, UART/I2C/SPI/BLE/CAN/GPIO, RTOS tasks, or hardware-in-loop validation:

1. Keep it in `app-development-guidance` as an app/product development flow unless the user needs a separate skill for hardware lab operations, flashing automation, schematic/PCB work, or toolchain-specific runbooks.
2. Add embedded guidance under `platforms/embedded/` and implementation details under `implementation/embedded/`.
3. Require datasheet/protocol contracts, hardware context contracts, driver/service/application boundaries, host fixtures, and hardware-in-loop evidence.
4. Do not rename the skill only because one hardware-backed project appears; prefer broadening language from app-only to app/product while preserving existing cross-skill links.

## Required Linked Updates

- `SKILL.md`: updated to include embedded, firmware, and hardware-product triggers.
- `README.md`: updated to describe app/product scope and embedded linked docs.
- `WORKFLOW.md`: updated with embedded owner layers, blockers, classification, and validation.
- `process/README.md`: updated with Embedded / Hardware Product Flow.
- `platforms/embedded/README.md`: added platform guidance.
- `implementation/embedded/README.md`: added contract-to-code implementation guidance.
- `CHECKLIST.md` and `checklists/embedded-firmware-review.md`: updated with review gates.
- `templates/initial-development-docs.md`: updated with Hardware / Firmware Contract fields.

## Validation

Use an existing embedded project as evidence only for reusable patterns. Do not copy project-specific pin choices, logs, board conclusions, or private hardware details into the reusable skill.

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。

#### One-line Summary

既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。

#### Evidence

既有工具輸出、觀察與專案證據已記於本 lesson 的原始段落；未新增或推論額外證據。

#### Generalized Lesson

將本條的具體情境視為候選通則；未在獨立情境重複驗證前，維持 candidate。

#### Agent Action

重用前先核對本條既有前提、限制與驗證方法；前提不符時重新取證。

#### Goal / Action / Validation

- Goal: 保留既有 lesson 的可驗證結論。
- Action: 依原始 Evidence 與 Trigger 重做相關檢查。
- Validation: 結果與原始結論一致才可重用。

#### Applies / Does Not Apply

- Applies: 本條既有 Trigger、前提與 Evidence 相符時。
- Does not apply: 跨專案、前提不明或缺少原始證據時；先重新驗證。

#### Promotion Target

- 尚未 promotion；保留為 candidate history，需有獨立 reuse evidence 才可提升。
