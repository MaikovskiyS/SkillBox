package updateHandler

import (
	"net/http"
	"skillbox/internal/domain/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	userID uint64
	newAge uint64
	dto    *DTO
	router *gin.Engine
	svc    service.UserService
	l      *logrus.Logger
}

func New(r *gin.Engine, svc service.UserService, l *logrus.Logger) *Handler {
	return &Handler{
		dto:    &DTO{},
		router: r,
		svc:    svc,
		l:      l,
	}
}
func (h *Handler) User(c *gin.Context) {
	h.dto.Id = c.Query("id")
	err := c.ShouldBindJSON(h.dto)
	if err != nil {
		h.l.Info("updateAge parse params err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer c.Request.Body.Close()
	err = h.dto.Validate()
	if err != nil {
		h.l.Info("updateAge validation err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = h.ConvertDTOToModel()
	if err != nil {
		h.l.Info("updateAge convert err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = h.svc.UpdateUserAge(c, h.userID, h.newAge)
	if err != nil {
		h.l.Info("updateAge err from service:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, "Age updated.")
}

func (h *Handler) ConvertDTOToModel() error {
	age, err := strconv.Atoi(h.dto.Age)
	if err != nil {
		h.l.Info("CreateHandler: cant convert dto age", err)
		return err
	}
	id, err := strconv.Atoi(h.dto.Id)
	if err != nil {
		h.l.Info("CreateHandler: cant convert dto age", err)
		return err
	}
	h.userID = uint64(id)
	h.newAge = uint64(age)
	return nil
}
