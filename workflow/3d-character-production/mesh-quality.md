# Mesh Quality

僅當 Identity 閘允許（`decision==accepted` 且 `validity==current`）。

填 [`records/mesh-qa-report.yaml`](records/mesh-qa-report.yaml)。
`decision` 只表示 geometry 是否可進 rig，必須滿足 topology、separation、holes、
scale/orientation 與 repairability。`uv_material` 另由 facial／export gate 讀取；
fail 或 not-evaluated 只能具名 defer 到 prototype，不能 silent pass 到 runtime-ready。

失敗 rollback：`mesh_or_candidate`。不得用 rig 權重掩蓋 mesh fail。
