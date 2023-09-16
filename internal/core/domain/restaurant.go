package domain

type Restaurant struct {
	Id           uint
	UserId       string
	Name         string
	Description  string
	LocationLat  float64
	LocationLong float64
	PhoneNumber  string
	AveragePrice AveragePrice
	ImageUrl     string
	IsInService  bool
	Type         RestaurantType
}

type AveragePrice string

const (
	LowerThanHundread         AveragePrice = "LowerThanHundread"
	HundreadToTwoHundread     AveragePrice = "HundreadToTwoHundread"
	TwoHundreadToFiveHundread AveragePrice = "TwoHundreadToFiveHundread"
	MoreThanFiveHundread      AveragePrice = "MoreThanFiveHundread"
	MoreThanOneThousand       AveragePrice = "MoreThanOneThousand"
)

type RestaurantType struct {
	Id   uint
	Name string
}
