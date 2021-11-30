package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	r "github.com/iamargus95/restaurant_rest_api_design/resources"
	s "github.com/iamargus95/restaurant_rest_api_design/services"
	"github.com/iamargus95/restaurant_rest_api_design/utils"
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

	userUUID, serviceErr := s.Signup(body)
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

func ListRestaurants(ctx *gin.Context) {

	pagination := utils.GeneratePaginationFromRequest(ctx)
	content, err := s.ListRestaurants(&pagination)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list of restaurants": content,
	})
}

func Menu(ctx *gin.Context) {

	restaurantDetails := utils.GetRestaurantIDFromRequest(ctx)
	content, err := s.GetMenuItems(&restaurantDetails)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Menu Items": content,
	})
}
