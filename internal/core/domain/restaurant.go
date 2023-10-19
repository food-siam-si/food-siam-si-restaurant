package domain

type Restaurant struct {
	Id           uint32
	UserId       uint32
	Name         string
	Description  string
	LocationLat  float32
	LocationLong float32
	PhoneNumber  string
	AveragePrice AveragePrice
	ImageUrl     string
	IsInService  bool
	Types        []RestaurantType
	AverageScore float32
}

type AveragePrice string

const (
	LowerThanHundred        AveragePrice = "LowerThanHundred"
	HundredToTwoHundred     AveragePrice = "HundredToTwoHundred"
	TwoHundredToFiveHundred AveragePrice = "TwoHundredToFiveHundred"
	MoreThanFiveHundred     AveragePrice = "MoreThanFiveHundred"
	MoreThanOneThousand     AveragePrice = "MoreThanOneThousand"
)
