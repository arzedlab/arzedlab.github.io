---
title: "V8 Internals: Tree"
description: "A conceptual overview of V8’s internal structures ..."
date: 2023-08-05
draft: false 
tags: [v8, tree, internals] 
toc: true 
---

# V8 source tree

```cpp
v8/
├─ src/
│  ├─ [api]
│  ├─ [asmjs]
│  ├─ [ast]
│  ├─ [base]
│  ├─ [baseline]
│  ├─ [bigint]
│  ├─ [builtins]
│  ├─ [codegen]
│  ├─ [common]
│  ├─ [compiler-dispatcher]
│  ├─ [compiler]
│  ├─ [d8]
│  ├─ [date]
│  ├─ [debug]
│  ├─ [deoptimizer]
│  ├─ [diagnostics]
│  ├─ [execution]
│  ├─ [extensions]
│  ├─ [flags]
│  ├─ [fuzzilli]
│  ├─ [handles]
│  ├─ [heap]
│  ├─ [ic]
│  ├─ [init]
│  ├─ [inspector]
│  ├─ [interpreter]
│  ├─ [json]
│  ├─ [libplatform]
│  ├─ [libsampler]
│  ├─ [logging]
│  ├─ [maglev]
│  ├─ [numbers]
│  ├─ [objects]
│  ├─ [parsing]
│  ├─ [profiler]
│  ├─ [protobuf]
│  ├─ [regexp]
│  ├─ [roots]
│  ├─ [runtime]
│  ├─ [sandbox]
│  ├─ [snapshot]
│  ├─ [strings]
│  ├─ [tasks]
│  ├─ [temporal]
│  ├─ [third_party]
│  ├─ [torque]
│  ├─ [tracing]
│  ├─ [trap-handler]
│  ├─ [utils]
│  ├─ [wasm]
│  ├─ [zone]
├─ third_party/
```

## `v8/src/api`

The `v8/src/api` directory in the V8 codebase is part of the V8 API. The V8 API provides functions for compiling and executing scripts, accessing C++ methods and data structures, handling errors, and enabling security checks¹.

Your application can use V8 just like any other C++ library. Your C++ code accesses V8 through the V8 API by including the header `include/v8.h`¹. This allows you to embed V8 in your C++ program and run a JavaScript program from it³.

For example, you can make your own application’s C++ objects and methods available to JavaScript, and make JavaScript objects and functions available to your C++ application¹. This is particularly useful for applications that want to provide a scripting interface or take advantage of the V8 JavaScript engine³.

(1) Getting started with embedding V8 · V8. [https://v8.dev/docs/embed](https://v8.dev/docs/embed).
(2) Explain V8 engine in Node.js - GeeksforGeeks. [https://www.geeksforgeeks.org/explain-v8-engine-in-node-js/](https://www.geeksforgeeks.org/explain-v8-engine-in-node-js/).

## `v8/src/asmjs`

The `v8/src/asmjs` directory in the V8 JavaScript engine's codebase is related to `asm.js`. `asm.js` is a subset of JavaScript designed to allow computer software written in languages such as C to be run as web applications while maintaining performance characteristics considerably better than standard JavaScript².

It's typically used as an intermediate language, generated through the use of a compiler that takes source code in a language such as C++ and outputs `asm.js`. This allows the code to run with performance characteristics closer to that of native code than standard JavaScript².

(1) asm.js - Wikipedia. [https://en.wikipedia.org/wiki/Asmjs](https://en.wikipedia.org/wiki/Asmjs).
(2) GitHub - v8/v8: The official mirror of the V8 Git repository. [https://github.com/v8/v8](https://github.com/v8/v8).

## `v8/src/ast`

The `v8/src/ast` directory of the V8 JavaScript engine contains the code related to the Abstract Syntax Tree (AST). The AST is a crucial part of the V8 engine as it represents the structure of the JavaScript code².

When V8 compiles JavaScript code, it first transforms the code into an AST. This AST is then used to generate machine code³. If you're willing to modify V8 and compile your own version of it, you can access the AST in the `compiler.cc` file where `MakeCode` is called throughout `MakeFunctionInfo`. This function transforms the AST that is stored in the passed-in `CompilationInfo` object into native code².

(1) Access the Abstract Syntax Tree of V8 Engine - Stack Overflow. [https://stackoverflow.com/questions/9451067/access-the-abstract-syntax-tree-of-v8-engine](https://stackoverflow.com/questions/9451067/access-the-abstract-syntax-tree-of-v8-engine).
(2) How V8 JavaScript engine works step by step [with diagram] - Medium. [https://cabulous.medium.com/how-v8-javascript-engine-works-5393832d80a7](https://cabulous.medium.com/how-v8-javascript-engine-works-5393832d80a7).
(3) GitHub - v8/v8: The official mirror of the V8 Git repository. [https://github.com/v8/v8](https://github.com/v8/v8).
(4) V8 docs. [https://v8.dev/docs](https://v8.dev/docs).
(5) Source Code. [https://chromium.googlesource.com/v8/v8/+/main/src/ast/ast.h](https://chromium.googlesource.com/v8/v8/+/main/src/ast/ast.h).

## `v8/src/base`

The `v8/src/base` directory in the V8 codebase typically contains base utility classes and functions that are used across the V8 engine. However, the specific contents and their purposes can vary as the codebase evolves. 

(1) GitHub - v8/v8: The official mirror of the V8 Git repository. [https://github.com/v8/v8](https://github.com/v8/v8).
(2) Documentation · V8. [https://v8.dev/docs](https://v8.dev/docs).
(3) Checking out the V8 source code · V8. [https://v8.dev/docs/source-code](https://v8.dev/docs/source-code).

## `v8/src/baseline`

The `v8/src/baseline` directory in the V8 codebase is related to the Liftoff compiler, which is a baseline compiler for WebAssembly in V8¹.

The goal of Liftoff is to reduce startup time for WebAssembly-based apps by generating code as fast as possible. Code quality is secondary, as hot code is eventually recompiled with TurboFan anyway¹. Liftoff avoids the time and memory overhead of constructing an intermediate representation (IR) and generates machine code in a single pass over the bytecode of a WebAssembly function¹.

With Liftoff and TurboFan, V8 now has two compilation tiers for WebAssembly: Liftoff as the baseline compiler for fast startup and TurboFan as optimizing compiler for maximum performance¹. This is likely the purpose of the `v8/src/baseline` directory. 

(1) Liftoff: a new baseline compiler for WebAssembly in V8 · V8. [https://v8.dev/blog/liftoff](https://v8.dev/blog/liftoff).
(2) GitHub - v8/v8: The official mirror of the V8 Git repository. [https://github.com/v8/v8](https://github.com/v8/v8).
(3) WebAssembly compilation pipeline · V8. [https://v8.dev/docs/wasm-compilation-pipeline](https://v8.dev/docs/wasm-compilation-pipeline).

## `v8/src/bigint`

The `v8/src/bigint` directory in V8 is dedicated to the implementation of BigInts¹². BigInts are a numeric primitive in JavaScript that can represent integers with arbitrary precision². This means they can safely store and operate on large integers, even beyond the safe integer limit for Numbers².

The BigInts are stored in memory as objects, large enough to hold all the BigInt's bits, in a series of chunks, which are called "digits"¹. These digits range from 0 to 4294967295.

(1) Adding BigInts to V8 · V8. [https://v8.dev/blog/bigint](https://v8.dev/blog/bigint).
(2) BigInt: arbitrary-precision integers in JavaScript · V8. [https://v8.dev/features/bigint](https://v8.dev/features/bigint).

## `v8/src/builtins`

The `v8/src/builtins` directory in V8 is dedicated to the implementation of built-in functions¹². Built-in functions in V8 come in different flavors with respect to implementation, depending on their functionality, performance requirements, and sometimes plain historical development¹.

Built-ins in V8 can be seen as chunks of code that are executable by the VM at runtime². A common use case is to implement the functions of built-in objects (such as RegExp or Promise), but built-ins can also be used to provide other internal functionality (e.g., as part of the IC system)².

V8’s built-ins can be implemented using a number of different methods, each with different trade-offs²:

- **Platform-dependent assembly language**: Can be highly efficient, but need manual ports to all platforms and are difficult to maintain.
- **C++**: Very similar in style to runtime functions and have access to V8’s powerful runtime functionality, but usually not suited to performance-sensitive areas.
- **JavaScript**: Concise and readable code, access to fast intrinsics, but frequent usage of slow runtime calls, subject to unpredictable performance through type pollution, and subtle issues around (complicated and non-obvious) JS semantics.
- **CodeStubAssembler**: Provides efficient low-level functionality that is very close to assembly language while remaining platform-independent and preserving readability².

(1) Built-in functions · V8. [https://v8.dev/docs/builtin-functions](https://v8.dev/docs/builtin-functions).
(2) Builtins - Google Open Source. [https://chromium.googlesource.com/external/github.com/v8/v8.wiki/+/30daab8a92562c331c93470f54877fa02b9422b5/CodeStubAssembler-Builtins.md](https://chromium.googlesource.com/external/github.com/v8/v8.wiki/+/30daab8a92562c331c93470f54877fa02b9422b5/CodeStubAssembler-Builtins.md).
(3) V8 Torque builtins · V8. [https://v8.dev/docs/torque-builtins](https://v8.dev/docs/torque-builtins).
(4) undefined. [https://tc39.github.io/ecma262/](https://tc39.github.io/ecma262/).

## `v8/src/codegen`

The `v8/src/codegen` directory in the V8 codebase is part of the V8 JavaScript Engine¹.

The term 'codegen' typically refers to the process of generating machine code from an intermediate representation (like bytecode). This is a crucial part of how JavaScript engines like V8 work. They take JavaScript code, convert it into an intermediate form, then generate machine code that can be executed by the host system.

## `v8/src/common`

The `v8/src/common` directory is part of the V8 JavaScript Engine's source code¹. 

In general, directories within a project's source code serve to organize related files and code modules. For instance, the `src` directory typically contains the source files for the project. A `common` subdirectory would usually contain code modules or files that are shared across multiple parts of the project.

(1) Getting Around the Chromium Source Code Directory Structure. [https://www.chromium.org/developers/how-tos/getting-around-the-chrome-source-code/](https://www.chromium.org/developers/how-tos/getting-around-the-chrome-source-code/).

## `v8/src/compiler-dispatcher`

The `v8/src/compiler-dispatcher` directory is part of the V8 JavaScript engine's source code. 

The V8 engine compiles JavaScript into machine code before executing it, rather than interpreting it in real-time. This leads to faster execution and better performance. The `compiler-dispatcher` directory likely contains code related to managing and dispatching these compilation tasks¹.

(1) src/compiler-dispatcher/compiler-dispatcher.h - v8/v8 - Git at Google. [https://chromium.googlesource.com/v8/v8/+/973ec26eb174ad434a04836b93aa739e320ae89d/src/compiler-dispatcher/compiler-dispatcher.h](https://chromium.googlesource.com/v8/v8/+/973ec26eb174ad434a04836b93aa739e320ae89d/src/compiler-dispatcher/compiler-dispatcher.h).

## `v8/src/compiler`

The `v8/src/compiler` directory in the V8 JavaScript Engine's source code is dedicated to the implementation of the V8 compiler¹².

V8 is Google's open-source JavaScript and WebAssembly engine, written in C++¹². It compiles and executes JavaScript source code, handles memory allocation for objects, and garbage collects objects it no longer needs².

The `compiler` directory contains the source code for the V8 compiler, which is responsible for translating JavaScript code into machine code that can be executed by the host system¹². This process involves several steps, including parsing the JavaScript code, optimizing it for performance, and finally generating the machine code¹².

## `v8/src/d8`

The `v8/src/d8` directory in the V8 JavaScript Engine's source code is dedicated to the implementation of the `d8` shell¹².

`d8` is a simple shell for V8 JavaScript engine¹. It's used for debugging and testing JavaScript and WebAssembly¹. The `d8` shell can be run standalone, or it can be embedded into any C++ application².

(1) V8 Installation and d8 shell usage · GitHub. [https://gist.github.com/kevincennis/0cd2138c78a07412ef21](https://gist.github.com/kevincennis/0cd2138c78a07412ef21).

## `v8/src/date`

Date and timezone functions

## `v8/src/debug`

The `v8/src/debug` directory in V8 is related to the debugging functionality provided by V8. Users usually interact with the V8 debugger through the Chrome DevTools interface. Embedders (including DevTools) need to rely directly on the Inspector Protocol¹.

The `v8/src/debug` directory contains the necessary code to establish communication channels for messages sent from the embedder to V8 and from V8 to the embedder¹. This allows for interaction with the Inspector through d8¹.

For example, d8 forwards inspector messages to JavaScript. The following code implements a basic, but fully functional interaction with the Inspector through d8¹:

```jsx
// inspector-demo.js
// Receiver function called by d8.
function receive(message) {
  print(message)
}

const msg = JSON.stringify({
  id: 0,
  method: 'Debugger.enable',
});

// Call the function provided by d8.
send(msg);

// Run this file by executing 'd8 --enable-inspector inspector-demo.js'.

```

This directory is crucial for providing debugging support in V8¹.

(1) Debugging over the V8 Inspector Protocol · V8. [https://v8.dev/docs/inspector](https://v8.dev/docs/inspector).

## `v8/src/deoptimizer`

The `v8/src/deoptimizer` in the V8 Chrome engine is related to the process of deoptimization. Deoptimization is a mechanism in the V8 JavaScript engine that allows it to optimize JavaScript execution by making certain assumptions about the code it's executing. If these assumptions turn out to be incorrect, the engine needs to revert, or "deoptimize", the code¹.

There are two types of deoptimization: eager and lazy¹. Eager deoptimization happens when the currently executing function has to be deoptimized. Lazy deoptimization, which is what `v8/src/deoptimizer` is primarily concerned with, is a "scheduled" deoptimization of a function that currently has one or more activations on the stack, but isn't the currently executing function¹.

In lazy deoptimization, the code has been marked as dependent on some assumption which is checked elsewhere and can trigger deoptimization the next time the code is executed¹. Deoptimizing implies having to rewrite the stack frame's contents, which is prohibitively difficult to do for any non-topmost stack frames, so such functions are marked for deoptimization, and will get deoptimized as soon as control returns to them (i.e., when they become the topmost stack frame)¹.

In summary, the `v8/src/deoptimizer` plays a crucial role in the V8 engine's ability to optimize JavaScript execution by providing a mechanism to safely revert optimizations when necessary. This contributes to the V8 engine's overall performance and efficiency in executing JavaScript code.

(1) In V8, what is lazy deoptimization, and how does it happen?. [https://stackoverflow.com/questions/70514415/in-v8-what-is-lazy-deoptimization-and-how-does-it-happen](https://stackoverflow.com/questions/70514415/in-v8-what-is-lazy-deoptimization-and-how-does-it-happen).

## `v8/src/diagnostics`

The `v8/src/diagnostics` directory in the V8 Chrome engine is related to the diagnostic tools and functionalities of the V8 engine. 

Diagnostic tools are crucial for the maintenance and optimization of a JavaScript engine like V8. They can provide insights into the performance of the engine, help identify bottlenecks or inefficiencies, and assist developers in debugging issues.

## `v8/src/execution`

The `v8/src/execution` directory in the V8 Chrome engine is likely related to the execution of JavaScript code.

## `v8/src/extensions`

The `v8/src/extensions` directory in the V8 Chrome engine is likely related to the extensions of the V8 engine. While there isn't specific documentation available for this directory, it's safe to assume that it contains code related to the extension mechanisms provided by the V8 engine.

Extensions in V8 could refer to additional functionalities or modules that can be added to the V8 engine to extend its capabilities. These could include additional APIs, debugging tools, or other features that are not part of the core V8 engine but can be added as needed.

## `v8/src/flags`

The `v8/src/flags` directory in the V8 Chrome engine is likely related to the handling of flags in the V8 engine. Flags in V8 are used to enable or disable certain features or behaviors of the engine⁵. They can be used for debugging, development, and performance tuning⁴.

Flags can be passed to the V8 engine at startup to configure its behavior⁴. For example, a flag could be used to enable a new optimization feature, to set the size of the memory heap, or to enable additional logging or profiling⁴.

[https://chromium.googlesource.com/v8/v8/+/master/src/flag-definitions.h](https://chromium.googlesource.com/v8/v8/+/master/src/flag-definitions.h)

Spoiler: there's nothing in there that you'd want to mess with. Most flags are for debugging and development. When it makes sense to turn on a flag, we turn it on by default.

You can also have a look at

[https://www.chromium.org/developers/how-tos/run-chromium-with-flags](https://www.chromium.org/developers/how-tos/run-chromium-with-flags)

Normally you don't want to mess with these if you are not actively working on V8/Chromium.

Running d8 with '--help' will also show you the available flags with a short description.

The `v8/src/flags` directory likely contains the code that defines these flags and handles their parsing and application. 

(1) V8 Flags - Google Groups. [https://groups.google.com/g/v8-users/c/VLZxrhv-wks](https://groups.google.com/g/v8-users/c/VLZxrhv-wks).
(2) undefined. [https://chromium.googlesource.com/v8/v8/+/master/src/flag-definitions.h](https://chromium.googlesource.com/v8/v8/+/master/src/flag-definitions.h).

## `v8/src/fuzzilli`

The `v8/src/fuzzilli` directory in the V8 Chrome engine is related to Fuzzilli, a JavaScript engine fuzzer¹. Fuzzilli is a (coverage-)guided fuzzer for dynamic language interpreters based on a custom intermediate language ("FuzzIL") which can be mutated and translated to JavaScript¹.

Fuzzilli is used to find bugs in JavaScript engines¹. It generates JavaScript samples that are specifically designed to discover type confusion bugs³. This design allows the fuzzer to scale to many cores on a single machine as well as to many different machines¹.

In summary, the `v8/src/fuzzilli` directory contains the Fuzzilli fuzzer.

(1) GitHub - googleprojectzero/fuzzilli: A JavaScript Engine Fuzzer. [https://github.com/googleprojectzero/fuzzilli](https://github.com/googleprojectzero/fuzzilli).
(2) Fuzzing JavaScript Engines with Fuzzilli · Doyensec's Blog. [https://blog.doyensec.com/2020/09/09/fuzzilli-jerryscript.html](https://blog.doyensec.com/2020/09/09/fuzzilli-jerryscript.html).

## `v8/src/handles`

The `v8/src/handles` directory in the V8 Chrome engine is related to the handling of handles in the V8 engine. In V8, handles are pointers used to manage JavaScript objects¹. They are a key part of V8's garbage collection mechanism¹.

Handles provide a way for the V8 engine to track and manage JavaScript objects. When a JavaScript object is created, a handle is also created for that object¹. This handle is then used by the V8 engine to reference and manipulate the JavaScript object¹.

The `v8/src/handles` directory likely contains the code that manages these handles. This could include the creation and deletion of handles, as well as other operations that involve handles¹.

## `v8/src/heap`

The `v8/src/heap` directory in the V8 Chrome engine is related to the management of the heap memory allocation³. The V8 engine's heap is where it allocates memory for objects.

V8 has a hard limit on its heap size¹. This serves as a safeguard against applications with memory leaks¹. When an application reaches this hard limit, V8 does a series of last resort garbage collections¹. If the garbage collections do not help to free memory, V8 stops execution and reports an out-of-memory failure¹.

The `v8/src/heap` directory likely contains the code that manages these aspects of memory management. This could include the creation and deletion of memory spaces, as well as other operations that involve memory management¹.

However, the exact contents and functionalities provided by `v8/src/heap` can vary between different versions of the V8 engine. For the most accurate information, it's recommended to refer to the source code or the official V8 documentation¹².

(1) Working of JavaScript’s Chrome V8 Engine - Medium. [https://medium.com/jsblend/working-of-javascripts-chrome-v8-engine-8eb5ca102a2a](https://medium.com/jsblend/working-of-javascripts-chrome-v8-engine-8eb5ca102a2a).
(2) One small step for Chrome, one giant heap for V8 · V8. [https://v8.dev/blog/heap-size-limit](https://v8.dev/blog/heap-size-limit).

## `v8/src/ic`

The `v8/src/ic` directory in the V8 Chrome engine is related to Inline Caches (ICs). ICs are a fundamental optimization in V8 used to speed up property access¹.

When V8 is executing JavaScript code and encounters a property access, it doesn't know the shape of the object being accessed. To optimize this, V8 uses ICs to remember the shape of objects from previous property accesses, which can then be used to speed up future accesses¹.

The `v8/src/ic` directory likely contains the code that manages these ICs. This could include the creation and deletion of ICs, as well as other operations that involve ICs¹.

## `v8/src/init`

The `v8/src/init` directory in the V8 Chrome engine is likely related to the initialization of the V8 engine. Initialization is a crucial step in the life cycle of any software component, and for a JavaScript engine like V8, this would involve setting up the necessary state and resources for the engine to execute JavaScript code¹.

This could include tasks such as setting up memory spaces, initializing internal data structures, setting up the JavaScript global environment, and more¹. The exact tasks performed during initialization can vary between different versions of the V8 engine¹.

## `v8/src/inspector`

The `v8/src/inspector` directory in the V8 Chrome engine is related to the V8 Inspector Protocol, which provides extensive debugging functionality¹. Users typically interact with the V8 debugger through the Chrome DevTools interface, while embedders (including DevTools) rely directly on the Inspector Protocol¹.

The Inspector Protocol allows for communication between V8 and the embedder, enabling the sending and receiving of debugging information¹. This includes setting up communication channels for messages sent from the embedder to V8 and from V8 to the embedder¹.

(1) Debugging over the V8 Inspector Protocol · V8. [https://v8.dev/docs/inspector](https://v8.dev/docs/inspector).

## `v8/src/interpreter`

The `v8/src/interpreter` directory in the V8 Chrome engine is related to the V8's interpreter, called Ignition. The interpreter is a key component of V8 that compiles JavaScript code and generates non-optimized machine code.

Ignition generates byte-code, which is good for code that only needs to run once. The byte-code runs inside the JavaScript engine itself. Interpreted code is faster to get something running but is a bit slower.

On runtime, the machine code is analyzed and re-compiled for optimal performance. This process is part of the Just-in-Time (JIT) compilation, which is how V8 achieves its speed.

(1) Deep Dive in to JavaScript Engine - (Chrome V8). [https://dev.to/edisonpappi/how-javascript-engines-chrome-v8-works-50if](https://dev.to/edisonpappi/how-javascript-engines-chrome-v8-works-50if).
(2) What is V8 JavaScript Engine? - StackPath. [https://www.stackpath.com/edge-academy/what-is-v8-javascript-engine/](https://www.stackpath.com/edge-academy/what-is-v8-javascript-engine/).

## `v8/src/json`

The `v8/src/json` directory in the V8 Chrome engine is likely related to the handling of JSON (JavaScript Object Notation) data. JSON is a lightweight data-interchange format that is easy for humans to read and write and easy for machines to parse and generate¹.

In the context of a JavaScript engine like V8, JSON data handling would involve parsing JSON data into JavaScript objects, and stringifying JavaScript objects into JSON data¹. These operations are fundamental for any JavaScript engine, as they enable JavaScript code to interact with JSON data, which is commonly used in web applications for data storage and transfer¹.

The `v8/src/json` directory likely contains the code that manages these JSON-related operations.

## `v8/src/libplatform`

The `v8/src/libplatform` directory in the V8 Chrome engine is likely related to the platform-specific code in the V8 engine. This could include code that deals with threading, event loops, and other platform-specific operations¹².

## `v8/src/libsampler`

The `v8/src/libsampler` directory in the V8 Chrome engine is likely related to the sampling profiler in the V8 engine. A sampling profiler periodically samples what code is currently being executed by the engine¹². This can be used to understand where the engine spends most of its time, which can be useful for identifying performance bottlenecks¹.

## `v8/src/logging`

The `v8/src/logging` directory in the V8 Chrome engine is related to the logging functionality of the V8 engine¹. Logging is a crucial aspect of any software system, and in the context of a JavaScript engine like V8, it can provide valuable insights into the engine's operation¹.

The logging functionality can be used for various purposes, such as debugging, performance tuning, and understanding the behavior of the engine¹. For example, V8’s logging code contains some optimizations to simplify logging state checks¹.

(1) Profiling Chromium with V8 · V8. [https://v8.dev/docs/profile-chromium](https://v8.dev/docs/profile-chromium).

## `v8/src/maglev`

The `v8/src/maglev` directory in the V8 Chrome engine is related to Maglev, a new optimizing compiler introduced in Chrome M117¹. Maglev is designed to be a fast optimizing compiler that generates good enough code, fast enough¹.

Maglev sits between the existing Sparkplug and TurboFan compilers in V8's architecture¹. All JavaScript code is first compiled to Ignition bytecode and executed by interpreting it¹. During execution, V8 tracks how the program behaves, including tracking object shapes and types¹. Both the runtime execution metadata and bytecode are fed into the optimizing compiler to generate high-performance, often speculative, machine code that runs significantly faster than the interpreter can¹.

Maglev aims to generate very performant machine code very quickly that is faster and can help save battery life³. Google reports that Maglev yields around a 7.5% boost to the JetStream benchmark and around a 5% boost for the Speedometer benchmark³.

In summary, the `v8/src/maglev` directory plays a crucial role in V8's performance optimization strategy, contributing to the speed and efficiency of JavaScript execution¹. 

(1) Maglev - V8’s Fastest Optimizing JIT · V8. [https://v8.dev/blog/maglev](https://v8.dev/blog/maglev).
(2) Google Chrome Begins Rollout Of New "Maglev" Mid-Tier Compiler. [https://www.phoronix.com/news/Google-Chrome-Maglev-Compiler](https://www.phoronix.com/news/Google-Chrome-Maglev-Compiler).
(3) Design Doc: Maglev - a new mid-tier optimising compiler - Google Groups. [https://groups.google.com/g/v8-dev/c/pcmkHmznjPM](https://groups.google.com/g/v8-dev/c/pcmkHmznjPM).
(4) New Chrome Maglev compiler boosts Speedometer, Jetstream - 9to5Google. [https://9to5google.com/2023/06/02/chrome-maglev-benchmarks/](https://9to5google.com/2023/06/02/chrome-maglev-benchmarks/).

## `v8/src/numbers`

The `v8/src/numbers` directory in the V8 Chrome engine likely contains the implementation of number-related functionalities¹. This could include the handling of different number types in JavaScript, such as integers and floating-point numbers, and the implementation of number-related operations¹.

## `v8/src/objects`

The `v8/src/objects` directory in the V8 JavaScript engine source code contains the implementation of all the data types, operators, objects, and functions specified in the ECMA standard.

## `v8/src/parsing`

The `v8/src/parsing` directory in the V8 engine source code is primarily responsible for parsing JavaScript code. Here's a brief overview of its role:

1. **Parsing JavaScript Code**: The first step in executing JavaScript code is to parse it into an Abstract Syntax Tree (AST). The V8's parser, which resides in the `v8/src/parsing` directory, does this job²⁴.
2. **Lexical Analysis**: Before the code is parsed into an AST, it is first converted into tokens in a process known as Lexical Analysis². A Scanner consumes a stream of Unicode characters, combines them into tokens, and removes all whitespace, newlines, and comments².
3. **Syntactical Analysis**: Once the engine converts your code into tokens, it's time to convert it into an Abstract Syntax Tree. This phase is called Syntax Analysis². The tokens are converted into an Abstract Syntax Tree using V8's Parser and the language syntax validation also happens during this phase².
4. **Generating Bytecode**: Once the AST is generated, it is sent to Ignition (another component of the V8 engine) which converts it into bytecode². This bytecode is then interpreted and executed².

In summary, the `v8/src/parsing` directory contains the code that is responsible for the initial stages of processing JavaScript code, converting it from raw text into a format that can be efficiently executed by the V8 engine. It's a crucial part of how V8 works².

(1) Chrome V8 Engine - Working - DEV Community. [https://dev.to/khattakdev/chrome-v8-engine-working-1lgi](https://dev.to/khattakdev/chrome-v8-engine-working-1lgi).
(2) Basics of understanding Chrome’s V8 Engine - Medium. [https://medium.com/@duartekevin91/basics-of-understanding-chromes-v8-engine-c5c8ec61fa6b](https://medium.com/@duartekevin91/basics-of-understanding-chromes-v8-engine-c5c8ec61fa6b).
(3) Documentation · V8. [https://v8.dev/docs](https://v8.dev/docs).
(4) What is Chrome V8? | Cloudflare. [https://www.cloudflare.com/learning/serverless/glossary/what-is-chrome-v8/](https://www.cloudflare.com/learning/serverless/glossary/what-is-chrome-v8/).
(5) Let’s Understand Chrome V8: Compiler Workflow, Parser. [https://javascript.plainenglish.io/lets-understand-chrome-v8-compiler-workflow-parser-36941d0ff204](https://javascript.plainenglish.io/lets-understand-chrome-v8-compiler-workflow-parser-36941d0ff204).

## `v8/src/profiler`

The `v8/src/profiler` in the V8 Chrome engine is responsible for profiling the execution of JavaScript and C/C++ code. It's a built-in sample-based profiler that records stacks of both JavaScript and C/C++ code¹.

Profiling is turned off by default, but can be enabled via the `--prof` command-line option¹. When profiling, V8 generates a `v8.log` file which contains profiling data¹. This data can be used to find bottlenecks and spot things that are unexpected, such as too much time spent in unoptimized code³.

The profiler monitors and watches code to optimize it⁴. It includes a data collector called the monitor or profiler which looks for the same code that is executed multiple times like a looping function⁵. This information is crucial for Just-In-Time (JIT) compilers which run in real time, interpret code, and execute it⁵.

(1) Using V8’s sample-based profiler · V8. [https://v8.dev/docs/profile](https://v8.dev/docs/profile).
(2) Creating V8 profiling timeline plots - The Chromium Projects. [https://www.chromium.org/developers/creating-v8-profiling-timeline-plots/](https://www.chromium.org/developers/creating-v8-profiling-timeline-plots/).
(3) Profiling Chromium with V8 · V8. [https://v8.dev/docs/profile-chromium](https://v8.dev/docs/profile-chromium).

## `v8/src/protobuf`

The `v8/src/protobuf` directory in the V8 Chrome engine likely contains the source code for Protocol Buffers (protobuf), a language-neutral, platform-neutral, extensible mechanism for serializing structured data.

## `v8/src/regexp`

The `v8/src/regexp` directory in the V8 Chrome engine contains the implementation of regular expressions (RegExp).

In its default configuration, V8 compiles regular expressions to native code upon the first execution². As part of the work on JIT-less V8, an interpreter for regular expressions was introduced². Interpreting regular expressions has the advantage of using less memory, but it comes with a performance penalty².

To mitigate this, V8 uses a 'tier-up' strategy for RegExp². Initially, all regular expressions are compiled to bytecode and interpreted, which saves a lot of memory². If a regular expression with the same pattern is used again, it is considered 'hot' and is recompiled to native code². From this point on, the execution continues as fast as possible².

(1) Improving V8 regular expressions · V8. [https://v8.dev/blog/regexp-tier-up](https://v8.dev/blog/regexp-tier-up).
(2) Speeding up V8 regular expressions · V8. [https://v8.dev/blog/speeding-up-regular-expressions](https://v8.dev/blog/speeding-up-regular-expressions).

## `v8/src/roots`

The `v8/src/roots` directory in the V8 JavaScript engine source code is part of the internal implementation of V8. *GC roots* are the special group of objects that are used by the garbage collector as a starting point to determine which objects are eligible for garbage collection. **A “root” is simply an object that the garbage collector assumes is reachable by default**, which then has its references traced in order to find all other current objects that are reachable. Any object that is not reachable through any reference chain of any of the root objects is considered unreachable and will eventually be destroyed by the garbage collector. **In V8, roots consist of objects in the current call stack** (i.e. local variables and parameters of the currently executing function), **active V8 handle scopes, global handles, and objects in the compilation cache**.

## `v8/src/runtime`

The `v8/src/runtime` directory in the V8 Chrome engine is part of the V8 runtime. The V8 runtime is the JavaScript engine that parses and executes script code¹. It provides rules for how memory is accessed, how the program can interact with the computer's operating system, and what program syntax is legal¹.

The runtime environment allows developers to take advantage of modern JavaScript features¹⁴. It also improves function detection¹. For example, the new runtime recognizes various function definition formats¹.

In summary, the `v8/src/runtime` directory is a crucial part of the V8 engine that enables the parsing and execution of JavaScript code, providing the runtime environment for JavaScript in both browser and server contexts.

(1) V8 Runtime Overview | Apps Script | Google for Developers. [https://developers.google.com/apps-script/guides/v8-runtime](https://developers.google.com/apps-script/guides/v8-runtime).
(2) Chrome V8 Engine - Javascript runtime for Node.js - Janishar Ali. [https://janisharali.com/blog/chrome-v8-engine-javascript-runtime-for-node-js](https://janisharali.com/blog/chrome-v8-engine-javascript-runtime-for-node-js).
(3) Apps Script V8 Runtime Explained For Non-Professional Developers. [https://www.benlcollins.com/apps-script/apps-script-v8-runtime/](https://www.benlcollins.com/apps-script/apps-script-v8-runtime/).
(4) GitHub - v8/v8: The official mirror of the V8 Git repository. [https://github.com/v8/v8](https://github.com/v8/v8).
(5) How Node.js V8 runtime is different from what we have on chrome console .... [https://www.geeksforgeeks.org/how-node-js-v8-runtime-is-different-from-what-we-have-on-chrome-console/](https://www.geeksforgeeks.org/how-node-js-v8-runtime-is-different-from-what-we-have-on-chrome-console/).

## `v8/src/sandbox`

The `v8/src/sandbox` in the V8 Chrome engine is related to the sandboxing feature of Chrome V8³. Sandboxing is a key feature that ensures each process is isolated. This means that JavaScript functions run separately on it and the execution of one piece of code does not affect any other piece of code³. This isolation helps to maintain the integrity and security of the overall system.

For instance, in Chromium, each renderer is a separate process, and the sandbox built around the renderer process prevents it from writing to a disk⁴. This helps to ensure that potentially harmful or malicious code cannot affect other processes or the system as a whole⁴.

## `v8/src/snapshot`

The `v8/src/snapshot` directory in the V8 JavaScript engine is related to the concept of **startup snapshots**².

V8 uses startup snapshots to speed up the initialization of a new JavaScript context². When a new context is created, all the built-in JavaScript functionality must be set up and initialized into V8’s heap². This process can be time-consuming if done from scratch².

To speed things up, V8 uses a shortcut: it deserializes a previously-prepared snapshot directly into the heap to get an initialized context². This is similar to thawing a frozen pizza for a quick dinner². On a regular desktop computer, this can bring the time to create a context from 40 ms down to less than 2 ms². On an average mobile phone, this could mean a difference between 270 ms and 10 ms².

Applications other than Chrome that embed V8 may require more than vanilla JavaScript². Many load additional library scripts at startup, before the “actual” application runs². For example, a simple TypeScript VM based on V8 would have to load the TypeScript compiler on startup in order to translate TypeScript source code into JavaScript on-the-fly². As of the release of V8 v4.3, embedders can utilize snapshotting to skip over the startup time incurred by such an initialization².

However, there are some limitations to this approach. The snapshot can only capture V8’s heap². Any interaction from V8 with the outside is off-limits when creating the snapshot². Such interactions include defining and calling API callbacks, creating typed arrays, and values derived from sources such as `Math.random` or `Date.now` are fixed once the snapshot has been captured². They are no longer really random nor reflect the current time².

(1) Custom startup snapshots · V8. [https://v8.dev/blog/custom-startup-snapshots](https://v8.dev/blog/custom-startup-snapshots).
(2) GitHub - v8/v8: The official mirror of the V8 Git repository. [https://github.com/v8/v8](https://github.com/v8/v8).
(3) V8 Snapshots Questions and Clarifications - Google Groups. [https://groups.google.com/g/v8-users/c/SfD5YSpU3Gw](https://groups.google.com/g/v8-users/c/SfD5YSpU3Gw).
(4) What is Chrome V8? | Cloudflare. [https://www.cloudflare.com/learning/serverless/glossary/what-is-chrome-v8/](https://www.cloudflare.com/learning/serverless/glossary/what-is-chrome-v8/).
(5) What is V8 JavaScript Engine? - StackPath. [https://www.stackpath.com/edge-academy/what-is-v8-javascript-engine/](https://www.stackpath.com/edge-academy/what-is-v8-javascript-engine/).

## `v8/src/strings`

The `v8/src/strings` directory in the V8 engine source code is likely to contain the implementation of string operations in V8.

## `v8/src/tasks`

The `v8/src/tasks` directory in the V8 Chrome engine likely contains the implementation of task-related functionalities¹. This could include task scheduling, task queues, and other related operations¹.

Tasks in V8 could refer to units of work that the engine needs to perform. These tasks can be scheduled to run at specific times or under specific conditions, and can be used to manage the execution of JavaScript code, garbage collection, and other internal operations¹.

## `v8/src/temporal`

The `v8/src/temporal` directory in the V8 JavaScript engine source code is likely related to the implementation of the Temporal API, a robust, modern, and feature-rich date/time library for JavaScript¹².

The Temporal API is a proposal for ECMAScript that aims to resolve long-standing issues with `Date` and libraries like `moment.js`, `date-fns`, and `luxon`¹². It provides standard objects and functions for working with dates, times, time zones, and durations¹².

The `v8/src/temporal` directory would contain the C++ implementations of these objects and functions, allowing JavaScript code running on the V8 engine to use the Temporal API¹².

The Temporal API was at Stage 3 of the TC39 process, which means it was a candidate for future inclusion in the ECMAScript standard¹². For the most up-to-date information, please refer to the official V8 and TC39 resources¹².

## `v8/src/third_party`

The `v8/src/third_party` directory in the V8 engine source code typically contains third-party libraries or components that the V8 engine depends on². These are external pieces of software that are not developed by the V8 team but are essential for the functionality of V8².

These third-party components can serve a variety of purposes, depending on what specific functionality they provide. They could offer utility functions, data structures, testing frameworks, or other features that the V8 engine uses².

However, without specific details about the contents of the `v8/src/third_party` directory, it's difficult to provide a more precise explanation of its purpose. The contents can vary based on the version of V8 and the specific third-party dependencies it uses at that time².

In general, the use of third-party code allows the V8 team to leverage existing, well-tested functionality instead of having to develop everything from scratch. This can lead to more robust and efficient code, and allows the team to focus on the core functionality of the V8 engine².

## `v8/src/torque`

The `v8/src/torque` in the V8 Chrome engine is a domain-specific language that is used to implement builtins². It was designed to be simple enough to directly translate the ECMAScript specification into an implementation in V8¹.

Torque allows developers contributing to the V8 project to express changes in the VM by focusing on the effects of their changes to the VM, rather than preoccupying themselves with unrelated implementation details¹. It combines a TypeScript-like syntax that eases both writing and understanding V8 code with syntax and types that reflect concepts that are already common in the CodeStubAssembler¹.

Torque provides language constructs to represent high-level, semantically-rich tidbits of V8 implementation, and the Torque compiler converts these morsels into efficient assembly code using the CodeStubAssembler¹. Both Torque’s language structure and the Torque compiler’s error checking ensure correctness in ways that were previously laborious and error-prone with direct usage of the CodeStubAssembler¹.

Most source written in Torque is checked into the V8 repository under the `src/builtins` directory, with the file extension `.tq`¹. Torque definitions of V8's heap-allocated classes are found alongside their C++ definitions, in `.tq` files with the same name as corresponding C++ files in `src/objects`¹.

(1) V8 Torque builtins · V8. [https://v8.dev/docs/torque-builtins](https://v8.dev/docs/torque-builtins).
(2) V8 Torque user manual · V8. [https://v8.dev/docs/torque](https://v8.dev/docs/torque).

## `v8/src/tracing`

The `v8/src/tracing` directory in the V8 Chrome engine is related to the tracing functionality provided by V8. Tracing in V8 works automatically when V8 is embedded in Chrome through the Chrome tracing system¹.

Here are some key points about tracing in V8¹:

- You can enable it in any standalone V8 or within an embedder that uses the Default Platform.
- To start tracing, use the `-enable-tracing` option. V8 generates a `v8_trace.json` that you can open in Chrome.
- Each trace event is associated with a set of categories, you can enable/disable the recording of trace events based on their categories.
- To enable more categories and have more fine control of the different parameters, you need to pass a config file.
- Enabling Runtime Call Statistics (RCS) in tracing requires recording the trace with the following two categories enabled: `v8` and `disabled-by-default-v8.runtime_stats`.
- To get the GC Object Statistics in tracing, you need to collect a trace with `disabled-by-default-v8.gc_stats` category enabled also you need to use the following `-js-flags`: `-track_gc_object_stats --noincremental-marking`.

For more detailed information, you may want to refer to the official V8 documentation¹. Please note that understanding the purpose of specific directories or files in a large codebase like V8 typically requires a deep understanding of the project's architecture and design. If you're interested in contributing to the project, the V8 documentation provides a good starting point¹.

(1) Tracing V8 · V8. [https://v8.dev/docs/trace](https://v8.dev/docs/trace).
(2) Stack trace API · V8. [https://v8.dev/docs/stack-trace-api](https://v8.dev/docs/stack-trace-api).

## `v8/src/trap-handler`

The `v8/src/trap-handler` directory in the V8 Chrome engine is responsible for handling traps in WebAssembly.

WebAssembly (often abbreviated as wasm) is a binary instruction format for a stack-based virtual machine¹. It is designed as a portable target for the compilation of high-level languages like C, C++, and Rust, enabling deployment on the web for client and server applications¹.

A trap in computing refers to a mechanism that interrupts the normal execution flow of a program². This can occur due to various reasons such as illegal operations, privilege violations, or other conditions that require special processing².

The trap handler in V8's WebAssembly implementation is designed to handle these traps and provide a safe execution environment². It ensures that if a WebAssembly program encounters an error condition, it doesn't crash the entire browser or lead to a security vulnerability².

In summary, the `v8/src/trap-handler` directory plays a crucial role in ensuring the robustness and security of WebAssembly execution in the V8 engine².

## `v8/src/utils`

The `v8/src/utils` directory in the V8 JavaScript engine source code is likely to contain utility functions and classes that are used across the V8 codebase. These utilities can include data structures, debugging aids, or other common functionality.

V8 is Google's open-source JavaScript engine that implements ECMAScript as specified in ECMA-262¹. It's written in C++ and is used in Google Chrome and other Chromium-based web browsers¹³⁴⁵. V8 can run standalone or can be embedded into any C++ application¹².

V8 compiles and executes JavaScript source code, handles memory allocation for objects, and garbage collects objects it no longer needs². Its stop-the-world, generational, accurate garbage collector is one of the keys to V8's performance².

## `v8/src/wasm`

The `v8/src/wasm` directory in the V8 Chrome engine is part of the WebAssembly (Wasm) implementation in V8¹. WebAssembly is a binary format that allows you to run code from programming languages other than JavaScript on the web efficiently and securely¹.

Here's a brief overview of how it works:

1. **Liftoff**: Initially, V8 does not compile any functions in a WebAssembly module. Instead, functions get compiled lazily with the baseline compiler Liftoff when the function gets called for the first time¹. Liftoff is a one-pass compiler, which means it iterates over the WebAssembly code once and emits machine code immediately for each WebAssembly instruction¹. Liftoff can compile WebAssembly code very fast, tens of megabytes per second¹.
2. **TurboFan**: Hot functions, which are functions that get executed often, get re-compiled with TurboFan, the optimizing compiler in V8 for both WebAssembly and JavaScript¹. TurboFan is a multi-pass compiler, which means that it builds multiple internal representations of the compiled code before emitting machine code¹. These additional internal representations allow optimizations and better register allocations, resulting in significantly faster code¹.
3. **Code caching**: If the WebAssembly module was compiled with `WebAssembly.compileStreaming`, then the TurboFan-generated machine code will also get cached¹. When the same WebAssembly module is fetched again from the same URL then the cached code can be used immediately without additional compilation¹.

In summary, the `v8/src/wasm` directory is a crucial part of the V8 engine that enables the parsing, compilation, and execution of WebAssembly code, providing a runtime environment for WebAssembly in both browser and server contexts.

(1) WebAssembly compilation pipeline · V8. [https://v8.dev/docs/wasm-compilation-pipeline](https://v8.dev/docs/wasm-compilation-pipeline).
(2) Experimental support for WebAssembly in V8 · V8. [https://v8.dev/blog/webassembly-experimental](https://v8.dev/blog/webassembly-experimental).
(3) What’s new in V8 8.6? Better WebAssembly! - Dev Genius. [https://blog.devgenius.io/whats-new-in-v8-8-6-better-webassembly-2a67abd766fa](https://blog.devgenius.io/whats-new-in-v8-8-6-better-webassembly-2a67abd766fa).

## `v8/src/zone`

The `v8/src/zone` directory in the V8 Chrome engine is related to the concept of a "Zone" in V8. A Zone is a region of memory used for temporary data that is allocated in the V8 engine during the compilation process¹.

Zones are used as a quick, efficient mechanism for allocating memory for temporary data structures while the engine is running. When the data is no longer needed (typically after the compilation of a function is finished), instead of deallocating each individual object, the entire Zone is discarded at once, which is faster¹.

This is part of V8's strategy to manage memory and improve performance. However, please note that this is a high-level explanation. For more detailed information, you may want to refer to the official V8 documentation or the V8 source code¹².
