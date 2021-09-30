# bufio

**写：**

问题：很多一点点数据的写操作会很损耗性能。

解决：bufio 包在写方面的作用是可以将数据先缓冲在内存中，等到缓冲到一定量后一次写入。从而减少写的次数，降低 cpu 的负担和操作开销。

**读：**

问题：同写的问题，如果一个很大文本，每次一个字符一个字符的读，很损耗性能。

解决：bufio 包在读方面的作用是可以将数据块一次读取出来，从而减少读的次数。

https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762