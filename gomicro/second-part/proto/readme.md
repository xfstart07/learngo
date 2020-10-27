# 生产接口

### 普通

```bash
protoc --proto_path=. --go_out=. --micro_out=. */*.proto
```

带引用：

```bash
protoc --proto_path=${GOPATH}/src:. --go_out=. --micro_out=. api/api.proto
```