> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-21 - Slot in-cabinet audio is a fourth stack; win-tier SFX often share BHS letter names

Status: candidate

#### One-line Summary

進桌重建除了 Spine／Animator／Particle，還必須盤點 `AudioClip`＋`sound_store`；慶祝窗字級（BIG／HUGE／SUPER）常與同名 SFX 一一對齊，不要另發明「吼叫」音檔名。

#### Human Explanation

視覺層分析完後很容易漏音效：bundle 裡已有數十個 `AudioClip`，runtime 也有 `*_sound_store`／line sound module，但 inventory 若只列 Texture／Sprite／Spine，重做時會缺床聲、線獎聲與慶祝窗聲。

常見慶祝窗路徑會同時出現：

1. L6 金字級視窗（BIG／HUGE／SUPER…）
2. 同字級的 `*_big_win`／`*_huge_win`／`*_super_win` AudioClip
3. 符號中獎另走 `reward_high_*`／`reward_low`／`reward_wild` 等床聲

動物張嘴是 Spine／合成圖層；音效名稱通常不叫 `roar`／動物種名。

#### Trigger

- 只做畫面 layer-stack，宣稱「動畫齊了」但未列 AudioClip 計數或 sound_store。
- 看到動物張嘴慶祝窗，去找 `cougar_roar` 這類檔名。
- 把 UI button SFX 與 feature-pack 慶祝 SFX 混成同一清單且不分 prefix。

#### Evidence

- Tool: UnityCache／Addressable object type counts；UnityPy `AudioClip` name＋duration metadata（不強制 commit 解碼 wav）。
- Sanitized pattern: feature pack 內同時有 `AudioClip` 計數、`*_sound_store` runtime 物件、以及與慶祝窗同字的 `*_big_win`／`*_huge_win`／`*_super_win`（若該 cabinet 有多字級窗）。
- Evidence path: `<PROJECT_ROOT>` cabinet `audio.md`／`interface-analysis/audio-inventory.json`（project-local）。

#### Generalized Lesson

1. **動畫四棧**：Spine（符號）＋ Animator（狀態／窗）＋ Particle（瞬間 FX）＋ **Audio（SFX／BGM bed）**。缺音效棧＝重建不完整。
2. **先列名再聽**：inventory 以 clip 名＋時長為主；解碼 sample 可選，預設不要把大型 wav 推進文件庫。
3. **字級對齊優先**：若 live 窗顯示 BIG／HUGE／SUPER，先用同字 AudioClip 名對上，再談動物圖層。
4. **符號聲 ≠ 慶祝窗聲**：`reward_*` 跟窗級 `*_big_win` 是不同事件；線聲模組（`lines_*`／combination／win_coins）又是另一路。
5. **Shared UI SFX**：`button_*`／`status_bar_*` 常缺 feature prefix，可能來自 common pack。

#### Agent Action

開進桌 UI／動畫任務時，在 Spine／Animator／Particle checklist 旁加 AudioClip 計數與 sound_store 名；慶祝窗文件必須寫 visual tier ↔ audio name 對照表（或明確 not-applicable）。

#### Goal / Action / Validation

- Goal: 音效納入進桌重建範圍，並與慶祝窗字級對齊。
- Action: 讀 bundle AudioClip 名／時長 → 分組（tier／reward／feature／reel／UI）→ 寫進 cabinet docs。
- Validation: inventory 含 AudioClip 計數；慶祝窗文件有 audio 對照或標 missing。

#### Applies / Does Not Apply

- Applies: Unity IL2CPP slot feature packs 含 AudioClip＋sound_store。
- Does not apply: 純靜態示意頁；無獨立 AudioClip、只靠系統預設音的極簡 client。

#### Related

- `2026-09-21_091600-slot-three-visual-planes-and-in-cabinet-anim-stack.md`
- `2026-09-21_111500-slot-celebration-window-hunt-continue-cta-not-midreel-gold.md`

#### Promotion Target

- `intelligence` / slot UI reconstruction checklist（若有）：補 Audio 為第四棧。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。

#### Validation

依本條既有 Evidence 重做對應觀察或檢查；結果與原始結論一致才可採用。

#### Required Linked Updates

- Not applicable: 這是歷史 closure 修復；未新增可安全提升的 canonical guidance，category index 已存在。
