package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sreedharputtu/timeslothub-service/internal/handler"
)

func main() {
	r := gin.Default()
	r.GET("/health", handler.HealthHandler)
	r.Run()
}
