package main

import (
	"github.com/gin-gonic/gin"
	handler "github.com/sreedharputtu/timeslothub-service/internal"
)

func main() {
	r := gin.Default()
	r.GET("/health", handler.HealthHandler)
	r.Run()
}
