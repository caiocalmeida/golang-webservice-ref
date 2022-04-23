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

// @securityDefinitions.apikey  API Key
// @in                          header
// @name                        X-API-KEY
func (r *router) Start() {
	e := gin.Default()

	docs.SwaggerInfo.BasePath = ""

	apiKeyRequired := e.Group("/", requireApiKey())
	apiKeyRequired.GET("/product", r.pc.getProducts)
	apiKeyRequired.GET("/product/:id", r.pc.getProduct)

	forbidden := e.Group("/", forbidRequest())
	forbidden.POST("/product", r.pc.postProduct)
	forbidden.PUT("/product/:id", r.pc.putProduct)
	forbidden.DELETE("/product/:id", r.pc.deleteProduct)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	e.Run()
}
