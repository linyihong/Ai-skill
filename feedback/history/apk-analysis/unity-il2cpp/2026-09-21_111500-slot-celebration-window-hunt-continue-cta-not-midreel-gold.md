> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-21 - Hunt celebration windows with Continue CTA; mid-reel gold is often Wild coins

Status: candidate

#### One-line Summary

獵取 BIG／HUGE／SUPER 等 L6 慶祝窗時，以底部綠色 Continue／继续 CTA（可加廣告倍率鈕）當主偵測器；旋轉中轉輪上的大片金色像素多半是 Wild 金幣，不能當慶祝窗命中。

#### Human Explanation

自動旋轉截圖時，用「畫面中央金色面積」當分數會大量誤報：Wild 金幣、倍率幣、線獎籌碼浮標都會拉高 gold score。真正的慶祝窗通常伴隨：

- 底部 Continue／继续（常倒數）與可選「乘以 N 倍」廣告 CTA
- 中央動物群／單獸特寫＋金字級（BIG／HUGE／SUPER…）
- 金額條需去敏後才能入庫

字級構圖常遞增：單獸 → 三獸 → 四獸＋額外光條／portal（具體動物依 cabinet theme）。

線獎（L5）則是折線＋符號 Spine 演出＋可選籌碼浮標，**不一定**開 L6 窗。

#### Trigger

- gold-pixel scorer 在旋轉中狂報 candidate，但幀裡只有 Wild 金幣。
- 截到 Continue 後立刻 dismiss，沒有 hold 多幀，錯過字級／動物構圖。
- 把線獎籌碼浮標誤標成 classic payout 模態窗。

#### Evidence

- Tool: device screencap burst；CTA 色帶偵測；可選 runtime Animator 名（idle dump 常有 window controller，playing dump 可能為 0）。
- Sanitized pattern: Continue CTA 與 mid-reel Wild gold 可分離；多字級窗共用同一 window atlas／controller 家族。
- Evidence path: `<PROJECT_ROOT>` cabinet win-FX live notes（project-local）。

#### Generalized Lesson

1. **主偵測器用 Continue CTA**，金色面積只作輔助。
2. **命中後 hold 數幀再 dismiss**，用來分辨 BIG／HUGE／SUPER 構圖。
3. **線獎 ≠ 慶祝窗**：折線＋籌碼浮標可單獨入庫；classic payout *modal* 需另 state-match。
4. **入庫前遮罩**：金額、餘額、下注、廣告倍率數字。
5. **Viewer HTML 保持 technique-free**：偵測腳本與 hook 名放 SOP／JSON，不寫進畫面說明頁。

#### Agent Action

寫慶祝窗獵取腳本時先實作 Continue 偵測與 hold；gold scorer 必須排除或降權「旋轉中／Stop 鈕仍顯示」的幀。

#### Goal / Action / Validation

- Goal: 高召回慶祝窗、低 Wild 金幣誤報。
- Action: CTA 偵測 → hold → 字級分類 → 去敏入庫。
- Validation: candidate 幀可見 Continue；文件標明誤報來源（Wild gold）。

#### Applies / Does Not Apply

- Applies: 有 Continue／继续 慶祝窗的 Unity slot cabinets。
- Does not apply: 無 CTA、純自動跳過的微慶祝；非 slot UI。

#### Related

- `2026-09-24_135500-slot-continue-cta-scorer-needs-matte-green-and-left-band.md`
- `2026-09-21_111200-slot-audio-fourth-stack-and-win-tier-sfx-names.md`
- `2026-09-21_091600-slot-three-visual-planes-and-in-cabinet-anim-stack.md`

#### Promotion Target

- apk-analysis slot capture SOP：慶祝窗獵取偵測器選擇。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
