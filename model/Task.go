package model

type Task struct {
	ID     int    `json:"id"`
	TaskID string `json:"task_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Params struct {
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
	} `json:"params"`
	CreateTime string `json:"create_time"`
	Result     string `json:"result"`
	GroupID    string `json:"group_id"`
	Priority   string `json:"priority"`
	UserID     int    `json:"user_id"`
}
