# D9 review — Continue maintenance, close delivery plan

Run ID: ui-pattern-d9-closeout-2026-09-09

## Preflight

| Item | Decision |
| --- | --- |
| Trigger | 使用者要求收尾其他計畫；D9 T4（五個完整 Core entries）已成立 |
| Sources | 主計畫 D9／完成條件；entry-schema、五個 entries、pattern-index；RC1／RC2 closure；stakeholder research-line closure；maintenance-governance；consumer maintenance overlay |
| Scope | D9 裁決、維護交接、歸檔與引用修復 |
| Conflict | research line 已 closed，但 plan D9 仍 deferred；不能把經過時間當 maintenance 實測 |
| Decision | **Continue**：維持 doc-only knowledge 與 governed maintenance；本 delivery plan completed |
| Non-goals | L2 runtime wiring、ADR promotion、glossary promotion、新研究 cycle、產品 UI 驗收 |
| Validation | 本輪 Core schema 結構核對、檔案／引用讀回；歷史研究結果只引用原紀錄 |

## D9 trigger readback

| Trigger | 本輪判定 |
| --- | --- |
| T1：三個月 | 未到；2026-07-14 起算至 2026-10-14。不得用 09-09 代替原起算點 |
| T2：第一次跨專案 Core 採用 | 有 consumer interaction／maintenance 使用紀錄，但本輪不把它等同完整 Core 採用證明 |
| T3：至少兩個 consumers | 本輪未證明；同一 consumer 的兩個 cases 不算兩個 projects |
| T4：至少五個完整 Core entries | **成立**：scrim、modal_dialog、bottom_sheet、drawer、toast，5/5 結構核對通過 |
| T5：abandoned | 未見書面放棄；既有 stakeholder 指令為 governed maintenance |

T4 已足以啟動本次 review，不需要補造 T2/T3 或等待 T1。

## Evidence and limits

- [Phase 2 summary](phase2-summary.md)：歷史 selection evidence；非本輪重新盲測。
- [RC1](research-cycle-1.md) 與 [RC2](research-cycle-2.md)：既有 research closure；本輪沒有宣稱再次驗證全套研究假說。
- [Stakeholder closure](stakeholder-research-line-closure-2026-07-15.md)：已要求停止研究擴張、進入維護。
- Canonical 五個 entries 保留 Core required keys、when/not_when、neighbors、intent examples 與合法 enum；Extended 不作完整性必要條件。
- Consumer 的 `.ai-skill/project/rules/stable-maintenance-dogfood.md` 仍保留 layer-first → mapping → archive 的使用入口。
- 本輪在提供的 consumer evidence 目錄未找到 `stable-maintenance-*.md` run；**不宣稱**三至四週觀察已完成、不推算 mapping 成功率或零 boundary breaks。
- 本輪只讀 consumer；consumer overlay 中舊 Ai-skill active pointer 留作 consumer maintainer 的導航更新事項，實際 canonical 入口是 maintenance-governance。

## Disposition

**Continue** 延續的是知識維護，不是未完成的研究實作。
原 plan 完成條件要求「至少一次 D9 書面裁決」，未要求必須 Promote。
RC1/RC2 與維護交接已完成，因此本計畫歸檔；維護 owner 為 software-delivery workflow maintainer，consumer incident owner 負責本地 evidence。

| Item | 處置與重啟條件 |
| --- | --- |
| Optional L2 | deferred；只有後續 D9 明確 Promote 且有 named consumer、load_when、scenario 時另開有界 wiring 工作 |
| Project alias／recipe partial、outer-inner 雙鏈 | deferred 至 consumer 實際實作；consumer UI maintainer owns，不是本 plan gate |
| 三角色入口 | deferred／非本 plan scope；沿用既有 delegation workflow，不新增 loop |
| Glossary candidates | 不 promotion；維持候選清單與現有 workflow 用詞，不藉結案升格 |
| ADR | 不 promotion；workflow/templates 是足夠且較輕的 target，跨專案 Core/expansion 證據未在本輪補足 |
| Method maturity | 保持 Replicated once；同一 UI family 不能當作跨 family replication |

下一次 D9 review：原 T1 滿三個月時（2026-10-14），或此前新確認的 Core 跨專案採用／新增第二 consumer／書面放棄，任一先到即人工 review。
本次已審的五個 entries 不因再次計數而重複觸發 review。真實 Boundary Break 依 maintenance-governance 的 Vocabulary Exit／Readiness 規則處理。
此處記錄維護責任與條件，沒有建立排程或自動啟動研究。

## Closure checks

主計畫與 evidence 整體移至 archived；inbound links、workflow 維護入口、plans index 同步。
本輪新增證據與結論不回填為 2026-07-15 的原始量測。
Repository validation 與 git push 結果以本輪命令輸出與最終 refs readback 為準。
