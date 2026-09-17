# Observation — editorial vs narrative transition

**Run ID**：2026-09-17-editorial-vs-narrative-transition  
**Kind**：Phase 3 observation（**不是** `contract_gap`；**不**擴 workflow）  
**原則**：不要一開始讓 AI 判斷「這裡是劇情轉場」。

## 兩層

| 層 | 是什麼 | 誰做 |
| --- | --- | --- |
| Editorial transition | 片子怎麼從 A 切到 B（hard cut／fade／dissolve／wipe／black／match cut） | CV／video analysis |
| Narrative transition | 故事從 A 到 B 變了什麼（地點／時間／角色群／事件／視角／回憶／夢／平行） | 語義推理 |

機械層產出 `transition_candidate`（相鄰 shot 的 visual／location／cast／audio／OCR／ASR 變化）。AI 再填 `narrative_transition`（type、from/to context、function）。OCR 突變（「三年後」「東京」）是 candidate evidence，見 [`2026-09-17-visual-text-evidence.md`](2026-09-17-visual-text-evidence.md)。

## 與 catalog

Clip 列仍是素材。shot／scene **關係**較適合獨立 graph 載體，不把全部 transition metadata 塞進每一列 clip。Graph 是關聯載體，不是 architecture core。

**Shot ≠ Scene。** 一個 scene 可含很多 shot。對劇情分析／matching 更有價值的常是 scene（location、characters、time、events、dialogue、`shots[]`）。聚合路徑：shot boundary → scene candidate → narrative transition。

## 第一版建議收（外部工具，非 contract）

非 AI：shot boundary、keyframes、duration、visual difference、audio／ASR／OCR change、black／fade。  
之後 AI：scene boundary、location／time／character transition、narrative function、flashback／dream／parallel。

真實片子要數：哪些 candidate 真被 matching／EDR 消費；有多少「畫面切了」不是敘事轉場。
