package routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"outstagram/server/configs"
	"outstagram/server/middlewares"
)

func PostAPIRouter(router *gin.RouterGroup) {
	postController, err := configs.InitializePostController()
	if err != nil {
		log.Fatal(err.Error())
	}

	router.Use(middlewares.VerifyToken)

	router.GET("/", postController.GetPosts)
	router.POST("/", postController.CreatePost)

	router.GET("/:postID", postController.GetPost)

	router.POST("/:postID/views", postController.ViewPost)
}