package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ctrl "github.com/iamargus95/restaurant_rest_api_design/controllers"
	"github.com/iamargus95/restaurant_rest_api_design/middlewares"
)

func StartGin() {
	r := gin.Default() // Init router
	r.GET("/", Welcome)
	r.NoRoute(NotFound)
	r.POST("/signup", ctrl.Signup)
	r.GET("/list-restaurants", ctrl.ListRestaurants).Use(middlewares.EnsureLoggedIn()) //Middlewares currently not working just for this route
	r.GET("/menu", ctrl.Menu).Use(middlewares.EnsureLoggedIn())

	log.Fatal(r.Run("localhost:8080"))
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Restaurant API",
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "route not found",
	})
}
