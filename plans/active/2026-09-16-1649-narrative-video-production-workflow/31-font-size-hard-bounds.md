# Candidate: Font size hard bounds

Companion to [`27-typography-layout-profile.md`](27-typography-layout-profile.md)、[`30-max-lines-is-bound.md`](30-max-lines-is-bound.md)。**workflow 已補硬閘**（[`subtitle-layout.md`](../../../workflow/narrative-video-production/subtitle-layout.md)）。  
觀察：[`evidence/2026-09-24-font-size-hard-bounds.md`](evidence/2026-09-24-font-size-hard-bounds.md)。

Layout 為了 fit 而無限縮字，是 Constraint 沒封下限。`min`／`max`／`step`（或 `allowed` 離散集）由 mechanical engine 強制：`min ≤ size ≤ max`。到 min 仍放不下 → 重切 Speech Unit，禁止 36→32→24。短句不得因畫面空而超過 max。Scene `max_delta_from_scene` 過大視為應切句，不是再縮。AI 不得選 min 以下字級。
