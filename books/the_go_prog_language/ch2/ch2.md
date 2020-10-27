每个中文字符在 UTF-8 中占 3 个字节。

## 数组切片

Go 语言支持用 myArray[first:last] 方式来基于数组生成一个数组切片。

有几种写法:

所有元素创建数组切片
mySlice := myArray[:]

前 5 个元素创建数组切片
mySlice := myArray[:5]

从第 5 个元素创建数组切片
mySlice := myArray[5:]

#### 直接创建切片

Go 语言提供的内置函数 make() 创建数组切片。

mySlice := make([]int, 5)

mySlice := make([]int, 5, 10) // 创建 5 个数组，预留 10 个元素的存储空间

mySlice := []int{1,2,3,4}

#### 动态增减元素

#### 元素遍历

Go 提供 range 关键字来快速遍历所有元素。
