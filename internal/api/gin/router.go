package api

import "github.com/gin-gonic/gin"

type Router interface {
	Start()
}

type router struct {
	uc UserController
}

func NewRouter(uc UserController) Router {
	return &router{uc: uc}
}

func (r *router) Start() {
	e := gin.Default()

	e.GET("/user", r.uc.getUsers)
	e.GET("/user/:id", r.uc.getUser)
	e.POST("/user", r.uc.postUser)
	e.PUT("/user/:id", r.uc.putUser)
	e.DELETE("/user/:id", r.uc.deleteUser)

	e.Run()
}
