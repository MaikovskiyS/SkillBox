package handler

import (
	"skillbox/internal/domain/service"
	createHandler "skillbox/internal/transport/http/handler/createUser"
	deleteHandler "skillbox/internal/transport/http/handler/deleteUser"
	getFriendsHandler "skillbox/internal/transport/http/handler/getFriends"
	createFriendship "skillbox/internal/transport/http/handler/makeFriend"
	updateHandler "skillbox/internal/transport/http/handler/updateUserAge"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Engine        *gin.Engine
	Creator       createHandler.Handler
	Deletor       deleteHandler.Handler
	Updator       updateHandler.Handler
	Friendcreator createFriendship.Handler
	Getter        getFriendsHandler.Handler
}

func New(svc service.UserService, l *logrus.Logger) *Router {
	engine := gin.New()

	return &Router{
		Engine:        engine,
		Creator:       *createHandler.New(engine, svc, l),
		Deletor:       *deleteHandler.New(engine, svc, l),
		Updator:       *updateHandler.New(engine, svc, l),
		Friendcreator: *createFriendship.New(engine, svc, l),
		Getter:        *getFriendsHandler.New(engine, svc, l),
	}

}
func (r *Router) RegisterRoutes() {
	r.Engine.POST("/user", r.Creator.User)
	r.Engine.DELETE("/user", r.Deletor.User)
	r.Engine.PUT("/user", r.Updator.User)
	r.Engine.POST("/friends", r.Friendcreator.Friendship)
	r.Engine.GET("/friends", r.Getter.Friends)

}
