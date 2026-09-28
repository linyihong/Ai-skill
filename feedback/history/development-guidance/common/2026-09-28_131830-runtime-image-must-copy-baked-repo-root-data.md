> 遵守 [共用規則索引](../../../../enforcement/README.md)、[dependency-reading](../../../../enforcement/dependency-reading.md)、[neutral-language](../../../../enforcement/neutral-language.md)、[goal-action-validation](../../../../enforcement/goal-action-validation.md) 與 [feedback-lessons](../../../feedback-lessons.md)；本檔只寫本條 lesson，不重複貼上共用政策全文。

### 2026-09-28 - Runtime image must ship data under baked repoRoot

Status: candidate

#### Reuse Evidence

尚未有獨立重用證據，維持 candidate。本條來自一次多階段前端映像部署後「靜態資源 OK、檔案型 API 全 404」的對照。

#### One-line Summary

若 build 把 `repoRoot` 烤成絕對路徑，runtime 映像必須把伺服器仍會 `fs.read` 的資料樹 COPY 到同一路徑，不能只帶編譯產物目錄。

#### Human Explanation

多階段 Docker 常在 build stage 掛完整 repo（含 docs），讓 bundler 把公開資產打進 `.output/public`，但 final stage 只 COPY 編譯輸出。若 runtimeConfig 仍保留 build 時的絕對 `repoRoot`，API 用 `join(repoRoot, 'docs/...')` 讀檔會全部 miss，而靜態 URL 仍 200——表面像「某 cabinet 沒上線」，其實是全庫 docs-store 斷路徑。

#### Trigger

- 健康檢查顯示 `repoRoot` 為 build 路徑（例如 `/src`），且 registry / cabinet 計數為 null 或 API 回 `cabinet not found`。
- 同一部署下公開 slot-assets 正常，但所有檔案型 cabinet / spin fixture API 404。
- 本機（完整 checkout）200，容器（僅 `.output`）全掛。

#### Evidence

- Tool: container inspect + health JSON + path probe under baked `repoRoot`
- Sanitized excerpt: health reported baked `repoRoot` with null registry count; runtime image lacked `docs/` under that root while public assets existed under the Nitro output tree.
- Evidence path: keep project-specific Dockerfile / deploy notes under `<PROJECT_ROOT>` platform docs.

#### Generalized Lesson

1. 對照 health（或等效 runtimeConfig dump）裡的 `repoRoot` 與容器內實際目錄。
2. Final image 必須在 **baked path** 提供仍被 server-side `fs` 讀取的資料（docs、fixture、farm JSON 等），或改為 runtime env override 並掛 volume。
3. 不要用「單一資源 404」當 onboard 失敗結論——先用第二個已知資源 / registry 計數排除全域路徑問題。

#### Agent Action

部署後若檔案型 API 全 404：先查 baked `repoRoot` 與 runtime COPY/mount，再查個別 package。修 Dockerfile 或掛載後重建 web 映像並用兩個 cabinet id 做對照驗證。

#### Goal / Action / Validation

- Goal: 容器內 docs-backed API 與本機相同可讀。
- Action: 在 final stage `COPY` 資料樹到 baked `repoRoot`（或注入可覆寫的 `repoRoot` + volume）。
- Validation or reference source: health registry count non-null；至少兩個 cabinet + 一個 spin fixture API 回 200。

#### Applies When

- 伺服器在 runtime 用絕對 `repoRoot` 讀 repo 相對檔案。
- 多階段映像只複製 bundler 輸出。

#### Does Not Apply When

- 資料已全部嵌進 bundle / DB，不再 fs 讀 repo。
- `repoRoot` 由 runtime env 指向已掛載 volume。

#### Validation

重建映像後：health 計數恢復；先前 404 的檔案型 API 變 200；靜態資產仍可用。

#### Promotion Target

- `workflow/development-guidance/`（deploy / container checklist，若日後整理）

#### Promotion Record

尚未 promotion；保留為 candidate history。

#### Required Linked Updates

N/A — candidate lesson only; no workflow promotion this turn.
