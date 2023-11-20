1.  hash

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