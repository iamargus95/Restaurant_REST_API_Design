package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamargus95/restaurant_rest_api_design/models"
)

func GetRestaurantIDFromRequest(c *gin.Context) models.Menu {
	// Initializing default
	restaurantid := 1
	sort := "menu_category asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "restaurantid":
			restaurantid, _ = strconv.Atoi(queryValue)
		}
	}
	return models.Menu{
		Restaurant_ID: restaurantid,
		Sort:          sort,
	}
}
