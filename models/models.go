package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name  string
	Email string
	Phone string
	Guid  uuid.UUID
}

func (t *Users) TableName() string {
	return "users"
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

type Restaurant struct {
	gorm.Model
	Name        string
	Description string
	Star_rating float32
	Address     string
}

// type Menu struct {
// 	gorm.Model
// 	Restaurant_id Restaurant
// 	Name          string
// 	Description   string
// 	Menu_category int
// 	Price         float32
// }
