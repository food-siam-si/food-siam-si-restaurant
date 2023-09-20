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
	res, err := handler.svc.FindAllType()
	if err != nil {
		return nil, err
	}

	var restaurantTypes []*proto.RestaurantType
	for _, r := range res {
		restaurantTypes = append(restaurantTypes, &proto.RestaurantType{
			Id:   r.Id,
			Name: r.Name,
		})
	}

	return &proto.GetRestaurantTypeResponse{
		RestaurantTypes: restaurantTypes,
	}, nil
}
