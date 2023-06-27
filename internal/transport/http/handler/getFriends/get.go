package getFriendsHandler

import (
	"net/http"
	"skillbox/internal/domain/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	userId uint64
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
func (h *Handler) Friends(c *gin.Context) {
	h.dto.Id = c.Query("id")
	err := h.dto.Validate()
	if err != nil {
		h.l.Info("getfirends validation err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = h.ConvertDTOToModel()
	if err != nil {
		h.l.Info("Get friends convert err:", err)
		c.String(http.StatusBadRequest, "create User convert err:", err.Error())
		return
	}
	friends, err := h.svc.GetFriends(c, h.userId)
	if err != nil {
		h.l.Info("create User err from service:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, friends)
}
func (h *Handler) ConvertDTOToModel() error {
	id, err := strconv.Atoi(h.dto.Id)
	if err != nil {
		h.l.Info("GetFriends convert err", err)
		return err
	}
	h.userId = uint64(id)
	return nil
}
