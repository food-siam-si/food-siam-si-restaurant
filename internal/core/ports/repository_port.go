package ports

import "food-siam-si-restaurant/internal/core/domain"

type RestaurantRepository interface {
	Create(restaurant *domain.Restaurant) error
	Update(id uint32, restaurant *domain.Restaurant) error
	FindById(id uint32) (domain.Restaurant, error)
	FindAll() ([]domain.Restaurant, error)
	FindAllType() ([]domain.RestaurantType, error)
	FindTypeById(id uint32) (domain.RestaurantType, error)
	FindByUserId(userId uint32) (domain.Restaurant, error)
}
