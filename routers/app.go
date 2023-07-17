package app

import (
	controller "github.com/abhinavramkumar/go-rss/controllers"
	middleware_auth "github.com/abhinavramkumar/go-rss/middlewares"
	"github.com/gin-gonic/gin"
)

func AppRouter(rg *gin.RouterGroup) {
	rg.GET("/ping", controller.PingController)
	rg.POST("/login", controller.LoginController)

	rg.POST("/createUser", controller.CreateUserController)
	rg.Use(middleware_auth.AuthMiddleware())
	rg.GET("/suggestedFeeds", controller.GetSuggestedFeeds)
	rg.GET("/getFeed", controller.GetFeed)
	rg.POST("/uploadOPML", controller.UploadOPML)
}
