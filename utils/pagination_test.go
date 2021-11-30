package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGeneratePaginationFromRequest(t *testing.T) {

	var testData = []struct {
		urlWithParams string
		page          int
		per_page      int
	}{
		{
			urlWithParams: "/list-restaurants?page=2&per_page=10",
			page:          2,
			per_page:      10,
		},
		{
			// defaults to page 1
			urlWithParams: "list-restaurants?per_page=20",
			page:          1,
			per_page:      20,
		},
		{
			// defaults to 10 per_page
			urlWithParams: "list-restaurants?page=3",
			page:          3,
			per_page:      10,
		},
	}

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	for _, test := range testData {
		c.Request, _ = http.NewRequest(http.MethodPost, test.urlWithParams, nil)
		c.Request.Header.Set("Content-Type", "application/json")

		got := GeneratePaginationFromRequest(c)
		fmt.Println(got)
		if got.Per_Page != test.per_page && got.Page != test.page {
			t.Fail()
		}
	}

}
