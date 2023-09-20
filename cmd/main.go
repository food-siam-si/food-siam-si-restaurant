package main

import (
	"food-siam-si-restaurant/config"
	"food-siam-si-restaurant/infrastructure/database"
	"food-siam-si-restaurant/internal/core/services"
	"food-siam-si-restaurant/internal/handlers"
	"food-siam-si-restaurant/internal/handlers/proto"
	"food-siam-si-restaurant/internal/repositories"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config.Load()

	db := database.NewGorm()
	restaurantRepo := repositories.NewRestaurantRepository(db)

	restaurantService := services.NewRestaurantService(restaurantRepo)

	grpcServer := grpc.NewServer()
	restaurantTypeHdl := handlers.NewRestaurantTypeHandler(restaurantService)
	restaurantHdl := handlers.NewRestaurantHandler(restaurantService)
	proto.RegisterRestaurantTypeServiceServer(grpcServer, restaurantTypeHdl)
	proto.RegisterRestaurantServiceServer(grpcServer, restaurantHdl)

	lis, err := net.Listen("tcp", config.Get().App.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)

}
