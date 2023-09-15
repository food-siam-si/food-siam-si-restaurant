package database

import (
	"fmt"
	"food-siam-si-restaurant/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.GetDb().Host, config.GetDb().User, config.GetDb().Password, config.GetDb().Name, config.GetDb().Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
