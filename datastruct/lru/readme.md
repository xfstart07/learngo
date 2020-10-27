# LRU 最近最少使用算法

LRU（Least Recently Used）思路：基本思路是在获取缓存的同时将缓存移到最新上，写入缓存时，当缓存满了，将最后（最少使用）的缓存删掉，然后将缓存写在最新上。缓存通过双向链表和 Map 来组织，优点是移动和查询都能在 O(1) 的时间复杂度完成。

## 例子

* v1 双向链表，使用循环查询版本
* v2 双向链表，使用Map查询版本
* v3 标准库双向链表，使用 Map 查询版本

# Ref

* [lru 练习题](https://leetcode-cn.com/problems/lru-cache/)
* [快取文件置换机制](https://zh.wikipedia.org/wiki/%E5%BF%AB%E5%8F%96%E6%96%87%E4%BB%B6%E7%BD%AE%E6%8F%9B%E6%A9%9F%E5%88%B6)
* [算法讲解](https://www.youtube.com/watch?v=O1Glb6FDZhk&list=PLItNLuX80m9DSNLBTLno6ifA2Vk9jpk4I&index=47)

# Redis

### 近似LRU算法

Redis的LRU算法不是一个严格的LRU实现。这意味着Redis不能选择最佳候选键来回收，也就是最久未被访问的那些键。相反，Redis 会尝试执行一个近似的LRU算法，通过采样一小部分键，然后在采样键中回收最适合(拥有最久访问时间)的那个。

### 回收策略

当内存达到限制时，Redis 具体的回收策略是通过 maxmemory-policy 配置项配置的。

以下的策略都是可用的：

- `noenviction`：不清除数据，只是返回错误，这样会导致浪费掉更多的内存，对大多数写命令（DEL 命令和其他的少数命令例外）；
- `allkeys-lru`：从所有的数据集（server.db[i].dict）中挑选最近最少使用的数据淘汰，以供新数据使用；
- `volatile-lru`：从已设置过期时间的数据集（server.db[i].expires）中挑选最近最少使用的数据淘汰，以供新数据使用；
- `allkeys-random`：从所有数据集（server.db[i].dict）中任意选择数据淘汰，以供新数据使用；
- `volatile-random`：从已设置过期时间的数据集（server.db[i].expires）中任意选择数据淘汰，以供新数据使用；
- `volatile-ttl`：从已设置过期时间的数据集（server.db[i].expires）中挑选将要过期的数据淘汰，以供新数据使用。

当 cache 中没有符合清除条件的 key 时，回收策略 volatile-lru, volatile-random 和 volatile-ttl 将会和策略 noeviction 一样返回错误。选择正确的回收策略是很重要的，取决于你的应用程序的访问模式。但是，你可以在程序运行时重新配置策略，使用 INFO 输出来监控缓存命中和错过的次数，以调优你的设置。

## Ref

* http://ifeve.com/redis-lru/
* [Redis 数据淘汰机制](https://wiki.jikexueyuan.com/project/redis/data-elimination-mechanism.html)
