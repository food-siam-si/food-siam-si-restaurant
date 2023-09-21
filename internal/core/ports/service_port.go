package ports

import "food-siam-si-restaurant/internal/core/domain"

type RestaurantService interface {
	VerifyRestaurantIdentity(id uint32, userId uint32) (bool, error)
	Create(restaurant domain.Restaurant) error
	FindById(id uint32) (domain.Restaurant, error)
	UpdateCurrent(userId uint32, restaurant domain.Restaurant) error
	GetCurrent(userId uint32) (domain.Restaurant, error)
	RandomRestaurant() (domain.Restaurant, error)
	FindAllType() ([]domain.RestaurantType, error)
}
