package ports

import "food-siam-si-restaurant/internal/core/domain"

type RestaurantService interface {
	VerifyRestaurantIdentity(id uint) error
	Create(restaurant domain.Restaurant) error
	FindById(id uint) (domain.Restaurant, error)
	Update(id uint, restaurant domain.Restaurant) error
	RandomRestaurant() (domain.Restaurant, error)
	FindAllType() ([]domain.RestaurantType, error)
}
