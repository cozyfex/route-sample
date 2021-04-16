package post

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	post := r.Group("/post")
	{
		post.GET("/1", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"post_id": 1,
			})
		})
	}
}
