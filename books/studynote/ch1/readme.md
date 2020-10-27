# Go 编译

```
go build -gcflags "-N -l" -ldflags=-compressdwarf=false -o main main.go

```

> `-gcflags "-N -l"` 用于关闭编译器代码优化与函数内联。

> 此外还需注意，Go 1.11 开始将调试信息压缩为 DWARF，macOS 下的 gdb 不能解释 DWARF。 因此需要使用 GDB 调试需要增加 `-ldflags=-compressdwarf=false`。

[DWARF 文件格式](https://www.ibm.com/developerworks/cn/aix/library/au-dwarf-debug-format/index.html)