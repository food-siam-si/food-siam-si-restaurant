package ports

import "food-siam-si-restaurant/internal/core/domain"

type RestaurantRepository interface {
	Create(restaurant *domain.Restaurant) error
	Update(id string, restaurant *domain.Restaurant) error
	FindById(id string) (domain.Restaurant, error)
	FindAll() ([]domain.Restaurant, error)
	FindAllType() ([]domain.RestaurantType, error)
}
