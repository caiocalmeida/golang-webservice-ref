package api

import (
	"net/http"

	"github.com/caiocalmeida/go-webservice-ref/internal/data"
	"github.com/caiocalmeida/go-webservice-ref/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, data.GetUsers())
}

func getUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if u := data.GetUserBy(id); u != nil {
		c.JSON(http.StatusOK, u)
		return
	}

	c.Status(http.StatusNotFound)
}

func postUser(c *gin.Context) {
	userDto := &UserDto{}

	if err := c.ShouldBind(userDto); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newUser := userDto.CreateUser()

	addedUser := data.AddUser(newUser)

	c.JSON(http.StatusCreated, addedUser)
}

func putUser(c *gin.Context) {
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

	c.JSON(http.StatusOK, data.UpdateUser(userDto.ToUser(id)))
}

func deleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if data.DeleteUser(id) {
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
