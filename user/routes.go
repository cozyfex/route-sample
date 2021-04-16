package user

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"username": "User ID",
			})
		})
	}
}
