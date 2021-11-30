package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetRestaurantIDFromRequest(t *testing.T) {

	var testData = []struct {
		urlWithParams string
		restaurant_id int
	}{
		{
			urlWithParams: "/menu?restaurantid=2",
			restaurant_id: 2,
		},
		{
			// defaults to restaurantid 1
			urlWithParams: "/menu",
			restaurant_id: 1,
		},
		{
			urlWithParams: "/menu?restaurantid=15",
			restaurant_id: 15,
		},
	}

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	for _, test := range testData {
		c.Request, _ = http.NewRequest(http.MethodPost, test.urlWithParams, nil)
		c.Request.Header.Set("Content-Type", "application/json")

		got := GetRestaurantIDFromRequest(c)
		if got.Restaurant_ID != test.restaurant_id {
			t.Fail()
		}
	}
}
