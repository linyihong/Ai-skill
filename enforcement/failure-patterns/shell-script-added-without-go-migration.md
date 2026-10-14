# Non-Go Script Added Without Go Migration（新增非 Go 腳本入口而非 Go 實作）

Status: validated
Class: `process-gap` / `platform-governance-miss`

## Trigger

Agent 在新增 hook、工具腳本、CLI 輔助功能時，直接建立 `.sh` / `.py` / `.rb` / `.js` 等非 Go 入口，未先查閱平台治理政策（`plans/archived/2026-05-21-0834-cross-platform-go-script-runtime.md`、`scripts/README.md` §Go-first），導致跨平台問題被引入，且 Go CLI 遷移路徑未被規劃。

具體觸發信號：

- `git diff --cached --diff-filter=A --name-only` 顯示新增 forbidden extension（見 `kge.ForbiddenNonGoScriptExts`）
- Commit message 無說明「為何不用 Go 實作」或 Go migration plan
- 新增 `.claude/hooks/*.sh`、`scripts/*.sh`、`workflow/**/tools/*.py` 等平台相依／腳本入口

## Failure Mode

非 Go 腳本入口破壞 Go-first 與跨平台可重現性：

1. **跨平台 breakage**：`.sh` 在 Windows 無法直接執行；Python/Ruby 依賴本機 interpreter
2. **遷移債務累積**：每新增一個腳本入口，Go migration 工作量就增加
3. **治理政策被繞過**：platform governance 形同虛設；phase gate 腳本變成產品而非方法

## Required Agent Action

1. **新增任何 automation 入口前**，先讀 [`scripts/README.md`](../../scripts/README.md) §Go-first 與 [`plans/archived/2026-05-21-0834-cross-platform-go-script-runtime.md`](../../plans/archived/2026-05-21-0834-cross-platform-go-script-runtime.md) §platform-governance
2. 新邏輯必須實作在 `scripts/ai-skill-cli/internal/app/`（subcommand 或 hooks）
3. `ai-skill hooks install` 負責在 `.claude/settings.json` 寫入平台正確的 binary 路徑
4. Workflow 只提供 **方法／流程／注意事項**；不要預先提交 `run_phaseN_gate.*` 產品腳本——需要時依流程用 CLI methods 組裝
5. 若確有短暫過渡需要，在 commit message 加 `[skip-go-migration]` 並說明 Go migration plan

## Prevention Gate

Pre-commit hook `validateNoNewShellScripts`（`scripts/ai-skill-cli/internal/app/hooks.go` → `portable/kge.NoNewShellScriptsRule`）：

- 偵測 Added non-Go script entrypoints → block with `new_non_go_script_forbidden`
- Opt-out：`[skip-go-migration]` 獨立行於 commit message body
- 修改既有 forbidden 檔（pending migration）不觸發

## Validation

符合下列條件時，此 pattern 已被防止：

- staged Added paths 不含 `ForbiddenNonGoScriptExts`，或
- Commit message 含 `[skip-go-migration]` 且有 Go migration plan 說明
- 新邏輯在 `hooks.go` 或 subcommand 中實作

## Source

- 2026-05-27 session：修正 double-bootstrap 時直接改 `.claude/hooks/*.sh`
- 2026-10-14 session：image-processing 暫用 Python phase gates；擴充 gate 覆蓋 `.py` 等並改為 Go methods

## Related

- [`plans/archived/2026-05-21-0834-cross-platform-go-script-runtime.md`](../../plans/archived/2026-05-21-0834-cross-platform-go-script-runtime.md)
- [`runtime/cli-modification-policy.yaml`](../../runtime/cli-modification-policy.yaml) §`gate.cli.no_new_shell_scripts`
- [`scripts/README.md`](../../scripts/README.md) §Go-first
