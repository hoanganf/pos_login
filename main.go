package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hoanganf/pos_login/src"
)

func main() {
	//connect to DB
	bean, err := src.InitBean()
	defer bean.DestroyBean()
	if err != nil {
		log.Fatalln("can not create bean", err)
	}
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/login", bean.LoginService.GetLogin)
	router.GET("/logout", bean.LoginService.GetLogout)
	router.POST("/login", bean.LoginService.Post)

	v1 := router.Group("v1")
	{
		v1.POST("/user", bean.UserService.Login)
	}
	router.Run(":8081")
}
