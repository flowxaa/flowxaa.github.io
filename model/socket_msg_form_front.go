package model

type ConnMsg struct {
	Msg ChatMsg `json:"msg,omitempty"`
	//FormUserID int     `json:"form_user_id,omitempty"`
	T         string `json:"t,omitempty"`
	Useridx   string `json:"useridx,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Token     string `json:"token,omitempty"`
}

// ChatMsgType = 1 群聊信息 ChatMsgType = 2 一对一信息 ...
type ChatMsg struct {
	ChatMsgType int                    `json:"chat_msg_type,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}
