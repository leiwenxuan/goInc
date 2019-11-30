package router

import (
	"rds/goinception/v1/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// 健康检查节点
	//r.GET("/v1/check")

	// v1 版本
	groupV1 := r.Group("/v1/rds")
	{

		groupV1.POST("/inception", handler.Inception)
	}
	return
}
