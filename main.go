package main

import (
	"fmt"
	"log"
	"net/http"

	approuter "github.com/abhinavramkumar/go-rss/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Starting up go server")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	initializeRoutes(r)
	r.Run(":3333")
}

func initializeRouter(group *gin.RouterGroup, f func(*gin.RouterGroup)) {
	f(group)
}

func initializeRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		app := v1.Group("/app")
		initializeRouter(app, approuter.AppRouter)
	}
}
