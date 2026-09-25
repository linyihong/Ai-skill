> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-24 - Continue CTA scorers need matte green + left band; gold-alone false+ on reels

Status: candidate

#### One-line Summary

綠色 Continue／继续 色帶偵測勿用「高亮 g>160 + 畫面正中央」：真實 CTA 常是 matte green（約 g=115–140）且偏左；金色面積單獨達標會把轉輪 Wild／光柱誤標成 TOTAL WIN。

#### Human Explanation

慶祝窗視覺獵取若把 Continue 門檻設成高飽和綠、又只掃中下正中央，會在真實 Continue 幀上得到接近 0 的分數（只吃到高光邊緣），同時把 mid-reel 金色誤當 TOTAL WIN。

可分離訊號：

- **Continue**：左中下帶（避開最右 Spin／停止鍵），matte green（g 明顯大於 r/b，但不必極亮）
- **TOTAL WIN／字級窗**：上中金字帶 **且** Continue 已上色
- **Flash**：中藍 spark／標題 **且** Continue
- Idle HUD 左條也可能有微量綠（遠低於真實 CTA）

ADB `screencap -p` 偶發 truncated PNG：讀圖前要 `load()`／重試，否則長跑 hunt 會中斷。

#### Trigger

- 肉眼可見 继续，但 scorer cont≈0.00x。
- `total` 高、`cont` 低卻標成 TOTAL WIN，幀裡是旋轉中 Wild／光柱。
- TOTAL WIN＋Continue 被誤標成 classic_candidate（未硬排除金字帶）。

#### Evidence

- Tool: dense screencap hunt；RGBA pixel band scorers；retry on truncated screencap.
- Sanitized pattern: matte left Continue vs neon highlight；gold band requires Continue for TOTAL WIN tag；classic candidate hard-excludes TOTAL WIN gold.
- Evidence path: project-local visual-hunt JSON（presence-only；amount PNGs not committed）.

#### Generalized Lesson

1. **先用真實 Continue 幀校準色帶**（取樣 CTA 本體 RGB），不要抄高亮門檻。
2. **Continue crop 避開 Spin／停止鍵**（通常在最右）。
3. **金字／藍 spark 標籤必須搭配 Continue**（或等價 CTA），禁止 gold-alone。
4. **互斥窗要硬排除**：TOTAL WIN 金字帶不可進 classic／line_win 候選。
5. **screencap 要 retry**：truncated PNG 會讓長跑腳本崩掉。

#### Agent Action

寫／修慶祝窗視覺 scorer 時：用真實 Continue 幀驗 cont 分離（idle≪Continue）；TOTAL WIN／Flash 標籤加 Continue 閘；classic 候選排除已知金字帶；ADB screencap 加截斷重試。

#### Goal / Action / Validation

- Goal: Continue 高召回、Wild gold 低誤報、窗類互斥正確。
- Action: 校準 matte+left Continue → 標籤加 CTA 閘 → 互斥排除 → screencap retry。
- Validation: idle cont 遠低於 Continue；TOTAL WIN／Flash 幀 cont 高且 tag 正確；無 amount PNG 入 git。

#### Applies / Does Not Apply

- Applies: 有綠色 Continue／继续 的 Unity slot 慶祝／派彩窗視覺獵取。
- Does not apply: 無 CTA 的微慶祝；非綠色 CTA；純 IL2CPP hook 路徑。

#### Related

- `2026-09-21_111500-slot-celebration-window-hunt-continue-cta-not-midreel-gold.md`
- `2026-09-24_114500-slot-base-line-wins-may-bypass-payout-window-show.md`
- `2026-09-24_132000-slot-classic-chip-fx-is-not-classic-payout-window.md`

#### Validation

- 用獨立 idle、Continue、TOTAL WIN／Flash 幀比較 scorer；確認 idle 顯著較低、CTA tag 正確、且 amount 圖像不入庫。

#### Promotion Target

- apk-analysis slot capture SOP：慶祝窗 Continue CTA 色帶校準與互斥標籤。

#### Reuse Evidence

- 尚未有獨立 cabinet／專案的重用證據；色彩與版面閾值特別需要第二個來源才能升格。

#### Promotion Record

- 尚未 promotion；僅保留為 feedback history candidate。

#### Required Linked Updates

- 已更新 `unity-il2cpp/README.md` category index；尚未更新 capture SOP，避免把單一 CTA 配色寫成通用預設。
