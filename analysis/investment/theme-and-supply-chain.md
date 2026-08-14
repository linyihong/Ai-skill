# Theme & Supply-Chain Decomposition（主題／供應鏈拆解）

把「看好某主題」拆成可驗證的節點、choke-point 與標的曝險，避免直接跳到喊單。

對齊 Phase 1 ①（CPO／optical）形狀：主題地圖 → 節點證據 → 標的候選（可選）→ uncertainty。

## 何時使用

- `theme-research`、主題型 `periodic-sweep`  
- 配置前需要「主題曝險從哪來」  
- 使用者丟單一敘事（「AI 光通訊」）但未指定 ticker

## 拆解步驟

1. **定主題邊界**  
   - 一句定義＋明確排除（例：含 CPO／光模組；排除純終端消費電子）  
   - 市場範圍：預設 TW＋US（intake 可改）

2. **畫供應鏈節點**（最少一層上下游）  

   | 層級 | 問什麼 |
   | --- | --- |
   | Demand | 誰在買？Capex／部署節奏？ |
   | Mid | 模組／交換／光學／封裝等中游 |
   | Supply choke | 設備、原料、製程、專利、產能瓶頸 |
   | Substitutes | 可替代技術／路徑 |

3. **標 choke-point**  
   - 每個 choke：為何是瓶頸、證據 authority、是否可緩解、時效  
   - 無 A／B 支撐的 choke → 標 `hypothesis`，不得當配置主因

4. **節點 → 曝險類型**（尚未到下單）  

   | 曝險類型 | 含義 |
   | --- | --- |
   | Pure-play | 營收高度綁該節點 |
   | Partial | 多業務之一 |
   | Enabler | 賣鏟人／設備／材料 |
   | Derivative | ETF／主題基金（另記費用） |

5. **輸出主題地圖 artifact 形狀**  

   ```text
   Theme definition + exclusions
   Node table (layer / role / choke? / evidence grade / as-of)
   Open questions
   Candidate names (optional) — 每名附：為何掛此節點、authority、uncertainty label
   What would falsify the theme thesis
   ```

## 品質檢查

| 檢查 | Fail 若… |
| --- | --- |
| 邊界清楚 | 主題與排除無法一句說清 |
| 證據分級 | 地圖主幹只靠 D／無 as-of |
| 無假精確 | 出現未校準「60% 會爆發」類點估計 |
| 可 falsify | 沒有「何種證據會推翻主題」 |
| 不越權 | 地圖直接變成「必買清單」且無 human selection |

## 與 name-diligence 的銜接

- 主題地圖產出 **候選標的清單** 後，單一名字進 [`news-trend-summary.md`](news-trend-summary.md)＋公開財報 diligence。  
- 不可用主題敘事替代公司 A 級揭露。

## 參考風格（非事實來源）

公開研究帳的供應鏈敘事節奏可當 **discovery 風格參考**；其結論一律當 **D**，見 [`researcher-note-contrast.md`](researcher-note-contrast.md)。
