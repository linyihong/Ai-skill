# Software Delivery Release Workflow

Release stage 檢查 behavior contracts、domain invariants、tests、performance evidence、known risks 與 rollback / recovery plan。

發布前必須執行 [`相依元件公告查核`](../dependency-advisory-review.md)，重新取得官方公告及掃描結果，並核對部署 runtime／映像版本；失敗或範圍內高風險未明時不得宣稱通過。

候選補充：保護性措辭逐句對應控制與證據，見 [宣稱驗證候選](../../../feedback/history/development-guidance/common/2026-10-20_150000-claim-vs-proof-audit-before-a-protection-wording-switch.md)；部署現況區分讀回、陳述與 unknown，見 [部署讀回候選](../../../feedback/history/development-guidance/common/2026-10-22_093000-deployment-state-is-readback-not-narrative.md)。具體帳本、tag、通知及閘設計由專案政策決定，並非通用必備實作。
