package main

import (
	"food-siam-si-restaurant/config"
	"food-siam-si-restaurant/infrastructure/database"
	"food-siam-si-restaurant/infrastructure/messagequeue"
	"food-siam-si-restaurant/internal/core/services"
	"food-siam-si-restaurant/internal/handlers"
	"food-siam-si-restaurant/internal/handlers/proto"
	"food-siam-si-restaurant/internal/repositories"
	"log"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	config.Load()

	db := database.NewGorm()
	restaurantRepo := repositories.NewRestaurantRepository(db)

	restaurantService := services.NewRestaurantService(restaurantRepo)

	customFunc := func(p any) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []recovery.Option{
		recovery.WithRecoveryHandler(customFunc),
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(
		recovery.UnaryServerInterceptor(opts...),
	),
		grpc.StreamInterceptor(
			recovery.StreamServerInterceptor(opts...),
		))
	restaurantTypeHdl := handlers.NewRestaurantTypeHandler(restaurantService)
	restaurantHdl := handlers.NewRestaurantHandler(restaurantService)
	proto.RegisterRestaurantTypeServiceServer(grpcServer, restaurantTypeHdl)
	proto.RegisterRestaurantServiceServer(grpcServer, restaurantHdl)

	lis, err := net.Listen("tcp", config.Get().App.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	kafkaReader := messagequeue.NewKafkaReader()
	defer kafkaReader.Close()

	kafkaHdl := handlers.NewKafkaHandler(kafkaReader, restaurantService)

	go kafkaHdl.Listen()
	grpcServer.Serve(lis)

}
