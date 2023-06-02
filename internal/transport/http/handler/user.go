package handler

import (
	"context"
	"fmt"
	"net/http"
	"skillbox/internal/domain/model"
	"skillbox/internal/transport/http/middleware"

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
	mw     *middleware.Middleware
	svc    UserService
	router *gin.Engine
}

func New(svc UserService, router *gin.Engine, mw *middleware.Middleware, logger *logrus.Logger) *Handler {
	router.Use(gin.Recovery())
	return &Handler{
		l:      logger,
		mw:     mw,
		svc:    svc,
		router: router,
	}
}

//RegisterRoutes...
func (h *Handler) RegisterRoutes() {
	h.l.Info("register Routes")
	h.router.GET("/", h.Hello)
	h.router.POST("/create", h.mw.CreateUserValidation(h.CreateUser))
	h.router.POST("/make_friends", h.mw.MakeFriendValidation(h.MakeFriend))
	h.router.DELETE("/user", h.mw.DeleteUserValidation(h.DeleteUser))
	h.router.GET("/friends", h.mw.GetFriendsValidation(h.GetFriends))
	h.router.PUT("/user_id", h.mw.UpdateUserAgeValidation(h.UpdateUserAge))

}

//Hello...
func (h *Handler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello from app on 8083 port")
}

//CreateUser...
func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User
	name := c.Keys["name"]
	age := c.Keys["age"]
	user.Name = name.(string)
	user.Age = age.(uint64)
	defer c.Request.Body.Close()
	id, err := h.svc.CreateUser(c, user)
	if err != nil {
		c.String(http.StatusBadRequest, "cant create User")
		return
	}
	c.JSON(http.StatusCreated, fmt.Sprintf("User created. Id:%v", id))
}

//MakeFriend...
func (h *Handler) MakeFriend(c *gin.Context) {
	value1 := c.Keys["target"]
	value2 := c.Keys["source"]
	target := value1.(uint64)
	source := value2.(uint64)
	defer c.Request.Body.Close()
	targetName, sourceName, err := h.svc.MakeFriend(c, source, target)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Users %s and %s became friends", sourceName, targetName))
}

//DeleteUser...
func (h *Handler) DeleteUser(c *gin.Context) {
	value := c.Keys["id"]
	id, ok := value.(uint64)
	if !ok {
		return
	}
	defer c.Request.Body.Close()
	err := h.svc.DeleteUser(c, id)
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

//GetFriends...
func (h *Handler) GetFriends(c *gin.Context) {
	value := c.Keys["id"]
	id := value.(uint64)
	defer c.Request.Body.Close()
	friends, err := h.svc.GetFriends(c, id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, friends)
}

//UpdateAge...
func (h *Handler) UpdateUserAge(c *gin.Context) {
	value1 := c.Keys["id"]
	value2 := c.Keys["age"]
	id := value1.(uint64)
	age := value2.(uint64)
	defer c.Request.Body.Close()
	err := h.svc.UpdateUserAge(c, id, age)
	if err != nil {
		c.String(http.StatusInternalServerError, "cant update age")
		return
	}
	c.JSON(http.StatusOK, "age changed")
}

// func validateUser(c *gin.Context) {
// 	var u User
// 	if err := c.ShouldBindJSON(&u); err == nil {
// 		c.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "User validation failed!",
// 			"error":   err.Error(),
// 		})
// 	}
