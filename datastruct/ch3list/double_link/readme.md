# 双向链表

```
p --> s --> p.next
在 p 中插入 s 结点

s->prior = p
s->next = p.next
p->next->prior = s
p->next = s
```
