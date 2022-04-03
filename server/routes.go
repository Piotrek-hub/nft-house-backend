package server

import (
	"github.com/gin-gonic/gin"
	"github.com/piotrek-hub/nft_house_backend/db"
)

func MainRouteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to API",
	})
}

func AddHouseRoute(c *gin.Context) {
	var house db.House
	c.Bind(&house)

	c.JSON(200, gin.H{
		"message": "all goood!",
	})
}

func GetHousesRoute(c *gin.Context) {
	houses := db.GetHouses()
	c.JSON(200, gin.H{
		"data": houses,
	})
}

func GetHouseRoute(c *gin.Context) {
	id := c.Param("id")
	houses := db.GetHouses(id)
	c.JSON(200, gin.H{
		"data": houses,
	})
}
