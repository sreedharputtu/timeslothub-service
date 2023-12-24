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
	ur := repository.NewUserRepository(db)
	ssr := repository.NewSlotSettingsRepository(db)
	cri := repository.NewCalendarRepoImpl(db)
	rh := handler.NewRequestHandler(ur, ssr, cri)
	router.NewRouter(rh).Run()
}
