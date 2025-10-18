package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		fmt.Println("Start Serving Request ...")
		fmt.Println("Request URL:", c.Request.URL.Path)
		c.Next()
		duration := time.Since(start)
		fmt.Println("Request Duration:", duration)
		fmt.Println("Done Serving Request!")
	}
}
