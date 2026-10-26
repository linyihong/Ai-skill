# Unity 動畫／特效驅動追蹤摘要

| 欄位 | 值 |
| --- | --- |
| Atom ID | `unity-fx` |
| Source path | [Unity 驅動追蹤](../../analysis/apk/workflows/unity-fx-driver-trace.md) |
| Lifecycle | candidate |
| Summary | Unity FX Animator.Play 視覺缺口先追 feature FX Setup／Play 與 m_Payloads writer；再依證據追 PlayableDirector／PlayableGraph、ParticleSystem、SetActive、Transform、Native backtrace 與 hardware watchpoint。 |
| When to read | Unity 特效、動畫、payload、Animator.Play、PlayableDirector、fx driver、visual gap 或骨骼位移不同步；需要靜態 xref、AssetRipper binding 與動態事件／工作幀交叉驗證時。 |
| Do not use for | 不推定 Timeline／插件存在；不把欄位參照監看當成集合全內容監看；不把 hook 成功當完整還原。 |
| Validation signal | 同 fixture 事件與畫面對齊；native 位置按 module／build ID 映射；discovery bridge 與 RAG 能回到 canonical source。 |
| Last checked | 2026-10-26 |
