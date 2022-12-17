package main

import (
	reqapi "chatGPT/reqApi"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {

		// 获取传入的 Query 字段并拼接
		prompt := ctx.Query("prompt")

		// 调用 api
		a, err := reqapi.ReqApi(prompt)
		if err != nil {
			return
		}

		// 将 openAI 回复的内容写到 “/” 路由上
		ctx.Writer.Write([]byte(a))
	})

	app.Run(":80")
}
