package services

import (
	"errors"
	"food-siam-si-restaurant/internal/core/domain"
	"food-siam-si-restaurant/internal/core/ports"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return false, status.Error(codes.NotFound, err.Error())
	}

	return res.UserId == userId, nil
}

func (svc *restaurantService) Create(restaurant domain.Restaurant) error {

	_, err := svc.repo.FindByUserId(restaurant.UserId)
	if err == nil {
		return status.Error(codes.AlreadyExists, "restaurant already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := svc.isValidType(restaurant.Types); err != nil {
		return status.Error(codes.InvalidArgument, "invalid restaurant type")
	}

	return svc.repo.Create(&restaurant)
}

func (svc *restaurantService) FindById(id uint32) (domain.Restaurant, error) {
	res, err := svc.repo.FindById(id)
	if err == gorm.ErrRecordNotFound {
		return res, status.Error(codes.NotFound, "restaurant not found")
	}
	return res, err
}

func (svc *restaurantService) GetCurrent(userId uint32) (domain.Restaurant, error) {
	res, err := svc.repo.FindByUserId(userId)
	if err == gorm.ErrRecordNotFound {
		return res, status.Error(codes.NotFound, "restaurant not found")
	}
	return res, err
}

func (svc *restaurantService) UpdateCurrent(userId uint32, restaurant domain.Restaurant) error {
	_, err := svc.repo.FindByUserId(userId)
	if err != nil {
		return status.Error(codes.NotFound, "restaurant not found")
	}

	if err := svc.isValidType(restaurant.Types); err != nil {
		return status.Error(codes.InvalidArgument, "invalid restaurant type")
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

func (svc *restaurantService) UpdateAverageScore(restaurantId uint32, score float32) error {

	if err := svc.repo.UpdateScoreById(restaurantId, score); err == gorm.ErrRecordNotFound {
		return status.Error(codes.NotFound, "restaurant not found")
	}

	return nil
}
