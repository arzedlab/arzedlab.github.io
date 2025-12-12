---
title: Android Cheatsheet
description: Android cheatsheet for pentesting
date: 2025-01-17T11:13:32+05:00
draft: false 
tags: [android, frida, ssl-pinning,apk] 
toc: true 
---

### Android root in Genymotion

```bash
adb shell setprop persist.sys.root_access 3
```

### **Setting up your Android device**

```bash
$ adb shell getprop ro.product.cpu.abilist # check your device cpu type

$ unxz frida-server.xz
```

```bash
$ adb root # might be required
$ adb push frida-server /data/local/tmp/
$ adb shell "chmod 755 /data/local/tmp/frida-server"
$ adb shell "/data/local/tmp/frida-server &"
```

```bash
$ frida-ps -U
```

```bash
frida -U -l multi-bypass.js -f uz.paynet.flagship_mobile
```

### Downloading And Merging APKs

```java
adb shell pm list packages -3 | grep telegram
```

```bash
$ adb
```

One command

```bash
$ adb shell pm path org.telegram.messenger | sed 's/package://g' | xargs -L 1 adb pull
```

Merge https://github.com/REAndroid/APKEditor 

```bash
$ java -jar APKEditor.jar m -i apk_files
```

### Sign Apks https://github.com/patrickfav/uber-apk-signer

```
$ java -jar uber-apk-signer-1.3.0.jar --apk release.RE.apk
```

### Bypassing Android SSL Pinning Flutter

reFlutter: [https://ayoubnajim.medium.com/bypass-ssl-pinning-for-flutter-apps-using-reflutter-framework-f77b858919b7](https://ayoubnajim.medium.com/bypass-ssl-pinning-for-flutter-apps-using-reflutter-framework-f77b858919b7)

```bash
$ pip3 install reflutter
```

```bash
$ reflutter apk_name.apk
```

### Bypassing Android SSL Pinning Java/Kotlin https://github.com/ilya-kozyr/android-ssl-pinning-bypass

```bash
$ python3 apk-rebuild.py input.apk
```

### Bypassing ssl pinning with Ghidra

First we find offset by searching the strings `ssl_client` and `ssl_server` 

Then run this script

```bash
var lib_loaded = 0;
var do_dlopen = null;
var call_constructor = null;

var linker = Process.findModuleByName("linker64");
if (linker === null) {
    console.error("Module 'linker64' not found!");
} else {
    linker.enumerateSymbols().forEach(function(symbol) {
        if (symbol.name.indexOf("do_dlopen") >= 0) {
            do_dlopen = symbol.address;
        }
        if (symbol.name.indexOf("call_constructor") >= 0) {
            call_constructor = symbol.address;
        }
    });
}

if (do_dlopen === null) {
    console.error("Symbol 'do_dlopen' not found!");
} else {
    Interceptor.attach(do_dlopen, {
        onEnter: function(args) {
            // Try to get the library path from context.x0 or fallback to args[0]
            var libPath;
            if (this.context && this.context.x0 !== undefined) {
                libPath = this.context.x0;
            } else if (args[0] !== undefined) {
                libPath = args[0];
            } else {
                console.error("Unable to determine library path pointer.");
                return;
            }
            // Ensure libPath is valid
            if (libPath.isNull()) {
                console.error("Library path pointer is null!");
                return;
            }
            var library_path = libPath.readCString();
            if (library_path.indexOf("libflutter.so") >= 0) {
                console.log(`[+] Detected loading of ${library_path}`);
                if (call_constructor !== null) {
                    Interceptor.attach(call_constructor, {
                        onEnter: function() {
                            if (lib_loaded === 0) {
                                lib_loaded = 1;
                                var module = Process.findModuleByName("libflutter.so");
                                if (module) {
                                    console.log(`[+] libflutter is loaded at ${module.base}`);
                                    // Adjust offset as needed
                                    session_verify_cert_chain(module.base.add(0x79af3e));
                                } else {
                                    console.error("libflutter.so module not found!");
                                }
                            }
                        }
                    });
                } else {
                    console.error("Symbol 'call_constructor' not found!");
                }
            }
        }
    });
}

function session_verify_cert_chain(address) {
    if (!address || address.isNull()) {
        console.error("Invalid address for session_verify_cert_chain");
        return;
    }
    Interceptor.attach(address, {
        onLeave: function(retval) {
            retval.replace(0x1);
            console.log(`[+] session_verify_cert_chain retval replaced with: ${retval}`);
        }
    });
}

```
