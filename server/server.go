package server

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()
	router.GET("/", MainRouteHandler)

	// API
	router.POST("/api", AddHouseRoute)
	router.GET("/api", GetHousesRoute)
	router.GET("/api/:id", GetHouseRoute)
	router.Run()
}