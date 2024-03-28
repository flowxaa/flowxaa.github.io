package api

import (
	"github.com/gin-gonic/gin"
)

func Chat(ctx *gin.Context) {
	ctx.String(200, "hello world")
}
