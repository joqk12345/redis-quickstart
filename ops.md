## ops
1. connection
```shell
redis-cli
```

## mutiply 
3. hash

```sh
hkeys global_tasks
hget global_tasks task_123
```
2. zset

```sh

zadd user_priority 1 2 
zrange user_priority 0 5 
zrange user_priority 0 5 withscores

```
3. user_tasks

```shell
keys user_tasks*

zrange user_tasks_1 0 5 withscores

```

4. task_result_list

* 查看指定范围的内容
```shell
LRANGE task_result_list 0 10
```
* 查看队列长度
```shell
LLEN task_result_list
```