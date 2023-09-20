package services

import (
	"errors"
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
	if err := svc.isValidType(restaurant.Types); err != nil {
		return err
	}

	return svc.repo.Create(&restaurant)
}

func (svc *restaurantService) FindById(id uint32) (domain.Restaurant, error) {
	return svc.repo.FindById(id)
}

func (svc *restaurantService) Update(id uint32, restaurant domain.Restaurant) error {
	current, err := svc.repo.FindById(id)
	if err != nil {
		return err
	}
	if current.UserId != restaurant.UserId {
		return errors.New("forbidden")
	}

	if err := svc.isValidType(restaurant.Types); err != nil {
		return err
	}

	return svc.repo.Update(id, &restaurant)
}

func (svc *restaurantService) RandomRestaurant() (domain.Restaurant, error) {
	return domain.Restaurant{}, nil
}

func (svc *restaurantService) FindAllType() ([]domain.RestaurantType, error) {
	return svc.repo.FindAllType()
}

func (svc *restaurantService) isValidType(restaurantTypes []domain.RestaurantType) error {
	for _, restaurantType := range restaurantTypes {
		if _, err := svc.repo.FindTypeById(restaurantType.Id); err != nil {
			return err
		}
	}

	return nil
}
