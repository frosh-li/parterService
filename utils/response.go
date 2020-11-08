package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Context *gin.Context
}

func (this *Response) Send(status int,message string,  data interface{})  {
	this.Context.JSON(200, gin.H{
		"status": status,
		"message": message,
		"data": data,
	})
}