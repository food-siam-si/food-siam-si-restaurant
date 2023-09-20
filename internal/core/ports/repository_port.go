package ports

import "food-siam-si-restaurant/internal/core/domain"

type RestaurantRepository interface {
	Create(restaurant *domain.Restaurant) error
	Update(id uint, restaurant *domain.Restaurant) error
	FindById(id uint) (domain.Restaurant, error)
	FindAll() ([]domain.Restaurant, error)
	FindAllType() ([]domain.RestaurantType, error)
	FindTypeById(id uint) (domain.RestaurantType, error)
}
