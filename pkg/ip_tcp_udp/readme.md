# UDP 面向数据报协议(User Datagram Protocol)

又称用户数据报文协议，是一个简单的面向数据报(package-oriented)的传输层协议。UDP 只提供数据的不可靠传递，它一旦把应用程序发给网络层的数据发送出去，就不保留数据备份（所以 UDP 有时候也被认为是不可靠的数据报协议）。UDP 在 IP 数据报的头部仅仅加入了复用和数据校验。

* UDP 报头 32 位，4 个字节

### TCP 和 UDP 区别

| 空         | TCP(传输控制协议) | UDP(用户数据报协议) |
| ---------- | ----------------- | ------------------- |
| 是否连接   | 面向连接          | 面向非连接          |
| 传输可靠性 | 可靠的            | 不可靠的            |
| 应用场合   | 传输大量的数据    | 少量数据            |
| 速度       | 慢                | 快                  |

### Golang 使用 UDP

#### API

* DialUDP 创建一个 udp 的客户端-服务端连接
* ListenUDP 监听一个 ip 地址的 udp
* Write 写数据
* WriteToUDP 写数据并指定发送的 ip 地址
* ReadFromUDP 读取 udp 数据

## 资源

* 深入 Go UDP 编程 http://colobu.com/2016/10/19/Go-UDP-Programming/
