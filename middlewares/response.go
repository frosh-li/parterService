package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GinBodyLogMiddleware(c *gin.Context) {
	fmt.Println("logger middlewares ginBodyLogMiddleware")
}