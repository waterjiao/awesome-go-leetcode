# 散列表

## 散列函数
该如何构造散列函数呢？散列函数设计的基本要求：

1. 散列函数计算得到的散列值是一个非负整数
2. 如果 key1 = key2，那 hash(key1) == hash(key2)；
3. 如果 key1 ≠ key2，那 hash(key1) ≠ hash(key2)。

## 解决散列冲突
1. 开放寻址法
   
   当数据量小，装载因子小，适合采用开放寻址法
   
2. 链表法
   
   当存储数据为大对象（比链表开销大），大数据量时，适合采用链表法，比开放寻址更加灵活，可以支持更多优化策略，
   比如用红黑树代替链表
   
## LRU
### 单链表

维护有序单链表，越靠近链尾的节点是越早之前访问的节点。当有新节点访问时，从头遍历链表

1. 如果在链表中，从当前节点删除，插入链头
2. 如果不在链表中，缓存未满，插入链头；缓存已满，删除链尾，插入链头

## 哈希算法
用途：
1. 安全加密 
   
2. 唯一标识
   
3. 数据校验
   
4. 散列函数
   
5. 负载均衡
   
6. 数据分片
   
7. 分布式存储