package routes

import (
	"login-system/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB) {
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running successfully"})
	})

	r.POST("/register", func(c *gin.Context) {
		user, err := controllers.RegisterUser(c, db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
	})
}
