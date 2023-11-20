package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"redis-quickstart/model"
	"strconv"
)

func main() {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password (if any)
		DB:       0,                // Redis database number
	})

	// Create a new task
	task := model.Task{
		ID:     1,
		TaskID: "task_123",
		Name:   "content2text",
		Type:   "single",
		Params: struct {
			URL             string `json:"url"`
			Title           string `json:"title"`
			CharacterLength int    `json:"character_len"`
			TTSConf         struct {
				Timer          string `json:"timer"`
				Speed          string `json:"speed"`
				Volume         string `json:"volume"`
				Gender         string `json:"gender"`
				TargetLanguage string `json:"target_language"`
			} `json:"tts_conf"`
		}{
			URL:             "https://www.wsj.com/tech/ai/openai-leadership-hangs-in-balance-as-sam-altmans-counte-rebellion-gains-steam-47276fa8?mod=hp_lead_pos1",
			Title:           "Sam Altman to Join Microsoft Following OpenAI Ouster",
			CharacterLength: 190,
			TTSConf: struct {
				Timer          string `json:"timer"`
				Speed          string `json:"speed"`
				Volume         string `json:"volume"`
				Gender         string `json:"gender"`
				TargetLanguage string `json:"target_language"`
			}{
				Timer:          "lisa",
				Speed:          "",
				Volume:         "",
				Gender:         "",
				TargetLanguage: "",
			},
		},
		CreateTime: "",
		Result:     "",
		GroupID:    "",
		Priority:   "",
		UserID:     1,
	}

	// Convert task to JSON
	taskJSON, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Failed to marshal task to JSON:", err)
		return
	}

	// Set the task in Redis
	//err = client.Set("task_id", taskJSON, 0).Err()
	//if err != nil {
	//	fmt.Println("Failed to set task in Redis:", err)
	//	return
	//}
	//fmt.Println("Task written to Redis")
	// 使用哈希数据结构存储任务
	err = client.HSet("global_tasks", task.TaskID, taskJSON).Err()
	if err != nil {
		fmt.Println("Failed to store task in Redis hash:", err)
		return
	}
	fmt.Println("Task stored in Redis hash: global_tasks")

	// 下发用户优先级队
	//err = client.ZAdd("user_priority", float64(1), float64(task.UserID).Err()
	z := redis.Z{Score: 0.5, Member: task.UserID}
	result := client.ZAdd("user_priority", z)
	if result.Err() != nil {
		// Handle the error
		fmt.Println("Error:", result.Err())
	} else {
		// No error, ZAdd operation was successful
		fmt.Println("ZAdd operation successful")
	}

	//下发任务队列
	t := redis.Z{Score: 0.5, Member: task.TaskID}
	user_task := "user_tasks_" + strconv.Itoa(task.UserID)
	result = client.ZAdd(user_task, t)
	if result.Err() != nil {
		// Handle the error
		fmt.Println("Error:", result.Err())
	} else {
		// No error, ZAdd operation was successful
		fmt.Println("ZAdd user_task successful")
	}

	// Close the Redis client connection
	err = client.Close()
	if err != nil {
		fmt.Println("Failed to close Redis client:", err)
		return
	}
	fmt.Println("Redis client closed")
}
