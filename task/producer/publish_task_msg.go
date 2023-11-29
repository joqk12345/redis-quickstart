package main

import (
	"fmt"
	"redis-quickstart/task/model"
	"redis-quickstart/task/rdb"
)

func put_task() {
	// Create a new task
	rc := rdb.NewRedisClient()
	defer rc.Close()
	var record model.TaskRecord
	record.TaskID = "test2"
	record.Name = "content2text"
	record.Type = "http"

	// 创建一个空的Params JSONMap
	// 定义结构
	type TTSConf struct {
		Timer          string `json:"timer"`
		Speed          string `json:"speed"`
		Volume         string `json:"volume"`
		Gender         string `json:"gender"`
		TargetLanguage string `json:"target_language"`
	}

	data := struct {
		URL             string  `json:"url"`
		Title           string  `json:"title"`
		CharacterLength int     `json:"character_len"`
		TTSConf         TTSConf `json:"tts_conf"`
	}{
		URL:             "https://www.wsj.com/tech/ai/openai-leadership-hangs-in-balance-as-sam-altmans-counte-rebellion-gains-steam-47276fa8?mod=hp_lead_pos1",
		Title:           "Sam Altman to Join Microsoft Following OpenAI Ouster",
		CharacterLength: 190,
		TTSConf: TTSConf{
			Timer:          "zh-CN-XiaoyiNeural",
			Speed:          "",
			Volume:         "",
			Gender:         "",
			TargetLanguage: "",
		},
	}

	// 将结构转为 map
	result := make(map[string]interface{})
	result["url"] = data.URL
	result["title"] = data.Title
	result["character_len"] = data.CharacterLength

	ttsConf := make(map[string]string)
	ttsConf["timer"] = data.TTSConf.Timer
	ttsConf["speed"] = data.TTSConf.Speed
	ttsConf["volume"] = data.TTSConf.Volume
	ttsConf["gender"] = data.TTSConf.Gender
	ttsConf["target_language"] = data.TTSConf.TargetLanguage
	result["tts_conf"] = ttsConf

	record.Params = result
	record.UserId = 2
	fmt.Println(record)
	err := rc.PutTaskRecord(record)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}

func put_task_result() {
	data := make(map[string]interface{})
	data["osspath"] = "s3://test/123.mp3"
	data["link"] = "s3://test/123/asdfasd/asdfasdf"

	taskResult := model.TaskResult{
		TaskID:    "test123",
		ErrMsg:    "",
		ErrorCode: 0,
		Data:      data,
	}
	rc := rdb.NewRedisClient()
	defer rc.Close()
	rc.PutTaskResult(taskResult)
}

func main() {
	put_task_result()
}
