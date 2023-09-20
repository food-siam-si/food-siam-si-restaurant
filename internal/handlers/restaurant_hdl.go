package handlers

import (
	"context"
	"food-siam-si-restaurant/internal/core/ports"
	"food-siam-si-restaurant/internal/handlers/proto"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type RestaurantHandler struct {
	svc ports.RestaurantService
	proto.RestaurantServiceServer
}

func NewRestaurantHandler(svc ports.RestaurantService) proto.RestaurantServiceServer {
	return RestaurantHandler{
		svc: svc,
	}
}

func (handler RestaurantHandler) VerifyIdentity(ctx context.Context, req *proto.VerifyRestaurantIdentityRequest) (*wrapperspb.BoolValue, error) {
	return wrapperspb.Bool(true), nil
}

func (handler RestaurantHandler) Create(ctx context.Context, req *proto.CreateRestaurantRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (handler RestaurantHandler) FindById(ctx context.Context, req *wrapperspb.UInt32Value) (*proto.Restaurant, error) {
	return &proto.Restaurant{}, nil
}

func (handler RestaurantHandler) Update(ctx context.Context, req *proto.UpdateRestaurantRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (handler RestaurantHandler) Random(ctx context.Context, req *proto.RandomRestaurantRequest) (*proto.Restaurant, error) {
	return &proto.Restaurant{}, nil
}
