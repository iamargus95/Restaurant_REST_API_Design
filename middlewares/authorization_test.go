package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestEnsureLoggedIn(t *testing.T) {

	var testCases = []struct {
		name          string
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			r := gin.New()
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			token := "beabdff8-5101-11ec-9b40-38f3abdee1f2"

			c.Request.Header.Set("X-Auth_Token", token)
			authPath := "/list-restautants"
			c.Request, _ = http.NewRequest(http.MethodGet, authPath, nil)

			r.GET(
				authPath,
				EnsureLoggedIn(),
			)

			r.ServeHTTP(w, c.Request)
			tc.checkResponse(t, w)
		})
	}

}
