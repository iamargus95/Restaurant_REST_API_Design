package services

import (
	"github.com/google/uuid"
	"github.com/iamargus95/restaurant_rest_api_design/models"
	r "github.com/iamargus95/restaurant_rest_api_design/resources"
	"gorm.io/gorm"
)

func Signup(db *gorm.DB, body r.SignupPayload) (uuid.UUID, error) {

	var nilUUID uuid.UUID
	var newUser models.Users

	userUUID, _ := uuid.NewUUID()
	newUser = models.Users{
		Name:  body.Name,
		Email: body.Email,
		Phone: body.Phone,
		Guid:  userUUID,
	}

	dbtranx := db.Create(&newUser)
	if dbtranx.Error != nil {
		return nilUUID, dbtranx.Error
	}

	db.Save(&newUser)
	return userUUID, nil
}
