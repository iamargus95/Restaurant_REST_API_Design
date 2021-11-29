package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	conn "github.com/iamargus95/restaurant_rest_api_design/dbconn"
	r "github.com/iamargus95/restaurant_rest_api_design/resources"
	s "github.com/iamargus95/restaurant_rest_api_design/services"
)

func Signup(ctx *gin.Context) {

	var body r.SignupPayload
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	dbconn := conn.GetDB()
	userUUID, serviceErr := s.Signup(dbconn, body)
	if serviceErr != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": serviceErr.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"GUID": userUUID,
	})
}
