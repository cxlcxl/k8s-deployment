package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/sysx"
)

func main() {
	r := gin.Default()

	r.GET("/hostname", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"hostname": sysx.Hostname()})
	})

	r.Run(":8080")
}
