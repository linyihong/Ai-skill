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
