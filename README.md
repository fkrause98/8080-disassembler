## 8080-disassembler
Simple 8080 disassembler in Go, following [Emulator 101](http://emulator101.com/),
still a WIP.
## Output example:
Using a hex dump example named "invaders.bin":
```shell
❯ go run main.go
0x0: NOP 
0x2: NOP 
0x4: NOP 
0x6: JMP 0x18d4 
0xC: NOP 
0xE: NOP 
0x10: PUSH PSW 
0x12: PUSH B 
0x14: PUSH D 
0x16: PUSH H 
0x18: JMP 0x008c 
0x1E: NOP 
0x20: PUSH PSW 
0x22: PUSH B 
0x24: PUSH D 
0x26: PUSH H 
0x28: MVI A,0x80 
0x2C: STA 0x2072 
0x32: LXI H,0x20c0 
0x38: DCR M 
0x3A: CALL 0x17cd 
0x40: IN 0x01 
0x44: RRC 
0x46: JC 0x0067 
0x4C: LDA 0x20ea 
0x52: ANA A 
0x54: JZ 0x0042 
0x5A: LDA 0x20eb 
0x60: CPI 0x99 
0x64: JZ 0x003e 
0x6A: ADI 0x01 
0x6E: DAA 
0x70: STA 0x20eb 
0x76: CALL 0x1947 
0x7C: XRA A 
0x7E: STA 0xea20 
```
