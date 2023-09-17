package domain

type Restaurant struct {
	Id           uint
	UserId       uint
	Name         string
	Description  string
	LocationLat  float64
	LocationLong float64
	PhoneNumber  string
	AveragePrice AveragePrice
	ImageUrl     string
	IsInService  bool
	Types        []RestaurantType
}

type AveragePrice string

const (
	LowerThanHundread         AveragePrice = "LowerThanHundread"
	HundreadToTwoHundread     AveragePrice = "HundreadToTwoHundread"
	TwoHundreadToFiveHundread AveragePrice = "TwoHundreadToFiveHundread"
	MoreThanFiveHundread      AveragePrice = "MoreThanFiveHundread"
	MoreThanOneThousand       AveragePrice = "MoreThanOneThousand"
)
