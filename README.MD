1. 基于redis的任务调度策略
  - [x] 保证按照user_id的多租户调度
  - [x] 保证按照任务的优先级调度
  - [ ] lru调度
2. 任务下发使用golang实现
   - [x] 下发任务
3. 任务消费使用Python实现
   - [ ] 消费任务

## 设计原理
### 通信原理
## 消息设计

1. global_tasks:
   1. type:(hashmap) 
   2. key：task_id 
   3. value：task_info
2. user_priority:
   1. type:(zset)
   2. key：user_id
   3. value: score
   4. 备注： 用户的优先级可以不同
3. user_tasks
   1. type:(zset)
   2. task_{user_id}
   3. key：task_id
   4. value: score
   5. 备注： 一个用户一个zset 存储改用户任务优先级

## 任务下发
## 任务消费


## 为什么使用redis
- Redis（Remote Dictionary Server）是一个开源的基于内存的数据结构存储系统，可以作为数据库、缓存和消息代理使用。它支持多种数据结构，如字符串（Strings）、哈希（Hashes）、列表（Lists）、集合（Sets）、有序集合（Sorted Sets）等，并提供了丰富的功能和灵活的用途。
- ### 关键特性：

  1. **内存存储**：数据存储在内存中，因此读写速度非常快。
  2. **持久化**：支持不同级别的持久化方式，可将数据保存到磁盘以便重启后恢复数据。
  3. **多种数据结构**：支持多种复杂的数据结构，每种结构都有自己的操作命令，可以进行灵活的数据操作。
  4. **分布式**：可以设置成分布式架构，通过复制或分区来实现高可用性。
  5. **发布/订阅**：支持发布/订阅模式，使得多个客户端可以订阅频道并接收消息。
  6. **事务**：支持事务操作，可以批量执行一系列命令，并保证原子性。
  7. **Lua脚本**：允许使用Lua脚本执行复杂的操作。
  8. **高性能**：由于数据存储在内存中，并且使用C语言编写，因此具有出色的性能。

- ### 使用场景：

  1. **缓存**：作为缓存存储常用数据，加速数据访问速度。
  2. **会话存储**：存储用户会话信息，如登录状态等。
  3. **计数器**：记录访问次数、点赞数等计数功能。
  4. **消息队列**：作为消息代理进行消息发布和订阅。
  5. **实时排行榜**：存储排名信息，如网站访问量排名等。
  6. **分布式锁**：通过Redis实现分布式锁，保证多个进程或服务的同步性。

Redis是一个功能强大、灵活多样的存储系统，由于其高性能和多种数据结构的支持，在许多应用程序中被广泛使用。

## 任务结果字段

```json
{
  "task_id": "test123",
  "error_code": 0,
  "err_msg": "",
  "data": {
    "link": "s3://test/123/asdfasd/asdfasdf",
    "osspath": "s3://test/123.mp3"
  }
}
```