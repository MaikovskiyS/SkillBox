package app

import (
	"skillbox/internal/adapters/storage/postgres"
	"skillbox/internal/config"
	"skillbox/internal/domain/service"
	"skillbox/internal/transport/http/handler"
	"skillbox/internal/transport/http/server"

	"github.com/sirupsen/logrus"
)

//TODO Gracefull
func Run(cfg *config.Config) {
	l := logrus.New()
	//router := gin.Default()
	client := postgres.NewClient(l, cfg)
	err := client.Connect()
	if err != nil {
		l.Fatal(err, "cant connect to DB")
	}
	defer client.Close()
	userRepo := postgres.NewUserRepository(client)
	userSvc := service.NewUserService(userRepo, l)
	router := handler.New(userSvc, l)
	httpServer := server.New(router)
	router.RegisterRoutes()
	httpServer.Start()

}
