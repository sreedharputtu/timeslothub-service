package main

import (
	"github.com/sreedharputtu/timeslothub-service/config"
	"github.com/sreedharputtu/timeslothub-service/handler"
	"github.com/sreedharputtu/timeslothub-service/repository"
	"github.com/sreedharputtu/timeslothub-service/router"
)

func main() {
	db, err := config.Database()
	if err != nil {
		//panic
	}
	userRepo := repository.NewUserRepository(db)
	rh := handler.NewRequestHandler(userRepo)
	router.NewRouter(rh).Run()
}
