# Repo 命名候選 — AI-native Cognitive Execution System

> **產出脈絡**：repo slug `Ai-skill` 為早期暫名；2026-05-26 landing page 改版時
> 已確定正式定位為「AI-native Cognitive Execution System」（AI-native 認知執行系統 /
> knowledge OS：治理、記憶、驗證閉環），並註明 slug 遲早要改（見根 `README.md` L11 說明）。
> 本檔為改名前置的候選發散，由 delegation loop 的隔離 worktree 產出。
> **最終選名決定保留給 maintainer** — 本報告只提供候選、理由與撞名風險註記。
> 產出日：2026-07-09。

## 篩選準則

- 好念、好記、slug 短（單字優先）。
- 呼應系統氣質：**知識比工具長壽**（portable / versionable）、**可執行治理**（contracts / gates）、
  **驗證閉環**（obligation ledger / validation loop）、**跨 agent 可攜**（vendor-independent）。
- 撞名檢查基於模型知識（cutoff 2026-01）之已知專案；**定名前仍應做一次
  GitHub / npm / PyPI / crates.io / 商標即時搜尋**，此為候選階段的合理盡職邊界。
- 光譜刻意發散：嚴肅（哲學詞源）→ 工程隱喻 → 玩味。

## 候選清單（12）

### 偏嚴肅：詞源直指定位

| # | Slug | 顯示名稱 | 為什麼適合 | 撞名風險 |
|---|------|----------|-----------|---------|
| 1 | `phronesis` | **Phronesis｜實踐智** | 亞里斯多德的「實踐智慧」：只能從經驗累積、且直接指導行動的判斷力——正是 `intelligence/` 層與「人的經驗隨時間複利」原則的一詞概括。 | 低。學術詞，無知名同名工具；發音（fro-NEE-sis）需適應。 |
| 2 | `noema` | **Noema｜所思** | 現象學術語：思想的**結構化對象**，獨立於思考者存在——「知識獨立於任何單一 agent / 模型」的哲學版本。短、好念。 | 中低。Noema Magazine（Berggruen Institute 刊物）同名但領域不同。 |
| 3 | `vademecum` | **Vade Mecum｜隨行典** | 拉丁文「與我同行」：隨身攜帶的參考手冊。Agent 換工具、換模型時「帶著走的那本典」——可攜性的最直白隱喻。 | 低。古詞，無知名軟體專案佔用。 |
| 4 | `exocortex` | **Exocortex｜外腦** | 「體外皮質」：掛在人與 agent 之外、可檢查可版本化的認知層。科幻感適中，工程師一聽就懂。 | 中。Exocortex（加拿大 3D 軟體公司，Clara.io）同名，但領域遠、近年不活躍。 |

### 中間帶：工程 / 航海 / 記憶隱喻

| # | Slug | 顯示名稱 | 為什麼適合 | 撞名風險 |
|---|------|----------|-----------|---------|
| 5 | `engram` | **Engram｜憶痕** | 神經科學的「記憶痕跡」：記憶的物理實體。知識不存在 hosted memory 而是留下可版本化的痕跡——memory 子系統的精準對應。 | 中低。有零星同名小型 ML/memory 實驗專案與 Destiny 2 遊戲用語；無壟斷性大專案。 |
| 6 | `cairn` | **Cairn｜疊石** | 山徑上的疊石路標：前人堆給**後來的旅人**（下一個 agent、下一代模型）指路。一音節、好拼、意象完全命中「知識比工具長壽」。 | 中低。有零星同名小工具；注意勿與圖形庫 cairo 混淆（拼字不同）。 |
| 7 | `plumbline` | **Plumbline｜準繩** | 建築工的鉛垂線：**外部的、不可協商的垂直標準**，每次施工都拿它校驗——gates / validators / 驗證閉環的隱喻。中文「準繩」古雅精準。 | 低。無知名同名專案。 |
| 8 | `keelson` | **Keelson｜內龍骨** | 船的內龍骨：鎖在龍骨上、看不見卻讓整艘船保持結構的縱梁。無論誰掌舵（哪個 agent），航向穩定性來自這根梁——governance spine 的意象。 | 低。冷僻船舶詞，避開了 `keel` 本身的零星撞名。 |
| 9 | `sextant` | **Sextant｜六分儀** | 六分儀讓任何水手對著**同一片星空**定位自己。不同 agent 讀同一套 canonical source、以 routing 定位該讀哪裡——「共同可信來源」的儀器版。 | 中低。有零星同名 devtools 小專案，無大型佔用。 |

### 偏玩味：自嘲與魔法

| # | Slug | 顯示名稱 | 為什麼適合 | 撞名風險 |
|---|------|----------|-----------|---------|
| 10 | `orrery` | **Orrery｜行星儀** | 十八世紀的太陽系機械模型：每個齒輪、每顆行星的運動**全部可見、可檢查**——「executable governance 不是黑盒」的優雅玩笑。 | 低。冷僻古董詞，無知名專案。 |
| 11 | `canonry` | **Canonry｜典章司** | 自造詞：保管並**強制執行** canon（可信正典）的衙門。把 gates / obligation ledger 的官僚感自嘲進名字裡，反而誠實。 | 低。自造詞，幾乎必然無撞名。 |
| 12 | `grimoire` | **Grimoire｜工程魔導書** | 魔法師的咒語書：workflow 是咒語、validator 是封印、failure pattern 是被記下來的反噬。對「給 agent 讀的書」這件事最有記憶點的玩法。 | 中。同名散落多個小專案與 GPT 應用；無單一壟斷者，但辨識度會被稀釋。 |

## 收斂建議（非決定）

- **最穩三選**：`cairn`（短、好念、意象直中核心）、`plumbline`（驗證閉環 + 中文名極佳）、
  `phronesis`（哲學上最準，但發音門檻）。
- **顯示名稱可雙層**：短 slug 當品牌（如 Cairn），完整定位語保留為副標
  （Cairn — AI-native Cognitive Execution System），README 首行即現行格式，遷移成本低。
- 定名後的 slug 遷移（GitHub redirect、CLI binary 名、`ai-skill` 命令前綴、文件內硬編碼路徑）
  是獨立工程，應另開 plan；本檔不展開。
