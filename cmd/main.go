package main

import (
	"log"
	"skillbox/internal/app"
	"skillbox/internal/config"
)

//войти в базу из контейнера
//psql -h localhost -p 5432 -U postgres -W
//Создать контейнер
//docker run --name mypostgres -p 5432:5432 -e POSTGRES_PASSWORD=Wild54323 -d postgres
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)

}
