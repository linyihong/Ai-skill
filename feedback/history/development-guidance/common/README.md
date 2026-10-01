# Development Guidance — Common Lessons

| 檔名 | Status | 標題 | 一句話摘要 |
|------|--------|------|-----------|
| `2026-10-01_093000-ocr-latin-boundary-geometry-before-lexical.md` | candidate | OCR Latin boundary: geometry before closed-class | boxes→geometry→lexical candidates；closed-class 非第一刀；exact-match 單字不拆 |
| `2026-09-30_155700-release-pool-client-before-nested-checkout.md` | candidate | Release pool client before nested same-pool checkout | Never hold a DB pool client across an await that also checks out from the same pool; saturation hang |
| `2026-05-05_194400-contract-first-development-flow.md` | candidate | Contract-first development flow | Start development from product intent, split bounded contexts, write BDD, define Domain, Architectur |
| `2026-05-05_200500-existing-project-doc-backfill-bdd-required.md` | candidate | Existing project doc backfill requires complete BDD | When opening app-development-guidance on an already implemented project, audit and backfill missing  |
| `2026-05-05_201000-missing-requirements-block-development.md` | candidate | Missing requirements block development | If required behavior, contracts, errors, security, storage, ownership, or tests are missing, the age |
| `2026-05-06_081600-change-intake-before-code.md` | candidate | Change intake before code | Before code changes, review the planning artifact and classify the work as new requirement, bug fix, |
| `2026-05-06_082000-separate-regression-from-new-code-validation.md` | candidate | Separate regression from new code validation | Separate tests that guard existing behavior from tests that prove new code; high total coverage does |
| `2026-05-06_083000-embedded-hardware-product-flow.md` | candidate | 2026-05-06_083000-embedded-hardware-product-flow | 既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。 |
| `2026-05-06_083200-implemented-first-contract-governance.md` | candidate | 2026-05-06_083200-implemented-first-contract-governance | 既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。 |
| `2026-05-06_103200-product-brief-validation-gate.md` | candidate | Product Brief validation gate | Product Brief / 企劃書 claims must be validated, labeled as assumptions, asked as open questions, scope |
| `2026-05-06_150000-same-session-doc-sync-after-code-fix.md` | candidate | 2026-05-06_150000-same-session-doc-sync-after-code-fix | 既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。 |
| `2026-05-07_081800-keep-project-incidents-out-of-skills.md` | candidate | Keep Project Incidents Out Of Skills | Reusable skills should capture generalized causes, decisions, and validation loops, while concrete p |
| `2026-05-07_122800-performance-test-release-gate.md` | candidate | Performance Test Release Gate | 功能正確不代表效能可上線；會影響容量、延遲或資源的變更需要效能預算與對應測試。 |
| `2026-05-07_152100-private-live-adapter-smoke-gate.md` | candidate | Private Live Adapter Smoke Gate | When analysis supports SDK core design but live replay needs private host, service, session, signing |
| `2026-05-07_153600-schema-derived-synthetic-fixtures.md` | candidate | Schema-Derived Synthetic Fixtures | When live payloads are private but response schemas are known, create synthetic schema-compatible fi |
| `2026-05-07_154400-analysis-sdk-contract-drift-gate.md` | candidate | Analysis-To-SDK Contract Drift Gate | When APK analysis closes or reclassifies live-readiness boundaries, immediately audit downstream SDK |
| `2026-05-07_160200-media-metadata-private-decrypt-boundary.md` | candidate | Media Metadata Private Decrypt Boundary | When live media URLs, signed query values, encrypted files, or wrapped media keys are private, SDK c |
| `2026-05-11_093100-public-mirror-drift-gate.md` | candidate | Public mirror drift gate | Public SDK delivery repositories can accumulate release-facing runtime or README fixes; check both d |
| `2026-05-11_093200-session-login-concurrency-matrix.md` | candidate | Session login concurrency matrix | SDKs that self-login must test multi-process startup, per-device session reuse, refresh single-fligh |
| `2026-05-12_013400-private-adapter-inside-test-module.md` | candidate | Private Adapter Inside Test Module | 當 live smoke 測試需要 private adapter（簽章、解密、不透明參數提供者），但這些實作不能進入公開 SDK reactor 時，應將它們以 test-scoped class  |
| `2026-05-12_043500-dart-io-httpclient-bypasses-java-frida-hooks.md` | candidate | Dart `dart:io` HttpClient Bypasses Java-Level Frida Hooks | 當 Flutter/Dart 應用使用 `dart:io` HttpClient 發送 HTTP 請求時，這些請求會完全繞過 Java 層（OkHttp、Netty、`java.net.Socket` |
| `2026-05-13_094500-anti-bot-gateway-blocks-external-sdk.md` | candidate | Anti-Bot Gateway Blocks External SDK Calls via TLS Fingerprint | 當目標 API 使用 anti-bot gateway（如 PerimeterX）時，外部 JVM SDK 無法直接呼叫 API，因為 gateway 會驗證 TLS fingerprint；唯一繞過 |
| `2026-05-13_095400-empty-post-body-bypasses-anti-bot-via-device-proxy.md` | candidate | Empty POST Body Bypasses Anti-Bot Gateway When Routed Through Device Proxy | 當 App 使用 device proxy 轉發請求時，anti-bot gateway（如 PerimeterX）的阻擋可以透過「空 POST body + 有效 `eh` header」繞過，因為 |
| `2026-05-15_173100-write-to-file-truncation-large-file.md` | candidate | write_to_file 大檔案截斷風險與防護 | 當使用 `write_to_file` 寫入超過 ~700 行的 Java 原始檔時，內容可能被截斷（truncated），導致編譯錯誤。應改用 `apply_diff` 分段寫入，或寫完後立即編譯驗 |
| `2026-05-18_141900-json-substring-matching-numeric-trap.md` | candidate | 2026-05-18_141900-json-substring-matching-numeric-trap |  |
| `2026-05-18_142700-list-files-verify-domain-path-before-write.md` | candidate | 2026-05-18_142700-list-files-verify-domain-path-before-write | 當要寫入 feedback lesson 到 `feedback/history/<domain>/` 時，**必須先 `list_files` 確認目標 domain 目錄確實存在**，不能依賴內部 |
| `2026-05-18_160000-registry-first-workflow-activation.md` | candidate | 2026-05-18_160000-registry-first-workflow-activation | Workflow 觸發條件與依賴應寫在 `routing-registry.yaml` 各 `route.workflow.*` 的 `activation_triggers` / `required |
| `2026-05-19_160800-evidence-first-before-live-blocker.md` | candidate | Evidence-first before live blocker | Before declaring a live/integration prerequisite missing, first search project docs, sanitized API e |
| `2026-05-21_142300-document-migration-map-before-deleting-legacy-surface.md` | candidate | Document Migration Map Before Deleting Legacy Surface | Legacy script deletion must be preceded by a developer-facing migration map that names the new Go ow |
| `2026-05-21_172702-migration-seeder-anti-patterns.md` | candidate | 2026-05-21_172702-migration-seeder-anti-patterns | 把大量業務資料以巨型 `INSERT` 包進 schema migration，會強制資料 lifecycle 與 schema lifecycle 綁定，造成部署、review、業務速度與環境漂移等 |
| `2026-05-21_172703-vendor-integration-architecture.md` | candidate | 2026-05-21_172703-vendor-integration-architecture | 整合超過 3 個外部廠商時，需在 Adapter / Compile-time submodule / Plugin SPI / Out-of-process service / Hybrid 之間選 |
| `2026-05-21_172704-dual-token-security-audit.md` | candidate | 2026-05-21_172704-dual-token-security-audit | 系統同時存在兩套以上 token 機制（JWT + JWE、HMAC + 對稱加密、平台 token + 廠商回調 token）時，每一道接縫都是簽章誤用、replay、key 混用、降級攻擊的高風險 |
| `2026-05-21_220226-knowledge-update-flow-master-doc-required.md` | candidate | 2026-05-21_220226-knowledge-update-flow-master-doc-required | 當「有新知識要寫入 Ai-skill」時，必須先讀 `governance/lifecycle/knowledge-update-flow.md` 作為流程總索引；**不得只讀子流程文件（如 `int |
| `2026-05-21_231500-executable-contract-boundary.md` | candidate | 2026-05-21_231500-executable-contract-boundary | 流程、gate、activation、blocking condition、required evidence 或 failure action 若會影響 agent 執行，必須提供 owner-la |
| `2026-05-22_092400-test-first-framework-upgrade.md` | candidate | 2026-05-22_092400-test-first-framework-upgrade | Framework / runtime / governance 升級時 validation scenarios 必須寫在 runtime 實作之前；scenarios 是 acceptance c |
| `2026-05-22_102818-migration-feature-bundling-antipattern.md` | candidate | 2026-05-22_102818-migration-feature-bundling-antipattern | 大型系統改版（rewrite / migration / platform 升級）時，**新版必須先達成 parity 才能加新功能**；同時混入新功能會讓驗證失去 ground truth，bug  |
| `2026-05-22_104646-module-count-discipline.md` | candidate | 2026-05-22_104646-module-count-discipline | Repo 內模組（build system module / workspace package / subproject）數量隨業務需求增長，但「畫一個模組」有工程成本；N 失控時 build ti |
| `2026-05-22_plan-first-decision-promotion.md` | candidate | 2026-05-22_plan-first-decision-promotion | 架構決策的提案、討論、alternatives 評估在 `plans/active/<plan>.md` §Decision Rationale 完成；只有 plan completed 且通過 AD |
| `2026-05-27_092801-intelligence-layer-bypass-via-tool-adapter.md` | candidate | 2026-05-27_092801-intelligence-layer-bypass-via-tool-adapter | 當 agent 取得跨工具可重用的 agent 行為洞見（如「多層強制執行設計模式」）時，不得因為主題「關於某工具」就直接寫進 `ai-tools/<tool>.md`（P3 tool adapter |
| `2026-05-27_154500-ai-codegen-validation-rate-parity.md` | candidate | 2026-05-27_154500-ai-codegen-validation-rate-parity | 既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。 |
| `2026-06-08_103300-component-traceability-marker-depth.md` | candidate | Component traceability marker depth | 新增或提升共用 UI component 時，BDD / contract 不應只驗 component 名稱或檔案存在，還要驗 feature refs 與最小實作語意 marker。 |
| `2026-06-08_141100-blob-manifest-uri-rewrite-test.md` | candidate | Blob manifest URI rewrite test | 把 HLS / manifest 內容轉成 `blob:` URL 前，必須先把 manifest 內會被播放器再次請求的 key、segment 或 asset URI 改成可解析的絕對 URL，並 |
| `2026-06-08_141200-browser-layout-overflow-integration.md` | candidate | Browser layout overflow integration | RWD 或視覺縮放問題不能只靠 CSS marker、snapshot 或人工感覺驗證；要用真實瀏覽器 viewport 量 `document`、`body`、app shell、固定導覽與 scr |
| `2026-06-18_play-view-kpi-sql-pass-dom-fail-browser-gate.md` | candidate | Play-view KPI: SQL/API pass, DOM still wrong | 用户可见的播放量 / KPI 标签争议时，**API 或 listing SQL 与 DB 一致不足以结案**；须在同一 `video_id`、同一路径上增加 **browser DOM 对照**，才 |
| `2026-07-15_140200-correct-sign-length-still-403-means-anti-bot-not-sign.md` | candidate | Correct offline sign length but HTTP 403 usually means anti-bot, not sign formula | 主機離線 signer 產出的 `sign` **長度與 in-app fingerprint 一致**，但 API 仍 **HTTP 403**（edgesuite／challenge HTML）時 |
| `2026-07-17_171500-mpegurl-maybe-managed-mediasource-prefer-hlsjs.md` | candidate | mpegurl "maybe" + ManagedMediaSource prefer hls.js | 對 Apple WebKit，`canPlayType("application/vnd.apple.mpegurl") === "maybe"` 且存在 `ManagedMediaSource` 時 |
| `2026-07-23_071200-injectable-wall-clock-for-day-boundary-logic.md` | candidate | 2026-07-23_071200-injectable-wall-clock-for-day-boundary-logic | Time-dependent business logic (local day, near-midnight windows, quotas) must take an injectable clo |
| `2026-07-30_135700-unreachable-test-surface-delete-over-repair.md` | candidate | 2026-07-30_135700-unreachable-test-surface-delete-over-repair | 當某個測試 surface 無法被任何 gate 執行時，先確認 canonical 策略是否早已把測試放在別處；若是，正解是「保留斷言、刪除 surface、加機械禁令」，而不是修好設定讓它跑起來。 |
| `2026-07-30_135800-run-new-ban-against-existing-tree.md` | candidate | 2026-07-30_135800-run-new-ban-against-existing-tree | 新增機械檢查後，先對整棵既有樹跑一次再接進 verification chain：同一 root cause 的兄弟實例會一起浮現，而那會改變這次任務的 scope。 |
| `2026-07-30_221700-docx-text-extraction-without-converter.md` | candidate | 缺 document converter 時的 .docx 文字抽取備援 | `.docx` 是 ZIP + XML，所以在沒有 `pandoc` / office suite / Python 的環境裡，仍可用平台內建解壓 + 任一可用 runtime 做標籤剝除來取得可審閱 |
| `2026-08-03_211615-windows-cmd-shim-needs-shell-not-cmd-suffix.md` | candidate | Windows 上的 .cmd shim 要靠 shell 啟動，不是補 .cmd 副檔名 | Node 在 Windows spawn 套件管理器的 wrapper 指令時，`shell: false` 找不到 shim（ENOENT），改指名 `<cmd>.cmd` 也會被拒（EINVAL） |
| `2026-09-04_101559-build-gate-must-restore-not-only-build.md` | candidate | 以 `--no-restore` 建置的 gate 驗的是還原圖，不是工作樹 | 當套件／專案的相依關係來自 restore 產生的鎖定檔而非組件 metadata 時，用 `--no-restore` 建置的 gate 驗證的是「上次 restore 時的相依圖」，同一棵樹在不同 |
| `2026-09-04_101800-derived-state-before-blaming-shared-baseline.md` | candidate | 歸因給共用基線之前，先排除本地衍生狀態 | 「原始碼輸入相同」加上「可穩定重現」不足以證明共用基線壞掉；這兩件事在本地衍生狀態過期時同樣成立，而衍生狀態不會出現在原始碼 diff 裡。 |
| `2026-09-04_135746-gate-that-runs-only-at-commit-is-one-bypass-from-gone.md` | candidate | 只在 commit 階段跑的 gate，繞過一次就等於永久失效 | 若某個 gate 只掛在 commit 而不掛在 push，它檢查的是「這次提交的內容」而不是「分支的狀態」，任何一次繞過或跳過都不會再被任何後續動作抓回來。 |
| `2026-09-04_135900-lazy-migration-on-background-flush-needs-tolerant-reads.md` | candidate | 惰性遷移若由背景批次觸發，讀取端必須自行容忍未遷移狀態 | 由「第一次寫入」觸發的就地遷移，會在服務啟動到第一次寫入之間留下一段讀取視窗；若讀取端不容忍未遷移狀態，升級後的第一批查詢會回 5xx，而症狀會以 flaky test 的形式出現。 |
| `2026-09-04_140100-index-link-gate-must-accept-the-resolvable-form.md` | candidate | 檢查索引連結的 gate，必須接受從該索引真的能解析的那種寫法 | 要求某個路徑字串出現在一份 Markdown 索引裡的 gate，若比對的是 repo 根相對路徑，就會拒絕唯一正確的相對連結，並接受寫上去會壞掉的那幾種。 |
| `2026-09-04_142000-derived-closure-must-not-expand-its-own-derivation-table.md` | candidate | 機械推導的相依封包，不可展開推導規則本身 | 用「跟著引用展開」推導每個單位的相依封包時，若把記錄各分支入口的對照表也一起展開，封包會坍縮回原本要消除的全域集合——而且數字看起來完全正常。 |
| `2026-09-04_142100-one-shot-watch-must-terminate-on-its-own-condition.md` | candidate | 只需要一次通知的監看，必須能自己結束 | 用 `tail -f` 這類永不結束的指令去等一個一次性事件，事件發生後監看仍會保持武裝到逾時，留下無人察覺的殘留程序。 |
| `2026-09-15_012600-stale-peer-refused-refresh-not-fail-job.md` | candidate | Stale peer refused → refresh discovery; do not fail the job | When a long-lived TCP peer returns `ECONNREFUSED` / host-unreachable, refresh the official discovery |
| `2026-09-15_110000-multi-seat-s2c-must-not-assume-fixed-row-glass.md` | candidate | 2026-09-15_110000-multi-seat-s2c-must-not-assume-fixed-row-glass | 既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。 |
| `2026-09-15_112000-wire-rewards-parse-collapse-richest-block.md` | candidate | 2026-09-15_112000-wire-rewards-parse-collapse-richest-block | 既有 lesson 的結論維持於本檔原始內容；此 closure 補記不新增專案事實。 |
| `2026-09-28_131830-runtime-image-must-copy-baked-repo-root-data.md` | candidate | Runtime image must ship data under baked repoRoot | 若 build 把 `repoRoot` 烤成絕對路徑，runtime 映像必須把伺服器仍會 `fs.read` 的資料樹 COPY 到同一路徑，不能只帶編譯產物目錄。 |

| `2026-09-28_133500-payline-chips-mirrors-vs-scaled-tags.md` | candidate | Payline CHIPS: strip line mirrors before scaling tags | Wire 常為每條線贏一筆 `<CHIPS>`；對獎時先按 lineWin 1:1 剝離镜像，剩餘標籤才當 extras |
| `2026-09-29_105651-translation-decision-selection-adapter-finality-gates.md` | candidate | Translation Decision: Selection adapter + mechanical Finality | LLM 只當 Selection Actor；截斷源文／親屬殘留用機械 Finality／Validation，禁止 case map 與 sole gloss |
| `2026-09-29_132500-same-lang-skip-must-block-pretranslate-cache.md` | candidate | Same-lang skip must block pretranslate cache (OCR Latin junk) | `skip_translate`／源≈目標時必須短路整集預譯與 `translations/{lang}` 寫入；不可只靠 per-line detect，否則 OCR 拉丁垃圾會被當成外語翻譯 |
| `2026-09-30_081600-ocr-recognition-lang-must-not-erase-latin-word-boundaries.md` | candidate | OCR recognition language must not erase Latin word boundaries | `recognition_language`≠`observed_script`；Latin 單框無空格要 boundary／derived，不可因 ch 路徑刪空格或覆蓋 raw |
| `2026-10-01_090000-ocr-join-token-seam-not-cumulative-script.md` | candidate | OCR join must use token seam not cumulative script | parts 已切、derived 黏：Latin\|Latin 必空格；禁止 cumulative mixed 否決；parts 不可被 text 覆蓋 |

Total: 68
