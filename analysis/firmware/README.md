# Firmware Analysis（stub）

Firmware 映像的區域辨識、抽取與 handoff 到 native 分析的最小入口。

**狀態**：stub。深度 dogfood 另開 spike（parent plan Q6）。

前置工具思路：Binwalk／Unblob（caller-supplied；本庫不安裝）。抽取後的 native blob 走 [`../binary/`](../binary/README.md)。

Evidence 紀律：[`../reverse-engineering/evidence-contract.md`](../reverse-engineering/evidence-contract.md)。授權：[`../../enforcement/authorization-scope.md`](../../enforcement/authorization-scope.md)。

計劃：[`../../plans/active/2026-10-10-1328-rea-analysis-capability-hardening/_plan.md`](../../plans/active/2026-10-10-1328-rea-analysis-capability-hardening/_plan.md)

← [回到 analysis/](../README.md)
