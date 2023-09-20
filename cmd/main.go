package main

import (
	"food-siam-si-restaurant/config"
	"food-siam-si-restaurant/internal/driver/handlers"
	"food-siam-si-restaurant/internal/driver/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config.Load()

	// db := database.NewGorm()
	// restaurantRepo := repositories.NewRestaurantRepository(db)

	grpcServer := grpc.NewServer()
	restaurantTypeHdl := handlers.NewRestaurantTypeHandler()
	proto.RegisterRestaurantTypeServiceServer(grpcServer, restaurantTypeHdl)

	lis, err := net.Listen("tcp", config.Get().App.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)

}
