package server

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", MainRouteHandler)

	api := r.Group("/api")
	{
		api.POST("/", AddHouseRoute)
		api.GET("/", GetHousesRoute)
		api.GET("/:id", GetHouseRoute)
	}

	r.Run()
}
