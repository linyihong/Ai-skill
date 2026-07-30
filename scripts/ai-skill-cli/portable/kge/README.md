# Knowledge Governance Engine（portable copy unit）

**這整個目錄就是可複製單元。**

不引用 Ai-skill 全庫、不依賴 `runtime.db`、不呼叫 git／hooks。  
外部專案若只要「組裝好的驗證引擎 + 規則」，把本目錄拷進自己的 Go module 即可。

Ai-skill 本體只負責 **Adapter**（組 `Context`、接 commit-msg／CI／CLI），不把流程鎖死在這包裡。

## 邊界

| 屬於本目錄（可 copy） | 不屬於本目錄（留在宿主專案） |
| --- | --- |
| `Context` / `Finding` / `Rule` / `Capability` | git hook、CI YAML |
| `Engine.Run` | 組裝 staged paths／diff／commit msg |
| 各 `rule_*.go`（Validate 純函式） | Ai-skill `runtime.db`、bootstrap、enforcement prose |
| 單測（stdlib only） | 專案特有 policy 文件路徑以外的 wiring |

**硬規則**：Rule **不得** `exec.Command("git"…)`，也不得 import 宿主的 hook 包。缺能力時回 `capability_missing`，由 Adapter 補 Context。

## 複製步驟（外部專案）

```text
1. 複製本目錄 → <THEIR_MODULE>/kge/   （或任意路徑，保持 package kge）
2. 在他們的 adapter 裡：
     eng := kge.NewEngine(kge.CognitiveCostRule{}, kge.CLIDocSyncRule{})
     findings := eng.Run(ctx)
3. Adapter 負責填 Context + Provided capabilities
4. 依 findings 決定 exit code（本包不強制 block 政策）
```

可刪掉不需要的 `rule_*.go`，或新增自己的 `rule_xxx.go`（實作 `kge.Rule`）。

## CLI（Ai-skill adapter）

```bash
ai-skill kge check [--root PATH]              # validation + advisory summary (D9)
ai-skill kge validate [--root PATH] [--advisory]  # validation; --advisory = full list
```

## 與 Ai-skill 的關係

| 角色 | 位置 |
| --- | --- |
| Portable engine | `scripts/ai-skill-cli/portable/kge/`（本目錄） |
| Ai-skill adapter | `scripts/ai-skill-cli/internal/app/kge_adapter.go` + `kge_cmd.go` |

Plan：`plans/active/2026-07-30-0950-knowledge-governance-runtime/`

## 驗證

```bash
cd scripts/ai-skill-cli && go test ./portable/kge/ -count=1
```

## 非目標

- 不是「再包一層 Ai-skill CLI 才能用」
- 不是把 enforcement Markdown 一併複製
- 不是單一巨大 `.go` 檔（刻意多檔：contracts / engine / rules，方便抽換規則）
