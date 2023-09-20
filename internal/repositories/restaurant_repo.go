package repositories

import (
	"food-siam-si-restaurant/internal/core/domain"
	"food-siam-si-restaurant/internal/core/ports"
	"food-siam-si-restaurant/internal/repositories/models"

	"gorm.io/gorm"
)

type restaurantRepository struct {
	db *gorm.DB
}

func NewRestaurantRepository(db *gorm.DB) ports.RestaurantRepository {
	return &restaurantRepository{db}
}

func (r *restaurantRepository) Create(payload *domain.Restaurant) error {
	restaurant := models.ParseRestaurant(payload)

	return r.db.Create(restaurant).Error
}

func (r *restaurantRepository) Update(id uint, payload *domain.Restaurant) error {
	return nil
}

func (r *restaurantRepository) FindById(id uint) (domain.Restaurant, error) {
	return domain.Restaurant{}, nil
}

func (r *restaurantRepository) FindAll() ([]domain.Restaurant, error) {
	return []domain.Restaurant{}, nil
}

func (r *restaurantRepository) FindAllType() ([]domain.RestaurantType, error) {
	return []domain.RestaurantType{}, nil
}

func (r *restaurantRepository) FindTypeById(id uint) (domain.RestaurantType, error) {
	return domain.RestaurantType{}, nil
}
