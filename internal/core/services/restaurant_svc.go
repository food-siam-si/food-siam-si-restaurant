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

func (svc *restaurantService) VerifyRestaurantIdentity(id uint32, userId uint32) (bool, error) {
	res, err := svc.repo.FindById(id)
	if err != nil {
		return false, err
	}

	return res.UserId == userId, nil
}

func (svc *restaurantService) Create(restaurant domain.Restaurant) error {
	// check if type is valid or not
	for _, restarestaurantType := range restaurant.Types {
		if _, err := svc.repo.FindTypeById(restarestaurantType.Id); err != nil {
			return err
		}
	}

	return svc.repo.Create(&restaurant)
}

func (svc *restaurantService) FindById(id uint32) (domain.Restaurant, error) {
	return svc.repo.FindById(id)
}

func (svc *restaurantService) Update(id uint32, restaurant domain.Restaurant) error {
	return nil
}

func (svc *restaurantService) RandomRestaurant() (domain.Restaurant, error) {
	return domain.Restaurant{}, nil
}

func (svc *restaurantService) FindAllType() ([]domain.RestaurantType, error) {
	return svc.repo.FindAllType()
}
