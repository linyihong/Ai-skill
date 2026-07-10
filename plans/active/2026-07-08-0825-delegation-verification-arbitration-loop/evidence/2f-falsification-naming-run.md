# 2f — falsification run（預註冊 2026-07-09；F2 待使用者選名）

> **還原註記（2026-07-10）**：本 run 紀錄（含預註冊判準）原載於 kit §2f，於併發 plan 回寫（b6481e5）中遺失，自 commit 66f58ed 還原。**預註冊時序不受影響**：判準的原始 commit 錨點為 2026-07-09（f167123 / 081aa90），git 歷史可證。


> **預註冊時點先於任務存在**（anti-reconstruction：判準 commit 於 run 前，run 後不得修改本節判準去救假說；要改判準只能預註冊給再下一個 run，本 run 依原判準判定）。

- **目標**：預期失敗的工作性質——**preference-allowed**（brainstorming / creative writing / open-ended design / ideation）。精確判準是**工作性質**非 domain 名：同 domain 可同時含 justification-required 與 preference-allowed 工作（architecture 安全審查 vs 概念發想）。
- **任務等待條件**：對話或專案中**自然出現**的真實 ideation / creative 需求；出現時標記 2f、用 loop 跑；**brief 不得誘導 evidence 結構**（不預塞 evidence requirement，觀察它自己長不長）。
- **預註冊 falsification 判準（兩個獨立觀察）**：
  | 觀察 | 問題 | fail 樣態 |
  |---|---|---|
  | **F1** | acceptance 是否自然形成 evidence requirement？ | 完全沒有 grounded 判準自然出現 |
  | **F2** | closure 是否**真的依賴** independent evidence？ | 證據存在但最終以偏好/品味關閉（「我喜歡第 7 個」）——**證據裝飾性** |
- **判讀表（run 前定死）**：F1✓F2✓ = ERA 成立於 preference 域（假說極強）；F1✓F2✗ = 邊界訊號（evidence 裝飾性）；F1✗ = 邊界訊號。失敗 = 依「Justification Required vs Preference Allowed」畫出 ERA 適用邊界，與成功同等有效。
- **ERA 單一問句**（本 run 要回答的）：**該任務的 Decision 是否必須依賴 Independent Evidence？**
- **任務選定（2026-07-09）**：**repo / 系統改名候選**——真實性錨點：2026-05-26 landing-page plan 明文「`Ai-skill` 僅作為尚未改名的 repo slug」（plans/README.md L275 + archived plan L8，T0 早於本 run 與預註冊）。工作性質 = preference-allowed（最終選名 = 使用者品味裁決）。**誠實標記**：使用者要求加速、由 orchestrator 主動排程（比照 2e 先例：真實任務可刻意排程，不可捏造）；brief 以自然委託語氣撰寫，**無 acceptance-evidence 標準、無 verification 段、無 loop 模板結構**——F1/F2 依預註冊判準觀察。
- **結果（期中，2026-07-09；F2 待使用者選名）**：
  - **Producer**（fresh agent、自然語氣 brief、零結構）：交付 12 候選（`03-repo-naming-candidates.md`，commit `0c35d92`）。**F1 觀察**：brief 未要求任何證據標準，producer **自發**附上篩選準則、逐候選撞名風險註記、並誠實標明知識截點限制建議定名前實查——evidence requirement **自然形成**（但為斷言式 grounding，非執行式查證）。
  - **Reviewer**（fresh agent、自然語氣 brief、零模板）：**自發產出** 結構性 findings（全清單被動隱喻偏誤、缺「execution」維度候選）+ 逐案 grounded 批評（含 rationale 事實錯誤查核：canonry 非自造詞、noema 現象學誤用；撞名評級擴格反證 ×5）+ 明確自我節制「最終選名不是我做」。**證據責任結構在無模板下自然重現**：找碴附依據、不做決定。
  - **期中新假說（比二元 F2 更細）**：preference-allowed 工作裡，evidence 自然出現在 **filter 層**（review 實質改變決策空間：淘汰 grimoire、揭露 vademecum 文化坑、修正 cairn 評級——證據非裝飾性），而 **selector 層**（存活候選中挑哪個）留給偏好。若 F2 最終量測證實此形態 → ERA 邊界不是「有無證據責任」，是「**證據責任止於 filter、closure 屬 preference**」——比全有/全無更精確的邊界。
  - **F2 量測點**：使用者選名的裁決依據——待選。**補充量測尺度（2026-07-09 第六輪 review 提出，時點先於選名，記為 supplementary；原始二元 F1/F2 判準仍為本 run 正式判定基礎）**：三模式分類 Evidence-determined / Evidence-constrained / Preference-determined + **反事實鑑別問題**（「若 review 不存在，你會選同一個嗎？」——排除後最佳剩餘 vs 本來就會選，同一選擇對 ERA 意義不同）。

