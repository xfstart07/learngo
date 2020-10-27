# Snowflake 算法

* 使用ID的41位存储毫秒精度的时间戳, 表示（1L<<41）/(1000L*3600*24*365)=69年的时间。
* 然后将NodeID添加到后续位中。
* 然后添加序列号，从0开始，并针对同一毫秒中生成的每个ID递增。如果您在同一毫秒内生成了足够的ID，以至于序列将滚动或溢出，则generate函数将暂停直到下一个毫秒。

```
+--------------------------------------------------------------------------+
| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
+--------------------------------------------------------------------------+
```

使用默认设置，snowflake 在大多数系统上应足够快，以每毫秒生成4096个唯一ID，每秒能产生4百万个ID。这是雪花ID格式支持的最大值。也就是说，每个操作大约243-244纳秒。

**缺点**：

* 强依赖机器时钟，如果机器上时钟回拨，会导致发号重复或者服务会处于不可用状态。

### Ref

* https://github.com/bwmarrin/snowflake
* https://juejin.im/post/5a7f9176f265da4e721c73a8
* https://tech.meituan.com/2017/04/21/mt-leaf.html
* https://github.com/sony/sonyflake
* https://github.com/baidu/uid-generator/blob/master/README.zh_cn.md
