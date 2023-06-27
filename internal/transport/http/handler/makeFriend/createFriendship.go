package createFriendship

import (
	"fmt"
	"net/http"
	"skillbox/internal/domain/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	targetID uint64
	sourceID uint64
	dto      *DTO
	router   *gin.Engine
	svc      service.UserService
	l        *logrus.Logger
}

func New(r *gin.Engine, svc service.UserService, l *logrus.Logger) *Handler {
	return &Handler{
		dto:    &DTO{},
		router: r,
		svc:    svc,
		l:      l,
	}
}
func (h *Handler) Friendship(c *gin.Context) {
	err := c.ShouldBindJSON(h.dto)
	if err != nil {
		h.l.Info("friendship parse params err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	defer c.Request.Body.Close()
	err = h.dto.Validate()
	if err != nil {
		h.l.Info("friendship validation err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = h.ConvertDTOToModel()
	if err != nil {
		h.l.Info("friendship convert err:", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	targetName, sourceName, err := h.svc.MakeFriend(c, h.sourceID, h.targetID)
	if err != nil {
		h.l.Info("friendship err from service:", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Users %s and %s became friends", sourceName, targetName))
}

func (h *Handler) ConvertDTOToModel() error {

	targetid, err := strconv.Atoi(h.dto.Target)
	if err != nil {
		h.l.Info("FriendshipHandler: cant convert dto targetid", err)
		return err
	}
	sourceid, err := strconv.Atoi(h.dto.Source)
	if err != nil {
		h.l.Info("FriendshipHandler: cant convert dto sourceid", err)
		return err
	}
	h.targetID = uint64(targetid)
	h.sourceID = uint64(sourceid)
	return nil
}
