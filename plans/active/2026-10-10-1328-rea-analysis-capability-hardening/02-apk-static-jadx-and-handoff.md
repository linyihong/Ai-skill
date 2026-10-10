---
id: 2026-10-10-1328-rea-02-apk-static-jadx-and-handoff
plan_kind: sub
status: completed
owner: larrylin/cursor-session
created: 2026-10-10
parent: 2026-10-10-1328-rea-analysis-capability-hardening
required_for_completion: true
sub_plan_reason: >
  Strengthen APK static JADX agent path and explicit handoff to existing
  dynamic Frida/traffic mainline without replacing it.
---

# Sub-plan 02 — APK static JADX + dynamic handoff

Parent: [`_plan.md`](_plan.md) Phase 2.

## Acceptance

- [x] `analysis/apk/static-jadx-path.md`
- [x] tools-and-failures / traffic-triage / apk README 更新
- [x] workflow apk execution-flow 加 static-first 分支（§2.0）
- [x] 明確標示 Android static 引擎不做 runtime／native-lib／Frida
