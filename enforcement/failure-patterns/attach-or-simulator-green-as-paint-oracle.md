# Attach / Simulator Green as Paint Oracle（附加成功或模擬器綠燈當成畫面成功）

Status: validated  
Class: `validation-gap` / `false-green`

## Trigger

Agent 或測試在下列任一條件下宣稱「影片可看 / 直屏修復完成」：

- iOS Simulator（或無 HEVC 硬解環境）`readyState` / duration / playlist attach 成功
- `video.currentTime` 前進但未檢查 `videoWidth` / `videoHeight`
- 為通過 Simulator，把 `canPlayType("…mpegurl") === "maybe"` 提升為 native HLS preference
- 只用 headless Chrome stylesheet 規則存在，證明 iOS 原生控制列已隱藏

## Failure Mode

1. **Attach ≠ paint**：音訊軌與時間可前進，視訊層 `videoWidth=0`，使用者看到封面或黑屏。
2. **Simulator false green**：模擬器對 HEVC 常無法硬解；同一路徑在實體裝置上的行為不同。
3. **Capability mismatch**：實體 Safari `"maybe"` + `ManagedMediaSource` 應走 hls.js/`blob`；誤走 native CDN src 後與 known-good production 分叉。
4. **Chrome hide false green**：iOS 常忽略 `::-webkit-media-controls { display:none }`；屬性仍在時原生 ±10 仍可見。

## Risk

- 把未驗證的修復推到測試機 / 正式機，使用者仍無法觀看
- 為修 Simulator 引入真機回歸
- 反覆「部署 → 真機失敗 → 再猜」循環

## Required Agent Action

1. 用**同一台實體裝置**對照 known-good production 與 candidate build。
2. Paint oracle 至少其一：`videoWidth>0 && videoHeight>0`、或可信的中心區 paint 截圖分析（且非僅封面）。
3. Native HLS preference 預設僅 `"probably"`；`"maybe"` + ManagedMediaSource 對齊 production 的 hls.js 路徑。
4. 隱藏原生控制：在確認 custom-player / hls.js 路徑後 strip `controls`；不要把 CSS hide 當唯一手段。
5. Simulator attach 測試可保留作負向門禁，但不得單獨關閉 user-visible paint claim。

## Prevention Gate

| 層級 | Gate |
| --- | --- |
| Capability branch | Document + contract：native only on `"probably"` |
| Physical paint | Remote automation / Web Inspector：`videoWidth>0` |
| Chrome hide | Assert `controls` stripped + marker on custom-player path |
| Deploy close | Physical A/B or stamped evidence before promote |

## Related

- Lesson: [`feedback/history/development-guidance/common/2026-07-17_171500-mpegurl-maybe-managed-mediasource-prefer-hlsjs.md`](../../feedback/history/development-guidance/common/2026-07-17_171500-mpegurl-maybe-managed-mediasource-prefer-hlsjs.md)
- Sibling: [`ai-codegen-passes-ci-fails-production.md`](ai-codegen-passes-ci-fails-production.md)（同屬 false-green / downstream 驗證不足家族）
