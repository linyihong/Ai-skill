# APK Analysis Feedback History

## 分類

| 分類 | 數量 | 說明 |
|------|------|------|
| [`common/`](common/) | 68 | 跨分類或通用 lesson（工具選擇、流程、UI、分析策略等） |
| [`flutter-dart-aot/`](flutter-dart-aot/) | 23 | Flutter/Dart AOT 相關 hook 與分析 |
| [`http-api/`](http-api/) | 34 | HTTP API 分析、文件化、UI 操作流程 |
| [`local-proxy/`](local-proxy/) | 11 | Local proxy 設定、診斷、hook |
| [`media-hls/`](media-hls/) | 3 | Media/HLS 串流分析 |
| [`unity-il2cpp/`](unity-il2cpp/) | 38 | Unity IL2CPP、AssetBundle、UnityCache |
| [`dynamic-capture/`](dynamic-capture/) | 2 | 動態捕獲相關 |

## Recent (2026-09-10)

| Slug | Category |
|------|----------|
| `common/2026-09-10_143500-slot-dual-wild-sum-and-desk-round` | Dual mult-wild sum; desk ROUND without symbols ≠ missing lines |
| `common/2026-09-10_142327-bison-round-jackpot-no-symbols` | ROUND jackpot RESULT may omit Symbols |
| `common/2026-09-10_141507-bison-lineset-frida-dump` | Live GetLineSet dump for payline geometry |
| `common/2026-09-10_042103-payline-geometry-spot-verify` | Payline geometry: spot-verify paths; do not invent full maps |
| `unity-il2cpp/2026-09-10_084100-session-scoped-c2s-command-synthesis-vs-wire-replay` | Session-scoped C2S command synthesis vs wire replay |

## Recent (2026-09-09)

| Slug | Category |
|------|----------|
| `unity-il2cpp/2026-09-09_112800-unity-ui-identity-from-loaded-objects-not-screenshot` | Unity on-screen UI identity from loaded objects, not screenshot matching |
| `common/2026-09-09_100200-wellknown-http-ports-may-not-fail-named-host-resolve` | Blocking well-known HTTP(S) ports may still hit resolve Success |
| `common/2026-09-09_094550-reset-level-may-change-login-packet-class` | force-stop vs clear-data can send different login packet classes |
| `common/2026-09-09_094540-prefs-key-names-on-device-without-values` | Extract prefs `name=` on-device; Magisk `su -c` |
| `common/2026-09-09_094520-radio-toggle-may-not-drop-custom-tcp` | Radio/`svc` may leave custom TCP ESTAB |
| `unity-il2cpp/2026-09-09_094510-cleardata-spawn-delay-thread-attach` | Clear-data spawn needs delayed `thread_attach` |
| `common/2026-09-09_094500-frida-quiet-mode-needs-timeout` | Frida 17 `-q` needs `-t` or the session exits |

## Recent (2026-09-08)

| Slug | Category |
|------|----------|
| `unity-il2cpp/2026-09-08_171200-slot-play-rules-compare-cabinets-without-rng` | Per-cabinet play-rules keys to compare games without RNG |
| `unity-il2cpp/2026-09-08_165800-slot-spin-fixtures-pair-json-manifest-and-reel-crop` | Pair each slot spin's JSON, reels-only crop, and manifest validation |
| `unity-il2cpp/2026-09-08_162200-symbol-dto-id-may-be-one-char-strip-not-resource-name` | Parsed symbol ID may be a one-character strip code |
| `common/2026-09-08_161200-curated-docs-assets-for-shareable-html-maps` | Curated docs assets for shareable HTML maps |
| `unity-il2cpp/2026-09-08_154800-wrapper-extras-may-be-ctor-invoke-clone-not-parseall` | Event wrapper extras may be only ctor/Invoke/Clone, not nested ParseAll |
| `unity-il2cpp/2026-09-08_154000-nested-dto-may-fill-via-ctor-without-parse` | Nested DTO may fill via ctor without Parse* |
| `unity-il2cpp/2026-09-08_153400-nested-dtos-may-have-own-parse-beside-wrapper` | Nested DTOs may have their own Parse/ParseAll beside the wrapper |
| `unity-il2cpp/2026-09-08_152200-nested-dto-ctors-may-sit-before-after-parse-not-apply` | Nested DTO ctors may sit before/after Parse, not in apply |
| `unity-il2cpp/2026-09-08_151000-parse-may-construct-wrapper-not-nested-dtos` | Parse may construct only the event wrapper, not nested DTOs |
| `unity-il2cpp/2026-09-08_150200-parse-may-store-dto-fields-without-setters` | Parse may store DTO fields without calling set_* |
| `unity-il2cpp/2026-09-08_144800-parse-may-convert-xml-via-enum-parse-without-get-value` | Parse may convert XML via Enum.Parse without get_Value |
| `unity-il2cpp/2026-09-08_144000-nested-element-attribute-may-not-share-load-tree-xattribute` | Nested Element Attribute may not share Load-tree XAttribute |
| `unity-il2cpp/2026-09-08_143200-elements-iterator-current-may-be-nested-xelement-ctor` | Elements iterator Current may be nested XElement ctor |
| `unity-il2cpp/2026-09-08_142800-parse-xcontainer-element-retval-may-be-nested-ctor-not-load-tree` | Parse XContainer.Element retval may be nested ctor, not Load tree |
| `unity-il2cpp/2026-09-08_141000-parse-xelement-args-may-navigate-via-xcontainer-element` | Parse XElement args may navigate via XContainer.Element |
| `unity-il2cpp/2026-09-08_140200-parse-wrapper-attribute-may-share-load-tree-xattribute` | Parse wrapper Attribute may share Load-tree XAttribute |
| `unity-il2cpp/2026-09-08_135800-xname-get-substring-this-may-not-be-xmlreader-name` | XName.Get Substring.this may not be XmlReader/XName getters |
| `unity-il2cpp/2026-09-08_135500-xname-get-substring-this-may-not-be-decrypt-getstring` | XName.Get Substring.this may not be Decrypt GetString |
| `unity-il2cpp/2026-09-08_135200-xname-get-string-may-be-substring-not-getstring` | XName.Get string may be String.Substring, not Decrypt GetString |
| `unity-il2cpp/2026-09-08_134800-xname-get-string-may-not-be-decrypt-getstring` | XName.Get string may not be the Decrypt GetString object |
| `unity-il2cpp/2026-09-08_134500-ctor-xname-object-may-not-equal-load-element-get-name` | Ctor XName object may not equal Load Element/get_Name XName |
| `unity-il2cpp/2026-09-08_134200-skipnotify-child-xelement-may-be-another-ctor-not-clonenode` | SkipNotify child XElement may be another ctor, not CloneNode or Load |
| `unity-il2cpp/2026-09-08_134000-xelement-xname-fill-may-use-skipnotify-not-public-add` | XElement(XName) fill may use SkipNotify, not public Add |
| `unity-il2cpp/2026-09-08_133000-parse-xelement-args-may-be-ctor-copies-not-load-tree` | Parse XElement args may be ctor copies, not Load tree |
| `unity-il2cpp/2026-09-08_131500-il2cpp-object-identity-uses-return-register-not-interceptor-retval` | IL2CPP object identity uses return register, not Interceptor retval |
| `unity-il2cpp/2026-09-08_130800-same-thread-dt-window-after-decrypt-not-global-live` | Same-thread dt window after Decrypt, not global LIVE |
| `unity-il2cpp/2026-09-08_130500-xmlreader-create-first-arg-not-nearby-stream-ctors` | XmlReader.Create first-arg class, not nearby stream ctors |
| `unity-il2cpp/2026-09-08_112201-client-il2cpp-may-embed-unused-server-factory-types` | Client IL2CPP may embed unused server-named factory types |
| `unity-il2cpp/2026-09-08_112200-live-cipher-mode-from-backing-enum-not-string-table` | Live cipher mode from backing enum, not string table |
| `unity-il2cpp/2026-09-08_094000-inspect-live-cryptor-object-before-key-recovery` | Inspect live cryptor object before key recovery |
| `unity-il2cpp/2026-09-08_083000-same-packet-class-encryptdata-length-can-vary` | Same packet class EncryptData length can vary |

## Recent (2026-09-07)

| Slug | Category |
|------|----------|
| `common/2026-09-07_173500-frida-attach-android-process-label-not-package` | Frida attach uses launcher label / PID, not package |
| `unity-il2cpp/2026-09-07_175500-leave-live-table-may-send-offers-then-leave-then-menu-layout` | Leave may be offers then leave then menu-layout |
| `unity-il2cpp/2026-09-07_172200-slot-s2c-may-apply-via-processspin-gameaction-not-parseresult` | S2C apply may be ProcessSpin GameAction[] not ParseResult |
| `unity-il2cpp/2026-09-07_170500-parseresult-arg-class-can-be-xdocument-when-c2s-is-not-xml` | ParseResult arg1 class (XDocument) ≠ C2S first-char |
| `unity-il2cpp/2026-09-07_163400-il2cpp-sendpacket-arg-class-is-the-opcode` | SendPacket: log packet class, not EncryptData string |
| `unity-il2cpp/2026-09-07_144800-unity-game-session-may-be-custom-tcp-not-tls443` | Custom TCP game session vs HTTPS CDN |
| `unity-il2cpp/2026-09-07_143800-unity-feature-art-lives-in-unitycache-not-apk` | Feature art in UnityCache `__data`, not APK |

## Recent (2026-09-01)

| Slug | Category |
|------|----------|
| `http-api/2026-09-01_093500-entitlement-grant-is-playable-field-presence-not-client-flag` | Playable field presence vs client lock flag |

## Recent (2026-08-28)

| Slug | Category |
|------|----------|
| `common/2026-08-28_152800-flutter-dio-device-autologin-standalone-sdk-bootstrap` | Device autologin as SDK bootstrap (not refresh-harvest) |
| `common/2026-08-28_151800-tasker-adb-play-focus-suppress-automation-guards` | Tasker ADB import / trial / onboarding guards for Play focus suppress |

## Recent (2026-07-15)

| Slug | Category |
|------|----------|
| `common/2026-07-15_132500-frida-e2e-compact-send-and-off-main-thread-http` | Frida E2E compact send + off-main-thread HTTP |
| `common/2026-07-15_140400-httpglobal-header-setter-hook-for-token-when-okhttp-introspect-fails` | HttpGlobal token hook when OkHttp introspect fails |
| `http-api/2026-07-15_140000-cold-start-token-may-come-from-bootstrap-not-login` | Guest token from bootstrap not login |
| `http-api/2026-07-15_140100-dual-sign-canonical-h5-json-vs-retrofit-empty-bodystr` | Dual sign path A/B |
| `http-api/2026-07-15_140300-bootstrap-body-may-carry-analytics-ids-not-device-gaid` | Bootstrap analytics IDs |
| `http-api/2026-07-15_140500-bootstrap-retrofit-sign-may-include-wire-json-in-bodystr` | Bootstrap bodyStr may be wire JSON (W27b) |
| `http-api/2026-07-15_140600-invalid-sign-json-200-may-mean-missing-initheaders-not-bad-rsa` | JSON Invalid sign → check initHeaders (W27c) |

## Recent (2026-07-14)

| Slug | Category |
|------|----------|
| `common/2026-07-14_171000-pairip-application-licenseactivity-cold-start-play-gate` | Play focus-steal A/B; promoted → analysis workflow + intelligence heuristic |
| `http-api/2026-07-14_141700-crosscheck-request-sign-offline-vs-inapp-fingerprint` | Offline vs in-app sign FP dry-run |
| `http-api/2026-07-14_141500-request-sign-may-be-sha256withrsa-pkcs8-header-concat` | SHA256withRSA + PKCS8 + header concat |
| `common/2026-07-14_135210-prove-wire-json-names-via-jsonreader-nextname` | Wire names via JsonReader.nextName (R8) |
| `common/2026-07-14_135200-sign-canonical-structural-probe-allowlist-keys` | Canonical structure probe：allowlist keys |
| `http-api/2026-07-14_135220-request-sign-interceptor-may-outnumber-signer-calls` | Interceptor count ≠ signer count |
| `http-api/2026-07-14_133500-request-sign-may-live-in-vendored-httpglobal-not-app-intercept` | Sign in vendored HttpGlobal, not app intercept |
| `http-api/2026-07-14_131600-request-sign-fixed-b64-length-may-be-rsa-scale` | sign fixed base64 len ≈ RSA-2048 scale |
| `http-api/2026-07-14_131500-ondisk-encrypted-chapter-blob-length-matches-decrypt-input` | On-disk blob len ≡ decrypt inLen |
| `http-api/2026-07-14_112000-rsa-public-unwrap-to-short-key-then-content-decrypt` | RSA unwrap → short key → content decrypt |
| `http-api/2026-07-14_105600-named-getcontentbody-decrypt-yields-utf8-chapter-plaintext` | getContentBody → UTF-8 chapter plaintext |
| `local-proxy/2026-07-14_105100-column-c-bypass-may-be-okhttp-no-proxy-not-cronet` | Column-C bypass may be OkHttp no-proxy |
| `local-proxy/2026-07-14_095200-on-device-ssl-certificate-ui-is-tls-path-signal` | On-device SSL UI = TLS signal |
| `local-proxy/2026-07-14_095210-mixed-mitm-ads-ok-firstparty-ssl-api-bypass` | Mixed MITM: ads OK / SSL UI / API bypass |
| `common/2026-07-14_095220-apk-analysis-window-closeout-cognitive-and-feedback` | Window close-out: Cognitive + Feedback |

## Recent (2026-06-22)

| Slug | Category |
|------|----------|
| `common/2026-06-22_120000-okhttp-r8-obfuscated-request-realinterceptorchain-proceed` | OkHttp R8 hook overload |
| `local-proxy/2026-06-22_120100-mitm-cdn-visible-primary-api-invisible-pinning-tier` | MITM CDN vs API pinning tier |
| `http-api/2026-06-22_120200-static-ktor-strings-not-dynamic-business-api-client` | Retrofit vs Ktor dynamic client |
| `http-api/2026-06-22_130000-in-session-api-host-differs-from-cold-start-primary` | Cold vs in-session API host failover |
| `common/2026-06-22_130100-r8-obfuscated-okhttp-response-needs-converter-hook` | Obfuscated Response / converter hook |
| `common/2026-06-22_131000-retrofit-gson-fromjson-hook-api-response-plaintext` | Gson.fromJson API JSON capture |
| `http-api/2026-06-22_141500-static-list-endpoint-zero-hit-check-detail-embedded-catalog` | List path 0 hit → check embedded catalog |
| `http-api/2026-06-22_141600-static-waterfall-path-zero-hit-check-shelf-layout-gating` | Waterfall path 0 hit → shelf layout gating |
| `http-api/2026-06-22_141700-custom-request-signatures-block-standalone-sdk-until-re-or-relay` | Custom sign blocks standalone SDK |
| `http-api/2026-06-22_141800-okhttp-interceptor-sorted-map-canonical-sign-pattern` | Interceptor sorted-map canonical for sign |
| `common/2026-06-22_141900-crypto-util-name-sha256-does-not-imply-plain-sha256` | sha256* name ≠ plain SHA256 verify gate |
| `common/2026-06-22_142000-jni-dynamic-registernatives-no-standard-java-com-export` | Dynamic JNI / no Java_com_* export |
| `http-api/2026-06-22_142100-encrypted-request-time-double-native-call-second-value-in-sign-map` | Encrypted-time header double native call |
| `common/2026-06-22_142200-frida-enumerate-loaded-classes-causes-script-load-timeout` | Avoid enumerateLoadedClasses timeout |
| `common/2026-06-22_142300-hmac-sha256-per-mode-keys-in-native-rodata` | Plain SHA256 fail → HMAC + per-mode rodata keys |
| `common/2026-06-22_142400-system-loadlibrary-name-overrides-static-so-assumption` | loadLibrary name vs wrong protection .so |
| `http-api/2026-06-22_142500-partial-offline-sign-hmac-solved-requesttime-session-remain-relay` | Hybrid SDK when sign offline, encrypted-time native |
| `common/2026-06-22_142600-frida-python-attach-java-bridge-use-cli-subprocess-for-rpc` | Python attach Java undefined → Frida CLI RPC |
| `http-api/2026-06-22_142700-opaque-api-blob-native-decrypt-secondary-compression-json` | Opaque blob → decryptStr + zlib → JSON |
| `http-api/2026-06-22_142800-native-encrypted-time-getter-may-cache-first-rotates-consecutive-same` | Encrypted-time getter: 1st rotates, 2nd+ same in burst |
| `http-api/2026-06-22_142900-wire-json-field-names-differ-from-gson-bean-paths` | Wire JSON keys ≠ Gson bean field names |
| `http-api/2026-06-22_143000-encrypt-plaintext-may-be-millis-plus-fingerprint-not-surface-seed` | Encrypt plaintext may be millis + fingerprint, not surface seed |
| `http-api/2026-06-22_143100-apk-signing-cert-sha1-may-be-stable-encrypt-plaintext-suffix` | Signing cert SHA1 as stable colon-hex suffix |
| `http-api/2026-06-22_143200-custom-f3aes-label-may-still-be-standard-aes128-cbc-with-rodata-key-iv` | Custom crypto label may still be standard AES after key hook |
| `http-api/2026-06-22_143300-native-decrypt-mode-may-reuse-encrypt-mode-key-material` | Decrypt mode N may share encrypt mode N keys |
| `http-api/2026-06-22_143400-login-may-bootstrap-with-sentinel-uid-session-headers` | Login may bootstrap with sentinel uid/session |
| `http-api/2026-06-22_143500-paywall-entitlement-may-bind-to-account-session-not-guest-bootstrap` | Paywall grant may need account session, not guest |

## 來源

所有 lesson 原位於 `skills/apk-analysis/feedback_history/`，已於 2026-05-13 搬遷至此，舊目錄已刪除。
