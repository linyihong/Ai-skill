# 檢索與採用計畫證據

## 引用規則

記錄全文放各 Run 檔，主計畫只引用檔案路徑；不使用行號作證據定位。新增 Run 必須同步更新本索引，原始專案資料須去敏；觀察不等於採用或完成證據。

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
| --- | --- | --- | --- |
| phase0-2026-10-23 | [phase0-inventory.md](phase0-inventory.md) | verified inventory | 來源覆蓋、metadata、consumer、Bridge 決策與 baseline 入口；非檢索品質驗收 |
| phase1-protocol | [phase1-protocol.md](phase1-protocol.md) | frozen protocol | 18 歷史重播 + 2 人工控制、評分、切分與成本門檻 |
| phase1-baseline | [phase1-baseline.md](phase1-baseline.md) | measured baseline | Diagnostic 1/14、holdout 0/4、人工負面 0/2；採用層 not evaluated |
| phase1-audit | [phase1-audit.md](phase1-audit.md) | gate pending verification | 公開固定版本、雙指標分母、粗原因分類、分離對照與預留資料曝光紀錄 |
| prephase2-control | [phase2-execution-contract.md](phase2-execution-contract.md) | controls verified / experiments pending | Binary pin、index manifest、四組配對與候選審閱；B0 重播一致，未執行 Q/E/QE |
| phase2-v1-ablation | [phase2-v1-ablation.md](phase2-v1-ablation.md) | measured / candidate rejected | 八份 replay、兩次 matrix 通過；QE 14/14 但 confuser 6/6、preview 超額；queue 未結清，production 未採用 |
| phase2-v2-protocol | [phase2-v2-protocol.md](phase2-v2-protocol.md) | frozen before replay | 代表案例欄位命中、phrase 優先、固定權重、top1 / 2000 runes；只評最終輸出 |
| phase2-v2-ablation | [phase2-v2-ablation.md](phase2-v2-ablation.md) | measured / candidate rejected | 精確優先、加權、top1 與預算後重評；QE 12/14、overlap 3/5 / confuser 3/6；輸出 ≤841 runes，但品質未通過 |
| phase2-v3-protocol | [phase2-v3-protocol.md](phase2-v3-protocol.md) | frozen offline design | 固定 broad-64 pool；排序與兩種 packer 分開比較，不預設 top1 |
| phase2-v3-semantic-contract | [phase2-v3-semantic-contract.md](phase2-v3-semantic-contract.md) | existing-agent enabled / API optional | 使用者已選現有 runner；不可得 usage 明列，美元價格／tokenizer gates 僅適用額外 API |
| phase2-v3-preflight | [phase2-v3-preflight.md](phase2-v3-preflight.md) | offline measured / semantic pending | 26 固定 pool、156 分離比較；多筆 full 可保留命中但 confuser 6/6；六份 label-free requests 與成本契約待執行 |
| phase2-v3-agent-runner | [phase2-v3-agent-runner.md](phase2-v3-agent-runner.md) | frozen dispatch protocol | 六個 fork-none subagents、單題 request 隔離；保留原始輸出與不可得 usage，不另購 API |
| phase2-v3-agent-pilot | [phase2-v3-agent-pilot.md](phase2-v3-agent-pilot.md) | six attempts measured / gate failed | Fresh subagents：5/5 正向、2/4 confuser；原始輸出與 unavailable usage 保留，時間監測未通過；Phase 2 未完成 |
| phase2-v3-disposition-protocol | [phase2-v3-disposition-protocol.md](phase2-v3-disposition-protocol.md) | frozen before D1 scoring | 只加 disposition admission；retrieved／preview／verified／adopted 分層，原成績不改 |
| phase2-v3-disposition | [phase2-v3-disposition.md](phase2-v3-disposition.md) | selection slice passed / Phase 2 incomplete | 同六題 direct-only preview 5/5、0/4、≤1922 runes；retrieved 仍 4/4；9 份讀回身份一致，適用／採用未驗收 |
| phase2-readiness-contract | [phase2-readiness-contract.md](phase2-readiness-contract.md) | frozen offline preview contract | unknown coverage 必須可見、P0 預設、分開 envelope 成本；列剩餘 Phase 2 gates 與責任 |
| phase2-readiness | [phase2-readiness.md](phase2-readiness.md) | coverage measured / Phase 2 incomplete | 全 26 題 770 source occurrences，738 unknown；P0 array 最大 1922，但可見 envelope R08 2208，整體成本未通過 |
