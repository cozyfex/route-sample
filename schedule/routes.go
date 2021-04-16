package schedule

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	schedule := r.Group("/schedule")
	{
		schedule.GET("/day", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"day": "today",
			})
		})
	}
}
