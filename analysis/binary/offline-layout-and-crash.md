# 離線 ELF layout 與 recorded crash

不開完整反編譯資料庫、不 live attach 時的取證。

## ELF layout

記錄 sections／segments／symbols／relocations／靜態 mitigation 候選。完整結果可再投影摘要，避免重跑解碼。推論（ASLR／NX 等）標 inference；缺旗標標 unknown。

## Recorded crash（Linux x86-64 core）

讀 thread registers、signals、raw notes。PID 是歷史 metadata，**不**授權對 live 行程操作。可選 debugger context 僅限 core-only；失敗給 setup guidance。

## 授權

core／ELF 必須是已授權標的的產物。不得用「拿到 core」反推可對線上系統除錯。

← [binary/](README.md)
