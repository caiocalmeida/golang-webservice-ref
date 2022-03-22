package api

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()

	r.GET("/user", getUsers)
	r.GET("/user/:id", getUser)
	r.POST("/user", postUser)
	r.PUT("/user/:id", putUser)
	r.DELETE("/user/:id", deleteUser)

	r.Run()
}
