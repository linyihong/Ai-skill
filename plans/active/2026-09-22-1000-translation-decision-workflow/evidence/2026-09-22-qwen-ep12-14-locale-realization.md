# Run — Qwen ep12–14 locale realization dogfood

**Run ID**: `2026-09-22-qwen-ep12-14`  
**Status**: observed（契約回饋；非 provider 評分）  
**Date**: 2026-09-22  
**Consumer**: narrative-video locale packs（en／ja／id）via pre-translate cache  
**Related fixtures**:  
[`../../workflow/translation/examples/address-title-chen-xiaojie-id.yaml`](../../../workflow/translation/examples/address-title-chen-xiaojie-id.yaml)、  
[`../../workflow/translation/examples/address-title-chen-xiaojie-ja.yaml`](../../../workflow/translation/examples/address-title-chen-xiaojie-ja.yaml)、  
[`../07-target-locale-realization.md`](../07-target-locale-realization.md)、  
[`../08-static-walkthrough-pass.md`](../08-static-walkthrough-pass.md)

## Scope

整集預譯覆蓋率（session 回報；數字為作業狀態，非 quality score）：

| 語種 | ep12 | ep13 | ep14 |
| --- | --- | --- | --- |
| en | 51/51 | 49/49 | 35/36（同源 skip 未寫 cache） |
| ja | 51/51 | 49/49 | 36/36 |
| id | 51/51 | 48/49 | 36/36 |

作業備註（去敏）：日文曾大量為空，因 strip 邏輯誤把帶漢字的日文當「中文残留」刪除；已修並補滿。**此屬 consumer pipeline bug，不是 Translation Decision 契約失敗。**

## 称谓对照（契約相關）

| source | en | ja | id |
| --- | --- | --- | --- |
| 陈小姐 | Ms. Chen | Chenさん | Nona Chen |
| 兄弟们 | Brothers | みんな | Saudara-saudara |
| 我先走了 | I'm out | 先に参ります | Saya harus pergi duluan |
| 我不是发烧我是被下药 | I'm not sick, I was drugged | 私は熱を出したわけじゃない、薬を飲まされたんだ | Saya tidak demam, saya diracun |

## 契約判讀

### 陈小姐

| Locale | Observed | Judgment |
| --- | --- | --- |
| en-US-ish | Ms. Chen | Title + Latin name — structure OK under en |
| id-ID | Nona Chen | Locale-aware title — aligns P0 id fixture |
| ja-JP | Chenさん | **Title realization OK**（さん）；**name realization incomplete** vs preferred katakana（チェン）→ `name_realization=review`（I11） |

**Not**: generic mistranslation.  
**Not**: only a Qwen prompt fix.  
**Is**: evidence that Target-Locale Realization must stay in Translation Decision.

### Other rows（observation only）

- `兄弟们` → ja `みんな`：register／audience adaptation；可日後當 slang／address Candidate Space 案例，本 run **不**升格新 invariant。
- 長句三語：literal／semantic path 粗看合理；無新 blocking contract gap。

## Feedback to SoT（已落地對照）

| Gap | Contract landing |
| --- | --- |
| Latin+honorific ≠ complete ja name | I11 + `validation.name_realization` |
| id English title residue | I5／I8 + `target_locale_residue` review |
| Locales authoritative from pack | I1／I2 + subtitle adapter |
| Surname seeds ≠ final answer | `knowledge/.../ja-JP/name-realization.yaml`（陈 only） |

## Acceptance for this run

- [x] Observed outputs mapped to I1–I11 without inventing quality scores  
- [x] Chenさん classified as name_realization incomplete, not locale_consistency residue  
- [x] No provider／prompt text committed  
- [ ] Live TDR files per cue in consumer project（out of Ai-skill scope this run）

## Next

- Consumer：對 `Chenさん` 類 cue 標 `name_realization=review` 或重跑 Selection 偏好片假  
- Ai-skill：title-mapping 最小種子（本 Phase 3 同批）；完整姓氏庫仍不做  
- Phase 4：document／UI adapter 仍 optional  
