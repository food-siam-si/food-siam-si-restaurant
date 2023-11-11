package ports

import "food-siam-si-restaurant/internal/core/domain"

type RestaurantRepository interface {
	Create(restaurant *domain.Restaurant) error
	UpdateByUserId(id uint32, restaurant *domain.Restaurant) error
	FindById(id uint32) (domain.Restaurant, error)
	FindAll(restaurantTypeIds []uint32, currentLat float32, currentLong float32, maxDistanceKm uint32) ([]domain.Restaurant, error)
	FindAllType() ([]domain.RestaurantType, error)
	FindTypeById(id uint32) (domain.RestaurantType, error)
	FindByUserId(userId uint32) (domain.Restaurant, error)
	UpdateScoreById(id uint32, score float32) error
}
