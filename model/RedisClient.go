package model

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 密码（如果有的话）
		DB:       0,                // Redis 数据库编号
	})
	return &RedisClient{
		client: client,
	}
}

func (rc *RedisClient) Close() error {
	return rc.client.Close()
}

// 按照用户公平调度
func (rc *RedisClient) GetFaireMember() (string, float64, string, string) {
	//1. 查询用户，
	// 调用 ZPopMin 方法从 zset 中查询并删除分数最低的成员
	result, err := rc.client.ZPopMin("user_priority", 1).Result()
	if err != nil {
		panic(err)
	}
	defer rc.Close() // 在方法返回之前关闭连接
	if len(result) > 0 {
		member := result[0]
		user_id := member.Member.(string)
		user_score := member.Score
		fmt.Println("user id:{}, score", user_id, user_score)
		//2. 查询用户对应的任务
		task_result, err := rc.client.ZPopMax("user_tasks_" + user_id).Result()
		if err != nil {
			panic(err)
		}
		if len(result) > 0 {
			task_id := task_result[0].Member.(string)
			task_score := task_result[0].Score
			fmt.Println("Popped task_id:", task_id)
			fmt.Println("Task Score:", task_score)

			// 从 global_tasks 中获取任务信息
			task_info, err := rc.client.HGet("global_tasks", task_id).Result()
			if err != nil {
				panic(err)
			}
			// 执行任务
			// 任务执行失败的话，需要将user、task_id、global_tasks 恢复
			// 任务执行成功的话，需要将user、task_id、global_tasks 删除

			// 检查 user_tasks 是否还有其他内容， 扫尾程序， 去掉无效用户，节省系统的调度资源
			tasks, err := rc.client.ZRange("user_tasks", 0, -1).Result()
			if err != nil {
				panic(err)
			}
			if len(tasks) == 0 {
				// 从 user_priority 中删除该用户
				rc.client.ZRem("user_priority", user_id)
			} else {
				z := redis.Z{Score: user_score, Member: user_id}
				rc.client.ZAdd("user_priority", z)
			}
			return user_id, user_score, task_id, task_info
		} else {
			fmt.Println("no memeber found in the zset")
		}
	}
	return "", 0, "", ""
}

// 按照用户分数优先
func (rc *RedisClient) GetMinPriorityMember() (string, float64, string, string) {
	// 查询分数最小的成员
	result, err := rc.client.ZRangeByScoreWithScores("user_priority", redis.ZRangeBy{
		Min:    "-inf", // 最小分数
		Max:    "+inf", // 最大分数
		Offset: 0,      // 偏移量，从第一个成员开始
		Count:  1,      // 查询的成员数量，只获取一个成员
	}).Result()

	if err != nil {
		panic(err)
	}
	defer rc.Close() // 在方法返回之前关闭连接
	if len(result) > 0 {
		member := result[0]
		user_id := member.Member.(string)
		score := member.Score
		fmt.Println("user id:{}, score", user_id, score)
		//2. 查询用户对应的任务
		task_result, err := rc.client.ZPopMax("user_tasks_" + user_id).Result()
		if err != nil {
			panic(err)
		}
		if len(result) > 0 {
			task_id := task_result[0].Member.(string)
			score := task_result[0].Score
			fmt.Println("Popped task_id:", task_id)
			fmt.Println("Score:", score)

			// 从 global_tasks 中获取任务信息
			task_info, err := rc.client.HGet("global_tasks", task_id).Result()
			if err != nil {
				panic(err)
			}
			// 执行任务

			// 检查 user_tasks 是否还有其他内容
			tasks, err := rc.client.ZRange("user_tasks", 0, -1).Result()
			if err != nil {
				panic(err)
			}
			if len(tasks) == 0 {
				// 从 user_priority 中删除该用户
				rc.client.ZRem("user_priority", user_id)
			}
			return user_id, score, task_id, task_info
		} else {
			fmt.Println("no memeber found in the zset")
		}
	}
	return "", 0, "", ""
}
