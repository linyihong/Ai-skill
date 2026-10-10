# JavaScript / Electron 靜態發行物分析

對授權的 ASAR 或已解出應用樹：模組圖、imports、source map、路由、IPC／preload 邊界、native addon 宣告關係。

## 與 web scraping 分界

| 本域 | [`../web/`](../web/README.md) |
| --- | --- |
| 問「發行物裡功能怎麼接」 | 問「網站怎麼抓資料」 |
| 輸入是本地 app 樹／ASAR | 輸入是 URL／瀏覽器頁 |

## 步驟

1. 授權 + digest（樹可對關鍵檔做 digest 集合）。
2. 建應用 graph（模組、邊界、IPC 字面量）。
3. 自 seed（字串／路由／IPC channel／export）追 feature。
4. 版本對照時：完整覆蓋才可稱 added／removed；minify 名不作穩定身分。
5. 靜態所見 ≠ runtime 註冊；動態 channel 保持 unknown。

Credentials／cookie／Authorization header **不**寫入 reusable docs。

← [desktop/](README.md)
