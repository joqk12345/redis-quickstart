package main

import (
	"fmt"
	"redis-quickstart/model"
)

func main() {
	redisClient := model.NewRedisClient()
	user_id, score, task_id, task_info := redisClient.GetFaireMember()
	if user_id != "" {
		fmt.Println("User ID:", user_id)
		fmt.Println("Score:", score)
		fmt.Println("Task ID:", task_id)
		fmt.Println("Task Info:", task_info)
	} else {
		fmt.Println("No user_id found in the zset")
	}
}
