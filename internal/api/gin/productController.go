package api

import (
	"net/http"

	"github.com/caiocalmeida/go-webservice-ref/internal/data"
	"github.com/caiocalmeida/go-webservice-ref/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController interface {
	getProducts(c *gin.Context)
	getProduct(c *gin.Context)
	postProduct(c *gin.Context)
	putProduct(c *gin.Context)
	deleteProduct(c *gin.Context)
}

type productController struct {
	pr data.ProductRepository
}

func NewProductController(pr data.ProductRepository) ProductController {
	return &productController{pr: pr}
}

// @Tags     Product
// @Produce  json
// @Success  200  {array}   ProductDtoOut  "OK"
// @Failure  404  {string}  string         "Not Found"
// @Failure  500  {string}  string         "Internal Error"
// @Router   /product [get]
func (pc *productController) getProducts(c *gin.Context) {
	products := pc.pr.GetProducts()

	productDtos := make([]*ProductDtoOut, len(products))
	for i, v := range products {
		productDtos[i] = ConvertProductToDto(v)
	}

	c.JSON(http.StatusOK, productDtos)
}

// @Tags     Product
// @Produce  json
// @Param    id   path      string         true  "Product UUID"
// @Success  200  {object}  ProductDtoOut  "OK"
// @Failure  400  {string}  string         "Bad Request"
// @Failure  404  {string}  string         "Not Found"
// @Failure  500  {string}  string         "Internal Error"
// @Router   /product/{id} [get]
func (pc *productController) getProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if p := pc.pr.GetProductBy(id); p != nil {
		c.JSON(http.StatusOK, ConvertProductToDto(p))
		return
	}

	c.Status(http.StatusNotFound)
}

// @Tags     Product
// @Accept   json
// @Produce  json
// @Param    productDTO  body      ProductDtoIn   true  "Product data"
// @Success  201         {object}  ProductDtoOut  "OK"
// @Failure  400         {string}  string         "Bad Request"
// @Failure  500         {string}  string         "Internal Error"
// @Router   /product [post]
func (pc *productController) postProduct(c *gin.Context) {
	productDto := &ProductDtoIn{}

	if err := c.ShouldBind(productDto); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newProduct := productDto.CreateProduct()

	addedProduct := pc.pr.AddProduct(newProduct)

	c.JSON(http.StatusCreated, ConvertProductToDto(addedProduct))
}

// @Tags     Product
// @Accept   json
// @Produce  json
// @Param    id           path      string         true  "Product UUID"
// @Param    productData  body      ProductDtoIn   true  "Product data"
// @Success  200          {object}  ProductDtoOut  "OK"
// @Failure  400          {string}  string         "Bad Request"
// @Failure  404          {string}  string         "Not Found"
// @Failure  500          {string}  string         "Internal Error"
// @Router   /product/{id} [put]
func (pc *productController) putProduct(c *gin.Context) {
	productDto := &ProductDtoIn{}
	if err := c.ShouldBind(&productDto); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product := pc.pr.UpdateProduct(productDto.ToProduct(id))

	c.JSON(http.StatusOK, ConvertProductToDto(product))
}

// @Tags     Product
// @Accept   json
// @Produce  json
// @Param    id   path      string  true  "Product UUID"
// @Success  200  {string}  string  "OK"
// @Failure  400  {string}  string  "Bad Request"
// @Failure  404  {string}  string  "Not Found"
// @Failure  500  {string}  string  "Internal Error"
// @Router   /product/{id} [delete]
func (pc *productController) deleteProduct(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if pc.pr.DeleteProduct(id) {
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusNotFound)
}

type ProductDtoIn struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (p *ProductDtoIn) CreateProduct() *domain.Product {
	return &domain.Product{Id: uuid.New(), Name: p.Name, Description: p.Description}
}

func (p *ProductDtoIn) ToProduct(id uuid.UUID) *domain.Product {
	return &domain.Product{Id: id, Name: p.Name, Description: p.Description}
}

type ProductDtoOut struct {
	Id          uuid.UUID `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

func ConvertProductToDto(p *domain.Product) *ProductDtoOut {
	return &ProductDtoOut{Id: p.Id, Name: p.Name, Description: p.Description}
}
