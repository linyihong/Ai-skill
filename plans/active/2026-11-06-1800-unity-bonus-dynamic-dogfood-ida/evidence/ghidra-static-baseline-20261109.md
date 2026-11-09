# Ghidra 靜態基線

2026-11-09，Ghidra 12.2 與 Temurin 25.0.5 已在 Apple Silicon 啟用。使用官方發行包的 buildNatives 編譯本機反編譯器，成功匯入 ARM64 ELF 並完成四個函式的批次匯出。固定輸入、工具與範圍的兩次唯讀匯出逐位元一致。

以 metadata 與 LLVM 核對完成派發的欄位讀取、間接跳轉及位置收尾的狀態清除。反編譯結果仍有跨函式控制流警告，未將該段推論直接當成事件順序。具體 RVA、輸出指紋、重跑方法與兩假說預測留在 `<PROJECT_ROOT>/docs/slots/el-dorado-dices/lab/bonus-ghidra-static-baseline-20261109.md`。

靜態環境與樣本重跑通過。Runtime 位址、實際自然事件次序與動態回放尚待驗證，計畫維持 draft。
