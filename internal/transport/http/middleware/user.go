package middleware

import (
	"fmt"
	"net/http"
	"skillbox/internal/transport/http/dto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	logger *logrus.Logger
}

func New(l *logrus.Logger) *Middleware {
	return &Middleware{
		logger: l,
	}
}

//DeleteUserValidation...
func (m *Middleware) DeleteUserValidation(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		m.logger.Info("DeleteUser middleware")
		var data dto.DeleteUserDto
		err := c.Bind(&data)
		if err != nil {
			c.String(http.StatusBadRequest, "cant parse params.Err:", err.Error())
			return
		}
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
		fmt.Println("ID:", id)
		c.Keys["id"] = uint64(id)
		next(c)
	}

}

//CreateUserValidation...
func (m *Middleware) CreateUserValidation(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		var data dto.CreateUserDTO
		err := c.Bind(&data)
		if err != nil {
			c.String(http.StatusBadRequest, "cant parse params. Err:", err.Error())
			return
		}
		err = data.Validate()
		if err != nil {
			c.String(http.StatusBadRequest, "validation failed. Err:", err.Error())
			return
		}
		age, err := strconv.Atoi(data.Age)
		if err != nil {
			c.String(http.StatusBadRequest, "cant convert params. Err:", err.Error())
			return
		}
		c.Keys["name"] = data.Name
		c.Keys["age"] = uint64(age)
		next(c)
	}
}

//MakeFriendValidation...
func (m *Middleware) MakeFriendValidation(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		var data dto.MakeFriendDTO
		err := c.Bind(&data)
		if err != nil {
			c.String(http.StatusBadRequest, "ivalid params. Err:", err.Error())
			return
		}
		err = data.Validate()
		if err != nil {
			c.String(http.StatusBadRequest, "validation failed. Err:", err.Error())
			return
		}
		targetId, err := strconv.Atoi(data.Target)
		if err != nil {
			c.String(http.StatusBadRequest, "ivalid params. Err:", err.Error())
			return
		}
		sourceId, err := strconv.Atoi(data.Source)
		if err != nil {
			c.String(http.StatusBadRequest, "ivalid params. Err:", err.Error())
			return
		}
		c.Keys["target"] = uint64(targetId)
		c.Keys["source"] = uint64(sourceId)
		next(c)
	}
}

//UpdateUserAgeValidation...
func (m *Middleware) UpdateUserAgeValidation(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		var data dto.UpdateUserAgeDTO
		err := c.Bind(&data)
		if err != nil {
			c.String(http.StatusBadRequest, "ivalid params. Err:", err.Error())
			return
		}
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
		c.Keys["id"] = uint64(id)
		c.Keys["age"] = uint64(age)
		next(c)
	}
}

//GetFriendsValidation...
func (m *Middleware) GetFriendsValidation(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		var data dto.GetFriendsDto
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
		c.Keys["id"] = uint64(id)
		next(c)
	}
}
