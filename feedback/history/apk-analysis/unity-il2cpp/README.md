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
| [2026-09-08_083000-same-packet-class-encryptdata-length-can-vary.md](2026-09-08_083000-same-packet-class-encryptdata-length-can-vary.md) | validated | 同一 packet class 的 EncryptData 長度可隨桌／下注變 |
