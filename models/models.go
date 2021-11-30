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

func (t *Restaurant) TableName() string {
	return "restaurants"
}

type Menu struct {
	Restaurant_ID int    `json:"restaurantid"`
	Sort          string `json:"sort"`
}

type Menu_Category int

const (
	Starter Menu_Category = iota
	Side_Dish
	Main_Course
	Beverage
	Dessert
)

type Menu_Item struct {
	gorm.Model
	Restaurant_Id Restaurant
	Name          string
	Description   string
	Menu_Category Menu_Category
	Price         float32
}

func (t *Menu_Item) TableName() string {
	return "Menu_Items"
}
