package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog_boke/controller"
)

func Start() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	//r.GET("/index", controller.ListUser)

	r.GET("/login", controller.GoLogin)
	r.POST("/login", controller.Login)

	r.GET("/", controller.Index)

	r.GET("/register", controller.GoRegister)
	r.POST("/register", controller.Register)

	r.GET("/post_index", controller.GoPostIndex)
	r.POST("/post", controller.AddPost)
	r.GET("/post", controller.GoAddPost)
	r.GET("/post_index/detail", controller.PostDetail)

	if err := r.Run(); err != nil {
		fmt.Println(err)
	}
}
