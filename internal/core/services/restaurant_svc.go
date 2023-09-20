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

func (svc *restaurantService) VerifyRestaurantIdentity(id uint) error {
	return nil
}

func (svc *restaurantService) Create(restaurant domain.Restaurant) error {
	return nil
}

func (svc *restaurantService) FindById(id uint) (domain.Restaurant, error) {
	return domain.Restaurant{}, nil
}

func (svc *restaurantService) Update(id uint, restaurant domain.Restaurant) error {
	return nil
}

func (svc *restaurantService) RandomRestaurant() (domain.Restaurant, error) {
	return domain.Restaurant{}, nil
}

func (svc *restaurantService) FindAllType() ([]domain.RestaurantType, error) {
	return []domain.RestaurantType{}, nil
}