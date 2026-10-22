# Software Delivery Release Workflow

Release stage 檢查 behavior contracts、domain invariants、tests、performance evidence、known risks 與 rollback / recovery plan。

發布前必須執行 [`相依元件公告查核`](../dependency-advisory-review.md)，重新取得官方公告及掃描結果，並核對部署 runtime／映像版本；失敗或範圍內高風險未明時不得宣稱通過。
