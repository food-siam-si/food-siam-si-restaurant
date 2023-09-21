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

func (r *restaurantRepository) UpdateByUserId(userId uint32, payload *domain.Restaurant) error {
	domainRestaurant, err := r.FindByUserId(userId)
	if err != nil {
		return err
	}
	currentRestaurant := models.ParseRestaurant(&domainRestaurant)

	restaurantMap := models.ParseRestaurantToMap(payload)
	types := models.ParseRestaurant(payload).Types

	if err := r.db.Model(&currentRestaurant).Association("Types").Replace(types); err != nil {
		return err
	}

	return r.db.Model(&currentRestaurant).Updates(restaurantMap).Error
}

func (r *restaurantRepository) FindById(id uint32) (domain.Restaurant, error) {
	restaurant := models.Restaurant{}

	err := r.db.Where("id = ?", id).Preload("Types").First(&restaurant).Error

	return *restaurant.ToDomain(), err
}

func (r *restaurantRepository) FindAll() ([]domain.Restaurant, error) {
	restaurants := []models.Restaurant{}

	err := r.db.Preload("Types").Find(&restaurants).Error

	if err != nil {
		return []domain.Restaurant{}, err
	}

	result := make([]domain.Restaurant, len(restaurants))

	for i, restaurant := range restaurants {
		result[i] = *restaurant.ToDomain()
	}

	return result, nil
}

func (r *restaurantRepository) FindAllType() ([]domain.RestaurantType, error) {
	restaurantTypes := []models.RestaurantType{}

	err := r.db.Find(&restaurantTypes).Error

	if err != nil {
		return []domain.RestaurantType{}, err
	}

	result := make([]domain.RestaurantType, len(restaurantTypes))

	for i, restaurantType := range restaurantTypes {
		result[i] = *restaurantType.ToDomain()
	}

	return result, nil
}

func (r *restaurantRepository) FindTypeById(id uint32) (domain.RestaurantType, error) {
	restaurantType := models.RestaurantType{}

	err := r.db.Where("id = ?", id).First(&restaurantType).Error

	return *restaurantType.ToDomain(), err
}

func (r *restaurantRepository) FindByUserId(userId uint32) (domain.Restaurant, error) {
	restaurant := models.Restaurant{}

	err := r.db.Where("user_id = ?", userId).Preload("Types").First(&restaurant).Error

	return *restaurant.ToDomain(), err
}
