package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog_boke/dao"
	"go-blog_boke/model"
	"net/http"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println("username", username)
	fmt.Println("password", password)

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GoRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)
	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}
}
