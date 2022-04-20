package api

import (
	docs "github.com/caiocalmeida/go-webservice-ref/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router interface {
	Start()
}

type router struct {
	uc UserController
}

func NewRouter(uc UserController) Router {
	return &router{uc: uc}
}

// @Title         Go Web API
// @Version       0.1
// @License.name  MIT
// @License.url   https://opensource.org/licenses/MIT
func (r *router) Start() {
	e := gin.Default()

	docs.SwaggerInfo.BasePath = ""

	e.GET("/user", r.uc.getUsers)
	e.GET("/user/:id", r.uc.getUser)
	e.POST("/user", r.uc.postUser)
	e.PUT("/user/:id", r.uc.putUser)
	e.DELETE("/user/:id", r.uc.deleteUser)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	e.Run()
}
