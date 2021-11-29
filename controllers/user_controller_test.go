package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/iamargus95/restaurant_rest_api_design/dbconn"
	"github.com/iamargus95/restaurant_rest_api_design/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

var DB *gorm.DB

var signupRequestTests = []struct {
	bodyData     []byte
	expectedCode int
}{
	{
		// OK
		bodyData:     []byte(`{"name": "testClient","email": "testing@test.in","phone": "9999999999"}`),
		expectedCode: 202,
	},
	{
		//invalid Email
		bodyData:     []byte(`{"name": "testClient","email": "testing@test","phone": "9999999999"}`),
		expectedCode: 400,
	},
	{
		//invalid PhoneNumber (length!=10)
		bodyData:     []byte(`{"name": "testClient","email": "testing@test","phone": "999999999"}`),
		expectedCode: 400,
	},
}

func TestSignup(t *testing.T) {

	asserts := assert.New(t)
	gin.SetMode(gin.TestMode)
	r := gin.New()

	for _, testdata := range signupRequestTests {

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		requestData := testdata.bodyData

		c.Request, _ = http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(requestData))
		c.Request.Header.Set("Content-Type", "application/json")

		Signup(c)

		r.ServeHTTP(w, c.Request)

		asserts.Equal(testdata.expectedCode, w.Code)
		//cleanup testdata from db
		cleanup()
	}
}

func cleanup() {
	var U models.Users
	db := dbconn.GetDB()
	db.Where("name = ?", "testClient").Delete(&U)
	db.Unscoped().Where("name = ?", "testClient").Delete(&U)
}
