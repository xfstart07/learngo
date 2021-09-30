# sync.atomic

### 实现原理

交换操作：

```
old = *addr
*addr = new
return old
```

比较并交换：

```
if *addr == old {
    *addr = new
    return true
}
return false
```

加操作：

```
*addr += delta
return *addr
```

加载和存储操作：

```
return *addr
*addr = val
```

## 常用场景

1. 在可能存在并发修改配置的情况下，用来存储系统配置。
2. 并发修改Int相关的类型。例如 Int32、Int64、uint32 。