# Narrative Video Production Workflow

`workflow/narrative-video-production/` 負責**敘事向影音**從 brief 到發布 evidence 的執行順序與 gate。
Acquisition（生成／實拍／剪輯／未來工具）可換；不得反向定義本 workflow。

> **狀態**：Phase 2 workflow contract。**沒有** `route.workflow.narrative-video-production`（Phase 4 才考慮）。
> YAML **不**投影（`runtime_projection.enabled: false`）。Detector 應維持 no_match。
> 架構護欄：plan companion [`03-architecture-invariants.md`](../../plans/active/2026-09-16-1649-narrative-video-production-workflow/03-architecture-invariants.md)。**驗收看各檔 gate，不只看 03。**

> **執行契約，不重新解釋契約。** 欄位在 [`records/`](records/README.md)。Markdown 說明何時讀、推進條件；heuristic 數字（CPS、band 秒數）留給 profile／dogfood。

## 何時讀哪個檔

| 認知階段 | 檔案 | load_when |
| --- | --- | --- |
| Lifecycle | [`execution-flow.md`](execution-flow.md) | 任何本 domain 任務 |
| Brief | [`intake.md`](intake.md) | 平台／時長／承諾未鎖 |
| 源世界 | [`source-bible.md`](source-bible.md) | 需要人物／集數穩定 id |
| 可剪素材 | [`material-clip-catalog.md`](material-clip-catalog.md) | 進庫、查找、時長 band |
| 敘事模板 | [`narrative-template-catalog.md`](narrative-template-catalog.md) | 選主模板 |
| 匹配腳本 | [`matching-script.md`](matching-script.md) | 每個 shot 選 clip |
| 鏡頭表 | [`script-and-shot-list.md`](script-and-shot-list.md) | 模板 beat ↔ `shot_id`；可選 `dialogue.semantic_context` |
| 決策 SoT | [`edit-decision-record.md`](edit-decision-record.md) | 開 EDR、mutation、人讀投影 |
| 字幕語系 | [`captions-and-locales.md`](captions-and-locales.md) | locale pack／三閘 |
| 成片對帳 | [`assemble-and-qc.md`](assemble-and-qc.md) | 時間線 vs EDR；fresh verifier |
| 發布證據 | [`publish-outcome.md`](publish-outcome.md) | 窗口後填 evidence status |
| Eligibility | [`artifact-gates.md`](artifact-gates.md) | 每一 stage 推進與 completion |
| Profile | [`profiles/README.md`](profiles/README.md) | recap／original／split-promo；首輪空 |

## Invariant 落點（驗收用）

| # | Invariant | 可驗證落點 |
| --- | --- | --- |
| 1 | EDR 是決策 SoT；mp4 是產出 | [`edit-decision-record.md`](edit-decision-record.md)、[`assemble-and-qc.md`](assemble-and-qc.md) |
| 2 | Bible ≠ Catalog | [`source-bible.md`](source-bible.md)、[`material-clip-catalog.md`](material-clip-catalog.md) |
| 3 | 查找只回既有 `clip_id` | [`material-clip-catalog.md`](material-clip-catalog.md) `retrieval_contract` |
| 4 | Constraints 定義可行集 | [`matching-script.md`](matching-script.md)、[`records/matching-script.yaml`](records/matching-script.yaml) |
| 5 | Selection 是明示 policy | 同上；缺 `selection.policy` = 閘失敗 |
| 6 | Assembly 對得上 EDR | [`assemble-and-qc.md`](assemble-and-qc.md) |
| 7 | Locale 三閘獨立 | [`captions-and-locales.md`](captions-and-locales.md) |
| 8 | publish-ready 需 fresh verifier | [`artifact-gates.md`](artifact-gates.md) |
| 9 | Outcome 是 evidence 不是真理 | [`publish-outcome.md`](publish-outcome.md) |
| 10 | Runtime 延後 | 本檔狀態列 + records `runtime_projection.enabled: false` |

## 核心原則

1. 填 [`records/`](records/)；禁止從成片反推剪輯理由。
2. 空可行集 → blocked；禁止從集外硬挑。
3. 作者自驗可到 `cut-ready`；`publish-ready` 必須獨立 verifier。
4. 工具名、TTS／翻譯 provider、ffmpeg 步驟不進 canonical 流程。
5. 與 [`3d-character-production`](../3d-character-production/README.md) 分界：那邊是可驅動角色資產；這邊是敘事成片。寫產片器走 `software-delivery`。
