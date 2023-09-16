package ports

import "food-siam-si-restaurant/internal/core/domain"

type RestaurantService interface {
	VerifyRestaurantIdentity(restaurantId string) error
	Create(restaurant domain.Restaurant) error
	FindById(id string) (domain.Restaurant, error)
	Update(id string, restaurant domain.Restaurant) error
	RandomRestaurant() (domain.Restaurant, error)
	FindAllType() ([]domain.RestaurantType, error)
}
