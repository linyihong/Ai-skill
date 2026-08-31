# Mesh Quality

僅當 Identity 閘允許（`decision==accepted` 且 `validity==current`）。

填 [`records/mesh-qa-report.yaml`](records/mesh-qa-report.yaml)。
`decision` 由該檔 `fail_when_any` 計算。未檢查的項不得記 PASS。

失敗 rollback：`mesh_or_candidate`。不得用 rig 權重掩蓋 mesh fail。
