package services

import (
	"github.com/google/uuid"
	conn "github.com/iamargus95/restaurant_rest_api_design/dbconn"
	"github.com/iamargus95/restaurant_rest_api_design/models"
	r "github.com/iamargus95/restaurant_rest_api_design/resources"
)

func Signup(body r.SignupPayload) (uuid.UUID, error) {

	var nilUUID uuid.UUID
	var newUser models.Users

	userUUID, _ := uuid.NewUUID()
	newUser = models.Users{
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
		Guid:  userUUID,
	}

	db := conn.GetDB()
	dbtranx := db.Create(&newUser)
	if dbtranx.Error != nil {
		return nilUUID, dbtranx.Error
	}

	db.Save(&newUser)
	return userUUID, nil
}

func ListRestaurants(pagination *models.Pagination) (*[]models.Restaurant, error) {

	var listOfRestaurants []models.Restaurant
	db := conn.GetDB()

	offset := (pagination.Page - 1) * pagination.Per_Page
	queryBuilder := db.Limit(pagination.Per_Page).Offset(offset).Order(pagination.Sort)
	dbtranx := queryBuilder.Model(&models.Restaurant{}).Find(&listOfRestaurants)
	if dbtranx.Error != nil {
		return nil, dbtranx.Error
	}

	return &listOfRestaurants, nil
}

func GetMenuItems(restaurant *models.Menu) (*[]models.Menu_Item, error) {

	var menuItems []models.Menu_Item
	db := conn.GetDB()

	dbtranx := db.Find(&menuItems).Where("restaurant_id = ?", restaurant.Restaurant_ID).Order(restaurant.Sort)
	if dbtranx.Error != nil {
		return nil, dbtranx.Error
	}

	return &menuItems, nil
}
