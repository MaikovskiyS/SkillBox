package deleteHandler

import (
	"net/http"
	"skillbox/internal/domain/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	userID uint64
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
	err := c.ShouldBindJSON(h.dto)
	if err != nil {
		h.l.Info("delete User parse params err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer c.Request.Body.Close()
	err = h.dto.Validate()
	if err != nil {
		h.l.Info("delete User validation err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = h.ConvertDTOToModel()
	if err != nil {
		h.l.Info("delete User convert err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = h.svc.DeleteUser(c, h.userID)
	if err != nil {
		h.l.Info("create User err from service:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "user deleted")

}
func (h *Handler) ConvertDTOToModel() error {
	id, err := strconv.Atoi(h.dto.Id)
	if err != nil {
		h.l.Info("DeleteHandler: cant convert dto id", err)
		return err
	}
	h.userID = uint64(id)
	return nil
}
