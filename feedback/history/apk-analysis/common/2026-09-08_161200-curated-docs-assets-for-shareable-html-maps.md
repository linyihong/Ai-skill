> 遵守 [共用規則索引](../../../enforcement/README.md)、[dependency-reading](../../../enforcement/dependency-reading.md)、[neutral-language](../../../enforcement/neutral-language.md)、[goal-action-validation](../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-08 - Curated docs assets for shareable HTML maps

Status: validated

#### One-line Summary

Shareable visual maps belong next to a **curated** `docs/assets/` copy with relative `src`. Full UnityCache／APK dumps stay gitignored.

#### Human Explanation

Analysis dumps are large and third-party IP, so they are gitignored. Putting an HTML sheet beside that dump (or using `file://` absolute paths) works on one machine and breaks for everyone else. The fix is not to un-ignore the dump: copy **only** the files the page references into a tracked docs folder and keep relative paths.

#### Trigger

- A visual map (HTML or markdown) needs extracted sprites／sheets to remain viewable after clone.
- The only copies live under a gitignored dump directory.
- Reviewers open the HTML and images 404.

#### Evidence

- Tool: gitignore vs tracked `docs/` layout.
- Sanitized excerpt: a feature map HTML sat next to a gitignored cache export; clone lost every `src`.
- Evidence path: target `docs/` under `<PROJECT_ROOT>`

#### Generalized Lesson

1. Keep bulk dumps gitignored (`slot/`, `exported-art/`, UnityCache pulls).
2. If a document must show images, copy the referenced files only into `docs/assets/<topic>/`.
3. HTML `src` is relative to that HTML (e.g. `assets/<topic>/file.png`). No `file://` and no absolute machine paths.
4. Document the exception in the dump README so later agents do not commit the whole export.

#### Agent Action

When asked to make a dump-backed map shareable: copy the referenced subset into tracked docs, retarget `src`, update the dump policy note. Do not lift `.gitignore` for the whole dump.

#### Goal / Action / Validation

- Goal: clone shows the same images as the author.
- Action: curated copy + relative paths + dump stays ignored.
- Validation: `git check-ignore` on dump files still matches; `git ls-files` lists the HTML and only the referenced assets; opening the HTML without a server still loads images.

#### Applies When

- Authorized APK／Unity analysis repos with gitignored art dumps.
- Owner asked for a shareable visual map.

#### Does Not Apply When

- Raw pcap, Frida logs, keys, or full AssetBundle blobs.
- Redistributing publisher art outside the analysis repo.

#### Validation

HTML and listed PNGs are tracked; dump directory remains ignored; relative `src` has no host path.

#### Promotion Target

- `analysis/apk/tools-and-failures.md`

#### Required Linked Updates

- 必須：tools-and-failures Unity 表加精選 assets 列，失敗表加「clone 圖裂」
- 必須：`feedback/history/apk-analysis/README.md` Recent
- 專案 dump 路徑與檔名留 `<PROJECT_ROOT>`
