# channel 通道

### 单向通道

**接口定义**

请想象一下， 如果一个接口声明包含了这样的声明， 会起到什么样的作用？例如：

```go
type SignalNotifier interface { Notify(c chan<- os.Signal, sig ...os.Signal) }
```

这种方式明确表达了一个实现规则：在该接口的所 有实现类型的 Notify 方法内部只能向 c 发送元素值。 这就相当于利用语法级别的约束避 免实现类型对参数 c 进行错误的操作。

**方法定义**

对 SignalNotifier 接口的声明稍作改变， 如下：

```go
type SignalNotifier interface { Notify(sig ...os.Signal) <-chan os.Signal }
```

此方法声明的约束目标是方法的调用方， 而非方法的实现方。 Notify 方法的调用方只能从作 为结果的通道中接收元素值， 而不能向其发送元素值。

**小结**

前一个版本的方法声明更适合存在于接口类型中， 因为它可以作为该接口的实现规则之一。 后一个版本的声明更适用于函数或结构体的方法， 原因是它可以约束对函数或方法的结果值的使用方式。

### 非缓冲的 Channel

**规则**

1. 发送时会阻塞，直到有接收者接收了数据
2. 接收时会阻塞，直到有发送过来的数据。

#### 同步的特性

### 源码

channel 的数据结构：

```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
```

`runtime/chan.go`

`hchan` 是 channel 的实际数据结构。