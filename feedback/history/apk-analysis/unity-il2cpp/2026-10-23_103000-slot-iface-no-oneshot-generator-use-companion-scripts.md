> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-23 - Slot interface-analysis: no one-shot generator; use companion scripts

Status: validated

#### Reuse Evidence

同一日 seated §3e 交件後，使用者指出包「不太一樣」並問 onboarding 是否有生產器。對照業務 onboarding §3e + `scripts/slots/README.md`：確認**沒有**整包 HTML 生成器；既有配套為 atlas／resolve／merge／thumbs。補契約欄位（`inheritsShared`、`sharedChrome`→`.json`）並把腳本表寫進 onboarding 與本 workflow 後，可重複套用到下一台 §3e。

#### One-line Summary

§3e 沒有一鍵 `generate_interface_analysis`；HTML 手寫對齊參考包，素材線必須跑 `export_sprite_atlas_manifest` → `resolve_ui_sprites` → `merge_sprite_resolution` → `build_sprite_display_thumbs`，且 entrance map 要有 `inheritsShared` + `.json` sharedChrome。

#### Human Explanation

Agent 容易把 §3d 的 `generate_symbol_map.py` 心智套到 §3e，或只交 UnityPy／薄 HTML，結果檔案集合與 bison／buffalo-rush「看起來不一樣」：缺 Pass-5 產物、缺 `inheritsShared`、`sharedChrome` 指到 `.md`。真相是契約本來就要求手寫 HTML + 明確配套腳本；漏跑腳本或漏契約欄才是問題。

#### Trigger

- 使用者問「onboarding 上面沒有生產器腳本嗎／怎麼不太一樣」。
- 新建或補 `interface-analysis/` 時只寫 HTML／inventory，未查 `scripts/slots/`。
- 把 `generate_symbol_map.py` 當 §3e 工具。

#### Evidence

- Tool: 對照業務 onboarding §3e、`scripts/slots/README.md` UI restoration 段、參考 `interface-analysis/` 目錄差集。
- Sanitized excerpt: no one-shot §3e HTML generator；companion scripts listed；contract fields `inheritsShared` + `sharedChrome` JSON。
- Evidence path: `<PROJECT_ROOT>/docs/slots/onboarding.md` §3e「配套脚本」；`<PROJECT_ROOT>/scripts/slots/README.md`；本 workflow 同節。

#### Generalized Lesson

1. **先讀腳本索引再 scaffold**：§3e ≠ §3d 生成器模型。
2. **Atlas 幾乎必跑**；resolve／merge／thumbs 依賴 splash 仍開的 Frida 捕獲——deferred 時寫清，不要默默省略。
3. **契約硬欄**：`inheritsShared` + `sharedChrome`→`.json`（可加 `sharedChromeDoc`）。
4. **深度對齊參考包**：薄骨架可抄薄示例；完整 UI 還原深度對齊完整參考實現 + Pass-5 產物。

#### Agent Action

開 §3e → 讀 onboarding「配套脚本」+ `scripts/slots/README` → 跑／標記 companion scripts → 填 `inheritsShared` → HTML 對齊參考深度 → 分段 commit。下次不要再「手搓完就交」。

#### Goal / Action / Validation

- Goal: §3e 交付與參考包可比，且腳本／契約不漏。
- Action: 本 lesson + onboarding 腳本表 + workflow 節；補缺漏欄位。
- Validation: entrance map 有 `inheritsShared` 與 `.json` sharedChrome；README／resource-map 對未跑腳本有 deferred 說明；workflow 有 companion scripts 節。

#### Applies When

- Unity slot／類似 UnityCache UI 還原，業務有 `interface-analysis/` + `scripts/slots/` atlas 管線。

#### Does Not Apply When

- 純協議／worker 對獎、不做 UI 還原；或專案明確無 atlas／Frida 管線。

#### Validation

新 §3e PR／commit：檔案清單應能對上「已跑／deferred」腳本表；不得只有手寫 JSON 卻宣稱 UI 還原完成。

#### Promotion Target

- [`analysis/apk/workflows/unity-asset-parse-flow.md`](../../../analysis/apk/workflows/unity-asset-parse-flow.md)（§3e companion scripts）
- 業務 slot onboarding §3e（腳本表；非 Ai-skill 路徑）

#### Promotion Record

2026-10-23：已寫入 `unity-asset-parse-flow.md`「§3e companion scripts」；業務 onboarding §3e 新增「配套脚本」表。

#### Required Linked Updates

- Updated: 本 lesson；`unity-asset-parse-flow.md`；`feedback/history/apk-analysis/unity-il2cpp/README.md`（若有索引列）
- Updated (business repo): `docs/slots/onboarding.md` §3e；目標機台 entrance map／README
- Not applicable: enforcement weight／failure-pattern（一次契約遺漏，已用 lesson + onboarding 表吸收）
