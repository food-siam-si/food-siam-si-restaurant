//	type RestaurantService interface {
//		VerifyRestaurantIdentity(restaurantId string) error
//		Create(restaurant domain.Restaurant) error
//		FindById(id string) (domain.Restaurant, error)
//		Update(id string, restaurant domain.Restaurant) error
//		RandomRestaurant() (domain.Restaurant, error)
//		FindAllType() ([]domain.RestaurantType, error)
//	}
package services

import (
	"food-siam-si-restaurant/internal/core/domain"
	"food-siam-si-restaurant/internal/core/ports"
)

type restaurantService struct {
	repo ports.RestaurantRepository
}

func NewRestaurantService(repo ports.RestaurantRepository) ports.RestaurantService {
	return &restaurantService{repo}
}

func (svc *restaurantService) VerifyRestaurantIdentity(restaurantId string) error {
	return nil
}

func (svc *restaurantService) Create(restaurant domain.Restaurant) error {
	return nil
}

func (svc *restaurantService) FindById(id string) (domain.Restaurant, error) {
	return domain.Restaurant{}, nil
}

func (svc *restaurantService) Update(id string, restaurant domain.Restaurant) error {
	return nil
}

func (svc *restaurantService) RandomRestaurant() (domain.Restaurant, error) {
	return domain.Restaurant{}, nil
}

func (svc *restaurantService) FindAllType() ([]domain.RestaurantType, error) {
	return []domain.RestaurantType{}, nil
}
