# Go 语言编程

Go 语言最主要的特性：

* 自动垃圾回收
* 更丰富的内置类型
* 函数多返回值
* 错误处理
* 匿名函数和闭包
* 类型和接口
* 并发编程
* 反射
* 语言交互性

## 并发编程

`goroutine` 是一种比线程更加轻盈、更省资源的协程。

Go 语言实现了 CSP(通信顺序进程，Communicating Sequential Process) 模型来作为 goroutine 间的推荐通信方式。在 CSP 模型中，一个并发系统由若干并行运行的顺序进程组成，每个进程不能对其他进程的变量赋值。进程之间只能通过一个通信原语实现协作。Go 语言用 channel（通道）这个概念来轻巧的实现了 CSP 模型。channel 的使用方式比较接近 Unix 系统中的管道(pipe) 概念，可以方便地进行跨 goroutine 的通信。

由于一个进程内创建的所有 goroutine 运行在同一个内存地址空间中，因此如果不同的 goroutine 不得不去访问共享的内存变量，访问前应该先获取相应的读写锁。Go 语言标准库中的 sync 包提供了完备的读写锁功能。
