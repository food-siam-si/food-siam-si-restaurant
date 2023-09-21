package services

import (
	"errors"
	"food-siam-si-restaurant/internal/core/domain"
	"food-siam-si-restaurant/internal/core/ports"

	"gorm.io/gorm"
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

	_, err := svc.repo.FindByUserId(restaurant.UserId)
	if err == nil {
		return errors.New("restaurant already exist")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := svc.isValidType(restaurant.Types); err != nil {
		return err
	}

	return svc.repo.Create(&restaurant)
}

func (svc *restaurantService) FindById(id uint32) (domain.Restaurant, error) {
	return svc.repo.FindById(id)
}

func (svc *restaurantService) GetCurrent(userId uint32) (domain.Restaurant, error) {
	return svc.repo.FindByUserId(userId)
}

func (svc *restaurantService) UpdateCurrent(userId uint32, restaurant domain.Restaurant) error {
	_, err := svc.repo.FindByUserId(userId)
	if err != nil {
		return err
	}

	if err := svc.isValidType(restaurant.Types); err != nil {
		return err
	}

	return svc.repo.UpdateByUserId(userId, &restaurant)
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
