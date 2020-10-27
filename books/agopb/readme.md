# Go 高级编程(advanced go programming book)

## 第四章 RPC和Protobuf

### 例子

* [hellorpc](ch4/hellorpc) rpc例子
* [tokengrpc](ch4/tokengrpc) token认证 GRPC，带反射服务

### grcpcurl

安装

```bash
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

mac 安装

```bash
brew install grpcurl
```

查看

```bash
grpcurl weixinote.dev:1234 list
```
