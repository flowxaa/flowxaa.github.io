package route

import (
	"WebSocketDemo/api"
	"WebSocketDemo/middleware"

	"github.com/gin-gonic/gin"
)

func GetRoute() *gin.Engine {
	ro := gin.Default()
	ro.Use(middleware.Cors)
	ro.Any("ws", api.ConCreateConn)
	ro.Any("/", api.Chat)

	ro.Any("send", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})
	ro.LoadHTMLGlob("view/*")
	return ro
}
