# Dogfood 證據索引（Delegation Verification & Arbitration Loop）

本目錄存放 **dogfood 量測與契約回饋** 的全文；[`01-dogfood-prompt-kit.md`](../01-dogfood-prompt-kit.md) 保留 **傳輸模板 A/B/C** 與精簡指標。

> **Canonical 規則**：[`governance/lifecycle/plan-evidence.md`](../../../governance/lifecycle/plan-evidence.md)（commit-msg `validatePlanEvidenceConvention` 機械強制）

## 引用規則（避免行號漂移）

| 做法 | 說明 |
|---|---|
| **用檔案路徑** | `evidence/<slug>.md` 或相對連結，**不要**寫「kit L449」類行號 |
| **用標題錨點** | 需要定位時寫 `evidence/foo.md` 內 `## 量測欄` 或表格 `#` 欄，不用絕對行號 |
| **專案細節** | inner commit、class 名、live 環境 → 留 consumer `<PROJECT_ROOT>` plan §執行紀錄；本目錄只留 generalized metrics（[`enforcement/sanitization.md`](../../../../enforcement/sanitization.md)） |
| **新 run** | 一律新增 `evidence/<run-id>-<slug>.md` + 更新本索引；kit 只留一行指標摘要 + 連結 |

## Run 索引

| Run ID | 檔案 | 狀態 | 摘要 |
|---|---|---|---|
| **2a** | [`2a-software-delivery-review-invoke.md`](2a-software-delivery-review-invoke.md) | 完成 | software-delivery review invoke dogfood |
| **2a-external** | [`2a-external-sync-adapter-step6.md`](2a-external-sync-adapter-step6.md) | 完成 | 外部 sync adapter step6 |
| **2b** | [`2b-plans-sop-expansion.md`](2b-plans-sop-expansion.md) | 完成 | plans SOP 擴充 + brief v2 回饋 |
| **2c** | [`2c-tiered-archive-platform.md`](2c-tiered-archive-platform.md) | 證據 only | tiered archive platform dogfood |
| **2d** | [`2d-outbound-sync-phase3.md`](2d-outbound-sync-phase3.md) | 證據 only | outbound sync Phase 3（4 slices） |
| **2d′** | [`2d-prime-externalrepoc-module-alignment.md`](2d-prime-externalrepoc-module-alignment.md) | 證據 only | ExternalRepoC 9j2 模組 01/02 對齊 follow-on；integration gate、remote_absent_delete、live teardown、release-time gate |
| **2f** | [`2f-falsification-naming-run.md`](2f-falsification-naming-run.md) | 預註冊 | falsification naming run |
| **2g** | [`2g-externalrepoa-server-doc-placement.md`](2g-externalrepoa-server-doc-placement.md) | 證據 only | ExternalRepoA server_doc test placement + overlay |
| **2h** | [`2h-externalrepoc-common-url-verification-gaps.md`](2h-externalrepoc-common-url-verification-gaps.md) | 證據 only | ExternalRepoC common-url Execute 验证不严 |
| **2i** | [`2i-externalrepoc-user-feedback-pull-execute.md`](2i-externalrepoc-user-feedback-pull-execute.md) | 證據 only | ExternalRepoC user-feedback S0–S4 Execute |
| **2j** | [`2j-externalrepoc-push-execute-skip-verifier-loop.md`](2j-externalrepoc-push-execute-skip-verifier-loop.md) | 负向证据 | ExternalRepoC 05 push Execute：**0 Verifier**、单 Task 包办、`delegation.enabled:false` 误豁免；consumer verifier-after-executor gate 回饋 |
| **2k** | [`2k-externalrepoc-push-post-close-runtime-gaps.md`](2k-externalrepoc-push-post-close-runtime-gaps.md) | 證據 only | ExternalRepoC 05 push **2j 纠偏后**：slice 关闭 vs 用户手验（模版/商户/远程同步）；Worker 拓扑、pull 映射、post-close surgical debt |
| **2l** | [`2l-externalrepoc-common-url-s2-mirror-skip-loop.md`](2l-externalrepoc-common-url-s2-mirror-skip-loop.md) | 负向证据 | ExternalRepoC 03 S2′ mirror：**0 Executor/Verifier**、surgical bypass 滥用、Shell 绕过 preToolUse；2j/2k 教训未内化 |
| **2m** | [`2m-externalrepoc-phase-g-mirror-batch-retrofit.md`](2m-externalrepoc-phase-g-mirror-batch-retrofit.md) | 正负对照 | ExternalRepoC **Phase G-mirror** 批量 retrofit：V-m1–V-m5 + 登记总表；02/01 合规 loop vs 03/2l；stale JVM V5-A 复发 |
| **2n** | [`2n-externalrepoc-push-delivery-s1-s6-compliant-loop.md`](2n-externalrepoc-push-delivery-s1-s6-compliant-loop.md) | 正向证据 | ExternalRepoC **07 push delivery** DEL-S1–S6：6/6 E+V loop、sub-plan `completed`、零 post-close bypass；对照 2j/2k/2l |
| **2o** | [`2o-consumer-tab-scroll-single-vs-delegation.md`](2o-consumer-tab-scroll-single-vs-delegation.md) | 正负对照 | `<PROJECT_ROOT>` tab-scroll 单 session vs 三角色 |
| **2p** | [`2p-externalrepoc-integration-default-cutover-d0-d5.md`](2p-externalrepoc-integration-default-cutover-d0-d5.md) | 正向证据 | Integration 默认切流 INT-D0–D5：6/6 E+V |
| **2q** | [`2q-externalrepoc-transport-inner-only-runtime-gap.md`](2q-externalrepoc-transport-inner-only-runtime-gap.md) | 负向/纠偏 | transport inner-only 假绿 |
| **2r** | [`2r-consumer-player-overlay-mode-a-hit-trap.md`](2r-consumer-player-overlay-mode-a-hit-trap.md) | 负向证据 | player overlay Mode A hit-trap |
| **2s** | [`2s-architecture-ui-pattern-knowledge-plan-review.md`](2s-architecture-ui-pattern-knowledge-plan-review.md) | 跨域 Architecture | UI Pattern Knowledge **plan** review：四責任成立；Knowledge 格仍空；orchestrator 越界疤保留 |
| **2t** | [`2t-apk-capability-handoff.md`](2t-apk-capability-handoff.md) | **partial / 2t-A observational** | 真實 APK Discovery consumer：overlay→[`04`](../04-apk-capability-handoff-boundary.md)；Assessment 持續 `no`（含可重放 sign 後）；**未**套三角色、**未** 2t-B |
| **2u** | [`2u-externalrepoc-p12-r1-mapping-removal-impl-done.md`](2u-externalrepoc-p12-r1-mapping-removal-impl-done.md) | 正负对照 | ExternalRepoC **P12-R1** 删 merchant-product-mapping：完整 E+V；`implementation_done`；V5-M linked、V5-A captcha/Redis defer；merge rebase + post-merge build fix |
| **2v** | [`2v-external-greenfield-consumer-phase2-preflight.md`](2v-external-greenfield-consumer-phase2-preflight.md) | preflight | Greenfield consumer Phase 2 scaffold：brief+backfill；0 E+V |
| **2w** | [`2w-travel-planning-independent-verification.md`](2w-travel-planning-independent-verification.md) | **完成（一輪）** | Travel Planning：分宿 2+3＋車停城堡；獨立 Verifier findings×13；仲裁 fix 預算標題／司機歸屬／09:00 受付／早接駁；不填 Knowledge |
| **2x** | [`2x-consumer-player-variant-matrix.md`](2x-consumer-player-variant-matrix.md) | 負向證據 | Player 單 fixture 假綠與 A-fix-B-break；回饋 acceptance-equivalence gate、方向×裝置×media-path 最小矩陣、decoded-frame oracle |
| **2y** | [`2y-kaizenwms-phase2-spa-scaffold-c1b.md`](2y-kaizenwms-phase2-spa-scaffold-c1b.md) | 正向＋C1b 纠偏 | KaizenWMS Phase 2 SPA scaffold：完整 O→E→V；Verifier 擋 deliverable-only A3；Orchestrator integration smoke 後 `slice_compliant_closed` |
| **2z** | [`2z-kaizenwms-phase3-karma-stale-serve.md`](2z-kaizenwms-phase3-karma-stale-serve.md) | 负向／纠偏 | KaizenWMS Phase 3：Karma-only 假綠＋stale `ng serve`／HMR 失效；回饋 browser e2e DoD、V5-U stale-dev-server、與 2q/2k 同構 |
| **3a** | [`3a-kaizenwms-spawn-friction-skip-loop.md`](3a-kaizenwms-spawn-friction-skip-loop.md) | 负向＋同日正向 | **主反例**：Cursor 機械 bootstrap／primary_source gate 讓 fresh Task 跑不起來 → 誤讀成「三角色被擋」→ 放棄 E／V；consumer §2.1 fallback＋shell-nav Phase 1 正向對照 |
| 2e | [`2e-grandfather-sunset-audit.md`](2e-grandfather-sunset-audit.md) | 完成 | Research 域 grandfather sunset；Q6/Q7/Q8 跨域观察 |

> **漸進遷移**：2026-07-09 起新證據進本目錄；kit 保留傳輸模板與精簡指標。
