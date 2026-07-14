> 遵守 [共用規則索引](../../../../enforcement/README.md)、[sanitization](../../../../enforcement/sanitization.md)、[reusable-guidance-boundary](../../../../enforcement/reusable-guidance-boundary.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson。

### 2026-07-14 - APK analysis window close-out needs Cognitive + Feedback

Status: candidate

#### One-line Summary

每個 APK 分析視窗（static／MITM／Frida／handoff assessment）結束時，必須同時輸出 **Cognitive Mode 報告**、**Feedback / Learning Report**，並在有可重用 lesson 時寫回 `<AI_SKILL_REPO>/feedback/history/`——stop-hook 會機械檢查，不可只留專案筆記。

#### Human Explanation

分析回合容易「做完動態就關掉」，漏掉：

1. Compact／full `Cognitive:` block（per-turn obligation）。
2. `FeedbackDecision` 區塊（有 lesson 時 `NEEDED` + `Target` + 實際 writeback）。
3. 去敏 lesson 檔與 history index 更新。

使用者已明確要求：分析完要回饋 Ai-skill；忘記 Cognitive／Feedback 會被機械檢查擋下。專案 overlay 應把「每窗 check」寫死，避免再開下一窗才補。

Close-out 仍遵守 Domain 邊界：Discovery Evidence 留 `<PROJECT_ROOT>`；reusable method 才進 feedback history。**不要**把目標 App 的真實 path／host 寫進 lesson。

#### Trigger

- User says analysis window is done, or agent is about to switch tools (MITM→Frida→SDK).
- Stop-hook missing Cognitive / Feedback.
- User reminds to write back to Ai-skill.

#### Evidence

- Obligation contract: `<AI_SKILL_REPO>/runtime/core-bootstrap.yaml` per-turn obligations.
- Project overlay: `<PROJECT_ROOT>/.ai-skill/project/rules/ai-skill-session-feedback-writeback.md`（強化後）.

#### Generalized Lesson

Treat each analysis technique window as a mini close-loop:

```text
Window evidence → project docs
Reusable method? → sanitized lesson + FeedbackDecision:NEEDED
Always → Cognitive: line (+ full table if non-default)
```

#### Agent Action

Before final response of an analysis turn:

1. Emit Cognitive report.
2. Decide FeedbackDecision；if NEEDED, write lesson then `Writeback: COMPLETED`.
3. Update domain README index row.
4. Keep Capability Assessment separate（Discovery vs Delivery）—do not auto-handoff to software-delivery.

#### Applies When

Any authorized APK analysis session under Ai-skill bootstrap.

#### Does Not Apply When

Pure Q&A with no new method and no runtime work（FeedbackDecision may be NONE）.

#### Promotion Target

- Project overlay：`<PROJECT_ROOT>/.ai-skill/project/rules/ai-skill-session-feedback-writeback.md`（linked）
- Linked-update reminder：lesson naming `analysis/` / `intelligence/` targets must be executed in the same writeback transaction（見 [`enforcement/linked-updates.md`](../../../../enforcement/linked-updates.md)）
