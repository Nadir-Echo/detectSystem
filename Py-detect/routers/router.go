package routers

import (
	"github.com/gin-gonic/gin"

	"Demo/pkg/setting"
	"Demo/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	gin.ForceConsoleColor()
	gin.SetMode(setting.RunMode)

	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})

	apiv1 := r.Group("api/v1")
	{
		//获取对应生产线的检测结果列表
		apiv1.GET("/results", v1.GetResult)
		//添加新的检测结果
		apiv1.POST("/results", v1.AddResult)
	}

	return r
}
