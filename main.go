package main

import (
	"github.com/gin-gonic/gin"
	"route-sample/post"
	"route-sample/schedule"
	"route-sample/user"
)

func main() {
	r := gin.Default()

	post.Routes(r)
	user.Routes(r)
	schedule.Routes(r)

	err := r.Run()
	if err != nil {
		return
	}
}
