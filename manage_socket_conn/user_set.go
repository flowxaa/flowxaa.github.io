package manage_socket_conn

import (
	"github.com/gorilla/websocket"
)

func init() {
	GetUserSet()
}

// 用户map 用来存储每个在线的用户id与对应的conn
type userSet struct {
	//	用户链接集  用户id => 链接对象
	users map[int]*websocket.Conn
	//lock  sync.Mutex
	//once  sync.Once
}

var us = new(userSet)

// 单例模式
func GetUserSet() *userSet {
	//us.once.Do(func() {

	us.users = make(map[int]*websocket.Conn)
	us.users[-1] = nil
	//us.lock = sync.Mutex{}
	//})
	return us
}
