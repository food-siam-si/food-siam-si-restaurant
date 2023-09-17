package main

import (
	"food-siam-si-restaurant/config"
	"food-siam-si-restaurant/infrastructure/database"
	"food-siam-si-restaurant/internal/core/domain"
	"food-siam-si-restaurant/internal/repositories"
)

func main() {
	config.Load()

	db := database.NewGorm()
	restaurantRepo := repositories.NewRestaurantRepository(db)

	restaurantRepo.Create(&domain.Restaurant{
		Name:         "Siam Si Restaurant",
		UserId:       1,
		Description:  "Siam Si Restaurant",
		LocationLat:  1.5,
		LocationLong: 2.5,
		PhoneNumber:  "08123456789",
		AveragePrice: domain.AveragePrice("LowerThanHundread"),
		ImageUrl:     "www.google.com",
		IsInService:  true,
		Types: []domain.RestaurantType{
			{
				Id: 2,
			},
		},
	})

}
