package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Demo/docs"
	"Demo/middleware"

	"Demo/pkg/setting"
	"Demo/pkg/util"
	"Demo/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// gin.ForceConsoleColor()
	// 设置CORS头
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:9528"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Token"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//r.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	//登录并获取token
	r.POST("user/login", api.Login)
	// r.POST("users/login", api.UserLogin)
	//获取对应的角色信息
	r.GET("user/info", api.Getinfo)

	//logout
	r.POST("user/logout", api.Logout)

	//use middleware
	
	//detect
	// r.POST("defectimg", util.Defectimg)
	r.GET("auth/detector",util.Auth_detector)
	r.GET("auth/admin",util.Auth_admin)
	r.GET("data/form",api.GetDataForm)
	r.GET("data/get-img",api.GetDataImg)
	r.GET("data/chart",api.GetChartData)

	r.GET("manage/getUserInfo",api.GetUserInfo)

	r.POST("defectimg",middleware.JWT_detector(), util.Defectimg)

	r.POST("/manage/addUser",api.AddNewUser)

	return r
}
