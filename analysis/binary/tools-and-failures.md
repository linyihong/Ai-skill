# Binary 工具與失敗判讀

| 工具族 | 用途 | 備註 |
| --- | --- | --- |
| Ghidra／IDA／Hopper／radare2 | 反編譯、xref | 選一即可；可選 MCP 橋接 |
| pwntools（caller-supplied） | ELF layout／部分 crash | Linux x64 常見 |
| GDB／pwndbg | core-only context | 非 live 除非另授權 |
| 可選 REA MCP／CLI | 同上能力的 agent 介面 | 見 `ai-tools/rea-mcp.md` |

| 失敗 | 處理 |
| --- | --- |
| 引擎未安裝 | engine_unavailable；改手動／CLI |
| 大檔啟動慢 | 拉長 deadline；先 overview |
| Windows／arch 不支援 | 記 host boundary |

← [binary/](README.md)
