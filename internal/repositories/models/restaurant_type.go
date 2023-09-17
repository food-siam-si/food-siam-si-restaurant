package models

import "food-siam-si-restaurant/internal/core/domain"

type RestaurantType struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
}

func ParseRestaurantType(restaurantType *domain.RestaurantType) *RestaurantType {
	return &RestaurantType{
		Id:   restaurantType.Id,
		Name: restaurantType.Name,
	}
}

func (r *RestaurantType) ToDomain() *domain.RestaurantType {
	return &domain.RestaurantType{
		Id:   r.Id,
		Name: r.Name,
	}
}
