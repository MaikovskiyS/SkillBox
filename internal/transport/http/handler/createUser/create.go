package createHandler

import (
	"fmt"
	"net/http"
	"skillbox/internal/domain/model"
	"skillbox/internal/domain/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	UserModel *model.User
	Dto       *DTO
	router    *gin.Engine
	svc       service.UserService
	l         *logrus.Logger
}

func New(r *gin.Engine, svc service.UserService, l *logrus.Logger) *Handler {
	return &Handler{
		UserModel: &model.User{},
		Dto:       &DTO{},
		router:    r,
		svc:       svc,
		l:         l,
	}
}
func (h *Handler) User(c *gin.Context) {
	err := c.ShouldBindJSON(h.Dto)
	if err != nil {
		h.l.Info("create User parse params err:", err)
		c.String(http.StatusBadRequest, "create User parse params err:", err.Error())
		return
	}
	defer c.Request.Body.Close()
	err = h.Dto.Validate()
	if err != nil {
		h.l.Info("create User validation err:", err)
		c.String(http.StatusBadRequest, "create User validation err:", err.Error())
		return
	}
	err = h.ConvertDTOToModel()
	if err != nil {
		h.l.Info("create User convert err:", err)
		c.String(http.StatusBadRequest, "create User convert err:", err.Error())
		return
	}
	id, err := h.svc.CreateUser(c, h.UserModel)
	if err != nil {
		h.l.Info("create User err from service:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, fmt.Sprintf(`User created. ID:%v`, id))
}

func (h *Handler) ConvertDTOToModel() error {
	age, err := strconv.Atoi(h.Dto.Age)
	if err != nil {
		h.l.Info("CreateHandler: cant convert dto age", err)
		return err
	}
	h.UserModel.Name = h.Dto.Name
	h.UserModel.Age = uint64(age)
	return nil
}
