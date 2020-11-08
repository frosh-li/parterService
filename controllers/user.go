package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"parterService/utils"
)


type UserCenter struct {
}

type UserSession struct {
	Username string `json:"username"`
	Phone string `json:"phone"`
}

func (this *UserCenter) Login(c *gin.Context) {
	session := sessions.Default(c)
	userSession := UserSession {
		Username: "lijunliang",
		Phone: "15330066919",
	}
	jsonBytes, _ := json.Marshal(userSession)

	session.Set("user", string(jsonBytes))
	session.Save()
	fmt.Println("session", string(jsonBytes))
	response := utils.Response{c}
	response.Send(200, "", userSession)
	//c.JSON(200, gin.H{
	//	"message": "success",
	//	"data": userSession,
	//})
}

func (this *UserCenter) UserInfo(c *gin.Context) {
	session := sessions.Default(c)
	//var userinfo UserSession
	cUser := session.Get("user").(string)
	var userinfo UserSession
	json.Unmarshal([]byte(cUser), &userinfo)
	fmt.Println("userinfo", cUser)
	c.JSON(200, gin.H{
		"data": userinfo,
	})
}

func (this *UserCenter) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "success",
		"data": "from user logout",
	})
}
