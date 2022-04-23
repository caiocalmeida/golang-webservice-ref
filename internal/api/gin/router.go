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
	pc ProductController
}

func NewRouter(pc ProductController) Router {
	return &router{pc: pc}
}

// @Title         Go Web API
// @Version       0.1
// @License.name  MIT
// @License.url   https://opensource.org/licenses/MIT
func (r *router) Start() {
	e := gin.Default()

	docs.SwaggerInfo.BasePath = ""

	e.GET("/product", r.pc.getProducts)
	e.GET("/product/:id", r.pc.getProduct)
	e.POST("/product", r.pc.postProduct)
	e.PUT("/product/:id", r.pc.putProduct)
	e.DELETE("/product/:id", r.pc.deleteProduct)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	e.Run()
}
