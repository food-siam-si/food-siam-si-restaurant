package handlers

import (
	"context"
	"food-siam-si-restaurant/internal/core/ports"
	"food-siam-si-restaurant/internal/handlers/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type RestaurantTypeHandler struct {
	svc ports.RestaurantService
	proto.RestaurantTypeServiceServer
}

func NewRestaurantTypeHandler(svc ports.RestaurantService) proto.RestaurantTypeServiceServer {
	return RestaurantTypeHandler{
		svc: svc,
	}
}

func (handler RestaurantTypeHandler) GetAll(context.Context, *emptypb.Empty) (*proto.GetRestaurantTypeResponse, error) {
	return &proto.GetRestaurantTypeResponse{
		RestaurantTypes: []*proto.RestaurantType{
			{
				Id:   1,
				Name: "Thai",
			},
			{
				Id:   2,
				Name: "Chinese",
			},
		},
	}, nil
}
