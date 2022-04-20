package api

import (
	"net/http"

	"github.com/caiocalmeida/go-webservice-ref/internal/data"
	"github.com/caiocalmeida/go-webservice-ref/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	getUsers(c *gin.Context)
	getUser(c *gin.Context)
	postUser(c *gin.Context)
	putUser(c *gin.Context)
	deleteUser(c *gin.Context)
}

type userController struct {
	ur data.UserRepository
}

func NewUserController(ur data.UserRepository) UserController {
	return &userController{ur: ur}
}

// @Tags     User
// @Produce  json
// @Success  200  {array}   UserDto  "OK"
// @Failure  404  {string}  string   "Not Found"
// @Failure  500  {string}  string   "Internal Error"
// @Router   /user [get]
func (uc *userController) getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, uc.ur.GetUsers())
}

// @Tags     User
// @Produce  json
// @Param    id   path      string   true  "User UUID"
// @Success  200  {object}  UserDto  "OK"
// @Failure  400  {string}  string   "Bad Request"
// @Failure  404  {string}  string   "Not Found"
// @Failure  500  {string}  string   "Internal Error"
// @Router   /user/{id} [get]
func (uc *userController) getUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if u := uc.ur.GetUserBy(id); u != nil {
		c.JSON(http.StatusOK, u)
		return
	}

	c.Status(http.StatusNotFound)
}

// @Tags     User
// @Accept   json
// @Produce  json
// @Param    userDTO  body      UserDto  true  "User data"
// @Success  200      {object}  UserDto  "OK"
// @Failure  400      {string}  string   "Bad Request"
// @Failure  500      {string}  string   "Internal Error"
// @Router   /user [post]
func (uc *userController) postUser(c *gin.Context) {
	userDto := &UserDto{}

	if err := c.ShouldBind(userDto); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newUser := userDto.CreateUser()

	addedUser := uc.ur.AddUser(newUser)

	c.JSON(http.StatusCreated, addedUser)
}

// @Tags     User
// @Accept   json
// @Produce  json
// @Param    id        path      string   true  "User UUID"
// @Param    userData  body      UserDto  true  "User data"
// @Success  200       {object}  UserDto  "OK"
// @Failure  400       {string}  string   "Bad Request"
// @Failure  404       {string}  string   "Not Found"
// @Failure  500       {string}  string   "Internal Error"
// @Router   /user/{id} [put]
func (uc *userController) putUser(c *gin.Context) {
	userDto := &UserDto{}
	if err := c.ShouldBind(&userDto); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, uc.ur.UpdateUser(userDto.ToUser(id)))
}

// @Tags     User
// @Accept   json
// @Produce  json
// @Param    id   path      string   true  "User UUID"
// @Success  200  {object}  UserDto  "OK"
// @Failure  400  {string}  string   "Bad Request"
// @Failure  404  {string}  string   "Not Found"
// @Failure  500  {string}  string   "Internal Error"
// @Router   /user/{id} [delete]
func (uc *userController) deleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if uc.ur.DeleteUser(id) {
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusNotFound)
}

type UserDto struct {
	Name string `json:"name" binding:"required"`
}

func (u *UserDto) CreateUser() *domain.User {
	return &domain.User{Id: uuid.New(), Name: u.Name}
}

func (u *UserDto) ToUser(id uuid.UUID) *domain.User {
	return &domain.User{Id: id, Name: u.Name}
}
