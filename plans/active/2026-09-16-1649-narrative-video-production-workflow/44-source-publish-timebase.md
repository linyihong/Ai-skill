# 44 — Source vs publish timebase（contract refinement）

Status: candidate invariant（未凍結進 Phase 1 十條）
Companion to [`03-architecture-invariants.md`](03-architecture-invariants.md)
Workflow SoT: [`../../../workflow/narrative-video-production/source-publish-timebase.md`](../../../workflow/narrative-video-production/source-publish-timebase.md)
Evidence: [`evidence/2026-10-01-source-publish-timebase.md`](evidence/2026-10-01-source-publish-timebase.md)

## Why now

Phase 3 dogfood 出現「1× 參考片 vs sped 發布片」對照時，wall-clock 同秒互比會製造假錯位。這不要求重開 Phase 或重做 Timeline IR 大局；只要補 **timebase／transform** invariant。

## Candidate invariant 17

**All source understanding and evidence extraction MUST use the canonical source timebase. Publish-time speed transforms MUST NOT be treated as new evidence.**

EDR／Timeline IR 應能保存：

- canonical／source timebase 時戳
- publish timebase 投影（若已 assemble）
- `timeline_transform`（至少 `constant_speed` + `rate`）

## Explicitly out of scope

- 不改十條凍結 invariant 編號語意
- 不接 runtime／不註冊 route
- 不把 ffmpeg `setpts`／`atempo` 寫進 execution-flow 步驟
- 不強制某一固定 rate（1.5 只是 dogfood 例）

## Acceptance（本 companion）

- [x] Workflow 契約檔存在並從 README／assemble／captions 可達
- [x] EDR record 有 optional `timeline_transform`／timebase 註記
- [x] Plan evidence 索引已列本 run
- [ ] Product adapter 明示 transform metadata（下一輪；見 `<PROJECT_ROOT>`）
