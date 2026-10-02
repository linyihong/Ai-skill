> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-10-02 - OCR layout and typography are evidence features, not hard rules

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。來自畫面具大量非對白文字（平台 UI／聊天／禮物）與硬字幕並存的 dogfood 觀察：只靠 `position: bottom` 或絕對字級閾值會誤殺或誤收。

#### One-line Summary

把 OCR 從「文字」升級成「文字＋相對幾何＋字級帶＋時序行為＋區域角色」；`font_size` 只作 likelihood evidence，禁止 `height > X → subtitle` 硬規則。

#### Human Explanation

同一幀常同時有頂欄 UI、左下滾動聊天、禮物條，以及真正硬字幕。它們的幾何／字級／時序行為不同：聊天多行、小字、高替換率；硬字幕常是較大字帶、同文穩定數秒、空間穩。機械層應先形成 region 並標 `region_role`（含 negative evidence），LLM 只對少量代表幀做 region classification／probe advisor，不當 OCR engine。

#### Trigger

- 畫面有密集平台 UI／直播聊天／禮物條，OCR 把留言當對白進 corpus
- 或因害怕誤收而把可能硬字幕整區誤殺
- 討論「要不要加字體大小判斷」時，有人想寫絕對 px／pt 閾值

#### Evidence

- Tool: product OCR／subtitle-candidate dogfood frame review
- Sanitized：同幀 top UI、lower-left multi-line changing chat、center overlay vs 典型硬字幕區的幾何差異
- Evidence path: `<PROJECT_ROOT>` analysis／assets only

#### Generalized Lesson

1. **相對尺寸，不是絕對字級**：使用 `height_ratio = bbox_h / frame_h`、`width_ratio`、`estimated_font_size_band: small|medium|large`；禁止跨解析度／直橫屏硬編碼 `font_size > X`。
2. **多證據聯合**：typography + position + persistence + spatial stability + morphology；單特徵不得定案。
3. **Region first**：OCR boxes → geometry clustering → region behavior → `region_role`（`subtitle_candidate`｜`livestream_chat`｜`platform_ui`｜`watermark`｜`badge`｜`unknown`）→ 再進 subtitle candidate。
4. **Temporal behavior 權重高**：`persistence_s`、`text_stability`、`box_stability`、`replacement_rate`；聊天＝高替換；硬字幕＝同文穩定。
5. **Negative evidence 必留**：為何不是字幕（small_typography、rapidly_changing_text、avatar／badge_present 等）寫進 exclusion_evidence，不當「垃圾資料丟棄」。
6. **LLM 位置**：只看少量代表幀＋mechanical boxes，輸出 region／layout hypothesis；不直接 crop、不宣稱「沒有字幕」。
7. **不開新 Phase**：補進既有 OCR mechanical evidence schema＋probe／escalation／subtitle-candidate 契約。

#### Agent Action

遇到「加字體大小」需求：先擴 schema（geometry／typography／temporal／region_role／negative evidence），用相對比例＋多證據分類；用真實混雜 UI 幀驗證，不為單幀加硬規則；勿開新 Phase。

#### Goal / Action / Validation

- Goal: 混雜 UI／聊天幀仍能區分 subtitle_candidate vs non-subtitle，且可審計。
- Action: 寫入／強化 NVP OCR layout＆typography evidence；產品 schema 對齊；單幀／短窗驗證。
- Validation or reference source: 聊天區多數 `livestream_chat`／non_subtitle；硬字幕帶可為 `subtitle_candidate`／uncertain；doc 含 exclusion_evidence；無絕對 px 閾值。

#### Applies When

- 機械 OCR／subtitle candidate／layout discovery
- 畫面可能含平台 UI、滾動聊天、浮水印與硬字幕

#### Does Not Apply When

- 純排版／burn 的 cue typography（那是 publish layout，不是 OCR evidence）
- 尚未有任何 OCR box（仍屬 acquisition miss）

#### Validation

- 同幀多樣文字：role／status 可區分；negative evidence 可讀。
- 改解析度後相對 ratio 仍可用；絕對字級規則不存在於契約／代碼預設。
- LLM 僅作 region advisor，不單獨決定 OCR crop。

#### Promotion Target

- `workflow/narrative-video-production/text-evidence-subtitle-candidate.md`
- `workflow/narrative-video-production/text-evidence-ocr-discovery.md`

#### Promotion Record

尚未 promotion；本輪同步強化上述兩份 workflow。

#### Required Linked Updates

- 更新 `feedback/history/narrative-video-production/common/README.md` 索引
- 強化 subtitle-candidate（typography／temporal／region_role／negative evidence）
- 在 ocr-discovery 加 Layout＆Typography Evidence 交叉連結
- 具體專案幀證據留 `<PROJECT_ROOT>`；本 lesson 不含專案片名／本機路徑
