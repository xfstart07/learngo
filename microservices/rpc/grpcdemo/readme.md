# 如何写一个 gRPC 例子

* 怎么起 gRPC 服务
* gRPC 客户端怎么调用服务

## 安装 gRPC 和 proto

安装 protocol buffers

``` 
brew install protobuf
```

安装 grpc

```
go get -u google.golang.org/grpc 
```

当安装中报错可能 `proto,protoc-gen-go` 命令并没有安装，可能会报 `protoc-gen-go: program not found or is not executable` 的错， 需要手动安装一下

```
go install github.com/golang/protobuf/{proto,protoc-gen-go}
```

## 根据 proto 生成代码

命令：

```sh
 protoc --go_out=plugins=grpc:. *.proto
```

### proto 生成的 go 文件

```go
// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}
```

