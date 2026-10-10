# Managed / .NET 靜態取證

PE／CLI assembly 的 metadata、CIL、宣告的 native 依賴與建置對照。預設**靜態 only**。

## 步驟

1. 授權 + digest。
2. 清單：assemblies、types、methods、P/Invoke／native deps。
3. 需要時讀 CIL；與另一建置對照時用完整、相容的頁集合。
4. NativeAOT：native 偽碼走 [`native-investigation.md`](native-investigation.md)；metadata 恢復有 host／layout 邊界，缺則 unknown。

## 邊界

靜態所見 ≠ runtime 註冊或政策強制。模糊 native 連結保持 unknown。

← [binary/](README.md)
