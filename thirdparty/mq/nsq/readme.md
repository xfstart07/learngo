# nsq 发布订阅例子

运行订阅者：

```
go run client.go
```

可以运行多个.

运行发布者:

```
go run server.go
```

## 测试结果

多次执行 server，consumer 接收到的数据并不全，nsq 存在丢失的情况。