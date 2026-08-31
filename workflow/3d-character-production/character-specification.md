# Character Specification（欄位入口）

只說明去哪填。分類（must-preserve / allowed-variation / forbidden-drift）的
**欄位定義**在 [`records/character-specification.yaml`](records/character-specification.yaml)。
`outfit_is_must_preserve` 由該檔 derived_flags 計算，供 identity invalidation 讀取——
不要在本檔重寫 outfit 規則。
