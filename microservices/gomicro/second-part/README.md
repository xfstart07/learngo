# 微服务例子

### api-apigateway 

把内部 RPC 接口暴露成 HTTP 接口，它接收 HTTP 请求，再转成 RCP 请求转向具体服务，最后返回 RPC 服务的响应结果。

启动 micro api 代理服务

```shell script
micro api --handler=api --namespace=go.micro.learning.api
```

请求：

```shell script
curl -i "http://weixinote.dev:8080/open/fetch?sum=12&prime=123"
```

### web-apigateway

Http 处理器是 HTTP 的反向代理，其内置有服务发现。

启动 micro web 代理服务

```shell script
micro api --handler=web --namespace=go.micro.learning.web
```

请求：

```shell script
curl 'http://weixinote.dev:8080/portal/sum?input=11'
```

### 服务

* sum-srv
* prime-srv
* log-srv
