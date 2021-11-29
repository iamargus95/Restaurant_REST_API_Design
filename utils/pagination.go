package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iamargus95/restaurant_rest_api_design/models"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	// Initializing default
	//	var mode string
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
