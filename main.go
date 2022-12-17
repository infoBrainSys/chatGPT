package main

import (
	reqapi "Learn/reqApi"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {

		a, err := reqapi.ReqApi()
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": a,
		})
	})

	app.Run(":80")
}
