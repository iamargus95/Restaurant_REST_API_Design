package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamargus95/restaurant_rest_api_design/models"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	// Initializing default
	//	var mode string
	per_page := 10
	page := 1
	sort := "star_rating desc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "per_page":
			per_page, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		}
	}
	return models.Pagination{
		Per_Page: per_page,
		Page:     page,
		Sort:     sort,
	}
}
