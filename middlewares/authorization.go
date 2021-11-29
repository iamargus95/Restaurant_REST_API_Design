package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamargus95/restaurant_rest_api_design/dbconn"
	"github.com/iamargus95/restaurant_rest_api_design/models"
)

func EnsureLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		validateToken(ctx)
	}
}

func validateToken(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("X-Auth-Token")
	fmt.Println(auth)

	if auth == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "no X-Auth-Token provided",
		})
		ctx.Abort()
		return
	}

	if len(auth) != 36 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid X-Auth-Token provided",
		})
		ctx.Abort()
		return
	}

	var user models.Users
	db := dbconn.GetDB()
	result := db.First(&user, "guid = ?", auth)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid X-Auth-Token Provided",
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
	ctx.Next()
}
