package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

var testWelcome = struct {
	url          string
	expectedCode int
	responseData []byte
}{
	url:          "/",
	expectedCode: http.StatusOK,
	responseData: []byte(`"message": "Welcome to the Restaurant API"`),
}

func TestWelcome(t *testing.T) {

	gin.SetMode(gin.TestMode)
	asserts := assert.New(t)
	r := gin.New()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, testWelcome.url, nil)

	Welcome(c)

	r.ServeHTTP(w, c.Request)
	asserts.Equal(testWelcome.expectedCode, w.Code)
}

var testNotFound = struct {
	url          string
	expectedCode int
	responseData []byte
}{
	url:          "/nonexistent",
	expectedCode: http.StatusNotFound,
	responseData: []byte(`"message": "route not found"`),
}

func TestNotFound(t *testing.T) {

	gin.SetMode(gin.TestMode)
	asserts := assert.New(t)
	r := gin.New()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, testNotFound.url, nil)

	NotFound(c)

	r.ServeHTTP(w, c.Request)
	asserts.Equal(testNotFound.expectedCode, w.Code)
}
