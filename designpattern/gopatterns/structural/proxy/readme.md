# 代理模式

代理模式提供了一个对象，用于控制对另一个对象的访问，拦截所有调用。

## 实现

代理可以与任何东西连接：网络连接，内存中的大对象，文件或其他昂贵或无法复制的资源。

```go 
  // 代理和对象必须实现一样的接口方法
  // To use proxy and to object they must implement same methods
    type IObject interface {
        ObjDo(action string)
    }

    // 真实对象，委托代理的数据
    // Object represents real objects which proxy will delegate data
    type Object struct {
        action string
    }

    // 真实对象实现接口方法，处理所有的逻辑
    // ObjDo implements IObject interface and handel's all logic
    func (obj *Object) ObjDo(action string) {
        // Action behavior
        fmt.Printf("I can, %s", action)
    }

    // 代理对象
    // ProxyObject represents proxy object with intercepts actions
    type ProxyObject struct {
        object *Object
    }

    // 代理实现接口方法，调用真实对象前的拦截对象
    // ObjDo are implemented IObject and intercept action before send in real Object
    func (p *ProxyObject) ObjDo(action string) {
        if p.object == nil {
            p.object = new(Object)
        }
        if action == "Run" {
            p.object.ObjDo(action) // Prints: I can, Run
        }
    }
```