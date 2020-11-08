package routers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	myConfig "parterService/config"
	"parterService/controllers"
)

func InitRouter() {
	r := gin.Default()
	store, redisError := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	fmt.Println("redis error")
	fmt.Println(redisError)
	var UserCenter controllers.UserCenter
	r.Use(sessions.Sessions("BwPartner", store))
	//r.Use(middlewares.GinBodyLogMiddleware)
	ApiService := r.Group("/v1/api")
	{
		UserCenterService := ApiService.Group("/user")
		{
			UserCenterService.POST("/login", UserCenter.Login)
			UserCenterService.POST("/logout", UserCenter.Logout)
			UserCenterService.GET("/userinfo", UserCenter.UserInfo)
		}
	}

	r.Run(myConfig.GetString("server.address")) // listen and serve on 0.0.0.0:8080
}
