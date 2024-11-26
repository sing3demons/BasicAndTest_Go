package api

import (
	"app/config"
	"app/models"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(r *gin.Engine) {
	authenAPI := r.Group("api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(401, gin.H{"result": "Login"})
}

func register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(401, gin.H{"status": "unable to bind data"})
	} 
		// user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		if err := config.GetDB().Create(&user).Error; err != nil {
			c.JSON(200, gin.H{"result": "nok", "error": err})
		} else {
			c.JSON(200, gin.H{"result": "ok", "data": user})
		}
	
}
