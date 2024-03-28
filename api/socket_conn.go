package api

import (
	"WebSocketDemo/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// websocket配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func checkOrigin(r *http.Request) bool {
	return true
}

// 用户申请创建socket链接
func ConCreateConn(ctx *gin.Context) {
	var (
		conn    *websocket.Conn
		err     error
		user_id int
	)
	//	获取user_id  这里可以是token,经过中间件解析后的存在 ctx 的user_id
	//	为方便演示 这里直接请求头带user_id,正常开发不建议
	user_id, err = strconv.Atoi("666")
	if err != nil && user_id <= 0 {
		ctx.JSON(200, model.ResDatas(500, "请求必须带id"+err.Error(), nil))
		return
	}
	fmt.Println("user_id", user_id)
	//	判断请求过来的链接是否要升级为websocket
	if websocket.IsWebSocketUpgrade(ctx.Request) {
		//	将请求升级为 websocket链接
		conn, err = upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
		if err != nil {
			ctx.JSON(200, model.ResDatas(500, "创建链接失败"+err.Error(), nil))
			return
		}
	} else {
		return
	}
	//	用户断开销毁
	defer func() {
		_ = conn.Close()
	}()
	for {
		var msg model.ConnMsg
		//	ReadJSON 获取值的方式类似于gin的 ctx.ShouldBind() 通过结构体的json映射值
		//	如果读不到值 则堵塞在此处
		err = conn.ReadJSON(&msg)
		if err != nil {
			// 写回错误信息
			err = conn.WriteJSON(model.ResDatas(400, "获取数据错误："+err.Error(), nil))
			if err != nil {
				fmt.Println("用户断开")
				return
			}
		}
		// do something.....
		//user_id, _ = strconv.Atoi(msg.Useridx)
		//msg.T = user_id
		fmt.Println("msg = ", msg)

		//	发送回信息
		err = conn.WriteJSON(msg)
		if err != nil {
			fmt.Println("用户断开")
			return
		}
	}

}
