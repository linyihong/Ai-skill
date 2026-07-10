# 2g — ExternalRepoA server_doc test placement + delegation overlay（2026-07-09，證據 only）

> **還原註記（2026-07-10）**：原載於 kit §2g（commit 66f58ed），於併發回寫中遺失，還原至此。


> **專案證據邊界**：class 名、live 路径留于 consumer workspace；Ai-skill 只保留 generalized metrics。

- **任務**：`server_doc` 内共置 `*.test.*` 不得新增/改内容，仅 delete；验证一律外层 `tests/`；policy 落 `.ai-skill/project/rules/`，**不**写 `server_doc/docs/`。
- **Delegation 落地**：`plan-delegation-execution-loop.md` + `delegation-verification-backfill.md` template；ERA 分工（Orchestrator overlay / Executor 改 guard 模板 / Verifier 跑 BDD live probe）。
- **机械 gate**：`gate.short_drama.server_doc_test_placement` — 外层 Cursor `check-plan-phase-before-commit.py` + template `.ai-skill/project/templates/sibling-repo-githooks/` → sibling 本地 `.githooks/`（`install-sibling-repo-githooks.sh`；`core.hooksPath`）。**Concealment 仅 sibling app repo 远程**（README 一行「提交规范」）；**Ai-skill / 外层 workspace 可完整写 install 路径**。
- **BDD**：7/7 pass — 增/改拦截、纯删放行；三角：overlay doc ↔ template pattern ↔ self-test + staged probe。
- **相对 ExternalRepoC 2d**：同 consumer overlay + backfill 模式；差异 = filename+diff 渐进迁移（非 Java `src/test/` 整目录 block）。
- **不视为** Phase 3 closure；support doc-only + consumer mechanical gate 证据累積。

