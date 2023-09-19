package models

import "food-siam-si-restaurant/internal/core/domain"

type Restaurant struct {
	Id           uint   `gorm:"primaryKey;autoIncrement"`
	UserId       uint   `gorm:"not null"`
	Name         string `gorm:"not null"`
	Description  string
	LocationLat  float64             `gorm:"not null"`
	LocationLong float64             `gorm:"not null"`
	PhoneNumber  string              `gorm:"not null"`
	AveragePrice domain.AveragePrice `sql:"type:ENUM('LowerThanHundread','HundreadToTwoHundread','TwoHundreadToFiveHundread','MoreThanFiveHundread','MoreThanOneThousand')"`
	ImageUrl     string              `gorm:"not null"`
	IsInService  bool                `gorm:"not null;default:true"`
	Types        []RestaurantType    `gorm:"many2many:restaurant_restaurant_type"`
}

func ParseRestaurant(restaurant *domain.Restaurant) *Restaurant {
	var restaurantTypes []RestaurantType
	for _, restaurantType := range restaurant.Types {
		restaurantTypes = append(restaurantTypes, *ParseRestaurantType(&restaurantType))
	}

	return &Restaurant{
		UserId:       restaurant.UserId,
		Name:         restaurant.Name,
		Description:  restaurant.Description,
		LocationLat:  restaurant.LocationLat,
		LocationLong: restaurant.LocationLong,
		PhoneNumber:  restaurant.PhoneNumber,
		AveragePrice: restaurant.AveragePrice,
		ImageUrl:     restaurant.ImageUrl,
		IsInService:  restaurant.IsInService,
		Types:        restaurantTypes,
	}
}

func (r *Restaurant) ToDomain() *domain.Restaurant {
	var restaurantTypes []domain.RestaurantType
	for _, restaurantType := range r.Types {
		restaurantTypes = append(restaurantTypes, *restaurantType.ToDomain())
	}

	return &domain.Restaurant{
		Id:           r.Id,
		UserId:       r.UserId,
		Name:         r.Name,
		Description:  r.Description,
		LocationLat:  r.LocationLat,
		LocationLong: r.LocationLong,
		PhoneNumber:  r.PhoneNumber,
		AveragePrice: r.AveragePrice,
		ImageUrl:     r.ImageUrl,
		IsInService:  r.IsInService,
		Types:        restaurantTypes,
	}
}
