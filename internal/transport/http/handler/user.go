package handler

import (
	"context"
	"fmt"
	"net/http"
	"skillbox/internal/domain/model"
	"skillbox/internal/transport/http/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(ctx context.Context, data model.User) (uint64, error)
	MakeFriend(ctx context.Context, source, target uint64) (string, string, error)
	GetFriends(ctx context.Context, id uint64) ([]model.User, error)
	UpdateUserAge(ctx context.Context, id, age uint64) error
	DeleteUser(ctx context.Context, id uint64) error
}

//go:generate mockgen -source=user.go -destination=mocks/mock.go

type Handler struct {
	l      *logrus.Logger
	svc    UserService
	router *gin.Engine
}

func New(svc UserService, router *gin.Engine, logger *logrus.Logger) *Handler {
	router.Use(gin.Recovery())
	return &Handler{
		l:      logger,
		svc:    svc,
		router: router,
	}
}

// RegisterRoutes registrating routes.
func (h *Handler) RegisterRoutes() {
	h.l.Info("register Routes")
	h.router.POST("/user", (h.CreateUser))
	h.router.POST("/friends", (h.MakeFriend))
	h.router.DELETE("/user", (h.DeleteUser))
	h.router.GET("/friends", (h.GetFriends))
	h.router.PUT("/user_id", (h.UpdateUserAge))

}

// Create creating new user.
func (h *Handler) CreateUser(c *gin.Context) {
	var data dto.CreateUserDTO
	err := c.Bind(&data)
	if err != nil {
		c.String(http.StatusBadRequest, "cant parse params.Err:", err.Error())
		return
	}
	fmt.Println(data)
	defer c.Request.Body.Close()
	err = data.Validate()
	if err != nil {
		c.String(http.StatusBadRequest, "validation failed. Err:", err.Error())
		return
	}
	age, err := strconv.Atoi(data.Age)
	if err != nil {
		c.String(http.StatusBadRequest, "cant parse params", err.Error())
	}
	var user model.User
	user.Age = uint64(age)
	user.Name = data.Name
	id, err := h.svc.CreateUser(c, user)
	if err != nil {
		c.String(http.StatusBadRequest, "cant create User")
		return
	}
	c.JSON(http.StatusCreated, fmt.Sprintf("User created. Id:%v", id))
}

// MakeFriend make friendship between 2 users.
func (h *Handler) MakeFriend(c *gin.Context) {
	var data dto.MakeFriendDTO
	err := c.Bind(&data)
	if err != nil {
		c.String(http.StatusBadRequest, "cant parse params.Err:", err.Error())
		return
	}
	defer c.Request.Body.Close()
	err = data.Validate()
	if err != nil {
		c.String(http.StatusBadRequest, "validation failed. Err:", err.Error())
		return
	}
	target, err := strconv.Atoi(data.Target)
	if err != nil {
		c.String(http.StatusInternalServerError, "server error", err.Error())
	}
	source, err := strconv.Atoi(data.Source)
	if err != nil {
		c.String(http.StatusInternalServerError, "server error", err.Error())
	}
	targetName, sourceName, err := h.svc.MakeFriend(c, uint64(source), uint64(target))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Users %s and %s became friends", sourceName, targetName))
}

// Delete user.
func (h *Handler) DeleteUser(c *gin.Context) {
	var data dto.DeleteUserDto
	err := c.Bind(&data)
	if err != nil {
		c.String(http.StatusBadRequest, "cant parse params.Err:", err.Error())
		return
	}
	defer c.Request.Body.Close()
	err = data.Validate()
	if err != nil {
		c.String(http.StatusBadRequest, "validation failed. Err:", err.Error())
		return
	}
	id, err := strconv.Atoi(data.Id)
	if err != nil {
		c.String(http.StatusBadRequest, "validation failed. Err:", err.Error())
		return
	}
	err = h.svc.DeleteUser(c, uint64(id))
	if err != nil {
		switch err.Error() {
		case "user not found":
			c.String(http.StatusNotFound, err.Error())
			return
		default:
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

	}
	c.JSON(http.StatusOK, "user deleted")
}

// GetFriends of a user.
func (h *Handler) GetFriends(c *gin.Context) {
	var data dto.GetFriendsDto
	idstr := c.Query("id")
	defer c.Request.Body.Close()
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.String(http.StatusBadRequest, "ivalid params. Err:", err.Error())
		return
	}
	data.Id = uint64(id)
	err = data.Validate()
	if err != nil {
		c.String(http.StatusBadRequest, "validation failed. Err:", err.Error())
		return
	}
	friends, err := h.svc.GetFriends(c, uint64(id))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, friends)
}

// UpdateAge of a user.
func (h *Handler) UpdateUserAge(c *gin.Context) {
	var data dto.UpdateUserAgeDTO
	err := c.Bind(&data)
	if err != nil {
		c.String(http.StatusBadRequest, "ivalid params. Err:", err.Error())
		return
	}
	defer c.Request.Body.Close()
	idstr := c.Query("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.String(http.StatusBadRequest, "ivalid params. Err:", err.Error())
		return
	}

	data.Id = uint64(id)
	err = data.Validate()
	if err != nil {
		c.String(http.StatusBadRequest, "validation failed. Err:", err.Error())
		return
	}
	age, err := strconv.Atoi(data.Age)
	if err != nil {
		c.String(http.StatusBadRequest, "ivalid params. Err:", err.Error())
		return
	}

	err = h.svc.UpdateUserAge(c, uint64(id), uint64(age))
	if err != nil {
		c.String(http.StatusInternalServerError, "cant update age")
		return
	}
	c.JSON(http.StatusOK, "age changed")
}
