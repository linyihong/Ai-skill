# Rigging and Deformation

填 [`records/deformation-acceptance-set.yaml`](records/deformation-acceptance-set.yaml)
body deformation 最小充分 set。集合只含 rig／weights 的 body pose；extreme facial
expression 由下一個獨立 stage 擁有。任一 observation `fail` → `decision: fail` → 不得進表情。

推進前再讀 `identity_acceptance.decision == accepted` 且 `validity == current`。
失敗 rollback：`rig_weights`。
