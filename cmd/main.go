package main

import (
	"food-siam-si-restaurant/config"
	"food-siam-si-restaurant/infrastructure/database"
)

func main() {
	config.Load()

	database.NewGorm()
}
