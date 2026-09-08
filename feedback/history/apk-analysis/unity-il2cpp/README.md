# unity-il2cpp feedback

Unity IL2CPP / AssetBundle / UnityCache 相關 lesson。

| File | Status | Summary |
| --- | --- | --- |
| [2026-09-07_143800-unity-feature-art-lives-in-unitycache-not-apk.md](2026-09-07_143800-unity-feature-art-lives-in-unitycache-not-apk.md) | validated | 功能美術常在 UnityCache `__data`，不在 APK |
| [2026-09-07_144800-unity-game-session-may-be-custom-tcp-not-tls443.md](2026-09-07_144800-unity-game-session-may-be-custom-tcp-not-tls443.md) | validated | 業務 session 常是非 443 自訂 TCP，不是只 MITM HTTPS |
| [2026-09-07_163400-il2cpp-sendpacket-arg-class-is-the-opcode.md](2026-09-07_163400-il2cpp-sendpacket-arg-class-is-the-opcode.md) | validated | `SendPacket` 記 packet class 名，不要 dump EncryptData 字串 |
| [2026-09-07_170500-parseresult-arg-class-can-be-xdocument-when-c2s-is-not-xml.md](2026-09-07_170500-parseresult-arg-class-can-be-xdocument-when-c2s-is-not-xml.md) | validated | ParseResult arg1 可為 XDocument，與 C2S first-char 無關 |
| [2026-09-07_172200-slot-s2c-may-apply-via-processspin-gameaction-not-parseresult.md](2026-09-07_172200-slot-s2c-may-apply-via-processspin-gameaction-not-parseresult.md) | validated | 旋轉 S2C 可能走 ProcessSpin / GameAction[]，不是下注封包的 ParseResult |
| [2026-09-07_175500-leave-live-table-may-send-offers-then-leave-then-menu-layout.md](2026-09-07_175500-leave-live-table-may-send-offers-then-leave-then-menu-layout.md) | validated | 離桌可能先 offers、再 leave、再 menu-layout |
| [2026-09-08_083000-same-packet-class-encryptdata-length-can-vary.md](2026-09-08_083000-same-packet-class-encryptdata-length-can-vary.md) | validated | 同一 packet class 的 EncryptData 長度可隨桌／下注／featured vs grid 漂移 |
| [2026-09-08_094000-inspect-live-cryptor-object-before-key-recovery.md](2026-09-08_094000-inspect-live-cryptor-object-before-key-recovery.md) | validated | 先解析 interface field 的 live cryptor class 與 pre-encrypt packet shape |
| [2026-09-08_112200-live-cipher-mode-from-backing-enum-not-string-table.md](2026-09-08_112200-live-cipher-mode-from-backing-enum-not-string-table.md) | validated | Live mode／padding 讀 algorithm backing enum 名，不是字串表 |
| [2026-09-08_112201-client-il2cpp-may-embed-unused-server-factory-types.md](2026-09-08_112201-client-il2cpp-may-embed-unused-server-factory-types.md) | validated | Client 內嵌的 Server.* factory 可能從未被 live 呼叫 |
| [2026-09-08_130500-xmlreader-create-first-arg-not-nearby-stream-ctors.md](2026-09-08_130500-xmlreader-create-first-arg-not-nearby-stream-ctors.md) | validated | XmlReader.Create 第一參數類型勝過同窗口 MemoryStream／GetString |
| [2026-09-08_130800-same-thread-dt-window-after-decrypt-not-global-live.md](2026-09-08_130800-same-thread-dt-window-after-decrypt-not-global-live.md) | validated | Decrypt 後同 thread 短 dt 才算 hop，不是全域 LIVE |
| [2026-09-08_131500-il2cpp-object-identity-uses-return-register-not-interceptor-retval.md](2026-09-08_131500-il2cpp-object-identity-uses-return-register-not-interceptor-retval.md) | validated | IL2CPP 物件 identity 用 return register，不是 Interceptor retval |
| [2026-09-08_134800-xname-get-string-may-not-be-decrypt-getstring.md](2026-09-08_134800-xname-get-string-may-not-be-decrypt-getstring.md) | validated | XName.Get 的 String 不必是 Decrypt GetString 物件 |
| [2026-09-08_134500-ctor-xname-object-may-not-equal-load-element-get-name.md](2026-09-08_134500-ctor-xname-object-may-not-equal-load-element-get-name.md) | validated | ctor XName 物件不必等於 Load Element/get_Name |
| [2026-09-08_134200-skipnotify-child-xelement-may-be-another-ctor-not-clonenode.md](2026-09-08_134200-skipnotify-child-xelement-may-be-another-ctor-not-clonenode.md) | validated | SkipNotify 子 XElement 可能是另一個 ctor，不是 CloneNode/Load |
| [2026-09-08_134000-xelement-xname-fill-may-use-skipnotify-not-public-add.md](2026-09-08_134000-xelement-xname-fill-may-use-skipnotify-not-public-add.md) | validated | XElement(XName) 填充可能走 SkipNotify，不是 public Add |
| [2026-09-08_133000-parse-xelement-args-may-be-ctor-copies-not-load-tree.md](2026-09-08_133000-parse-xelement-args-may-be-ctor-copies-not-load-tree.md) | validated | Parse 的 XElement 可能是 ctor 複本，不是 Load 樹節點 |
