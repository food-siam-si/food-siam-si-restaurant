package handlers

import (
	"context"
	"food-siam-si-restaurant/internal/core/domain"
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
	res, err := handler.svc.VerifyRestaurantIdentity(req.Id, req.User.Id)
	if err != nil {
		return wrapperspb.Bool(false), err
	}

	return wrapperspb.Bool(res), nil
}

func (handler RestaurantHandler) Create(ctx context.Context, req *proto.CreateRestaurantRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (handler RestaurantHandler) FindById(ctx context.Context, req *wrapperspb.UInt32Value) (*proto.Restaurant, error) {
	res, err := handler.svc.FindById(req.Value)
	if err != nil {
		return nil, err
	}

	var restaurantType []*proto.RestaurantType

	for _, v := range res.Types {
		restaurantType = append(restaurantType, &proto.RestaurantType{
			Id:   v.Id,
			Name: v.Name,
		})
	}

	return &proto.Restaurant{
		Id:             res.Id,
		Name:           res.Name,
		Description:    &res.Description,
		PhoneNumber:    res.PhoneNumber,
		UserId:         res.UserId,
		LocationLat:    res.LocationLat,
		LocationLong:   res.LocationLong,
		ImageUrl:       res.ImageUrl,
		IsInService:    res.IsInService,
		RestaurantType: restaurantType,
		AveragePrice:   parseAveragePrice(&res.AveragePrice),
	}, nil
}

func (handler RestaurantHandler) Update(ctx context.Context, req *proto.UpdateRestaurantRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (handler RestaurantHandler) Random(ctx context.Context, req *proto.RandomRestaurantRequest) (*proto.Restaurant, error) {
	return &proto.Restaurant{}, nil
}

func parseAveragePrice(averagePrice *domain.AveragePrice) proto.AveragePrice {
	switch *averagePrice {
	case domain.LowerThanHundred:
		return proto.AveragePrice_LowerThanHundreds
	case domain.HundredToTwoHundred:
		return proto.AveragePrice_HundredToTwoHundred
	case domain.TwoHundredToFiveHundred:
		return proto.AveragePrice_TwoHundredToFiveHundred
	case domain.MoreThanFiveHundred:
		return proto.AveragePrice_MoreThanFiveHundred
	case domain.MoreThanOneThousand:
		return proto.AveragePrice_MoreThanOneThousand
	default:
		return proto.AveragePrice_LowerThanHundreds
	}
}
