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
	Per_Page int    `json:"per_page"`
	Page     int    `json:"page"`
	Sort     string `json:"sort"`
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

type Category int

const (
	Starter Category = iota
	Side_Dish
	Main_Course
	Beverage
	Dessert
)

type Menu_Item struct {
	gorm.Model
	RestaurantID  int
	Restaurant    Restaurant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name          string
	Description   string
	Menu_Category Category
	Price         float32
}

func (t *Menu_Item) TableName() string {
	return "menu_items"
}
