package handlers

import (
	"context"
	"food-siam-si-restaurant/internal/driver/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type RestaurantTypeHandler struct {
	proto.RestaurantTypeServiceServer
}

func NewRestaurantTypeHandler() proto.RestaurantTypeServiceServer {
	return RestaurantTypeHandler{}
}

func (handler RestaurantTypeHandler) GetRestaurantTypes(context.Context, *emptypb.Empty) (*proto.GetRestaurantTypeResponse, error) {
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
