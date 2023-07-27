package main

import (
	"database/sql"
	"fmt"
	"iot_devices/pkg/api"
	"iot_devices/pkg/config"
	"iot_devices/pkg/repository"
	"iot_devices/pkg/service"
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	cfg, err := config.LoadConfigurations()
	if err != nil {
		log.Println(err)
	}
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tempRepo := repository.NewTemperatureRepository(db)
	tempService := service.NewTemperatureService(tempRepo)
	velRepo := repository.NewVelocityRepository(db)
	velService := service.NewVelocityService(velRepo)
	trafficRepo := repository.NewTrafficRepository(db)
	trafficService := service.NewTrafficService(trafficRepo)
	handler := api.NewHandler(tempService, velService, trafficService)
	fmt.Println("running")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
