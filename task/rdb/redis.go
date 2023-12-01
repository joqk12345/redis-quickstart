package rdb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"redis-quickstart/task/model"
	"strconv"
	"time"
)

const global_tasks = "global_tasks"
const user_priority = "user_priority"
const user_task_prefix = "user_tasks_"
const task_result_list = "task_result_list"

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // Redis password (if any)
		DB:       0,                // Redis database number
	})
	return &RedisClient{
		client: client,
	}
}

func (rc *RedisClient) PutTaskRecord(record model.TaskRecord) error {
	if record.TaskID == "" {
		fmt.Println(" record taskId can't empty")
		return nil
	}
	if record.UserId == 0 {
		fmt.Println(" record userId can't zero")
		return nil
	}
	taskJSON, err := json.Marshal(record)
	if len(record.Params) == 0 {
		fmt.Println(" record  parms can't empty")
		return nil
	}
	if err != nil {
		fmt.Println("Failed to marshal task to JSON:", err)
		return nil
	}
	err = rc.client.HSet(context.Background(), global_tasks, record.TaskID, taskJSON).Err()
	if err != nil {
		fmt.Println("Failed to store task in Redis hash:", err)
		return err
	}
	// 下发用户优先级队
	//z := redis.Z{Score: 0.5, Member: record.UserId}
	z := &redis.Z{Score: 0.5, Member: strconv.Itoa(int(record.UserId))}
	result := rc.client.ZAdd(context.Background(), user_priority, z)
	if result.Err() != nil {
		fmt.Println("Error:", result.Err())
	} else {
		fmt.Println("ZAdd operation successful")
	}
	//下发任务队列
	//t := redis.Z{Score: 0.5, Member: record.TaskID}
	t := &redis.Z{Score: 0.5, Member: record.TaskID}
	user_task := user_task_prefix + strconv.Itoa(int(record.UserId))
	result = rc.client.ZAdd(context.Background(), user_task, t)
	if result.Err() != nil {
		fmt.Println("Error:", result.Err())
	} else {
		fmt.Println("ZAdd user_task successful")
	}
	return nil
}

func (rc *RedisClient) PullTaskResult() (string, error) {
	result, err := rc.client.RPop(context.Background(), task_result_list).Result()
	if err == redis.Nil {
		time.Sleep(5 * time.Second)
	} else if err != nil {
		fmt.Println("fetch data error:", err)
	}
	return result, err
}

func (rc *RedisClient) GetTaskResultLen() (int64, error) {
	result_len, err := rc.client.LLen(context.Background(), task_result_list).Result()
	if err == redis.Nil {
		time.Sleep(5 * time.Second)
	} else if err != nil {
		fmt.Println("get Result length error:", err)
	}
	return result_len, err
}

func (rc *RedisClient) PutTaskResult(result model.TaskResult) (int64, error) {

	taskResultJSON, err := json.Marshal(result)
	result_len, err := rc.client.LPush(context.Background(), task_result_list, taskResultJSON).Result()
	if err != nil {
		fmt.Println("put task Result  error:", err)
	}
	return result_len, err
}

func (rc *RedisClient) Close() error {
	err := rc.client.Close()
	if err != nil {
		return err
	}
	return nil
}
