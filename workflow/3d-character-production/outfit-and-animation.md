# Outfit and Animation

僅當 deformation 與 facial-expression acceptance 都 pass，且 identity 仍為
`accepted + current`。換裝若發生：**寫 mutation_event**，由
[`records/identity-acceptance.yaml`](records/identity-acceptance.yaml) 更新 `validity`；
更新完成前不得進 export；本檔不寫「換裝是否重審」。

動作／clip 證據進 pack 的 `actions`／`motion_play`。
