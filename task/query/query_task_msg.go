package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"redis-quickstart/model"
)

func main() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 密码（如果有的话）
		DB:       0,                // Redis 数据库编号
	})

	// 要查询的任务的 task_id
	taskID := "task_123"

	// 从哈希中获取任务的值
	taskJSON, err := client.HGet("global_tasks", taskID).Result()
	if err == redis.Nil {
		fmt.Println("Task not found in Redis hash")
		return
	} else if err != nil {
		fmt.Println("Failed to get task from Redis hash:", err)
		return
	}

	// 将任务的 JSON 值解析为 Task 结构体
	var task model.Task
	err = json.Unmarshal([]byte(taskJSON), &task)
	if err != nil {
		fmt.Println("Failed to unmarshal task JSON:", err)
		return
	}

	// 打印任务信息
	fmt.Println("Task found in Redis hash:")
	fmt.Println("Task ID:", task.TaskID)
	fmt.Println("Name:", task.Name)
	fmt.Println(taskJSON)
	// 关闭 Redis 客户端连接
	err = client.Close()
	if err != nil {
		fmt.Println("Failed to close Redis client:", err)
		return
	}
	fmt.Println("Redis client closed")
}
